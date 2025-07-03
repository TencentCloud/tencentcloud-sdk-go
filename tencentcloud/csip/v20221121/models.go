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

package v20221121

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddNewBindRoleUserRequestParams struct {

}

type AddNewBindRoleUserRequest struct {
	*tchttp.BaseRequest
	
}

func (r *AddNewBindRoleUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNewBindRoleUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNewBindRoleUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNewBindRoleUserResponseParams struct {
	// 0成功，其他失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddNewBindRoleUserResponse struct {
	*tchttp.BaseResponse
	Response *AddNewBindRoleUserResponseParams `json:"Response"`
}

func (r *AddNewBindRoleUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNewBindRoleUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertExtraInfo struct {
	// 相关攻击事件
	RelateEvent *RelatedEvent `json:"RelateEvent,omitnil,omitempty" name:"RelateEvent"`

	// 泄漏内容
	LeakContent *string `json:"LeakContent,omitnil,omitempty" name:"LeakContent"`

	// 泄漏API
	LeakAPI *string `json:"LeakAPI,omitnil,omitempty" name:"LeakAPI"`

	// secretID
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// 命中规则
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 规则描述
	RuleDesc *string `json:"RuleDesc,omitnil,omitempty" name:"RuleDesc"`

	// 协议端口
	ProtocolPort *string `json:"ProtocolPort,omitnil,omitempty" name:"ProtocolPort"`

	// 攻击内容
	AttackContent *string `json:"AttackContent,omitnil,omitempty" name:"AttackContent"`

	// 攻击IP画像
	AttackIPProfile *string `json:"AttackIPProfile,omitnil,omitempty" name:"AttackIPProfile"`

	// 攻击IP标签
	AttackIPTags *string `json:"AttackIPTags,omitnil,omitempty" name:"AttackIPTags"`

	// 请求方式
	RequestMethod *string `json:"RequestMethod,omitnil,omitempty" name:"RequestMethod"`

	// HTTP日志
	HttpLog *string `json:"HttpLog,omitnil,omitempty" name:"HttpLog"`

	// 被攻击域名
	AttackDomain *string `json:"AttackDomain,omitnil,omitempty" name:"AttackDomain"`

	// 文件路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// user_agent
	UserAgent *string `json:"UserAgent,omitnil,omitempty" name:"UserAgent"`

	// 请求头
	RequestHeaders *string `json:"RequestHeaders,omitnil,omitempty" name:"RequestHeaders"`

	// 登录用户名
	LoginUserName *string `json:"LoginUserName,omitnil,omitempty" name:"LoginUserName"`

	// 漏洞名称
	VulnerabilityName *string `json:"VulnerabilityName,omitnil,omitempty" name:"VulnerabilityName"`

	// 公共漏洞和暴露
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// 服务进程
	ServiceProcess *string `json:"ServiceProcess,omitnil,omitempty" name:"ServiceProcess"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 文件MD5
	FileMD5 *string `json:"FileMD5,omitnil,omitempty" name:"FileMD5"`

	// 文件最近访问时间
	FileLastAccessTime *string `json:"FileLastAccessTime,omitnil,omitempty" name:"FileLastAccessTime"`

	// 文件修改时间
	FileModifyTime *string `json:"FileModifyTime,omitnil,omitempty" name:"FileModifyTime"`

	// 最近访问时间
	RecentAccessTime *string `json:"RecentAccessTime,omitnil,omitempty" name:"RecentAccessTime"`

	// 最近修改时间
	RecentModifyTime *string `json:"RecentModifyTime,omitnil,omitempty" name:"RecentModifyTime"`

	// 病毒名
	VirusName *string `json:"VirusName,omitnil,omitempty" name:"VirusName"`

	// 病毒文件标签
	VirusFileTags *string `json:"VirusFileTags,omitnil,omitempty" name:"VirusFileTags"`

	// 行为特征
	BehavioralCharacteristics *string `json:"BehavioralCharacteristics,omitnil,omitempty" name:"BehavioralCharacteristics"`

	// 进程名（PID）
	ProcessNamePID *string `json:"ProcessNamePID,omitnil,omitempty" name:"ProcessNamePID"`

	// 进程路径
	ProcessPath *string `json:"ProcessPath,omitnil,omitempty" name:"ProcessPath"`

	// 进程命令行
	ProcessCommandLine *string `json:"ProcessCommandLine,omitnil,omitempty" name:"ProcessCommandLine"`

	// 进程权限
	ProcessPermissions *string `json:"ProcessPermissions,omitnil,omitempty" name:"ProcessPermissions"`

	// 执行命令
	ExecutedCommand *string `json:"ExecutedCommand,omitnil,omitempty" name:"ExecutedCommand"`

	// 受影响文件名
	AffectedFileName *string `json:"AffectedFileName,omitnil,omitempty" name:"AffectedFileName"`

	// 诱饵路径
	DecoyPath *string `json:"DecoyPath,omitnil,omitempty" name:"DecoyPath"`

	// 恶意进程文件大小
	MaliciousProcessFileSize *string `json:"MaliciousProcessFileSize,omitnil,omitempty" name:"MaliciousProcessFileSize"`

	// 恶意进程文件MD5
	MaliciousProcessFileMD5 *string `json:"MaliciousProcessFileMD5,omitnil,omitempty" name:"MaliciousProcessFileMD5"`

	// 恶意进程名（PID）
	MaliciousProcessNamePID *string `json:"MaliciousProcessNamePID,omitnil,omitempty" name:"MaliciousProcessNamePID"`

	// 恶意进程路径
	MaliciousProcessPath *string `json:"MaliciousProcessPath,omitnil,omitempty" name:"MaliciousProcessPath"`

	// 恶意进程启动时间
	MaliciousProcessStartTime *string `json:"MaliciousProcessStartTime,omitnil,omitempty" name:"MaliciousProcessStartTime"`

	// 命令内容
	CommandContent *string `json:"CommandContent,omitnil,omitempty" name:"CommandContent"`

	// 启动用户
	StartupUser *string `json:"StartupUser,omitnil,omitempty" name:"StartupUser"`

	// 用户所属组
	UserGroup *string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// 新增权限
	NewPermissions *string `json:"NewPermissions,omitnil,omitempty" name:"NewPermissions"`

	// 父进程
	ParentProcess *string `json:"ParentProcess,omitnil,omitempty" name:"ParentProcess"`

	// 类名
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// 所属类加载器
	ClassLoader *string `json:"ClassLoader,omitnil,omitempty" name:"ClassLoader"`

	// 类文件大小
	ClassFileSize *string `json:"ClassFileSize,omitnil,omitempty" name:"ClassFileSize"`

	// 类文件MD5
	ClassFileMD5 *string `json:"ClassFileMD5,omitnil,omitempty" name:"ClassFileMD5"`

	// 父类名
	ParentClassName *string `json:"ParentClassName,omitnil,omitempty" name:"ParentClassName"`

	// 继承接口
	InheritedInterface *string `json:"InheritedInterface,omitnil,omitempty" name:"InheritedInterface"`

	// 注释
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 载荷内容
	PayloadContent *string `json:"PayloadContent,omitnil,omitempty" name:"PayloadContent"`

	// 回连地址画像
	CallbackAddressPortrait *string `json:"CallbackAddressPortrait,omitnil,omitempty" name:"CallbackAddressPortrait"`

	// 回连地址标签
	CallbackAddressTag *string `json:"CallbackAddressTag,omitnil,omitempty" name:"CallbackAddressTag"`

	// 进程MD5
	ProcessMD5 *string `json:"ProcessMD5,omitnil,omitempty" name:"ProcessMD5"`

	// 文件权限
	FilePermission *string `json:"FilePermission,omitnil,omitempty" name:"FilePermission"`

	// 来源于日志分析的信息字段
	FromLogAnalysisData []*KeyValue `json:"FromLogAnalysisData,omitnil,omitempty" name:"FromLogAnalysisData"`

	// 命中探针
	HitProbe *string `json:"HitProbe,omitnil,omitempty" name:"HitProbe"`

	// 命中蜜罐
	HitHoneyPot *string `json:"HitHoneyPot,omitnil,omitempty" name:"HitHoneyPot"`

	// 命令列表
	CommandList *string `json:"CommandList,omitnil,omitempty" name:"CommandList"`

	// 攻击事件描述
	AttackEventDesc *string `json:"AttackEventDesc,omitnil,omitempty" name:"AttackEventDesc"`

	// 进程信息
	ProcessInfo *string `json:"ProcessInfo,omitnil,omitempty" name:"ProcessInfo"`

	// 使用用户名&密码
	UserNameAndPwd *string `json:"UserNameAndPwd,omitnil,omitempty" name:"UserNameAndPwd"`

	// 主机防护策略ID
	StrategyID *string `json:"StrategyID,omitnil,omitempty" name:"StrategyID"`

	// 主机防护策略名称
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 主机防护命中策略，是策略ID和策略名称的组合
	HitStrategy *string `json:"HitStrategy,omitnil,omitempty" name:"HitStrategy"`

	// 进程名
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// PID
	PID *string `json:"PID,omitnil,omitempty" name:"PID"`

	// 容器Pod名
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 容器PodID
	PodID *string `json:"PodID,omitnil,omitempty" name:"PodID"`

	// Http响应
	Response *string `json:"Response,omitnil,omitempty" name:"Response"`

	// 系统调用
	SystemCall *string `json:"SystemCall,omitnil,omitempty" name:"SystemCall"`

	// 操作类型verb
	Verb *string `json:"Verb,omitnil,omitempty" name:"Verb"`

	// 日志ID
	LogID *string `json:"LogID,omitnil,omitempty" name:"LogID"`

	// 变更内容
	Different *string `json:"Different,omitnil,omitempty" name:"Different"`

	// 事件类型
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 事件描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 目标地址(容器反弹shell)
	TargetAddress *string `json:"TargetAddress,omitnil,omitempty" name:"TargetAddress"`

	// 恶意请求域名(容器恶意外联)
	MaliciousRequestDomain *string `json:"MaliciousRequestDomain,omitnil,omitempty" name:"MaliciousRequestDomain"`

	// 规则类型(容器K8sAPI异常请求)
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 请求资源(容器K8sAPI异常请求)
	RequestURI *string `json:"RequestURI,omitnil,omitempty" name:"RequestURI"`

	// 发起请求用户(容器K8sAPI异常请求)
	RequestUser *string `json:"RequestUser,omitnil,omitempty" name:"RequestUser"`

	// 请求对象(容器K8sAPI异常请求)
	RequestObject *string `json:"RequestObject,omitnil,omitempty" name:"RequestObject"`

	// 响应对象(容器K8sAPI异常请求)
	ResponseObject *string `json:"ResponseObject,omitnil,omitempty" name:"ResponseObject"`

	// 文件类型(容器文件篡改)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 标签特征(容器恶意外联)
	TIType *string `json:"TIType,omitnil,omitempty" name:"TIType"`

	// 来源IP(容器K8sAPI异常请求)
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`
}

type AlertInfo struct {
	// 告警ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 告警名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 告警来源
	// CFW:云防火墙
	// WAF:Web应用防火墙
	// CWP:主机安全
	// CSIP:云安全中心
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 告警等级
	// 1:提示
	// 2:低危
	// 3:中危
	// 4:高危
	// 5:严重
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 攻击者
	Attacker *RoleInfo `json:"Attacker,omitnil,omitempty" name:"Attacker"`

	// 受害者
	Victim *RoleInfo `json:"Victim,omitnil,omitempty" name:"Victim"`

	// 证据数据(例如攻击内容等，base64编码)
	EvidenceData *string `json:"EvidenceData,omitnil,omitempty" name:"EvidenceData"`

	// 证据位置(例如协议端口)
	EvidenceLocation *string `json:"EvidenceLocation,omitnil,omitempty" name:"EvidenceLocation"`

	// 证据路径
	EvidencePath *string `json:"EvidencePath,omitnil,omitempty" name:"EvidencePath"`

	// 首次告警时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近告警时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 告警次数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 紧急缓解建议
	UrgentSuggestion *string `json:"UrgentSuggestion,omitnil,omitempty" name:"UrgentSuggestion"`

	// 根治建议
	RemediationSuggestion *string `json:"RemediationSuggestion,omitnil,omitempty" name:"RemediationSuggestion"`

	// 处理状态
	// 0：未处置，1：已忽略，2：已处置
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 告警处理类型
	ProcessType *string `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// 告警大类
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警小类
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 下拉字段
	ExtraInfo *AlertExtraInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 聚合字段
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 告警日期
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// appid
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 账户名称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 账户ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 行为
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 风险排查
	RiskInvestigation *string `json:"RiskInvestigation,omitnil,omitempty" name:"RiskInvestigation"`

	// 风险处置
	RiskTreatment *string `json:"RiskTreatment,omitnil,omitempty" name:"RiskTreatment"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 语句检索
	LogSearch *string `json:"LogSearch,omitnil,omitempty" name:"LogSearch"`
}

type AssetBaseInfoResponse struct {
	// vpc-id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc-name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 操作系统
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 公网ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号数量
	AccountNum *uint64 `json:"AccountNum,omitnil,omitempty" name:"AccountNum"`

	// 端口数量
	PortNum *uint64 `json:"PortNum,omitnil,omitempty" name:"PortNum"`

	// 进程数量
	ProcessNum *uint64 `json:"ProcessNum,omitnil,omitempty" name:"ProcessNum"`

	// 软件应用数量
	SoftApplicationNum *uint64 `json:"SoftApplicationNum,omitnil,omitempty" name:"SoftApplicationNum"`

	// 数据库数量
	DatabaseNum *uint64 `json:"DatabaseNum,omitnil,omitempty" name:"DatabaseNum"`

	// Web应用数量
	WebApplicationNum *uint64 `json:"WebApplicationNum,omitnil,omitempty" name:"WebApplicationNum"`

	// 服务数量
	ServiceNum *uint64 `json:"ServiceNum,omitnil,omitempty" name:"ServiceNum"`

	// web框架数量
	WebFrameworkNum *uint64 `json:"WebFrameworkNum,omitnil,omitempty" name:"WebFrameworkNum"`

	// Web站点数量
	WebSiteNum *uint64 `json:"WebSiteNum,omitnil,omitempty" name:"WebSiteNum"`

	// Jar包数量
	JarPackageNum *uint64 `json:"JarPackageNum,omitnil,omitempty" name:"JarPackageNum"`

	// 启动服务数量
	StartServiceNum *uint64 `json:"StartServiceNum,omitnil,omitempty" name:"StartServiceNum"`

	// 计划任务数量
	ScheduledTaskNum *uint64 `json:"ScheduledTaskNum,omitnil,omitempty" name:"ScheduledTaskNum"`

	// 环境变量数量
	EnvironmentVariableNum *uint64 `json:"EnvironmentVariableNum,omitnil,omitempty" name:"EnvironmentVariableNum"`

	// 内核模块数量
	KernelModuleNum *uint64 `json:"KernelModuleNum,omitnil,omitempty" name:"KernelModuleNum"`

	// 系统安装包数量
	SystemInstallationPackageNum *uint64 `json:"SystemInstallationPackageNum,omitnil,omitempty" name:"SystemInstallationPackageNum"`

	// 剩余防护时长
	SurplusProtectDay *uint64 `json:"SurplusProtectDay,omitnil,omitempty" name:"SurplusProtectDay"`

	// 客户端是否安装  1 已安装 0 未安装
	CWPStatus *uint64 `json:"CWPStatus,omitnil,omitempty" name:"CWPStatus"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 防护等级
	ProtectLevel *string `json:"ProtectLevel,omitnil,omitempty" name:"ProtectLevel"`

	// 防护时长
	ProtectedDay *uint64 `json:"ProtectedDay,omitnil,omitempty" name:"ProtectedDay"`
}

type AssetCluster struct {
	// 租户id
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 租户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 租户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 集群id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 集群名称
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 集群类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 集群创建时间
	InstanceCreateTime *string `json:"InstanceCreateTime,omitnil,omitempty" name:"InstanceCreateTime"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群防护状态，左边枚举,右边为显示
	// 集群防护状态 
	// 0:未接入
	// 1:未防护 
	// 2:部分防护 
	// 3:防护中 
	// 4:接入异常 
	// 5:接入中 
	// 6:卸载中 
	// 7:卸载异常
	ProtectStatus *int64 `json:"ProtectStatus,omitnil,omitempty" name:"ProtectStatus"`

	// 接入信息，不为空表示有接入异常信息
	ProtectInfo *string `json:"ProtectInfo,omitnil,omitempty" name:"ProtectInfo"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// kubernetes版本
	KubernetesVersion *string `json:"KubernetesVersion,omitnil,omitempty" name:"KubernetesVersion"`

	// 运行时组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 运行时组件版本
	ComponentVersion *string `json:"ComponentVersion,omitnil,omitempty" name:"ComponentVersion"`

	// 组件状态
	ComponentStatus *string `json:"ComponentStatus,omitnil,omitempty" name:"ComponentStatus"`

	// 体检时间
	CheckTime *string `json:"CheckTime,omitnil,omitempty" name:"CheckTime"`

	// 关联主机数
	MachineCount *int64 `json:"MachineCount,omitnil,omitempty" name:"MachineCount"`

	// 关联pod数
	PodCount *int64 `json:"PodCount,omitnil,omitempty" name:"PodCount"`

	// 关联service数
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// 漏洞风险
	VulRisk *int64 `json:"VulRisk,omitnil,omitempty" name:"VulRisk"`

	// 配置风险
	CFGRisk *int64 `json:"CFGRisk,omitnil,omitempty" name:"CFGRisk"`

	// 体检数
	CheckCount *int64 `json:"CheckCount,omitnil,omitempty" name:"CheckCount"`

	// 是否核心：1:核心，2:非核心
	IsCore *int64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 是否新资产 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// 云资产类型：0：腾讯云，1：aws，2：azure
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

type AssetClusterPod struct {
	// 租户id
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 租户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 租户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// pod id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// pod名称
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// pod创建时间
	InstanceCreateTime *string `json:"InstanceCreateTime,omitnil,omitempty" name:"InstanceCreateTime"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 主机id
	MachineId *string `json:"MachineId,omitnil,omitempty" name:"MachineId"`

	// 主机名
	MachineName *string `json:"MachineName,omitnil,omitempty" name:"MachineName"`

	// pod ip
	PodIp *string `json:"PodIp,omitnil,omitempty" name:"PodIp"`

	// 关联service数
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// 关联容器数
	ContainerCount *int64 `json:"ContainerCount,omitnil,omitempty" name:"ContainerCount"`

	// 公网ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 是否核心：1:核心，2:非核心
	IsCore *int64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 是否新资产 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type AssetInfoDetail struct {
	// 用户appid
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// CVE编号
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`

	// 是扫描，0默认未扫描，1正在扫描，2扫描完成，3扫描出错
	IsScan *int64 `json:"IsScan,omitnil,omitempty" name:"IsScan"`

	// 影响资产数目
	InfluenceAsset *int64 `json:"InfluenceAsset,omitnil,omitempty" name:"InfluenceAsset"`

	// 未修复资产数目
	NotRepairAsset *int64 `json:"NotRepairAsset,omitnil,omitempty" name:"NotRepairAsset"`

	// 未防护资产数目
	NotProtectAsset *int64 `json:"NotProtectAsset,omitnil,omitempty" name:"NotProtectAsset"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务百分比
	TaskPercent *int64 `json:"TaskPercent,omitnil,omitempty" name:"TaskPercent"`

	// 任务时间
	TaskTime *int64 `json:"TaskTime,omitnil,omitempty" name:"TaskTime"`

	// 扫描时间
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`
}

type AssetInstanceTypeMap struct {
	// 资产类型
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 资产类型
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 资产类型和实例类型映射关系
	InstanceTypeList []*FilterDataObject `json:"InstanceTypeList,omitnil,omitempty" name:"InstanceTypeList"`
}

type AssetProcessItem struct {
	// 云账号ID
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 租户ID
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 云账号名称
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 进程ID
	ProcessID *string `json:"ProcessID,omitnil,omitempty" name:"ProcessID"`

	// 进程名称
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// 命令行
	CmdLine *string `json:"CmdLine,omitnil,omitempty" name:"CmdLine"`

	// 监听端口列表
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

type AssetRiskItem struct {
	// 租户ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 云厂商名称
	ProviderName *string `json:"ProviderName,omitnil,omitempty" name:"ProviderName"`

	// 云账号名称
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// 云账号ID
	CloudAccountId *string `json:"CloudAccountId,omitnil,omitempty" name:"CloudAccountId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 首次发现时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 风险状态
	RiskStatus *int64 `json:"RiskStatus,omitnil,omitempty" name:"RiskStatus"`

	// 风险名称
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// 检查类型
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// 风险等级
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 风险规则ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`
}

type AssetTag struct {
	// 标签的key值,可以是字母、数字、下划线
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的vale值,可以是字母、数字、下划线
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type AssetViewCFGRisk struct {
	// 唯一id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 配置名
	CFGName *string `json:"CFGName,omitnil,omitempty" name:"CFGName"`

	// 检查类型
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 来源
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 相关规范
	CFGSTD *string `json:"CFGSTD,omitnil,omitempty" name:"CFGSTD"`

	// 配置详情
	CFGDescribe *string `json:"CFGDescribe,omitnil,omitempty" name:"CFGDescribe"`

	// 修复建议
	CFGFix *string `json:"CFGFix,omitnil,omitempty" name:"CFGFix"`

	// 帮助文档链接
	CFGHelpURL *string `json:"CFGHelpURL,omitnil,omitempty" name:"CFGHelpURL"`

	// 前端使用索引
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 当资产类型为LBL的时候，展示该字段，方便定位具体的LB
	ClbId *string `json:"ClbId,omitnil,omitempty" name:"ClbId"`
}

type AssetViewPortRisk struct {
	// 端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 服务
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 处置建议,0保持现状、1限制访问、2封禁端口
	Suggestion *uint64 `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 状态，0未处理、1已处置、2已忽略、3云防已防护
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 风险ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 前端索引
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 识别来源，详细看枚举返回。
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 服务判定,high_risk_service 高危服务 web_service web服务 other_service 其他服务
	ServiceJudge *string `json:"ServiceJudge,omitnil,omitempty" name:"ServiceJudge"`

	// 状态，0未处理、1已处置、2已忽略、3云防已防护、4无需处理
	XspmStatus *uint64 `json:"XspmStatus,omitnil,omitempty" name:"XspmStatus"`
}

type AssetViewVULRisk struct {
	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 服务
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 状态，0未处理、1已处置、2已忽略
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 风险ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 前端索引
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 漏洞类型
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 漏洞描述
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 漏洞影响组件
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 技术参考
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// 漏洞影响版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 风险点
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// 漏洞名称
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// 修复方案
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// pocid
	POCId *string `json:"POCId,omitnil,omitempty" name:"POCId"`

	// 扫描来源
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 主机版本
	CWPVersion *int64 `json:"CWPVersion,omitnil,omitempty" name:"CWPVersion"`

	// 是否支持修复
	IsSupportRepair *bool `json:"IsSupportRepair,omitnil,omitempty" name:"IsSupportRepair"`

	// 是否支持扫描
	IsSupportDetect *bool `json:"IsSupportDetect,omitnil,omitempty" name:"IsSupportDetect"`

	// 实例uuid
	InstanceUUID *string `json:"InstanceUUID,omitnil,omitempty" name:"InstanceUUID"`

	// 攻击载荷
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`
}

type AssetViewVULRiskData struct {
	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 状态，0未处理、1标记已处置、2已忽略，3已处置 ，4 处置中 ，5 检测中 ，6部分已处置
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 风险ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 漏洞类型
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 漏洞影响组件
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 漏洞影响版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 风险点
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// 漏洞名称
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// pocid
	POCId *string `json:"POCId,omitnil,omitempty" name:"POCId"`

	// 扫描来源
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 主机版本
	CWPVersion *int64 `json:"CWPVersion,omitnil,omitempty" name:"CWPVersion"`

	// 实例uuid
	InstanceUUID *string `json:"InstanceUUID,omitnil,omitempty" name:"InstanceUUID"`

	// 攻击载荷
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`

	// CVSS评分
	CVSS *float64 `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// 前端索引id
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// pcmgrId
	PCMGRId *string `json:"PCMGRId,omitnil,omitempty" name:"PCMGRId"`

	// 报告id
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 漏洞标签
	VulTag []*string `json:"VulTag,omitnil,omitempty" name:"VulTag"`

	// 漏洞披露时间
	DisclosureTime *string `json:"DisclosureTime,omitnil,omitempty" name:"DisclosureTime"`

	// 攻击热度
	AttackHeat *uint64 `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// 是否必修漏洞1是，0不是
	IsSuggest *int64 `json:"IsSuggest,omitnil,omitempty" name:"IsSuggest"`

	// 处置任务ID
	HandleTaskId *string `json:"HandleTaskId,omitnil,omitempty" name:"HandleTaskId"`

	// 引擎来源
	EngineSource *string `json:"EngineSource,omitnil,omitempty" name:"EngineSource"`

	// 新的漏洞风险id(同全网漏洞表的RiskId)
	VulRiskId *string `json:"VulRiskId,omitnil,omitempty" name:"VulRiskId"`

	// 新版漏洞id
	TvdID *string `json:"TvdID,omitnil,omitempty" name:"TvdID"`

	// 是否可以一键体检，1-可以，0-不可以
	IsOneClick *uint64 `json:"IsOneClick,omitnil,omitempty" name:"IsOneClick"`

	// 是否POC扫描，0-非POC，1-POC
	IsPOC *uint64 `json:"IsPOC,omitnil,omitempty" name:"IsPOC"`
}

type AssetViewWeakPassRisk struct {
	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 服务
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 状态，0未处理、1已处置、2已忽略
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// ID，处理风险使用
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 前端索引
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 弱口令类型
	PasswordType *string `json:"PasswordType,omitnil,omitempty" name:"PasswordType"`

	// 来源
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 漏洞类型
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// 漏洞url
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// 修复建议
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// 证明
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type AttributeOptionSet struct {
	// cvm实例类型
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// cvm实例名称
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type BugInfoDetail struct {
	// 漏洞编号
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 漏洞对应pocId
	PatchId *string `json:"PatchId,omitnil,omitempty" name:"PatchId"`

	// 漏洞名称
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// 漏洞严重性：high,middle，low，info
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// cvss评分
	CVSSScore *string `json:"CVSSScore,omitnil,omitempty" name:"CVSSScore"`

	// cve编号
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`

	// 漏洞标签
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 漏洞种类，1:web应用，2:系统组件漏洞，3:配置风险
	VULCategory *uint64 `json:"VULCategory,omitnil,omitempty" name:"VULCategory"`

	// 漏洞影响系统
	ImpactOs *string `json:"ImpactOs,omitnil,omitempty" name:"ImpactOs"`

	// 漏洞影响组件
	ImpactCOMPENT *string `json:"ImpactCOMPENT,omitnil,omitempty" name:"ImpactCOMPENT"`

	// 漏洞影响版本
	ImpactVersion *string `json:"ImpactVersion,omitnil,omitempty" name:"ImpactVersion"`

	// 链接
	Reference *string `json:"Reference,omitnil,omitempty" name:"Reference"`

	// 漏洞描述
	VULDescribe *string `json:"VULDescribe,omitnil,omitempty" name:"VULDescribe"`

	// 修复建议
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// 产品支持状态，实时返回
	ProSupport *uint64 `json:"ProSupport,omitnil,omitempty" name:"ProSupport"`

	// 是否公开，0为未发布，1为发布
	IsPublish *uint64 `json:"IsPublish,omitnil,omitempty" name:"IsPublish"`

	// 释放时间
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 漏洞子类别
	SubCategory *string `json:"SubCategory,omitnil,omitempty" name:"SubCategory"`
}

type CFGViewCFGRisk struct {
	// 影响资产
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 状态，0未处理、1已处置、2已忽略
	AffectAssetCount *uint64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// 资产唯一id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 资产子类型
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 前端索引
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 配置名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFGName *string `json:"CFGName,omitnil,omitempty" name:"CFGName"`

	// 检查类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFGSTD *string `json:"CFGSTD,omitnil,omitempty" name:"CFGSTD"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFGDescribe *string `json:"CFGDescribe,omitnil,omitempty" name:"CFGDescribe"`

	// 修复建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFGFix *string `json:"CFGFix,omitnil,omitempty" name:"CFGFix"`

	// 帮助文档
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFGHelpURL *string `json:"CFGHelpURL,omitnil,omitempty" name:"CFGHelpURL"`
}

type CVMAssetVO struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 防护状态
	CWPStatus *uint64 `json:"CWPStatus,omitnil,omitempty" name:"CWPStatus"`

	// 资产创建时间
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 私网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc 名
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// appid信息
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 可用区
	AvailableArea *string `json:"AvailableArea,omitnil,omitempty" name:"AvailableArea"`

	// 是否核心
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网名
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 主机安全Agent UUID
	InstanceUuid *string `json:"InstanceUuid,omitnil,omitempty" name:"InstanceUuid"`

	// 云主机 CVM UUID
	InstanceQUuid *string `json:"InstanceQUuid,omitnil,omitempty" name:"InstanceQUuid"`

	// os名
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 分区
	PartitionCount *uint64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// cpu信息
	CPUInfo *string `json:"CPUInfo,omitnil,omitempty" name:"CPUInfo"`

	// cpu大小
	CPUSize *uint64 `json:"CPUSize,omitnil,omitempty" name:"CPUSize"`

	// cpu负载
	CPULoad *string `json:"CPULoad,omitnil,omitempty" name:"CPULoad"`

	// 内存大小
	MemorySize *string `json:"MemorySize,omitnil,omitempty" name:"MemorySize"`

	// 内存负载
	MemoryLoad *string `json:"MemoryLoad,omitnil,omitempty" name:"MemoryLoad"`

	// 硬盘大小
	DiskSize *string `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 硬盘负载
	DiskLoad *string `json:"DiskLoad,omitnil,omitempty" name:"DiskLoad"`

	// 账号数
	AccountCount *string `json:"AccountCount,omitnil,omitempty" name:"AccountCount"`

	// 进程数
	ProcessCount *string `json:"ProcessCount,omitnil,omitempty" name:"ProcessCount"`

	// 软件应用
	AppCount *string `json:"AppCount,omitnil,omitempty" name:"AppCount"`

	// 监听端口
	PortCount *uint64 `json:"PortCount,omitnil,omitempty" name:"PortCount"`

	// 网络攻击
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// 网络访问
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// 网络拦截
	Intercept *uint64 `json:"Intercept,omitnil,omitempty" name:"Intercept"`

	// 入向峰值带宽
	InBandwidth *string `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`

	// 出向峰值带宽
	OutBandwidth *string `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// 入向累计流量
	InFlow *string `json:"InFlow,omitnil,omitempty" name:"InFlow"`

	// 出向累计流量
	OutFlow *string `json:"OutFlow,omitnil,omitempty" name:"OutFlow"`

	// 最近扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 恶意主动外联
	NetWorkOut *uint64 `json:"NetWorkOut,omitnil,omitempty" name:"NetWorkOut"`

	// 端口风险
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// 漏洞风险
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil,omitempty" name:"VulnerabilityRisk"`

	// 配置风险
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// 扫描任务数
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// memberId
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// os全称
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 风险服务暴露
	RiskExposure *int64 `json:"RiskExposure,omitnil,omitempty" name:"RiskExposure"`

	// 模拟攻击工具状态。0代表未安装，1代表已安装，2代表已离线
	BASAgentStatus *int64 `json:"BASAgentStatus,omitnil,omitempty" name:"BASAgentStatus"`

	// 1新资产；0 非新资产
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// 0 未安装  1安装 2:安装中
	CVMAgentStatus *int64 `json:"CVMAgentStatus,omitnil,omitempty" name:"CVMAgentStatus"`

	// 1:开启 0:未开启
	CVMStatus *int64 `json:"CVMStatus,omitnil,omitempty" name:"CVMStatus"`

	// 1:客户端已安装 0：未安装 2: Agentless
	DefenseModel *int64 `json:"DefenseModel,omitnil,omitempty" name:"DefenseModel"`

	// 1:已安装 0:未安装
	TatStatus *int64 `json:"TatStatus,omitnil,omitempty" name:"TatStatus"`

	// cpu趋势图
	CpuTrend []*Element `json:"CpuTrend,omitnil,omitempty" name:"CpuTrend"`

	// 内存趋势图
	MemoryTrend []*Element `json:"MemoryTrend,omitnil,omitempty" name:"MemoryTrend"`

	// 1:agent在线 0:agent离线 2:主机离线
	AgentStatus *int64 `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`

	// 本月防护关闭次数
	CloseDefenseCount *int64 `json:"CloseDefenseCount,omitnil,omitempty" name:"CloseDefenseCount"`

	// 运行状态
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 安全组数据
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 物理内存占用KB
	AgentMemRss *int64 `json:"AgentMemRss,omitnil,omitempty" name:"AgentMemRss"`

	// CPU使用率百分比
	AgentCpuPer *float64 `json:"AgentCpuPer,omitnil,omitempty" name:"AgentCpuPer"`

	// cvm真正所属的appid
	RealAppid *int64 `json:"RealAppid,omitnil,omitempty" name:"RealAppid"`

	// 云资产类型：0：腾讯云，1：aws，2：azure
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 主机防护状态枚举
	// 0：未安装
	// 1：基础版防护中
	// 2：普惠版防护中
	// 3：专业版防护中
	// 4：旗舰版防护中
	// 5：已离线
	// 6：已关机
	ProtectStatus *int64 `json:"ProtectStatus,omitnil,omitempty" name:"ProtectStatus"`

	// 最后离线时间
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`
}

type CheckViewRiskItem struct {
	// 检查项规则ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// 风险名称
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// 检查类型
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// 风险等级
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 存在1个风险项
	RiskDesc *string `json:"RiskDesc,omitnil,omitempty" name:"RiskDesc"`

	// 首次发现时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 风险更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 风险状态
	RiskStatus *uint64 `json:"RiskStatus,omitnil,omitempty" name:"RiskStatus"`

	// 受影响资产数量
	AssetCount *uint64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// 风险数量
	RiskCount *uint64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 事件类型
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`
}

type ClbListenerListInfo struct {
	// 监听器id
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 负载均衡Id
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡名称
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 负载均衡ip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 区域
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络id
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil,omitempty" name:"NumericalVpcId"`

	// 负载均衡类型
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 监听器域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 负载均衡域名
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil,omitempty" name:"LoadBalancerDomain"`
}

type CloudCountDesc struct {
	// 0表示腾讯云
	// 1表示AWS
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 账户数量
	CloudCount *int64 `json:"CloudCount,omitnil,omitempty" name:"CloudCount"`

	// 该云账号类型描述
	CloudDesc *string `json:"CloudDesc,omitnil,omitempty" name:"CloudDesc"`
}

// Predefined struct for user
type CreateDomainAndIpRequestParams struct {
	// 公网IP/域名
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDomainAndIpRequest struct {
	*tchttp.BaseRequest
	
	// 公网IP/域名
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateDomainAndIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainAndIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "MemberId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainAndIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainAndIpResponseParams struct {
	// 返回创建成功的数量
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainAndIpResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainAndIpResponseParams `json:"Response"`
}

func (r *CreateDomainAndIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainAndIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRiskCenterScanTaskRequestParams struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// 扫描项目；port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 扫描资产信息列表
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// 扫描计划详情
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// ip/域名/url数组
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// 请求发起源，vss表示漏洞扫描服务，云安全中心的用户请填充csip，默认csip
	ScanFrom *string `json:"ScanFrom,omitnil,omitempty" name:"ScanFrom"`

	// 高级配置
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// 资产标签
	Tags *AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 任务完成回调webhook地址
	FinishWebHook *string `json:"FinishWebHook,omitnil,omitempty" name:"FinishWebHook"`
}

type CreateRiskCenterScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// 扫描项目；port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 扫描资产信息列表
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// 扫描计划详情
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// ip/域名/url数组
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// 请求发起源，vss表示漏洞扫描服务，云安全中心的用户请填充csip，默认csip
	ScanFrom *string `json:"ScanFrom,omitnil,omitempty" name:"ScanFrom"`

	// 高级配置
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// 资产标签
	Tags *AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 任务完成回调webhook地址
	FinishWebHook *string `json:"FinishWebHook,omitnil,omitempty" name:"FinishWebHook"`
}

func (r *CreateRiskCenterScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRiskCenterScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "ScanAssetType")
	delete(f, "ScanItem")
	delete(f, "ScanPlanType")
	delete(f, "MemberId")
	delete(f, "Assets")
	delete(f, "ScanPlanContent")
	delete(f, "SelfDefiningAssets")
	delete(f, "ScanFrom")
	delete(f, "TaskAdvanceCFG")
	delete(f, "TaskMode")
	delete(f, "Tags")
	delete(f, "FinishWebHook")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRiskCenterScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRiskCenterScanTaskResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 0,任务创建成功；小于0失败；-1为存在资产未认证
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 未认证资产列表
	UnAuthAsset []*string `json:"UnAuthAsset,omitnil,omitempty" name:"UnAuthAsset"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRiskCenterScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRiskCenterScanTaskResponseParams `json:"Response"`
}

func (r *CreateRiskCenterScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRiskCenterScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CsipRiskCenterStatistics struct {
	// 端口风险总数
	PortTotal *uint64 `json:"PortTotal,omitnil,omitempty" name:"PortTotal"`

	// 端口风险高危数量
	PortHighLevel *uint64 `json:"PortHighLevel,omitnil,omitempty" name:"PortHighLevel"`

	// 	弱口令风险总数
	WeakPasswordTotal *uint64 `json:"WeakPasswordTotal,omitnil,omitempty" name:"WeakPasswordTotal"`

	// 弱口令风险高危数量
	WeakPasswordHighLevel *uint64 `json:"WeakPasswordHighLevel,omitnil,omitempty" name:"WeakPasswordHighLevel"`

	// 网站风险数量
	WebsiteTotal *uint64 `json:"WebsiteTotal,omitnil,omitempty" name:"WebsiteTotal"`

	// 网站高危风险数量
	WebsiteHighLevel *uint64 `json:"WebsiteHighLevel,omitnil,omitempty" name:"WebsiteHighLevel"`

	// 最新的扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 漏洞风险数
	VULTotal *int64 `json:"VULTotal,omitnil,omitempty" name:"VULTotal"`

	// 高危漏洞风险数
	VULHighLevel *int64 `json:"VULHighLevel,omitnil,omitempty" name:"VULHighLevel"`

	// 配置项风险数量
	CFGTotal *int64 `json:"CFGTotal,omitnil,omitempty" name:"CFGTotal"`

	// 高危配置项风险数量
	CFGHighLevel *int64 `json:"CFGHighLevel,omitnil,omitempty" name:"CFGHighLevel"`

	// 测绘服务风险数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerTotal *int64 `json:"ServerTotal,omitnil,omitempty" name:"ServerTotal"`

	// 测绘服务高危数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerHighLevel *int64 `json:"ServerHighLevel,omitnil,omitempty" name:"ServerHighLevel"`

	// 主机基线风险数量
	HostBaseLineRiskTotal *int64 `json:"HostBaseLineRiskTotal,omitnil,omitempty" name:"HostBaseLineRiskTotal"`

	// 主机基线高危风险数量
	HostBaseLineRiskHighLevel *int64 `json:"HostBaseLineRiskHighLevel,omitnil,omitempty" name:"HostBaseLineRiskHighLevel"`

	// 容器基线风险数量
	PodBaseLineRiskTotal *int64 `json:"PodBaseLineRiskTotal,omitnil,omitempty" name:"PodBaseLineRiskTotal"`

	// 容器基线高危风险数量
	PodBaseLineRiskHighLevel *int64 `json:"PodBaseLineRiskHighLevel,omitnil,omitempty" name:"PodBaseLineRiskHighLevel"`
}

type DBAssetVO struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// vpcid
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc标签
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 资产创建时间
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// 最近扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 配置风险
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// 网络攻击
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// 网络访问
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// 扫描任务
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// 用户appid
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 昵称别名
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 内网ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 公网ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否核心
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 是否新资产: 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type DataSearchBug struct {
	// 返回查询状态
	StateCode *string `json:"StateCode,omitnil,omitempty" name:"StateCode"`

	// 漏洞详情
	DataBug []*BugInfoDetail `json:"DataBug,omitnil,omitempty" name:"DataBug"`

	// 漏洞影响资产详情
	DataAsset []*AssetInfoDetail `json:"DataAsset,omitnil,omitempty" name:"DataAsset"`

	// true支持扫描。false不支持扫描
	VSSScan *bool `json:"VSSScan,omitnil,omitempty" name:"VSSScan"`

	// 0不支持，1支持
	CWPScan *string `json:"CWPScan,omitnil,omitempty" name:"CWPScan"`

	// 1支持虚拟补丁，0或空不支持
	CFWPatch *string `json:"CFWPatch,omitnil,omitempty" name:"CFWPatch"`

	// 0不支持，1支持
	WafPatch *int64 `json:"WafPatch,omitnil,omitempty" name:"WafPatch"`

	// 0不支持，1支持
	CWPFix *int64 `json:"CWPFix,omitnil,omitempty" name:"CWPFix"`

	// 产品支持状态
	DataSupport []*ProductSupport `json:"DataSupport,omitnil,omitempty" name:"DataSupport"`

	// cveId
	CveId *string `json:"CveId,omitnil,omitempty" name:"CveId"`
}

type DbAssetInfo struct {
	// 云防状态
	CFWStatus *uint64 `json:"CFWStatus,omitnil,omitempty" name:"CFWStatus"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// vpc信息
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 公网ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 私网ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// vpc信息
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 云防保护版本
	CFWProtectLevel *uint64 `json:"CFWProtectLevel,omitnil,omitempty" name:"CFWProtectLevel"`

	// tag信息
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`
}

// Predefined struct for user
type DeleteDomainAndIpRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产
	Content []*PublicIpDomainListKey `json:"Content,omitnil,omitempty" name:"Content"`

	// 是否保留路径配置，1：保留，其他：不保留，默认不传为不保留
	RetainPath *int64 `json:"RetainPath,omitnil,omitempty" name:"RetainPath"`

	// 以后是否忽略该资产，，1：忽略，其他：不忽略，默认不传为忽略
	IgnoreAsset *int64 `json:"IgnoreAsset,omitnil,omitempty" name:"IgnoreAsset"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 删除类型，取值： ALL， 删除全部，将直接忽略Content的内容；                                           其他值 ,非全部，则Centent必填，  默认为其他值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DeleteDomainAndIpRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产
	Content []*PublicIpDomainListKey `json:"Content,omitnil,omitempty" name:"Content"`

	// 是否保留路径配置，1：保留，其他：不保留，默认不传为不保留
	RetainPath *int64 `json:"RetainPath,omitnil,omitempty" name:"RetainPath"`

	// 以后是否忽略该资产，，1：忽略，其他：不忽略，默认不传为忽略
	IgnoreAsset *int64 `json:"IgnoreAsset,omitnil,omitempty" name:"IgnoreAsset"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 删除类型，取值： ALL， 删除全部，将直接忽略Content的内容；                                           其他值 ,非全部，则Centent必填，  默认为其他值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DeleteDomainAndIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainAndIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Content")
	delete(f, "RetainPath")
	delete(f, "IgnoreAsset")
	delete(f, "Tags")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainAndIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainAndIpResponseParams struct {
	// 删除的资产数量
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainAndIpResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainAndIpResponseParams `json:"Response"`
}

func (r *DeleteDomainAndIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainAndIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskScanTaskRequestParams struct {
	// 任务id 列表
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DeleteRiskScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id 列表
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DeleteRiskScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRiskScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskScanTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRiskScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRiskScanTaskResponseParams `json:"Response"`
}

func (r *DeleteRiskScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertListRequestParams struct {
	// 标签搜索筛选
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 被调用的集团账号的成员id
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`

	// 0:默认全部 1:资产ID 2:域名
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

type DescribeAlertListRequest struct {
	*tchttp.BaseRequest
	
	// 标签搜索筛选
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 被调用的集团账号的成员id
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`

	// 0:默认全部 1:资产ID 2:域名
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

func (r *DescribeAlertListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	delete(f, "OperatedMemberId")
	delete(f, "AssetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlertListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlertListResponseParams struct {
	// 全量告警列表
	AlertList []*AlertInfo `json:"AlertList,omitnil,omitempty" name:"AlertList"`

	// 告警大类数量
	AlertTypeCount []*TagCount `json:"AlertTypeCount,omitnil,omitempty" name:"AlertTypeCount"`

	// 告警总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 0：succeed 1：timeout
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 返回状态信息
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlertListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlertListResponseParams `json:"Response"`
}

func (r *DescribeAlertListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetProcessListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

type DescribeAssetProcessListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
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
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "Provider")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetProcessListResponseParams struct {
	// 进程数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 进程列表
	AssetProcessList []*AssetProcessItem `json:"AssetProcessList,omitnil,omitempty" name:"AssetProcessList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAssetRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeAssetRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeAssetRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetRiskListResponseParams struct {
	// 资产视角下风险数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产视角下风险列表
	AssetRiskList []*AssetRiskItem `json:"AssetRiskList,omitnil,omitempty" name:"AssetRiskList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetRiskListResponseParams `json:"Response"`
}

func (r *DescribeAssetRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetViewVulRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeAssetViewVulRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeAssetViewVulRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetViewVulRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetViewVulRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetViewVulRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产视角的漏洞风险列表
	Data []*AssetViewVULRiskData `json:"Data,omitnil,omitempty" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// 漏洞类型列表
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// tag枚举
	Tags []*FilterDataObject `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetViewVulRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetViewVulRiskListResponseParams `json:"Response"`
}

func (r *DescribeAssetViewVulRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetViewVulRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCFWAssetStatisticsRequestParams struct {

}

type DescribeCFWAssetStatisticsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCFWAssetStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCFWAssetStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCFWAssetStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCFWAssetStatisticsResponseParams struct {
	// 网络资产总数
	NetworkTotal *int64 `json:"NetworkTotal,omitnil,omitempty" name:"NetworkTotal"`

	// 资产clb数量
	ClbTotal *int64 `json:"ClbTotal,omitnil,omitempty" name:"ClbTotal"`

	// nat数量
	NatTotal *int64 `json:"NatTotal,omitnil,omitempty" name:"NatTotal"`

	// 公网ip数量
	PublicAssetTotal *int64 `json:"PublicAssetTotal,omitnil,omitempty" name:"PublicAssetTotal"`

	// 主机数量
	CVMAssetTotal *int64 `json:"CVMAssetTotal,omitnil,omitempty" name:"CVMAssetTotal"`

	// 配置风险
	CFGTotal *int64 `json:"CFGTotal,omitnil,omitempty" name:"CFGTotal"`

	// 端口风险
	PortTotal *int64 `json:"PortTotal,omitnil,omitempty" name:"PortTotal"`

	// 内容风险
	WebsiteTotal *int64 `json:"WebsiteTotal,omitnil,omitempty" name:"WebsiteTotal"`

	// 风险服务暴露
	ServerTotal *int64 `json:"ServerTotal,omitnil,omitempty" name:"ServerTotal"`

	// 弱口令风险
	WeakPasswordTotal *int64 `json:"WeakPasswordTotal,omitnil,omitempty" name:"WeakPasswordTotal"`

	// 漏洞风险
	VULTotal *int64 `json:"VULTotal,omitnil,omitempty" name:"VULTotal"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCFWAssetStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCFWAssetStatisticsResponseParams `json:"Response"`
}

func (r *DescribeCFWAssetStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCFWAssetStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCSIPRiskStatisticsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCSIPRiskStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCSIPRiskStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCSIPRiskStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCSIPRiskStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCSIPRiskStatisticsResponseParams struct {
	// 资产概况数据
	Data *CsipRiskCenterStatistics `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCSIPRiskStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCSIPRiskStatisticsResponseParams `json:"Response"`
}

func (r *DescribeCSIPRiskStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCSIPRiskStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetInfoRequestParams struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type DescribeCVMAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *DescribeCVMAssetInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCVMAssetInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetInfoResponseParams struct {
	// 数据
	Data *AssetBaseInfoResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCVMAssetInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCVMAssetInfoResponseParams `json:"Response"`
}

func (r *DescribeCVMAssetInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCVMAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCVMAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCVMAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 机器列表
	Data []*CVMAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 防护状态
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil,omitempty" name:"DefenseStatusList"`

	// vpc枚举
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 资产类型枚举
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// 操作系统枚举
	SystemTypeList []*FilterDataObject `json:"SystemTypeList,omitnil,omitempty" name:"SystemTypeList"`

	// ip列表
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitnil,omitempty" name:"IpTypeList"`

	// appid列表
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 可用区列表
	ZoneList []*FilterDataObject `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// os列表
	OsList []*FilterDataObject `json:"OsList,omitnil,omitempty" name:"OsList"`

	// 资产类型和实例类型的对应关系
	AssetMapInstanceTypeList []*AssetInstanceTypeMap `json:"AssetMapInstanceTypeList,omitnil,omitempty" name:"AssetMapInstanceTypeList"`

	// 公网内网枚举
	PublicPrivateAttr []*FilterDataObject `json:"PublicPrivateAttr,omitnil,omitempty" name:"PublicPrivateAttr"`

	// 主机防护状态
	ProtectStatusList []*FilterDataObject `json:"ProtectStatusList,omitnil,omitempty" name:"ProtectStatusList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCVMAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCVMAssetsResponseParams `json:"Response"`
}

func (r *DescribeCVMAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckViewRisksRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeCheckViewRisksRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeCheckViewRisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckViewRisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCheckViewRisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckViewRisksResponseParams struct {
	// 检查视角下风险数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 检查视角下风险列表
	CheckViewRiskList []*CheckViewRiskItem `json:"CheckViewRiskList,omitnil,omitempty" name:"CheckViewRiskList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCheckViewRisksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCheckViewRisksResponseParams `json:"Response"`
}

func (r *DescribeCheckViewRisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckViewRisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeClusterAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeClusterAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAssetsResponseParams struct {
	// 列表
	Data []*AssetCluster `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群类型枚举
	ClusterTypeList []*FilterDataObject `json:"ClusterTypeList,omitnil,omitempty" name:"ClusterTypeList"`

	// 集群状态枚举
	ClusterStatusList []*FilterDataObject `json:"ClusterStatusList,omitnil,omitempty" name:"ClusterStatusList"`

	// 组件状态枚举
	ComponentStatusList []*FilterDataObject `json:"ComponentStatusList,omitnil,omitempty" name:"ComponentStatusList"`

	// 私有网络枚举
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 地域枚举
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 租户枚举
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 集群防护状态枚举
	ProtectStatusList []*FilterDataObject `json:"ProtectStatusList,omitnil,omitempty" name:"ProtectStatusList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAssetsResponseParams `json:"Response"`
}

func (r *DescribeClusterAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPodAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeClusterPodAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeClusterPodAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPodAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPodAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPodAssetsResponseParams struct {
	// 列表
	Data []*AssetClusterPod `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群pod状态枚举
	PodStatusList []*FilterDataObject `json:"PodStatusList,omitnil,omitempty" name:"PodStatusList"`

	// 命名空间枚举
	NamespaceList []*FilterDataObject `json:"NamespaceList,omitnil,omitempty" name:"NamespaceList"`

	// 地域枚举
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 租户枚举
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterPodAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterPodAssetsResponseParams `json:"Response"`
}

func (r *DescribeClusterPodAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPodAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetInfoRequestParams struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type DescribeDbAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *DescribeDbAssetInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbAssetInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetInfoResponseParams struct {
	// db资产详情
	Data *DbAssetInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDbAssetInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbAssetInfoResponseParams `json:"Response"`
}

func (r *DescribeDbAssetInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产类型:MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

type DescribeDbAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产类型:MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

func (r *DescribeDbAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "AssetTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 资产总数
	Data []*DBAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 地域枚举
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 资产类型枚举
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// Vpc枚举
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// Appid枚举
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 公网内网枚举
	PublicPrivateAttr []*FilterDataObject `json:"PublicPrivateAttr,omitnil,omitempty" name:"PublicPrivateAttr"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDbAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbAssetsResponseParams `json:"Response"`
}

func (r *DescribeDbAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 安全中心自定义标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeDomainAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 安全中心自定义标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeDomainAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainAssetsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 域名列表
	Data []*DomainAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 防护状态列表
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil,omitempty" name:"DefenseStatusList"`

	// 资产归属地列表
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitnil,omitempty" name:"AssetLocationList"`

	// 资产类型列表
	SourceTypeList []*FilterDataObject `json:"SourceTypeList,omitnil,omitempty" name:"SourceTypeList"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainAssetsResponseParams `json:"Response"`
}

func (r *DescribeDomainAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposeAssetCategoryRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeExposeAssetCategoryRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeExposeAssetCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposeAssetCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExposeAssetCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposeAssetCategoryResponseParams struct {
	// 云边界分析资产分类列表
	ExposeAssetTypeList []*ExposeAssetTypeItem `json:"ExposeAssetTypeList,omitnil,omitempty" name:"ExposeAssetTypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExposeAssetCategoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExposeAssetCategoryResponseParams `json:"Response"`
}

func (r *DescribeExposeAssetCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposeAssetCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposePathRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 资产域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 端口或端口范围
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

type DescribeExposePathRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 资产域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 端口或端口范围
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

func (r *DescribeExposePathRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposePathRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "AssetId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExposePathRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposePathResponseParams struct {
	// 云边界分析路径节点内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExposePathResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExposePathResponseParams `json:"Response"`
}

func (r *DescribeExposePathResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposePathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposuresRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeExposuresRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeExposuresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposuresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExposuresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposuresResponseParams struct {
	// 云边界分析资产数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云边界分析资产列表
	ExposeList []*ExposesItem `json:"ExposeList,omitnil,omitempty" name:"ExposeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExposuresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExposuresResponseParams `json:"Response"`
}

func (r *DescribeExposuresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExposuresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeGatewayAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeGatewayAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayAssetsResponseParams struct {
	// 列表
	Data []*GateWayAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 资产类型列表
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// vpc列表
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// appid列表
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatewayAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayAssetsResponseParams `json:"Response"`
}

func (r *DescribeGatewayAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHighBaseLineRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 云账号ID
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

type DescribeHighBaseLineRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 云账号ID
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

func (r *DescribeHighBaseLineRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHighBaseLineRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "CloudAccountID")
	delete(f, "Provider")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHighBaseLineRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHighBaseLineRiskListResponseParams struct {
	// 高危基线风险数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 高危基线风险列表
	HighBaseLineRiskList []*HighBaseLineRiskItem `json:"HighBaseLineRiskList,omitnil,omitempty" name:"HighBaseLineRiskList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHighBaseLineRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHighBaseLineRiskListResponseParams `json:"Response"`
}

func (r *DescribeHighBaseLineRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHighBaseLineRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeListenerListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeListenerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerListResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 监听器列表
	Data []*ClbListenerListInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeListenerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenerListResponseParams `json:"Response"`
}

func (r *DescribeListenerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNICAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeNICAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeNICAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNICAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNICAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNICAssetsResponseParams struct {
	// 列表
	Data []*NICAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 资产类型列表
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// vpc列表
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// appid列表
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNICAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNICAssetsResponseParams `json:"Response"`
}

func (r *DescribeNICAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNICAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationInfoRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeOrganizationInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeOrganizationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationInfoResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集团用户列表
	Data []*OrganizationInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationInfoResponseParams `json:"Response"`
}

func (r *DescribeOrganizationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationUserInfoRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 不支持多云
	NotSupportCloud *bool `json:"NotSupportCloud,omitnil,omitempty" name:"NotSupportCloud"`
}

type DescribeOrganizationUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 不支持多云
	NotSupportCloud *bool `json:"NotSupportCloud,omitnil,omitempty" name:"NotSupportCloud"`
}

func (r *DescribeOrganizationUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "NotSupportCloud")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationUserInfoResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集团用户列表
	Data []*OrganizationUserInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 加入方式枚举
	JoinTypeLst []*FilterDataObject `json:"JoinTypeLst,omitnil,omitempty" name:"JoinTypeLst"`

	// 云厂商枚举
	CloudTypeLst []*FilterDataObject `json:"CloudTypeLst,omitnil,omitempty" name:"CloudTypeLst"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationUserInfoResponseParams `json:"Response"`
}

func (r *DescribeOrganizationUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOtherCloudAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// -
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产类型:MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

type DescribeOtherCloudAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// -
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产类型:MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

func (r *DescribeOtherCloudAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOtherCloudAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "AssetTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOtherCloudAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOtherCloudAssetsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 资产总数
	Data []*DBAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 地域枚举
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 资产类型枚举
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// Vpc枚举
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// Appid枚举
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOtherCloudAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOtherCloudAssetsResponseParams `json:"Response"`
}

func (r *DescribeOtherCloudAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOtherCloudAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicIpAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 安全中心自定义标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribePublicIpAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 安全中心自定义标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribePublicIpAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicIpAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicIpAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicIpAssetsResponseParams struct {
	// 列表
	Data []*IpAssetListVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 资产归属地
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitnil,omitempty" name:"AssetLocationList"`

	// ip列表枚举
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitnil,omitempty" name:"IpTypeList"`

	// 地域列表枚举
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 防护枚举
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil,omitempty" name:"DefenseStatusList"`

	// 资产类型枚举
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// AppId枚举
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublicIpAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicIpAssetsResponseParams `json:"Response"`
}

func (r *DescribePublicIpAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicIpAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryImageAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// filter过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeRepositoryImageAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// filter过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeRepositoryImageAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryImageAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepositoryImageAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryImageAssetsResponseParams struct {
	// 仓库镜像列表
	Data []*RepositoryImageVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// region列表
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRepositoryImageAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoryImageAssetsResponseParams `json:"Response"`
}

func (r *DescribeRepositoryImageAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryImageAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewCFGRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewCFGRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewCFGRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewCFGRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewCFGRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewCFGRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产视角的配置风险列表
	Data []*AssetViewCFGRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 配置名列表
	CFGNameLists []*FilterDataObject `json:"CFGNameLists,omitnil,omitempty" name:"CFGNameLists"`

	// 检查类型列表
	CheckTypeLists []*FilterDataObject `json:"CheckTypeLists,omitnil,omitempty" name:"CheckTypeLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewCFGRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewCFGRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewCFGRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewCFGRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewPortRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewPortRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewPortRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewPortRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewPortRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewPortRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产视角的配置风险列表
	Data []*AssetViewPortRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 建议列表
	SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitnil,omitempty" name:"SuggestionLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewPortRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewPortRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewPortRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewPortRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewVULRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewVULRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewVULRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewVULRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewVULRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewVULRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产视角的漏洞风险列表
	Data []*AssetViewVULRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// 漏洞类型列表
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewVULRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewVULRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewVULRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewVULRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewWeakPasswordRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewWeakPasswordRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewWeakPasswordRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewWeakPasswordRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 风险列表
	Data []*AssetViewWeakPassRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// 弱口令类型列表
	PasswordTypeLists []*FilterDataObject `json:"PasswordTypeLists,omitnil,omitempty" name:"PasswordTypeLists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewWeakPasswordRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterCFGViewCFGRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeRiskCenterCFGViewCFGRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeRiskCenterCFGViewCFGRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterCFGViewCFGRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterCFGViewCFGRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterCFGViewCFGRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产视角的配置风险列表
	Data []*CFGViewCFGRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 配置名列表
	CFGNameLists []*FilterDataObject `json:"CFGNameLists,omitnil,omitempty" name:"CFGNameLists"`

	// 检查类型列表
	CheckTypeLists []*FilterDataObject `json:"CheckTypeLists,omitnil,omitempty" name:"CheckTypeLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterCFGViewCFGRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterCFGViewCFGRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterCFGViewCFGRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterCFGViewCFGRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterPortViewPortRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterPortViewPortRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterPortViewPortRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterPortViewPortRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterPortViewPortRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterPortViewPortRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 端口视角的端口风险列表
	Data []*PortViewPortRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 处置建议列表
	SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitnil,omitempty" name:"SuggestionLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterPortViewPortRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterPortViewPortRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterPortViewPortRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterPortViewPortRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterServerRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterServerRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterServerRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterServerRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterServerRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterServerRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 风险服务列表
	Data []*ServerRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 资产类型枚举
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterServerRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterServerRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterServerRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterServerRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterVULViewVULRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterVULViewVULRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterVULViewVULRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterVULViewVULRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterVULViewVULRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterVULViewVULRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 漏洞产视角的漏洞风险列表
	Data []*VULViewVULRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// 漏洞类型列表
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterVULViewVULRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterVULViewVULRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterVULViewVULRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterVULViewVULRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterWebsiteRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterWebsiteRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterWebsiteRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterWebsiteRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterWebsiteRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterWebsiteRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 内容风险列表
	Data []*WebsiteRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// 风险类型列表
	DetectEngineLists []*FilterDataObject `json:"DetectEngineLists,omitnil,omitempty" name:"DetectEngineLists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterWebsiteRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterWebsiteRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterWebsiteRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterWebsiteRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDetailListRequestParams struct {
	// 风险规则ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRiskDetailListRequest struct {
	*tchttp.BaseRequest
	
	// 风险规则ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeRiskDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskRuleId")
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskDetailListResponseParams struct {
	// 资产视角下风险详情数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产视角下风险详情列表
	AssetRiskDetailList []*RiskDetailItem `json:"AssetRiskDetailList,omitnil,omitempty" name:"AssetRiskDetailList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskDetailListResponseParams `json:"Response"`
}

func (r *DescribeRiskDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskRuleDetailRequestParams struct {
	// 风险规则ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`
}

type DescribeRiskRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 风险规则ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`
}

func (r *DescribeRiskRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskRuleDetailResponseParams struct {
	// 风险规则ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 风险名称
	RiskName *string `json:"RiskName,omitnil,omitempty" name:"RiskName"`

	// 风险危害
	RiskInfluence *string `json:"RiskInfluence,omitnil,omitempty" name:"RiskInfluence"`

	// 修复指引
	RiskFixAdvice *string `json:"RiskFixAdvice,omitnil,omitempty" name:"RiskFixAdvice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeRiskRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskRulesRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeRiskRulesRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeRiskRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskRulesResponseParams struct {
	// 风险规则数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 风险规则列表
	RiskRuleList []*RiskRuleItem `json:"RiskRuleList,omitnil,omitempty" name:"RiskRuleList"`

	// 实例类型选项
	InstanceTypeList []*AttributeOptionSet `json:"InstanceTypeList,omitnil,omitempty" name:"InstanceTypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskRulesResponseParams `json:"Response"`
}

func (r *DescribeRiskRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanReportListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 列表过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeScanReportListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 列表过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeScanReportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanReportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanReportListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务日志列表
	Data []*ScanTaskInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 主账户ID列表
	UINList []*string `json:"UINList,omitnil,omitempty" name:"UINList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanReportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanReportListResponseParams `json:"Response"`
}

func (r *DescribeScanReportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanStatisticRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 体检任务id
	TaskLogId *string `json:"TaskLogId,omitnil,omitempty" name:"TaskLogId"`
}

type DescribeScanStatisticRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 体检任务id
	TaskLogId *string `json:"TaskLogId,omitnil,omitempty" name:"TaskLogId"`
}

func (r *DescribeScanStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "TaskLogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanStatisticResponseParams struct {
	// 端口服务数量
	PortServiceCount *uint64 `json:"PortServiceCount,omitnil,omitempty" name:"PortServiceCount"`

	// Web服务数量
	WebAppCount *uint64 `json:"WebAppCount,omitnil,omitempty" name:"WebAppCount"`

	// 弱口令风险数量
	WeakPasswordCount *uint64 `json:"WeakPasswordCount,omitnil,omitempty" name:"WeakPasswordCount"`

	// 漏洞风险数量
	VulCount *uint64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`

	// 高危端口服务数量
	HighRiskPortServiceCount *uint64 `json:"HighRiskPortServiceCount,omitnil,omitempty" name:"HighRiskPortServiceCount"`

	// 风险Web服务数量
	RiskWebAppCount *uint64 `json:"RiskWebAppCount,omitnil,omitempty" name:"RiskWebAppCount"`

	// 端口服务近7天新增数量
	PortServiceIncrement *uint64 `json:"PortServiceIncrement,omitnil,omitempty" name:"PortServiceIncrement"`

	// Web服务近7天新增数量
	WebAppIncrement *uint64 `json:"WebAppIncrement,omitnil,omitempty" name:"WebAppIncrement"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanStatisticResponseParams `json:"Response"`
}

func (r *DescribeScanStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 标签
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeScanTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 标签
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeScanTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskListResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务日志列表
	Data []*ScanTaskInfoList `json:"Data,omitnil,omitempty" name:"Data"`

	// 主账户ID列表
	UINList []*string `json:"UINList,omitnil,omitempty" name:"UINList"`

	// 体检模式过滤列表
	TaskModeList []*FilterDataObject `json:"TaskModeList,omitnil,omitempty" name:"TaskModeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanTaskListResponseParams `json:"Response"`
}

func (r *DescribeScanTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchBugInfoRequestParams struct {
	// 1的时候返回应急漏洞，2的时候返回应急漏洞列表，3的时候搭配输入CVEId字段展示该漏洞数据
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// id=3时传入该参数
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`
}

type DescribeSearchBugInfoRequest struct {
	*tchttp.BaseRequest
	
	// 1的时候返回应急漏洞，2的时候返回应急漏洞列表，3的时候搭配输入CVEId字段展示该漏洞数据
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// id=3时传入该参数
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`
}

func (r *DescribeSearchBugInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchBugInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "CVEId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchBugInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchBugInfoResponseParams struct {
	// 漏洞信息和资产信息
	Data *DataSearchBug `json:"Data,omitnil,omitempty" name:"Data"`

	// 状态值，0：查询成功，非0：查询失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 状态信息，success：查询成功，fail：查询失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSearchBugInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchBugInfoResponseParams `json:"Response"`
}

func (r *DescribeSearchBugInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchBugInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubUserInfoRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeSubUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeSubUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubUserInfoResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 子用户列表
	Data []*SubUserInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 厂商枚举列表
	CloudTypeLst []*FilterDataObject `json:"CloudTypeLst,omitnil,omitempty" name:"CloudTypeLst"`

	// 所属主账号appid枚举
	OwnerAppIDLst []*FilterDataObject `json:"OwnerAppIDLst,omitnil,omitempty" name:"OwnerAppIDLst"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubUserInfoResponseParams `json:"Response"`
}

func (r *DescribeSubUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeSubnetAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeSubnetAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetAssetsResponseParams struct {
	// 列表
	Data []*SubnetAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// vpc列表
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// appid列表
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 可用区列表
	ZoneList []*FilterDataObject `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubnetAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetAssetsResponseParams `json:"Response"`
}

func (r *DescribeSubnetAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeTaskLogListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeTaskLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 报告列表
	Data []*TaskLogInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 待查看数量
	NotViewNumber *int64 `json:"NotViewNumber,omitnil,omitempty" name:"NotViewNumber"`

	// 报告模板数
	ReportTemplateNumber *int64 `json:"ReportTemplateNumber,omitnil,omitempty" name:"ReportTemplateNumber"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLogListResponseParams `json:"Response"`
}

func (r *DescribeTaskLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogURLRequestParams struct {
	// 0: 预览， 1: 下载
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 任务报告Id 列表
	ReportItemKeyList []*ReportItemKey `json:"ReportItemKeyList,omitnil,omitempty" name:"ReportItemKeyList"`

	// 报告中任务id列表
	ReportTaskIdList []*ReportTaskIdList `json:"ReportTaskIdList,omitnil,omitempty" name:"ReportTaskIdList"`
}

type DescribeTaskLogURLRequest struct {
	*tchttp.BaseRequest
	
	// 0: 预览， 1: 下载
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 任务报告Id 列表
	ReportItemKeyList []*ReportItemKey `json:"ReportItemKeyList,omitnil,omitempty" name:"ReportItemKeyList"`

	// 报告中任务id列表
	ReportTaskIdList []*ReportTaskIdList `json:"ReportTaskIdList,omitnil,omitempty" name:"ReportTaskIdList"`
}

func (r *DescribeTaskLogURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "MemberId")
	delete(f, "ReportItemKeyList")
	delete(f, "ReportTaskIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogURLResponseParams struct {
	// 返回报告临时下载url
	Data []*TaskLogURL `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskLogURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLogURLResponseParams `json:"Response"`
}

func (r *DescribeTaskLogURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopAttackInfoRequestParams struct {
	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 1:攻击类型 2:攻击者
	QueryType *int64 `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 被调用的集团账号的成员id
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`

	// 资产名称
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 0: 默认全部 1:资产ID 2:域名
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

type DescribeTopAttackInfoRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 1:攻击类型 2:攻击者
	QueryType *int64 `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 被调用的集团账号的成员id
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`

	// 资产名称
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 0: 默认全部 1:资产ID 2:域名
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

func (r *DescribeTopAttackInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopAttackInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "MemberId")
	delete(f, "OperatedMemberId")
	delete(f, "AssetName")
	delete(f, "AssetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopAttackInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopAttackInfoResponseParams struct {
	// Top攻击类型/攻击者次数
	TopAttackInfo []*TagCount `json:"TopAttackInfo,omitnil,omitempty" name:"TopAttackInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopAttackInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopAttackInfoResponseParams `json:"Response"`
}

func (r *DescribeTopAttackInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopAttackInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUebaRuleRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeUebaRuleRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeUebaRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUebaRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUebaRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUebaRuleResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 策略列表
	Data []*UebaRule `json:"Data,omitnil,omitempty" name:"Data"`

	// 自定义策略对应的告警类别枚举
	AlterType []*FilterDataObject `json:"AlterType,omitnil,omitempty" name:"AlterType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUebaRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUebaRuleResponseParams `json:"Response"`
}

func (r *DescribeUebaRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUebaRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 查询条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeVULListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 查询条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeVULListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVULListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULListResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 漏洞列表
	Data []*VULBaseInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 漏洞类型列表
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// 风险等级列表
	RiskLevels []*FilterDataObject `json:"RiskLevels,omitnil,omitempty" name:"RiskLevels"`

	// 标签
	Tags []*FilterDataObject `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 产品支持情况
	ProductSupport []*FilterDataObject `json:"ProductSupport,omitnil,omitempty" name:"ProductSupport"`

	// 产品支持情况
	CheckStatus []*FilterDataObject `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// 攻击热度枚举
	AttackHeat []*FilterDataObject `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVULListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVULListResponseParams `json:"Response"`
}

func (r *DescribeVULListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskAdvanceCFGListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeVULRiskAdvanceCFGListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeVULRiskAdvanceCFGListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskAdvanceCFGListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "TaskId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVULRiskAdvanceCFGListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskAdvanceCFGListResponseParams struct {
	// 配置项列表
	Data []*VULRiskAdvanceCFGList `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 风险等级过滤列表
	RiskLevelLists []*FilterDataObject `json:"RiskLevelLists,omitnil,omitempty" name:"RiskLevelLists"`

	// 漏洞类型过滤列表
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// 识别来源过滤列表
	CheckFromLists []*FilterDataObject `json:"CheckFromLists,omitnil,omitempty" name:"CheckFromLists"`

	// 漏洞标签列表
	VulTagList []*FilterDataObject `json:"VulTagList,omitnil,omitempty" name:"VulTagList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVULRiskAdvanceCFGListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVULRiskAdvanceCFGListResponseParams `json:"Response"`
}

func (r *DescribeVULRiskAdvanceCFGListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskAdvanceCFGListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskDetailRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// pcMgrId
	PCMGRId *string `json:"PCMGRId,omitnil,omitempty" name:"PCMGRId"`
}

type DescribeVULRiskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// pcMgrId
	PCMGRId *string `json:"PCMGRId,omitnil,omitempty" name:"PCMGRId"`
}

func (r *DescribeVULRiskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "RiskId")
	delete(f, "PCMGRId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVULRiskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskDetailResponseParams struct {
	// 安全产品支持情况
	ServiceSupport []*ServiceSupport `json:"ServiceSupport,omitnil,omitempty" name:"ServiceSupport"`

	// 漏洞趋势
	VulTrend []*VulTrend `json:"VulTrend,omitnil,omitempty" name:"VulTrend"`

	// 漏洞补充信息
	VulData *VULRiskInfo `json:"VulData,omitnil,omitempty" name:"VulData"`

	// 小助手问答id
	QuestionId *string `json:"QuestionId,omitnil,omitempty" name:"QuestionId"`

	// 会话id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVULRiskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVULRiskDetailResponseParams `json:"Response"`
}

func (r *DescribeVULRiskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeVpcAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤参数
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeVpcAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAssetsResponseParams struct {
	// 列表
	Data []*Vpc `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// vpc列表
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// appid列表
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcAssetsResponseParams `json:"Response"`
}

func (r *DescribeVpcAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 云账号ID
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

type DescribeVulRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 云账号ID
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
}

func (r *DescribeVulRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "CloudAccountID")
	delete(f, "Provider")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulRiskListResponseParams struct {
	// 漏洞数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 漏洞列表
	VulRiskList []*VulRiskItem `json:"VulRiskList,omitnil,omitempty" name:"VulRiskList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVulRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulRiskListResponseParams `json:"Response"`
}

func (r *DescribeVulRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulViewVulRiskListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeVulViewVulRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 资产标签
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeVulViewVulRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulViewVulRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulViewVulRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulViewVulRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 漏洞产视角的漏洞风险列表
	Data []*VULViewVULRiskData `json:"Data,omitnil,omitempty" name:"Data"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// 漏洞类型列表
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// tag枚举
	Tags []*FilterDataObject `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVulViewVulRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulViewVulRiskListResponseParams `json:"Response"`
}

func (r *DescribeVulViewVulRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulViewVulRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAssetVO struct {
	// 资产id
	AssetId []*string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	AssetName []*string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产类型
	AssetType []*string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 地域
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// Waf状态
	WAFStatus *uint64 `json:"WAFStatus,omitnil,omitempty" name:"WAFStatus"`

	// 资产创建时间
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// Appid
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号id
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 账号名称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 是否核心
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 是否云上资产
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// 网络攻击
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// 网络访问
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// 网络拦截
	Intercept *uint64 `json:"Intercept,omitnil,omitempty" name:"Intercept"`

	// 入站峰值带宽
	InBandwidth *string `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`

	// 出站峰值带宽
	OutBandwidth *string `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// 入站累计流量
	InFlow *string `json:"InFlow,omitnil,omitempty" name:"InFlow"`

	// 出站累计流量
	OutFlow *string `json:"OutFlow,omitnil,omitempty" name:"OutFlow"`

	// 最近扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 端口风险
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// 漏洞风险
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil,omitempty" name:"VulnerabilityRisk"`

	// 配置风险
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// 扫描任务
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// 域名
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 解析ip
	SeverIp []*string `json:"SeverIp,omitnil,omitempty" name:"SeverIp"`

	// bot攻击数量
	BotCount *uint64 `json:"BotCount,omitnil,omitempty" name:"BotCount"`

	// 弱口令风险
	WeakPassword *uint64 `json:"WeakPassword,omitnil,omitempty" name:"WeakPassword"`

	// 内容风险
	WebContentRisk *uint64 `json:"WebContentRisk,omitnil,omitempty" name:"WebContentRisk"`

	// tag标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 关联实例类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// memberId信息
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// cc攻击
	CCAttack *int64 `json:"CCAttack,omitnil,omitempty" name:"CCAttack"`

	// web攻击
	WebAttack *int64 `json:"WebAttack,omitnil,omitempty" name:"WebAttack"`

	// 风险服务暴露数量
	ServiceRisk *uint64 `json:"ServiceRisk,omitnil,omitempty" name:"ServiceRisk"`

	// 是否新资产 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// 待确认资产的随机三级域名
	VerifyDomain *string `json:"VerifyDomain,omitnil,omitempty" name:"VerifyDomain"`

	// 待确认资产的TXT记录内容
	VerifyTXTRecord *string `json:"VerifyTXTRecord,omitnil,omitempty" name:"VerifyTXTRecord"`

	// 待确认资产的认证状态，0-待认证，1-认证成功，2-认证中，3-txt认证失败，4-人工认证失败
	VerifyStatus *int64 `json:"VerifyStatus,omitnil,omitempty" name:"VerifyStatus"`

	// bot访问数量
	BotAccessCount *int64 `json:"BotAccessCount,omitnil,omitempty" name:"BotAccessCount"`
}

type Element struct {
	// 统计类型
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 统计对象
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ExposeAssetTypeItem struct {
	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 云厂商名称
	ProviderName *string `json:"ProviderName,omitnil,omitempty" name:"ProviderName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 资产类型名称
	AssetTypeName *string `json:"AssetTypeName,omitnil,omitempty" name:"AssetTypeName"`
}

type ExposesItem struct {
	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 云账号名称
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// 云账号
	CloudAccountId *string `json:"CloudAccountId,omitnil,omitempty" name:"CloudAccountId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 端口或者端口范围
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 开放
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 风险类型
	RiskType *string `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// acl类型
	AclType *string `json:"AclType,omitnil,omitempty" name:"AclType"`

	// acl列表
	AclList *string `json:"AclList,omitnil,omitempty" name:"AclList"`

	// 资产ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 端口服务数量
	PortServiceCount *uint64 `json:"PortServiceCount,omitnil,omitempty" name:"PortServiceCount"`

	// 高危端口数量
	HighRiskPortServiceCount *uint64 `json:"HighRiskPortServiceCount,omitnil,omitempty" name:"HighRiskPortServiceCount"`

	// web应用数量
	WebAppCount *uint64 `json:"WebAppCount,omitnil,omitempty" name:"WebAppCount"`

	// 有风险web应用数量
	RiskWebAppCount *uint64 `json:"RiskWebAppCount,omitnil,omitempty" name:"RiskWebAppCount"`

	// 弱口令数量
	WeakPasswordCount *uint64 `json:"WeakPasswordCount,omitnil,omitempty" name:"WeakPasswordCount"`

	// 漏洞数量
	VulCount *uint64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`

	// 首次发现时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 实例类型名称
	AssetTypeName *string `json:"AssetTypeName,omitnil,omitempty" name:"AssetTypeName"`

	// 开放状态
	DisplayStatus *string `json:"DisplayStatus,omitnil,omitempty" name:"DisplayStatus"`

	// 端口状态
	DisplayRiskType *string `json:"DisplayRiskType,omitnil,omitempty" name:"DisplayRiskType"`

	// 扫描任务状态
	ScanTaskStatus *string `json:"ScanTaskStatus,omitnil,omitempty" name:"ScanTaskStatus"`

	// uuid
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 是否进行过安全体检
	HasScan *string `json:"HasScan,omitnil,omitempty" name:"HasScan"`

	// 租户ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 租户ID字符串
	AppIdStr *string `json:"AppIdStr,omitnil,omitempty" name:"AppIdStr"`
}

type Filter struct {
	// 查询数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询偏移位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序采用升序还是降序 升:asc 降 desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 需排序的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 过滤的列及内容
	Filters []*WhereFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 可填无， 日志使用查询时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 可填无， 日志使用查询时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type FilterDataObject struct {
	// 英文翻译
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 中文翻译
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type Filters struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例ID内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 模糊匹配
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExactMatch *string `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type GateWayAsset struct {
	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 资产ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 私有ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 公网ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络名
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 出向峰值带宽
	OutboundPeakBandwidth *string `json:"OutboundPeakBandwidth,omitnil,omitempty" name:"OutboundPeakBandwidth"`

	// 入向峰值带宽
	InboundPeakBandwidth *string `json:"InboundPeakBandwidth,omitnil,omitempty" name:"InboundPeakBandwidth"`

	// 出站累计流量
	OutboundCumulativeFlow *string `json:"OutboundCumulativeFlow,omitnil,omitempty" name:"OutboundCumulativeFlow"`

	// 入站累计流量
	InboundCumulativeFlow *string `json:"InboundCumulativeFlow,omitnil,omitempty" name:"InboundCumulativeFlow"`

	// 网络攻击
	NetworkAttack *int64 `json:"NetworkAttack,omitnil,omitempty" name:"NetworkAttack"`

	// 暴露端口
	ExposedPort *int64 `json:"ExposedPort,omitnil,omitempty" name:"ExposedPort"`

	// 暴露漏洞
	ExposedVUL *int64 `json:"ExposedVUL,omitnil,omitempty" name:"ExposedVUL"`

	// 配置风险
	ConfigureRisk *int64 `json:"ConfigureRisk,omitnil,omitempty" name:"ConfigureRisk"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务数
	ScanTask *int64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// 最后扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// ipv6地址
	AddressIPV6 *string `json:"AddressIPV6,omitnil,omitempty" name:"AddressIPV6"`

	// 是否核心
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 风险服务暴露
	RiskExposure *int64 `json:"RiskExposure,omitnil,omitempty" name:"RiskExposure"`

	// 是否新资产 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// 网关状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// TSE的网关真实地域
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`
}

type HighBaseLineRiskItem struct {
	// 云账号ID
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// 实例ID
	AssetID *string `json:"AssetID,omitnil,omitempty" name:"AssetID"`

	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 风险名称
	RiskName *string `json:"RiskName,omitnil,omitempty" name:"RiskName"`

	// 风险分类
	RiskCategory *string `json:"RiskCategory,omitnil,omitempty" name:"RiskCategory"`

	// 风险等级
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 风险描述
	RiskDesc *string `json:"RiskDesc,omitnil,omitempty" name:"RiskDesc"`

	// 风险结果
	RiskResult *string `json:"RiskResult,omitnil,omitempty" name:"RiskResult"`

	// 修复建议
	FixAdvice *string `json:"FixAdvice,omitnil,omitempty" name:"FixAdvice"`

	// Linux漏洞
	RiskCategoryName *string `json:"RiskCategoryName,omitnil,omitempty" name:"RiskCategoryName"`

	// 风险等级名称
	RiskLevelName *string `json:"RiskLevelName,omitnil,omitempty" name:"RiskLevelName"`

	// 实例状态
	InstanceStatusName *string `json:"InstanceStatusName,omitnil,omitempty" name:"InstanceStatusName"`

	// 首次发现时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近发现时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 租户ID
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`
}

type IpAssetListVO struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 云防状态
	CFWStatus *uint64 `json:"CFWStatus,omitnil,omitempty" name:"CFWStatus"`

	// 资产创建时间
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 公网ip类型
	PublicIpType *uint64 `json:"PublicIpType,omitnil,omitempty" name:"PublicIpType"`

	// vpc
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc名
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// appid
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 名称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 核心
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 云上
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// 网络攻击
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// 网络访问
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// 网络拦截
	Intercept *uint64 `json:"Intercept,omitnil,omitempty" name:"Intercept"`

	// 入向带宽
	InBandwidth *string `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`

	// 出向带宽
	OutBandwidth *string `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// 入向流量
	InFlow *string `json:"InFlow,omitnil,omitempty" name:"InFlow"`

	// 出向流量
	OutFlow *string `json:"OutFlow,omitnil,omitempty" name:"OutFlow"`

	// 最近扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 端口风险
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// 漏洞风险
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil,omitempty" name:"VulnerabilityRisk"`

	// 配置风险
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// 扫描任务
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// 弱口令
	WeakPassword *uint64 `json:"WeakPassword,omitnil,omitempty" name:"WeakPassword"`

	// 内容风险
	WebContentRisk *uint64 `json:"WebContentRisk,omitnil,omitempty" name:"WebContentRisk"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// eip主键
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// MemberId信息
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 风险服务暴露
	RiskExposure *int64 `json:"RiskExposure,omitnil,omitempty" name:"RiskExposure"`

	// 是否新资产 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// 资产认证状态，0-待认证，1-认证成功，2-认证中，3+-认证失败
	VerifyStatus *int64 `json:"VerifyStatus,omitnil,omitempty" name:"VerifyStatus"`
}

type KeyValue struct {
	// 字段
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyOrganizationAccountStatusRequestParams struct {
	// 修改集团账号状态，1 开启， 2关闭
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyOrganizationAccountStatusRequest struct {
	*tchttp.BaseRequest
	
	// 修改集团账号状态，1 开启， 2关闭
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *ModifyOrganizationAccountStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrganizationAccountStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOrganizationAccountStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrganizationAccountStatusResponseParams struct {
	// 返回值为0，则修改成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyOrganizationAccountStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOrganizationAccountStatusResponseParams `json:"Response"`
}

func (r *ModifyOrganizationAccountStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrganizationAccountStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskCenterRiskStatusRequestParams struct {
	// 风险资产相关数据
	RiskStatusKeys []*RiskCenterStatusKey `json:"RiskStatusKeys,omitnil,omitempty" name:"RiskStatusKeys"`

	// 处置状态，1为已处置、2为已忽略，3为取消已处置，4为取消已忽略
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 风险类型，0-端口风险， 1-漏洞风险，2-弱口令风险， 3-网站内容风险，4-配置风险，5-风险服务暴露
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyRiskCenterRiskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 风险资产相关数据
	RiskStatusKeys []*RiskCenterStatusKey `json:"RiskStatusKeys,omitnil,omitempty" name:"RiskStatusKeys"`

	// 处置状态，1为已处置、2为已忽略，3为取消已处置，4为取消已忽略
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 风险类型，0-端口风险， 1-漏洞风险，2-弱口令风险， 3-网站内容风险，4-配置风险，5-风险服务暴露
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *ModifyRiskCenterRiskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterRiskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskStatusKeys")
	delete(f, "Status")
	delete(f, "Type")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRiskCenterRiskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskCenterRiskStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRiskCenterRiskStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRiskCenterRiskStatusResponseParams `json:"Response"`
}

func (r *ModifyRiskCenterRiskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterRiskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskCenterScanTaskRequestParams struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// 扫描项目；port/poc/weakpass/webcontent/configrisk
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// 要修改的任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 扫描资产信息列表
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// 扫描计划详情
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// ip/域名/url数组
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// 高级配置
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// 任务完成回调webhook地址
	FinishWebHook *string `json:"FinishWebHook,omitnil,omitempty" name:"FinishWebHook"`
}

type ModifyRiskCenterScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// 扫描项目；port/poc/weakpass/webcontent/configrisk
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// 要修改的任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 扫描资产信息列表
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// 扫描计划详情
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// ip/域名/url数组
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// 高级配置
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// 任务完成回调webhook地址
	FinishWebHook *string `json:"FinishWebHook,omitnil,omitempty" name:"FinishWebHook"`
}

func (r *ModifyRiskCenterScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "ScanAssetType")
	delete(f, "ScanItem")
	delete(f, "ScanPlanType")
	delete(f, "TaskId")
	delete(f, "MemberId")
	delete(f, "Assets")
	delete(f, "ScanPlanContent")
	delete(f, "SelfDefiningAssets")
	delete(f, "TaskAdvanceCFG")
	delete(f, "TaskMode")
	delete(f, "FinishWebHook")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRiskCenterScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskCenterScanTaskResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 0，修改成功，其他失败；-1为存在资产未认证
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 未认证资产列表
	UnAuthAsset []*string `json:"UnAuthAsset,omitnil,omitempty" name:"UnAuthAsset"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRiskCenterScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRiskCenterScanTaskResponseParams `json:"Response"`
}

func (r *ModifyRiskCenterScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUebaRuleSwitchRequestParams struct {
	// 策略ID
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 开关状态
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyUebaRuleSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 开关状态
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *ModifyUebaRuleSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUebaRuleSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	delete(f, "Status")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUebaRuleSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUebaRuleSwitchResponseParams struct {
	// 0成功，1失败
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 返回信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUebaRuleSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUebaRuleSwitchResponseParams `json:"Response"`
}

func (r *ModifyUebaRuleSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUebaRuleSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NICAsset struct {
	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 资产ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 私有ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 公网ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络名
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 出向峰值带宽
	OutboundPeakBandwidth *string `json:"OutboundPeakBandwidth,omitnil,omitempty" name:"OutboundPeakBandwidth"`

	// 入向峰值带宽
	InboundPeakBandwidth *string `json:"InboundPeakBandwidth,omitnil,omitempty" name:"InboundPeakBandwidth"`

	// 出站累计流量
	OutboundCumulativeFlow *string `json:"OutboundCumulativeFlow,omitnil,omitempty" name:"OutboundCumulativeFlow"`

	// 入站累计流量
	InboundCumulativeFlow *string `json:"InboundCumulativeFlow,omitnil,omitempty" name:"InboundCumulativeFlow"`

	// 网络攻击
	NetworkAttack *int64 `json:"NetworkAttack,omitnil,omitempty" name:"NetworkAttack"`

	// 暴露端口
	ExposedPort *int64 `json:"ExposedPort,omitnil,omitempty" name:"ExposedPort"`

	// 暴露漏洞
	ExposedVUL *int64 `json:"ExposedVUL,omitnil,omitempty" name:"ExposedVUL"`

	// 配置风险
	ConfigureRisk *int64 `json:"ConfigureRisk,omitnil,omitempty" name:"ConfigureRisk"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务数
	ScanTask *int64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// 最后扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 是否核心
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 是否新资产 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type NewAlertKey struct {
	// 需要更改的用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 告警类别
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警子类别
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 告警来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 告警名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 告警key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 时间
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type OrganizationInfo struct {
	// 成员账号名称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 部门节点名称，账号所属部门
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Member/Admin/DelegatedAdmin/EntityAdmin; 成员/管理员/委派管理员/主体管理员
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 成员账号id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 账号加入方式,create/invite
	JoinType *string `json:"JoinType,omitnil,omitempty" name:"JoinType"`

	// 集团名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 管理员账号名称
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 管理员Uin
	AdminUin *string `json:"AdminUin,omitnil,omitempty" name:"AdminUin"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 部门数
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 成员数
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`

	// 子账号数
	SubAccountCount *int64 `json:"SubAccountCount,omitnil,omitempty" name:"SubAccountCount"`

	// 异常子账号数量
	AbnormalSubUserCount *int64 `json:"AbnormalSubUserCount,omitnil,omitempty" name:"AbnormalSubUserCount"`

	// 集团关系策略权限
	GroupPermission []*string `json:"GroupPermission,omitnil,omitempty" name:"GroupPermission"`

	// 成员关系策略权限
	MemberPermission []*string `json:"MemberPermission,omitnil,omitempty" name:"MemberPermission"`

	// 集团付费模式；0/自付费，1/代付费
	GroupPayMode *int64 `json:"GroupPayMode,omitnil,omitempty" name:"GroupPayMode"`

	// 个人付费模式；0/自付费，1/代付费
	MemberPayMode *int64 `json:"MemberPayMode,omitnil,omitempty" name:"MemberPayMode"`

	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本
	CFWProtect *string `json:"CFWProtect,omitnil,omitempty" name:"CFWProtect"`

	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本
	WAFProtect *string `json:"WAFProtect,omitnil,omitempty" name:"WAFProtect"`

	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本
	CWPProtect *string `json:"CWPProtect,omitnil,omitempty" name:"CWPProtect"`

	// 所有部门的集合数组
	Departments []*string `json:"Departments,omitnil,omitempty" name:"Departments"`

	// 成员创建时间
	MemberCreateTime *string `json:"MemberCreateTime,omitnil,omitempty" name:"MemberCreateTime"`

	// Advanced/Enterprise/Ultimate 
	CSIPProtect *string `json:"CSIPProtect,omitnil,omitempty" name:"CSIPProtect"`

	// 1表示配额消耗方
	QuotaConsumer *int64 `json:"QuotaConsumer,omitnil,omitempty" name:"QuotaConsumer"`

	// 管理员/委派管理员 已开启数量
	EnableAdminCount *int64 `json:"EnableAdminCount,omitnil,omitempty" name:"EnableAdminCount"`

	// 账户多云信息统计，数组形式，具体参考CloudCountDesc描述
	CloudCountDesc []*CloudCountDesc `json:"CloudCountDesc,omitnil,omitempty" name:"CloudCountDesc"`

	// 管理员/委派管理员 总数量
	AdminCount *int64 `json:"AdminCount,omitnil,omitempty" name:"AdminCount"`
}

type OrganizationUserInfo struct {
	// 成员账号Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 成员账号名称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 部门节点名称，账号所属部门
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 资产数量
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// 风险数量
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 攻击数量
	AttackCount *int64 `json:"AttackCount,omitnil,omitempty" name:"AttackCount"`

	// Member/Admin/;成员或者管理员
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 成员账号id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 成员账号Appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号加入方式,create/invite
	JoinType *string `json:"JoinType,omitnil,omitempty" name:"JoinType"`

	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本
	CFWProtect *string `json:"CFWProtect,omitnil,omitempty" name:"CFWProtect"`

	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本
	WAFProtect *string `json:"WAFProtect,omitnil,omitempty" name:"WAFProtect"`

	// 空则未开启，否则不同字符串对应不同版本，common为通用，不区分版本
	CWPProtect *string `json:"CWPProtect,omitnil,omitempty" name:"CWPProtect"`

	// 1启用，0未启用
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// "Free"       //免费版  "Advanced"   //高级版 "Enterprise" //企业版 "Ultimate"   //旗舰版
	CSIPProtect *string `json:"CSIPProtect,omitnil,omitempty" name:"CSIPProtect"`

	// 1为配额消耗者
	QuotaConsumer *int64 `json:"QuotaConsumer,omitnil,omitempty" name:"QuotaConsumer"`

	// 账户类型，0为腾讯云账户，1为AWS账户
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 0为缺省值，1为10分钟，2为1小时，3为24小时
	SyncFrequency *int64 `json:"SyncFrequency,omitnil,omitempty" name:"SyncFrequency"`

	// 多云账户是否过期
	IsExpired *bool `json:"IsExpired,omitnil,omitempty" name:"IsExpired"`

	// 多云账户 权限列表
	PermissionList []*string `json:"PermissionList,omitnil,omitempty" name:"PermissionList"`

	// 1
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 0 腾讯云集团账户
	// 1 腾讯云接入账户
	// 2 非腾讯云
	TcMemberType *int64 `json:"TcMemberType,omitnil,omitempty" name:"TcMemberType"`

	// 子账号数量
	SubUserCount *int64 `json:"SubUserCount,omitnil,omitempty" name:"SubUserCount"`

	// 加入方式详细信息
	JoinTypeInfo *string `json:"JoinTypeInfo,omitnil,omitempty" name:"JoinTypeInfo"`
}

type PortRiskAdvanceCFGParamItem struct {
	// 端口集合,以逗号分隔
	PortSets *string `json:"PortSets,omitnil,omitempty" name:"PortSets"`

	// 检测项类型，0-系统定义，1-用户自定义
	CheckType *int64 `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// 检测项描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 是否启用，1-启用，0-禁用
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type PortViewPortRisk struct {
	// 未处理数量
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 处置建议,0保持现状、1限制访问、2封禁端口
	Suggestion *uint64 `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 影响资产数量
	AffectAssetCount *string `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 识别来源
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 前端索引
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 服务
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`
}

type ProductSupport struct {
	// true支持扫描。false不支持扫描
	VSSScan *bool `json:"VSSScan,omitnil,omitempty" name:"VSSScan"`

	// 0不支持，1支持
	CWPScan *string `json:"CWPScan,omitnil,omitempty" name:"CWPScan"`

	// 1支持虚拟补丁，0或空不支持
	CFWPatch *string `json:"CFWPatch,omitnil,omitempty" name:"CFWPatch"`

	// 0不支持，1支持	
	WafPatch *int64 `json:"WafPatch,omitnil,omitempty" name:"WafPatch"`

	// 0不支持，1支持	
	CWPFix *int64 `json:"CWPFix,omitnil,omitempty" name:"CWPFix"`

	// cveid
	CveId *string `json:"CveId,omitnil,omitempty" name:"CveId"`
}

type PublicIpDomainListKey struct {
	// 资产值
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`
}

type RelatedEvent struct {
	// 事件ID
	EventID *string `json:"EventID,omitnil,omitempty" name:"EventID"`

	// 事件描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 与事件关联的告警数量
	RelatedCount *int64 `json:"RelatedCount,omitnil,omitempty" name:"RelatedCount"`
}

type ReportItemKey struct {
	// 日志Id列表
	TaskLogList []*string `json:"TaskLogList,omitnil,omitempty" name:"TaskLogList"`
}

type ReportTaskIdList struct {
	// 任务id列表
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 租户ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type RepositoryImageVO struct {
	// 用户appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 镜像id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 镜像名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 镜像创建时间
	InstanceCreateTime *string `json:"InstanceCreateTime,omitnil,omitempty" name:"InstanceCreateTime"`

	// 镜像大小带单位
	InstanceSize *string `json:"InstanceSize,omitnil,omitempty" name:"InstanceSize"`

	// 构建次数
	BuildCount *int64 `json:"BuildCount,omitnil,omitempty" name:"BuildCount"`

	// 镜像类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 授权状态
	AuthStatus *int64 `json:"AuthStatus,omitnil,omitempty" name:"AuthStatus"`

	// 镜像版本
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 仓库地址
	RepositoryUrl *string `json:"RepositoryUrl,omitnil,omitempty" name:"RepositoryUrl"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 是否核心
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 漏洞风险
	VulRisk *int64 `json:"VulRisk,omitnil,omitempty" name:"VulRisk"`

	// 检查任务
	CheckCount *int64 `json:"CheckCount,omitnil,omitempty" name:"CheckCount"`

	// 体检时间
	CheckTime *string `json:"CheckTime,omitnil,omitempty" name:"CheckTime"`

	// 是否新资产 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type RiskCenterStatusKey struct {
	// 风险ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 公网IP/域名
	PublicIPDomain *string `json:"PublicIPDomain,omitnil,omitempty" name:"PublicIPDomain"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// APP ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type RiskDetailItem struct {
	// 首次发现时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 风险状态
	RiskStatus *int64 `json:"RiskStatus,omitnil,omitempty" name:"RiskStatus"`

	// 风险内容
	RiskContent *string `json:"RiskContent,omitnil,omitempty" name:"RiskContent"`

	// 云厂商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 云厂商名称
	ProviderName *string `json:"ProviderName,omitnil,omitempty" name:"ProviderName"`

	// 云账号
	CloudAccountId *string `json:"CloudAccountId,omitnil,omitempty" name:"CloudAccountId"`

	// 云账号名称
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 风险ID
	RiskId *uint64 `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 风险规则ID
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// 风险验证状态
	CheckStatus *string `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`
}

type RiskRuleItem struct {
	// 风险检查项ID
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// 云厂商名称
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例类型名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 风险名称
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// 检查类型
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// 风险等级
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 风险危害
	RiskInfluence *string `json:"RiskInfluence,omitnil,omitempty" name:"RiskInfluence"`
}

type RoleInfo struct {
	// IP
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// HostIP
	HostIP *string `json:"HostIP,omitnil,omitempty" name:"HostIP"`

	// 原始IP
	OriginIP *string `json:"OriginIP,omitnil,omitempty" name:"OriginIP"`

	// 端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 资产ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 纬度
	Latitude *string `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// 经度
	Longitude *string `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// 信息
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 企业名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 家族团伙
	Family *string `json:"Family,omitnil,omitempty" name:"Family"`

	// 病毒名
	VirusName *string `json:"VirusName,omitnil,omitempty" name:"VirusName"`

	// MD5值
	MD5 *string `json:"MD5,omitnil,omitempty" name:"MD5"`

	// 恶意进程文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 1:主机资产 2:域名资产 3:网络资产
	AssetType *int64 `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 来源日志分析的信息字段
	FromLogAnalysisData []*KeyValue `json:"FromLogAnalysisData,omitnil,omitempty" name:"FromLogAnalysisData"`

	// 容器名
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`

	// 容器ID
	ContainerID *string `json:"ContainerID,omitnil,omitempty" name:"ContainerID"`
}

type ScanTaskInfo struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务状态码：1等待开始  2正在扫描  3扫描出错 4扫描完成
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务进度
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务完成时间
	TaskTime *string `json:"TaskTime,omitnil,omitempty" name:"TaskTime"`

	// 报告ID
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// 报告名称
	ReportName *string `json:"ReportName,omitnil,omitempty" name:"ReportName"`

	// 扫描计划，0-周期任务,1-立即扫描,2-定时扫描,3-自定义
	ScanPlan *int64 `json:"ScanPlan,omitnil,omitempty" name:"ScanPlan"`

	// 关联的资产数
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// APP ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户主账户ID
	UIN *string `json:"UIN,omitnil,omitempty" name:"UIN"`

	// 用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type ScanTaskInfoList struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// cron格式
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 创建时间
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 自定义指定扫描资产信息
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// 预估时间
	PredictTime *int64 `json:"PredictTime,omitnil,omitempty" name:"PredictTime"`

	// 预估完成时间
	PredictEndTime *string `json:"PredictEndTime,omitnil,omitempty" name:"PredictEndTime"`

	// 报告数量
	ReportNumber *int64 `json:"ReportNumber,omitnil,omitempty" name:"ReportNumber"`

	// 资产数量
	AssetNumber *int64 `json:"AssetNumber,omitnil,omitempty" name:"AssetNumber"`

	// 扫描状态, 0-初始值，1-正在扫描，2-扫描完成，3-扫描出错，4-停止扫描
	ScanStatus *int64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// 任务进度
	Percent *float64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// port/poc/weakpass/webcontent/configrisk
	ScanItem *string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// 0-全扫，1-指定资产扫，2-排除资产扫，3-自定义指定资产扫描
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// vss子任务id
	VSSTaskId *string `json:"VSSTaskId,omitnil,omitempty" name:"VSSTaskId"`

	// cspm子任务id
	CSPMTaskId *string `json:"CSPMTaskId,omitnil,omitempty" name:"CSPMTaskId"`

	// 主机漏扫子任务id
	CWPPOCId *string `json:"CWPPOCId,omitnil,omitempty" name:"CWPPOCId"`

	// 主机基线子任务id
	CWPBlId *string `json:"CWPBlId,omitnil,omitempty" name:"CWPBlId"`

	// vss子任务进度
	VSSTaskProcess *int64 `json:"VSSTaskProcess,omitnil,omitempty" name:"VSSTaskProcess"`

	// cspm子任务进度
	CSPMTaskProcess *uint64 `json:"CSPMTaskProcess,omitnil,omitempty" name:"CSPMTaskProcess"`

	// 主机漏扫子任务进度
	CWPPOCProcess *int64 `json:"CWPPOCProcess,omitnil,omitempty" name:"CWPPOCProcess"`

	// 主机基线子任务进度
	CWPBlProcess *uint64 `json:"CWPBlProcess,omitnil,omitempty" name:"CWPBlProcess"`

	// 异常状态码
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 异常信息
	ErrorInfo *string `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// 周期任务开始的天数
	StartDay *int64 `json:"StartDay,omitnil,omitempty" name:"StartDay"`

	// 扫描频率,单位天,1-每天,7-每周,30-月,0-扫描一次
	Frequency *int64 `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// 完成次数
	CompleteNumber *int64 `json:"CompleteNumber,omitnil,omitempty" name:"CompleteNumber"`

	// 已完成资产个数
	CompleteAssetNumber *int64 `json:"CompleteAssetNumber,omitnil,omitempty" name:"CompleteAssetNumber"`

	// 风险数
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 资产
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// 用户Appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户主账户ID
	UIN *string `json:"UIN,omitnil,omitempty" name:"UIN"`

	// 用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// 扫描来源
	ScanFrom *string `json:"ScanFrom,omitnil,omitempty" name:"ScanFrom"`

	// 是否限免体检0不是，1是
	IsFree *int64 `json:"IsFree,omitnil,omitempty" name:"IsFree"`

	// 是否可以删除，1-可以，0-不可以，对应多账户管理使用
	IsDelete *int64 `json:"IsDelete,omitnil,omitempty" name:"IsDelete"`

	// 任务源类型，0-默认，1-小助手，2-体检项
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type ServerRisk struct {
	// 测绘标签
	ServiceTag *string `json:"ServiceTag,omitnil,omitempty" name:"ServiceTag"`

	// 端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 风险等级 low:低危 high:高危 middle:中危 info:提示 extreme:严重
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 服务
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 风险详情
	RiskDetails *string `json:"RiskDetails,omitnil,omitempty" name:"RiskDetails"`

	// 处置建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 状态，0未处理、1已处置、2已忽略、3云防已防护
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 资产唯一id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 服务快照
	ServiceSnapshot *string `json:"ServiceSnapshot,omitnil,omitempty" name:"ServiceSnapshot"`

	// 服务访问的url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 列表索引值
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 风险列表
	RiskList []*ServerRiskSuggestion `json:"RiskList,omitnil,omitempty" name:"RiskList"`

	// 建议列表
	SuggestionList []*ServerRiskSuggestion `json:"SuggestionList,omitnil,omitempty" name:"SuggestionList"`

	// HTTP响应状态码
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 新风险等级,high_risk 高危 suspect 疑似 Normal 暂无风险
	NewLevel *string `json:"NewLevel,omitnil,omitempty" name:"NewLevel"`

	// 状态，0未处理、1已处置、2已忽略、3云防已防护、4无需处理
	XspmStatus *uint64 `json:"XspmStatus,omitnil,omitempty" name:"XspmStatus"`
}

type ServerRiskSuggestion struct {
	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 详情
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`
}

type ServiceSupport struct {
	// 产品名称:
	// "cfw_waf_virtual", "cwp_detect", "cwp_defense", "cwp_fix"
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 已处理的资产总数
	SupportHandledCount *int64 `json:"SupportHandledCount,omitnil,omitempty" name:"SupportHandledCount"`

	// 支持的资产总数
	SupportTotalCount *int64 `json:"SupportTotalCount,omitnil,omitempty" name:"SupportTotalCount"`

	// 是否支持该产品1支持；0不支持
	IsSupport *bool `json:"IsSupport,omitnil,omitempty" name:"IsSupport"`
}

type StatisticalFilter struct {
	// 0:不基于统计检测
	// 1:发生次数高于固定值
	// 2:发生次数高于周期平均值的百分之
	// 3:发生次数高于用户平均值的百分之
	OperatorType *int64 `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`

	// 统计值
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type StopRiskCenterTaskRequestParams struct {
	// 任务id 列表
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type StopRiskCenterTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id 列表
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *StopRiskCenterTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRiskCenterTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRiskCenterTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRiskCenterTaskResponseParams struct {
	// Status为0， 停止成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopRiskCenterTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopRiskCenterTaskResponseParams `json:"Response"`
}

func (r *StopRiskCenterTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRiskCenterTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubUserInfo struct {
	// 主键ID，无业务意义仅作为唯一键
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 子账号Appid
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 子账号UIn
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号名称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 主账号Appid
	OwnerAppID *string `json:"OwnerAppID,omitnil,omitempty" name:"OwnerAppID"`

	// 主账号Uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 主账号名称
	OwnerNickName *string `json:"OwnerNickName,omitnil,omitempty" name:"OwnerNickName"`

	// 所属主账号memberId信息
	OwnerMemberID *string `json:"OwnerMemberID,omitnil,omitempty" name:"OwnerMemberID"`

	// 账户类型，0为腾讯云账户，1为AWS账户
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 可访问服务数量
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// 可访问接口数量
	InterfaceCount *int64 `json:"InterfaceCount,omitnil,omitempty" name:"InterfaceCount"`

	// 可访问资源数量
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// 访问/行为日志数量
	LogCount *int64 `json:"LogCount,omitnil,omitempty" name:"LogCount"`

	// 权限配置风险
	ConfigRiskCount *int64 `json:"ConfigRiskCount,omitnil,omitempty" name:"ConfigRiskCount"`

	// 危险行为告警
	ActionRiskCount *int64 `json:"ActionRiskCount,omitnil,omitempty" name:"ActionRiskCount"`

	// 是否接入云审计日志
	IsAccessCloudAudit *bool `json:"IsAccessCloudAudit,omitnil,omitempty" name:"IsAccessCloudAudit"`

	// 是否配置风险的安全体检
	IsAccessCheck *bool `json:"IsAccessCheck,omitnil,omitempty" name:"IsAccessCheck"`

	// 是否配置用户行为管理策略
	IsAccessUeba *bool `json:"IsAccessUeba,omitnil,omitempty" name:"IsAccessUeba"`
}

type SubnetAsset struct {
	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 资产ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络名
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// cidr
	CIDR *string `json:"CIDR,omitnil,omitempty" name:"CIDR"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// cvm数
	CVM *int64 `json:"CVM,omitnil,omitempty" name:"CVM"`

	// 可用ip数
	AvailableIp *int64 `json:"AvailableIp,omitnil,omitempty" name:"AvailableIp"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 配置风险
	ConfigureRisk *int64 `json:"ConfigureRisk,omitnil,omitempty" name:"ConfigureRisk"`

	// 任务数
	ScanTask *int64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// 最后扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 是否核心
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// 是否新资产 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type Tag struct {
	// 标签名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签内容
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TagCount struct {
	// 产品名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志条数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type Tags struct {
	// 主机标签key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 主机标签value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskAdvanceCFG struct {
	// 端口风险高级配置
	PortRisk []*PortRiskAdvanceCFGParamItem `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// 漏洞风险高级配置
	VulRisk []*TaskCenterVulRiskInputParam `json:"VulRisk,omitnil,omitempty" name:"VulRisk"`

	// 弱口令风险高级配置
	WeakPwdRisk []*TaskCenterWeakPwdRiskInputParam `json:"WeakPwdRisk,omitnil,omitempty" name:"WeakPwdRisk"`

	// 配置风险高级配置
	CFGRisk []*TaskCenterCFGRiskInputParam `json:"CFGRisk,omitnil,omitempty" name:"CFGRisk"`
}

type TaskAssetObject struct {
	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 资产分类
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// ip/域名/资产id，数据库id等
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 多云资产唯一id
	Arn *string `json:"Arn,omitnil,omitempty" name:"Arn"`
}

type TaskCenterCFGRiskInputParam struct {
	// 检测项ID
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// 是否开启，0-不开启，1-开启
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type TaskCenterVulRiskInputParam struct {
	// 风险ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 是否开启，0-不开启，1-开启
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type TaskCenterWeakPwdRiskInputParam struct {
	// 检测项ID
	CheckItemId *int64 `json:"CheckItemId,omitnil,omitempty" name:"CheckItemId"`

	// 是否开启，0-不开启，1-开启
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type TaskIdListKey struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type TaskLogInfo struct {
	// 报告名称
	TaskLogName *string `json:"TaskLogName,omitnil,omitempty" name:"TaskLogName"`

	// 报告ID
	TaskLogId *string `json:"TaskLogId,omitnil,omitempty" name:"TaskLogId"`

	// 关联资产个数
	AssetsNumber *int64 `json:"AssetsNumber,omitnil,omitempty" name:"AssetsNumber"`

	// 安全风险数量
	RiskNumber *int64 `json:"RiskNumber,omitnil,omitempty" name:"RiskNumber"`

	// 报告生成时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 任务状态码：0 初始值  1正在扫描  2扫描完成  3扫描出错，4停止，5暂停，6该任务已被重启过
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 关联任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 扫描开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务中心扫描任务ID
	TaskCenterTaskId *string `json:"TaskCenterTaskId,omitnil,omitempty" name:"TaskCenterTaskId"`

	// 租户ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账户ID
	UIN *string `json:"UIN,omitnil,omitempty" name:"UIN"`

	// 用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 报告类型： 1安全体检 2日报 3周报 4月报
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 报告模板id
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type TaskLogURL struct {
	// 报告下载临时链接
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 任务报告id
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 任务报告名称
	TaskLogName *string `json:"TaskLogName,omitnil,omitempty" name:"TaskLogName"`

	// APP ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type UebaCustomRule struct {
	// 策略名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 1: 云账号
	// 2: 自定义用户
	UserType *int64 `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 发生时间
	// 1：10分钟
	// 2：1小时
	// 3：一天
	// 4：一周
	// 5：一个月
	TimeInterval *int64 `json:"TimeInterval,omitnil,omitempty" name:"TimeInterval"`

	// 发生事件
	EventContent *UebaEventContent `json:"EventContent,omitnil,omitempty" name:"EventContent"`

	// 告警名称
	AlertName *string `json:"AlertName,omitnil,omitempty" name:"AlertName"`

	// 告警类型
	// 0:  提示
	// 1:  低危
	// 2:  中危
	// 3:  高危
	// 4:  严重
	AlterLevel *int64 `json:"AlterLevel,omitnil,omitempty" name:"AlterLevel"`

	// 操作者
	Operator []*string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 操作对象
	OperateObject []*string `json:"OperateObject,omitnil,omitempty" name:"OperateObject"`

	// 操作方式
	OperateMethod []*string `json:"OperateMethod,omitnil,omitempty" name:"OperateMethod"`

	// 日志类型
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 日志中文名
	LogTypeStr *string `json:"LogTypeStr,omitnil,omitempty" name:"LogTypeStr"`
}

type UebaEventContent struct {
	// 发生事件类型
	// 1:语句检索
	// 2:过滤检索
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 语句检索内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 检索条件
	Filters []*WhereFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 统计条件
	StatisticalFilter *StatisticalFilter `json:"StatisticalFilter,omitnil,omitempty" name:"StatisticalFilter"`
}

type UebaRule struct {
	// 策略id
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 策略类型
	// 0:系统策略
	// 1:自定义策略
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 策略等级
	// 0:提示
	// 1:低危
	// 2:中危
	// 3:高危
	// 4:严重
	RuleLevel *int64 `json:"RuleLevel,omitnil,omitempty" name:"RuleLevel"`

	// 策略内容
	RuleContent *string `json:"RuleContent,omitnil,omitempty" name:"RuleContent"`

	// 策略开关
	RuleStatus *bool `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// 命中次数
	HitCount *uint64 `json:"HitCount,omitnil,omitempty" name:"HitCount"`

	// 所属账号Appid
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 多账号，成员ID
	MemberID *string `json:"MemberID,omitnil,omitempty" name:"MemberID"`

	// Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 自定义规则具体内容
	CustomRuleDetail *UebaCustomRule `json:"CustomRuleDetail,omitnil,omitempty" name:"CustomRuleDetail"`

	// 云类型
	// 腾讯云：0
	// aws：1
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

// Predefined struct for user
type UpdateAlertStatusListRequestParams struct {
	// 告警ID列表
	ID []*NewAlertKey `json:"ID,omitnil,omitempty" name:"ID"`

	// 操作类型 
	// 1:撤销处置 
	// 2:标记为已处置 
	// 3:标记忽略 
	// 4:取消标记处置
	// 5:取消标记忽略
	OperateType *int64 `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 被调用的集团账号的成员id
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`
}

type UpdateAlertStatusListRequest struct {
	*tchttp.BaseRequest
	
	// 告警ID列表
	ID []*NewAlertKey `json:"ID,omitnil,omitempty" name:"ID"`

	// 操作类型 
	// 1:撤销处置 
	// 2:标记为已处置 
	// 3:标记忽略 
	// 4:取消标记处置
	// 5:取消标记忽略
	OperateType *int64 `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 被调用的集团账号的成员id
	OperatedMemberId []*string `json:"OperatedMemberId,omitnil,omitempty" name:"OperatedMemberId"`
}

func (r *UpdateAlertStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "OperateType")
	delete(f, "MemberId")
	delete(f, "OperatedMemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAlertStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlertStatusListResponseParams struct {
	// 结果信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 结果代码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAlertStatusListResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAlertStatusListResponseParams `json:"Response"`
}

func (r *UpdateAlertStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VULBaseInfo struct {
	// 风险等级
	// high 高危/ middle 中危 / low 低危 /info 提示
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 漏洞发布时间
	PublishTime *string `json:"PublishTime,omitnil,omitempty" name:"PublishTime"`

	// 最近扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 影响资产数量
	AffectAssetCount *uint64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// 风险ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 漏洞类型
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// 漏洞名
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// 描述
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 漏洞payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 漏洞影响组件
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 技术参考
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// 漏洞影响版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 风险点
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 修复建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞
	// 注意：此字段可能返回 null，表示取不到有效值。
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`

	// CVSS评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSS *float64 `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// 攻击热度
	// 0/1/2/3 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackHeat *int64 `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// 检测状态 0 未扫描 1扫描中 2 扫描完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *int64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// 1/0是否必修
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSuggest *int64 `json:"IsSuggest,omitnil,omitempty" name:"IsSuggest"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulTag []*string `json:"VulTag,omitnil,omitempty" name:"VulTag"`

	// 支持产品 逗号分隔  "cfw_waf_virtual", "cwp_detect", "cwp_defense", "cwp_fix"
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportProduct *string `json:"SupportProduct,omitnil,omitempty" name:"SupportProduct"`

	// 漏洞检测任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 漏洞id 旧版
	// 注意：此字段可能返回 null，表示取不到有效值。
	PcmgrID *string `json:"PcmgrID,omitnil,omitempty" name:"PcmgrID"`

	// 漏洞id  新版
	// 注意：此字段可能返回 null，表示取不到有效值。
	TvdID *string `json:"TvdID,omitnil,omitempty" name:"TvdID"`
}

type VULRiskAdvanceCFGList struct {
	// 风险ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 漏洞名称
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 识别来源
	CheckFrom *string `json:"CheckFrom,omitnil,omitempty" name:"CheckFrom"`

	// 是否启用，1-启用，0-禁用
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 风险类型
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// 影响版本
	ImpactVersion *string `json:"ImpactVersion,omitnil,omitempty" name:"ImpactVersion"`

	// CVE
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// 漏洞标签
	VULTag []*string `json:"VULTag,omitnil,omitempty" name:"VULTag"`

	// 修复方式
	FixMethod []*string `json:"FixMethod,omitnil,omitempty" name:"FixMethod"`

	// 披露时间
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`

	// 漏洞描述
	VULDescribe *string `json:"VULDescribe,omitnil,omitempty" name:"VULDescribe"`

	// 影响组件
	ImpactComponent *string `json:"ImpactComponent,omitnil,omitempty" name:"ImpactComponent"`

	// 漏洞Payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 技术参考
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// cvss评分
	CVSS *string `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// 攻击热度
	AttackHeat *string `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// 安全产品支持情况
	ServiceSupport []*ServiceSupport `json:"ServiceSupport,omitnil,omitempty" name:"ServiceSupport"`

	// 最新检测时间
	RecentScanTime *string `json:"RecentScanTime,omitnil,omitempty" name:"RecentScanTime"`
}

type VULRiskInfo struct {
	// 修复建议
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// 技术参考/参考链接
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// 漏洞描述
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 受影响组件
	ImpactComponent []*VulImpactComponentInfo `json:"ImpactComponent,omitnil,omitempty" name:"ImpactComponent"`
}

type VULViewVULRisk struct {
	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 影响资产
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 影响资产数量
	AffectAssetCount *uint64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// 风险ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 扫描来源，具体看接口返回枚举类型
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 前端索引
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 漏洞类型
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// 漏洞名
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// 描述
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 漏洞payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 漏洞影响组件
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 技术参考
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// 漏洞影响版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 风险点
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 修复建议
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`
}

type VULViewVULRiskData struct {
	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 影响资产
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 影响资产数量
	AffectAssetCount *uint64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// 风险ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 扫描来源，具体看接口返回枚举类型
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 前端索引
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 漏洞类型
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// 漏洞名
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// 漏洞payload
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 漏洞影响组件
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 漏洞影响版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 风险点
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`

	// CVSS评分
	CVSS *float64 `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// PCMGRId
	PCMGRId *string `json:"PCMGRId,omitnil,omitempty" name:"PCMGRId"`

	// 漏洞标签。搜索时应急 必修传参VulTag=SuggestRepair/EMGCVul
	VulTag []*string `json:"VulTag,omitnil,omitempty" name:"VulTag"`

	// 漏洞披露时间
	DisclosureTime *string `json:"DisclosureTime,omitnil,omitempty" name:"DisclosureTime"`

	// 攻击热度
	AttackHeat *uint64 `json:"AttackHeat,omitnil,omitempty" name:"AttackHeat"`

	// 是否必修漏洞，1-是，0-不是
	IsSuggest *int64 `json:"IsSuggest,omitnil,omitempty" name:"IsSuggest"`

	// 处置任务id
	HandleTaskId *string `json:"HandleTaskId,omitnil,omitempty" name:"HandleTaskId"`

	// 引擎来源
	EngineSource *string `json:"EngineSource,omitnil,omitempty" name:"EngineSource"`

	// 新的漏洞风险id
	VulRiskId *string `json:"VulRiskId,omitnil,omitempty" name:"VulRiskId"`

	// 新版漏洞id
	TvdID *string `json:"TvdID,omitnil,omitempty" name:"TvdID"`

	// 是否可以一键体检，1-可以，0-不可以
	IsOneClick *uint64 `json:"IsOneClick,omitnil,omitempty" name:"IsOneClick"`
}

type Vpc struct {
	// 子网(只支持32位)
	Subnet *uint64 `json:"Subnet,omitnil,omitempty" name:"Subnet"`

	// 互通vpc(只支持32位)
	ConnectedVpc *uint64 `json:"ConnectedVpc,omitnil,omitempty" name:"ConnectedVpc"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// region区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 云服务器(只支持32位)
	CVM *uint64 `json:"CVM,omitnil,omitempty" name:"CVM"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// dns域名
	DNS []*string `json:"DNS,omitnil,omitempty" name:"DNS"`

	// 资产名称
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// cidr网段
	CIDR *string `json:"CIDR,omitnil,omitempty" name:"CIDR"`

	// 资产创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 是否新资产 1新
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// 是否核心资产1是 2不是
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`
}

type VulImpactComponentInfo struct {
	// 组件名称
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 版本名称
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type VulRiskItem struct {
	// 云账号ID
	CloudAccountID *string `json:"CloudAccountID,omitnil,omitempty" name:"CloudAccountID"`

	// 实例ID
	AssetID *string `json:"AssetID,omitnil,omitempty" name:"AssetID"`

	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 漏洞名称
	VulName *string `json:"VulName,omitnil,omitempty" name:"VulName"`

	// 漏洞类型
	VulCategory *string `json:"VulCategory,omitnil,omitempty" name:"VulCategory"`

	// 漏洞等级
	VulLevel *string `json:"VulLevel,omitnil,omitempty" name:"VulLevel"`

	// CVE编号
	CveID *string `json:"CveID,omitnil,omitempty" name:"CveID"`

	// 漏洞描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 容器ID
	ContainerID *string `json:"ContainerID,omitnil,omitempty" name:"ContainerID"`

	// 漏洞风险修复建议
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// Linux漏洞
	VulCategoryName *string `json:"VulCategoryName,omitnil,omitempty" name:"VulCategoryName"`

	// 漏洞等级名称
	VulLevelName *string `json:"VulLevelName,omitnil,omitempty" name:"VulLevelName"`

	// 实例状态中文信息
	InstanceStatusName *string `json:"InstanceStatusName,omitnil,omitempty" name:"InstanceStatusName"`

	// 租户ID
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`
}

type VulTrend struct {
	// 影响的资产数
	AffectAssetCount *int64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// 影响的用户数
	AffectUserCount *int64 `json:"AffectUserCount,omitnil,omitempty" name:"AffectUserCount"`

	// 攻击数
	AttackCount *int64 `json:"AttackCount,omitnil,omitempty" name:"AttackCount"`

	// 时间
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type WebsiteRisk struct {
	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// 风险等级，low-低危，high-高危，middle-中危，info-提示，extreme-严重。
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 状态，0未处理、1已处置、2已忽略
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// ID,处理风险使用
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 前端索引
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 风险链接
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 风险文件地址
	URLPath *string `json:"URLPath,omitnil,omitempty" name:"URLPath"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 类型
	DetectEngine *string `json:"DetectEngine,omitnil,omitempty" name:"DetectEngine"`

	// 结果描述
	ResultDescribe *string `json:"ResultDescribe,omitnil,omitempty" name:"ResultDescribe"`

	// 源地址url
	SourceURL *string `json:"SourceURL,omitnil,omitempty" name:"SourceURL"`

	// 源文件地址
	SourceURLPath *string `json:"SourceURLPath,omitnil,omitempty" name:"SourceURLPath"`
}

type WhereFilter struct {
	// 过滤的项
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤的值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 中台定义：
	// 1等于 2大于 3小于 4大于等于 5小于等于 6不等于 9模糊匹配 13非模糊匹配 14按位与
	// 精确匹配填 7 模糊匹配填9 
	OperatorType *int64 `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`
}