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

	// 命中规则名字
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 命中规则的id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 事件最后一次处理的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`
}

type AbnormalProcessEventInfo struct {

	// 进程目录
	ProcessPath *string `json:"ProcessPath,omitempty" name:"ProcessPath"`

	// 事件类型，MALICE_PROCESS_START:恶意进程启动
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 命中规则
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

type AddAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type AddCompliancePolicyItemToWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type AddEditAbnormalProcessRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type AddEditAccessControlRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type AddEditReverseShellWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type AddEditRiskSyscallWhiteListRequest struct {
	*tchttp.BaseRequest

	// 仅在添加白名单时候使用
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 增加白名单信息，白名单id为空，编辑白名单id不能为空
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

type AddEditRiskSyscallWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type AddEditWarningRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CheckRepeatAssetImageRegistryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否重复
	// 注意：此字段可能返回 null，表示取不到有效值。
		IsRepeat *bool `json:"IsRepeat,omitempty" name:"IsRepeat"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateAssetImageRegistryScanTaskOneKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateAssetImageRegistryScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateAssetImageScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateAssetImageScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
		TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateCheckComponentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// "InstallSucc"表示安装成功，"InstallFailed"表示安装失败
		InstallResult *string `json:"InstallResult,omitempty" name:"InstallResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateClusterCheckTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回创建的集群检查任务的ID，为0表示创建失败。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 创建检查任务的结果，"Succ"为成功，其他的为失败原因
		CreateResult *string `json:"CreateResult,omitempty" name:"CreateResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateComplianceTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回创建的合规检查任务的ID。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateExportComplianceStatusListJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回创建的导出任务的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateOrModifyPostPayCoresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateRefreshTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回创建的集群检查任务的ID，为0表示创建失败。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 创建检查任务的结果，"Succ"为成功，"Failed"为失败
		CreateResult *string `json:"CreateResult,omitempty" name:"CreateResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateVirusScanAgainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateVirusScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
		TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteAbnormalProcessRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteAccessControlRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteCompliancePolicyItemFromWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteReverseShellWhiteListsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteRiskSyscallWhiteListsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAbnormalProcessDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAbnormalProcessEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAbnormalProcessEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 异常进程数组
		EventSet []*AbnormalProcessEventInfo `json:"EventSet,omitempty" name:"EventSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAbnormalProcessRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异常进程策略详细信息
		RuleDetail *AbnormalProcessRuleInfo `json:"RuleDetail,omitempty" name:"RuleDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAbnormalProcessRulesExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAbnormalProcessRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 异常进程策略信息列表
		RuleSet []*RuleBaseInfo `json:"RuleSet,omitempty" name:"RuleSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAccessControlDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAccessControlEventsExportRequest struct {
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
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessControlEventsExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessControlEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAccessControlEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 访问控制事件数组
		EventSet []*AccessControlEventInfo `json:"EventSet,omitempty" name:"EventSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAccessControlRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 运行时策略详细信息
		RuleDetail *AccessControlRuleInfo `json:"RuleDetail,omitempty" name:"RuleDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAccessControlRulesExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAccessControlRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 访问控制策略信息列表
		RuleSet []*RuleBaseInfo `json:"RuleSet,omitempty" name:"RuleSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAffectedClusterCountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAffectedNodeListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 受影响的节点总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 受影响的节点列表
		AffectedNodeList []*AffectedNodeItem `json:"AffectedNodeList,omitempty" name:"AffectedNodeList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAffectedWorkloadListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 受影响的workload列表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 受影响的workload列表
		AffectedWorkloadList []*AffectedWorkloadItem `json:"AffectedWorkloadList,omitempty" name:"AffectedWorkloadList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetAppServiceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// db服务列表
		List []*ServiceInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetComponentListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组件列表
		List []*ComponentInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetContainerDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetContainerListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 容器列表
		List []*ContainerInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetDBServiceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// db服务列表
		List []*ServiceInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetHostDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetHostListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主机列表
		List []*HostInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageBindRuleInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 镜像绑定规则列表
		ImageBindRuleSet []*ImagesBindRuleInfo `json:"ImageBindRuleSet,omitempty" name:"ImageBindRuleSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAssetImageHostListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像列表
		List []*ImageHost `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageListExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// excel文件下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像列表
		List []*ImagesInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryAssetStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新进度状态,doing更新中，success更新成功，failed失败
		Status *string `json:"Status,omitempty" name:"Status"`

		// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Err *string `json:"Err,omitempty" name:"Err"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAssetImageRegistryListExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// excel文件下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像仓库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*ImageRepoInfo `json:"List,omitempty" name:"List"`

		// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAssetImageRegistryRegistryListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryRiskInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像漏洞列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*ImageRisk `json:"List,omitempty" name:"List"`

		// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryRiskListExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// excel文件下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryScanStatusOneKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAssetImageRegistrySummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryVirusListExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// excel文件下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryVirusListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像漏洞列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*ImageVirus `json:"List,omitempty" name:"List"`

		// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryVulListExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// excel文件下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRegistryVulListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像漏洞列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*ImageVul `json:"List,omitempty" name:"List"`

		// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRiskListExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// excel文件下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageRiskListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像病毒列表
		List []*ImageRiskInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAssetImageScanStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAssetImageScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
		TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageSimpleListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像列表
		List []*AssetSimpleImageInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageVirusListExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// excel文件下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageVirusListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeAssetImageVulListExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// excel文件下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetImageVulListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像漏洞列表
		List []*ImagesVul `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetPortListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 端口列表
		List []*PortInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetProcessListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 端口列表
		List []*ProcessInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAssetWebServiceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主机列表
		List []*ServiceInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCheckItemListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检查项详情数组
		ClusterCheckItems []*ClusterCheckItem `json:"ClusterCheckItems,omitempty" name:"ClusterCheckItems"`

		// 检查项总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeClusterSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeComplianceAssetDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeComplianceAssetListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回资产的总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回各类资产的列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AssetInfoList []*ComplianceAssetInfo `json:"AssetInfoList,omitempty" name:"AssetInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeComplianceAssetPolicyItemListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回检测项的总数。如果用户未启用基线检查，此处返回0。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回某个资产下的检测项的列表。
		AssetPolicyItemList []*ComplianceAssetPolicyItem `json:"AssetPolicyItemList,omitempty" name:"AssetPolicyItemList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCompliancePeriodTaskListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 定时任务的总量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 定时任务信息的列表
		PeriodTaskSet []*CompliancePeriodTask `json:"PeriodTaskSet,omitempty" name:"PeriodTaskSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCompliancePolicyItemAffectedAssetListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回各检测项所影响的资产的列表。
		AffectedAssetList []*ComplianceAffectedAsset `json:"AffectedAssetList,omitempty" name:"AffectedAssetList"`

		// 检测项影响的资产的总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCompliancePolicyItemAffectedSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回各检测项影响的资产的汇总信息。
		PolicyItemSummary *CompliancePolicyItemSummary `json:"PolicyItemSummary,omitempty" name:"PolicyItemSummary"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeComplianceScanFailedAssetListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回检测失败的资产的总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回各类检测失败的资产的汇总信息的列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ScanFailedAssetList []*ComplianceScanFailedAsset `json:"ScanFailedAssetList,omitempty" name:"ScanFailedAssetList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeComplianceTaskAssetSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeComplianceTaskPolicyItemSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回最近一次合规检查任务的ID。这个任务为本次所展示数据的来源。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 返回检测项的总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回各检测项对应的汇总信息的列表。
		PolicyItemSummaryList []*CompliancePolicyItemSummary `json:"PolicyItemSummaryList,omitempty" name:"PolicyItemSummaryList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeComplianceWhitelistItemListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 白名单项的列表。
		WhitelistItemSet []*ComplianceWhitelistItem `json:"WhitelistItemSet,omitempty" name:"WhitelistItemSet"`

		// 白名单项的总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeContainerAssetSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeContainerSecEventSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEscapeEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeEscapeEventInfoRequest struct {
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

type DescribeEscapeEventInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 逃逸事件数组
		EventSet []*EscapeEventInfo `json:"EventSet,omitempty" name:"EventSet"`

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEscapeEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEscapeRuleInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则信息
		RuleSet []*EscapeRule `json:"RuleSet,omitempty" name:"RuleSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEscapeSafeStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unsafe：存在风险，Safe：暂无风险,UnKnown:未知风险
		IsSafe *string `json:"IsSafe,omitempty" name:"IsSafe"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeExportJobResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeImageAuthorizedInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeImageRegistryTimingScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeImageRiskSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全漏洞
		VulnerabilityCnt []*RunTimeRiskInfo `json:"VulnerabilityCnt,omitempty" name:"VulnerabilityCnt"`

		// 木马病毒
		MalwareVirusCnt []*RunTimeRiskInfo `json:"MalwareVirusCnt,omitempty" name:"MalwareVirusCnt"`

		// 敏感信息
		RiskCnt []*RunTimeRiskInfo `json:"RiskCnt,omitempty" name:"RiskCnt"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeImageRiskTendencyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本地镜像新增风险趋势信息列表
		ImageRiskTendencySet []*ImageRiskTendencyInfo `json:"ImageRiskTendencySet,omitempty" name:"ImageRiskTendencySet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeImageSimpleListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像列表
		ImageList []*ImageSimpleInfo `json:"ImageList,omitempty" name:"ImageList"`

		// 镜像数
		ImageCnt *uint64 `json:"ImageCnt,omitempty" name:"ImageCnt"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribePostPayDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 弹性计费扣费详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		SoftQuotaDayDetail []*SoftQuotaDayInfo `json:"SoftQuotaDayDetail,omitempty" name:"SoftQuotaDayDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeProVersionInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribePurchaseStateInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRefreshTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 刷新任务状态，可能为：Task_New,Task_Running,Task_Finish,Task_Error,Task_NoExist
		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeReverseShellDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeReverseShellEventsExportRequest struct {
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
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellEventsExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// execle下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 反弹shell数组
		EventSet []*ReverseShellEventInfo `json:"EventSet,omitempty" name:"EventSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeReverseShellWhiteListDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件基本信息
		WhiteListDetailInfo *ReverseShellWhiteListInfo `json:"WhiteListDetailInfo,omitempty" name:"WhiteListDetailInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeReverseShellWhiteListsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 白名单信息列表
		WhiteListSet []*ReverseShellWhiteListBaseInfo `json:"WhiteListSet,omitempty" name:"WhiteListSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRiskListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 风险详情数组
		ClusterRiskItems []*ClusterRiskItem `json:"ClusterRiskItems,omitempty" name:"ClusterRiskItems"`

		// 风险项的总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRiskSyscallDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeRiskSyscallEventsExportRequest struct {
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
	delete(f, "ExportField")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskSyscallEventsExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskSyscallEventsExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Excel下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRiskSyscallEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 高危系统调用数组
		EventSet []*RiskSyscallEventInfo `json:"EventSet,omitempty" name:"EventSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRiskSyscallNamesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 系统调用名称列表
		SyscallNames []*string `json:"SyscallNames,omitempty" name:"SyscallNames"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRiskSyscallWhiteListDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 白名单基本信息
		WhiteListDetailInfo *RiskSyscallWhiteListInfo `json:"WhiteListDetailInfo,omitempty" name:"WhiteListDetailInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRiskSyscallWhiteListsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 白名单信息列表
		WhiteListSet []*RiskSyscallWhiteListBaseInfo `json:"WhiteListSet,omitempty" name:"WhiteListSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSecEventsTendencyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 运行时安全事件趋势信息列表
		EventTendencySet []*SecTendencyEventInfo `json:"EventTendencySet,omitempty" name:"EventTendencySet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeTaskResultSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeUnfinishRefreshTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回最近的一次任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 任务状态，为Task_New,Task_Running,Task_Finish,Task_Error,Task_NoExist.Task_New,Task_Running表示有任务存在，不允许新下发
		TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeUserClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群的详细信息
		ClusterInfoList []*ClusterInfoItem `json:"ClusterInfoList,omitempty" name:"ClusterInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeValueAddedSrvInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 仓库镜像未授权数量
		RegistryImageCnt *uint64 `json:"RegistryImageCnt,omitempty" name:"RegistryImageCnt"`

		// 本地镜像未授权数量
		LocalImageCnt *uint64 `json:"LocalImageCnt,omitempty" name:"LocalImageCnt"`

		// 未使用的镜像安全扫描授权数
		UnusedAuthorizedCnt *uint64 `json:"UnusedAuthorizedCnt,omitempty" name:"UnusedAuthorizedCnt"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeVirusDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeVirusListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 木马列表
		List []*VirusInfo `json:"List,omitempty" name:"List"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeVirusMonitorSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeVirusScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeVirusScanTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeVirusScanTimeoutSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 超时时长单位小时
	// 注意：此字段可能返回 null，表示取不到有效值。
		Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeVirusSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	// <li>HostIp- String - 是否必填：是 - 容器名称</li>
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

type DescribeVirusTaskListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件查杀列表
		List []*VirusTaskInfo `json:"List,omitempty" name:"List"`

		// 总数量(容器任务数量)
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeWarningRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略列表
		WarningRules []*WarningRule `json:"WarningRules,omitempty" name:"WarningRules"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 状态
	//      EVENT_UNDEAL:事件未处理
	//      EVENT_DEALED:事件已经处理
	//      EVENT_INGNORE：事件忽略
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

type ExportVirusListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出任务ID，前端拿着任务ID查询任务进度
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type InitializeUserComplianceEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyAbnormalProcessRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyAbnormalProcessStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyAccessControlRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyAccessControlStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyAssetImageRegistryScanStopOneKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyAssetImageRegistryScanStopResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyAssetImageScanStopResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 停止状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 同步任务发送结果
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyCompliancePeriodTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyEscapeEventStatusRequest struct {
	*tchttp.BaseRequest

	// 处理事件ids
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态
	//    EVENT_DEALED:事件已经处理
	//      EVENT_INGNORE：事件忽略
	//      EVENT_DEL:事件删除
	Status *string `json:"Status,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEscapeEventStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEscapeEventStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyEscapeRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyReverseShellStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyRiskSyscallStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyVirusFileStatusRequest struct {
	*tchttp.BaseRequest

	// 处理事件id
	EventIdSet []*string `json:"EventIdSet,omitempty" name:"EventIdSet"`

	// 标记事件的状态，   
	//     EVENT_DEALED:事件处理
	//     EVENT_INGNORE"：事件忽略
	//      EVENT_DEL:事件删除
	//      EVENT_ADD_WHITE:事件加白
	//      EVENT_PENDING: 事件待处理
	Status *string `json:"Status,omitempty" name:"Status"`

	// 事件备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVirusFileStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVirusFileStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyVirusMonitorSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyVirusScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyVirusScanTimeoutSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type RemoveAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type RenewImageAuthorizeStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ScanComplianceAssetsByPolicyItemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回重新检测任务的ID。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ScanComplianceAssetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回重新检测任务的ID。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ScanCompliancePolicyItemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回重新检测任务的ID。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ScanComplianceScanFailedAssetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回重新检测任务的ID。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type SecTendencyEventInfo struct {

	// 趋势列表
	EventSet []*RunTimeTendencyInfo `json:"EventSet,omitempty" name:"EventSet"`

	// 事件类型：
	// ET_ESCAPE : 容器逃逸
	// ET_REVERSE_SHELL: 反弹shell
	// ET_RISK_SYSCALL:高危系统调用
	// ET_ABNORMAL_PROCESS: 异常进程
	// ET_ACCESS_CONTROL 文件篡改
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

type SetCheckModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// "Succ"表示设置成功，"Failed"表示设置失败
		SetCheckResult *string `json:"SetCheckResult,omitempty" name:"SetCheckResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type StopVirusScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type SyncAssetImageRegistryAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type UpdateAssetImageRegistryRegistryDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type UpdateImageRegistryTimingScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 容器状态，CS_RUNING:运行， CS_PAUSE:暂停，CS_STOP:停止，
	// 												       CS_CREATE:已经创建， CS_DESTORY:销毁
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
