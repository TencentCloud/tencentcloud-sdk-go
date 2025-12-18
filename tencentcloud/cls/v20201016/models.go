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

package v20201016

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddMachineGroupInfoRequestParams struct {
	// 机器组Id
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/product/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 机器组类型
	// 目前type支持 ip 和 label
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`
}

type AddMachineGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 机器组Id
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/product/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 机器组类型
	// 目前type支持 ip 和 label
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`
}

func (r *AddMachineGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMachineGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "MachineGroupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddMachineGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddMachineGroupInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddMachineGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *AddMachineGroupInfoResponseParams `json:"Response"`
}

func (r *AddMachineGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMachineGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvanceFilterRuleInfo struct {
	// 过滤字段
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 过滤规则，0:等于，1:字段存在，2:字段不存在
	Rule *uint64 `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 过滤值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AdvancedConsumerConfiguration struct {
	// Ckafka分区hash状态。 默认 false
	// 
	// - true：开启根据字段 Hash 值结果相等的信息投递到统一 ckafka 分区
	// - false：关闭根据字段 Hash 值结果相等的信息投递到统一 ckafka 分区
	PartitionHashStatus *bool `json:"PartitionHashStatus,omitnil,omitempty" name:"PartitionHashStatus"`

	// 需要计算 hash 的字段列表。最大支持5个字段。
	PartitionFields []*string `json:"PartitionFields,omitnil,omitempty" name:"PartitionFields"`
}

type AlarmAnalysisConfig struct {
	// 键。支持以下key：
	// SyntaxRule：语法规则，value支持 0：Lucene语法；1： CQL语法。
	// QueryIndex：执行语句序号。value支持  -1：自定义； 1：执行语句1； 2：执行语句2。
	// CustomQuery：检索语句。 QueryIndex为-1时有效且必填，value示例： "* | select count(*) as count"。
	// Fields：字段。value支持 __SOURCE__；__FILENAME__；__HOSTNAME__；__TIMESTAMP__；__INDEX_STATUS__；__PKG_LOGID__；__TOPIC__。
	// Format：显示形式。value支持 1：每条日志一行；2：每条日志每个字段一行。
	// Limit：最大日志条数。 value示例： 5。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值。
	// 键对应值如下：
	// SyntaxRule：语法规则，value支持 0：Lucene语法；1： CQL语法。
	// QueryIndex：执行语句序号。value支持  -1：自定义； 1：执行语句1； 2：执行语句2。
	// CustomQuery：检索语句。 QueryIndex为-1时有效且必填，value示例： "* | select count(*) as count"。
	// Fields：字段。value支持 __SOURCE__；__FILENAME__；__HOSTNAME__；__TIMESTAMP__；__INDEX_STATUS__；__PKG_LOGID__；__TOPIC__。
	// Format：显示形式。value支持 1：每条日志一行；2：每条日志每个字段一行。
	// Limit：最大日志条数。 value示例： 5。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AlarmClassification struct {
	// 分类键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 分类值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AlarmInfo struct {
	// 告警策略名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 监控对象列表。
	AlarmTargets []*AlarmTargetInfo `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// 是否触发告警的单触发条件。与MultiConditions参数互斥。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// 告警重复的周期。单位是min。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知渠道组列表。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/product/614/56462)获取关联的告警通知渠道组列表，和MonitorNotice互斥
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// 开启状态。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 创建时间。格式： YYYY-MM-DD HH:MM:SS
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近更新时间。格式： YYYY-MM-DD HH:MM:SS
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 自定义通知模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// 自定义回调模板
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// 多维分析设置
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// 分组触发状态。true：开启，false：关闭（默认）
	GroupTriggerStatus *bool `json:"GroupTriggerStatus,omitnil,omitempty" name:"GroupTriggerStatus"`

	// 分组触发条件。
	GroupTriggerCondition []*string `json:"GroupTriggerCondition,omitnil,omitempty" name:"GroupTriggerCondition"`

	// 告警策略绑定的标签信息。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 监控对象类型。0:执行语句共用监控对象;1:每个执行语句单独选择监控对象。 
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 告警级别。0:警告(Warn);1:提醒(Info);2:紧急 (Critical)。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警附加分类字段。
	Classifications []*AlarmClassification `json:"Classifications,omitnil,omitempty" name:"Classifications"`

	// 多触发条件。与
	// Condition互斥。
	MultiConditions []*MultiCondition `json:"MultiConditions,omitnil,omitempty" name:"MultiConditions"`

	// 腾讯云可观测平台通知渠道相关信息，和AlarmNoticeIds互斥
	MonitorNotice *MonitorNotice `json:"MonitorNotice,omitnil,omitempty" name:"MonitorNotice"`
}

type AlarmNotice struct {
	// 告警通知渠道组名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 告警通知渠道组绑定的标签信息。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 告警模板的类型。可选值：
	// <br><li> Trigger - 告警触发</li>
	// <br><li> Recovery - 告警恢复</li>
	// <br><li> All - 告警触发和告警恢复</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警通知模板接收者信息。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// 告警通知模板回调信息。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`

	// 告警通知模板ID。
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// 通知规则。
	NoticeRules []*NoticeRule `json:"NoticeRules,omitnil,omitempty" name:"NoticeRules"`

	// 免登录操作告警开关。
	// 参数值： 1：关闭 2：开启（默认开启）
	AlarmShieldStatus *uint64 `json:"AlarmShieldStatus,omitnil,omitempty" name:"AlarmShieldStatus"`

	// 调用链接域名。http:// 或者 https:// 开头，不能/结尾
	JumpDomain *string `json:"JumpDomain,omitnil,omitempty" name:"JumpDomain"`

	// 投递相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmNoticeDeliverConfig *AlarmNoticeDeliverConfig `json:"AlarmNoticeDeliverConfig,omitnil,omitempty" name:"AlarmNoticeDeliverConfig"`

	// 创建时间。格式： YYYY-MM-DD HH:MM:SS
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近更新时间。格式： YYYY-MM-DD HH:MM:SS
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 投递日志开关。
	// 
	// 参数值：
	// 
	// 1：关闭
	// 
	// 2：开启 
	DeliverStatus *uint64 `json:"DeliverStatus,omitnil,omitempty" name:"DeliverStatus"`

	// 投递日志标识。
	// 
	// 参数值：
	// 
	// 1：未启用
	// 
	// 2：已启用
	// 
	// 3：投递异常
	DeliverFlag *uint64 `json:"DeliverFlag,omitnil,omitempty" name:"DeliverFlag"`

	// 通知渠道组配置的告警屏蔽统计状态数量信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmShieldCount *AlarmShieldCount `json:"AlarmShieldCount,omitnil,omitempty" name:"AlarmShieldCount"`

	// 统一设定自定义回调参数。
	// -  true: 使用通知内容模板中的自定义回调参数覆盖告警策略中单独配置的请求头及请求内容。
	// -  false:优先使用告警策略中单独配置的请求头及请求内容。
	CallbackPrioritize *bool `json:"CallbackPrioritize,omitnil,omitempty" name:"CallbackPrioritize"`
}

type AlarmNoticeDeliverConfig struct {
	// 通知渠道投递日志配置信息。
	DeliverConfig *DeliverConfig `json:"DeliverConfig,omitnil,omitempty" name:"DeliverConfig"`

	// 投递失败原因。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

type AlarmShieldCount struct {
	// 符合检索条件的告警屏蔽总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 告警屏蔽未生效数量
	InvalidCount *uint64 `json:"InvalidCount,omitnil,omitempty" name:"InvalidCount"`

	// 告警屏蔽生效中数量
	ValidCount *uint64 `json:"ValidCount,omitnil,omitempty" name:"ValidCount"`

	// 告警屏蔽已过期数量
	ExpireCount *uint64 `json:"ExpireCount,omitnil,omitempty" name:"ExpireCount"`
}

type AlarmShieldInfo struct {
	// 通知渠道组Id
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// 屏蔽规则id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 屏蔽开始时间（秒级时间戳）。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 屏蔽结束时间（秒级时间戳）。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 屏蔽类型。1：屏蔽所有通知，2：按照Rule参数屏蔽匹配规则的通知。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 屏蔽规则，当Type为2时必填。规则填写方式详见[产品文档](https://cloud.tencent.com/document/product/614/103178#rule)。
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 屏蔽原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 规则创建来源。
	// 1. 控制台，2.api，3.告警通知
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 操作者。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 规则状态。
	// 0：暂未生效，1：生效中，2：已失效
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则创建时间。秒级时间戳(s)
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则更新时间。秒级时间戳(s)
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AlarmTarget struct {
	// 日志主题ID。-通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 查询语句。
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 告警对象序号；从1开始递增。
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`

	// 查询范围起始时间相对于告警执行时间的偏移，单位为分钟，取值为非正，最大值为0，最小值为-1440。
	StartTimeOffset *int64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 查询范围终止时间相对于告警执行时间的偏移，单位为分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。
	EndTimeOffset *int64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 日志集ID。通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 检索语法规则，默认值为0。
	// 0：Lucene语法，1：CQL语法。
	// 详细说明参见<a href="https://cloud.tencent.com/document/product/614/47044#RetrievesConditionalRules" target="_blank">检索条件语法规则</a>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

type AlarmTargetInfo struct {
	// 日志集ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名称。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 查询语句。
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 告警对象序号。
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`

	// 查询范围起始时间相对于告警执行时间的偏移，单位为分钟，取值为非正，最大值为0，最小值为-1440。
	StartTimeOffset *int64 `json:"StartTimeOffset,omitnil,omitempty" name:"StartTimeOffset"`

	// 查询范围终止时间相对于告警执行时间的偏移，单位为分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。
	EndTimeOffset *int64 `json:"EndTimeOffset,omitnil,omitempty" name:"EndTimeOffset"`

	// 检索语法规则，默认值为0。
	// 0：Lucene语法，1：CQL语法。
	// 详细说明参见<a href="https://cloud.tencent.com/document/product/614/47044#RetrievesConditionalRules" target="_blank">检索条件语法规则</a>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// 主题类型。
	// 0: 日志主题，1: 指标主题
	BizType *uint64 `json:"BizType,omitnil,omitempty" name:"BizType"`
}

type AlertHistoryNotice struct {
	// 通知渠道组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 通知渠道组ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`
}

type AlertHistoryRecord struct {
	// 告警历史ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 告警策略ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 告警策略名称
	AlarmName *string `json:"AlarmName,omitnil,omitempty" name:"AlarmName"`

	// 监控对象ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 监控对象名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 监控对象所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 触发条件
	Trigger *string `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// 持续周期，持续满足触发条件TriggerCount个周期后，再进行告警
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// 告警通知发送频率，单位为分钟(min)
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// 通知渠道组
	Notices []*AlertHistoryNotice `json:"Notices,omitnil,omitempty" name:"Notices"`

	// 告警持续时间，单位为分钟(min)
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 告警状态，0代表未恢复，1代表已恢复，2代表已失效
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 告警发生时间，毫秒级Unix时间戳(ms)
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 告警分组触发时对应的分组信息
	GroupTriggerCondition []*GroupTriggerConditionInfo `json:"GroupTriggerCondition,omitnil,omitempty" name:"GroupTriggerCondition"`

	// 告警级别，0代表警告(Warn)，1代表提醒(Info)，2代表紧急 (Critical)
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 监控对象类型。
	// 0:执行语句共用监控对象; 1:每个执行语句单独选择监控对象。 
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 通知渠道类型，0默认代表cls内部通知渠道，1代表腾讯云可观测平台通知渠道
	SendType *uint64 `json:"SendType,omitnil,omitempty" name:"SendType"`
}

type AnalysisDimensional struct {
	// 分析名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分析类型：query(自定义检索分析)，field(字段TOP5及占比统计) ，original(相关原始日志)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分析内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 多维分析配置。
	// 
	// 当Analysis的Type字段为query（自定义）时，支持
	// {
	// "Key": "SyntaxRule",  // 语法规则
	// "Value": "1"  //0：Lucene语法 ，1： CQL语法
	// }
	// 
	// 当Analysis的Type字段为field（top5）时,  支持
	//  {
	//     "Key": "QueryIndex",
	//     "Value": "-1" //  -1：自定义， 1：执行语句1， 2：执行语句2
	// },{
	//     "Key": "CustomQuery", //检索语句。 QueryIndex为-1时有效且必填
	//     "Value": "* | select count(*) as count"
	// },{
	//     "Key": "SyntaxRule", // 查不到这个字段也是老语法（Lucene）
	//     "Value": "0"//0:Lucene, 1:CQL
	// }       
	// 
	// 当Analysis的Type字段为original（原始日志）时,  支持
	// {
	//     "Key": "Fields",
	//     "Value": "__SOURCE__,__HOSTNAME__,__TIMESTAMP__,__PKG_LOGID__,__TAG__.pod_ip"
	// }, {
	//     "Key": "QueryIndex",
	//     "Value": "-1" //  -1：自定义， 1：执行语句1， 2：执行语句2
	// },{
	//     "Key": "CustomQuery", //  //检索语句。 QueryIndex为-1时有效且必填
	//     "Value": "* | select count(*) as count"
	// },{
	//     "Key": "Format", //显示形式。1：每条日志一行，2：每条日志每个字段一行
	//     "Value": "2"
	// },
	// {
	//     "Key": "Limit", //最大日志条数
	//     "Value": "5"
	// },{
	//     "Key": "SyntaxRule", // 查不到这个字段也是老语法
	//     "Value": "0"//0:Lucene, 1:CQL
	// }
	ConfigInfo []*AlarmAnalysisConfig `json:"ConfigInfo,omitnil,omitempty" name:"ConfigInfo"`
}

type AnonymousInfo struct {
	// 操作列表，支持trackLog(JS/HTTP上传日志  )和realtimeProducer(kafka协议上传日志)
	Operations []*string `json:"Operations,omitnil,omitempty" name:"Operations"`

	// 条件列表
	Conditions []*ConditionInfo `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

// Predefined struct for user
type ApplyConfigToMachineGroupRequestParams struct {
	// 采集配置ID
	//  - 通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)获取采集配置Id。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type ApplyConfigToMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// 采集配置ID
	//  - 通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)获取采集配置Id。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *ApplyConfigToMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConfigToMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyConfigToMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyConfigToMachineGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyConfigToMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *ApplyConfigToMachineGroupResponseParams `json:"Response"`
}

func (r *ApplyConfigToMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConfigToMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallBackInfo struct {
	// 回调时的Body。
	// 可将各类告警变量放在请求内容中，详见[帮助文档](https://cloud.tencent.com/document/product/614/74718)。
	// 如下示例：
	// 
	// ```
	// {
	// "TopicId": "{{ .QueryLog[0][0].topicId }}",
	// "key": "{{.Alarm}}",
	// "time": "{{ .QueryLog[0][0].time }}",
	// "log": "{{ .QueryLog[0][0].content.__CONTENT__ }}",
	// "namespace": "{{ .QueryLog[0][0].content.__TAG__.namespace }}"
	// }
	// ```
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 回调时的HTTP请求头部字段。
	// 例如：下面请求头部字段来告知服务器请求主体的内容类型为JSON。
	// ```
	// "Content-Type: application/json"
	// ```
	Headers []*string `json:"Headers,omitnil,omitempty" name:"Headers"`
}

// Predefined struct for user
type CheckFunctionRequestParams struct {
	// 加工语句。 当FuncType为2时，EtlContent必须使用[log_auto_output](https://cloud.tencent.com/document/product/614/70733#b3c58797-4825-4807-bef4-68106e25024f) 
	// 
	// 其他参考文档：
	// 
	// - [创建加工任务](https://cloud.tencent.com/document/product/614/63940) 
	// -  [函数总览](https://cloud.tencent.com/document/product/614/70395)
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// 加工任务目标topic_id以及别名，当 FuncType 为 1 时，必填。
	// 目标日志主题ID通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// 数据加工目标主题的类型. 1 固定主题 2动态创建
	FuncType *int64 `json:"FuncType,omitnil,omitempty" name:"FuncType"`
}

type CheckFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 加工语句。 当FuncType为2时，EtlContent必须使用[log_auto_output](https://cloud.tencent.com/document/product/614/70733#b3c58797-4825-4807-bef4-68106e25024f) 
	// 
	// 其他参考文档：
	// 
	// - [创建加工任务](https://cloud.tencent.com/document/product/614/63940) 
	// -  [函数总览](https://cloud.tencent.com/document/product/614/70395)
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// 加工任务目标topic_id以及别名，当 FuncType 为 1 时，必填。
	// 目标日志主题ID通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// 数据加工目标主题的类型. 1 固定主题 2动态创建
	FuncType *int64 `json:"FuncType,omitnil,omitempty" name:"FuncType"`
}

func (r *CheckFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EtlContent")
	delete(f, "DstResources")
	delete(f, "FuncType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckFunctionResponseParams struct {
	// 失败错误码
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 失败错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckFunctionResponse struct {
	*tchttp.BaseResponse
	Response *CheckFunctionResponseParams `json:"Response"`
}

func (r *CheckFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckRechargeKafkaServerRequestParams struct {
	// 导入Kafka类型，0: 腾讯云CKafka；1: 用户自建Kafka。
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// 腾讯云CKafka实例ID。
	// KafkaType为0时，KafkaInstance必填
	// 
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取实例id。
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// 服务地址。
	// KafkaType为1时，ServerAddr必填
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// ServerAddr是否为加密连接，默认值false。当KafkaType为1用户自建kafka时生效。
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// 加密访问协议。KafkaType参数为1并且IsEncryptionAddr参数为true时必填。
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 用户kafka拓展信息
	UserKafkaMeta *UserKafkaMeta `json:"UserKafkaMeta,omitnil,omitempty" name:"UserKafkaMeta"`
}

type CheckRechargeKafkaServerRequest struct {
	*tchttp.BaseRequest
	
	// 导入Kafka类型，0: 腾讯云CKafka；1: 用户自建Kafka。
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// 腾讯云CKafka实例ID。
	// KafkaType为0时，KafkaInstance必填
	// 
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取实例id。
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// 服务地址。
	// KafkaType为1时，ServerAddr必填
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// ServerAddr是否为加密连接，默认值false。当KafkaType为1用户自建kafka时生效。
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// 加密访问协议。KafkaType参数为1并且IsEncryptionAddr参数为true时必填。
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 用户kafka拓展信息
	UserKafkaMeta *UserKafkaMeta `json:"UserKafkaMeta,omitnil,omitempty" name:"UserKafkaMeta"`
}

func (r *CheckRechargeKafkaServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRechargeKafkaServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KafkaType")
	delete(f, "KafkaInstance")
	delete(f, "ServerAddr")
	delete(f, "IsEncryptionAddr")
	delete(f, "Protocol")
	delete(f, "UserKafkaMeta")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckRechargeKafkaServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckRechargeKafkaServerResponseParams struct {
	// Kafka集群可访问状态。
	// 
	// - 0：可正常访问 
	// - -1：broker 连接失败
	// - -2：sasl 鉴权失败
	// - -3：ckafka 角色未授权
	// - -4：topic 列表不存在
	// - -5：topic 内暂无数据
	// - -6：用户没有 ckafka 权限
	// - -7：消费组已经存在
	// - -8：kafka 实例不存在或已销毁
	// - -9：Broker 列表为空
	// - -10：Broker 地址格式不正确
	// - -11：Broker 端口非整型
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckRechargeKafkaServerResponse struct {
	*tchttp.BaseResponse
	Response *CheckRechargeKafkaServerResponseParams `json:"Response"`
}

func (r *CheckRechargeKafkaServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRechargeKafkaServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ckafka struct {
	// Ckafka 的 InstanceId。
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取实例id。
	// - 通过 [创建实例](https://cloud.tencent.com/document/product/597/53207) 获取实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Ckafka 的 TopicName。
	// - 通过 [创建 Topic](https://cloud.tencent.com/document/product/597/73566) 获得TopicName。
	// - 通过 [获取主题列表](https://cloud.tencent.com/document/product/597/40847) 获得TopicName。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Ckafka 的 Vip。
	// - 通过 [获取实例属性 ](https://cloud.tencent.com/document/product/597/40836) 获取vip信息。
	// - 如果是通过 角色ARN 方式创建投递任务，则Vip字段可为空。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Ckafka 的 Vport。
	// - 通过 [获取实例属性 ](https://cloud.tencent.com/document/product/597/40836) 获取vip port信息。
	// - 如果是通过 角色ARN 方式创建投递任务，则Vport字段可为空。
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Ckafka 的 InstanceName。
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取InstanceName。
	// - 通过 [创建实例](https://cloud.tencent.com/document/product/597/53207) 获取InstanceName。
	// - 如果是通过 角色ARN 方式创建投递任务，则InstanceName字段可为空。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Ckafka 的 TopicId。
	// - 通过 [创建 Topic](https://cloud.tencent.com/document/product/597/73566) 获得TopicId。
	// - 通过 [获取主题列表](https://cloud.tencent.com/document/product/597/40847) 获得TopicId。
	// - 如果是通过 角色ARN 方式创建投递任务，则TopicId字段可为空。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

// Predefined struct for user
type CloseKafkaConsumerRequestParams struct {
	// 日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`
}

type CloseKafkaConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`
}

func (r *CloseKafkaConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseKafkaConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseKafkaConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseKafkaConsumerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseKafkaConsumerResponse struct {
	*tchttp.BaseResponse
	Response *CloseKafkaConsumerResponseParams `json:"Response"`
}

func (r *CloseKafkaConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudProductLogTaskInfo struct {
	// 日志服务地域
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志配置拓展信息， 一般用于存储额外的日志投递配置
	Extend *string `json:"Extend,omitnil,omitempty" name:"Extend"`

	// 日志类型，支持枚举：CDS-AUDIT、CDS-RISK、CDB-AUDIT、TDSQL-C-AUDIT、MongoDB-AUDIT、MongoDB-SlowLog、MongoDB-ErrorLog、TDMYSQL-SLOW、DCDB-AUDIT、DCDB-SLOW、DCDB-ERROR、MariaDB-AUDIT、MariaDB-SLOW、MariaDB-ERROR、PostgreSQL-SLOW、PostgreSQL-ERROR、PostgreSQL-AUDIT、BH-FILELOG、BH-COMMANDLOG、APIS-ACCESS
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 任务状态， 0创建中 1创建完成 2 删除中 
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type CollectConfig struct {
	// 指定采集类型的采集配置名称信息。
	// <li>当CollectInfo中Type为0：表示元数据配置，name为元数据名称。
	// 目前支持"container_id"，"container_name"，"image_name"，"namespace"，"pod_uid"，"pod_name"，"pod_ip"。
	// </li>
	// <li>当CollectInfo中Type为1：指定pod label，name为指定pod label名称。</li>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CollectInfo struct {
	// 采集类型，必填字段。
	// <li>0：元数据配置。</li>
	// <li>1：指定Pod Label。</li>
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 指定采集类型的采集配置信息。
	// <li>当Type为0时，CollectConfigs不允许为空。</li>
	// <li>当Type为1时，CollectConfigs为空时，表示选择所有Pod Label；否则CollectConfigs为指定Pod Label。</li>
	CollectConfigs []*CollectConfig `json:"CollectConfigs,omitnil,omitempty" name:"CollectConfigs"`
}

type Column struct {
	// 列的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 列的属性
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CompressInfo struct {
	// 压缩格式，支持gzip、lzop、snappy和none不压缩
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type ConditionInfo struct {
	// 条件属性，目前只支持VpcID
	Attributes *string `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// 条件规则，1:等于，2:不等于
	Rule *uint64 `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 对应条件属性的值
	ConditionValue *string `json:"ConditionValue,omitnil,omitempty" name:"ConditionValue"`
}

type ConfigExtraInfo struct {
	// 采集规则扩展配置ID
	ConfigExtraId *string `json:"ConfigExtraId,omitnil,omitempty" name:"ConfigExtraId"`

	// 采集规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 自建k8s集群日志采集类型，支持
	// - container_stdout 标准输出
	// - container_file 标准文件
	// - host_file 节点文件
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点文件配置信息
	HostFile *HostFileInfo `json:"HostFile,omitnil,omitempty" name:"HostFile"`

	// 容器文件路径信息
	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitnil,omitempty" name:"ContainerFile"`

	// 容器标准输出信息
	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitnil,omitempty" name:"ContainerStdout"`

	// 日志格式化方式
	LogFormat *string `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 更新时间
	// - 时间格式：yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建时间
	// - 时间格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户自定义解析字符串
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 自建采集配置标
	ConfigFlag *string `json:"ConfigFlag,omitnil,omitempty" name:"ConfigFlag"`

	// 日志集ID
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集name
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题name
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 采集相关配置信息。详情见 CollectInfo复杂类型配置。
	CollectInfos []*CollectInfo `json:"CollectInfos,omitnil,omitempty" name:"CollectInfos"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// 样例：{"ClsAgentFileTimeout":0,"ClsAgentMaxDepth":10,"ClsAgentParseFailMerge":true}
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type ConfigInfo struct {
	// 采集规则配置ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 采集规则配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志格式化方式
	LogFormat *string `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// 日志采集路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 采集的日志类型。
	// - json_log代表：JSON-文件日志（详见[使用 JSON 提取模式采集日志](https://cloud.tencent.com/document/product/614/17419)）；
	// - delimiter_log代表：分隔符-文件日志（详见[使用分隔符提取模式采集日志](https://cloud.tencent.com/document/product/614/17420)）；
	// - minimalist_log代表：单行全文-文件日志（详见[使用单行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17421)）；
	// - fullregex_log代表：单行完全正则-文件日志（详见[使用单行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52365)）；
	// - multiline_log代表：多行全文-文件日志（详见[使用多行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17422)）；
	// - multiline_fullregex_log代表：多行完全正则-文件日志（详见[使用多行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52366)）；
	// - user_define_log代表：组合解析（适用于多格式嵌套的日志，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)）；
	// - service_syslog代表：syslog 采集（详见[采集 Syslog](https://cloud.tencent.com/document/product/614/81454)）；
	// - windows_event_log代表：Windows事件日志（详见[采集 Windows 事件日志](https://cloud.tencent.com/document/product/614/96678)）。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 采集配置所属日志主题ID即TopicId
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 更新时间
	// - 时间格式：yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建时间
	// - 时间格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户自定义解析字符串，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)。
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// 样例：
	// `{\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true}`
	// 
	// 控制台默认占位值：`{\"ClsAgentDefault\":0}`
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`

	// 日志输入类型（<span style="color:red; font-weight:bold">注：windows场景必填且仅支持file和windows_event类型</span>）
	// - file: 文件类型采集
	// - windows_event：windows事件采集
	// - syslog：系统日志采集
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`
}

type ConsoleSharingConfig struct {
	// 分享链接名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 仪表盘: 1; 检索页:2
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 分享链接有效期，单位：毫秒，最长支持30天
	DurationMilliseconds *uint64 `json:"DurationMilliseconds,omitnil,omitempty" name:"DurationMilliseconds"`

	// 允许访问的资源列表，目前仅支持一个Resource
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 分享链接域名，可选范围
	// - 公网匿名分享：填写clsshare.com
	// - datasight内网匿名分享(若开启)：datasight内网域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分享链接加密访问验证码。支持0-9和a-z(不区分大小写)在内的6个字符，可为空，代表免验证码访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyCode *string `json:"VerifyCode,omitnil,omitempty" name:"VerifyCode"`

	// 默认查询范围的开始时间点，支持绝对时间(13位Unix时间戳)或相对时间表达式
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 默认查询范围的结束时间点，支持绝对时间(13位Unix时间戳)或相对时间表达式。注意，结束时间点要大于开始时间点
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 仅当StartTime/EndTime为相对时间时使用，基于NowTime计算绝对时间，默认为创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NowTime *uint64 `json:"NowTime,omitnil,omitempty" name:"NowTime"`

	// 默认的检索分析语句，仅当Type为2时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*ConsoleSharingParam `json:"Params,omitnil,omitempty" name:"Params"`

	// 是否允许访问者自行修改检索分析时间范围。默认不锁定（false）
	IsLockTimeRange *bool `json:"IsLockTimeRange,omitnil,omitempty" name:"IsLockTimeRange"`

	// 是否允许访问者自行修改日志检索语句。在检索页分享中表示检索语句锁定状态；在仪表盘中表示过滤变量锁定状态。默认不锁定（false）
	IsLockQuery *bool `json:"IsLockQuery,omitnil,omitempty" name:"IsLockQuery"`

	// 检索页分享是否允许访问者下载日志，默认不允许（false）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportLogExport *bool `json:"IsSupportLogExport,omitnil,omitempty" name:"IsSupportLogExport"`
}

type ConsoleSharingInfo struct {
	// 分享ID
	SharingId *string `json:"SharingId,omitnil,omitempty" name:"SharingId"`

	// 分享链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	SharingUrl *string `json:"SharingUrl,omitnil,omitempty" name:"SharingUrl"`

	// 匿名分享配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SharingConfig *ConsoleSharingConfig `json:"SharingConfig,omitnil,omitempty" name:"SharingConfig"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *uint64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 分享链接状态
	// 1: 正常 
	// -1: 因内容安全审查异常导致被封禁(存在于使用公网域名分享时)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 10001-广告 20001-政治 20002-色情 20004-社会事件 20011-暴力 20012-低俗 20006-违法犯罪 20007-谩骂 20008-欺诈 20013-版权 20104-谣言 21000-其他, 10086-聚合, 24001-暴恐（天御独有恶意类型），20472-违法，
	// 24005-社会
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentSafetyCode *uint64 `json:"ContentSafetyCode,omitnil,omitempty" name:"ContentSafetyCode"`
}

type ConsoleSharingParam struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ConsumerContent struct {
	// 是否投递 TAG 信息。
	// 当EnableTag为true时，表示投递TAG元信息。
	EnableTag *bool `json:"EnableTag,omitnil,omitempty" name:"EnableTag"`

	// 需要投递的元数据列表，目前仅支持：\_\_SOURCE\_\_，\_\_FILENAME\_\_，\_\_TIMESTAMP\_\_，\_\_HOSTNAME\_\_和\_\_PKGID\_\_
	MetaFields []*string `json:"MetaFields,omitnil,omitempty" name:"MetaFields"`

	// 当EnableTag为true时，必须填写TagJsonNotTiled字段。
	// TagJsonNotTiled用于标识tag信息是否json平铺。
	// 
	// TagJsonNotTiled为true时不平铺，示例：
	// TAG信息：`{"__TAG__":{"fieldA":200,"fieldB":"text"}}`
	// 不平铺：`{"__TAG__":{"fieldA":200,"fieldB":"text"}}`
	// 
	// TagJsonNotTiled为false时平铺，示例：
	// TAG信息：`{"__TAG__":{"fieldA":200,"fieldB":"text"}}`
	// 平铺：`{"__TAG__.fieldA":200,"__TAG__.fieldB":"text"}`
	TagJsonNotTiled *bool `json:"TagJsonNotTiled,omitnil,omitempty" name:"TagJsonNotTiled"`

	// 投递时间戳精度，可选项 [1：秒；2：毫秒] ，默认是1。
	TimestampAccuracy *int64 `json:"TimestampAccuracy,omitnil,omitempty" name:"TimestampAccuracy"`

	// 投递Json格式。
	// JsonType为0：和原始日志一致，不转义。示例：
	// 日志原文：`{"a":"aa", "b":{"b1":"b1b1", "c1":"c1c1"}}`
	// 投递到Ckafka：`{"a":"aa", "b":{"b1":"b1b1", "c1":"c1c1"}}`
	// 
	// JsonType为1：转义。示例：
	// 日志原文：`{"a":"aa", "b":{"b1":"b1b1", "c1":"c1c1"}}`
	// 投递到Ckafka：`{"a":"aa","b":"{\"b1\":\"b1b1\", \"c1\":\"c1c1\"}"}`
	JsonType *int64 `json:"JsonType,omitnil,omitempty" name:"JsonType"`
}

type ConsumerGroup struct {
	// 消费组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 状态。
	// 
	// - Empty：组内没有成员，但存在已提交的偏移量。所有消费者都离开但保留了偏移量
	// - Dead：组内没有成员，且没有已提交的偏移量。组被删除或长时间无活动
	// - Stable：组内成员正常消费，分区分配平衡。正常运行状态
	// - PreparingRebalance：组正在准备重新平衡。有新成员加入或现有成员离开
	// - CompletingRebalance：组正在准备重新平衡。有新成员加入或现有成员离开
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 分区分配策略均衡算法名称。
	// 
	// - 常见均衡算法如下：
	//     - range:按分区范围分配
	//     - roundrobin:轮询式分配
	//     - sticky:粘性分配（避免不必要的重平衡）
	ProtocolName *string `json:"ProtocolName,omitnil,omitempty" name:"ProtocolName"`
}

type ContainerFileInfo struct {
	// namespace可以多个，用分隔号分割,例如A,B
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 容器名称
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 日志文件夹
	LogPath *string `json:"LogPath,omitnil,omitempty" name:"LogPath"`

	// 日志名称
	FilePattern *string `json:"FilePattern,omitnil,omitempty" name:"FilePattern"`

	// 日志文件信息
	FilePaths []*FilePathInfo `json:"FilePaths,omitnil,omitempty" name:"FilePaths"`

	// pod标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeLabels []*string `json:"IncludeLabels,omitnil,omitempty" name:"IncludeLabels"`

	// 工作负载信息
	WorkLoad *ContainerWorkLoadInfo `json:"WorkLoad,omitnil,omitempty" name:"WorkLoad"`

	// 需要排除的namespace可以多个，用分隔号分割,例如A,B
	ExcludeNamespace *string `json:"ExcludeNamespace,omitnil,omitempty" name:"ExcludeNamespace"`

	// 需要排除的pod标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeLabels []*string `json:"ExcludeLabels,omitnil,omitempty" name:"ExcludeLabels"`

	// metadata信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomLabels []*string `json:"CustomLabels,omitnil,omitempty" name:"CustomLabels"`
}

type ContainerStdoutInfo struct {
	// 是否所有容器
	AllContainers *bool `json:"AllContainers,omitnil,omitempty" name:"AllContainers"`

	// container为空表所有的，不为空采集指定的容器
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// namespace可以多个，用分隔号分割,例如A,B；为空或者没有这个字段，表示所有namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// pod标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeLabels []*string `json:"IncludeLabels,omitnil,omitempty" name:"IncludeLabels"`

	// 工作负载信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkLoads []*ContainerWorkLoadInfo `json:"WorkLoads,omitnil,omitempty" name:"WorkLoads"`

	// 需要排除的namespace可以多个，用分隔号分割,例如A,B
	ExcludeNamespace *string `json:"ExcludeNamespace,omitnil,omitempty" name:"ExcludeNamespace"`

	// 需要排除的pod标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeLabels []*string `json:"ExcludeLabels,omitnil,omitempty" name:"ExcludeLabels"`

	// metadata信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomLabels []*string `json:"CustomLabels,omitnil,omitempty" name:"CustomLabels"`
}

type ContainerWorkLoadInfo struct {
	// 工作负载的类型，支持的取值如下：
	// - deployment：无状态应用部署控制器。可管理无状态应用的多副本 Pod；支持滚动更新和回滚，实现无缝发布；通过 ReplicaSet 确保指定数量的 Pod 副本持续运行。适合 Web 服务、API 服务、微服务等不需要持久化存储或唯一标识的应用。
	// 
	// - statefulset：用于管理有状态应用的控制器。可以为每个 Pod 提供稳定的、唯一的标识（有序的主机名和网络标识）；能提供稳定的持久化存储（PVC 随 Pod 的迁移而保留）；Pod 的创建、扩展、删除和更新都遵循确定的顺序。适合数据库（MySQL、MongoDB）、消息队列（Kafka、RabbitMQ）、集群应用等需要稳定网络标识和持久化存储的应用。
	// 
	// - daemonset：确保所有（或特定）节点上都运行一个 Pod 副本的控制器。会在每个符合条件的 Node 上自动部署并运行一个 Pod；当新节点加入集群时，会自动在新节点上创建 Pod；适合运行节点级别的后台任务或服务，例如日志收集（Fluentd、Filebeat）、节点监控（Node Exporter）、网络插件（Calico、Weave Net）等场景。
	// 
	// - job：用于运行一次性任务的控制器。可创建一个或多个 Pod 来执行任务，直到成功完成；当任务完成后，Pod 不会重启（除非配置了重启策略）；可以指定任务的并行度和重试次数。适合数据处理、批量任务、数据库迁移、离线计算等一次性执行完成后就退出的任务。
	// 
	// - cronjob：基于时间调度的 Job 控制器。类似于 Linux 的 cron，按照预定的时间表周期性地创建并运行 Job。每个调度周期都会创建一个新的 Job 来执行任务。适合定期备份、发送报告、数据清理、定时同步等需要周期性执行的任务。
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 工作负载的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 容器名
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type ContentInfo struct {
	// 内容格式，支持json，csv，parquet
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// csv格式内容描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Csv *CsvInfo `json:"Csv,omitnil,omitempty" name:"Csv"`

	// json格式内容描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *JsonInfo `json:"Json,omitnil,omitempty" name:"Json"`

	// parquet格式内容描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parquet *ParquetInfo `json:"Parquet,omitnil,omitempty" name:"Parquet"`
}

type CosRechargeInfo struct {
	// COS导入配置ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// COS导入任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// COS存储桶
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// COS存储桶所在地域。
	// 
	// - 通过[地域和访问域名](https://cloud.tencent.com/document/product/436/6224)获取地域信息。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// COS文件所在文件夹的前缀
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表单行全文；
	// 默认为minimalist_log
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 状态   status 0: 已创建, 1: 运行中, 2: 已停止, 3: 已完成, 4: 运行失败。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否启用:   0： 未启用  ， 1：启用
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 创建时间。时间格式：YYYY-MM-DD HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。时间格式：YYYY-MM-DD HH:mm:ss
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 进度条百分值
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 压缩方式supported: "", "gzip", "lzop", "snappy”;  默认空不压缩
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 见： ExtractRuleInfo 结构描述
	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitnil,omitempty" name:"ExtractRuleInfo"`

	// COS导入任务类型。1：一次性导入任务；2：持续性导入任务。
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 元数据。支持 bucket，object。
	Metadata []*string `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

// Predefined struct for user
type CreateAlarmNoticeRequestParams struct {
	// 通知渠道组名称。最大支持255个字节。 不支持 '|'。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的通知渠道组。最大支持50个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 【简易模式】（简易模式/告警模式二选一，分别配置相应参数）
	// 需要发送通知的告警类型。可选值：
	// - Trigger - 告警触发
	// - Recovery - 告警恢复
	// - All - 告警触发和告警恢复
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 【简易模式】（简易模式/告警模式二选一，分别配置相应参数）
	// 通知接收对象。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// 【简易模式】（简易模式/告警模式二选一，分别配置相应参数）
	// 接口回调信息（包括企业微信、钉钉、飞书）。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`

	// 【高级模式】（简易模式/告警模式二选一，分别配置相应参数）
	// 通知规则。
	NoticeRules []*NoticeRule `json:"NoticeRules,omitnil,omitempty" name:"NoticeRules"`

	// 查询数据链接。http:// 或者 https:// 开头，不能/结尾
	JumpDomain *string `json:"JumpDomain,omitnil,omitempty" name:"JumpDomain"`

	// 投递日志开关。可取值如下：
	// 1：关闭（默认值）；
	// 2：开启 
	// 投递日志开关开启时， DeliverConfig参数必填。
	DeliverStatus *uint64 `json:"DeliverStatus,omitnil,omitempty" name:"DeliverStatus"`

	// 投递日志配置参数。当DeliverStatus开启时，必填。
	DeliverConfig *DeliverConfig `json:"DeliverConfig,omitnil,omitempty" name:"DeliverConfig"`

	// 免登录操作告警开关。可取值如下：
	// -      1：关闭
	// -      2：开启（默认值）
	AlarmShieldStatus *uint64 `json:"AlarmShieldStatus,omitnil,omitempty" name:"AlarmShieldStatus"`

	// 统一设定自定义回调参数。
	// -  true: 使用通知内容模板中的自定义回调参数覆盖告警策略中单独配置的请求头及请求内容。
	// -  false:优先使用告警策略中单独配置的请求头及请求内容。
	CallbackPrioritize *bool `json:"CallbackPrioritize,omitnil,omitempty" name:"CallbackPrioritize"`
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest
	
	// 通知渠道组名称。最大支持255个字节。 不支持 '|'。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的通知渠道组。最大支持50个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 【简易模式】（简易模式/告警模式二选一，分别配置相应参数）
	// 需要发送通知的告警类型。可选值：
	// - Trigger - 告警触发
	// - Recovery - 告警恢复
	// - All - 告警触发和告警恢复
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 【简易模式】（简易模式/告警模式二选一，分别配置相应参数）
	// 通知接收对象。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// 【简易模式】（简易模式/告警模式二选一，分别配置相应参数）
	// 接口回调信息（包括企业微信、钉钉、飞书）。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`

	// 【高级模式】（简易模式/告警模式二选一，分别配置相应参数）
	// 通知规则。
	NoticeRules []*NoticeRule `json:"NoticeRules,omitnil,omitempty" name:"NoticeRules"`

	// 查询数据链接。http:// 或者 https:// 开头，不能/结尾
	JumpDomain *string `json:"JumpDomain,omitnil,omitempty" name:"JumpDomain"`

	// 投递日志开关。可取值如下：
	// 1：关闭（默认值）；
	// 2：开启 
	// 投递日志开关开启时， DeliverConfig参数必填。
	DeliverStatus *uint64 `json:"DeliverStatus,omitnil,omitempty" name:"DeliverStatus"`

	// 投递日志配置参数。当DeliverStatus开启时，必填。
	DeliverConfig *DeliverConfig `json:"DeliverConfig,omitnil,omitempty" name:"DeliverConfig"`

	// 免登录操作告警开关。可取值如下：
	// -      1：关闭
	// -      2：开启（默认值）
	AlarmShieldStatus *uint64 `json:"AlarmShieldStatus,omitnil,omitempty" name:"AlarmShieldStatus"`

	// 统一设定自定义回调参数。
	// -  true: 使用通知内容模板中的自定义回调参数覆盖告警策略中单独配置的请求头及请求内容。
	// -  false:优先使用告警策略中单独配置的请求头及请求内容。
	CallbackPrioritize *bool `json:"CallbackPrioritize,omitnil,omitempty" name:"CallbackPrioritize"`
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
	delete(f, "Name")
	delete(f, "Tags")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	delete(f, "NoticeRules")
	delete(f, "JumpDomain")
	delete(f, "DeliverStatus")
	delete(f, "DeliverConfig")
	delete(f, "AlarmShieldStatus")
	delete(f, "CallbackPrioritize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmNoticeResponseParams struct {
	// 告警模板ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateAlarmRequestParams struct {
	// 告警策略名称。最大支持255个字节。 不支持 '|'。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 监控对象列表。
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为2000。
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// 告警重复的周期，单位是分钟。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知渠道组列表。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/product/614/56462)获取关联的告警通知渠道组列表，和MonitorNotice互斥
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// 告警发送通知的触发条件
	//  注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 告警级别
	// 0:警告(Warn); 1:提醒(Info); 2:紧急 (Critical)。
	// 注意:  
	// - 不填则默认为0。
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 多触发条件
	//  注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	// 
	// 
	MultiConditions []*MultiCondition `json:"MultiConditions,omitnil,omitempty" name:"MultiConditions"`

	// 是否开启告警策略。
	// 默认值为true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 请使用Status参数控制是否开启告警策略。
	//
	// Deprecated: Enable is deprecated.
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 用户自定义告警内容
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// 用户自定义回调
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// 多维分析
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// 分组触发状态。
	// 默认值false
	GroupTriggerStatus *bool `json:"GroupTriggerStatus,omitnil,omitempty" name:"GroupTriggerStatus"`

	// 分组触发条件。
	GroupTriggerCondition []*string `json:"GroupTriggerCondition,omitnil,omitempty" name:"GroupTriggerCondition"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的告警策略。
	// 
	// 最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 监控对象类型。0:执行语句共用监控对象; 1:每个执行语句单独选择监控对象。 
	// 不填则默认为0。
	// 当值为1时，AlarmTargets元素个数不能超过10个，AlarmTargets中的Number必须是从1开始的连续正整数，不能重复。
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 告警附加分类信息列表。
	// Classifications元素个数不能超过20个。
	// Classifications元素的Key不能为空，不能重复，长度不能超过50个字符，符合正则 `^[a-z]([a-z0-9_]{0,49})$`。
	// Classifications元素的Value长度不能超过200个字符。
	Classifications []*AlarmClassification `json:"Classifications,omitnil,omitempty" name:"Classifications"`
}

type CreateAlarmRequest struct {
	*tchttp.BaseRequest
	
	// 告警策略名称。最大支持255个字节。 不支持 '|'。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 监控对象列表。
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为2000。
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// 告警重复的周期，单位是分钟。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知渠道组列表。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/product/614/56462)获取关联的告警通知渠道组列表，和MonitorNotice互斥
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// 告警发送通知的触发条件
	//  注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 告警级别
	// 0:警告(Warn); 1:提醒(Info); 2:紧急 (Critical)。
	// 注意:  
	// - 不填则默认为0。
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 多触发条件
	//  注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	// 
	// 
	MultiConditions []*MultiCondition `json:"MultiConditions,omitnil,omitempty" name:"MultiConditions"`

	// 是否开启告警策略。
	// 默认值为true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 请使用Status参数控制是否开启告警策略。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 用户自定义告警内容
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// 用户自定义回调
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// 多维分析
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// 分组触发状态。
	// 默认值false
	GroupTriggerStatus *bool `json:"GroupTriggerStatus,omitnil,omitempty" name:"GroupTriggerStatus"`

	// 分组触发条件。
	GroupTriggerCondition []*string `json:"GroupTriggerCondition,omitnil,omitempty" name:"GroupTriggerCondition"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的告警策略。
	// 
	// 最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 监控对象类型。0:执行语句共用监控对象; 1:每个执行语句单独选择监控对象。 
	// 不填则默认为0。
	// 当值为1时，AlarmTargets元素个数不能超过10个，AlarmTargets中的Number必须是从1开始的连续正整数，不能重复。
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 告警附加分类信息列表。
	// Classifications元素个数不能超过20个。
	// Classifications元素的Key不能为空，不能重复，长度不能超过50个字符，符合正则 `^[a-z]([a-z0-9_]{0,49})$`。
	// Classifications元素的Value长度不能超过200个字符。
	Classifications []*AlarmClassification `json:"Classifications,omitnil,omitempty" name:"Classifications"`
}

func (r *CreateAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AlarmTargets")
	delete(f, "MonitorTime")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "Condition")
	delete(f, "AlarmLevel")
	delete(f, "MultiConditions")
	delete(f, "Status")
	delete(f, "Enable")
	delete(f, "MessageTemplate")
	delete(f, "CallBack")
	delete(f, "Analysis")
	delete(f, "GroupTriggerStatus")
	delete(f, "GroupTriggerCondition")
	delete(f, "Tags")
	delete(f, "MonitorObjectType")
	delete(f, "Classifications")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmResponseParams struct {
	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlarmResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlarmResponseParams `json:"Response"`
}

func (r *CreateAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmShieldRequestParams struct {
	// 通知渠道组id。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/product/614/56462)获取通知渠道组id
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// 屏蔽规则开始时间，秒级(s)时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 屏蔽规则结束时间，秒级(s)时间戳。结束时间需要大于当前时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 屏蔽类型。1：屏蔽所有通知，2：按照Rule参数屏蔽匹配规则的通知。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 屏蔽原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 屏蔽规则，当Type为2时必填。规则填写方式详见[产品文档](https://cloud.tencent.com/document/product/614/103178#rule)。
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type CreateAlarmShieldRequest struct {
	*tchttp.BaseRequest
	
	// 通知渠道组id。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/product/614/56462)获取通知渠道组id
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// 屏蔽规则开始时间，秒级(s)时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 屏蔽规则结束时间，秒级(s)时间戳。结束时间需要大于当前时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 屏蔽类型。1：屏蔽所有通知，2：按照Rule参数屏蔽匹配规则的通知。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 屏蔽原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 屏蔽规则，当Type为2时必填。规则填写方式详见[产品文档](https://cloud.tencent.com/document/product/614/103178#rule)。
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *CreateAlarmShieldRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmShieldRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Reason")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmShieldRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmShieldResponseParams struct {
	// 屏蔽规则ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlarmShieldResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlarmShieldResponseParams `json:"Response"`
}

func (r *CreateAlarmShieldResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmShieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudProductLogCollectionRequestParams struct {
	// 实例ID
	// - 通过各个接入云产品官方文档获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云产品标识，支持枚举：CDS、CWP、CDB、TDSQL-C、MongoDB、TDStore、DCDB、MariaDB、PostgreSQL、BH、APIS
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// 日志类型，支持枚举：CDS-AUDIT、CDS-RISK、CDB-AUDIT、TDSQL-C-AUDIT、MongoDB-AUDIT、MongoDB-SlowLog、MongoDB-ErrorLog、TDMYSQL-SLOW、DCDB-AUDIT、DCDB-SLOW、DCDB-ERROR、MariaDB-AUDIT、MariaDB-SLOW、MariaDB-ERROR、PostgreSQL-SLOW、PostgreSQL-ERROR、PostgreSQL-AUDIT、BH-FILELOG、BH-COMMANDLOG、APIS-ACCESS
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 云产品地域。 不同日志类型(LogType)地域入参格式存在差异， 请参考如下示例：
	// - CDS所有日志类型：ap-guangzhou
	// - CDB-AUDIT: gz
	// - TDSQL-C-AUDIT:  gz
	// - MongoDB-AUDIT:  gz
	// - MongoDB-SlowLog：ap-guangzhou
	// - MongoDB-ErrorLog：ap-guangzhou
	// - TDMYSQL-SLOW：gz
	// - DCDB所有日志类型：gz
	// - MariaDB所有日志类型：gz
	// - PostgreSQL所有日志类型：gz
	// - BH所有日志类型：overseas-polaris(中国香港地区和其他)/fsi-polaris(金融区)/general-polaris(普通区)/intl-sg-prod(国际站)
	// - APIS所有日志类型：gz
	CloudProductRegion *string `json:"CloudProductRegion,omitnil,omitempty" name:"CloudProductRegion"`

	// CLS目标地域
	// - 支持地域参考  [地域列表](https://cloud.tencent.com/document/api/614/56474#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8) 文档   
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`

	// 日志集名称，未填LogsetId时必填。若日志集不存在, 将自动创建
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题名称，在未填TopicId时必填。 若日志主题不存在，将自动创建
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志配置扩展信息， 一般用于存储额外的日志投递配置
	Extend *string `json:"Extend,omitnil,omitempty" name:"Extend"`

	// 日志集id
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志主题id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type CreateCloudProductLogCollectionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	// - 通过各个接入云产品官方文档获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云产品标识，支持枚举：CDS、CWP、CDB、TDSQL-C、MongoDB、TDStore、DCDB、MariaDB、PostgreSQL、BH、APIS
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// 日志类型，支持枚举：CDS-AUDIT、CDS-RISK、CDB-AUDIT、TDSQL-C-AUDIT、MongoDB-AUDIT、MongoDB-SlowLog、MongoDB-ErrorLog、TDMYSQL-SLOW、DCDB-AUDIT、DCDB-SLOW、DCDB-ERROR、MariaDB-AUDIT、MariaDB-SLOW、MariaDB-ERROR、PostgreSQL-SLOW、PostgreSQL-ERROR、PostgreSQL-AUDIT、BH-FILELOG、BH-COMMANDLOG、APIS-ACCESS
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 云产品地域。 不同日志类型(LogType)地域入参格式存在差异， 请参考如下示例：
	// - CDS所有日志类型：ap-guangzhou
	// - CDB-AUDIT: gz
	// - TDSQL-C-AUDIT:  gz
	// - MongoDB-AUDIT:  gz
	// - MongoDB-SlowLog：ap-guangzhou
	// - MongoDB-ErrorLog：ap-guangzhou
	// - TDMYSQL-SLOW：gz
	// - DCDB所有日志类型：gz
	// - MariaDB所有日志类型：gz
	// - PostgreSQL所有日志类型：gz
	// - BH所有日志类型：overseas-polaris(中国香港地区和其他)/fsi-polaris(金融区)/general-polaris(普通区)/intl-sg-prod(国际站)
	// - APIS所有日志类型：gz
	CloudProductRegion *string `json:"CloudProductRegion,omitnil,omitempty" name:"CloudProductRegion"`

	// CLS目标地域
	// - 支持地域参考  [地域列表](https://cloud.tencent.com/document/api/614/56474#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8) 文档   
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`

	// 日志集名称，未填LogsetId时必填。若日志集不存在, 将自动创建
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题名称，在未填TopicId时必填。 若日志主题不存在，将自动创建
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志配置扩展信息， 一般用于存储额外的日志投递配置
	Extend *string `json:"Extend,omitnil,omitempty" name:"Extend"`

	// 日志集id
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志主题id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *CreateCloudProductLogCollectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudProductLogCollectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AssumerName")
	delete(f, "LogType")
	delete(f, "CloudProductRegion")
	delete(f, "ClsRegion")
	delete(f, "LogsetName")
	delete(f, "TopicName")
	delete(f, "Extend")
	delete(f, "LogsetId")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudProductLogCollectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudProductLogCollectionResponseParams struct {
	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名称
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// -1 创建中，1创建完成 
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudProductLogCollectionResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudProductLogCollectionResponseParams `json:"Response"`
}

func (r *CreateCloudProductLogCollectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudProductLogCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigExtraRequestParams struct {
	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志主题id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志源类型。支持 container_stdout：容器标准输出；container_file：容器文件路径；host_file：节点文件路径。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 采集的日志类型，默认为minimalist_log。支持以下类型：
	// - json_log代表：JSON-文件日志（详见[使用 JSON 提取模式采集日志](https://cloud.tencent.com/document/product/614/17419)）；
	// - delimiter_log代表：分隔符-文件日志（详见[使用分隔符提取模式采集日志](https://cloud.tencent.com/document/product/614/17420)）；
	// - minimalist_log代表：单行全文-文件日志（详见[使用单行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17421)）；
	// - fullregex_log代表：单行完全正则-文件日志（详见[使用单行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52365)）；
	// - multiline_log代表：多行全文-文件日志（详见[使用多行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17422)）；
	// - multiline_fullregex_log代表：多行完全正则-文件日志（详见[使用多行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52366)）；
	// - user_define_log代表：组合解析（适用于多格式嵌套的日志，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)）。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 采集配置标记。
	// - 目前只支持label_k8s，用于标记自建k8s集群使用的采集配置
	ConfigFlag *string `json:"ConfigFlag,omitnil,omitempty" name:"ConfigFlag"`

	// 日志集id
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名称
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集名称。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题名称
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 自建k8s-节点文件配置信息,包括文件路径、名称及元数据相关信息。
	// 
	// - 详细参考 [HostFileInfo](https://cloud.tencent.com/document/api/614/56471#HostFileInfo) 信息。
	HostFile *HostFileInfo `json:"HostFile,omitnil,omitempty" name:"HostFile"`

	// 容器文件路径类型配置。
	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitnil,omitempty" name:"ContainerFile"`

	// 自建k8s-容器标准输出信息，包括容器、命名空间等。
	// 
	// - 详细参考 [ContainerStdoutInfo](https://cloud.tencent.com/document/api/614/56471#ContainerStdoutInfo) 信息。
	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitnil,omitempty" name:"ContainerStdout"`

	// 日志格式化方式，用于容器采集场景。
	// - stdout-docker-json：用于docker容器采集场景
	// - stdout-containerd：用于containerd容器采集场景
	//
	// Deprecated: LogFormat is deprecated.
	LogFormat *string `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 组合解析采集规则，用于复杂场景下的日志采集。
	// - 取值参考：[使用组合解析提取模式采集日志
	// ](https://cloud.tencent.com/document/product/614/61310)
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 绑定的机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	// - GroupId 与 GroupIds 选择其一即可，不可同时为空。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 绑定的机器组ID列表
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id信息。
	// - GroupId 与 GroupIds 选择其一即可，不可同时为空。
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 采集相关配置信息。详细参考 [CollectInfo](https://cloud.tencent.com/document/api/614/56471#CollectInfo) 信息。
	CollectInfos []*CollectInfo `json:"CollectInfos,omitnil,omitempty" name:"CollectInfos"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// - ClsAgentDefault(自定义默认值，无特殊含义，用于清空其他选项)，建议取值0
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type CreateConfigExtraRequest struct {
	*tchttp.BaseRequest
	
	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志主题id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志源类型。支持 container_stdout：容器标准输出；container_file：容器文件路径；host_file：节点文件路径。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 采集的日志类型，默认为minimalist_log。支持以下类型：
	// - json_log代表：JSON-文件日志（详见[使用 JSON 提取模式采集日志](https://cloud.tencent.com/document/product/614/17419)）；
	// - delimiter_log代表：分隔符-文件日志（详见[使用分隔符提取模式采集日志](https://cloud.tencent.com/document/product/614/17420)）；
	// - minimalist_log代表：单行全文-文件日志（详见[使用单行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17421)）；
	// - fullregex_log代表：单行完全正则-文件日志（详见[使用单行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52365)）；
	// - multiline_log代表：多行全文-文件日志（详见[使用多行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17422)）；
	// - multiline_fullregex_log代表：多行完全正则-文件日志（详见[使用多行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52366)）；
	// - user_define_log代表：组合解析（适用于多格式嵌套的日志，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)）。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 采集配置标记。
	// - 目前只支持label_k8s，用于标记自建k8s集群使用的采集配置
	ConfigFlag *string `json:"ConfigFlag,omitnil,omitempty" name:"ConfigFlag"`

	// 日志集id
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名称
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集名称。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题名称
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 自建k8s-节点文件配置信息,包括文件路径、名称及元数据相关信息。
	// 
	// - 详细参考 [HostFileInfo](https://cloud.tencent.com/document/api/614/56471#HostFileInfo) 信息。
	HostFile *HostFileInfo `json:"HostFile,omitnil,omitempty" name:"HostFile"`

	// 容器文件路径类型配置。
	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitnil,omitempty" name:"ContainerFile"`

	// 自建k8s-容器标准输出信息，包括容器、命名空间等。
	// 
	// - 详细参考 [ContainerStdoutInfo](https://cloud.tencent.com/document/api/614/56471#ContainerStdoutInfo) 信息。
	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitnil,omitempty" name:"ContainerStdout"`

	// 日志格式化方式，用于容器采集场景。
	// - stdout-docker-json：用于docker容器采集场景
	// - stdout-containerd：用于containerd容器采集场景
	LogFormat *string `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 组合解析采集规则，用于复杂场景下的日志采集。
	// - 取值参考：[使用组合解析提取模式采集日志
	// ](https://cloud.tencent.com/document/product/614/61310)
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 绑定的机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	// - GroupId 与 GroupIds 选择其一即可，不可同时为空。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 绑定的机器组ID列表
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id信息。
	// - GroupId 与 GroupIds 选择其一即可，不可同时为空。
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 采集相关配置信息。详细参考 [CollectInfo](https://cloud.tencent.com/document/api/614/56471#CollectInfo) 信息。
	CollectInfos []*CollectInfo `json:"CollectInfos,omitnil,omitempty" name:"CollectInfos"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// - ClsAgentDefault(自定义默认值，无特殊含义，用于清空其他选项)，建议取值0
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

func (r *CreateConfigExtraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigExtraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "TopicId")
	delete(f, "Type")
	delete(f, "LogType")
	delete(f, "ConfigFlag")
	delete(f, "LogsetId")
	delete(f, "LogsetName")
	delete(f, "TopicName")
	delete(f, "HostFile")
	delete(f, "ContainerFile")
	delete(f, "ContainerStdout")
	delete(f, "LogFormat")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "UserDefineRule")
	delete(f, "GroupId")
	delete(f, "GroupIds")
	delete(f, "CollectInfos")
	delete(f, "AdvancedConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigExtraResponseParams struct {
	// 采集配置扩展信息ID
	ConfigExtraId *string `json:"ConfigExtraId,omitnil,omitempty" name:"ConfigExtraId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConfigExtraResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigExtraResponseParams `json:"Response"`
}

func (r *CreateConfigExtraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigRequestParams struct {
	// 采集配置名称
	// - 名称种不得包含特殊字符｜
	// - 名称最长255字符，超过截断
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 采集配置所属日志主题ID即TopicId
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 日志采集路径，包含文件名，支持多个路径，多个路径之间英文逗号分隔，文件采集情况下必填
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 采集的日志类型，默认为minimalist_log。支持以下类型：
	// - json_log代表：JSON-文件日志（详见[使用 JSON 提取模式采集日志](https://cloud.tencent.com/document/product/614/17419)）；
	// - delimiter_log代表：分隔符-文件日志（详见[使用分隔符提取模式采集日志](https://cloud.tencent.com/document/product/614/17420)）；
	// - minimalist_log代表：单行全文-文件日志（详见[使用单行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17421)）；
	// - fullregex_log代表：单行完全正则-文件日志（详见[使用单行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52365)）；
	// - multiline_log代表：多行全文-文件日志（详见[使用多行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17422)）；
	// - multiline_fullregex_log代表：多行完全正则-文件日志（详见[使用多行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52366)）；
	// - user_define_log代表：组合解析（适用于多格式嵌套的日志，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)）；
	// - service_syslog代表：syslog 采集（详见[采集 Syslog](https://cloud.tencent.com/document/product/614/81454)）；
	// - windows_event_log代表：Windows事件日志（详见[采集 Windows 事件日志](https://cloud.tencent.com/document/product/614/96678)）。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 用户自定义采集规则，Json格式序列化的字符串。当LogType为user_define_log时，必填。
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// 样例：
	// `{\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true}`
	// 
	// 控制台默认占位值：`{\"ClsAgentDefault\":0}`
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`

	// 日志输入类型（<span style="color:red; font-weight:bold">注：windows场景必填且仅支持file和windows_event类型</span>）
	// - file: 文件类型采集
	// - windows_event：windows事件采集
	// - syslog：系统日志采集
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`
}

type CreateConfigRequest struct {
	*tchttp.BaseRequest
	
	// 采集配置名称
	// - 名称种不得包含特殊字符｜
	// - 名称最长255字符，超过截断
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 采集配置所属日志主题ID即TopicId
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 日志采集路径，包含文件名，支持多个路径，多个路径之间英文逗号分隔，文件采集情况下必填
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 采集的日志类型，默认为minimalist_log。支持以下类型：
	// - json_log代表：JSON-文件日志（详见[使用 JSON 提取模式采集日志](https://cloud.tencent.com/document/product/614/17419)）；
	// - delimiter_log代表：分隔符-文件日志（详见[使用分隔符提取模式采集日志](https://cloud.tencent.com/document/product/614/17420)）；
	// - minimalist_log代表：单行全文-文件日志（详见[使用单行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17421)）；
	// - fullregex_log代表：单行完全正则-文件日志（详见[使用单行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52365)）；
	// - multiline_log代表：多行全文-文件日志（详见[使用多行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17422)）；
	// - multiline_fullregex_log代表：多行完全正则-文件日志（详见[使用多行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52366)）；
	// - user_define_log代表：组合解析（适用于多格式嵌套的日志，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)）；
	// - service_syslog代表：syslog 采集（详见[采集 Syslog](https://cloud.tencent.com/document/product/614/81454)）；
	// - windows_event_log代表：Windows事件日志（详见[采集 Windows 事件日志](https://cloud.tencent.com/document/product/614/96678)）。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 用户自定义采集规则，Json格式序列化的字符串。当LogType为user_define_log时，必填。
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// 样例：
	// `{\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true}`
	// 
	// 控制台默认占位值：`{\"ClsAgentDefault\":0}`
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`

	// 日志输入类型（<span style="color:red; font-weight:bold">注：windows场景必填且仅支持file和windows_event类型</span>）
	// - file: 文件类型采集
	// - windows_event：windows事件采集
	// - syslog：系统日志采集
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`
}

func (r *CreateConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Output")
	delete(f, "Path")
	delete(f, "LogType")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "UserDefineRule")
	delete(f, "AdvancedConfig")
	delete(f, "InputType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigResponseParams struct {
	// 采集配置ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigResponseParams `json:"Response"`
}

func (r *CreateConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsoleSharingRequestParams struct {
	// 免密分享配置
	SharingConfig *ConsoleSharingConfig `json:"SharingConfig,omitnil,omitempty" name:"SharingConfig"`
}

type CreateConsoleSharingRequest struct {
	*tchttp.BaseRequest
	
	// 免密分享配置
	SharingConfig *ConsoleSharingConfig `json:"SharingConfig,omitnil,omitempty" name:"SharingConfig"`
}

func (r *CreateConsoleSharingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsoleSharingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SharingConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsoleSharingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsoleSharingResponseParams struct {
	// 免密分享链接
	SharingUrl *string `json:"SharingUrl,omitnil,omitempty" name:"SharingUrl"`

	// 免密分享链接ID
	SharingId *string `json:"SharingId,omitnil,omitempty" name:"SharingId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConsoleSharingResponse struct {
	*tchttp.BaseResponse
	Response *CreateConsoleSharingResponseParams `json:"Response"`
}

func (r *CreateConsoleSharingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsoleSharingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerRequestParams struct {
	// 投递任务绑定的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 是否投递日志的元数据信息，默认为 true。
	// 当NeedContent为true时：字段Content有效。
	// 当NeedContent为false时：字段Content无效。
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// 如果需要投递元数据信息，元数据信息的描述
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka的描述
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// 投递时压缩方式，取值0，2，3。[0：NONE；2：SNAPPY；3：LZ4]
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// 角色访问描述名 [创建角色](https://cloud.tencent.com/document/product/598/19381)
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 外部ID
	ExternalId *string `json:"ExternalId,omitnil,omitempty" name:"ExternalId"`

	// 高级配置项
	AdvancedConfig *AdvancedConsumerConfiguration `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type CreateConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 投递任务绑定的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 是否投递日志的元数据信息，默认为 true。
	// 当NeedContent为true时：字段Content有效。
	// 当NeedContent为false时：字段Content无效。
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// 如果需要投递元数据信息，元数据信息的描述
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka的描述
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// 投递时压缩方式，取值0，2，3。[0：NONE；2：SNAPPY；3：LZ4]
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// 角色访问描述名 [创建角色](https://cloud.tencent.com/document/product/598/19381)
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 外部ID
	ExternalId *string `json:"ExternalId,omitnil,omitempty" name:"ExternalId"`

	// 高级配置项
	AdvancedConfig *AdvancedConsumerConfiguration `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

func (r *CreateConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "NeedContent")
	delete(f, "Content")
	delete(f, "Ckafka")
	delete(f, "Compression")
	delete(f, "RoleArn")
	delete(f, "ExternalId")
	delete(f, "AdvancedConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConsumerResponse struct {
	*tchttp.BaseResponse
	Response *CreateConsumerResponseParams `json:"Response"`
}

func (r *CreateConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosRechargeRequestParams struct {
	// 日志主题Id。
	// 
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志集Id。
	// 
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// COS导入任务名称,最大支持128个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// COS存储桶所在地域，详见产品支持的[地域列表](https://cloud.tencent.com/document/product/436/6224)。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表单行全文；
	// 默认为minimalist_log
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// COS文件所在文件夹的前缀。默认为空，投递存储桶下所有的文件。
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// supported: "", "gzip", "lzop", "snappy"。默认空，不压缩。
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitnil,omitempty" name:"ExtractRuleInfo"`

	// COS导入任务类型。1：一次性导入任务；2：持续性导入任务。默认为1：一次性导入任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 元数据。
	Metadata []*string `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type CreateCosRechargeRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// 
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志集Id。
	// 
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// COS导入任务名称,最大支持128个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// COS存储桶所在地域，详见产品支持的[地域列表](https://cloud.tencent.com/document/product/436/6224)。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表单行全文；
	// 默认为minimalist_log
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// COS文件所在文件夹的前缀。默认为空，投递存储桶下所有的文件。
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// supported: "", "gzip", "lzop", "snappy"。默认空，不压缩。
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitnil,omitempty" name:"ExtractRuleInfo"`

	// COS导入任务类型。1：一次性导入任务；2：持续性导入任务。默认为1：一次性导入任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 元数据。
	Metadata []*string `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

func (r *CreateCosRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "LogsetId")
	delete(f, "Name")
	delete(f, "Bucket")
	delete(f, "BucketRegion")
	delete(f, "LogType")
	delete(f, "Prefix")
	delete(f, "Compress")
	delete(f, "ExtractRuleInfo")
	delete(f, "TaskType")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosRechargeResponseParams struct {
	// COS导入任务id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCosRechargeResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosRechargeResponseParams `json:"Response"`
}

func (r *CreateCosRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDashboardSubscribeRequestParams struct {
	// 仪表盘订阅名称。
	// 输入限制：
	// - 不能为空
	// - 长度不能超过128字节
	// - 不能包含字符'|'
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 仪表盘Id。
	// - 通过[获取仪表盘](https://cloud.tencent.com/document/product/614/95636)获取仪表盘Id。
	DashboardId *string `json:"DashboardId,omitnil,omitempty" name:"DashboardId"`

	// 订阅时间cron表达式，格式为：{秒数} {分钟} {小时} {日期} {月份} {星期}；（有效数据为：{分钟} {小时} {日期} {月份} {星期}）。
	// - {秒数} 取值范围： 0 ~ 59 
	// - {分钟} 取值范围： 0 ~ 59 
	// - {小时} 取值范围： 0 ~ 23 
	// - {日期} 取值范围： 1 ~ 31 AND (dayOfMonth最后一天： L) 
	// - {月份} 取值范围： 1 ~ 12 
	// - {星期} 取值范围： 0 ~ 6 【0:星期日， 6星期六】
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 仪表盘订阅数据。
	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitnil,omitempty" name:"SubscribeData"`
}

type CreateDashboardSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 仪表盘订阅名称。
	// 输入限制：
	// - 不能为空
	// - 长度不能超过128字节
	// - 不能包含字符'|'
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 仪表盘Id。
	// - 通过[获取仪表盘](https://cloud.tencent.com/document/product/614/95636)获取仪表盘Id。
	DashboardId *string `json:"DashboardId,omitnil,omitempty" name:"DashboardId"`

	// 订阅时间cron表达式，格式为：{秒数} {分钟} {小时} {日期} {月份} {星期}；（有效数据为：{分钟} {小时} {日期} {月份} {星期}）。
	// - {秒数} 取值范围： 0 ~ 59 
	// - {分钟} 取值范围： 0 ~ 59 
	// - {小时} 取值范围： 0 ~ 23 
	// - {日期} 取值范围： 1 ~ 31 AND (dayOfMonth最后一天： L) 
	// - {月份} 取值范围： 1 ~ 12 
	// - {星期} 取值范围： 0 ~ 6 【0:星期日， 6星期六】
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 仪表盘订阅数据。
	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitnil,omitempty" name:"SubscribeData"`
}

func (r *CreateDashboardSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDashboardSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DashboardId")
	delete(f, "Cron")
	delete(f, "SubscribeData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDashboardSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDashboardSubscribeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDashboardSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *CreateDashboardSubscribeResponseParams `json:"Response"`
}

func (r *CreateDashboardSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDashboardSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataTransformRequestParams struct {
	// 任务类型. 1: 指定主题；2:动态创建。详情请参考[创建加工任务文档](https://cloud.tencent.com/document/product/614/63940)。
	FuncType *int64 `json:"FuncType,omitnil,omitempty" name:"FuncType"`

	// 日志主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// 加工任务名称
	// 名称限制
	// - 不能为空字符串
	// - 不能包含字符'|'
	// - 最长 255 个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 加工语句。 当FuncType为2时，EtlContent必须使用[log_auto_output](https://cloud.tencent.com/document/product/614/70733#b3c58797-4825-4807-bef4-68106e25024f) 
	// 
	// 其他参考文档：
	// 
	// - [创建加工任务](https://cloud.tencent.com/document/product/614/63940) 
	// -  [函数总览](https://cloud.tencent.com/document/product/614/70395)
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// 加工类型。
	// 1：使用源日志主题中的随机数据，进行加工预览；2：使用用户自定义测试数据，进行加工预览；3：创建真实加工任务。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 加工任务目标topic_id以及别名,当FuncType=1时，该参数必填，当FuncType=2时，无需填写。
	// 目标topic_id，通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	// 别名限制 1.不能为空字符串，2. 不能包含字符'|'。
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// 任务启动状态.   默认为1:开启,  2:关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 用于预览加工结果的测试数据
	// 目标日志主题ID通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	PreviewLogStatistics []*PreviewLogStatistic `json:"PreviewLogStatistics,omitnil,omitempty" name:"PreviewLogStatistics"`

	// 当FuncType为2时，动态创建的日志集、日志主题的个数超出产品规格限制是否丢弃数据， 默认为false。
	// 
	// false：创建兜底日志集、日志主题并将日志写入兜底主题；
	// true：丢弃日志数据。
	BackupGiveUpData *bool `json:"BackupGiveUpData,omitnil,omitempty" name:"BackupGiveUpData"`

	// 是否开启投递服务日志。1：关闭，2：开启。
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`

	// 数据加工类型。0：标准加工任务； 1：前置加工任务。前置加工任务将采集的日志处理完成后，再写入日志主题。
	DataTransformType *uint64 `json:"DataTransformType,omitnil,omitempty" name:"DataTransformType"`

	// 保留失败日志状态，1:不保留(默认)，2:保留。
	KeepFailureLog *uint64 `json:"KeepFailureLog,omitnil,omitempty" name:"KeepFailureLog"`

	// 失败日志的字段名称
	FailureLogKey *string `json:"FailureLogKey,omitnil,omitempty" name:"FailureLogKey"`

	// 指定加工数据的开始时间, 秒级时间戳。
	//  - 日志主题生命周期内的任意时间范围，如果超出了生命周期,只处理生命周期内有数据的部分。
	ProcessFromTimestamp *uint64 `json:"ProcessFromTimestamp,omitnil,omitempty" name:"ProcessFromTimestamp"`

	// 指定加工数据的结束时间，秒级时间戳。
	// 
	// -  不可指定未来的时间
	// -  不填则表示持续执行
	ProcessToTimestamp *uint64 `json:"ProcessToTimestamp,omitnil,omitempty" name:"ProcessToTimestamp"`

	// 对已经创建的并且使用了关联外部数据库能力的任务预览（TaskType 为 1 或 2）时，该值必传
	// 数据加工任务ID- 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 关联的数据源信息
	DataTransformSqlDataSources []*DataTransformSqlDataSource `json:"DataTransformSqlDataSources,omitnil,omitempty" name:"DataTransformSqlDataSources"`

	// 设置的环境变量
	EnvInfos []*EnvInfo `json:"EnvInfos,omitnil,omitempty" name:"EnvInfos"`
}

type CreateDataTransformRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型. 1: 指定主题；2:动态创建。详情请参考[创建加工任务文档](https://cloud.tencent.com/document/product/614/63940)。
	FuncType *int64 `json:"FuncType,omitnil,omitempty" name:"FuncType"`

	// 日志主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// 加工任务名称
	// 名称限制
	// - 不能为空字符串
	// - 不能包含字符'|'
	// - 最长 255 个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 加工语句。 当FuncType为2时，EtlContent必须使用[log_auto_output](https://cloud.tencent.com/document/product/614/70733#b3c58797-4825-4807-bef4-68106e25024f) 
	// 
	// 其他参考文档：
	// 
	// - [创建加工任务](https://cloud.tencent.com/document/product/614/63940) 
	// -  [函数总览](https://cloud.tencent.com/document/product/614/70395)
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// 加工类型。
	// 1：使用源日志主题中的随机数据，进行加工预览；2：使用用户自定义测试数据，进行加工预览；3：创建真实加工任务。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 加工任务目标topic_id以及别名,当FuncType=1时，该参数必填，当FuncType=2时，无需填写。
	// 目标topic_id，通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	// 别名限制 1.不能为空字符串，2. 不能包含字符'|'。
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// 任务启动状态.   默认为1:开启,  2:关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 用于预览加工结果的测试数据
	// 目标日志主题ID通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	PreviewLogStatistics []*PreviewLogStatistic `json:"PreviewLogStatistics,omitnil,omitempty" name:"PreviewLogStatistics"`

	// 当FuncType为2时，动态创建的日志集、日志主题的个数超出产品规格限制是否丢弃数据， 默认为false。
	// 
	// false：创建兜底日志集、日志主题并将日志写入兜底主题；
	// true：丢弃日志数据。
	BackupGiveUpData *bool `json:"BackupGiveUpData,omitnil,omitempty" name:"BackupGiveUpData"`

	// 是否开启投递服务日志。1：关闭，2：开启。
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`

	// 数据加工类型。0：标准加工任务； 1：前置加工任务。前置加工任务将采集的日志处理完成后，再写入日志主题。
	DataTransformType *uint64 `json:"DataTransformType,omitnil,omitempty" name:"DataTransformType"`

	// 保留失败日志状态，1:不保留(默认)，2:保留。
	KeepFailureLog *uint64 `json:"KeepFailureLog,omitnil,omitempty" name:"KeepFailureLog"`

	// 失败日志的字段名称
	FailureLogKey *string `json:"FailureLogKey,omitnil,omitempty" name:"FailureLogKey"`

	// 指定加工数据的开始时间, 秒级时间戳。
	//  - 日志主题生命周期内的任意时间范围，如果超出了生命周期,只处理生命周期内有数据的部分。
	ProcessFromTimestamp *uint64 `json:"ProcessFromTimestamp,omitnil,omitempty" name:"ProcessFromTimestamp"`

	// 指定加工数据的结束时间，秒级时间戳。
	// 
	// -  不可指定未来的时间
	// -  不填则表示持续执行
	ProcessToTimestamp *uint64 `json:"ProcessToTimestamp,omitnil,omitempty" name:"ProcessToTimestamp"`

	// 对已经创建的并且使用了关联外部数据库能力的任务预览（TaskType 为 1 或 2）时，该值必传
	// 数据加工任务ID- 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 关联的数据源信息
	DataTransformSqlDataSources []*DataTransformSqlDataSource `json:"DataTransformSqlDataSources,omitnil,omitempty" name:"DataTransformSqlDataSources"`

	// 设置的环境变量
	EnvInfos []*EnvInfo `json:"EnvInfos,omitnil,omitempty" name:"EnvInfos"`
}

func (r *CreateDataTransformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTransformRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FuncType")
	delete(f, "SrcTopicId")
	delete(f, "Name")
	delete(f, "EtlContent")
	delete(f, "TaskType")
	delete(f, "DstResources")
	delete(f, "EnableFlag")
	delete(f, "PreviewLogStatistics")
	delete(f, "BackupGiveUpData")
	delete(f, "HasServicesLog")
	delete(f, "DataTransformType")
	delete(f, "KeepFailureLog")
	delete(f, "FailureLogKey")
	delete(f, "ProcessFromTimestamp")
	delete(f, "ProcessToTimestamp")
	delete(f, "TaskId")
	delete(f, "DataTransformSqlDataSources")
	delete(f, "EnvInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataTransformRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataTransformResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataTransformResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataTransformResponseParams `json:"Response"`
}

func (r *CreateDataTransformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeliverCloudFunctionRequestParams struct {
	// 投递规则属于的TopicId。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 投递的云函数名字。仅支持[事件函数](https://cloud.tencent.com/document/product/583/9694#scf-.E4.BA.8B.E4.BB.B6.E5.87.BD.E6.95.B0) （[函数类型选型](https://cloud.tencent.com/document/product/583/73483)）
	// 通过 [获取函数列表](https://cloud.tencent.com/document/product/583/18582) 获取函数信息。
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// 命名空间。参考 [命名空间管理](https://cloud.tencent.com/document/product/583/35913)
	// - 通过 [列出命名空间列表](https://cloud.tencent.com/document/product/583/37158) 获取Name。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 函数版本。
	// - 通过 [查询函数版本 ](https://cloud.tencent.com/document/product/583/37162) 获取函数版本。
	Qualifier *string `json:"Qualifier,omitnil,omitempty" name:"Qualifier"`

	// 投递最长等待时间，单位：秒。 默认：60
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 投递最大消息数。默认为100。支持范围[1,10000]
	MaxMsgNum *uint64 `json:"MaxMsgNum,omitnil,omitempty" name:"MaxMsgNum"`
}

type CreateDeliverCloudFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 投递规则属于的TopicId。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 投递的云函数名字。仅支持[事件函数](https://cloud.tencent.com/document/product/583/9694#scf-.E4.BA.8B.E4.BB.B6.E5.87.BD.E6.95.B0) （[函数类型选型](https://cloud.tencent.com/document/product/583/73483)）
	// 通过 [获取函数列表](https://cloud.tencent.com/document/product/583/18582) 获取函数信息。
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// 命名空间。参考 [命名空间管理](https://cloud.tencent.com/document/product/583/35913)
	// - 通过 [列出命名空间列表](https://cloud.tencent.com/document/product/583/37158) 获取Name。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 函数版本。
	// - 通过 [查询函数版本 ](https://cloud.tencent.com/document/product/583/37162) 获取函数版本。
	Qualifier *string `json:"Qualifier,omitnil,omitempty" name:"Qualifier"`

	// 投递最长等待时间，单位：秒。 默认：60
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 投递最大消息数。默认为100。支持范围[1,10000]
	MaxMsgNum *uint64 `json:"MaxMsgNum,omitnil,omitempty" name:"MaxMsgNum"`
}

func (r *CreateDeliverCloudFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeliverCloudFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "FunctionName")
	delete(f, "Namespace")
	delete(f, "Qualifier")
	delete(f, "Timeout")
	delete(f, "MaxMsgNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeliverCloudFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeliverCloudFunctionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeliverCloudFunctionResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeliverCloudFunctionResponseParams `json:"Response"`
}

func (r *CreateDeliverCloudFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeliverCloudFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDlcDeliverRequestParams struct {
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z, _,-,中文字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 投递类型。0:批投递,1:实时投递
	DeliverType *uint64 `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 投递时间范围的开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// dlc配置信息
	DlcInfo *DlcInfo `json:"DlcInfo,omitnil,omitempty" name:"DlcInfo"`

	// 投递文件大小,单位MB。 DeliverType=0时必填，范围 5<= MaxSize <= 256。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 投递间隔，单位秒。 DeliverType=0时必填，范围 300<= Interval <=900。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 投递时间范围的结束时间。 如果为空，则表示不限时。EndTime不为空时，需要大于StartTime。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否开启投递服务日志。1关闭，2开启。默认开启
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`
}

type CreateDlcDeliverRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z, _,-,中文字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 投递类型。0:批投递,1:实时投递
	DeliverType *uint64 `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 投递时间范围的开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// dlc配置信息
	DlcInfo *DlcInfo `json:"DlcInfo,omitnil,omitempty" name:"DlcInfo"`

	// 投递文件大小,单位MB。 DeliverType=0时必填，范围 5<= MaxSize <= 256。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 投递间隔，单位秒。 DeliverType=0时必填，范围 300<= Interval <=900。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 投递时间范围的结束时间。 如果为空，则表示不限时。EndTime不为空时，需要大于StartTime。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否开启投递服务日志。1关闭，2开启。默认开启
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`
}

func (r *CreateDlcDeliverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDlcDeliverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Name")
	delete(f, "DeliverType")
	delete(f, "StartTime")
	delete(f, "DlcInfo")
	delete(f, "MaxSize")
	delete(f, "Interval")
	delete(f, "EndTime")
	delete(f, "HasServicesLog")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDlcDeliverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDlcDeliverResponseParams struct {
	// 配置id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDlcDeliverResponse struct {
	*tchttp.BaseResponse
	Response *CreateDlcDeliverResponseParams `json:"Response"`
}

func (r *CreateDlcDeliverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDlcDeliverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportRequestParams struct {
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志导出数量,  最大值5000万
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 日志导出检索语句，不支持<a href="https://cloud.tencent.com/document/product/614/44061" target="_blank">[SQL语句]</a>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 日志导出起始时间，毫秒时间戳
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 日志导出结束时间，毫秒时间戳
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 语法规则,  默认值为0。
	// 0：Lucene语法，1：CQL语法。
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// 导出字段
	DerivedFields []*string `json:"DerivedFields,omitnil,omitempty" name:"DerivedFields"`
}

type CreateExportRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志导出数量,  最大值5000万
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 日志导出检索语句，不支持<a href="https://cloud.tencent.com/document/product/614/44061" target="_blank">[SQL语句]</a>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 日志导出起始时间，毫秒时间戳
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 日志导出结束时间，毫秒时间戳
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 语法规则,  默认值为0。
	// 0：Lucene语法，1：CQL语法。
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// 导出字段
	DerivedFields []*string `json:"DerivedFields,omitnil,omitempty" name:"DerivedFields"`
}

func (r *CreateExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Count")
	delete(f, "Query")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Order")
	delete(f, "Format")
	delete(f, "SyntaxRule")
	delete(f, "DerivedFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportResponseParams struct {
	// 日志导出ID。
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExportResponse struct {
	*tchttp.BaseResponse
	Response *CreateExportResponseParams `json:"Response"`
}

func (r *CreateExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIndexRequestParams struct {
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 索引规则
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 是否生效，默认为true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 内置保留字段（`__FILENAME__`，`__HOSTNAME__`及`__SOURCE__`）是否包含至全文索引，默认为false，推荐设置为true
	// * false:不包含
	// * true:包含
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// 元数据字段（前缀为`__TAG__`的字段）是否包含至全文索引，默认为0，推荐设置为1
	// * 0:仅包含开启键值索引的元数据字段
	// * 1:包含所有元数据字段
	// * 2:不包含任何元数据字段
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`

	// 自定义日志解析异常存储字段。
	CoverageField *string `json:"CoverageField,omitnil,omitempty" name:"CoverageField"`
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 索引规则
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 是否生效，默认为true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 内置保留字段（`__FILENAME__`，`__HOSTNAME__`及`__SOURCE__`）是否包含至全文索引，默认为false，推荐设置为true
	// * false:不包含
	// * true:包含
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// 元数据字段（前缀为`__TAG__`的字段）是否包含至全文索引，默认为0，推荐设置为1
	// * 0:仅包含开启键值索引的元数据字段
	// * 1:包含所有元数据字段
	// * 2:不包含任何元数据字段
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`

	// 自定义日志解析异常存储字段。
	CoverageField *string `json:"CoverageField,omitnil,omitempty" name:"CoverageField"`
}

func (r *CreateIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Rule")
	delete(f, "Status")
	delete(f, "IncludeInternalFields")
	delete(f, "MetadataFlag")
	delete(f, "CoverageField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIndexResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIndexResponse struct {
	*tchttp.BaseResponse
	Response *CreateIndexResponseParams `json:"Response"`
}

func (r *CreateIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKafkaRechargeRequestParams struct {
	// 导入CLS目标TopicId。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka导入配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 导入Kafka类型，0: 腾讯云CKafka，1: 用户自建Kafka
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// 导入数据位置，-2:最早（默认），-1：最晚
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 日志导入规则。
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`

	// 腾讯云CKafka实例ID，KafkaType为0时必填。
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取实例id。
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// 服务地址，KafkaType为1时必填。
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// ServerAddr是否为加密连接，KafkaType为1时必填。
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// 加密访问协议。
	// KafkaType为1并且IsEncryptionAddr为true时Protocol必填。
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 用户Kafka消费组名称。
	// - 消费组是 Kafka 提供的可扩展且具有容错性的消费者机制，一个消费组中存在多个消费者，组内的所有消费者共同消费订阅 Topic 中的消息。一个消费者可同时消费多个 Partition，但一个 Partition 只能被消费组内的一个消费者消费。
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 用户kafka拓展信息
	UserKafkaMeta *UserKafkaMeta `json:"UserKafkaMeta,omitnil,omitempty" name:"UserKafkaMeta"`
}

type CreateKafkaRechargeRequest struct {
	*tchttp.BaseRequest
	
	// 导入CLS目标TopicId。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka导入配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 导入Kafka类型，0: 腾讯云CKafka，1: 用户自建Kafka
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// 导入数据位置，-2:最早（默认），-1：最晚
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 日志导入规则。
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`

	// 腾讯云CKafka实例ID，KafkaType为0时必填。
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取实例id。
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// 服务地址，KafkaType为1时必填。
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// ServerAddr是否为加密连接，KafkaType为1时必填。
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// 加密访问协议。
	// KafkaType为1并且IsEncryptionAddr为true时Protocol必填。
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 用户Kafka消费组名称。
	// - 消费组是 Kafka 提供的可扩展且具有容错性的消费者机制，一个消费组中存在多个消费者，组内的所有消费者共同消费订阅 Topic 中的消息。一个消费者可同时消费多个 Partition，但一个 Partition 只能被消费组内的一个消费者消费。
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 用户kafka拓展信息
	UserKafkaMeta *UserKafkaMeta `json:"UserKafkaMeta,omitnil,omitempty" name:"UserKafkaMeta"`
}

func (r *CreateKafkaRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKafkaRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Name")
	delete(f, "KafkaType")
	delete(f, "UserKafkaTopics")
	delete(f, "Offset")
	delete(f, "LogRechargeRule")
	delete(f, "KafkaInstance")
	delete(f, "ServerAddr")
	delete(f, "IsEncryptionAddr")
	delete(f, "Protocol")
	delete(f, "ConsumerGroupName")
	delete(f, "UserKafkaMeta")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKafkaRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKafkaRechargeResponseParams struct {
	// Kafka导入配置ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKafkaRechargeResponse struct {
	*tchttp.BaseResponse
	Response *CreateKafkaRechargeResponseParams `json:"Response"`
}

func (r *CreateKafkaRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogsetRequestParams struct {
	// 日志集名字。
	// 
	// - 最大支持255个字符。不支持`|`字符。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 标签描述列表。最大支持10个标签键值对，并且不能有重复的键值对
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 日志集ID，格式为：用户自定义部分-用户APPID。未填写该参数时将自动生成ID。
	// 
	// - 用户自定义部分仅支持小写字母、数字和-，且不能以-开头和结尾，长度为3至40字符。
	// - 尾部需要使用-拼接用户APPID，APPID可在https://console.cloud.tencent.com/developer页面查询。
	// - 如果指定该字段，需保证全地域唯一
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

type CreateLogsetRequest struct {
	*tchttp.BaseRequest
	
	// 日志集名字。
	// 
	// - 最大支持255个字符。不支持`|`字符。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 标签描述列表。最大支持10个标签键值对，并且不能有重复的键值对
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 日志集ID，格式为：用户自定义部分-用户APPID。未填写该参数时将自动生成ID。
	// 
	// - 用户自定义部分仅支持小写字母、数字和-，且不能以-开头和结尾，长度为3至40字符。
	// - 尾部需要使用-拼接用户APPID，APPID可在https://console.cloud.tencent.com/developer页面查询。
	// - 如果指定该字段，需保证全地域唯一
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

func (r *CreateLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetName")
	delete(f, "Tags")
	delete(f, "LogsetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogsetResponseParams struct {
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLogsetResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogsetResponseParams `json:"Response"`
}

func (r *CreateLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMachineGroupRequestParams struct {
	// 机器组名字。
	// 输入限制：
	// - 最大支持255个字符，不能为空字符串
	// - 不能包含字符'|'
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 创建机器组类型。取值如下：
	// - Type：ip，Values中为ip字符串列表创建机器组
	// - Type：label，Values中为标签字符串列表创建机器组
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的机器组。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启机器组自动更新。默认false
	AutoUpdate *bool `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费。默认false
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// 机器组中机器离线清理时间。单位：天
	// 
	// - 大于0时生效。
	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitnil,omitempty" name:"DelayCleanupTime"`

	// 机器组元数据信息列表
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`

	// 系统类型，取值如下：
	// - 0：Linux （默认值）
	// - 1：Windows
	OSType *uint64 `json:"OSType,omitnil,omitempty" name:"OSType"`
}

type CreateMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// 机器组名字。
	// 输入限制：
	// - 最大支持255个字符，不能为空字符串
	// - 不能包含字符'|'
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 创建机器组类型。取值如下：
	// - Type：ip，Values中为ip字符串列表创建机器组
	// - Type：label，Values中为标签字符串列表创建机器组
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的机器组。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启机器组自动更新。默认false
	AutoUpdate *bool `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费。默认false
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// 机器组中机器离线清理时间。单位：天
	// 
	// - 大于0时生效。
	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitnil,omitempty" name:"DelayCleanupTime"`

	// 机器组元数据信息列表
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`

	// 系统类型，取值如下：
	// - 0：Linux （默认值）
	// - 1：Windows
	OSType *uint64 `json:"OSType,omitnil,omitempty" name:"OSType"`
}

func (r *CreateMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	delete(f, "DelayCleanupTime")
	delete(f, "MetaTags")
	delete(f, "OSType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMachineGroupResponseParams struct {
	// 机器组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateMachineGroupResponseParams `json:"Response"`
}

func (r *CreateMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNoticeContentRequestParams struct {
	// 模板名称。最大支持255个字节
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板内容语言。0：中文1：英文
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 模板详细配置。
	NoticeContents []*NoticeContent `json:"NoticeContents,omitnil,omitempty" name:"NoticeContents"`
}

type CreateNoticeContentRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称。最大支持255个字节
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板内容语言。0：中文1：英文
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 模板详细配置。
	NoticeContents []*NoticeContent `json:"NoticeContents,omitnil,omitempty" name:"NoticeContents"`
}

func (r *CreateNoticeContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNoticeContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeContents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNoticeContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNoticeContentResponseParams struct {
	// 通知内容配置ID
	NoticeContentId *string `json:"NoticeContentId,omitnil,omitempty" name:"NoticeContentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNoticeContentResponse struct {
	*tchttp.BaseResponse
	Response *CreateNoticeContentResponseParams `json:"Response"`
}

func (r *CreateNoticeContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNoticeContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScheduledSqlRequestParams struct {
	// 源日志主题ID- 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// 任务名称，0~255字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务启动状态.  1开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 定时SQL分析目标日志主题
	DstResource *ScheduledSqlResouceInfo `json:"DstResource,omitnil,omitempty" name:"DstResource"`

	// 查询语句
	ScheduledSqlContent *string `json:"ScheduledSqlContent,omitnil,omitempty" name:"ScheduledSqlContent"`

	// 调度开始时间,Unix时间戳，单位ms
	ProcessStartTime *uint64 `json:"ProcessStartTime,omitnil,omitempty" name:"ProcessStartTime"`

	// 调度类型，1:持续运行 2:指定时间范围
	ProcessType *int64 `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// 调度周期(分钟)，1~1440分钟
	ProcessPeriod *int64 `json:"ProcessPeriod,omitnil,omitempty" name:"ProcessPeriod"`

	// 单次查询的时间窗口,如果您的目标主题为指标主题，建议该参数的大小不超过30分钟，否则可能转指标失败。 
	ProcessTimeWindow *string `json:"ProcessTimeWindow,omitnil,omitempty" name:"ProcessTimeWindow"`

	// 执行延迟(秒)，0~120秒，默认60秒
	ProcessDelay *int64 `json:"ProcessDelay,omitnil,omitempty" name:"ProcessDelay"`

	// 源topicId的地域信息,支持地域见 [地域列表](https://cloud.tencent.com/document/api/614/56474#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8) 文档
	SrcTopicRegion *string `json:"SrcTopicRegion,omitnil,omitempty" name:"SrcTopicRegion"`

	// 调度结束时间，当ProcessType=2时为必传字段, Unix时间戳，单位ms
	ProcessEndTime *uint64 `json:"ProcessEndTime,omitnil,omitempty" name:"ProcessEndTime"`

	// 查询语法规则。 默认值为0。0：Lucene语法，1：CQL语法  
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

type CreateScheduledSqlRequest struct {
	*tchttp.BaseRequest
	
	// 源日志主题ID- 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// 任务名称，0~255字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务启动状态.  1开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 定时SQL分析目标日志主题
	DstResource *ScheduledSqlResouceInfo `json:"DstResource,omitnil,omitempty" name:"DstResource"`

	// 查询语句
	ScheduledSqlContent *string `json:"ScheduledSqlContent,omitnil,omitempty" name:"ScheduledSqlContent"`

	// 调度开始时间,Unix时间戳，单位ms
	ProcessStartTime *uint64 `json:"ProcessStartTime,omitnil,omitempty" name:"ProcessStartTime"`

	// 调度类型，1:持续运行 2:指定时间范围
	ProcessType *int64 `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// 调度周期(分钟)，1~1440分钟
	ProcessPeriod *int64 `json:"ProcessPeriod,omitnil,omitempty" name:"ProcessPeriod"`

	// 单次查询的时间窗口,如果您的目标主题为指标主题，建议该参数的大小不超过30分钟，否则可能转指标失败。 
	ProcessTimeWindow *string `json:"ProcessTimeWindow,omitnil,omitempty" name:"ProcessTimeWindow"`

	// 执行延迟(秒)，0~120秒，默认60秒
	ProcessDelay *int64 `json:"ProcessDelay,omitnil,omitempty" name:"ProcessDelay"`

	// 源topicId的地域信息,支持地域见 [地域列表](https://cloud.tencent.com/document/api/614/56474#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8) 文档
	SrcTopicRegion *string `json:"SrcTopicRegion,omitnil,omitempty" name:"SrcTopicRegion"`

	// 调度结束时间，当ProcessType=2时为必传字段, Unix时间戳，单位ms
	ProcessEndTime *uint64 `json:"ProcessEndTime,omitnil,omitempty" name:"ProcessEndTime"`

	// 查询语法规则。 默认值为0。0：Lucene语法，1：CQL语法  
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

func (r *CreateScheduledSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduledSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcTopicId")
	delete(f, "Name")
	delete(f, "EnableFlag")
	delete(f, "DstResource")
	delete(f, "ScheduledSqlContent")
	delete(f, "ProcessStartTime")
	delete(f, "ProcessType")
	delete(f, "ProcessPeriod")
	delete(f, "ProcessTimeWindow")
	delete(f, "ProcessDelay")
	delete(f, "SrcTopicRegion")
	delete(f, "ProcessEndTime")
	delete(f, "SyntaxRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScheduledSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScheduledSqlResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateScheduledSqlResponse struct {
	*tchttp.BaseResponse
	Response *CreateScheduledSqlResponseParams `json:"Response"`
}

func (r *CreateScheduledSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduledSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShipperRequestParams struct {
	// 创建的投递规则所属的日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 投递规则投递的新的目录前缀。
	// - 仅支持0-9A-Za-z-_/
	// - 最大支持256个字符
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 投递规则的名字。最大支持255个字符
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// 投递的时间间隔，单位 秒，默认300，范围 300-900
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 投递的文件的最大值，单位 MB，默认256，范围 5-256
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// 投递日志的分区规则，支持strftime的时间格式表示
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 投递日志的压缩配置
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 投递日志的内容格式配置
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// 投递文件命名配置，0：随机数命名，1：投递时间命名，默认0（随机数命名）
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`

	// 投递数据范围的开始时间点（秒级时间戳），不能超出日志主题的生命周期起点。
	// 如果用户不填写，默认为用户新建投递任务的时间。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 投递数据范围的结束时间点（秒级时间戳），不能填写未来时间。
	// 如果用户不填写，默认为持续投递，即无限。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 对象存储类型，默认值为 STANDARD。枚举值请参见[ 存储类型概述](https://cloud.tencent.com/document/product/436/33417) 文档。
	// 参考值有：
	// 
	// - STANDARD：标准存储
	// - STANDARD_IA：低频存储
	// - ARCHIVE：归档存储
	// - DEEP_ARCHIVE：深度归档存储
	// - MAZ_STANDARD：标准存储（多 AZ）
	// - MAZ_STANDARD_IA：低频存储（多 AZ）
	// - INTELLIGENT_TIERING：智能分层存储
	// - MAZ_INTELLIGENT_TIERING：智能分层存储（多 AZ）
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 角色访问描述名 [创建角色](https://cloud.tencent.com/document/product/598/19381)
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 外部ID
	ExternalId *string `json:"ExternalId,omitnil,omitempty" name:"ExternalId"`
}

type CreateShipperRequest struct {
	*tchttp.BaseRequest
	
	// 创建的投递规则所属的日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 投递规则投递的新的目录前缀。
	// - 仅支持0-9A-Za-z-_/
	// - 最大支持256个字符
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 投递规则的名字。最大支持255个字符
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// 投递的时间间隔，单位 秒，默认300，范围 300-900
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 投递的文件的最大值，单位 MB，默认256，范围 5-256
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// 投递日志的分区规则，支持strftime的时间格式表示
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 投递日志的压缩配置
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 投递日志的内容格式配置
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// 投递文件命名配置，0：随机数命名，1：投递时间命名，默认0（随机数命名）
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`

	// 投递数据范围的开始时间点（秒级时间戳），不能超出日志主题的生命周期起点。
	// 如果用户不填写，默认为用户新建投递任务的时间。
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 投递数据范围的结束时间点（秒级时间戳），不能填写未来时间。
	// 如果用户不填写，默认为持续投递，即无限。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 对象存储类型，默认值为 STANDARD。枚举值请参见[ 存储类型概述](https://cloud.tencent.com/document/product/436/33417) 文档。
	// 参考值有：
	// 
	// - STANDARD：标准存储
	// - STANDARD_IA：低频存储
	// - ARCHIVE：归档存储
	// - DEEP_ARCHIVE：深度归档存储
	// - MAZ_STANDARD：标准存储（多 AZ）
	// - MAZ_STANDARD_IA：低频存储（多 AZ）
	// - INTELLIGENT_TIERING：智能分层存储
	// - MAZ_INTELLIGENT_TIERING：智能分层存储（多 AZ）
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 角色访问描述名 [创建角色](https://cloud.tencent.com/document/product/598/19381)
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 外部ID
	ExternalId *string `json:"ExternalId,omitnil,omitempty" name:"ExternalId"`
}

func (r *CreateShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Bucket")
	delete(f, "Prefix")
	delete(f, "ShipperName")
	delete(f, "Interval")
	delete(f, "MaxSize")
	delete(f, "FilterRules")
	delete(f, "Partition")
	delete(f, "Compress")
	delete(f, "Content")
	delete(f, "FilenameMode")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "StorageType")
	delete(f, "RoleArn")
	delete(f, "ExternalId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateShipperResponseParams struct {
	// 投递任务ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateShipperResponse struct {
	*tchttp.BaseResponse
	Response *CreateShipperResponseParams `json:"Response"`
}

func (r *CreateShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRequestParams struct {
	// 日志集ID
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 主题名称
	// 名称限制
	// - 不能为空字符串
	// - 不能包含字符'|'
	// - 不能使用以下名称["cls_service_log","loglistener_status","loglistener_alarm","loglistener_business","cls_service_metric"]
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题分区个数。默认创建1个，最大支持创建10个分区。
	PartitionCount *int64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启自动分裂，默认值为true
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// 开启自动分裂后，每个主题能够允许的最大分区数，默认值为50
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// 日志主题的存储类型，可选值 hot（标准存储），cold（低频存储）；默认为hot。指标主题不支持该配置。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 存储时间，单位天。
	// - 日志主题：日志接入标准存储时，支持1至3600天，值为3640时代表永久保存。
	// - 日志主题：日志接入低频存储时，支持7至3600天，值为3640时代表永久保存。
	// - 指标主题：支持1至3600天，值为3640时代表永久保存。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 主题描述
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// 0：日志主题关闭日志沉降。
	// 非0：日志主题开启日志沉降后标准存储的天数，HotPeriod需要大于等于7，且小于Period。
	// 仅在StorageType为 hot 时生效，指标主题不支持该配置。
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`

	// 加密相关参数。 支持加密地域并且开白用户可以传此参数，其他场景不能传递该参数。
	// 0或者不传： 不加密
	// 1：kms-cls 云产品密钥加密
	// 
	// 支持地域：ap-beijing,ap-guangzhou,ap-shanghai,ap-singapore,ap-bangkok,ap-jakarta,eu-frankfurt,ap-seoul,ap-tokyo
	Encryption *uint64 `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 主题类型
	// - 0:日志主题，默认值
	// - 1:指标主题
	BizType *uint64 `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 主题自定义ID，格式为：用户自定义部分-用户APPID。未填写该参数时将自动生成ID。
	// - 用户自定义部分仅支持小写字母、数字和-，且不能以-开头和结尾，长度为3至40字符
	// - 尾部需要使用-拼接用户APPID，APPID可在https://console.cloud.tencent.com/developer页面查询。
	// - 如果指定该字段，需保证全地域唯一
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 免鉴权开关。 false：关闭； true：开启。默认为false。
	// 开启后将支持指定操作匿名访问该日志主题。详情请参见[日志主题](https://cloud.tencent.com/document/product/614/41035)。指标主题不支持该配置。
	IsWebTracking *bool `json:"IsWebTracking,omitnil,omitempty" name:"IsWebTracking"`

	// 主题扩展信息
	Extends *TopicExtendInfo `json:"Extends,omitnil,omitempty" name:"Extends"`

	// 开启记录公网来源ip和服务端接收时间
	IsSourceFrom *bool `json:"IsSourceFrom,omitnil,omitempty" name:"IsSourceFrom"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// 日志集ID
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 主题名称
	// 名称限制
	// - 不能为空字符串
	// - 不能包含字符'|'
	// - 不能使用以下名称["cls_service_log","loglistener_status","loglistener_alarm","loglistener_business","cls_service_metric"]
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题分区个数。默认创建1个，最大支持创建10个分区。
	PartitionCount *int64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启自动分裂，默认值为true
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// 开启自动分裂后，每个主题能够允许的最大分区数，默认值为50
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// 日志主题的存储类型，可选值 hot（标准存储），cold（低频存储）；默认为hot。指标主题不支持该配置。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 存储时间，单位天。
	// - 日志主题：日志接入标准存储时，支持1至3600天，值为3640时代表永久保存。
	// - 日志主题：日志接入低频存储时，支持7至3600天，值为3640时代表永久保存。
	// - 指标主题：支持1至3600天，值为3640时代表永久保存。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 主题描述
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// 0：日志主题关闭日志沉降。
	// 非0：日志主题开启日志沉降后标准存储的天数，HotPeriod需要大于等于7，且小于Period。
	// 仅在StorageType为 hot 时生效，指标主题不支持该配置。
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`

	// 加密相关参数。 支持加密地域并且开白用户可以传此参数，其他场景不能传递该参数。
	// 0或者不传： 不加密
	// 1：kms-cls 云产品密钥加密
	// 
	// 支持地域：ap-beijing,ap-guangzhou,ap-shanghai,ap-singapore,ap-bangkok,ap-jakarta,eu-frankfurt,ap-seoul,ap-tokyo
	Encryption *uint64 `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 主题类型
	// - 0:日志主题，默认值
	// - 1:指标主题
	BizType *uint64 `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 主题自定义ID，格式为：用户自定义部分-用户APPID。未填写该参数时将自动生成ID。
	// - 用户自定义部分仅支持小写字母、数字和-，且不能以-开头和结尾，长度为3至40字符
	// - 尾部需要使用-拼接用户APPID，APPID可在https://console.cloud.tencent.com/developer页面查询。
	// - 如果指定该字段，需保证全地域唯一
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 免鉴权开关。 false：关闭； true：开启。默认为false。
	// 开启后将支持指定操作匿名访问该日志主题。详情请参见[日志主题](https://cloud.tencent.com/document/product/614/41035)。指标主题不支持该配置。
	IsWebTracking *bool `json:"IsWebTracking,omitnil,omitempty" name:"IsWebTracking"`

	// 主题扩展信息
	Extends *TopicExtendInfo `json:"Extends,omitnil,omitempty" name:"Extends"`

	// 开启记录公网来源ip和服务端接收时间
	IsSourceFrom *bool `json:"IsSourceFrom,omitnil,omitempty" name:"IsSourceFrom"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicName")
	delete(f, "PartitionCount")
	delete(f, "Tags")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	delete(f, "StorageType")
	delete(f, "Period")
	delete(f, "Describes")
	delete(f, "HotPeriod")
	delete(f, "Encryption")
	delete(f, "BizType")
	delete(f, "TopicId")
	delete(f, "IsWebTracking")
	delete(f, "Extends")
	delete(f, "IsSourceFrom")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// 主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicResponseParams `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebCallbackRequestParams struct {
	// 通知内容名称。最大支持255个字节
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 渠道类型。
	// 
	// WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Webhook地址。
	Webhook *string `json:"Webhook,omitnil,omitempty" name:"Webhook"`

	// 请求方式。 支持POST、PUT。
	// 
	// 当Type为Http时，必填。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 秘钥。最大支持1024个字节
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

type CreateWebCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 通知内容名称。最大支持255个字节
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 渠道类型。
	// 
	// WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Webhook地址。
	Webhook *string `json:"Webhook,omitnil,omitempty" name:"Webhook"`

	// 请求方式。 支持POST、PUT。
	// 
	// 当Type为Http时，必填。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 秘钥。最大支持1024个字节
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

func (r *CreateWebCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Webhook")
	delete(f, "Method")
	delete(f, "Key")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWebCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebCallbackResponseParams struct {
	// 回调配置ID。
	WebCallbackId *string `json:"WebCallbackId,omitnil,omitempty" name:"WebCallbackId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWebCallbackResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebCallbackResponseParams `json:"Response"`
}

func (r *CreateWebCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CsvInfo struct {
	// csv首行是否打印key
	PrintKey *bool `json:"PrintKey,omitnil,omitempty" name:"PrintKey"`

	// 每列key的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 各字段间的分隔符
	Delimiter *string `json:"Delimiter,omitnil,omitempty" name:"Delimiter"`

	// 若字段内容中包含分隔符，则使用该转义符包裹改字段，只能填写单引号、双引号、空字符串
	EscapeChar *string `json:"EscapeChar,omitnil,omitempty" name:"EscapeChar"`

	// 对于上面指定的不存在字段使用该内容填充
	NonExistingField *string `json:"NonExistingField,omitnil,omitempty" name:"NonExistingField"`
}

type DashboardInfo struct {
	// 仪表盘id
	DashboardId *string `json:"DashboardId,omitnil,omitempty" name:"DashboardId"`

	// 仪表盘名字
	DashboardName *string `json:"DashboardName,omitnil,omitempty" name:"DashboardName"`

	// 仪表盘数据
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 创建仪表盘的时间。格式：YYYY-MM-DD HH:MM:SS
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// AssumerUin非空则表示创建该日志主题的服务方Uin
	AssumerUin *uint64 `json:"AssumerUin,omitnil,omitempty" name:"AssumerUin"`

	// RoleName非空则表示创建该日志主题的服务方使用的角色
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// AssumerName非空则表示创建该日志主题的服务方名称
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// 日志主题绑定的标签信息
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 仪表盘所在地域： 为了兼容老的地域。
	DashboardRegion *string `json:"DashboardRegion,omitnil,omitempty" name:"DashboardRegion"`

	// 修改仪表盘的时间。格式：YYYY-MM-DD HH:MM:SS
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 仪表盘对应的topic相关信息
	DashboardTopicInfos []*DashboardTopicInfo `json:"DashboardTopicInfos,omitnil,omitempty" name:"DashboardTopicInfos"`
}

type DashboardNoticeMode struct {
	// 仪表盘通知方式。
	// 
	// - Uin：腾讯云用户
	// - Group：腾讯云用户组
	// - WeCom：企业微信回调
	// - Email：自定义邮件
	// - DingTalk：钉钉
	// - Lark：飞书
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// 知方式对应的值。
	// - 当ReceiverType为：`WeCom`、`DingTalk`、`Lark` 时，Values必须为空，且Url字段必填。
	// - 当ReceiverType为：`Uin`、`Group`、`Email` 时，Values必填，且Url字段必须为空。
	// - 当ReceiverType为：`Uin ` 时，Values为用户id，通过 [拉取子用户](https://cloud.tencent.com/document/product/598/34587) 获取子用户 UID 。
	// - 当ReceiverType为：`Group` 时，Values为用户组id，通过 [查询用户组列表](https://cloud.tencent.com/document/product/598/34589) 获取用户组 ID 。
	// - 当ReceiverType为：`Email` 时，Values为用户邮箱信息。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 仪表盘通知渠道。
	// 
	// -  支持：["Email","Sms","WeChat","Phone"]。
	// -  当ReceiverType为 `Email` 或 `WeCom` 时，ReceiverChannels无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverChannels []*string `json:"ReceiverChannels,omitnil,omitempty" name:"ReceiverChannels"`

	// 订阅方式	- 回调地址。
	// - 当ReceiverType为：`WeCom`、`DingTalk`、`Lark` 时，Url字段必填为各渠道的回调地址。
	//     - 为：`WeCom` 时，Url为 企业微信回调地址。
	//     - 为：`DingTalk` 时，Url为 钉钉机器人Webhook地址。
	//     - 为：`Lark` 时，Url为 飞书机器人Webhook地址。
	// - 当ReceiverType为：`Uin`、`Group`、`Email` 时，Url字段必须为空。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DashboardSubscribeData struct {
	// 仪表盘订阅通知方式。
	NoticeModes []*DashboardNoticeMode `json:"NoticeModes,omitnil,omitempty" name:"NoticeModes"`

	// 仪表盘订阅时间，为空标识取仪表盘默认的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DashboardTime []*string `json:"DashboardTime,omitnil,omitempty" name:"DashboardTime"`

	// 仪表盘订阅模板变量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateVariables []*DashboardTemplateVariable `json:"TemplateVariables,omitnil,omitempty" name:"TemplateVariables"`

	// 时区。参考：https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#SHANGHAI
	Timezone *string `json:"Timezone,omitnil,omitempty" name:"Timezone"`

	// 语言。 zh 中文、en`英文。
	SubscribeLanguage *string `json:"SubscribeLanguage,omitnil,omitempty" name:"SubscribeLanguage"`

	// 调用链接域名。http:// 或者 https:// 开头，不能/结尾
	JumpDomain *string `json:"JumpDomain,omitnil,omitempty" name:"JumpDomain"`

	// 自定义跳转链接。
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`
}

type DashboardSubscribeInfo struct {
	// 仪表盘订阅id。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 仪表盘订阅名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 仪表盘id。
	DashboardId *string `json:"DashboardId,omitnil,omitempty" name:"DashboardId"`

	// 仪表盘订阅时间。
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 仪表盘订阅数据。
	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitnil,omitempty" name:"SubscribeData"`

	// 仪表盘订阅记录创建时间。格式：`YYYY-MM-DD HH:MM:SS`
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 仪表盘订阅记录更新时间。格式：`YYYY-MM-DD HH:MM:SS`
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 仪表盘订阅记录最后一次发送成功时间。格式：`YYYY-MM-DD HH:MM:SS`
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// 腾讯云主账号Id。
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 腾讯云主账号下的子账号Id。
	SubUin *uint64 `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 仪表盘订阅记录最后一次发送的状态。success：全部发送成功，fail：未发送， partialSuccess：部分发送成功。
	LastStatus *string `json:"LastStatus,omitnil,omitempty" name:"LastStatus"`
}

type DashboardTemplateVariable struct {
	// key的值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// key对应的values取值values
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type DashboardTopicInfo struct {
	// 主题id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// topic所在的地域。
	// - 1:广州
	// - 4:上海
	// - 5:中国香港
	// - 7:上海金融
	// - 8:北京
	// - 9:新加坡
	// - 11:深圳金融
	// - 15:硅谷
	// - 16:成都
	// - 17:法兰克福
	// - 18:首尔
	// - 19:重庆
	// - 22:弗吉尼亚
	// - 23:曼谷
	// - 25:东京
	// - 33:南京
	// - 36:天津
	// - 39:台北
	// - 46:北京金融
	// - 72:雅加达
	// - 74:圣保罗
	// - 78:上海自动驾驶云
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type DataTransformResouceInfo struct {
	// 日志主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 别名
	// 限制：不能包含字符 |。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type DataTransformSqlDataSource struct {
	// 数据源类型 1:MySql;2:自建mysql;3:pgsql
	DataSource *uint64 `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// InstanceId所属地域。例如：ap-guangzhou
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例Id。
	// - 当DataSource为1时，表示云数据库Mysql 实例id，如：cdb-zxcvbnm
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// mysql访问用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 别名。数据加工语句中使用
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// mysql访问密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type DataTransformTaskInfo struct {
	// 数据加工任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据加工任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务启用状态，默认为1，正常开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 加工任务类型，1： DSL(使用自定义加工语言的加工任务)， 2：SQL(使用sql的加工任务)
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 源日志主题
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// 当前加工任务状态（1准备中/2运行中/3停止中/4已停止）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 时间格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近修改时间
	// 时间格式：yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 最后启用时间，如果需要重建集群，修改该时间
	// 时间格式：yyyy-MM-dd HH:mm:ss
	LastEnableTime *string `json:"LastEnableTime,omitnil,omitempty" name:"LastEnableTime"`

	// 日志主题名称
	SrcTopicName *string `json:"SrcTopicName,omitnil,omitempty" name:"SrcTopicName"`

	// 日志集id
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 加工任务目的topic_id以及别名
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// 加工逻辑函数。
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// 兜底topic_id
	BackupTopicID *string `json:"BackupTopicID,omitnil,omitempty" name:"BackupTopicID"`

	// 超限之后是否丢弃日志数据
	BackupGiveUpData *bool `json:"BackupGiveUpData,omitnil,omitempty" name:"BackupGiveUpData"`

	// 是否开启投递服务日志。 1关闭,2开启
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`

	// 任务目标日志主题数量
	TaskDstCount *uint64 `json:"TaskDstCount,omitnil,omitempty" name:"TaskDstCount"`

	// 数据加工类型。0：标准加工任务；1：前置加工任务。
	DataTransformType *uint64 `json:"DataTransformType,omitnil,omitempty" name:"DataTransformType"`

	// 保留失败日志状态。 1:不保留，2:保留
	KeepFailureLog *uint64 `json:"KeepFailureLog,omitnil,omitempty" name:"KeepFailureLog"`

	// 失败日志的字段名称
	FailureLogKey *string `json:"FailureLogKey,omitnil,omitempty" name:"FailureLogKey"`

	// 指定加工数据的开始时间，秒级时间戳。
	// - 日志主题生命周期内的任意时间范围，如果超出了生命周期,只处理生命周期内有数据的部分。
	ProcessFromTimestamp *uint64 `json:"ProcessFromTimestamp,omitnil,omitempty" name:"ProcessFromTimestamp"`

	// 指定加工数据的结束时间，秒级时间戳。
	// 1. 不可指定未来的时间
	// 2. 不填则表示持续执行
	ProcessToTimestamp *uint64 `json:"ProcessToTimestamp,omitnil,omitempty" name:"ProcessToTimestamp"`

	// sql数据源信息
	DataTransformSqlDataSources []*DataTransformSqlDataSource `json:"DataTransformSqlDataSources,omitnil,omitempty" name:"DataTransformSqlDataSources"`

	// 环境变量
	EnvInfos []*EnvInfo `json:"EnvInfos,omitnil,omitempty" name:"EnvInfos"`
}

// Predefined struct for user
type DeleteAlarmNoticeRequestParams struct {
	// 通知渠道组ID。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/api/614/56462)获取通知渠道组ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`
}

type DeleteAlarmNoticeRequest struct {
	*tchttp.BaseRequest
	
	// 通知渠道组ID。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/api/614/56462)获取通知渠道组ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`
}

func (r *DeleteAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmNoticeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmNoticeResponseParams `json:"Response"`
}

func (r *DeleteAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmRequestParams struct {
	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`
}

type DeleteAlarmRequest struct {
	*tchttp.BaseRequest
	
	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`
}

func (r *DeleteAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAlarmResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmResponseParams `json:"Response"`
}

func (r *DeleteAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmShieldRequestParams struct {
	// 屏蔽规则id。通过[获取告警屏蔽配置规则](https://cloud.tencent.com/document/api/614/103650)获取屏蔽规则ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 通知渠道组id。通过[获取告警屏蔽配置规则](https://cloud.tencent.com/document/api/614/103650)获取通知渠道组id
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`
}

type DeleteAlarmShieldRequest struct {
	*tchttp.BaseRequest
	
	// 屏蔽规则id。通过[获取告警屏蔽配置规则](https://cloud.tencent.com/document/api/614/103650)获取屏蔽规则ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 通知渠道组id。通过[获取告警屏蔽配置规则](https://cloud.tencent.com/document/api/614/103650)获取通知渠道组id
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`
}

func (r *DeleteAlarmShieldRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmShieldRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "AlarmNoticeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmShieldRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmShieldResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAlarmShieldResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmShieldResponseParams `json:"Response"`
}

func (r *DeleteAlarmShieldResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmShieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudProductLogCollectionRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云产品标识，支持枚举：CDS、CWP、CDB、TDSQL-C、MongoDB、TDStore、DCDB、MariaDB、PostgreSQL、BH、APIS
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// 日志类型，支持枚举：CDS-AUDIT、CDS-RISK、CDB-AUDIT、TDSQL-C-AUDIT、MongoDB-AUDIT、MongoDB-SlowLog、MongoDB-ErrorLog、TDMYSQL-SLOW、DCDB-AUDIT、DCDB-SLOW、DCDB-ERROR、MariaDB-AUDIT、MariaDB-SLOW、MariaDB-ERROR、PostgreSQL-SLOW、PostgreSQL-ERROR、PostgreSQL-AUDIT、BH-FILELOG、BH-COMMANDLOG、APIS-ACCESS
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 云产品地域。 不同日志类型(LogType)地域入参格式存在差异， 请参考如下示例：
	// - CDS所有日志类型：ap-guangzhou
	// - CDB-AUDIT: gz
	// - TDSQL-C-AUDIT: gz
	// - MongoDB-AUDIT: gz
	// - MongoDB-SlowLog：ap-guangzhou
	// - MongoDB-ErrorLog：ap-guangzhou
	// - TDMYSQL-SLOW：gz
	// - DCDB所有日志类型：gz
	// - MariaDB所有日志类型：gz
	// - PostgreSQL所有日志类型：gz
	// - BH所有日志类型：overseas-polaris(中国香港地区和其他)/fsi-polaris(金融区)/general-polaris(普通区)/intl-sg-prod(国际站)
	// - APIS所有日志类型：gz
	CloudProductRegion *string `json:"CloudProductRegion,omitnil,omitempty" name:"CloudProductRegion"`
}

type DeleteCloudProductLogCollectionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云产品标识，支持枚举：CDS、CWP、CDB、TDSQL-C、MongoDB、TDStore、DCDB、MariaDB、PostgreSQL、BH、APIS
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// 日志类型，支持枚举：CDS-AUDIT、CDS-RISK、CDB-AUDIT、TDSQL-C-AUDIT、MongoDB-AUDIT、MongoDB-SlowLog、MongoDB-ErrorLog、TDMYSQL-SLOW、DCDB-AUDIT、DCDB-SLOW、DCDB-ERROR、MariaDB-AUDIT、MariaDB-SLOW、MariaDB-ERROR、PostgreSQL-SLOW、PostgreSQL-ERROR、PostgreSQL-AUDIT、BH-FILELOG、BH-COMMANDLOG、APIS-ACCESS
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 云产品地域。 不同日志类型(LogType)地域入参格式存在差异， 请参考如下示例：
	// - CDS所有日志类型：ap-guangzhou
	// - CDB-AUDIT: gz
	// - TDSQL-C-AUDIT: gz
	// - MongoDB-AUDIT: gz
	// - MongoDB-SlowLog：ap-guangzhou
	// - MongoDB-ErrorLog：ap-guangzhou
	// - TDMYSQL-SLOW：gz
	// - DCDB所有日志类型：gz
	// - MariaDB所有日志类型：gz
	// - PostgreSQL所有日志类型：gz
	// - BH所有日志类型：overseas-polaris(中国香港地区和其他)/fsi-polaris(金融区)/general-polaris(普通区)/intl-sg-prod(国际站)
	// - APIS所有日志类型：gz
	CloudProductRegion *string `json:"CloudProductRegion,omitnil,omitempty" name:"CloudProductRegion"`
}

func (r *DeleteCloudProductLogCollectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudProductLogCollectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AssumerName")
	delete(f, "LogType")
	delete(f, "CloudProductRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudProductLogCollectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudProductLogCollectionResponseParams struct {
	// 枚举值，0创建中 1创建完成 2删除中 3删除完成
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudProductLogCollectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudProductLogCollectionResponseParams `json:"Response"`
}

func (r *DeleteCloudProductLogCollectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudProductLogCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigExtraRequestParams struct {
	// 特殊采集规则扩展配置ID
	// - 通过[获取特殊采集配置](https://cloud.tencent.com/document/api/614/71164)特殊采集规则扩展配置ID。
	ConfigExtraId *string `json:"ConfigExtraId,omitnil,omitempty" name:"ConfigExtraId"`
}

type DeleteConfigExtraRequest struct {
	*tchttp.BaseRequest
	
	// 特殊采集规则扩展配置ID
	// - 通过[获取特殊采集配置](https://cloud.tencent.com/document/api/614/71164)特殊采集规则扩展配置ID。
	ConfigExtraId *string `json:"ConfigExtraId,omitnil,omitempty" name:"ConfigExtraId"`
}

func (r *DeleteConfigExtraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigExtraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigExtraId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigExtraResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigExtraResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigExtraResponseParams `json:"Response"`
}

func (r *DeleteConfigExtraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigFromMachineGroupRequestParams struct {
	// 机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 采集配置ID
	//  - 通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)获取采集配置Id。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DeleteConfigFromMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// 机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 采集配置ID
	//  - 通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)获取采集配置Id。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigFromMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFromMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigFromMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigFromMachineGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigFromMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigFromMachineGroupResponseParams `json:"Response"`
}

func (r *DeleteConfigFromMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFromMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigRequestParams struct {
	// 采集配置ID
	//  - 通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)获取采集配置Id。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DeleteConfigRequest struct {
	*tchttp.BaseRequest
	
	// 采集配置ID
	//  - 通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)获取采集配置Id。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigResponseParams `json:"Response"`
}

func (r *DeleteConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsoleSharingRequestParams struct {
	// 免密分享Id。
	// - 通过 [获取免密分享列表](https://cloud.tencent.com/document/product/614/109798) 获取免密分享Id。 
	// - 通过 [创建免密分享](https://cloud.tencent.com/document/product/614/109800) 获取免密分享Id。
	SharingId *string `json:"SharingId,omitnil,omitempty" name:"SharingId"`
}

type DeleteConsoleSharingRequest struct {
	*tchttp.BaseRequest
	
	// 免密分享Id。
	// - 通过 [获取免密分享列表](https://cloud.tencent.com/document/product/614/109798) 获取免密分享Id。 
	// - 通过 [创建免密分享](https://cloud.tencent.com/document/product/614/109800) 获取免密分享Id。
	SharingId *string `json:"SharingId,omitnil,omitempty" name:"SharingId"`
}

func (r *DeleteConsoleSharingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsoleSharingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SharingId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConsoleSharingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsoleSharingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConsoleSharingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConsoleSharingResponseParams `json:"Response"`
}

func (r *DeleteConsoleSharingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsoleSharingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsumerRequestParams struct {
	// 投递任务绑定的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 投递任务绑定的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsumerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConsumerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConsumerResponseParams `json:"Response"`
}

func (r *DeleteConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCosRechargeRequestParams struct {
	// COS导入配置Id。
	// - 通过 [获取cos导入配置](https://cloud.tencent.com/document/product/614/88099) 获取COS导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 日志主题Id。
	// -  通过[获取日志主题列表](https://cloud.tencent.com/document/api/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteCosRechargeRequest struct {
	*tchttp.BaseRequest
	
	// COS导入配置Id。
	// - 通过 [获取cos导入配置](https://cloud.tencent.com/document/product/614/88099) 获取COS导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 日志主题Id。
	// -  通过[获取日志主题列表](https://cloud.tencent.com/document/api/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteCosRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCosRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCosRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCosRechargeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCosRechargeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCosRechargeResponseParams `json:"Response"`
}

func (r *DeleteCosRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDashboardSubscribeRequestParams struct {
	// 仪表盘订阅记录id。通过 [获取仪表盘订阅列表](https://cloud.tencent.com/document/api/614/105779)接口获取Id。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteDashboardSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 仪表盘订阅记录id。通过 [获取仪表盘订阅列表](https://cloud.tencent.com/document/api/614/105779)接口获取Id。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteDashboardSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDashboardSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDashboardSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDashboardSubscribeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDashboardSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDashboardSubscribeResponseParams `json:"Response"`
}

func (r *DeleteDashboardSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDashboardSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataTransformRequestParams struct {
	// 数据加工任务ID- 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteDataTransformRequest struct {
	*tchttp.BaseRequest
	
	// 数据加工任务ID- 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteDataTransformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataTransformRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataTransformRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataTransformResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDataTransformResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataTransformResponseParams `json:"Response"`
}

func (r *DeleteDataTransformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDlcDeliverRequestParams struct {
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 任务id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteDlcDeliverRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 任务id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteDlcDeliverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDlcDeliverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDlcDeliverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDlcDeliverResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDlcDeliverResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDlcDeliverResponseParams `json:"Response"`
}

func (r *DeleteDlcDeliverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDlcDeliverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExportRequestParams struct {
	// 日志导出任务Id
	// - 通过[获取日志下载任务列表](https://cloud.tencent.com/document/product/614/56449)获取日志导出任务Id。
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`
}

type DeleteExportRequest struct {
	*tchttp.BaseRequest
	
	// 日志导出任务Id
	// - 通过[获取日志下载任务列表](https://cloud.tencent.com/document/product/614/56449)获取日志导出任务Id。
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`
}

func (r *DeleteExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExportResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteExportResponse struct {
	*tchttp.BaseResponse
	Response *DeleteExportResponseParams `json:"Response"`
}

func (r *DeleteExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIndexRequestParams struct {
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIndexResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIndexResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIndexResponseParams `json:"Response"`
}

func (r *DeleteIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKafkaRechargeRequestParams struct {
	// Kafka导入配置Id。
	// 
	// - 通过 [创建Kafka数据订阅任务](https://cloud.tencent.com/document/product/614/94448)获取Kafka导入配置Id。
	// - 通过 [获取Kafka数据订阅任务列表](https://cloud.tencent.com/document/product/614/94446)获取Kafka导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 导入CLS目标日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteKafkaRechargeRequest struct {
	*tchttp.BaseRequest
	
	// Kafka导入配置Id。
	// 
	// - 通过 [创建Kafka数据订阅任务](https://cloud.tencent.com/document/product/614/94448)获取Kafka导入配置Id。
	// - 通过 [获取Kafka数据订阅任务列表](https://cloud.tencent.com/document/product/614/94446)获取Kafka导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 导入CLS目标日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteKafkaRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKafkaRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKafkaRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKafkaRechargeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteKafkaRechargeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKafkaRechargeResponseParams `json:"Response"`
}

func (r *DeleteKafkaRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogsetRequestParams struct {
	// 日志集Id。通过 [获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

type DeleteLogsetRequest struct {
	*tchttp.BaseRequest
	
	// 日志集Id。通过 [获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

func (r *DeleteLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogsetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLogsetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogsetResponseParams `json:"Response"`
}

func (r *DeleteLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineGroupInfoRequestParams struct {
	// 机器组Id
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/product/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 机器组类型
	// 目前type支持 ip 和 label
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`
}

type DeleteMachineGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 机器组Id
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/product/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 机器组类型
	// 目前type支持 ip 和 label
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`
}

func (r *DeleteMachineGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "MachineGroupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineGroupInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMachineGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineGroupInfoResponseParams `json:"Response"`
}

func (r *DeleteMachineGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineGroupRequestParams struct {
	// 机器组Id
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/product/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// 机器组Id
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/product/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeleteMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineGroupResponseParams `json:"Response"`
}

func (r *DeleteMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNoticeContentRequestParams struct {
	// 通知内容模板ID。-通过[获取通知内容模板](https://cloud.tencent.com/document/api/614/111714)获取通知内容模版ID
	NoticeContentId *string `json:"NoticeContentId,omitnil,omitempty" name:"NoticeContentId"`
}

type DeleteNoticeContentRequest struct {
	*tchttp.BaseRequest
	
	// 通知内容模板ID。-通过[获取通知内容模板](https://cloud.tencent.com/document/api/614/111714)获取通知内容模版ID
	NoticeContentId *string `json:"NoticeContentId,omitnil,omitempty" name:"NoticeContentId"`
}

func (r *DeleteNoticeContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNoticeContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NoticeContentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNoticeContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNoticeContentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNoticeContentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNoticeContentResponseParams `json:"Response"`
}

func (r *DeleteNoticeContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNoticeContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScheduledSqlRequestParams struct {
	// 任务ID，通过[获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519)获取
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 源日志主题ID，通过[获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519)获取
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`
}

type DeleteScheduledSqlRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，通过[获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519)获取
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 源日志主题ID，通过[获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519)获取
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`
}

func (r *DeleteScheduledSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduledSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "SrcTopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScheduledSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScheduledSqlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteScheduledSqlResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScheduledSqlResponseParams `json:"Response"`
}

func (r *DeleteScheduledSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduledSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShipperRequestParams struct {
	// 投递规则Id。
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745)获取ShipperId。
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`
}

type DeleteShipperRequest struct {
	*tchttp.BaseRequest
	
	// 投递规则Id。
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745)获取ShipperId。
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`
}

func (r *DeleteShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShipperResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteShipperResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShipperResponseParams `json:"Response"`
}

func (r *DeleteShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRequestParams struct {
	// 主题ID- 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	
	// 主题ID- 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicResponseParams `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWebCallbackRequestParams struct {
	// 告警渠道回调配置ID。-通过[获取告警渠道回调配置列表](https://cloud.tencent.com/document/api/614/115229)获取告警渠道回调配置ID
	WebCallbackId *string `json:"WebCallbackId,omitnil,omitempty" name:"WebCallbackId"`
}

type DeleteWebCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 告警渠道回调配置ID。-通过[获取告警渠道回调配置列表](https://cloud.tencent.com/document/api/614/115229)获取告警渠道回调配置ID
	WebCallbackId *string `json:"WebCallbackId,omitnil,omitempty" name:"WebCallbackId"`
}

func (r *DeleteWebCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebCallbackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWebCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWebCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWebCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWebCallbackResponseParams `json:"Response"`
}

func (r *DeleteWebCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeliverConfig struct {
	// 地域信息。
	// 
	// 示例：
	//  ap-guangzhou  广州地域；
	// ap-nanjing 南京地域。
	// 
	// 详细信息请查看官网[地域和访问域名](https://cloud.tencent.com/document/product/614/18940)
	// 
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 日志主题ID。-通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 投递数据范围。
	// 
	// 0: 全部日志, 包括告警策略日常周期执行的所有日志，也包括告警策略变更产生的日志，默认值
	// 
	// 1:仅告警触发及恢复日志
	Scope *uint64 `json:"Scope,omitnil,omitempty" name:"Scope"`
}

// Predefined struct for user
type DescribeAlarmNoticesRequestParams struct {
	// name
	// 按照【通知渠道组名称】进行过滤。
	// 类型：String
	// 示例："Filters":[{"Key":"name","Values":["test-notice"]}]
	// 必选：否
	// alarmNoticeId
	// 按照【通知渠道组ID】进行过滤。
	// 类型：String
	// 示例："Filters": [{Key: "alarmNoticeId", Values: ["notice-5281f1d2-6275-4e56-9ec3-a1eb19d8bc2f"]}]
	// 必选：否
	// uid
	// 按照【接收用户ID】进行过滤。
	// 类型：String
	// 示例："Filters": [{Key: "uid", Values: ["1137546"]}]
	// 必选：否
	// groupId
	// 按照【接收用户组ID】进行过滤。
	// 类型：String
	// 示例："Filters": [{Key: "groupId", Values: ["344098"]}]
	// 必选：否
	// 
	// deliverFlag
	// 按照【投递状态】进行过滤。
	// 类型：String
	// 必选：否
	// 可选值： "1":未启用,  "2": 已启用, "3":投递异常
	// 示例："Filters":[{"Key":"deliverFlag","Values":["2"]}]
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否需要返回通知渠道组配置的告警屏蔽统计状态数量信息。
	// - true：需要返回；
	// - false：不返回（默认false）。
	HasAlarmShieldCount *bool `json:"HasAlarmShieldCount,omitnil,omitempty" name:"HasAlarmShieldCount"`
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest
	
	// name
	// 按照【通知渠道组名称】进行过滤。
	// 类型：String
	// 示例："Filters":[{"Key":"name","Values":["test-notice"]}]
	// 必选：否
	// alarmNoticeId
	// 按照【通知渠道组ID】进行过滤。
	// 类型：String
	// 示例："Filters": [{Key: "alarmNoticeId", Values: ["notice-5281f1d2-6275-4e56-9ec3-a1eb19d8bc2f"]}]
	// 必选：否
	// uid
	// 按照【接收用户ID】进行过滤。
	// 类型：String
	// 示例："Filters": [{Key: "uid", Values: ["1137546"]}]
	// 必选：否
	// groupId
	// 按照【接收用户组ID】进行过滤。
	// 类型：String
	// 示例："Filters": [{Key: "groupId", Values: ["344098"]}]
	// 必选：否
	// 
	// deliverFlag
	// 按照【投递状态】进行过滤。
	// 类型：String
	// 必选：否
	// 可选值： "1":未启用,  "2": 已启用, "3":投递异常
	// 示例："Filters":[{"Key":"deliverFlag","Values":["2"]}]
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否需要返回通知渠道组配置的告警屏蔽统计状态数量信息。
	// - true：需要返回；
	// - false：不返回（默认false）。
	HasAlarmShieldCount *bool `json:"HasAlarmShieldCount,omitnil,omitempty" name:"HasAlarmShieldCount"`
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
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "HasAlarmShieldCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticesResponseParams struct {
	// 告警通知模板列表。
	AlarmNotices []*AlarmNotice `json:"AlarmNotices,omitnil,omitempty" name:"AlarmNotices"`

	// 符合条件的告警通知模板总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAlarmShieldsRequestParams struct {
	// 通知渠道组id。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/api/614/56462)获取通知渠道组id
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// - taskId:按照【规则id】进行过滤。类型：String  必选：否
	// - status:按照【规则状态】进行过滤。类型：String。 支持 0:暂未生效，1:生效中，2:已失效。 必选：否
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAlarmShieldsRequest struct {
	*tchttp.BaseRequest
	
	// 通知渠道组id。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/api/614/56462)获取通知渠道组id
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// - taskId:按照【规则id】进行过滤。类型：String  必选：否
	// - status:按照【规则状态】进行过滤。类型：String。 支持 0:暂未生效，1:生效中，2:已失效。 必选：否
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAlarmShieldsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmShieldsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmShieldsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmShieldsResponseParams struct {
	// 符合条件的规则总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 告警屏蔽规则详情
	AlarmShields []*AlarmShieldInfo `json:"AlarmShields,omitnil,omitempty" name:"AlarmShields"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmShieldsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmShieldsResponseParams `json:"Response"`
}

func (r *DescribeAlarmShieldsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmShieldsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmsRequestParams struct {
	// name
	// - 按照【告警策略名称】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：test-alarm
	// 
	// alarmId
	// - 按照【告警策略ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：alarm-b60cf034-c3d6-4b01-xxxx-4e877ebb4751
	// 
	// topicId
	// - 按照【监控对象的日志主题ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：6766f83d-659e-xxxx-a8f7-9104a1012743
	// 
	// enable
	// - 按照【启用状态】进行过滤。
	// - 类型：String
	// - 备注：enable参数值范围: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False。 其它值将返回参数错误信息.
	// - 必选：否
	// - 示例：true
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAlarmsRequest struct {
	*tchttp.BaseRequest
	
	// name
	// - 按照【告警策略名称】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：test-alarm
	// 
	// alarmId
	// - 按照【告警策略ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：alarm-b60cf034-c3d6-4b01-xxxx-4e877ebb4751
	// 
	// topicId
	// - 按照【监控对象的日志主题ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：6766f83d-659e-xxxx-a8f7-9104a1012743
	// 
	// enable
	// - 按照【启用状态】进行过滤。
	// - 类型：String
	// - 备注：enable参数值范围: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False。 其它值将返回参数错误信息.
	// - 必选：否
	// - 示例：true
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAlarmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmsResponseParams struct {
	// 告警策略列表。
	Alarms []*AlarmInfo `json:"Alarms,omitnil,omitempty" name:"Alarms"`

	// 符合查询条件的告警策略数目。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmsResponseParams `json:"Response"`
}

func (r *DescribeAlarmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertRecordHistoryRequestParams struct {
	// 查询时间范围启始时间，毫秒级unix时间戳(ms)
	From *uint64 `json:"From,omitnil,omitempty" name:"From"`

	// 查询时间范围结束时间，毫秒级unix时间戳(ms)
	To *uint64 `json:"To,omitnil,omitempty" name:"To"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// - alertId：按照告警策略ID进行过滤。类型：String 必选：否
	// - topicId：按照监控对象ID进行过滤。类型：String 必选：否
	// - status：按照告警状态进行过滤。类型：String 必选：否，0代表未恢复，1代表已恢复，2代表已失效
	// - alarmLevel：按照告警等级进行过滤。类型：String 必选：否，0代表警告，1代表提醒，2代表紧急
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAlertRecordHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 查询时间范围启始时间，毫秒级unix时间戳(ms)
	From *uint64 `json:"From,omitnil,omitempty" name:"From"`

	// 查询时间范围结束时间，毫秒级unix时间戳(ms)
	To *uint64 `json:"To,omitnil,omitempty" name:"To"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// - alertId：按照告警策略ID进行过滤。类型：String 必选：否
	// - topicId：按照监控对象ID进行过滤。类型：String 必选：否
	// - status：按照告警状态进行过滤。类型：String 必选：否，0代表未恢复，1代表已恢复，2代表已失效
	// - alarmLevel：按照告警等级进行过滤。类型：String 必选：否，0代表警告，1代表提醒，2代表紧急
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAlertRecordHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertRecordHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlertRecordHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertRecordHistoryResponseParams struct {
	// 告警历史总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 告警历史详情
	Records []*AlertHistoryRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlertRecordHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlertRecordHistoryResponseParams `json:"Response"`
}

func (r *DescribeAlertRecordHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertRecordHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudProductLogTasksRequestParams struct {
	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为100，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// - assumerName
	//   - 按照【云产品标识】进行过滤。
	//   - 类型：String
	//   - 必选：否
	//   - 枚举：CDS、CWP、CDB、TDSQL-C、MongoDB、TDStore、DCDB、MariaDB、PostgreSQL、BH、APIS
	// - logType
	//   - 按照【日志类型】进行过滤。
	//   - 类型：String
	//   - 必选：否
	//   - 枚举：CDS-AUDIT、CDS-RISK、CDB-AUDIT、TDSQL-C-AUDIT、MongoDB-AUDIT、MongoDB-SlowLog、MongoDB-ErrorLog、TDMYSQL-SLOW、DCDB-AUDIT、DCDB-SLOW、DCDB-ERROR、MariaDB-AUDIT、MariaDB-SLOW、MariaDB-ERROR、PostgreSQL-SLOW、PostgreSQL-ERROR、PostgreSQL-AUDIT、BH-FILELOG、BH-COMMANDLOG、APIS-ACCESS
	// - instanceId
	//   - 按照【实例ID】进行过滤。
	//   - 类型：String
	//   - 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCloudProductLogTasksRequest struct {
	*tchttp.BaseRequest
	
	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为100，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// - assumerName
	//   - 按照【云产品标识】进行过滤。
	//   - 类型：String
	//   - 必选：否
	//   - 枚举：CDS、CWP、CDB、TDSQL-C、MongoDB、TDStore、DCDB、MariaDB、PostgreSQL、BH、APIS
	// - logType
	//   - 按照【日志类型】进行过滤。
	//   - 类型：String
	//   - 必选：否
	//   - 枚举：CDS-AUDIT、CDS-RISK、CDB-AUDIT、TDSQL-C-AUDIT、MongoDB-AUDIT、MongoDB-SlowLog、MongoDB-ErrorLog、TDMYSQL-SLOW、DCDB-AUDIT、DCDB-SLOW、DCDB-ERROR、MariaDB-AUDIT、MariaDB-SLOW、MariaDB-ERROR、PostgreSQL-SLOW、PostgreSQL-ERROR、PostgreSQL-AUDIT、BH-FILELOG、BH-COMMANDLOG、APIS-ACCESS
	// - instanceId
	//   - 按照【实例ID】进行过滤。
	//   - 类型：String
	//   - 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCloudProductLogTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudProductLogTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudProductLogTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudProductLogTasksResponseParams struct {
	// 日志配置详情列表
	Tasks []*CloudProductLogTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 日志配置总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudProductLogTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudProductLogTasksResponseParams `json:"Response"`
}

func (r *DescribeCloudProductLogTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudProductLogTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigExtrasRequestParams struct {
	// 过滤器，支持如下选项：
	// name
	// - 按照【特殊采集配置名称】进行模糊匹配过滤。
	// - 类型：String
	// - 示例：test-config
	// 
	// configExtraId
	// - 按照【特殊采集配置ID】进行过滤。
	// - 类型：String
	// - 示例：3b83f9d6-3a4d-47f9-9b7f-285c868b2f9a
	// 
	// topicId
	// - 按照【日志主题】进行过滤。
	// - 类型：String
	// - 示例：3581a3be-aa41-423b-995a-54ec84da6264
	// 
	// machineGroupId
	// - 按照【机器组ID】进行过滤。
	// - 类型：String
	// - 示例：f948972f-a063-408c-a59f-8c3230bddaf6
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeConfigExtrasRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器，支持如下选项：
	// name
	// - 按照【特殊采集配置名称】进行模糊匹配过滤。
	// - 类型：String
	// - 示例：test-config
	// 
	// configExtraId
	// - 按照【特殊采集配置ID】进行过滤。
	// - 类型：String
	// - 示例：3b83f9d6-3a4d-47f9-9b7f-285c868b2f9a
	// 
	// topicId
	// - 按照【日志主题】进行过滤。
	// - 类型：String
	// - 示例：3581a3be-aa41-423b-995a-54ec84da6264
	// 
	// machineGroupId
	// - 按照【机器组ID】进行过滤。
	// - 类型：String
	// - 示例：f948972f-a063-408c-a59f-8c3230bddaf6
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeConfigExtrasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigExtrasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigExtrasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigExtrasResponseParams struct {
	// 采集配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configs []*ConfigExtraInfo `json:"Configs,omitnil,omitempty" name:"Configs"`

	// 过滤到的总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigExtrasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigExtrasResponseParams `json:"Response"`
}

func (r *DescribeConfigExtrasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigExtrasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMachineGroupsRequestParams struct {
	// 采集配置ID
	// - 通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)获取采集配置Id。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DescribeConfigMachineGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 采集配置ID
	// - 通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)获取采集配置Id。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

func (r *DescribeConfigMachineGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMachineGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigMachineGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigMachineGroupsResponseParams struct {
	// 采集规则配置绑定的机器组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitnil,omitempty" name:"MachineGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigMachineGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigMachineGroupsResponseParams `json:"Response"`
}

func (r *DescribeConfigMachineGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigsRequestParams struct {
	// configName
	// - 按照【采集配置名称】进行模糊匹配过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：test-config
	// 
	// configId
	// - 按照【采集配置ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：3581a3be-aa41-423b-995a-54ec84da6264
	// 
	// topicId
	// - 按照【日志主题】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：3b83f9d6-3a4d-47f9-9b7f-285c868b2f9a
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest
	
	// configName
	// - 按照【采集配置名称】进行模糊匹配过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：test-config
	// 
	// configId
	// - 按照【采集配置ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：3581a3be-aa41-423b-995a-54ec84da6264
	// 
	// topicId
	// - 按照【日志主题】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 示例：3b83f9d6-3a4d-47f9-9b7f-285c868b2f9a
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigsResponseParams struct {
	// 采集配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configs []*ConfigInfo `json:"Configs,omitnil,omitempty" name:"Configs"`

	// 过滤到的总数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigsResponseParams `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsoleSharingListRequestParams struct {

}

type DescribeConsoleSharingListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeConsoleSharingListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsoleSharingListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsoleSharingListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsoleSharingListResponseParams struct {
	// 分页的总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 控制台免密分享列表
	ConsoleSharingInfos []*ConsoleSharingInfo `json:"ConsoleSharingInfos,omitnil,omitempty" name:"ConsoleSharingInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsoleSharingListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsoleSharingListResponseParams `json:"Response"`
}

func (r *DescribeConsoleSharingListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsoleSharingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerRequestParams struct {
	// 投递任务绑定的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DescribeConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 投递任务绑定的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DescribeConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerResponseParams struct {
	// 投递任务是否生效
	Effective *bool `json:"Effective,omitnil,omitempty" name:"Effective"`

	// 是否投递日志的元数据信息
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// 如果需要投递元数据信息，元数据信息的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka的描述
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// 压缩方式[0:NONE；2:SNAPPY；3:LZ4]
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerResponseParams `json:"Response"`
}

func (r *DescribeConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRechargesRequestParams struct {
	// 日志主题Id。
	// -  通过[获取日志主题列表](https://cloud.tencent.com/document/api/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 状态   status 0: 已创建, 1: 运行中, 2: 已停止, 3: 已完成, 4: 运行失败。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否启用:   0： 未启用  ， 1：启用
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type DescribeCosRechargesRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// -  通过[获取日志主题列表](https://cloud.tencent.com/document/api/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 状态   status 0: 已创建, 1: 运行中, 2: 已停止, 3: 已完成, 4: 运行失败。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否启用:   0： 未启用  ， 1：启用
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *DescribeCosRechargesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRechargesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Status")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosRechargesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRechargesResponseParams struct {
	// 见: CosRechargeInfo 结构描述
	Data []*CosRechargeInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosRechargesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosRechargesResponseParams `json:"Response"`
}

func (r *DescribeCosRechargesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRechargesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDashboardSubscribesRequestParams struct {
	// dashboardId：按照【仪表盘id】进行过滤。类型：String必选：否
	// 
	// - 仪表盘id。通过 [获取仪表盘](https://cloud.tencent.com/document/api/614/95636)接口获取DashboardId。
	// - 入参示例：dashboard-522a5609-1f41-4b11-8086-5afd1d7574f5
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDashboardSubscribesRequest struct {
	*tchttp.BaseRequest
	
	// dashboardId：按照【仪表盘id】进行过滤。类型：String必选：否
	// 
	// - 仪表盘id。通过 [获取仪表盘](https://cloud.tencent.com/document/api/614/95636)接口获取DashboardId。
	// - 入参示例：dashboard-522a5609-1f41-4b11-8086-5afd1d7574f5
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDashboardSubscribesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardSubscribesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDashboardSubscribesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDashboardSubscribesResponseParams struct {
	// 仪表盘订阅列表
	DashboardSubscribeInfos []*DashboardSubscribeInfo `json:"DashboardSubscribeInfos,omitnil,omitempty" name:"DashboardSubscribeInfos"`

	// 总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDashboardSubscribesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDashboardSubscribesResponseParams `json:"Response"`
}

func (r *DescribeDashboardSubscribesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardSubscribesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDashboardsRequestParams struct {
	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// - dashboardId 按照【仪表盘id】进行过滤，类型：String， 必选：否。
	//     - 示例值：dashboard-522a5609-1f41-4b11-8086-5afd1d7574f5
	// - dashboardName 按照【仪表盘名字】进行模糊搜索过滤，类型：String，必选：否。
	//     - 示例值：业务大盘
	// - dashboardRegion 按照【仪表盘地域】进行过滤（兼容老的仪表盘），通过云API创建的仪表盘该属性，类型：String，必选：否。
	//     - 参考  [地域和访问域名](https://cloud.tencent.com/document/product/614/18940)
	//     - 示例：ap-guangzhou
	// - tagKey 按照【标签键】进行过滤，类型：String，必选：否。
	//     - 示例值：
	//     ```
	//     "Filters":[
	//         {
	//             "Key": "tagKey",
	//             "Values": [
	//                 "tag-key-test"
	//             ]
	//         }
	//     ]
	//     ```
	// 
	// - tag:tagKey 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换，类型：String，必选：否，
	//     - 参考 [示例1](https://cloud.tencent.com/document/api/614/95636#4.-.E7.A4.BA.E4.BE.8B) 使用。
	//     ```
	//     "Filters": [
	//         {
	//             "Key": "tag:tag-key-test",
	//             "Values": [
	//                 "12"
	//             ]
	//         }
	//     ]
	//     ```
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 按照topicId和regionId过滤。
	// - topicId:日志主题Id。
	//     -  通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	//     - 示例值：439a5304-08f9-484b-9c4d-46ff57133816
	// - regionId
	//     - 1:广州
	//     - 4:上海
	//     - 5:中国香港
	//     - 7:上海金融
	//     - 8:北京
	//     - 9:新加坡
	//     - 11:深圳金融
	//     - 15:硅谷
	//     - 16:成都
	//     - 17:法兰克福
	//     - 18:首尔
	//     - 19:重庆
	//     - 22:弗吉尼亚
	//     - 23:曼谷
	//     - 25:东京
	//     - 33:南京
	//     - 36:天津
	//     - 39:台北
	//     - 46:北京金融
	//     - 72:雅加达
	//     - 74:圣保罗
	//     - 78:上海自动驾驶云
	TopicIdRegionFilter []*TopicIdAndRegion `json:"TopicIdRegionFilter,omitnil,omitempty" name:"TopicIdRegionFilter"`
}

type DescribeDashboardsRequest struct {
	*tchttp.BaseRequest
	
	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// - dashboardId 按照【仪表盘id】进行过滤，类型：String， 必选：否。
	//     - 示例值：dashboard-522a5609-1f41-4b11-8086-5afd1d7574f5
	// - dashboardName 按照【仪表盘名字】进行模糊搜索过滤，类型：String，必选：否。
	//     - 示例值：业务大盘
	// - dashboardRegion 按照【仪表盘地域】进行过滤（兼容老的仪表盘），通过云API创建的仪表盘该属性，类型：String，必选：否。
	//     - 参考  [地域和访问域名](https://cloud.tencent.com/document/product/614/18940)
	//     - 示例：ap-guangzhou
	// - tagKey 按照【标签键】进行过滤，类型：String，必选：否。
	//     - 示例值：
	//     ```
	//     "Filters":[
	//         {
	//             "Key": "tagKey",
	//             "Values": [
	//                 "tag-key-test"
	//             ]
	//         }
	//     ]
	//     ```
	// 
	// - tag:tagKey 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换，类型：String，必选：否，
	//     - 参考 [示例1](https://cloud.tencent.com/document/api/614/95636#4.-.E7.A4.BA.E4.BE.8B) 使用。
	//     ```
	//     "Filters": [
	//         {
	//             "Key": "tag:tag-key-test",
	//             "Values": [
	//                 "12"
	//             ]
	//         }
	//     ]
	//     ```
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 按照topicId和regionId过滤。
	// - topicId:日志主题Id。
	//     -  通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	//     - 示例值：439a5304-08f9-484b-9c4d-46ff57133816
	// - regionId
	//     - 1:广州
	//     - 4:上海
	//     - 5:中国香港
	//     - 7:上海金融
	//     - 8:北京
	//     - 9:新加坡
	//     - 11:深圳金融
	//     - 15:硅谷
	//     - 16:成都
	//     - 17:法兰克福
	//     - 18:首尔
	//     - 19:重庆
	//     - 22:弗吉尼亚
	//     - 23:曼谷
	//     - 25:东京
	//     - 33:南京
	//     - 36:天津
	//     - 39:台北
	//     - 46:北京金融
	//     - 72:雅加达
	//     - 74:圣保罗
	//     - 78:上海自动驾驶云
	TopicIdRegionFilter []*TopicIdAndRegion `json:"TopicIdRegionFilter,omitnil,omitempty" name:"TopicIdRegionFilter"`
}

func (r *DescribeDashboardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "TopicIdRegionFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDashboardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDashboardsResponseParams struct {
	// 仪表盘的数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 仪表盘详细明细
	DashboardInfos []*DashboardInfo `json:"DashboardInfos,omitnil,omitempty" name:"DashboardInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDashboardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDashboardsResponseParams `json:"Response"`
}

func (r *DescribeDashboardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataTransformInfoRequestParams struct {
	// - taskName
	// 按照【加工任务名称】进行过滤。
	// 类型：String
	// 必选：否
	// 示例：test-task
	// 
	// - taskId
	// 按照【加工任务id】进行过滤。
	// 类型：String
	// 必选：否
	// 示例：a3622556-6402-4942-b4ff-5ae32ec29810
	// 数据加工任务ID- 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	// 
	// - topicId
	// 按照【源topicId】进行过滤。
	// 类型：String
	// 必选：否
	// 示例：756cec3e-a0a5-44c3-85a8-090870582000
	// 日志主题ID
	// 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	// 
	// - status
	// 按照【 任务运行状态】进行过滤。 1：准备中，2：运行中，3：停止中，4：已停止
	// 类型：String
	// 必选：否
	// 示例：1
	// 
	// - hasServiceLog
	// 按照【是否开启服务日志】进行过滤。 1：未开启，2：已开启
	// 类型：String
	// 必选：否
	// 示例：1
	// 
	// - dstTopicType
	// 按照【目标topic类型】进行过滤。  1：固定，2：动态
	// 类型：String
	// 必选：否
	// 示例：1
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 默认值为2.   1: 获取单个任务的详细信息 2：获取任务列表
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Type为1， 此参数必填
	// 数据加工任务ID- 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeDataTransformInfoRequest struct {
	*tchttp.BaseRequest
	
	// - taskName
	// 按照【加工任务名称】进行过滤。
	// 类型：String
	// 必选：否
	// 示例：test-task
	// 
	// - taskId
	// 按照【加工任务id】进行过滤。
	// 类型：String
	// 必选：否
	// 示例：a3622556-6402-4942-b4ff-5ae32ec29810
	// 数据加工任务ID- 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	// 
	// - topicId
	// 按照【源topicId】进行过滤。
	// 类型：String
	// 必选：否
	// 示例：756cec3e-a0a5-44c3-85a8-090870582000
	// 日志主题ID
	// 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	// 
	// - status
	// 按照【 任务运行状态】进行过滤。 1：准备中，2：运行中，3：停止中，4：已停止
	// 类型：String
	// 必选：否
	// 示例：1
	// 
	// - hasServiceLog
	// 按照【是否开启服务日志】进行过滤。 1：未开启，2：已开启
	// 类型：String
	// 必选：否
	// 示例：1
	// 
	// - dstTopicType
	// 按照【目标topic类型】进行过滤。  1：固定，2：动态
	// 类型：String
	// 必选：否
	// 示例：1
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 默认值为2.   1: 获取单个任务的详细信息 2：获取任务列表
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Type为1， 此参数必填
	// 数据加工任务ID- 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeDataTransformInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataTransformInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataTransformInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataTransformInfoResponseParams struct {
	// 数据加工任务列表信息
	DataTransformTaskInfos []*DataTransformTaskInfo `json:"DataTransformTaskInfos,omitnil,omitempty" name:"DataTransformTaskInfos"`

	// 任务总次数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataTransformInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataTransformInfoResponseParams `json:"Response"`
}

func (r *DescribeDataTransformInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataTransformInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDlcDeliversRequestParams struct {
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// - taskId
	// 按照【任务id】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - name
	// 按照【任务名称】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - tableName
	// 按照【数据表】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - statusFlag
	// 按照【状态】进行过滤。支持："1","2","3","4"。含义：1：RUNNING，2：STOPPED，3：FINISHED，4：FAILED
	// 类型：String
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为10。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDlcDeliversRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// - taskId
	// 按照【任务id】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - name
	// 按照【任务名称】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - tableName
	// 按照【数据表】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - statusFlag
	// 按照【状态】进行过滤。支持："1","2","3","4"。含义：1：RUNNING，2：STOPPED，3：FINISHED，4：FAILED
	// 类型：String
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为10。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDlcDeliversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDlcDeliversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDlcDeliversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDlcDeliversResponseParams struct {
	// 告警渠道回调配置列表。
	Infos []*DlcDeliverInfo `json:"Infos,omitnil,omitempty" name:"Infos"`

	// 符合条件的通知内容配置总数。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDlcDeliversResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDlcDeliversResponseParams `json:"Response"`
}

func (r *DescribeDlcDeliversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDlcDeliversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportsRequestParams struct {
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeExportsRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportsResponseParams struct {
	// 日志导出列表
	Exports []*ExportInfo `json:"Exports,omitnil,omitempty" name:"Exports"`

	// 总数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExportsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExportsResponseParams `json:"Response"`
}

func (r *DescribeExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexRequestParams struct {
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DescribeIndexRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DescribeIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexResponseParams struct {
	// 日志主题Id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 索引状态。true：开启状态，false：关闭状态
	// 开启后可对日志进行检索分析，将产生索引流量、索引存储及相应费用。[费用详情](https://cloud.tencent.com/document/product/614/45802#.E8.AE.A1.E8.B4.B9.E9.A1.B9)
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 索引配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 索引修改时间，初始值为索引创建时间。格式 `YYYY-MM-DD HH:MM:SS`
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 内置保留字段（`__FILENAME__`，`__HOSTNAME__`及`__SOURCE__`）是否包含至全文索引
	// * false:不包含
	// * true:包含
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// 元数据字段（前缀为`__TAG__`的字段）是否包含至全文索引
	// * 0:仅包含开启键值索引的元数据字段
	// * 1:包含所有元数据字段
	// * 2:不包含任何元数据字段
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`

	// 自定义日志解析异常存储字段。
	CoverageField *string `json:"CoverageField,omitnil,omitempty" name:"CoverageField"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIndexResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexResponseParams `json:"Response"`
}

func (r *DescribeIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaConsumerGroupDetailRequestParams struct {
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 消费组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

type DescribeKafkaConsumerGroupDetailRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 消费组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

func (r *DescribeKafkaConsumerGroupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaConsumerGroupDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Group")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKafkaConsumerGroupDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaConsumerGroupDetailResponseParams struct {
	// 日志集id
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 消费组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 消费组信息列表
	PartitionInfos []*GroupPartitionInfo `json:"PartitionInfos,omitnil,omitempty" name:"PartitionInfos"`

	// Empty：组内没有成员，但存在已提交的偏移量。所有消费者都离开但保留了偏移量
	// Dead：组内没有成员，且没有已提交的偏移量。组被删除或长时间无活动
	// Stable：组内成员正常消费，分区分配平衡。正常运行状态
	// PreparingRebalance：组正在准备重新平衡。有新成员加入或现有成员离开
	// CompletingRebalance：组正在准备重新平衡。有新成员加入或现有成员离开
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKafkaConsumerGroupDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKafkaConsumerGroupDetailResponseParams `json:"Response"`
}

func (r *DescribeKafkaConsumerGroupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaConsumerGroupDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaConsumerGroupListRequestParams struct {
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// - group
	// 按照【消费组名称】进行过滤。
	// 类型：String
	// 必选：否
	// 示例：消费组1
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为10。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeKafkaConsumerGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// - group
	// 按照【消费组名称】进行过滤。
	// 类型：String
	// 必选：否
	// 示例：消费组1
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为10。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeKafkaConsumerGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaConsumerGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKafkaConsumerGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaConsumerGroupListResponseParams struct {
	// 日志主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志集id
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 总个数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 消费组信息列表
	Groups []*ConsumerGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKafkaConsumerGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKafkaConsumerGroupListResponseParams `json:"Response"`
}

func (r *DescribeKafkaConsumerGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaConsumerGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaConsumerRequestParams struct {
	// 日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`
}

type DescribeKafkaConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`
}

func (r *DescribeKafkaConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKafkaConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaConsumerResponseParams struct {
	// Kafka协议消费是否打开
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// KafkaConsumer 消费时使用的Topic参数
	TopicID *string `json:"TopicID,omitnil,omitempty" name:"TopicID"`

	// 压缩方式[0:NONE；2:SNAPPY；3:LZ4]
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// kafka协议消费数据格式
	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitnil,omitempty" name:"ConsumerContent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKafkaConsumerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKafkaConsumerResponseParams `json:"Response"`
}

func (r *DescribeKafkaConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaRechargesRequestParams struct {
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 导入配置Id。
	// - 通过 [创建Kafka数据订阅任务](https://cloud.tencent.com/document/product/614/94448)获取Kafka导入配置Id。
	// - 通过 [获取Kafka数据订阅任务列表](https://cloud.tencent.com/document/product/614/94446)获取Kafka导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 状态。1: 运行中，2: 暂停，3：错误
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeKafkaRechargesRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 导入配置Id。
	// - 通过 [创建Kafka数据订阅任务](https://cloud.tencent.com/document/product/614/94448)获取Kafka导入配置Id。
	// - 通过 [获取Kafka数据订阅任务列表](https://cloud.tencent.com/document/product/614/94446)获取Kafka导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 状态。1: 运行中，2: 暂停，3：错误
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeKafkaRechargesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaRechargesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Id")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKafkaRechargesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaRechargesResponseParams struct {
	// KafkaRechargeInfo 信息列表
	Infos []*KafkaRechargeInfo `json:"Infos,omitnil,omitempty" name:"Infos"`

	// Kafka导入信息总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKafkaRechargesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKafkaRechargesResponseParams `json:"Response"`
}

func (r *DescribeKafkaRechargesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaRechargesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogContextRequestParams struct {
	// 要查询的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志时间,  即 [检索分析日志](https://cloud.tencent.com/document/product/614/56447) 接口返回信息中Results结构体中的Time，需按照 UTC+8 时区将该毫秒级Unix时间戳转换为 YYYY-mm-dd HH:MM:SS.FFF 格式的字符串。
	BTime *string `json:"BTime,omitnil,omitempty" name:"BTime"`

	// 日志包序号，即 [检索分析日志](https://cloud.tencent.com/document/product/614/56447) 接口返回信息中Results结构体中的PkgId。
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// 日志包内一条日志的序号，即 [检索分析日志](https://cloud.tencent.com/document/product/614/56447) 接口返回信息中Results结构中的PkgLogId。
	PkgLogId *int64 `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// 前${PrevLogs}条日志，默认值10。
	PrevLogs *int64 `json:"PrevLogs,omitnil,omitempty" name:"PrevLogs"`

	// 后${NextLogs}条日志，默认值10。
	NextLogs *int64 `json:"NextLogs,omitnil,omitempty" name:"NextLogs"`

	// 检索语句，对日志上下文进行过滤，最大长度为12KB
	// 语句由 <a href="https://cloud.tencent.com/document/product/614/47044" target="_blank">[检索条件]</a>构成，不支持SQL语句
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 上下文检索的开始时间，单位：毫秒级时间戳
	// 注意：
	// - From为空时，表示上下文检索的开始时间不做限制
	// - From和To非空时，From < To
	// - 暂时仅支持上海 / 弗吉尼亚/ 新加坡地域
	From *uint64 `json:"From,omitnil,omitempty" name:"From"`

	// 上下文检索的结束时间，单位：毫秒级时间戳。
	// 注意：
	// - To为空时，表示上下文检索的结束时间不做限制
	// - From和To非空时，From < To
	// - 暂时仅支持上海 / 弗吉尼亚/ 新加坡地域
	To *uint64 `json:"To,omitnil,omitempty" name:"To"`
}

type DescribeLogContextRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志时间,  即 [检索分析日志](https://cloud.tencent.com/document/product/614/56447) 接口返回信息中Results结构体中的Time，需按照 UTC+8 时区将该毫秒级Unix时间戳转换为 YYYY-mm-dd HH:MM:SS.FFF 格式的字符串。
	BTime *string `json:"BTime,omitnil,omitempty" name:"BTime"`

	// 日志包序号，即 [检索分析日志](https://cloud.tencent.com/document/product/614/56447) 接口返回信息中Results结构体中的PkgId。
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// 日志包内一条日志的序号，即 [检索分析日志](https://cloud.tencent.com/document/product/614/56447) 接口返回信息中Results结构中的PkgLogId。
	PkgLogId *int64 `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// 前${PrevLogs}条日志，默认值10。
	PrevLogs *int64 `json:"PrevLogs,omitnil,omitempty" name:"PrevLogs"`

	// 后${NextLogs}条日志，默认值10。
	NextLogs *int64 `json:"NextLogs,omitnil,omitempty" name:"NextLogs"`

	// 检索语句，对日志上下文进行过滤，最大长度为12KB
	// 语句由 <a href="https://cloud.tencent.com/document/product/614/47044" target="_blank">[检索条件]</a>构成，不支持SQL语句
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 上下文检索的开始时间，单位：毫秒级时间戳
	// 注意：
	// - From为空时，表示上下文检索的开始时间不做限制
	// - From和To非空时，From < To
	// - 暂时仅支持上海 / 弗吉尼亚/ 新加坡地域
	From *uint64 `json:"From,omitnil,omitempty" name:"From"`

	// 上下文检索的结束时间，单位：毫秒级时间戳。
	// 注意：
	// - To为空时，表示上下文检索的结束时间不做限制
	// - From和To非空时，From < To
	// - 暂时仅支持上海 / 弗吉尼亚/ 新加坡地域
	To *uint64 `json:"To,omitnil,omitempty" name:"To"`
}

func (r *DescribeLogContextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "BTime")
	delete(f, "PkgId")
	delete(f, "PkgLogId")
	delete(f, "PrevLogs")
	delete(f, "NextLogs")
	delete(f, "Query")
	delete(f, "From")
	delete(f, "To")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogContextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogContextResponseParams struct {
	// 日志上下文信息集合
	LogContextInfos []*LogContextInfo `json:"LogContextInfos,omitnil,omitempty" name:"LogContextInfos"`

	// 上文日志是否已经返回完成（当PrevOver为false，表示有上文日志还未全部返回）。
	PrevOver *bool `json:"PrevOver,omitnil,omitempty" name:"PrevOver"`

	// 下文日志是否已经返回完成（当NextOver为false，表示有下文日志还未全部返回）。
	NextOver *bool `json:"NextOver,omitnil,omitempty" name:"NextOver"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogContextResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogContextResponseParams `json:"Response"`
}

func (r *DescribeLogContextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogHistogramRequestParams struct {
	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 检索分析语句。
	// 语句由 [检索条件] | [SQL语句]构成，无需对日志进行统计分析时，可省略其中的管道符 | 及SQL语句。
	// 使用*或空字符串可查询所有日志。
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 要查询的日志主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 时间间隔: 单位ms  限制性条件：(To-From) / interval <= 200
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 检索语法规则，默认值为0。
	// 0：Lucene语法，1：CQL语法。
	// 详细说明参见<a href="https://cloud.tencent.com/document/product/614/47044#RetrievesConditionalRules" target="_blank">检索条件语法规则</a>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

type DescribeLogHistogramRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 检索分析语句。
	// 语句由 [检索条件] | [SQL语句]构成，无需对日志进行统计分析时，可省略其中的管道符 | 及SQL语句。
	// 使用*或空字符串可查询所有日志。
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 要查询的日志主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 时间间隔: 单位ms  限制性条件：(To-From) / interval <= 200
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 检索语法规则，默认值为0。
	// 0：Lucene语法，1：CQL语法。
	// 详细说明参见<a href="https://cloud.tencent.com/document/product/614/47044#RetrievesConditionalRules" target="_blank">检索条件语法规则</a>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

func (r *DescribeLogHistogramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogHistogramRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "TopicId")
	delete(f, "Interval")
	delete(f, "SyntaxRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogHistogramRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogHistogramResponseParams struct {
	// 统计周期： 单位ms
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 命中关键字的日志总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 周期内统计结果详情
	HistogramInfos []*HistogramInfo `json:"HistogramInfos,omitnil,omitempty" name:"HistogramInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogHistogramResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogHistogramResponseParams `json:"Response"`
}

func (r *DescribeLogHistogramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsetsRequestParams struct {
	// logsetName
	// - 按照【日志集名称】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 通过 [获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集名称。
	// logsetId
	// - 按照【日志集ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 通过 [获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	// 
	// tagKey
	// - 按照【标签键】进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// tag:tagKey
	// - 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。
	// - 类型：String
	// - 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeLogsetsRequest struct {
	*tchttp.BaseRequest
	
	// logsetName
	// - 按照【日志集名称】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 通过 [获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集名称。
	// logsetId
	// - 按照【日志集ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// - 通过 [获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	// 
	// tagKey
	// - 按照【标签键】进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// tag:tagKey
	// - 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。
	// - 类型：String
	// - 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeLogsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsetsResponseParams struct {
	// 分页的总数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志集列表
	Logsets []*LogsetInfo `json:"Logsets,omitnil,omitempty" name:"Logsets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogsetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogsetsResponseParams `json:"Response"`
}

func (r *DescribeLogsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineGroupConfigsRequestParams struct {
	// 机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DescribeMachineGroupConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DescribeMachineGroupConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineGroupConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineGroupConfigsResponseParams struct {
	// 采集规则配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configs []*ConfigInfo `json:"Configs,omitnil,omitempty" name:"Configs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMachineGroupConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineGroupConfigsResponseParams `json:"Response"`
}

func (r *DescribeMachineGroupConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineGroupsRequestParams struct {
	// 过滤条件
	// machineGroupName
	// - 按照【机器组名称】进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// machineGroupId
	// - 按照【机器组ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// osType
	// - 按照【操作系统类型】进行过滤。0： Linux；1： Windows
	// - 类型：Int
	// - 必选：否
	// 
	// tagKey
	// - 按照【标签键】进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// tag:tagKey
	// - 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。
	// - 类型：String
	// - 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMachineGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	// machineGroupName
	// - 按照【机器组名称】进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// machineGroupId
	// - 按照【机器组ID】进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// osType
	// - 按照【操作系统类型】进行过滤。0： Linux；1： Windows
	// - 类型：Int
	// - 必选：否
	// 
	// tagKey
	// - 按照【标签键】进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// tag:tagKey
	// - 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。
	// - 类型：String
	// - 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMachineGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineGroupsResponseParams struct {
	// 机器组信息列表
	MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitnil,omitempty" name:"MachineGroups"`

	// 分页的总数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMachineGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineGroupsResponseParams `json:"Response"`
}

func (r *DescribeMachineGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesRequestParams struct {
	// 查询的机器组ID。
	// 
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// ip
	// - 按照ip进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// instance
	// - 按照实例id进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// version
	// - 按照LogListener版本进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// status
	// - 按照机器状态进行过滤。
	// - 类型：String
	// - 必选：否
	// - 可选值：0：离线，1：正常
	// 
	// offlineTime
	// - 按照机器离线时间进行过滤。
	// - 类型：String
	// - 必选：否
	// - -可选值：0：无离线时间，12：12小时内，24：一天内，48：两天内，99：两天前
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目。最大支持100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest
	
	// 查询的机器组ID。
	// 
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// ip
	// - 按照ip进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// instance
	// - 按照实例id进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// version
	// - 按照LogListener版本进行过滤。
	// - 类型：String
	// - 必选：否
	// 
	// status
	// - 按照机器状态进行过滤。
	// - 类型：String
	// - 必选：否
	// - 可选值：0：离线，1：正常
	// 
	// offlineTime
	// - 按照机器离线时间进行过滤。
	// - 类型：String
	// - 必选：否
	// - -可选值：0：无离线时间，12：12小时内，24：一天内，48：两天内，99：两天前
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目。最大支持100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachinesResponseParams struct {
	// 机器状态信息组
	Machines []*MachineInfo `json:"Machines,omitnil,omitempty" name:"Machines"`

	// 机器组是否开启自动升级功能。 0：未开启自动升级；1：开启了自动升级。
	AutoUpdate *int64 `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// 机器组自动升级功能预设开始时间
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// 机器组自动升级功能预设结束时间
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// 当前用户可用最新的Loglistener版本
	LatestAgentVersion *string `json:"LatestAgentVersion,omitnil,omitempty" name:"LatestAgentVersion"`

	// 是否开启服务日志
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// 总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachinesResponseParams `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoticeContentsRequestParams struct {
	// <li> name
	// 按照【通知内容模板名称】进行过滤。
	// 类型：String
	// 必选：否
	// </li>
	// <li> noticeContentId
	// 按照【通知内容模板ID】进行过滤。
	// 类型：String
	// 必选：否
	// </li>
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNoticeContentsRequest struct {
	*tchttp.BaseRequest
	
	// <li> name
	// 按照【通知内容模板名称】进行过滤。
	// 类型：String
	// 必选：否
	// </li>
	// <li> noticeContentId
	// 按照【通知内容模板ID】进行过滤。
	// 类型：String
	// 必选：否
	// </li>
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNoticeContentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoticeContentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNoticeContentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoticeContentsResponseParams struct {
	// 通知内容模板列表。
	NoticeContents []*NoticeContentTemplate `json:"NoticeContents,omitnil,omitempty" name:"NoticeContents"`

	// 符合条件的通知内容模板总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNoticeContentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNoticeContentsResponseParams `json:"Response"`
}

func (r *DescribeNoticeContentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoticeContentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePartitionsRequestParams struct {
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DescribePartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

func (r *DescribePartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePartitionsResponseParams struct {
	// 分区列表
	Partitions []*PartitionInfo `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePartitionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePartitionsResponseParams `json:"Response"`
}

func (r *DescribePartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScheduledSqlInfoRequestParams struct {
	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// - srcTopicName按照【源日志主题名称】进行过滤，模糊匹配。类型：String。必选：否。示例：业务日志主题1 ，通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题名称。
	// - dstTopicName按照【目标日志主题名称】进行过滤，模糊匹配。类型：String。必选：否。示例：业务日志主题 2，通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题名称。
	// - srcTopicId按照【源日志主题ID】进行过滤。类型：String。必选：否。示例：a4478687-2382-4486-9692-de7986350f6b ，通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题id。
	// - dstTopicId按照【目标日志主题ID】进行过滤。类型：String。必选：否。示例：bd4d3375-d72a-4cd2-988d-d8eda2bd62b0，通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题id。
	// - bizType按照【主题类型】进行过滤，0：日志主题；1：指标主题。类型：String。必选：否
	// - status按照【任务状态】进行过滤，1：运行；2：停止；3：异常。类型：String。必选：否
	// - taskName按照【任务名称】进行过滤，模糊匹配。类型：String。必选：否。示例：metricTask ，通过 [获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519) 获取任务名称。
	// - taskId按照【任务ID】进行过滤，模糊匹配。类型：String。必选：否。示例：9c64f9c1-a14e-4b59-b074-5b73cac3dd66 ，通过 [获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519) 获取任务id。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeScheduledSqlInfoRequest struct {
	*tchttp.BaseRequest
	
	// 分页的偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// - srcTopicName按照【源日志主题名称】进行过滤，模糊匹配。类型：String。必选：否。示例：业务日志主题1 ，通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题名称。
	// - dstTopicName按照【目标日志主题名称】进行过滤，模糊匹配。类型：String。必选：否。示例：业务日志主题 2，通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题名称。
	// - srcTopicId按照【源日志主题ID】进行过滤。类型：String。必选：否。示例：a4478687-2382-4486-9692-de7986350f6b ，通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题id。
	// - dstTopicId按照【目标日志主题ID】进行过滤。类型：String。必选：否。示例：bd4d3375-d72a-4cd2-988d-d8eda2bd62b0，通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题id。
	// - bizType按照【主题类型】进行过滤，0：日志主题；1：指标主题。类型：String。必选：否
	// - status按照【任务状态】进行过滤，1：运行；2：停止；3：异常。类型：String。必选：否
	// - taskName按照【任务名称】进行过滤，模糊匹配。类型：String。必选：否。示例：metricTask ，通过 [获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519) 获取任务名称。
	// - taskId按照【任务ID】进行过滤，模糊匹配。类型：String。必选：否。示例：9c64f9c1-a14e-4b59-b074-5b73cac3dd66 ，通过 [获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519) 获取任务id。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeScheduledSqlInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScheduledSqlInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Name")
	delete(f, "TaskId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScheduledSqlInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScheduledSqlInfoResponseParams struct {
	// ScheduledSQL任务列表信息
	ScheduledSqlTaskInfos []*ScheduledSqlTaskInfo `json:"ScheduledSqlTaskInfos,omitnil,omitempty" name:"ScheduledSqlTaskInfos"`

	// 任务总次数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScheduledSqlInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScheduledSqlInfoResponseParams `json:"Response"`
}

func (r *DescribeScheduledSqlInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScheduledSqlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShipperTasksRequestParams struct {
	// 投递规则Id。
	// 
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745)获取ShipperId。
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// 查询的开始时间戳，支持最近3天的查询， 毫秒。
	// StartTime必须小于EndTime
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询的结束时间戳， 毫秒。
	// StartTime必须小于EndTime
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeShipperTasksRequest struct {
	*tchttp.BaseRequest
	
	// 投递规则Id。
	// 
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745)获取ShipperId。
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// 查询的开始时间戳，支持最近3天的查询， 毫秒。
	// StartTime必须小于EndTime
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询的结束时间戳， 毫秒。
	// StartTime必须小于EndTime
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeShipperTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShipperTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShipperTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShipperTasksResponseParams struct {
	// 投递任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*ShipperTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShipperTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShipperTasksResponseParams `json:"Response"`
}

func (r *DescribeShipperTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShipperTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShippersRequestParams struct {
	// - shipperName：按照【投递规则名称】进行过滤。
	//     类型：String。
	//     必选：否
	// - shipperId：按照【投递规则ID】进行过滤。
	//     类型：String。
	//     必选：否
	// - topicId：按照【日志主题】进行过滤。
	//     类型：String。
	//     必选：否
	// - taskStatus：按照【任务运行状态】进行过滤。 支持`0`：停止，`1`：运行中，`2`：异常
	//     类型：String
	//     必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为10。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 控制Filters相关字段是否为精确匹配。  0: 默认值，shipperName模糊匹配 1: shipperName 精确匹配
	PreciseSearch *uint64 `json:"PreciseSearch,omitnil,omitempty" name:"PreciseSearch"`
}

type DescribeShippersRequest struct {
	*tchttp.BaseRequest
	
	// - shipperName：按照【投递规则名称】进行过滤。
	//     类型：String。
	//     必选：否
	// - shipperId：按照【投递规则ID】进行过滤。
	//     类型：String。
	//     必选：否
	// - topicId：按照【日志主题】进行过滤。
	//     类型：String。
	//     必选：否
	// - taskStatus：按照【任务运行状态】进行过滤。 支持`0`：停止，`1`：运行中，`2`：异常
	//     类型：String
	//     必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为10。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 控制Filters相关字段是否为精确匹配。  0: 默认值，shipperName模糊匹配 1: shipperName 精确匹配
	PreciseSearch *uint64 `json:"PreciseSearch,omitnil,omitempty" name:"PreciseSearch"`
}

func (r *DescribeShippersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShippersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PreciseSearch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShippersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShippersResponseParams struct {
	// 投递规则列表
	Shippers []*ShipperInfo `json:"Shippers,omitnil,omitempty" name:"Shippers"`

	// 本次查询获取到的总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShippersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShippersResponseParams `json:"Response"`
}

func (r *DescribeShippersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShippersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicsRequestParams struct {
	// <ul><li>topicName 按照【主题名称】进行过滤，默认为模糊匹配，可使用 PreciseSearch 参数设置为精确匹配。类型：String。必选：否</li>
	// <li>logsetName 按照【日志集名称】进行过滤，默认为模糊匹配，可使用 PreciseSearch 参数设置为精确匹配。类型：String。必选：否</li>
	// <li>topicId 按照【主题ID】进行过滤。类型：String。必选：否</li>
	// <li>logsetId 按照【日志集ID】进行过滤，可通过调用 <a href="https://cloud.tencent.com/document/product/614/58624">DescribeLogsets</a> 查询已创建的日志集列表或登录控制台进行查看；也可以调用<a href="https://cloud.tencent.com/document/product/614/58626">CreateLogset</a> 创建新的日志集。类型：String。必选：否</li>
	// <li>tagKey 按照【标签键】进行过滤。类型：String。必选：否</li>
	// <li>tag:tagKey 按照【标签键值对】进行过滤。tagKey 使用具体的标签键进行替换，例如 tag:exampleKey。类型：String。必选：否</li>
	// <li>storageType 按照【主题的存储类型】进行过滤。可选值 hot（标准存储），cold（低频存储）类型：String。必选：否</li></ul>
	// 注意：每次请求的 Filters 的上限为10，Filter.Values 的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 控制Filters相关字段是否为精确匹配。
	// <ul><li>0: 默认值，topicName 和 logsetName 模糊匹配</li>
	// <li>1: topicName   精确匹配</li>
	// <li>2: logsetName精确匹配</li>
	// <li>3: topicName 和logsetName 都精确匹配</li></ul>
	PreciseSearch *uint64 `json:"PreciseSearch,omitnil,omitempty" name:"PreciseSearch"`

	// 主题类型
	// - 0:日志主题，默认值
	// - 1:指标主题
	BizType *uint64 `json:"BizType,omitnil,omitempty" name:"BizType"`
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest
	
	// <ul><li>topicName 按照【主题名称】进行过滤，默认为模糊匹配，可使用 PreciseSearch 参数设置为精确匹配。类型：String。必选：否</li>
	// <li>logsetName 按照【日志集名称】进行过滤，默认为模糊匹配，可使用 PreciseSearch 参数设置为精确匹配。类型：String。必选：否</li>
	// <li>topicId 按照【主题ID】进行过滤。类型：String。必选：否</li>
	// <li>logsetId 按照【日志集ID】进行过滤，可通过调用 <a href="https://cloud.tencent.com/document/product/614/58624">DescribeLogsets</a> 查询已创建的日志集列表或登录控制台进行查看；也可以调用<a href="https://cloud.tencent.com/document/product/614/58626">CreateLogset</a> 创建新的日志集。类型：String。必选：否</li>
	// <li>tagKey 按照【标签键】进行过滤。类型：String。必选：否</li>
	// <li>tag:tagKey 按照【标签键值对】进行过滤。tagKey 使用具体的标签键进行替换，例如 tag:exampleKey。类型：String。必选：否</li>
	// <li>storageType 按照【主题的存储类型】进行过滤。可选值 hot（标准存储），cold（低频存储）类型：String。必选：否</li></ul>
	// 注意：每次请求的 Filters 的上限为10，Filter.Values 的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 控制Filters相关字段是否为精确匹配。
	// <ul><li>0: 默认值，topicName 和 logsetName 模糊匹配</li>
	// <li>1: topicName   精确匹配</li>
	// <li>2: logsetName精确匹配</li>
	// <li>3: topicName 和logsetName 都精确匹配</li></ul>
	PreciseSearch *uint64 `json:"PreciseSearch,omitnil,omitempty" name:"PreciseSearch"`

	// 主题类型
	// - 0:日志主题，默认值
	// - 1:指标主题
	BizType *uint64 `json:"BizType,omitnil,omitempty" name:"BizType"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PreciseSearch")
	delete(f, "BizType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicsResponseParams struct {
	// 主题列表
	Topics []*TopicInfo `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 总数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicsResponseParams `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebCallbacksRequestParams struct {
	// - name
	// 按照【告警渠道回调配置名称】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - webCallbackId
	// 按照【告警渠道回调配置ID】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - type
	// 按照【告警渠道回调配置渠道类型】进行过滤。
	// 类型：String
	// 必选：否
	// 枚举值：WeCom，DingTalk，Lark，Http
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeWebCallbacksRequest struct {
	*tchttp.BaseRequest
	
	// - name
	// 按照【告警渠道回调配置名称】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - webCallbackId
	// 按照【告警渠道回调配置ID】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// - type
	// 按照【告警渠道回调配置渠道类型】进行过滤。
	// 类型：String
	// 必选：否
	// 枚举值：WeCom，DingTalk，Lark，Http
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeWebCallbacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebCallbacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebCallbacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebCallbacksResponseParams struct {
	// 告警渠道回调配置列表。
	WebCallbacks []*WebCallbackInfo `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`

	// 符合条件的通知内容配置总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWebCallbacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebCallbacksResponseParams `json:"Response"`
}

func (r *DescribeWebCallbacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebCallbacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DlcDeliverInfo struct {
	// 任务id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 账号id。
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 日志主题id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 任务名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 投递类型，0：实时投递，1：历史投递
	DeliverType *uint64 `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 投递文件大小，单位MB
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 投递间隔 单位秒
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 投递时间范围的开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 投递时间范围的结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// dlc配置信息
	DlcInfo *DlcInfo `json:"DlcInfo,omitnil,omitempty" name:"DlcInfo"`

	// 是否开启投递服务日志。1关闭，2开启
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`

	// 任务状态。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务进度。历史投递任务生效。
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 日志主题类型。0:标准主题，1:指标主题
	BizType *uint64 `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 任务创建时间。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务修改时间。
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DlcFiledInfo struct {
	// cls日志中的字段名
	ClsField *string `json:"ClsField,omitnil,omitempty" name:"ClsField"`

	// 数据湖计算服务表的列名
	DlcField *string `json:"DlcField,omitnil,omitempty" name:"DlcField"`

	// 数据湖计算服务字段类型
	DlcFieldType *string `json:"DlcFieldType,omitnil,omitempty" name:"DlcFieldType"`

	// 解析失败填充字段
	FillField *string `json:"FillField,omitnil,omitempty" name:"FillField"`

	// 是否禁用
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`
}

type DlcInfo struct {
	// dlc表信息
	TableInfo *DlcTableInfo `json:"TableInfo,omitnil,omitempty" name:"TableInfo"`

	// dlc数据字段信息
	FieldInfos []*DlcFiledInfo `json:"FieldInfos,omitnil,omitempty" name:"FieldInfos"`

	// dlc分区信息
	PartitionInfos []*DlcPartitionInfo `json:"PartitionInfos,omitnil,omitempty" name:"PartitionInfos"`

	// dlc分区额外信息
	PartitionExtra *DlcPartitionExtra `json:"PartitionExtra,omitnil,omitempty" name:"PartitionExtra"`
}

type DlcPartitionExtra struct {
	// 时间格式	eg: %Y-%m-%d %H:%M:%S.%f
	TimeFormat *string `json:"TimeFormat,omitnil,omitempty" name:"TimeFormat"`

	// 时间时区
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type DlcPartitionInfo struct {
	// cls日志中的字段名
	ClsField *string `json:"ClsField,omitnil,omitempty" name:"ClsField"`

	// dlc表的列名
	DlcField *string `json:"DlcField,omitnil,omitempty" name:"DlcField"`

	// dlc字段类型
	DlcFieldType *string `json:"DlcFieldType,omitnil,omitempty" name:"DlcFieldType"`
}

type DlcTableInfo struct {
	// 数据目录
	DataDirectory *string `json:"DataDirectory,omitnil,omitempty" name:"DataDirectory"`

	// 数据库
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据表
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type DynamicIndex struct {
	// 键值索引自动配置开关
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type EnvInfo struct {
	// 环境变量名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 环境变量值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type EscalateNoticeInfo struct {
	// 告警通知模板接收者信息。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// 告警通知模板回调信息。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`

	// 告警升级开关。`true`：开启告警升级、`false`：关闭告警升级，默认：false
	Escalate *bool `json:"Escalate,omitnil,omitempty" name:"Escalate"`

	// 告警升级间隔。单位：分钟，范围`[1，14400]`
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 升级条件。`1`：无人认领且未恢复、`2`：未恢复，默认为1
	// - 无人认领且未恢复：告警没有恢复并且没有人认领则升级
	// - 未恢复：当前告警持续未恢复则升级
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警升级后下一个环节的通知渠道配置，最多可配置5个环节。
	EscalateNotice *EscalateNoticeInfo `json:"EscalateNotice,omitnil,omitempty" name:"EscalateNotice"`
}

type EventLog struct {
	// 事件通道，支持
	// - Application 应用日志
	// - Security 安全日志
	// - Setup 启动日志
	// - System 系统日志
	// - ALL 所有日志
	EventChannel *string `json:"EventChannel,omitnil,omitempty" name:"EventChannel"`

	// 时间字段（Timestamp）支持的类型
	// - 1（用户自定义时间）
	// - 2（当前时间）
	TimeType *uint64 `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 时间，用户选择自定义时间类型时，需要指定时间，单位秒
	// 格式：时间戳，1754897446
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 事件ID过滤列表
	// 	
	// 选填，为空表示不做过滤
	// 支持正向过滤单个值（例：20）或范围（例：0-20），也支持反向过滤单个值(例：-20)
	// 多个过滤项之间可由逗号隔开，例：1-200,-100表示采集1-200范围内除了100以外的事件日志
	EventIDs []*string `json:"EventIDs,omitnil,omitempty" name:"EventIDs"`
}

type ExcludePathInfo struct {
	// 类型，选填File或Path
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Type对应的具体内容
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ExportInfo struct {
	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志导出任务ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// 日志导出查询语句
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 日志导出文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 日志文件大小
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 日志导出时间排序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 日志导出格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 日志导出数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 日志下载状态。Processing:导出正在进行中，Completed:导出完成，Failed:导出失败，Expired:日志导出已过期(三天有效期), Queuing 排队中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 日志导出起始时间，毫秒时间戳
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 日志导出结束时间，毫秒时间戳
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 日志导出路径,有效期一个小时，请尽快使用该路径下载。
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`

	// 日志导出创建时间
	// 时间格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 语法规则。 默认值为0。
	// 0：Lucene语法，1：CQL语法。
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// 导出字段
	DerivedFields []*string `json:"DerivedFields,omitnil,omitempty" name:"DerivedFields"`
}

type ExtractRuleInfo struct {
	// 时间字段的key名字，TimeKey和TimeFormat必须成对出现
	TimeKey *string `json:"TimeKey,omitnil,omitempty" name:"TimeKey"`

	// 时间字段的格式，参考c语言的strftime函数对于时间的格式说明输出参数
	// - 参考 [配置时间格式](https://cloud.tencent.com/document/product/614/38614) 文档 
	TimeFormat *string `json:"TimeFormat,omitnil,omitempty" name:"TimeFormat"`

	// 分隔符类型日志的分隔符，只有LogType为delimiter_log时有效
	Delimiter *string `json:"Delimiter,omitnil,omitempty" name:"Delimiter"`

	// 整条日志匹配规则，只有LogType为fullregex_log时有效
	LogRegex *string `json:"LogRegex,omitnil,omitempty" name:"LogRegex"`

	// 行首匹配规则，只有LogType为multiline_log或fullregex_log时有效
	BeginRegex *string `json:"BeginRegex,omitnil,omitempty" name:"BeginRegex"`

	// 取的每个字段的key名字，为空的key代表丢弃这个字段，只有LogType为delimiter_log时有效，json_log的日志使用json本身的key。限制100个。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 日志过滤规则列表（旧版），需要过滤日志的key，及其对应的regex。
	//  注意：2.9.3及以上版本LogListener ，建议使用AdvanceFilterRules配置日志过滤规则。
	FilterKeyRegex []*KeyRegexInfo `json:"FilterKeyRegex,omitnil,omitempty" name:"FilterKeyRegex"`

	// 解析失败日志是否上传，true表示上传，false表示不上传
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnMatchUpLoadSwitch *bool `json:"UnMatchUpLoadSwitch,omitnil,omitempty" name:"UnMatchUpLoadSwitch"`

	// 失败日志的key，当UnMatchUpLoadSwitch为true时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnMatchLogKey *string `json:"UnMatchLogKey,omitnil,omitempty" name:"UnMatchLogKey"`

	// 增量采集模式下的回溯数据量，默认：-1（全量采集）；其他非负数表示增量采集（从最新的位置，往前采集${Backtracking}字节（Byte）的日志）最大支持1073741824（1G）。
	// 注意：
	// - COS导入不支持此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Backtracking *int64 `json:"Backtracking,omitnil,omitempty" name:"Backtracking"`

	// 是否为Gbk编码。 0：否；1：是。
	// 注意
	// - 目前取0值时，表示UTF-8编码
	// - COS导入不支持此字段。
	IsGBK *int64 `json:"IsGBK,omitnil,omitempty" name:"IsGBK"`

	// 是否为标准json。  0：否； 1：是。
	// - 标准json指采集器使用业界标准开源解析器进行json解析，非标json指采集器使用CLS自研json解析器进行解析，两种解析器没有本质区别，建议客户使用标准json进行解析。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JsonStandard *int64 `json:"JsonStandard,omitnil,omitempty" name:"JsonStandard"`

	// syslog传输协议，取值为tcp或者udp，只有在LogType为service_syslog时生效，其余类型无需填写。
	// 注意：
	// - 该字段适用于：创建采集规则配置、修改采集规则配置。
	// - COS导入不支持此字段。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// syslog系统日志采集指定采集器监听的地址和端口 ，形式：[ip]:[port]，只有在LogType为service_syslog时生效，其余类型无需填写。
	// 注意：
	// - 该字段适用于：创建采集规则配置、修改采集规则配置。
	// - COS导入不支持此字段。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// rfc3164：指定系统日志采集使用RFC3164协议解析日志。
	// rfc5424：指定系统日志采集使用RFC5424协议解析日志。
	// auto：自动匹配rfc3164或者rfc5424其中一种协议。
	// 只有在LogType为service_syslog时生效，其余类型无需填写。
	// 注意：
	// - 该字段适用于：创建采集规则配置、修改采集规则配置
	// - COS导入不支持此字段。
	ParseProtocol *string `json:"ParseProtocol,omitnil,omitempty" name:"ParseProtocol"`

	// 元数据类型。0: 不使用元数据信息；1:使用机器组元数据；2:使用用户自定义元数据；3:使用采集配置路径。
	// 注意：
	// - COS导入不支持此字段。
	MetadataType *int64 `json:"MetadataType,omitnil,omitempty" name:"MetadataType"`

	// 采集配置路径正则表达式。
	// 
	// ```
	// 请用"()"标识路径中目标字段对应的正则表达式，解析时将"()"视为捕获组，并以__TAG__.{i}:{目标字段}的形式与日志一起上报，其中i为捕获组的序号。
	// 若不希望以序号为键名，可以通过命名捕获组"(?<{键名}>{正则})"自定义键名，并以__TAG__.{键名}:{目标字段}的形式与日志一起上报。最多支持5个捕获组
	// ```
	// 
	// 注意：
	// - MetadataType为3时必填。
	// - COS导入不支持此字段。
	PathRegex *string `json:"PathRegex,omitnil,omitempty" name:"PathRegex"`

	// 用户自定义元数据信息。
	// 注意：
	// - MetadataType为2时必填。
	// - COS导入不支持此字段。
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`

	// Windows事件日志采集规则，只有在LogType为windows_event_log时生效，其余类型无需填写。
	EventLogRules []*EventLog `json:"EventLogRules,omitnil,omitempty" name:"EventLogRules"`

	// 日志过滤规则列表（新版）。
	// 注意：
	// - 2.9.3以下版本LogListener不支持， 请使用FilterKeyRegex配置日志过滤规则。
	// - 自建k8s采集配置（CreateConfigExtra、ModifyConfigExtra）不支持此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceFilterRules []*AdvanceFilterRuleInfo `json:"AdvanceFilterRules,omitnil,omitempty" name:"AdvanceFilterRules"`
}

type FilePathInfo struct {
	// 文件路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 文件名称
	File *string `json:"File,omitnil,omitempty" name:"File"`
}

type Filter struct {
	// 需要过滤的字段。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 需要过滤的值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FilterRuleInfo struct {
	// 过滤规则Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 过滤规则
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`

	// 过滤规则Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type FullTextInfo struct {
	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 全文索引的分词符，其中的每个字符代表一个分词符；
	// 仅支持英文符号、\n\t\r及转义符\；
	// 注意：\n\t\r本身已被转义，直接使用双引号包裹即可作为入参，无需再次转义。使用API Explorer进行调试时请使用JSON参数输入方式，以避免\n\t\r被重复转义
	Tokenizer *string `json:"Tokenizer,omitnil,omitempty" name:"Tokenizer"`

	// 是否包含中文
	ContainZH *bool `json:"ContainZH,omitnil,omitempty" name:"ContainZH"`
}

// Predefined struct for user
type GetAlarmLogRequestParams struct {
	// 要查询的执行详情的起始时间，Unix时间戳，单位ms。
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的执行详情的结束时间，Unix时间戳，单位ms。
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 查询过滤条件，例如：
	// - 按告警策略ID查询：`alert_id:"alarm-0745ec00-e605-xxxx-b50b-54afe61fc971"`
	//    - 通过[获取告警策略列表](https://cloud.tencent.com/document/api/614/56461)获取告警策略ID
	// - 按监控对象ID查询：`monitored_object:"823d8bfa-76a7-xxxx-8399-8cda74d4009b" `
	//   - 通过[获取告警策略列表](https://cloud.tencent.com/document/api/614/56461)获取监控对象ID
	// - 按告警策略ID及监控对象ID查询：`alert_id:"alarm-0745ec00-e605-xxxx-b50b-54afe61fc971" AND monitored_object:"823d8bfa-76a7-xxxx-8399-8cda74d4009b"`
	// - 按告警策略ID及监控对象ID查询支持SQL语句：`(alert_id:"alarm-5ce45495-09e8-4d58-xxxx-768134bf330c") AND (monitored_object:"3c514e84-6f1f-46ec-xxxx-05de6163f7fe") AND NOT condition_evaluate_result: "Skip" AND condition_evaluate_result:[* TO *] | SELECT count(*) as top50StatisticsTotalCount, count_if(condition_evaluate_result='ProcessError') as top50StatisticsFailureCount, count_if(notification_send_result!='NotSend') as top50NoticeTotalCount, count_if(notification_send_result='SendPartFail' or notification_send_result='SendFail') as top50NoticeFailureCount, alert_id, alert_name, monitored_object, topic_type, happen_threshold, alert_threshold, notify_template group by alert_id, alert_name, monitored_object,topic_type, happen_threshold, alert_threshold, notify_template order by top50StatisticsTotalCount desc limit 1`
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 单次查询返回的执行详情条数，最大值为1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。
	// 注意：
	// * 透传该参数时，请勿修改除该参数外的其它参数
	// * 仅当检索分析语句(Query)不包含SQL时有效，SQL获取后续结果参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果排序方式参考<a href="https://cloud.tencent.com/document/product/614/58978" target="_blank">SQL ORDER BY语法</a>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// true：代表使用新的检索结果返回方式，输出参数AnalysisRecords和Columns有效；
	// false：代表使用老的检索结果返回方式，输出AnalysisResults和ColNames有效；
	// 两种返回方式在编码格式上有少量区别，建议使用true。
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`
}

type GetAlarmLogRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的执行详情的起始时间，Unix时间戳，单位ms。
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的执行详情的结束时间，Unix时间戳，单位ms。
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 查询过滤条件，例如：
	// - 按告警策略ID查询：`alert_id:"alarm-0745ec00-e605-xxxx-b50b-54afe61fc971"`
	//    - 通过[获取告警策略列表](https://cloud.tencent.com/document/api/614/56461)获取告警策略ID
	// - 按监控对象ID查询：`monitored_object:"823d8bfa-76a7-xxxx-8399-8cda74d4009b" `
	//   - 通过[获取告警策略列表](https://cloud.tencent.com/document/api/614/56461)获取监控对象ID
	// - 按告警策略ID及监控对象ID查询：`alert_id:"alarm-0745ec00-e605-xxxx-b50b-54afe61fc971" AND monitored_object:"823d8bfa-76a7-xxxx-8399-8cda74d4009b"`
	// - 按告警策略ID及监控对象ID查询支持SQL语句：`(alert_id:"alarm-5ce45495-09e8-4d58-xxxx-768134bf330c") AND (monitored_object:"3c514e84-6f1f-46ec-xxxx-05de6163f7fe") AND NOT condition_evaluate_result: "Skip" AND condition_evaluate_result:[* TO *] | SELECT count(*) as top50StatisticsTotalCount, count_if(condition_evaluate_result='ProcessError') as top50StatisticsFailureCount, count_if(notification_send_result!='NotSend') as top50NoticeTotalCount, count_if(notification_send_result='SendPartFail' or notification_send_result='SendFail') as top50NoticeFailureCount, alert_id, alert_name, monitored_object, topic_type, happen_threshold, alert_threshold, notify_template group by alert_id, alert_name, monitored_object,topic_type, happen_threshold, alert_threshold, notify_template order by top50StatisticsTotalCount desc limit 1`
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 单次查询返回的执行详情条数，最大值为1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。
	// 注意：
	// * 透传该参数时，请勿修改除该参数外的其它参数
	// * 仅当检索分析语句(Query)不包含SQL时有效，SQL获取后续结果参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果排序方式参考<a href="https://cloud.tencent.com/document/product/614/58978" target="_blank">SQL ORDER BY语法</a>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// true：代表使用新的检索结果返回方式，输出参数AnalysisRecords和Columns有效；
	// false：代表使用老的检索结果返回方式，输出AnalysisResults和ColNames有效；
	// 两种返回方式在编码格式上有少量区别，建议使用true。
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`
}

func (r *GetAlarmLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseNewAnalysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAlarmLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAlarmLogResponseParams struct {
	// 加载后续详情的Context
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 指定时间范围内的告警执行详情是否完整返回
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 返回的结果是否为SQL分析结果
	Analysis *bool `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// 分析结果的列名，如果Query语句有SQL查询，则返回查询字段的列名；
	// 否则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColNames []*string `json:"ColNames,omitnil,omitempty" name:"ColNames"`

	// 执行详情查询结果。
	// 当Query字段无SQL语句时，返回查询结果。
	// 当Query字段有SQL语句时，可能返回null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*LogInfo `json:"Results,omitnil,omitempty" name:"Results"`

	// 执行详情统计分析结果。当Query字段有SQL语句时，返回SQL统计结果，否则可能返回null。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisResults []*LogItems `json:"AnalysisResults,omitnil,omitempty" name:"AnalysisResults"`

	// 执行详情统计分析结果；UseNewAnalysis为true有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisRecords []*string `json:"AnalysisRecords,omitnil,omitempty" name:"AnalysisRecords"`

	// 分析结果的列名， UseNewAnalysis为true有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAlarmLogResponse struct {
	*tchttp.BaseResponse
	Response *GetAlarmLogResponseParams `json:"Response"`
}

func (r *GetAlarmLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupPartitionInfo struct {
	// 分区id
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`

	// 分区最新数据时间戳，单位：s
	CommitTimestamp *int64 `json:"CommitTimestamp,omitnil,omitempty" name:"CommitTimestamp"`

	// 消费者
	Consumer *string `json:"Consumer,omitnil,omitempty" name:"Consumer"`
}

type GroupTriggerConditionInfo struct {
	// 分组触发字段名称
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 分组触发字段值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type HighLightItem struct {
	// 高亮的日志字段名称
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 高亮的关键词
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type HistogramInfo struct {
	// 统计周期内的日志条数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 按 period 取整后的 unix timestamp： 单位毫秒
	BTime *int64 `json:"BTime,omitnil,omitempty" name:"BTime"`
}

type HostFileInfo struct {
	// 日志文件夹
	LogPath *string `json:"LogPath,omitnil,omitempty" name:"LogPath"`

	// 日志文件名
	FilePattern *string `json:"FilePattern,omitnil,omitempty" name:"FilePattern"`

	// metadata信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomLabels []*string `json:"CustomLabels,omitnil,omitempty" name:"CustomLabels"`
}

type JsonInfo struct {
	// 启用标志
	EnableTag *bool `json:"EnableTag,omitnil,omitempty" name:"EnableTag"`

	// 元数据信息列表, 可选值为 __SOURCE__、__FILENAME__、__TIMESTAMP__、__HOSTNAME__。
	// 
	// - __SOURCE__：日志采集的源 IP，示例：10.0.1.2
	// - __FILENAME__：日志采集的文件名，示例：/data/log/nginx/access.log
	// - __TIMESTAMP__：日志时间戳（毫秒级别 Unix 时间戳），按时间范围检索日志时，将自动使用该时间对日志进行检索，在控制台显示为“日志时间”，示例：1640005601188
	// - __HOSTNAME__：日志来源机器名称，需使用2.7.4及以上版本的 Loglistener 才会采集该字段，示例：localhost
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaFields []*string `json:"MetaFields,omitnil,omitempty" name:"MetaFields"`

	// 投递Json格式，0：字符串方式投递；1:以结构化方式投递
	JsonType *int64 `json:"JsonType,omitnil,omitempty" name:"JsonType"`
}

type KafkaConsumerContent struct {
	// 消费数据格式。 0：原始内容；1：JSON。
	Format *int64 `json:"Format,omitnil,omitempty" name:"Format"`

	// 是否投递 TAG 信息
	// Format为0时，此字段不需要赋值
	EnableTag *bool `json:"EnableTag,omitnil,omitempty" name:"EnableTag"`

	// 元数据信息列表, 可选值为：\_\_SOURCE\_\_、\_\_FILENAME\_\_
	// 、\_\_TIMESTAMP\_\_、\_\_HOSTNAME\_\_、\_\_PKGID\_\_
	// Format为0时，此字段不需要赋值
	MetaFields []*string `json:"MetaFields,omitnil,omitempty" name:"MetaFields"`

	// tag数据处理方式：1:不平铺（默认值）；2:平铺。
	// 
	// 不平铺示例：
	// TAG信息：`{"__TAG__":{"fieldA":200,"fieldB":"text"}}`
	// 不平铺：`{"__TAG__":{"fieldA":200,"fieldB":"text"}}`
	// 
	// 平铺示例：
	// TAG信息：`{"__TAG__":{"fieldA":200,"fieldB":"text"}}`
	// 平铺：`{"__TAG__.fieldA":200,"__TAG__.fieldB":"text"}`
	TagTransaction *int64 `json:"TagTransaction,omitnil,omitempty" name:"TagTransaction"`

	// 消费数据Json格式：
	// 1：不转义（默认格式）
	// 2：转义
	// 
	// 投递Json格式。
	// JsonType为1：和原始日志一致，不转义。示例：
	// 日志原文：`{"a":"aa", "b":{"b1":"b1b1", "c1":"c1c1"}}`
	// 投递到Ckafka：`{"a":"aa", "b":{"b1":"b1b1", "c1":"c1c1"}}`
	// 
	// JsonType为2：转义。示例：
	// 日志原文：`{"a":"aa", "b":{"b1":"b1b1", "c1":"c1c1"}}`
	// 投递到Ckafka：`{"a":"aa","b":"{\"b1\":\"b1b1\", \"c1\":\"c1c1\"}"}`
	JsonType *int64 `json:"JsonType,omitnil,omitempty" name:"JsonType"`
}

type KafkaProtocolInfo struct {
	// 协议类型，支持的协议类型包括 plaintext、sasl_plaintext 或 sasl_ssl。建议使用 sasl_ssl，此协议会进行连接加密同时需要用户认证。
	// 
	// - 当IsEncryptionAddr为true时，Protocol必填。
	// - 支持的协议类型如下：
	//     - plaintext：纯文本无加密协议
	//     - sasl_ssl：SASL 认证 + SSL 加密
	//     - ssl：纯 SSL/TLS 加密协议
	//     - sasl_plaintext：SASL 认证 + 非加密通道
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 加密类型，支持 PLAIN、SCRAM-SHA-256 或 SCRAM-SHA-512。
	// 
	// - 当Protocol为  `sasl_plaintext` 或 `sasl_ssl` 时 Mechanism 必填。
	// - 支持加密类型如下
	//     -  PLAIN：明文认证
	//     -  SCRAM-SHA-256：基于挑战-响应机制，使用PBKDF2-HMAC-SHA256算法
	//     -  SCRAM-SHA-512：增强版SCRAM，使用PBKDF2-HMAC-SHA512算法
	Mechanism *string `json:"Mechanism,omitnil,omitempty" name:"Mechanism"`

	// 用户名。
	// 当Protocol为sasl_plaintext或sasl_ssl时必填
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户密码。
	// 当Protocol为sasl_plaintext或sasl_ssl时必填
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type KafkaRechargeInfo struct {
	// Kafka数据订阅配置的ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka导入任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 导入Kafka类型，0: 腾讯云CKafka，1: 用户自建Kafka
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// 腾讯云CKafka实例ID，KafkaType为0时必填
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// 服务地址
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// ServerAddr是否为加密连接	
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// 加密访问协议，IsEncryptionAddr参数为true时必填
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// 用户Kafka消费组名称	
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 状态 ，1：运行中；2：暂停。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 导入数据位置，-2:最早（默认），-1：最晚
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 创建时间。格式`YYYY-MM-DD HH:MM:SS`
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。格式`YYYY-MM-DD HH:MM:SS`
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 日志导入规则
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`

	// 用户kafka拓展信息
	UserKafkaMeta *UserKafkaMeta `json:"UserKafkaMeta,omitnil,omitempty" name:"UserKafkaMeta"`
}

type KeyRegexInfo struct {
	// 需要过滤日志的key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// key对应的过滤规则regex
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`
}

type KeyValueInfo struct {
	// 需要配置键值或者元字段索引的字段名称，仅支持字母、数字、下划线和-./@，且不能以下划线开头
	// 
	// 注意：
	// 1，元字段（tag）的Key无需额外添加`__TAG__.`前缀，与上传日志时对应的字段Key一致即可，腾讯云控制台展示时将自动添加`__TAG__.`前缀
	// 2，键值索引（KeyValue）及元字段索引（Tag）中的Key总数不能超过300
	// 3，Key的层级不能超过10层，例如a.b.c.d.e.f.g.h.j.k
	// 4，不允许同时包含json父子级字段，例如a及a.b
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 字段的索引描述信息
	Value *ValueInfo `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogContextInfo struct {
	// 日志来源设备
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 采集路径
	Filename *string `json:"Filename,omitnil,omitempty" name:"Filename"`

	// 日志内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 日志包序号
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// 日志包内一条日志的序号
	PkgLogId *int64 `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// 日志时间戳
	BTime *int64 `json:"BTime,omitnil,omitempty" name:"BTime"`

	// 日志来源主机名称
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 原始日志(仅在日志创建索引异常时有值)
	RawLog *string `json:"RawLog,omitnil,omitempty" name:"RawLog"`

	// 日志创建索引异常原因(仅在日志创建索引异常时有值)
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`

	// 日志内容的高亮描述信息
	HighLights []*HighLightItem `json:"HighLights,omitnil,omitempty" name:"HighLights"`
}

type LogInfo struct {
	// 日志时间，单位ms
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志来源IP
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 日志文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 日志上报请求包的ID
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// 请求包内日志的ID
	PkgLogId *string `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// 符合检索条件的关键词，一般用于高亮显示。仅支持键值检索，不支持全文检索	
	HighLights []*HighLightItem `json:"HighLights,omitnil,omitempty" name:"HighLights"`

	// 日志内容的Json序列化字符串
	LogJson *string `json:"LogJson,omitnil,omitempty" name:"LogJson"`

	// 日志来源主机名称
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 原始日志(仅在日志创建索引异常时有值)
	RawLog *string `json:"RawLog,omitnil,omitempty" name:"RawLog"`

	// 日志创建索引异常原因(仅在日志创建索引异常时有值)
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`
}

type LogItem struct {
	// 日志Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 日志Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogItems struct {
	// 分析结果返回的KV数据对
	Data []*LogItem `json:"Data,omitnil,omitempty" name:"Data"`
}

type LogRechargeRuleInfo struct {
	// 导入类型，支持json_log：json格式日志，minimalist_log: 单行全文，fullregex_log: 单行完全正则
	RechargeType *string `json:"RechargeType,omitnil,omitempty" name:"RechargeType"`

	// 解析编码格式，0: UTF-8（默认值），1: GBK
	EncodingFormat *uint64 `json:"EncodingFormat,omitnil,omitempty" name:"EncodingFormat"`

	// 使用默认时间状态。true：开启后将使用系统当前时间或 Kafka 消息时间戳作为日志时间戳；false：关闭将使用日志中的时间字段作为日志时间戳。 默认：true
	DefaultTimeSwitch *bool `json:"DefaultTimeSwitch,omitnil,omitempty" name:"DefaultTimeSwitch"`

	// 整条日志匹配规则，只有RechargeType为fullregex_log时有效
	LogRegex *string `json:"LogRegex,omitnil,omitempty" name:"LogRegex"`

	// 解析失败日志是否上传，true表示上传，false表示不上传
	UnMatchLogSwitch *bool `json:"UnMatchLogSwitch,omitnil,omitempty" name:"UnMatchLogSwitch"`

	// 解析失败日志的键名称
	UnMatchLogKey *string `json:"UnMatchLogKey,omitnil,omitempty" name:"UnMatchLogKey"`

	// 解析失败日志时间来源，0: 系统当前时间，1: Kafka消息时间戳
	UnMatchLogTimeSrc *uint64 `json:"UnMatchLogTimeSrc,omitnil,omitempty" name:"UnMatchLogTimeSrc"`

	// 默认时间来源，0: 系统当前时间，1: Kafka消息时间戳
	DefaultTimeSrc *uint64 `json:"DefaultTimeSrc,omitnil,omitempty" name:"DefaultTimeSrc"`

	// 时间字段，日志中代表时间的字段名。
	// 
	// - 当DefaultTimeSwitch为false，且RechargeType数据提取模式为 `json_log` JSON-文件日志 或 `fullregex_log` 单行完全正则-文件日志时， TimeKey不能为空。
	TimeKey *string `json:"TimeKey,omitnil,omitempty" name:"TimeKey"`

	// 时间提取正则表达式。
	// - 当DefaultTimeSwitch为false，且RechargeType数据提取模式为 `minimalist_log` 单行全文-文件日志时， TimeRegex不能为空。
	// - 仅需输入日志中代表时间的字段的正则表达式即可；若匹配到多个字段，将使用第一个。
	//    例：日志原文为：message with time 2022-08-08 14:20:20，则您可以设置提取时间正则为\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d
	TimeRegex *string `json:"TimeRegex,omitnil,omitempty" name:"TimeRegex"`

	// 时间字段格式。
	// - 当DefaultTimeSwitch为false时， TimeFormat不能为空。
	TimeFormat *string `json:"TimeFormat,omitnil,omitempty" name:"TimeFormat"`

	// 时间字段时区。
	// - 当DefaultTimeSwitch为false时， TimeZone不能为空。
	// - 时区格式规则
	// ​前缀​：使用 GMT 或 UTC 作为时区基准
	// ​偏移量​：
	//     - `-` 表示西时区（比基准时间晚）
	//     - `+` 表示东时区（比基准时间早）
	//     -  格式为 ±HH:MM（小时:分钟）
	// 
	// - 当前支持：
	// ```
	// "GMT-12:00" 
	// "GMT-11:00" 
	// "GMT-10:00" 
	// "GMT-09:30" 
	// "GMT-09:00" 
	// "GMT-08:00" 
	// "GMT-07:00" 
	// "GMT-06:00" 
	// "GMT-05:00" 
	// "GMT-04:00" 
	// "GMT-03:30" 
	// "GMT-03:00" 
	// "GMT-02:00" 
	// "GMT-01:00" 
	// "GMT+00:00"
	// "GMT+01:00"
	// "GMT+02:00"
	// "GMT+03:30"
	// "GMT+04:00"
	// "GMT+04:30"
	// "GMT+05:00"
	// "GMT+05:30"
	// "GMT+05:45"
	// "GMT+06:00"
	// "GMT+06:30"
	// "GMT+07:00"
	// "GMT+08:00"
	// "GMT+09:00"
	// "GMT+09:30"
	// "GMT+10:00"
	// "GMT+10:30"
	// "GMT+11:00"
	// "GMT+11:30"
	// "GMT+12:00"
	// "GMT+12:45"
	// "GMT+13:00"
	// "GMT+14:00"
	// "UTC-11:00"
	// "UTC-10:00"
	// "UTC-09:00"
	// "UTC-08:00"
	// "UTC-12:00"
	// "UTC-07:00"
	// "UTC-06:00"
	// "UTC-05:00"
	// "UTC-04:30"
	// "UTC-04:00"
	// "UTC-03:30"
	// "UTC-03:00"
	// "UTC-02:00"
	// "UTC-01:00"
	// "UTC+00:00"
	// "UTC+01:00"
	// "UTC+02:00"
	// "UTC+03:00"
	// "UTC+03:30"
	// "UTC+04:00"
	// "UTC+04:30"
	// "UTC+05:00"
	// "UTC+05:45"
	// "UTC+06:00"
	// "UTC+06:30"
	// "UTC+07:00"
	// "UTC+08:00"
	// "UTC+09:00"
	// "UTC+09:30"
	// "UTC+10:00"
	// "UTC+11:00"
	// "UTC+12:00"
	// "UTC+13:00"
	// ```
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// 元数据信息，Kafka导入支持kafka_topic,kafka_partition,kafka_offset,kafka_timestamp
	Metadata []*string `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// 日志Key列表，RechargeType为full_regex_log、delimiter_log时必填
	Keys []*string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// json解析模式，开启首层数据解析
	ParseArray *bool `json:"ParseArray,omitnil,omitempty" name:"ParseArray"`

	// 分隔符解析模式-分隔符
	// 当解析格式为分隔符提取时，该字段必填
	Delimiter *string `json:"Delimiter,omitnil,omitempty" name:"Delimiter"`
}

type LogsetInfo struct {
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名称
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 创建时间。格式 `YYYY-MM-DD HH:MM:SS`
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 若AssumerUin非空，则表示创建该日志集的服务方Uin
	AssumerUin *uint64 `json:"AssumerUin,omitnil,omitempty" name:"AssumerUin"`

	// 云产品标识，日志集由其它云产品创建时，该字段会显示云产品名称，例如CDN、TKE
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// 日志集绑定的标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 日志集下日志主题的数目
	TopicCount *int64 `json:"TopicCount,omitnil,omitempty" name:"TopicCount"`

	// 若AssumerName非空，则表示创建该日志集的服务方角色
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 日志集下指标主题的数目
	MetricTopicCount *int64 `json:"MetricTopicCount,omitnil,omitempty" name:"MetricTopicCount"`
}

type MachineGroupInfo struct {
	// 机器组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 机器组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 机器组类型
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// 创建时间
	// 时间格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 机器组绑定的标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启机器组自动更新
	AutoUpdate *string `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// 机器组中机器离线定期清理时间，单位天，默认设置30天。
	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitnil,omitempty" name:"DelayCleanupTime"`

	// 机器组元数据信息列表
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`

	// 操作系统类型，0: Linux，1: windows
	OSType *uint64 `json:"OSType,omitnil,omitempty" name:"OSType"`
}

type MachineGroupTypeInfo struct {
	// 机器组类型。支持 ip 和 label。
	// - ip：表示该机器组Values中存的是采集机器的ip地址
	// - label：表示该机器组Values中存储的是机器的标签
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 机器描述列表。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type MachineInfo struct {
	// 机器的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 机器实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 机器状态，0:异常，1:正常
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 机器离线时间，空为正常，异常返回具体时间。时间格式：YYYY-MM-DD HH:mm:ss
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// 机器是否开启自动升级。0:关闭，1:开启
	AutoUpdate *int64 `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// 机器当前版本号。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 机器升级功能状态。 0：升级成功；1：升级中；-1：升级失败。
	UpdateStatus *int64 `json:"UpdateStatus,omitnil,omitempty" name:"UpdateStatus"`

	// 机器升级结果标识。
	// 0：成功；1200：升级成功；其他值表示异常。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 机器升级结果信息。
	// “ok”：成功；“update success”：升级成功；其他值为失败原因。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

// Predefined struct for user
type MergePartitionRequestParams struct {
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 合并的PartitionId（找到下一个分区InclusiveBeginKey与入参PartitionId对应的ExclusiveEndKey相等，且找到的分区必须是读写分区（Status:readwrite），入参PartitionId与找到的PartitionId设置为只读分区（Status:readonly）,再新建一个新的读写分区） 。[获取分区列表](https://cloud.tencent.com/document/product/614/56470)
	// 
	// 1. 入参PartitionId只能是读写分区（Status的值有readonly，readwrite），且能找到入参PartitionId的下一个可读写分区（找到下一个分区InclusiveBeginKey与入参PartitionId对应的ExclusiveEndKey相等）；
	// 2. 入参PartitionId不能是最后一个分区（PartitionId的ExclusiveEndKey不能是ffffffffffffffffffffffffffffffff）；
	// 3. topic的分区数量是有限制的（默认50个），合并之后不能超过最大分区，否则不能合并。
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`
}

type MergePartitionRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 合并的PartitionId（找到下一个分区InclusiveBeginKey与入参PartitionId对应的ExclusiveEndKey相等，且找到的分区必须是读写分区（Status:readwrite），入参PartitionId与找到的PartitionId设置为只读分区（Status:readonly）,再新建一个新的读写分区） 。[获取分区列表](https://cloud.tencent.com/document/product/614/56470)
	// 
	// 1. 入参PartitionId只能是读写分区（Status的值有readonly，readwrite），且能找到入参PartitionId的下一个可读写分区（找到下一个分区InclusiveBeginKey与入参PartitionId对应的ExclusiveEndKey相等）；
	// 2. 入参PartitionId不能是最后一个分区（PartitionId的ExclusiveEndKey不能是ffffffffffffffffffffffffffffffff）；
	// 3. topic的分区数量是有限制的（默认50个），合并之后不能超过最大分区，否则不能合并。
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`
}

func (r *MergePartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MergePartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MergePartitionResponseParams struct {
	// 合并结果集
	Partitions []*PartitionInfo `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MergePartitionResponse struct {
	*tchttp.BaseResponse
	Response *MergePartitionResponseParams `json:"Response"`
}

func (r *MergePartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaTagInfo struct {
	// 元数据key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 元数据value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type MetricLabel struct {
	// 指标名称
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 指标内容
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyAlarmNoticeRequestParams struct {
	// 通知渠道组ID。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/api/614/56462)获取通知渠道组ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的通知渠道组。最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 通知渠道组名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 通知类型。可选值：
	// <li> Trigger - 告警触发</li>
	// <li> Recovery - 告警恢复</li>
	// <li> All - 告警触发和告警恢复</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 通知接收对象。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// 接口回调信息（包括企业微信等）。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`

	// 通知规则。
	// 
	// 注意: 
	// 
	// - Type、NoticeReceivers和WebCallbacks是一组配置，NoticeRules是另一组配置，2组配置互斥。
	// - 传其中一组数据，则另一组数据置空。
	NoticeRules []*NoticeRule `json:"NoticeRules,omitnil,omitempty" name:"NoticeRules"`

	// 调用链接域名。http:// 或者 https:// 开头，不能/结尾
	JumpDomain *string `json:"JumpDomain,omitnil,omitempty" name:"JumpDomain"`

	// 投递日志开关。
	// 
	// 参数值：
	// 1：关闭；
	// 
	// 2：开启 
	DeliverStatus *uint64 `json:"DeliverStatus,omitnil,omitempty" name:"DeliverStatus"`

	// 投递日志配置。
	DeliverConfig *DeliverConfig `json:"DeliverConfig,omitnil,omitempty" name:"DeliverConfig"`

	// 免登录操作告警开关。
	// 
	// 参数值： 
	//         1：关闭
	//         2：开启（默认开启）
	AlarmShieldStatus *uint64 `json:"AlarmShieldStatus,omitnil,omitempty" name:"AlarmShieldStatus"`

	// 统一设定自定义回调参数。
	// -  true: 使用通知内容模板中的自定义回调参数覆盖告警策略中单独配置的请求头及请求内容。
	// -  false:优先使用告警策略中单独配置的请求头及请求内容。
	CallbackPrioritize *bool `json:"CallbackPrioritize,omitnil,omitempty" name:"CallbackPrioritize"`
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest
	
	// 通知渠道组ID。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/api/614/56462)获取通知渠道组ID
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的通知渠道组。最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 通知渠道组名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 通知类型。可选值：
	// <li> Trigger - 告警触发</li>
	// <li> Recovery - 告警恢复</li>
	// <li> All - 告警触发和告警恢复</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 通知接收对象。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// 接口回调信息（包括企业微信等）。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`

	// 通知规则。
	// 
	// 注意: 
	// 
	// - Type、NoticeReceivers和WebCallbacks是一组配置，NoticeRules是另一组配置，2组配置互斥。
	// - 传其中一组数据，则另一组数据置空。
	NoticeRules []*NoticeRule `json:"NoticeRules,omitnil,omitempty" name:"NoticeRules"`

	// 调用链接域名。http:// 或者 https:// 开头，不能/结尾
	JumpDomain *string `json:"JumpDomain,omitnil,omitempty" name:"JumpDomain"`

	// 投递日志开关。
	// 
	// 参数值：
	// 1：关闭；
	// 
	// 2：开启 
	DeliverStatus *uint64 `json:"DeliverStatus,omitnil,omitempty" name:"DeliverStatus"`

	// 投递日志配置。
	DeliverConfig *DeliverConfig `json:"DeliverConfig,omitnil,omitempty" name:"DeliverConfig"`

	// 免登录操作告警开关。
	// 
	// 参数值： 
	//         1：关闭
	//         2：开启（默认开启）
	AlarmShieldStatus *uint64 `json:"AlarmShieldStatus,omitnil,omitempty" name:"AlarmShieldStatus"`

	// 统一设定自定义回调参数。
	// -  true: 使用通知内容模板中的自定义回调参数覆盖告警策略中单独配置的请求头及请求内容。
	// -  false:优先使用告警策略中单独配置的请求头及请求内容。
	CallbackPrioritize *bool `json:"CallbackPrioritize,omitnil,omitempty" name:"CallbackPrioritize"`
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
	delete(f, "AlarmNoticeId")
	delete(f, "Tags")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	delete(f, "NoticeRules")
	delete(f, "JumpDomain")
	delete(f, "DeliverStatus")
	delete(f, "DeliverConfig")
	delete(f, "AlarmShieldStatus")
	delete(f, "CallbackPrioritize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmNoticeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyAlarmRequestParams struct {
	// 告警策略ID。-通过[获取告警策略列表](https://cloud.tencent.com/document/product/614/56461)获取告警策略ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 告警策略名称。最大支持255个字节，不支持 '|'。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// 告警信息发送的触发条件。
	// 
	// 注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 告警级别。
	// 
	// 0:警告(Warn);1:提醒(Info);2:紧急 (Critical)
	// 
	// 注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 多触发条件。 
	// 
	// 注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	MultiConditions []*MultiCondition `json:"MultiConditions,omitnil,omitempty" name:"MultiConditions"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为2000。
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// 告警重复的周期。单位是分钟。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知渠道组列表。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/product/614/56462)获取关联的告警通知渠道组列表，和MonitorNotice互斥
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// 监控对象列表。
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// 是否开启告警策略。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 该参数已废弃，请使用Status参数控制是否开启告警策略。
	//
	// Deprecated: Enable is deprecated.
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 用户自定义告警内容
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// 用户自定义回调
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// 多维分析
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// 分组触发状态。true：开启，false：关闭（默认）
	GroupTriggerStatus *bool `json:"GroupTriggerStatus,omitnil,omitempty" name:"GroupTriggerStatus"`

	// 分组触发条件。
	GroupTriggerCondition []*string `json:"GroupTriggerCondition,omitnil,omitempty" name:"GroupTriggerCondition"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的告警策略。最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 监控对象类型。0:执行语句共用监控对象; 1:每个执行语句单独选择监控对象。 
	// 当值为1时，AlarmTargets元素个数不能超过10个，AlarmTargets中的Number必须是从1开始的连续正整数，不能重复。
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 告警附加分类信息列表。
	// Classifications元素个数不能超过20个。
	// Classifications元素的Key不能为空，不能重复，长度不能超过50个字符，符合正则 `^[a-z]([a-z0-9_]{0,49})$`。
	// Classifications元素的Value长度不能超过200个字符。
	Classifications []*AlarmClassification `json:"Classifications,omitnil,omitempty" name:"Classifications"`
}

type ModifyAlarmRequest struct {
	*tchttp.BaseRequest
	
	// 告警策略ID。-通过[获取告警策略列表](https://cloud.tencent.com/document/product/614/56461)获取告警策略ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 告警策略名称。最大支持255个字节，不支持 '|'。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitnil,omitempty" name:"MonitorTime"`

	// 告警信息发送的触发条件。
	// 
	// 注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 告警级别。
	// 
	// 0:警告(Warn);1:提醒(Info);2:紧急 (Critical)
	// 
	// 注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 多触发条件。 
	// 
	// 注意:  
	// - Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。
	MultiConditions []*MultiCondition `json:"MultiConditions,omitnil,omitempty" name:"MultiConditions"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为2000。
	TriggerCount *int64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// 告警重复的周期。单位是分钟。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitnil,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知渠道组列表。-通过[获取通知渠道组列表](https://cloud.tencent.com/document/product/614/56462)获取关联的告警通知渠道组列表，和MonitorNotice互斥
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitnil,omitempty" name:"AlarmNoticeIds"`

	// 监控对象列表。
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitnil,omitempty" name:"AlarmTargets"`

	// 是否开启告警策略。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 该参数已废弃，请使用Status参数控制是否开启告警策略。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 用户自定义告警内容
	MessageTemplate *string `json:"MessageTemplate,omitnil,omitempty" name:"MessageTemplate"`

	// 用户自定义回调
	CallBack *CallBackInfo `json:"CallBack,omitnil,omitempty" name:"CallBack"`

	// 多维分析
	Analysis []*AnalysisDimensional `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// 分组触发状态。true：开启，false：关闭（默认）
	GroupTriggerStatus *bool `json:"GroupTriggerStatus,omitnil,omitempty" name:"GroupTriggerStatus"`

	// 分组触发条件。
	GroupTriggerCondition []*string `json:"GroupTriggerCondition,omitnil,omitempty" name:"GroupTriggerCondition"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的告警策略。最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 监控对象类型。0:执行语句共用监控对象; 1:每个执行语句单独选择监控对象。 
	// 当值为1时，AlarmTargets元素个数不能超过10个，AlarmTargets中的Number必须是从1开始的连续正整数，不能重复。
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 告警附加分类信息列表。
	// Classifications元素个数不能超过20个。
	// Classifications元素的Key不能为空，不能重复，长度不能超过50个字符，符合正则 `^[a-z]([a-z0-9_]{0,49})$`。
	// Classifications元素的Value长度不能超过200个字符。
	Classifications []*AlarmClassification `json:"Classifications,omitnil,omitempty" name:"Classifications"`
}

func (r *ModifyAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	delete(f, "Name")
	delete(f, "MonitorTime")
	delete(f, "Condition")
	delete(f, "AlarmLevel")
	delete(f, "MultiConditions")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "AlarmTargets")
	delete(f, "Status")
	delete(f, "Enable")
	delete(f, "MessageTemplate")
	delete(f, "CallBack")
	delete(f, "Analysis")
	delete(f, "GroupTriggerStatus")
	delete(f, "GroupTriggerCondition")
	delete(f, "Tags")
	delete(f, "MonitorObjectType")
	delete(f, "Classifications")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmResponseParams `json:"Response"`
}

func (r *ModifyAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmShieldRequestParams struct {
	// 屏蔽规则ID。-通过[获取告警屏蔽配置规则](https://cloud.tencent.com/document/api/614/103650)获取屏蔽规则ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 通知渠道组id。-通过[获取告警屏蔽配置规则](https://cloud.tencent.com/document/api/614/103650)获取通知渠道组id
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// 屏蔽开始时间，秒级(s)时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 屏蔽结束时间，秒级(s)时间戳。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 屏蔽类型。1：屏蔽所有通知，2：按照Rule参数屏蔽匹配规则的通知。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 屏蔽规则，当Type为2时必填。规则填写方式详见[产品文档](https://cloud.tencent.com/document/product/614/103178#rule)。
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 屏蔽原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 规则状态。只有规则状态为生效中（status:1）时，才能将其修改为已失效（status:2）。
	// 枚举：0（未生效），1（生效中），2（已失效）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyAlarmShieldRequest struct {
	*tchttp.BaseRequest
	
	// 屏蔽规则ID。-通过[获取告警屏蔽配置规则](https://cloud.tencent.com/document/api/614/103650)获取屏蔽规则ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 通知渠道组id。-通过[获取告警屏蔽配置规则](https://cloud.tencent.com/document/api/614/103650)获取通知渠道组id
	AlarmNoticeId *string `json:"AlarmNoticeId,omitnil,omitempty" name:"AlarmNoticeId"`

	// 屏蔽开始时间，秒级(s)时间戳。
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 屏蔽结束时间，秒级(s)时间戳。
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 屏蔽类型。1：屏蔽所有通知，2：按照Rule参数屏蔽匹配规则的通知。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 屏蔽规则，当Type为2时必填。规则填写方式详见[产品文档](https://cloud.tencent.com/document/product/614/103178#rule)。
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 屏蔽原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 规则状态。只有规则状态为生效中（status:1）时，才能将其修改为已失效（status:2）。
	// 枚举：0（未生效），1（生效中），2（已失效）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyAlarmShieldRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmShieldRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "AlarmNoticeId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Rule")
	delete(f, "Reason")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmShieldRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmShieldResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmShieldResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmShieldResponseParams `json:"Response"`
}

func (r *ModifyAlarmShieldResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmShieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudProductLogCollectionRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云产品标识，支持枚举：CDS、CWP、CDB、TDSQL-C、MongoDB、TDStore、DCDB、MariaDB、PostgreSQL、BH、APIS
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// 日志类型，支持枚举：CDS-AUDIT、CDS-RISK、CDB-AUDIT、TDSQL-C-AUDIT、MongoDB-AUDIT、MongoDB-SlowLog、MongoDB-ErrorLog、TDMYSQL-SLOW、DCDB-AUDIT、DCDB-SLOW、DCDB-ERROR、MariaDB-AUDIT、MariaDB-SLOW、MariaDB-ERROR、PostgreSQL-SLOW、PostgreSQL-ERROR、PostgreSQL-AUDIT、BH-FILELOG、BH-COMMANDLOG、APIS-ACCESS
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 云产品地域。 不同日志类型(LogType)地域入參格式存在差异， 请参考如下示例：
	// - CDS所有日志类型：ap-guangzhou
	// - CDB-AUDIT: gz
	// - TDSQL-C-AUDIT: gz
	// - MongoDB-AUDIT: gz
	// - MongoDB-SlowLog：ap-guangzhou
	// - MongoDB-ErrorLog：ap-guangzhou
	// - TDMYSQL-SLOW：gz
	// - DCDB所有日志类型：gz
	// - MariaDB所有日志类型：gz
	// - PostgreSQL所有日志类型：gz
	// - BH所有日志类型：overseas-polaris(中国香港地区和其他)/fsi-polaris(金融区)/general-polaris(普通区)/intl-sg-prod(国际站)
	// - APIS所有日志类型：gz
	CloudProductRegion *string `json:"CloudProductRegion,omitnil,omitempty" name:"CloudProductRegion"`

	// 日志配置拓展信息， 一般用于存储额外的日志投递配置
	Extend *string `json:"Extend,omitnil,omitempty" name:"Extend"`
}

type ModifyCloudProductLogCollectionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 云产品标识，支持枚举：CDS、CWP、CDB、TDSQL-C、MongoDB、TDStore、DCDB、MariaDB、PostgreSQL、BH、APIS
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// 日志类型，支持枚举：CDS-AUDIT、CDS-RISK、CDB-AUDIT、TDSQL-C-AUDIT、MongoDB-AUDIT、MongoDB-SlowLog、MongoDB-ErrorLog、TDMYSQL-SLOW、DCDB-AUDIT、DCDB-SLOW、DCDB-ERROR、MariaDB-AUDIT、MariaDB-SLOW、MariaDB-ERROR、PostgreSQL-SLOW、PostgreSQL-ERROR、PostgreSQL-AUDIT、BH-FILELOG、BH-COMMANDLOG、APIS-ACCESS
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 云产品地域。 不同日志类型(LogType)地域入參格式存在差异， 请参考如下示例：
	// - CDS所有日志类型：ap-guangzhou
	// - CDB-AUDIT: gz
	// - TDSQL-C-AUDIT: gz
	// - MongoDB-AUDIT: gz
	// - MongoDB-SlowLog：ap-guangzhou
	// - MongoDB-ErrorLog：ap-guangzhou
	// - TDMYSQL-SLOW：gz
	// - DCDB所有日志类型：gz
	// - MariaDB所有日志类型：gz
	// - PostgreSQL所有日志类型：gz
	// - BH所有日志类型：overseas-polaris(中国香港地区和其他)/fsi-polaris(金融区)/general-polaris(普通区)/intl-sg-prod(国际站)
	// - APIS所有日志类型：gz
	CloudProductRegion *string `json:"CloudProductRegion,omitnil,omitempty" name:"CloudProductRegion"`

	// 日志配置拓展信息， 一般用于存储额外的日志投递配置
	Extend *string `json:"Extend,omitnil,omitempty" name:"Extend"`
}

func (r *ModifyCloudProductLogCollectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudProductLogCollectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AssumerName")
	delete(f, "LogType")
	delete(f, "CloudProductRegion")
	delete(f, "Extend")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudProductLogCollectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudProductLogCollectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudProductLogCollectionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudProductLogCollectionResponseParams `json:"Response"`
}

func (r *ModifyCloudProductLogCollectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudProductLogCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigExtraRequestParams struct {
	// 采集配置扩展信息id
	// - 通过[获取特殊采集配置](https://cloud.tencent.com/document/api/614/71164)获取采集配置扩展信息id。
	ConfigExtraId *string `json:"ConfigExtraId,omitnil,omitempty" name:"ConfigExtraId"`

	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志主题id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 自建k8s-节点文件配置信息,包括文件路径、名称及元数据相关信息。
	// 
	// - 详情参考  [HostFileInfo](https://cloud.tencent.com/document/api/614/56471#HostFileInfo) 文档。
	HostFile *HostFileInfo `json:"HostFile,omitnil,omitempty" name:"HostFile"`

	// 采集配置标记。
	// - 目前只支持label_k8s，用于标记自建k8s集群使用的采集配置
	// - 详情参考 [ ContainerFileInfo](https://cloud.tencent.com/document/api/614/56471#ContainerFileInfo) 文档
	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitnil,omitempty" name:"ContainerFile"`

	// 自建k8s-容器标准输出信息，包括容器、命名空间等，
	// 
	// - 详情参考 [ContainerStdoutInfo]( https://cloud.tencent.com/document/api/614/56471#ContainerStdoutInfo) 文档
	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitnil,omitempty" name:"ContainerStdout"`

	// 采集的日志类型，默认为minimalist_log。支持以下类型：
	// - json_log代表：JSON-文件日志（详见[使用 JSON 提取模式采集日志](https://cloud.tencent.com/document/product/614/17419)）；
	// - delimiter_log代表：分隔符-文件日志（详见[使用分隔符提取模式采集日志](https://cloud.tencent.com/document/product/614/17420)）；
	// - minimalist_log代表：单行全文-文件日志（详见[使用单行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17421)）；
	// - fullregex_log代表：单行完全正则-文件日志（详见[使用单行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52365)）；
	// - multiline_log代表：多行全文-文件日志（详见[使用多行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17422)）；
	// - multiline_fullregex_log代表：多行完全正则-文件日志（详见[使用多行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52366)）；
	// - user_define_log代表：组合解析（适用于多格式嵌套的日志，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)）。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 日志格式化方式，用于容器采集场景。
	// - stdout-docker-json：用于docker容器采集场景
	// - stdout-containerd：用于containerd容器采集场景
	//
	// Deprecated: LogFormat is deprecated.
	LogFormat *string `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType。
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 组合解析采集规则，用于复杂场景下的日志采集。
	// - 取值参考：[使用组合解析提取模式采集日志
	// ](https://cloud.tencent.com/document/product/614/61310)
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 容器场景，日志采集输入类型，支持以下三种类型
	// - container_stdout 标准输出
	// - container_file 容器文件
	// - host_file 主机节点文件
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 自建采集配置标
	ConfigFlag *string `json:"ConfigFlag,omitnil,omitempty" name:"ConfigFlag"`

	// 日志集ID
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名称
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集名称。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题名称
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// - ClsAgentDefault(自定义默认值，无特殊含义，用于清空其他选项)，建议取值0
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type ModifyConfigExtraRequest struct {
	*tchttp.BaseRequest
	
	// 采集配置扩展信息id
	// - 通过[获取特殊采集配置](https://cloud.tencent.com/document/api/614/71164)获取采集配置扩展信息id。
	ConfigExtraId *string `json:"ConfigExtraId,omitnil,omitempty" name:"ConfigExtraId"`

	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志主题id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 自建k8s-节点文件配置信息,包括文件路径、名称及元数据相关信息。
	// 
	// - 详情参考  [HostFileInfo](https://cloud.tencent.com/document/api/614/56471#HostFileInfo) 文档。
	HostFile *HostFileInfo `json:"HostFile,omitnil,omitempty" name:"HostFile"`

	// 采集配置标记。
	// - 目前只支持label_k8s，用于标记自建k8s集群使用的采集配置
	// - 详情参考 [ ContainerFileInfo](https://cloud.tencent.com/document/api/614/56471#ContainerFileInfo) 文档
	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitnil,omitempty" name:"ContainerFile"`

	// 自建k8s-容器标准输出信息，包括容器、命名空间等，
	// 
	// - 详情参考 [ContainerStdoutInfo]( https://cloud.tencent.com/document/api/614/56471#ContainerStdoutInfo) 文档
	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitnil,omitempty" name:"ContainerStdout"`

	// 采集的日志类型，默认为minimalist_log。支持以下类型：
	// - json_log代表：JSON-文件日志（详见[使用 JSON 提取模式采集日志](https://cloud.tencent.com/document/product/614/17419)）；
	// - delimiter_log代表：分隔符-文件日志（详见[使用分隔符提取模式采集日志](https://cloud.tencent.com/document/product/614/17420)）；
	// - minimalist_log代表：单行全文-文件日志（详见[使用单行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17421)）；
	// - fullregex_log代表：单行完全正则-文件日志（详见[使用单行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52365)）；
	// - multiline_log代表：多行全文-文件日志（详见[使用多行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17422)）；
	// - multiline_fullregex_log代表：多行完全正则-文件日志（详见[使用多行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52366)）；
	// - user_define_log代表：组合解析（适用于多格式嵌套的日志，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)）。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 日志格式化方式，用于容器采集场景。
	// - stdout-docker-json：用于docker容器采集场景
	// - stdout-containerd：用于containerd容器采集场景
	LogFormat *string `json:"LogFormat,omitnil,omitempty" name:"LogFormat"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType。
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 组合解析采集规则，用于复杂场景下的日志采集。
	// - 取值参考：[使用组合解析提取模式采集日志
	// ](https://cloud.tencent.com/document/product/614/61310)
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 容器场景，日志采集输入类型，支持以下三种类型
	// - container_stdout 标准输出
	// - container_file 容器文件
	// - host_file 主机节点文件
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 机器组ID
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/api/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 自建采集配置标
	ConfigFlag *string `json:"ConfigFlag,omitnil,omitempty" name:"ConfigFlag"`

	// 日志集ID
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名称
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/api/614/58624)获取日志集名称。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题名称
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// - ClsAgentDefault(自定义默认值，无特殊含义，用于清空其他选项)，建议取值0
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

func (r *ModifyConfigExtraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigExtraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigExtraId")
	delete(f, "Name")
	delete(f, "TopicId")
	delete(f, "HostFile")
	delete(f, "ContainerFile")
	delete(f, "ContainerStdout")
	delete(f, "LogType")
	delete(f, "LogFormat")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "UserDefineRule")
	delete(f, "Type")
	delete(f, "GroupId")
	delete(f, "ConfigFlag")
	delete(f, "LogsetId")
	delete(f, "LogsetName")
	delete(f, "TopicName")
	delete(f, "AdvancedConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigExtraResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConfigExtraResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConfigExtraResponseParams `json:"Response"`
}

func (r *ModifyConfigExtraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigRequestParams struct {
	// 采集规则配置ID，通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)返回信息获取。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 采集规则配置名称
	// - 不能包含特殊字符｜
	// - 长度不能超过255字符，超过会被截断
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志采集路径，包含文件名
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 采集的日志类型。支持以下类型：
	// - json_log代表：JSON-文件日志（详见[使用 JSON 提取模式采集日志](https://cloud.tencent.com/document/product/614/17419)）；
	// - delimiter_log代表：分隔符-文件日志（详见[使用分隔符提取模式采集日志](https://cloud.tencent.com/document/product/614/17420)）；
	// - minimalist_log代表：单行全文-文件日志（详见[使用单行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17421)）；
	// - fullregex_log代表：单行完全正则-文件日志（详见[使用单行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52365)）；
	// - multiline_log代表：多行全文-文件日志（详见[使用多行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17422)）；
	// - multiline_fullregex_log代表：多行完全正则-文件日志（详见[使用多行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52366)）；
	// - user_define_log代表：组合解析（适用于多格式嵌套的日志，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)）；
	// - service_syslog代表：syslog 采集（详见[采集 Syslog](https://cloud.tencent.com/document/product/614/81454)）；
	// - windows_event_log代表：Windows事件日志（详见[采集 Windows 事件日志](https://cloud.tencent.com/document/product/614/96678)）。
	// 
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 采集配置关联的日志主题（TopicId）
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 用户自定义解析字符串，Json格式序列化的字符串。
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// 样例：
	// `{\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true}`
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`

	// 日志输入类型（<span style="color:red; font-weight:bold">注：windows场景必填且仅支持file和windows_event类型</span>）
	// - file: 文件类型采集
	// - windows_event：windows事件采集
	// - syslog：系统日志采集
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`
}

type ModifyConfigRequest struct {
	*tchttp.BaseRequest
	
	// 采集规则配置ID，通过[获取采集规则配置](https://cloud.tencent.com/document/product/614/58616)返回信息获取。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 采集规则配置名称
	// - 不能包含特殊字符｜
	// - 长度不能超过255字符，超过会被截断
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志采集路径，包含文件名
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 采集的日志类型。支持以下类型：
	// - json_log代表：JSON-文件日志（详见[使用 JSON 提取模式采集日志](https://cloud.tencent.com/document/product/614/17419)）；
	// - delimiter_log代表：分隔符-文件日志（详见[使用分隔符提取模式采集日志](https://cloud.tencent.com/document/product/614/17420)）；
	// - minimalist_log代表：单行全文-文件日志（详见[使用单行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17421)）；
	// - fullregex_log代表：单行完全正则-文件日志（详见[使用单行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52365)）；
	// - multiline_log代表：多行全文-文件日志（详见[使用多行全文提取模式采集日志](https://cloud.tencent.com/document/product/614/17422)）；
	// - multiline_fullregex_log代表：多行完全正则-文件日志（详见[使用多行-完全正则提取模式采集日志](https://cloud.tencent.com/document/product/614/52366)）；
	// - user_define_log代表：组合解析（适用于多格式嵌套的日志，详见[使用组合解析提取模式采集日志](https://cloud.tencent.com/document/product/614/61310)）；
	// - service_syslog代表：syslog 采集（详见[采集 Syslog](https://cloud.tencent.com/document/product/614/81454)）；
	// - windows_event_log代表：Windows事件日志（详见[采集 Windows 事件日志](https://cloud.tencent.com/document/product/614/96678)）。
	// 
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitnil,omitempty" name:"ExtractRule"`

	// 采集黑名单路径列表
	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitnil,omitempty" name:"ExcludePaths"`

	// 采集配置关联的日志主题（TopicId）
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 用户自定义解析字符串，Json格式序列化的字符串。
	UserDefineRule *string `json:"UserDefineRule,omitnil,omitempty" name:"UserDefineRule"`

	// 高级采集配置。 Json字符串， Key/Value定义为如下：
	// - ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时
	// - ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数
	// - ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false
	// 样例：
	// `{\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true}`
	AdvancedConfig *string `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`

	// 日志输入类型（<span style="color:red; font-weight:bold">注：windows场景必填且仅支持file和windows_event类型</span>）
	// - file: 文件类型采集
	// - windows_event：windows事件采集
	// - syslog：系统日志采集
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`
}

func (r *ModifyConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "Name")
	delete(f, "Path")
	delete(f, "LogType")
	delete(f, "ExtractRule")
	delete(f, "ExcludePaths")
	delete(f, "Output")
	delete(f, "UserDefineRule")
	delete(f, "AdvancedConfig")
	delete(f, "InputType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConfigResponseParams `json:"Response"`
}

func (r *ModifyConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsoleSharingRequestParams struct {
	// 免密分享Id。
	// - 通过 [获取免密分享列表](https://cloud.tencent.com/document/product/614/109798) 获取免密分享Id。 
	// - 通过 [创建免密分享](https://cloud.tencent.com/document/product/614/109800) 获取免密分享Id。
	SharingId *string `json:"SharingId,omitnil,omitempty" name:"SharingId"`

	// 指定分享链接有效期，单位：毫秒，最长可设定有效期为30天
	DurationMilliseconds *uint64 `json:"DurationMilliseconds,omitnil,omitempty" name:"DurationMilliseconds"`
}

type ModifyConsoleSharingRequest struct {
	*tchttp.BaseRequest
	
	// 免密分享Id。
	// - 通过 [获取免密分享列表](https://cloud.tencent.com/document/product/614/109798) 获取免密分享Id。 
	// - 通过 [创建免密分享](https://cloud.tencent.com/document/product/614/109800) 获取免密分享Id。
	SharingId *string `json:"SharingId,omitnil,omitempty" name:"SharingId"`

	// 指定分享链接有效期，单位：毫秒，最长可设定有效期为30天
	DurationMilliseconds *uint64 `json:"DurationMilliseconds,omitnil,omitempty" name:"DurationMilliseconds"`
}

func (r *ModifyConsoleSharingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsoleSharingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SharingId")
	delete(f, "DurationMilliseconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConsoleSharingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsoleSharingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConsoleSharingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConsoleSharingResponseParams `json:"Response"`
}

func (r *ModifyConsoleSharingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsoleSharingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsumerRequestParams struct {
	// 投递任务绑定的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 投递任务是否生效，默认不生效
	Effective *bool `json:"Effective,omitnil,omitempty" name:"Effective"`

	// 是否投递日志的元数据信息，默认为 true。
	// 当NeedContent为true时：字段Content有效。
	// 当NeedContent为false时：字段Content无效。
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// 如果需要投递元数据信息，元数据信息的描述
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka的描述
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// 投递时压缩方式，取值0，2，3。[0：NONE；2：SNAPPY；3：LZ4]
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// 角色访问描述名 [创建角色](https://cloud.tencent.com/document/product/598/19381)
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 外部ID
	ExternalId *string `json:"ExternalId,omitnil,omitempty" name:"ExternalId"`

	// 高级配置
	AdvancedConfig *AdvancedConsumerConfiguration `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type ModifyConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 投递任务绑定的日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 投递任务是否生效，默认不生效
	Effective *bool `json:"Effective,omitnil,omitempty" name:"Effective"`

	// 是否投递日志的元数据信息，默认为 true。
	// 当NeedContent为true时：字段Content有效。
	// 当NeedContent为false时：字段Content无效。
	NeedContent *bool `json:"NeedContent,omitnil,omitempty" name:"NeedContent"`

	// 如果需要投递元数据信息，元数据信息的描述
	Content *ConsumerContent `json:"Content,omitnil,omitempty" name:"Content"`

	// CKafka的描述
	Ckafka *Ckafka `json:"Ckafka,omitnil,omitempty" name:"Ckafka"`

	// 投递时压缩方式，取值0，2，3。[0：NONE；2：SNAPPY；3：LZ4]
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// 角色访问描述名 [创建角色](https://cloud.tencent.com/document/product/598/19381)
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 外部ID
	ExternalId *string `json:"ExternalId,omitnil,omitempty" name:"ExternalId"`

	// 高级配置
	AdvancedConfig *AdvancedConsumerConfiguration `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

func (r *ModifyConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Effective")
	delete(f, "NeedContent")
	delete(f, "Content")
	delete(f, "Ckafka")
	delete(f, "Compression")
	delete(f, "RoleArn")
	delete(f, "ExternalId")
	delete(f, "AdvancedConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsumerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConsumerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConsumerResponseParams `json:"Response"`
}

func (r *ModifyConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCosRechargeRequestParams struct {
	// COS导入配置Id。
	// 
	// - 通过[获取cos导入配置](https://cloud.tencent.com/document/product/614/88099)接口获取COS导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 日志主题Id。
	// 
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// COS导入任务名称,最大支持128个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务状态   0： 停用 ， 1：启用
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// COS存储桶所在地域，详见产品支持的[地域列表](https://cloud.tencent.com/document/product/436/6224)。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// COS文件所在文件夹的前缀。为空串时投递存储桶下所有的文件。
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表单行全文； 默认为minimalist_log
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 解析格式。supported: "", "gzip", "lzop", "snappy"。空串表示不压缩。
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitnil,omitempty" name:"ExtractRuleInfo"`

	// COS导入任务类型。1：一次性导入任务；2：持续性导入任务。
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 元数据。支持 bucket，object。
	Metadata []*string `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type ModifyCosRechargeRequest struct {
	*tchttp.BaseRequest
	
	// COS导入配置Id。
	// 
	// - 通过[获取cos导入配置](https://cloud.tencent.com/document/product/614/88099)接口获取COS导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 日志主题Id。
	// 
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// COS导入任务名称,最大支持128个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务状态   0： 停用 ， 1：启用
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// COS存储桶所在地域，详见产品支持的[地域列表](https://cloud.tencent.com/document/product/436/6224)。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// COS文件所在文件夹的前缀。为空串时投递存储桶下所有的文件。
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表单行全文； 默认为minimalist_log
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 解析格式。supported: "", "gzip", "lzop", "snappy"。空串表示不压缩。
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 提取规则，如果设置了ExtractRule，则必须设置LogType
	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitnil,omitempty" name:"ExtractRuleInfo"`

	// COS导入任务类型。1：一次性导入任务；2：持续性导入任务。
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 元数据。支持 bucket，object。
	Metadata []*string `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

func (r *ModifyCosRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCosRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TopicId")
	delete(f, "Name")
	delete(f, "Enable")
	delete(f, "Bucket")
	delete(f, "BucketRegion")
	delete(f, "Prefix")
	delete(f, "LogType")
	delete(f, "Compress")
	delete(f, "ExtractRuleInfo")
	delete(f, "TaskType")
	delete(f, "Metadata")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCosRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCosRechargeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCosRechargeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCosRechargeResponseParams `json:"Response"`
}

func (r *ModifyCosRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDashboardSubscribeRequestParams struct {
	// 仪表盘订阅id。通过 [获取仪表盘订阅列表](https://cloud.tencent.com/document/api/614/105779)接口获取Id。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 仪表盘id。通过 [获取仪表盘](https://cloud.tencent.com/document/api/614/95636)接口获取DashboardId。
	DashboardId *string `json:"DashboardId,omitnil,omitempty" name:"DashboardId"`

	// 仪表盘订阅名称。最大支持128个字符，且不支持'|'字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 订阅时间cron表达式，格式为：{秒数} {分钟} {小时} {日期} {月份} {星期}；（有效数据为：{分钟} {小时} {日期} {月份} {星期}）。
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 仪表盘订阅数据。
	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitnil,omitempty" name:"SubscribeData"`
}

type ModifyDashboardSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 仪表盘订阅id。通过 [获取仪表盘订阅列表](https://cloud.tencent.com/document/api/614/105779)接口获取Id。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 仪表盘id。通过 [获取仪表盘](https://cloud.tencent.com/document/api/614/95636)接口获取DashboardId。
	DashboardId *string `json:"DashboardId,omitnil,omitempty" name:"DashboardId"`

	// 仪表盘订阅名称。最大支持128个字符，且不支持'|'字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 订阅时间cron表达式，格式为：{秒数} {分钟} {小时} {日期} {月份} {星期}；（有效数据为：{分钟} {小时} {日期} {月份} {星期}）。
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 仪表盘订阅数据。
	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitnil,omitempty" name:"SubscribeData"`
}

func (r *ModifyDashboardSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDashboardSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "DashboardId")
	delete(f, "Name")
	delete(f, "Cron")
	delete(f, "SubscribeData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDashboardSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDashboardSubscribeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDashboardSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDashboardSubscribeResponseParams `json:"Response"`
}

func (r *ModifyDashboardSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDashboardSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataTransformRequestParams struct {
	// 数据加工任务ID
	// - 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 加工任务名称
	// - 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务名称。
	// 
	// 名称限制
	// - 不能为空字符串
	// - 不能包含字符'|'
	// - 最长 255 个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 加工语句。 当FuncType为2时，EtlContent必须使用[log_auto_output](https://cloud.tencent.com/document/product/614/70733#b3c58797-4825-4807-bef4-68106e25024f) 
	// 
	// 其他参考文档：
	// 
	// - [创建加工任务](https://cloud.tencent.com/document/product/614/63940) 
	// -  [函数总览](https://cloud.tencent.com/document/product/614/70395)
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// 任务启动状态. 默认为1，开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 加工任务目的topic_id以及别名
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// 超限之后是否丢弃日志数据
	BackupGiveUpData *bool `json:"BackupGiveUpData,omitnil,omitempty" name:"BackupGiveUpData"`

	// 是否开启投递服务日志。1关闭，2开启
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`

	// 保留失败日志状态。 1:不保留，2:保留
	KeepFailureLog *uint64 `json:"KeepFailureLog,omitnil,omitempty" name:"KeepFailureLog"`

	// 失败日志的字段名称
	FailureLogKey *string `json:"FailureLogKey,omitnil,omitempty" name:"FailureLogKey"`

	// 外部数据源信息
	DataTransformSqlDataSources []*DataTransformSqlDataSource `json:"DataTransformSqlDataSources,omitnil,omitempty" name:"DataTransformSqlDataSources"`

	// 设置的环境变量
	EnvInfos []*EnvInfo `json:"EnvInfos,omitnil,omitempty" name:"EnvInfos"`
}

type ModifyDataTransformRequest struct {
	*tchttp.BaseRequest
	
	// 数据加工任务ID
	// - 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 加工任务名称
	// - 通过[获取数据加工任务列表基本信息](https://cloud.tencent.com/document/product/614/72182)获取数据加工任务名称。
	// 
	// 名称限制
	// - 不能为空字符串
	// - 不能包含字符'|'
	// - 最长 255 个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 加工语句。 当FuncType为2时，EtlContent必须使用[log_auto_output](https://cloud.tencent.com/document/product/614/70733#b3c58797-4825-4807-bef4-68106e25024f) 
	// 
	// 其他参考文档：
	// 
	// - [创建加工任务](https://cloud.tencent.com/document/product/614/63940) 
	// -  [函数总览](https://cloud.tencent.com/document/product/614/70395)
	EtlContent *string `json:"EtlContent,omitnil,omitempty" name:"EtlContent"`

	// 任务启动状态. 默认为1，开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 加工任务目的topic_id以及别名
	DstResources []*DataTransformResouceInfo `json:"DstResources,omitnil,omitempty" name:"DstResources"`

	// 超限之后是否丢弃日志数据
	BackupGiveUpData *bool `json:"BackupGiveUpData,omitnil,omitempty" name:"BackupGiveUpData"`

	// 是否开启投递服务日志。1关闭，2开启
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`

	// 保留失败日志状态。 1:不保留，2:保留
	KeepFailureLog *uint64 `json:"KeepFailureLog,omitnil,omitempty" name:"KeepFailureLog"`

	// 失败日志的字段名称
	FailureLogKey *string `json:"FailureLogKey,omitnil,omitempty" name:"FailureLogKey"`

	// 外部数据源信息
	DataTransformSqlDataSources []*DataTransformSqlDataSource `json:"DataTransformSqlDataSources,omitnil,omitempty" name:"DataTransformSqlDataSources"`

	// 设置的环境变量
	EnvInfos []*EnvInfo `json:"EnvInfos,omitnil,omitempty" name:"EnvInfos"`
}

func (r *ModifyDataTransformRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataTransformRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "EtlContent")
	delete(f, "EnableFlag")
	delete(f, "DstResources")
	delete(f, "BackupGiveUpData")
	delete(f, "HasServicesLog")
	delete(f, "KeepFailureLog")
	delete(f, "FailureLogKey")
	delete(f, "DataTransformSqlDataSources")
	delete(f, "EnvInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataTransformRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataTransformResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDataTransformResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDataTransformResponseParams `json:"Response"`
}

func (r *ModifyDataTransformResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDlcDeliverRequestParams struct {
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 任务id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z, _,-,中文字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 投递类型。0:批投递,1:实时投递
	DeliverType *uint64 `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 投递时间范围的开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 投递时间范围的结束时间。 如果为空，则表示不限时
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 投递文件大小,单位MB。 DeliverType=0时必填，范围 5<= MaxSize <= 256。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 投递间隔，单位秒。 DeliverType=0时必填，范围 300<= Interval <=900。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// dlc配置信息
	DlcInfo *DlcInfo `json:"DlcInfo,omitnil,omitempty" name:"DlcInfo"`

	// 是否开启投递服务日志。1关闭，2开启。默认开启
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`

	// 任务状态。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyDlcDeliverRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 任务id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z, _,-,中文字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 投递类型。0:批投递,1:实时投递
	DeliverType *uint64 `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 投递时间范围的开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 投递时间范围的结束时间。 如果为空，则表示不限时
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 投递文件大小,单位MB。 DeliverType=0时必填，范围 5<= MaxSize <= 256。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 投递间隔，单位秒。 DeliverType=0时必填，范围 300<= Interval <=900。
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// dlc配置信息
	DlcInfo *DlcInfo `json:"DlcInfo,omitnil,omitempty" name:"DlcInfo"`

	// 是否开启投递服务日志。1关闭，2开启。默认开启
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`

	// 任务状态。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyDlcDeliverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDlcDeliverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "DeliverType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MaxSize")
	delete(f, "Interval")
	delete(f, "DlcInfo")
	delete(f, "HasServicesLog")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDlcDeliverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDlcDeliverResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDlcDeliverResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDlcDeliverResponseParams `json:"Response"`
}

func (r *ModifyDlcDeliverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDlcDeliverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIndexRequestParams struct {
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 索引状态。false：关闭索引， true：开启索引
	// 开启后可对日志进行检索分析，将产生索引流量、索引存储及相应费用。[费用详情](https://cloud.tencent.com/document/product/614/45802#.E8.AE.A1.E8.B4.B9.E9.A1.B9)
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 索引规则
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 内置保留字段（`__FILENAME__`，`__HOSTNAME__`及`__SOURCE__`）是否包含至全文索引，默认为false，推荐设置为true
	// * false:不包含
	// * true:包含
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// 元数据字段（前缀为`__TAG__`的字段）是否包含至全文索引，默认为0，推荐设置为1
	// * 0:仅包含开启键值索引的元数据字段
	// * 1:包含所有元数据字段
	// * 2:不包含任何元数据字段
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`

	// 自定义日志解析异常存储字段。
	CoverageField *string `json:"CoverageField,omitnil,omitempty" name:"CoverageField"`
}

type ModifyIndexRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 索引状态。false：关闭索引， true：开启索引
	// 开启后可对日志进行检索分析，将产生索引流量、索引存储及相应费用。[费用详情](https://cloud.tencent.com/document/product/614/45802#.E8.AE.A1.E8.B4.B9.E9.A1.B9)
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 索引规则
	Rule *RuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 内置保留字段（`__FILENAME__`，`__HOSTNAME__`及`__SOURCE__`）是否包含至全文索引，默认为false，推荐设置为true
	// * false:不包含
	// * true:包含
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// 元数据字段（前缀为`__TAG__`的字段）是否包含至全文索引，默认为0，推荐设置为1
	// * 0:仅包含开启键值索引的元数据字段
	// * 1:包含所有元数据字段
	// * 2:不包含任何元数据字段
	MetadataFlag *uint64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`

	// 自定义日志解析异常存储字段。
	CoverageField *string `json:"CoverageField,omitnil,omitempty" name:"CoverageField"`
}

func (r *ModifyIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Status")
	delete(f, "Rule")
	delete(f, "IncludeInternalFields")
	delete(f, "MetadataFlag")
	delete(f, "CoverageField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIndexResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIndexResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIndexResponseParams `json:"Response"`
}

func (r *ModifyIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKafkaConsumerGroupOffsetRequestParams struct {

}

type ModifyKafkaConsumerGroupOffsetRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyKafkaConsumerGroupOffsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKafkaConsumerGroupOffsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKafkaConsumerGroupOffsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKafkaConsumerGroupOffsetResponseParams struct {
	// 状态码。0：成功，-1：失败
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKafkaConsumerGroupOffsetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKafkaConsumerGroupOffsetResponseParams `json:"Response"`
}

func (r *ModifyKafkaConsumerGroupOffsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKafkaConsumerGroupOffsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKafkaConsumerRequestParams struct {
	// 日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`

	// 压缩方式。0：不压缩；2：使用Snappy压缩；3：使用LZ4压缩
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// kafka协议消费数据格式
	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitnil,omitempty" name:"ConsumerContent"`
}

type ModifyKafkaConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`

	// 压缩方式。0：不压缩；2：使用Snappy压缩；3：使用LZ4压缩
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// kafka协议消费数据格式
	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitnil,omitempty" name:"ConsumerContent"`
}

func (r *ModifyKafkaConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKafkaConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTopicId")
	delete(f, "Compression")
	delete(f, "ConsumerContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKafkaConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKafkaConsumerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKafkaConsumerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKafkaConsumerResponseParams `json:"Response"`
}

func (r *ModifyKafkaConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKafkaRechargeRequestParams struct {
	// 导入配置Id。
	// - 通过 [创建Kafka数据订阅任务](https://cloud.tencent.com/document/product/614/94448)获取Kafka导入配置Id。
	// - 通过 [获取Kafka数据订阅任务列表](https://cloud.tencent.com/document/product/614/94446)获取Kafka导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 导入CLS目标TopicId。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka导入配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 导入Kafka类型，0：腾讯云CKafka：1：用户自建Kafka。
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// 腾讯云CKafka实例ID，KafkaType为0时必填。
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取实例id。
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// 服务地址，KafkaType为1时必填。
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// ServerAddr是否为加密连接，KafkaType为1时必填。
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// 加密访问协议，KafkaType参数为1并且IsEncryptionAddr参数为true时必填。
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开。
	// 
	// - Kafka类型为腾讯云CKafka时：通过 [获取主题列表](https://cloud.tencent.com/document/product/597/40847) 获取TopicName。
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// 用户Kafka消费组名称
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 日志导入规则
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`

	// 导入控制，1：暂停；2：启动。
	StatusControl *uint64 `json:"StatusControl,omitnil,omitempty" name:"StatusControl"`

	// 用户kafka拓展信息
	UserKafkaMeta *UserKafkaMeta `json:"UserKafkaMeta,omitnil,omitempty" name:"UserKafkaMeta"`
}

type ModifyKafkaRechargeRequest struct {
	*tchttp.BaseRequest
	
	// 导入配置Id。
	// - 通过 [创建Kafka数据订阅任务](https://cloud.tencent.com/document/product/614/94448)获取Kafka导入配置Id。
	// - 通过 [获取Kafka数据订阅任务列表](https://cloud.tencent.com/document/product/614/94446)获取Kafka导入配置Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 导入CLS目标TopicId。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// Kafka导入配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 导入Kafka类型，0：腾讯云CKafka：1：用户自建Kafka。
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// 腾讯云CKafka实例ID，KafkaType为0时必填。
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取实例id。
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// 服务地址，KafkaType为1时必填。
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// ServerAddr是否为加密连接，KafkaType为1时必填。
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// 加密访问协议，KafkaType参数为1并且IsEncryptionAddr参数为true时必填。
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开。
	// 
	// - Kafka类型为腾讯云CKafka时：通过 [获取主题列表](https://cloud.tencent.com/document/product/597/40847) 获取TopicName。
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// 用户Kafka消费组名称
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 日志导入规则
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`

	// 导入控制，1：暂停；2：启动。
	StatusControl *uint64 `json:"StatusControl,omitnil,omitempty" name:"StatusControl"`

	// 用户kafka拓展信息
	UserKafkaMeta *UserKafkaMeta `json:"UserKafkaMeta,omitnil,omitempty" name:"UserKafkaMeta"`
}

func (r *ModifyKafkaRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKafkaRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TopicId")
	delete(f, "Name")
	delete(f, "KafkaType")
	delete(f, "KafkaInstance")
	delete(f, "ServerAddr")
	delete(f, "IsEncryptionAddr")
	delete(f, "Protocol")
	delete(f, "UserKafkaTopics")
	delete(f, "ConsumerGroupName")
	delete(f, "LogRechargeRule")
	delete(f, "StatusControl")
	delete(f, "UserKafkaMeta")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKafkaRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKafkaRechargeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKafkaRechargeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKafkaRechargeResponseParams `json:"Response"`
}

func (r *ModifyKafkaRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogsetRequestParams struct {
	// 日志集Id。通过 [获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名字。- 最大支持255个字符。不支持`|`字符。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志集的绑定的标签键值对。最大支持10个标签键值对，同一个资源只能同时绑定一个标签键。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyLogsetRequest struct {
	*tchttp.BaseRequest
	
	// 日志集Id。通过 [获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志集名字。- 最大支持255个字符。不支持`|`字符。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志集的绑定的标签键值对。最大支持10个标签键值对，同一个资源只能同时绑定一个标签键。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyLogsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "LogsetName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogsetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLogsetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLogsetResponseParams `json:"Response"`
}

func (r *ModifyLogsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMachineGroupRequestParams struct {
	// 机器组Id
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/product/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 机器组名称
	// 输入限制：
	// - 不能为空字符串
	// - 不能包含字符'|'
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 机器组类型。 
	// Type：ip，Values中为ip字符串列表机器组；
	// Type：label，Values中为标签字符串列表机器组。
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启机器组自动更新
	AutoUpdate *bool `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// 机器组中机器定期离线清理时间。单位：天
	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitnil,omitempty" name:"DelayCleanupTime"`

	// 机器组元数据信息列表
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`
}

type ModifyMachineGroupRequest struct {
	*tchttp.BaseRequest
	
	// 机器组Id
	// - 通过[获取机器组列表](https://cloud.tencent.com/document/product/614/56438)获取机器组Id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 机器组名称
	// 输入限制：
	// - 不能为空字符串
	// - 不能包含字符'|'
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 机器组类型。 
	// Type：ip，Values中为ip字符串列表机器组；
	// Type：label，Values中为标签字符串列表机器组。
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitnil,omitempty" name:"MachineGroupType"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启机器组自动更新
	AutoUpdate *bool `json:"AutoUpdate,omitnil,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateStartTime *string `json:"UpdateStartTime,omitnil,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	// 时间格式：HH:mm:ss
	UpdateEndTime *string `json:"UpdateEndTime,omitnil,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费
	ServiceLogging *bool `json:"ServiceLogging,omitnil,omitempty" name:"ServiceLogging"`

	// 机器组中机器定期离线清理时间。单位：天
	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitnil,omitempty" name:"DelayCleanupTime"`

	// 机器组元数据信息列表
	MetaTags []*MetaTagInfo `json:"MetaTags,omitnil,omitempty" name:"MetaTags"`
}

func (r *ModifyMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	delete(f, "DelayCleanupTime")
	delete(f, "MetaTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMachineGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMachineGroupResponseParams `json:"Response"`
}

func (r *ModifyMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNoticeContentRequestParams struct {
	// 通知内容模板ID。-通过[获取通知内容模板](https://cloud.tencent.com/document/api/614/111714)获取通知内容模板ID
	NoticeContentId *string `json:"NoticeContentId,omitnil,omitempty" name:"NoticeContentId"`

	// 通知内容模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 通知内容语言。
	// 
	// 0：中文 1：英文
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 通知内容模板详细信息。
	NoticeContents []*NoticeContent `json:"NoticeContents,omitnil,omitempty" name:"NoticeContents"`
}

type ModifyNoticeContentRequest struct {
	*tchttp.BaseRequest
	
	// 通知内容模板ID。-通过[获取通知内容模板](https://cloud.tencent.com/document/api/614/111714)获取通知内容模板ID
	NoticeContentId *string `json:"NoticeContentId,omitnil,omitempty" name:"NoticeContentId"`

	// 通知内容模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 通知内容语言。
	// 
	// 0：中文 1：英文
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 通知内容模板详细信息。
	NoticeContents []*NoticeContent `json:"NoticeContents,omitnil,omitempty" name:"NoticeContents"`
}

func (r *ModifyNoticeContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNoticeContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NoticeContentId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeContents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNoticeContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNoticeContentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNoticeContentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNoticeContentResponseParams `json:"Response"`
}

func (r *ModifyNoticeContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNoticeContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScheduledSqlRequestParams struct {
	// 任务ID，通过[获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519)获取
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 源日志主题，通过[获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519)获取
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// 任务启动状态.   1开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 定时SQL分析的目标日志主题
	DstResource *ScheduledSqlResouceInfo `json:"DstResource,omitnil,omitempty" name:"DstResource"`

	// 查询语句
	ScheduledSqlContent *string `json:"ScheduledSqlContent,omitnil,omitempty" name:"ScheduledSqlContent"`

	// 调度周期(分钟)，1~1440分钟
	ProcessPeriod *int64 `json:"ProcessPeriod,omitnil,omitempty" name:"ProcessPeriod"`

	// 单次查询的时间窗口. 例子中为近15分钟
	ProcessTimeWindow *string `json:"ProcessTimeWindow,omitnil,omitempty" name:"ProcessTimeWindow"`

	// 执行延迟(秒)，0~120秒，默认60秒
	ProcessDelay *int64 `json:"ProcessDelay,omitnil,omitempty" name:"ProcessDelay"`

	// 源topicId的地域信息,支持地域见 [地域列表](https://cloud.tencent.com/document/api/614/56474#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8) 文档
	SrcTopicRegion *string `json:"SrcTopicRegion,omitnil,omitempty" name:"SrcTopicRegion"`

	// 任务名称，0~255字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 语法规则。 默认值为0。 0：Lucene语法，1：CQL语法
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

type ModifyScheduledSqlRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，通过[获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519)获取
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 源日志主题，通过[获取定时SQL分析任务列表](https://cloud.tencent.com/document/product/614/95519)获取
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// 任务启动状态.   1开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 定时SQL分析的目标日志主题
	DstResource *ScheduledSqlResouceInfo `json:"DstResource,omitnil,omitempty" name:"DstResource"`

	// 查询语句
	ScheduledSqlContent *string `json:"ScheduledSqlContent,omitnil,omitempty" name:"ScheduledSqlContent"`

	// 调度周期(分钟)，1~1440分钟
	ProcessPeriod *int64 `json:"ProcessPeriod,omitnil,omitempty" name:"ProcessPeriod"`

	// 单次查询的时间窗口. 例子中为近15分钟
	ProcessTimeWindow *string `json:"ProcessTimeWindow,omitnil,omitempty" name:"ProcessTimeWindow"`

	// 执行延迟(秒)，0~120秒，默认60秒
	ProcessDelay *int64 `json:"ProcessDelay,omitnil,omitempty" name:"ProcessDelay"`

	// 源topicId的地域信息,支持地域见 [地域列表](https://cloud.tencent.com/document/api/614/56474#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8) 文档
	SrcTopicRegion *string `json:"SrcTopicRegion,omitnil,omitempty" name:"SrcTopicRegion"`

	// 任务名称，0~255字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 语法规则。 默认值为0。 0：Lucene语法，1：CQL语法
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`
}

func (r *ModifyScheduledSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduledSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "SrcTopicId")
	delete(f, "EnableFlag")
	delete(f, "DstResource")
	delete(f, "ScheduledSqlContent")
	delete(f, "ProcessPeriod")
	delete(f, "ProcessTimeWindow")
	delete(f, "ProcessDelay")
	delete(f, "SrcTopicRegion")
	delete(f, "Name")
	delete(f, "SyntaxRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyScheduledSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScheduledSqlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyScheduledSqlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyScheduledSqlResponseParams `json:"Response"`
}

func (r *ModifyScheduledSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduledSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyShipperRequestParams struct {
	// 投递规则Id。
	// 
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745)获取ShipperId。
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 投递规则投递的新的目录前缀。
	// - 仅支持0-9A-Za-z-_/
	// - 最大支持256个字符
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 投递规则的开关状态。true：开启投递任务；false：关闭投递任务。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 投递规则的名字
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// 投递的时间间隔，单位 秒，默认300，范围 300-900
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 投递的文件的最大值，单位 MB，默认256，范围 5-256
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// 投递日志的分区规则，支持strftime的时间格式表示
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 投递日志的压缩配置
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 投递日志的内容格式配置
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// 投递文件命名配置，0：随机数命名，1：投递时间命名。
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`

	// 对象存储类型，默认值为 STANDARD。枚举值请参见[ 存储类型概述](https://cloud.tencent.com/document/product/436/33417) 文档。
	// 参考值有：
	// 
	// - STANDARD：标准存储
	// - STANDARD_IA：低频存储
	// - ARCHIVE：归档存储
	// - DEEP_ARCHIVE：深度归档存储
	// - MAZ_STANDARD：标准存储（多 AZ）
	// - MAZ_STANDARD_IA：低频存储（多 AZ）
	// - INTELLIGENT_TIERING：智能分层存储
	// - MAZ_INTELLIGENT_TIERING：智能分层存储（多 AZ）
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 角色访问描述名 [创建角色](https://cloud.tencent.com/document/product/598/19381)
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 外部ID
	ExternalId *string `json:"ExternalId,omitnil,omitempty" name:"ExternalId"`
}

type ModifyShipperRequest struct {
	*tchttp.BaseRequest
	
	// 投递规则Id。
	// 
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745)获取ShipperId。
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 投递规则投递的新的目录前缀。
	// - 仅支持0-9A-Za-z-_/
	// - 最大支持256个字符
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 投递规则的开关状态。true：开启投递任务；false：关闭投递任务。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 投递规则的名字
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// 投递的时间间隔，单位 秒，默认300，范围 300-900
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 投递的文件的最大值，单位 MB，默认256，范围 5-256
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// 投递日志的分区规则，支持strftime的时间格式表示
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 投递日志的压缩配置
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 投递日志的内容格式配置
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// 投递文件命名配置，0：随机数命名，1：投递时间命名。
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`

	// 对象存储类型，默认值为 STANDARD。枚举值请参见[ 存储类型概述](https://cloud.tencent.com/document/product/436/33417) 文档。
	// 参考值有：
	// 
	// - STANDARD：标准存储
	// - STANDARD_IA：低频存储
	// - ARCHIVE：归档存储
	// - DEEP_ARCHIVE：深度归档存储
	// - MAZ_STANDARD：标准存储（多 AZ）
	// - MAZ_STANDARD_IA：低频存储（多 AZ）
	// - INTELLIGENT_TIERING：智能分层存储
	// - MAZ_INTELLIGENT_TIERING：智能分层存储（多 AZ）
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 角色访问描述名 [创建角色](https://cloud.tencent.com/document/product/598/19381)
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 外部ID
	ExternalId *string `json:"ExternalId,omitnil,omitempty" name:"ExternalId"`
}

func (r *ModifyShipperRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyShipperRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "Bucket")
	delete(f, "Prefix")
	delete(f, "Status")
	delete(f, "ShipperName")
	delete(f, "Interval")
	delete(f, "MaxSize")
	delete(f, "FilterRules")
	delete(f, "Partition")
	delete(f, "Compress")
	delete(f, "Content")
	delete(f, "FilenameMode")
	delete(f, "StorageType")
	delete(f, "RoleArn")
	delete(f, "ExternalId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyShipperRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyShipperResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyShipperResponse struct {
	*tchttp.BaseResponse
	Response *ModifyShipperResponseParams `json:"Response"`
}

func (r *ModifyShipperResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicRequestParams struct {
	//  主题ID- 通过[获取主题列表](https://cloud.tencent.com/document/product/614/56454)获取主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题名称
	// 输入限制：
	// - 不能为空字符串
	// - 不能包含字符'|'
	// - 不能使用以下名称["cls_service_log","loglistener_status","loglistener_alarm","loglistener_business","cls_service_metric"]
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的主题。最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 主题是否开启采集，true：开启采集；false：关闭采集。
	// 控制台目前不支持修改此参数。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否开启自动分裂
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// 若开启最大分裂，该主题能够允许的最大分区数；
	// 默认为50；必须为正数
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// 生命周期，单位天，标准存储取值范围1\~3600，低频存储取值范围7\~3600。取值为3640时代表永久保存
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 存储类型：cold 低频存储，hot 标准存储
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 主题描述
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// 0：日志主题关闭日志沉降。
	// 非0：日志主题开启日志沉降后标准存储的天数。HotPeriod需要大于等于7，且小于Period。
	// 仅在StorageType为 hot 时生效，指标主题不支持该配置。
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`

	// 免鉴权开关。 false：关闭； true：开启。
	// 开启后将支持指定操作匿名访问该日志主题。详情请参见[日志主题](https://cloud.tencent.com/document/product/614/41035)。
	IsWebTracking *bool `json:"IsWebTracking,omitnil,omitempty" name:"IsWebTracking"`

	// 主题扩展信息
	Extends *TopicExtendInfo `json:"Extends,omitnil,omitempty" name:"Extends"`

	// 主题分区数量。
	// 默认为1；
	// 取值范围及约束：
	// - 当输入值<=0，系统自动调整为1。
	// - 如果未传MaxSplitPartitions，需要PartitionCount<=50；
	// - 如果传递了MaxSplitPartitions，需要PartitionCount<=MaxSplitPartitions；
	PartitionCount *uint64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// 取消切换存储任务的id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取取消切换存储任务的id【Topics中的TopicAsyncTaskID字段】。
	CancelTopicAsyncTaskID *string `json:"CancelTopicAsyncTaskID,omitnil,omitempty" name:"CancelTopicAsyncTaskID"`

	// 加密相关参数。 支持加密地域并且开白用户可以传此参数，其他场景不能传递该参数。
	// 只支持传入1：kms-cls 云产品秘钥加密
	Encryption *uint64 `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 开启记录公网来源ip和服务端接收时间
	IsSourceFrom *bool `json:"IsSourceFrom,omitnil,omitempty" name:"IsSourceFrom"`
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest
	
	//  主题ID- 通过[获取主题列表](https://cloud.tencent.com/document/product/614/56454)获取主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题名称
	// 输入限制：
	// - 不能为空字符串
	// - 不能包含字符'|'
	// - 不能使用以下名称["cls_service_log","loglistener_status","loglistener_alarm","loglistener_business","cls_service_metric"]
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的主题。最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 主题是否开启采集，true：开启采集；false：关闭采集。
	// 控制台目前不支持修改此参数。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否开启自动分裂
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// 若开启最大分裂，该主题能够允许的最大分区数；
	// 默认为50；必须为正数
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// 生命周期，单位天，标准存储取值范围1\~3600，低频存储取值范围7\~3600。取值为3640时代表永久保存
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 存储类型：cold 低频存储，hot 标准存储
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 主题描述
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// 0：日志主题关闭日志沉降。
	// 非0：日志主题开启日志沉降后标准存储的天数。HotPeriod需要大于等于7，且小于Period。
	// 仅在StorageType为 hot 时生效，指标主题不支持该配置。
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`

	// 免鉴权开关。 false：关闭； true：开启。
	// 开启后将支持指定操作匿名访问该日志主题。详情请参见[日志主题](https://cloud.tencent.com/document/product/614/41035)。
	IsWebTracking *bool `json:"IsWebTracking,omitnil,omitempty" name:"IsWebTracking"`

	// 主题扩展信息
	Extends *TopicExtendInfo `json:"Extends,omitnil,omitempty" name:"Extends"`

	// 主题分区数量。
	// 默认为1；
	// 取值范围及约束：
	// - 当输入值<=0，系统自动调整为1。
	// - 如果未传MaxSplitPartitions，需要PartitionCount<=50；
	// - 如果传递了MaxSplitPartitions，需要PartitionCount<=MaxSplitPartitions；
	PartitionCount *uint64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// 取消切换存储任务的id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取取消切换存储任务的id【Topics中的TopicAsyncTaskID字段】。
	CancelTopicAsyncTaskID *string `json:"CancelTopicAsyncTaskID,omitnil,omitempty" name:"CancelTopicAsyncTaskID"`

	// 加密相关参数。 支持加密地域并且开白用户可以传此参数，其他场景不能传递该参数。
	// 只支持传入1：kms-cls 云产品秘钥加密
	Encryption *uint64 `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 开启记录公网来源ip和服务端接收时间
	IsSourceFrom *bool `json:"IsSourceFrom,omitnil,omitempty" name:"IsSourceFrom"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "TopicName")
	delete(f, "Tags")
	delete(f, "Status")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	delete(f, "Period")
	delete(f, "StorageType")
	delete(f, "Describes")
	delete(f, "HotPeriod")
	delete(f, "IsWebTracking")
	delete(f, "Extends")
	delete(f, "PartitionCount")
	delete(f, "CancelTopicAsyncTaskID")
	delete(f, "Encryption")
	delete(f, "IsSourceFrom")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicResponseParams `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebCallbackRequestParams struct {
	// 告警渠道回调配置ID。-通过[获取告警渠道回调配置列表](https://cloud.tencent.com/document/api/614/115229)获取告警渠道回调配置ID
	WebCallbackId *string `json:"WebCallbackId,omitnil,omitempty" name:"WebCallbackId"`

	// 告警渠道回调配置名称。最大支持255个字节
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 渠道类型
	// 
	// WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调;
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 回调地址。
	Webhook *string `json:"Webhook,omitnil,omitempty" name:"Webhook"`

	// 请求方式。
	// 
	// 支持POST、PUT。
	// 
	// 注意：当Type为Http时，必填。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 秘钥信息。最大支持1024个字节
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

type ModifyWebCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 告警渠道回调配置ID。-通过[获取告警渠道回调配置列表](https://cloud.tencent.com/document/api/614/115229)获取告警渠道回调配置ID
	WebCallbackId *string `json:"WebCallbackId,omitnil,omitempty" name:"WebCallbackId"`

	// 告警渠道回调配置名称。最大支持255个字节
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 渠道类型
	// 
	// WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调;
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 回调地址。
	Webhook *string `json:"Webhook,omitnil,omitempty" name:"Webhook"`

	// 请求方式。
	// 
	// 支持POST、PUT。
	// 
	// 注意：当Type为Http时，必填。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 秘钥信息。最大支持1024个字节
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

func (r *ModifyWebCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebCallbackId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Webhook")
	delete(f, "Method")
	delete(f, "Key")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWebCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWebCallbackResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWebCallbackResponseParams `json:"Response"`
}

func (r *ModifyWebCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorNotice struct {
	// 以数组的形式提供MonitorNoticeRule
	Notices []*MonitorNoticeRule `json:"Notices,omitnil,omitempty" name:"Notices"`
}

type MonitorNoticeRule struct {
	// 腾讯云可观测平台通知模板 ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`

	// 腾讯云可观测平台内容模板ID，不传默认内容模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentTmplId *string `json:"ContentTmplId,omitnil,omitempty" name:"ContentTmplId"`

	// 告警级别,0:警告(Warn); 1:提醒(Info); 2:紧急 (Critical)
	AlarmLevels []*uint64 `json:"AlarmLevels,omitnil,omitempty" name:"AlarmLevels"`
}

type MonitorTime struct {
	// 执行周期， 可选值：`Period`、`Fixed`、`Cron`。
	// 
	// - Period：固定频率
	// - Fixed：固定时间
	// - Cron：Cron表达式
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 执行的周期，或者定制执行的时间节点。单位为分钟，取值范围为1~1440。
	// 当type为`Period`,`Fixed`时，time字段生效。
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 执行的周期cron表达式。示例：`"*/1 * * * *"` 从左到右每个field的含义 Minutes field(分钟), Hours field(小时),Day of month field(日期),Month field(月份),Day of week field(星期)， 不支持秒级别。
	// 当type为`Cron`时，CronExpression字段生效。
	CronExpression *string `json:"CronExpression,omitnil,omitempty" name:"CronExpression"`
}

type MultiCondition struct {
	// 触发条件。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 告警级别。0:警告(Warn); 1:提醒(Info); 2:紧急 (Critical)。
	// 
	// - 不填则默认为0。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`
}

type MultiTopicSearchInformation struct {
	// 要检索分析的日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type NoticeContent struct {
	// 渠道类型
	// 
	// Email:邮件;Sms:短信;WeChat:微信;Phone:电话;WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调;
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警触发通知内容模板。
	TriggerContent *NoticeContentInfo `json:"TriggerContent,omitnil,omitempty" name:"TriggerContent"`

	// 告警恢复通知内容模板。
	RecoveryContent *NoticeContentInfo `json:"RecoveryContent,omitnil,omitempty" name:"RecoveryContent"`
}

type NoticeContentInfo struct {
	// 通知内容模板标题信息。
	// 部分通知渠道类型不支持“标题”，请参照[腾讯云控制台页面](https://console.cloud.tencent.com/cls/alarm/notice-template)。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 通知内容模板正文信息。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 请求头（Request Headers）：在HTTP请求中，请求头包含了客户端向服务器发送的附加信息，如用户代理、授权凭证、期望的响应格式等。
	// 仅“自定义回调”支持该配置。
	Headers []*string `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type NoticeContentTemplate struct {
	// 通知内容模板ID。
	NoticeContentId *string `json:"NoticeContentId,omitnil,omitempty" name:"NoticeContentId"`

	// 通知内容模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 语言类型。
	// 
	// 0： 中文
	// 1： 英文
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 通知内容模板信息。
	NoticeContents []*NoticeContent `json:"NoticeContents,omitnil,omitempty" name:"NoticeContents"`

	// 通知内容模板标记。
	// 
	// 0： 用户自定义
	// 1： 系统内置
	Flag *uint64 `json:"Flag,omitnil,omitempty" name:"Flag"`

	// 创建者主账号。
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 创建/修改者子账号。
	SubUin *uint64 `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 创建时间 秒级(s)时间戳。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间 秒级(s)时间戳。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type NoticeReceiver struct {
	// 接受者类型。可选值：
	// -  Uin - 用户ID
	// - Group - 用户组ID
	// 暂不支持其余接收者类型。
	ReceiverType *string `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// 接收者。
	// 当ReceiverType为Uin时，ReceiverIds的值为用户uid。[子用户信息查询](https://cloud.tencent.com/document/api/598/53486)
	// 当ReceiverType为Group时，ReceiverIds的值为用户组id。[CAM用户组](https://cloud.tencent.com/document/product/598/34589)
	ReceiverIds []*int64 `json:"ReceiverIds,omitnil,omitempty" name:"ReceiverIds"`

	// 通知接收渠道。
	// - Email - 邮件
	// - Sms - 短信
	// - WeChat - 微信
	// - Phone - 电话
	ReceiverChannels []*string `json:"ReceiverChannels,omitnil,omitempty" name:"ReceiverChannels"`

	// 通知内容模板ID，使用Default-zh引用默认模板（中文），使用Default-en引用DefaultTemplate(English)。-通过[获取通知内容模板](https://cloud.tencent.com/document/product/614/111714)获取通知内容模板ID
	NoticeContentId *string `json:"NoticeContentId,omitnil,omitempty" name:"NoticeContentId"`

	// 允许接收信息的开始时间。格式：`15:04:05`。必填
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 允许接收信息的结束时间。格式：`15:04:05`。必填
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 位序。
	// 
	// - 入参时无效。
	// - 出参时有效。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

type NoticeRule struct {
	// 匹配规则 JSON串。
	// **rule规则树格式为嵌套结构体JSON字符串**
	// `{"Value":"AND","Type":"Operation","Children":[{"Value":"OR","Type":"Operation","Children":[{"Type":"Condition","Value":"Level","Children":[{"Value":"In","Type":"Compare"},{"Value":"[1,0]","Type":"Value"}]},{"Type":"Condition","Value":"Level","Children":[{"Value":"NotIn","Type":"Compare"},{"Value":"[2]","Type":"Value"}]}]}]}`
	// 
	// **rule规则树限制规则如下**：
	// - 顶层rule中Type可取值：`Condition`，`Operation`
	// - Type为`Operation`的子节点支持的Type可取值：`Condition`，`Operation`
	// - Type为`Condition`的子节点支持的Type可取值：`String`，`Compare`，`Array`，`TimeRange`，`Value`，`Key`
	// - 其他Type无子节点
	// - 当rule Type为`Operation`时，value可取值：`AND`，`OR`
	// - 当rule Type为`Condition`时，value不可为空，子节点个数不能小于2
	//     - 当子节点Type为  `Compare` 时，value可取值：`>`，`<`，`>=`，`<=`，`=`，`!=`，`Between`，`NotBetween`，`=~`，`!=~`，`In`，`NotIn`
	//     - value为`Between`，`NotBetween`时，下一个子节点value必须是长度为2的数组
	//     - value为`=~`，`!=~`时，下一个子节点value必须是一个正则表达式
	//     - value为`In`，`NotIn`时， 下一个子节点value必须是一个数组
	// 
	// **业务参数含义**：
	// - Type：Condition 表示是规则条件，Value：Level 表示告警等级
	//     - 子节点Type支持`Compare`，Value支持`In`，`NotIn`
	//     - 下一个子节点value支持的值：0（警告），1（提醒），2 （紧急）
	// 以下示例表示：告警等级属于提醒
	// `{\"Value\":\"AND\",\"Type\":\"Operation\",\"Children\":[{\"Type\":\"Condition\",\"Value\":\"Level\",\"Children\":[{\"Value\":\"In\",\"Type\":\"Compare\"},{\"Value\":\"[1]\",\"Type\":\"Value\"}]}]}`
	// 
	// - Type：Condition 表示是规则条件，Value：NotifyType 表示通知类型
	//     - 子节点Type支持`Compare`，Value支持`In`，`NotIn`
	//     - 下一个子节点value支持的值：1（告警通知），2 （恢复通知）
	// 以下示例表示：通知类型属于告警通知或通知类型不属于恢复通知
	// `{\"Value\":\"AND\",\"Type\":\"Operation\",\"Children\":[{\"Value\":\"OR\",\"Type\":\"Operation\",\"Children\":[{\"Type\":\"Condition\",\"Value\":\"NotifyType\",\"Children\":[{\"Value\":\"In\",\"Type\":\"Compare\"},{\"Value\":\"[1]\",\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"NotifyType\",\"Children\":[{\"Value\":\"NotIn\",\"Type\":\"Compare\"},{\"Value\":\"[2]\",\"Type\":\"Value\"}]}]}]}`
	// 
	// - Type：Condition 表示是规则条件，Value：AlarmID 表示告警策略
	//     - 子节点Type支持`Compare`，Value支持`In`，`NotIn`
	//     - 下一个子节点value支持的值：告警策略id数组
	// 以下示例表示：告警策略属于alarm-53af048c-254b-4c73-bb48-xxx,alarm-6dfa8bc5-08da-4d64-b6cb-xxx或告警策略不属于alarm-1036314c-1e49-4cee-a8fb-xxx
	// `"{\"Value\":\"AND\",\"Type\":\"Operation\",\"Children\":[{\"Value\":\"OR\",\"Type\":\"Operation\",\"Children\":[{\"Type\":\"Condition\",\"Value\":\"AlarmID\",\"Children\":[{\"Value\":\"In\",\"Type\":\"Compare\"},{\"Value\":\"[\\\"alarm-53af048c-254b-4c73-bb48-xxx\\\",\\\"alarm-6dfa8bc5-08da-4d64-b6cb-xxx\\\"]\",\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"AlarmID\",\"Children\":[{\"Value\":\"NotIn\",\"Type\":\"Compare\"},{\"Value\":\"[\\\"alarm-1036314c-1e49-4cee-a8fb-xxx\\\"]\",\"Type\":\"Value\"}]}]}]}"`
	// 
	// - Type：Condition 表示是规则条件，Value：AlarmName 表示告警策略名称
	//     - 子节点Type支持`Compare`，Value支持`=~`，`!=~`
	//     - 下一个子节点value支持的值：必须是正则表达式
	// 以下示例表示：告警策略名称正则匹配^test$或告警策略名称正则不匹配^hahaha$
	// `{\"Value\":\"AND\",\"Type\":\"Operation\",\"Children\":[{\"Value\":\"OR\",\"Type\":\"Operation\",\"Children\":[{\"Type\":\"Condition\",\"Value\":\"AlarmName\",\"Children\":[{\"Value\":\"=~\",\"Type\":\"Compare\"},{\"Value\":\"^test$\",\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"AlarmName\",\"Children\":[{\"Value\":\"!=~\",\"Type\":\"Compare\"},{\"Value\":\"^hahaha$\",\"Type\":\"Value\"}]}]}]}`
	// 
	// - Type：Condition 表示是规则条件，Value：Label 表示告警分类字段
	//     - 子节点Type支持`Compare`，Value支持`In`，`NotIn`，`=~`，`!=~`
	//     - 下一个子节点value支持的值：`In`，`NotIn` 时value是数组，`=~`，`!=~`时value是正则表达式
	// 以下示例表示：告警分类字段key1属于v1或告警分类字段key2不属于v2或告警分类字段key3正则匹配^test$或告警分类字段key4正则不匹配^hahaha$
	// `{\"Value\":\"AND\",\"Type\":\"Operation\",\"Children\":[{\"Value\":\"OR\",\"Type\":\"Operation\",\"Children\":[{\"Type\":\"Condition\",\"Value\":\"Label\",\"Children\":[{\"Value\":\"key1\",\"Type\":\"Key\"},{\"Value\":\"In\",\"Type\":\"Compare\"},{\"Value\":\"[\\\"v1\\\"]\",\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"Label\",\"Children\":[{\"Value\":\"key2\",\"Type\":\"Key\"},{\"Value\":\"NotIn\",\"Type\":\"Compare\"},{\"Value\":\"[\\\"v2\\\"]\",\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"Label\",\"Children\":[{\"Value\":\"key3\",\"Type\":\"Key\"},{\"Value\":\"=~\",\"Type\":\"Compare\"},{\"Value\":\"^test$\",\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"Label\",\"Children\":[{\"Value\":\"key4\",\"Type\":\"Key\"},{\"Value\":\"!=~\",\"Type\":\"Compare\"},{\"Value\":\"^hahaha$\",\"Type\":\"Value\"}]}]}]}`
	// 
	// - Type：Condition 表示是规则条件，Value：NotifyTime 表示通知时间
	//     - 子节点Type支持`Compare`，Value支持`Between `，`NotBetween `
	//     - 下一个子节点value支持的值：长度为2，格式为`14:20:36`的字符串数组
	// 以下示例表示：通知时间在指定范围内14:18:36至14:33:36或通知时间不在指定范围内14:20:36至14:30:36
	// `{\"Value\":\"AND\",\"Type\":\"Operation\",\"Children\":[{\"Value\":\"OR\",\"Type\":\"Operation\",\"Children\":[{\"Type\":\"Condition\",\"Value\":\"NotifyTime\",\"Children\":[{\"Value\":\"Between\",\"Type\":\"Compare\"},{\"Value\":\"[\\\"14:18:36\\\",\\\"14:33:36\\\"]\",\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"NotifyTime\",\"Children\":[{\"Value\":\"NotBetween\",\"Type\":\"Compare\"},{\"Value\":\"[\\\"14:20:36\\\",\\\"14:30:36\\\"]\",\"Type\":\"Value\"}]}]}]}`
	// 
	// - Type：Condition 表示是规则条件，Value：Duration 表示告警持续时间
	//     - 子节点Type支持`Compare`，Value支持`>`，`<`，`>=`，`<=`
	//     - 下一个子节点value支持的值：整型值单位分钟
	// 以下示例表示：告警持续时间大于1分钟或告警持续时间大于等于2分钟或告警持续时间小于3分钟或告警持续时间小于等于4分钟
	// `{\"Value\":\"AND\",\"Type\":\"Operation\",\"Children\":[{\"Value\":\"OR\",\"Type\":\"Operation\",\"Children\":[{\"Type\":\"Condition\",\"Value\":\"Duration\",\"Children\":[{\"Value\":\">\",\"Type\":\"Compare\"},{\"Value\":1,\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"Duration\",\"Children\":[{\"Value\":\">=\",\"Type\":\"Compare\"},{\"Value\":2,\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"Duration\",\"Children\":[{\"Value\":\"<\",\"Type\":\"Compare\"},{\"Value\":3,\"Type\":\"Value\"}]},{\"Type\":\"Condition\",\"Value\":\"Duration\",\"Children\":[{\"Value\":\"<=\",\"Type\":\"Compare\"},{\"Value\":4,\"Type\":\"Value\"}]}]}]}`
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 告警通知接收者信息。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitnil,omitempty" name:"NoticeReceivers"`

	// 告警通知模板回调信息，包括企业微信、钉钉、飞书。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitnil,omitempty" name:"WebCallbacks"`

	// 告警升级开关。`true`：开启告警升级、`false`：关闭告警升级，默认：false
	Escalate *bool `json:"Escalate,omitnil,omitempty" name:"Escalate"`

	// 告警升级条件。`1`：无人认领且未恢复、`2`：未恢复，默认为1
	// - 无人认领且未恢复：告警没有恢复并且没有人认领则升级
	// - 未恢复：当前告警持续未恢复则升级
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警升级间隔。单位：分钟，范围`[1，14400]`
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 告警升级后下一个环节的通知渠道配置
	EscalateNotice *EscalateNoticeInfo `json:"EscalateNotice,omitnil,omitempty" name:"EscalateNotice"`
}

// Predefined struct for user
type OpenKafkaConsumerRequestParams struct {
	// 日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`

	// 压缩方式[0:NONE；2:SNAPPY；3:LZ4]，默认：0
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// kafka协议消费数据格式
	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitnil,omitempty" name:"ConsumerContent"`
}

type OpenKafkaConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	// - 通过 [创建日志主题](https://cloud.tencent.com/document/product/614/56456) 获取日志主题Id。
	FromTopicId *string `json:"FromTopicId,omitnil,omitempty" name:"FromTopicId"`

	// 压缩方式[0:NONE；2:SNAPPY；3:LZ4]，默认：0
	Compression *int64 `json:"Compression,omitnil,omitempty" name:"Compression"`

	// kafka协议消费数据格式
	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitnil,omitempty" name:"ConsumerContent"`
}

func (r *OpenKafkaConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenKafkaConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTopicId")
	delete(f, "Compression")
	delete(f, "ConsumerContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenKafkaConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenKafkaConsumerResponseParams struct {
	// KafkaConsumer 消费时使用的Topic参数
	TopicID *string `json:"TopicID,omitnil,omitempty" name:"TopicID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenKafkaConsumerResponse struct {
	*tchttp.BaseResponse
	Response *OpenKafkaConsumerResponseParams `json:"Response"`
}

func (r *OpenKafkaConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParquetInfo struct {
	// ParquetKeyInfo数组
	ParquetKeyInfo []*ParquetKeyInfo `json:"ParquetKeyInfo,omitnil,omitempty" name:"ParquetKeyInfo"`
}

type ParquetKeyInfo struct {
	// 键值名称
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 数据类型，目前支持6种类型：string、boolean、int32、int64、float、double
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// 解析失败赋值信息
	KeyNonExistingField *string `json:"KeyNonExistingField,omitnil,omitempty" name:"KeyNonExistingField"`
}

type PartitionInfo struct {
	// 分区ID
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`

	// 分区的状态（readwrite或者是readonly）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 分区哈希键起始key
	InclusiveBeginKey *string `json:"InclusiveBeginKey,omitnil,omitempty" name:"InclusiveBeginKey"`

	// 分区哈希键结束key
	ExclusiveEndKey *string `json:"ExclusiveEndKey,omitnil,omitempty" name:"ExclusiveEndKey"`

	// 分区创建时间
	// 时间格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 只读分区数据停止写入时间
	// 时间格式：yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastWriteTime *string `json:"LastWriteTime,omitnil,omitempty" name:"LastWriteTime"`
}

// Predefined struct for user
type PreviewKafkaRechargeRequestParams struct {
	// 预览类型，1：源数据预览；2：导出结果预览。
	PreviewType *uint64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 导入Kafka类型，0：腾讯云CKafka；1：用户自建Kafka。
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开。
	// 最多支持100个。
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// 导入数据位置，-2：最早；-1：最晚。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 腾讯云CKafka实例ID，当KafkaType为0时参数KafkaInstance有效且必填。
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取实例id。
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// 服务地址。
	// KafkaType为1时ServerAddr必填。
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// ServerAddr是否为加密连接。
	// KafkaType为1时有效。
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// 加密访问协议。
	// KafkaType为1并且IsEncryptionAddr为true时Protocol必填。
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 用户Kafka消费组。
	// 
	// - 消费组是 Kafka 提供的可扩展且具有容错性的消费者机制，一个消费组中存在多个消费者，组内的所有消费者共同消费订阅 Topic 中的消息。一个消费者可同时消费多个 Partition，但一个 Partition 只能被消费组内的一个消费者消费。
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 日志导入规则
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`

	// 用户kafka拓展信息
	UserKafkaMeta *UserKafkaMeta `json:"UserKafkaMeta,omitnil,omitempty" name:"UserKafkaMeta"`
}

type PreviewKafkaRechargeRequest struct {
	*tchttp.BaseRequest
	
	// 预览类型，1：源数据预览；2：导出结果预览。
	PreviewType *uint64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 导入Kafka类型，0：腾讯云CKafka；1：用户自建Kafka。
	KafkaType *uint64 `json:"KafkaType,omitnil,omitempty" name:"KafkaType"`

	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开。
	// 最多支持100个。
	UserKafkaTopics *string `json:"UserKafkaTopics,omitnil,omitempty" name:"UserKafkaTopics"`

	// 导入数据位置，-2：最早；-1：最晚。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 腾讯云CKafka实例ID，当KafkaType为0时参数KafkaInstance有效且必填。
	// - 通过 [获取实例列表信息](https://cloud.tencent.com/document/product/597/40835) 获取实例id。
	KafkaInstance *string `json:"KafkaInstance,omitnil,omitempty" name:"KafkaInstance"`

	// 服务地址。
	// KafkaType为1时ServerAddr必填。
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// ServerAddr是否为加密连接。
	// KafkaType为1时有效。
	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitnil,omitempty" name:"IsEncryptionAddr"`

	// 加密访问协议。
	// KafkaType为1并且IsEncryptionAddr为true时Protocol必填。
	Protocol *KafkaProtocolInfo `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 用户Kafka消费组。
	// 
	// - 消费组是 Kafka 提供的可扩展且具有容错性的消费者机制，一个消费组中存在多个消费者，组内的所有消费者共同消费订阅 Topic 中的消息。一个消费者可同时消费多个 Partition，但一个 Partition 只能被消费组内的一个消费者消费。
	ConsumerGroupName *string `json:"ConsumerGroupName,omitnil,omitempty" name:"ConsumerGroupName"`

	// 日志导入规则
	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitnil,omitempty" name:"LogRechargeRule"`

	// 用户kafka拓展信息
	UserKafkaMeta *UserKafkaMeta `json:"UserKafkaMeta,omitnil,omitempty" name:"UserKafkaMeta"`
}

func (r *PreviewKafkaRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewKafkaRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PreviewType")
	delete(f, "KafkaType")
	delete(f, "UserKafkaTopics")
	delete(f, "Offset")
	delete(f, "KafkaInstance")
	delete(f, "ServerAddr")
	delete(f, "IsEncryptionAddr")
	delete(f, "Protocol")
	delete(f, "ConsumerGroupName")
	delete(f, "LogRechargeRule")
	delete(f, "UserKafkaMeta")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PreviewKafkaRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PreviewKafkaRechargeResponseParams struct {
	// 日志样例，PreviewType为2时返回
	LogSample *string `json:"LogSample,omitnil,omitempty" name:"LogSample"`

	// 日志预览结果
	LogData *string `json:"LogData,omitnil,omitempty" name:"LogData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PreviewKafkaRechargeResponse struct {
	*tchttp.BaseResponse
	Response *PreviewKafkaRechargeResponseParams `json:"Response"`
}

func (r *PreviewKafkaRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreviewLogStatistic struct {
	// 日志内容
	LogContent *string `json:"LogContent,omitnil,omitempty" name:"LogContent"`

	// 行号。从0开始
	LineNum *int64 `json:"LineNum,omitnil,omitempty" name:"LineNum"`

	// 目标日志主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	DstTopicId *string `json:"DstTopicId,omitnil,omitempty" name:"DstTopicId"`

	// 失败错误信息， 空字符串""表示正常
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 日志时间，格式：`2024-05-07 17:13:17.105`
	// 
	// - 入参时无效
	// - 出参时有效，为日志中的时间格式
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 目标topic-name
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DstTopicName is deprecated.
	DstTopicName *string `json:"DstTopicName,omitnil,omitempty" name:"DstTopicName"`
}

// Predefined struct for user
type QueryMetricRequestParams struct {
	// 查询语句，使用PromQL语法	
	// - 参考 [语法规则](https://cloud.tencent.com/document/product/614/90334) 文档
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 指标主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 查询时间，秒级Unix时间戳。为空时代表当前时间戳。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`
}

type QueryMetricRequest struct {
	*tchttp.BaseRequest
	
	// 查询语句，使用PromQL语法	
	// - 参考 [语法规则](https://cloud.tencent.com/document/product/614/90334) 文档
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 指标主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 查询时间，秒级Unix时间戳。为空时代表当前时间戳。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`
}

func (r *QueryMetricRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMetricRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "TopicId")
	delete(f, "Time")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMetricRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMetricResponseParams struct {
	// 指标查询结果类型，支持
	// - scalar 标量值
	// - string 字符串值
	// - vector 瞬时向量
	// - matrix 区间向量
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// 指标查询结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryMetricResponse struct {
	*tchttp.BaseResponse
	Response *QueryMetricResponseParams `json:"Response"`
}

func (r *QueryMetricResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRangeMetricRequestParams struct {
	// 指标主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 查询语句，使用PromQL语法
	// - 参考 [语法规则](https://cloud.tencent.com/document/product/614/90334) 文档
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 查询起始时间，秒级Unix时间戳
	Start *uint64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 查询结束时间，秒级Unix时间戳
	End *uint64 `json:"End,omitnil,omitempty" name:"End"`

	// 查询时间间隔，单位秒
	Step *uint64 `json:"Step,omitnil,omitempty" name:"Step"`
}

type QueryRangeMetricRequest struct {
	*tchttp.BaseRequest
	
	// 指标主题ID
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 查询语句，使用PromQL语法
	// - 参考 [语法规则](https://cloud.tencent.com/document/product/614/90334) 文档
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 查询起始时间，秒级Unix时间戳
	Start *uint64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 查询结束时间，秒级Unix时间戳
	End *uint64 `json:"End,omitnil,omitempty" name:"End"`

	// 查询时间间隔，单位秒
	Step *uint64 `json:"Step,omitnil,omitempty" name:"Step"`
}

func (r *QueryRangeMetricRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRangeMetricRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Query")
	delete(f, "Start")
	delete(f, "End")
	delete(f, "Step")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryRangeMetricRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRangeMetricResponseParams struct {
	// 指标查询结果类型，支持
	// - scalar 标量值
	// - string 字符串值
	// - vector 瞬时向量
	// - matrix 区间向量
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// 指标查询结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryRangeMetricResponse struct {
	*tchttp.BaseResponse
	Response *QueryRangeMetricResponseParams `json:"Response"`
}

func (r *QueryRangeMetricResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRangeMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryShipperTaskRequestParams struct {
	// 投递规则Id。
	// 
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745)获取ShipperId。
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// 投递任务Id。
	// 
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745) 获取TaskId。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type RetryShipperTaskRequest struct {
	*tchttp.BaseRequest
	
	// 投递规则Id。
	// 
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745)获取ShipperId。
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// 投递任务Id。
	// 
	// - 通过 [获取投递任务列表](https://cloud.tencent.com/document/product/614/58745) 获取TaskId。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *RetryShipperTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryShipperTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShipperId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryShipperTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryShipperTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryShipperTaskResponse struct {
	*tchttp.BaseResponse
	Response *RetryShipperTaskResponseParams `json:"Response"`
}

func (r *RetryShipperTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryShipperTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleInfo struct {
	// 全文索引配置, 为空时代表未开启全文索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullText *FullTextInfo `json:"FullText,omitnil,omitempty" name:"FullText"`

	// 键值索引配置，为空时代表未开启键值索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValue *RuleKeyValueInfo `json:"KeyValue,omitnil,omitempty" name:"KeyValue"`

	// 元字段索引配置，为空时代表未开启元字段索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *RuleTagInfo `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 键值索引自动配置，为空时代表未开启该功能。
	// 启用后自动将日志内的字段添加到键值索引中，包括日志中后续新增的字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DynamicIndex *DynamicIndex `json:"DynamicIndex,omitnil,omitempty" name:"DynamicIndex"`
}

type RuleKeyValueInfo struct {
	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 需要建立索引的键值对信息
	KeyValues []*KeyValueInfo `json:"KeyValues,omitnil,omitempty" name:"KeyValues"`
}

type RuleTagInfo struct {
	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 元字段索引配置中的字段信息
	KeyValues []*KeyValueInfo `json:"KeyValues,omitnil,omitempty" name:"KeyValues"`
}

type ScheduledSqlResouceInfo struct {
	// 通过 [获取日志主题列表](https://cloud.tencent.com/document/product/614/56454) 获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题的地域信息，当前不支持跨地域，支持地域参考 [地域列表](https://cloud.tencent.com/document/api/614/56474#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8) 文档。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 主题类型：0为日志主题，1为指标主题
	BizType *int64 `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 指标名称。当BizType为1时，MetricName需要填写
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 指标名称
	// BizType为1时，优先使用MetricNames字段多指标只能填充到MetricNames字段，单指标建议填充到MetricName字段
	MetricNames []*string `json:"MetricNames,omitnil,omitempty" name:"MetricNames"`

	// 指标维度，不接受时间类型。
	MetricLabels []*string `json:"MetricLabels,omitnil,omitempty" name:"MetricLabels"`

	// 指标时间戳，默认值为SQL查询时间范围的左侧时间点，您也可以指定其他字段（类型为uinx时间、TimeStamp，精度毫秒）为指标时间戳。
	CustomTime *string `json:"CustomTime,omitnil,omitempty" name:"CustomTime"`

	// 除了MetricLabels，您还可以使用该参数，为指标补充静态的维度。
	// 维度名以字母或下划线开头，后面可以跟字母、数字或下划线，长度小于等于1024 字节
	CustomMetricLabels []*MetricLabel `json:"CustomMetricLabels,omitnil,omitempty" name:"CustomMetricLabels"`
}

type ScheduledSqlTaskInfo struct {
	// ScheduledSql任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// ScheduledSql任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源日志主题id
	SrcTopicId *string `json:"SrcTopicId,omitnil,omitempty" name:"SrcTopicId"`

	// 源日志主题名称
	SrcTopicName *string `json:"SrcTopicName,omitnil,omitempty" name:"SrcTopicName"`

	// 定时SQL分析目标主题
	DstResource *ScheduledSqlResouceInfo `json:"DstResource,omitnil,omitempty" name:"DstResource"`

	// 任务创建时间。格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务更新时间，格式：yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 任务状态，1:运行 2:停止 3:异常-找不到源日志主题 4:异常-找不到目标主题
	// 
	// 5: 访问权限问题 6:内部故障 7:其他故障
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务启用状态，1开启,  2关闭
	EnableFlag *int64 `json:"EnableFlag,omitnil,omitempty" name:"EnableFlag"`

	// 查询语句
	ScheduledSqlContent *string `json:"ScheduledSqlContent,omitnil,omitempty" name:"ScheduledSqlContent"`

	// 调度开始时间，格式：yyyy-MM-dd HH:mm:ss
	ProcessStartTime *string `json:"ProcessStartTime,omitnil,omitempty" name:"ProcessStartTime"`

	// 调度类型，1:持续运行 2:指定时间范围
	ProcessType *int64 `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// 调度结束时间，格式：yyyy-MM-dd HH:mm:ss，当process_type=2时为必传字段
	ProcessEndTime *string `json:"ProcessEndTime,omitnil,omitempty" name:"ProcessEndTime"`

	// 调度周期(分钟)，1~1440分钟
	ProcessPeriod *int64 `json:"ProcessPeriod,omitnil,omitempty" name:"ProcessPeriod"`

	// 查询的时间窗口. @m-15m, @m，意为近15分钟
	ProcessTimeWindow *string `json:"ProcessTimeWindow,omitnil,omitempty" name:"ProcessTimeWindow"`

	// 执行延迟(秒)，0~120秒，默认60秒
	ProcessDelay *int64 `json:"ProcessDelay,omitnil,omitempty" name:"ProcessDelay"`

	// 源topicId的地域信息，支持地域见 [地域列表](https://cloud.tencent.com/document/api/614/56474#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8) 文档。
	SrcTopicRegion *string `json:"SrcTopicRegion,omitnil,omitempty" name:"SrcTopicRegion"`

	// 语法规则，0：Lucene语法，1：CQL语法
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// 是否开启投递服务日志。1：关闭，2：开启。
	HasServicesLog *uint64 `json:"HasServicesLog,omitnil,omitempty" name:"HasServicesLog"`

	// 全文检索标记。1：关闭，2：打开。
	FullQuery *uint64 `json:"FullQuery,omitnil,omitempty" name:"FullQuery"`
}

// Predefined struct for user
type SearchCosRechargeInfoRequestParams struct {
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志集Id。
	// 
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// COS导入任务名称,最大支持128个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// COS存储桶所在地域，详见产品支持的[地域列表](https://cloud.tencent.com/document/product/436/6224)。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// COS文件所在文件夹的前缀。默认为空，投递存储桶下所有的文件。
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 压缩模式:   "", "gzip", "lzop", "snappy"。  默认："" 不压缩
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`
}

type SearchCosRechargeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id。
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志集Id。
	// 
	// - 通过[获取日志集列表](https://cloud.tencent.com/document/product/614/58624)获取日志集Id。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// COS导入任务名称,最大支持128个字节。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// COS存储桶，详见产品支持的[存储桶命名规范](https://cloud.tencent.com/document/product/436/13312)。	
	// 
	// - 通过[GET Service（List Buckets）](https://cloud.tencent.com/document/product/436/8291)获取COS存储桶。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// COS存储桶所在地域，详见产品支持的[地域列表](https://cloud.tencent.com/document/product/436/6224)。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// COS文件所在文件夹的前缀。默认为空，投递存储桶下所有的文件。
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 压缩模式:   "", "gzip", "lzop", "snappy"。  默认："" 不压缩
	Compress *string `json:"Compress,omitnil,omitempty" name:"Compress"`
}

func (r *SearchCosRechargeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchCosRechargeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "LogsetId")
	delete(f, "Name")
	delete(f, "Bucket")
	delete(f, "BucketRegion")
	delete(f, "Prefix")
	delete(f, "Compress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchCosRechargeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchCosRechargeInfoResponseParams struct {
	// 匹配到的存储桶下的某个文件的前几行数据
	Data []*string `json:"Data,omitnil,omitempty" name:"Data"`

	// 匹配到的存储桶下的文件个数
	Sum *uint64 `json:"Sum,omitnil,omitempty" name:"Sum"`

	// 当前预览文件路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 预览获取数据失败原因
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 状态。
	// - 0：成功
	// - 10000：参数错误，请确认参数
	// - 10001：授权失败，请确认授权
	// - 10002：获取文件列表失败，请稍后再试。若无法解决，请咨询 [在线支持](https://cloud.tencent.com/online-service) 或 [提交工单](https://console.cloud.tencent.com/workorder/category?level1_id=83&level2_id=469&source=14&data_title=%E6%97%A5%E5%BF%97%E6%9C%8D%E5%8A%A1&step=1) 处理。
	// - 10003：桶内无相应前缀文件，请使用正确的桶、文件前缀和压缩方式
	// - 10004：文件下载失败，请稍后再试。若无法解决，请咨询 [在线支持](https://cloud.tencent.com/online-service) 或 [提交工单](https://console.cloud.tencent.com/workorder/category?level1_id=83&level2_id=469&source=14&data_title=%E6%97%A5%E5%BF%97%E6%9C%8D%E5%8A%A1&step=1) 处理。
	// - 10005：文件解压缩失败，请选择正确的压缩方式然后再试
	// - 10006：读取文件内容失败，请确认文件可读
	// - 10007：文件预览失败，请稍后再试。若无法解决，请咨询 [在线支持](https://cloud.tencent.com/online-service) 或 [提交工单](https://console.cloud.tencent.com/workorder/category?level1_id=83&level2_id=469&source=14&data_title=%E6%97%A5%E5%BF%97%E6%9C%8D%E5%8A%A1&step=1) 处理。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchCosRechargeInfoResponse struct {
	*tchttp.BaseResponse
	Response *SearchCosRechargeInfoResponseParams `json:"Response"`
}

func (r *SearchCosRechargeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchCosRechargeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchDashboardSubscribeRequestParams struct {
	// 仪表盘id。通过 [获取仪表盘](https://cloud.tencent.com/document/api/614/95636)接口获取DashboardId。
	DashboardId *string `json:"DashboardId,omitnil,omitempty" name:"DashboardId"`

	// 仪表盘订阅数据。
	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitnil,omitempty" name:"SubscribeData"`

	// 仪表盘订阅Id。通过 [获取仪表盘订阅列表](https://cloud.tencent.com/document/api/614/105779)接口获取Id。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 仪表盘订阅名称。通过 [获取仪表盘订阅列表](https://cloud.tencent.com/document/api/614/105779)接口获取Name。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type SearchDashboardSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 仪表盘id。通过 [获取仪表盘](https://cloud.tencent.com/document/api/614/95636)接口获取DashboardId。
	DashboardId *string `json:"DashboardId,omitnil,omitempty" name:"DashboardId"`

	// 仪表盘订阅数据。
	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitnil,omitempty" name:"SubscribeData"`

	// 仪表盘订阅Id。通过 [获取仪表盘订阅列表](https://cloud.tencent.com/document/api/614/105779)接口获取Id。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 仪表盘订阅名称。通过 [获取仪表盘订阅列表](https://cloud.tencent.com/document/api/614/105779)接口获取Name。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *SearchDashboardSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchDashboardSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DashboardId")
	delete(f, "SubscribeData")
	delete(f, "Id")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchDashboardSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchDashboardSubscribeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchDashboardSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *SearchDashboardSubscribeResponseParams `json:"Response"`
}

func (r *SearchDashboardSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchDashboardSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogErrors struct {
	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 错误码
	ErrorCodeStr *string `json:"ErrorCodeStr,omitnil,omitempty" name:"ErrorCodeStr"`
}

type SearchLogInfos struct {
	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志存储生命周期
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 透传本次接口返回的Context值，可获取后续更多日志，过期时间1小时
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

// Predefined struct for user
type SearchLogRequestParams struct {
	// 要检索分析的日志的起始时间，**Unix时间戳（毫秒）**
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要检索分析的日志的结束时间，**Unix时间戳（毫秒）**
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 检索分析语句，最大长度为12KB
	// 语句由 <a href="https://cloud.tencent.com/document/product/614/47044" target="_blank">[检索条件]</a> | <a href="https://cloud.tencent.com/document/product/614/44061" target="_blank">[SQL语句]</a>构成，无需对日志进行统计分析时，可省略其中的管道符<code> | </code>及SQL语句
	// 使用*或空字符串可查询所有日志
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 检索语法规则，默认值为0，推荐使用1 。
	// - 0：Lucene语法
	// - 1：CQL语法（CLS Query Language，日志服务专用检索语法）
	// 
	//  ⚠️⚠️ **注意**
	//  **该参数值建议设置为 1，使用 CQL 语法规则，控制台日志检索及仪表盘默认均使用该语法规则。**
	//  该参数值未指定或者为 0 时，将使用 Lucene 语法，语法容易报错且查询结果与控制台默认语法规则不一致。详细说明参见<a href="https://cloud.tencent.com/document/product/614/47044#RetrievesConditionalRules" target="_blank">检索条件语法规则</a>。
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// - 要检索分析的日志主题ID，仅能指定一个日志主题。
	// - 如需同时检索多个日志主题，请使用Topics参数。
	// - TopicId 和 Topics 不能同时使用，在一次请求中有且只能选择一个。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// - 要检索分析的日志主题列表，最大支持50个日志主题。
	// - 检索单个日志主题时请使用TopicId。
	// - TopicId 和 Topics 不能同时使用，在一次请求中有且只能选择一个。
	Topics []*MultiTopicSearchInformation `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果排序方式参考<a href="https://cloud.tencent.com/document/product/614/58978" target="_blank">SQL ORDER BY语法</a>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 表示单次查询返回的原始日志条数，默认为100，最大值为1000。
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果条数指定方式参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	// 
	// 可通过两种方式获取后续更多日志：
	// * Context:透传上次接口返回的Context值，获取后续更多日志，总计最多可获取1万条原始日志
	// * Offset:偏移量，表示从第几行开始返回原始日志，无日志条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询原始日志的偏移量，表示从第几行开始返回原始日志，默认为0。 
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * 不能与Context参数同时使用
	// * 仅适用于单日志主题检索
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。
	// 注意：
	// * 透传该参数时，请勿修改除该参数外的其它参数
	// * 仅适用于单日志主题检索，检索多个日志主题时，请使用Topics中的Context
	// * 仅当检索分析语句(Query)不包含SQL时有效，SQL获取后续结果参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 执行统计分析（Query中包含SQL）时，是否对原始日志先进行采样，再进行统计分析。
	// 0：自动采样;
	// 0～1：按指定采样率采样，例如0.02;
	// 1：不采样，即精确分析
	// 默认值为1
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// 为true代表使用新的检索结果返回方式，输出参数AnalysisRecords和Columns有效
	// 为false时代表使用老的检索结果返回方式, 输出AnalysisResults和ColNames有效
	// 两种返回方式在编码格式上有少量区别，建议使用true
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`

	// 是否高亮符合检索条件的关键词，一般用于高亮显示。仅支持键值检索，不支持全文检索
	HighLight *bool `json:"HighLight,omitnil,omitempty" name:"HighLight"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest
	
	// 要检索分析的日志的起始时间，**Unix时间戳（毫秒）**
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要检索分析的日志的结束时间，**Unix时间戳（毫秒）**
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 检索分析语句，最大长度为12KB
	// 语句由 <a href="https://cloud.tencent.com/document/product/614/47044" target="_blank">[检索条件]</a> | <a href="https://cloud.tencent.com/document/product/614/44061" target="_blank">[SQL语句]</a>构成，无需对日志进行统计分析时，可省略其中的管道符<code> | </code>及SQL语句
	// 使用*或空字符串可查询所有日志
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 检索语法规则，默认值为0，推荐使用1 。
	// - 0：Lucene语法
	// - 1：CQL语法（CLS Query Language，日志服务专用检索语法）
	// 
	//  ⚠️⚠️ **注意**
	//  **该参数值建议设置为 1，使用 CQL 语法规则，控制台日志检索及仪表盘默认均使用该语法规则。**
	//  该参数值未指定或者为 0 时，将使用 Lucene 语法，语法容易报错且查询结果与控制台默认语法规则不一致。详细说明参见<a href="https://cloud.tencent.com/document/product/614/47044#RetrievesConditionalRules" target="_blank">检索条件语法规则</a>。
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// - 要检索分析的日志主题ID，仅能指定一个日志主题。
	// - 如需同时检索多个日志主题，请使用Topics参数。
	// - TopicId 和 Topics 不能同时使用，在一次请求中有且只能选择一个。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// - 要检索分析的日志主题列表，最大支持50个日志主题。
	// - 检索单个日志主题时请使用TopicId。
	// - TopicId 和 Topics 不能同时使用，在一次请求中有且只能选择一个。
	Topics []*MultiTopicSearchInformation `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果排序方式参考<a href="https://cloud.tencent.com/document/product/614/58978" target="_blank">SQL ORDER BY语法</a>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 表示单次查询返回的原始日志条数，默认为100，最大值为1000。
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果条数指定方式参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	// 
	// 可通过两种方式获取后续更多日志：
	// * Context:透传上次接口返回的Context值，获取后续更多日志，总计最多可获取1万条原始日志
	// * Offset:偏移量，表示从第几行开始返回原始日志，无日志条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询原始日志的偏移量，表示从第几行开始返回原始日志，默认为0。 
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * 不能与Context参数同时使用
	// * 仅适用于单日志主题检索
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。
	// 注意：
	// * 透传该参数时，请勿修改除该参数外的其它参数
	// * 仅适用于单日志主题检索，检索多个日志主题时，请使用Topics中的Context
	// * 仅当检索分析语句(Query)不包含SQL时有效，SQL获取后续结果参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 执行统计分析（Query中包含SQL）时，是否对原始日志先进行采样，再进行统计分析。
	// 0：自动采样;
	// 0～1：按指定采样率采样，例如0.02;
	// 1：不采样，即精确分析
	// 默认值为1
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// 为true代表使用新的检索结果返回方式，输出参数AnalysisRecords和Columns有效
	// 为false时代表使用老的检索结果返回方式, 输出AnalysisResults和ColNames有效
	// 两种返回方式在编码格式上有少量区别，建议使用true
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`

	// 是否高亮符合检索条件的关键词，一般用于高亮显示。仅支持键值检索，不支持全文检索
	HighLight *bool `json:"HighLight,omitnil,omitempty" name:"HighLight"`
}

func (r *SearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "SyntaxRule")
	delete(f, "TopicId")
	delete(f, "Topics")
	delete(f, "Sort")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Context")
	delete(f, "SamplingRate")
	delete(f, "UseNewAnalysis")
	delete(f, "HighLight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchLogResponseParams struct {
	// 透传本次接口返回的Context值，可获取后续更多日志，过期时间1小时。
	// 注意：
	// * 仅适用于单日志主题检索，检索多个日志主题时，请使用Topics中的Context
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 符合检索条件的日志是否已全部返回，如未全部返回可使用Context参数获取后续更多日志
	// 注意：仅当检索分析语句(Query)不包含SQL时有效
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 返回的是否为统计分析（即SQL）结果
	Analysis *bool `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// 匹配检索条件的原始日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*LogInfo `json:"Results,omitnil,omitempty" name:"Results"`

	// 日志统计分析结果的列名
	// 当UseNewAnalysis为false时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColNames []*string `json:"ColNames,omitnil,omitempty" name:"ColNames"`

	// 日志统计分析结果
	// 当UseNewAnalysis为false时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisResults []*LogItems `json:"AnalysisResults,omitnil,omitempty" name:"AnalysisResults"`

	// 日志统计分析结果
	// 当UseNewAnalysis为true时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisRecords []*string `json:"AnalysisRecords,omitnil,omitempty" name:"AnalysisRecords"`

	// 日志统计分析结果的列属性
	// 当UseNewAnalysis为true时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 本次统计分析使用的采样率
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// 使用多日志主题检索时，各个日志主题的基本信息，例如报错信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topics *SearchLogTopics `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchLogResponseParams `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogTopics struct {
	// 多日志主题检索对应的错误信息
	Errors []*SearchLogErrors `json:"Errors,omitnil,omitempty" name:"Errors"`

	// 多日志主题检索各日志主题信息
	Infos []*SearchLogInfos `json:"Infos,omitnil,omitempty" name:"Infos"`
}

type ShipperInfo struct {
	// 投递规则ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 投递的bucket地址
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 投递的前缀目录
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`

	// 投递规则的名字
	ShipperName *string `json:"ShipperName,omitnil,omitempty" name:"ShipperName"`

	// 投递的时间间隔，单位 秒
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 投递的文件的最大值，单位 MB
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 是否生效
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 投递日志的过滤规则
	FilterRules []*FilterRuleInfo `json:"FilterRules,omitnil,omitempty" name:"FilterRules"`

	// 投递日志的分区规则，支持strftime的时间格式表示
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 投递日志的压缩配置
	Compress *CompressInfo `json:"Compress,omitnil,omitempty" name:"Compress"`

	// 投递日志的内容格式配置
	Content *ContentInfo `json:"Content,omitnil,omitempty" name:"Content"`

	// 投递日志的创建时间。格式：YYYY-MM-DD HH:MM:SS
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 投递文件命名配置，0：随机数命名，1：投递时间命名，默认0（随机数命名）
	FilenameMode *uint64 `json:"FilenameMode,omitnil,omitempty" name:"FilenameMode"`

	// 投递数据范围的开始时间点
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 投递数据范围的结束时间点
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 历史数据投递的进度（仅当用户选择的数据内中历史数据时才有效）
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 历史数据全部投递完成剩余的时间（仅当用户选择的数据中有历史数据时才有效）
	RemainTime *int64 `json:"RemainTime,omitnil,omitempty" name:"RemainTime"`

	// 历史任务状态：
	// 0：实时任务
	// 1：任务准备中
	// 2：任务运行中
	// 3：任务运行异常
	// 4：任务运行结束
	HistoryStatus *int64 `json:"HistoryStatus,omitnil,omitempty" name:"HistoryStatus"`

	// 对象存储类型，默认值为 STANDARD。枚举值请参见[ 存储类型概述](https://cloud.tencent.com/document/product/436/33417) 文档。
	// 参考值有：
	// STANDARD：标准存储
	// STANDARD_IA：低频存储
	// ARCHIVE：归档存储
	// DEEP_ARCHIVE：深度归档存储
	// MAZ_STANDARD：标准存储（多 AZ）
	// MAZ_STANDARD_IA：低频存储（多 AZ）
	// INTELLIGENT_TIERING：智能分层存储
	// MAZ_INTELLIGENT_TIERING：智能分层存储（多 AZ）
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 角色访问描述名 [创建角色](https://cloud.tencent.com/document/product/598/19381)
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 外部ID
	ExternalId *string `json:"ExternalId,omitnil,omitempty" name:"ExternalId"`

	// 任务运行状态。支持`0`,`1`,`2`
	// 
	// - `0`: 停止
	// - `1`: 运行中
	// - `2`: 异常
	TaskStatus *uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`
}

type ShipperTaskInfo struct {
	// 投递任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 投递信息ID
	ShipperId *string `json:"ShipperId,omitnil,omitempty" name:"ShipperId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 本批投递的日志的开始时间戳，毫秒
	RangeStart *int64 `json:"RangeStart,omitnil,omitempty" name:"RangeStart"`

	// 本批投递的日志的结束时间戳， 毫秒
	RangeEnd *int64 `json:"RangeEnd,omitnil,omitempty" name:"RangeEnd"`

	// 本次投递任务的开始时间戳， 毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 本次投递任务的结束时间戳， 毫秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 本次投递的结果。"success"，"running"，"failed"
	// 
	// - success：任务成功。
	// - running：任务处理中。
	// - failed：任务失败。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 结果的详细信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type SplitPartitionRequestParams struct {
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 待分裂分区Id
	// - 通过[获取分区列表](https://cloud.tencent.com/document/product/614/56470)获取待分裂分区Id。
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`

	// 分区切分的哈希key的位置，只在Number=2时有意义
	SplitKey *string `json:"SplitKey,omitnil,omitempty" name:"SplitKey"`

	// 分区分裂个数(可选)，默认等于2
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`
}

type SplitPartitionRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题Id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 待分裂分区Id
	// - 通过[获取分区列表](https://cloud.tencent.com/document/product/614/56470)获取待分裂分区Id。
	PartitionId *int64 `json:"PartitionId,omitnil,omitempty" name:"PartitionId"`

	// 分区切分的哈希key的位置，只在Number=2时有意义
	SplitKey *string `json:"SplitKey,omitnil,omitempty" name:"SplitKey"`

	// 分区分裂个数(可选)，默认等于2
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`
}

func (r *SplitPartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	delete(f, "SplitKey")
	delete(f, "Number")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SplitPartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SplitPartitionResponseParams struct {
	// 分裂结果集
	Partitions []*PartitionInfo `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SplitPartitionResponse struct {
	*tchttp.BaseResponse
	Response *SplitPartitionResponseParams `json:"Response"`
}

func (r *SplitPartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TopicExtendInfo struct {
	// 日志主题免鉴权配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnonymousAccess *AnonymousInfo `json:"AnonymousAccess,omitnil,omitempty" name:"AnonymousAccess"`
}

type TopicIdAndRegion struct {
	// 日志主题id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题id所在的地域id。
	// 
	// id,地域,简称信息如下：
	// - 1,   广州,ap-guangzhou
	// - 4,   上海,ap-shanghai
	// - 5,   中国香港,ap-hongkong
	// - 7,   上海金融,ap-shanghai-fsi
	// - 8,   北京,ap-beijing
	// - 9,   新加坡,ap-singapore
	// - 11,  深圳金融,ap-shenzhen-fsi
	// - 15,  硅谷,na-siliconvalley
	// - 16,  成都,ap-chengdu
	// - 17,  法兰克福,eu-frankfurt
	// - 18,  首尔,ap-seoul
	// - 19,  重庆,ap-chongqing
	// - 22,  弗吉尼亚,na-ashburn
	// - 23,  曼谷,ap-bangkok
	// - 25,  东京,ap-tokyo
	// - 33,  南京,ap-nanjing
	// - 46,  北京金融,ap-beijing-fsi
	// - 72,  雅加达,ap-jakarta
	// - 74,  圣保罗,sa-saopaulo
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

type TopicInfo struct {
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题分区个数
	PartitionCount *int64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// 主题是否开启索引（主题类型需为日志主题）
	Index *bool `json:"Index,omitnil,omitempty" name:"Index"`

	// AssumerUin非空则表示创建该日志主题的服务方Uin
	AssumerUin *uint64 `json:"AssumerUin,omitnil,omitempty" name:"AssumerUin"`

	// 云产品标识，主题由其它云产品创建时，该字段会显示云产品名称，例如CDN、TKE
	AssumerName *string `json:"AssumerName,omitnil,omitempty" name:"AssumerName"`

	// 创建时间。格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 主题是否开启采集，true：开启采集；false：关闭采集。
	// 创建日志主题时默认开启，可通过SDK调用ModifyTopic修改此字段。
	// 控制台目前不支持修改此参数。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 主题绑定的标签信息
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// RoleName非空则表示创建该日志主题的服务方使用的角色
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 该主题是否开启自动分裂
	AutoSplit *bool `json:"AutoSplit,omitnil,omitempty" name:"AutoSplit"`

	// 若开启自动分裂的话，该主题能够允许的最大分区数
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitnil,omitempty" name:"MaxSplitPartitions"`

	// 主题的存储类型
	// 
	// - hot: 标准存储
	// - cold: 低频存储
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 生命周期，单位天，可取值范围1~3600。取值为3640时代表永久保存
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 云产品二级标识，日志主题由其它云产品创建时，该字段会显示云产品名称及其日志类型的二级分类，例如TKE-Audit、TKE-Event。部分云产品仅有云产品标识(AssumerName)，无该字段。
	SubAssumerName *string `json:"SubAssumerName,omitnil,omitempty" name:"SubAssumerName"`

	// 主题描述
	Describes *string `json:"Describes,omitnil,omitempty" name:"Describes"`

	// 开启日志沉降，标准存储的生命周期， hotPeriod < Period。
	// 标准存储为 hotPeriod, 低频存储则为 Period-hotPeriod。（主题类型需为日志主题）
	// HotPeriod=0为没有开启日志沉降。
	HotPeriod *uint64 `json:"HotPeriod,omitnil,omitempty" name:"HotPeriod"`

	// kms-cls服务秘钥id
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 主题类型。
	// - 0: 日志主题 
	// - 1: 指标主题
	BizType *uint64 `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 免鉴权开关。 false：关闭； true：开启。
	// 开启后将支持指定操作匿名访问该日志主题。详情请参见[日志主题](https://cloud.tencent.com/document/product/614/41035)。
	IsWebTracking *bool `json:"IsWebTracking,omitnil,omitempty" name:"IsWebTracking"`

	// 日志主题扩展信息
	Extends *TopicExtendInfo `json:"Extends,omitnil,omitempty" name:"Extends"`

	// 异步迁移任务ID
	TopicAsyncTaskID *string `json:"TopicAsyncTaskID,omitnil,omitempty" name:"TopicAsyncTaskID"`

	// 异步迁移状态
	// - 1：进行中
	// - 2：已完成
	// - 3：失败
	// - 4：已取消
	MigrationStatus *uint64 `json:"MigrationStatus,omitnil,omitempty" name:"MigrationStatus"`

	// 异步迁移完成后，预计生效日期
	// 时间格式：yyyy-MM-dd HH:mm:ss
	EffectiveDate *string `json:"EffectiveDate,omitnil,omitempty" name:"EffectiveDate"`

	// IsSourceFrom 开启记录公网来源ip和服务端接收时间
	IsSourceFrom *bool `json:"IsSourceFrom,omitnil,omitempty" name:"IsSourceFrom"`
}

// Predefined struct for user
type UploadLogRequestParams struct {
	// 日志主题id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 该参数已废弃，请勿使用
	//
	// Deprecated: HashKey is deprecated.
	HashKey *string `json:"HashKey,omitnil,omitempty" name:"HashKey"`

	// 压缩方法，目前支持
	// - lz4
	// - zstd
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
}

type UploadLogRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题id
	// - 通过[获取日志主题列表](https://cloud.tencent.com/document/product/614/56454)获取日志主题Id。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 该参数已废弃，请勿使用
	HashKey *string `json:"HashKey,omitnil,omitempty" name:"HashKey"`

	// 压缩方法，目前支持
	// - lz4
	// - zstd
	CompressType *string `json:"CompressType,omitnil,omitempty" name:"CompressType"`
}

func (r *UploadLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "HashKey")
	delete(f, "CompressType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadLogResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadLogResponse struct {
	*tchttp.BaseResponse
	Response *UploadLogResponseParams `json:"Response"`
}

func (r *UploadLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserKafkaMeta struct {
	// 用户kafka version
	// 支持如下版本：
	//   - 0.10.2.0
	//   - 1.0.0
	//   - 2.0.0
	//   - 2.2.0
	//   - 2.4.0
	//   - 2.6.0
	//   - 2.7.0
	//   - 2.8.0
	//   - 3.0.0
	//   - 3.2.0
	KafkaVersion *string `json:"KafkaVersion,omitnil,omitempty" name:"KafkaVersion"`
}

type ValueInfo struct {
	// 字段类型，目前支持的类型有：long、text、double
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 字段的分词符，其中的每个字符代表一个分词符；
	// 仅支持英文符号、\n\t\r及转义符\；
	// long及double类型字段需为空；
	// 注意：\n\t\r本身已被转义，直接使用双引号包裹即可作为入参，无需再次转义。使用API Explorer进行调试时请使用JSON参数输入方式，以避免\n\t\r被重复转义
	Tokenizer *string `json:"Tokenizer,omitnil,omitempty" name:"Tokenizer"`

	// 字段是否开启分析功能
	SqlFlag *bool `json:"SqlFlag,omitnil,omitempty" name:"SqlFlag"`

	// 是否包含中文，long及double类型字段需为false
	ContainZH *bool `json:"ContainZH,omitnil,omitempty" name:"ContainZH"`

	// 字段别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type WebCallback struct {
	// 回调的类型。可选值：
	// - Http(自定义接口回调)
	// - WeCom(企业微信)
	// - DingTalk(钉钉)
	// - Lark(飞书)
	CallbackType *string `json:"CallbackType,omitnil,omitempty" name:"CallbackType"`

	// 回调地址，最大支持1024个字节。
	// 也可使用WebCallbackId引用集成配置中的URL，此时该字段请填写为空字符串。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 集成配置ID。-通过[获取告警渠道回调配置列表](https://cloud.tencent.com/document/product/614/115229)获取集成配置ID
	WebCallbackId *string `json:"WebCallbackId,omitnil,omitempty" name:"WebCallbackId"`

	// 回调方法。可选值：
	// - POST（默认值）
	// - PUT
	// 
	// 注意：
	// - 参数CallbackType为Http时为必选，其它回调方式无需填写。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 通知内容模板ID，使用Default-zh引用默认模板（中文），使用Default-en引用DefaultTemplate(English)。
	NoticeContentId *string `json:"NoticeContentId,omitnil,omitempty" name:"NoticeContentId"`

	// 提醒类型。
	// 
	// 0：不提醒；1：指定人；2：所有人
	RemindType *uint64 `json:"RemindType,omitnil,omitempty" name:"RemindType"`

	// 电话列表。
	Mobiles []*string `json:"Mobiles,omitnil,omitempty" name:"Mobiles"`

	// 用户ID列表。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 该参数已废弃，请使用NoticeContentId。
	//
	// Deprecated: Headers is deprecated.
	Headers []*string `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 该参数已废弃，请使用NoticeContentId。
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Body is deprecated.
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 序号。
	// - 入参无效。
	// - 出参有效。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

type WebCallbackInfo struct {
	// 告警渠道回调配置id。
	WebCallbackId *string `json:"WebCallbackId,omitnil,omitempty" name:"WebCallbackId"`

	// 告警渠道回调配置名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 渠道类型
	// 
	// WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调;
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 回调地址。
	Webhook *string `json:"Webhook,omitnil,omitempty" name:"Webhook"`

	// 请求方式。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 秘钥信息。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 主账号。
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号。
	SubUin *uint64 `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 创建时间。秒(s)级时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。秒(s)级时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}