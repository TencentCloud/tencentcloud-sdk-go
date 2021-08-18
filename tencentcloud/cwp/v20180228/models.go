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

package v20180228

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Account struct {

	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 帐号名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 帐号所属组。
	Groups *string `json:"Groups,omitempty" name:"Groups"`

	// 帐号类型。
	// <li>ORDINARY：普通帐号</li>
	// <li>SUPPER：超级管理员帐号</li>
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// 帐号创建时间。
	AccountCreateTime *string `json:"AccountCreateTime,omitempty" name:"AccountCreateTime"`

	// 帐号最后登录时间。
	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`
}

type AccountStatistics struct {

	// 用户名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type AssetFilters struct {

	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否模糊查询
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type AssetKeyVal struct {

	// 标签
	Key *string `json:"Key,omitempty" name:"Key"`

	// 数量
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type BashEvent struct {

	// 数据ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 执行用户名
	User *string `json:"User,omitempty" name:"User"`

	// 平台类型
	Platform *uint64 `json:"Platform,omitempty" name:"Platform"`

	// 执行命令
	BashCmd *string `json:"BashCmd,omitempty" name:"BashCmd"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则等级：1-高 2-中 3-低
	RuleLevel *uint64 `json:"RuleLevel,omitempty" name:"RuleLevel"`

	// 处理状态： 0 = 待处理 1= 已处理, 2 = 已加白
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 0: bash日志 1: 实时监控(雷霆版)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectBy *uint64 `json:"DetectBy,omitempty" name:"DetectBy"`

	// 进程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pid *string `json:"Pid,omitempty" name:"Pid"`

	// 进程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exe *string `json:"Exe,omitempty" name:"Exe"`

	// 处理时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 规则类别  0=系统规则，1=用户规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCategory *uint64 `json:"RuleCategory,omitempty" name:"RuleCategory"`

	// 自动生成的正则表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegexBashCmd *string `json:"RegexBashCmd,omitempty" name:"RegexBashCmd"`
}

type BashRule struct {

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 危险等级(0 ：无 1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 正则表达式
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// 规则描述
	Decription *string `json:"Decription,omitempty" name:"Decription"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 状态 (0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 生效服务器的uuid数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`

	// 0=黑名单 1=白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	White *uint64 `json:"White,omitempty" name:"White"`

	// 是否处理之前的事件 0: 不处理 1:处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealOldEvents *uint64 `json:"DealOldEvents,omitempty" name:"DealOldEvents"`
}

type BruteAttackInfo struct {

	// 唯一Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一标识UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 来源ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// SUCCESS：破解成功；FAILED：破解失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 国家id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 城市id
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *uint64 `json:"City,omitempty" name:"City"`

	// 省份id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *uint64 `json:"Province,omitempty" name:"Province"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 阻断状态：1-阻断成功；非1-阻断失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	BanStatus *uint64 `json:"BanStatus,omitempty" name:"BanStatus"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *uint64 `json:"EventType,omitempty" name:"EventType"`

	// 发生次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 机器UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 是否为专业版（true/false）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

	// 被攻击的服务的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 最近攻击时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type CloseProVersionRequest struct {
	*tchttp.BaseRequest

	// 主机唯一标识Uuid。
	// 黑石的InstanceId，CVM的Uuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *CloseProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseProVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseProVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentStatistics struct {

	// 组件ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`

	// 组件名称。
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 组件类型。
	// <li>WEB：Web组件</li>
	// <li>SYSTEM：系统组件</li>
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 组件描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateBaselineStrategyRequest struct {
	*tchttp.BaseRequest

	// 策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 检测周期, 表示每隔多少天进行检测.示例: 2, 表示每2天进行检测一次.
	ScanCycle *uint64 `json:"ScanCycle,omitempty" name:"ScanCycle"`

	// 定期检测时间，该时间下发扫描. 示例:“22:00”, 表示在22:00下发检测
	ScanAt *string `json:"ScanAt,omitempty" name:"ScanAt"`

	// 该策略下选择的基线id数组. 示例: [1,3,5,7]
	CategoryIds []*uint64 `json:"CategoryIds,omitempty" name:"CategoryIds"`

	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 云主机类型：“CVM”：虚拟主机，"BMS"：裸金属，"ECM"：边缘计算主机
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 主机地域. 示例: "ap-bj"
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 主机id数组. 示例: ["quuid1","quuid2"]
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *CreateBaselineStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaselineStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyName")
	delete(f, "ScanCycle")
	delete(f, "ScanAt")
	delete(f, "CategoryIds")
	delete(f, "IsGlobal")
	delete(f, "MachineType")
	delete(f, "RegionCode")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBaselineStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateBaselineStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBaselineStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaselineStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProtectServerRequest struct {
	*tchttp.BaseRequest

	// 防护目录地址
	ProtectDir *string `json:"ProtectDir,omitempty" name:"ProtectDir"`

	// 防护机器 信息
	ProtectHostConfig []*ProtectHostConfig `json:"ProtectHostConfig,omitempty" name:"ProtectHostConfig"`
}

func (r *CreateProtectServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtectServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectDir")
	delete(f, "ProtectHostConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProtectServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProtectServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProtectServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtectServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScanMalwareSettingRequest struct {
	*tchttp.BaseRequest

	// 扫描模式 0 全盘扫描, 1 快速扫描
	ScanPattern *uint64 `json:"ScanPattern,omitempty" name:"ScanPattern"`

	// 服务器分类：1:专业版服务器；2:自选服务器
	HostType *int64 `json:"HostType,omitempty" name:"HostType"`

	// 自选服务器时生效，主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`

	// 超时时间单位 秒
	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitempty" name:"TimeoutPeriod"`
}

func (r *CreateScanMalwareSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanMalwareSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScanPattern")
	delete(f, "HostType")
	delete(f, "QuuidList")
	delete(f, "TimeoutPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScanMalwareSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateScanMalwareSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScanMalwareSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanMalwareSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSearchLogRequest struct {
	*tchttp.BaseRequest

	// 搜索内容
	SearchContent *string `json:"SearchContent,omitempty" name:"SearchContent"`
}

func (r *CreateSearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSearchLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSearchLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：成功，非0：失败
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type CreateSearchTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：成功，非0：失败
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DefendAttackLog struct {

	// 日志ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 来源IP
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 来源端口
	SrcPort *uint64 `json:"SrcPort,omitempty" name:"SrcPort"`

	// 攻击方式
	HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`

	// 攻击描述
	HttpCgi *string `json:"HttpCgi,omitempty" name:"HttpCgi"`

	// 攻击参数
	HttpParam *string `json:"HttpParam,omitempty" name:"HttpParam"`

	// 威胁类型
	VulType *string `json:"VulType,omitempty" name:"VulType"`

	// 攻击时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 目标服务器IP
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 目标服务器名称
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

	// 攻击内容
	HttpContent *string `json:"HttpContent,omitempty" name:"HttpContent"`
}

type DeleteAttackLogsRequest struct {
	*tchttp.BaseRequest

	// 日志ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAttackLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashEventsRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBashEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashRulesRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 暴力破解事件Id数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 白名单ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端Uuid。
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

type DeleteMachineResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteMachineTagRequest struct {
	*tchttp.BaseRequest

	// 关联的标签ID
	Rid *uint64 `json:"Rid,omitempty" name:"Rid"`
}

func (r *DeleteMachineTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaliciousRequestsRequest struct {
	*tchttp.BaseRequest

	// 恶意请求记录ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马记录ID数组 (最大100条)
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 删除异地登录事件的方式，可选值："Ids"、"Ip"、"All"，默认为Ids
	DelType *string `json:"DelType,omitempty" name:"DelType"`

	// 异地登录事件ID数组。DelType为Ids或DelType未填时此项必填
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`

	// 异地登录事件的Ip。DelType为Ip时必填
	Ip []*string `json:"Ip,omitempty" name:"Ip"`

	// 主机Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DeleteNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DelType")
	delete(f, "Ids")
	delete(f, "Ip")
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeletePrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivilegeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeRulesRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeletePrivilegeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivilegeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivilegeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivilegeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivilegeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
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
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteReverseShellRulesRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteReverseShellRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReverseShellRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteReverseShellRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteReverseShellRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReverseShellRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DeleteSearchTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：成功，非0：失败
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteTagsRequest struct {
	*tchttp.BaseRequest

	// 标签ID
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWebPageEventLogRequest struct {
	*tchttp.BaseRequest
}

func (r *DeleteWebPageEventLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebPageEventLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWebPageEventLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWebPageEventLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWebPageEventLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebPageEventLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatisticsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号用户名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccountStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 帐号统计列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 帐号统计列表。
		AccountStatistics []*AccountStatistics `json:"AccountStatistics,omitempty" name:"AccountStatistics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Uuid表示，查询该主机下列表信息。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Username表示，查询该用户名下列表信息。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	// <li>Privilege - String - 是否必填：否 - 帐号类型（ORDINARY: 普通帐号 | SUPPER: 超级管理员帐号）</li>
	// <li>MachineIp - String - 是否必填：否 - 主机内网IP</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Username")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 帐号列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 帐号数据列表。
		Accounts []*Account `json:"Accounts,omitempty" name:"Accounts"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAssetInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主机数
		MachineCount *uint64 `json:"MachineCount,omitempty" name:"MachineCount"`

		// 账号数
		AccountCount *uint64 `json:"AccountCount,omitempty" name:"AccountCount"`

		// 端口数
		PortCount *uint64 `json:"PortCount,omitempty" name:"PortCount"`

		// 进程数
		ProcessCount *uint64 `json:"ProcessCount,omitempty" name:"ProcessCount"`

		// 软件数
		SoftwareCount *uint64 `json:"SoftwareCount,omitempty" name:"SoftwareCount"`

		// 数据库数
		DatabaseCount *uint64 `json:"DatabaseCount,omitempty" name:"DatabaseCount"`

		// Web应用数
		WebAppCount *uint64 `json:"WebAppCount,omitempty" name:"WebAppCount"`

		// Web框架数
		WebFrameCount *uint64 `json:"WebFrameCount,omitempty" name:"WebFrameCount"`

		// Web服务数
		WebServiceCount *uint64 `json:"WebServiceCount,omitempty" name:"WebServiceCount"`

		// Web站点数
		WebLocationCount *uint64 `json:"WebLocationCount,omitempty" name:"WebLocationCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetRecentMachineInfoRequest struct {
	*tchttp.BaseRequest

	// 开始时间，如：2020-09-22
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 结束时间，如：2020-09-22
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeAssetRecentMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetRecentMachineInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetRecentMachineInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetRecentMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalList []*AssetKeyVal `json:"TotalList,omitempty" name:"TotalList"`

		// 在线数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		LiveList []*AssetKeyVal `json:"LiveList,omitempty" name:"LiveList"`

		// 离线数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		OfflineList []*AssetKeyVal `json:"OfflineList,omitempty" name:"OfflineList"`

		// 风险数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		RiskList []*AssetKeyVal `json:"RiskList,omitempty" name:"RiskList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetRecentMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetRecentMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogInfoRequest struct {
	*tchttp.BaseRequest

	// 日志ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAttackLogInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackLogInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志ID
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 主机ID
		Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

		// 攻击来源端口
		SrcPort *uint64 `json:"SrcPort,omitempty" name:"SrcPort"`

		// 攻击来源IP
		SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

		// 攻击目标端口
		DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

		// 攻击目标IP
		DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

		// 攻击方法
		HttpMethod *string `json:"HttpMethod,omitempty" name:"HttpMethod"`

		// 攻击目标主机
		HttpHost *string `json:"HttpHost,omitempty" name:"HttpHost"`

		// 攻击头信息
		HttpHead *string `json:"HttpHead,omitempty" name:"HttpHead"`

		// 攻击者浏览器标识
		HttpUserAgent *string `json:"HttpUserAgent,omitempty" name:"HttpUserAgent"`

		// 请求源
		HttpReferer *string `json:"HttpReferer,omitempty" name:"HttpReferer"`

		// 威胁类型
		VulType *string `json:"VulType,omitempty" name:"VulType"`

		// 攻击路径
		HttpCgi *string `json:"HttpCgi,omitempty" name:"HttpCgi"`

		// 攻击参数
		HttpParam *string `json:"HttpParam,omitempty" name:"HttpParam"`

		// 攻击时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 攻击内容
		HttpContent *string `json:"HttpContent,omitempty" name:"HttpContent"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackLogInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	// <li>VulType - String 威胁类型 - 是否必填: 否</li>
	// <li>SrcIp - String 攻击源IP - 是否必填: 否</li>
	// <li>DstIp - String 攻击目标IP - 是否必填: 否</li>
	// <li>SrcPort - String 攻击源端口 - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 主机安全客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 云主机机器ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Uuid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		AttackLogs []*DefendAttackLog `json:"AttackLogs,omitempty" name:"AttackLogs"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackVulTypeListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAttackVulTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackVulTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackVulTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttackVulTypeListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 威胁类型列表
		List []*string `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttackVulTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackVulTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机内网IP)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBashEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 高危命令事件列表
		List []*BashEvent `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashRulesRequest struct {
	*tchttp.BaseRequest

	// 0-系统规则; 1-用户规则
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(规则名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表内容
		List []*BashRule `json:"List,omitempty" name:"List"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttackListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Status - String - 是否必填：否 - 状态筛选：失败：FAILED 成功：SUCCESS</li>
	// <li>UserName - String - 是否必填：否 - UserName筛选</li>
	// <li>SrcIp - String - 是否必填：否 - 来源ip筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选，开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选，结束时间</li>
	// <li>Banned - String - 是否必填：否 - 阻断状态筛选，多个用","分割：0-未阻断（全局ZK开关关闭），82-未阻断(非专业版)，83-未阻断(已加白名单)，1-已阻断，2-未阻断-程序异常，3-未阻断-内网攻击暂不支持阻断，4-未阻断-安平暂不支持阻断</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeBruteAttackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttackListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBruteAttackListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBruteAttackListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 密码破解列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		BruteAttackList []*BruteAttackInfo `json:"BruteAttackList,omitempty" name:"BruteAttackList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBruteAttackListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBruteAttackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentStatisticsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// ComponentName - String - 是否必填：否 - 组件名称
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeComponentStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComponentStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组件统计列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 组件统计列表数据数组。
		ComponentStatistics []*ComponentStatistics `json:"ComponentStatistics,omitempty" name:"ComponentStatistics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeComponentStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComponentStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DescribeESAggregationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ES聚合结果JSON
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeESHitsRequest struct {
	*tchttp.BaseRequest

	// ES查询条件JSON
	Query *string `json:"Query,omitempty" name:"Query"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100。
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

type DescribeESHitsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ES查询结果JSON
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeExportMachinesRequest struct {
	*tchttp.BaseRequest

	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 机器所属业务ID列表
	ProjectIds []*uint64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
}

func (r *DescribeExportMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ProjectIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralStatRequest struct {
	*tchttp.BaseRequest

	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`
}

func (r *DescribeGeneralStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralStatResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云主机总数
		MachinesAll *uint64 `json:"MachinesAll,omitempty" name:"MachinesAll"`

		// 云主机没有安装主机安全客户端的总数
		MachinesUninstalled *uint64 `json:"MachinesUninstalled,omitempty" name:"MachinesUninstalled"`

		// 主机安全客户端总数的总数
		AgentsAll *uint64 `json:"AgentsAll,omitempty" name:"AgentsAll"`

		// 主机安全客户端在线的总数
		AgentsOnline *uint64 `json:"AgentsOnline,omitempty" name:"AgentsOnline"`

		// 主机安全客户端离线的总数
		AgentsOffline *uint64 `json:"AgentsOffline,omitempty" name:"AgentsOffline"`

		// 主机安全客户端专业版的总数
		AgentsPro *uint64 `json:"AgentsPro,omitempty" name:"AgentsPro"`

		// 主机安全客户端基础版的总数
		AgentsBasic *uint64 `json:"AgentsBasic,omitempty" name:"AgentsBasic"`

		// 7天内到期的预付费专业版总数
		AgentsProExpireWithInSevenDays *uint64 `json:"AgentsProExpireWithInSevenDays,omitempty" name:"AgentsProExpireWithInSevenDays"`

		// 风险主机总数
		RiskMachine *uint64 `json:"RiskMachine,omitempty" name:"RiskMachine"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryAccountsRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Username - String - 是否必填：否 - 帐号名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeHistoryAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHistoryAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 帐号变更历史列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 帐号变更历史数据数组。
		HistoryAccounts []*HistoryAccount `json:"HistoryAccounts,omitempty" name:"HistoryAccounts"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeHistoryServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHistoryServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHistoryServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 1 可购买 2 只能升降配 3 只能跳到续费管理页
		BuyStatus *uint64 `json:"BuyStatus,omitempty" name:"BuyStatus"`

		// 用户已购容量 单位 G
		InquireNum *uint64 `json:"InquireNum,omitempty" name:"InquireNum"`

		// 到期时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 是否自动续费,0 初始值, 1 开通 2 没开通
		IsAutoOpenRenew *uint64 `json:"IsAutoOpenRenew,omitempty" name:"IsAutoOpenRenew"`

		// 资源ID
		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

		// 0 没开通 1 正常 2隔离 3销毁
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHistoryServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportMachineInfoRequest struct {
	*tchttp.BaseRequest

	// 服务器内网IP（默认）/ 服务器名称 / 服务器ID 数组
	MachineList []*string `json:"MachineList,omitempty" name:"MachineList"`

	// 批量导入的数据类型：Ip、Name、Id 三选一
	ImportType *string `json:"ImportType,omitempty" name:"ImportType"`

	// 是否仅支持专业版机器的查询（true：仅专业版   false：专业版+基础版）
	IsQueryProMachine *bool `json:"IsQueryProMachine,omitempty" name:"IsQueryProMachine"`
}

func (r *DescribeImportMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportMachineInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineList")
	delete(f, "ImportType")
	delete(f, "IsQueryProMachine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImportMachineInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImportMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 有效的机器信息列表：机器名称、机器公网/内网ip、机器标签
	// 注意：此字段可能返回 null，表示取不到有效值。
		EffectiveMachineInfoList []*EffectiveMachineInfo `json:"EffectiveMachineInfoList,omitempty" name:"EffectiveMachineInfoList"`

		// 用户批量导入失败的机器列表（比如机器不存在等...）
	// 注意：此字段可能返回 null，表示取不到有效值。
		InvalidMachineList []*string `json:"InvalidMachineList,omitempty" name:"InvalidMachineList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DescribeIndexListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ES 索引信息
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeLogStorageStatisticResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总容量（单位：GB）
		TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`

		// 已使用容量（单位：GB）
		UsedSize *uint64 `json:"UsedSize,omitempty" name:"UsedSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeLoginWhiteListRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLoginWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoginWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLoginWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 异地登录白名单数组
		LoginWhiteLists []*LoginWhiteLists `json:"LoginWhiteLists,omitempty" name:"LoginWhiteLists"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLoginWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineInfoRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// cvm id， quuid、uuid必填一个
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *DescribeMachineInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器ip。
		MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

		// 受云镜保护天数。
		ProtectDays *uint64 `json:"ProtectDays,omitempty" name:"ProtectDays"`

		// 操作系统。
		MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`

		// 主机名称。
		MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

		// 在线状态。
	// <li>ONLINE： 在线</li>
	// <li>OFFLINE：离线</li>
		MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`

		// CVM或BM主机唯一标识。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 主机外网IP。
		MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

		// CVM或BM主机唯一Uuid。
		Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

		// 云镜客户端唯一Uuid。
		Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

		// 是否开通专业版。
	// <li>true：是</li>
	// <li>false：否</li>
		IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

		// 专业版开通时间。
		ProVersionOpenDate *string `json:"ProVersionOpenDate,omitempty" name:"ProVersionOpenDate"`

		// 云主机类型。
	// <li>CVM: 虚拟主机</li>
	// <li>BM: 黑石物理机</li>
		MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

		// 机器所属地域。如：ap-guangzhou，ap-shanghai
		MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

		// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>
		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

		// 免费木马剩余检测数量。
		FreeMalwaresLeft *uint64 `json:"FreeMalwaresLeft,omitempty" name:"FreeMalwaresLeft"`

		// 免费漏洞剩余检测数量。
		FreeVulsLeft *uint64 `json:"FreeVulsLeft,omitempty" name:"FreeVulsLeft"`

		// agent版本号
		AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`

		// 专业版到期时间(仅预付费)
		ProVersionDeadline *string `json:"ProVersionDeadline,omitempty" name:"ProVersionDeadline"`

		// 是否有资产扫描记录，0无，1有
		HasAssetScan *uint64 `json:"HasAssetScan,omitempty" name:"HasAssetScan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineListRequest struct {
	*tchttp.BaseRequest

	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMachineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主机列表
		Machines []*Machine `json:"Machines,omitempty" name:"Machines"`

		// 主机数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineOsListRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMachineOsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineOsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineOsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineOsListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作系统列表
		List []*OsName `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineOsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineOsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMachineRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CVM 云服务器地域列表
		CVM []*RegionInfo `json:"CVM,omitempty" name:"CVM"`

		// BM 黑石机器地域列表
		BM []*RegionInfo `json:"BM,omitempty" name:"BM"`

		// LH 轻量应用服务器地域列表
		LH []*RegionInfo `json:"LH,omitempty" name:"LH"`

		// ECM 边缘计算服务器地域列表
		ECM []*RegionInfo `json:"ECM,omitempty" name:"ECM"`

		// Other 混合云地域列表
		Other []*RegionInfo `json:"Other,omitempty" name:"Other"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字 </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线/关机 | ONLINE: 在线 | UNINSTALLED：未安装 | AGENT_OFFLINE 离线| AGENT_SHUTDOWN 已关机）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// <li>Risk - String 是否必填: 否 - 风险主机( yes ) </li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 机器所属业务ID列表
	ProjectIds []*uint64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
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
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ProjectIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主机列表
		Machines []*Machine `json:"Machines,omitempty" name:"Machines"`

		// 主机数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeMalwareInfoRequest struct {
	*tchttp.BaseRequest

	// 唯一ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeMalwareInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMalwareInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 恶意文件详情信息
		MalwareInfo *MalwareInfo `json:"MalwareInfo,omitempty" name:"MalwareInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareTimingScanSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMalwareTimingScanSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareTimingScanSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMalwareTimingScanSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMalwareTimingScanSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测模式 0 全盘检测  1快速检测
		CheckPattern *uint64 `json:"CheckPattern,omitempty" name:"CheckPattern"`

		// 检测周期 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 检测周期 超时结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 是否全部服务器 1 全部 2 自选
		IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

		// 自选服务器时必须 主机quuid的string数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`

		// 监控模式 0 标准 1深度
		MonitoringPattern *uint64 `json:"MonitoringPattern,omitempty" name:"MonitoringPattern"`

		// 周期 1每天
		Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`

		// 定时检测开关 0 关闭1 开启
		EnableScan *int64 `json:"EnableScan,omitempty" name:"EnableScan"`

		// 唯一ID
		Id *int64 `json:"Id,omitempty" name:"Id"`

		// 实时监控0 关闭 1开启
		RealTimeMonitoring *int64 `json:"RealTimeMonitoring,omitempty" name:"RealTimeMonitoring"`

		// 是否自动隔离：1-是，0-否
		AutoIsolation *uint64 `json:"AutoIsolation,omitempty" name:"AutoIsolation"`

		// 一键扫描超时时长，如：1800秒（s）
		ClickTimeout *uint64 `json:"ClickTimeout,omitempty" name:"ClickTimeout"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMalwareTimingScanSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMalwareTimingScanSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortStatisticsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Port - Uint64 - 是否必填：否 - 端口号</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOpenPortStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpenPortStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOpenPortStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 端口统计列表总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 端口统计数据列表
		OpenPortStatistics []*OpenPortStatistics `json:"OpenPortStatistics,omitempty" name:"OpenPortStatistics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOpenPortStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPortStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatisticsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeOverviewStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOverviewStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器在线数。
		OnlineMachineNum *uint64 `json:"OnlineMachineNum,omitempty" name:"OnlineMachineNum"`

		// 专业服务器数。
		ProVersionMachineNum *uint64 `json:"ProVersionMachineNum,omitempty" name:"ProVersionMachineNum"`

		// 木马文件数。
		MalwareNum *uint64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

		// 异地登录数。
		NonlocalLoginNum *uint64 `json:"NonlocalLoginNum,omitempty" name:"NonlocalLoginNum"`

		// 暴力破解成功数。
		BruteAttackSuccessNum *uint64 `json:"BruteAttackSuccessNum,omitempty" name:"BruteAttackSuccessNum"`

		// 漏洞数。
		VulNum *uint64 `json:"VulNum,omitempty" name:"VulNum"`

		// 安全基线数。
		BaseLineNum *uint64 `json:"BaseLineNum,omitempty" name:"BaseLineNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOverviewStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键词(主机IP)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivilegeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据列表
		List []*PrivilegeEscalationProcess `json:"List,omitempty" name:"List"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeRulesRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrivilegeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivilegeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivilegeRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表内容
		List []*PrivilegeRule `json:"List,omitempty" name:"List"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivilegeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivilegeRulesResponse) FromJsonString(s string) error {
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

		// 后付费昨日扣费
		PostPayCost *uint64 `json:"PostPayCost,omitempty" name:"PostPayCost"`

		// 新增主机是否自动开通专业版
		IsAutoOpenProVersion *bool `json:"IsAutoOpenProVersion,omitempty" name:"IsAutoOpenProVersion"`

		// 开通专业版主机数
		ProVersionNum *uint64 `json:"ProVersionNum,omitempty" name:"ProVersionNum"`

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

type DescribeProcessStatisticsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>ProcessName - String - 是否必填：否 - 进程名</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeProcessStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProcessStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProcessStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 进程统计列表记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 进程统计列表数据数组。
		ProcessStatistics []*ProcessStatistics `json:"ProcessStatistics,omitempty" name:"ProcessStatistics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcessStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProcessStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(主机内网IP|进程名)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表内容
		List []*ReverseShell `json:"List,omitempty" name:"List"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

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

type DescribeReverseShellRulesRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 关键字(进程名称)</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeReverseShellRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReverseShellRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReverseShellRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表内容
		List []*ReverseShellRule `json:"List,omitempty" name:"List"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReverseShellRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReverseShellRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsListRequest struct {
	*tchttp.BaseRequest

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>Url - String - 是否必填：否 - Url筛选</li>
	// <li>Status - String - 是否必填：否 - 状态筛选0:待处理；2:信任；3:不信任</li>
	// <li>MergeBeginTime - String - 是否必填：否 - 最近访问开始时间</li>
	// <li>MergeEndTime - String - 是否必填：否 - 最近访问结束时间</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序方式：根据请求次数排序：asc-升序/desc-降序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段：AccessCount-请求次数
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeRiskDnsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDnsListRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskDnsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRiskDnsListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 恶意请求列表数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		RiskDnsList []*RiskDnsList `json:"RiskDnsList,omitempty" name:"RiskDnsList"`

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRiskDnsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDnsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanMalwareScheduleRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScanMalwareScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanMalwareScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanMalwareScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanMalwareScheduleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 扫描进度（单位：%）
		Schedule *int64 `json:"Schedule,omitempty" name:"Schedule"`

		// 风险文件数,当进度满了以后才有该值
		RiskFileNumber *int64 `json:"RiskFileNumber,omitempty" name:"RiskFileNumber"`

		// 是否正在扫描中
		IsSchedule *bool `json:"IsSchedule,omitempty" name:"IsSchedule"`

		// 0 从未扫描过、 1 扫描中、 2扫描完成、 3停止中、 4停止完成
		ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanMalwareScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanMalwareScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanTaskDetailsRequest struct {
	*tchttp.BaseRequest

	// 模块类型 当前提供 Malware 木马 , Vul 漏洞 , Baseline 基线
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 过滤参数
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeScanTaskDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleType")
	delete(f, "TaskId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanTaskDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 扫描任务信息列表
		ScanTaskDetailList []*ScanTaskDetails `json:"ScanTaskDetailList,omitempty" name:"ScanTaskDetailList"`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 扫描机器总数
		ScanMachineCount *uint64 `json:"ScanMachineCount,omitempty" name:"ScanMachineCount"`

		// 发现风险机器数
		RiskMachineCount *uint64 `json:"RiskMachineCount,omitempty" name:"RiskMachineCount"`

		// 扫描开始时间
		ScanBeginTime *string `json:"ScanBeginTime,omitempty" name:"ScanBeginTime"`

		// 扫描结束时间
		ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`

		// 检测时间
		ScanTime *uint64 `json:"ScanTime,omitempty" name:"ScanTime"`

		// 扫描进度
		ScanProgress *uint64 `json:"ScanProgress,omitempty" name:"ScanProgress"`

		// 扫描剩余时间
		ScanLeftTime *uint64 `json:"ScanLeftTime,omitempty" name:"ScanLeftTime"`

		// 扫描内容
		ScanContent []*string `json:"ScanContent,omitempty" name:"ScanContent"`

		// 漏洞信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		VulInfo []*VulDetailInfo `json:"VulInfo,omitempty" name:"VulInfo"`

		// 风险事件个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		RiskEventCount *uint64 `json:"RiskEventCount,omitempty" name:"RiskEventCount"`

		// 0一键检测 1定时检测
	// 注意：此字段可能返回 null，表示取不到有效值。
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// 任务是否全部正在被停止 ture是
	// 注意：此字段可能返回 null，表示取不到有效值。
		StoppingAll *bool `json:"StoppingAll,omitempty" name:"StoppingAll"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanTaskDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanVulSettingRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeScanVulSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanVulSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanVulSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanVulSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 漏洞类型：1: web应用漏洞 2:系统组件漏洞 (多选英文逗号分隔)
		VulCategories *string `json:"VulCategories,omitempty" name:"VulCategories"`

		// 危害等级：1-低危；2-中危；3-高危；4-严重 (多选英文逗号分隔)
		VulLevels *string `json:"VulLevels,omitempty" name:"VulLevels"`

		// 定期检测间隔时间（天）
		TimerInterval *uint64 `json:"TimerInterval,omitempty" name:"TimerInterval"`

		// 定期检测时间，如：00:00
		TimerTime *string `json:"TimerTime,omitempty" name:"TimerTime"`

		// 是否紧急漏洞：0-否 1-是
		VulEmergency *uint64 `json:"VulEmergency,omitempty" name:"VulEmergency"`

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 是否开启
		EnableScan *uint64 `json:"EnableScan,omitempty" name:"EnableScan"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 一键扫描超时时长，如：1800秒（s）
		ClickTimeout *uint64 `json:"ClickTimeout,omitempty" name:"ClickTimeout"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScanVulSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanVulSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DescribeSearchExportListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出的任务号
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSearchLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 历史搜索记录
		Data []*string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSearchTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 模板列表
		List []*SearchTemplate `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSecurityDynamicsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSecurityDynamicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityDynamicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityDynamicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全事件消息数组。
		SecurityDynamics []*SecurityDynamic `json:"SecurityDynamics,omitempty" name:"SecurityDynamics"`

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityDynamicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityDynamicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityEventsCntRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSecurityEventsCntRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityEventsCntRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityEventsCntRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityEventsCntResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 木马文件相关风险事件
		Malware *SecurityEventInfo `json:"Malware,omitempty" name:"Malware"`

		// 登录审计相关风险事件
		HostLogin *SecurityEventInfo `json:"HostLogin,omitempty" name:"HostLogin"`

		// 密码破解相关风险事件
		BruteAttack *SecurityEventInfo `json:"BruteAttack,omitempty" name:"BruteAttack"`

		// 恶意请求相关风险事件
		RiskDns *SecurityEventInfo `json:"RiskDns,omitempty" name:"RiskDns"`

		// 高危命令相关风险事件
		Bash *SecurityEventInfo `json:"Bash,omitempty" name:"Bash"`

		// 本地提权相关风险事件
		PrivilegeRules *SecurityEventInfo `json:"PrivilegeRules,omitempty" name:"PrivilegeRules"`

		// 反弹Shell相关风险事件
		ReverseShell *SecurityEventInfo `json:"ReverseShell,omitempty" name:"ReverseShell"`

		// 系统组件相关风险事件
		SysVul *SecurityEventInfo `json:"SysVul,omitempty" name:"SysVul"`

		// Web应用漏洞相关风险事件
		WebVul *SecurityEventInfo `json:"WebVul,omitempty" name:"WebVul"`

		// 应急漏洞相关风险事件
		EmergencyVul *SecurityEventInfo `json:"EmergencyVul,omitempty" name:"EmergencyVul"`

		// 安全基线相关风险事件
		BaseLine *SecurityEventInfo `json:"BaseLine,omitempty" name:"BaseLine"`

		// 攻击检测相关风险事件
		AttackLogs *SecurityEventInfo `json:"AttackLogs,omitempty" name:"AttackLogs"`

		// 受影响机器数
		EffectMachineCount *uint64 `json:"EffectMachineCount,omitempty" name:"EffectMachineCount"`

		// 所有事件总数
		EventsCount *uint64 `json:"EventsCount,omitempty" name:"EventsCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityEventsCntResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityEventsCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendsRequest struct {
	*tchttp.BaseRequest

	// 开始时间，如：2021-07-10
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 结束时间，如：2021-07-10
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeSecurityTrendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityTrendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityTrendsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 木马事件统计数据数组。
		Malwares []*SecurityTrend `json:"Malwares,omitempty" name:"Malwares"`

		// 异地登录事件统计数据数组。
		NonLocalLoginPlaces []*SecurityTrend `json:"NonLocalLoginPlaces,omitempty" name:"NonLocalLoginPlaces"`

		// 密码破解事件统计数据数组。
		BruteAttacks []*SecurityTrend `json:"BruteAttacks,omitempty" name:"BruteAttacks"`

		// 漏洞统计数据数组。
		Vuls []*SecurityTrend `json:"Vuls,omitempty" name:"Vuls"`

		// 基线统计数据数组。
		BaseLines []*SecurityTrend `json:"BaseLines,omitempty" name:"BaseLines"`

		// 恶意请求统计数据数组。
		MaliciousRequests []*SecurityTrend `json:"MaliciousRequests,omitempty" name:"MaliciousRequests"`

		// 高危命令统计数据数组。
		HighRiskBashs []*SecurityTrend `json:"HighRiskBashs,omitempty" name:"HighRiskBashs"`

		// 反弹shell统计数据数组。
		ReverseShells []*SecurityTrend `json:"ReverseShells,omitempty" name:"ReverseShells"`

		// 本地提权统计数据数组。
		PrivilegeEscalations []*SecurityTrend `json:"PrivilegeEscalations,omitempty" name:"PrivilegeEscalations"`

		// 网络攻击统计数据数组。
		CyberAttacks []*SecurityTrend `json:"CyberAttacks,omitempty" name:"CyberAttacks"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSecurityTrendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityTrendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagMachinesRequest struct {
	*tchttp.BaseRequest

	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTagMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表数据
		List []*TagMachine `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest

	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 过滤条件。
	// <li>Keywords - String - 是否必填：否 - 查询关键字(机器名称/机器IP </li>
	// <li>Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线 | UNINSTALLED：未安装 | SHUTDOWN 已关机）</li>
	// <li>Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）</li>
	// <li>Risk - String 是否必填: 否 - 风险主机( yes ) </li>
	// <li>Os -String 是否必填: 否 - 操作系统( DescribeMachineOsList 接口 值 )
	// 每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表信息
		List []*Tag `json:"List,omitempty" name:"List"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsualLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// 云镜客户端UUID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

func (r *DescribeUsualLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsualLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsualLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 常用登录地数组
		UsualLoginPlaces []*UsualPlace `json:"UsualLoginPlaces,omitempty" name:"UsualLoginPlaces"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsualLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsualLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebPageGeneralizeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWebPageGeneralizeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageGeneralizeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebPageGeneralizeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebPageGeneralizeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 防护监测 0 未开启 1 已开启 2 异常
		ProtectMonitor *uint64 `json:"ProtectMonitor,omitempty" name:"ProtectMonitor"`

		// 防护目录数
		ProtectDirNum *uint64 `json:"ProtectDirNum,omitempty" name:"ProtectDirNum"`

		// 防护文件数
		ProtectFileNum *uint64 `json:"ProtectFileNum,omitempty" name:"ProtectFileNum"`

		// 篡改文件数
		TamperFileNum *uint64 `json:"TamperFileNum,omitempty" name:"TamperFileNum"`

		// 篡改数
		TamperNum *uint64 `json:"TamperNum,omitempty" name:"TamperNum"`

		// 今日防护数
		ProtectToday *uint64 `json:"ProtectToday,omitempty" name:"ProtectToday"`

		// 防护主机数
		ProtectHostNum *uint64 `json:"ProtectHostNum,omitempty" name:"ProtectHostNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebPageGeneralizeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebPageGeneralizeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditBashRulesRequest struct {
	*tchttp.BaseRequest

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 正则表达式
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// 规则ID（新增时不填）
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID数组
	Uuids []*string `json:"Uuids,omitempty" name:"Uuids"`

	// 主机IP
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 危险等级(0:无，1: 高危 2:中危 3: 低危)
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 是否全局规则(默认否)：1-全局，0-非全局
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 0=黑名单， 1=白名单
	White *uint64 `json:"White,omitempty" name:"White"`

	// 事件列表点击“加入白名单”时,需要传EventId 事件的id
	EventId *uint64 `json:"EventId,omitempty" name:"EventId"`

	// 是否处理旧事件为白名单 0=不处理 1=处理
	DealOldEvents *uint64 `json:"DealOldEvents,omitempty" name:"DealOldEvents"`
}

func (r *EditBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Rule")
	delete(f, "Id")
	delete(f, "Uuids")
	delete(f, "HostIp")
	delete(f, "Level")
	delete(f, "IsGlobal")
	delete(f, "White")
	delete(f, "EventId")
	delete(f, "DealOldEvents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EditBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EditTagsRequest struct {
	*tchttp.BaseRequest

	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// CVM主机ID
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *EditTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Id")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EditTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EffectiveMachineInfo struct {

	// 机器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 机器公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachinePublicIp *string `json:"MachinePublicIp,omitempty" name:"MachinePublicIp"`

	// 机器内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachinePrivateIp *string `json:"MachinePrivateIp,omitempty" name:"MachinePrivateIp"`

	// 机器标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineTag []*MachineTag `json:"MachineTag,omitempty" name:"MachineTag"`

	// 机器Quuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 云镜Uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
}

type ExportAssetCoreModuleListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>Name- string - 是否必填：否 - 包名</li>
	// <li>User- string - 是否必填：否 - 用户</li>
	Filters []*AssetFilters `json:"Filters,omitempty" name:"Filters"`

	// 排序方式，asc升序 或 desc降序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序依据:Size,ProcessCount,ModuleCount
	By *string `json:"By,omitempty" name:"By"`

	// 服务器Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 服务器Quuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAssetCoreModuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAssetCoreModuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "Uuid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportAssetCoreModuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportAssetCoreModuleListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步下载任务ID，需要配合ExportTasks接口使用
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAssetCoreModuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAssetCoreModuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportAttackLogsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// <li>HttpMethod - String - 是否必填：否 - 攻击方法(POST|GET)</li>
	// <li>DateRange - String - 是否必填：否 - 时间范围(存储最近3个月的数据)，如最近一个月["2019-11-17", "2019-12-17"]</li>
	// <li>VulType - String 威胁类型 - 是否必填: 否</li>
	// <li>SrcIp - String 攻击源IP - 是否必填: 否</li>
	// <li>DstIp - String 攻击目标IP - 是否必填: 否</li>
	// <li>SrcPort - String 攻击源端口 - 是否必填: 否</li>
	// <li>DstPort - String 攻击目标端口 - 是否必填: 否</li>
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`

	// 主机安全客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 云主机机器ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ExportAttackLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttackLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Uuid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportAttackLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportAttackLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 导出任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportAttackLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttackLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBashEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBashEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBashEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportBashEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 导出任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBashEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBashEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportBruteAttacksRequest struct {
	*tchttp.BaseRequest

	// 过滤参数
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportBruteAttacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBruteAttacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportBruteAttacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 导出任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportBruteAttacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBruteAttacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportMaliciousRequestsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportMaliciousRequestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportMaliciousRequestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportMaliciousRequestsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportMaliciousRequestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMaliciousRequestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportMalwaresRequest struct {
	*tchttp.BaseRequest

	// 限制条数,默认10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量 默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤参数。
	// <li>IpOrAlias - String - 是否必填：否 - 主机ip或别名筛选</li>
	// <li>FilePath - String - 是否必填：否 - 路径筛选</li>
	// <li>VirusName - String - 是否必填：否 - 描述筛选</li>
	// <li>CreateBeginTime - String - 是否必填：否 - 创建时间筛选-开始时间</li>
	// <li>CreateEndTime - String - 是否必填：否 - 创建时间筛选-结束时间</li>
	// <li>Status - String - 是否必填：否 - 状态筛选</li>
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 任务id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportNonlocalLoginPlacesRequest struct {
	*tchttp.BaseRequest

	// <li>Status - int - 是否必填：否 - 状态筛选1:正常登录；2：异地登录</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportNonlocalLoginPlacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportNonlocalLoginPlacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportNonlocalLoginPlacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 导出任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportNonlocalLoginPlacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportNonlocalLoginPlacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportPrivilegeEventsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportPrivilegeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportPrivilegeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportPrivilegeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportPrivilegeEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 导出任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportPrivilegeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportPrivilegeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportReverseShellEventsRequest struct {
	*tchttp.BaseRequest

	// 过滤参数
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`
}

func (r *ExportReverseShellEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportReverseShellEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportReverseShellEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportReverseShellEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 任务id
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportReverseShellEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportReverseShellEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportTasksRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// PENDING：正在生成下载链接，FINISHED：下载链接已生成，ERROR：网络异常等异常情况
		Status *string `json:"Status,omitempty" name:"Status"`

		// 下载链接
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDetectionExcelRequest struct {
	*tchttp.BaseRequest

	// 本次漏洞检测任务id（不同于出参的导出本次漏洞检测Excel的任务Id）
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExportVulDetectionExcelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulDetectionExcelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVulDetectionExcelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDetectionExcelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 导出本次漏洞检测Excel的任务Id（不同于入参的本次漏洞检测任务id）
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDetectionExcelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulDetectionExcelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDetectionReportRequest struct {
	*tchttp.BaseRequest

	// 漏洞扫描任务id（不同于出参的导出检测报告的任务Id）
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 过滤参数
	Filters []*Filters `json:"Filters,omitempty" name:"Filters"`

	// 需要返回的数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ExportVulDetectionReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulDetectionReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportVulDetectionReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExportVulDetectionReportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出文件下载链接地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 导出检测报告的任务Id（不同于入参的漏洞扫描任务id）
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportVulDetectionReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportVulDetectionReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 模糊搜索
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type Filters struct {

	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否模糊匹配，前端框架会带上，可以不管
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type HistoryAccount struct {

	// 唯一ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一Uuid。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机内网IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机名。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 帐号名。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 帐号变更类型。
	// <li>CREATE：表示新增帐号</li>
	// <li>MODIFY：表示修改帐号</li>
	// <li>DELETE：表示删除帐号</li>
	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`

	// 变更时间。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type IgnoreImpactedHostsRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID数组。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *IgnoreImpactedHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IgnoreImpactedHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IgnoreImpactedHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IgnoreImpactedHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreImpactedHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceOpenProVersionPrepaidRequest struct {
	*tchttp.BaseRequest

	// 预付费模式(包年包月)参数设置。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 需要开通专业版机器列表数组。
	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines"`
}

func (r *InquiryPriceOpenProVersionPrepaidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceOpenProVersionPrepaidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargePrepaid")
	delete(f, "Machines")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceOpenProVersionPrepaidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceOpenProVersionPrepaidResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 预支费用的原价，单位：元。
		OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 预支费用的折扣价，单位：元。
		DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceOpenProVersionPrepaidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceOpenProVersionPrepaidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginWhiteLists struct {

	// 记录ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 白名单地域
	Places []*Place `json:"Places,omitempty" name:"Places"`

	// 白名单用户（多个用户逗号隔开）
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 白名单IP（多个IP逗号隔开）
	SrcIp *string `json:"SrcIp,omitempty" name:"SrcIp"`

	// 是否为全局规则
	IsGlobal *bool `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 创建白名单时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改白名单时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 机器名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 机器IP
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type Machine struct {

	// 主机名称。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 主机系统。
	MachineOs *string `json:"MachineOs,omitempty" name:"MachineOs"`

	// 主机状态。
	// <li>OFFLINE: 离线  </li>
	// <li>ONLINE: 在线</li>
	// <li>SHUTDOWN: 已关机</li>
	MachineStatus *string `json:"MachineStatus,omitempty" name:"MachineStatus"`

	// 云镜客户端唯一Uuid，若客户端长时间不在线将返回空字符。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// CVM或BM机器唯一Uuid。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 漏洞数。
	VulNum *int64 `json:"VulNum,omitempty" name:"VulNum"`

	// 主机IP。
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 是否是专业版。
	// <li>true： 是</li>
	// <li>false：否</li>
	IsProVersion *bool `json:"IsProVersion,omitempty" name:"IsProVersion"`

	// 主机外网IP。
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// 主机状态。
	// <li>POSTPAY: 表示后付费，即按量计费  </li>
	// <li>PREPAY: 表示预付费，即包年包月</li>
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 木马数。
	MalwareNum *int64 `json:"MalwareNum,omitempty" name:"MalwareNum"`

	// 标签信息
	Tag []*MachineTag `json:"Tag,omitempty" name:"Tag"`

	// 基线风险数。
	BaselineNum *int64 `json:"BaselineNum,omitempty" name:"BaselineNum"`

	// 网络风险数。
	CyberAttackNum *int64 `json:"CyberAttackNum,omitempty" name:"CyberAttackNum"`

	// 风险状态。
	// <li>SAFE：安全</li>
	// <li>RISK：风险</li>
	// <li>UNKNOWN：未知</li>
	SecurityStatus *string `json:"SecurityStatus,omitempty" name:"SecurityStatus"`

	// 入侵事件数
	InvasionNum *int64 `json:"InvasionNum,omitempty" name:"InvasionNum"`

	// 地域信息
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`

	// 实例状态 TERMINATED_PRO_VERSION 已销毁
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// 防篡改 授权状态 1 授权 0 未授权
	LicenseStatus *uint64 `json:"LicenseStatus,omitempty" name:"LicenseStatus"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否有资产扫描接口，0无，1有
	HasAssetScan *uint64 `json:"HasAssetScan,omitempty" name:"HasAssetScan"`

	// 机器所属专区类型 CVM 云服务器, BM 黑石, ECM 边缘计算, LH 轻量应用服务器 ,Other 混合云专区
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type MachineTag struct {

	// 关联标签ID
	Rid *int64 `json:"Rid,omitempty" name:"Rid"`

	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签ID
	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`
}

type MalwareInfo struct {

	// 病毒名称
	VirusName *string `json:"VirusName,omitempty" name:"VirusName"`

	// 文件大小
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件MD5
	MD5 *string `json:"MD5,omitempty" name:"MD5"`

	// 文件地址
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 首次运行时间
	FileCreateTime *string `json:"FileCreateTime,omitempty" name:"FileCreateTime"`

	// 最近一次运行时间
	FileModifierTime *string `json:"FileModifierTime,omitempty" name:"FileModifierTime"`

	// 危害描述
	HarmDescribe *string `json:"HarmDescribe,omitempty" name:"HarmDescribe"`

	// 建议方案
	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`

	// 服务器名称
	ServersName *string `json:"ServersName,omitempty" name:"ServersName"`

	// 服务器IP
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程ID
	ProcessID *string `json:"ProcessID,omitempty" name:"ProcessID"`

	// 标签特性
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 影响广度 // 暂时不提供
	// 注意：此字段可能返回 null，表示取不到有效值。
	Breadth *string `json:"Breadth,omitempty" name:"Breadth"`

	// 查询热度 // 暂时不提供
	// 注意：此字段可能返回 null，表示取不到有效值。
	Heat *string `json:"Heat,omitempty" name:"Heat"`

	// 唯一ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 首次发现时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近扫描时间
	LatestScanTime *string `json:"LatestScanTime,omitempty" name:"LatestScanTime"`

	// 参考链接
	Reference *string `json:"Reference,omitempty" name:"Reference"`
}

type ModifyAutoOpenProVersionConfigRequest struct {
	*tchttp.BaseRequest

	// 设置自动开通状态。
	// <li>CLOSE：关闭</li>
	// <li>OPEN：打开</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAutoOpenProVersionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoOpenProVersionConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoOpenProVersionConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoOpenProVersionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoOpenProVersionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMalwareTimingScanSettingsRequest struct {
	*tchttp.BaseRequest

	// 检测模式 0 全盘检测  1快速检测
	CheckPattern *uint64 `json:"CheckPattern,omitempty" name:"CheckPattern"`

	// 检测周期 开始时间，如：02:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 检测周期 超时结束时间，如：04:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否全部服务器 1 全部 2 自选
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 定时检测开关 0 关闭 1开启
	EnableScan *uint64 `json:"EnableScan,omitempty" name:"EnableScan"`

	// 监控模式 0 标准 1深度
	MonitoringPattern *uint64 `json:"MonitoringPattern,omitempty" name:"MonitoringPattern"`

	// 扫描周期 默认每天 1
	Cycle *uint64 `json:"Cycle,omitempty" name:"Cycle"`

	// 实时监控 0 关闭 1开启
	RealTimeMonitoring *uint64 `json:"RealTimeMonitoring,omitempty" name:"RealTimeMonitoring"`

	// 自选服务器时必须 主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`

	// 是否自动隔离 1隔离 0 不隔离
	AutoIsolation *uint64 `json:"AutoIsolation,omitempty" name:"AutoIsolation"`
}

func (r *ModifyMalwareTimingScanSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMalwareTimingScanSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CheckPattern")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IsGlobal")
	delete(f, "EnableScan")
	delete(f, "MonitoringPattern")
	delete(f, "Cycle")
	delete(f, "RealTimeMonitoring")
	delete(f, "QuuidList")
	delete(f, "AutoIsolation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMalwareTimingScanSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMalwareTimingScanSettingsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMalwareTimingScanSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMalwareTimingScanSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProVersionRenewFlagRequest struct {
	*tchttp.BaseRequest

	// 自动续费标识。取值范围：
	// <li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li>
	// <li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li>
	// <li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 主机唯一ID，对应CVM的uuid、BM的instanceId。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *ModifyProVersionRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProVersionRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RenewFlag")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProVersionRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyProVersionRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProVersionRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProVersionRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWarningSettingRequest struct {
	*tchttp.BaseRequest

	// 告警设置的修改内容
	WarningObjects []*WarningObject `json:"WarningObjects,omitempty" name:"WarningObjects"`
}

func (r *ModifyWarningSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWarningSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WarningObjects")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWarningSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWarningSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWarningSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWarningSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebPageProtectSettingRequest struct {
	*tchttp.BaseRequest

	// 需要操作的类型1 目录名称 2 防护文件类型
	ModifyType *uint64 `json:"ModifyType,omitempty" name:"ModifyType"`

	// 提交值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 配置对应的protect_path
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *ModifyWebPageProtectSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebPageProtectSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModifyType")
	delete(f, "Value")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWebPageProtectSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebPageProtectSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWebPageProtectSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebPageProtectSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenPortStatistics struct {

	// 端口号
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 主机数量
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type OpenProVersionPrepaidRequest struct {
	*tchttp.BaseRequest

	// 购买相关参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 需要开通专业版主机信息数组。
	Machines []*ProVersionMachine `json:"Machines,omitempty" name:"Machines"`
}

func (r *OpenProVersionPrepaidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionPrepaidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargePrepaid")
	delete(f, "Machines")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenProVersionPrepaidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionPrepaidResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID列表。
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionPrepaidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionPrepaidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionRequest struct {
	*tchttp.BaseRequest

	// 云主机类型。
	// <li>CVM：表示虚拟主机</li>
	// <li>BM:  表示黑石物理机</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器所属地域。
	// 如：ap-guangzhou，ap-shanghai
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM的Uuid
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`

	// 活动ID。
	ActivityId *uint64 `json:"ActivityId,omitempty" name:"ActivityId"`
}

func (r *OpenProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MachineType")
	delete(f, "MachineRegion")
	delete(f, "Quuids")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenProVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OpenProVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OsName struct {

	// 系统名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 操作系统类型枚举值
	MachineOSType *uint64 `json:"MachineOSType,omitempty" name:"MachineOSType"`
}

type Place struct {

	// 城市 ID。
	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`

	// 省份 ID。
	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// 国家ID，暂只支持国内：1。
	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`

	// 位置名称
	Location *string `json:"Location,omitempty" name:"Location"`
}

type PrivilegeEscalationProcess struct {

	// 数据ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`

	// 执行命令
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户组
	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`

	// 进程文件权限
	ProcFilePrivilege *string `json:"ProcFilePrivilege,omitempty" name:"ProcFilePrivilege"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`

	// 父进程用户名
	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`

	// 进程树
	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`

	// 处理状态：0-待处理 2-白名单
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 发生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 机器名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
}

type PrivilegeRule struct {

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 是否S权限
	SMode *uint64 `json:"SMode,omitempty" name:"SMode"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 状态(0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
}

type ProVersionMachine struct {

	// 主机类型。
	// <li>CVM: 云服务器</li>
	// <li>BM: 黑石物理机</li>
	// <li>ECM: 边缘计算服务器</li>
	// <li>LH: 轻量应用服务器</li>
	// <li>Other: 混合云机器</li>
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 主机所在地域。
	// 如：ap-guangzhou、ap-beijing
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机唯一标识Uuid数组。
	// 黑石的InstanceId，CVM的Uuid ,边缘计算的Uuid , 轻量应用服务器的Uuid ,混合云机器的Quuid 。 当前参数最大长度限制20
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

type ProcessStatistics struct {

	// 进程名。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 主机数量。
	MachineNum *uint64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type ProtectHostConfig struct {

	// 机器唯一ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 防护开关 0  关闭 1开启
	ProtectSwitch *uint64 `json:"ProtectSwitch,omitempty" name:"ProtectSwitch"`

	// 自动恢复开关 0 关闭 1开启
	AutoRecovery *uint64 `json:"AutoRecovery,omitempty" name:"AutoRecovery"`
}

type RecoverMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马Id数组（最大100条）
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *RecoverMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RecoverMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 恢复成功id数组，若无则返回空数组
		SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds"`

		// 恢复失败id数组，若无则返回空数组
		FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地域标志，如 ap-guangzhou，ap-shanghai，ap-beijing
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域中文名，如华南地区（广州），华东地区（上海金融），华北地区（北京）
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域代码，如 gz，sh，bj
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 地域英文名
	RegionNameEn *string `json:"RegionNameEn,omitempty" name:"RegionNameEn"`
}

type RenewProVersionRequest struct {
	*tchttp.BaseRequest

	// 购买相关参数。
	ChargePrepaid *ChargePrepaid `json:"ChargePrepaid,omitempty" name:"ChargePrepaid"`

	// 主机唯一ID，对应CVM的uuid、BM的InstanceId。
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`
}

func (r *RenewProVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewProVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargePrepaid")
	delete(f, "Quuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewProVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RenewProVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewProVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewProVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RescanImpactedHostRequest struct {
	*tchttp.BaseRequest

	// 漏洞ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *RescanImpactedHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RescanImpactedHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RescanImpactedHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RescanImpactedHostResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RescanImpactedHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RescanImpactedHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReverseShell struct {

	// ID 主键
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜UUID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机内网IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`

	// 目标IP
	DstIp *string `json:"DstIp,omitempty" name:"DstIp"`

	// 目标端口
	DstPort *uint64 `json:"DstPort,omitempty" name:"DstPort"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程路径
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`

	// 命令详情
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// 执行用户
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 执行用户组
	UserGroup *string `json:"UserGroup,omitempty" name:"UserGroup"`

	// 父进程名
	ParentProcName *string `json:"ParentProcName,omitempty" name:"ParentProcName"`

	// 父进程用户
	ParentProcUser *string `json:"ParentProcUser,omitempty" name:"ParentProcUser"`

	// 父进程用户组
	ParentProcGroup *string `json:"ParentProcGroup,omitempty" name:"ParentProcGroup"`

	// 父进程路径
	ParentProcPath *string `json:"ParentProcPath,omitempty" name:"ParentProcPath"`

	// 处理状态：0-待处理 2-白名单
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 产生时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 主机名
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 进程树
	ProcTree *string `json:"ProcTree,omitempty" name:"ProcTree"`

	// 检测方法
	DetectBy *uint64 `json:"DetectBy,omitempty" name:"DetectBy"`
}

type ReverseShellRule struct {

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 客户端ID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 目标IP
	DestIp *string `json:"DestIp,omitempty" name:"DestIp"`

	// 目标端口
	DestPort *string `json:"DestPort,omitempty" name:"DestPort"`

	// 操作人
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 是否全局规则
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 状态 (0: 有效 1: 无效)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 主机IP
	Hostip *string `json:"Hostip,omitempty" name:"Hostip"`
}

type RiskDnsList struct {

	// 对外访问域名
	Url *string `json:"Url,omitempty" name:"Url"`

	// 访问次数
	AccessCount *uint64 `json:"AccessCount,omitempty" name:"AccessCount"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 进程MD5
	ProcessMd5 *string `json:"ProcessMd5,omitempty" name:"ProcessMd5"`

	// 是否为全局规则，0否，1是
	GlobalRuleId *uint64 `json:"GlobalRuleId,omitempty" name:"GlobalRuleId"`

	// 用户规则id
	UserRuleId *uint64 `json:"UserRuleId,omitempty" name:"UserRuleId"`

	// 状态；0-待处理，2-已加白，3-非信任状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 首次访问时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近访问时间
	MergeTime *string `json:"MergeTime,omitempty" name:"MergeTime"`

	// 唯一 Quuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机ip
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 参考
	Reference *string `json:"Reference,omitempty" name:"Reference"`

	// 命令行
	CmdLine *string `json:"CmdLine,omitempty" name:"CmdLine"`

	// 进程号
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 唯一UUID
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 建议方案
	SuggestScheme *string `json:"SuggestScheme,omitempty" name:"SuggestScheme"`

	// 标签特性
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type ScanAssetRequest struct {
	*tchttp.BaseRequest

	// 资产指纹类型id列表
	AssetTypeIds []*uint64 `json:"AssetTypeIds,omitempty" name:"AssetTypeIds"`

	// Quuid列表
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *ScanAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetTypeIds")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ScanAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanTaskDetails struct {

	// 服务器IP
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// 服务器名称
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 操作系统
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 风险数量
	RiskNum *uint64 `json:"RiskNum,omitempty" name:"RiskNum"`

	// 扫描开始时间
	ScanBeginTime *string `json:"ScanBeginTime,omitempty" name:"ScanBeginTime"`

	// 扫描结束时间
	ScanEndTime *string `json:"ScanEndTime,omitempty" name:"ScanEndTime"`

	// 唯一Uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 唯一Quuid
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 状态码
	Status *string `json:"Status,omitempty" name:"Status"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// id唯一
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 失败详情
	FailType *uint64 `json:"FailType,omitempty" name:"FailType"`
}

type ScanVulAgainRequest struct {
	*tchttp.BaseRequest

	// 漏洞事件id串，多个用英文逗号分隔
	EventIds *string `json:"EventIds,omitempty" name:"EventIds"`

	// 重新检查的机器uuid,多个逗号分隔
	Uuids *string `json:"Uuids,omitempty" name:"Uuids"`
}

func (r *ScanVulAgainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulAgainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventIds")
	delete(f, "Uuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanVulAgainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulAgainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanVulAgainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulAgainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulRequest struct {
	*tchttp.BaseRequest

	// 漏洞类型：1: web应用漏洞 2:系统组件漏洞 (多选英文;分隔)
	VulCategories *string `json:"VulCategories,omitempty" name:"VulCategories"`

	// 危害等级：1-低危；2-中危；3-高危；4-严重 (多选英文;分隔)
	VulLevels *string `json:"VulLevels,omitempty" name:"VulLevels"`

	// 服务器分类：1:专业版服务器；2:自选服务器
	HostType *uint64 `json:"HostType,omitempty" name:"HostType"`

	// 自选服务器时生效，主机quuid的string数组
	QuuidList []*string `json:"QuuidList,omitempty" name:"QuuidList"`

	// 是否是应急漏洞 0 否 1 是
	VulEmergency *uint64 `json:"VulEmergency,omitempty" name:"VulEmergency"`

	// 超时时长 单位秒
	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitempty" name:"TimeoutPeriod"`

	// 需要扫描的漏洞id
	VulIds []*uint64 `json:"VulIds,omitempty" name:"VulIds"`
}

func (r *ScanVulRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulCategories")
	delete(f, "VulLevels")
	delete(f, "HostType")
	delete(f, "QuuidList")
	delete(f, "VulEmergency")
	delete(f, "TimeoutPeriod")
	delete(f, "VulIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanVulRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanVulResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulSettingRequest struct {
	*tchttp.BaseRequest

	// 定期检测间隔时间（天）
	TimerInterval *uint64 `json:"TimerInterval,omitempty" name:"TimerInterval"`

	// 漏洞类型：1: web应用漏洞 2:系统组件漏洞, 以数组方式传参[1,2]
	VulCategories []*uint64 `json:"VulCategories,omitempty" name:"VulCategories"`

	// 危害等级：1-低危；2-中危；3-高危；4-严重,以数组方式传参[1,2,3]
	VulLevels []*uint64 `json:"VulLevels,omitempty" name:"VulLevels"`

	// 定期检测时间，如：02:10:50
	TimerTime *string `json:"TimerTime,omitempty" name:"TimerTime"`

	// 是否是应急漏洞 0 否 1 是
	VulEmergency *uint64 `json:"VulEmergency,omitempty" name:"VulEmergency"`

	// 扫描开始时间，如：00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 扫描结束时间，如：08:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否开启扫描 1开启 0不开启
	EnableScan *uint64 `json:"EnableScan,omitempty" name:"EnableScan"`
}

func (r *ScanVulSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimerInterval")
	delete(f, "VulCategories")
	delete(f, "VulLevels")
	delete(f, "TimerTime")
	delete(f, "VulEmergency")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "EnableScan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanVulSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ScanVulSettingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScanVulSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVulSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type SecurityDynamic struct {

	// 云镜客户端UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 安全事件发生时间。
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// 安全事件类型。
	// <li>MALWARE：木马事件</li>
	// <li>NON_LOCAL_LOGIN：异地登录</li>
	// <li>BRUTEATTACK_SUCCESS：密码破解成功</li>
	// <li>VUL：漏洞</li>
	// <li>BASELINE：安全基线</li>
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 安全事件消息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 安全事件等级。
	// <li>RISK: 严重</li>
	// <li>HIGH: 高危</li>
	// <li>NORMAL: 中危</li>
	// <li>LOW: 低危</li>
	SecurityLevel *string `json:"SecurityLevel,omitempty" name:"SecurityLevel"`
}

type SecurityEventInfo struct {

	// 安全事件数
	EventCnt *uint64 `json:"EventCnt,omitempty" name:"EventCnt"`

	// 受影响机器数
	UuidCnt *uint64 `json:"UuidCnt,omitempty" name:"UuidCnt"`
}

type SecurityTrend struct {

	// 事件时间。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 事件数量。
	EventNum *uint64 `json:"EventNum,omitempty" name:"EventNum"`
}

type SeparateMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马事件ID数组。(最大100条)
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *SeparateMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SeparateMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SeparateMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 隔离成功的id数组，若无则返回空数组
		SuccessIds []*uint64 `json:"SuccessIds,omitempty" name:"SuccessIds"`

		// 隔离失败的id数组，若无则返回空数组
		FailedIds []*uint64 `json:"FailedIds,omitempty" name:"FailedIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SeparateMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SeparateMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBashEventsStatusRequest struct {
	*tchttp.BaseRequest

	// ID数组，最大100条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`

	// 新状态(0-待处理 1-高危 2-正常)
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetBashEventsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetBashEventsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetBashEventsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetBashEventsStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBashEventsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetBashEventsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchBashRulesRequest struct {
	*tchttp.BaseRequest

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 是否禁用
	Disabled *uint64 `json:"Disabled,omitempty" name:"Disabled"`
}

func (r *SwitchBashRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchBashRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Disabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchBashRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchBashRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchBashRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchBashRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetScanRequest struct {
	*tchttp.BaseRequest

	// 是否同步
	Sync *bool `json:"Sync,omitempty" name:"Sync"`
}

func (r *SyncAssetScanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncAssetScanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sync")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncAssetScanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SyncAssetScanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncAssetScanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncAssetScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 标签名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务器数
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type TagMachine struct {

	// ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 主机ID
	Quuid *string `json:"Quuid,omitempty" name:"Quuid"`

	// 主机名称
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// 主机内网IP
	MachineIp *string `json:"MachineIp,omitempty" name:"MachineIp"`

	// 主机外网IP
	MachineWanIp *string `json:"MachineWanIp,omitempty" name:"MachineWanIp"`

	// 主机区域
	MachineRegion *string `json:"MachineRegion,omitempty" name:"MachineRegion"`

	// 主机区域类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type TrustMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马ID数组（单次不超过的最大条数：100）
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *TrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrustMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UntrustMalwaresRequest struct {
	*tchttp.BaseRequest

	// 木马ID数组，单次最大处理不能超过200条。
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *UntrustMalwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UntrustMalwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UntrustMalwaresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UntrustMalwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntrustMalwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBaselineStrategyRequest struct {
	*tchttp.BaseRequest

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 检测周期
	ScanCycle *uint64 `json:"ScanCycle,omitempty" name:"ScanCycle"`

	// 定期检测时间，该时间下发扫描
	ScanAt *string `json:"ScanAt,omitempty" name:"ScanAt"`

	// 该策略下选择的基线id数组
	CategoryIds []*string `json:"CategoryIds,omitempty" name:"CategoryIds"`

	// 扫描范围是否全部服务器, 1:是  0:否, 为1则为全部专业版主机
	IsGlobal *uint64 `json:"IsGlobal,omitempty" name:"IsGlobal"`

	// 云主机类型：cvm：虚拟主机，bms：裸金属，ecm：边缘计算主机
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 主机地域
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 主机id数组
	Quuids []*string `json:"Quuids,omitempty" name:"Quuids"`
}

func (r *UpdateBaselineStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBaselineStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	delete(f, "StrategyName")
	delete(f, "ScanCycle")
	delete(f, "ScanAt")
	delete(f, "CategoryIds")
	delete(f, "IsGlobal")
	delete(f, "MachineType")
	delete(f, "RegionCode")
	delete(f, "Quuids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateBaselineStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateBaselineStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateBaselineStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBaselineStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsualPlace struct {

	// ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云镜客户端唯一标识UUID。
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// 国家 ID。
	CountryId *uint64 `json:"CountryId,omitempty" name:"CountryId"`

	// 省份 ID。
	ProvinceId *uint64 `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// 城市 ID。
	CityId *uint64 `json:"CityId,omitempty" name:"CityId"`
}

type VulDetailInfo struct {

	// 漏洞ID
	VulId *uint64 `json:"VulId,omitempty" name:"VulId"`

	// 漏洞级别
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 漏洞名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// cve编号
	CveId *string `json:"CveId,omitempty" name:"CveId"`

	// 漏洞分类
	VulCategory *uint64 `json:"VulCategory,omitempty" name:"VulCategory"`

	// 漏洞描述
	Descript *string `json:"Descript,omitempty" name:"Descript"`

	// 修复建议
	Fix *string `json:"Fix,omitempty" name:"Fix"`

	// 参考链接
	Reference *string `json:"Reference,omitempty" name:"Reference"`

	// CVSS评分
	CvssScore *float64 `json:"CvssScore,omitempty" name:"CvssScore"`

	// CVSS详情
	Cvss *string `json:"Cvss,omitempty" name:"Cvss"`

	// 发布时间
	PublishTime *string `json:"PublishTime,omitempty" name:"PublishTime"`
}

type WarningObject struct {

	// 事件告警类型；1：离线，2：木马，3：异常登录，4：爆破，5：漏洞（已拆分为9-12四种类型）6：高位命令，7：反弹sell，8：本地提权，9：系统组件漏洞，10：web应用漏洞，11：应急漏洞，12：安全基线
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 1: 关闭告警 0: 开启告警
	DisablePhoneWarning *uint64 `json:"DisablePhoneWarning,omitempty" name:"DisablePhoneWarning"`

	// 开始时间，格式: HH:mm
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间，格式: HH:mm
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 漏洞等级控制位二进制，每一位对应页面漏洞等级的开启关闭：低中高（0:关闭；1：开启），例如：101 → 同时勾选低+高；01→(登录审计)疑似不告警，高危告警
	ControlBits *string `json:"ControlBits,omitempty" name:"ControlBits"`
}
