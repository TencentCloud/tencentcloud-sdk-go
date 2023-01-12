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

package v20201101

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ABTestConfig struct {
	// 灰度项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// true：正在灰度，false：不在灰度
	Status *bool `json:"Status,omitempty" name:"Status"`
}

type AbnormalProcessChildRuleInfo struct {
	// 策略模式，   RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截
	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`

	// 进程路径
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 子策略id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 威胁等级，HIGH:高，MIDDLE:中，LOW:低
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleLevel *string `json:"RuleLevel,omitempty" name:"RuleLevel"`
}

type AbnormalProcessEventDescription struct {
	// 事件规则
	Description *string `json:"Description,omitempty" name:"Description"`

	// 解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 命中规则详细信息
	MatchRule *AbnormalProcessChildRuleInfo `json:"MatchRule,omitempty" name:"MatchRule"`

	// 命中规则名称，PROXY_TOOL：代理软件，TRANSFER_CONTROL：横向渗透，ATTACK_CMD：恶意命令，REVERSE_SHELL：反弹shell，FILELESS：无文件程序执行，RISK_CMD：高危命令，ABNORMAL_CHILD_PROC：敏感服务异常子进程启动，USER_DEFINED_RULE：用户自定义规则
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 命中规则的id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 事件最后一次处理的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`

	// 命中策略名称：SYSTEM_DEFINED_RULE （系统策略）或  用户自定义的策略名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type AbnormalProcessEventInfo struct {
	// 进程目录
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 事件类型，MALICE_PROCESS_START:恶意进程启动
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 命中规则名称，PROXY_TOOL：代理软件，TRANSFER_CONTROL：横向渗透，ATTACK_CMD：恶意命令，REVERSE_SHELL：反弹shell，FILELESS：无文件程序执行，RISK_CMD：高危命令，ABNORMAL_CHILD_PROC：敏感服务异常子进程启动，USER_DEFINED_RULE：用户自定义规则
	MatchRuleName *string `json:"MatchRuleName,omitempty" name:"MatchRuleName"`

	// 生成时间
	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 镜像名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 动作执行结果，    BEHAVIOR_NONE: 无
	//     BEHAVIOR_ALERT: 告警
	//     BEHAVIOR_RELEASE：放行
	//     BEHAVIOR_HOLDUP_FAILED:拦截失败
	//     BEHAVIOR_HOLDUP_SUCCESSED：拦截失败
	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`

	// 状态，EVENT_UNDEAL:事件未处理
	//     EVENT_DEALED:事件已经处理
	//     EVENT_INGNORE：事件已经忽略
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件记录的唯一id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 镜像id，用于跳转
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 容器id，用于跳转
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 事件解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件详细描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 命中策略id
	MatchRuleId *string `json:"MatchRuleId,omitempty" name:"MatchRuleId"`

	// 命中规则行为：
	// RULE_MODE_RELEASE 放行
	// RULE_MODE_ALERT  告警
	// RULE_MODE_HOLDUP 拦截
	MatchAction *string `json:"MatchAction,omitempty" name:"MatchAction"`

	// 命中规则进程信息
	MatchProcessPath *string `json:"MatchProcessPath,omitempty" name:"MatchProcessPath"`

	// 规则是否存在
	RuleExist *bool `json:"RuleExist,omitempty" name:"RuleExist"`

	// 事件数量
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 最近生成时间
	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`

	// 规则组Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 命中策略名称：SYSTEM_DEFINED_RULE （系统策略）或  用户自定义的策略名字
	MatchGroupName *string `json:"MatchGroupName,omitempty" name:"MatchGroupName"`

	// 命中规则等级，HIGH：高危，MIDDLE：中危，LOW：低危。
	MatchRuleLevel *string `json:"MatchRuleLevel,omitempty" name:"MatchRuleLevel"`

	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`

	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING
	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type AbnormalProcessEventTendencyInfo struct {
	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 待处理代理软件事件数
	ProxyToolEventCount *int64 `json:"ProxyToolEventCount,omitempty" name:"ProxyToolEventCount"`

	// 待处理横向参透事件数
	TransferControlEventCount *int64 `json:"TransferControlEventCount,omitempty" name:"TransferControlEventCount"`

	// 待处理恶意命令事件数
	AttackCmdEventCount *int64 `json:"AttackCmdEventCount,omitempty" name:"AttackCmdEventCount"`

	// 待处理反弹shell事件数
	ReverseShellEventCount *int64 `json:"ReverseShellEventCount,omitempty" name:"ReverseShellEventCount"`

	// 待处理无文件程序执行事件数
	FilelessEventCount *int64 `json:"FilelessEventCount,omitempty" name:"FilelessEventCount"`

	// 待处理高危命令事件数
	RiskCmdEventCount *int64 `json:"RiskCmdEventCount,omitempty" name:"RiskCmdEventCount"`

	// 待处理敏感服务异常子进程启动事件数
	AbnormalChildProcessEventCount *int64 `json:"AbnormalChildProcessEventCount,omitempty" name:"AbnormalChildProcessEventCount"`

	// 待处理自定义规则事件数
	UserDefinedRuleEventCount *int64 `json:"UserDefinedRuleEventCount,omitempty" name:"UserDefinedRuleEventCount"`
}

type AbnormalProcessRuleInfo struct {
	// true:策略启用，false:策略禁用
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`

	// 生效镜像id，空数组代表全部镜像
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// 用户策略的子策略数组
	ChildRules []*AbnormalProcessChildRuleInfo `json:"ChildRules,omitempty" name:"ChildRules"`

	// 策略名字
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 策略id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 系统策略的子策略数组
	SystemChildRules []*AbnormalProcessSystemChildRuleInfo `json:"SystemChildRules,omitempty" name:"SystemChildRules"`

	// 是否是系统默认策略
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type AbnormalProcessSystemChildRuleInfo struct {
	// 子策略Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 子策略状态，true为开启，false为关闭
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`

	// 策略模式,  RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截
	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`

	// 子策略检测的行为类型
	// PROXY_TOOL： 代理软件
	// TRANSFER_CONTROL：横向渗透
	// ATTACK_CMD： 恶意命令
	// REVERSE_SHELL：反弹shell
	// FILELESS：无文件程序执行
	// RISK_CMD：高危命令
	// ABNORMAL_CHILD_PROC: 敏感服务异常子进程启动
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 威胁等级，HIGH:高，MIDDLE:中，LOW:低
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleLevel *string `json:"RuleLevel,omitempty" name:"RuleLevel"`
}

type AccessControlChildRuleInfo struct {
	// 策略模式,  RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截
	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`

	// 进程路径
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 被访问文件路径，仅仅在访问控制生效
	TargetFilePath *string `json:"TargetFilePath,omitempty" name:"TargetFilePath"`

	// 子策略id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type AccessControlEventDescription struct {
	// 事件规则
	Description *string `json:"Description,omitempty" name:"Description"`

	// 解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 命中规则详细信息
	MatchRule *AccessControlChildRuleInfo `json:"MatchRule,omitempty" name:"MatchRule"`

	// 命中规则名字
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 命中规则id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 事件最后一次处理的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type AccessControlEventInfo struct {
	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 命中规则名称
	MatchRuleName *string `json:"MatchRuleName,omitempty" name:"MatchRuleName"`

	// 生成时间
	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 镜像名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 动作执行结果，   BEHAVIOR_NONE: 无
	//     BEHAVIOR_ALERT: 告警
	//     BEHAVIOR_RELEASE：放行
	//     BEHAVIOR_HOLDUP_FAILED:拦截失败
	//     BEHAVIOR_HOLDUP_SUCCESSED：拦截失败
	Behavior *string `json:"Behavior,omitempty" name:"Behavior"`

	// 状态0:未处理  “EVENT_UNDEAL”:事件未处理
	//     "EVENT_DEALED":事件已经处理
	//     "EVENT_INGNORE"：事件已经忽略
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件记录的唯一id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 事件类型， FILE_ABNORMAL_READ:文件异常读取
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 镜像id, 用于跳转
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 容器id, 用于跳转
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 事件解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件详细描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 命中策略id
	MatchRuleId *string `json:"MatchRuleId,omitempty" name:"MatchRuleId"`

	// 命中规则行为：
	// RULE_MODE_RELEASE 放行
	// RULE_MODE_ALERT  告警
	// RULE_MODE_HOLDUP 拦截
	MatchAction *string `json:"MatchAction,omitempty" name:"MatchAction"`

	// 命中规则进程信息
	MatchProcessPath *string `json:"MatchProcessPath,omitempty" name:"MatchProcessPath"`

	// 命中规则文件信息
	MatchFilePath *string `json:"MatchFilePath,omitempty" name:"MatchFilePath"`

	// 文件路径，包含名字
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 规则是否存在
	RuleExist *bool `json:"RuleExist,omitempty" name:"RuleExist"`

	// 事件数量
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 最近生成时间
	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`

	// 规则组id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`

	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING
	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type AccessControlRuleInfo struct {
	// 开关,true:开启，false:禁用
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`

	// 生效惊现id，空数组代表全部镜像
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// 用户策略的子策略数组
	ChildRules []*AccessControlChildRuleInfo `json:"ChildRules,omitempty" name:"ChildRules"`

	// 策略名字
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 策略id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 系统策略的子策略数组
	SystemChildRules []*AccessControlSystemChildRuleInfo `json:"SystemChildRules,omitempty" name:"SystemChildRules"`

	// 是否是系统默认策略
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type AccessControlSystemChildRuleInfo struct {
	// 子策略Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 策略模式,  RULE_MODE_RELEASE: 放行
	//    RULE_MODE_ALERT: 告警
	//    RULE_MODE_HOLDUP:拦截
	RuleMode *string `json:"RuleMode,omitempty" name:"RuleMode"`

	// 子策略状态，true为开启，false为关闭
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`

	// 子策略检测的入侵行为类型
	// CHANGE_CRONTAB：篡改计划任务
	// CHANGE_SYS_BIN：篡改系统程序
	// CHANGE_USRCFG：篡改用户配置
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

// Predefined struct for user
type AddAndPublishNetworkFirewallPolicyDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 出站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// pod选择器
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义规则
	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

type AddAndPublishNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 出站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// pod选择器
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义规则
	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

func (r *AddAndPublishNetworkFirewallPolicyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAndPublishNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PolicyName")
	delete(f, "FromPolicyRule")
	delete(f, "ToPolicyRule")
	delete(f, "PodSelector")
	delete(f, "Namespace")
	delete(f, "Description")
	delete(f, "CustomPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAndPublishNetworkFirewallPolicyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAndPublishNetworkFirewallPolicyDetailResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddAndPublishNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse
	Response *AddAndPublishNetworkFirewallPolicyDetailResponseParams `json:"Response"`
}

func (r *AddAndPublishNetworkFirewallPolicyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAndPublishNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAndPublishNetworkFirewallPolicyYamlDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type AddAndPublishNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *AddAndPublishNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAndPublishNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PolicyName")
	delete(f, "Yaml")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAndPublishNetworkFirewallPolicyYamlDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAndPublishNetworkFirewallPolicyYamlDetailResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddAndPublishNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse
	Response *AddAndPublishNetworkFirewallPolicyYamlDetailResponseParams `json:"Response"`
}

func (r *AddAndPublishNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAndPublishNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAssetImageRegistryRegistryDetailRequestParams struct {
	// 仓库名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 仓库url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 仓库类型，列表：harbor
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 网络类型，列表：public（公网）
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 仓库版本
	RegistryVersion *string `json:"RegistryVersion,omitempty" name:"RegistryVersion"`

	// 区域，列表：default（默认）
	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`

	// 限速
	SpeedLimit *int64 `json:"SpeedLimit,omitempty" name:"SpeedLimit"`

	// 安全模式（证书校验）：0（默认） 非安全模式（跳过证书校验）：1
	Insecure *uint64 `json:"Insecure,omitempty" name:"Insecure"`
}

type AddAssetImageRegistryRegistryDetailRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 仓库url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 仓库类型，列表：harbor
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 网络类型，列表：public（公网）
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 仓库版本
	RegistryVersion *string `json:"RegistryVersion,omitempty" name:"RegistryVersion"`

	// 区域，列表：default（默认）
	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`

	// 限速
	SpeedLimit *int64 `json:"SpeedLimit,omitempty" name:"SpeedLimit"`

	// 安全模式（证书校验）：0（默认） 非安全模式（跳过证书校验）：1
	Insecure *uint64 `json:"Insecure,omitempty" name:"Insecure"`
}

func (r *AddAssetImageRegistryRegistryDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAssetImageRegistryRegistryDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "Url")
	delete(f, "RegistryType")
	delete(f, "NetType")
	delete(f, "RegistryVersion")
	delete(f, "RegistryRegion")
	delete(f, "SpeedLimit")
	delete(f, "Insecure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAssetImageRegistryRegistryDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAssetImageRegistryRegistryDetailResponseParams struct {
	// 连接错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckErr *string `json:"HealthCheckErr,omitempty" name:"HealthCheckErr"`

	// 名称错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameRepeatErr *string `json:"NameRepeatErr,omitempty" name:"NameRepeatErr"`

	// 仓库唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *AddAssetImageRegistryRegistryDetailResponseParams `json:"Response"`
}

func (r *AddAssetImageRegistryRegistryDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAssetImageRegistryRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddComplianceAssetPolicySetToWhitelistRequestParams struct {
	// 资产ID+检查项IDs. 列表
	AssetPolicySetList []*ComplianceAssetPolicySetItem `json:"AssetPolicySetList,omitempty" name:"AssetPolicySetList"`
}

type AddComplianceAssetPolicySetToWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// 资产ID+检查项IDs. 列表
	AssetPolicySetList []*ComplianceAssetPolicySetItem `json:"AssetPolicySetList,omitempty" name:"AssetPolicySetList"`
}

func (r *AddComplianceAssetPolicySetToWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddComplianceAssetPolicySetToWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetPolicySetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddComplianceAssetPolicySetToWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddComplianceAssetPolicySetToWhitelistResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddComplianceAssetPolicySetToWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *AddComplianceAssetPolicySetToWhitelistResponseParams `json:"Response"`
}

func (r *AddComplianceAssetPolicySetToWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddComplianceAssetPolicySetToWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCompliancePolicyAssetSetToWhitelistRequestParams struct {
	// 检查项ID
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 需要忽略指定检查项内的资产ID列表
	CustomerAssetItemIdSet []*uint64 `json:"CustomerAssetItemIdSet,omitempty" name:"CustomerAssetItemIdSet"`
}

type AddCompliancePolicyAssetSetToWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// 检查项ID
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 需要忽略指定检查项内的资产ID列表
	CustomerAssetItemIdSet []*uint64 `json:"CustomerAssetItemIdSet,omitempty" name:"CustomerAssetItemIdSet"`
}

func (r *AddCompliancePolicyAssetSetToWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCompliancePolicyAssetSetToWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerPolicyItemId")
	delete(f, "CustomerAssetItemIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCompliancePolicyAssetSetToWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCompliancePolicyAssetSetToWhitelistResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddCompliancePolicyAssetSetToWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *AddCompliancePolicyAssetSetToWhitelistResponseParams `json:"Response"`
}

func (r *AddCompliancePolicyAssetSetToWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCompliancePolicyAssetSetToWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCompliancePolicyItemToWhitelistRequestParams struct {
	// 要忽略的检测项的ID的列表
	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

type AddCompliancePolicyItemToWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// 要忽略的检测项的ID的列表
	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

func (r *AddCompliancePolicyItemToWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCompliancePolicyItemToWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerPolicyItemIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCompliancePolicyItemToWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCompliancePolicyItemToWhitelistResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddCompliancePolicyItemToWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *AddCompliancePolicyItemToWhitelistResponseParams `json:"Response"`
}

func (r *AddCompliancePolicyItemToWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCompliancePolicyItemToWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditAbnormalProcessRuleRequestParams struct {
	// 增加策略信息，策略id为空，编辑策略是id不能为空
	RuleInfo *AbnormalProcessRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

	// 仅在加白的时候带上
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type AddEditAbnormalProcessRuleRequest struct {
	*tchttp.BaseRequest
	
	// 增加策略信息，策略id为空，编辑策略是id不能为空
	RuleInfo *AbnormalProcessRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

	// 仅在加白的时候带上
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *AddEditAbnormalProcessRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditAbnormalProcessRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleInfo")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEditAbnormalProcessRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditAbnormalProcessRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddEditAbnormalProcessRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddEditAbnormalProcessRuleResponseParams `json:"Response"`
}

func (r *AddEditAbnormalProcessRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditAbnormalProcessRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditAccessControlRuleRequestParams struct {
	// 增加策略信息，策略id为空，编辑策略是id不能为空
	RuleInfo *AccessControlRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

	// 仅在白名单时候使用
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type AddEditAccessControlRuleRequest struct {
	*tchttp.BaseRequest
	
	// 增加策略信息，策略id为空，编辑策略是id不能为空
	RuleInfo *AccessControlRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

	// 仅在白名单时候使用
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *AddEditAccessControlRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditAccessControlRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleInfo")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEditAccessControlRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditAccessControlRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddEditAccessControlRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddEditAccessControlRuleResponseParams `json:"Response"`
}

func (r *AddEditAccessControlRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditImageAutoAuthorizedRuleRequestParams struct {
	// 授权范围类别，MANUAL:自选主机节点，ALL:全部镜像
	RangeType *string `json:"RangeType,omitempty" name:"RangeType"`

	// 每天最大的镜像授权数限制, 0表示无限制
	MaxDailyCount *int64 `json:"MaxDailyCount,omitempty" name:"MaxDailyCount"`

	// 规则是否生效，0:不生效，1:已生效
	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 自选主机id，当授权范围为MANUAL:自选主机时且HostIdFilters为空时，必填
	HostIdSet []*string `json:"HostIdSet,omitempty" name:"HostIdSet"`

	// 规则id，在编辑时，必填
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 根据条件过滤，当授权范围为MANUAL:自选主机时且HostIdSet为空时，必填
	HostIdFilters []*AssetFilters `json:"HostIdFilters,omitempty" name:"HostIdFilters"`

	// 根据条件过滤而且排除指定主机id
	ExcludeHostIdSet []*string `json:"ExcludeHostIdSet,omitempty" name:"ExcludeHostIdSet"`
}

type AddEditImageAutoAuthorizedRuleRequest struct {
	*tchttp.BaseRequest
	
	// 授权范围类别，MANUAL:自选主机节点，ALL:全部镜像
	RangeType *string `json:"RangeType,omitempty" name:"RangeType"`

	// 每天最大的镜像授权数限制, 0表示无限制
	MaxDailyCount *int64 `json:"MaxDailyCount,omitempty" name:"MaxDailyCount"`

	// 规则是否生效，0:不生效，1:已生效
	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 自选主机id，当授权范围为MANUAL:自选主机时且HostIdFilters为空时，必填
	HostIdSet []*string `json:"HostIdSet,omitempty" name:"HostIdSet"`

	// 规则id，在编辑时，必填
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 根据条件过滤，当授权范围为MANUAL:自选主机时且HostIdSet为空时，必填
	HostIdFilters []*AssetFilters `json:"HostIdFilters,omitempty" name:"HostIdFilters"`

	// 根据条件过滤而且排除指定主机id
	ExcludeHostIdSet []*string `json:"ExcludeHostIdSet,omitempty" name:"ExcludeHostIdSet"`
}

func (r *AddEditImageAutoAuthorizedRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditImageAutoAuthorizedRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RangeType")
	delete(f, "MaxDailyCount")
	delete(f, "IsEnabled")
	delete(f, "HostIdSet")
	delete(f, "RuleId")
	delete(f, "HostIdFilters")
	delete(f, "ExcludeHostIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEditImageAutoAuthorizedRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditImageAutoAuthorizedRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddEditImageAutoAuthorizedRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddEditImageAutoAuthorizedRuleResponseParams `json:"Response"`
}

func (r *AddEditImageAutoAuthorizedRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditImageAutoAuthorizedRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditReverseShellWhiteListRequestParams struct {
	// 增加或编辑白名单信息。新增白名单时WhiteListInfo.id为空，编辑白名单WhiteListInfo.id不能为空。
	WhiteListInfo *ReverseShellWhiteListInfo `json:"WhiteListInfo,omitempty" name:"WhiteListInfo"`

	// 仅在添加事件白名单时候使用
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type AddEditReverseShellWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 增加或编辑白名单信息。新增白名单时WhiteListInfo.id为空，编辑白名单WhiteListInfo.id不能为空。
	WhiteListInfo *ReverseShellWhiteListInfo `json:"WhiteListInfo,omitempty" name:"WhiteListInfo"`

	// 仅在添加事件白名单时候使用
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *AddEditReverseShellWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditReverseShellWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteListInfo")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEditReverseShellWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditReverseShellWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddEditReverseShellWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *AddEditReverseShellWhiteListResponseParams `json:"Response"`
}

func (r *AddEditReverseShellWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditReverseShellWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditRiskSyscallWhiteListRequestParams struct {
	// 仅在添加事件白名单时候使用
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 增加或编辑白名单信。新增白名单时WhiteListInfo.id为空，编辑白名单WhiteListInfo.id不能为空.
	WhiteListInfo *RiskSyscallWhiteListInfo `json:"WhiteListInfo,omitempty" name:"WhiteListInfo"`
}

type AddEditRiskSyscallWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 仅在添加事件白名单时候使用
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 增加或编辑白名单信。新增白名单时WhiteListInfo.id为空，编辑白名单WhiteListInfo.id不能为空.
	WhiteListInfo *RiskSyscallWhiteListInfo `json:"WhiteListInfo,omitempty" name:"WhiteListInfo"`
}

func (r *AddEditRiskSyscallWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditRiskSyscallWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "WhiteListInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEditRiskSyscallWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditRiskSyscallWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddEditRiskSyscallWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *AddEditRiskSyscallWhiteListResponseParams `json:"Response"`
}

func (r *AddEditRiskSyscallWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditRiskSyscallWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditWarningRulesRequestParams struct {
	// 告警开关策略
	WarningRules []*WarningRule `json:"WarningRules,omitempty" name:"WarningRules"`
}

type AddEditWarningRulesRequest struct {
	*tchttp.BaseRequest
	
	// 告警开关策略
	WarningRules []*WarningRule `json:"WarningRules,omitempty" name:"WarningRules"`
}

func (r *AddEditWarningRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditWarningRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WarningRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEditWarningRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEditWarningRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddEditWarningRulesResponse struct {
	*tchttp.BaseResponse
	Response *AddEditWarningRulesResponseParams `json:"Response"`
}

func (r *AddEditWarningRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEditWarningRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEscapeWhiteListRequestParams struct {
	// 加白名单事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸
	EventType []*string `json:"EventType,omitempty" name:"EventType"`

	// 加白名单镜像ID数组
	ImageIDs []*string `json:"ImageIDs,omitempty" name:"ImageIDs"`
}

type AddEscapeWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 加白名单事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸
	EventType []*string `json:"EventType,omitempty" name:"EventType"`

	// 加白名单镜像ID数组
	ImageIDs []*string `json:"ImageIDs,omitempty" name:"ImageIDs"`
}

func (r *AddEscapeWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEscapeWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventType")
	delete(f, "ImageIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEscapeWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEscapeWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddEscapeWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *AddEscapeWhiteListResponseParams `json:"Response"`
}

func (r *AddEscapeWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEscapeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddIgnoreVulRequestParams struct {
	// 漏洞PocID信息列表
	List []*ModifyIgnoreVul `json:"List,omitempty" name:"List"`
}

type AddIgnoreVulRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞PocID信息列表
	List []*ModifyIgnoreVul `json:"List,omitempty" name:"List"`
}

func (r *AddIgnoreVulRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddIgnoreVulRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "List")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddIgnoreVulRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddIgnoreVulResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddIgnoreVulResponse struct {
	*tchttp.BaseResponse
	Response *AddIgnoreVulResponseParams `json:"Response"`
}

func (r *AddIgnoreVulResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddIgnoreVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNetworkFirewallPolicyDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 出站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// pod选择器
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义规则
	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

type AddNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 出站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// pod选择器
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义规则
	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

func (r *AddNetworkFirewallPolicyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PolicyName")
	delete(f, "FromPolicyRule")
	delete(f, "ToPolicyRule")
	delete(f, "PodSelector")
	delete(f, "Namespace")
	delete(f, "Description")
	delete(f, "CustomPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNetworkFirewallPolicyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNetworkFirewallPolicyDetailResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse
	Response *AddNetworkFirewallPolicyDetailResponseParams `json:"Response"`
}

func (r *AddNetworkFirewallPolicyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNetworkFirewallPolicyYamlDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type AddNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *AddNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PolicyName")
	delete(f, "Yaml")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNetworkFirewallPolicyYamlDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNetworkFirewallPolicyYamlDetailResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse
	Response *AddNetworkFirewallPolicyYamlDetailResponseParams `json:"Response"`
}

func (r *AddNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AffectedNodeItem struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名字
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 内网ip地址
	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 节点的角色，Master、Work等
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// k8s版本
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 运行时组件,docker或者containerd
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 检查结果的验证信息
	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
}

type AffectedWorkloadItem struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名字
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 工作负载名称
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`

	// 工作负载类型
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 检测结果的验证信息
	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
}

type AssetClusterListItem struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群状态
	// CSR_RUNNING: 运行中
	// CSR_EXCEPTION:异常
	// CSR_DEL:已经删除
	Status *string `json:"Status,omitempty" name:"Status"`

	// 绑定规则名称
	BindRuleName *string `json:"BindRuleName,omitempty" name:"BindRuleName"`

	// 集群类型:
	// CT_TKE: TKE集群
	// CT_USER_CREATE: 用户自建集群
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type AssetFilters struct {
	// 过滤键的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否模糊查询
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type AssetSimpleImageInfo struct {
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 关联容器个数
	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`

	// 最后扫描时间
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 镜像大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type AutoAuthorizedImageInfo struct {
	// 镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 授权时间
	AuthorizedTime *string `json:"AuthorizedTime,omitempty" name:"AuthorizedTime"`

	// 授权结果，SUCCESS:成功，REACH_LIMIT:达到授权上限，LICENSE_INSUFFICIENT:授权数不足'
	Status *string `json:"Status,omitempty" name:"Status"`

	// 是否授权，1：是，0：否
	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
}

type AutoAuthorizedRuleHostInfo struct {
	// 主机id
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 主机ip即内网ip
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 镜像个数
	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`

	// 容器个数
	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`

	// 外网ip
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 主机实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// docker 版本
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

	// agent运行状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type CKafkaInstanceInfo struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 主题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*CKafkaTopicInfo `json:"TopicList,omitempty" name:"TopicList"`

	// 路由列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteList []*CkafkaRouteInfo `json:"RouteList,omitempty" name:"RouteList"`

	// kafka版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaVersion *string `json:"KafkaVersion,omitempty" name:"KafkaVersion"`
}

type CKafkaTopicInfo struct {
	// 主题ID
	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

// Predefined struct for user
type CheckNetworkFirewallPolicyYamlRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CheckNetworkFirewallPolicyYamlRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CheckNetworkFirewallPolicyYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckNetworkFirewallPolicyYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "PolicyName")
	delete(f, "Yaml")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckNetworkFirewallPolicyYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckNetworkFirewallPolicyYamlResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckNetworkFirewallPolicyYamlResponse struct {
	*tchttp.BaseResponse
	Response *CheckNetworkFirewallPolicyYamlResponseParams `json:"Response"`
}

func (r *CheckNetworkFirewallPolicyYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckNetworkFirewallPolicyYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckRepeatAssetImageRegistryRequestParams struct {
	// 仓库名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CheckRepeatAssetImageRegistryRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckRepeatAssetImageRegistryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRepeatAssetImageRegistryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckRepeatAssetImageRegistryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckRepeatAssetImageRegistryResponseParams struct {
	// 是否重复
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRepeat *bool `json:"IsRepeat,omitempty" name:"IsRepeat"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckRepeatAssetImageRegistryResponse struct {
	*tchttp.BaseResponse
	Response *CheckRepeatAssetImageRegistryResponseParams `json:"Response"`
}

func (r *CheckRepeatAssetImageRegistryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRepeatAssetImageRegistryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CkafkaRouteInfo struct {
	// 路由ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteID *int64 `json:"RouteID,omitempty" name:"RouteID"`

	// 域名名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainPort *uint64 `json:"DomainPort,omitempty" name:"DomainPort"`

	// 虚拟ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 虚拟ip类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipType *int64 `json:"VipType,omitempty" name:"VipType"`

	// 接入类型
	// // 0：PLAINTEXT (明文方式，没有带用户信息老版本及社区版本都支持)
	// 	// 1：SASL_PLAINTEXT（明文方式，不过在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）
	// 	// 2：SSL（SSL加密通信，没有带用户信息，老版本及社区版本都支持）
	// 	// 3：SASL_SSL（SSL加密通信，在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`
}

type ClsLogsetInfo struct {
	// 日志集ID
	LogsetID *string `json:"LogsetID,omitempty" name:"LogsetID"`

	// 日志集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// cls主题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*ClsTopicInfo `json:"TopicList,omitempty" name:"TopicList"`
}

type ClsTopicInfo struct {
	// 主题ID
	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ClusterCheckItem struct {
	// 唯一的检测项的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`

	// 风险项的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 检测项详细描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemDetail *string `json:"ItemDetail,omitempty" name:"ItemDetail"`

	// 威胁等级。严重Serious,高危High,中危Middle,提示Hint
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 检查对象、风险对象.Runc,Kubelet,Containerd,Pods
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskTarget *string `json:"RiskTarget,omitempty" name:"RiskTarget"`

	// 风险类别,漏洞风险CVERisk,配置风险ConfigRisk
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType *string `json:"RiskType,omitempty" name:"RiskType"`

	// 检测项所属的风险类型,权限提升:PrivilegePromotion,拒绝服务:RefuseService,目录穿越:DirectoryEscape,未授权访问:UnauthorizedAccess,权限许可和访问控制问题:PrivilegeAndAccessControl,敏感信息泄露:SensitiveInfoLeak
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskAttribute *string `json:"RiskAttribute,omitempty" name:"RiskAttribute"`

	// 风险特征,Tag.存在EXP:ExistEXP,存在POD:ExistPOC,无需重启:NoNeedReboot, 服务重启:ServerRestart,远程信息泄露:RemoteInfoLeak,远程拒绝服务:RemoteRefuseService,远程利用:RemoteExploit,远程执行:RemoteExecute
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskProperty *string `json:"RiskProperty,omitempty" name:"RiskProperty"`

	// CVE编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVENumber *string `json:"CVENumber,omitempty" name:"CVENumber"`

	// 披露时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscoverTime *string `json:"DiscoverTime,omitempty" name:"DiscoverTime"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// CVSS信息,用于画图
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSS *string `json:"CVSS,omitempty" name:"CVSS"`

	// CVSS分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSSScore *string `json:"CVSSScore,omitempty" name:"CVSSScore"`

	// 参考连接
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelateLink *string `json:"RelateLink,omitempty" name:"RelateLink"`

	// 影响类型，为Node或者Workload
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectedType *string `json:"AffectedType,omitempty" name:"AffectedType"`

	// 受影响的版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectedVersion *string `json:"AffectedVersion,omitempty" name:"AffectedVersion"`

	// 忽略的资产数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoredAssetNum *int64 `json:"IgnoredAssetNum,omitempty" name:"IgnoredAssetNum"`

	// 是否忽略该检测项
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsIgnored *bool `json:"IsIgnored,omitempty" name:"IsIgnored"`

	// 受影响评估
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskAssessment *string `json:"RiskAssessment,omitempty" name:"RiskAssessment"`
}

type ClusterCheckTaskItem struct {
	// 指定要扫描的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群所属地域
	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`

	// 指定要扫描的节点IP
	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`

	// 按照要扫描的workload名字
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`
}

type ClusterCreateComponentItem struct {
	// 要安装组件的集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 该集群对应的地域
	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

type ClusterInfoItem struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名字
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 集群操作系统
	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群节点数
	ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`

	// 集群区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 监控组件的状态，为Defender_Uninstall、Defender_Normal、Defender_Error、Defender_Installing
	DefenderStatus *string `json:"DefenderStatus,omitempty" name:"DefenderStatus"`

	// 集群状态
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 集群的检测模式，为Cluster_Normal或者Cluster_Actived.
	ClusterCheckMode *string `json:"ClusterCheckMode,omitempty" name:"ClusterCheckMode"`

	// 是否自动定期检测
	ClusterAutoCheck *bool `json:"ClusterAutoCheck,omitempty" name:"ClusterAutoCheck"`

	// 防护容器部署失败原因，为UserDaemonSetNotReady时,和UnreadyNodeNum转成"N个节点防御容器为就绪"，其他错误直接展示
	DefenderErrorReason *string `json:"DefenderErrorReason,omitempty" name:"DefenderErrorReason"`

	// 防御容器没有ready状态的节点数量
	UnreadyNodeNum *uint64 `json:"UnreadyNodeNum,omitempty" name:"UnreadyNodeNum"`

	// 严重风险检查项的数量
	SeriousRiskCount *int64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`

	// 高风险检查项的数量
	HighRiskCount *int64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`

	// 中风险检查项的数量
	MiddleRiskCount *int64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`

	// 提示风险检查项的数量
	HintRiskCount *int64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`

	// 检查失败原因
	CheckFailReason *string `json:"CheckFailReason,omitempty" name:"CheckFailReason"`

	// 检查状态,为Task_Running, NoRisk, HasRisk, Uncheck, Task_Error
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 任务创建时间,检查时间
	TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`
}

type ClusterRiskItem struct {
	// 检测项相关信息
	CheckItem *ClusterCheckItem `json:"CheckItem,omitempty" name:"CheckItem"`

	// 验证信息
	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`

	// 事件描述,检查的错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 受影响的集群数量
	AffectedClusterCount *uint64 `json:"AffectedClusterCount,omitempty" name:"AffectedClusterCount"`

	// 受影响的节点数量
	AffectedNodeCount *uint64 `json:"AffectedNodeCount,omitempty" name:"AffectedNodeCount"`
}

type ComplianceAffectedAsset struct {
	// 为客户分配的唯一的资产项的ID。
	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`

	// 资产项的名称。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 资产项的类型
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 检测状态
	// 
	// CHECK_INIT, 待检测
	// 
	// CHECK_RUNNING, 检测中
	// 
	// CHECK_FINISHED, 检测完成
	// 
	// CHECK_FAILED, 检测失败
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 节点名称。
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 上次检测的时间，格式为”YYYY-MM-DD HH:m::SS“。
	// 
	// 如果没有检测过，此处为”0000-00-00 00:00:00“。
	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`

	// 检测结果。取值为：
	// 
	// RESULT_FAILED: 未通过
	// 
	// RESULT_PASSED: 通过
	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`

	// 主机IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 镜像的tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// 检查项验证信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`

	// 主机实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ComplianceAssetDetailInfo struct {
	// 客户资产的ID。
	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`

	// 资产类别。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 资产的名称。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 资产所属的节点的名称。
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 资产所在的主机的名称。
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 资产所在的主机的IP。
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 检测状态
	// CHECK_INIT, 待检测
	// CHECK_RUNNING, 检测中
	// CHECK_FINISHED, 检测完成
	// CHECK_FAILED, 检测失败
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 此类资产通过的检测项的数目。
	PassedPolicyItemCount *uint64 `json:"PassedPolicyItemCount,omitempty" name:"PassedPolicyItemCount"`

	// 此类资产未通过的检测的数目。
	FailedPolicyItemCount *uint64 `json:"FailedPolicyItemCount,omitempty" name:"FailedPolicyItemCount"`

	// 上次检测的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`

	// 检测结果：
	// RESULT_FAILED: 未通过。
	// RESULT_PASSED: 通过。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`

	// 资产的运行状态。
	AssetStatus *string `json:"AssetStatus,omitempty" name:"AssetStatus"`

	// 创建资产的时间。
	// ASSET_NORMAL: 正常运行，
	// ASSET_PAUSED: 暂停运行，
	// ASSET_STOPPED: 停止运行，
	// ASSET_ABNORMAL: 异常
	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`
}

type ComplianceAssetInfo struct {
	// 客户资产的ID。
	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`

	// 资产类别。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 资产的名称。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 当资产为镜像时，这个字段为镜像Tag。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// 资产所在的主机IP。
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 资产所属的节点的名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 检测状态
	// 
	// CHECK_INIT, 待检测
	// 
	// CHECK_RUNNING, 检测中
	// 
	// CHECK_FINISHED, 检测完成
	// 
	// CHECK_FAILED, 检测失败
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 此类资产通过的检测项的数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassedPolicyItemCount *uint64 `json:"PassedPolicyItemCount,omitempty" name:"PassedPolicyItemCount"`

	// 此类资产未通过的检测的数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedPolicyItemCount *uint64 `json:"FailedPolicyItemCount,omitempty" name:"FailedPolicyItemCount"`

	// 上次检测的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`

	// 检测结果：
	// RESULT_FAILED: 未通过。
	// RESULT_PASSED: 通过。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`

	// 主机节点的实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ComplianceAssetPolicyItem struct {
	// 为客户分配的唯一的检测项的ID。
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 检测项的原始ID
	BasePolicyItemId *uint64 `json:"BasePolicyItemId,omitempty" name:"BasePolicyItemId"`

	// 检测项的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 检测项所属的类型的名称
	Category *string `json:"Category,omitempty" name:"Category"`

	// 所属的合规标准的ID
	BenchmarkStandardId *uint64 `json:"BenchmarkStandardId,omitempty" name:"BenchmarkStandardId"`

	// 所属的合规标准的名称
	BenchmarkStandardName *string `json:"BenchmarkStandardName,omitempty" name:"BenchmarkStandardName"`

	// 威胁等级
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 检测状态
	// CHECK_INIT, 待检测
	// CHECK_RUNNING, 检测中
	// CHECK_FINISHED, 检测完成
	// CHECK_FAILED, 检测失败
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 检测结果
	// RESULT_PASSED: 通过
	// RESULT_FAILED: 未通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`

	// 检测项对应的白名单项的ID。如果存在且非0，表示检测项被用户忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhitelistId *uint64 `json:"WhitelistId,omitempty" name:"WhitelistId"`

	// 处理建议。
	FixSuggestion *string `json:"FixSuggestion,omitempty" name:"FixSuggestion"`

	// 最近检测的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`

	// 验证信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyInfo *string `json:"VerifyInfo,omitempty" name:"VerifyInfo"`
}

type ComplianceAssetPolicySetItem struct {
	// 资产ID
	CustomerAssetItemId *uint64 `json:"CustomerAssetItemId,omitempty" name:"CustomerAssetItemId"`

	// 需要忽略指定资产内的检查项ID列表，为空表示所有
	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

type ComplianceAssetSummary struct {
	// 资产类别。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 是否为客户的首次检测。与CheckStatus配合使用。
	IsCustomerFirstCheck *bool `json:"IsCustomerFirstCheck,omitempty" name:"IsCustomerFirstCheck"`

	// 检测状态
	// 
	// CHECK_UNINIT, 用户未启用此功能
	// 
	// CHECK_INIT, 待检测
	// 
	// CHECK_RUNNING, 检测中
	// 
	// CHECK_FINISHED, 检测完成
	// 
	// CHECK_FAILED, 检测失败
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 此类别的检测进度，为 0~100 的数。若未在检测中，无此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckProgress *float64 `json:"CheckProgress,omitempty" name:"CheckProgress"`

	// 此类资产通过的检测项的数目。
	PassedPolicyItemCount *uint64 `json:"PassedPolicyItemCount,omitempty" name:"PassedPolicyItemCount"`

	// 此类资产未通过的检测的数目。
	FailedPolicyItemCount *uint64 `json:"FailedPolicyItemCount,omitempty" name:"FailedPolicyItemCount"`

	// 此类资产下未通过的严重级别的检测项的数目。
	FailedCriticalPolicyItemCount *uint64 `json:"FailedCriticalPolicyItemCount,omitempty" name:"FailedCriticalPolicyItemCount"`

	// 此类资产下未通过的高危检测项的数目。
	FailedHighRiskPolicyItemCount *uint64 `json:"FailedHighRiskPolicyItemCount,omitempty" name:"FailedHighRiskPolicyItemCount"`

	// 此类资产下未通过的中危检测项的数目。
	FailedMediumRiskPolicyItemCount *uint64 `json:"FailedMediumRiskPolicyItemCount,omitempty" name:"FailedMediumRiskPolicyItemCount"`

	// 此类资产下未通过的低危检测项的数目。
	FailedLowRiskPolicyItemCount *uint64 `json:"FailedLowRiskPolicyItemCount,omitempty" name:"FailedLowRiskPolicyItemCount"`

	// 此类资产下提示级别的检测项的数目。
	NoticePolicyItemCount *uint64 `json:"NoticePolicyItemCount,omitempty" name:"NoticePolicyItemCount"`

	// 通过检测的资产的数目。
	PassedAssetCount *uint64 `json:"PassedAssetCount,omitempty" name:"PassedAssetCount"`

	// 未通过检测的资产的数目。
	FailedAssetCount *uint64 `json:"FailedAssetCount,omitempty" name:"FailedAssetCount"`

	// 此类资产的合规率，0~100的数。
	AssetPassedRate *float64 `json:"AssetPassedRate,omitempty" name:"AssetPassedRate"`

	// 检测失败的资产的数目。
	ScanFailedAssetCount *uint64 `json:"ScanFailedAssetCount,omitempty" name:"ScanFailedAssetCount"`

	// 上次检测的耗时，单位为秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckCostTime *float64 `json:"CheckCostTime,omitempty" name:"CheckCostTime"`

	// 上次检测的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`

	// 定时检测规则。
	PeriodRule *CompliancePeriodTaskRule `json:"PeriodRule,omitempty" name:"PeriodRule"`

	// 已开启的检查项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenPolicyItemCount *uint64 `json:"OpenPolicyItemCount,omitempty" name:"OpenPolicyItemCount"`

	// 已忽略的检查项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoredPolicyItemCount *uint64 `json:"IgnoredPolicyItemCount,omitempty" name:"IgnoredPolicyItemCount"`
}

type ComplianceBenchmarkStandard struct {
	// 合规标准的ID
	StandardId *uint64 `json:"StandardId,omitempty" name:"StandardId"`

	// 合规标准的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合规标准包含的数目
	PolicyItemCount *uint64 `json:"PolicyItemCount,omitempty" name:"PolicyItemCount"`

	// 是否启用此标准
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 标准的描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ComplianceBenchmarkStandardEnable struct {
	// 合规标准的ID。
	StandardId *uint64 `json:"StandardId,omitempty" name:"StandardId"`

	// 是否启用合规标准
	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type ComplianceContainerDetailInfo struct {
	// 容器在主机上的ID。
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 容器所属的Pod的名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitempty" name:"PodName"`
}

type ComplianceFilters struct {
	// 过滤键的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否模糊查询。默认为是。
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type ComplianceHostDetailInfo struct {
	// 主机上的Docker版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

	// 主机上的K8S的版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`
}

type ComplianceImageDetailInfo struct {
	// 镜像在主机上的ID。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像的名称。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像的Tag。
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// 镜像所在远程仓库的路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Repository *string `json:"Repository,omitempty" name:"Repository"`
}

type ComplianceK8SDetailInfo struct {
	// K8S集群的名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// K8S集群的版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
}

type CompliancePeriodTask struct {
	// 周期任务的ID
	PeriodTaskId *uint64 `json:"PeriodTaskId,omitempty" name:"PeriodTaskId"`

	// 资产类型。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 最近一次触发的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTriggerTime *string `json:"LastTriggerTime,omitempty" name:"LastTriggerTime"`

	// 总的检查项数目
	TotalPolicyItemCount *uint64 `json:"TotalPolicyItemCount,omitempty" name:"TotalPolicyItemCount"`

	// 周期设置
	PeriodRule *CompliancePeriodTaskRule `json:"PeriodRule,omitempty" name:"PeriodRule"`

	// 合规标准列表
	BenchmarkStandardSet []*ComplianceBenchmarkStandard `json:"BenchmarkStandardSet,omitempty" name:"BenchmarkStandardSet"`
}

type CompliancePeriodTaskRule struct {
	// 执行的频率（几天一次），取值为：1,3,7。
	Frequency *uint64 `json:"Frequency,omitempty" name:"Frequency"`

	// 在这天的什么时间执行，格式为：HH:mm:SS。
	ExecutionTime *string `json:"ExecutionTime,omitempty" name:"ExecutionTime"`

	// 是否开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type CompliancePolicyAssetSetItem struct {
	// 检查项ID
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 需要忽略指定检查项内的资产ID列表，为空表示所有
	CustomerAssetItemIdSet []*uint64 `json:"CustomerAssetItemIdSet,omitempty" name:"CustomerAssetItemIdSet"`
}

type CompliancePolicyItemSummary struct {
	// 为客户分配的唯一的检测项的ID。
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 检测项的原始ID。
	BasePolicyItemId *uint64 `json:"BasePolicyItemId,omitempty" name:"BasePolicyItemId"`

	// 检测项的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 检测项所属的类型，枚举字符串。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 所属的合规标准
	BenchmarkStandardName *string `json:"BenchmarkStandardName,omitempty" name:"BenchmarkStandardName"`

	// 威胁等级。RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 检测项所属的资产类型
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 最近检测的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastCheckTime *string `json:"LastCheckTime,omitempty" name:"LastCheckTime"`

	// 检测状态
	// 
	// CHECK_INIT, 待检测
	// 
	// CHECK_RUNNING, 检测中
	// 
	// CHECK_FINISHED, 检测完成
	// 
	// CHECK_FAILED, 检测失败
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 检测结果。RESULT_PASSED: 通过
	// 
	// RESULT_FAILED: 未通过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`

	// 通过检测的资产的数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassedAssetCount *uint64 `json:"PassedAssetCount,omitempty" name:"PassedAssetCount"`

	// 未通过检测的资产的数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedAssetCount *uint64 `json:"FailedAssetCount,omitempty" name:"FailedAssetCount"`

	// 检测项对应的白名单项的ID。如果存在且非0，表示检测项被用户忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhitelistId *uint64 `json:"WhitelistId,omitempty" name:"WhitelistId"`

	// 处理建议。
	FixSuggestion *string `json:"FixSuggestion,omitempty" name:"FixSuggestion"`

	// 所属的合规标准的ID
	BenchmarkStandardId *uint64 `json:"BenchmarkStandardId,omitempty" name:"BenchmarkStandardId"`

	// 检测项适用的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicableVersion *string `json:"ApplicableVersion,omitempty" name:"ApplicableVersion"`
}

type ComplianceScanFailedAsset struct {
	// 客户资产的ID。
	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`

	// 资产类别。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 检测状态
	// CHECK_INIT, 待检测
	// CHECK_RUNNING, 检测中
	// CHECK_FINISHED, 检测完成
	// CHECK_FAILED, 检测失败
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 资产的名称。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 资产检测失败的原因。
	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`

	// 检测失败的处理建议。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 检测的时间。
	CheckTime *string `json:"CheckTime,omitempty" name:"CheckTime"`
}

type ComplianceWhitelistItem struct {
	// 白名单项的ID。
	WhitelistItemId *uint64 `json:"WhitelistItemId,omitempty" name:"WhitelistItemId"`

	// 客户检测项的ID。
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 检测项的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 合规标准的名称。
	StandardName *string `json:"StandardName,omitempty" name:"StandardName"`

	// 合规标准的ID。
	StandardId *uint64 `json:"StandardId,omitempty" name:"StandardId"`

	// 检测项影响的资产的数目。
	AffectedAssetCount *uint64 `json:"AffectedAssetCount,omitempty" name:"AffectedAssetCount"`

	// 最后更新的时间
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 加入到白名单的时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
}

type ComponentInfo struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 版本
	Version *string `json:"Version,omitempty" name:"Version"`
}

type ComponentsInfo struct {
	// 组件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Component *string `json:"Component,omitempty" name:"Component"`

	// 组件版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
}

// Predefined struct for user
type ConfirmNetworkFirewallPolicyRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id数组
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

type ConfirmNetworkFirewallPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id数组
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ConfirmNetworkFirewallPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmNetworkFirewallPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfirmNetworkFirewallPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfirmNetworkFirewallPolicyResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建确认任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ConfirmNetworkFirewallPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ConfirmNetworkFirewallPolicyResponseParams `json:"Response"`
}

func (r *ConfirmNetworkFirewallPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmNetworkFirewallPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerInfo struct {
	// 容器id
	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 容器运行状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 运行用户
	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`

	// 命令行
	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`

	// CPU使用率 *1000
	CPUUsage *uint64 `json:"CPUUsage,omitempty" name:"CPUUsage"`

	// 内存使用 kb
	RamUsage *uint64 `json:"RamUsage,omitempty" name:"RamUsage"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像id
	POD *string `json:"POD,omitempty" name:"POD"`

	// 主机id
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 主机ip
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 外网ip
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`

	// 网络子状态
	NetSubStatus *string `json:"NetSubStatus,omitempty" name:"NetSubStatus"`

	// 隔离来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateSource *string `json:"IsolateSource,omitempty" name:"IsolateSource"`

	// 隔离时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
}

type ContainerMount struct {
	// 挂载类型 bind
	Type *string `json:"Type,omitempty" name:"Type"`

	// 宿主机路径
	Source *string `json:"Source,omitempty" name:"Source"`

	// 容器内路径
	Destination *string `json:"Destination,omitempty" name:"Destination"`

	// 模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 读写权限
	RW *bool `json:"RW,omitempty" name:"RW"`

	// 传播类型
	Propagation *string `json:"Propagation,omitempty" name:"Propagation"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 驱动
	Driver *string `json:"Driver,omitempty" name:"Driver"`
}

type ContainerNetwork struct {
	// endpoint id
	EndpointID *string `json:"EndpointID,omitempty" name:"EndpointID"`

	// 模式:bridge
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 网络名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网络ID
	NetworkID *string `json:"NetworkID,omitempty" name:"NetworkID"`

	// 网关
	Gateway *string `json:"Gateway,omitempty" name:"Gateway"`

	// IPV4地址
	Ipv4 *string `json:"Ipv4,omitempty" name:"Ipv4"`

	// IPV6地址
	Ipv6 *string `json:"Ipv6,omitempty" name:"Ipv6"`

	// MAC 地址
	MAC *string `json:"MAC,omitempty" name:"MAC"`
}

// Predefined struct for user
type CreateAbnormalProcessRulesExportJobRequestParams struct {
	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type CreateAbnormalProcessRulesExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateAbnormalProcessRulesExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAbnormalProcessRulesExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAbnormalProcessRulesExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAbnormalProcessRulesExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAbnormalProcessRulesExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateAbnormalProcessRulesExportJobResponseParams `json:"Response"`
}

func (r *CreateAbnormalProcessRulesExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAbnormalProcessRulesExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessControlsRuleExportJobRequestParams struct {
	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By []*string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type CreateAccessControlsRuleExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By []*string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateAccessControlsRuleExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessControlsRuleExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessControlsRuleExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessControlsRuleExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccessControlsRuleExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessControlsRuleExportJobResponseParams `json:"Response"`
}

func (r *CreateAccessControlsRuleExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessControlsRuleExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageRegistryScanTaskOneKeyRequestParams struct {
	// 是否扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 扫描类型数组
	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`

	// 扫描的镜像列表Id
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

type CreateAssetImageRegistryScanTaskOneKeyRequest struct {
	*tchttp.BaseRequest
	
	// 是否扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 扫描类型数组
	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`

	// 扫描的镜像列表Id
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CreateAssetImageRegistryScanTaskOneKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageRegistryScanTaskOneKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "All")
	delete(f, "Images")
	delete(f, "ScanType")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetImageRegistryScanTaskOneKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageRegistryScanTaskOneKeyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAssetImageRegistryScanTaskOneKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetImageRegistryScanTaskOneKeyResponseParams `json:"Response"`
}

func (r *CreateAssetImageRegistryScanTaskOneKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageRegistryScanTaskOneKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageRegistryScanTaskRequestParams struct {
	// 是否扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 扫描类型数组
	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`

	// 扫描的镜像列表
	Id []*uint64 `json:"Id,omitempty" name:"Id"`

	// 过滤条件
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 不需要扫描的镜像列表, 与Filters配合使用
	ExcludeImageList []*uint64 `json:"ExcludeImageList,omitempty" name:"ExcludeImageList"`

	// 是否仅扫描各repository最新版的镜像, 与Filters配合使用
	OnlyScanLatest *bool `json:"OnlyScanLatest,omitempty" name:"OnlyScanLatest"`
}

type CreateAssetImageRegistryScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 是否扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 扫描类型数组
	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`

	// 扫描的镜像列表
	Id []*uint64 `json:"Id,omitempty" name:"Id"`

	// 过滤条件
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 不需要扫描的镜像列表, 与Filters配合使用
	ExcludeImageList []*uint64 `json:"ExcludeImageList,omitempty" name:"ExcludeImageList"`

	// 是否仅扫描各repository最新版的镜像, 与Filters配合使用
	OnlyScanLatest *bool `json:"OnlyScanLatest,omitempty" name:"OnlyScanLatest"`
}

func (r *CreateAssetImageRegistryScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageRegistryScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "All")
	delete(f, "Images")
	delete(f, "ScanType")
	delete(f, "Id")
	delete(f, "Filters")
	delete(f, "ExcludeImageList")
	delete(f, "OnlyScanLatest")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetImageRegistryScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageRegistryScanTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAssetImageRegistryScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetImageRegistryScanTaskResponseParams `json:"Response"`
}

func (r *CreateAssetImageRegistryScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageRegistryScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageScanSettingRequestParams struct {
	// 开关
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 扫描时间
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 扫描周期
	ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`

	// 扫描木马
	ScanVirus *bool `json:"ScanVirus,omitempty" name:"ScanVirus"`

	// 扫描敏感信息
	ScanRisk *bool `json:"ScanRisk,omitempty" name:"ScanRisk"`

	// 扫描漏洞
	ScanVul *bool `json:"ScanVul,omitempty" name:"ScanVul"`

	// 全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 自定义镜像
	Images []*string `json:"Images,omitempty" name:"Images"`
}

type CreateAssetImageScanSettingRequest struct {
	*tchttp.BaseRequest
	
	// 开关
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 扫描时间
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 扫描周期
	ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`

	// 扫描木马
	ScanVirus *bool `json:"ScanVirus,omitempty" name:"ScanVirus"`

	// 扫描敏感信息
	ScanRisk *bool `json:"ScanRisk,omitempty" name:"ScanRisk"`

	// 扫描漏洞
	ScanVul *bool `json:"ScanVul,omitempty" name:"ScanVul"`

	// 全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 自定义镜像
	Images []*string `json:"Images,omitempty" name:"Images"`
}

func (r *CreateAssetImageScanSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageScanSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Enable")
	delete(f, "ScanTime")
	delete(f, "ScanPeriod")
	delete(f, "ScanVirus")
	delete(f, "ScanRisk")
	delete(f, "ScanVul")
	delete(f, "All")
	delete(f, "Images")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetImageScanSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageScanSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAssetImageScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetImageScanSettingResponseParams `json:"Response"`
}

func (r *CreateAssetImageScanSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageScanTaskRequestParams struct {
	// 是否扫描全部镜像；全部镜像，镜像列表和根据过滤条件筛选三选一。
	All *bool `json:"All,omitempty" name:"All"`

	// 需要扫描的镜像列表；全部镜像，镜像列表和根据过滤条件筛选三选一。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 扫描漏洞；漏洞，木马和风险需选其一
	ScanVul *bool `json:"ScanVul,omitempty" name:"ScanVul"`

	// 扫描木马；漏洞，木马和风险需选其一
	ScanVirus *bool `json:"ScanVirus,omitempty" name:"ScanVirus"`

	// 扫描风险；漏洞，木马和风险需选其一
	ScanRisk *bool `json:"ScanRisk,omitempty" name:"ScanRisk"`

	// 根据过滤条件筛选出镜像；全部镜像，镜像列表和根据过滤条件筛选三选一。
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 根据过滤条件筛选出镜像，再排除个别镜像
	ExcludeImageIds []*string `json:"ExcludeImageIds,omitempty" name:"ExcludeImageIds"`
}

type CreateAssetImageScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 是否扫描全部镜像；全部镜像，镜像列表和根据过滤条件筛选三选一。
	All *bool `json:"All,omitempty" name:"All"`

	// 需要扫描的镜像列表；全部镜像，镜像列表和根据过滤条件筛选三选一。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 扫描漏洞；漏洞，木马和风险需选其一
	ScanVul *bool `json:"ScanVul,omitempty" name:"ScanVul"`

	// 扫描木马；漏洞，木马和风险需选其一
	ScanVirus *bool `json:"ScanVirus,omitempty" name:"ScanVirus"`

	// 扫描风险；漏洞，木马和风险需选其一
	ScanRisk *bool `json:"ScanRisk,omitempty" name:"ScanRisk"`

	// 根据过滤条件筛选出镜像；全部镜像，镜像列表和根据过滤条件筛选三选一。
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 根据过滤条件筛选出镜像，再排除个别镜像
	ExcludeImageIds []*string `json:"ExcludeImageIds,omitempty" name:"ExcludeImageIds"`
}

func (r *CreateAssetImageScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "All")
	delete(f, "Images")
	delete(f, "ScanVul")
	delete(f, "ScanVirus")
	delete(f, "ScanRisk")
	delete(f, "Filters")
	delete(f, "ExcludeImageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetImageScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageScanTaskResponseParams struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAssetImageScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetImageScanTaskResponseParams `json:"Response"`
}

func (r *CreateAssetImageScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageVirusExportJobRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为10000
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type CreateAssetImageVirusExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为10000
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *CreateAssetImageVirusExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageVirusExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "ImageID")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetImageVirusExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetImageVirusExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAssetImageVirusExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetImageVirusExportJobResponseParams `json:"Response"`
}

func (r *CreateAssetImageVirusExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetImageVirusExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCheckComponentRequestParams struct {
	// 要安装的集群列表信息
	ClusterInfoList []*ClusterCreateComponentItem `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`
}

type CreateCheckComponentRequest struct {
	*tchttp.BaseRequest
	
	// 要安装的集群列表信息
	ClusterInfoList []*ClusterCreateComponentItem `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`
}

func (r *CreateCheckComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCheckComponentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterInfoList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCheckComponentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCheckComponentResponseParams struct {
	// "InstallSucc"表示安装成功，"InstallFailed"表示安装失败
	InstallResult *string `json:"InstallResult,omitempty" name:"InstallResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCheckComponentResponse struct {
	*tchttp.BaseResponse
	Response *CreateCheckComponentResponseParams `json:"Response"`
}

func (r *CreateCheckComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCheckComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterCheckTaskRequestParams struct {
	// 指定要扫描的集群信息
	ClusterCheckTaskList []*ClusterCheckTaskItem `json:"ClusterCheckTaskList,omitempty" name:"ClusterCheckTaskList"`
}

type CreateClusterCheckTaskRequest struct {
	*tchttp.BaseRequest
	
	// 指定要扫描的集群信息
	ClusterCheckTaskList []*ClusterCheckTaskItem `json:"ClusterCheckTaskList,omitempty" name:"ClusterCheckTaskList"`
}

func (r *CreateClusterCheckTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterCheckTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterCheckTaskList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterCheckTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterCheckTaskResponseParams struct {
	// 返回创建的集群检查任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建检查任务的结果，"Succ"为成功，其他的为失败原因
	CreateResult *string `json:"CreateResult,omitempty" name:"CreateResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterCheckTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterCheckTaskResponseParams `json:"Response"`
}

func (r *CreateClusterCheckTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterCheckTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateComplianceTaskRequestParams struct {
	// 指定要扫描的资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	// AssetTypeSet, PolicySetId, PeriodTaskId三个参数，必须要给其中一个参数填写有效的值。
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`

	// 按照策略集ID指定的策略执行合规检查。
	PolicySetId *uint64 `json:"PolicySetId,omitempty" name:"PolicySetId"`

	// 按照定时任务ID指定的策略执行合规检查。
	PeriodTaskId *uint64 `json:"PeriodTaskId,omitempty" name:"PeriodTaskId"`
}

type CreateComplianceTaskRequest struct {
	*tchttp.BaseRequest
	
	// 指定要扫描的资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	// AssetTypeSet, PolicySetId, PeriodTaskId三个参数，必须要给其中一个参数填写有效的值。
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`

	// 按照策略集ID指定的策略执行合规检查。
	PolicySetId *uint64 `json:"PolicySetId,omitempty" name:"PolicySetId"`

	// 按照定时任务ID指定的策略执行合规检查。
	PeriodTaskId *uint64 `json:"PeriodTaskId,omitempty" name:"PeriodTaskId"`
}

func (r *CreateComplianceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateComplianceTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetTypeSet")
	delete(f, "PolicySetId")
	delete(f, "PeriodTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateComplianceTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateComplianceTaskResponseParams struct {
	// 返回创建的合规检查任务的ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateComplianceTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateComplianceTaskResponseParams `json:"Response"`
}

func (r *CreateComplianceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateComplianceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateComponentExportJobRequestParams struct {
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10000，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式desc ，asc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type CreateComponentExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10000，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式desc ，asc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *CreateComponentExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateComponentExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageID")
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateComponentExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateComponentExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateComponentExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateComponentExportJobResponseParams `json:"Response"`
}

func (r *CreateComponentExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateComponentExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefenceVulExportJobRequestParams struct {
	// 需要返回的数量，默认为10000，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type CreateDefenceVulExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10000，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateDefenceVulExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefenceVulExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDefenceVulExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefenceVulExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDefenceVulExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateDefenceVulExportJobResponseParams `json:"Response"`
}

func (r *CreateDefenceVulExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefenceVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmergencyVulExportJobRequestParams struct {
	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type CreateEmergencyVulExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateEmergencyVulExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmergencyVulExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmergencyVulExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmergencyVulExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEmergencyVulExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmergencyVulExportJobResponseParams `json:"Response"`
}

func (r *CreateEmergencyVulExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmergencyVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEscapeEventsExportJobRequestParams struct {
	// 需要返回的数量，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：latest_found_time
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type CreateEscapeEventsExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：latest_found_time
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateEscapeEventsExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEscapeEventsExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEscapeEventsExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEscapeEventsExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEscapeEventsExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateEscapeEventsExportJobResponseParams `json:"Response"`
}

func (r *CreateEscapeEventsExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEscapeEventsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEscapeWhiteListExportJobRequestParams struct {
	// 过滤条件。
	// <li>EventType- String - 是否必填：否 - 加白事件类型，ESCAPE_CGROUPS：利用cgroup机制逃逸，ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸， ESCAPE_DOCKER_API：访问Docker API接口逃逸，ESCAPE_VUL_OCCURRED：逃逸漏洞利用，MOUNT_SENSITIVE_PTAH：敏感路径挂载，PRIVILEGE_CONTAINER_START：特权容器， PRIVILEGE：程序提权逃逸</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10000，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：主机数量：HostCount，容器数量：ContainerCount，更新时间：UpdateTime
	By *string `json:"By,omitempty" name:"By"`
}

type CreateEscapeWhiteListExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>EventType- String - 是否必填：否 - 加白事件类型，ESCAPE_CGROUPS：利用cgroup机制逃逸，ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸， ESCAPE_DOCKER_API：访问Docker API接口逃逸，ESCAPE_VUL_OCCURRED：逃逸漏洞利用，MOUNT_SENSITIVE_PTAH：敏感路径挂载，PRIVILEGE_CONTAINER_START：特权容器， PRIVILEGE：程序提权逃逸</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10000，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：主机数量：HostCount，容器数量：ContainerCount，更新时间：UpdateTime
	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateEscapeWhiteListExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEscapeWhiteListExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEscapeWhiteListExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEscapeWhiteListExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEscapeWhiteListExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateEscapeWhiteListExportJobResponseParams `json:"Response"`
}

func (r *CreateEscapeWhiteListExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEscapeWhiteListExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportComplianceStatusListJobRequestParams struct {
	// 要导出信息的资产类型
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 按照检测项导出，还是按照资产导出。true: 按照资产导出；false: 按照检测项导出。
	ExportByAsset *bool `json:"ExportByAsset,omitempty" name:"ExportByAsset"`

	// true, 全部导出；false, 根据IdList来导出数据。
	ExportAll *bool `json:"ExportAll,omitempty" name:"ExportAll"`

	// 要导出的资产ID列表或检测项ID列表，由ExportByAsset的取值决定。
	IdList []*uint64 `json:"IdList,omitempty" name:"IdList"`
}

type CreateExportComplianceStatusListJobRequest struct {
	*tchttp.BaseRequest
	
	// 要导出信息的资产类型
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 按照检测项导出，还是按照资产导出。true: 按照资产导出；false: 按照检测项导出。
	ExportByAsset *bool `json:"ExportByAsset,omitempty" name:"ExportByAsset"`

	// true, 全部导出；false, 根据IdList来导出数据。
	ExportAll *bool `json:"ExportAll,omitempty" name:"ExportAll"`

	// 要导出的资产ID列表或检测项ID列表，由ExportByAsset的取值决定。
	IdList []*uint64 `json:"IdList,omitempty" name:"IdList"`
}

func (r *CreateExportComplianceStatusListJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportComplianceStatusListJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetType")
	delete(f, "ExportByAsset")
	delete(f, "ExportAll")
	delete(f, "IdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportComplianceStatusListJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportComplianceStatusListJobResponseParams struct {
	// 返回创建的导出任务的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateExportComplianceStatusListJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateExportComplianceStatusListJobResponseParams `json:"Response"`
}

func (r *CreateExportComplianceStatusListJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportComplianceStatusListJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostExportJobRequestParams struct {
	// 过滤条件。
	// <li>Status - String - 是否必填：否 - agent状态筛选，"ALL":"全部"(或不传该字段),"UNINSTALL"："未安装","OFFLINE"："离线", "ONLINE"："防护中"</li>
	// <li>HostName - String - 是否必填：否 - 主机名筛选</li>
	// <li>Group- String - 是否必填：否 - 主机群组搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>HostID- string - 是否必填：否 - 主机id搜索</li>
	// <li>DockerVersion- string - 是否必填：否 - docker版本搜索</li>
	// <li>MachineType- string - 是否必填：否 - 主机来源MachineType搜索，"ALL":"全部"(或不传该字段),主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；</li>
	// <li>DockerStatus- string - 是否必填：否 - docker安装状态，"ALL":"全部"(或不传该字段),"INSTALL":"已安装","UNINSTALL":"未安装"</li>
	// <li>ProjectID- string - 是否必填：否 - 所属项目id搜索</li>
	// <li>Tag:xxx(tag:key)- string- 是否必填：否 - 标签键值搜索 示例Filters":[{"Name":"tag:tke-kind","Values":["service"]}]</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 需要返回的数量，默认为10，最大值为10000
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type CreateHostExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Status - String - 是否必填：否 - agent状态筛选，"ALL":"全部"(或不传该字段),"UNINSTALL"："未安装","OFFLINE"："离线", "ONLINE"："防护中"</li>
	// <li>HostName - String - 是否必填：否 - 主机名筛选</li>
	// <li>Group- String - 是否必填：否 - 主机群组搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>HostID- string - 是否必填：否 - 主机id搜索</li>
	// <li>DockerVersion- string - 是否必填：否 - docker版本搜索</li>
	// <li>MachineType- string - 是否必填：否 - 主机来源MachineType搜索，"ALL":"全部"(或不传该字段),主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；</li>
	// <li>DockerStatus- string - 是否必填：否 - docker安装状态，"ALL":"全部"(或不传该字段),"INSTALL":"已安装","UNINSTALL":"未安装"</li>
	// <li>ProjectID- string - 是否必填：否 - 所属项目id搜索</li>
	// <li>Tag:xxx(tag:key)- string- 是否必填：否 - 标签键值搜索 示例Filters":[{"Name":"tag:tke-kind","Values":["service"]}]</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 需要返回的数量，默认为10，最大值为10000
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateHostExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHostExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHostExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateHostExportJobResponseParams `json:"Response"`
}

func (r *CreateHostExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageExportJobRequestParams struct {
	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type CreateImageExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateImageExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// excel文件下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateImageExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageExportJobResponseParams `json:"Response"`
}

func (r *CreateImageExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateK8sApiAbnormalEventExportJobRequestParams struct {
	// 过滤条件。
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>MatchRules - string  - 是否必填: 否 -命中规则筛选</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	// <li>Status - string  - 是否必填: 否 -事件状态筛选</li>
	// <li>MatchRuleType - string  - 是否必填: 否 -命中规则类型筛选</li>
	// <li>ClusterRunningStatus - string  - 是否必填: 否 -集群运行状态</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type CreateK8sApiAbnormalEventExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>MatchRules - string  - 是否必填: 否 -命中规则筛选</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	// <li>Status - string  - 是否必填: 否 -事件状态筛选</li>
	// <li>MatchRuleType - string  - 是否必填: 否 -命中规则类型筛选</li>
	// <li>ClusterRunningStatus - string  - 是否必填: 否 -集群运行状态</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateK8sApiAbnormalEventExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateK8sApiAbnormalEventExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateK8sApiAbnormalEventExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateK8sApiAbnormalEventExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateK8sApiAbnormalEventExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateK8sApiAbnormalEventExportJobResponseParams `json:"Response"`
}

func (r *CreateK8sApiAbnormalEventExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateK8sApiAbnormalEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateK8sApiAbnormalRuleExportJobRequestParams struct {
	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By []*string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type CreateK8sApiAbnormalRuleExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By []*string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateK8sApiAbnormalRuleExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateK8sApiAbnormalRuleExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateK8sApiAbnormalRuleExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateK8sApiAbnormalRuleExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateK8sApiAbnormalRuleExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateK8sApiAbnormalRuleExportJobResponseParams `json:"Response"`
}

func (r *CreateK8sApiAbnormalRuleExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateK8sApiAbnormalRuleExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateK8sApiAbnormalRuleInfoRequestParams struct {
	// 规则详情
	RuleInfo *K8sApiAbnormalRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

	// 拷贝规则ID(适用于复制规则场景)
	CopySrcRuleID *string `json:"CopySrcRuleID,omitempty" name:"CopySrcRuleID"`

	// 事件ID(适用于事件加白场景)
	EventID *uint64 `json:"EventID,omitempty" name:"EventID"`
}

type CreateK8sApiAbnormalRuleInfoRequest struct {
	*tchttp.BaseRequest
	
	// 规则详情
	RuleInfo *K8sApiAbnormalRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`

	// 拷贝规则ID(适用于复制规则场景)
	CopySrcRuleID *string `json:"CopySrcRuleID,omitempty" name:"CopySrcRuleID"`

	// 事件ID(适用于事件加白场景)
	EventID *uint64 `json:"EventID,omitempty" name:"EventID"`
}

func (r *CreateK8sApiAbnormalRuleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateK8sApiAbnormalRuleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleInfo")
	delete(f, "CopySrcRuleID")
	delete(f, "EventID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateK8sApiAbnormalRuleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateK8sApiAbnormalRuleInfoResponseParams struct {
	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateK8sApiAbnormalRuleInfoResponse struct {
	*tchttp.BaseResponse
	Response *CreateK8sApiAbnormalRuleInfoResponseParams `json:"Response"`
}

func (r *CreateK8sApiAbnormalRuleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateK8sApiAbnormalRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkFirewallClusterRefreshRequestParams struct {

}

type CreateNetworkFirewallClusterRefreshRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateNetworkFirewallClusterRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkFirewallClusterRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkFirewallClusterRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkFirewallClusterRefreshResponseParams struct {
	// 返回创建的集群检查任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建检查任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNetworkFirewallClusterRefreshResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkFirewallClusterRefreshResponseParams `json:"Response"`
}

func (r *CreateNetworkFirewallClusterRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkFirewallClusterRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkFirewallPolicyDiscoverRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type CreateNetworkFirewallPolicyDiscoverRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateNetworkFirewallPolicyDiscoverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkFirewallPolicyDiscoverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkFirewallPolicyDiscoverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkFirewallPolicyDiscoverResponseParams struct {
	// 返回创建的集群检查任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建检查任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNetworkFirewallPolicyDiscoverResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkFirewallPolicyDiscoverResponseParams `json:"Response"`
}

func (r *CreateNetworkFirewallPolicyDiscoverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkFirewallPolicyDiscoverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkFirewallPublishRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id数组
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

type CreateNetworkFirewallPublishRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id数组
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CreateNetworkFirewallPublishRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkFirewallPublishRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkFirewallPublishRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkFirewallPublishResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNetworkFirewallPublishResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkFirewallPublishResponseParams `json:"Response"`
}

func (r *CreateNetworkFirewallPublishResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkFirewallPublishResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkFirewallUndoPublishRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id数组
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

type CreateNetworkFirewallUndoPublishRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id数组
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *CreateNetworkFirewallUndoPublishRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkFirewallUndoPublishRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkFirewallUndoPublishRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkFirewallUndoPublishResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNetworkFirewallUndoPublishResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkFirewallUndoPublishResponseParams `json:"Response"`
}

func (r *CreateNetworkFirewallUndoPublishResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkFirewallUndoPublishResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrModifyPostPayCoresRequestParams struct {
	// 弹性计费上限，最小值500
	CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`
}

type CreateOrModifyPostPayCoresRequest struct {
	*tchttp.BaseRequest
	
	// 弹性计费上限，最小值500
	CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`
}

func (r *CreateOrModifyPostPayCoresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrModifyPostPayCoresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CoresCnt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrModifyPostPayCoresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrModifyPostPayCoresResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrModifyPostPayCoresResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrModifyPostPayCoresResponseParams `json:"Response"`
}

func (r *CreateOrModifyPostPayCoresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrModifyPostPayCoresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProcessEventsExportJobRequestParams struct {
	// 需要返回的数量，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：latest_found_time
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type CreateProcessEventsExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：latest_found_time
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *CreateProcessEventsExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcessEventsExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProcessEventsExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProcessEventsExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProcessEventsExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateProcessEventsExportJobResponseParams `json:"Response"`
}

func (r *CreateProcessEventsExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProcessEventsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRefreshTaskRequestParams struct {

}

type CreateRefreshTaskRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateRefreshTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRefreshTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRefreshTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRefreshTaskResponseParams struct {
	// 返回创建的集群检查任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建检查任务的结果，"Succ"为成功，"Failed"为失败
	CreateResult *string `json:"CreateResult,omitempty" name:"CreateResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRefreshTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRefreshTaskResponseParams `json:"Response"`
}

func (r *CreateRefreshTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRefreshTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRiskDnsEventExportJobRequestParams struct {
	// 过滤条件。
	// <li>EventStatus- String - 是否必填：否 - 事件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_ADD_WHITE：已加白</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>RiskDns- string - 是否必填：否 - 恶意域名。</li>
	// <li>RiskIP- string - 是否必填：否 - 恶意IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，最大值为100000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：事件数量：EventCount
	By *string `json:"By,omitempty" name:"By"`
}

type CreateRiskDnsEventExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>EventStatus- String - 是否必填：否 - 事件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_ADD_WHITE：已加白</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  事件类型，恶意域名请求：DOMAIN，恶意IP请求：IP</li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>RiskDns- string - 是否必填：否 - 恶意域名。</li>
	// <li>RiskIP- string - 是否必填：否 - 恶意IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	// <li>PublicIP- string - 是否必填：否 - 外网IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，最大值为100000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：事件数量：EventCount
	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateRiskDnsEventExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRiskDnsEventExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRiskDnsEventExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRiskDnsEventExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRiskDnsEventExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateRiskDnsEventExportJobResponseParams `json:"Response"`
}

func (r *CreateRiskDnsEventExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRiskDnsEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSearchTemplateRequestParams struct {
	// 搜索模板
	SearchTemplate *SearchTemplate `json:"SearchTemplate,omitempty" name:"SearchTemplate"`
}

type CreateSearchTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 搜索模板
	SearchTemplate *SearchTemplate `json:"SearchTemplate,omitempty" name:"SearchTemplate"`
}

func (r *CreateSearchTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSearchTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSearchTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSearchTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSearchTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateSearchTemplateResponseParams `json:"Response"`
}

func (r *CreateSearchTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSearchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSystemVulExportJobRequestParams struct {
	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞子类型</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type CreateSystemVulExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞子类型</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateSystemVulExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSystemVulExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSystemVulExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSystemVulExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSystemVulExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateSystemVulExportJobResponseParams `json:"Response"`
}

func (r *CreateSystemVulExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSystemVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVirusScanAgainRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 需要扫描的容器id集合
	ContainerIds []*string `json:"ContainerIds,omitempty" name:"ContainerIds"`

	// 是否是扫描全部超时的
	TimeoutAll *bool `json:"TimeoutAll,omitempty" name:"TimeoutAll"`

	// 重新设置的超时时长
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
}

type CreateVirusScanAgainRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 需要扫描的容器id集合
	ContainerIds []*string `json:"ContainerIds,omitempty" name:"ContainerIds"`

	// 是否是扫描全部超时的
	TimeoutAll *bool `json:"TimeoutAll,omitempty" name:"TimeoutAll"`

	// 重新设置的超时时长
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
}

func (r *CreateVirusScanAgainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVirusScanAgainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ContainerIds")
	delete(f, "TimeoutAll")
	delete(f, "Timeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVirusScanAgainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVirusScanAgainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVirusScanAgainResponse struct {
	*tchttp.BaseResponse
	Response *CreateVirusScanAgainResponseParams `json:"Response"`
}

func (r *CreateVirusScanAgainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVirusScanAgainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVirusScanTaskRequestParams struct {
	// 是否扫描所有路径
	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`

	// 扫描范围0容器1主机节点
	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`

	// true 全选，false 自选
	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`

	// 超时时长，单位小时
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 当ScanPathAll为false生效 0扫描以下路径 1、扫描除以下路径
	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`

	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定
	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`

	// 自选排除或扫描的地址
	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
}

type CreateVirusScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 是否扫描所有路径
	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`

	// 扫描范围0容器1主机节点
	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`

	// true 全选，false 自选
	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`

	// 超时时长，单位小时
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 当ScanPathAll为false生效 0扫描以下路径 1、扫描除以下路径
	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`

	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定
	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`

	// 自选排除或扫描的地址
	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
}

func (r *CreateVirusScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVirusScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScanPathAll")
	delete(f, "ScanRangeType")
	delete(f, "ScanRangeAll")
	delete(f, "Timeout")
	delete(f, "ScanPathType")
	delete(f, "ScanIds")
	delete(f, "ScanPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVirusScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVirusScanTaskResponseParams struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVirusScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateVirusScanTaskResponseParams `json:"Response"`
}

func (r *CreateVirusScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVirusScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulContainerExportJobRequestParams struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID</li>
	// <li>ContainerName- String -是否必填: 否 - 容器名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

type CreateVulContainerExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID</li>
	// <li>ContainerName- String -是否必填: 否 - 容器名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *CreateVulContainerExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulContainerExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PocID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVulContainerExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulContainerExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVulContainerExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateVulContainerExportJobResponseParams `json:"Response"`
}

func (r *CreateVulContainerExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulContainerExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulDefenceEventExportJobRequestParams struct {
	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_DEFENDED：已防御</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  入侵状态，防御成功：EVENT_DEFENDED，尝试攻击：EVENT_ATTACK </li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号。</li>
	// <li>SourceIP- string - 是否必填：否 - 攻击源IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，最大值为100000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：事件数量：EventCount
	By *string `json:"By,omitempty" name:"By"`
}

type CreateVulDefenceEventExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_DEFENDED：已防御</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  入侵状态，防御成功：EVENT_DEFENDED，尝试攻击：EVENT_ATTACK </li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号。</li>
	// <li>SourceIP- string - 是否必填：否 - 攻击源IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，最大值为100000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：事件数量：EventCount
	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateVulDefenceEventExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulDefenceEventExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVulDefenceEventExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulDefenceEventExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVulDefenceEventExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateVulDefenceEventExportJobResponseParams `json:"Response"`
}

func (r *CreateVulDefenceEventExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulDefenceEventExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulDefenceHostExportJobRequestParams struct {
	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，正常：SUCCESS，异常：FAIL， NO_DEFENCE:未防御</li>
	// <li>KeyWords- string - 是否必填：否 - 主机名称/IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，最大值为100000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：更新时间：ModifyTime/首次开启时间：CreateTime
	By *string `json:"By,omitempty" name:"By"`
}

type CreateVulDefenceHostExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，正常：SUCCESS，异常：FAIL， NO_DEFENCE:未防御</li>
	// <li>KeyWords- string - 是否必填：否 - 主机名称/IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，最大值为100000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：更新时间：ModifyTime/首次开启时间：CreateTime
	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateVulDefenceHostExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulDefenceHostExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVulDefenceHostExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulDefenceHostExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVulDefenceHostExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateVulDefenceHostExportJobResponseParams `json:"Response"`
}

func (r *CreateVulDefenceHostExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulDefenceHostExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulExportJobRequestParams struct {
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10000，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式desc ，asc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type CreateVulExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10000，最大值为10000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式desc ，asc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *CreateVulExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageID")
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVulExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVulExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateVulExportJobResponseParams `json:"Response"`
}

func (r *CreateVulExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulImageExportJobRequestParams struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ClientIP- string -是否必填: 否 - 内网IP</li>
	// <li>PublicIP- string -是否必填: 否 - 外网IP</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>HostName- string -是否必填: 否 - 主机名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type CreateVulImageExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ClientIP- string -是否必填: 否 - 内网IP</li>
	// <li>PublicIP- string -是否必填: 否 - 外网IP</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>HostName- string -是否必填: 否 - 主机名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateVulImageExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulImageExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PocID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVulImageExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulImageExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVulImageExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateVulImageExportJobResponseParams `json:"Response"`
}

func (r *CreateVulImageExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulImageExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulScanTaskRequestParams struct {
	// 本地镜像扫描范围类型。ALL:全部本地镜像，NOT_SCAN：全部已授权未扫描本地镜像，IMAGEIDS:自选本地镜像ID
	LocalImageScanType *string `json:"LocalImageScanType,omitempty" name:"LocalImageScanType"`

	// 根据已授权的本地镜像IDs扫描，优先权高于根据满足条件的已授权的本地镜像。
	LocalImageIDs []*string `json:"LocalImageIDs,omitempty" name:"LocalImageIDs"`

	// 仓库镜像扫描范围类型。ALL:全部仓库镜像，NOT_SCAN：全部已授权未扫描仓库镜像，IMAGEIDS:自选仓库镜像ID
	RegistryImageScanType *string `json:"RegistryImageScanType,omitempty" name:"RegistryImageScanType"`

	// 根据已授权的仓库镜像IDs扫描，优先权高于根据满足条件的已授权的仓库镜像。
	RegistryImageIDs []*uint64 `json:"RegistryImageIDs,omitempty" name:"RegistryImageIDs"`

	// 本地镜像重新漏洞扫描时的任务ID
	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`

	// 仓库镜像重新漏洞扫描时的任务ID
	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
}

type CreateVulScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 本地镜像扫描范围类型。ALL:全部本地镜像，NOT_SCAN：全部已授权未扫描本地镜像，IMAGEIDS:自选本地镜像ID
	LocalImageScanType *string `json:"LocalImageScanType,omitempty" name:"LocalImageScanType"`

	// 根据已授权的本地镜像IDs扫描，优先权高于根据满足条件的已授权的本地镜像。
	LocalImageIDs []*string `json:"LocalImageIDs,omitempty" name:"LocalImageIDs"`

	// 仓库镜像扫描范围类型。ALL:全部仓库镜像，NOT_SCAN：全部已授权未扫描仓库镜像，IMAGEIDS:自选仓库镜像ID
	RegistryImageScanType *string `json:"RegistryImageScanType,omitempty" name:"RegistryImageScanType"`

	// 根据已授权的仓库镜像IDs扫描，优先权高于根据满足条件的已授权的仓库镜像。
	RegistryImageIDs []*uint64 `json:"RegistryImageIDs,omitempty" name:"RegistryImageIDs"`

	// 本地镜像重新漏洞扫描时的任务ID
	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`

	// 仓库镜像重新漏洞扫描时的任务ID
	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
}

func (r *CreateVulScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LocalImageScanType")
	delete(f, "LocalImageIDs")
	delete(f, "RegistryImageScanType")
	delete(f, "RegistryImageIDs")
	delete(f, "LocalTaskID")
	delete(f, "RegistryTaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVulScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulScanTaskResponseParams struct {
	// 本地镜像重新漏洞扫描时的任务ID
	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`

	// 仓库镜像重新漏洞扫描时的任务ID
	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVulScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateVulScanTaskResponseParams `json:"Response"`
}

func (r *CreateVulScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebVulExportJobRequestParams struct {
	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞子类型</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type CreateWebVulExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为50000，最大值为50000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞子类型</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *CreateWebVulExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebVulExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWebVulExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebVulExportJobResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWebVulExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebVulExportJobResponseParams `json:"Response"`
}

func (r *CreateWebVulExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebVulExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAbnormalProcessRulesRequestParams struct {
	// 策略的ids
	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`
}

type DeleteAbnormalProcessRulesRequest struct {
	*tchttp.BaseRequest
	
	// 策略的ids
	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`
}

func (r *DeleteAbnormalProcessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAbnormalProcessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAbnormalProcessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAbnormalProcessRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAbnormalProcessRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAbnormalProcessRulesResponseParams `json:"Response"`
}

func (r *DeleteAbnormalProcessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAbnormalProcessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessControlRulesRequestParams struct {
	// 策略的ids
	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`
}

type DeleteAccessControlRulesRequest struct {
	*tchttp.BaseRequest
	
	// 策略的ids
	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`
}

func (r *DeleteAccessControlRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessControlRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccessControlRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessControlRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccessControlRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccessControlRulesResponseParams `json:"Response"`
}

func (r *DeleteAccessControlRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessControlRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteComplianceAssetPolicySetFromWhitelistRequestParams struct {
	// 资产ID
	AssetItemId *uint64 `json:"AssetItemId,omitempty" name:"AssetItemId"`

	// 需要忽略指定资产内的检查项ID列表
	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

type DeleteComplianceAssetPolicySetFromWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// 资产ID
	AssetItemId *uint64 `json:"AssetItemId,omitempty" name:"AssetItemId"`

	// 需要忽略指定资产内的检查项ID列表
	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

func (r *DeleteComplianceAssetPolicySetFromWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteComplianceAssetPolicySetFromWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetItemId")
	delete(f, "CustomerPolicyItemIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteComplianceAssetPolicySetFromWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteComplianceAssetPolicySetFromWhitelistResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteComplianceAssetPolicySetFromWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *DeleteComplianceAssetPolicySetFromWhitelistResponseParams `json:"Response"`
}

func (r *DeleteComplianceAssetPolicySetFromWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteComplianceAssetPolicySetFromWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCompliancePolicyAssetSetFromWhitelistRequestParams struct {
	// （检查项ID+资产ID列表）的列表
	PolicyAssetSetList []*CompliancePolicyAssetSetItem `json:"PolicyAssetSetList,omitempty" name:"PolicyAssetSetList"`
}

type DeleteCompliancePolicyAssetSetFromWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// （检查项ID+资产ID列表）的列表
	PolicyAssetSetList []*CompliancePolicyAssetSetItem `json:"PolicyAssetSetList,omitempty" name:"PolicyAssetSetList"`
}

func (r *DeleteCompliancePolicyAssetSetFromWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompliancePolicyAssetSetFromWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyAssetSetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCompliancePolicyAssetSetFromWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCompliancePolicyAssetSetFromWhitelistResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCompliancePolicyAssetSetFromWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCompliancePolicyAssetSetFromWhitelistResponseParams `json:"Response"`
}

func (r *DeleteCompliancePolicyAssetSetFromWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompliancePolicyAssetSetFromWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCompliancePolicyItemFromWhitelistRequestParams struct {
	// 指定的白名单项的ID的列表
	WhitelistIdSet []*uint64 `json:"WhitelistIdSet,omitempty" name:"WhitelistIdSet"`
}

type DeleteCompliancePolicyItemFromWhitelistRequest struct {
	*tchttp.BaseRequest
	
	// 指定的白名单项的ID的列表
	WhitelistIdSet []*uint64 `json:"WhitelistIdSet,omitempty" name:"WhitelistIdSet"`
}

func (r *DeleteCompliancePolicyItemFromWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompliancePolicyItemFromWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhitelistIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCompliancePolicyItemFromWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCompliancePolicyItemFromWhitelistResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCompliancePolicyItemFromWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCompliancePolicyItemFromWhitelistResponseParams `json:"Response"`
}

func (r *DeleteCompliancePolicyItemFromWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompliancePolicyItemFromWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEscapeWhiteListRequestParams struct {
	// 白名单记录ID数组
	IDSet []*int64 `json:"IDSet,omitempty" name:"IDSet"`
}

type DeleteEscapeWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 白名单记录ID数组
	IDSet []*int64 `json:"IDSet,omitempty" name:"IDSet"`
}

func (r *DeleteEscapeWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEscapeWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IDSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEscapeWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEscapeWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEscapeWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEscapeWhiteListResponseParams `json:"Response"`
}

func (r *DeleteEscapeWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEscapeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIgnoreVulRequestParams struct {
	// 漏洞PocID 信息列表
	List []*ModifyIgnoreVul `json:"List,omitempty" name:"List"`
}

type DeleteIgnoreVulRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞PocID 信息列表
	List []*ModifyIgnoreVul `json:"List,omitempty" name:"List"`
}

func (r *DeleteIgnoreVulRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIgnoreVulRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "List")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIgnoreVulRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIgnoreVulResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIgnoreVulResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIgnoreVulResponseParams `json:"Response"`
}

func (r *DeleteIgnoreVulResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIgnoreVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteK8sApiAbnormalRuleRequestParams struct {
	// 规则ID集合
	RuleIDSet []*string `json:"RuleIDSet,omitempty" name:"RuleIDSet"`
}

type DeleteK8sApiAbnormalRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID集合
	RuleIDSet []*string `json:"RuleIDSet,omitempty" name:"RuleIDSet"`
}

func (r *DeleteK8sApiAbnormalRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteK8sApiAbnormalRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIDSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteK8sApiAbnormalRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteK8sApiAbnormalRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteK8sApiAbnormalRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteK8sApiAbnormalRuleResponseParams `json:"Response"`
}

func (r *DeleteK8sApiAbnormalRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteK8sApiAbnormalRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineRequestParams struct {
	// 客户端Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type DeleteMachineRequest struct {
	*tchttp.BaseRequest
	
	// 客户端Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMachineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMachineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMachineResponseParams `json:"Response"`
}

func (r *DeleteMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkFirewallPolicyDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id数组
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

type DeleteNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id数组
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteNetworkFirewallPolicyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkFirewallPolicyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkFirewallPolicyDetailResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建删除任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetworkFirewallPolicyDetailResponseParams `json:"Response"`
}

func (r *DeleteNetworkFirewallPolicyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellEventsRequestParams struct {
	// 事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
}

type DeleteReverseShellEventsRequest struct {
	*tchttp.BaseRequest
	
	// 事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
}

func (r *DeleteReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellEventsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReverseShellEventsResponseParams `json:"Response"`
}

func (r *DeleteReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellWhiteListsRequestParams struct {
	// 白名单ids
	WhiteListIdSet []*string `json:"WhiteListIdSet,omitempty" name:"WhiteListIdSet"`
}

type DeleteReverseShellWhiteListsRequest struct {
	*tchttp.BaseRequest
	
	// 白名单ids
	WhiteListIdSet []*string `json:"WhiteListIdSet,omitempty" name:"WhiteListIdSet"`
}

func (r *DeleteReverseShellWhiteListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellWhiteListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteListIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReverseShellWhiteListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReverseShellWhiteListsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteReverseShellWhiteListsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReverseShellWhiteListsResponseParams `json:"Response"`
}

func (r *DeleteReverseShellWhiteListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskSyscallEventsRequestParams struct {
	// 事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
}

type DeleteRiskSyscallEventsRequest struct {
	*tchttp.BaseRequest
	
	// 事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`
}

func (r *DeleteRiskSyscallEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskSyscallEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRiskSyscallEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskSyscallEventsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRiskSyscallEventsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRiskSyscallEventsResponseParams `json:"Response"`
}

func (r *DeleteRiskSyscallEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskSyscallEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskSyscallWhiteListsRequestParams struct {
	// 白名单ids
	WhiteListIdSet []*string `json:"WhiteListIdSet,omitempty" name:"WhiteListIdSet"`
}

type DeleteRiskSyscallWhiteListsRequest struct {
	*tchttp.BaseRequest
	
	// 白名单ids
	WhiteListIdSet []*string `json:"WhiteListIdSet,omitempty" name:"WhiteListIdSet"`
}

func (r *DeleteRiskSyscallWhiteListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskSyscallWhiteListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteListIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRiskSyscallWhiteListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskSyscallWhiteListsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRiskSyscallWhiteListsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRiskSyscallWhiteListsResponseParams `json:"Response"`
}

func (r *DeleteRiskSyscallWhiteListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskSyscallWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSearchTemplateRequestParams struct {
	// 模板ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DeleteSearchTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteSearchTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSearchTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSearchTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSearchTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSearchTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSearchTemplateResponseParams `json:"Response"`
}

func (r *DeleteSearchTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSearchTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeABTestConfigRequestParams struct {
	// 灰度项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type DescribeABTestConfigRequest struct {
	*tchttp.BaseRequest
	
	// 灰度项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *DescribeABTestConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeABTestConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeABTestConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeABTestConfigResponseParams struct {
	// 灰度项目配置
	Config []*ABTestConfig `json:"Config,omitempty" name:"Config"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeABTestConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeABTestConfigResponseParams `json:"Response"`
}

func (r *DescribeABTestConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeABTestConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessDetailRequestParams struct {
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DescribeAbnormalProcessDetailRequest struct {
	*tchttp.BaseRequest
	
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeAbnormalProcessDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalProcessDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessDetailResponseParams struct {
	// 事件基本信息
	EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`

	// 进程信息
	ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`

	// 父进程信息
	ParentProcessInfo *ProcessDetailBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`

	// 事件描述
	EventDetail *AbnormalProcessEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`

	// 祖先进程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAbnormalProcessDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalProcessDetailResponseParams `json:"Response"`
}

func (r *DescribeAbnormalProcessDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessEventTendencyRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeAbnormalProcessEventTendencyRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAbnormalProcessEventTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessEventTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalProcessEventTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessEventTendencyResponseParams struct {
	// 待处理异常进程事件趋势
	List []*AbnormalProcessEventTendencyInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAbnormalProcessEventTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalProcessEventTendencyResponseParams `json:"Response"`
}

func (r *DescribeAbnormalProcessEventTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessEventsExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAbnormalProcessEventsExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAbnormalProcessEventsExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessEventsExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalProcessEventsExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessEventsExportResponseParams struct {
	// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAbnormalProcessEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalProcessEventsExportResponseParams `json:"Response"`
}

func (r *DescribeAbnormalProcessEventsExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessEventsRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAbnormalProcessEventsRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAbnormalProcessEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalProcessEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessEventsResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 异常进程数组
	EventSet []*AbnormalProcessEventInfo `json:"EventSet,omitempty" name:"EventSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAbnormalProcessEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalProcessEventsResponseParams `json:"Response"`
}

func (r *DescribeAbnormalProcessEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessLevelSummaryRequestParams struct {

}

type DescribeAbnormalProcessLevelSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAbnormalProcessLevelSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessLevelSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalProcessLevelSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessLevelSummaryResponseParams struct {
	// 异常进程高危待处理事件数
	HighLevelEventCount *int64 `json:"HighLevelEventCount,omitempty" name:"HighLevelEventCount"`

	// 异常进程中危待处理事件数
	MediumLevelEventCount *int64 `json:"MediumLevelEventCount,omitempty" name:"MediumLevelEventCount"`

	// 异常进程低危待处理事件数
	LowLevelEventCount *int64 `json:"LowLevelEventCount,omitempty" name:"LowLevelEventCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAbnormalProcessLevelSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalProcessLevelSummaryResponseParams `json:"Response"`
}

func (r *DescribeAbnormalProcessLevelSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessLevelSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessRuleDetailRequestParams struct {
	// 策略唯一id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 镜像id, 在添加白名单的时候使用
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAbnormalProcessRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 策略唯一id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 镜像id, 在添加白名单的时候使用
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAbnormalProcessRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "ImageId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalProcessRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessRuleDetailResponseParams struct {
	// 异常进程策略详细信息
	RuleDetail *AbnormalProcessRuleInfo `json:"RuleDetail,omitempty" name:"RuleDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAbnormalProcessRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalProcessRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeAbnormalProcessRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessRulesExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAbnormalProcessRulesExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAbnormalProcessRulesExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessRulesExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalProcessRulesExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessRulesExportResponseParams struct {
	// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAbnormalProcessRulesExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalProcessRulesExportResponseParams `json:"Response"`
}

func (r *DescribeAbnormalProcessRulesExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessRulesExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessRulesRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAbnormalProcessRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAbnormalProcessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalProcessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalProcessRulesResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 异常进程策略信息列表
	RuleSet []*RuleBaseInfo `json:"RuleSet,omitempty" name:"RuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAbnormalProcessRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalProcessRulesResponseParams `json:"Response"`
}

func (r *DescribeAbnormalProcessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalProcessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlDetailRequestParams struct {
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DescribeAccessControlDetailRequest struct {
	*tchttp.BaseRequest
	
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeAccessControlDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessControlDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlDetailResponseParams struct {
	// 事件基本信息
	EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`

	// 进程信息
	ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`

	// 被篡改信息
	TamperedFileInfo *FileAttributeInfo `json:"TamperedFileInfo,omitempty" name:"TamperedFileInfo"`

	// 事件描述
	EventDetail *AccessControlEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`

	// 父进程信息
	ParentProcessInfo *ProcessBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`

	// 祖先进程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessControlDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessControlDetailResponseParams `json:"Response"`
}

func (r *DescribeAccessControlDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlEventsExportRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type DescribeAccessControlEventsExportRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeAccessControlEventsExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlEventsExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessControlEventsExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlEventsExportResponseParams struct {
	// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessControlEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessControlEventsExportResponseParams `json:"Response"`
}

func (r *DescribeAccessControlEventsExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlEventsRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAccessControlEventsRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAccessControlEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessControlEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlEventsResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 访问控制事件数组
	EventSet []*AccessControlEventInfo `json:"EventSet,omitempty" name:"EventSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessControlEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessControlEventsResponseParams `json:"Response"`
}

func (r *DescribeAccessControlEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlRuleDetailRequestParams struct {
	// 策略唯一id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 镜像id, 仅仅在事件加白的时候使用
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAccessControlRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 策略唯一id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 镜像id, 仅仅在事件加白的时候使用
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAccessControlRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "ImageId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessControlRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlRuleDetailResponseParams struct {
	// 运行时策略详细信息
	RuleDetail *AccessControlRuleInfo `json:"RuleDetail,omitempty" name:"RuleDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessControlRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessControlRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeAccessControlRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlRulesExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAccessControlRulesExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAccessControlRulesExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlRulesExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessControlRulesExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlRulesExportResponseParams struct {
	// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessControlRulesExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessControlRulesExportResponseParams `json:"Response"`
}

func (r *DescribeAccessControlRulesExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlRulesExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlRulesRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAccessControlRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAccessControlRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessControlRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessControlRulesResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 访问控制策略信息列表
	RuleSet []*RuleBaseInfo `json:"RuleSet,omitempty" name:"RuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessControlRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessControlRulesResponseParams `json:"Response"`
}

func (r *DescribeAccessControlRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessControlRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAffectedClusterCountRequestParams struct {

}

type DescribeAffectedClusterCountRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAffectedClusterCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAffectedClusterCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAffectedClusterCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAffectedClusterCountResponseParams struct {
	// 严重风险的集群数量
	SeriousRiskClusterCount *uint64 `json:"SeriousRiskClusterCount,omitempty" name:"SeriousRiskClusterCount"`

	// 高危风险的集群数量
	HighRiskClusterCount *uint64 `json:"HighRiskClusterCount,omitempty" name:"HighRiskClusterCount"`

	// 中危风险的集群数量
	MiddleRiskClusterCount *uint64 `json:"MiddleRiskClusterCount,omitempty" name:"MiddleRiskClusterCount"`

	// 低危风险的集群数量
	HintRiskClusterCount *uint64 `json:"HintRiskClusterCount,omitempty" name:"HintRiskClusterCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAffectedClusterCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAffectedClusterCountResponseParams `json:"Response"`
}

func (r *DescribeAffectedClusterCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAffectedClusterCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAffectedNodeListRequestParams struct {
	// 唯一的检测项的ID
	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName, ClusterId,InstanceId,PrivateIpAddresses
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAffectedNodeListRequest struct {
	*tchttp.BaseRequest
	
	// 唯一的检测项的ID
	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName, ClusterId,InstanceId,PrivateIpAddresses
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAffectedNodeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAffectedNodeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CheckItemId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAffectedNodeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAffectedNodeListResponseParams struct {
	// 受影响的节点总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 受影响的节点列表
	AffectedNodeList []*AffectedNodeItem `json:"AffectedNodeList,omitempty" name:"AffectedNodeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAffectedNodeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAffectedNodeListResponseParams `json:"Response"`
}

func (r *DescribeAffectedNodeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAffectedNodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAffectedWorkloadListRequestParams struct {
	// 唯一的检测项的ID
	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：WorkloadType,ClusterId
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAffectedWorkloadListRequest struct {
	*tchttp.BaseRequest
	
	// 唯一的检测项的ID
	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：WorkloadType,ClusterId
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAffectedWorkloadListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAffectedWorkloadListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CheckItemId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAffectedWorkloadListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAffectedWorkloadListResponseParams struct {
	// 受影响的workload列表数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 受影响的workload列表
	AffectedWorkloadList []*AffectedWorkloadItem `json:"AffectedWorkloadList,omitempty" name:"AffectedWorkloadList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAffectedWorkloadListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAffectedWorkloadListResponseParams `json:"Response"`
}

func (r *DescribeAffectedWorkloadListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAffectedWorkloadListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDaemonSetCmdRequestParams struct {
	// 是否是腾讯云
	IsCloud *bool `json:"IsCloud,omitempty" name:"IsCloud"`

	// 网络类型：basic-基础网络，private-VPC, public-公网，direct-专线
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 地域标示, NetType=direct时必填
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// VpcId, NetType=direct时必填
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 命令有效期，非腾讯云时必填
	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
}

type DescribeAgentDaemonSetCmdRequest struct {
	*tchttp.BaseRequest
	
	// 是否是腾讯云
	IsCloud *bool `json:"IsCloud,omitempty" name:"IsCloud"`

	// 网络类型：basic-基础网络，private-VPC, public-公网，direct-专线
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 地域标示, NetType=direct时必填
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// VpcId, NetType=direct时必填
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 命令有效期，非腾讯云时必填
	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`
}

func (r *DescribeAgentDaemonSetCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDaemonSetCmdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsCloud")
	delete(f, "NetType")
	delete(f, "RegionCode")
	delete(f, "VpcId")
	delete(f, "ExpireDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentDaemonSetCmdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDaemonSetCmdResponseParams struct {
	// 安装命令
	Command *string `json:"Command,omitempty" name:"Command"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentDaemonSetCmdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentDaemonSetCmdResponseParams `json:"Response"`
}

func (r *DescribeAgentDaemonSetCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDaemonSetCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentInstallCommandRequestParams struct {
	// 是否是腾讯云
	IsCloud *bool `json:"IsCloud,omitempty" name:"IsCloud"`

	// 网络类型：basic-基础网络，private-VPC, public-公网，direct-专线
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 地域标示, NetType=direct时必填
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// VpcId, NetType=direct时必填
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 命令有效期，非腾讯云时必填
	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`

	// 标签ID列表，IsCloud=false时才会生效
	TagIds []*uint64 `json:"TagIds,omitempty" name:"TagIds"`
}

type DescribeAgentInstallCommandRequest struct {
	*tchttp.BaseRequest
	
	// 是否是腾讯云
	IsCloud *bool `json:"IsCloud,omitempty" name:"IsCloud"`

	// 网络类型：basic-基础网络，private-VPC, public-公网，direct-专线
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 地域标示, NetType=direct时必填
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// VpcId, NetType=direct时必填
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 命令有效期，非腾讯云时必填
	ExpireDate *string `json:"ExpireDate,omitempty" name:"ExpireDate"`

	// 标签ID列表，IsCloud=false时才会生效
	TagIds []*uint64 `json:"TagIds,omitempty" name:"TagIds"`
}

func (r *DescribeAgentInstallCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentInstallCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsCloud")
	delete(f, "NetType")
	delete(f, "RegionCode")
	delete(f, "VpcId")
	delete(f, "ExpireDate")
	delete(f, "TagIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentInstallCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentInstallCommandResponseParams struct {
	// linux系统安装命令
	LinuxCommand *string `json:"LinuxCommand,omitempty" name:"LinuxCommand"`

	// windows系统安装命令（windows2008及以上）
	WindowsCommand *string `json:"WindowsCommand,omitempty" name:"WindowsCommand"`

	// windows系统安装命令第一步（windows2003）
	WindowsStepOne *string `json:"WindowsStepOne,omitempty" name:"WindowsStepOne"`

	// windows系统安装命令第二步（windows2003）
	WindowsStepTwo *string `json:"WindowsStepTwo,omitempty" name:"WindowsStepTwo"`

	// windows版agent下载链接
	WindowsDownloadUrl *string `json:"WindowsDownloadUrl,omitempty" name:"WindowsDownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentInstallCommandResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentInstallCommandResponseParams `json:"Response"`
}

func (r *DescribeAgentInstallCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentInstallCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetAppServiceListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetAppServiceListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetAppServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetAppServiceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetAppServiceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetAppServiceListResponseParams struct {
	// db服务列表
	List []*ServiceInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetAppServiceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetAppServiceListResponseParams `json:"Response"`
}

func (r *DescribeAssetAppServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetAppServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetClusterListRequestParams struct {
	// 过滤条件。
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>Status - string  - 是否必填: 否 -集群状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段。
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAssetClusterListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>Status - string  - 是否必填: 否 -集群状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段。
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetClusterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetClusterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetClusterListResponseParams struct {
	// 集群列表
	List []*AssetClusterListItem `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetClusterListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetClusterListResponseParams `json:"Response"`
}

func (r *DescribeAssetClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetComponentListRequestParams struct {
	// 容器id
	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetComponentListRequest struct {
	*tchttp.BaseRequest
	
	// 容器id
	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetComponentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetComponentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContainerID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetComponentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetComponentListResponseParams struct {
	// 组件列表
	List []*ComponentInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetComponentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetComponentListResponseParams `json:"Response"`
}

func (r *DescribeAssetComponentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetComponentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetContainerDetailRequestParams struct {
	// 容器id
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
}

type DescribeAssetContainerDetailRequest struct {
	*tchttp.BaseRequest
	
	// 容器id
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
}

func (r *DescribeAssetContainerDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetContainerDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContainerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetContainerDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetContainerDetailResponseParams struct {
	// 主机id
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 主机ip
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 运行状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 运行账户
	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`

	// 命令行
	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`

	// CPU使用率 * 1000
	CPUUsage *uint64 `json:"CPUUsage,omitempty" name:"CPUUsage"`

	// 内存使用 KB
	RamUsage *uint64 `json:"RamUsage,omitempty" name:"RamUsage"`

	// 镜像名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 归属POD
	POD *string `json:"POD,omitempty" name:"POD"`

	// k8s 主节点
	K8sMaster *string `json:"K8sMaster,omitempty" name:"K8sMaster"`

	// 容器内进程数
	ProcessCnt *uint64 `json:"ProcessCnt,omitempty" name:"ProcessCnt"`

	// 容器内端口数
	PortCnt *uint64 `json:"PortCnt,omitempty" name:"PortCnt"`

	// 组件数
	ComponentCnt *uint64 `json:"ComponentCnt,omitempty" name:"ComponentCnt"`

	// app数
	AppCnt *uint64 `json:"AppCnt,omitempty" name:"AppCnt"`

	// websvc数
	WebServiceCnt *uint64 `json:"WebServiceCnt,omitempty" name:"WebServiceCnt"`

	// 挂载
	Mounts []*ContainerMount `json:"Mounts,omitempty" name:"Mounts"`

	// 容器网络信息
	Network *ContainerNetwork `json:"Network,omitempty" name:"Network"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像创建时间
	ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`

	// 镜像大小
	ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 主机状态 offline,online,pause
	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`

	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`

	// 网络子状态
	NetSubStatus *string `json:"NetSubStatus,omitempty" name:"NetSubStatus"`

	// 隔离来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateSource *string `json:"IsolateSource,omitempty" name:"IsolateSource"`

	// 隔离时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetContainerDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetContainerDetailResponseParams `json:"Response"`
}

func (r *DescribeAssetContainerDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetContainerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetContainerListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ContainerName - String - 是否必填：否 - 容器名称模糊搜索</li>
	// <li>Status - String - 是否必填：否 - 容器运行状态筛选，0："created",1："running", 2："paused", 3："restarting", 4："removing", 5："exited", 6："dead" </li>
	// <li>Runas - String - 是否必填：否 - 运行用户筛选</li>
	// <li>ImageName- String - 是否必填：否 - 镜像名称搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>OrderBy - String 是否必填：否 -排序字段，支持：cpu_usage, mem_usage的动态排序 ["cpu_usage","+"]  '+'升序、'-'降序</li>
	// <li>NetStatus - String -是否必填: 否 -  容器网络状态筛选 normal isolated isolating isolate_failed restoring restore_failed</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAssetContainerListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ContainerName - String - 是否必填：否 - 容器名称模糊搜索</li>
	// <li>Status - String - 是否必填：否 - 容器运行状态筛选，0："created",1："running", 2："paused", 3："restarting", 4："removing", 5："exited", 6："dead" </li>
	// <li>Runas - String - 是否必填：否 - 运行用户筛选</li>
	// <li>ImageName- String - 是否必填：否 - 镜像名称搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>OrderBy - String 是否必填：否 -排序字段，支持：cpu_usage, mem_usage的动态排序 ["cpu_usage","+"]  '+'升序、'-'降序</li>
	// <li>NetStatus - String -是否必填: 否 -  容器网络状态筛选 normal isolated isolating isolate_failed restoring restore_failed</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetContainerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetContainerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetContainerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetContainerListResponseParams struct {
	// 容器列表
	List []*ContainerInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetContainerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetContainerListResponseParams `json:"Response"`
}

func (r *DescribeAssetContainerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetContainerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDBServiceListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetDBServiceListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetDBServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDBServiceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetDBServiceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDBServiceListResponseParams struct {
	// db服务列表
	List []*ServiceInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetDBServiceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetDBServiceListResponseParams `json:"Response"`
}

func (r *DescribeAssetDBServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDBServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetHostDetailRequestParams struct {
	// 主机id
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

type DescribeAssetHostDetailRequest struct {
	*tchttp.BaseRequest
	
	// 主机id
	HostId *string `json:"HostId,omitempty" name:"HostId"`
}

func (r *DescribeAssetHostDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetHostDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetHostDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetHostDetailResponseParams struct {
	// 云镜uuid
	UUID *string `json:"UUID,omitempty" name:"UUID"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 主机名
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 主机分组
	Group *string `json:"Group,omitempty" name:"Group"`

	// 主机IP
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 操作系统
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// agent版本
	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`

	// 内核版本
	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`

	// docker版本
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

	// docker api版本
	DockerAPIVersion *string `json:"DockerAPIVersion,omitempty" name:"DockerAPIVersion"`

	// docker go 版本
	DockerGoVersion *string `json:"DockerGoVersion,omitempty" name:"DockerGoVersion"`

	// docker 文件系统类型
	DockerFileSystemDriver *string `json:"DockerFileSystemDriver,omitempty" name:"DockerFileSystemDriver"`

	// docker root 目录
	DockerRootDir *string `json:"DockerRootDir,omitempty" name:"DockerRootDir"`

	// 镜像数
	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`

	// 容器数
	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`

	// k8s IP
	K8sMasterIP *string `json:"K8sMasterIP,omitempty" name:"K8sMasterIP"`

	// k8s version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// kube proxy
	KubeProxyVersion *string `json:"KubeProxyVersion,omitempty" name:"KubeProxyVersion"`

	// "UNINSTALL"："未安装","OFFLINE"："离线", "ONLINE"："防护中
	Status *string `json:"Status,omitempty" name:"Status"`

	// 是否Containerd
	IsContainerd *bool `json:"IsContainerd,omitempty" name:"IsContainerd"`

	// 主机来源;"TENCENTCLOUD":"腾讯云服务器","OTHERCLOUD":"非腾讯云服务器"
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 外网ip
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 主机实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 地域ID
	RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`

	// 所属项目
	Project *ProjectInfo `json:"Project,omitempty" name:"Project"`

	// 标签
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetHostDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetHostDetailResponseParams `json:"Response"`
}

func (r *DescribeAssetHostDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetHostDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetHostListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - agent状态筛选，"ALL":"全部"(或不传该字段),"UNINSTALL"："未安装","OFFLINE"："离线", "ONLINE"："防护中"</li>
	// <li>HostName - String - 是否必填：否 - 主机名筛选</li>
	// <li>Group- String - 是否必填：否 - 主机群组搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>HostID- string - 是否必填：否 - 主机id搜索</li>
	// <li>DockerVersion- string - 是否必填：否 - docker版本搜索</li>
	// <li>MachineType- string - 是否必填：否 - 主机来源MachineType搜索，"ALL":"全部"(或不传该字段),主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；</li>
	// <li>DockerStatus- string - 是否必填：否 - docker安装状态，"ALL":"全部"(或不传该字段),"INSTALL":"已安装","UNINSTALL":"未安装"</li>
	// <li>ProjectID- string - 是否必填：否 - 所属项目id搜索</li>
	// <li>Tag:xxx(tag:key)- string- 是否必填：否 - 标签键值搜索 示例Filters":[{"Name":"tag:tke-kind","Values":["service"]}]</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAssetHostListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status - String - 是否必填：否 - agent状态筛选，"ALL":"全部"(或不传该字段),"UNINSTALL"："未安装","OFFLINE"："离线", "ONLINE"："防护中"</li>
	// <li>HostName - String - 是否必填：否 - 主机名筛选</li>
	// <li>Group- String - 是否必填：否 - 主机群组搜索</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>HostID- string - 是否必填：否 - 主机id搜索</li>
	// <li>DockerVersion- string - 是否必填：否 - docker版本搜索</li>
	// <li>MachineType- string - 是否必填：否 - 主机来源MachineType搜索，"ALL":"全部"(或不传该字段),主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；</li>
	// <li>DockerStatus- string - 是否必填：否 - docker安装状态，"ALL":"全部"(或不传该字段),"INSTALL":"已安装","UNINSTALL":"未安装"</li>
	// <li>ProjectID- string - 是否必填：否 - 所属项目id搜索</li>
	// <li>Tag:xxx(tag:key)- string- 是否必填：否 - 标签键值搜索 示例Filters":[{"Name":"tag:tke-kind","Values":["service"]}]</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetHostListResponseParams struct {
	// 主机列表
	List []*HostInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetHostListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetHostListResponseParams `json:"Response"`
}

func (r *DescribeAssetHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageBindRuleInfoRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"EventType","Values":[""]}]
	// EventType取值：
	// "FILE_ABNORMAL_READ" 访问控制
	// "MALICE_PROCESS_START" 恶意进程启动
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAssetImageBindRuleInfoRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"EventType","Values":[""]}]
	// EventType取值：
	// "FILE_ABNORMAL_READ" 访问控制
	// "MALICE_PROCESS_START" 恶意进程启动
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetImageBindRuleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageBindRuleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageBindRuleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageBindRuleInfoResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像绑定规则列表
	ImageBindRuleSet []*ImagesBindRuleInfo `json:"ImageBindRuleSet,omitempty" name:"ImageBindRuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageBindRuleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageBindRuleInfoResponseParams `json:"Response"`
}

func (r *DescribeAssetImageBindRuleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageBindRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageDetailRequestParams struct {
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

type DescribeAssetImageDetailRequest struct {
	*tchttp.BaseRequest
	
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

func (r *DescribeAssetImageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageDetailResponseParams struct {
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 关联主机个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`

	// 关联容器个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 漏洞个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`

	// 风险行为数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`

	// 敏感信息数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveInfoCnt *uint64 `json:"SensitiveInfoCnt,omitempty" name:"SensitiveInfoCnt"`

	// 是否信任镜像
	IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`

	// 镜像系统
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// agent镜像扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentError *string `json:"AgentError,omitempty" name:"AgentError"`

	// 后端镜像扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanError *string `json:"ScanError,omitempty" name:"ScanError"`

	// 系统架构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 作者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Author *string `json:"Author,omitempty" name:"Author"`

	// 构建历史
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildHistory *string `json:"BuildHistory,omitempty" name:"BuildHistory"`

	// 木马扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVirusProgress *uint64 `json:"ScanVirusProgress,omitempty" name:"ScanVirusProgress"`

	// 漏洞扫进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVulProgress *uint64 `json:"ScanVulProgress,omitempty" name:"ScanVulProgress"`

	// 敏感信息扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRiskProgress *uint64 `json:"ScanRiskProgress,omitempty" name:"ScanRiskProgress"`

	// 木马扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`

	// 漏洞扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`

	// 敏感信息错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`

	// 镜像扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 木马病毒数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`

	// 镜像扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 剩余扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainScanTime *uint64 `json:"RemainScanTime,omitempty" name:"RemainScanTime"`

	// 授权为：1，未授权为：0
	IsAuthorized *int64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageDetailResponseParams `json:"Response"`
}

func (r *DescribeAssetImageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageHostListRequestParams struct {
	// 过滤条件 支持ImageID,HostID
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetImageHostListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件 支持ImageID,HostID
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetImageHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageHostListResponseParams struct {
	// 镜像列表
	List []*ImageHost `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageHostListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageHostListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageListExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAssetImageListExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageListExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageListExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageListExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageListExportResponseParams struct {
	// excel文件下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageListExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageListExportResponseParams `json:"Response"`
}

func (r *DescribeAssetImageListExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAssetImageListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ImageName- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>ScanStatus - String - 是否必填：否 - 镜像扫描状态notScan，scanning，scanned，scanErr</li>
	// <li>ImageID- String - 是否必填：否 - 镜像ID筛选，</li>
	// <li>SecurityRisk- String - 是否必填：否 - 安全风险，VulCnt 、VirusCnt、RiskCnt、IsTrustImage</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageListResponseParams struct {
	// 镜像列表
	List []*ImagesInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryAssetStatusRequestParams struct {

}

type DescribeAssetImageRegistryAssetStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAssetImageRegistryAssetStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryAssetStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryAssetStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryAssetStatusResponseParams struct {
	// 更新进度状态,doing更新中，success更新成功，failed失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Err *string `json:"Err,omitempty" name:"Err"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryAssetStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryAssetStatusResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryAssetStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryAssetStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryDetailRequestParams struct {
	// 仓库列表id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

type DescribeAssetImageRegistryDetailRequest struct {
	*tchttp.BaseRequest
	
	// 仓库列表id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeAssetImageRegistryDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryDetailResponseParams struct {
	// 镜像Digest
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`

	// 镜像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`

	// 镜像类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 仓库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// 扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 安全漏洞数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`

	// 木马病毒数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`

	// 风险行为数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`

	// 敏感信息数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SentiveInfoCnt *uint64 `json:"SentiveInfoCnt,omitempty" name:"SentiveInfoCnt"`

	// 镜像系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 木马扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`

	// 漏洞扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`

	// 层文件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LayerInfo *string `json:"LayerInfo,omitempty" name:"LayerInfo"`

	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 高危扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`

	// 木马信息扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVirusProgress *uint64 `json:"ScanVirusProgress,omitempty" name:"ScanVirusProgress"`

	// 漏洞扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVulProgress *uint64 `json:"ScanVulProgress,omitempty" name:"ScanVulProgress"`

	// 敏感扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRiskProgress *uint64 `json:"ScanRiskProgress,omitempty" name:"ScanRiskProgress"`

	// 剩余扫描时间秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRemainTime *uint64 `json:"ScanRemainTime,omitempty" name:"ScanRemainTime"`

	// cve扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CveStatus *string `json:"CveStatus,omitempty" name:"CveStatus"`

	// 高危扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskStatus *string `json:"RiskStatus,omitempty" name:"RiskStatus"`

	// 木马扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusStatus *string `json:"VirusStatus,omitempty" name:"VirusStatus"`

	// 总进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 授权状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`

	// 镜像大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 镜像Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`

	// 镜像创建的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryDetailResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryListExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式，asc，desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 是否仅展示repository版本最新的镜像，默认为false
	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
}

type DescribeAssetImageRegistryListExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式，asc，desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 是否仅展示repository版本最新的镜像，默认为false
	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
}

func (r *DescribeAssetImageRegistryListExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryListExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "OnlyShowLatest")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryListExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryListExportResponseParams struct {
	// excel文件下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryListExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryListExportResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryListExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤字段
	// IsAuthorized是否授权，取值全部all，未授权0，已授权1
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式，asc，desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 是否仅展示各repository最新的镜像, 默认为false
	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
}

type DescribeAssetImageRegistryListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤字段
	// IsAuthorized是否授权，取值全部all，未授权0，已授权1
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式，asc，desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 是否仅展示各repository最新的镜像, 默认为false
	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
}

func (r *DescribeAssetImageRegistryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "OnlyShowLatest")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryListResponseParams struct {
	// 镜像仓库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ImageRepoInfo `json:"List,omitempty" name:"List"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryRegistryDetailRequestParams struct {
	// 仓库唯一id
	RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
}

type DescribeAssetImageRegistryRegistryDetailRequest struct {
	*tchttp.BaseRequest
	
	// 仓库唯一id
	RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeAssetImageRegistryRegistryDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryRegistryDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryRegistryDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryRegistryDetailResponseParams struct {
	// 仓库名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 仓库url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 仓库类型，列表：harbor
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 仓库版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryVersion *string `json:"RegistryVersion,omitempty" name:"RegistryVersion"`

	// 网络类型，列表：public（公网）
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 区域，列表:default（默认）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`

	// 限速
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeedLimit *uint64 `json:"SpeedLimit,omitempty" name:"SpeedLimit"`

	// 安全模式（证书校验）：0（默认） 非安全模式（跳过证书校验）：1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Insecure *uint64 `json:"Insecure,omitempty" name:"Insecure"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryRegistryDetailResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryRegistryDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryRegistryListRequestParams struct {

}

type DescribeAssetImageRegistryRegistryListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAssetImageRegistryRegistryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryRegistryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryRegistryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryRegistryListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryRegistryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryRegistryListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryRegistryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryRegistryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryRiskInfoListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像id
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 排序字段（Level）
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 + -
	Order *string `json:"Order,omitempty" name:"Order"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAssetImageRegistryRiskInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像id
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 排序字段（Level）
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 + -
	Order *string `json:"Order,omitempty" name:"Order"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryRiskInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryRiskInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ImageInfo")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryRiskInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryRiskInfoListResponseParams struct {
	// 镜像漏洞列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ImageRisk `json:"List,omitempty" name:"List"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryRiskInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryRiskInfoListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryRiskInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryRiskInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryRiskListExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAssetImageRegistryRiskListExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryRiskListExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryRiskListExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ImageInfo")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryRiskListExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryRiskListExportResponseParams struct {
	// excel文件下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryRiskListExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryRiskListExportResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryRiskListExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryRiskListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryScanStatusOneKeyRequestParams struct {
	// 需要获取进度的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 是否获取全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 需要获取进度的镜像列表Id
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAssetImageRegistryScanStatusOneKeyRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取进度的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 是否获取全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 需要获取进度的镜像列表Id
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryScanStatusOneKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryScanStatusOneKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Images")
	delete(f, "All")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryScanStatusOneKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryScanStatusOneKeyResponseParams struct {
	// 镜像个数
	ImageTotal *uint64 `json:"ImageTotal,omitempty" name:"ImageTotal"`

	// 扫描镜像个数
	ImageScanCnt *uint64 `json:"ImageScanCnt,omitempty" name:"ImageScanCnt"`

	// 扫描进度列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageStatus []*ImageProgress `json:"ImageStatus,omitempty" name:"ImageStatus"`

	// 安全个数
	SuccessCount *uint64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 风险个数
	RiskCount *uint64 `json:"RiskCount,omitempty" name:"RiskCount"`

	// 总的扫描进度
	Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`

	// 总的扫描状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 扫描剩余时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRemainTime *uint64 `json:"ScanRemainTime,omitempty" name:"ScanRemainTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryScanStatusOneKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryScanStatusOneKeyResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryScanStatusOneKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryScanStatusOneKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistrySummaryRequestParams struct {

}

type DescribeAssetImageRegistrySummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAssetImageRegistrySummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistrySummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistrySummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistrySummaryResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistrySummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistrySummaryResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistrySummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistrySummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryVirusListExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAssetImageRegistryVirusListExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVirusListExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryVirusListExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ImageInfo")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryVirusListExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryVirusListExportResponseParams struct {
	// excel文件下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryVirusListExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryVirusListExportResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryVirusListExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryVirusListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryVirusListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAssetImageRegistryVirusListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVirusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryVirusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ImageInfo")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryVirusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryVirusListResponseParams struct {
	// 镜像漏洞列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ImageVirus `json:"List,omitempty" name:"List"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryVirusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryVirusListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryVirusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryVulListExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAssetImageRegistryVulListExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVulListExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryVulListExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ImageInfo")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryVulListExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryVulListExportResponseParams struct {
	// excel文件下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryVulListExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryVulListExportResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryVulListExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryVulListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryVulListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeAssetImageRegistryVulListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 漏洞级别筛选，</li>
	// <li>Name - String - 是否必填：否 - 漏洞名称</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 镜像标识Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAssetImageRegistryVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ImageInfo")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRegistryVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRegistryVulListResponseParams struct {
	// 镜像漏洞列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ImageVul `json:"List,omitempty" name:"List"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRegistryVulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRegistryVulListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRegistryVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRegistryVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRiskListExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 风险级别 1,2,3,4，</li>
	// <li>Behavior - String - 是否必填：否 - 风险行为 1,2,3,4</li>
	// <li>Type - String - 是否必填：否 - 风险类型  1,2,</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetImageRiskListExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 风险级别 1,2,3,4，</li>
	// <li>Behavior - String - 是否必填：否 - 风险行为 1,2,3,4</li>
	// <li>Type - String - 是否必填：否 - 风险类型  1,2,</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetImageRiskListExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRiskListExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "ImageID")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRiskListExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRiskListExportResponseParams struct {
	// excel文件下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRiskListExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRiskListExportResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRiskListExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRiskListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRiskListRequestParams struct {
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 风险级别 1,2,3,4，</li>
	// <li>Behavior - String - 是否必填：否 - 风险行为 1,2,3,4</li>
	// <li>Type - String - 是否必填：否 - 风险类型  1,2,</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAssetImageRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 风险级别 1,2,3,4，</li>
	// <li>Behavior - String - 是否必填：否 - 风险行为 1,2,3,4</li>
	// <li>Type - String - 是否必填：否 - 风险类型  1,2,</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageRiskListResponseParams struct {
	// 镜像病毒列表
	List []*ImageRiskInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageRiskListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageScanSettingRequestParams struct {

}

type DescribeAssetImageScanSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAssetImageScanSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageScanSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageScanSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageScanSettingResponseParams struct {
	// 开关
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 扫描时刻(完整时间;后端按0时区解析时分秒)
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 扫描间隔
	ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`

	// 扫描木马
	ScanVirus *bool `json:"ScanVirus,omitempty" name:"ScanVirus"`

	// 扫描敏感信息
	ScanRisk *bool `json:"ScanRisk,omitempty" name:"ScanRisk"`

	// 扫描漏洞
	ScanVul *bool `json:"ScanVul,omitempty" name:"ScanVul"`

	// 扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 自定义扫描镜像
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageScanSettingResponseParams `json:"Response"`
}

func (r *DescribeAssetImageScanSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageScanStatusRequestParams struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

type DescribeAssetImageScanStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *DescribeAssetImageScanStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageScanStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageScanStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageScanStatusResponseParams struct {
	// 镜像个数
	ImageTotal *uint64 `json:"ImageTotal,omitempty" name:"ImageTotal"`

	// 扫描镜像个数
	ImageScanCnt *uint64 `json:"ImageScanCnt,omitempty" name:"ImageScanCnt"`

	// 扫描状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 扫描进度 ImageScanCnt/ImageTotal *100
	Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`

	// 安全个数
	SuccessCount *uint64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 风险个数
	RiskCount *uint64 `json:"RiskCount,omitempty" name:"RiskCount"`

	// 剩余扫描时间
	LeftSeconds *uint64 `json:"LeftSeconds,omitempty" name:"LeftSeconds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageScanStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageScanStatusResponseParams `json:"Response"`
}

func (r *DescribeAssetImageScanStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageScanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageScanTaskRequestParams struct {

}

type DescribeAssetImageScanTaskRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAssetImageScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageScanTaskResponseParams struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageScanTaskResponseParams `json:"Response"`
}

func (r *DescribeAssetImageScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageSimpleListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 镜像名、镜像id 称筛选，</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAssetImageSimpleListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 镜像名、镜像id 称筛选，</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageSimpleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageSimpleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageSimpleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageSimpleListResponseParams struct {
	// 镜像列表
	List []*AssetSimpleImageInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageSimpleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageSimpleListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageSimpleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageSimpleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageVirusListExportRequestParams struct {
	// 列表支持字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>RiskLevel - String - 是否必填：否 - 风险等级  1,2,3,4</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetImageVirusListExportRequest struct {
	*tchttp.BaseRequest
	
	// 列表支持字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>RiskLevel - String - 是否必填：否 - 风险等级  1,2,3,4</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetImageVirusListExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageVirusListExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "ImageID")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageVirusListExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageVirusListExportResponseParams struct {
	// excel文件下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageVirusListExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageVirusListExportResponseParams `json:"Response"`
}

func (r *DescribeAssetImageVirusListExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageVirusListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageVirusListRequestParams struct {
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>RiskLevel - String - 是否必填：否 - 风险等级  1,2,3,4</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序 asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAssetImageVirusListRequest struct {
	*tchttp.BaseRequest
	
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 镜像名称筛选，</li>
	// <li>RiskLevel - String - 是否必填：否 - 风险等级  1,2,3,4</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序 asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAssetImageVirusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageVirusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageVirusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageVirusListResponseParams struct {
	// 镜像病毒列表
	List []*ImageVirusInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 病毒扫描状态
	// 0:未扫描
	// 1:扫描中
	// 2:扫描完成
	// 3:扫描出错
	// 4:扫描取消
	VirusScanStatus *uint64 `json:"VirusScanStatus,omitempty" name:"VirusScanStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageVirusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageVirusListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageVirusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageVulListExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 漏洞名称名称筛选，</li>
	// <li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetImageVulListExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 漏洞名称名称筛选，</li>
	// <li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetImageVulListExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageVulListExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "ImageID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageVulListExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageVulListExportResponseParams struct {
	// excel文件下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageVulListExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageVulListExportResponseParams `json:"Response"`
}

func (r *DescribeAssetImageVulListExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageVulListExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageVulListRequestParams struct {
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 漏洞名称名称筛选，</li>
	// <li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段（Level）
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 + -
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeAssetImageVulListRequest struct {
	*tchttp.BaseRequest
	
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Name- String - 是否必填：否 - 漏洞名称名称筛选，</li>
	// <li>Level - String - 是否必填：否 - 风险等级  1,2,3,4</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段（Level）
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 + -
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeAssetImageVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetImageVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetImageVulListResponseParams struct {
	// 镜像漏洞列表
	List []*ImagesVul `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetImageVulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetImageVulListResponseParams `json:"Response"`
}

func (r *DescribeAssetImageVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetImageVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetPortListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>All - String - 是否必填：否 - 模糊查询可选字段</li>
	// <li>RunAs - String - 是否必填：否 - 运行用户筛选，</li>
	// <li>ContainerID - String - 是否必填：否 - 容器id</li>
	// <li>HostID- String - 是否必填：是 - 主机id</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>ProcessName- string - 是否必填：否 - 进程名搜索</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetPortListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>All - String - 是否必填：否 - 模糊查询可选字段</li>
	// <li>RunAs - String - 是否必填：否 - 运行用户筛选，</li>
	// <li>ContainerID - String - 是否必填：否 - 容器id</li>
	// <li>HostID- String - 是否必填：是 - 主机id</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>ProcessName- string - 是否必填：否 - 进程名搜索</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetPortListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetPortListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetPortListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetPortListResponseParams struct {
	// 端口列表
	List []*PortInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetPortListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetPortListResponseParams `json:"Response"`
}

func (r *DescribeAssetPortListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetPortListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetProcessListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>RunAs - String - 是否必填：否 - 运行用户筛选，</li>
	// <li>ContainerID - String - 是否必填：否 - 容器id</li>
	// <li>HostID- String - 是否必填：是 - 主机id</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>ProcessName- string - 是否必填：否 - 进程名搜索</li>
	// <li>Pid- string - 是否必填：否 - 进程id搜索(关联进程)</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetProcessListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>RunAs - String - 是否必填：否 - 运行用户筛选，</li>
	// <li>ContainerID - String - 是否必填：否 - 容器id</li>
	// <li>HostID- String - 是否必填：是 - 主机id</li>
	// <li>HostIP- string - 是否必填：否 - 主机ip搜索</li>
	// <li>ProcessName- string - 是否必填：否 - 进程名搜索</li>
	// <li>Pid- string - 是否必填：否 - 进程id搜索(关联进程)</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetProcessListResponseParams struct {
	// 端口列表
	List []*ProcessInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetProcessListResponseParams `json:"Response"`
}

func (r *DescribeAssetProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSummaryRequestParams struct {

}

type DescribeAssetSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAssetSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSummaryResponseParams struct {
	// 应用个数
	AppCnt *uint64 `json:"AppCnt,omitempty" name:"AppCnt"`

	// 容器个数
	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`

	// 暂停的容器个数
	ContainerPause *uint64 `json:"ContainerPause,omitempty" name:"ContainerPause"`

	// 运行的容器个数
	ContainerRunning *uint64 `json:"ContainerRunning,omitempty" name:"ContainerRunning"`

	// 停止运行的容器个数
	ContainerStop *uint64 `json:"ContainerStop,omitempty" name:"ContainerStop"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据库个数
	DbCnt *uint64 `json:"DbCnt,omitempty" name:"DbCnt"`

	// 镜像个数
	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`

	// 主机在线个数
	HostOnline *uint64 `json:"HostOnline,omitempty" name:"HostOnline"`

	// 主机个数
	HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`

	// 有风险的镜像个数
	ImageHasRiskInfoCnt *uint64 `json:"ImageHasRiskInfoCnt,omitempty" name:"ImageHasRiskInfoCnt"`

	// 有病毒的镜像个数
	ImageHasVirusCnt *uint64 `json:"ImageHasVirusCnt,omitempty" name:"ImageHasVirusCnt"`

	// 有漏洞的镜像个数
	ImageHasVulsCnt *uint64 `json:"ImageHasVulsCnt,omitempty" name:"ImageHasVulsCnt"`

	// 不受信任的镜像个数
	ImageUntrustCnt *uint64 `json:"ImageUntrustCnt,omitempty" name:"ImageUntrustCnt"`

	// 监听的端口个数
	ListenPortCnt *uint64 `json:"ListenPortCnt,omitempty" name:"ListenPortCnt"`

	// 进程个数
	ProcessCnt *uint64 `json:"ProcessCnt,omitempty" name:"ProcessCnt"`

	// web服务个数
	WebServiceCnt *uint64 `json:"WebServiceCnt,omitempty" name:"WebServiceCnt"`

	// 最近镜像扫描时间
	LatestImageScanTime *string `json:"LatestImageScanTime,omitempty" name:"LatestImageScanTime"`

	// 风险镜像个数
	ImageUnsafeCnt *uint64 `json:"ImageUnsafeCnt,omitempty" name:"ImageUnsafeCnt"`

	// 主机未安装agent数量
	HostUnInstallCnt *uint64 `json:"HostUnInstallCnt,omitempty" name:"HostUnInstallCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetSummaryResponseParams `json:"Response"`
}

func (r *DescribeAssetSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSyncLastTimeRequestParams struct {

}

type DescribeAssetSyncLastTimeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAssetSyncLastTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSyncLastTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetSyncLastTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSyncLastTimeResponseParams struct {
	// 资产最近同步时间
	AssetSyncLastTime *string `json:"AssetSyncLastTime,omitempty" name:"AssetSyncLastTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetSyncLastTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetSyncLastTimeResponseParams `json:"Response"`
}

func (r *DescribeAssetSyncLastTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSyncLastTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebServiceListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>
	// <li>Type- String - 是否必填：否 - 主机运行状态筛选，"Apache"
	// "Jboss"
	// "lighttpd"
	// "Nginx"
	// "Tomcat"</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAssetWebServiceListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords- String - 是否必填：否 - 模糊查询可选字段</li>
	// <li>Type- String - 是否必填：否 - 主机运行状态筛选，"Apache"
	// "Jboss"
	// "lighttpd"
	// "Nginx"
	// "Tomcat"</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAssetWebServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebServiceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetWebServiceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetWebServiceListResponseParams struct {
	// 主机列表
	List []*ServiceInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAssetWebServiceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetWebServiceListResponseParams `json:"Response"`
}

func (r *DescribeAssetWebServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetWebServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoAuthorizedRuleHostRequestParams struct {
	// 规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 需要返回的数量，默认为全部；
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序方式，asc，desc
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeAutoAuthorizedRuleHostRequest struct {
	*tchttp.BaseRequest
	
	// 规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 需要返回的数量，默认为全部；
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序方式，asc，desc
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeAutoAuthorizedRuleHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoAuthorizedRuleHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoAuthorizedRuleHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoAuthorizedRuleHostResponseParams struct {
	// 镜像自动授权规则授权范围主机列表
	List []*AutoAuthorizedRuleHostInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAutoAuthorizedRuleHostResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoAuthorizedRuleHostResponseParams `json:"Response"`
}

func (r *DescribeAutoAuthorizedRuleHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoAuthorizedRuleHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckItemListRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name 可取值：risk_level风险等级, risk_target检查对象，风险对象,risk_type风险类别,risk_attri检测项所属的风险类型
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCheckItemListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name 可取值：risk_level风险等级, risk_target检查对象，风险对象,risk_type风险类别,risk_attri检测项所属的风险类型
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCheckItemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckItemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCheckItemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckItemListResponseParams struct {
	// 检查项详情数组
	ClusterCheckItems []*ClusterCheckItem `json:"ClusterCheckItems,omitempty" name:"ClusterCheckItems"`

	// 检查项总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCheckItemListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCheckItemListResponseParams `json:"Response"`
}

func (r *DescribeCheckItemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterDetailResponseParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名字
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 当前集群扫描任务的进度，100表示扫描完成.
	ScanTaskProgress *int64 `json:"ScanTaskProgress,omitempty" name:"ScanTaskProgress"`

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 运行时组件
	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`

	// 集群节点数
	ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`

	// 集群状态 (Running 运行中 Creating 创建中 Abnormal 异常 )
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 集群类型：为托管集群MANAGED_CLUSTER、独立集群INDEPENDENT_CLUSTER
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 严重风险检查项的数量
	SeriousRiskCount *uint64 `json:"SeriousRiskCount,omitempty" name:"SeriousRiskCount"`

	// 高风险检查项的数量
	HighRiskCount *uint64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`

	// 中风险检查项的数量
	MiddleRiskCount *uint64 `json:"MiddleRiskCount,omitempty" name:"MiddleRiskCount"`

	// 提示风险检查项的数量
	HintRiskCount *uint64 `json:"HintRiskCount,omitempty" name:"HintRiskCount"`

	// 检查任务的状态
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 防御容器状态
	DefenderStatus *string `json:"DefenderStatus,omitempty" name:"DefenderStatus"`

	// 扫描任务创建时间
	TaskCreateTime *string `json:"TaskCreateTime,omitempty" name:"TaskCreateTime"`

	// 网络类型.PublicNetwork为公网类型,VPCNetwork为VPC网络
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// API Server地址
	ApiServerAddress *string `json:"ApiServerAddress,omitempty" name:"ApiServerAddress"`

	// 节点数
	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// 命名空间数
	NamespaceCount *uint64 `json:"NamespaceCount,omitempty" name:"NamespaceCount"`

	// 工作负载数
	WorkloadCount *uint64 `json:"WorkloadCount,omitempty" name:"WorkloadCount"`

	// Pod数量
	PodCount *uint64 `json:"PodCount,omitempty" name:"PodCount"`

	// Service数量
	ServiceCount *uint64 `json:"ServiceCount,omitempty" name:"ServiceCount"`

	// Ingress数量
	IngressCount *uint64 `json:"IngressCount,omitempty" name:"IngressCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDetailResponseParams `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterSummaryRequestParams struct {

}

type DescribeClusterSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeClusterSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterSummaryResponseParams struct {
	// 集群总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 有风险的集群数量
	RiskClusterCount *uint64 `json:"RiskClusterCount,omitempty" name:"RiskClusterCount"`

	// 未检查的集群数量
	UncheckClusterCount *uint64 `json:"UncheckClusterCount,omitempty" name:"UncheckClusterCount"`

	// 托管集群数量
	ManagedClusterCount *uint64 `json:"ManagedClusterCount,omitempty" name:"ManagedClusterCount"`

	// 独立集群数量
	IndependentClusterCount *uint64 `json:"IndependentClusterCount,omitempty" name:"IndependentClusterCount"`

	// 无风险的集群数量
	NoRiskClusterCount *uint64 `json:"NoRiskClusterCount,omitempty" name:"NoRiskClusterCount"`

	// 已经检查集群数
	CheckedClusterCount *uint64 `json:"CheckedClusterCount,omitempty" name:"CheckedClusterCount"`

	// 自动检查集群数
	AutoCheckClusterCount *uint64 `json:"AutoCheckClusterCount,omitempty" name:"AutoCheckClusterCount"`

	// 手动检查集群数
	ManualCheckClusterCount *uint64 `json:"ManualCheckClusterCount,omitempty" name:"ManualCheckClusterCount"`

	// 检查失败集群数
	FailedClusterCount *uint64 `json:"FailedClusterCount,omitempty" name:"FailedClusterCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterSummaryResponseParams `json:"Response"`
}

func (r *DescribeClusterSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceAssetDetailInfoRequestParams struct {
	// 客户资产ID。
	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
}

type DescribeComplianceAssetDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// 客户资产ID。
	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`
}

func (r *DescribeComplianceAssetDetailInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceAssetDetailInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerAssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceAssetDetailInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceAssetDetailInfoResponseParams struct {
	// 某资产的详情。
	AssetDetailInfo *ComplianceAssetDetailInfo `json:"AssetDetailInfo,omitempty" name:"AssetDetailInfo"`

	// 当资产为容器时，返回此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerDetailInfo *ComplianceContainerDetailInfo `json:"ContainerDetailInfo,omitempty" name:"ContainerDetailInfo"`

	// 当资产为镜像时，返回此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageDetailInfo *ComplianceImageDetailInfo `json:"ImageDetailInfo,omitempty" name:"ImageDetailInfo"`

	// 当资产为主机时，返回此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostDetailInfo *ComplianceHostDetailInfo `json:"HostDetailInfo,omitempty" name:"HostDetailInfo"`

	// 当资产为K8S时，返回此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	K8SDetailInfo *ComplianceK8SDetailInfo `json:"K8SDetailInfo,omitempty" name:"K8SDetailInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComplianceAssetDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComplianceAssetDetailInfoResponseParams `json:"Response"`
}

func (r *DescribeComplianceAssetDetailInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceAssetDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceAssetListRequestParams struct {
	// 资产类型列表。
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回的数据量，默认为10，最大为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询过滤器
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeComplianceAssetListRequest struct {
	*tchttp.BaseRequest
	
	// 资产类型列表。
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回的数据量，默认为10，最大为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询过滤器
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceAssetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceAssetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetTypeSet")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceAssetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceAssetListResponseParams struct {
	// 返回资产的总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回各类资产的列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetInfoList []*ComplianceAssetInfo `json:"AssetInfoList,omitempty" name:"AssetInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComplianceAssetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComplianceAssetListResponseParams `json:"Response"`
}

func (r *DescribeComplianceAssetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceAssetPolicyItemListRequestParams struct {
	// 客户资产的ID。
	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 要获取的数据量，默认为10，最大为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。Name字段支持
	// RiskLevel
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeComplianceAssetPolicyItemListRequest struct {
	*tchttp.BaseRequest
	
	// 客户资产的ID。
	CustomerAssetId *uint64 `json:"CustomerAssetId,omitempty" name:"CustomerAssetId"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 要获取的数据量，默认为10，最大为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤器列表。Name字段支持
	// RiskLevel
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceAssetPolicyItemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceAssetPolicyItemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerAssetId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceAssetPolicyItemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceAssetPolicyItemListResponseParams struct {
	// 返回检测项的总数。如果用户未启用基线检查，此处返回0。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回某个资产下的检测项的列表。
	AssetPolicyItemList []*ComplianceAssetPolicyItem `json:"AssetPolicyItemList,omitempty" name:"AssetPolicyItemList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComplianceAssetPolicyItemListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComplianceAssetPolicyItemListResponseParams `json:"Response"`
}

func (r *DescribeComplianceAssetPolicyItemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceAssetPolicyItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompliancePeriodTaskListRequestParams struct {
	// 资产的类型，取值为：
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeCompliancePeriodTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 资产的类型，取值为：
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCompliancePeriodTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompliancePeriodTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompliancePeriodTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompliancePeriodTaskListResponseParams struct {
	// 定时任务的总量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 定时任务信息的列表
	PeriodTaskSet []*CompliancePeriodTask `json:"PeriodTaskSet,omitempty" name:"PeriodTaskSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCompliancePeriodTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompliancePeriodTaskListResponseParams `json:"Response"`
}

func (r *DescribeCompliancePeriodTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompliancePeriodTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompliancePolicyItemAffectedAssetListRequestParams struct {
	// DescribeComplianceTaskPolicyItemSummaryList返回的CustomerPolicyItemId，表示检测项的ID。
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// Name - String
	// Name 可取值：NodeName, CheckResult
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCompliancePolicyItemAffectedAssetListRequest struct {
	*tchttp.BaseRequest
	
	// DescribeComplianceTaskPolicyItemSummaryList返回的CustomerPolicyItemId，表示检测项的ID。
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// Name - String
	// Name 可取值：NodeName, CheckResult
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCompliancePolicyItemAffectedAssetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompliancePolicyItemAffectedAssetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerPolicyItemId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompliancePolicyItemAffectedAssetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompliancePolicyItemAffectedAssetListResponseParams struct {
	// 返回各检测项所影响的资产的列表。
	AffectedAssetList []*ComplianceAffectedAsset `json:"AffectedAssetList,omitempty" name:"AffectedAssetList"`

	// 检测项影响的资产的总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCompliancePolicyItemAffectedAssetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompliancePolicyItemAffectedAssetListResponseParams `json:"Response"`
}

func (r *DescribeCompliancePolicyItemAffectedAssetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompliancePolicyItemAffectedAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompliancePolicyItemAffectedSummaryRequestParams struct {
	// DescribeComplianceTaskPolicyItemSummaryList返回的CustomerPolicyItemId，表示检测项的ID。
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
}

type DescribeCompliancePolicyItemAffectedSummaryRequest struct {
	*tchttp.BaseRequest
	
	// DescribeComplianceTaskPolicyItemSummaryList返回的CustomerPolicyItemId，表示检测项的ID。
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`
}

func (r *DescribeCompliancePolicyItemAffectedSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompliancePolicyItemAffectedSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerPolicyItemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompliancePolicyItemAffectedSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompliancePolicyItemAffectedSummaryResponseParams struct {
	// 返回各检测项影响的资产的汇总信息。
	PolicyItemSummary *CompliancePolicyItemSummary `json:"PolicyItemSummary,omitempty" name:"PolicyItemSummary"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCompliancePolicyItemAffectedSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompliancePolicyItemAffectedSummaryResponseParams `json:"Response"`
}

func (r *DescribeCompliancePolicyItemAffectedSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompliancePolicyItemAffectedSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceScanFailedAssetListRequestParams struct {
	// 资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回的数据量，默认为10，最大为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询过滤器
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeComplianceScanFailedAssetListRequest struct {
	*tchttp.BaseRequest
	
	// 资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回的数据量，默认为10，最大为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询过滤器
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceScanFailedAssetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceScanFailedAssetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetTypeSet")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceScanFailedAssetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceScanFailedAssetListResponseParams struct {
	// 返回检测失败的资产的总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回各类检测失败的资产的汇总信息的列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanFailedAssetList []*ComplianceScanFailedAsset `json:"ScanFailedAssetList,omitempty" name:"ScanFailedAssetList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComplianceScanFailedAssetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComplianceScanFailedAssetListResponseParams `json:"Response"`
}

func (r *DescribeComplianceScanFailedAssetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceScanFailedAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceTaskAssetSummaryRequestParams struct {
	// 资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
}

type DescribeComplianceTaskAssetSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 资产类型列表。
	// ASSET_CONTAINER, 容器
	// ASSET_IMAGE, 镜像
	// ASSET_HOST, 主机
	// ASSET_K8S, K8S资产
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`
}

func (r *DescribeComplianceTaskAssetSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceTaskAssetSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetTypeSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceTaskAssetSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceTaskAssetSummaryResponseParams struct {
	// 返回用户的状态，
	// 
	// USER_UNINIT: 用户未初始化。
	// USER_INITIALIZING，表示用户正在初始化环境。
	// USER_NORMAL: 正常状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 返回各类资产的汇总信息的列表。
	AssetSummaryList []*ComplianceAssetSummary `json:"AssetSummaryList,omitempty" name:"AssetSummaryList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComplianceTaskAssetSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComplianceTaskAssetSummaryResponseParams `json:"Response"`
}

func (r *DescribeComplianceTaskAssetSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceTaskAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceTaskPolicyItemSummaryListRequestParams struct {
	// 资产类型。仅查询与指定资产类型相关的检测项。
	// 
	// ASSET_CONTAINER, 容器
	// 
	// ASSET_IMAGE, 镜像
	// 
	// ASSET_HOST, 主机
	// 
	// ASSET_K8S, K8S资产
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// Name - String
	// Name 可取值：ItemType, StandardId,  RiskLevel。
	// 当为K8S资产时，还可取ClusterName。
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeComplianceTaskPolicyItemSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// 资产类型。仅查询与指定资产类型相关的检测项。
	// 
	// ASSET_CONTAINER, 容器
	// 
	// ASSET_IMAGE, 镜像
	// 
	// ASSET_HOST, 主机
	// 
	// ASSET_K8S, K8S资产
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// Name - String
	// Name 可取值：ItemType, StandardId,  RiskLevel。
	// 当为K8S资产时，还可取ClusterName。
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComplianceTaskPolicyItemSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceTaskPolicyItemSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceTaskPolicyItemSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceTaskPolicyItemSummaryListResponseParams struct {
	// 返回最近一次合规检查任务的ID。这个任务为本次所展示数据的来源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 返回检测项的总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回各检测项对应的汇总信息的列表。
	PolicyItemSummaryList []*CompliancePolicyItemSummary `json:"PolicyItemSummaryList,omitempty" name:"PolicyItemSummaryList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComplianceTaskPolicyItemSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComplianceTaskPolicyItemSummaryListResponseParams `json:"Response"`
}

func (r *DescribeComplianceTaskPolicyItemSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceTaskPolicyItemSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceWhitelistItemListRequestParams struct {
	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 要获取的数量，默认为10，最大为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资产类型列表。
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`

	// 查询过滤器
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 desc asc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeComplianceWhitelistItemListRequest struct {
	*tchttp.BaseRequest
	
	// 起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 要获取的数量，默认为10，最大为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资产类型列表。
	AssetTypeSet []*string `json:"AssetTypeSet,omitempty" name:"AssetTypeSet"`

	// 查询过滤器
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 desc asc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeComplianceWhitelistItemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceWhitelistItemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AssetTypeSet")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceWhitelistItemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceWhitelistItemListResponseParams struct {
	// 白名单项的列表。
	WhitelistItemSet []*ComplianceWhitelistItem `json:"WhitelistItemSet,omitempty" name:"WhitelistItemSet"`

	// 白名单项的总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeComplianceWhitelistItemListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComplianceWhitelistItemListResponseParams `json:"Response"`
}

func (r *DescribeComplianceWhitelistItemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceWhitelistItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerAssetSummaryRequestParams struct {

}

type DescribeContainerAssetSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeContainerAssetSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerAssetSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerAssetSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerAssetSummaryResponseParams struct {
	// 容器总数
	ContainerTotalCnt *uint64 `json:"ContainerTotalCnt,omitempty" name:"ContainerTotalCnt"`

	// 正在运行容器数量
	ContainerRunningCnt *uint64 `json:"ContainerRunningCnt,omitempty" name:"ContainerRunningCnt"`

	// 暂停运行容器数量
	ContainerPauseCnt *uint64 `json:"ContainerPauseCnt,omitempty" name:"ContainerPauseCnt"`

	// 停止运行容器数量
	ContainerStopped *uint64 `json:"ContainerStopped,omitempty" name:"ContainerStopped"`

	// 本地镜像数量
	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`

	// 主机节点数量
	HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`

	// 主机正在运行节点数量
	HostRunningCnt *uint64 `json:"HostRunningCnt,omitempty" name:"HostRunningCnt"`

	// 主机离线节点数量
	HostOfflineCnt *uint64 `json:"HostOfflineCnt,omitempty" name:"HostOfflineCnt"`

	// 镜像仓库数量
	ImageRegistryCnt *uint64 `json:"ImageRegistryCnt,omitempty" name:"ImageRegistryCnt"`

	// 镜像总数
	ImageTotalCnt *uint64 `json:"ImageTotalCnt,omitempty" name:"ImageTotalCnt"`

	// 主机未安装agent数量
	HostUnInstallCnt *uint64 `json:"HostUnInstallCnt,omitempty" name:"HostUnInstallCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeContainerAssetSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerAssetSummaryResponseParams `json:"Response"`
}

func (r *DescribeContainerAssetSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerAssetSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerSecEventSummaryRequestParams struct {

}

type DescribeContainerSecEventSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeContainerSecEventSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerSecEventSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerSecEventSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerSecEventSummaryResponseParams struct {
	// 未处理逃逸事件
	UnhandledEscapeCnt *uint64 `json:"UnhandledEscapeCnt,omitempty" name:"UnhandledEscapeCnt"`

	// 未处理反弹shell事件
	UnhandledReverseShellCnt *uint64 `json:"UnhandledReverseShellCnt,omitempty" name:"UnhandledReverseShellCnt"`

	// 未处理高危系统调用
	UnhandledRiskSyscallCnt *uint64 `json:"UnhandledRiskSyscallCnt,omitempty" name:"UnhandledRiskSyscallCnt"`

	// 未处理异常进程
	UnhandledAbnormalProcessCnt *uint64 `json:"UnhandledAbnormalProcessCnt,omitempty" name:"UnhandledAbnormalProcessCnt"`

	// 未处理文件篡改
	UnhandledFileCnt *uint64 `json:"UnhandledFileCnt,omitempty" name:"UnhandledFileCnt"`

	// 未处理木马事件
	UnhandledVirusEventCnt *uint64 `json:"UnhandledVirusEventCnt,omitempty" name:"UnhandledVirusEventCnt"`

	// 未处理恶意外连事件
	UnhandledMaliciousConnectionEventCnt *uint64 `json:"UnhandledMaliciousConnectionEventCnt,omitempty" name:"UnhandledMaliciousConnectionEventCnt"`

	// 未处理k8sApi事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnhandledK8sApiEventCnt *uint64 `json:"UnhandledK8sApiEventCnt,omitempty" name:"UnhandledK8sApiEventCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeContainerSecEventSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerSecEventSummaryResponseParams `json:"Response"`
}

func (r *DescribeContainerSecEventSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerSecEventSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeESAggregationsRequestParams struct {
	// ES聚合条件JSON
	Query *string `json:"Query,omitempty" name:"Query"`
}

type DescribeESAggregationsRequest struct {
	*tchttp.BaseRequest
	
	// ES聚合条件JSON
	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeESAggregationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeESAggregationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeESAggregationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeESAggregationsResponseParams struct {
	// ES聚合结果JSON
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeESAggregationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeESAggregationsResponseParams `json:"Response"`
}

func (r *DescribeESAggregationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeESAggregationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeESHitsRequestParams struct {
	// ES查询条件JSON
	Query *string `json:"Query,omitempty" name:"Query"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeESHitsRequest struct {
	*tchttp.BaseRequest
	
	// ES查询条件JSON
	Query *string `json:"Query,omitempty" name:"Query"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeESHitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeESHitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeESHitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeESHitsResponseParams struct {
	// ES查询结果JSON
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeESHitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeESHitsResponseParams `json:"Response"`
}

func (r *DescribeESHitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeESHitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmergencyVulListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeEmergencyVulListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEmergencyVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmergencyVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEmergencyVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmergencyVulListResponseParams struct {
	// 漏洞总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 漏洞列表
	List []*EmergencyVulInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEmergencyVulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEmergencyVulListResponseParams `json:"Response"`
}

func (r *DescribeEmergencyVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmergencyVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventDetailRequestParams struct {
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DescribeEscapeEventDetailRequest struct {
	*tchttp.BaseRequest
	
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeEscapeEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEscapeEventDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventDetailResponseParams struct {
	// 事件基本信息
	EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`

	// 进程信息
	ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`

	// 事件描述
	EventDetail *EscapeEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`

	// 父进程信息
	ParentProcessInfo *ProcessBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`

	// 祖先进程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEscapeEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEscapeEventDetailResponseParams `json:"Response"`
}

func (r *DescribeEscapeEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventInfoRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeEscapeEventInfoRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,Status：EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEscapeEventInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEscapeEventInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventInfoResponseParams struct {
	// 逃逸事件数组
	EventSet []*EscapeEventInfo `json:"EventSet,omitempty" name:"EventSet"`

	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEscapeEventInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEscapeEventInfoResponseParams `json:"Response"`
}

func (r *DescribeEscapeEventInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventTendencyRequestParams struct {
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

type DescribeEscapeEventTendencyRequest struct {
	*tchttp.BaseRequest
	
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeEscapeEventTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndTime")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEscapeEventTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventTendencyResponseParams struct {
	// 待处理逃逸事件趋势
	List []*EscapeEventTendencyInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEscapeEventTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEscapeEventTendencyResponseParams `json:"Response"`
}

func (r *DescribeEscapeEventTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventTypeSummaryRequestParams struct {

}

type DescribeEscapeEventTypeSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEscapeEventTypeSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventTypeSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEscapeEventTypeSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventTypeSummaryResponseParams struct {
	// 容器逃逸事件数
	ContainerEscapeEventCount *int64 `json:"ContainerEscapeEventCount,omitempty" name:"ContainerEscapeEventCount"`

	// 程序提权事件数
	ProcessPrivilegeEventCount *int64 `json:"ProcessPrivilegeEventCount,omitempty" name:"ProcessPrivilegeEventCount"`

	// 风险容器事件数
	RiskContainerEventCount *int64 `json:"RiskContainerEventCount,omitempty" name:"RiskContainerEventCount"`

	// 逃逸事件待处理数
	PendingEscapeEventCount *int64 `json:"PendingEscapeEventCount,omitempty" name:"PendingEscapeEventCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEscapeEventTypeSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEscapeEventTypeSummaryResponseParams `json:"Response"`
}

func (r *DescribeEscapeEventTypeSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventTypeSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventsExportRequestParams struct {
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeEscapeEventsExportRequest struct {
	*tchttp.BaseRequest
	
	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEscapeEventsExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventsExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEscapeEventsExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeEventsExportResponseParams struct {
	// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEscapeEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEscapeEventsExportResponseParams `json:"Response"`
}

func (r *DescribeEscapeEventsExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeRuleInfoRequestParams struct {

}

type DescribeEscapeRuleInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEscapeRuleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeRuleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEscapeRuleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeRuleInfoResponseParams struct {
	// 规则信息
	RuleSet []*EscapeRule `json:"RuleSet,omitempty" name:"RuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEscapeRuleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEscapeRuleInfoResponseParams `json:"Response"`
}

func (r *DescribeEscapeRuleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeSafeStateRequestParams struct {

}

type DescribeEscapeSafeStateRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEscapeSafeStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeSafeStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEscapeSafeStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeSafeStateResponseParams struct {
	// Unsafe：存在风险，Safe：暂无风险,UnKnown:未知风险
	IsSafe *string `json:"IsSafe,omitempty" name:"IsSafe"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEscapeSafeStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEscapeSafeStateResponseParams `json:"Response"`
}

func (r *DescribeEscapeSafeStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeSafeStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeWhiteListRequestParams struct {
	// 过滤条件。
	// <li>EventType- String - 是否必填：否 - 加白事件类型，ESCAPE_CGROUPS：利用cgroup机制逃逸，ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸， ESCAPE_DOCKER_API：访问Docker API接口逃逸，ESCAPE_VUL_OCCURRED：逃逸漏洞利用，MOUNT_SENSITIVE_PTAH：敏感路径挂载，PRIVILEGE_CONTAINER_START：特权容器， PRIVILEGE：程序提权逃逸</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：主机数量：HostCount，容器数量：ContainerCount，更新时间：UpdateTime
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeEscapeWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>EventType- String - 是否必填：否 - 加白事件类型，ESCAPE_CGROUPS：利用cgroup机制逃逸，ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸， ESCAPE_DOCKER_API：访问Docker API接口逃逸，ESCAPE_VUL_OCCURRED：逃逸漏洞利用，MOUNT_SENSITIVE_PTAH：敏感路径挂载，PRIVILEGE_CONTAINER_START：特权容器， PRIVILEGE：程序提权逃逸</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：主机数量：HostCount，容器数量：ContainerCount，更新时间：UpdateTime
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeEscapeWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEscapeWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEscapeWhiteListResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 逃逸白名单列表
	List []*EscapeWhiteListInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEscapeWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEscapeWhiteListResponseParams `json:"Response"`
}

func (r *DescribeEscapeWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEscapeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportJobDownloadURLRequestParams struct {
	// 任务ID
	JobID *string `json:"JobID,omitempty" name:"JobID"`
}

type DescribeExportJobDownloadURLRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	JobID *string `json:"JobID,omitempty" name:"JobID"`
}

func (r *DescribeExportJobDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportJobDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportJobDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportJobDownloadURLResponseParams struct {
	// 下载链接
	DownloadURL *string `json:"DownloadURL,omitempty" name:"DownloadURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExportJobDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExportJobDownloadURLResponseParams `json:"Response"`
}

func (r *DescribeExportJobDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportJobDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportJobManageListRequestParams struct {
	// 过滤条件。
	// <li>ExportStatus- string -是否必填: 否 - 导出状态 RUNNING: 导出中 SUCCESS:导出完成 FAILURE:失败
	// <li>ExportSource- string -是否必填: 否 - 导出来源 LocalImage: 本地镜像
	// </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	// InsertTime: 创建时间
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeExportJobManageListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>ExportStatus- string -是否必填: 否 - 导出状态 RUNNING: 导出中 SUCCESS:导出完成 FAILURE:失败
	// <li>ExportSource- string -是否必填: 否 - 导出来源 LocalImage: 本地镜像
	// </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	// InsertTime: 创建时间
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeExportJobManageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportJobManageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportJobManageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportJobManageListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务列表
	List []*ExportJobInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExportJobManageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExportJobManageListResponseParams `json:"Response"`
}

func (r *DescribeExportJobManageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportJobManageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportJobResultRequestParams struct {
	// CreateExportComplianceStatusListJob返回的JobId字段的值
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DescribeExportJobResultRequest struct {
	*tchttp.BaseRequest
	
	// CreateExportComplianceStatusListJob返回的JobId字段的值
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeExportJobResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportJobResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportJobResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExportJobResultResponseParams struct {
	// 导出的状态。取值为, SUCCESS:成功、FAILURE:失败，RUNNING: 进行中。
	ExportStatus *string `json:"ExportStatus,omitempty" name:"ExportStatus"`

	// 返回下载URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadURL *string `json:"DownloadURL,omitempty" name:"DownloadURL"`

	// 当ExportStatus为RUNNING时，返回导出进度。0~100范围的浮点数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExportProgress *float64 `json:"ExportProgress,omitempty" name:"ExportProgress"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureMsg *string `json:"FailureMsg,omitempty" name:"FailureMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExportJobResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExportJobResultResponseParams `json:"Response"`
}

func (r *DescribeExportJobResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportJobResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAuthorizedInfoRequestParams struct {

}

type DescribeImageAuthorizedInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImageAuthorizedInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAuthorizedInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageAuthorizedInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAuthorizedInfoResponseParams struct {
	// 总共有效的镜像授权数
	TotalAuthorizedCnt *uint64 `json:"TotalAuthorizedCnt,omitempty" name:"TotalAuthorizedCnt"`

	// 已使用镜像授权数
	UsedAuthorizedCnt *uint64 `json:"UsedAuthorizedCnt,omitempty" name:"UsedAuthorizedCnt"`

	// 已开启扫描镜像数
	ScannedImageCnt *uint64 `json:"ScannedImageCnt,omitempty" name:"ScannedImageCnt"`

	// 未开启扫描镜像数
	NotScannedImageCnt *uint64 `json:"NotScannedImageCnt,omitempty" name:"NotScannedImageCnt"`

	// 本地未开启扫描镜像数
	NotScannedLocalImageCnt *uint64 `json:"NotScannedLocalImageCnt,omitempty" name:"NotScannedLocalImageCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageAuthorizedInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageAuthorizedInfoResponseParams `json:"Response"`
}

func (r *DescribeImageAuthorizedInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAuthorizedInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAutoAuthorizedLogListRequestParams struct {
	// 自动授权任务Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// Status授权结果，SUCCESS:成功，REACH_LIMIT:达到授权上限，LICENSE_INSUFFICIENT:授权数不足
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段：AuthorizedTime
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式，asc，desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeImageAutoAuthorizedLogListRequest struct {
	*tchttp.BaseRequest
	
	// 自动授权任务Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// Status授权结果，SUCCESS:成功，REACH_LIMIT:达到授权上限，LICENSE_INSUFFICIENT:授权数不足
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段：AuthorizedTime
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式，asc，desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeImageAutoAuthorizedLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAutoAuthorizedLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageAutoAuthorizedLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAutoAuthorizedLogListResponseParams struct {
	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 自动授权结果镜像列表
	List []*AutoAuthorizedImageInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageAutoAuthorizedLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageAutoAuthorizedLogListResponseParams `json:"Response"`
}

func (r *DescribeImageAutoAuthorizedLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAutoAuthorizedLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAutoAuthorizedRuleRequestParams struct {

}

type DescribeImageAutoAuthorizedRuleRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImageAutoAuthorizedRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAutoAuthorizedRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageAutoAuthorizedRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAutoAuthorizedRuleResponseParams struct {
	// 规则是否生效，0:不生效，1:已生效
	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 授权范围类别，MANUAL:自选主机节点，ALL:全部镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeType *string `json:"RangeType,omitempty" name:"RangeType"`

	// 授权范围是自选主机时的主机数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`

	// 每天最大的镜像授权数限制, 0表示无限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDailyCount *int64 `json:"MaxDailyCount,omitempty" name:"MaxDailyCount"`

	// 规则id，用未设置时为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageAutoAuthorizedRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageAutoAuthorizedRuleResponseParams `json:"Response"`
}

func (r *DescribeImageAutoAuthorizedRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAutoAuthorizedRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAutoAuthorizedTaskListRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤字段
	// Status授权结果，全部授权成功：ALLSUCCSESS，部分授权失败：PARTIALFAIL,全部授权失败：ALLFAIL
	// Type授权方式，AUTO:自动授权，MANUAL:手动授权
	// Source 镜像来源，LOCAL:本地镜像，REGISTRY:仓库镜像
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeImageAutoAuthorizedTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤字段
	// Status授权结果，全部授权成功：ALLSUCCSESS，部分授权失败：PARTIALFAIL,全部授权失败：ALLFAIL
	// Type授权方式，AUTO:自动授权，MANUAL:手动授权
	// Source 镜像来源，LOCAL:本地镜像，REGISTRY:仓库镜像
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeImageAutoAuthorizedTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAutoAuthorizedTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageAutoAuthorizedTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAutoAuthorizedTaskListResponseParams struct {
	// 自动授权任务列表
	List []*ImageAutoAuthorizedTask `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageAutoAuthorizedTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageAutoAuthorizedTaskListResponseParams `json:"Response"`
}

func (r *DescribeImageAutoAuthorizedTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAutoAuthorizedTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageComponentListRequestParams struct {
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式desc ，asc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeImageComponentListRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ComponentName- String - 是否必填：否 - 镜像组件名称</li><li>ComponentVersion- String - 是否必填：否 - 镜像组件版本</li><li>ComponentType- String - 是否必填：否 - 镜像组件类型</li><li>VulLevel- String - 是否必填：否 - 漏洞威胁等级</li><li>HasVul- String - 是否必填：否 -是否有漏洞；true：是，false，否；不传或ALL ：全部</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式desc ，asc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeImageComponentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageComponentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageComponentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageComponentListResponseParams struct {
	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像组件列表
	List []*ImageComponent `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageComponentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageComponentListResponseParams `json:"Response"`
}

func (r *DescribeImageComponentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageComponentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRegistryNamespaceListRequestParams struct {
	// 本次查询的起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 本次查询的数据量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询的过滤条件。Name字段可取值"Namespace"。
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeImageRegistryNamespaceListRequest struct {
	*tchttp.BaseRequest
	
	// 本次查询的起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 本次查询的数据量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询的过滤条件。Name字段可取值"Namespace"。
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeImageRegistryNamespaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRegistryNamespaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageRegistryNamespaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRegistryNamespaceListResponseParams struct {
	// 可返回的项目空间的总量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的项目空间列表
	NamespaceList []*string `json:"NamespaceList,omitempty" name:"NamespaceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageRegistryNamespaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageRegistryNamespaceListResponseParams `json:"Response"`
}

func (r *DescribeImageRegistryNamespaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRegistryNamespaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRegistryTimingScanTaskRequestParams struct {

}

type DescribeImageRegistryTimingScanTaskRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImageRegistryTimingScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRegistryTimingScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageRegistryTimingScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRegistryTimingScanTaskResponseParams struct {
	// 定时扫描开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 定时任务扫描时间
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 定时扫描间隔
	ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`

	// 扫描类型数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`

	// 扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 自定义扫描镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 自动以扫描镜像Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id []*uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageRegistryTimingScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageRegistryTimingScanTaskResponseParams `json:"Response"`
}

func (r *DescribeImageRegistryTimingScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRegistryTimingScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRiskSummaryRequestParams struct {

}

type DescribeImageRiskSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImageRiskSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRiskSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageRiskSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRiskSummaryResponseParams struct {
	// 安全漏洞
	VulnerabilityCnt []*RunTimeRiskInfo `json:"VulnerabilityCnt,omitempty" name:"VulnerabilityCnt"`

	// 木马病毒
	MalwareVirusCnt []*RunTimeRiskInfo `json:"MalwareVirusCnt,omitempty" name:"MalwareVirusCnt"`

	// 敏感信息
	RiskCnt []*RunTimeRiskInfo `json:"RiskCnt,omitempty" name:"RiskCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageRiskSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageRiskSummaryResponseParams `json:"Response"`
}

func (r *DescribeImageRiskSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRiskSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRiskTendencyRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeImageRiskTendencyRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeImageRiskTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRiskTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageRiskTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRiskTendencyResponseParams struct {
	// 本地镜像新增风险趋势信息列表
	ImageRiskTendencySet []*ImageRiskTendencyInfo `json:"ImageRiskTendencySet,omitempty" name:"ImageRiskTendencySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageRiskTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageRiskTendencyResponseParams `json:"Response"`
}

func (r *DescribeImageRiskTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRiskTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSimpleListRequestParams struct {
	// IsAuthorized 是否已经授权, 0:否 1:是 无:全部
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeImageSimpleListRequest struct {
	*tchttp.BaseRequest
	
	// IsAuthorized 是否已经授权, 0:否 1:是 无:全部
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeImageSimpleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSimpleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageSimpleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageSimpleListResponseParams struct {
	// 镜像列表
	ImageList []*ImageSimpleInfo `json:"ImageList,omitempty" name:"ImageList"`

	// 镜像数
	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageSimpleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageSimpleListResponseParams `json:"Response"`
}

func (r *DescribeImageSimpleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageSimpleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexListRequestParams struct {

}

type DescribeIndexListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIndexListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexListResponseParams struct {
	// ES 索引信息
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIndexListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexListResponseParams `json:"Response"`
}

func (r *DescribeIndexListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInspectionReportRequestParams struct {

}

type DescribeInspectionReportRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInspectionReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInspectionReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInspectionReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInspectionReportResponseParams struct {
	// 报告名称
	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`

	// 下载链接
	ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInspectionReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInspectionReportResponseParams `json:"Response"`
}

func (r *DescribeInspectionReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInspectionReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalEventInfoRequestParams struct {
	// 事件ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

type DescribeK8sApiAbnormalEventInfoRequest struct {
	*tchttp.BaseRequest
	
	// 事件ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeK8sApiAbnormalEventInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalEventInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeK8sApiAbnormalEventInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalEventInfoResponseParams struct {
	// 事件详情
	Info *K8sApiAbnormalEventInfo `json:"Info,omitempty" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeK8sApiAbnormalEventInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeK8sApiAbnormalEventInfoResponseParams `json:"Response"`
}

func (r *DescribeK8sApiAbnormalEventInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalEventInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalEventListRequestParams struct {
	// 过滤条件。
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>MatchRules - string  - 是否必填: 否 -命中规则筛选</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	// <li>Status - string  - 是否必填: 否 -事件状态筛选</li>
	// <li>MatchRuleType - string  - 是否必填: 否 -命中规则类型筛选</li>
	// <li>ClusterRunningStatus - string  - 是否必填: 否 -集群运行状态</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	// LatestFoundTime: 最近生成时间
	// AlarmCount: 告警数量
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeK8sApiAbnormalEventListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>MatchRules - string  - 是否必填: 否 -命中规则筛选</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	// <li>Status - string  - 是否必填: 否 -事件状态筛选</li>
	// <li>MatchRuleType - string  - 是否必填: 否 -命中规则类型筛选</li>
	// <li>ClusterRunningStatus - string  - 是否必填: 否 -集群运行状态</li>
	// <li>ClusterName - string  - 是否必填: 否 -集群名称</li>
	// <li>ClusterID - string  - 是否必填: 否 -集群ID</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	// LatestFoundTime: 最近生成时间
	// AlarmCount: 告警数量
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeK8sApiAbnormalEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeK8sApiAbnormalEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalEventListResponseParams struct {
	// 事件列表
	List []*K8sApiAbnormalEventListItem `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeK8sApiAbnormalEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeK8sApiAbnormalEventListResponseParams `json:"Response"`
}

func (r *DescribeK8sApiAbnormalEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalRuleInfoRequestParams struct {
	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

type DescribeK8sApiAbnormalRuleInfoRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *DescribeK8sApiAbnormalRuleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalRuleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeK8sApiAbnormalRuleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalRuleInfoResponseParams struct {
	// 规则详情
	Info *K8sApiAbnormalRuleInfo `json:"Info,omitempty" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeK8sApiAbnormalRuleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeK8sApiAbnormalRuleInfoResponseParams `json:"Response"`
}

func (r *DescribeK8sApiAbnormalRuleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalRuleListRequestParams struct {
	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段。
	// <li>UpdateTime - string  - 是否必填: 否 -最后更新时间</li>
	// <li>EffectClusterCount - string  - 是否必填: 否 -影响集群数</li>
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeK8sApiAbnormalRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>RuleType - string  - 是否必填: 否 -规则类型</li>
	// <li>Status - string  - 是否必填: 否 -状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段。
	// <li>UpdateTime - string  - 是否必填: 否 -最后更新时间</li>
	// <li>EffectClusterCount - string  - 是否必填: 否 -影响集群数</li>
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeK8sApiAbnormalRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeK8sApiAbnormalRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalRuleListResponseParams struct {
	// 规则列表
	List []*K8sApiAbnormalRuleListItem `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeK8sApiAbnormalRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeK8sApiAbnormalRuleListResponseParams `json:"Response"`
}

func (r *DescribeK8sApiAbnormalRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalRuleScopeListRequestParams struct {
	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// <li>Action - string -是否必填: 否 - 执行动作</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeK8sApiAbnormalRuleScopeListRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// <li>Action - string -是否必填: 否 - 执行动作</li>
	// <li>RiskLevel - string  - 是否必填: 否 -威胁等级筛选</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeK8sApiAbnormalRuleScopeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalRuleScopeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeK8sApiAbnormalRuleScopeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalRuleScopeListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表
	List []*K8sApiAbnormalRuleScopeInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeK8sApiAbnormalRuleScopeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeK8sApiAbnormalRuleScopeListResponseParams `json:"Response"`
}

func (r *DescribeK8sApiAbnormalRuleScopeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalRuleScopeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalSummaryRequestParams struct {

}

type DescribeK8sApiAbnormalSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeK8sApiAbnormalSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeK8sApiAbnormalSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalSummaryResponseParams struct {
	// 待处理事件个数
	UnhandleEventCount *uint64 `json:"UnhandleEventCount,omitempty" name:"UnhandleEventCount"`

	// 待处理高危事件个数
	UnhandleHighLevelEventCount *uint64 `json:"UnhandleHighLevelEventCount,omitempty" name:"UnhandleHighLevelEventCount"`

	// 待处理中危事件个数
	UnhandleMediumLevelEventCount *uint64 `json:"UnhandleMediumLevelEventCount,omitempty" name:"UnhandleMediumLevelEventCount"`

	// 待处理低危事件个数
	UnhandleLowLevelEventCount *uint64 `json:"UnhandleLowLevelEventCount,omitempty" name:"UnhandleLowLevelEventCount"`

	// 待处理提示级别事件个数
	UnhandleNoticeLevelEventCount *uint64 `json:"UnhandleNoticeLevelEventCount,omitempty" name:"UnhandleNoticeLevelEventCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeK8sApiAbnormalSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeK8sApiAbnormalSummaryResponseParams `json:"Response"`
}

func (r *DescribeK8sApiAbnormalSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalTendencyRequestParams struct {
	// 趋势周期(默认为7天)
	TendencyPeriod *uint64 `json:"TendencyPeriod,omitempty" name:"TendencyPeriod"`
}

type DescribeK8sApiAbnormalTendencyRequest struct {
	*tchttp.BaseRequest
	
	// 趋势周期(默认为7天)
	TendencyPeriod *uint64 `json:"TendencyPeriod,omitempty" name:"TendencyPeriod"`
}

func (r *DescribeK8sApiAbnormalTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TendencyPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeK8sApiAbnormalTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeK8sApiAbnormalTendencyResponseParams struct {
	// 趋势列表
	List []*K8sApiAbnormalTendencyItem `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeK8sApiAbnormalTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeK8sApiAbnormalTendencyResponseParams `json:"Response"`
}

func (r *DescribeK8sApiAbnormalTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeK8sApiAbnormalTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageStatisticRequestParams struct {

}

type DescribeLogStorageStatisticRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLogStorageStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogStorageStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageStatisticResponseParams struct {
	// 总容量（单位：B）
	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`

	// 已使用容量（单位：B）
	UsedSize *uint64 `json:"UsedSize,omitempty" name:"UsedSize"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogStorageStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogStorageStatisticResponseParams `json:"Response"`
}

func (r *DescribeLogStorageStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallAuditRecordRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - Action
	// Name 可取值：publish，unpublish，confirm，add，update，delete
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeNetworkFirewallAuditRecordRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - Action
	// Name 可取值：publish，unpublish，confirm，add，update，delete
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkFirewallAuditRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallAuditRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallAuditRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallAuditRecordResponseParams struct {
	// 集群审计总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群的审计详细信息
	AuditList []*NetworkAuditRecord `json:"AuditList,omitempty" name:"AuditList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallAuditRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallAuditRecordResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallAuditRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallAuditRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallClusterListRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeNetworkFirewallClusterListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkFirewallClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallClusterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallClusterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallClusterListResponseParams struct {
	// 集群总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群的详细信息
	ClusterInfoList []*NetworkClusterInfoItem `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallClusterListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallClusterListResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallClusterRefreshStatusRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeNetworkFirewallClusterRefreshStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeNetworkFirewallClusterRefreshStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallClusterRefreshStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallClusterRefreshStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallClusterRefreshStatusResponseParams struct {
	// 任务状态，可能为：Task_Running,Task_Succ,Task_Error,Task_NoExist
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallClusterRefreshStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallClusterRefreshStatusResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallClusterRefreshStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallClusterRefreshStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallNamespaceLabelListRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeNetworkFirewallNamespaceLabelListRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkFirewallNamespaceLabelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallNamespaceLabelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallNamespaceLabelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallNamespaceLabelListResponseParams struct {
	// 集群总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群空间标签详细信息
	ClusterNamespaceLabelList []*NetworkClusterNamespaceLabelInfo `json:"ClusterNamespaceLabelList,omitempty" name:"ClusterNamespaceLabelList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallNamespaceLabelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallNamespaceLabelListResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallNamespaceLabelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallNamespaceLabelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallNamespaceListRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeNetworkFirewallNamespaceListRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkFirewallNamespaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallNamespaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallNamespaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallNamespaceListResponseParams struct {
	// 集群总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群的详细信息
	ClusterNamespaceList []*NetworkClusterNamespaceInfo `json:"ClusterNamespaceList,omitempty" name:"ClusterNamespaceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallNamespaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallNamespaceListResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallNamespaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallNamespaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPodLabelsListRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeNetworkFirewallPodLabelsListRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkFirewallPodLabelsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPodLabelsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallPodLabelsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPodLabelsListResponseParams struct {
	// 集群pod总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群pod详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodList []*NetworkClusterPodInfo `json:"PodList,omitempty" name:"PodList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallPodLabelsListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallPodLabelsListResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallPodLabelsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPodLabelsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyDetailRequestParams struct {
	// 策略Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeNetworkFirewallPolicyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallPolicyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyDetailResponseParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 入站类型
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 出站类型
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// 自定义规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`

	// pod选择器
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 策略创建时间
	PolicyCreateTime *string `json:"PolicyCreateTime,omitempty" name:"PolicyCreateTime"`

	// 策略源类型，分为System和Manual，分别代表手动和系统添加
	PolicySourceType *string `json:"PolicySourceType,omitempty" name:"PolicySourceType"`

	// 网络策略对应的网络插件
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkPolicyPlugin *string `json:"NetworkPolicyPlugin,omitempty" name:"NetworkPolicyPlugin"`

	// 网络策略状态
	PublishStatus *string `json:"PublishStatus,omitempty" name:"PublishStatus"`

	// 网络发布结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishResult *string `json:"PublishResult,omitempty" name:"PublishResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallPolicyDetailResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyDiscoverRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeNetworkFirewallPolicyDiscoverRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeNetworkFirewallPolicyDiscoverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyDiscoverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallPolicyDiscoverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyDiscoverResponseParams struct {
	// 任务状态，可能为：Task_Running,Task_Succ,Task_Error,Task_NoExist
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallPolicyDiscoverResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallPolicyDiscoverResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyDiscoverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyDiscoverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyListRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeNetworkFirewallPolicyListRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeNetworkFirewallPolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallPolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyListResponseParams struct {
	// 集群总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群的详细信息
	NetPolicy []*NetworkPolicyInfoItem `json:"NetPolicy,omitempty" name:"NetPolicy"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallPolicyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallPolicyListResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyStatusRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeNetworkFirewallPolicyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeNetworkFirewallPolicyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallPolicyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyStatusResponseParams struct {
	// 任务状态，可能为：Task_Running,Task_Succ,Task_Error,Task_NoExist
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// NameRepeat,K8sRuleIngressPortError等
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResult []*string `json:"TaskResult,omitempty" name:"TaskResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallPolicyStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallPolicyStatusResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyYamlDetailRequestParams struct {
	// 策略Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkFirewallPolicyYamlDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkFirewallPolicyYamlDetailResponseParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// base64编码的yaml字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 策略创建时间
	PolicyCreateTime *string `json:"PolicyCreateTime,omitempty" name:"PolicyCreateTime"`

	// 策略源类型，分为System和Manual，分别代表手动和系统添加
	PolicySourceType *string `json:"PolicySourceType,omitempty" name:"PolicySourceType"`

	// 网络策略对应的网络插件
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkPolicyPlugin *string `json:"NetworkPolicyPlugin,omitempty" name:"NetworkPolicyPlugin"`

	// 网络策略状态
	PublishStatus *string `json:"PublishStatus,omitempty" name:"PublishStatus"`

	// 网络发布结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishResult *string `json:"PublishResult,omitempty" name:"PublishResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkFirewallPolicyYamlDetailResponseParams `json:"Response"`
}

func (r *DescribeNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewestVulRequestParams struct {

}

type DescribeNewestVulRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNewestVulRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewestVulRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewestVulRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewestVulResponseParams struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 漏洞名称
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 披露时间
	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`

	// 应急漏洞风险情况：NOT_SCAN：未扫描，SCANNING：扫描中，SCANNED：已扫描
	Status *string `json:"Status,omitempty" name:"Status"`

	// 漏洞CVEID
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNewestVulResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewestVulResponseParams `json:"Response"`
}

func (r *DescribeNewestVulResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewestVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostPayDetailRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribePostPayDetailRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePostPayDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostPayDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePostPayDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostPayDetailResponseParams struct {
	// 弹性计费扣费详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftQuotaDayDetail []*SoftQuotaDayInfo `json:"SoftQuotaDayDetail,omitempty" name:"SoftQuotaDayDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePostPayDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePostPayDetailResponseParams `json:"Response"`
}

func (r *DescribePostPayDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostPayDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionInfoRequestParams struct {

}

type DescribeProVersionInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProVersionInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProVersionInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProVersionInfoResponseParams struct {
	// 专业版开始时间，补充购买时才不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 专业版结束时间，补充购买时才不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 需购买的机器核数
	CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`

	// 弹性计费上限
	MaxPostPayCoresCnt *uint64 `json:"MaxPostPayCoresCnt,omitempty" name:"MaxPostPayCoresCnt"`

	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 购买状态
	// 待购: Pending
	// 已购: Normal
	// 隔离: Isolate
	BuyStatus *string `json:"BuyStatus,omitempty" name:"BuyStatus"`

	// 是否曾经购买过(false:未曾 true:曾经购买过)
	IsPurchased *bool `json:"IsPurchased,omitempty" name:"IsPurchased"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProVersionInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProVersionInfoResponseParams `json:"Response"`
}

func (r *DescribeProVersionInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProVersionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePromotionActivityRequestParams struct {
	// 活动ID
	ActiveID *uint64 `json:"ActiveID,omitempty" name:"ActiveID"`
}

type DescribePromotionActivityRequest struct {
	*tchttp.BaseRequest
	
	// 活动ID
	ActiveID *uint64 `json:"ActiveID,omitempty" name:"ActiveID"`
}

func (r *DescribePromotionActivityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePromotionActivityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActiveID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePromotionActivityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePromotionActivityResponseParams struct {
	// 促销活动内容
	List []*PromotionActivityContent `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePromotionActivityResponse struct {
	*tchttp.BaseResponse
	Response *DescribePromotionActivityResponseParams `json:"Response"`
}

func (r *DescribePromotionActivityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePromotionActivityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicKeyRequestParams struct {

}

type DescribePublicKeyRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicKeyResponseParams struct {
	// 公钥
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicKeyResponseParams `json:"Response"`
}

func (r *DescribePublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurchaseStateInfoRequestParams struct {

}

type DescribePurchaseStateInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePurchaseStateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurchaseStateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurchaseStateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurchaseStateInfoResponseParams struct {
	// 0：可申请试用可购买；1：只可购买(含试用审核不通过和试用过期)；2：试用生效中；3：专业版生效中；4：专业版过期
	State *int64 `json:"State,omitempty" name:"State"`

	// 总核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`

	// 已购买核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedCoresCnt *uint64 `json:"AuthorizedCoresCnt,omitempty" name:"AuthorizedCoresCnt"`

	// 镜像数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`

	// 已授权镜像数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedImageCnt *uint64 `json:"AuthorizedImageCnt,omitempty" name:"AuthorizedImageCnt"`

	// 已购买镜像授权数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PurchasedAuthorizedCnt *uint64 `json:"PurchasedAuthorizedCnt,omitempty" name:"PurchasedAuthorizedCnt"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

	// 0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费(用户设置)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutomaticRenewal *int64 `json:"AutomaticRenewal,omitempty" name:"AutomaticRenewal"`

	// 试用期间赠送镜像授权数，可能会过期
	// 注意：此字段可能返回 null，表示取不到有效值。
	GivenAuthorizedCnt *uint64 `json:"GivenAuthorizedCnt,omitempty" name:"GivenAuthorizedCnt"`

	// 起始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 子状态(具体意义依据State字段而定)
	// State为4时，有效值为: ISOLATE(隔离) DESTROED(已销毁)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubState *string `json:"SubState,omitempty" name:"SubState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePurchaseStateInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurchaseStateInfoResponseParams `json:"Response"`
}

func (r *DescribePurchaseStateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurchaseStateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRefreshTaskRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeRefreshTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeRefreshTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRefreshTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRefreshTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRefreshTaskResponseParams struct {
	// 刷新任务状态，可能为：Task_New,Task_Running,Task_Finish,Task_Error,Task_NoExist
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRefreshTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRefreshTaskResponseParams `json:"Response"`
}

func (r *DescribeRefreshTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRefreshTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellDetailRequestParams struct {
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DescribeReverseShellDetailRequest struct {
	*tchttp.BaseRequest
	
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeReverseShellDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellDetailResponseParams struct {
	// 事件基本信息
	EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`

	// 进程信息
	ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`

	// 父进程信息
	ParentProcessInfo *ProcessDetailBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`

	// 事件描述
	EventDetail *ReverseShellEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`

	// 祖先进程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReverseShellDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellDetailResponseParams `json:"Response"`
}

func (r *DescribeReverseShellDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventsExportRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type DescribeReverseShellEventsExportRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeReverseShellEventsExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventsExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellEventsExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventsExportResponseParams struct {
	// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReverseShellEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellEventsExportResponseParams `json:"Response"`
}

func (r *DescribeReverseShellEventsExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventsRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeReverseShellEventsRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellEventsResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 反弹shell数组
	EventSet []*ReverseShellEventInfo `json:"EventSet,omitempty" name:"EventSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellEventsResponseParams `json:"Response"`
}

func (r *DescribeReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellWhiteListDetailRequestParams struct {
	// 白名单id
	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
}

type DescribeReverseShellWhiteListDetailRequest struct {
	*tchttp.BaseRequest
	
	// 白名单id
	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
}

func (r *DescribeReverseShellWhiteListDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellWhiteListDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteListId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellWhiteListDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellWhiteListDetailResponseParams struct {
	// 事件基本信息
	WhiteListDetailInfo *ReverseShellWhiteListInfo `json:"WhiteListDetailInfo,omitempty" name:"WhiteListDetailInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReverseShellWhiteListDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellWhiteListDetailResponseParams `json:"Response"`
}

func (r *DescribeReverseShellWhiteListDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellWhiteListDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellWhiteListsRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeReverseShellWhiteListsRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeReverseShellWhiteListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellWhiteListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellWhiteListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReverseShellWhiteListsResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 白名单信息列表
	WhiteListSet []*ReverseShellWhiteListBaseInfo `json:"WhiteListSet,omitempty" name:"WhiteListSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReverseShellWhiteListsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReverseShellWhiteListsResponseParams `json:"Response"`
}

func (r *DescribeReverseShellWhiteListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskListRequestParams struct {
	// 要查询的集群ID，如果不指定，则查询用户所有的风险项
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：RiskLevel风险等级, RiskTarget检查对象，风险对象,RiskType风险类别,RiskAttribute检测项所属的风险类型,Name
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的集群ID，如果不指定，则查询用户所有的风险项
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：RiskLevel风险等级, RiskTarget检查对象，风险对象,RiskType风险类别,RiskAttribute检测项所属的风险类型,Name
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskListResponseParams struct {
	// 风险详情数组
	ClusterRiskItems []*ClusterRiskItem `json:"ClusterRiskItems,omitempty" name:"ClusterRiskItems"`

	// 风险项的总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallDetailRequestParams struct {
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

type DescribeRiskSyscallDetailRequest struct {
	*tchttp.BaseRequest
	
	// 事件唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeRiskSyscallDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskSyscallDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallDetailResponseParams struct {
	// 事件基本信息
	EventBaseInfo *RunTimeEventBaseInfo `json:"EventBaseInfo,omitempty" name:"EventBaseInfo"`

	// 进程信息
	ProcessInfo *ProcessDetailInfo `json:"ProcessInfo,omitempty" name:"ProcessInfo"`

	// 父进程信息
	ParentProcessInfo *ProcessDetailBaseInfo `json:"ParentProcessInfo,omitempty" name:"ParentProcessInfo"`

	// 事件描述
	EventDetail *RiskSyscallEventDescription `json:"EventDetail,omitempty" name:"EventDetail"`

	// 祖先进程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AncestorProcessInfo *ProcessBaseInfo `json:"AncestorProcessInfo,omitempty" name:"AncestorProcessInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRiskSyscallDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskSyscallDetailResponseParams `json:"Response"`
}

func (r *DescribeRiskSyscallDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallEventsExportRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type DescribeRiskSyscallEventsExportRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *DescribeRiskSyscallEventsExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallEventsExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskSyscallEventsExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallEventsExportResponseParams struct {
	// Excel下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRiskSyscallEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskSyscallEventsExportResponseParams `json:"Response"`
}

func (r *DescribeRiskSyscallEventsExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallEventsExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallEventsRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeRiskSyscallEventsRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeRiskSyscallEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskSyscallEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallEventsResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 高危系统调用数组
	EventSet []*RiskSyscallEventInfo `json:"EventSet,omitempty" name:"EventSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRiskSyscallEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskSyscallEventsResponseParams `json:"Response"`
}

func (r *DescribeRiskSyscallEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallNamesRequestParams struct {

}

type DescribeRiskSyscallNamesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRiskSyscallNamesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallNamesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskSyscallNamesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallNamesResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 系统调用名称列表
	SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRiskSyscallNamesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskSyscallNamesResponseParams `json:"Response"`
}

func (r *DescribeRiskSyscallNamesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallWhiteListDetailRequestParams struct {
	// 白名单id
	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
}

type DescribeRiskSyscallWhiteListDetailRequest struct {
	*tchttp.BaseRequest
	
	// 白名单id
	WhiteListId *string `json:"WhiteListId,omitempty" name:"WhiteListId"`
}

func (r *DescribeRiskSyscallWhiteListDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallWhiteListDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteListId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskSyscallWhiteListDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallWhiteListDetailResponseParams struct {
	// 白名单基本信息
	WhiteListDetailInfo *RiskSyscallWhiteListInfo `json:"WhiteListDetailInfo,omitempty" name:"WhiteListDetailInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRiskSyscallWhiteListDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskSyscallWhiteListDetailResponseParams `json:"Response"`
}

func (r *DescribeRiskSyscallWhiteListDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallWhiteListDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallWhiteListsRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeRiskSyscallWhiteListsRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数,"Filters":[{"Name":"Status","Values":["2"]}]
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 升序降序,asc desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeRiskSyscallWhiteListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallWhiteListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskSyscallWhiteListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskSyscallWhiteListsResponseParams struct {
	// 事件总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 白名单信息列表
	WhiteListSet []*RiskSyscallWhiteListBaseInfo `json:"WhiteListSet,omitempty" name:"WhiteListSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRiskSyscallWhiteListsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskSyscallWhiteListsResponseParams `json:"Response"`
}

func (r *DescribeRiskSyscallWhiteListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskSyscallWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanIgnoreVulListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式:DESC,ACS
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段 UpdateTime
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeScanIgnoreVulListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式:DESC,ACS
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段 UpdateTime
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeScanIgnoreVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanIgnoreVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanIgnoreVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanIgnoreVulListResponseParams struct {
	// 总是
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 漏洞列表
	List []*ScanIgnoreVul `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeScanIgnoreVulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanIgnoreVulListResponseParams `json:"Response"`
}

func (r *DescribeScanIgnoreVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanIgnoreVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchExportListRequestParams struct {
	// ES查询条件JSON
	Query *string `json:"Query,omitempty" name:"Query"`
}

type DescribeSearchExportListRequest struct {
	*tchttp.BaseRequest
	
	// ES查询条件JSON
	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeSearchExportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchExportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchExportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchExportListResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSearchExportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchExportListResponseParams `json:"Response"`
}

func (r *DescribeSearchExportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchExportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchLogsRequestParams struct {

}

type DescribeSearchLogsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSearchLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchLogsResponseParams struct {
	// 历史搜索记录 保留最新的10条
	Data []*string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSearchLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchLogsResponseParams `json:"Response"`
}

func (r *DescribeSearchLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchTemplatesRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSearchTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSearchTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchTemplatesResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 模板列表
	List []*SearchTemplate `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSearchTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchTemplatesResponseParams `json:"Response"`
}

func (r *DescribeSearchTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecEventsTendencyRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeSecEventsTendencyRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSecEventsTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecEventsTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecEventsTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecEventsTendencyResponseParams struct {
	// 运行时安全事件趋势信息列表
	EventTendencySet []*SecTendencyEventInfo `json:"EventTendencySet,omitempty" name:"EventTendencySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecEventsTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecEventsTendencyResponseParams `json:"Response"`
}

func (r *DescribeSecEventsTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecEventsTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogAlertMsgRequestParams struct {
	// 告警类型
	// 日志储量告警: log_reserve_full
	// 日志存储时间告警: log_save_day_limit
	// kafka实例/公网域名不可用: kafka_instance_domain_unavailable
	// kafka 用户名密码错误: kafka_user_passwd_wrong
	// kafka后台报错字段: kafka_field_wrong
	Type []*string `json:"Type,omitempty" name:"Type"`
}

type DescribeSecLogAlertMsgRequest struct {
	*tchttp.BaseRequest
	
	// 告警类型
	// 日志储量告警: log_reserve_full
	// 日志存储时间告警: log_save_day_limit
	// kafka实例/公网域名不可用: kafka_instance_domain_unavailable
	// kafka 用户名密码错误: kafka_user_passwd_wrong
	// kafka后台报错字段: kafka_field_wrong
	Type []*string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSecLogAlertMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogAlertMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogAlertMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogAlertMsgResponseParams struct {
	// 告警消息队列
	List []*SecLogAlertMsgInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogAlertMsgResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogAlertMsgResponseParams `json:"Response"`
}

func (r *DescribeSecLogAlertMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogAlertMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogCleanSettingInfoRequestParams struct {

}

type DescribeSecLogCleanSettingInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecLogCleanSettingInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogCleanSettingInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogCleanSettingInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogCleanSettingInfoResponseParams struct {
	// 触发清理的储量底线
	ReservesLimit *uint64 `json:"ReservesLimit,omitempty" name:"ReservesLimit"`

	// 清理停止时的储量截至线
	ReservesDeadline *uint64 `json:"ReservesDeadline,omitempty" name:"ReservesDeadline"`

	// 触发清理的存储天数
	DayLimit *uint64 `json:"DayLimit,omitempty" name:"DayLimit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogCleanSettingInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogCleanSettingInfoResponseParams `json:"Response"`
}

func (r *DescribeSecLogCleanSettingInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogCleanSettingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogDeliveryClsOptionsRequestParams struct {
	// 地域
	ClsRegion *string `json:"ClsRegion,omitempty" name:"ClsRegion"`
}

type DescribeSecLogDeliveryClsOptionsRequest struct {
	*tchttp.BaseRequest
	
	// 地域
	ClsRegion *string `json:"ClsRegion,omitempty" name:"ClsRegion"`
}

func (r *DescribeSecLogDeliveryClsOptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogDeliveryClsOptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClsRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogDeliveryClsOptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogDeliveryClsOptionsResponseParams struct {
	// cls可选日志集合列表(仅当入参ClsRegion不为空时返回)
	LogSetList []*ClsLogsetInfo `json:"LogSetList,omitempty" name:"LogSetList"`

	// 可选地域列表(仅当入参ClsRegion为空时返回)
	RegionList []*RegionInfo `json:"RegionList,omitempty" name:"RegionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogDeliveryClsOptionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogDeliveryClsOptionsResponseParams `json:"Response"`
}

func (r *DescribeSecLogDeliveryClsOptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogDeliveryClsOptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogDeliveryClsSettingRequestParams struct {

}

type DescribeSecLogDeliveryClsSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecLogDeliveryClsSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogDeliveryClsSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogDeliveryClsSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogDeliveryClsSettingResponseParams struct {
	// 日志类型列表
	LogTypeList []*SecLogDeliveryClsSettingInfo `json:"LogTypeList,omitempty" name:"LogTypeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogDeliveryClsSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogDeliveryClsSettingResponseParams `json:"Response"`
}

func (r *DescribeSecLogDeliveryClsSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogDeliveryClsSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogDeliveryKafkaOptionsRequestParams struct {
	// 地域，若为空则返回所有可选地域
	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
}

type DescribeSecLogDeliveryKafkaOptionsRequest struct {
	*tchttp.BaseRequest
	
	// 地域，若为空则返回所有可选地域
	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
}

func (r *DescribeSecLogDeliveryKafkaOptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogDeliveryKafkaOptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegionID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogDeliveryKafkaOptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogDeliveryKafkaOptionsResponseParams struct {
	// 实例列表
	InstanceList []*CKafkaInstanceInfo `json:"InstanceList,omitempty" name:"InstanceList"`

	// 地域列表
	RegionList []*RegionInfo `json:"RegionList,omitempty" name:"RegionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogDeliveryKafkaOptionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogDeliveryKafkaOptionsResponseParams `json:"Response"`
}

func (r *DescribeSecLogDeliveryKafkaOptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogDeliveryKafkaOptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogDeliveryKafkaSettingRequestParams struct {

}

type DescribeSecLogDeliveryKafkaSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecLogDeliveryKafkaSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogDeliveryKafkaSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogDeliveryKafkaSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogDeliveryKafkaSettingResponseParams struct {
	// 消息队列实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 消息队列实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 日志类型队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTypeList []*SecLogDeliveryKafkaSettingInfo `json:"LogTypeList,omitempty" name:"LogTypeList"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitempty" name:"User"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogDeliveryKafkaSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogDeliveryKafkaSettingResponseParams `json:"Response"`
}

func (r *DescribeSecLogDeliveryKafkaSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogDeliveryKafkaSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogJoinObjectListRequestParams struct {
	// 日志类型
	// bash: "container_bash",
	// 启动: "container_launch",
	// "k8s": "k8s_api"
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 主机状态 </li>
	// <li>HostIP- String - 是否必填：否 - 主机内网IP </li>
	// <li>PublicIP- String - 是否必填：否 - 主机外网IP </li>
	// <li>HostName- String - 是否必填：否 - 主机名称 </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeSecLogJoinObjectListRequest struct {
	*tchttp.BaseRequest
	
	// 日志类型
	// bash: "container_bash",
	// 启动: "container_launch",
	// "k8s": "k8s_api"
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 主机状态 </li>
	// <li>HostIP- String - 是否必填：否 - 主机内网IP </li>
	// <li>PublicIP- String - 是否必填：否 - 主机外网IP </li>
	// <li>HostName- String - 是否必填：否 - 主机名称 </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeSecLogJoinObjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogJoinObjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogJoinObjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogJoinObjectListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 接入对象列表
	List []*SecLogJoinObjectInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogJoinObjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogJoinObjectListResponseParams `json:"Response"`
}

func (r *DescribeSecLogJoinObjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogJoinObjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogJoinTypeListRequestParams struct {

}

type DescribeSecLogJoinTypeListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecLogJoinTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogJoinTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogJoinTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogJoinTypeListResponseParams struct {
	// 接入日志列表
	List []*SecLogJoinInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogJoinTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogJoinTypeListResponseParams `json:"Response"`
}

func (r *DescribeSecLogJoinTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogJoinTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogKafkaUINRequestParams struct {

}

type DescribeSecLogKafkaUINRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecLogKafkaUINRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogKafkaUINRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogKafkaUINRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogKafkaUINResponseParams struct {
	// 目标UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstUIN *string `json:"DstUIN,omitempty" name:"DstUIN"`

	// 授权状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogKafkaUINResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogKafkaUINResponseParams `json:"Response"`
}

func (r *DescribeSecLogKafkaUINResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogKafkaUINResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogVasInfoRequestParams struct {

}

type DescribeSecLogVasInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecLogVasInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogVasInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecLogVasInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecLogVasInfoResponseParams struct {
	// 购买状态
	// 待购: Pending
	// 已购: Normal
	// 隔离: Isolate
	BuyStatus *string `json:"BuyStatus,omitempty" name:"BuyStatus"`

	// 存储时长(月)
	LogSaveMonth *int64 `json:"LogSaveMonth,omitempty" name:"LogSaveMonth"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 截止时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 存储容量(GB)
	LogCapacity *uint64 `json:"LogCapacity,omitempty" name:"LogCapacity"`

	// 资源ID
	ResourceID *string `json:"ResourceID,omitempty" name:"ResourceID"`

	// 是否曾经购买过(false:未曾 true:曾经购买过)
	IsPurchased *bool `json:"IsPurchased,omitempty" name:"IsPurchased"`

	// 试用存储容量(GB)
	TrialCapacity *uint64 `json:"TrialCapacity,omitempty" name:"TrialCapacity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecLogVasInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecLogVasInfoResponseParams `json:"Response"`
}

func (r *DescribeSecLogVasInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecLogVasInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportDefenceVulRequestParams struct {
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：披露时间：SubmitTime
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeSupportDefenceVulRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：披露时间：SubmitTime
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeSupportDefenceVulRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportDefenceVulRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportDefenceVulRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportDefenceVulResponseParams struct {
	// 支持防御的漏洞列表
	List []*SupportDefenceVul `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSupportDefenceVulResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportDefenceVulResponseParams `json:"Response"`
}

func (r *DescribeSupportDefenceVulResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportDefenceVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemVulListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeSystemVulListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeSystemVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSystemVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemVulListResponseParams struct {
	// 漏洞总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 漏洞列表
	List []*VulInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSystemVulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSystemVulListResponseParams `json:"Response"`
}

func (r *DescribeSystemVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultSummaryRequestParams struct {

}

type DescribeTaskResultSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTaskResultSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultSummaryResponseParams struct {
	// 严重风险影响的节点数量,返回7天数据
	SeriousRiskNodeCount []*uint64 `json:"SeriousRiskNodeCount,omitempty" name:"SeriousRiskNodeCount"`

	// 高风险影响的节点的数量,返回7天数据
	HighRiskNodeCount []*uint64 `json:"HighRiskNodeCount,omitempty" name:"HighRiskNodeCount"`

	// 中风险检查项的节点数量,返回7天数据
	MiddleRiskNodeCount []*uint64 `json:"MiddleRiskNodeCount,omitempty" name:"MiddleRiskNodeCount"`

	// 提示风险检查项的节点数量,返回7天数据
	HintRiskNodeCount []*uint64 `json:"HintRiskNodeCount,omitempty" name:"HintRiskNodeCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskResultSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResultSummaryResponseParams `json:"Response"`
}

func (r *DescribeTaskResultSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTcssSummaryRequestParams struct {

}

type DescribeTcssSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTcssSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTcssSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTcssSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTcssSummaryResponseParams struct {
	// 镜像总数
	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`

	// 已扫描镜像数
	ScannedImageCnt *uint64 `json:"ScannedImageCnt,omitempty" name:"ScannedImageCnt"`

	// 待扫描镜像个数
	UnScannedImageCnt *uint64 `json:"UnScannedImageCnt,omitempty" name:"UnScannedImageCnt"`

	// 本地镜像个数
	LocalImageCnt *uint64 `json:"LocalImageCnt,omitempty" name:"LocalImageCnt"`

	// 仓库镜像个数
	RepositoryImageCnt *uint64 `json:"RepositoryImageCnt,omitempty" name:"RepositoryImageCnt"`

	// 风险本地镜像个数
	RiskLocalImageCnt *uint64 `json:"RiskLocalImageCnt,omitempty" name:"RiskLocalImageCnt"`

	// 风险仓库镜像个数
	RiskRepositoryImageCnt *uint64 `json:"RiskRepositoryImageCnt,omitempty" name:"RiskRepositoryImageCnt"`

	// 容器个数
	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`

	// 风险容器个数
	RiskContainerCnt *uint64 `json:"RiskContainerCnt,omitempty" name:"RiskContainerCnt"`

	// 集群个数
	ClusterCnt *uint64 `json:"ClusterCnt,omitempty" name:"ClusterCnt"`

	// 风险集群个数
	RiskClusterCnt *uint64 `json:"RiskClusterCnt,omitempty" name:"RiskClusterCnt"`

	// 待扫描漏洞个数
	UnScannedVulCnt *uint64 `json:"UnScannedVulCnt,omitempty" name:"UnScannedVulCnt"`

	// 风险漏洞个数
	RiskVulCnt *uint64 `json:"RiskVulCnt,omitempty" name:"RiskVulCnt"`

	// 安全基线待扫描项个数
	UnScannedBaseLineCnt *uint64 `json:"UnScannedBaseLineCnt,omitempty" name:"UnScannedBaseLineCnt"`

	// 安全基线风险个数
	RiskBaseLineCnt *uint64 `json:"RiskBaseLineCnt,omitempty" name:"RiskBaseLineCnt"`

	// 运行时(高危)待处理事件个数
	RuntimeUnhandleEventCnt *uint64 `json:"RuntimeUnhandleEventCnt,omitempty" name:"RuntimeUnhandleEventCnt"`

	// 待扫描集群个数
	UnScannedClusterCnt *uint64 `json:"UnScannedClusterCnt,omitempty" name:"UnScannedClusterCnt"`

	// 是否已扫描镜像
	ScanImageStatus *bool `json:"ScanImageStatus,omitempty" name:"ScanImageStatus"`

	// 是否已扫描集群
	ScanClusterStatus *bool `json:"ScanClusterStatus,omitempty" name:"ScanClusterStatus"`

	// 是否已扫描基线
	ScanBaseLineStatus *bool `json:"ScanBaseLineStatus,omitempty" name:"ScanBaseLineStatus"`

	// 是否已扫描漏洞
	ScanVulStatus *bool `json:"ScanVulStatus,omitempty" name:"ScanVulStatus"`

	// 漏洞影响镜像数
	VulRiskImageCnt *uint64 `json:"VulRiskImageCnt,omitempty" name:"VulRiskImageCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTcssSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTcssSummaryResponseParams `json:"Response"`
}

func (r *DescribeTcssSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTcssSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnauthorizedCoresTendencyRequestParams struct {

}

type DescribeUnauthorizedCoresTendencyRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUnauthorizedCoresTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnauthorizedCoresTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnauthorizedCoresTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnauthorizedCoresTendencyResponseParams struct {
	// 未授权核数趋势
	List []*UnauthorizedCoresTendency `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnauthorizedCoresTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnauthorizedCoresTendencyResponseParams `json:"Response"`
}

func (r *DescribeUnauthorizedCoresTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnauthorizedCoresTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnfinishRefreshTaskRequestParams struct {

}

type DescribeUnfinishRefreshTaskRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUnfinishRefreshTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnfinishRefreshTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnfinishRefreshTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnfinishRefreshTaskResponseParams struct {
	// 返回最近的一次任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态，为Task_New,Task_Running,Task_Finish,Task_Error,Task_NoExist.Task_New,Task_Running表示有任务存在，不允许新下发
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnfinishRefreshTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnfinishRefreshTaskResponseParams `json:"Response"`
}

func (r *DescribeUnfinishRefreshTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnfinishRefreshTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserClusterRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeUserClusterRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每次查询的最大记录数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Name - String
	// Name 可取值：ClusterName,ClusterId,ClusterType,Region,ClusterCheckMode,ClusterAutoCheck
	Filters []*ComplianceFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式 asc,desc
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeUserClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserClusterResponseParams struct {
	// 集群总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群的详细信息
	ClusterInfoList []*ClusterInfoItem `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserClusterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserClusterResponseParams `json:"Response"`
}

func (r *DescribeUserClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeValueAddedSrvInfoRequestParams struct {

}

type DescribeValueAddedSrvInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeValueAddedSrvInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeValueAddedSrvInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeValueAddedSrvInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeValueAddedSrvInfoResponseParams struct {
	// 仓库镜像未授权数量
	RegistryImageCnt *uint64 `json:"RegistryImageCnt,omitempty" name:"RegistryImageCnt"`

	// 本地镜像未授权数量
	LocalImageCnt *uint64 `json:"LocalImageCnt,omitempty" name:"LocalImageCnt"`

	// 未使用的镜像安全扫描授权数
	UnusedAuthorizedCnt *uint64 `json:"UnusedAuthorizedCnt,omitempty" name:"UnusedAuthorizedCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeValueAddedSrvInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeValueAddedSrvInfoResponseParams `json:"Response"`
}

func (r *DescribeValueAddedSrvInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeValueAddedSrvInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusAutoIsolateSampleDetailRequestParams struct {
	// 文件MD5值
	MD5 *string `json:"MD5,omitempty" name:"MD5"`
}

type DescribeVirusAutoIsolateSampleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 文件MD5值
	MD5 *string `json:"MD5,omitempty" name:"MD5"`
}

func (r *DescribeVirusAutoIsolateSampleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusAutoIsolateSampleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MD5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusAutoIsolateSampleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusAutoIsolateSampleDetailResponseParams struct {
	// 文件Md5值
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// 文件大小(B)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 病毒名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 查杀引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	KillEngine []*string `json:"KillEngine,omitempty" name:"KillEngine"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 事件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`

	// 建议方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`

	// 参考链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReferenceLink *string `json:"ReferenceLink,omitempty" name:"ReferenceLink"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusAutoIsolateSampleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusAutoIsolateSampleDetailResponseParams `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSampleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusAutoIsolateSampleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusAutoIsolateSampleDownloadURLRequestParams struct {
	// 样本Md5值
	MD5 *string `json:"MD5,omitempty" name:"MD5"`
}

type DescribeVirusAutoIsolateSampleDownloadURLRequest struct {
	*tchttp.BaseRequest
	
	// 样本Md5值
	MD5 *string `json:"MD5,omitempty" name:"MD5"`
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusAutoIsolateSampleDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MD5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusAutoIsolateSampleDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusAutoIsolateSampleDownloadURLResponseParams struct {
	// 样本下载链接
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusAutoIsolateSampleDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusAutoIsolateSampleDownloadURLResponseParams `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSampleDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusAutoIsolateSampleDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusAutoIsolateSampleListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>MD5- String - 是否必填：否 - md5 </li>
	// <li>AutoIsolateSwitch- String - 是否必填：否 - 自动隔离开关 </li>
	// <li>VirusName- String - 是否必填：否 - 病毒名 </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeVirusAutoIsolateSampleListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>MD5- String - 是否必填：否 - md5 </li>
	// <li>AutoIsolateSwitch- String - 是否必填：否 - 自动隔离开关 </li>
	// <li>VirusName- String - 是否必填：否 - 病毒名 </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVirusAutoIsolateSampleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusAutoIsolateSampleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusAutoIsolateSampleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusAutoIsolateSampleListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 木马自动隔离样本列表
	List []*VirusAutoIsolateSampleInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusAutoIsolateSampleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusAutoIsolateSampleListResponseParams `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSampleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusAutoIsolateSampleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusAutoIsolateSettingRequestParams struct {

}

type DescribeVirusAutoIsolateSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVirusAutoIsolateSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusAutoIsolateSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusAutoIsolateSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusAutoIsolateSettingResponseParams struct {
	// 自动隔离开关(true:开 false:关)
	AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`

	// 是否中断隔离文件关联的进程(true:是 false:否)
	IsKillProgress *bool `json:"IsKillProgress,omitempty" name:"IsKillProgress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusAutoIsolateSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusAutoIsolateSettingResponseParams `json:"Response"`
}

func (r *DescribeVirusAutoIsolateSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusAutoIsolateSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusDetailRequestParams struct {
	// 木马文件id
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeVirusDetailRequest struct {
	*tchttp.BaseRequest
	
	// 木马文件id
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeVirusDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusDetailResponseParams struct {
	// 镜像ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 木马文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 木马文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 最近生成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 病毒名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 容器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 容器id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 主机名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 主机id
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostId *string `json:"HostId,omitempty" name:"HostId"`

	// 进程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 进程md5
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`

	// 进程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessId *uint64 `json:"ProcessId,omitempty" name:"ProcessId"`

	// 进程参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessArgv *string `json:"ProcessArgv,omitempty" name:"ProcessArgv"`

	// 进程链
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessChan *string `json:"ProcessChan,omitempty" name:"ProcessChan"`

	// 进程组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessAccountGroup *string `json:"ProcessAccountGroup,omitempty" name:"ProcessAccountGroup"`

	// 进程启动用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessStartAccount *string `json:"ProcessStartAccount,omitempty" name:"ProcessStartAccount"`

	// 进程文件权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessFileAuthority *string `json:"ProcessFileAuthority,omitempty" name:"ProcessFileAuthority"`

	// 来源：0：一键扫描， 1：定时扫描 2：实时监控
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 事件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`

	// 建议方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mark *string `json:"Mark,omitempty" name:"Mark"`

	// 风险文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// DEAL_NONE:文件待处理
	// DEAL_IGNORE:已经忽略
	// DEAL_ADD_WHITELIST:加白
	// DEAL_DEL:文件已经删除
	// DEAL_ISOLATE:已经隔离
	// DEAL_ISOLATING:隔离中
	// DEAL_ISOLATE_FAILED:隔离失败
	// DEAL_RECOVERING:恢复中
	// DEAL_RECOVER_FAILED: 恢复失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 失败子状态:
	// FILE_NOT_FOUND:文件不存在
	// FILE_ABNORMAL:文件异常
	// FILE_ABNORMAL_DEAL_RECOVER:恢复文件时，文件异常
	// BACKUP_FILE_NOT_FOUND:备份文件不存在
	// CONTAINER_NOT_FOUND_DEAL_ISOLATE:隔离时，容器不存在
	// CONTAINER_NOT_FOUND_DEAL_RECOVER:恢复时，容器不存在
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubStatus *string `json:"SubStatus,omitempty" name:"SubStatus"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 外网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

	// 父进程启动用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	PProcessStartUser *string `json:"PProcessStartUser,omitempty" name:"PProcessStartUser"`

	// 父进程用户组
	// 注意：此字段可能返回 null，表示取不到有效值。
	PProcessUserGroup *string `json:"PProcessUserGroup,omitempty" name:"PProcessUserGroup"`

	// 父进程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PProcessPath *string `json:"PProcessPath,omitempty" name:"PProcessPath"`

	// 父进程命令行参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PProcessParam *string `json:"PProcessParam,omitempty" name:"PProcessParam"`

	// 祖先进程启动用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	AncestorProcessStartUser *string `json:"AncestorProcessStartUser,omitempty" name:"AncestorProcessStartUser"`

	// 祖先进程用户组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AncestorProcessUserGroup *string `json:"AncestorProcessUserGroup,omitempty" name:"AncestorProcessUserGroup"`

	// 祖先进程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	AncestorProcessPath *string `json:"AncestorProcessPath,omitempty" name:"AncestorProcessPath"`

	// 祖先进程命令行参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AncestorProcessParam *string `json:"AncestorProcessParam,omitempty" name:"AncestorProcessParam"`

	// 事件最后一次处理的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`

	// 容器隔离状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 容器隔离子状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`

	// 检测平台
	// 1: 云查杀引擎
	// 2: tav
	// 3: binaryAi
	// 4: 异常行为
	// 5: 威胁情报
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusDetailResponseParams `json:"Response"`
}

func (r *DescribeVirusDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusEventTendencyRequestParams struct {
	// 趋势周期(默认为7天)
	TendencyPeriod *uint64 `json:"TendencyPeriod,omitempty" name:"TendencyPeriod"`
}

type DescribeVirusEventTendencyRequest struct {
	*tchttp.BaseRequest
	
	// 趋势周期(默认为7天)
	TendencyPeriod *uint64 `json:"TendencyPeriod,omitempty" name:"TendencyPeriod"`
}

func (r *DescribeVirusEventTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusEventTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TendencyPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusEventTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusEventTendencyResponseParams struct {
	// 趋势列表
	List []*VirusTendencyInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusEventTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusEventTendencyResponseParams `json:"Response"`
}

func (r *DescribeVirusEventTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>FileName - String - 是否必填：否 - 文件名称</li>
	// <li>FilePath - String - 是否必填：否 - 文件路径</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名称</li>
	// <li>ContainerName- String - 是否必填：是 - 容器名称</li>
	// <li>ContainerId- string - 是否必填：否 - 容器id</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称</li>
	// <li>ImageId- string - 是否必填：否 - 镜像id</li>
	// <li>IsRealTime- int - 是否必填：否 - 过滤是否实时监控数据</li>
	// <li>TaskId- string - 是否必填：否 - 任务ID</li>
	// <li>ContainerNetStatus - String -是否必填: 否 -  容器网络状态筛选 NORMAL ISOLATED ISOLATING RESTORING RESTORE_FAILED</li>
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>ContainerStatus - string -是否必填: 否 - 容器状态 RUNNING PAUSED STOPPED CREATED DESTROYED RESTARTING REMOVING</li>
	// <li>AutoIsolateMode - string -是否必填: 否 - 隔离方式 MANUAL AUTO</li>
	// <li>MD5 - string -是否必填: 否 - md5 </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeVirusListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>FileName - String - 是否必填：否 - 文件名称</li>
	// <li>FilePath - String - 是否必填：否 - 文件路径</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名称</li>
	// <li>ContainerName- String - 是否必填：是 - 容器名称</li>
	// <li>ContainerId- string - 是否必填：否 - 容器id</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称</li>
	// <li>ImageId- string - 是否必填：否 - 镜像id</li>
	// <li>IsRealTime- int - 是否必填：否 - 过滤是否实时监控数据</li>
	// <li>TaskId- string - 是否必填：否 - 任务ID</li>
	// <li>ContainerNetStatus - String -是否必填: 否 -  容器网络状态筛选 NORMAL ISOLATED ISOLATING RESTORING RESTORE_FAILED</li>
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>ContainerStatus - string -是否必填: 否 - 容器状态 RUNNING PAUSED STOPPED CREATED DESTROYED RESTARTING REMOVING</li>
	// <li>AutoIsolateMode - string -是否必填: 否 - 隔离方式 MANUAL AUTO</li>
	// <li>MD5 - string -是否必填: 否 - md5 </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVirusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusListResponseParams struct {
	// 木马列表
	List []*VirusInfo `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusListResponseParams `json:"Response"`
}

func (r *DescribeVirusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusManualScanEstimateTimeoutRequestParams struct {
	// 扫描范围0容器1主机节点
	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`

	// true 全选，false 自选
	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`

	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定
	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`
}

type DescribeVirusManualScanEstimateTimeoutRequest struct {
	*tchttp.BaseRequest
	
	// 扫描范围0容器1主机节点
	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`

	// true 全选，false 自选
	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`

	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定
	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`
}

func (r *DescribeVirusManualScanEstimateTimeoutRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusManualScanEstimateTimeoutRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScanRangeType")
	delete(f, "ScanRangeAll")
	delete(f, "ScanIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusManualScanEstimateTimeoutRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusManualScanEstimateTimeoutResponseParams struct {
	// 预估超时时间(h)
	Timeout *float64 `json:"Timeout,omitempty" name:"Timeout"`

	// 单台主机并行扫描容器数
	ContainerScanConcurrencyCount *uint64 `json:"ContainerScanConcurrencyCount,omitempty" name:"ContainerScanConcurrencyCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusManualScanEstimateTimeoutResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusManualScanEstimateTimeoutResponseParams `json:"Response"`
}

func (r *DescribeVirusManualScanEstimateTimeoutResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusManualScanEstimateTimeoutResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusMonitorSettingRequestParams struct {

}

type DescribeVirusMonitorSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVirusMonitorSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusMonitorSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusMonitorSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusMonitorSettingResponseParams struct {
	// 是否开启实时监控
	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`

	// 扫描全部路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`

	// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`

	// 自选排除或扫描的地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusMonitorSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusMonitorSettingResponseParams `json:"Response"`
}

func (r *DescribeVirusMonitorSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusMonitorSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusSampleDownloadUrlRequestParams struct {
	// 木马id
	ID *string `json:"ID,omitempty" name:"ID"`
}

type DescribeVirusSampleDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 木马id
	ID *string `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeVirusSampleDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusSampleDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusSampleDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusSampleDownloadUrlResponseParams struct {
	// 样本下载地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusSampleDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusSampleDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeVirusSampleDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusSampleDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusScanSettingRequestParams struct {

}

type DescribeVirusScanSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVirusScanSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusScanSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusScanSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusScanSettingResponseParams struct {
	// 是否开启定期扫描
	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`

	// 检测周期每隔多少天
	Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`

	// 扫描开始时间
	BeginScanAt *string `json:"BeginScanAt,omitempty" name:"BeginScanAt"`

	// 扫描全部路径
	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`

	// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径
	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`

	// 超时时长，单位小时
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 扫描范围0容器1主机节点
	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`

	// true 全选，false 自选
	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`

	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定
	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`

	// 自选排除或扫描的地址
	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`

	// 一键检测的超时设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClickTimeout *uint64 `json:"ClickTimeout,omitempty" name:"ClickTimeout"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusScanSettingResponseParams `json:"Response"`
}

func (r *DescribeVirusScanSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusScanTaskStatusRequestParams struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

type DescribeVirusScanTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *DescribeVirusScanTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusScanTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusScanTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusScanTaskStatusResponseParams struct {
	// 查杀容器个数
	ContainerTotal *uint64 `json:"ContainerTotal,omitempty" name:"ContainerTotal"`

	// 风险容器个数
	RiskContainerCnt *uint64 `json:"RiskContainerCnt,omitempty" name:"RiskContainerCnt"`

	// 扫描状态 任务状态:
	// SCAN_NONE:无， 
	// SCAN_SCANNING:正在扫描中，
	// SCAN_FINISH：扫描完成， 
	// SCAN_TIMEOUT：扫描超时
	// SCAN_CANCELING: 取消中
	// SCAN_CANCELED:已取消
	Status *string `json:"Status,omitempty" name:"Status"`

	// 扫描进度 I
	Schedule *uint64 `json:"Schedule,omitempty" name:"Schedule"`

	// 已经扫描了的容器个数
	ContainerScanCnt *uint64 `json:"ContainerScanCnt,omitempty" name:"ContainerScanCnt"`

	// 风险个数
	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`

	// 剩余扫描时间
	LeftSeconds *uint64 `json:"LeftSeconds,omitempty" name:"LeftSeconds"`

	// 扫描开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 扫描结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 扫描类型，"CYCLE"：周期扫描， "MANUAL"：手动扫描
	ScanType *string `json:"ScanType,omitempty" name:"ScanType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusScanTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusScanTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeVirusScanTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusScanTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusScanTimeoutSettingRequestParams struct {
	// 设置类型0一键检测，1定时检测
	ScanType *uint64 `json:"ScanType,omitempty" name:"ScanType"`
}

type DescribeVirusScanTimeoutSettingRequest struct {
	*tchttp.BaseRequest
	
	// 设置类型0一键检测，1定时检测
	ScanType *uint64 `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *DescribeVirusScanTimeoutSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusScanTimeoutSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusScanTimeoutSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusScanTimeoutSettingResponseParams struct {
	// 超时时长单位小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusScanTimeoutSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusScanTimeoutSettingResponseParams `json:"Response"`
}

func (r *DescribeVirusScanTimeoutSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusScanTimeoutSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusSummaryRequestParams struct {

}

type DescribeVirusSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVirusSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusSummaryResponseParams struct {
	// 最近的一次扫描任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 木马影响容器个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskContainerCnt *uint64 `json:"RiskContainerCnt,omitempty" name:"RiskContainerCnt"`

	// 待处理风险个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`

	// 病毒库更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusDataBaseModifyTime *string `json:"VirusDataBaseModifyTime,omitempty" name:"VirusDataBaseModifyTime"`

	// 木马影响容器个数较昨日增长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskContainerIncrease *int64 `json:"RiskContainerIncrease,omitempty" name:"RiskContainerIncrease"`

	// 待处理风险个数较昨日增长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskIncrease *int64 `json:"RiskIncrease,omitempty" name:"RiskIncrease"`

	// 隔离事件个数较昨日新增
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateIncrease *int64 `json:"IsolateIncrease,omitempty" name:"IsolateIncrease"`

	// 隔离事件总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateCnt *int64 `json:"IsolateCnt,omitempty" name:"IsolateCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusSummaryResponseParams `json:"Response"`
}

func (r *DescribeVirusSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusTaskListRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ContainerName - String - 是否必填：否 - 容器名称</li>
	// <li>ContainerId - String - 是否必填：否 - 容器id</li>
	// <li>Hostname - String - 是否必填：否 - 主机名称</li>
	// <li>HostIp- String - 是否必填：否 - 主机IP</li>
	// <li>ImageId- String - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String - 是否必填：否 - 镜像名称</li>
	// <li>Status- String - 是否必填：否 - 状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeVirusTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ContainerName - String - 是否必填：否 - 容器名称</li>
	// <li>ContainerId - String - 是否必填：否 - 容器id</li>
	// <li>Hostname - String - 是否必填：否 - 主机名称</li>
	// <li>HostIp- String - 是否必填：否 - 主机IP</li>
	// <li>ImageId- String - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String - 是否必填：否 - 镜像名称</li>
	// <li>Status- String - 是否必填：否 - 状态</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeVirusTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVirusTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVirusTaskListResponseParams struct {
	// 文件查杀列表
	List []*VirusTaskInfo `json:"List,omitempty" name:"List"`

	// 总数量(容器任务数量)
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVirusTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVirusTaskListResponseParams `json:"Response"`
}

func (r *DescribeVirusTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVirusTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulContainerListRequestParams struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID</li>
	// <li>ContainerName- String -是否必填: 否 - 容器名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeVulContainerListRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID</li>
	// <li>ContainerName- String -是否必填: 否 - 容器名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulContainerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulContainerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PocID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulContainerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulContainerListResponseParams struct {
	// 容器列表
	List []*VulAffectedContainerInfo `json:"List,omitempty" name:"List"`

	// 容器总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulContainerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulContainerListResponseParams `json:"Response"`
}

func (r *DescribeVulContainerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulContainerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceEventDetailRequestParams struct {
	// 事件ID
	EventID *int64 `json:"EventID,omitempty" name:"EventID"`
}

type DescribeVulDefenceEventDetailRequest struct {
	*tchttp.BaseRequest
	
	// 事件ID
	EventID *int64 `json:"EventID,omitempty" name:"EventID"`
}

func (r *DescribeVulDefenceEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceEventDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulDefenceEventDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceEventDetailResponseParams struct {
	// 漏洞防御事件详细
	EventDetail *VulDefenceEventDetail `json:"EventDetail,omitempty" name:"EventDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulDefenceEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulDefenceEventDetailResponseParams `json:"Response"`
}

func (r *DescribeVulDefenceEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceEventRequestParams struct {
	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_DEFENDED：已防御</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  入侵状态，防御成功：EVENT_DEFENDED，尝试攻击：EVENT_ATTACK </li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号。</li>
	// <li>SourceIP- string - 是否必填：否 - 攻击源IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：事件数量：EventCount
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeVulDefenceEventRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，待处理：EVENT_UNDEAL，EVENT_DEALED：已处理，已忽略：EVENT_IGNORE， EVENT_DEFENDED：已防御</li>
	// <li>ContainerStatus- String - 是否必填：否 - 容器运行状态筛选，已创建：CREATED,正常运行：RUNNING, 暂定运行：PAUSED, 停止运行：	STOPPED，重启中：RESTARTING, 迁移中：REMOVING, 销毁：DESTROYED </li>
	// <li>ContainerNetStatus- String -是否必填: 否 -  容器网络状态筛选 未隔离：NORMAL，已隔离：ISOLATED，隔离失败：ISOLATE_FAILED，解除隔离失败：RESTORE_FAILED，解除隔离中：RESTORING，隔离中：ISOLATING</li>
	// <li>EventType - String -是否必填: 否 -  入侵状态，防御成功：EVENT_DEFENDED，尝试攻击：EVENT_ATTACK </li>
	// <li>TimeRange- String -是否必填: 否 -  时间范围，第一个值表示开始时间，第二个值表示结束时间 </li>
	// <li>VulName- string - 是否必填：否 - 漏洞名称。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号。</li>
	// <li>SourceIP- string - 是否必填：否 - 攻击源IP。</li>
	// <li>ContainerName- string - 是否必填：否 - 容器名称。</li>
	// <li>ContainerID- string - 是否必填：否 - 容器ID。</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称。</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID。</li>
	// <li>HostName- string - 是否必填：否 - 主机名称。</li>
	// <li>HostIP- string - 是否必填：否 - 内网IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：事件数量：EventCount
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulDefenceEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulDefenceEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceEventResponseParams struct {
	// 漏洞防御事件列表
	List []*VulDefenceEvent `json:"List,omitempty" name:"List"`

	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulDefenceEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulDefenceEventResponseParams `json:"Response"`
}

func (r *DescribeVulDefenceEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceEventTendencyRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeVulDefenceEventTendencyRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeVulDefenceEventTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceEventTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulDefenceEventTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceEventTendencyResponseParams struct {
	// 漏洞防御事件趋势
	DefendedList []*VulDefenceEventTendency `json:"DefendedList,omitempty" name:"DefendedList"`

	// 漏洞攻击事件趋势
	AttackList []*VulDefenceEventTendency `json:"AttackList,omitempty" name:"AttackList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulDefenceEventTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulDefenceEventTendencyResponseParams `json:"Response"`
}

func (r *DescribeVulDefenceEventTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceEventTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceHostRequestParams struct {
	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，正常：SUCCESS，异常：FAIL， NO_DEFENCE:未防御</li>
	// <li>KeyWords- string - 是否必填：否 - 主机名称/IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：更新时间：ModifyTime/首次开启时间：CreateTime
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeVulDefenceHostRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>Status- String - 是否必填：否 - 插件状态，正常：SUCCESS，异常：FAIL， NO_DEFENCE:未防御</li>
	// <li>KeyWords- string - 是否必填：否 - 主机名称/IP。</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式：asc/desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：更新时间：ModifyTime/首次开启时间：CreateTime
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulDefenceHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulDefenceHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceHostResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 漏洞防御的主机列表
	List []*VulDefenceHost `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulDefenceHostResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulDefenceHostResponseParams `json:"Response"`
}

func (r *DescribeVulDefenceHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefencePluginRequestParams struct {
	// 主机HostID即quuid
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status- String - 是否必填：否 -插件运行状态：注入中:INJECTING，注入成功：SUCCESS，注入失败：FAIL，插件超时：TIMEOUT，插件退出：QUIT</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeVulDefencePluginRequest struct {
	*tchttp.BaseRequest
	
	// 主机HostID即quuid
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Status- String - 是否必填：否 -插件运行状态：注入中:INJECTING，注入成功：SUCCESS，注入失败：FAIL，插件超时：TIMEOUT，插件退出：QUIT</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulDefencePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefencePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulDefencePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefencePluginResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 漏洞防御插件列表
	List []*VulDefencePlugin `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulDefencePluginResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulDefencePluginResponseParams `json:"Response"`
}

func (r *DescribeVulDefencePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefencePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceSettingRequestParams struct {

}

type DescribeVulDefenceSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVulDefenceSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulDefenceSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDefenceSettingResponseParams struct {
	// 是否开启:0: 关闭 1:开启
	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 漏洞防御主机范围: 0:自选主机节点，1:全部
	Scope *int64 `json:"Scope,omitempty" name:"Scope"`

	// 漏洞防御主机数量
	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`

	// 开启漏洞防御异常主机数量
	ExceptionHostCount *int64 `json:"ExceptionHostCount,omitempty" name:"ExceptionHostCount"`

	// 自选漏洞防御主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIDs []*string `json:"HostIDs,omitempty" name:"HostIDs"`

	// 开通容器安全的主机总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostTotalCount *int64 `json:"HostTotalCount,omitempty" name:"HostTotalCount"`

	// 支持防御的漏洞数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportDefenseVulCount *int64 `json:"SupportDefenseVulCount,omitempty" name:"SupportDefenseVulCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulDefenceSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulDefenceSettingResponseParams `json:"Response"`
}

func (r *DescribeVulDefenceSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDefenceSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDetailRequestParams struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

type DescribeVulDetailRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

func (r *DescribeVulDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PocID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDetailResponseParams struct {
	// 漏洞详情信息
	VulInfo *VulDetailInfo `json:"VulInfo,omitempty" name:"VulInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulDetailResponseParams `json:"Response"`
}

func (r *DescribeVulDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulIgnoreLocalImageListRequestParams struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式:DESC,ACS
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段 ImageSize
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeVulIgnoreLocalImageListRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式:DESC,ACS
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段 ImageSize
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulIgnoreLocalImageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulIgnoreLocalImageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PocID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulIgnoreLocalImageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulIgnoreLocalImageListResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像列表
	List []*VulIgnoreLocalImage `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulIgnoreLocalImageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulIgnoreLocalImageListResponseParams `json:"Response"`
}

func (r *DescribeVulIgnoreLocalImageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulIgnoreLocalImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulIgnoreRegistryImageListRequestParams struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeVulIgnoreRegistryImageListRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVulIgnoreRegistryImageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulIgnoreRegistryImageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PocID")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulIgnoreRegistryImageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulIgnoreRegistryImageListResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像列表
	List []*VulIgnoreRegistryImage `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulIgnoreRegistryImageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulIgnoreRegistryImageListResponseParams `json:"Response"`
}

func (r *DescribeVulIgnoreRegistryImageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulIgnoreRegistryImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulImageListRequestParams struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>HostIP- string -是否必填: 否 - 内网IP</li>
	// <li>PublicIP- string -是否必填: 否 - 外网IP</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>HostName- string -是否必填: 否 - 主机名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeVulImageListRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>HostIP- string -是否必填: 否 - 内网IP</li>
	// <li>PublicIP- string -是否必填: 否 - 外网IP</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>HostName- string -是否必填: 否 - 主机名称</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulImageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulImageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PocID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulImageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulImageListResponseParams struct {
	// 受影响的镜像列表
	List []*VulAffectedImageInfo `json:"List,omitempty" name:"List"`

	// 镜像总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulImageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulImageListResponseParams `json:"Response"`
}

func (r *DescribeVulImageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulImageSummaryRequestParams struct {

}

type DescribeVulImageSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVulImageSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulImageSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulImageSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulImageSummaryResponseParams struct {
	// 受严重或高危漏洞影响的镜像数
	SeriousVulImageCount *int64 `json:"SeriousVulImageCount,omitempty" name:"SeriousVulImageCount"`

	// 已扫描的镜像数
	ScannedImageCount *int64 `json:"ScannedImageCount,omitempty" name:"ScannedImageCount"`

	// 漏洞总数量
	VulTotalCount *int64 `json:"VulTotalCount,omitempty" name:"VulTotalCount"`

	// 系统漏洞数
	SysTemVulCount *int64 `json:"SysTemVulCount,omitempty" name:"SysTemVulCount"`

	// web应用漏洞数
	WebVulCount *int64 `json:"WebVulCount,omitempty" name:"WebVulCount"`

	// 已授权镜像数
	AllAuthorizedImageCount *int64 `json:"AllAuthorizedImageCount,omitempty" name:"AllAuthorizedImageCount"`

	// 应急漏洞数
	EmergencyVulCount *int64 `json:"EmergencyVulCount,omitempty" name:"EmergencyVulCount"`

	// 支持扫描的漏洞总数量
	SupportVulTotalCount *int64 `json:"SupportVulTotalCount,omitempty" name:"SupportVulTotalCount"`

	// 漏洞库更新时间
	VulLibraryUpdateTime *string `json:"VulLibraryUpdateTime,omitempty" name:"VulLibraryUpdateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulImageSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulImageSummaryResponseParams `json:"Response"`
}

func (r *DescribeVulImageSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulImageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulLevelImageSummaryRequestParams struct {

}

type DescribeVulLevelImageSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVulLevelImageSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulLevelImageSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulLevelImageSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulLevelImageSummaryResponseParams struct {
	// 高危漏洞最新本地镜像占比
	HighLevelVulLocalImagePercent *float64 `json:"HighLevelVulLocalImagePercent,omitempty" name:"HighLevelVulLocalImagePercent"`

	// 中危漏洞最新本地镜像占比
	MediumLevelVulLocalImagePercent *float64 `json:"MediumLevelVulLocalImagePercent,omitempty" name:"MediumLevelVulLocalImagePercent"`

	// 低危漏洞最新本地镜像占比
	LowLevelVulLocalImagePercent *float64 `json:"LowLevelVulLocalImagePercent,omitempty" name:"LowLevelVulLocalImagePercent"`

	// 严重漏洞最新本地镜像占比
	CriticalLevelVulLocalImagePercent *float64 `json:"CriticalLevelVulLocalImagePercent,omitempty" name:"CriticalLevelVulLocalImagePercent"`

	// 影响的最新版本本地镜像数
	LocalNewestImageCount *int64 `json:"LocalNewestImageCount,omitempty" name:"LocalNewestImageCount"`

	// 影响的最新版本仓库镜像数
	RegistryNewestImageCount *int64 `json:"RegistryNewestImageCount,omitempty" name:"RegistryNewestImageCount"`

	// 高危漏洞最新仓库镜像占比
	HighLevelVulRegistryImagePercent *float64 `json:"HighLevelVulRegistryImagePercent,omitempty" name:"HighLevelVulRegistryImagePercent"`

	// 中危漏洞最新仓库镜像占比
	MediumLevelVulRegistryImagePercent *float64 `json:"MediumLevelVulRegistryImagePercent,omitempty" name:"MediumLevelVulRegistryImagePercent"`

	// 低危漏洞最新仓库镜像占比
	LowLevelVulRegistryImagePercent *float64 `json:"LowLevelVulRegistryImagePercent,omitempty" name:"LowLevelVulRegistryImagePercent"`

	// 严重漏洞最新仓库镜像占比
	CriticalLevelVulRegistryImagePercent *float64 `json:"CriticalLevelVulRegistryImagePercent,omitempty" name:"CriticalLevelVulRegistryImagePercent"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulLevelImageSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulLevelImageSummaryResponseParams `json:"Response"`
}

func (r *DescribeVulLevelImageSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulLevelImageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulLevelSummaryRequestParams struct {
	// 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 EMERGENCY:应急漏洞
	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
}

type DescribeVulLevelSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 EMERGENCY:应急漏洞
	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
}

func (r *DescribeVulLevelSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulLevelSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CategoryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulLevelSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulLevelSummaryResponseParams struct {
	// 高危漏洞数
	HighLevelVulCount *int64 `json:"HighLevelVulCount,omitempty" name:"HighLevelVulCount"`

	// 中危漏洞数
	MediumLevelVulCount *int64 `json:"MediumLevelVulCount,omitempty" name:"MediumLevelVulCount"`

	// 低危漏洞数
	LowLevelVulCount *int64 `json:"LowLevelVulCount,omitempty" name:"LowLevelVulCount"`

	// 严重漏洞数
	CriticalLevelVulCount *int64 `json:"CriticalLevelVulCount,omitempty" name:"CriticalLevelVulCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulLevelSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulLevelSummaryResponseParams `json:"Response"`
}

func (r *DescribeVulLevelSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulLevelSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanAuthorizedImageSummaryRequestParams struct {

}

type DescribeVulScanAuthorizedImageSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVulScanAuthorizedImageSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanAuthorizedImageSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulScanAuthorizedImageSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanAuthorizedImageSummaryResponseParams struct {
	// 全部已授权的本地镜像数
	AllAuthorizedImageCount *int64 `json:"AllAuthorizedImageCount,omitempty" name:"AllAuthorizedImageCount"`

	// 已授权未扫描的本地镜像数
	UnScanAuthorizedImageCount *int64 `json:"UnScanAuthorizedImageCount,omitempty" name:"UnScanAuthorizedImageCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulScanAuthorizedImageSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulScanAuthorizedImageSummaryResponseParams `json:"Response"`
}

func (r *DescribeVulScanAuthorizedImageSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanAuthorizedImageSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanInfoRequestParams struct {
	// 本地镜像漏洞扫描任务ID，无则返回最近一次的漏洞任务扫描
	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`

	// 仓库镜像漏洞扫描任务ID，无则返回最近一次的漏洞任务扫描
	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
}

type DescribeVulScanInfoRequest struct {
	*tchttp.BaseRequest
	
	// 本地镜像漏洞扫描任务ID，无则返回最近一次的漏洞任务扫描
	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`

	// 仓库镜像漏洞扫描任务ID，无则返回最近一次的漏洞任务扫描
	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
}

func (r *DescribeVulScanInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LocalTaskID")
	delete(f, "RegistryTaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulScanInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanInfoResponseParams struct {
	// 本次扫描的本地镜像总数
	LocalImageScanCount *int64 `json:"LocalImageScanCount,omitempty" name:"LocalImageScanCount"`

	// 忽略的漏洞数量
	IgnoreVulCount *int64 `json:"IgnoreVulCount,omitempty" name:"IgnoreVulCount"`

	// 漏洞扫描的开始时间
	ScanStartTime *string `json:"ScanStartTime,omitempty" name:"ScanStartTime"`

	// 漏洞扫描的结束时间
	ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`

	// 发现风险镜像数量
	FoundRiskImageCount *int64 `json:"FoundRiskImageCount,omitempty" name:"FoundRiskImageCount"`

	// 本地发现漏洞数量
	FoundVulCount *int64 `json:"FoundVulCount,omitempty" name:"FoundVulCount"`

	// 扫描进度
	ScanProgress *float64 `json:"ScanProgress,omitempty" name:"ScanProgress"`

	// 本次扫描的仓库镜像总数
	RegistryImageScanCount *int64 `json:"RegistryImageScanCount,omitempty" name:"RegistryImageScanCount"`

	// 本地镜像最近一次的漏洞任务扫描ID
	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`

	// 扫描状态:NOT_SCAN:未扫描，SCANNING:扫描中,SCANNED:完成，CANCELED：扫描停止
	Status *string `json:"Status,omitempty" name:"Status"`

	// 剩余时间，秒
	RemainingTime *float64 `json:"RemainingTime,omitempty" name:"RemainingTime"`

	// 仓库镜像最近一次的漏洞任务扫描ID
	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`

	// 仓库发现漏洞数量
	RegistryFoundVulCount *int64 `json:"RegistryFoundVulCount,omitempty" name:"RegistryFoundVulCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulScanInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulScanInfoResponseParams `json:"Response"`
}

func (r *DescribeVulScanInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanLocalImageListRequestParams struct {
	// 漏洞扫描任务ID
	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ScanStatus- string -是否必填: 否 - 检测状态。WAIT_SCAN：待检测，SCANNING：检查中，SCANNED：检查完成，SCAN_ERR：检查失败，CANCELED：检测停止</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeVulScanLocalImageListRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞扫描任务ID
	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`

	// 过滤条件。
	// <li>OnlyAffectedNewestImage- Bool- 是否必填：否 - 仅展示影响最新版本镜像的漏洞</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ScanStatus- string -是否必填: 否 - 检测状态。WAIT_SCAN：待检测，SCANNING：检查中，SCANNED：检查完成，SCAN_ERR：检查失败，CANCELED：检测停止</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeVulScanLocalImageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanLocalImageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulScanLocalImageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulScanLocalImageListResponseParams struct {
	// 镜像总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像列表
	List []*VulScanImageInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulScanLocalImageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulScanLocalImageListResponseParams `json:"Response"`
}

func (r *DescribeVulScanLocalImageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulScanLocalImageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulSummaryRequestParams struct {
	// 过滤条件。
	// <li>OnlyAffectedNewestImage- string- 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>OnlyAffectedContainer-string- 是否必填：否 - 仅展示影响容器的漏洞,true,false</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 ALL:全部漏洞</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

type DescribeVulSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>OnlyAffectedNewestImage- string- 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>OnlyAffectedContainer-string- 是否必填：否 - 仅展示影响容器的漏洞,true,false</li>
	// <li>CategoryType- string - 是否必填：否 - 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 ALL:全部漏洞</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeVulSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulSummaryResponseParams struct {
	// 漏洞总数量
	VulTotalCount *int64 `json:"VulTotalCount,omitempty" name:"VulTotalCount"`

	// 严重及高危漏洞数量
	SeriousVulCount *int64 `json:"SeriousVulCount,omitempty" name:"SeriousVulCount"`

	// 重点关注漏洞数量
	SuggestVulCount *int64 `json:"SuggestVulCount,omitempty" name:"SuggestVulCount"`

	// 有Poc或者Exp的漏洞数量
	PocExpLevelVulCount *int64 `json:"PocExpLevelVulCount,omitempty" name:"PocExpLevelVulCount"`

	// 有远程Exp的漏洞数量
	RemoteExpLevelVulCount *int64 `json:"RemoteExpLevelVulCount,omitempty" name:"RemoteExpLevelVulCount"`

	// 受严重或高危漏洞影响的最新版本镜像数
	SeriousVulNewestImageCount *int64 `json:"SeriousVulNewestImageCount,omitempty" name:"SeriousVulNewestImageCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulSummaryResponseParams `json:"Response"`
}

func (r *DescribeVulSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulTendencyRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 枚举类型：
	// LATEST：最新版本
	// CONTAINER: 运行容器
	SphereOfInfluence *string `json:"SphereOfInfluence,omitempty" name:"SphereOfInfluence"`
}

type DescribeVulTendencyRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 枚举类型：
	// LATEST：最新版本
	// CONTAINER: 运行容器
	SphereOfInfluence *string `json:"SphereOfInfluence,omitempty" name:"SphereOfInfluence"`
}

func (r *DescribeVulTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SphereOfInfluence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulTendencyResponseParams struct {
	// 漏洞趋势列表
	VulTendencySet []*VulTendencyInfo `json:"VulTendencySet,omitempty" name:"VulTendencySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulTendencyResponseParams `json:"Response"`
}

func (r *DescribeVulTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulTopRankingRequestParams struct {
	// 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 EMERGENCY:应急漏洞
	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
}

type DescribeVulTopRankingRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞分类: SYSTEM:系统漏洞 WEB:web应用漏洞 EMERGENCY:应急漏洞
	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`
}

func (r *DescribeVulTopRankingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulTopRankingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CategoryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulTopRankingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulTopRankingResponseParams struct {
	// 漏洞Top排名信息列表
	List []*VulTopRankingInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulTopRankingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulTopRankingResponseParams `json:"Response"`
}

func (r *DescribeVulTopRankingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulTopRankingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWarningRulesRequestParams struct {

}

type DescribeWarningRulesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWarningRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarningRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWarningRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWarningRulesResponseParams struct {
	// 告警策略列表
	WarningRules []*WarningRule `json:"WarningRules,omitempty" name:"WarningRules"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWarningRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWarningRulesResponseParams `json:"Response"`
}

func (r *DescribeWarningRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarningRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebVulListRequestParams struct {
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

type DescribeWebVulListRequest struct {
	*tchttp.BaseRequest
	
	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>OnlyAffectedContainer- string - 是否必填：否 - 仅展示影响容器的漏洞true,false</li>
	// <li>OnlyAffectedNewestImage-string - 是否必填：否 - 仅展示影响最新版本镜像的漏洞true,false</li>
	// <li>Level- String - 是否必填：否 - 威胁等级，CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低</li>
	// <li>Tags- string - 是否必填：否 - 漏洞标签，POC，EXP。</li>
	// <li>CanBeFixed- string - 是否必填：否 - 是否可修复true,false。</li>
	// <li>CVEID- string - 是否必填：否 - CVE编号</li>
	// <li>ImageID- string - 是否必填：否 - 镜像ID</li>
	// <li>ImageName- String -是否必填: 否 - 镜像名称</li>
	// <li>ContainerID- string -是否必填: 否 - 容器ID</li>
	// <li>ContainerName- string -是否必填: 否 - 容器名称</li>
	// <li>ComponentName- string -是否必填: 否 - 组件名称</li>
	// <li>ComponentVersion- string -是否必填: 否 - 组件版本</li>
	// <li>Name- string -是否必填: 否 - 漏洞名称</li>
	// <li>FocusOnType - string - 是否必填：否 -关注紧急度类型 。ALL :全部，SERIOUS_LEVEL： 严重和高危 ，IS_SUGGEST： 重点关注，POC_EXP 有Poc或Exp ，NETWORK_EXP: 远程Exp</li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeWebVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebVulListResponseParams struct {
	// 漏洞总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 漏洞列表
	List []*VulInfo `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebVulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebVulListResponseParams `json:"Response"`
}

func (r *DescribeWebVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmergencyVulInfo struct {
	// 漏洞名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 漏洞标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// CVSS V3分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`

	// 风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// CVE编号
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 漏洞类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 漏洞披露时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`

	// 最近发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`

	// 应急漏洞风险情况：NOT_SCAN：未扫描，SCANNING：扫描中，SCANNED_NOT_RISK：已扫描，暂未风险 ，SCANNED_RISK：已扫描存在风险
	Status *string `json:"Status,omitempty" name:"Status"`

	// 漏洞ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 防御状态，NO_DEFENDED:未防御，DEFENDED:已防御
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenceStatus *string `json:"DefenceStatus,omitempty" name:"DefenceStatus"`

	// 漏洞防御主机范围: MANUAL:自选主机节点，ALL:全部
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenceScope *string `json:"DefenceScope,omitempty" name:"DefenceScope"`

	// 漏洞防御主机数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenceHostCount *int64 `json:"DefenceHostCount,omitempty" name:"DefenceHostCount"`

	// 已防御攻击次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefendedCount *int64 `json:"DefendedCount,omitempty" name:"DefendedCount"`
}

type EscapeEventDescription struct {
	// 事件规则
	Description *string `json:"Description,omitempty" name:"Description"`

	// 解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 事件最后一次处理的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type EscapeEventInfo struct {
	// 事件类型
	//    ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	//    ESCAPE_SYSCALL:Syscall逃逸
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 镜像名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 状态，EVENT_UNDEAL:未处理，EVENT_DEALED:已处理，EVENT_INGNORE:忽略
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件记录的唯一id
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// pod(实例)的名字
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 生成时间
	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`

	// 事件名字，
	// 宿主机文件访问逃逸、
	// Syscall逃逸、
	// MountNamespace逃逸、
	// 程序提权逃逸、
	// 特权容器启动逃逸、
	// 敏感路径挂载
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 镜像id，用于跳转
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 容器id，用于跳转
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 事件解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 事件数量
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 最近生成时间
	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`

	// 节点IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`

	// 主机IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`

	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING
	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type EscapeEventTendencyInfo struct {
	// 待处理风险容器事件总数
	RiskContainerEventCount *int64 `json:"RiskContainerEventCount,omitempty" name:"RiskContainerEventCount"`

	// 待处理程序特权事件总数
	ProcessPrivilegeEventCount *int64 `json:"ProcessPrivilegeEventCount,omitempty" name:"ProcessPrivilegeEventCount"`

	// 待处理容器逃逸事件总数
	ContainerEscapeEventCount *int64 `json:"ContainerEscapeEventCount,omitempty" name:"ContainerEscapeEventCount"`

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`
}

type EscapeRule struct {
	// 规则类型   
	// ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	// ESCAPE_SYSCALL:Syscall逃逸
	Type *string `json:"Type,omitempty" name:"Type"`

	// 规则名称
	// 宿主机文件访问逃逸、
	// Syscall逃逸、
	// MountNamespace逃逸、
	// 程序提权逃逸、
	// 特权容器启动逃逸、
	// 敏感路径挂载
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否打开：false否 ，true是
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`

	// 规则组别。RISK_CONTAINER：风险容器，PROCESS_PRIVILEGE：程序特权，CONTAINER_ESCAPE：容器逃逸
	Group *string `json:"Group,omitempty" name:"Group"`
}

type EscapeRuleEnabled struct {
	// 规则类型
	//    ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	//    ESCAPE_SYSCALL:Syscall逃逸
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否打开：false否 ，true是
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

type EscapeWhiteListInfo struct {
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 白名单记录ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 关联主机数量
	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`

	// 关联容器数量
	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`

	// 加白事件类型
	EventType []*string `json:"EventType,omitempty" name:"EventType"`

	// 创建时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 镜像大小
	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`
}

type ExportJobInfo struct {
	// 任务ID
	JobID *string `json:"JobID,omitempty" name:"JobID"`

	// 任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// 导出状态
	ExportStatus *string `json:"ExportStatus,omitempty" name:"ExportStatus"`

	// 导出进度
	ExportProgress *int64 `json:"ExportProgress,omitempty" name:"ExportProgress"`

	// 失败原因
	FailureMsg *string `json:"FailureMsg,omitempty" name:"FailureMsg"`

	// 超时时间
	Timeout *string `json:"Timeout,omitempty" name:"Timeout"`

	// 插入时间
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`
}

// Predefined struct for user
type ExportVirusListRequestParams struct {
	// 过滤条件。
	// <li>FileName - String - 是否必填：否 - 文件名称</li>
	// <li>FilePath - String - 是否必填：否 - 文件路径</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名称</li>
	// <li>ContainerName- String - 是否必填：是 - 容器名称</li>
	// <li>ContainerId- string - 是否必填：否 - 容器id</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称</li>
	// <li>ImageId- string - 是否必填：否 - 镜像id</li>
	// <li>IsRealTime- int - 是否必填：否 - 过滤是否实时监控数据</li>
	// <li>TaskId- string - 是否必填：否 - 任务ID</li>
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>ContainerNetStatus - String -是否必填: 否 -  容器网络状态筛选 NORMAL ISOLATED ISOLATING RESTORING RESTORE_FAILED</li>
	// <li>ContainerStatus - string -是否必填: 否 - 容器状态 RUNNING PAUSED STOPPED CREATED DESTROYED RESTARTING REMOVING</li>
	// <li>AutoIsolateMode - string -是否必填: 否 - 隔离方式 MANUAL AUTO</li>
	// <li>MD5 - string -是否必填: 否 - md5 </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

type ExportVirusListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>FileName - String - 是否必填：否 - 文件名称</li>
	// <li>FilePath - String - 是否必填：否 - 文件路径</li>
	// <li>VirusName - String - 是否必填：否 - 病毒名称</li>
	// <li>ContainerName- String - 是否必填：是 - 容器名称</li>
	// <li>ContainerId- string - 是否必填：否 - 容器id</li>
	// <li>ImageName- string - 是否必填：否 - 镜像名称</li>
	// <li>ImageId- string - 是否必填：否 - 镜像id</li>
	// <li>IsRealTime- int - 是否必填：否 - 过滤是否实时监控数据</li>
	// <li>TaskId- string - 是否必填：否 - 任务ID</li>
	// <li>TimeRange - string -是否必填: 否 - 时间范围筛选 ["2022-03-31 16:55:00", "2022-03-31 17:00:00"]</li>
	// <li>ContainerNetStatus - String -是否必填: 否 -  容器网络状态筛选 NORMAL ISOLATED ISOLATING RESTORING RESTORE_FAILED</li>
	// <li>ContainerStatus - string -是否必填: 否 - 容器状态 RUNNING PAUSED STOPPED CREATED DESTROYED RESTARTING REMOVING</li>
	// <li>AutoIsolateMode - string -是否必填: 否 - 隔离方式 MANUAL AUTO</li>
	// <li>MD5 - string -是否必填: 否 - md5 </li>
	Filters []*RunTimeFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 导出字段
	ExportField []*string `json:"ExportField,omitempty" name:"ExportField"`
}

func (r *ExportVirusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVirusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "ExportField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVirusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportVirusListResponseParams struct {
	// 导出任务ID，前端拿着任务ID查询任务进度
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportVirusListResponse struct {
	*tchttp.BaseResponse
	Response *ExportVirusListResponseParams `json:"Response"`
}

func (r *ExportVirusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVirusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileAttributeInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件大小(字节)
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件路径
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 文件创建时间
	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`

	// 最近被篡改文件创建时间
	LatestTamperedFileMTime *string `json:"LatestTamperedFileMTime,omitempty" name:"LatestTamperedFileMTime"`

	// 新文件内容
	NewFile *string `json:"NewFile,omitempty" name:"NewFile"`

	// 新旧文件的差异
	FileDiff *string `json:"FileDiff,omitempty" name:"FileDiff"`
}

type HostInfo struct {
	// 主机id
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 主机ip即内网ip
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 业务组
	Group *string `json:"Group,omitempty" name:"Group"`

	// docker 版本
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

	// docker 文件系统类型
	DockerFileSystemDriver *string `json:"DockerFileSystemDriver,omitempty" name:"DockerFileSystemDriver"`

	// 镜像个数
	ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`

	// 容器个数
	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`

	// agent运行状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 是否是Containerd
	IsContainerd *bool `json:"IsContainerd,omitempty" name:"IsContainerd"`

	// 主机来源：["CVM", "ECM", "LH", "BM"]  中的之一为腾讯云服务器；["Other"]之一非腾讯云服务器；
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 外网ip
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 主机uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 地域ID
	RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`

	// 所属项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	Project *ProjectInfo `json:"Project,omitempty" name:"Project"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
}

type ImageAutoAuthorizedTask struct {
	// 任务id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 授权方式，AUTO:自动授权，MANUAL:手动授权
	Type *string `json:"Type,omitempty" name:"Type"`

	// 任务日期
	AuthorizedDate *string `json:"AuthorizedDate,omitempty" name:"AuthorizedDate"`

	// 镜像来源，LOCAL:本地镜像，REGISTRY:仓库镜像
	Source *string `json:"Source,omitempty" name:"Source"`

	// 最近授权时间
	LastAuthorizedTime *string `json:"LastAuthorizedTime,omitempty" name:"LastAuthorizedTime"`

	// 自动授权成功数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 自动授权失败数
	FailCount *int64 `json:"FailCount,omitempty" name:"FailCount"`

	// 最近任务失败码，REACH_LIMIT:达到授权上限，LICENSE_INSUFFICIENT:授权数不足
	LatestFailCode *string `json:"LatestFailCode,omitempty" name:"LatestFailCode"`
}

type ImageComponent struct {
	// 组件名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 组件版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 组件路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 组件类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 组件漏洞数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulCount *uint64 `json:"VulCount,omitempty" name:"VulCount"`

	// 镜像ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`
}

type ImageHost struct {
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 主机id
	HostID *string `json:"HostID,omitempty" name:"HostID"`
}

type ImageInfo struct {
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像tag
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// 强制扫描
	Force *string `json:"Force,omitempty" name:"Force"`

	// 镜像id
	ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`

	// 仓库类型
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 镜像仓库地址
	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ImageProgress struct {
	// 镜像id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 镜像仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`

	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 仓库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// 镜像扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 镜像cve扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	CveProgress *uint64 `json:"CveProgress,omitempty" name:"CveProgress"`

	// 镜像敏感扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskProgress *uint64 `json:"RiskProgress,omitempty" name:"RiskProgress"`

	// 镜像木马扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusProgress *uint64 `json:"VirusProgress,omitempty" name:"VirusProgress"`
}

type ImageRepoInfo struct {
	// 镜像Digest
	ImageDigest *string `json:"ImageDigest,omitempty" name:"ImageDigest"`

	// 镜像仓库地址
	ImageRepoAddress *string `json:"ImageRepoAddress,omitempty" name:"ImageRepoAddress"`

	// 仓库类型
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像版本
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// 镜像大小
	ImageSize *uint64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 最近扫描时间
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 扫描状态
	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 安全漏洞数
	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`

	// 木马病毒数
	VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`

	// 风险行为数
	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`

	// 敏感信息数
	SentiveInfoCnt *uint64 `json:"SentiveInfoCnt,omitempty" name:"SentiveInfoCnt"`

	// 是否可信镜像
	IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`

	// 镜像系统
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 木马扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`

	// 漏洞扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 高危扫描错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`

	// 敏感信息扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVirusProgress *uint64 `json:"ScanVirusProgress,omitempty" name:"ScanVirusProgress"`

	// 木马扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVulProgress *uint64 `json:"ScanVulProgress,omitempty" name:"ScanVulProgress"`

	// 漏洞扫描进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRiskProgress *uint64 `json:"ScanRiskProgress,omitempty" name:"ScanRiskProgress"`

	// 剩余扫描时间秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRemainTime *uint64 `json:"ScanRemainTime,omitempty" name:"ScanRemainTime"`

	// cve扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CveStatus *string `json:"CveStatus,omitempty" name:"CveStatus"`

	// 高危扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskStatus *string `json:"RiskStatus,omitempty" name:"RiskStatus"`

	// 木马扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusStatus *string `json:"VirusStatus,omitempty" name:"VirusStatus"`

	// 总进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 授权状态
	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`

	// 仓库区域
	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`

	// 列表id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 镜像Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像创建的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`

	// 是否为镜像的最新版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLatestImage *bool `json:"IsLatestImage,omitempty" name:"IsLatestImage"`
}

type ImageRisk struct {
	// 高危行为
	// 注意：此字段可能返回 null，表示取不到有效值。
	Behavior *uint64 `json:"Behavior,omitempty" name:"Behavior"`

	// 种类
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstructionContent *string `json:"InstructionContent,omitempty" name:"InstructionContent"`
}

type ImageRiskInfo struct {
	// 行为
	Behavior *uint64 `json:"Behavior,omitempty" name:"Behavior"`

	// 类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 级别
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 详情
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 解决方案
	InstructionContent *string `json:"InstructionContent,omitempty" name:"InstructionContent"`
}

type ImageRiskTendencyInfo struct {
	// 趋势列表
	ImageRiskSet []*RunTimeTendencyInfo `json:"ImageRiskSet,omitempty" name:"ImageRiskSet"`

	// 风险类型：
	// IRT_VULNERABILITY : 安全漏洞
	// IRT_MALWARE_VIRUS: 木马病毒
	// IRT_RISK:敏感信息
	ImageRiskType *string `json:"ImageRiskType,omitempty" name:"ImageRiskType"`
}

type ImageSimpleInfo struct {
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 类型
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// 关联容器数
	ContainerCnt *int64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`
}

type ImageVirus struct {
	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 病毒名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件md5
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 首次发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstScanTime *string `json:"FirstScanTime,omitempty" name:"FirstScanTime"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`
}

type ImageVirusInfo struct {
	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 病毒名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 修护建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 首次发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstScanTime *string `json:"FirstScanTime,omitempty" name:"FirstScanTime"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`

	// 文件md5
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 检测平台
	// 1: 云查杀引擎
	// 2: tav
	// 3: binaryAi
	// 4: 异常行为
	// 5: 威胁情报
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
}

type ImageVul struct {
	// 漏洞id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 观点验证程序id
	// 注意：此字段可能返回 null，表示取不到有效值。
	POCID *string `json:"POCID,omitempty" name:"POCID"`

	// 漏洞名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 涉及组件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*ComponentsInfo `json:"Components,omitempty" name:"Components"`

	// 分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 分类2
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`

	// 风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Des *string `json:"Des,omitempty" name:"Des"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`

	// 引用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reference *string `json:"Reference,omitempty" name:"Reference"`

	// 防御方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`

	// 提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`

	// Cvss分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvssScore *string `json:"CvssScore,omitempty" name:"CvssScore"`

	// Cvss信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvssVector *string `json:"CvssVector,omitempty" name:"CvssVector"`

	// 是否建议修复
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSuggest *string `json:"IsSuggest,omitempty" name:"IsSuggest"`

	// 修复版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixedVersions *string `json:"FixedVersions,omitempty" name:"FixedVersions"`

	// 漏洞标签:"CanBeFixed","DynamicLevelPoc","DynamicLevelExp"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*string `json:"Tag,omitempty" name:"Tag"`

	// 组件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Component *string `json:"Component,omitempty" name:"Component"`

	// 组件版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
}

type ImagesBindRuleInfo struct {
	// 镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 关联容器数量
	ContainerCnt *int64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`

	// 绑定规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 镜像大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`
}

type ImagesInfo struct {
	// 镜像id
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 主机个数
	HostCnt *uint64 `json:"HostCnt,omitempty" name:"HostCnt"`

	// 容器个数
	ContainerCnt *uint64 `json:"ContainerCnt,omitempty" name:"ContainerCnt"`

	// 扫描时间
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 漏洞个数
	VulCnt *uint64 `json:"VulCnt,omitempty" name:"VulCnt"`

	// 病毒个数
	VirusCnt *uint64 `json:"VirusCnt,omitempty" name:"VirusCnt"`

	// 敏感信息个数
	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`

	// 是否信任镜像
	IsTrustImage *bool `json:"IsTrustImage,omitempty" name:"IsTrustImage"`

	// 镜像系统
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// agent镜像扫描错误
	AgentError *string `json:"AgentError,omitempty" name:"AgentError"`

	// 后端镜像扫描错误
	ScanError *string `json:"ScanError,omitempty" name:"ScanError"`

	// 扫描状态
	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 木马扫描错误信息
	ScanVirusError *string `json:"ScanVirusError,omitempty" name:"ScanVirusError"`

	// 漏洞扫描错误信息
	ScanVulError *string `json:"ScanVulError,omitempty" name:"ScanVulError"`

	// 风险扫描错误信息
	ScanRiskError *string `json:"ScanRiskError,omitempty" name:"ScanRiskError"`

	// 是否是重点关注镜像，为0不是，非0是
	IsSuggest *uint64 `json:"IsSuggest,omitempty" name:"IsSuggest"`

	// 是否授权，1是0否
	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`

	// 组件个数
	ComponentCnt *uint64 `json:"ComponentCnt,omitempty" name:"ComponentCnt"`
}

type ImagesVul struct {
	// 漏洞id
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 漏洞名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 组件
	Component *string `json:"Component,omitempty" name:"Component"`

	// 版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 分类
	Category *string `json:"Category,omitempty" name:"Category"`

	// 分类2
	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`

	// 风险等级
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 描述
	Des *string `json:"Des,omitempty" name:"Des"`

	// 解决方案
	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`

	// 引用
	Reference *string `json:"Reference,omitempty" name:"Reference"`

	// 防御方案
	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`

	// 提交时间
	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`

	// CVSS V3分数
	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`

	// CVSS V3描述
	CVSSV3Desc *string `json:"CVSSV3Desc,omitempty" name:"CVSSV3Desc"`

	// 是否是重点关注：true：是，false：不是
	IsSuggest *bool `json:"IsSuggest,omitempty" name:"IsSuggest"`

	// 修复版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixedVersions *string `json:"FixedVersions,omitempty" name:"FixedVersions"`

	// 漏洞标签:"CanBeFixed","DynamicLevelPoc","DynamicLevelExp"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*string `json:"Tag,omitempty" name:"Tag"`
}

// Predefined struct for user
type InitializeUserComplianceEnvironmentRequestParams struct {

}

type InitializeUserComplianceEnvironmentRequest struct {
	*tchttp.BaseRequest
	
}

func (r *InitializeUserComplianceEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitializeUserComplianceEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitializeUserComplianceEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitializeUserComplianceEnvironmentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InitializeUserComplianceEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *InitializeUserComplianceEnvironmentResponseParams `json:"Response"`
}

func (r *InitializeUserComplianceEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitializeUserComplianceEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8sApiAbnormalEventInfo struct {
	// 命中规则名称
	MatchRuleName *string `json:"MatchRuleName,omitempty" name:"MatchRuleName"`

	// 命中规则类型
	MatchRuleType *string `json:"MatchRuleType,omitempty" name:"MatchRuleType"`

	// 告警等级
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 集群ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群运行状态
	ClusterRunningStatus *string `json:"ClusterRunningStatus,omitempty" name:"ClusterRunningStatus"`

	// 初次生成时间
	FirstCreateTime *string `json:"FirstCreateTime,omitempty" name:"FirstCreateTime"`

	// 最近一次生成时间
	LastCreateTime *string `json:"LastCreateTime,omitempty" name:"LastCreateTime"`

	// 告警数量
	AlarmCount *uint64 `json:"AlarmCount,omitempty" name:"AlarmCount"`

	// 状态
	// "EVENT_UNDEAL":未处理
	// "EVENT_DEALED": 已处理
	// "EVENT_IGNORE": 忽略
	// "EVENT_DEL": 删除
	// "EVENT_ADD_WHITE": 加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 集群masterIP
	ClusterMasterIP *string `json:"ClusterMasterIP,omitempty" name:"ClusterMasterIP"`

	// k8s版本
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// 运行时组件
	RunningComponent []*string `json:"RunningComponent,omitempty" name:"RunningComponent"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 建议
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 请求信息
	Info *string `json:"Info,omitempty" name:"Info"`

	// 规则ID
	MatchRuleID *string `json:"MatchRuleID,omitempty" name:"MatchRuleID"`

	// 高亮字段数组
	HighLightFields []*string `json:"HighLightFields,omitempty" name:"HighLightFields"`

	// 命中规则
	MatchRule *K8sApiAbnormalRuleScopeInfo `json:"MatchRule,omitempty" name:"MatchRule"`
}

type K8sApiAbnormalEventListItem struct {
	// 事件ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 命中规则类型
	MatchRuleType *string `json:"MatchRuleType,omitempty" name:"MatchRuleType"`

	// 威胁等级
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 集群ID
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群运行状态
	ClusterRunningStatus *string `json:"ClusterRunningStatus,omitempty" name:"ClusterRunningStatus"`

	// 初次生成时间
	FirstCreateTime *string `json:"FirstCreateTime,omitempty" name:"FirstCreateTime"`

	// 最近一次生成时间
	LastCreateTime *string `json:"LastCreateTime,omitempty" name:"LastCreateTime"`

	// 告警数量
	AlarmCount *uint64 `json:"AlarmCount,omitempty" name:"AlarmCount"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 规则类型
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 描述信息
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 解决方案
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 命中规则
	MatchRule *K8sApiAbnormalRuleScopeInfo `json:"MatchRule,omitempty" name:"MatchRule"`
}

type K8sApiAbnormalRuleInfo struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 状态
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 规则信息列表
	RuleInfoList []*K8sApiAbnormalRuleScopeInfo `json:"RuleInfoList,omitempty" name:"RuleInfoList"`

	// 生效集群IDSet
	EffectClusterIDSet []*string `json:"EffectClusterIDSet,omitempty" name:"EffectClusterIDSet"`

	// 规则类型
	// RT_SYSTEM 系统规则
	// RT_USER 用户自定义
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 是否所有集群生效
	EffectAllCluster *bool `json:"EffectAllCluster,omitempty" name:"EffectAllCluster"`

	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

type K8sApiAbnormalRuleListItem struct {
	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则类型
	// RT_SYSTEM 系统规则
	// RT_USER 用户自定义
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 受影响集群总数
	EffectClusterCount *uint64 `json:"EffectClusterCount,omitempty" name:"EffectClusterCount"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 编辑账号
	OprUin *string `json:"OprUin,omitempty" name:"OprUin"`

	// 状态
	Status *bool `json:"Status,omitempty" name:"Status"`
}

type K8sApiAbnormalRuleScopeInfo struct {
	// 范围
	// 系统事件:
	// ANONYMOUS_ACCESS: 匿名访问
	// ABNORMAL_UA_REQ: 异常UA请求
	// ANONYMOUS_ABNORMAL_PERMISSION: 匿名用户权限异动
	// GET_CREDENTIALS: 凭据信息获取
	// MOUNT_SENSITIVE_PATH: 敏感路径挂载
	// COMMAND_RUN: 命令执行
	// PRIVILEGE_CONTAINER: 特权容器
	// EXCEPTION_CRONTAB_TASK: 异常定时任务
	// STATICS_POD: 静态pod创建
	// ABNORMAL_CREATE_POD: 异常pod创建
	// USER_DEFINED: 用户自定义
	Scope *string `json:"Scope,omitempty" name:"Scope"`

	// 动作(RULE_MODE_ALERT: 告警 RULE_MODE_RELEASE:放行)
	Action *string `json:"Action,omitempty" name:"Action"`

	// 威胁等级 HIGH:高级 MIDDLE: 中级 LOW:低级 NOTICE:提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 开关状态(true:开 false:关) 适用于系统规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 是否被删除 适用于自定义规则入参
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDelete *bool `json:"IsDelete,omitempty" name:"IsDelete"`
}

type K8sApiAbnormalTendencyItem struct {
	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 异常UA请求事件数
	ExceptionUARequestCount *uint64 `json:"ExceptionUARequestCount,omitempty" name:"ExceptionUARequestCount"`

	// 匿名用户权限事件数
	AnonymousUserRightCount *uint64 `json:"AnonymousUserRightCount,omitempty" name:"AnonymousUserRightCount"`

	// 凭据信息获取事件数
	CredentialInformationObtainCount *uint64 `json:"CredentialInformationObtainCount,omitempty" name:"CredentialInformationObtainCount"`

	// 敏感数据挂载事件数
	SensitiveDataMountCount *uint64 `json:"SensitiveDataMountCount,omitempty" name:"SensitiveDataMountCount"`

	// 命令执行事件数
	CmdExecCount *uint64 `json:"CmdExecCount,omitempty" name:"CmdExecCount"`

	// 异常定时任务事件数
	AbnormalScheduledTaskCount *uint64 `json:"AbnormalScheduledTaskCount,omitempty" name:"AbnormalScheduledTaskCount"`

	// 静态Pod创建数
	StaticsPodCreateCount *uint64 `json:"StaticsPodCreateCount,omitempty" name:"StaticsPodCreateCount"`

	// 可疑容器创建数
	DoubtfulContainerCreateCount *uint64 `json:"DoubtfulContainerCreateCount,omitempty" name:"DoubtfulContainerCreateCount"`

	// 自定义规则事件数
	UserDefinedRuleCount *uint64 `json:"UserDefinedRuleCount,omitempty" name:"UserDefinedRuleCount"`

	// 匿名访问事件数
	AnonymousAccessCount *uint64 `json:"AnonymousAccessCount,omitempty" name:"AnonymousAccessCount"`

	// 特权容器事件数
	PrivilegeContainerCount *uint64 `json:"PrivilegeContainerCount,omitempty" name:"PrivilegeContainerCount"`
}

// Predefined struct for user
type ModifyAbnormalProcessRuleStatusRequestParams struct {
	// 策略的ids
	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`

	// 策略开关，true开启，false关闭
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

type ModifyAbnormalProcessRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 策略的ids
	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`

	// 策略开关，true开启，false关闭
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

func (r *ModifyAbnormalProcessRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAbnormalProcessRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIdSet")
	delete(f, "IsEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAbnormalProcessRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAbnormalProcessRuleStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAbnormalProcessRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAbnormalProcessRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyAbnormalProcessRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAbnormalProcessRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAbnormalProcessStatusRequestParams struct {
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，   
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyAbnormalProcessStatusRequest struct {
	*tchttp.BaseRequest
	
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，   
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAbnormalProcessStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAbnormalProcessStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIdSet")
	delete(f, "Status")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAbnormalProcessStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAbnormalProcessStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAbnormalProcessStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAbnormalProcessStatusResponseParams `json:"Response"`
}

func (r *ModifyAbnormalProcessStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAbnormalProcessStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessControlRuleStatusRequestParams struct {
	// 策略的ids
	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`

	// 策略开关，true:代表开启， false代表关闭
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

type ModifyAccessControlRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 策略的ids
	RuleIdSet []*string `json:"RuleIdSet,omitempty" name:"RuleIdSet"`

	// 策略开关，true:代表开启， false代表关闭
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

func (r *ModifyAccessControlRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessControlRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIdSet")
	delete(f, "IsEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessControlRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessControlRuleStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccessControlRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessControlRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyAccessControlRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessControlRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessControlStatusRequestParams struct {
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，     
	// EVENT_DEALED:事件已经处理
	//      EVENT_INGNORE：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注事件信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyAccessControlStatusRequest struct {
	*tchttp.BaseRequest
	
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，     
	// EVENT_DEALED:事件已经处理
	//      EVENT_INGNORE：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注事件信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAccessControlStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessControlStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIdSet")
	delete(f, "Status")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessControlStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessControlStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccessControlStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessControlStatusResponseParams `json:"Response"`
}

func (r *ModifyAccessControlStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessControlStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetImageRegistryScanStopOneKeyRequestParams struct {
	// 是否扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 扫描的镜像列表Id
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

type ModifyAssetImageRegistryScanStopOneKeyRequest struct {
	*tchttp.BaseRequest
	
	// 是否扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 扫描的镜像列表Id
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyAssetImageRegistryScanStopOneKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetImageRegistryScanStopOneKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "All")
	delete(f, "Images")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssetImageRegistryScanStopOneKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetImageRegistryScanStopOneKeyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAssetImageRegistryScanStopOneKeyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAssetImageRegistryScanStopOneKeyResponseParams `json:"Response"`
}

func (r *ModifyAssetImageRegistryScanStopOneKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetImageRegistryScanStopOneKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetImageRegistryScanStopRequestParams struct {
	// 是否扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 扫描的镜像列表
	Id []*uint64 `json:"Id,omitempty" name:"Id"`

	// 过滤条件
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 不要扫描的镜像列表，与Filters配合使用
	ExcludeImageList []*uint64 `json:"ExcludeImageList,omitempty" name:"ExcludeImageList"`

	// 是否仅扫描各repository最新版本的镜像
	OnlyScanLatest *bool `json:"OnlyScanLatest,omitempty" name:"OnlyScanLatest"`
}

type ModifyAssetImageRegistryScanStopRequest struct {
	*tchttp.BaseRequest
	
	// 是否扫描全部镜像
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描的镜像列表
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 扫描的镜像列表
	Id []*uint64 `json:"Id,omitempty" name:"Id"`

	// 过滤条件
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 不要扫描的镜像列表，与Filters配合使用
	ExcludeImageList []*uint64 `json:"ExcludeImageList,omitempty" name:"ExcludeImageList"`

	// 是否仅扫描各repository最新版本的镜像
	OnlyScanLatest *bool `json:"OnlyScanLatest,omitempty" name:"OnlyScanLatest"`
}

func (r *ModifyAssetImageRegistryScanStopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetImageRegistryScanStopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "All")
	delete(f, "Images")
	delete(f, "Id")
	delete(f, "Filters")
	delete(f, "ExcludeImageList")
	delete(f, "OnlyScanLatest")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssetImageRegistryScanStopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetImageRegistryScanStopResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAssetImageRegistryScanStopResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAssetImageRegistryScanStopResponseParams `json:"Response"`
}

func (r *ModifyAssetImageRegistryScanStopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetImageRegistryScanStopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetImageScanStopRequestParams struct {
	// 任务id；任务id，镜像id和根据过滤条件筛选三选一。
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 镜像id；任务id，镜像id和根据过滤条件筛选三选一。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 根据过滤条件筛选出镜像；任务id，镜像id和根据过滤条件筛选三选一。
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 根据过滤条件筛选出镜像，再排除个别镜像
	ExcludeImageIds *string `json:"ExcludeImageIds,omitempty" name:"ExcludeImageIds"`
}

type ModifyAssetImageScanStopRequest struct {
	*tchttp.BaseRequest
	
	// 任务id；任务id，镜像id和根据过滤条件筛选三选一。
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 镜像id；任务id，镜像id和根据过滤条件筛选三选一。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 根据过滤条件筛选出镜像；任务id，镜像id和根据过滤条件筛选三选一。
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 根据过滤条件筛选出镜像，再排除个别镜像
	ExcludeImageIds *string `json:"ExcludeImageIds,omitempty" name:"ExcludeImageIds"`
}

func (r *ModifyAssetImageScanStopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetImageScanStopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	delete(f, "Images")
	delete(f, "Filters")
	delete(f, "ExcludeImageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssetImageScanStopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetImageScanStopResponseParams struct {
	// 停止状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAssetImageScanStopResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAssetImageScanStopResponseParams `json:"Response"`
}

func (r *ModifyAssetImageScanStopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetImageScanStopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetRequestParams struct {
	// 全部同步
	All *bool `json:"All,omitempty" name:"All"`

	// 要同步的主机列表 两个参数必选一个 All优先
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
}

type ModifyAssetRequest struct {
	*tchttp.BaseRequest
	
	// 全部同步
	All *bool `json:"All,omitempty" name:"All"`

	// 要同步的主机列表 两个参数必选一个 All优先
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
}

func (r *ModifyAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "All")
	delete(f, "Hosts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetResponseParams struct {
	// 同步任务发送结果
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAssetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAssetResponseParams `json:"Response"`
}

func (r *ModifyAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompliancePeriodTaskRequestParams struct {
	// 要修改的定时任务的ID，由DescribeCompliancePeriodTaskList接口返回。
	PeriodTaskId *uint64 `json:"PeriodTaskId,omitempty" name:"PeriodTaskId"`

	// 定时任务的周期规则。不填时，不修改。
	PeriodRule *CompliancePeriodTaskRule `json:"PeriodRule,omitempty" name:"PeriodRule"`

	// 设置合规标准。不填时，不修改。
	StandardSettings []*ComplianceBenchmarkStandardEnable `json:"StandardSettings,omitempty" name:"StandardSettings"`
}

type ModifyCompliancePeriodTaskRequest struct {
	*tchttp.BaseRequest
	
	// 要修改的定时任务的ID，由DescribeCompliancePeriodTaskList接口返回。
	PeriodTaskId *uint64 `json:"PeriodTaskId,omitempty" name:"PeriodTaskId"`

	// 定时任务的周期规则。不填时，不修改。
	PeriodRule *CompliancePeriodTaskRule `json:"PeriodRule,omitempty" name:"PeriodRule"`

	// 设置合规标准。不填时，不修改。
	StandardSettings []*ComplianceBenchmarkStandardEnable `json:"StandardSettings,omitempty" name:"StandardSettings"`
}

func (r *ModifyCompliancePeriodTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompliancePeriodTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PeriodTaskId")
	delete(f, "PeriodRule")
	delete(f, "StandardSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCompliancePeriodTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompliancePeriodTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCompliancePeriodTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCompliancePeriodTaskResponseParams `json:"Response"`
}

func (r *ModifyCompliancePeriodTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompliancePeriodTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContainerNetStatusRequestParams struct {
	// 容器ID
	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`

	// 状态(
	// 隔离容器: EVENT_ISOLATE_CONTAINER
	// 恢复容器: EVENT_RESOTRE_CONTAINER
	// )
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyContainerNetStatusRequest struct {
	*tchttp.BaseRequest
	
	// 容器ID
	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`

	// 状态(
	// 隔离容器: EVENT_ISOLATE_CONTAINER
	// 恢复容器: EVENT_RESOTRE_CONTAINER
	// )
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyContainerNetStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContainerNetStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContainerID")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyContainerNetStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContainerNetStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyContainerNetStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyContainerNetStatusResponseParams `json:"Response"`
}

func (r *ModifyContainerNetStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContainerNetStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEscapeEventStatusRequestParams struct {
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态：
	// EVENT_UNDEAL:未处理（取消忽略），
	// EVENT_DEALED:已处理，
	// EVENT_IGNORE:忽略，
	// EVENT_DELETE：已删除
	// EVENT_ADD_WHITE：加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 加白镜像ID数组
	ImageIDs []*string `json:"ImageIDs,omitempty" name:"ImageIDs"`

	// 加白事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸
	EventType []*string `json:"EventType,omitempty" name:"EventType"`
}

type ModifyEscapeEventStatusRequest struct {
	*tchttp.BaseRequest
	
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态：
	// EVENT_UNDEAL:未处理（取消忽略），
	// EVENT_DEALED:已处理，
	// EVENT_IGNORE:忽略，
	// EVENT_DELETE：已删除
	// EVENT_ADD_WHITE：加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 加白镜像ID数组
	ImageIDs []*string `json:"ImageIDs,omitempty" name:"ImageIDs"`

	// 加白事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸
	EventType []*string `json:"EventType,omitempty" name:"EventType"`
}

func (r *ModifyEscapeEventStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEscapeEventStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIdSet")
	delete(f, "Status")
	delete(f, "Remark")
	delete(f, "ImageIDs")
	delete(f, "EventType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEscapeEventStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEscapeEventStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEscapeEventStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEscapeEventStatusResponseParams `json:"Response"`
}

func (r *ModifyEscapeEventStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEscapeEventStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEscapeRuleRequestParams struct {
	// 需要修改的数组
	RuleSet []*EscapeRuleEnabled `json:"RuleSet,omitempty" name:"RuleSet"`
}

type ModifyEscapeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的数组
	RuleSet []*EscapeRuleEnabled `json:"RuleSet,omitempty" name:"RuleSet"`
}

func (r *ModifyEscapeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEscapeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEscapeRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEscapeRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEscapeRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEscapeRuleResponseParams `json:"Response"`
}

func (r *ModifyEscapeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEscapeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEscapeWhiteListRequestParams struct {
	// 加白名单事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸
	EventType []*string `json:"EventType,omitempty" name:"EventType"`

	// 白名单记录ID
	IDSet []*int64 `json:"IDSet,omitempty" name:"IDSet"`
}

type ModifyEscapeWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 加白名单事件类型
	//    ESCAPE_CGROUPS：利用cgroup机制逃逸
	//    ESCAPE_TAMPER_SENSITIVE_FILE：篡改敏感文件逃逸
	//    ESCAPE_DOCKER_API：访问Docker API接口逃逸
	//    ESCAPE_VUL_OCCURRED：逃逸漏洞利用
	//    MOUNT_SENSITIVE_PTAH：敏感路径挂载
	//    PRIVILEGE_CONTAINER_START：特权容器
	//    PRIVILEGE：程序提权逃逸
	EventType []*string `json:"EventType,omitempty" name:"EventType"`

	// 白名单记录ID
	IDSet []*int64 `json:"IDSet,omitempty" name:"IDSet"`
}

func (r *ModifyEscapeWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEscapeWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventType")
	delete(f, "IDSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEscapeWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEscapeWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEscapeWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEscapeWhiteListResponseParams `json:"Response"`
}

func (r *ModifyEscapeWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEscapeWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIgnoreVul struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 忽略的镜像ID，空表示全部
	ImageIDs []*string `json:"ImageIDs,omitempty" name:"ImageIDs"`

	// 当有镜像时
	// 镜像类型: LOCAL 本地镜像 REGISTRY 仓库镜像
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
}

// Predefined struct for user
type ModifyImageAuthorizedRequestParams struct {
	// 本地镜像是否全部授权的标识，优先权高于根据本地镜像ids授权。等于true时需UpdatedLocalImageCnt大于0。
	AllLocalImages *bool `json:"AllLocalImages,omitempty" name:"AllLocalImages"`

	// 仓库镜像是否全部授权的标识，优先权高于根据镜像ids授权。等于true时需UpdatedRegistryImageCnt大于0。
	AllRegistryImages *bool `json:"AllRegistryImages,omitempty" name:"AllRegistryImages"`

	// 指定操作授权的本地镜像数量，判断优先权最高，实际多出的镜像随机忽略，实际不足的部分也忽略。
	UpdatedLocalImageCnt *uint64 `json:"UpdatedLocalImageCnt,omitempty" name:"UpdatedLocalImageCnt"`

	// 指定操作授权的仓库镜像数量，判断优先权最高，实际多出的镜像随机忽略，实际不足的部分也忽略；
	UpdatedRegistryImageCnt *uint64 `json:"UpdatedRegistryImageCnt,omitempty" name:"UpdatedRegistryImageCnt"`

	// 根据满足条件的本地镜像授权,本地镜像来源；ASSETIMAGE:本地镜像列表；IMAGEALL:同步本地镜像；AllLocalImages为false且LocalImageIds为空和UpdatedLocalImageCnt大于0时，需要
	ImageSourceType *string `json:"ImageSourceType,omitempty" name:"ImageSourceType"`

	// 根据满足条件的本地镜像授权，AllLocalImages为false且LocalImageIds为空和UpdatedLocalImageCnt大于0时，需要。
	LocalImageFilter []*AssetFilters `json:"LocalImageFilter,omitempty" name:"LocalImageFilter"`

	// 根据满足条件的仓库镜像授权，AllRegistryImages为false且RegistryImageIds为空和UpdatedRegistryImageCnt大于0时，需要。
	RegistryImageFilter []*AssetFilters `json:"RegistryImageFilter,omitempty" name:"RegistryImageFilter"`

	// 根据满足条件的镜像授权,同时排除的本地镜像。
	ExcludeLocalImageIds []*string `json:"ExcludeLocalImageIds,omitempty" name:"ExcludeLocalImageIds"`

	// 根据满足条件的镜像授权,同时排除的仓库镜像。
	ExcludeRegistryImageIds []*string `json:"ExcludeRegistryImageIds,omitempty" name:"ExcludeRegistryImageIds"`

	// 根据本地镜像ids授权，优先权高于根据满足条件的镜像授权。AllLocalImages为false且LocalImageFilter为空和UpdatedLocalImageCnt大于0时，需要。
	LocalImageIds []*string `json:"LocalImageIds,omitempty" name:"LocalImageIds"`

	// 根据仓库镜像Ids授权，优先权高于根据满足条件的镜像授。AllRegistryImages为false且RegistryImageFilter为空和UpdatedRegistryImageCnt大于0时，需要。
	RegistryImageIds []*string `json:"RegistryImageIds,omitempty" name:"RegistryImageIds"`

	// 是否仅最新的镜像；RegistryImageFilter不为空且UpdatedRegistryImageCnt大于0时仓库镜像需要。
	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
}

type ModifyImageAuthorizedRequest struct {
	*tchttp.BaseRequest
	
	// 本地镜像是否全部授权的标识，优先权高于根据本地镜像ids授权。等于true时需UpdatedLocalImageCnt大于0。
	AllLocalImages *bool `json:"AllLocalImages,omitempty" name:"AllLocalImages"`

	// 仓库镜像是否全部授权的标识，优先权高于根据镜像ids授权。等于true时需UpdatedRegistryImageCnt大于0。
	AllRegistryImages *bool `json:"AllRegistryImages,omitempty" name:"AllRegistryImages"`

	// 指定操作授权的本地镜像数量，判断优先权最高，实际多出的镜像随机忽略，实际不足的部分也忽略。
	UpdatedLocalImageCnt *uint64 `json:"UpdatedLocalImageCnt,omitempty" name:"UpdatedLocalImageCnt"`

	// 指定操作授权的仓库镜像数量，判断优先权最高，实际多出的镜像随机忽略，实际不足的部分也忽略；
	UpdatedRegistryImageCnt *uint64 `json:"UpdatedRegistryImageCnt,omitempty" name:"UpdatedRegistryImageCnt"`

	// 根据满足条件的本地镜像授权,本地镜像来源；ASSETIMAGE:本地镜像列表；IMAGEALL:同步本地镜像；AllLocalImages为false且LocalImageIds为空和UpdatedLocalImageCnt大于0时，需要
	ImageSourceType *string `json:"ImageSourceType,omitempty" name:"ImageSourceType"`

	// 根据满足条件的本地镜像授权，AllLocalImages为false且LocalImageIds为空和UpdatedLocalImageCnt大于0时，需要。
	LocalImageFilter []*AssetFilters `json:"LocalImageFilter,omitempty" name:"LocalImageFilter"`

	// 根据满足条件的仓库镜像授权，AllRegistryImages为false且RegistryImageIds为空和UpdatedRegistryImageCnt大于0时，需要。
	RegistryImageFilter []*AssetFilters `json:"RegistryImageFilter,omitempty" name:"RegistryImageFilter"`

	// 根据满足条件的镜像授权,同时排除的本地镜像。
	ExcludeLocalImageIds []*string `json:"ExcludeLocalImageIds,omitempty" name:"ExcludeLocalImageIds"`

	// 根据满足条件的镜像授权,同时排除的仓库镜像。
	ExcludeRegistryImageIds []*string `json:"ExcludeRegistryImageIds,omitempty" name:"ExcludeRegistryImageIds"`

	// 根据本地镜像ids授权，优先权高于根据满足条件的镜像授权。AllLocalImages为false且LocalImageFilter为空和UpdatedLocalImageCnt大于0时，需要。
	LocalImageIds []*string `json:"LocalImageIds,omitempty" name:"LocalImageIds"`

	// 根据仓库镜像Ids授权，优先权高于根据满足条件的镜像授。AllRegistryImages为false且RegistryImageFilter为空和UpdatedRegistryImageCnt大于0时，需要。
	RegistryImageIds []*string `json:"RegistryImageIds,omitempty" name:"RegistryImageIds"`

	// 是否仅最新的镜像；RegistryImageFilter不为空且UpdatedRegistryImageCnt大于0时仓库镜像需要。
	OnlyShowLatest *bool `json:"OnlyShowLatest,omitempty" name:"OnlyShowLatest"`
}

func (r *ModifyImageAuthorizedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAuthorizedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AllLocalImages")
	delete(f, "AllRegistryImages")
	delete(f, "UpdatedLocalImageCnt")
	delete(f, "UpdatedRegistryImageCnt")
	delete(f, "ImageSourceType")
	delete(f, "LocalImageFilter")
	delete(f, "RegistryImageFilter")
	delete(f, "ExcludeLocalImageIds")
	delete(f, "ExcludeRegistryImageIds")
	delete(f, "LocalImageIds")
	delete(f, "RegistryImageIds")
	delete(f, "OnlyShowLatest")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageAuthorizedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAuthorizedResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyImageAuthorizedResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageAuthorizedResponseParams `json:"Response"`
}

func (r *ModifyImageAuthorizedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAuthorizedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyK8sApiAbnormalEventStatusRequestParams struct {
	// 事件ID集合
	EventIDSet []*uint64 `json:"EventIDSet,omitempty" name:"EventIDSet"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyK8sApiAbnormalEventStatusRequest struct {
	*tchttp.BaseRequest
	
	// 事件ID集合
	EventIDSet []*uint64 `json:"EventIDSet,omitempty" name:"EventIDSet"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyK8sApiAbnormalEventStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyK8sApiAbnormalEventStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIDSet")
	delete(f, "Status")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyK8sApiAbnormalEventStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyK8sApiAbnormalEventStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyK8sApiAbnormalEventStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyK8sApiAbnormalEventStatusResponseParams `json:"Response"`
}

func (r *ModifyK8sApiAbnormalEventStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyK8sApiAbnormalEventStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyK8sApiAbnormalRuleInfoRequestParams struct {
	// 规则详情
	RuleInfo *K8sApiAbnormalRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
}

type ModifyK8sApiAbnormalRuleInfoRequest struct {
	*tchttp.BaseRequest
	
	// 规则详情
	RuleInfo *K8sApiAbnormalRuleInfo `json:"RuleInfo,omitempty" name:"RuleInfo"`
}

func (r *ModifyK8sApiAbnormalRuleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyK8sApiAbnormalRuleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyK8sApiAbnormalRuleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyK8sApiAbnormalRuleInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyK8sApiAbnormalRuleInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyK8sApiAbnormalRuleInfoResponseParams `json:"Response"`
}

func (r *ModifyK8sApiAbnormalRuleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyK8sApiAbnormalRuleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyK8sApiAbnormalRuleStatusRequestParams struct {
	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`

	// 状态(true:开 false:关)
	Status *bool `json:"Status,omitempty" name:"Status"`
}

type ModifyK8sApiAbnormalRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`

	// 状态(true:开 false:关)
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyK8sApiAbnormalRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyK8sApiAbnormalRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyK8sApiAbnormalRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyK8sApiAbnormalRuleStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyK8sApiAbnormalRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyK8sApiAbnormalRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyK8sApiAbnormalRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyK8sApiAbnormalRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReverseShellStatusRequestParams struct {
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，   
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyReverseShellStatusRequest struct {
	*tchttp.BaseRequest
	
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，   
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyReverseShellStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReverseShellStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIdSet")
	delete(f, "Status")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyReverseShellStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReverseShellStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyReverseShellStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyReverseShellStatusResponseParams `json:"Response"`
}

func (r *ModifyReverseShellStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReverseShellStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskSyscallStatusRequestParams struct {
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，   
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyRiskSyscallStatusRequest struct {
	*tchttp.BaseRequest
	
	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，   
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyRiskSyscallStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskSyscallStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIdSet")
	delete(f, "Status")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRiskSyscallStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskSyscallStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRiskSyscallStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRiskSyscallStatusResponseParams `json:"Response"`
}

func (r *ModifyRiskSyscallStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskSyscallStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogCleanSettingInfoRequestParams struct {
	// 触发清理的储量底线(50-99)
	ReservesLimit *uint64 `json:"ReservesLimit,omitempty" name:"ReservesLimit"`

	// 清理停止时的储量截至线(>0,小于ReservesLimit)
	ReservesDeadline *uint64 `json:"ReservesDeadline,omitempty" name:"ReservesDeadline"`

	// 触发清理的存储天数(>=1)
	DayLimit *uint64 `json:"DayLimit,omitempty" name:"DayLimit"`
}

type ModifySecLogCleanSettingInfoRequest struct {
	*tchttp.BaseRequest
	
	// 触发清理的储量底线(50-99)
	ReservesLimit *uint64 `json:"ReservesLimit,omitempty" name:"ReservesLimit"`

	// 清理停止时的储量截至线(>0,小于ReservesLimit)
	ReservesDeadline *uint64 `json:"ReservesDeadline,omitempty" name:"ReservesDeadline"`

	// 触发清理的存储天数(>=1)
	DayLimit *uint64 `json:"DayLimit,omitempty" name:"DayLimit"`
}

func (r *ModifySecLogCleanSettingInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogCleanSettingInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReservesLimit")
	delete(f, "ReservesDeadline")
	delete(f, "DayLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecLogCleanSettingInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogCleanSettingInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecLogCleanSettingInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecLogCleanSettingInfoResponseParams `json:"Response"`
}

func (r *ModifySecLogCleanSettingInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogCleanSettingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogDeliveryClsSettingRequestParams struct {
	// 日志信息
	List []*SecLogDeliveryClsSettingInfo `json:"List,omitempty" name:"List"`
}

type ModifySecLogDeliveryClsSettingRequest struct {
	*tchttp.BaseRequest
	
	// 日志信息
	List []*SecLogDeliveryClsSettingInfo `json:"List,omitempty" name:"List"`
}

func (r *ModifySecLogDeliveryClsSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogDeliveryClsSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "List")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecLogDeliveryClsSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogDeliveryClsSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecLogDeliveryClsSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecLogDeliveryClsSettingResponseParams `json:"Response"`
}

func (r *ModifySecLogDeliveryClsSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogDeliveryClsSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogDeliveryKafkaSettingRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 用户名
	User *string `json:"User,omitempty" name:"User"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 日志类型队列
	LogTypeList []*SecLogDeliveryKafkaSettingInfo `json:"LogTypeList,omitempty" name:"LogTypeList"`

	// 接入类型
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// kafka版本号
	KafkaVersion *string `json:"KafkaVersion,omitempty" name:"KafkaVersion"`

	// 地域ID
	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
}

type ModifySecLogDeliveryKafkaSettingRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 用户名
	User *string `json:"User,omitempty" name:"User"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 日志类型队列
	LogTypeList []*SecLogDeliveryKafkaSettingInfo `json:"LogTypeList,omitempty" name:"LogTypeList"`

	// 接入类型
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// kafka版本号
	KafkaVersion *string `json:"KafkaVersion,omitempty" name:"KafkaVersion"`

	// 地域ID
	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
}

func (r *ModifySecLogDeliveryKafkaSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogDeliveryKafkaSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "InstanceName")
	delete(f, "Domain")
	delete(f, "User")
	delete(f, "Password")
	delete(f, "LogTypeList")
	delete(f, "AccessType")
	delete(f, "KafkaVersion")
	delete(f, "RegionID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecLogDeliveryKafkaSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogDeliveryKafkaSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecLogDeliveryKafkaSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecLogDeliveryKafkaSettingResponseParams `json:"Response"`
}

func (r *ModifySecLogDeliveryKafkaSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogDeliveryKafkaSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogJoinObjectsRequestParams struct {
	// 日志类型
	// bash日志: container_bash
	// 容器启动: container_launch
	// k8sApi: k8s_api
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 绑定主机quuid列表
	BindList []*string `json:"BindList,omitempty" name:"BindList"`

	// 待解绑主机quuid列表
	UnBindList []*string `json:"UnBindList,omitempty" name:"UnBindList"`
}

type ModifySecLogJoinObjectsRequest struct {
	*tchttp.BaseRequest
	
	// 日志类型
	// bash日志: container_bash
	// 容器启动: container_launch
	// k8sApi: k8s_api
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 绑定主机quuid列表
	BindList []*string `json:"BindList,omitempty" name:"BindList"`

	// 待解绑主机quuid列表
	UnBindList []*string `json:"UnBindList,omitempty" name:"UnBindList"`
}

func (r *ModifySecLogJoinObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogJoinObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogType")
	delete(f, "BindList")
	delete(f, "UnBindList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecLogJoinObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogJoinObjectsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecLogJoinObjectsResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecLogJoinObjectsResponseParams `json:"Response"`
}

func (r *ModifySecLogJoinObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogJoinObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogJoinStateRequestParams struct {
	// 日志类型
	// bash日志: container_bash
	// 容器启动: container_launch
	// k8sApi: k8s_api
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 状态(true:开 false:关)
	State *bool `json:"State,omitempty" name:"State"`
}

type ModifySecLogJoinStateRequest struct {
	*tchttp.BaseRequest
	
	// 日志类型
	// bash日志: container_bash
	// 容器启动: container_launch
	// k8sApi: k8s_api
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 状态(true:开 false:关)
	State *bool `json:"State,omitempty" name:"State"`
}

func (r *ModifySecLogJoinStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogJoinStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogType")
	delete(f, "State")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecLogJoinStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogJoinStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecLogJoinStateResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecLogJoinStateResponseParams `json:"Response"`
}

func (r *ModifySecLogJoinStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogJoinStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogKafkaUINRequestParams struct {
	// 目标UIN
	DstUIN *string `json:"DstUIN,omitempty" name:"DstUIN"`
}

type ModifySecLogKafkaUINRequest struct {
	*tchttp.BaseRequest
	
	// 目标UIN
	DstUIN *string `json:"DstUIN,omitempty" name:"DstUIN"`
}

func (r *ModifySecLogKafkaUINRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogKafkaUINRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DstUIN")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecLogKafkaUINRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecLogKafkaUINResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecLogKafkaUINResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecLogKafkaUINResponseParams `json:"Response"`
}

func (r *ModifySecLogKafkaUINResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecLogKafkaUINResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusAutoIsolateExampleSwitchRequestParams struct {
	// 文件Md5值
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// 开关(开:true 关: false)
	Status *bool `json:"Status,omitempty" name:"Status"`
}

type ModifyVirusAutoIsolateExampleSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 文件Md5值
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// 开关(开:true 关: false)
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyVirusAutoIsolateExampleSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusAutoIsolateExampleSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MD5")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVirusAutoIsolateExampleSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusAutoIsolateExampleSwitchResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVirusAutoIsolateExampleSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVirusAutoIsolateExampleSwitchResponseParams `json:"Response"`
}

func (r *ModifyVirusAutoIsolateExampleSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusAutoIsolateExampleSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusAutoIsolateSettingRequestParams struct {
	// 自动隔离开关(true:开 false:关)
	AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`

	// 是否中断隔离文件关联的进程(true:是 false:否)
	IsKillProgress *bool `json:"IsKillProgress,omitempty" name:"IsKillProgress"`
}

type ModifyVirusAutoIsolateSettingRequest struct {
	*tchttp.BaseRequest
	
	// 自动隔离开关(true:开 false:关)
	AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`

	// 是否中断隔离文件关联的进程(true:是 false:否)
	IsKillProgress *bool `json:"IsKillProgress,omitempty" name:"IsKillProgress"`
}

func (r *ModifyVirusAutoIsolateSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusAutoIsolateSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoIsolateSwitch")
	delete(f, "IsKillProgress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVirusAutoIsolateSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusAutoIsolateSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVirusAutoIsolateSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVirusAutoIsolateSettingResponseParams `json:"Response"`
}

func (r *ModifyVirusAutoIsolateSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusAutoIsolateSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusFileStatusRequestParams struct {
	// 处理事件id
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，   
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//     EVENT_DEL:事件删除
	//     EVENT_ADD_WHITE:事件加白
	//     EVENT_PENDING: 事件待处理
	// 	EVENT_ISOLATE_CONTAINER: 隔离容器
	// 	EVENT_RESOTRE_CONTAINER: 恢复容器
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 是否后续自动隔离相同MD5文件
	AutoIsolate *bool `json:"AutoIsolate,omitempty" name:"AutoIsolate"`
}

type ModifyVirusFileStatusRequest struct {
	*tchttp.BaseRequest
	
	// 处理事件id
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，   
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//     EVENT_DEL:事件删除
	//     EVENT_ADD_WHITE:事件加白
	//     EVENT_PENDING: 事件待处理
	// 	EVENT_ISOLATE_CONTAINER: 隔离容器
	// 	EVENT_RESOTRE_CONTAINER: 恢复容器
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 是否后续自动隔离相同MD5文件
	AutoIsolate *bool `json:"AutoIsolate,omitempty" name:"AutoIsolate"`
}

func (r *ModifyVirusFileStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusFileStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIdSet")
	delete(f, "Status")
	delete(f, "Remark")
	delete(f, "AutoIsolate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVirusFileStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusFileStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVirusFileStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVirusFileStatusResponseParams `json:"Response"`
}

func (r *ModifyVirusFileStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusFileStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusMonitorSettingRequestParams struct {
	// 是否开启定期扫描
	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`

	// 扫描全部路径
	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`

	// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径(扫描范围只能小于等于1)
	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`

	// 自选排除或扫描的地址
	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
}

type ModifyVirusMonitorSettingRequest struct {
	*tchttp.BaseRequest
	
	// 是否开启定期扫描
	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`

	// 扫描全部路径
	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`

	// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径(扫描范围只能小于等于1)
	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`

	// 自选排除或扫描的地址
	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
}

func (r *ModifyVirusMonitorSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusMonitorSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnableScan")
	delete(f, "ScanPathAll")
	delete(f, "ScanPathType")
	delete(f, "ScanPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVirusMonitorSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusMonitorSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVirusMonitorSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVirusMonitorSettingResponseParams `json:"Response"`
}

func (r *ModifyVirusMonitorSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusMonitorSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusScanSettingRequestParams struct {
	// 是否开启定期扫描
	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`

	// 检测周期每隔多少天(1|3|7)
	Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`

	// 扫描开始时间
	BeginScanAt *string `json:"BeginScanAt,omitempty" name:"BeginScanAt"`

	// 扫描全部路径(true:全选,false:自选)
	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`

	// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径
	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`

	// 超时时长(5~24h)
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 扫描范围0容器1主机节点
	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`

	// true 全选，false 自选
	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`

	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定
	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`

	// 扫描路径
	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
}

type ModifyVirusScanSettingRequest struct {
	*tchttp.BaseRequest
	
	// 是否开启定期扫描
	EnableScan *bool `json:"EnableScan,omitempty" name:"EnableScan"`

	// 检测周期每隔多少天(1|3|7)
	Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`

	// 扫描开始时间
	BeginScanAt *string `json:"BeginScanAt,omitempty" name:"BeginScanAt"`

	// 扫描全部路径(true:全选,false:自选)
	ScanPathAll *bool `json:"ScanPathAll,omitempty" name:"ScanPathAll"`

	// 当ScanPathAll为true 生效 0扫描以下路径 1、扫描除以下路径
	ScanPathType *uint64 `json:"ScanPathType,omitempty" name:"ScanPathType"`

	// 超时时长(5~24h)
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 扫描范围0容器1主机节点
	ScanRangeType *uint64 `json:"ScanRangeType,omitempty" name:"ScanRangeType"`

	// true 全选，false 自选
	ScanRangeAll *bool `json:"ScanRangeAll,omitempty" name:"ScanRangeAll"`

	// 自选扫描范围的容器id或者主机id 根据ScanRangeType决定
	ScanIds []*string `json:"ScanIds,omitempty" name:"ScanIds"`

	// 扫描路径
	ScanPath []*string `json:"ScanPath,omitempty" name:"ScanPath"`
}

func (r *ModifyVirusScanSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusScanSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnableScan")
	delete(f, "Cycle")
	delete(f, "BeginScanAt")
	delete(f, "ScanPathAll")
	delete(f, "ScanPathType")
	delete(f, "Timeout")
	delete(f, "ScanRangeType")
	delete(f, "ScanRangeAll")
	delete(f, "ScanIds")
	delete(f, "ScanPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVirusScanSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusScanSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVirusScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVirusScanSettingResponseParams `json:"Response"`
}

func (r *ModifyVirusScanSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusScanTimeoutSettingRequestParams struct {
	// 超时时长单位小时(5~24h)
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 设置类型0一键检测，1定时检测
	ScanType *uint64 `json:"ScanType,omitempty" name:"ScanType"`
}

type ModifyVirusScanTimeoutSettingRequest struct {
	*tchttp.BaseRequest
	
	// 超时时长单位小时(5~24h)
	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

	// 设置类型0一键检测，1定时检测
	ScanType *uint64 `json:"ScanType,omitempty" name:"ScanType"`
}

func (r *ModifyVirusScanTimeoutSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusScanTimeoutSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Timeout")
	delete(f, "ScanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVirusScanTimeoutSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVirusScanTimeoutSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVirusScanTimeoutSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVirusScanTimeoutSettingResponseParams `json:"Response"`
}

func (r *ModifyVirusScanTimeoutSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVirusScanTimeoutSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVulDefenceEventStatusRequestParams struct {
	// 事件IDs数组
	EventIDs []*int64 `json:"EventIDs,omitempty" name:"EventIDs"`

	// 操作状态：
	// EVENT_DEALED：已处理，EVENT_IGNORE：忽略，EVENT_ISOLATE_CONTAINER：隔离容器，EVENT_DEL：删除
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyVulDefenceEventStatusRequest struct {
	*tchttp.BaseRequest
	
	// 事件IDs数组
	EventIDs []*int64 `json:"EventIDs,omitempty" name:"EventIDs"`

	// 操作状态：
	// EVENT_DEALED：已处理，EVENT_IGNORE：忽略，EVENT_ISOLATE_CONTAINER：隔离容器，EVENT_DEL：删除
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyVulDefenceEventStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVulDefenceEventStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIDs")
	delete(f, "Status")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVulDefenceEventStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVulDefenceEventStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVulDefenceEventStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVulDefenceEventStatusResponseParams `json:"Response"`
}

func (r *ModifyVulDefenceEventStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVulDefenceEventStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVulDefenceSettingRequestParams struct {
	// 是否开启:0: 关闭 1:开启
	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 漏洞防御主机范围:0：自选 1: 全部主机。IsEnabled为1时必填
	Scope *int64 `json:"Scope,omitempty" name:"Scope"`

	// 自选漏洞防御主机,Scope为0时必填
	HostIDs []*string `json:"HostIDs,omitempty" name:"HostIDs"`
}

type ModifyVulDefenceSettingRequest struct {
	*tchttp.BaseRequest
	
	// 是否开启:0: 关闭 1:开启
	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 漏洞防御主机范围:0：自选 1: 全部主机。IsEnabled为1时必填
	Scope *int64 `json:"Scope,omitempty" name:"Scope"`

	// 自选漏洞防御主机,Scope为0时必填
	HostIDs []*string `json:"HostIDs,omitempty" name:"HostIDs"`
}

func (r *ModifyVulDefenceSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVulDefenceSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsEnabled")
	delete(f, "Scope")
	delete(f, "HostIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVulDefenceSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVulDefenceSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVulDefenceSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVulDefenceSettingResponseParams `json:"Response"`
}

func (r *ModifyVulDefenceSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVulDefenceSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkAuditRecord struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名字
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 动作
	Action *string `json:"Action,omitempty" name:"Action"`

	// 操作人
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 策略名
	NetworkPolicyName *string `json:"NetworkPolicyName,omitempty" name:"NetworkPolicyName"`

	// 操作时间
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`

	// 操作人appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 操作人uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type NetworkClusterInfoItem struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名字
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 集群操作系统
	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 集群网络插件
	NetworkPolicyPlugin *string `json:"NetworkPolicyPlugin,omitempty" name:"NetworkPolicyPlugin"`

	// 集群状态
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 总策略数量
	TotalRuleCount *int64 `json:"TotalRuleCount,omitempty" name:"TotalRuleCount"`

	// 已开启策略数量
	EnableRuleCount *int64 `json:"EnableRuleCount,omitempty" name:"EnableRuleCount"`

	// 集群网络插件状态，正常：Running 不正常：Error
	NetworkPolicyPluginStatus *string `json:"NetworkPolicyPluginStatus,omitempty" name:"NetworkPolicyPluginStatus"`

	// 集群网络插件错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkPolicyPluginError *string `json:"NetworkPolicyPluginError,omitempty" name:"NetworkPolicyPluginError"`
}

type NetworkClusterNamespaceInfo struct {
	// 网络空间标签
	Labels *string `json:"Labels,omitempty" name:"Labels"`

	// 网络空间名字
	Name *string `json:"Name,omitempty" name:"Name"`
}

type NetworkClusterNamespaceLabelInfo struct {
	// 网络空间标签
	Labels *string `json:"Labels,omitempty" name:"Labels"`

	// 网络空间名字
	Name *string `json:"Name,omitempty" name:"Name"`
}

type NetworkClusterPodInfo struct {
	// pod名字
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// pod空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// pod标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels *string `json:"Labels,omitempty" name:"Labels"`

	// pod类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadKind *string `json:"WorkloadKind,omitempty" name:"WorkloadKind"`
}

type NetworkCustomPolicy struct {
	// 网络策略方向，分为FROM和TO
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// 网络策略策略端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ports []*NetworkPorts `json:"Ports,omitempty" name:"Ports"`

	// 网络策略策略对象
	// 
	// 开启待确认：PublishedNoConfirm
	// 
	// 开启已确认：PublishedConfirmed
	// 
	// 关闭中：unPublishing
	// 
	// 开启中：Publishing
	// 
	// 待开启：unPublishEdit
	// 注意：此字段可能返回 null，表示取不到有效值。
	Peer []*NetworkPeer `json:"Peer,omitempty" name:"Peer"`
}

type NetworkPeer struct {
	// 对象类型：
	// 
	// 命名空间：NamespaceSelector，代表NamespaceSelector有值
	// 
	// pod类型：PodSelector，代表NamespaceSelector和PodSelector都有值
	// 
	// ip类型：IPBlock，代表只有IPBlock有值
	PeerType *string `json:"PeerType,omitempty" name:"PeerType"`

	// 空间选择器
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceSelector *string `json:"NamespaceSelector,omitempty" name:"NamespaceSelector"`

	// pod选择器
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// Ip选择器
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPBlock *string `json:"IPBlock,omitempty" name:"IPBlock"`
}

type NetworkPolicyInfoItem struct {
	// 网络策略名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网络策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 发布状态：
	// 
	// 开启待确认：PublishedNoConfirm
	// 
	// 开启已确认：PublishedConfirmed
	// 
	// 关闭中：unPublishing
	// 
	// 开启中：Publishing
	// 
	// 待开启：unPublishEdit
	PublishStatus *string `json:"PublishStatus,omitempty" name:"PublishStatus"`

	// 策略类型：
	// 
	// 自动发现：System
	// 
	// 手动添加：Manual
	PolicySourceType *string `json:"PolicySourceType,omitempty" name:"PolicySourceType"`

	// 策略空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 策略创建日期
	PolicyCreateTime *string `json:"PolicyCreateTime,omitempty" name:"PolicyCreateTime"`

	// 策略类型
	// 
	// kube-router：KubeRouter
	// 
	// cilium：Cilium
	NetworkPolicyPlugin *string `json:"NetworkPolicyPlugin,omitempty" name:"NetworkPolicyPlugin"`

	// 策略发布结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishResult *string `json:"PublishResult,omitempty" name:"PublishResult"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// 作用对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 网络策略Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type NetworkPorts struct {
	// 网络策略协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 网络策略策略端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`
}

// Predefined struct for user
type OpenTcssTrialRequestParams struct {

}

type OpenTcssTrialRequest struct {
	*tchttp.BaseRequest
	
}

func (r *OpenTcssTrialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenTcssTrialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenTcssTrialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenTcssTrialResponseParams struct {
	// 试用开通结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 试用开通开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenTcssTrialResponse struct {
	*tchttp.BaseResponse
	Response *OpenTcssTrialResponseParams `json:"Response"`
}

func (r *OpenTcssTrialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenTcssTrialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PortInfo struct {
	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 对外ip
	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`

	// 主机端口
	PublicPort *uint64 `json:"PublicPort,omitempty" name:"PublicPort"`

	// 容器端口
	ContainerPort *uint64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 容器Pid
	ContainerPID *uint64 `json:"ContainerPID,omitempty" name:"ContainerPID"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 主机id
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 主机ip
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 容器内监听地址
	ListenContainer *string `json:"ListenContainer,omitempty" name:"ListenContainer"`

	// 容器外监听地址
	ListenHost *string `json:"ListenHost,omitempty" name:"ListenHost"`

	// 运行账号
	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`

	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 外网ip
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
}

type ProcessBaseInfo struct {
	// 进程启动用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`

	// 进程用户组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`

	// 进程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 进程命令行参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
}

type ProcessDetailBaseInfo struct {
	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程pid
	ProcessId *uint64 `json:"ProcessId,omitempty" name:"ProcessId"`

	// 进程启动用户
	ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`

	// 进程用户组
	ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`

	// 进程路径
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 进程命令行参数
	ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
}

type ProcessDetailInfo struct {
	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程权限
	ProcessAuthority *string `json:"ProcessAuthority,omitempty" name:"ProcessAuthority"`

	// 进程pid
	ProcessId *uint64 `json:"ProcessId,omitempty" name:"ProcessId"`

	// 进程启动用户
	ProcessStartUser *string `json:"ProcessStartUser,omitempty" name:"ProcessStartUser"`

	// 进程用户组
	ProcessUserGroup *string `json:"ProcessUserGroup,omitempty" name:"ProcessUserGroup"`

	// 进程路径
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 进程树
	ProcessTree *string `json:"ProcessTree,omitempty" name:"ProcessTree"`

	// 进程md5
	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`

	// 进程命令行参数
	ProcessParam *string `json:"ProcessParam,omitempty" name:"ProcessParam"`
}

type ProcessInfo struct {
	// 进程启动时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 运行用户
	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`

	// 命令行参数
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// Exe路径
	Exe *string `json:"Exe,omitempty" name:"Exe"`

	// 主机PID
	PID *uint64 `json:"PID,omitempty" name:"PID"`

	// 容器内pid
	ContainerPID *uint64 `json:"ContainerPID,omitempty" name:"ContainerPID"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 主机id
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 主机ip
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 外网ip
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
}

type ProjectInfo struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`
}

type PromotionActivityContent struct {
	// 月数
	MonthNum *uint64 `json:"MonthNum,omitempty" name:"MonthNum"`

	// 核数最低限量
	CoresCountLimit *uint64 `json:"CoresCountLimit,omitempty" name:"CoresCountLimit"`

	// 专业版折扣
	ProfessionalDiscount *uint64 `json:"ProfessionalDiscount,omitempty" name:"ProfessionalDiscount"`

	// 附赠镜像数
	ImageAuthorizationNum *uint64 `json:"ImageAuthorizationNum,omitempty" name:"ImageAuthorizationNum"`
}

type RaspInfo struct {
	// rasp名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// rasp  描述
	Value *string `json:"Value,omitempty" name:"Value"`
}

type RegionInfo struct {
	// 地域标识
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

// Predefined struct for user
type RemoveAssetImageRegistryRegistryDetailRequestParams struct {
	// 仓库唯一id
	RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
}

type RemoveAssetImageRegistryRegistryDetailRequest struct {
	*tchttp.BaseRequest
	
	// 仓库唯一id
	RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *RemoveAssetImageRegistryRegistryDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAssetImageRegistryRegistryDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveAssetImageRegistryRegistryDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveAssetImageRegistryRegistryDetailResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *RemoveAssetImageRegistryRegistryDetailResponseParams `json:"Response"`
}

func (r *RemoveAssetImageRegistryRegistryDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAssetImageRegistryRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewImageAuthorizeStateRequestParams struct {
	// 是否全部未授权镜像
	AllImages *bool `json:"AllImages,omitempty" name:"AllImages"`

	// 镜像ids
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

type RenewImageAuthorizeStateRequest struct {
	*tchttp.BaseRequest
	
	// 是否全部未授权镜像
	AllImages *bool `json:"AllImages,omitempty" name:"AllImages"`

	// 镜像ids
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

func (r *RenewImageAuthorizeStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewImageAuthorizeStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AllImages")
	delete(f, "ImageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewImageAuthorizeStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewImageAuthorizeStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewImageAuthorizeStateResponse struct {
	*tchttp.BaseResponse
	Response *RenewImageAuthorizeStateResponseParams `json:"Response"`
}

func (r *RenewImageAuthorizeStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewImageAuthorizeStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetSecLogTopicConfigRequestParams struct {
	// 配置类型(ckafka/cls)
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// 日志类型
	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

type ResetSecLogTopicConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置类型(ckafka/cls)
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// 日志类型
	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

func (r *ResetSecLogTopicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSecLogTopicConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigType")
	delete(f, "LogType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetSecLogTopicConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetSecLogTopicConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetSecLogTopicConfigResponse struct {
	*tchttp.BaseResponse
	Response *ResetSecLogTopicConfigResponseParams `json:"Response"`
}

func (r *ResetSecLogTopicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetSecLogTopicConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReverseShellEventDescription struct {
	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 目标地址
	DstAddress *string `json:"DstAddress,omitempty" name:"DstAddress"`

	// 事件最后一次处理的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type ReverseShellEventInfo struct {
	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程路径
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 容器id
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 镜像名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 生成时间
	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`

	// 事件解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件详细描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 状态，EVENT_UNDEAL:事件未处理
	//     EVENT_DEALED:事件已经处理
	//     EVENT_INGNORE：事件已经忽略
	//     EVENT_ADD_WHITE：时间已经加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件id
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 父进程名
	PProcessName *string `json:"PProcessName,omitempty" name:"PProcessName"`

	// 事件数量
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 最近生成时间
	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`

	// 目标地址
	DstAddress *string `json:"DstAddress,omitempty" name:"DstAddress"`

	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`

	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING
	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type ReverseShellWhiteListBaseInfo struct {
	// 白名单id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 镜像数量
	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`

	// 连接进程名字
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 目标地址ip
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 目标端口
	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`

	// 是否是全局白名单，true全局
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 镜像id数组，为空代表全部
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

type ReverseShellWhiteListInfo struct {
	// 目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口
	DstPort *string `json:"DstPort,omitempty" name:"DstPort"`

	// 目标进程
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 镜像id数组，为空代表全部
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// 白名单id，如果新建则id为空
	Id *string `json:"Id,omitempty" name:"Id"`
}

type RiskSyscallEventDescription struct {
	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 系统调用名称
	SyscallName *string `json:"SyscallName,omitempty" name:"SyscallName"`

	// 事件最后一次处理的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type RiskSyscallEventInfo struct {
	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程路径
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 容器id
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 镜像名
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 生成时间
	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`

	// 事件解决方案
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 事件详细描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 系统调用名称
	SyscallName *string `json:"SyscallName,omitempty" name:"SyscallName"`

	// 状态，EVENT_UNDEAL:事件未处理
	//     EVENT_DEALED:事件已经处理
	//     EVENT_INGNORE：事件已经忽略
	//     EVENT_ADD_WHITE：时间已经加白
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件id
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// pod(实例)的名字
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 系统监控名称是否存在
	RuleExist *bool `json:"RuleExist,omitempty" name:"RuleExist"`

	// 事件数量
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 最近生成时间
	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`

	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`

	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING
	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`
}

type RiskSyscallWhiteListBaseInfo struct {
	// 白名单id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 镜像数量
	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`

	// 连接进程路径
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 系统调用名称列表
	SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 是否是全局白名单，true全局
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 镜像id数组
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`
}

type RiskSyscallWhiteListInfo struct {
	// 镜像id数组，为空代表全部
	ImageIds []*string `json:"ImageIds,omitempty" name:"ImageIds"`

	// 系统调用名称，通过DescribeRiskSyscallNames接口获取枚举列表
	SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`

	// 目标进程
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 白名单id，如果新建则id为空
	Id *string `json:"Id,omitempty" name:"Id"`
}

type RuleBaseInfo struct {
	// true: 默认策略，false:自定义策略
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 策略生效镜像数量
	EffectImageCount *uint64 `json:"EffectImageCount,omitempty" name:"EffectImageCount"`

	// 策略Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 策略更新时间, 存在为空的情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 策略名字
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 编辑用户名称
	EditUserName *string `json:"EditUserName,omitempty" name:"EditUserName"`

	// true: 策略启用，false：策略禁用
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`
}

type RunTimeEventBaseInfo struct {
	// 事件唯一ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 事件发现时间
	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`

	// 容器id
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// Pod名称
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 状态， “EVENT_UNDEAL”:事件未处理
	//     "EVENT_DEALED":事件已经处理
	//     "EVENT_INGNORE"：事件已经忽略
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件名称：
	// 宿主机文件访问逃逸、
	// Syscall逃逸、
	// MountNamespace逃逸、
	// 程序提权逃逸、
	// 特权容器启动逃逸、
	// 敏感路径挂载
	// 恶意进程启动
	// 文件篡改
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 事件类型
	//    ESCAPE_HOST_ACESS_FILE:宿主机文件访问逃逸
	//    ESCAPE_MOUNT_NAMESPACE:MountNamespace逃逸
	//    ESCAPE_PRIVILEDGE:程序提权逃逸
	//    ESCAPE_PRIVILEDGE_CONTAINER_START:特权容器启动逃逸
	//    ESCAPE_MOUNT_SENSITIVE_PTAH:敏感路径挂载
	//    ESCAPE_SYSCALL:Syscall逃逸
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 事件数量
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 最近生成时间
	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 外网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// "NODE_DESTROYED"      //节点已销毁
	// "CONTAINER_EXITED"    //容器已退出
	// "CONTAINER_DESTROYED" //容器已销毁
	// "SHARED_HOST"         // 容器与主机共享网络
	// "RESOURCE_LIMIT"      //隔离操作资源超限
	// "UNKNOW"              // 原因未知
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`
}

type RunTimeFilters struct {
	// 过滤键的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否模糊查询
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type RunTimeRiskInfo struct {
	// 数量
	Cnt *uint64 `json:"Cnt,omitempty" name:"Cnt"`

	// 风险等级：
	// CRITICAL: 严重
	// HIGH: 高
	// MEDIUM：中
	// LOW: 低
	Level *string `json:"Level,omitempty" name:"Level"`
}

type RunTimeTendencyInfo struct {
	// 当天时间
	CurTime *string `json:"CurTime,omitempty" name:"CurTime"`

	// 当前数量
	Cnt *uint64 `json:"Cnt,omitempty" name:"Cnt"`
}

// Predefined struct for user
type ScanComplianceAssetsByPolicyItemRequestParams struct {
	// 指定的检测项的ID
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 要重新扫描的客户资产项ID的列表。
	CustomerAssetIdSet []*uint64 `json:"CustomerAssetIdSet,omitempty" name:"CustomerAssetIdSet"`
}

type ScanComplianceAssetsByPolicyItemRequest struct {
	*tchttp.BaseRequest
	
	// 指定的检测项的ID
	CustomerPolicyItemId *uint64 `json:"CustomerPolicyItemId,omitempty" name:"CustomerPolicyItemId"`

	// 要重新扫描的客户资产项ID的列表。
	CustomerAssetIdSet []*uint64 `json:"CustomerAssetIdSet,omitempty" name:"CustomerAssetIdSet"`
}

func (r *ScanComplianceAssetsByPolicyItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanComplianceAssetsByPolicyItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerPolicyItemId")
	delete(f, "CustomerAssetIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanComplianceAssetsByPolicyItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanComplianceAssetsByPolicyItemResponseParams struct {
	// 返回重新检测任务的ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanComplianceAssetsByPolicyItemResponse struct {
	*tchttp.BaseResponse
	Response *ScanComplianceAssetsByPolicyItemResponseParams `json:"Response"`
}

func (r *ScanComplianceAssetsByPolicyItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanComplianceAssetsByPolicyItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanComplianceAssetsRequestParams struct {
	// 要重新扫描的客户资产项ID的列表。
	CustomerAssetIdSet []*uint64 `json:"CustomerAssetIdSet,omitempty" name:"CustomerAssetIdSet"`
}

type ScanComplianceAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 要重新扫描的客户资产项ID的列表。
	CustomerAssetIdSet []*uint64 `json:"CustomerAssetIdSet,omitempty" name:"CustomerAssetIdSet"`
}

func (r *ScanComplianceAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanComplianceAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerAssetIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanComplianceAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanComplianceAssetsResponseParams struct {
	// 返回重新检测任务的ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanComplianceAssetsResponse struct {
	*tchttp.BaseResponse
	Response *ScanComplianceAssetsResponseParams `json:"Response"`
}

func (r *ScanComplianceAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanComplianceAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanCompliancePolicyItemsRequestParams struct {
	// 要重新扫描的客户检测项的列表。
	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

type ScanCompliancePolicyItemsRequest struct {
	*tchttp.BaseRequest
	
	// 要重新扫描的客户检测项的列表。
	CustomerPolicyItemIdSet []*uint64 `json:"CustomerPolicyItemIdSet,omitempty" name:"CustomerPolicyItemIdSet"`
}

func (r *ScanCompliancePolicyItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanCompliancePolicyItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerPolicyItemIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanCompliancePolicyItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanCompliancePolicyItemsResponseParams struct {
	// 返回重新检测任务的ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanCompliancePolicyItemsResponse struct {
	*tchttp.BaseResponse
	Response *ScanCompliancePolicyItemsResponseParams `json:"Response"`
}

func (r *ScanCompliancePolicyItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanCompliancePolicyItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanComplianceScanFailedAssetsRequestParams struct {
	// 要重新扫描的客户资产项ID的列表。
	CustomerAssetIdSet []*uint64 `json:"CustomerAssetIdSet,omitempty" name:"CustomerAssetIdSet"`
}

type ScanComplianceScanFailedAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 要重新扫描的客户资产项ID的列表。
	CustomerAssetIdSet []*uint64 `json:"CustomerAssetIdSet,omitempty" name:"CustomerAssetIdSet"`
}

func (r *ScanComplianceScanFailedAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanComplianceScanFailedAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerAssetIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanComplianceScanFailedAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanComplianceScanFailedAssetsResponseParams struct {
	// 返回重新检测任务的ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanComplianceScanFailedAssetsResponse struct {
	*tchttp.BaseResponse
	Response *ScanComplianceScanFailedAssetsResponseParams `json:"Response"`
}

func (r *ScanComplianceScanFailedAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanComplianceScanFailedAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanIgnoreVul struct {
	// 漏洞名称
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞CVEID
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 忽略的仓库镜像数
	RegistryImageCount *int64 `json:"RegistryImageCount,omitempty" name:"RegistryImageCount"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 是否忽略所有镜像：0：否/1：是
	IsIgnoreAll *int64 `json:"IsIgnoreAll,omitempty" name:"IsIgnoreAll"`

	// 忽略的本地镜像数
	LocalImageCount *int64 `json:"LocalImageCount,omitempty" name:"LocalImageCount"`
}

type SearchTemplate struct {
	// 检索名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 检索索引类型
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 检索语句
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 时间范围
	TimeRange *string `json:"TimeRange,omitempty" name:"TimeRange"`

	// 转换的检索语句内容
	Query *string `json:"Query,omitempty" name:"Query"`

	// 检索方式。输入框检索：standard,过滤，检索：simple
	Flag *string `json:"Flag,omitempty" name:"Flag"`

	// 展示数据
	DisplayData *string `json:"DisplayData,omitempty" name:"DisplayData"`

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type SecLogAlertMsgInfo struct {
	// 告警类型
	MsgType *string `json:"MsgType,omitempty" name:"MsgType"`

	// 告警值
	MsgValue *string `json:"MsgValue,omitempty" name:"MsgValue"`

	// 状态(0:关闭 1:开启)
	State *bool `json:"State,omitempty" name:"State"`
}

type SecLogDeliveryClsSettingInfo struct {
	// 日志类型
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 投递状态(true:开启 false:关闭)
	State *bool `json:"State,omitempty" name:"State"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 日志集
	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`

	// 主题ID
	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`

	// 日志集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSetName *string `json:"LogSetName,omitempty" name:"LogSetName"`

	// 主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type SecLogDeliveryKafkaSettingInfo struct {
	// 日志类型
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 主题ID
	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`

	// 主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 投递状态(false:关 true:开)
	State *bool `json:"State,omitempty" name:"State"`
}

type SecLogJoinInfo struct {
	// 已接入数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 是否已接入(true:已接入 false:未接入)
	IsJoined *bool `json:"IsJoined,omitempty" name:"IsJoined"`

	// 日志类型(
	// 容器bash:  "container_bash"
	// 容器启动: "container_launch"
	// k8sApi: "k8s_api"
	// )
	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

type SecLogJoinObjectInfo struct {
	// 主机ID
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 主机名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 主机IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 主机状态
	HostStatus *string `json:"HostStatus,omitempty" name:"HostStatus"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`

	// 接入状态(true:已接入  false:未接入)
	JoinState *bool `json:"JoinState,omitempty" name:"JoinState"`

	// 集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 集群主节点地址
	ClusterMainAddress *string `json:"ClusterMainAddress,omitempty" name:"ClusterMainAddress"`
}

type SecTendencyEventInfo struct {
	// 趋势列表
	EventSet []*RunTimeTendencyInfo `json:"EventSet,omitempty" name:"EventSet"`

	// 事件类型：
	// ET_ESCAPE : 容器逃逸
	// ET_REVERSE_SHELL: 反弹shell
	// ET_RISK_SYSCALL:高危系统调用
	// ET_ABNORMAL_PROCESS: 异常进程
	// ET_ACCESS_CONTROL 文件篡改
	// ET_VIRUS 木马事件
	// ET_MALICIOUS_CONNECTION 恶意外连事件
	EventType *string `json:"EventType,omitempty" name:"EventType"`
}

type ServiceInfo struct {
	// 服务id
	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`

	// 主机id
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 主机ip
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 服务名 例如nginx/redis
	Type *string `json:"Type,omitempty" name:"Type"`

	// 版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 账号
	RunAs *string `json:"RunAs,omitempty" name:"RunAs"`

	// 监听端口
	Listen []*string `json:"Listen,omitempty" name:"Listen"`

	// 配置
	Config *string `json:"Config,omitempty" name:"Config"`

	// 关联进程数
	ProcessCnt *uint64 `json:"ProcessCnt,omitempty" name:"ProcessCnt"`

	// 访问日志
	AccessLog *string `json:"AccessLog,omitempty" name:"AccessLog"`

	// 错误日志
	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`

	// 数据目录
	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`

	// web目录
	WebRoot *string `json:"WebRoot,omitempty" name:"WebRoot"`

	// 关联的进程id
	Pids []*uint64 `json:"Pids,omitempty" name:"Pids"`

	// 服务类型 app,web,db
	MainType *string `json:"MainType,omitempty" name:"MainType"`

	// 执行文件
	Exe *string `json:"Exe,omitempty" name:"Exe"`

	// 服务命令行参数
	Parameter *string `json:"Parameter,omitempty" name:"Parameter"`

	// 容器id
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 外网ip
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`
}

// Predefined struct for user
type SetCheckModeRequestParams struct {
	// 要设置的集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 集群检查模式(正常模式"Cluster_Normal"、主动模式"Cluster_Actived"、不设置"Cluster_Unset")
	ClusterCheckMode *string `json:"ClusterCheckMode,omitempty" name:"ClusterCheckMode"`

	// 0不设置 1打开 2关闭
	ClusterAutoCheck *uint64 `json:"ClusterAutoCheck,omitempty" name:"ClusterAutoCheck"`
}

type SetCheckModeRequest struct {
	*tchttp.BaseRequest
	
	// 要设置的集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 集群检查模式(正常模式"Cluster_Normal"、主动模式"Cluster_Actived"、不设置"Cluster_Unset")
	ClusterCheckMode *string `json:"ClusterCheckMode,omitempty" name:"ClusterCheckMode"`

	// 0不设置 1打开 2关闭
	ClusterAutoCheck *uint64 `json:"ClusterAutoCheck,omitempty" name:"ClusterAutoCheck"`
}

func (r *SetCheckModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCheckModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "ClusterCheckMode")
	delete(f, "ClusterAutoCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCheckModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCheckModeResponseParams struct {
	// "Succ"表示设置成功，"Failed"表示设置失败
	SetCheckResult *string `json:"SetCheckResult,omitempty" name:"SetCheckResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetCheckModeResponse struct {
	*tchttp.BaseResponse
	Response *SetCheckModeResponseParams `json:"Response"`
}

func (r *SetCheckModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCheckModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SoftQuotaDayInfo struct {
	// 扣费时间
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 计费核数
	CoresCnt *uint64 `json:"CoresCnt,omitempty" name:"CoresCnt"`
}

// Predefined struct for user
type StopVirusScanTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 需要停止的容器id 为空默认停止整个任务
	ContainerIds []*string `json:"ContainerIds,omitempty" name:"ContainerIds"`
}

type StopVirusScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 需要停止的容器id 为空默认停止整个任务
	ContainerIds []*string `json:"ContainerIds,omitempty" name:"ContainerIds"`
}

func (r *StopVirusScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopVirusScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ContainerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopVirusScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopVirusScanTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopVirusScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopVirusScanTaskResponseParams `json:"Response"`
}

func (r *StopVirusScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopVirusScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopVulScanTaskRequestParams struct {
	// 本地镜像漏洞扫描任务ID
	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`

	// 本地镜像ID，无则全部
	LocalImageIDs []*string `json:"LocalImageIDs,omitempty" name:"LocalImageIDs"`

	// 仓库镜像ID，无则全部
	RegistryImageIDs []*uint64 `json:"RegistryImageIDs,omitempty" name:"RegistryImageIDs"`

	// 仓库镜像漏洞扫描任务ID
	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
}

type StopVulScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 本地镜像漏洞扫描任务ID
	LocalTaskID *int64 `json:"LocalTaskID,omitempty" name:"LocalTaskID"`

	// 本地镜像ID，无则全部
	LocalImageIDs []*string `json:"LocalImageIDs,omitempty" name:"LocalImageIDs"`

	// 仓库镜像ID，无则全部
	RegistryImageIDs []*uint64 `json:"RegistryImageIDs,omitempty" name:"RegistryImageIDs"`

	// 仓库镜像漏洞扫描任务ID
	RegistryTaskID *int64 `json:"RegistryTaskID,omitempty" name:"RegistryTaskID"`
}

func (r *StopVulScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopVulScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LocalTaskID")
	delete(f, "LocalImageIDs")
	delete(f, "RegistryImageIDs")
	delete(f, "RegistryTaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopVulScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopVulScanTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopVulScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopVulScanTaskResponseParams `json:"Response"`
}

func (r *StopVulScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopVulScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SupportDefenceVul struct {
	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 漏洞名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 漏洞标签
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 漏洞CVSS
	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`

	// 漏洞威胁等级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 漏洞CVEID
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 漏洞披露时间
	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`
}

// Predefined struct for user
type SwitchImageAutoAuthorizedRuleRequestParams struct {
	// 规则是否生效，0:不生效，1:已生效
	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type SwitchImageAutoAuthorizedRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则是否生效，0:不生效，1:已生效
	IsEnabled *int64 `json:"IsEnabled,omitempty" name:"IsEnabled"`

	// 规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *SwitchImageAutoAuthorizedRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchImageAutoAuthorizedRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsEnabled")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchImageAutoAuthorizedRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchImageAutoAuthorizedRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchImageAutoAuthorizedRuleResponse struct {
	*tchttp.BaseResponse
	Response *SwitchImageAutoAuthorizedRuleResponseParams `json:"Response"`
}

func (r *SwitchImageAutoAuthorizedRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchImageAutoAuthorizedRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncAssetImageRegistryAssetRequestParams struct {

}

type SyncAssetImageRegistryAssetRequest struct {
	*tchttp.BaseRequest
	
}

func (r *SyncAssetImageRegistryAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncAssetImageRegistryAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncAssetImageRegistryAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncAssetImageRegistryAssetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SyncAssetImageRegistryAssetResponse struct {
	*tchttp.BaseResponse
	Response *SyncAssetImageRegistryAssetResponseParams `json:"Response"`
}

func (r *SyncAssetImageRegistryAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncAssetImageRegistryAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type UnauthorizedCoresTendency struct {
	// 日期
	DateTime *string `json:"DateTime,omitempty" name:"DateTime"`

	// 未授权的核数
	CoresCount *int64 `json:"CoresCount,omitempty" name:"CoresCount"`
}

// Predefined struct for user
type UpdateAndPublishNetworkFirewallPolicyDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 出站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// pod选择器
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义规则
	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

type UpdateAndPublishNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 出站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// pod选择器
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义规则
	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

func (r *UpdateAndPublishNetworkFirewallPolicyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAndPublishNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Id")
	delete(f, "FromPolicyRule")
	delete(f, "ToPolicyRule")
	delete(f, "PodSelector")
	delete(f, "Namespace")
	delete(f, "Description")
	delete(f, "CustomPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAndPublishNetworkFirewallPolicyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAndPublishNetworkFirewallPolicyDetailResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAndPublishNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAndPublishNetworkFirewallPolicyDetailResponseParams `json:"Response"`
}

func (r *UpdateAndPublishNetworkFirewallPolicyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAndPublishNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAndPublishNetworkFirewallPolicyYamlDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Id")
	delete(f, "Yaml")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAndPublishNetworkFirewallPolicyYamlDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAndPublishNetworkFirewallPolicyYamlDetailResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponseParams `json:"Response"`
}

func (r *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAndPublishNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAssetImageRegistryRegistryDetailRequestParams struct {
	// 仓库名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 仓库url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 仓库类型，列表：harbor
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 网络类型，列表：public（公网）
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 仓库版本
	RegistryVersion *string `json:"RegistryVersion,omitempty" name:"RegistryVersion"`

	// 区域，列表：default（默认）
	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`

	// 限速
	SpeedLimit *int64 `json:"SpeedLimit,omitempty" name:"SpeedLimit"`

	// 安全模式（证书校验）：0（默认） 非安全模式（跳过证书校验）：1
	Insecure *uint64 `json:"Insecure,omitempty" name:"Insecure"`
}

type UpdateAssetImageRegistryRegistryDetailRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 仓库url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 仓库类型，列表：harbor
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 网络类型，列表：public（公网）
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 仓库版本
	RegistryVersion *string `json:"RegistryVersion,omitempty" name:"RegistryVersion"`

	// 区域，列表：default（默认）
	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`

	// 限速
	SpeedLimit *int64 `json:"SpeedLimit,omitempty" name:"SpeedLimit"`

	// 安全模式（证书校验）：0（默认） 非安全模式（跳过证书校验）：1
	Insecure *uint64 `json:"Insecure,omitempty" name:"Insecure"`
}

func (r *UpdateAssetImageRegistryRegistryDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAssetImageRegistryRegistryDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "Url")
	delete(f, "RegistryType")
	delete(f, "NetType")
	delete(f, "RegistryVersion")
	delete(f, "RegistryRegion")
	delete(f, "SpeedLimit")
	delete(f, "Insecure")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAssetImageRegistryRegistryDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAssetImageRegistryRegistryDetailResponseParams struct {
	// 连接错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckErr *string `json:"HealthCheckErr,omitempty" name:"HealthCheckErr"`

	// 名称错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameRepeatErr *string `json:"NameRepeatErr,omitempty" name:"NameRepeatErr"`

	// 仓库唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryId *int64 `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAssetImageRegistryRegistryDetailResponseParams `json:"Response"`
}

func (r *UpdateAssetImageRegistryRegistryDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAssetImageRegistryRegistryDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateImageRegistryTimingScanTaskRequestParams struct {
	// 定时扫描周期
	ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`

	// 定时扫描开关
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 定时扫描的时间
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 扫描木马类型数组
	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`

	// 扫描镜像
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 是否扫描所有
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描镜像Id
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

type UpdateImageRegistryTimingScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 定时扫描周期
	ScanPeriod *uint64 `json:"ScanPeriod,omitempty" name:"ScanPeriod"`

	// 定时扫描开关
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 定时扫描的时间
	ScanTime *string `json:"ScanTime,omitempty" name:"ScanTime"`

	// 扫描木马类型数组
	ScanType []*string `json:"ScanType,omitempty" name:"ScanType"`

	// 扫描镜像
	Images []*ImageInfo `json:"Images,omitempty" name:"Images"`

	// 是否扫描所有
	All *bool `json:"All,omitempty" name:"All"`

	// 扫描镜像Id
	Id []*uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *UpdateImageRegistryTimingScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateImageRegistryTimingScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScanPeriod")
	delete(f, "Enable")
	delete(f, "ScanTime")
	delete(f, "ScanType")
	delete(f, "Images")
	delete(f, "All")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateImageRegistryTimingScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateImageRegistryTimingScanTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateImageRegistryTimingScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *UpdateImageRegistryTimingScanTaskResponseParams `json:"Response"`
}

func (r *UpdateImageRegistryTimingScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateImageRegistryTimingScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNetworkFirewallPolicyDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 出站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// pod选择器
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义规则
	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

type UpdateNetworkFirewallPolicyDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 入站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	FromPolicyRule *int64 `json:"FromPolicyRule,omitempty" name:"FromPolicyRule"`

	// 出站规则
	// 
	// 全部允许：1
	// 
	// 全部拒绝 ：2
	// 
	// 自定义：3
	ToPolicyRule *int64 `json:"ToPolicyRule,omitempty" name:"ToPolicyRule"`

	// pod选择器
	PodSelector *string `json:"PodSelector,omitempty" name:"PodSelector"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义规则
	CustomPolicy []*NetworkCustomPolicy `json:"CustomPolicy,omitempty" name:"CustomPolicy"`
}

func (r *UpdateNetworkFirewallPolicyDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNetworkFirewallPolicyDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Id")
	delete(f, "FromPolicyRule")
	delete(f, "ToPolicyRule")
	delete(f, "PodSelector")
	delete(f, "Namespace")
	delete(f, "Description")
	delete(f, "CustomPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateNetworkFirewallPolicyDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNetworkFirewallPolicyDetailResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateNetworkFirewallPolicyDetailResponse struct {
	*tchttp.BaseResponse
	Response *UpdateNetworkFirewallPolicyDetailResponseParams `json:"Response"`
}

func (r *UpdateNetworkFirewallPolicyDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNetworkFirewallPolicyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNetworkFirewallPolicyYamlDetailRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type UpdateNetworkFirewallPolicyYamlDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 策略id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// base64编码的networkpolicy yaml字符串
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateNetworkFirewallPolicyYamlDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNetworkFirewallPolicyYamlDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Id")
	delete(f, "Yaml")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateNetworkFirewallPolicyYamlDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNetworkFirewallPolicyYamlDetailResponseParams struct {
	// 返回创建的任务的ID，为0表示创建失败。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 创建任务的结果，"Succ"为成功，"Failed"为失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateNetworkFirewallPolicyYamlDetailResponse struct {
	*tchttp.BaseResponse
	Response *UpdateNetworkFirewallPolicyYamlDetailResponseParams `json:"Response"`
}

func (r *UpdateNetworkFirewallPolicyYamlDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNetworkFirewallPolicyYamlDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirusAutoIsolateSampleInfo struct {
	// 文件MD5值
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// 病毒名
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 最近编辑时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 自动隔离开关(true:开 false:关)
	AutoIsolateSwitch *bool `json:"AutoIsolateSwitch,omitempty" name:"AutoIsolateSwitch"`
}

type VirusInfo struct {
	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件路径
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 病毒名称
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 容器id
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING
	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// DEAL_NONE:文件待处理
	// DEAL_IGNORE:已经忽略
	// DEAL_ADD_WHITELIST:加白
	// DEAL_DEL:文件已经删除
	// DEAL_ISOLATE:已经隔离
	// DEAL_ISOLATING:隔离中
	// DEAL_ISOLATE_FAILED:隔离失败
	// DEAL_RECOVERING:恢复中
	// DEAL_RECOVER_FAILED: 恢复失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 事件描述
	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`

	// 建议方案
	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`

	// 失败子状态:
	// FILE_NOT_FOUND:文件不存在
	// FILE_ABNORMAL:文件异常
	// FILE_ABNORMAL_DEAL_RECOVER:恢复文件时，文件异常
	// BACKUP_FILE_NOT_FOUND:备份文件不存在
	// CONTAINER_NOT_FOUND_DEAL_ISOLATE:隔离时，容器不存在
	// CONTAINER_NOT_FOUND_DEAL_RECOVER:恢复时，容器不存在
	// TIMEOUT: 超时
	// TOO_MANY: 任务过多
	// OFFLINE: 离线
	// INTERNAL: 服务内部错误
	// VALIDATION: 参数非法
	SubStatus *string `json:"SubStatus,omitempty" name:"SubStatus"`

	// 网络状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`

	// md5值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// 风险等级 RISK_CRITICAL, RISK_HIGH, RISK_MEDIUM, RISK_LOW, RISK_NOTICE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 检测平台
	// 1: 云查杀引擎
	// 2: tav
	// 3: binaryAi
	// 4: 异常行为
	// 5: 威胁情报
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPlatform []*string `json:"CheckPlatform,omitempty" name:"CheckPlatform"`
}

type VirusTaskInfo struct {
	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 容器id
	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 主机ip
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 扫描状态：
	// WAIT: 等待扫描
	// FAILED: 失败
	// SCANNING: 扫描中
	// FINISHED: 结束
	// CANCELING: 取消中
	// CANCELED: 已取消
	// CANCEL_FAILED： 取消失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 检测开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 检测结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 风险个数
	RiskCnt *uint64 `json:"RiskCnt,omitempty" name:"RiskCnt"`

	// 事件id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 错误原因:
	// SEND_SUCCESSED: 下发成功
	// SCAN_WAIT: agent排队扫描等待中
	// OFFLINE: 离线
	// SEND_FAILED:下发失败
	// TIMEOUT: 超时
	// LOW_AGENT_VERSION: 客户端版本过低
	// AGENT_NOT_FOUND: 镜像所属客户端版不存在
	// TOO_MANY: 任务过多
	// VALIDATION: 参数非法
	// INTERNAL: 服务内部错误
	// MISC: 其他错误
	// UNAUTH: 所在镜像未授权
	// SEND_CANCEL_SUCCESSED:下发成功
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

type VirusTendencyInfo struct {
	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 待处理事件总数
	PendingEventCount *uint64 `json:"PendingEventCount,omitempty" name:"PendingEventCount"`

	// 风险容器总数
	RiskContainerCount *uint64 `json:"RiskContainerCount,omitempty" name:"RiskContainerCount"`

	// 事件总数
	EventCount *uint64 `json:"EventCount,omitempty" name:"EventCount"`

	// 隔离事件总数
	IsolateEventCount *uint64 `json:"IsolateEventCount,omitempty" name:"IsolateEventCount"`
}

type VulAffectedComponentInfo struct {
	// 组件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 组件版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version []*string `json:"Version,omitempty" name:"Version"`

	// 组件修复版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixedVersion []*string `json:"FixedVersion,omitempty" name:"FixedVersion"`
}

type VulAffectedContainerInfo struct {
	// 内网IP
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 容器ID
	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// Pod名称
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// PodIP值
	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`

	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 主机ID
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 外网IP
	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`
}

type VulAffectedImageComponentInfo struct {
	// 组件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 组件版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 组件修复版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixedVersion *string `json:"FixedVersion,omitempty" name:"FixedVersion"`

	// 组件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`
}

type VulAffectedImageInfo struct {
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 关联的主机数
	HostCount *int64 `json:"HostCount,omitempty" name:"HostCount"`

	// 关联的容器数
	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`

	// 组件列表
	ComponentList []*VulAffectedImageComponentInfo `json:"ComponentList,omitempty" name:"ComponentList"`
}

type VulDefenceEvent struct {
	// 漏洞CVEID
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 漏洞名称
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 入侵状态
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 攻击源IP
	SourceIP *string `json:"SourceIP,omitempty" name:"SourceIP"`

	// 攻击源ip地址所在城市
	City *string `json:"City,omitempty" name:"City"`

	// 事件数量
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 容器ID
	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 处理状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件ID
	EventID *int64 `json:"EventID,omitempty" name:"EventID"`

	// 首次发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 隔离状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 最近发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`

	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`

	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`

	// 主机QUUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	QUUID *string `json:"QUUID,omitempty" name:"QUUID"`

	// 主机内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 主机名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

type VulDefenceEventDetail struct {
	// 漏洞CVEID
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 漏洞名称
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 入侵状态
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 攻击源IP
	SourceIP *string `json:"SourceIP,omitempty" name:"SourceIP"`

	// 攻击源ip地址所在城市
	City *string `json:"City,omitempty" name:"City"`

	// 事件数量
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 容器ID
	ContainerID *string `json:"ContainerID,omitempty" name:"ContainerID"`

	// 容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 处理状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 攻击源端口
	SourcePort []*string `json:"SourcePort,omitempty" name:"SourcePort"`

	// 事件ID
	EventID *int64 `json:"EventID,omitempty" name:"EventID"`

	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 主机内网IP
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 主机外网IP
	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`

	// Pod名称
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 危害描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 修复建议
	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`

	// 攻击包
	NetworkPayload *string `json:"NetworkPayload,omitempty" name:"NetworkPayload"`

	// 进程PID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PID *int64 `json:"PID,omitempty" name:"PID"`

	// 进程主类名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// 堆栈信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StackTrace *string `json:"StackTrace,omitempty" name:"StackTrace"`

	// 监听账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerAccount *string `json:"ServerAccount,omitempty" name:"ServerAccount"`

	// 监听端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerPort *string `json:"ServerPort,omitempty" name:"ServerPort"`

	// 进程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerExe *string `json:"ServerExe,omitempty" name:"ServerExe"`

	// 进程命令行参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerArg *string `json:"ServerArg,omitempty" name:"ServerArg"`

	// 主机QUUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	QUUID *string `json:"QUUID,omitempty" name:"QUUID"`

	// 隔离状态
	// 未隔离  	NORMAL
	// 已隔离		ISOLATED
	// 隔离中		ISOLATING
	// 隔离失败	ISOLATE_FAILED
	// 解除隔离中  RESTORING
	// 解除隔离失败 RESTORE_FAILED
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetStatus *string `json:"ContainerNetStatus,omitempty" name:"ContainerNetStatus"`

	// 容器子状态
	// "AGENT_OFFLINE"       //Agent离线
	// 	"NODE_DESTROYED"      //节点已销毁
	// 	"CONTAINER_EXITED"    //容器已退出
	// 	"CONTAINER_DESTROYED" //容器已销毁
	// 	"SHARED_HOST"         // 容器与主机共享网络
	// 	"RESOURCE_LIMIT"      //隔离操作资源超限
	// 	"UNKNOW"              // 原因未知
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetSubStatus *string `json:"ContainerNetSubStatus,omitempty" name:"ContainerNetSubStatus"`

	// 容器隔离操作来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerIsolateOperationSrc *string `json:"ContainerIsolateOperationSrc,omitempty" name:"ContainerIsolateOperationSrc"`

	// 容器状态
	// 正在运行: RUNNING
	// 暂停: PAUSED
	// 停止: STOPPED
	// 已经创建: CREATED
	// 已经销毁: DESTROYED
	// 正在重启中: RESTARTING
	// 迁移中: REMOVING
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerStatus *string `json:"ContainerStatus,omitempty" name:"ContainerStatus"`

	// 接口Url
	// 注意：此字段可能返回 null，表示取不到有效值。
	JNDIUrl *string `json:"JNDIUrl,omitempty" name:"JNDIUrl"`

	// rasp detail
	// 注意：此字段可能返回 null，表示取不到有效值。
	RaspDetail []*RaspInfo `json:"RaspDetail,omitempty" name:"RaspDetail"`
}

type VulDefenceEventTendency struct {
	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 事件数量
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`
}

type VulDefenceHost struct {
	// 主机名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 主机ip即内网ip
	HostIP *string `json:"HostIP,omitempty" name:"HostIP"`

	// 主机QUUID
	HostID *string `json:"HostID,omitempty" name:"HostID"`

	// 插件状态，正常：SUCCESS，异常：FAIL， NO_DEFENDED:未防御
	Status *string `json:"Status,omitempty" name:"Status"`

	// 外网ip
	PublicIP *string `json:"PublicIP,omitempty" name:"PublicIP"`

	// 首次开启时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type VulDefencePlugin struct {
	// java进程pid
	PID *int64 `json:"PID,omitempty" name:"PID"`

	// 进程主类名
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// 插件运行状态：注入中:INJECTING，注入成功：SUCCESS，注入失败：FAIL，插件超时：TIMEOUT，插件退出：QUIT
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误日志
	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`
}

type VulDetailInfo struct {
	// CVE编号
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 漏洞名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 漏洞标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 漏洞类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`

	// 漏洞威胁等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 漏洞披露时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`

	// 漏洞描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// CVSS V3描述
	CVSSV3Desc *string `json:"CVSSV3Desc,omitempty" name:"CVSSV3Desc"`

	// 漏洞修复建议
	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`

	// 缓解措施
	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`

	// 参考链接
	Reference []*string `json:"Reference,omitempty" name:"Reference"`

	// CVSS V3分数
	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`

	// 受漏洞影响的组件列表
	ComponentList []*VulAffectedComponentInfo `json:"ComponentList,omitempty" name:"ComponentList"`

	// 影响本地镜像数
	LocalImageCount *int64 `json:"LocalImageCount,omitempty" name:"LocalImageCount"`

	// 影响容器数
	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`

	// 影响仓库镜像数
	RegistryImageCount *int64 `json:"RegistryImageCount,omitempty" name:"RegistryImageCount"`

	// 漏洞子类型
	Category *string `json:"Category,omitempty" name:"Category"`

	// 影响最新本地镜像数
	LocalNewestImageCount *int64 `json:"LocalNewestImageCount,omitempty" name:"LocalNewestImageCount"`

	// 影响最新仓库镜像数
	RegistryNewestImageCount *int64 `json:"RegistryNewestImageCount,omitempty" name:"RegistryNewestImageCount"`

	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 防御状态，NO_DEFENDED:未防御，DEFENDED:已防御
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenceStatus *string `json:"DefenceStatus,omitempty" name:"DefenceStatus"`

	// 漏洞防御主机范围: MANUAL:自选主机节点，ALL:全部
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenceScope *string `json:"DefenceScope,omitempty" name:"DefenceScope"`

	// 漏洞防御主机数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenceHostCount *int64 `json:"DefenceHostCount,omitempty" name:"DefenceHostCount"`

	// 已防御攻击次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefendedCount *int64 `json:"DefendedCount,omitempty" name:"DefendedCount"`

	// 是否已扫描，NOT_SCAN:未扫描,SCANNED:已扫描
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`
}

type VulIgnoreLocalImage struct {
	// 记录ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像大小
	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

type VulIgnoreRegistryImage struct {
	// 记录ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 仓库名称
	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`

	// 仓库地址
	RegistryPath *string `json:"RegistryPath,omitempty" name:"RegistryPath"`

	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

type VulInfo struct {
	// 漏洞名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 漏洞标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// CVSS V3分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSSV3Score *float64 `json:"CVSSV3Score,omitempty" name:"CVSSV3Score"`

	// 风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitempty" name:"Level"`

	// CVE编号
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 漏洞子类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 首次发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FoundTime *string `json:"FoundTime,omitempty" name:"FoundTime"`

	// 最近发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestFoundTime *string `json:"LatestFoundTime,omitempty" name:"LatestFoundTime"`

	// 漏洞ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 影响本地镜像数
	LocalImageCount *int64 `json:"LocalImageCount,omitempty" name:"LocalImageCount"`

	// 影响容器数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`

	// 影响仓库镜像数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryImageCount *int64 `json:"RegistryImageCount,omitempty" name:"RegistryImageCount"`

	// 漏洞PocID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PocID *string `json:"PocID,omitempty" name:"PocID"`

	// 防御状态，NO_DEFENDED:未防御，DEFENDED:已防御
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenceStatus *string `json:"DefenceStatus,omitempty" name:"DefenceStatus"`

	// 漏洞防御主机范围: MANUAL:自选主机节点，ALL:全部
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenceScope *string `json:"DefenceScope,omitempty" name:"DefenceScope"`

	// 漏洞防御主机数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenceHostCount *int64 `json:"DefenceHostCount,omitempty" name:"DefenceHostCount"`

	// 已防御攻击次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefendedCount *int64 `json:"DefendedCount,omitempty" name:"DefendedCount"`
}

type VulScanImageInfo struct {
	// 镜像ID
	ImageID *string `json:"ImageID,omitempty" name:"ImageID"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像大小
	Size *float64 `json:"Size,omitempty" name:"Size"`

	// 任务状态:SCANNING:扫描中 FAILED:失败 FINISHED:完成 CANCELED:取消
	ScanStatus *string `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 扫描时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanDuration *float64 `json:"ScanDuration,omitempty" name:"ScanDuration"`

	// 高危漏洞数
	HighLevelVulCount *int64 `json:"HighLevelVulCount,omitempty" name:"HighLevelVulCount"`

	// 中危漏洞数
	MediumLevelVulCount *int64 `json:"MediumLevelVulCount,omitempty" name:"MediumLevelVulCount"`

	// 低危漏洞数
	LowLevelVulCount *int64 `json:"LowLevelVulCount,omitempty" name:"LowLevelVulCount"`

	// 严重漏洞数
	CriticalLevelVulCount *int64 `json:"CriticalLevelVulCount,omitempty" name:"CriticalLevelVulCount"`

	// 本地镜像漏洞扫描任务ID
	TaskID *int64 `json:"TaskID,omitempty" name:"TaskID"`

	// 漏洞扫描的开始时间
	ScanStartTime *string `json:"ScanStartTime,omitempty" name:"ScanStartTime"`

	// 漏洞扫描的结束时间
	ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`

	// 失败原因:TIMEOUT:超时 TOO_MANY:任务过多 OFFLINE:离线
	ErrorStatus *string `json:"ErrorStatus,omitempty" name:"ErrorStatus"`
}

type VulTendencyInfo struct {
	// 漏洞趋势列表
	VulSet []*RunTimeTendencyInfo `json:"VulSet,omitempty" name:"VulSet"`

	// 漏洞影响的镜像类型：
	// LOCAL：本地镜像
	// REGISTRY: 仓库镜像
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
}

type VulTopRankingInfo struct {
	// 漏洞名称
	VulName *string `json:"VulName,omitempty" name:"VulName"`

	// 威胁等级,CRITICAL:严重 HIGH:高/MIDDLE:中/LOW:低
	Level *string `json:"Level,omitempty" name:"Level"`

	// 影响的镜像数
	AffectedImageCount *int64 `json:"AffectedImageCount,omitempty" name:"AffectedImageCount"`

	// 影响的容器数
	AffectedContainerCount *int64 `json:"AffectedContainerCount,omitempty" name:"AffectedContainerCount"`

	// 漏洞ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 漏洞PocID
	PocID *string `json:"PocID,omitempty" name:"PocID"`
}

type WarningRule struct {
	// 告警事件类型：
	// 镜像仓库安全-木马：IMG_REG_VIRUS
	// 镜像仓库安全-漏洞：IMG_REG_VUL
	// 镜像仓库安全-敏感信息：IMG_REG_RISK
	// 镜像安全-木马：IMG_VIRUS
	// 镜像安全-漏洞：IMG_VUL
	// 镜像安全-敏感信息：IMG_RISK
	// 镜像安全-镜像拦截：IMG_INTERCEPT
	// 运行时安全-容器逃逸：RUNTIME_ESCAPE
	// 运行时安全-异常进程：RUNTIME_FILE
	// 运行时安全-异常文件访问：RUNTIME_PROCESS
	// 运行时安全-高危系统调用：RUNTIME_SYSCALL
	// 运行时安全-反弹Shell：RUNTIME_REVERSE_SHELL
	// 运行时安全-木马：RUNTIME_VIRUS
	Type *string `json:"Type,omitempty" name:"Type"`

	// 开关状态：
	// 打开：ON
	// 关闭：OFF
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 告警开始时间，格式: HH:mm
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 告警结束时间，格式: HH:mm
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 告警等级策略控制，二进制位每位代表一个含义，值以字符串类型传递
	// 控制开关分为高、中、低，则二进制位分别为：第1位:低，第2位:中，第3位:高，0表示关闭、1表示打开。
	// 如：高危和中危打开告警，低危关闭告警，则二进制值为：110
	// 告警类型不区分等级控制，则传1。
	ControlBits *string `json:"ControlBits,omitempty" name:"ControlBits"`
}