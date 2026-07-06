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

package v20221121

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AIAgentAsset struct {
	// <p>ID 标识</p>
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>agent 名称</p>
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// <p>agent 使用模型名称</p>
	AgentModel []*string `json:"AgentModel,omitnil,omitempty" name:"AgentModel"`

	// <p>实例 ID</p>
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// <p>实例名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>metadata 风险列表。有如下枚举值: 1. AK_TMP  2. USER_DATA</p>
	MetadataRiskList []*string `json:"MetadataRiskList,omitnil,omitempty" name:"MetadataRiskList"`

	// <p>首次检出时间</p>
	IdentityTimeFirst *string `json:"IdentityTimeFirst,omitnil,omitempty" name:"IdentityTimeFirst"`

	// <p>最近检出时间</p>
	IdentityTimeLast *string `json:"IdentityTimeLast,omitnil,omitempty" name:"IdentityTimeLast"`

	// <p>检出方式。有如下枚举值 1. FINGER 资产指纹方式检出 2. NETWORK 网络访问方式检出</p>
	IdentityMethod *string `json:"IdentityMethod,omitnil,omitempty" name:"IdentityMethod"`

	// <p>暴露状态。有如下枚举值。1. EXPOSED；2.UNEXPOSED；</p><ol start="3"><li>UNKNOWN;</li></ol>
	ExposureStatus *string `json:"ExposureStatus,omitnil,omitempty" name:"ExposureStatus"`

	// <p>metadata 有风险时对应路径</p>
	MetadataRiskURL *string `json:"MetadataRiskURL,omitnil,omitempty" name:"MetadataRiskURL"`

	// <p>无</p>
	SkillState *SkillState `json:"SkillState,omitnil,omitempty" name:"SkillState"`

	// <p>流量沙箱插件状态</p>
	TrafficPluginState *TrafficPluginState `json:"TrafficPluginState,omitnil,omitempty" name:"TrafficPluginState"`

	// <p>流量沙箱规则状态</p>
	TrafficRuleState []*TrafficRuleState `json:"TrafficRuleState,omitnil,omitempty" name:"TrafficRuleState"`

	// <p>命令沙箱插件状态</p>
	CommandPluginState *CommandPluginState `json:"CommandPluginState,omitnil,omitempty" name:"CommandPluginState"`
}

type AKInfo struct {
	// ak对应id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// ak具体值 临时密钥时返回临时密钥
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 所属账号
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type AccessCredentialOutput struct {
	// 凭据键名（原文），如SecretId、SecretKey、Token等
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 凭据键值（打码后）
	// 补充说明：保留前3后4位，中间用***替代；长度不足7位时全部替换为***
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AccessKeyAlarm struct {
	// <p>告警名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>告警等级<br>0-无效 1-提示 2-低危 3-中危 4-高危 5-严重</p>
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>告警记录ID</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>告警规则ID</p>
	AlarmRuleID *int64 `json:"AlarmRuleID,omitnil,omitempty" name:"AlarmRuleID"`

	// <p>告警类型<br>0 异常调用<br>1 泄漏监测</p>
	AlarmType *int64 `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// <p>访问密钥</p>
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// <p>访问密钥ID</p>
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// <p>访问密钥备注</p>
	AccessKeyRemark *string `json:"AccessKeyRemark,omitnil,omitempty" name:"AccessKeyRemark"`

	// <p>最后告警时间</p>
	LastAlarmTime *string `json:"LastAlarmTime,omitnil,omitempty" name:"LastAlarmTime"`

	// <p>告警状态<br>0-未处理 1-已处理 2-已忽略</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>聚合日期</p>
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// <p>告警标签</p>
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// <p>所属主账号Uin</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>所属主账号昵称</p>
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// <p>所属子账号Uin</p>
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// <p>所属子账号昵称</p>
	SubNickname *string `json:"SubNickname,omitnil,omitempty" name:"SubNickname"`

	// <p>账号类型<br>0 主账号AK 1 子账号AK 2 临时密钥</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>所属appid</p>
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>泄漏证据</p>
	LeakEvidence []*string `json:"LeakEvidence,omitnil,omitempty" name:"LeakEvidence"`

	// <p>是否支持编辑信任账号</p>
	IsSupportEditWhiteAccount *bool `json:"IsSupportEditWhiteAccount,omitnil,omitempty" name:"IsSupportEditWhiteAccount"`

	// <p>告警证据</p>
	Evidence *string `json:"Evidence,omitnil,omitempty" name:"Evidence"`

	// <p>告警规则标识</p>
	RuleKey *string `json:"RuleKey,omitnil,omitempty" name:"RuleKey"`

	// <p>云厂商类型 0:腾讯云 1:亚马逊云 2:微软云 3:谷歌云 4:阿里云 5:华为云</p>
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// <p>告警AI分析状态<br>-1 分析失败<br>0 未分析<br>1 分析中<br>2 分析成功，真实告警<br>3 分析成功，可疑告警</p>
	AIStatus *int64 `json:"AIStatus,omitnil,omitempty" name:"AIStatus"`

	// <p>首次告警时间戳（秒级）</p>
	FirstAlarmTimestamp *int64 `json:"FirstAlarmTimestamp,omitnil,omitempty" name:"FirstAlarmTimestamp"`

	// <p>最后告警时间戳（秒级）</p>
	LastAlarmTimestamp *int64 `json:"LastAlarmTimestamp,omitnil,omitempty" name:"LastAlarmTimestamp"`

	// <p>ai分析失败描述，未失败为空字符串</p>
	AIFailedReason *string `json:"AIFailedReason,omitnil,omitempty" name:"AIFailedReason"`
}

type AccessKeyAlarmCount struct {
	// 访问密钥的ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 访问密钥
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 告警数量
	AlarmCount *int64 `json:"AlarmCount,omitnil,omitempty" name:"AlarmCount"`

	// 访问密钥状态 0 禁用 1 已启用 2 已删除
	AccessKeyStatus *int64 `json:"AccessKeyStatus,omitnil,omitempty" name:"AccessKeyStatus"`

	// AK创建时间
	AccessKeyCreateTime *string `json:"AccessKeyCreateTime,omitnil,omitempty" name:"AccessKeyCreateTime"`

	// AK最后使用时间，从未使用过则返回“-”
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`
}

type AccessKeyAlarmInfo struct {
	// 告警类型/风险类型
	// 告警类型：
	// 0异常调用
	// 1泄漏检测
	// 2自定义
	// 
	// 风险类型：
	// 0：配置风险
	// 1: 自定义风险
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 告警数量/风险数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type AccessKeyAsset struct {
	// AK 的id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// AK名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 账号所属APPID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 所属主账号Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 主账号昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 所属子账号Uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 所属子账号昵称
	SubNickname *string `json:"SubNickname,omitnil,omitempty" name:"SubNickname"`

	// 0 主账号AK
	// 1 子账号AK
	// 2 临时密钥
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 安全建议 枚举
	// 0 正常
	// 1 立即处理
	// 2 建议加固
	Advice *int64 `json:"Advice,omitnil,omitempty" name:"Advice"`

	// 告警信息列表
	AccessKeyAlarmList []*AccessKeyAlarmInfo `json:"AccessKeyAlarmList,omitnil,omitempty" name:"AccessKeyAlarmList"`

	// 风险信息列表
	AccessKeyRiskList []*AccessKeyAlarmInfo `json:"AccessKeyRiskList,omitnil,omitempty" name:"AccessKeyRiskList"`

	// 源IP数量
	IPCount *int64 `json:"IPCount,omitnil,omitempty" name:"IPCount"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// AK状态 
	// 0:禁用
	// 1:已启用
	// 2:已删除(已在cam侧删除，安全中心仍然存留之前的记录)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 0 表示已检测
	// 1 表示检测中
	CheckStatus *int64 `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// 云厂商类型 0:腾讯云 1:亚马逊云 2:微软云 3:谷歌云 4:阿里云 5:华为云
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

type AccessKeyRisk struct {
	// 风险名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 风险等级
	// 0-无效 1-提示 2-低危 3-中危 4-高危 5-严重
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 风险记录ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 风险规则ID
	RiskRuleID *int64 `json:"RiskRuleID,omitnil,omitempty" name:"RiskRuleID"`

	// 风险类型
	// 0-配置风险
	RiskType *int64 `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// 访问密钥
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 访问密钥ID
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// 访问密钥备注
	AccessKeyRemark *string `json:"AccessKeyRemark,omitnil,omitempty" name:"AccessKeyRemark"`

	// 风险检出时间
	RiskTime *string `json:"RiskTime,omitnil,omitempty" name:"RiskTime"`

	// 风险状态
	// 0-未处理 2-已忽略 3-已收敛
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 风险标签
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 风险证据
	Evidence *string `json:"Evidence,omitnil,omitempty" name:"Evidence"`

	// 风险描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 所属主账号Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 所属主账号昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 所属子账号Uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 所属子账号昵称
	SubNickname *string `json:"SubNickname,omitnil,omitempty" name:"SubNickname"`

	// 账号类型
	// 0 主账号AK 1子账号AK
	// 2 临时密钥
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 检测状态
	// 0表示 已检测
	// 1表示 检测中
	CheckStatus *int64 `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// 所属appid
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 对应风险的查询参数
	QueryParam *string `json:"QueryParam,omitnil,omitempty" name:"QueryParam"`

	// 云类型 0-腾讯云 4-阿里云
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 相关的AK列表，包含AK名和AK备注
	RelatedAK []*AKInfo `json:"RelatedAK,omitnil,omitempty" name:"RelatedAK"`
}

type AccessKeyUser struct {
	// 账号ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 账号名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 0 主账号 1子账号
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 访问方式
	// 0 API
	// 1 控制台与API
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 安全建议 枚举 0 正常 1 立即处理 2 建议加固
	Advice *int64 `json:"Advice,omitnil,omitempty" name:"Advice"`

	// 告警信息列表
	AccessKeyAlarmList []*AccessKeyAlarmInfo `json:"AccessKeyAlarmList,omitnil,omitempty" name:"AccessKeyAlarmList"`

	// 风险信息列表
	AccessKeyRiskList []*AccessKeyAlarmInfo `json:"AccessKeyRiskList,omitnil,omitempty" name:"AccessKeyRiskList"`

	// 账号所属APPID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 主账号昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 子账号昵称
	SubNickname *string `json:"SubNickname,omitnil,omitempty" name:"SubNickname"`

	// 账号所属主账号Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 账号自身uin，主账号时与主账号uin一致
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 登录IP
	LoginIP *string `json:"LoginIP,omitnil,omitempty" name:"LoginIP"`

	// 登录地址
	LoginLocation *string `json:"LoginLocation,omitnil,omitempty" name:"LoginLocation"`

	// 登录时间
	LoginTime *string `json:"LoginTime,omitnil,omitempty" name:"LoginTime"`

	// 运营商名称
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// 操作保护是否开启
	// 0 未开启
	// 1 已开启
	ActionFlag *int64 `json:"ActionFlag,omitnil,omitempty" name:"ActionFlag"`

	// 登录保护是否开启
	// 0 未开启
	// 1 已开启
	LoginFlag *int64 `json:"LoginFlag,omitnil,omitempty" name:"LoginFlag"`

	// 0 表示已检测 1 表示检测中
	CheckStatus *int64 `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// 云厂商类型 0:腾讯云 1:亚马逊云 2:微软云 3:谷歌云 4:阿里云 5:华为云
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

// Predefined struct for user
type AddDspmAssetManagerRequestParams struct {
	// 管理员uin
	IdentifyIds []*string `json:"IdentifyIds,omitnil,omitempty" name:"IdentifyIds"`

	// 资产id
	AssetId []*string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type AddDspmAssetManagerRequest struct {
	*tchttp.BaseRequest
	
	// 管理员uin
	IdentifyIds []*string `json:"IdentifyIds,omitnil,omitempty" name:"IdentifyIds"`

	// 资产id
	AssetId []*string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *AddDspmAssetManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDspmAssetManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentifyIds")
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDspmAssetManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDspmAssetManagerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddDspmAssetManagerResponse struct {
	*tchttp.BaseResponse
	Response *AddDspmAssetManagerResponseParams `json:"Response"`
}

func (r *AddDspmAssetManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDspmAssetManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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
	// <p>租户ID</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>云厂商</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>云厂商名称</p>
	ProviderName *string `json:"ProviderName,omitnil,omitempty" name:"ProviderName"`

	// <p>云账号名称</p>
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// <p>云账号ID</p>
	CloudAccountId *string `json:"CloudAccountId,omitnil,omitempty" name:"CloudAccountId"`

	// <p>实例名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>首次发现时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>风险状态</p>
	RiskStatus *int64 `json:"RiskStatus,omitnil,omitempty" name:"RiskStatus"`

	// <p>风险名称</p>
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// <p>检查类型</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>风险等级</p>
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// <p>风险规则ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>处置分类</p>
	Classify *string `json:"Classify,omitnil,omitempty" name:"Classify"`

	// <p>等保合规</p>
	StandardTerms []*StandardTerm `json:"StandardTerms,omitnil,omitempty" name:"StandardTerms"`

	// <p>资产类型</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>资产类型图标</p>
	AssetTypeIconURL *string `json:"AssetTypeIconURL,omitnil,omitempty" name:"AssetTypeIconURL"`

	// <p>资产类型</p>
	AssetTypeName *string `json:"AssetTypeName,omitnil,omitempty" name:"AssetTypeName"`
}

type AssetTag struct {
	// 标签的key值,可以是字母、数字、下划线
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的vale值,可以是字母、数字、下划线
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type AssetTagModifyAssetItem struct {
	// <p>appid</p>
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>资产类型</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>实例ID</p>
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// <p>云厂商</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`
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

type AttackStageCount struct {
	// <p>攻击阶段</p>
	AttackStage *string `json:"AttackStage,omitnil,omitempty" name:"AttackStage"`

	// <p>策略数量</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type AttributeOptionSet struct {
	// cvm实例类型
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// cvm实例名称
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AuditLogInfo struct {
	// ai分数
	AiScore *float64 `json:"AiScore,omitnil,omitempty" name:"AiScore"`

	// 应用用户
	AppUser *string `json:"AppUser,omitnil,omitempty" name:"AppUser"`

	// 备份数据包
	BackPacket *string `json:"BackPacket,omitnil,omitempty" name:"BackPacket"`

	// 客户端 IP
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 客户端 Mac
	ClientMac *string `json:"ClientMac,omitnil,omitempty" name:"ClientMac"`

	// 终端名称，取值Proxy时为casb代理流量，其它为Agent流量
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// 客户端用户
	ClientUser *string `json:"ClientUser,omitnil,omitempty" name:"ClientUser"`

	// 客户端端口
	ClientPort *uint64 `json:"ClientPort,omitnil,omitempty" name:"ClientPort"`

	// 风险等级
	DangerLevel *uint64 `json:"DangerLevel,omitnil,omitempty" name:"DangerLevel"`

	// 数据库 IP
	DbIp *string `json:"DbIp,omitnil,omitempty" name:"DbIp"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 数据库端口
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// 数据库用户
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// 影响行数
	EffectRow *uint64 `json:"EffectRow,omitnil,omitempty" name:"EffectRow"`

	// 执行时间
	ExecTime *uint64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// 命中规则
	HitRule *string `json:"HitRule,omitnil,omitempty" name:"HitRule"`

	// 日志 ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据资产名称
	InstanceId *uint64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计单元名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 操作语句(sql 语句)
	OpSql *string `json:"OpSql,omitnil,omitempty" name:"OpSql"`

	// 操作时间(时间)
	OpTime *uint64 `json:"OpTime,omitnil,omitempty" name:"OpTime"`

	// 返回消息
	RetMsg *string `json:"RetMsg,omitnil,omitempty" name:"RetMsg"`

	// 返回码
	RetNo *uint64 `json:"RetNo,omitnil,omitempty" name:"RetNo"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 操作类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 数据资产名称
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 规则集合
	HitRules []*HitRules `json:"HitRules,omitnil,omitempty" name:"HitRules"`

	// 流量来源
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 单条审计日志id
	ReqId *string `json:"ReqId,omitnil,omitempty" name:"ReqId"`

	// SQL 主要类型，DML/DDL/DCL/TCL
	SqlMainType *string `json:"SqlMainType,omitnil,omitempty" name:"SqlMainType"`

	// 表名集合
	TableNames []*string `json:"TableNames,omitnil,omitempty" name:"TableNames"`

	// 字段名集合
	FieldNames []*string `json:"FieldNames,omitnil,omitempty" name:"FieldNames"`

	// 字段名
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 数据库类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 客户端工具
	ClientDriverName *string `json:"ClientDriverName,omitnil,omitempty" name:"ClientDriverName"`

	// 位置信息
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`

	// 字段信息（包含敏感信息）
	FieldDetails []*TableField `json:"FieldDetails,omitnil,omitempty" name:"FieldDetails"`

	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 资产所属账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type BackupLog struct {
	// 索引
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 索引开始时间
	IndexStartTime *int64 `json:"IndexStartTime,omitnil,omitempty" name:"IndexStartTime"`

	// 索引结束时间
	IndexEndTime *int64 `json:"IndexEndTime,omitnil,omitempty" name:"IndexEndTime"`

	// 备份后压缩的大小，单位M
	BackupSize *int64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// 日志状态 0备份未完成， 1备份文件，2恢复中，3已恢复，4.已删除
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 恢复剩余的分钟数，分钟，需要前端转换
	RestoreProcessRemindTime *int64 `json:"RestoreProcessRemindTime,omitnil,omitempty" name:"RestoreProcessRemindTime"`

	// 恢复日志保留的时间
	RestoreRemindTime *int64 `json:"RestoreRemindTime,omitnil,omitempty" name:"RestoreRemindTime"`

	// 恢复索引大小
	RestoreIndexSize *int64 `json:"RestoreIndexSize,omitnil,omitempty" name:"RestoreIndexSize"`

	// 恢复日志执行结束时间
	RestoreEndTime *int64 `json:"RestoreEndTime,omitnil,omitempty" name:"RestoreEndTime"`

	// 备份所属的appId
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 备份所属的资产ID
	AssetId *uint64 `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 资产所属账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
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

type CICDToken struct {
	// <p>ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>appid</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>CI/CD名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>用于接入的Token</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// <p>扫描结果存储时长</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>已扫描文件</p>
	FileCnt *uint64 `json:"FileCnt,omitnil,omitempty" name:"FileCnt"`

	// <p>最近扫描状态</p>
	LastScanStatus *string `json:"LastScanStatus,omitnil,omitempty" name:"LastScanStatus"`

	// <p>最近扫描时间</p>
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`
}

type CSIPTag struct {
	// <p>标签颜色</p>
	TagColor *string `json:"TagColor,omitnil,omitempty" name:"TagColor"`

	// <p>标签ID</p>
	TagID *uint64 `json:"TagID,omitnil,omitempty" name:"TagID"`

	// <p>标签键（根据语言环境返回中文或英文）</p>
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// <p>标签值（根据语言环境返回中文或英文）</p>
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
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

type CallRecord struct {
	// 调用记录ID
	CallID *string `json:"CallID,omitnil,omitempty" name:"CallID"`

	// 访问密钥
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 访问密钥备注
	AccessKeyRemark *string `json:"AccessKeyRemark,omitnil,omitempty" name:"AccessKeyRemark"`

	// 访问密钥ID
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// 调用源IP
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// 调用源IP备注
	SourceIPRemark *string `json:"SourceIPRemark,omitnil,omitempty" name:"SourceIPRemark"`

	// 调用源IP地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// IP类型 0:账号内（未备注） 1:账号外（未备注） 2:账号内 (已备注) 3:账号外 (已备注)
	IPType *int64 `json:"IPType,omitnil,omitempty" name:"IPType"`

	// 调用接口名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 调用产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 调用类型
	// 0:控制台调用
	// 1:API
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 用户类型CAMUser/root/AssumedRole
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 用户/角色名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 策略列表
	PolicySet []*string `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`

	// 调用次数
	CallCount *int64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// 调用错误码
	// 0表示成功
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 首次调用时间
	FirstCallTime *string `json:"FirstCallTime,omitnil,omitempty" name:"FirstCallTime"`

	// 最后调用时间
	LastCallTime *string `json:"LastCallTime,omitnil,omitempty" name:"LastCallTime"`

	// IP关联资产ID，如果为空字符串，表示没有关联
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// IP关联资产名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 聚合日期
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// appid
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 展示状态
	ShowStatus *bool `json:"ShowStatus,omitnil,omitempty" name:"ShowStatus"`

	// 运营商
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// 账号外vpc信息列表
	VpcInfo []*SourceIPVpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// 调用请求客户端列表
	ReqClient []*string `json:"ReqClient,omitnil,omitempty" name:"ReqClient"`
}

type CheckViewRiskItem struct {
	// <p>检查项规则ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>风险名称</p>
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// <p>检查类型</p>
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>风险等级</p>
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// <p>存在1个风险项</p>
	RiskDesc *string `json:"RiskDesc,omitnil,omitempty" name:"RiskDesc"`

	// <p>首次发现时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>风险更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>云厂商</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>风险状态</p>
	RiskStatus *uint64 `json:"RiskStatus,omitnil,omitempty" name:"RiskStatus"`

	// <p>受影响资产数量</p>
	AssetCount *uint64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// <p>风险数量</p>
	RiskCount *uint64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// <p>资产类型</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>事件类型</p>
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// <p>处置分类</p>
	Classify *string `json:"Classify,omitnil,omitempty" name:"Classify"`

	// <p>cspm规范条款</p>
	StandardTerms []*StandardTerm `json:"StandardTerms,omitnil,omitempty" name:"StandardTerms"`

	// <p>资产类型图标</p>
	AssetTypeIconURL *string `json:"AssetTypeIconURL,omitnil,omitempty" name:"AssetTypeIconURL"`
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

type CommandPluginState struct {
	// <p>插件安装状态（上层聚合）<br>枚举值：<br>NONE：未安装<br>INSTALLING：安装中<br>INSTALLED：已安装<br>INSTALL_FAIL：安装失败</p>
	InstallStatus *string `json:"InstallStatus,omitnil,omitempty" name:"InstallStatus"`
}

type ContainerEnvInfo struct {
	// <p>节点类型</p>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>docker版本</p>
	DockerVersion *string `json:"DockerVersion,omitnil,omitempty" name:"DockerVersion"`

	// <p>containerd版本</p>
	ContainerdVersion *string `json:"ContainerdVersion,omitnil,omitempty" name:"ContainerdVersion"`

	// <p>文件系统类型</p>
	FileSystemType *string `json:"FileSystemType,omitnil,omitempty" name:"FileSystemType"`
}

type CosAccessInfo struct {
	// 可访问账号uin
	AccessUin *string `json:"AccessUin,omitnil,omitempty" name:"AccessUin"`

	// 可访问账号uid
	AccessUid *string `json:"AccessUid,omitnil,omitempty" name:"AccessUid"`

	// 昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 身份标识 
	// 1 主账号
	// 2 子账号
	Identity *int64 `json:"Identity,omitnil,omitempty" name:"Identity"`

	// 所属主账号名称
	MainNickName *string `json:"MainNickName,omitnil,omitempty" name:"MainNickName"`

	// 可访问ak列表
	AkList []*string `json:"AkList,omitnil,omitempty" name:"AkList"`

	// 可访问权限数
	CamPolicyCount *int64 `json:"CamPolicyCount,omitnil,omitempty" name:"CamPolicyCount"`

	// 修改时间Unix时间单位毫秒
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type CosActionInfo struct {
	// 接口名
	ActionName *string `json:"ActionName,omitnil,omitempty" name:"ActionName"`

	// 接口中文名
	ActionNameCn *string `json:"ActionNameCn,omitnil,omitempty" name:"ActionNameCn"`

	// 接口描述
	ActionDescription *string `json:"ActionDescription,omitnil,omitempty" name:"ActionDescription"`
}

type CosAkAssetInfo struct {
	// appid
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// ak id
	AkId *string `json:"AkId,omitnil,omitempty" name:"AkId"`

	// ak名称
	AkName *string `json:"AkName,omitnil,omitempty" name:"AkName"`

	// ak备注
	AkRemark *string `json:"AkRemark,omitnil,omitempty" name:"AkRemark"`

	// ak所属uin
	AkOwnerUin *string `json:"AkOwnerUin,omitnil,omitempty" name:"AkOwnerUin"`

	// ak类型 1 主 2 子
	AkOwnerType *int64 `json:"AkOwnerType,omitnil,omitempty" name:"AkOwnerType"`

	// ak所属账号名
	AkOwnerName *string `json:"AkOwnerName,omitnil,omitempty" name:"AkOwnerName"`

	// ak主账号名
	AkMainOwnerName *string `json:"AkMainOwnerName,omitnil,omitempty" name:"AkMainOwnerName"`

	// ak关联桶集合
	AkRelBucketSet []*string `json:"AkRelBucketSet,omitnil,omitempty" name:"AkRelBucketSet"`

	// ak关联告警集合
	AkRelAlarmSet []*CosRiskInfo `json:"AkRelAlarmSet,omitnil,omitempty" name:"AkRelAlarmSet"`

	// Ak关联ip数
	AkRelIpCount *uint64 `json:"AkRelIpCount,omitnil,omitempty" name:"AkRelIpCount"`

	// ak状态 0 禁用 1 启用
	AkStatus *int64 `json:"AkStatus,omitnil,omitempty" name:"AkStatus"`

	// 创建时间
	CreateTimestamp *uint64 `json:"CreateTimestamp,omitnil,omitempty" name:"CreateTimestamp"`

	// 最后访问时间
	LastAccessTimestamp *uint64 `json:"LastAccessTimestamp,omitnil,omitempty" name:"LastAccessTimestamp"`
}

type CosAkSet struct {
	// ak所属appid
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// ak名称集合
	AkNameSet []*string `json:"AkNameSet,omitnil,omitempty" name:"AkNameSet"`
}

type CosAlarmInfo struct {
	// appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略类型 0-未知规则分类(Unknown), 1-异常行为(AbnormalBehavior), 2-权限过大(ExcessivePermission), 3-资源枚举(ResourceEnumerated), 4-匿名访问(AnonymousAccess)
	PolicyAbnormalType *int64 `json:"PolicyAbnormalType,omitnil,omitempty" name:"PolicyAbnormalType"`

	// 风险等级：0:Normal, 1:Tip, 2:Low, 3:Middle, 4:High, 5:Critical
	PolicyRiskLevel *int64 `json:"PolicyRiskLevel,omitnil,omitempty" name:"PolicyRiskLevel"`

	// 策略信息描述
	PolicyDescription *string `json:"PolicyDescription,omitnil,omitempty" name:"PolicyDescription"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 桶地域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 桶备注
	BucketMarker *string `json:"BucketMarker,omitnil,omitempty" name:"BucketMarker"`

	// 桶tag信息
	BucketTagInfo *string `json:"BucketTagInfo,omitnil,omitempty" name:"BucketTagInfo"`

	// 桶可访问属性
	BucketAccessWay *string `json:"BucketAccessWay,omitnil,omitempty" name:"BucketAccessWay"`

	// 所属账号uin
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 所属账号昵称
	AccountNickName *string `json:"AccountNickName,omitnil,omitempty" name:"AccountNickName"`

	// 所属账号社身份 1 主 2子
	AccountIdentify *int64 `json:"AccountIdentify,omitnil,omitempty" name:"AccountIdentify"`

	// 子账号所属主账号昵称
	AccountMainNickName *string `json:"AccountMainNickName,omitnil,omitempty" name:"AccountMainNickName"`

	// 告警时间戳Unix时间单位毫秒
	AlarmTimestamp *int64 `json:"AlarmTimestamp,omitnil,omitempty" name:"AlarmTimestamp"`

	// 处置状态 0 未处理 1 标记处置 2标记忽略
	HandleStatus *int64 `json:"HandleStatus,omitnil,omitempty" name:"HandleStatus"`

	// 告警对象id
	AlarmId *int64 `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 桶地域码值
	BucketRegionCode *string `json:"BucketRegionCode,omitnil,omitempty" name:"BucketRegionCode"`

	// 数据识别分类详情
	CategoryDetails []*CosIdentifyCategoryDetail `json:"CategoryDetails,omitnil,omitempty" name:"CategoryDetails"`
}

type CosAlarmRiskIdInfo struct {
	// 告警id
	AlarmRiskId *int64 `json:"AlarmRiskId,omitnil,omitempty" name:"AlarmRiskId"`

	// 租户id
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type CosAlarmTrendInfo struct {
	// 当前日期字符串格式
	CurrentDateStr *string `json:"CurrentDateStr,omitnil,omitempty" name:"CurrentDateStr"`

	// 当前日期总数
	CurrentDayCount *int64 `json:"CurrentDayCount,omitnil,omitempty" name:"CurrentDayCount"`

	// 当天告警分类详情
	CurrentDayOverView []*CosRiskInfo `json:"CurrentDayOverView,omitnil,omitempty" name:"CurrentDayOverView"`
}

type CosAssetDataScanDetail struct {
	// <p>识别任务状态 0:未识别 1:识别中 2:识别终止 3:识别成功 4:识别失败</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>识别进度</p>
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// <p>最近扫描时间</p>
	LatestScanTime *uint64 `json:"LatestScanTime,omitnil,omitempty" name:"LatestScanTime"`

	// <p>识别失败信息</p>
	ErrorInfo *string `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>识别结果分类详情</p>
	CategoryDetails []*CosIdentifyCategoryDetail `json:"CategoryDetails,omitnil,omitempty" name:"CategoryDetails"`
}

type CosAssetFileIdentifyInfo struct {
	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件路径
	DirName *string `json:"DirName,omitnil,omitempty" name:"DirName"`

	// 分类数据项详情
	CategoryDetails []*CosIdentifyCategoryDetail `json:"CategoryDetails,omitnil,omitempty" name:"CategoryDetails"`
}

type CosAssetInfo struct {
	// <p>appid</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>cos桶名</p>
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// <p>cos region名</p>
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// <p>地域码值</p>
	BucketRegionCode *string `json:"BucketRegionCode,omitnil,omitempty" name:"BucketRegionCode"`

	// <p>cos桶备注</p>
	BucketMarker *string `json:"BucketMarker,omitnil,omitempty" name:"BucketMarker"`

	// <p>cos桶主账号所属者</p>
	BucketOwnerUin *string `json:"BucketOwnerUin,omitnil,omitempty" name:"BucketOwnerUin"`

	// <p>cos主账号所属者昵称</p>
	BucketOwnerNickName *string `json:"BucketOwnerNickName,omitnil,omitempty" name:"BucketOwnerNickName"`

	// <p>cos桶标签详情</p>
	BucketTagInfo *string `json:"BucketTagInfo,omitnil,omitempty" name:"BucketTagInfo"`

	// <p>安全建议<br>1 暂无异常<br>2 建议加固<br>3 立即处理</p>
	BucketSecuritySuggestion *int64 `json:"BucketSecuritySuggestion,omitnil,omitempty" name:"BucketSecuritySuggestion"`

	// <p>告警列表</p>
	BucketAlarmList []*CosRiskAlarmInfo `json:"BucketAlarmList,omitnil,omitempty" name:"BucketAlarmList"`

	// <p>风险列表</p>
	BucketRiskList []*CosRiskAlarmInfo `json:"BucketRiskList,omitnil,omitempty" name:"BucketRiskList"`

	// <p>调用源ip数</p>
	BucketInvokeSourceIpCount *int64 `json:"BucketInvokeSourceIpCount,omitnil,omitempty" name:"BucketInvokeSourceIpCount"`

	// <p>访问策略</p>
	BucketAccessWay *CosBucketAccessWay `json:"BucketAccessWay,omitnil,omitempty" name:"BucketAccessWay"`

	// <p>创建时间Unix时间单位毫秒</p>
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>最后访问时间Unix时间单位毫秒</p>
	LastAccessTime *int64 `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// <p>存储桶id</p>
	BucketId *uint64 `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// <p>0 关闭<br>1 开启</p>
	MonitorStatus *uint64 `json:"MonitorStatus,omitnil,omitempty" name:"MonitorStatus"`

	// <p>数据识别扫描信息</p>
	DataScanInfo *CosAssetDataScanDetail `json:"DataScanInfo,omitnil,omitempty" name:"DataScanInfo"`

	// <p>存储桶Az类型</p><p>枚举值：</p><ul><li>MAZ： 多az</li><li>SAZ： 单az</li></ul>
	BucketAzType *string `json:"BucketAzType,omitnil,omitempty" name:"BucketAzType"`

	// <p>存储桶存储大小</p><p>默认值：0</p>
	BucketStorageSize *int64 `json:"BucketStorageSize,omitnil,omitempty" name:"BucketStorageSize"`

	// <p>存储桶对象个数</p><p>默认值：0</p>
	BucketObjectCount *int64 `json:"BucketObjectCount,omitnil,omitempty" name:"BucketObjectCount"`

	// <p>存储桶敏感识别采样率</p><p>取值范围：[0, 1]</p><p>默认值：0</p>
	IdentifySampleRate *float64 `json:"IdentifySampleRate,omitnil,omitempty" name:"IdentifySampleRate"`
}

type CosAssetSyncTaskInfo struct {
	// appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 同步任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 最后一次扫描时间
	LastScanTime *int64 `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`
}

type CosAuditPayInfo struct {
	// <p>APPID</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>订单状态 0未购买 1正常，2隔离，3销毁，6试用中，7到期</p>
	OrderStatus *uint64 `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// <p>已购对象存储数量</p>
	BucketNum *uint64 `json:"BucketNum,omitnil,omitempty" name:"BucketNum"`

	// <p>支付模式，0-后付费 1-预付费</p>
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>是否单独购买，1-单独购买，2-被其它账号共享</p>
	IsSelfBuy *uint64 `json:"IsSelfBuy,omitnil,omitempty" name:"IsSelfBuy"`

	// <p>订单开始时间</p>
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// <p>订单到期时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>0-用户未设置,1-用户设置自动续费,2-用户设置不自动续费</p>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>订单时长</p>
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>时长单位</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>资源id</p>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>公测结束时间</p>
	BetaEndTime *string `json:"BetaEndTime,omitnil,omitempty" name:"BetaEndTime"`

	// <p>系统当前时间</p>
	TimeNow *string `json:"TimeNow,omitnil,omitempty" name:"TimeNow"`

	// <p>是否分享给其它账号，1-是，2-否</p>
	IsShareToOther *uint64 `json:"IsShareToOther,omitnil,omitempty" name:"IsShareToOther"`

	// <p>uin</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>昵称</p>
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// <p>共享的bucketIdSet</p>
	BindBucket []*CosBucketId `json:"BindBucket,omitnil,omitempty" name:"BindBucket"`

	// <p>共享的appid</p>
	SharedAppIdSet []*uint64 `json:"SharedAppIdSet,omitnil,omitempty" name:"SharedAppIdSet"`

	// <p>是否已经开启后付费</p>
	PostPayStatus *uint64 `json:"PostPayStatus,omitnil,omitempty" name:"PostPayStatus"`

	// <p>0：未做过试用期试用   1 ：做过试用期试用</p>
	IsTestUser *uint64 `json:"IsTestUser,omitnil,omitempty" name:"IsTestUser"`

	// <p>剩余可用数</p>
	AvailableBucketNum *uint64 `json:"AvailableBucketNum,omitnil,omitempty" name:"AvailableBucketNum"`

	// <p>已开启的监测存储桶数</p>
	MonitorBucketNum *uint64 `json:"MonitorBucketNum,omitnil,omitempty" name:"MonitorBucketNum"`

	// <p>总的存储桶数</p>
	TotalBucketNum *uint64 `json:"TotalBucketNum,omitnil,omitempty" name:"TotalBucketNum"`

	// <p>后付费产品开关状态</p>
	PostProductStatusList []*uint64 `json:"PostProductStatusList,omitnil,omitempty" name:"PostProductStatusList"`

	// <p>后付费产品购买状态</p>
	PostProductBuyStatusList []*uint64 `json:"PostProductBuyStatusList,omitnil,omitempty" name:"PostProductBuyStatusList"`

	// <p>新后付费资源id</p>
	NewPostPayResourceId *string `json:"NewPostPayResourceId,omitnil,omitempty" name:"NewPostPayResourceId"`
}

type CosBucketAccessWay struct {
	// 可访问方式：
	// specify 指定用户
	// anonymous 可匿名访问
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 用户数
	AccessUserCount *int64 `json:"AccessUserCount,omitnil,omitempty" name:"AccessUserCount"`

	// ak数
	AccessAkCount *int64 `json:"AccessAkCount,omitnil,omitempty" name:"AccessAkCount"`

	// 角色数
	AccessRoleCount *int64 `json:"AccessRoleCount,omitnil,omitempty" name:"AccessRoleCount"`
}

type CosBucketBillingInfo struct {
	// <p>appid</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>uin</p>
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// <p>昵称</p>
	OwnerNickName *string `json:"OwnerNickName,omitnil,omitempty" name:"OwnerNickName"`

	// <p>存储桶数量</p>
	BucketCount *uint64 `json:"BucketCount,omitnil,omitempty" name:"BucketCount"`

	// <p>0 未购买 1  已单独购买 2 已被共享</p>
	BuyStatus *uint64 `json:"BuyStatus,omitnil,omitempty" name:"BuyStatus"`

	// <p>共享账号appid</p>
	ShareFromAppId *uint64 `json:"ShareFromAppId,omitnil,omitempty" name:"ShareFromAppId"`

	// <p>共享账号uin</p>
	ShareFromUin *string `json:"ShareFromUin,omitnil,omitempty" name:"ShareFromUin"`

	// <p>共享账号昵称</p>
	ShareFromNickName *string `json:"ShareFromNickName,omitnil,omitempty" name:"ShareFromNickName"`

	// <p>监控的存储桶数</p>
	MonitorBucketCount *uint64 `json:"MonitorBucketCount,omitnil,omitempty" name:"MonitorBucketCount"`

	// <p>0 关闭 1 开启</p>
	IsAutoMonitor *uint64 `json:"IsAutoMonitor,omitnil,omitempty" name:"IsAutoMonitor"`

	// <p>是否启用白名单日志功能</p>
	LogFeatureWhitelist *bool `json:"LogFeatureWhitelist,omitnil,omitempty" name:"LogFeatureWhitelist"`

	// <p>是否存在新的后付费订单</p>
	IsHaveNewPostOrder *bool `json:"IsHaveNewPostOrder,omitnil,omitempty" name:"IsHaveNewPostOrder"`

	// <p>是否存在旧后付费订单</p>
	IsHaveOldPostOrder *bool `json:"IsHaveOldPostOrder,omitnil,omitempty" name:"IsHaveOldPostOrder"`

	// <p>后付费产品列表</p>
	PostProductList []*int64 `json:"PostProductList,omitnil,omitempty" name:"PostProductList"`
}

type CosBucketId struct {
	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// bucket id集合
	BucketIdSet []*string `json:"BucketIdSet,omitnil,omitempty" name:"BucketIdSet"`
}

type CosBucketInfo struct {
	// appid信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 存储桶名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 地域信息
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 地域码值
	BucketRegionCode *string `json:"BucketRegionCode,omitnil,omitempty" name:"BucketRegionCode"`

	// 备注
	BucketMarker *string `json:"BucketMarker,omitnil,omitempty" name:"BucketMarker"`
}

type CosBucketTaskInfo struct {
	// appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 存储桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 最后一次扫描时间
	LastScanTime *int64 `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`
}

type CosDictionary struct {
	// <p>字典id</p>
	DictId *uint64 `json:"DictId,omitnil,omitempty" name:"DictId"`

	// <p>字典名称</p>
	DictName *string `json:"DictName,omitnil,omitempty" name:"DictName"`
}

type CosIdentifyCategoryDetail struct {
	// <p>分类id</p>
	CategoryId *uint64 `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>分类名称</p>
	CategoryName *string `json:"CategoryName,omitnil,omitempty" name:"CategoryName"`

	// <p>数据项集合</p>
	RuleSet []*CosIdentifyRuleDetail `json:"RuleSet,omitnil,omitempty" name:"RuleSet"`
}

type CosIdentifyRuleDetail struct {
	// 数据项id
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 数据项名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 敏感级别id
	LevelId *uint64 `json:"LevelId,omitnil,omitempty" name:"LevelId"`

	// 敏感级别名称
	LevelName *string `json:"LevelName,omitnil,omitempty" name:"LevelName"`

	// 敏感程度
	LevelScore *uint64 `json:"LevelScore,omitnil,omitempty" name:"LevelScore"`
}

type CosInvokeDetailInfo struct {
	// 调用时间
	InvokeTimestamp *int64 `json:"InvokeTimestamp,omitnil,omitempty" name:"InvokeTimestamp"`

	// 请求id
	InvokeRequestId *string `json:"InvokeRequestId,omitnil,omitempty" name:"InvokeRequestId"`

	// 调用内容
	InvokeContent *string `json:"InvokeContent,omitnil,omitempty" name:"InvokeContent"`
}

type CosInvokeIpVpcInfo struct {
	// vpc所属uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// vpc所属appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// vpcid信息
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`
}

type CosInvokeLog struct {
	// 调用时间戳
	InvokeTimestamp *uint64 `json:"InvokeTimestamp,omitnil,omitempty" name:"InvokeTimestamp"`

	// 请求id
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 请求内容 base64 json 结构
	RequestContent *string `json:"RequestContent,omitnil,omitempty" name:"RequestContent"`
}

type CosOverview struct {
	// 资产总数
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// 需要立即处理的资产数
	AlarmAssetCount *int64 `json:"AlarmAssetCount,omitnil,omitempty" name:"AlarmAssetCount"`

	// 需要加固的资产数
	RiskAssetCount *int64 `json:"RiskAssetCount,omitnil,omitempty" name:"RiskAssetCount"`

	// 告警总数
	AlarmCount *int64 `json:"AlarmCount,omitnil,omitempty" name:"AlarmCount"`

	// 当日新增告警总数
	IncrementAlarmCount *int64 `json:"IncrementAlarmCount,omitnil,omitempty" name:"IncrementAlarmCount"`

	// 风险总数
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 当日新增告警总数
	IncrementRiskCount *int64 `json:"IncrementRiskCount,omitnil,omitempty" name:"IncrementRiskCount"`

	// 风险top详情
	RiskTop []*CosRiskInfo `json:"RiskTop,omitnil,omitempty" name:"RiskTop"`

	// 告警风险top
	AlarmTop []*CosRiskInfo `json:"AlarmTop,omitnil,omitempty" name:"AlarmTop"`
}

type CosPermissionInfo struct {
	// 权限来源
	PermissionSource *string `json:"PermissionSource,omitnil,omitempty" name:"PermissionSource"`

	// 权限内容
	PermissionContent *string `json:"PermissionContent,omitnil,omitempty" name:"PermissionContent"`

	// 授权资源
	GrantResource *string `json:"GrantResource,omitnil,omitempty" name:"GrantResource"`

	// 授权动作
	GrantAction *string `json:"GrantAction,omitnil,omitempty" name:"GrantAction"`

	// 授权条件
	GrantCondition *string `json:"GrantCondition,omitnil,omitempty" name:"GrantCondition"`
}

type CosPolicyInfo struct {
	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略类型
	// PolicyType：1 告警策略 2 风险策略  3  白名单策略  4 ip隐藏策略
	PolicyType *int64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// system:系统内置 user:用户自定义
	PolicySource *int64 `json:"PolicySource,omitnil,omitempty" name:"PolicySource"`

	// 策略内容
	PolicyContent *string `json:"PolicyContent,omitnil,omitempty" name:"PolicyContent"`

	// 0 关闭
	// 1 开启
	PolicyStatus *int64 `json:"PolicyStatus,omitnil,omitempty" name:"PolicyStatus"`

	// 策略分类
	PolicyAbnormalType *int64 `json:"PolicyAbnormalType,omitnil,omitempty" name:"PolicyAbnormalType"`

	// 风险级别
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 创建时间
	PolicyCreateTime *int64 `json:"PolicyCreateTime,omitnil,omitempty" name:"PolicyCreateTime"`

	// 更新时间
	PolicyUpdateTime *int64 `json:"PolicyUpdateTime,omitnil,omitempty" name:"PolicyUpdateTime"`

	// 策略近七天命中次数
	PolicyHitCount *int64 `json:"PolicyHitCount,omitnil,omitempty" name:"PolicyHitCount"`

	// 告警内容hash
	PolicyContentHash *string `json:"PolicyContentHash,omitnil,omitempty" name:"PolicyContentHash"`

	// 关联账户数
	RelAccountCount *int64 `json:"RelAccountCount,omitnil,omitempty" name:"RelAccountCount"`

	// 关联账号uin
	RelAccountUin *string `json:"RelAccountUin,omitnil,omitempty" name:"RelAccountUin"`

	// 关联账号名
	RelAccountName *string `json:"RelAccountName,omitnil,omitempty" name:"RelAccountName"`

	// 描述信息
	PolicyDescription *string `json:"PolicyDescription,omitnil,omitempty" name:"PolicyDescription"`

	// 备注信息
	PolicyMarker *string `json:"PolicyMarker,omitnil,omitempty" name:"PolicyMarker"`

	// appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 多账号场景下的id集合
	PolicyIdSet []*int64 `json:"PolicyIdSet,omitnil,omitempty" name:"PolicyIdSet"`

	// 是否处置历史数据状态  0 无须处置 1 需要处置 2 已处置
	PolicyHistoryHandleStatus *uint64 `json:"PolicyHistoryHandleStatus,omitnil,omitempty" name:"PolicyHistoryHandleStatus"`

	// 系统策略编辑状态
	SystemPolicyEditStatus *int64 `json:"SystemPolicyEditStatus,omitnil,omitempty" name:"SystemPolicyEditStatus"`
}

type CosRiskActionInfo struct {
	// 接口名
	ActionName *string `json:"ActionName,omitnil,omitempty" name:"ActionName"`

	// 接口名中文
	ActionNameCn *string `json:"ActionNameCn,omitnil,omitempty" name:"ActionNameCn"`

	// 调用次数
	InvokeCount *int64 `json:"InvokeCount,omitnil,omitempty" name:"InvokeCount"`

	// 最后访问时间Unix时间单位毫秒
	ActionAccessTime *int64 `json:"ActionAccessTime,omitnil,omitempty" name:"ActionAccessTime"`
}

type CosRiskAlarmInfo struct {
	// 策略类型枚举值
	PolicyType *int64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略名
	PolicyTypeName *string `json:"PolicyTypeName,omitnil,omitempty" name:"PolicyTypeName"`

	// 策略类型对应的策略数量
	PolicyCount *int64 `json:"PolicyCount,omitnil,omitempty" name:"PolicyCount"`
}

type CosRiskBucketInfo struct {
	// appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 桶地域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 桶备注信息
	BucketMarker *string `json:"BucketMarker,omitnil,omitempty" name:"BucketMarker"`

	// 桶uin
	BucketUin *string `json:"BucketUin,omitnil,omitempty" name:"BucketUin"`

	// uin昵称
	BucketNickName *string `json:"BucketNickName,omitnil,omitempty" name:"BucketNickName"`

	// uin主账号昵称
	BucketMainNickName *string `json:"BucketMainNickName,omitnil,omitempty" name:"BucketMainNickName"`

	// uin身份
	BucketIdentify *int64 `json:"BucketIdentify,omitnil,omitempty" name:"BucketIdentify"`

	// 风险检出时间Unix时间单位毫秒
	LastScanTimestamp *int64 `json:"LastScanTimestamp,omitnil,omitempty" name:"LastScanTimestamp"`

	// 状态信息
	HandleStatus *int64 `json:"HandleStatus,omitnil,omitempty" name:"HandleStatus"`

	// 风险名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 风险类型
	PolicyType *int64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略级别
	PolicyLevel *int64 `json:"PolicyLevel,omitnil,omitempty" name:"PolicyLevel"`

	// 策略描述
	PolicyDescription *string `json:"PolicyDescription,omitnil,omitempty" name:"PolicyDescription"`

	// 访问方式
	BucketAccessWay *string `json:"BucketAccessWay,omitnil,omitempty" name:"BucketAccessWay"`

	// 标签信息
	BucketTagInfo *string `json:"BucketTagInfo,omitnil,omitempty" name:"BucketTagInfo"`

	// 风险id
	RiskId *int64 `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// cos地域码值
	BucketRegionCode *string `json:"BucketRegionCode,omitnil,omitempty" name:"BucketRegionCode"`

	// 是否开启自动监测状态 0 关闭 1 开启
	BucketMonitorStatus *uint64 `json:"BucketMonitorStatus,omitnil,omitempty" name:"BucketMonitorStatus"`
}

type CosRiskInfo struct {
	// 策略类型码值
	PolicyType *int64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略分类名
	PolicyTypeName *string `json:"PolicyTypeName,omitnil,omitempty" name:"PolicyTypeName"`

	// 命中策略总数
	PolicyCount *int64 `json:"PolicyCount,omitnil,omitempty" name:"PolicyCount"`
}

type CosRiskTrendInfo struct {
	// 当前日期
	CurrentDateStr *string `json:"CurrentDateStr,omitnil,omitempty" name:"CurrentDateStr"`

	// 风险数据信息
	RiskDataSet []*CosRiskInfo `json:"RiskDataSet,omitnil,omitempty" name:"RiskDataSet"`
}

type CosRiskViewInfo struct {
	// appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略id
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略分类
	PolicyType *int64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略风险等级
	PolicyRiskLevel *int64 `json:"PolicyRiskLevel,omitnil,omitempty" name:"PolicyRiskLevel"`

	// 策略描述
	PolicyDescription *string `json:"PolicyDescription,omitnil,omitempty" name:"PolicyDescription"`

	// 待处理的桶数
	HandleBucketCount *int64 `json:"HandleBucketCount,omitnil,omitempty" name:"HandleBucketCount"`

	// 最近风险检出时间Unix时间单位毫秒
	LastScanTimestamp *int64 `json:"LastScanTimestamp,omitnil,omitempty" name:"LastScanTimestamp"`
}

type CosRoleAccessInfo struct {
	// 角色ID
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色描述
	RoleDescription *string `json:"RoleDescription,omitnil,omitempty" name:"RoleDescription"`

	// 可访问权限数
	PermissionCount *uint64 `json:"PermissionCount,omitnil,omitempty" name:"PermissionCount"`

	// 策略创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type CosSourceIpInfo struct {
	// 调用UA
	UA []*string `json:"UA,omitnil,omitempty" name:"UA"`

	// 调用vpc信息
	VpcInfo *CosInvokeIpVpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`
}

// Predefined struct for user
type CreateAccessKeyCheckTaskRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 风险列表
	RiskIDList []*int64 `json:"RiskIDList,omitnil,omitempty" name:"RiskIDList"`

	// 访问密钥列表
	AccessKeyList []*string `json:"AccessKeyList,omitnil,omitempty" name:"AccessKeyList"`

	// 账号uin列表
	SubUinList []*string `json:"SubUinList,omitnil,omitempty" name:"SubUinList"`

	// 风险规则id列表
	RiskRuleIDList []*int64 `json:"RiskRuleIDList,omitnil,omitempty" name:"RiskRuleIDList"`
}

type CreateAccessKeyCheckTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 风险列表
	RiskIDList []*int64 `json:"RiskIDList,omitnil,omitempty" name:"RiskIDList"`

	// 访问密钥列表
	AccessKeyList []*string `json:"AccessKeyList,omitnil,omitempty" name:"AccessKeyList"`

	// 账号uin列表
	SubUinList []*string `json:"SubUinList,omitnil,omitempty" name:"SubUinList"`

	// 风险规则id列表
	RiskRuleIDList []*int64 `json:"RiskRuleIDList,omitnil,omitempty" name:"RiskRuleIDList"`
}

func (r *CreateAccessKeyCheckTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeyCheckTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "RiskIDList")
	delete(f, "AccessKeyList")
	delete(f, "SubUinList")
	delete(f, "RiskRuleIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessKeyCheckTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeyCheckTaskResponseParams struct {
	// 0表示成功 1表示失败
	Code *uint64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccessKeyCheckTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessKeyCheckTaskResponseParams `json:"Response"`
}

func (r *CreateAccessKeyCheckTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeyCheckTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeySyncTaskRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type CreateAccessKeySyncTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *CreateAccessKeySyncTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeySyncTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessKeySyncTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeySyncTaskResponseParams struct {
	// 发起同步任务
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 0成功 1失败
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccessKeySyncTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessKeySyncTaskResponseParams `json:"Response"`
}

func (r *CreateAccessKeySyncTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeySyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosAssetSyncTaskRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 1 同步所有 2 仅同步资产数
	SyncType *int64 `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

type CreateCosAssetSyncTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 1 同步所有 2 仅同步资产数
	SyncType *int64 `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

func (r *CreateCosAssetSyncTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosAssetSyncTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "SyncType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosAssetSyncTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosAssetSyncTaskResponseParams struct {
	// 同步任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCosAssetSyncTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosAssetSyncTaskResponseParams `json:"Response"`
}

func (r *CreateCosAssetSyncTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosAssetSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosObjectScanTaskRequestParams struct {
	// <p>1: 敏感数据识别 2:恶意文件扫描 3:批量扫描敏感数据</p>
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>存储桶列表</p>
	BucketSet []*string `json:"BucketSet,omitnil,omitempty" name:"BucketSet"`

	// <p>任务参数</p>
	TaskArgs *string `json:"TaskArgs,omitnil,omitempty" name:"TaskArgs"`

	// <p>是否全部扫描</p>
	IsScanAll *bool `json:"IsScanAll,omitnil,omitempty" name:"IsScanAll"`

	// <p>扫描时需要剔除的存储桶</p>
	DeleteBucketSet []*string `json:"DeleteBucketSet,omitnil,omitempty" name:"DeleteBucketSet"`
}

type CreateCosObjectScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>1: 敏感数据识别 2:恶意文件扫描 3:批量扫描敏感数据</p>
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>存储桶列表</p>
	BucketSet []*string `json:"BucketSet,omitnil,omitempty" name:"BucketSet"`

	// <p>任务参数</p>
	TaskArgs *string `json:"TaskArgs,omitnil,omitempty" name:"TaskArgs"`

	// <p>是否全部扫描</p>
	IsScanAll *bool `json:"IsScanAll,omitnil,omitempty" name:"IsScanAll"`

	// <p>扫描时需要剔除的存储桶</p>
	DeleteBucketSet []*string `json:"DeleteBucketSet,omitnil,omitempty" name:"DeleteBucketSet"`
}

func (r *CreateCosObjectScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosObjectScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "MemberId")
	delete(f, "BucketSet")
	delete(f, "TaskArgs")
	delete(f, "IsScanAll")
	delete(f, "DeleteBucketSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosObjectScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosObjectScanTaskResponseParams struct {
	// <p>任务id</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCosObjectScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosObjectScanTaskResponseParams `json:"Response"`
}

func (r *CreateCosObjectScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosObjectScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosPolicyRequestParams struct {
	// 策略信息
	CosPolicyInfo *CosPolicyInfo `json:"CosPolicyInfo,omitnil,omitempty" name:"CosPolicyInfo"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type CreateCosPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略信息
	CosPolicyInfo *CosPolicyInfo `json:"CosPolicyInfo,omitnil,omitempty" name:"CosPolicyInfo"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *CreateCosPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosPolicyInfo")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCosPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosPolicyResponseParams `json:"Response"`
}

func (r *CreateCosPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosRiskScanTaskRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 需要扫描的桶列表
	BucketNameSet []*CosBucketInfo `json:"BucketNameSet,omitnil,omitempty" name:"BucketNameSet"`

	// 是否扫描全部的桶
	IsScanAllBucket *bool `json:"IsScanAllBucket,omitnil,omitempty" name:"IsScanAllBucket"`
}

type CreateCosRiskScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 需要扫描的桶列表
	BucketNameSet []*CosBucketInfo `json:"BucketNameSet,omitnil,omitempty" name:"BucketNameSet"`

	// 是否扫描全部的桶
	IsScanAllBucket *bool `json:"IsScanAllBucket,omitnil,omitempty" name:"IsScanAllBucket"`
}

func (r *CreateCosRiskScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosRiskScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "BucketNameSet")
	delete(f, "IsScanAllBucket")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosRiskScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosRiskScanTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCosRiskScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosRiskScanTaskResponseParams `json:"Response"`
}

func (r *CreateCosRiskScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosRiskScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
type CreateDspmAccessExportJobRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 视图类型
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type CreateDspmAccessExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 视图类型
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *CreateDspmAccessExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmAccessExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "View")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmAccessExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmAccessExportJobResponseParams struct {
	// 任务ID
	JobID *string `json:"JobID,omitnil,omitempty" name:"JobID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmAccessExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmAccessExportJobResponseParams `json:"Response"`
}

func (r *CreateDspmAccessExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmAccessExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmApplyOrderRequestParams struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 申请类型。0-子账号授权 1-访客授权。
	ApplyType *int64 `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 权限信息。
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 主机地址。当前仅支持'%'。默认'%'。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 从审批完成后开始计算的访问权限失效时间，临时账号有效。单位毫秒。
	ValidatePeriod *int64 `json:"ValidatePeriod,omitnil,omitempty" name:"ValidatePeriod"`

	// 审批人列表。为空使用资产全部管理员。
	ApproverUin []*string `json:"ApproverUin,omitnil,omitempty" name:"ApproverUin"`

	// 申请原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 管理类型。0-普通成员 1-管理员
	ManagerType *int64 `json:"ManagerType,omitnil,omitempty" name:"ManagerType"`

	// 被授权者。子账号授权时，传目标uin，为空时默认使用当前uin；访客授权时传访客身份id。
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`
}

type CreateDspmApplyOrderRequest struct {
	*tchttp.BaseRequest
	
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 申请类型。0-子账号授权 1-访客授权。
	ApplyType *int64 `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 权限信息。
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 主机地址。当前仅支持'%'。默认'%'。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 从审批完成后开始计算的访问权限失效时间，临时账号有效。单位毫秒。
	ValidatePeriod *int64 `json:"ValidatePeriod,omitnil,omitempty" name:"ValidatePeriod"`

	// 审批人列表。为空使用资产全部管理员。
	ApproverUin []*string `json:"ApproverUin,omitnil,omitempty" name:"ApproverUin"`

	// 申请原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 管理类型。0-普通成员 1-管理员
	ManagerType *int64 `json:"ManagerType,omitnil,omitempty" name:"ManagerType"`

	// 被授权者。子账号授权时，传目标uin，为空时默认使用当前uin；访客授权时传访客身份id。
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`
}

func (r *CreateDspmApplyOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmApplyOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "ApplyType")
	delete(f, "Privilege")
	delete(f, "Host")
	delete(f, "ValidatePeriod")
	delete(f, "ApproverUin")
	delete(f, "Reason")
	delete(f, "ManagerType")
	delete(f, "Subject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmApplyOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmApplyOrderResponseParams struct {
	// 申请单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 自动审批
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoApproval *bool `json:"AutoApproval,omitnil,omitempty" name:"AutoApproval"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmApplyOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmApplyOrderResponseParams `json:"Response"`
}

func (r *CreateDspmApplyOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmApplyOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmApproveHistoryExportJobRequestParams struct {
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type CreateDspmApproveHistoryExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *CreateDspmApproveHistoryExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmApproveHistoryExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmApproveHistoryExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmApproveHistoryExportJobResponseParams struct {
	// 任务ID
	JobID *string `json:"JobID,omitnil,omitempty" name:"JobID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmApproveHistoryExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmApproveHistoryExportJobResponseParams `json:"Response"`
}

func (r *CreateDspmApproveHistoryExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmApproveHistoryExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmAssetAccessTopologyExportJobRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 视图类型。ip或instance
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type CreateDspmAssetAccessTopologyExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 视图类型。ip或instance
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *CreateDspmAssetAccessTopologyExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmAssetAccessTopologyExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "View")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmAssetAccessTopologyExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmAssetAccessTopologyExportJobResponseParams struct {
	// 任务ID
	JobID *string `json:"JobID,omitnil,omitempty" name:"JobID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmAssetAccessTopologyExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmAssetAccessTopologyExportJobResponseParams `json:"Response"`
}

func (r *CreateDspmAssetAccessTopologyExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmAssetAccessTopologyExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmAssetsExportJobRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type CreateDspmAssetsExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *CreateDspmAssetsExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmAssetsExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmAssetsExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmAssetsExportJobResponseParams struct {
	// 任务ID
	JobID *string `json:"JobID,omitnil,omitempty" name:"JobID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmAssetsExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmAssetsExportJobResponseParams `json:"Response"`
}

func (r *CreateDspmAssetsExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmAssetsExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmExportTaskRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>风险等级(0-安全,1-低风险,2-中风险,3-高风险,-1-全部)</p>
	DangerLevel *int64 `json:"DangerLevel,omitnil,omitempty" name:"DangerLevel"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库端口</p>
	DbPort *int64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// <p>数据库 IP</p>
	DbIp *string `json:"DbIp,omitnil,omitempty" name:"DbIp"`

	// <p>资产 ID</p>
	AssetsId *int64 `json:"AssetsId,omitnil,omitempty" name:"AssetsId"`

	// <p>会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>客户端 IP</p>
	ClientSideIp *string `json:"ClientSideIp,omitnil,omitempty" name:"ClientSideIp"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>命中规则</p>
	HitRule *int64 `json:"HitRule,omitnil,omitempty" name:"HitRule"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>模糊查询</p>
	FuzzySearch *string `json:"FuzzySearch,omitnil,omitempty" name:"FuzzySearch"`

	// <p>用户名</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>客户端</p>
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// <p>流量来源，取值 Agent/Proxy/空；传Agent会返回Agent的日志，传Proxy会返回Proxy日志，两都传或不传则返回所有</p>
	SourceTypes []*string `json:"SourceTypes,omitnil,omitempty" name:"SourceTypes"`

	// <p>表名，长度限制64，多个表名查询的话可以用空格连接</p>
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// <p>字段名，长度限制64，多个字段名查询的话可以用空格连接</p>
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// <p>SQL 主要类型，DDL, DML, DCL, TCL</p>
	SqlMainTypes []*string `json:"SqlMainTypes,omitnil,omitempty" name:"SqlMainTypes"`

	// <p>操作类型</p>
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// <p>影响行数最小值</p>
	RowNumMin *int64 `json:"RowNumMin,omitnil,omitempty" name:"RowNumMin"`

	// <p>影响行数最大值</p>
	RowNumMax *int64 `json:"RowNumMax,omitnil,omitempty" name:"RowNumMax"`

	// <p>数据库类型</p>
	DbTypes []*string `json:"DbTypes,omitnil,omitempty" name:"DbTypes"`

	// <p>返回码</p>
	RetNo *int64 `json:"RetNo,omitnil,omitempty" name:"RetNo"`

	// <p>客户端工具</p>
	ClientDriverName *string `json:"ClientDriverName,omitnil,omitempty" name:"ClientDriverName"`

	// <p>客户端端口</p>
	ClientPort *int64 `json:"ClientPort,omitnil,omitempty" name:"ClientPort"`

	// <p>审计日志 ID</p>
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// <p>风险等级数组(0-安全,1-低风险,2-中风险,3-高风险), 如果要全部，则需要将所有的值都传入。如果为空，则会参考：DangerLevel 入参</p>
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`

	// <p>字段分类</p>
	SensitiveCategoryRule *string `json:"SensitiveCategoryRule,omitnil,omitempty" name:"SensitiveCategoryRule"`

	// <p>字段分级</p>
	SensitiveLevelRisk *string `json:"SensitiveLevelRisk,omitnil,omitempty" name:"SensitiveLevelRisk"`

	// <p>事务Id</p>
	TrxId *int64 `json:"TrxId,omitnil,omitempty" name:"TrxId"`

	// <p>clientMac</p>
	ClientMac *string `json:"ClientMac,omitnil,omitempty" name:"ClientMac"`
}

type CreateDspmExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>风险等级(0-安全,1-低风险,2-中风险,3-高风险,-1-全部)</p>
	DangerLevel *int64 `json:"DangerLevel,omitnil,omitempty" name:"DangerLevel"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>数据库端口</p>
	DbPort *int64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// <p>数据库 IP</p>
	DbIp *string `json:"DbIp,omitnil,omitempty" name:"DbIp"`

	// <p>资产 ID</p>
	AssetsId *int64 `json:"AssetsId,omitnil,omitempty" name:"AssetsId"`

	// <p>会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>客户端 IP</p>
	ClientSideIp *string `json:"ClientSideIp,omitnil,omitempty" name:"ClientSideIp"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>命中规则</p>
	HitRule *int64 `json:"HitRule,omitnil,omitempty" name:"HitRule"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>模糊查询</p>
	FuzzySearch *string `json:"FuzzySearch,omitnil,omitempty" name:"FuzzySearch"`

	// <p>用户名</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// <p>客户端</p>
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// <p>流量来源，取值 Agent/Proxy/空；传Agent会返回Agent的日志，传Proxy会返回Proxy日志，两都传或不传则返回所有</p>
	SourceTypes []*string `json:"SourceTypes,omitnil,omitempty" name:"SourceTypes"`

	// <p>表名，长度限制64，多个表名查询的话可以用空格连接</p>
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// <p>字段名，长度限制64，多个字段名查询的话可以用空格连接</p>
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// <p>SQL 主要类型，DDL, DML, DCL, TCL</p>
	SqlMainTypes []*string `json:"SqlMainTypes,omitnil,omitempty" name:"SqlMainTypes"`

	// <p>操作类型</p>
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// <p>影响行数最小值</p>
	RowNumMin *int64 `json:"RowNumMin,omitnil,omitempty" name:"RowNumMin"`

	// <p>影响行数最大值</p>
	RowNumMax *int64 `json:"RowNumMax,omitnil,omitempty" name:"RowNumMax"`

	// <p>数据库类型</p>
	DbTypes []*string `json:"DbTypes,omitnil,omitempty" name:"DbTypes"`

	// <p>返回码</p>
	RetNo *int64 `json:"RetNo,omitnil,omitempty" name:"RetNo"`

	// <p>客户端工具</p>
	ClientDriverName *string `json:"ClientDriverName,omitnil,omitempty" name:"ClientDriverName"`

	// <p>客户端端口</p>
	ClientPort *int64 `json:"ClientPort,omitnil,omitempty" name:"ClientPort"`

	// <p>审计日志 ID</p>
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// <p>风险等级数组(0-安全,1-低风险,2-中风险,3-高风险), 如果要全部，则需要将所有的值都传入。如果为空，则会参考：DangerLevel 入参</p>
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`

	// <p>字段分类</p>
	SensitiveCategoryRule *string `json:"SensitiveCategoryRule,omitnil,omitempty" name:"SensitiveCategoryRule"`

	// <p>字段分级</p>
	SensitiveLevelRisk *string `json:"SensitiveLevelRisk,omitnil,omitempty" name:"SensitiveLevelRisk"`

	// <p>事务Id</p>
	TrxId *int64 `json:"TrxId,omitnil,omitempty" name:"TrxId"`

	// <p>clientMac</p>
	ClientMac *string `json:"ClientMac,omitnil,omitempty" name:"ClientMac"`
}

func (r *CreateDspmExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "DangerLevel")
	delete(f, "DbName")
	delete(f, "DbPort")
	delete(f, "DbIp")
	delete(f, "AssetsId")
	delete(f, "SessionId")
	delete(f, "ClientSideIp")
	delete(f, "EndTime")
	delete(f, "HitRule")
	delete(f, "StartTime")
	delete(f, "FuzzySearch")
	delete(f, "UserName")
	delete(f, "ClientName")
	delete(f, "SourceTypes")
	delete(f, "TableName")
	delete(f, "FieldName")
	delete(f, "SqlMainTypes")
	delete(f, "SqlType")
	delete(f, "RowNumMin")
	delete(f, "RowNumMax")
	delete(f, "DbTypes")
	delete(f, "RetNo")
	delete(f, "ClientDriverName")
	delete(f, "ClientPort")
	delete(f, "LogId")
	delete(f, "DangerLevels")
	delete(f, "SensitiveCategoryRule")
	delete(f, "SensitiveLevelRisk")
	delete(f, "TrxId")
	delete(f, "ClientMac")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmExportTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmExportTaskResponseParams `json:"Response"`
}

func (r *CreateDspmExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmIdentifyInfoListExportJobRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type CreateDspmIdentifyInfoListExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *CreateDspmIdentifyInfoListExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmIdentifyInfoListExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmIdentifyInfoListExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmIdentifyInfoListExportJobResponseParams struct {
	// 任务ID
	JobID *string `json:"JobID,omitnil,omitempty" name:"JobID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmIdentifyInfoListExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmIdentifyInfoListExportJobResponseParams `json:"Response"`
}

func (r *CreateDspmIdentifyInfoListExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmIdentifyInfoListExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmPersonalIdentifyRequestParams struct {
	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateDspmPersonalIdentifyRequest struct {
	*tchttp.BaseRequest
	
	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateDspmPersonalIdentifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmPersonalIdentifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Phone")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmPersonalIdentifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmPersonalIdentifyResponseParams struct {
	// 个人id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 身份id
	IdentifyId *string `json:"IdentifyId,omitnil,omitempty" name:"IdentifyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmPersonalIdentifyResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmPersonalIdentifyResponseParams `json:"Response"`
}

func (r *CreateDspmPersonalIdentifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmPersonalIdentifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmRiskExportJobRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type CreateDspmRiskExportJobRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *CreateDspmRiskExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmRiskExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmRiskExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmRiskExportJobResponseParams struct {
	// 任务ID
	JobID *string `json:"JobID,omitnil,omitempty" name:"JobID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmRiskExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmRiskExportJobResponseParams `json:"Response"`
}

func (r *CreateDspmRiskExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmRiskExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmWhitelistStrategyRequestParams struct {
	// 策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 白名单
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateDspmWhitelistStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 白名单
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateDspmWhitelistStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmWhitelistStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyType")
	delete(f, "MemberId")
	delete(f, "Name")
	delete(f, "Rule")
	delete(f, "AssetId")
	delete(f, "Account")
	delete(f, "Host")
	delete(f, "RiskId")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDspmWhitelistStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDspmWhitelistStrategyResponseParams struct {
	// 白名单id
	WhitelistStrategyId *string `json:"WhitelistStrategyId,omitnil,omitempty" name:"WhitelistStrategyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDspmWhitelistStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateDspmWhitelistStrategyResponseParams `json:"Response"`
}

func (r *CreateDspmWhitelistStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDspmWhitelistStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCAccessTokenRequestParams struct {
	// <p>CI/CD名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>扫描结果存储时长(30/60/90/120/150/180天)</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type CreateIaCAccessTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>CI/CD名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>扫描结果存储时长(30/60/90/120/150/180天)</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *CreateIaCAccessTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCAccessTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIaCAccessTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCAccessTokenResponseParams struct {
	// <p>接入Token</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIaCAccessTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateIaCAccessTokenResponseParams `json:"Response"`
}

func (r *CreateIaCAccessTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCFileExportJobRequestParams struct {
	// <p>过滤条件</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type CreateIaCFileExportJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>过滤条件</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *CreateIaCFileExportJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCFileExportJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIaCFileExportJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCFileExportJobResponseParams struct {
	// <p>任务ID</p>
	JobID *string `json:"JobID,omitnil,omitempty" name:"JobID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIaCFileExportJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateIaCFileExportJobResponseParams `json:"Response"`
}

func (r *CreateIaCFileExportJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCFileExportJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCFileReScanTaskRequestParams struct {
	// <p>文件ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type CreateIaCFileReScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>文件ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *CreateIaCFileReScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCFileReScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIaCFileReScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIaCFileReScanTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIaCFileReScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateIaCFileReScanTaskResponseParams `json:"Response"`
}

func (r *CreateIaCFileReScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIaCFileReScanTaskResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type CreateSkillScanRequestParams struct {
	// ZIP 文件内容的 Base64 编码
	// 入参限制：文件大小上限 7MB（编码前），仅接受有效的 ZIP 格式
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文件名，用于服务端日志记录
	// 参数格式：形如 my-skill.zip
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type CreateSkillScanRequest struct {
	*tchttp.BaseRequest
	
	// ZIP 文件内容的 Base64 编码
	// 入参限制：文件大小上限 7MB（编码前），仅接受有效的 ZIP 格式
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// 文件名，用于服务端日志记录
	// 参数格式：形如 my-skill.zip
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

func (r *CreateSkillScanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillScanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileBase64")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSkillScanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillScanResponseParams struct {
	// 文件的 SHA256 Hash，用于轮询 DescribeSkillScanResult 接口
	// 参数格式：sha256:<64位hex>
	ContentHash *string `json:"ContentHash,omitnil,omitempty" name:"ContentHash"`

	// 本次请求实际绑定的引擎版本号。调用方应保存并在后续 DescribeSkillScanResult 时显式传入
	EngineVersion *int64 `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 任务状态，固定为 SCANNING，表示任务已接收
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 可读的操作结果描述
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSkillScanResponse struct {
	*tchttp.BaseResponse
	Response *CreateSkillScanResponseParams `json:"Response"`
}

func (r *CreateSkillScanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialEffectScope struct {
	// 是否排除模式
	// 枚举值：
	// 0：包含模式（仅Instances中的机器生效），此时Instances必填
	// 1：排除模式（Instances中的机器不生效，其余机器生效），此时Instances可选（空列表表示全部机器生效）
	Exclude *int64 `json:"Exclude,omitnil,omitempty" name:"Exclude"`

	// 机器实例ID列表。Exclude为0时必填，表示仅这些机器可访问凭证；Exclude为1时可选，表示这些机器不可访问凭证（空列表表示全部机器生效）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*string `json:"Instances,omitnil,omitempty" name:"Instances"`
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
type DeleteCosAkAssetRequestParams struct {
	// 要删除的cos ak集合
	CosAkSet []*CosAkSet `json:"CosAkSet,omitnil,omitempty" name:"CosAkSet"`
}

type DeleteCosAkAssetRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的cos ak集合
	CosAkSet []*CosAkSet `json:"CosAkSet,omitnil,omitempty" name:"CosAkSet"`
}

func (r *DeleteCosAkAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCosAkAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosAkSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCosAkAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCosAkAssetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCosAkAssetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCosAkAssetResponseParams `json:"Response"`
}

func (r *DeleteCosAkAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCosAkAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCosPolicyRequestParams struct {
	// 要删除的策略集合
	PolicyIdSet []*int64 `json:"PolicyIdSet,omitnil,omitempty" name:"PolicyIdSet"`

	// 是否删除所有
	IsDeleteAll *int64 `json:"IsDeleteAll,omitnil,omitempty" name:"IsDeleteAll"`
}

type DeleteCosPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的策略集合
	PolicyIdSet []*int64 `json:"PolicyIdSet,omitnil,omitempty" name:"PolicyIdSet"`

	// 是否删除所有
	IsDeleteAll *int64 `json:"IsDeleteAll,omitnil,omitempty" name:"IsDeleteAll"`
}

func (r *DeleteCosPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCosPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyIdSet")
	delete(f, "IsDeleteAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCosPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCosPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCosPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCosPolicyResponseParams `json:"Response"`
}

func (r *DeleteCosPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCosPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
type DeleteDspmApplyOrderRequestParams struct {
	// 申请单id
	OrderId []*string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

type DeleteDspmApplyOrderRequest struct {
	*tchttp.BaseRequest
	
	// 申请单id
	OrderId []*string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

func (r *DeleteDspmApplyOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmApplyOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDspmApplyOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmApplyOrderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDspmApplyOrderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDspmApplyOrderResponseParams `json:"Response"`
}

func (r *DeleteDspmApplyOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmApplyOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmAssetAccountRequestParams struct {
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

type DeleteDspmAssetAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

func (r *DeleteDspmAssetAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmAssetAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "Account")
	delete(f, "Host")
	delete(f, "RiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDspmAssetAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmAssetAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDspmAssetAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDspmAssetAccountResponseParams `json:"Response"`
}

func (r *DeleteDspmAssetAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmAssetAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmBackupLogListRequestParams struct {
	// 备份日志Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DeleteDspmBackupLogListRequest struct {
	*tchttp.BaseRequest
	
	// 备份日志Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DeleteDspmBackupLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmBackupLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDspmBackupLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmBackupLogListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDspmBackupLogListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDspmBackupLogListResponseParams `json:"Response"`
}

func (r *DeleteDspmBackupLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmBackupLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmExportTaskRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 要删除的导出任务ID
	TaskIds []*int64 `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type DeleteDspmExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 要删除的导出任务ID
	TaskIds []*int64 `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *DeleteDspmExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDspmExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmExportTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDspmExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDspmExportTaskResponseParams `json:"Response"`
}

func (r *DeleteDspmExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmPersonalIdentifyRequestParams struct {
	// 个人id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

type DeleteDspmPersonalIdentifyRequest struct {
	*tchttp.BaseRequest
	
	// 个人id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`
}

func (r *DeleteDspmPersonalIdentifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmPersonalIdentifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDspmPersonalIdentifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmPersonalIdentifyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDspmPersonalIdentifyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDspmPersonalIdentifyResponseParams `json:"Response"`
}

func (r *DeleteDspmPersonalIdentifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmPersonalIdentifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmRestoreLogListRequestParams struct {
	// 日志Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DeleteDspmRestoreLogListRequest struct {
	*tchttp.BaseRequest
	
	// 日志Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DeleteDspmRestoreLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmRestoreLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDspmRestoreLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmRestoreLogListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDspmRestoreLogListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDspmRestoreLogListResponseParams `json:"Response"`
}

func (r *DeleteDspmRestoreLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmRestoreLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmWhitelistStrategyRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 白名单id
	WhitelistStrategyId []*string `json:"WhitelistStrategyId,omitnil,omitempty" name:"WhitelistStrategyId"`
}

type DeleteDspmWhitelistStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 白名单id
	WhitelistStrategyId []*string `json:"WhitelistStrategyId,omitnil,omitempty" name:"WhitelistStrategyId"`
}

func (r *DeleteDspmWhitelistStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmWhitelistStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "WhitelistStrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDspmWhitelistStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDspmWhitelistStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDspmWhitelistStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDspmWhitelistStrategyResponseParams `json:"Response"`
}

func (r *DeleteDspmWhitelistStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDspmWhitelistStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIaCAccessTokenRequestParams struct {
	// <p>删除ID列表</p>
	Id []*uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteIaCAccessTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>删除ID列表</p>
	Id []*uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteIaCAccessTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIaCAccessTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIaCAccessTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIaCAccessTokenResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIaCAccessTokenResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIaCAccessTokenResponseParams `json:"Response"`
}

func (r *DeleteIaCAccessTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIaCAccessTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIaCFileRequestParams struct {
	// <p>删除ID列表</p>
	Id []*uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteIaCFileRequest struct {
	*tchttp.BaseRequest
	
	// <p>删除ID列表</p>
	Id []*uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteIaCFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIaCFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIaCFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIaCFileResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIaCFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIaCFileResponseParams `json:"Response"`
}

func (r *DeleteIaCFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIaCFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskScanTaskRequestParams struct {
	// 任务id 和目标AppID列表
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DeleteRiskScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id 和目标AppID列表
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
type DescribeAIAgentAssetListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeAIAgentAssetListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeAIAgentAssetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAgentAssetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIAgentAssetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIAgentAssetListResponseParams struct {
	// 资产列表
	AssetList []*AIAgentAsset `json:"AssetList,omitnil,omitempty" name:"AssetList"`

	// 资产总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIAgentAssetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIAgentAssetListResponseParams `json:"Response"`
}

func (r *DescribeAIAgentAssetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIAgentAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAILinkSettingRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeAILinkSettingRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeAILinkSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAILinkSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAILinkSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAILinkSettingResponseParams struct {
	// <p>0 关闭AI-Link智链引擎，1 开启AI-Link智链引擎</p>
	AILinkEnable *uint64 `json:"AILinkEnable,omitnil,omitempty" name:"AILinkEnable"`

	// <p>深度模式 0-关闭 1-开启</p>
	RuleScopeDeep *uint64 `json:"RuleScopeDeep,omitnil,omitempty" name:"RuleScopeDeep"`

	// <p>均衡模式 0-关闭 1-开启</p>
	RuleScopeBalanced *uint64 `json:"RuleScopeBalanced,omitnil,omitempty" name:"RuleScopeBalanced"`

	// <p>精准模式 0-关闭 1-开启</p>
	RuleScopePrecise *uint64 `json:"RuleScopePrecise,omitnil,omitempty" name:"RuleScopePrecise"`

	// <p>1 全部专业/旗舰版主机，0 Quuids列表主机</p>
	Scope *uint64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>自选主机Quuid列表</p>
	Quuids []*string `json:"Quuids,omitnil,omitempty" name:"Quuids"`

	// <p>排除主机Quuid列表</p>
	ExcludeQuuids []*string `json:"ExcludeQuuids,omitnil,omitempty" name:"ExcludeQuuids"`

	// <p>新增资产自动包含 0 不包含 1包含</p>
	AutoInclude *uint64 `json:"AutoInclude,omitnil,omitempty" name:"AutoInclude"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAILinkSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAILinkSettingResponseParams `json:"Response"`
}

func (r *DescribeAILinkSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAILinkSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAKAnalysisDetailRequestParams struct {
	// 告警记录ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeAKAnalysisDetailRequest struct {
	*tchttp.BaseRequest
	
	// 告警记录ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeAKAnalysisDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAKAnalysisDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAKAnalysisDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAKAnalysisDetailResponseParams struct {
	// 告警AI分析状态 -1 分析失败 0 未分析 1 分析中 2 分析成功，真实告警 3 分析成功，可疑告警
	AIStatus *int64 `json:"AIStatus,omitnil,omitempty" name:"AIStatus"`

	// AI分析任务ID
	AITaskID *string `json:"AITaskID,omitnil,omitempty" name:"AITaskID"`

	// 告警AI分析结果，base64格式，避免数据被拦截
	AIResult *string `json:"AIResult,omitnil,omitempty" name:"AIResult"`

	// 反馈建议
	Feedback *string `json:"Feedback,omitnil,omitempty" name:"Feedback"`

	// 反馈状态  0表示没有反馈，1表示认可，2表示不认可
	FeedbackResult *int64 `json:"FeedbackResult,omitnil,omitempty" name:"FeedbackResult"`

	// 失败原因
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAKAnalysisDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAKAnalysisDetailResponseParams `json:"Response"`
}

func (r *DescribeAKAnalysisDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAKAnalysisDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalCallRecordRequestParams struct {
	// 告警规则ID
	AlarmRuleID *int64 `json:"AlarmRuleID,omitnil,omitempty" name:"AlarmRuleID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 访问密钥
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 调用源IP
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeAbnormalCallRecordRequest struct {
	*tchttp.BaseRequest
	
	// 告警规则ID
	AlarmRuleID *int64 `json:"AlarmRuleID,omitnil,omitempty" name:"AlarmRuleID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 访问密钥
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 调用源IP
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeAbnormalCallRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalCallRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmRuleID")
	delete(f, "MemberId")
	delete(f, "AccessKey")
	delete(f, "SourceIP")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAbnormalCallRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAbnormalCallRecordResponseParams struct {
	// 调用记录列表
	Data []*CallRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 调用记录总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAbnormalCallRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAbnormalCallRecordResponseParams `json:"Response"`
}

func (r *DescribeAbnormalCallRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAbnormalCallRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAlarmDetailRequestParams struct {
	// 告警记录ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeAccessKeyAlarmDetailRequest struct {
	*tchttp.BaseRequest
	
	// 告警记录ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeAccessKeyAlarmDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAlarmDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyAlarmDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAlarmDetailResponseParams struct {
	// 告警信息
	AlarmInfo *AccessKeyAlarm `json:"AlarmInfo,omitnil,omitempty" name:"AlarmInfo"`

	// 所属账号CAM策略数量
	CamCount *int64 `json:"CamCount,omitnil,omitempty" name:"CamCount"`

	// AK风险数量
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 告警策略描述
	AlarmDesc *string `json:"AlarmDesc,omitnil,omitempty" name:"AlarmDesc"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyAlarmDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyAlarmDetailResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyAlarmDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAlarmDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAlarmRequestParams struct {
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 访问密钥的ID
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// 源IP的ID
	SourceIPID *uint64 `json:"SourceIPID,omitnil,omitempty" name:"SourceIPID"`

	// 账号uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

type DescribeAccessKeyAlarmRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 访问密钥的ID
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// 源IP的ID
	SourceIPID *uint64 `json:"SourceIPID,omitnil,omitempty" name:"SourceIPID"`

	// 账号uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

func (r *DescribeAccessKeyAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	delete(f, "AccessKeyID")
	delete(f, "SourceIPID")
	delete(f, "SubUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAlarmResponseParams struct {
	// 告警列表
	Data []*AccessKeyAlarm `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyAlarmResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyAlarmResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAssetRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeAccessKeyAssetRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeAccessKeyAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyAssetResponseParams struct {
	// 访问密钥资产列表
	Data []*AccessKeyAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// 全部数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyAssetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyAssetResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyRiskDetailRequestParams struct {
	// 风险记录ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeAccessKeyRiskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 风险记录ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeAccessKeyRiskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyRiskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyRiskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyRiskDetailResponseParams struct {
	// 风险列表
	RiskInfo *AccessKeyRisk `json:"RiskInfo,omitnil,omitempty" name:"RiskInfo"`

	// CAM策略总数
	CamCount *int64 `json:"CamCount,omitnil,omitempty" name:"CamCount"`

	// 账号关联告警数量
	AlarmCount *int64 `json:"AlarmCount,omitnil,omitempty" name:"AlarmCount"`

	// 访问方式 0 API 1 控制台与API
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 访问密钥告警数量列表
	AccessKeyAlarmCount []*AccessKeyAlarmCount `json:"AccessKeyAlarmCount,omitnil,omitempty" name:"AccessKeyAlarmCount"`

	// 操作保护是否开启 0 未开启 1 已开启
	ActionFlag *int64 `json:"ActionFlag,omitnil,omitempty" name:"ActionFlag"`

	// 登录保护是否开启 0 未开启 1 已开启
	LoginFlag *int64 `json:"LoginFlag,omitnil,omitempty" name:"LoginFlag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyRiskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyRiskDetailResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyRiskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyRiskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyRiskRequestParams struct {
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 访问密钥的ID
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// 账号uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

type DescribeAccessKeyRiskRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 访问密钥的ID
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// 账号uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

func (r *DescribeAccessKeyRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	delete(f, "AccessKeyID")
	delete(f, "SubUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyRiskResponseParams struct {
	// 风险列表
	Data []*AccessKeyRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyRiskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyRiskResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyUserDetailRequestParams struct {
	// 账号自身uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeAccessKeyUserDetailRequest struct {
	*tchttp.BaseRequest
	
	// 账号自身uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeAccessKeyUserDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyUserDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubUin")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyUserDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyUserDetailResponseParams struct {
	// 账号详情信息
	User *AccessKeyUser `json:"User,omitnil,omitempty" name:"User"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyUserDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyUserDetailResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyUserDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyUserDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyUserListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeAccessKeyUserListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeAccessKeyUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessKeyUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessKeyUserListResponseParams struct {
	// 账号列表
	Data []*AccessKeyUser `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessKeyUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessKeyUserListResponseParams `json:"Response"`
}

func (r *DescribeAccessKeyUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessKeyUserListResponse) FromJsonString(s string) error {
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
	// <p>集团账号的成员id</p>
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
	
	// <p>集团账号的成员id</p>
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
	// <p>集团账号的成员id</p>
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
	
	// <p>集团账号的成员id</p>
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

	// 等保规范名称集合
	StandardNameList []*StandardItem `json:"StandardNameList,omitnil,omitempty" name:"StandardNameList"`

	// 资产类型集合
	AssetTypeList []*AttributeOptionSet `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

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
type DescribeAssumeRoleRequestParams struct {
	// 角色名
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type DescribeAssumeRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *DescribeAssumeRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssumeRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssumeRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssumeRoleResponseParams struct {
	// 是否绑定角色。0-未绑定 1-已绑定
	Bind *int64 `json:"Bind,omitnil,omitempty" name:"Bind"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssumeRoleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssumeRoleResponseParams `json:"Response"`
}

func (r *DescribeAssumeRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssumeRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBucketInvokeIpListRequestParams struct {
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeBucketInvokeIpListRequest struct {
	*tchttp.BaseRequest
	
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeBucketInvokeIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBucketInvokeIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "BucketName")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBucketInvokeIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBucketInvokeIpListResponseParams struct {
	// ip信息
	Data []*CosSourceIpInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBucketInvokeIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBucketInvokeIpListResponseParams `json:"Response"`
}

func (r *DescribeBucketInvokeIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBucketInvokeIpListResponse) FromJsonString(s string) error {
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
type DescribeCLSLogIndexV3RequestParams struct {
	// <p>过滤条件</p>
	Filters []*LogCLSFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>limit限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>offset</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCLSLogIndexV3Request struct {
	*tchttp.BaseRequest
	
	// <p>过滤条件</p>
	Filters []*LogCLSFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>limit限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>offset</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeCLSLogIndexV3Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSLogIndexV3Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "MemberId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCLSLogIndexV3Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCLSLogIndexV3ResponseParams struct {
	// <p>主题信息</p>
	TopicIndexInfos []*LogTopicIndexInfo `json:"TopicIndexInfos,omitnil,omitempty" name:"TopicIndexInfos"`

	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCLSLogIndexV3Response struct {
	*tchttp.BaseResponse
	Response *DescribeCLSLogIndexV3ResponseParams `json:"Response"`
}

func (r *DescribeCLSLogIndexV3Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSLogIndexV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCLSLogListV3RequestParams struct {
	// <p>开始时间</p>
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// <p>结束时间</p>
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// <p>查询条件</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>语法</p>
	SyntaxRule *int64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// <p>主题</p>
	Topics []*LogContextInfo `json:"Topics,omitnil,omitempty" name:"Topics"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>排序</p>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>limit</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>offset</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>采样</p>
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// <p>是否高亮</p>
	HighLight *bool `json:"HighLight,omitnil,omitempty" name:"HighLight"`

	// <p>是否采用新分析</p>
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`

	// <p>查询优化</p>
	QueryOptimize *int64 `json:"QueryOptimize,omitnil,omitempty" name:"QueryOptimize"`

	// <p>主题id</p>
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// <p>上下文信息</p>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>查询类型</p>
	SubQueryTypes []*string `json:"SubQueryTypes,omitnil,omitempty" name:"SubQueryTypes"`
}

type DescribeCLSLogListV3Request struct {
	*tchttp.BaseRequest
	
	// <p>开始时间</p>
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// <p>结束时间</p>
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// <p>查询条件</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>语法</p>
	SyntaxRule *int64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// <p>主题</p>
	Topics []*LogContextInfo `json:"Topics,omitnil,omitempty" name:"Topics"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>排序</p>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>limit</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>offset</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>采样</p>
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// <p>是否高亮</p>
	HighLight *bool `json:"HighLight,omitnil,omitempty" name:"HighLight"`

	// <p>是否采用新分析</p>
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`

	// <p>查询优化</p>
	QueryOptimize *int64 `json:"QueryOptimize,omitnil,omitempty" name:"QueryOptimize"`

	// <p>主题id</p>
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// <p>上下文信息</p>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>查询类型</p>
	SubQueryTypes []*string `json:"SubQueryTypes,omitnil,omitempty" name:"SubQueryTypes"`
}

func (r *DescribeCLSLogListV3Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSLogListV3Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "SyntaxRule")
	delete(f, "Topics")
	delete(f, "MemberId")
	delete(f, "Sort")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SamplingRate")
	delete(f, "HighLight")
	delete(f, "UseNewAnalysis")
	delete(f, "QueryOptimize")
	delete(f, "TopicId")
	delete(f, "Context")
	delete(f, "SubQueryTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCLSLogListV3Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCLSLogListV3ResponseParams struct {
	// <p>上下文</p>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>listover</p>
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// <p>是否采用分析</p>
	Analysis *bool `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// <p>结果</p>
	Results []*LogSearchResult `json:"Results,omitnil,omitempty" name:"Results"`

	// <p>列名</p>
	ColNames []*string `json:"ColNames,omitnil,omitempty" name:"ColNames"`

	// <p>分析结果</p>
	AnalysisResults []*LogItems `json:"AnalysisResults,omitnil,omitempty" name:"AnalysisResults"`

	// <p>分析记录</p>
	AnalysisRecords []*string `json:"AnalysisRecords,omitnil,omitempty" name:"AnalysisRecords"`

	// <p>列名</p>
	Columns []*LogColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// <p>采样</p>
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// <p>主题信息</p>
	Topics *LogSearchTopics `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCLSLogListV3Response struct {
	*tchttp.BaseResponse
	Response *DescribeCLSLogListV3ResponseParams `json:"Response"`
}

func (r *DescribeCLSLogListV3Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCLSLogListV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCSIPRiskStatisticsRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤内容
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCSIPRiskStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
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
type DescribeCWPMachineDetailRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeCWPMachineDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeCWPMachineDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCWPMachineDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCWPMachineDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCWPMachineDetailResponseParams struct {
	// <p>主机详情</p>
	MachineDetail *MachineDetail `json:"MachineDetail,omitnil,omitempty" name:"MachineDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCWPMachineDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCWPMachineDetailResponseParams `json:"Response"`
}

func (r *DescribeCWPMachineDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCWPMachineDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCWPMachinesRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>一、主表字段筛选（需要指定 OperatorType）<br>MachineName：主机名称，支持 OperatorType 9(模糊)、7(IN)，Values示例：[&quot;test-server&quot;]<br>MachineIp：内网IP，支持 OperatorType 9(模糊)、7(IN)，Values示例：[&quot;10.0.0.1&quot;]<br>MachineWanIp：外网IP，支持 OperatorType 9(模糊)、7(IN)，Values示例：[&quot;1.2.3.4&quot;]<br>InstanceID：实例ID，支持 OperatorType 9(模糊)、7(IN)，Values示例：[&quot;ins-xxxxx&quot;]<br>MachineStatus / InstanceStatus：实例状态，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;RUNNING&quot;]，可选值：RUNNING/STOPPED/EXPIRED<br>MachineOs：操作系统类型，支持 OperatorType 7(IN)，Values示例：[&quot;1&quot;]，值为数字编码，见下方OsType说明<br>VpcId：VPC ID，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;vpc-xxxxx&quot;]<br>CloudFromEnum：云服务商，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;0&quot;]，值为数字编码，见下方CloudFrom说明<br>Region ：地域，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;ap-guangzhou&quot;]<br>AppId：账号AppId，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;1234567890&quot;]<br>ProjectId：项目ID，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;0&quot;]</p><p>二、预筛选字段（不需要指定 OperatorType）<br>AgentStatus：Agent状态，单选，Values示例：[&quot;ONLINE&quot;]，可选值：ONLINE/OFFLINE/UNINSTALL<br>ProtectType：防护类型（综合），Values示例：[&quot;ULTIMATE&quot;]，可选值：BASIC/PRO/ULTIMATE/NONE<br>CsipProtectType：CSIP防护类型，Values示例：[&quot;ULTIMATE&quot;]，可选值：BASIC/PRO/ULTIMATE/NONE<br>CloudTags：云标签，Values示例：[&quot;tagKey$tagValue&quot;]，格式：tagKey$tagValue 或 tagKey（只匹配key），最多5个值<br>Tags：资产标签，Values示例：[&quot;123&quot;]，值为标签ID<br>ExposedStatus：暴露状态，单选，Values示例：[&quot;EXPOSED&quot;]，可选值：NOT_APPLICABLE/EXPOSED/UNEXPOSED</p><p>三、特殊筛选字段（不需要指定 OperatorType）<br>NetworkType：网络类型，单选，Values示例：[&quot;1&quot;]，1=VPC网络, 2=基础网络, 3=非腾讯云网络<br>MachineType：机器类型，可多选，Values示例：[&quot;CVM&quot;]，可选值：CVM/BM/ECM/LH/EKS-NATIVE/ECS/EC2/VMS<br>Common：通用搜索，单选，Values示例：[&quot;关键词&quot;]，同时对内网IP、外网IP、主机名称、实例ID做模糊匹配</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>是否需要tat状态信息</p>
	NeedTatStatus *bool `json:"NeedTatStatus,omitnil,omitempty" name:"NeedTatStatus"`

	// <p>是否需要额外信息，如安全中心标签、腾讯云标签</p>
	MoreInformation *bool `json:"MoreInformation,omitnil,omitempty" name:"MoreInformation"`

	// <p>是否需要容器信息，如容器数、核数、容器防护状态</p>
	NeedContainerInfo *bool `json:"NeedContainerInfo,omitnil,omitempty" name:"NeedContainerInfo"`
}

type DescribeCWPMachinesRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>一、主表字段筛选（需要指定 OperatorType）<br>MachineName：主机名称，支持 OperatorType 9(模糊)、7(IN)，Values示例：[&quot;test-server&quot;]<br>MachineIp：内网IP，支持 OperatorType 9(模糊)、7(IN)，Values示例：[&quot;10.0.0.1&quot;]<br>MachineWanIp：外网IP，支持 OperatorType 9(模糊)、7(IN)，Values示例：[&quot;1.2.3.4&quot;]<br>InstanceID：实例ID，支持 OperatorType 9(模糊)、7(IN)，Values示例：[&quot;ins-xxxxx&quot;]<br>MachineStatus / InstanceStatus：实例状态，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;RUNNING&quot;]，可选值：RUNNING/STOPPED/EXPIRED<br>MachineOs：操作系统类型，支持 OperatorType 7(IN)，Values示例：[&quot;1&quot;]，值为数字编码，见下方OsType说明<br>VpcId：VPC ID，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;vpc-xxxxx&quot;]<br>CloudFromEnum：云服务商，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;0&quot;]，值为数字编码，见下方CloudFrom说明<br>Region ：地域，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;ap-guangzhou&quot;]<br>AppId：账号AppId，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;1234567890&quot;]<br>ProjectId：项目ID，支持 OperatorType 7(IN)、1(等于)，Values示例：[&quot;0&quot;]</p><p>二、预筛选字段（不需要指定 OperatorType）<br>AgentStatus：Agent状态，单选，Values示例：[&quot;ONLINE&quot;]，可选值：ONLINE/OFFLINE/UNINSTALL<br>ProtectType：防护类型（综合），Values示例：[&quot;ULTIMATE&quot;]，可选值：BASIC/PRO/ULTIMATE/NONE<br>CsipProtectType：CSIP防护类型，Values示例：[&quot;ULTIMATE&quot;]，可选值：BASIC/PRO/ULTIMATE/NONE<br>CloudTags：云标签，Values示例：[&quot;tagKey$tagValue&quot;]，格式：tagKey$tagValue 或 tagKey（只匹配key），最多5个值<br>Tags：资产标签，Values示例：[&quot;123&quot;]，值为标签ID<br>ExposedStatus：暴露状态，单选，Values示例：[&quot;EXPOSED&quot;]，可选值：NOT_APPLICABLE/EXPOSED/UNEXPOSED</p><p>三、特殊筛选字段（不需要指定 OperatorType）<br>NetworkType：网络类型，单选，Values示例：[&quot;1&quot;]，1=VPC网络, 2=基础网络, 3=非腾讯云网络<br>MachineType：机器类型，可多选，Values示例：[&quot;CVM&quot;]，可选值：CVM/BM/ECM/LH/EKS-NATIVE/ECS/EC2/VMS<br>Common：通用搜索，单选，Values示例：[&quot;关键词&quot;]，同时对内网IP、外网IP、主机名称、实例ID做模糊匹配</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>是否需要tat状态信息</p>
	NeedTatStatus *bool `json:"NeedTatStatus,omitnil,omitempty" name:"NeedTatStatus"`

	// <p>是否需要额外信息，如安全中心标签、腾讯云标签</p>
	MoreInformation *bool `json:"MoreInformation,omitnil,omitempty" name:"MoreInformation"`

	// <p>是否需要容器信息，如容器数、核数、容器防护状态</p>
	NeedContainerInfo *bool `json:"NeedContainerInfo,omitnil,omitempty" name:"NeedContainerInfo"`
}

func (r *DescribeCWPMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCWPMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "NeedTatStatus")
	delete(f, "MoreInformation")
	delete(f, "NeedContainerInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCWPMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCWPMachinesResponseParams struct {
	// <p>主机列表</p>
	Machines []*Machine `json:"Machines,omitnil,omitempty" name:"Machines"`

	// <p>总数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCWPMachinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCWPMachinesResponseParams `json:"Response"`
}

func (r *DescribeCWPMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCWPMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallRecordRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 访问密钥的ID
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// 调用源IP的ID
	SourceIPID *uint64 `json:"SourceIPID,omitnil,omitempty" name:"SourceIPID"`

	// 访问账号uin
	AccUin *string `json:"AccUin,omitnil,omitempty" name:"AccUin"`

	// 访问密钥，注意：不支持临时密钥的情况
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCallRecordRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 访问密钥的ID
	AccessKeyID *uint64 `json:"AccessKeyID,omitnil,omitempty" name:"AccessKeyID"`

	// 调用源IP的ID
	SourceIPID *uint64 `json:"SourceIPID,omitnil,omitempty" name:"SourceIPID"`

	// 访问账号uin
	AccUin *string `json:"AccUin,omitnil,omitempty" name:"AccUin"`

	// 访问密钥，注意：不支持临时密钥的情况
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCallRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "AccessKeyID")
	delete(f, "SourceIPID")
	delete(f, "AccUin")
	delete(f, "AccessKey")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallRecordResponseParams struct {
	// 调用记录列表
	Data []*CallRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 调用记录总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCallRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallRecordResponseParams `json:"Response"`
}

func (r *DescribeCallRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckViewRisksRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>过滤内容</p>
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分页大小</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序类型</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>排序字段</p>
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeCheckViewRisksRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>过滤内容</p>
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分页大小</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序类型</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>排序字段</p>
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
	// <p>检查视角下风险数量</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>检查视角下风险列表</p>
	CheckViewRiskList []*CheckViewRiskItem `json:"CheckViewRiskList,omitnil,omitempty" name:"CheckViewRiskList"`

	// <p>检查视角下cspm规范标签列表</p>
	StandardNameList []*StandardItem `json:"StandardNameList,omitnil,omitempty" name:"StandardNameList"`

	// <p>资产类型集合</p>
	AssetTypeList []*AttributeOptionSet `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// <p>云厂商类型集合</p>
	ProviderList []*AttributeOptionSet `json:"ProviderList,omitnil,omitempty" name:"ProviderList"`

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
type DescribeConfigCheckRulesRequestParams struct {
	// <p>集团账号的成员id</p>
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

type DescribeConfigCheckRulesRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
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

func (r *DescribeConfigCheckRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigCheckRulesRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigCheckRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigCheckRulesResponseParams struct {
	// 风险规则数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 风险规则列表
	RuleList []*RiskRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// 云厂商类型选项
	ProviderList []*AttributeOptionSet `json:"ProviderList,omitnil,omitempty" name:"ProviderList"`

	// 风险等级类型选项
	RiskLevelList []*AttributeOptionSet `json:"RiskLevelList,omitnil,omitempty" name:"RiskLevelList"`

	// 处置分类选项
	DispositionTypeList []*AttributeOptionSet `json:"DispositionTypeList,omitnil,omitempty" name:"DispositionTypeList"`

	// 检查类型选项
	CheckTypeList []*AttributeOptionSet `json:"CheckTypeList,omitnil,omitempty" name:"CheckTypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigCheckRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigCheckRulesResponseParams `json:"Response"`
}

func (r *DescribeConfigCheckRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigCheckRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAccessPermissionRequestParams struct {
	// 关联的appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 需要查看的uin
	RelUin *string `json:"RelUin,omitnil,omitempty" name:"RelUin"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosAccessPermissionRequest struct {
	*tchttp.BaseRequest
	
	// 关联的appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 需要查看的uin
	RelUin *string `json:"RelUin,omitnil,omitempty" name:"RelUin"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosAccessPermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAccessPermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "BucketName")
	delete(f, "RelUin")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAccessPermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAccessPermissionResponseParams struct {
	// cos权限信息
	Data []*CosPermissionInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAccessPermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAccessPermissionResponseParams `json:"Response"`
}

func (r *DescribeCosAccessPermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAccessPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAccessPermissionsRequestParams struct {
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosAccessPermissionsRequest struct {
	*tchttp.BaseRequest
	
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosAccessPermissionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAccessPermissionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "BucketName")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAccessPermissionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAccessPermissionsResponseParams struct {
	// 返回数据列表
	Data []*CosAccessInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAccessPermissionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAccessPermissionsResponseParams `json:"Response"`
}

func (r *DescribeCosAccessPermissionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAccessPermissionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosActionListRequestParams struct {
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosActionListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosActionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosActionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosActionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosActionListResponseParams struct {
	// 列表
	Data []*CosActionInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosActionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosActionListResponseParams `json:"Response"`
}

func (r *DescribeCosActionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosActionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAkAssetRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 查询过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosAkAssetRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 查询过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosAkAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAkAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAkAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAkAssetResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// ak资产列表
	Data []*CosAkAssetInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAkAssetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAkAssetResponseParams `json:"Response"`
}

func (r *DescribeCosAkAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAkAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAkInvokeIpListRequestParams struct {
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// ak
	Ak *string `json:"Ak,omitnil,omitempty" name:"Ak"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosAkInvokeIpListRequest struct {
	*tchttp.BaseRequest
	
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// ak
	Ak *string `json:"Ak,omitnil,omitempty" name:"Ak"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosAkInvokeIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAkInvokeIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "Ak")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAkInvokeIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAkInvokeIpListResponseParams struct {
	// ip信息
	Data []*CosSourceIpInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAkInvokeIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAkInvokeIpListResponseParams `json:"Response"`
}

func (r *DescribeCosAkInvokeIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAkInvokeIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAlarmListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosAlarmListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosAlarmListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAlarmListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAlarmListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAlarmListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 告警列表
	Data []*CosAlarmInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAlarmListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAlarmListResponseParams `json:"Response"`
}

func (r *DescribeCosAlarmListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAlarmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAlarmTrendDataRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 需要查看多久的时间
	LastDays *int64 `json:"LastDays,omitnil,omitempty" name:"LastDays"`
}

type DescribeCosAlarmTrendDataRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 需要查看多久的时间
	LastDays *int64 `json:"LastDays,omitnil,omitempty" name:"LastDays"`
}

func (r *DescribeCosAlarmTrendDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAlarmTrendDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "LastDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAlarmTrendDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAlarmTrendDataResponseParams struct {
	// 告警趋势信息
	CosAlarmTrendInfo []*CosAlarmTrendInfo `json:"CosAlarmTrendInfo,omitnil,omitempty" name:"CosAlarmTrendInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAlarmTrendDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAlarmTrendDataResponseParams `json:"Response"`
}

func (r *DescribeCosAlarmTrendDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAlarmTrendDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAssetRequestParams struct {
	// 请求过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeCosAssetRequest struct {
	*tchttp.BaseRequest
	
	// 请求过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeCosAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAssetResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 桶信息
	Data []*CosAssetInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAssetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAssetResponseParams `json:"Response"`
}

func (r *DescribeCosAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAssetSyncTaskRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeCosAssetSyncTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeCosAssetSyncTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAssetSyncTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAssetSyncTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAssetSyncTaskResponseParams struct {
	// 数据信息
	Data []*CosAssetSyncTaskInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAssetSyncTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAssetSyncTaskResponseParams `json:"Response"`
}

func (r *DescribeCosAssetSyncTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAssetSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAuditAppIdListRequestParams struct {

}

type DescribeCosAuditAppIdListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCosAuditAppIdListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAuditAppIdListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAuditAppIdListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAuditAppIdListResponseParams struct {
	// 已购买appid集合
	Data []*int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAuditAppIdListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAuditAppIdListResponseParams `json:"Response"`
}

func (r *DescribeCosAuditAppIdListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAuditAppIdListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAuditDictionaryListRequestParams struct {
	// <p>字典类型（RootCategory：一级分类，IdentifyRule:敏感识别数据项）</p>
	DictType *string `json:"DictType,omitnil,omitempty" name:"DictType"`

	// <p>筛选条件</p>
	Filters []*WhereFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCosAuditDictionaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>字典类型（RootCategory：一级分类，IdentifyRule:敏感识别数据项）</p>
	DictType *string `json:"DictType,omitnil,omitempty" name:"DictType"`

	// <p>筛选条件</p>
	Filters []*WhereFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCosAuditDictionaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAuditDictionaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DictType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAuditDictionaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAuditDictionaryListResponseParams struct {
	// <p>结果集</p>
	DataSet []*CosDictionary `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAuditDictionaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAuditDictionaryListResponseParams `json:"Response"`
}

func (r *DescribeCosAuditDictionaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAuditDictionaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAuditPayInfoRequestParams struct {

}

type DescribeCosAuditPayInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCosAuditPayInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAuditPayInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosAuditPayInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosAuditPayInfoResponseParams struct {
	// cos审计支付信息
	CosAuditPayInfo *CosAuditPayInfo `json:"CosAuditPayInfo,omitnil,omitempty" name:"CosAuditPayInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosAuditPayInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosAuditPayInfoResponseParams `json:"Response"`
}

func (r *DescribeCosAuditPayInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosAuditPayInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosBucketBillingInfoRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeCosBucketBillingInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeCosBucketBillingInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosBucketBillingInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosBucketBillingInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosBucketBillingInfoResponseParams struct {
	// 存储桶计费信息
	CosBucketBillingInfoSet []*CosBucketBillingInfo `json:"CosBucketBillingInfoSet,omitnil,omitempty" name:"CosBucketBillingInfoSet"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosBucketBillingInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosBucketBillingInfoResponseParams `json:"Response"`
}

func (r *DescribeCosBucketBillingInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosBucketBillingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosBucketListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosBucketListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosBucketListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosBucketListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosBucketListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosBucketListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 资产信息
	Data []*CosAssetInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosBucketListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosBucketListResponseParams `json:"Response"`
}

func (r *DescribeCosBucketListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosBucketListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosBucketRiskRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosBucketRiskRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosBucketRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosBucketRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosBucketRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosBucketRiskResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表
	Data []*CosRiskBucketInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosBucketRiskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosBucketRiskResponseParams `json:"Response"`
}

func (r *DescribeCosBucketRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosBucketRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosIdentifyFileListRequestParams struct {
	// <p>存储桶名</p>
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>筛选项</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>0：没有识别结果 1：有识别结果</p>
	ResultStatus *int64 `json:"ResultStatus,omitnil,omitempty" name:"ResultStatus"`
}

type DescribeCosIdentifyFileListRequest struct {
	*tchttp.BaseRequest
	
	// <p>存储桶名</p>
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>筛选项</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// <p>0：没有识别结果 1：有识别结果</p>
	ResultStatus *int64 `json:"ResultStatus,omitnil,omitempty" name:"ResultStatus"`
}

func (r *DescribeCosIdentifyFileListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosIdentifyFileListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BucketName")
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "ResultStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosIdentifyFileListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosIdentifyFileListResponseParams struct {
	// <p>总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>结果集</p>
	DataSet []*CosAssetFileIdentifyInfo `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosIdentifyFileListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosIdentifyFileListResponseParams `json:"Response"`
}

func (r *DescribeCosIdentifyFileListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosIdentifyFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosInvokeUaRequestParams struct {
	// appid
	RelAppId *uint64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosInvokeUaRequest struct {
	*tchttp.BaseRequest
	
	// appid
	RelAppId *uint64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosInvokeUaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosInvokeUaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosInvokeUaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosInvokeUaResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 文件列表
	Data []*string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosInvokeUaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosInvokeUaResponseParams `json:"Response"`
}

func (r *DescribeCosInvokeUaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosInvokeUaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosIpInvokeLogRequestParams struct {
	// appid
	RelAppId *uint64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosIpInvokeLogRequest struct {
	*tchttp.BaseRequest
	
	// appid
	RelAppId *uint64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosIpInvokeLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosIpInvokeLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosIpInvokeLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosIpInvokeLogResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 请求日志数据
	Data []*CosInvokeLog `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosIpInvokeLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosIpInvokeLogResponseParams `json:"Response"`
}

func (r *DescribeCosIpInvokeLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosIpInvokeLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosIpInvokeRecordFileRequestParams struct {
	// appid
	RelAppId *uint64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosIpInvokeRecordFileRequest struct {
	*tchttp.BaseRequest
	
	// appid
	RelAppId *uint64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosIpInvokeRecordFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosIpInvokeRecordFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosIpInvokeRecordFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosIpInvokeRecordFileResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 文件列表
	//
	// Deprecated: Data is deprecated.
	Data []*string `json:"Data,omitnil,omitempty" name:"Data"`

	// 文件列表详情
	DataSet []*CosAssetFileIdentifyInfo `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosIpInvokeRecordFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosIpInvokeRecordFileResponseParams `json:"Response"`
}

func (r *DescribeCosIpInvokeRecordFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosIpInvokeRecordFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosOverviewRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤信息
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤信息
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosOverviewResponseParams struct {
	// cos概览
	CosOverview *CosOverview `json:"CosOverview,omitnil,omitempty" name:"CosOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosOverviewResponseParams `json:"Response"`
}

func (r *DescribeCosOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosPolicyRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosPolicyResponseParams struct {
	// 策略总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 策略信息
	Data []*CosPolicyInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosPolicyResponseParams `json:"Response"`
}

func (r *DescribeCosPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRiskActionListRequestParams struct {
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosRiskActionListRequest struct {
	*tchttp.BaseRequest
	
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosRiskActionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRiskActionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "PolicyId")
	delete(f, "BucketName")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosRiskActionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRiskActionListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表
	Data []*CosRiskActionInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosRiskActionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosRiskActionListResponseParams `json:"Response"`
}

func (r *DescribeCosRiskActionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRiskActionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRiskEvidenceRequestParams struct {
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 存储桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosRiskEvidenceRequest struct {
	*tchttp.BaseRequest
	
	// appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 策略id
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 存储桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosRiskEvidenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRiskEvidenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "PolicyId")
	delete(f, "BucketName")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosRiskEvidenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRiskEvidenceResponseParams struct {
	// 证据信息
	Evidences []*CosPermissionInfo `json:"Evidences,omitnil,omitempty" name:"Evidences"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosRiskEvidenceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosRiskEvidenceResponseParams `json:"Response"`
}

func (r *DescribeCosRiskEvidenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRiskEvidenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRiskScanTaskRequestParams struct {
	// 需要查看的存储桶详情
	BucketInfoSet []*CosBucketInfo `json:"BucketInfoSet,omitnil,omitempty" name:"BucketInfoSet"`
}

type DescribeCosRiskScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 需要查看的存储桶详情
	BucketInfoSet []*CosBucketInfo `json:"BucketInfoSet,omitnil,omitempty" name:"BucketInfoSet"`
}

func (r *DescribeCosRiskScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRiskScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BucketInfoSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosRiskScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRiskScanTaskResponseParams struct {
	// cos桶任务详情
	BucketTaskInfoSet []*CosBucketTaskInfo `json:"BucketTaskInfoSet,omitnil,omitempty" name:"BucketTaskInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosRiskScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosRiskScanTaskResponseParams `json:"Response"`
}

func (r *DescribeCosRiskScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRiskScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRoleAccessPermissionRequestParams struct {
	// 关联的appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 需要查看的角色id
	RelRoleId *string `json:"RelRoleId,omitnil,omitempty" name:"RelRoleId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosRoleAccessPermissionRequest struct {
	*tchttp.BaseRequest
	
	// 关联的appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 需要查看的角色id
	RelRoleId *string `json:"RelRoleId,omitnil,omitempty" name:"RelRoleId"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosRoleAccessPermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRoleAccessPermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "RelRoleId")
	delete(f, "BucketName")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosRoleAccessPermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRoleAccessPermissionResponseParams struct {
	// cos权限信息
	Data []*CosPermissionInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosRoleAccessPermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosRoleAccessPermissionResponseParams `json:"Response"`
}

func (r *DescribeCosRoleAccessPermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRoleAccessPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRoleAccessPermissionsRequestParams struct {
	// 存储桶所属appid
	RelAppId *uint64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 存储桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosRoleAccessPermissionsRequest struct {
	*tchttp.BaseRequest
	
	// 存储桶所属appid
	RelAppId *uint64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 存储桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosRoleAccessPermissionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRoleAccessPermissionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "BucketName")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosRoleAccessPermissionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosRoleAccessPermissionsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 角色详情
	Data []*CosRoleAccessInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosRoleAccessPermissionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosRoleAccessPermissionsResponseParams `json:"Response"`
}

func (r *DescribeCosRoleAccessPermissionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosRoleAccessPermissionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosSourceIpRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCosSourceIpRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCosSourceIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosSourceIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosSourceIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosSourceIpResponseParams struct {
	// 列表信息
	Data []*CosSourceIpInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosSourceIpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosSourceIpResponseParams `json:"Response"`
}

func (r *DescribeCosSourceIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosSourceIpResponse) FromJsonString(s string) error {
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
type DescribeDspmAccessRecordRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 视图类型。ip或instance
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAccessRecordRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 视图类型。ip或instance
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAccessRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAccessRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "View")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAccessRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAccessRecordResponseParams struct {
	// 访问记录
	AccessSet []*DspmAccessRecord `json:"AccessSet,omitnil,omitempty" name:"AccessSet"`

	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAccessRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAccessRecordResponseParams `json:"Response"`
}

func (r *DescribeDspmAccessRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAccessRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAccessTopologyAccountsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器。 支持的FilterName: Ip/AssetId
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAccessTopologyAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器。 支持的FilterName: Ip/AssetId
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAccessTopologyAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAccessTopologyAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAccessTopologyAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAccessTopologyAccountsResponseParams struct {
	// 资产账号列表
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAccessTopologyAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAccessTopologyAccountsResponseParams `json:"Response"`
}

func (r *DescribeDspmAccessTopologyAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAccessTopologyAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAccessTopologyAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器。 支持的FilterName:  Ip/Account
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAccessTopologyAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器。 支持的FilterName:  Ip/Account
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAccessTopologyAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAccessTopologyAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAccessTopologyAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAccessTopologyAssetsResponseParams struct {
	// 资产id列表
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAccessTopologyAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAccessTopologyAssetsResponseParams `json:"Response"`
}

func (r *DescribeDspmAccessTopologyAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAccessTopologyAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAccessTopologyIpsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器。 支持的FilterName: AssetId/Account
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAccessTopologyIpsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器。 支持的FilterName: AssetId/Account
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAccessTopologyIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAccessTopologyIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAccessTopologyIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAccessTopologyIpsResponseParams struct {
	// ip列表
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAccessTopologyIpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAccessTopologyIpsResponseParams `json:"Response"`
}

func (r *DescribeDspmAccessTopologyIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAccessTopologyIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmApplyHistoryRequestParams struct {
	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmApplyHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmApplyHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmApplyHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmApplyHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmApplyHistoryResponseParams struct {
	// 申请记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 申请记录信息
	ApplySet []*DspmApplyOrder `json:"ApplySet,omitnil,omitempty" name:"ApplySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmApplyHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmApplyHistoryResponseParams `json:"Response"`
}

func (r *DescribeDspmApplyHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmApplyHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmApplyOrderListRequestParams struct {
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmApplyOrderListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmApplyOrderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmApplyOrderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmApplyOrderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmApplyOrderListResponseParams struct {
	// 申请单总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 申请单详情
	OrderSet []*DspmApplyOrder `json:"OrderSet,omitnil,omitempty" name:"OrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmApplyOrderListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmApplyOrderListResponseParams `json:"Response"`
}

func (r *DescribeDspmApplyOrderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmApplyOrderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmApproveHistoryRequestParams struct {
	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmApproveHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmApproveHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmApproveHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmApproveHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmApproveHistoryResponseParams struct {
	// 审批记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审批记录信息
	ApproveSet []*DspmApplyOrder `json:"ApproveSet,omitnil,omitempty" name:"ApproveSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmApproveHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmApproveHistoryResponseParams `json:"Response"`
}

func (r *DescribeDspmApproveHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmApproveHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmApproveOrderListRequestParams struct {
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmApproveOrderListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmApproveOrderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmApproveOrderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmApproveOrderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmApproveOrderListResponseParams struct {
	// 审批单总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审批单详情
	OrderSet []*DspmApproverOrder `json:"OrderSet,omitnil,omitempty" name:"OrderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmApproveOrderListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmApproveOrderListResponseParams `json:"Response"`
}

func (r *DescribeDspmApproveOrderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmApproveOrderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccessTopologyRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 视图类型。ip或instance
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// - 来源ip方式查看
	// View: "ip"
	// 
	// Filter:{
	// 	{
	// 		"Name":"Ip",
	// 		"Values":["172.1.1.1"]
	// 	},
	// 	{
	// 		"Name":"AssetId",
	// 		"Values":["cdb-1111|ap-guangzhou","cdb-2222|ap-guangzhou","cdb-3333|ap-guangzhou"]
	// 	},
	// 	{
	// 		"Name":"Account",
	// 		"Values":["root|%","test|%"]
	// 	}
	// }
	// 
	// 
	// - 实例方式查看
	// View: "instance"
	// 
	// Filter:{
	// 	{
	// 		"Name":"AssetId",
	// 		"Values":["cdb-1111|ap-guangzhou"]
	// 	},
	// 	{
	// 		"Name":"Ip",
	// 		"Values":["172.1.1.1","172.1.1.2","172.1.1.3"]
	// 	},
	// 	{
	// 		"Name":"Account",
	// 		"Values":["root|%","test|%"]
	// 	}
	// }
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetAccessTopologyRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 视图类型。ip或instance
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// - 来源ip方式查看
	// View: "ip"
	// 
	// Filter:{
	// 	{
	// 		"Name":"Ip",
	// 		"Values":["172.1.1.1"]
	// 	},
	// 	{
	// 		"Name":"AssetId",
	// 		"Values":["cdb-1111|ap-guangzhou","cdb-2222|ap-guangzhou","cdb-3333|ap-guangzhou"]
	// 	},
	// 	{
	// 		"Name":"Account",
	// 		"Values":["root|%","test|%"]
	// 	}
	// }
	// 
	// 
	// - 实例方式查看
	// View: "instance"
	// 
	// Filter:{
	// 	{
	// 		"Name":"AssetId",
	// 		"Values":["cdb-1111|ap-guangzhou"]
	// 	},
	// 	{
	// 		"Name":"Ip",
	// 		"Values":["172.1.1.1","172.1.1.2","172.1.1.3"]
	// 	},
	// 	{
	// 		"Name":"Account",
	// 		"Values":["root|%","test|%"]
	// 	}
	// }
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetAccessTopologyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccessTopologyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "View")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetAccessTopologyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccessTopologyResponseParams struct {
	// 拓扑数据
	ItemSet []*DspmAssetAccessTopologyItem `json:"ItemSet,omitnil,omitempty" name:"ItemSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetAccessTopologyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetAccessTopologyResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetAccessTopologyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccessTopologyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccountIdentifyRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetAccountIdentifyRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetAccountIdentifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccountIdentifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "AssetId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetAccountIdentifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccountIdentifyResponseParams struct {
	// 资产账号身份总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产账号身份信息
	IdentifySet []*DspmAssetAccountIdentify `json:"IdentifySet,omitnil,omitempty" name:"IdentifySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetAccountIdentifyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetAccountIdentifyResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetAccountIdentifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccountIdentifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccountPresetPrivilegesRequestParams struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetAccountPresetPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetAccountPresetPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccountPresetPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "Account")
	delete(f, "Host")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetAccountPresetPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccountPresetPrivilegesResponseParams struct {
	// 权限信息
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetAccountPresetPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetAccountPresetPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetAccountPresetPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccountPresetPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccountRecycledPrivilegesRequestParams struct {
	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

type DescribeDspmAssetAccountRecycledPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

func (r *DescribeDspmAssetAccountRecycledPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccountRecycledPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetAccountRecycledPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccountRecycledPrivilegesResponseParams struct {
	// 权限信息
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetAccountRecycledPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetAccountRecycledPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetAccountRecycledPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccountRecycledPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccountsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "AssetId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetAccountsResponseParams struct {
	// 资产账号总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 账号信息
	AccountSet []*DspmAssetAccount `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetAccountsResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetDatabaseListRequestParams struct {
	// <p>资产实例id</p>
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>筛选项</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetDatabaseListRequest struct {
	*tchttp.BaseRequest
	
	// <p>资产实例id</p>
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>筛选项</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetDatabaseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetDatabaseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetDatabaseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetDatabaseListResponseParams struct {
	// <p>总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>结果集</p>
	DataSet []*DspmAssetDatabaseInfo `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetDatabaseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetDatabaseListResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetDatabaseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetDatabaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetDatabasesRequestParams struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type DescribeDspmAssetDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *DescribeDspmAssetDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetDatabasesResponseParams struct {
	// 数据库列表
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetDatabasesResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetFieldListRequestParams struct {
	// 资产实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetFieldListRequest struct {
	*tchttp.BaseRequest
	
	// 资产实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetFieldListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetFieldListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "DbName")
	delete(f, "TableName")
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetFieldListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetFieldListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 结果集
	DataSet []*DspmAssetFieldInfo `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetFieldListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetFieldListResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetFieldListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetFieldListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetIdsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetIdsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetIdsResponseParams struct {
	// 数据库资产总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产id信息
	AssetSet []*DspmDbAssetId `json:"AssetSet,omitnil,omitempty" name:"AssetSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetIdsResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetLoginCredentialRequestParams struct {
	// 数据库资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 主机。默认'%'
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 加密算法
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`
}

type DescribeDspmAssetLoginCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 数据库资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 主机。默认'%'
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 加密算法
	EncryptMethod *string `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`
}

func (r *DescribeDspmAssetLoginCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetLoginCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "Host")
	delete(f, "EncryptMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetLoginCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetLoginCredentialResponseParams struct {
	// 账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 密码信息
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 有效期开始时间
	ValidateStart *string `json:"ValidateStart,omitnil,omitempty" name:"ValidateStart"`

	// 有效期结束时间
	ValidateEnd *string `json:"ValidateEnd,omitnil,omitempty" name:"ValidateEnd"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetLoginCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetLoginCredentialResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetLoginCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetLoginCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetSecurityAnalyseStatusRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetSecurityAnalyseStatusRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetSecurityAnalyseStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetSecurityAnalyseStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetSecurityAnalyseStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetSecurityAnalyseStatusResponseParams struct {
	// 数据库资产总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产安全分析状态信息
	AssetSet []*DspmAssetSecurityAnalyseStatus `json:"AssetSet,omitnil,omitempty" name:"AssetSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetSecurityAnalyseStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetSecurityAnalyseStatusResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetSecurityAnalyseStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetSecurityAnalyseStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetSupportedPrivilegesRequestParams struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type DescribeDspmAssetSupportedPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *DescribeDspmAssetSupportedPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetSupportedPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetSupportedPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetSupportedPrivilegesResponseParams struct {
	// 实例支持的全局权限。
	GlobalSupportedPrivileges []*string `json:"GlobalSupportedPrivileges,omitnil,omitempty" name:"GlobalSupportedPrivileges"`

	// 实例支持的数据库权限。
	DatabaseSupportedPrivileges []*string `json:"DatabaseSupportedPrivileges,omitnil,omitempty" name:"DatabaseSupportedPrivileges"`

	// 实例支持的数据库表权限。
	TableSupportedPrivileges []*string `json:"TableSupportedPrivileges,omitnil,omitempty" name:"TableSupportedPrivileges"`

	// 实例支持的数据库列权限。
	ColumnSupportedPrivileges []*string `json:"ColumnSupportedPrivileges,omitnil,omitempty" name:"ColumnSupportedPrivileges"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetSupportedPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetSupportedPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetSupportedPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetSupportedPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetTableListRequestParams struct {
	// 资产实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetTableListRequest struct {
	*tchttp.BaseRequest
	
	// 资产实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetTableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetTableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "DbName")
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetTableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetTableListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 结果集
	DataSet []*DspmAssetTableInfo `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetTableListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetTableListResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetTableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetTableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmAssetsResponseParams struct {
	// 数据库资产总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产信息
	AssetSet []*DspmDbAsset `json:"AssetSet,omitnil,omitempty" name:"AssetSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmAssetsResponseParams `json:"Response"`
}

func (r *DescribeDspmAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmBackupLogListRequestParams struct {
	// <p>限制数目</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>排序方式(desc=倒叙,asc=升序)</p>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>排序字段(支持&#39;StartTime&#39;)</p>
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>备份日志状态0未完成的,1备份文件，2恢复中，3已恢复，4.已删除,全部查询-1</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>数据库类型,如：cdb, mariadb</p>
	DbTypes []*string `json:"DbTypes,omitnil,omitempty" name:"DbTypes"`
}

type DescribeDspmBackupLogListRequest struct {
	*tchttp.BaseRequest
	
	// <p>限制数目</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>排序方式(desc=倒叙,asc=升序)</p>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>排序字段(支持&#39;StartTime&#39;)</p>
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// <p>开始时间</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>备份日志状态0未完成的,1备份文件，2恢复中，3已恢复，4.已删除,全部查询-1</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>数据库类型,如：cdb, mariadb</p>
	DbTypes []*string `json:"DbTypes,omitnil,omitempty" name:"DbTypes"`
}

func (r *DescribeDspmBackupLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmBackupLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "MemberId")
	delete(f, "Sort")
	delete(f, "Field")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Status")
	delete(f, "DbTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmBackupLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmBackupLogListResponseParams struct {
	// <p>总共多少条</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>备份日志列表</p>
	List []*BackupLog `json:"List,omitnil,omitempty" name:"List"`

	// <p>当前是否存在恢复中任务</p>
	HasRestoringTask *bool `json:"HasRestoringTask,omitnil,omitempty" name:"HasRestoringTask"`

	// <p>最大恢复空间</p>
	MaxRestoreSizeInGB *int64 `json:"MaxRestoreSizeInGB,omitnil,omitempty" name:"MaxRestoreSizeInGB"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmBackupLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmBackupLogListResponseParams `json:"Response"`
}

func (r *DescribeDspmBackupLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmBackupLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmBackupSettingRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeDspmBackupSettingRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeDspmBackupSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmBackupSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmBackupSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmBackupSettingResponseParams struct {
	// 备份日志保留时长
	BackupLogSaveTime *int64 `json:"BackupLogSaveTime,omitnil,omitempty" name:"BackupLogSaveTime"`

	// 恢复日志保留时长
	RestoreLogSaveTime *int64 `json:"RestoreLogSaveTime,omitnil,omitempty" name:"RestoreLogSaveTime"`

	// 日志最大生命周期限制
	LogMaxSaveTime *int64 `json:"LogMaxSaveTime,omitnil,omitempty" name:"LogMaxSaveTime"`

	// 在线日志最大天数限制
	OnlineLogMaxSaveTime *int64 `json:"OnlineLogMaxSaveTime,omitnil,omitempty" name:"OnlineLogMaxSaveTime"`

	// 最大在线日志条数，单位是：个
	MaxOnlineLogCount *int64 `json:"MaxOnlineLogCount,omitnil,omitempty" name:"MaxOnlineLogCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmBackupSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmBackupSettingResponseParams `json:"Response"`
}

func (r *DescribeDspmBackupSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmBackupSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmDictionaryListRequestParams struct {
	// <p>字典类型（RootCategory：一级分类，IdentifyRule:敏感识别数据项）</p>
	DictType *string `json:"DictType,omitnil,omitempty" name:"DictType"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>筛选条件</p>
	Filters []*WhereFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDspmDictionaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>字典类型（RootCategory：一级分类，IdentifyRule:敏感识别数据项）</p>
	DictType *string `json:"DictType,omitnil,omitempty" name:"DictType"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>筛选条件</p>
	Filters []*WhereFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDspmDictionaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmDictionaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DictType")
	delete(f, "MemberId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmDictionaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmDictionaryListResponseParams struct {
	// <p>结果集</p>
	DataSet []*DspmDictionary `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmDictionaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmDictionaryListResponseParams `json:"Response"`
}

func (r *DescribeDspmDictionaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmDictionaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmExportTaskRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 任务状态：0.未开始 1.执行中 2.执行成功 3.执行超时 4.执行失败
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeDspmExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 任务状态：0.未开始 1.执行中 2.执行成功 3.执行超时 4.执行失败
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeDspmExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "TaskStatus")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmExportTaskResponseParams struct {
	// 任务列表
	List []*ExportTask `json:"List,omitnil,omitempty" name:"List"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmExportTaskResponseParams `json:"Response"`
}

func (r *DescribeDspmExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmIdentifyIdListRequestParams struct {
	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmIdentifyIdListRequest struct {
	*tchttp.BaseRequest
	
	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmIdentifyIdListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmIdentifyIdListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmIdentifyIdListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmIdentifyIdListResponseParams struct {
	// id总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 身份id列表
	IdSet []*DspmIdentifyIdItem `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmIdentifyIdListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmIdentifyIdListResponseParams `json:"Response"`
}

func (r *DescribeDspmIdentifyIdListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmIdentifyIdListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmIdentifyInfoListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmIdentifyInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmIdentifyInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmIdentifyInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmIdentifyInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmIdentifyInfoListResponseParams struct {
	// 身份总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 身份 信息
	InfoSet []*DspmIdentifyInfoItem `json:"InfoSet,omitnil,omitempty" name:"InfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmIdentifyInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmIdentifyInfoListResponseParams `json:"Response"`
}

func (r *DescribeDspmIdentifyInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmIdentifyInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmIdentifyInfoRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeDspmIdentifyInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeDspmIdentifyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmIdentifyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmIdentifyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmIdentifyInfoResponseParams struct {
	// 身份id
	IdentifyId *string `json:"IdentifyId,omitnil,omitempty" name:"IdentifyId"`

	// 身份统计信息
	IdentifyCount []*DspmIdentifyCount `json:"IdentifyCount,omitnil,omitempty" name:"IdentifyCount"`

	// 申请单个数
	ApplyOrderCount *int64 `json:"ApplyOrderCount,omitnil,omitempty" name:"ApplyOrderCount"`

	// 审批单个数
	ApproveOrderCount *int64 `json:"ApproveOrderCount,omitnil,omitempty" name:"ApproveOrderCount"`

	// 已审批个数
	ApproveHistoryCount *int64 `json:"ApproveHistoryCount,omitnil,omitempty" name:"ApproveHistoryCount"`

	// 资产总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// 云账号总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UinAccountCount *int64 `json:"UinAccountCount,omitnil,omitempty" name:"UinAccountCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmIdentifyInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmIdentifyInfoResponseParams `json:"Response"`
}

func (r *DescribeDspmIdentifyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmIdentifyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmLogListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序方式(desc=倒叙,asc=升序)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序字段(opTime=时间,dangerLvl=风险等级)
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 风险等级(0-安全,1-低风险,2-中风险,3-高风险,不传全部)
	DangerLevel *string `json:"DangerLevel,omitnil,omitempty" name:"DangerLevel"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 数据库端口
	DbPort *int64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// 数据库 IP
	DbIp *string `json:"DbIp,omitnil,omitempty" name:"DbIp"`

	// 资产 ID
	AssetsId *int64 `json:"AssetsId,omitnil,omitempty" name:"AssetsId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 客户端 IP
	ClientSideIp *string `json:"ClientSideIp,omitnil,omitempty" name:"ClientSideIp"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 命中规则
	HitRule *int64 `json:"HitRule,omitnil,omitempty" name:"HitRule"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 模糊查询
	FuzzySearch *string `json:"FuzzySearch,omitnil,omitempty" name:"FuzzySearch"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 恢复日志id
	RestoreLogId *int64 `json:"RestoreLogId,omitnil,omitempty" name:"RestoreLogId"`

	// 客户端
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// 流量来源，取值 Agent/Proxy/空；传Agent会返回Agent的日志，传Proxy会返回Proxy日志，两者都传或不传则返回所有
	SourceTypes []*string `json:"SourceTypes,omitnil,omitempty" name:"SourceTypes"`

	// 表名，长度限制64，多个表名查询的话可以用空格连接
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 字段名，长度限制64，多个字段名查询的话可以用空格连接
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// SQL 主要类型，DDL, DML, DCL, TCL
	SqlMainTypes []*string `json:"SqlMainTypes,omitnil,omitempty" name:"SqlMainTypes"`

	// 操作类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// 影响行数最小值 
	RowNumMin *int64 `json:"RowNumMin,omitnil,omitempty" name:"RowNumMin"`

	// 影响行数最大值
	RowNumMax *int64 `json:"RowNumMax,omitnil,omitempty" name:"RowNumMax"`

	// 数据库类型
	DbTypes []*string `json:"DbTypes,omitnil,omitempty" name:"DbTypes"`

	// 返回码
	RetNo *int64 `json:"RetNo,omitnil,omitempty" name:"RetNo"`

	// 客户端工具
	ClientDriverName *string `json:"ClientDriverName,omitnil,omitempty" name:"ClientDriverName"`

	// 客户端端口
	ClientPort *int64 `json:"ClientPort,omitnil,omitempty" name:"ClientPort"`

	// 审计日志 ID
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 风险等级数组(0-安全,1-低风险,2-中风险,3-高风险)
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`

	// 字段分类
	SensitiveCategoryRule *string `json:"SensitiveCategoryRule,omitnil,omitempty" name:"SensitiveCategoryRule"`

	// 字段分级
	SensitiveLevelRisk *string `json:"SensitiveLevelRisk,omitnil,omitempty" name:"SensitiveLevelRisk"`

	// 客户端MAC
	ClientMac *string `json:"ClientMac,omitnil,omitempty" name:"ClientMac"`
}

type DescribeDspmLogListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序方式(desc=倒叙,asc=升序)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序字段(opTime=时间,dangerLvl=风险等级)
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 风险等级(0-安全,1-低风险,2-中风险,3-高风险,不传全部)
	DangerLevel *string `json:"DangerLevel,omitnil,omitempty" name:"DangerLevel"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 数据库端口
	DbPort *int64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// 数据库 IP
	DbIp *string `json:"DbIp,omitnil,omitempty" name:"DbIp"`

	// 资产 ID
	AssetsId *int64 `json:"AssetsId,omitnil,omitempty" name:"AssetsId"`

	// 会话 ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 客户端 IP
	ClientSideIp *string `json:"ClientSideIp,omitnil,omitempty" name:"ClientSideIp"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 命中规则
	HitRule *int64 `json:"HitRule,omitnil,omitempty" name:"HitRule"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 模糊查询
	FuzzySearch *string `json:"FuzzySearch,omitnil,omitempty" name:"FuzzySearch"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 恢复日志id
	RestoreLogId *int64 `json:"RestoreLogId,omitnil,omitempty" name:"RestoreLogId"`

	// 客户端
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// 流量来源，取值 Agent/Proxy/空；传Agent会返回Agent的日志，传Proxy会返回Proxy日志，两者都传或不传则返回所有
	SourceTypes []*string `json:"SourceTypes,omitnil,omitempty" name:"SourceTypes"`

	// 表名，长度限制64，多个表名查询的话可以用空格连接
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 字段名，长度限制64，多个字段名查询的话可以用空格连接
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// SQL 主要类型，DDL, DML, DCL, TCL
	SqlMainTypes []*string `json:"SqlMainTypes,omitnil,omitempty" name:"SqlMainTypes"`

	// 操作类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// 影响行数最小值 
	RowNumMin *int64 `json:"RowNumMin,omitnil,omitempty" name:"RowNumMin"`

	// 影响行数最大值
	RowNumMax *int64 `json:"RowNumMax,omitnil,omitempty" name:"RowNumMax"`

	// 数据库类型
	DbTypes []*string `json:"DbTypes,omitnil,omitempty" name:"DbTypes"`

	// 返回码
	RetNo *int64 `json:"RetNo,omitnil,omitempty" name:"RetNo"`

	// 客户端工具
	ClientDriverName *string `json:"ClientDriverName,omitnil,omitempty" name:"ClientDriverName"`

	// 客户端端口
	ClientPort *int64 `json:"ClientPort,omitnil,omitempty" name:"ClientPort"`

	// 审计日志 ID
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 风险等级数组(0-安全,1-低风险,2-中风险,3-高风险)
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`

	// 字段分类
	SensitiveCategoryRule *string `json:"SensitiveCategoryRule,omitnil,omitempty" name:"SensitiveCategoryRule"`

	// 字段分级
	SensitiveLevelRisk *string `json:"SensitiveLevelRisk,omitnil,omitempty" name:"SensitiveLevelRisk"`

	// 客户端MAC
	ClientMac *string `json:"ClientMac,omitnil,omitempty" name:"ClientMac"`
}

func (r *DescribeDspmLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Sort")
	delete(f, "Field")
	delete(f, "DangerLevel")
	delete(f, "DbName")
	delete(f, "DbPort")
	delete(f, "DbIp")
	delete(f, "AssetsId")
	delete(f, "SessionId")
	delete(f, "ClientSideIp")
	delete(f, "EndTime")
	delete(f, "HitRule")
	delete(f, "StartTime")
	delete(f, "FuzzySearch")
	delete(f, "UserName")
	delete(f, "RestoreLogId")
	delete(f, "ClientName")
	delete(f, "SourceTypes")
	delete(f, "TableName")
	delete(f, "FieldName")
	delete(f, "SqlMainTypes")
	delete(f, "SqlType")
	delete(f, "RowNumMin")
	delete(f, "RowNumMax")
	delete(f, "DbTypes")
	delete(f, "RetNo")
	delete(f, "ClientDriverName")
	delete(f, "ClientPort")
	delete(f, "LogId")
	delete(f, "DangerLevels")
	delete(f, "SensitiveCategoryRule")
	delete(f, "SensitiveLevelRisk")
	delete(f, "ClientMac")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmLogListResponseParams struct {
	// 总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志信息列表
	List []*AuditLogInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmLogListResponseParams `json:"Response"`
}

func (r *DescribeDspmLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmPayInfoRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeDspmPayInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeDspmPayInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmPayInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmPayInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmPayInfoResponseParams struct {
	// APPID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 订单状态 0未购买 1正常，2隔离，3销毁，6试用中，7到期
	OrderStatus *int64 `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// 已购数据库实例数量
	AssetNum *int64 `json:"AssetNum,omitnil,omitempty" name:"AssetNum"`

	// 已购审计日志量（TB）
	LogStorage *float64 `json:"LogStorage,omitnil,omitempty" name:"LogStorage"`

	// 已使用数据库实例数量
	UsedAssetNum *int64 `json:"UsedAssetNum,omitnil,omitempty" name:"UsedAssetNum"`

	// 已使用审计日志量（TB）
	UsedLogStorage *float64 `json:"UsedLogStorage,omitnil,omitempty" name:"UsedLogStorage"`

	// 已购sql存储总量（单位百万）
	SqlTotal *uint64 `json:"SqlTotal,omitnil,omitempty" name:"SqlTotal"`

	// 已购sql qps
	SqlQps *uint64 `json:"SqlQps,omitnil,omitempty" name:"SqlQps"`

	// 支付模式，0-后付费 1-预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否单独购买，1-单独购买，2-被其它账号共享
	IsSelfBuy *int64 `json:"IsSelfBuy,omitnil,omitempty" name:"IsSelfBuy"`

	// 订单开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 订单到期时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 0-用户未设置,1-用户设置自动续费,2-用户设置不自动续费
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 订单时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 公测结束时间
	BetaEndTime *string `json:"BetaEndTime,omitnil,omitempty" name:"BetaEndTime"`

	// 系统当前时间
	TimeNow *string `json:"TimeNow,omitnil,omitempty" name:"TimeNow"`

	// 是否分享给其它账号，1-是，2-否
	IsShareToOther *int64 `json:"IsShareToOther,omitnil,omitempty" name:"IsShareToOther"`

	// uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 购买信息
	InquireData []*InquireInfo `json:"InquireData,omitnil,omitempty" name:"InquireData"`

	// 版本(专业版：professional 试用版：trial)
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmPayInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmPayInfoResponseParams `json:"Response"`
}

func (r *DescribeDspmPayInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmPayInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmPersonApplyHistoryRequestParams struct {
	// 对象
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmPersonApplyHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 对象
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmPersonApplyHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmPersonApplyHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Subject")
	delete(f, "AssetId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmPersonApplyHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmPersonApplyHistoryResponseParams struct {
	// 申请记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 申请记录信息
	ApplySet []*DspmPersonApplyHistoryItem `json:"ApplySet,omitnil,omitempty" name:"ApplySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmPersonApplyHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmPersonApplyHistoryResponseParams `json:"Response"`
}

func (r *DescribeDspmPersonApplyHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmPersonApplyHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmPersonalIdentifyListRequestParams struct {
	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmPersonalIdentifyListRequest struct {
	*tchttp.BaseRequest
	
	// 筛选项
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmPersonalIdentifyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmPersonalIdentifyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmPersonalIdentifyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmPersonalIdentifyListResponseParams struct {
	// 身份总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 个人身份信息
	InfoSet []*DspmPersonIdentifyItem `json:"InfoSet,omitnil,omitempty" name:"InfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmPersonalIdentifyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmPersonalIdentifyListResponseParams `json:"Response"`
}

func (r *DescribeDspmPersonalIdentifyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmPersonalIdentifyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskDetailRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

type DescribeDspmRiskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

func (r *DescribeDspmRiskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	delete(f, "RiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmRiskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskDetailResponseParams struct {
	// 风险等级
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 风险检出时间
	DetectTime *string `json:"DetectTime,omitnil,omitempty" name:"DetectTime"`

	// 资产实例Id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 公网访问地址
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网访问地址
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 身份id
	IdentifyId *string `json:"IdentifyId,omitnil,omitempty" name:"IdentifyId"`

	// 所属云账号uin用户
	OwnerUin *DspmUinUser `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 所属个人用户信息
	Person *DspmPersonUser `json:"Person,omitnil,omitempty" name:"Person"`

	// 风险名称
	RiskName *string `json:"RiskName,omitnil,omitempty" name:"RiskName"`

	// 风险英文名称
	RiskNameEn *string `json:"RiskNameEn,omitnil,omitempty" name:"RiskNameEn"`

	// 风险数据
	RiskData *string `json:"RiskData,omitnil,omitempty" name:"RiskData"`

	// 基线数据
	BaselineData *string `json:"BaselineData,omitnil,omitempty" name:"BaselineData"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 策略类别
	StrategyCategory *string `json:"StrategyCategory,omitnil,omitempty" name:"StrategyCategory"`

	// 账号类型
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 风险状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否资产管理员
	IsAssetManager *int64 `json:"IsAssetManager,omitnil,omitempty" name:"IsAssetManager"`

	// 数据起始时间
	DataBeginTime *string `json:"DataBeginTime,omitnil,omitempty" name:"DataBeginTime"`

	// 数据结束时间
	DataEndTime *string `json:"DataEndTime,omitnil,omitempty" name:"DataEndTime"`

	// 风险类型。risk-风险；alarm-告警。
	RiskType *string `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmRiskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmRiskDetailResponseParams `json:"Response"`
}

func (r *DescribeDspmRiskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmRiskRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskResponseParams struct {
	// 风险列表
	RiskSet []*DspmRisk `json:"RiskSet,omitnil,omitempty" name:"RiskSet"`

	// 风险总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmRiskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmRiskResponseParams `json:"Response"`
}

func (r *DescribeDspmRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskStrategyGroupRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmRiskStrategyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmRiskStrategyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskStrategyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmRiskStrategyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskStrategyGroupResponseParams struct {
	// 分组策略列表
	StrategyGroupSet []*DspmRiskStrategyGroup `json:"StrategyGroupSet,omitnil,omitempty" name:"StrategyGroupSet"`

	// 分组策略总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmRiskStrategyGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmRiskStrategyGroupResponseParams `json:"Response"`
}

func (r *DescribeDspmRiskStrategyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskStrategyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskStrategyRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmRiskStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmRiskStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmRiskStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskStrategyResponseParams struct {
	// 策略列表
	StrategySet []*DspmRiskStrategy `json:"StrategySet,omitnil,omitempty" name:"StrategySet"`

	// 策略总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmRiskStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmRiskStrategyResponseParams `json:"Response"`
}

func (r *DescribeDspmRiskStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskTendencyRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 起始日期
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DescribeDspmRiskTendencyRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 起始日期
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *DescribeDspmRiskTendencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskTendencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmRiskTendencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmRiskTendencyResponseParams struct {
	// 风险趋势
	RiskTendencySet []*DspmRiskTendency `json:"RiskTendencySet,omitnil,omitempty" name:"RiskTendencySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmRiskTendencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmRiskTendencyResponseParams `json:"Response"`
}

func (r *DescribeDspmRiskTendencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmRiskTendencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmStatisticsRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>按照资产类型过滤</p><p>枚举值：</p><ul><li>cdb： cdb</li><li>mariadb： mariadb</li><li>cynosdb： cynosdb</li></ul><p>默认值：默认值为空，即不按照资产类型过滤，返回所有资产信息</p>
	AssetType []*string `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

type DescribeDspmStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>按照资产类型过滤</p><p>枚举值：</p><ul><li>cdb： cdb</li><li>mariadb： mariadb</li><li>cynosdb： cynosdb</li></ul><p>默认值：默认值为空，即不按照资产类型过滤，返回所有资产信息</p>
	AssetType []*string `json:"AssetType,omitnil,omitempty" name:"AssetType"`
}

func (r *DescribeDspmStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "AssetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmStatisticsResponseParams struct {
	// <p>资产统计信息</p>
	AssetCount *DspmAssetCount `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// <p>访问Ip统计信息</p>
	IpCount *DspmIpCount `json:"IpCount,omitnil,omitempty" name:"IpCount"`

	// <p>用户账号统计信息</p>
	UserCount *DspmAccountCount `json:"UserCount,omitnil,omitempty" name:"UserCount"`

	// <p>风险统计信息</p>
	RiskCount *DspmRiskCount `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// <p>资产安全分析统计信息</p>
	AnalyseAssetStatusCount *DspmSecurityAnalyseStatusCount `json:"AnalyseAssetStatusCount,omitnil,omitempty" name:"AnalyseAssetStatusCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmStatisticsResponseParams `json:"Response"`
}

func (r *DescribeDspmStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmSupportedAssetTypeRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeDspmSupportedAssetTypeRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeDspmSupportedAssetTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmSupportedAssetTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmSupportedAssetTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmSupportedAssetTypeResponseParams struct {
	// <p>实例支持的全局权限。</p>
	AssetTypeSet []*DspmSupportedAssetType `json:"AssetTypeSet,omitnil,omitempty" name:"AssetTypeSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmSupportedAssetTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmSupportedAssetTypeResponseParams `json:"Response"`
}

func (r *DescribeDspmSupportedAssetTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmSupportedAssetTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmSyncAssetsStatusRequestParams struct {

}

type DescribeDspmSyncAssetsStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDspmSyncAssetsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmSyncAssetsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmSyncAssetsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmSyncAssetsStatusResponseParams struct {
	// 资产同步任务状态。0-未执行。1-执行中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmSyncAssetsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmSyncAssetsStatusResponseParams `json:"Response"`
}

func (r *DescribeDspmSyncAssetsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmSyncAssetsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmSyncUsersStatusRequestParams struct {

}

type DescribeDspmSyncUsersStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDspmSyncUsersStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmSyncUsersStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmSyncUsersStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmSyncUsersStatusResponseParams struct {
	// 用户同步任务状态。0-未执行。1-执行中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmSyncUsersStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmSyncUsersStatusResponseParams `json:"Response"`
}

func (r *DescribeDspmSyncUsersStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmSyncUsersStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmWhitelistStrategyRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDspmWhitelistStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDspmWhitelistStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmWhitelistStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDspmWhitelistStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDspmWhitelistStrategyResponseParams struct {
	// 白名单列表
	WhitelistSet []*DspmWhitelistStrategy `json:"WhitelistSet,omitnil,omitempty" name:"WhitelistSet"`

	// 白名单总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDspmWhitelistStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDspmWhitelistStrategyResponseParams `json:"Response"`
}

func (r *DescribeDspmWhitelistStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDspmWhitelistStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEDRRuleListRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>PolicyType - int - 是否必填：否 - 策略类型<br>PolicyName - string - 是否必填：否 - 策略名称<br>Domain - string - 是否必填：否 - 域名(先对域名做urlencode,再base64)<br>PolicyAction- int - 是否必填：否 - 策略动作<br>IsEnabled - int - 是否必填：否 - 是否生效</p>
	Filters []*EDRFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>限制条数,默认10,最大100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量,默认0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序方式: [ASC:升序|DESC:降序]</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>可选排序列: [ModifyTime]</p>
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeEDRRuleListRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>PolicyType - int - 是否必填：否 - 策略类型<br>PolicyName - string - 是否必填：否 - 策略名称<br>Domain - string - 是否必填：否 - 域名(先对域名做urlencode,再base64)<br>PolicyAction- int - 是否必填：否 - 策略动作<br>IsEnabled - int - 是否必填：否 - 是否生效</p>
	Filters []*EDRFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>限制条数,默认10,最大100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量,默认0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序方式: [ASC:升序|DESC:降序]</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>可选排序列: [ModifyTime]</p>
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeEDRRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEDRRuleListRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEDRRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEDRRuleListResponseParams struct {
	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>列表</p>
	List []*EDRRule `json:"List,omitnil,omitempty" name:"List"`

	// <p>攻击阶段对应的策略数量</p>
	AttackStageCounts []*AttackStageCount `json:"AttackStageCounts,omitnil,omitempty" name:"AttackStageCounts"`

	// <p>检测方式对应的策略数量</p>
	DetectTypeCounts []*DetectTypeCount `json:"DetectTypeCounts,omitnil,omitempty" name:"DetectTypeCounts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEDRRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEDRRuleListResponseParams `json:"Response"`
}

func (r *DescribeEDRRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEDRRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdrAlertInfoRequestParams struct {
	// <p>告警定位信息（含跨账号AppID）</p>
	Target *EdrAlertTarget `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeEdrAlertInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>告警定位信息（含跨账号AppID）</p>
	Target *EdrAlertTarget `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeEdrAlertInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdrAlertInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Target")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdrAlertInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdrAlertInfoResponseParams struct {
	// <p>告警详情</p>
	Alert *EdrAlertDetail `json:"Alert,omitnil,omitempty" name:"Alert"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdrAlertInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdrAlertInfoResponseParams `json:"Response"`
}

func (r *DescribeEdrAlertInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdrAlertInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdrAlertListRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>PolicyType - int - 是否必填：否 - 策略类型PolicyName - string - 是否必填：否 - 策略名称Domain - string - 是否必填：否 - 域名(先对域名做urlencode,再base64)PolicyAction- int - 是否必填：否 - 策略动作IsEnabled - int - 是否必填：否 - 是否生效</p>
	Filters []*EDRFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>限制条数,默认10,最大100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量,默认0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序方式: [ASC:升序|DESC:降序]</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>可选排序列: [LatestDetectTime]</p>
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeEdrAlertListRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>PolicyType - int - 是否必填：否 - 策略类型PolicyName - string - 是否必填：否 - 策略名称Domain - string - 是否必填：否 - 域名(先对域名做urlencode,再base64)PolicyAction- int - 是否必填：否 - 策略动作IsEnabled - int - 是否必填：否 - 是否生效</p>
	Filters []*EDRFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>限制条数,默认10,最大100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量,默认0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序方式: [ASC:升序|DESC:降序]</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>可选排序列: [LatestDetectTime]</p>
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeEdrAlertListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdrAlertListRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdrAlertListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdrAlertListResponseParams struct {
	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>列表</p>
	List []*EdrAlertItem `json:"List,omitnil,omitempty" name:"List"`

	// <p>攻击阶段对应的策略数量</p>
	AttackStageCounts []*AttackStageCount `json:"AttackStageCounts,omitnil,omitempty" name:"AttackStageCounts"`

	// <p>告警大类统计（随筛选变化，排除 AlertCategory filter）</p>
	AlertCategoryCounts []*EdrAlertCategoryCount `json:"AlertCategoryCounts,omitnil,omitempty" name:"AlertCategoryCounts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdrAlertListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdrAlertListResponseParams `json:"Response"`
}

func (r *DescribeEdrAlertListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdrAlertListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExposeAssetCategoryRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeExposeAssetCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
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
	// <p>集团账号的成员id</p>
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
	
	// <p>集团账号的成员id</p>
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
	// <p>集团账号的成员id</p>
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
	
	// <p>集团账号的成员id</p>
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
	// <p>集团账号的成员id</p>
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
	
	// <p>集团账号的成员id</p>
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
type DescribeIaCFileListRequestParams struct {
	// <p>过滤条件</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeIaCFileListRequest struct {
	*tchttp.BaseRequest
	
	// <p>过滤条件</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeIaCFileListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIaCFileListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileListResponseParams struct {
	// <p>列表</p>
	List []*IaCFile `json:"List,omitnil,omitempty" name:"List"`

	// <p>总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIaCFileListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIaCFileListResponseParams `json:"Response"`
}

func (r *DescribeIaCFileListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileOverviewRequestParams struct {
	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeIaCFileOverviewRequest struct {
	*tchttp.BaseRequest
	
	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeIaCFileOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIaCFileOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileOverviewResponseParams struct {
	// <p>文件数量</p>
	TotalFile *uint64 `json:"TotalFile,omitnil,omitempty" name:"TotalFile"`

	// <p>风险文件数量(1:Dockerfile,2:Terraform,3:KubernetesYaml)</p>
	RiskFile []*KeyValueInt `json:"RiskFile,omitnil,omitempty" name:"RiskFile"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIaCFileOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIaCFileOverviewResponseParams `json:"Response"`
}

func (r *DescribeIaCFileOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileReportRequestParams struct {
	// <p>资产ID</p>
	AssetId *uint64 `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeIaCFileReportRequest struct {
	*tchttp.BaseRequest
	
	// <p>资产ID</p>
	AssetId *uint64 `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeIaCFileReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIaCFileReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCFileReportResponseParams struct {
	// <p>检测文件</p>
	File *string `json:"File,omitnil,omitempty" name:"File"`

	// <p>检测状态(0:待扫描,1:检测中,2:已完成,3:检测异常)</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>检测时间</p>
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// <p>风险列表</p>
	Risks []*IaCFileRisk `json:"Risks,omitnil,omitempty" name:"Risks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIaCFileReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIaCFileReportResponseParams `json:"Response"`
}

func (r *DescribeIaCFileReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCFileReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCTokenListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>过滤条件</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeIaCTokenListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>过滤条件</p>
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeIaCTokenListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCTokenListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIaCTokenListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIaCTokenListResponseParams struct {
	// <p>列表</p>
	List []*CICDToken `json:"List,omitnil,omitempty" name:"List"`

	// <p>总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIaCTokenListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIaCTokenListResponseParams `json:"Response"`
}

func (r *DescribeIaCTokenListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIaCTokenListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpInvokeRecordDetailRequestParams struct {
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeIpInvokeRecordDetailRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeIpInvokeRecordDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpInvokeRecordDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpInvokeRecordDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpInvokeRecordDetailResponseParams struct {
	// 调用详情信息
	InvokeDetailInfo []*CosInvokeDetailInfo `json:"InvokeDetailInfo,omitnil,omitempty" name:"InvokeDetailInfo"`

	// 调用权限相关
	InvokePermission []*CosPermissionInfo `json:"InvokePermission,omitnil,omitempty" name:"InvokePermission"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIpInvokeRecordDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpInvokeRecordDetailResponseParams `json:"Response"`
}

func (r *DescribeIpInvokeRecordDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpInvokeRecordDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpInvokeRecordRequestParams struct {
	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeIpInvokeRecordRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeIpInvokeRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpInvokeRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpInvokeRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpInvokeRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIpInvokeRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpInvokeRecordResponseParams `json:"Response"`
}

func (r *DescribeIpInvokeRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpInvokeRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeySandboxCredentialListRequestParams struct {
	// 过滤条件列表，支持的过滤条件如下：
	// CredentialName - 凭证名称（模糊匹配）
	// CredentialType - 凭证类型（精确匹配），取值：access、sts
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeKeySandboxCredentialListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件列表，支持的过滤条件如下：
	// CredentialName - 凭证名称（模糊匹配）
	// CredentialType - 凭证类型（精确匹配），取值：access、sts
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeKeySandboxCredentialListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeySandboxCredentialListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeySandboxCredentialListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeySandboxCredentialListResponseParams struct {
	// 凭证数据列表
	Data []*KeySandboxCredential `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKeySandboxCredentialListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeySandboxCredentialListResponseParams `json:"Response"`
}

func (r *DescribeKeySandboxCredentialListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeySandboxCredentialListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeySandboxCredentialRequestParams struct {
	// 凭证ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeKeySandboxCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 凭证ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeKeySandboxCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeySandboxCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialId")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeySandboxCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeySandboxCredentialResponseParams struct {
	// 凭证ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// 凭证名称
	CredentialName *string `json:"CredentialName,omitnil,omitempty" name:"CredentialName"`

	// 凭证类型
	// 枚举值：
	// access：常规密钥
	// sts：STS临时密钥
	CredentialType *string `json:"CredentialType,omitnil,omitempty" name:"CredentialType"`

	// 生效机器范围
	CredentialEffectScope *CredentialEffectScope `json:"CredentialEffectScope,omitnil,omitempty" name:"CredentialEffectScope"`

	// 常规密钥凭据数据（打码后），CredentialType为access时返回
	// 补充说明：Key为原文，Value为打码后的值（保留前3后4位，中间用***替代）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Access []*AccessCredentialOutput `json:"Access,omitnil,omitempty" name:"Access"`

	// STS凭据数据（打码后），CredentialType为sts时返回
	// 补充说明：System为原文，SecretID和SecretKey为打码后的值（保留前3后4位，中间用***替代）
	// 注意：此字段可能返回 null，表示取不到有效值。
	STS *STSCredentialOutput `json:"STS,omitnil,omitempty" name:"STS"`

	// 创建时间
	// 参数格式：YYYY-MM-DDTHH:mm:ssZ（ISO8601格式）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 参数格式：YYYY-MM-DDTHH:mm:ssZ（ISO8601格式）
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKeySandboxCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeySandboxCredentialResponseParams `json:"Response"`
}

func (r *DescribeKeySandboxCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeySandboxCredentialResponse) FromJsonString(s string) error {
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
type DescribeNotifyAssetConfigRequestParams struct {
	// <p>模块名</p>
	Modules []*string `json:"Modules,omitnil,omitempty" name:"Modules"`
}

type DescribeNotifyAssetConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>模块名</p>
	Modules []*string `json:"Modules,omitnil,omitempty" name:"Modules"`
}

func (r *DescribeNotifyAssetConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotifyAssetConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Modules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotifyAssetConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotifyAssetConfigResponseParams struct {
	// <p>资产范围配置</p>
	Items []*NotifyAssetConfigItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotifyAssetConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotifyAssetConfigResponseParams `json:"Response"`
}

func (r *DescribeNotifyAssetConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotifyAssetConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotifySettingAlertRequestParams struct {

}

type DescribeNotifySettingAlertRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNotifySettingAlertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotifySettingAlertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotifySettingAlertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotifySettingAlertResponseParams struct {
	// <p>通知配置</p>
	Settings []*NotifySetting `json:"Settings,omitnil,omitempty" name:"Settings"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotifySettingAlertResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotifySettingAlertResponseParams `json:"Response"`
}

func (r *DescribeNotifySettingAlertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotifySettingAlertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotifySettingRequestParams struct {

}

type DescribeNotifySettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNotifySettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotifySettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotifySettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotifySettingResponseParams struct {
	// <p>通知设置列表</p>
	List []*NotifySetting `json:"List,omitnil,omitempty" name:"List"`

	// <p>成员账号Id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotifySettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotifySettingResponseParams `json:"Response"`
}

func (r *DescribeNotifySettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotifySettingResponse) FromJsonString(s string) error {
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
type DescribePolicyHitDataRequestParams struct {
	// 查看的日期时间戳
	IndexTimestamp *int64 `json:"IndexTimestamp,omitnil,omitempty" name:"IndexTimestamp"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribePolicyHitDataRequest struct {
	*tchttp.BaseRequest
	
	// 查看的日期时间戳
	IndexTimestamp *int64 `json:"IndexTimestamp,omitnil,omitempty" name:"IndexTimestamp"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribePolicyHitDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyHitDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IndexTimestamp")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyHitDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyHitDataResponseParams struct {
	// 策略命中详情信息
	PolicyHitDetail []*CosRiskInfo `json:"PolicyHitDetail,omitnil,omitempty" name:"PolicyHitDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyHitDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyHitDataResponseParams `json:"Response"`
}

func (r *DescribePolicyHitDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyHitDataResponse) FromJsonString(s string) error {
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
type DescribeRiskBucketListRequestParams struct {
	// 关联的appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 规则id
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeRiskBucketListRequest struct {
	*tchttp.BaseRequest
	
	// 关联的appid
	RelAppId *int64 `json:"RelAppId,omitnil,omitempty" name:"RelAppId"`

	// 规则id
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeRiskBucketListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskBucketListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RelAppId")
	delete(f, "PolicyId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskBucketListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskBucketListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 受影响的存储桶
	Data []*CosRiskBucketInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskBucketListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskBucketListResponseParams `json:"Response"`
}

func (r *DescribeRiskBucketListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskBucketListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCallRecordRequestParams struct {
	// 风险记录ID
	RiskID *int64 `json:"RiskID,omitnil,omitempty" name:"RiskID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeRiskCallRecordRequest struct {
	*tchttp.BaseRequest
	
	// 风险记录ID
	RiskID *int64 `json:"RiskID,omitnil,omitempty" name:"RiskID"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeRiskCallRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCallRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskID")
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCallRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCallRecordResponseParams struct {
	// 风险调用记录列表
	Data []*RiskCallRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 调用记录总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCallRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCallRecordResponseParams `json:"Response"`
}

func (r *DescribeRiskCallRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCallRecordResponse) FromJsonString(s string) error {
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

	// <p>集团账号的成员id</p>
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

	// <p>集团账号的成员id</p>
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
type DescribeRiskItemListRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeRiskItemListRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeRiskItemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskItemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskItemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskItemListResponseParams struct {
	// 列表信息
	Data []*CosRiskViewInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskItemListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskItemListResponseParams `json:"Response"`
}

func (r *DescribeRiskItemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskRuleDetailRequestParams struct {
	// <p>风险规则ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`
}

type DescribeRiskRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>风险规则ID</p>
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
	// <p>风险规则ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>云厂商</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>风险名称</p>
	RiskName *string `json:"RiskName,omitnil,omitempty" name:"RiskName"`

	// <p>风险危害</p>
	RiskInfluence *string `json:"RiskInfluence,omitnil,omitempty" name:"RiskInfluence"`

	// <p>修复指引</p>
	RiskFixAdvice *string `json:"RiskFixAdvice,omitnil,omitempty" name:"RiskFixAdvice"`

	// <p>资产类型</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

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
type DescribeRiskTrendDataRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 指定的日期
	LastDays *int64 `json:"LastDays,omitnil,omitempty" name:"LastDays"`
}

type DescribeRiskTrendDataRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 指定的日期
	LastDays *int64 `json:"LastDays,omitnil,omitempty" name:"LastDays"`
}

func (r *DescribeRiskTrendDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskTrendDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "LastDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskTrendDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskTrendDataResponseParams struct {
	// 风险趋势数据
	CosRiskTrendData []*CosRiskTrendInfo `json:"CosRiskTrendData,omitnil,omitempty" name:"CosRiskTrendData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskTrendDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskTrendDataResponseParams `json:"Response"`
}

func (r *DescribeRiskTrendDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskTrendDataResponse) FromJsonString(s string) error {
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
type DescribeSkillScanPayInfoRequestParams struct {

}

type DescribeSkillScanPayInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSkillScanPayInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillScanPayInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillScanPayInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillScanPayInfoResponseParams struct {
	// <p>订单所属租户 AppID</p>
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>订单状态<br>枚举值：<br>0：未购买<br>1：正常<br>2：隔离<br>6：试用中<br>7：已过期<br>8：试用到期</p>
	OrderStatus *int64 `json:"OrderStatus,omitnil,omitempty" name:"OrderStatus"`

	// <p>总配额</p>
	TotalQuota *int64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// <p>已消耗配额</p>
	UsedCount *int64 `json:"UsedCount,omitnil,omitempty" name:"UsedCount"`

	// <p>支付模式<br>枚举值：<br>0：后付费<br>1：预付费</p>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>自动续费标志<br>枚举值：<br>0：未设置<br>1：自动续费<br>2：不自动续费</p>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>资源ID</p>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>购买时长</p>
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// <p>时长单位</p>
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// <p>订单开始时间</p>
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// <p>订单到期时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>公测结束时间，固定为 2026-06-30 23:59:59</p>
	BetaEndTime *string `json:"BetaEndTime,omitnil,omitempty" name:"BetaEndTime"`

	// <p>服务器当前时间</p>
	TimeNow *string `json:"TimeNow,omitnil,omitempty" name:"TimeNow"`

	// <p>租户 Uin</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>租户昵称</p>
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillScanPayInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillScanPayInfoResponseParams `json:"Response"`
}

func (r *DescribeSkillScanPayInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillScanPayInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillScanResultRequestParams struct {
	// ZIP 文件的 SHA256 Hash
	// 参数格式：sha256:<64位hex>
	ContentHash *string `json:"ContentHash,omitnil,omitempty" name:"ContentHash"`

	// 指定引擎版本号
	// 取值参考：由 CreateSkillScan 接口返回
	EngineVersion *int64 `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 报告签名地址有效期
	// 单位：小时
	// 默认值：8760（1年）
	// 补充说明：对返回的 ReportURL 生效
	ReportURLExpireHours *int64 `json:"ReportURLExpireHours,omitnil,omitempty" name:"ReportURLExpireHours"`
}

type DescribeSkillScanResultRequest struct {
	*tchttp.BaseRequest
	
	// ZIP 文件的 SHA256 Hash
	// 参数格式：sha256:<64位hex>
	ContentHash *string `json:"ContentHash,omitnil,omitempty" name:"ContentHash"`

	// 指定引擎版本号
	// 取值参考：由 CreateSkillScan 接口返回
	EngineVersion *int64 `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 报告签名地址有效期
	// 单位：小时
	// 默认值：8760（1年）
	// 补充说明：对返回的 ReportURL 生效
	ReportURLExpireHours *int64 `json:"ReportURLExpireHours,omitnil,omitempty" name:"ReportURLExpireHours"`
}

func (r *DescribeSkillScanResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillScanResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContentHash")
	delete(f, "EngineVersion")
	delete(f, "ReportURLExpireHours")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillScanResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillScanResultResponseParams struct {
	// 检测状态
	// 枚举值：
	// SUCCESS：检测完成，有结果
	// SCANNING：检测进行中
	// NOT_FOUND：无检测记录
	// FAILED：检测失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 检测结果详情。Status=SUCCESS 时大部分字段有值；Status=SCANNING 时仅包含 ContentHash 和 CreatedAt；Status=FAILED 时仅包含 ContentHash、FailedAt 和 Message；Status=NOT_FOUND 时仅包含 ContentHash
	Data *SkillScanItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillScanResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillScanResultResponseParams `json:"Response"`
}

func (r *DescribeSkillScanResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillScanResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceIPAssetRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeSourceIPAssetRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeSourceIPAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceIPAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceIPAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceIPAssetResponseParams struct {
	// 访问密钥资产列表
	Data []*SourceIPAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// 全部数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSourceIPAssetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSourceIPAssetResponseParams `json:"Response"`
}

func (r *DescribeSourceIPAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceIPAssetResponse) FromJsonString(s string) error {
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
type DescribeUserCallRecordRequestParams struct {
	// 账号uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeUserCallRecordRequest struct {
	*tchttp.BaseRequest
	
	// 账号uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤器
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeUserCallRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCallRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubUin")
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserCallRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCallRecordResponseParams struct {
	// 账号调用记录列表
	Data []*UserCallRecord `json:"Data,omitnil,omitempty" name:"Data"`

	// 调用记录总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserCallRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserCallRecordResponseParams `json:"Response"`
}

func (r *DescribeUserCallRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCallRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDspmInfoListRequestParams struct {
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeUserDspmInfoListRequest struct {
	*tchttp.BaseRequest
	
	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeUserDspmInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDspmInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDspmInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDspmInfoListResponseParams struct {
	// 账号dspm信息列表
	List []*UserDspmInfo `json:"List,omitnil,omitempty" name:"List"`

	// 已勾选数据库资产总数
	SelectedAssetNum *int64 `json:"SelectedAssetNum,omitnil,omitempty" name:"SelectedAssetNum"`

	// 账号总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserDspmInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDspmInfoListResponseParams `json:"Response"`
}

func (r *DescribeUserDspmInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDspmInfoListResponse) FromJsonString(s string) error {
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
	// <p>集团账号的成员id</p>
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
	
	// <p>集团账号的成员id</p>
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

type DetectTypeCount struct {
	// <p>检测方式，0：主机检测，1：网络检测</p>
	DetectType *int64 `json:"DetectType,omitnil,omitempty" name:"DetectType"`

	// <p>策略数量</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DiskPartitionInfo struct {
	// <p>分区名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>挂载路径</p>
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>使用百分比</p>
	Percent *float64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// <p>分区大小(MB)</p>
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// <p>分区类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>已使用(MB)</p>
	Used *uint64 `json:"Used,omitnil,omitempty" name:"Used"`
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

type DomainInfo struct {
	// <p>域名</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>分析时间</p>
	AnalysisTime *string `json:"AnalysisTime,omitnil,omitempty" name:"AnalysisTime"`

	// <p>标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

// Predefined struct for user
type DownloadDspmExportLogRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 要下载的导出任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DownloadDspmExportLogRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 要下载的导出任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DownloadDspmExportLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadDspmExportLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadDspmExportLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadDspmExportLogResponseParams struct {
	// 下载URL
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DownloadDspmExportLogResponse struct {
	*tchttp.BaseResponse
	Response *DownloadDspmExportLogResponseParams `json:"Response"`
}

func (r *DownloadDspmExportLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadDspmExportLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DspmAccessRecord struct {
	// 资产信息
	Asset *DspmDbAsset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 账号
	Accounts []*DspmAssetAccount `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 来源ip信息
	SourceIpList []*DspmIp `json:"SourceIpList,omitnil,omitempty" name:"SourceIpList"`

	// 记录时间
	RecordTime *string `json:"RecordTime,omitnil,omitempty" name:"RecordTime"`

	// 登录成功次数
	LoginSuccessCount *int64 `json:"LoginSuccessCount,omitnil,omitempty" name:"LoginSuccessCount"`

	// 登录失败次数
	LoginFailedCount *int64 `json:"LoginFailedCount,omitnil,omitempty" name:"LoginFailedCount"`
}

type DspmAccessRecordId struct {
	// 来源ip
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 资产列表
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资产账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 记录时间
	RecordTime *string `json:"RecordTime,omitnil,omitempty" name:"RecordTime"`
}

type DspmAccountCount struct {
	// 云账号个数
	UinAccountCount *int64 `json:"UinAccountCount,omitnil,omitempty" name:"UinAccountCount"`

	// 访客账号个数
	PersonCount *int64 `json:"PersonCount,omitnil,omitempty" name:"PersonCount"`

	// 未管控账号个数
	UncontrolledAccountCount *int64 `json:"UncontrolledAccountCount,omitnil,omitempty" name:"UncontrolledAccountCount"`

	// 总账号个数
	TotalAccountCount *int64 `json:"TotalAccountCount,omitnil,omitempty" name:"TotalAccountCount"`
}

type DspmApplyOrder struct {
	// 申请单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 身份id。
	IdentifyId *string `json:"IdentifyId,omitnil,omitempty" name:"IdentifyId"`

	// 申请人账号uin
	ApplicantUin *DspmUinUser `json:"ApplicantUin,omitnil,omitempty" name:"ApplicantUin"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 申请类型。0-关联身份 1-编辑身份 2-创建临时身份
	ApplyType *int64 `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 申请权限。
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 从审批完成后开始计算的访问权限失效时间，临时账号有效。单位毫秒。
	ValidatePeriod *int64 `json:"ValidatePeriod,omitnil,omitempty" name:"ValidatePeriod"`

	// 申请原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 审批步骤
	ApproverSteps []*DspmApproverStep `json:"ApproverSteps,omitnil,omitempty" name:"ApproverSteps"`

	// 管理类型。0-普通成员 1-管理员
	ManagerType *int64 `json:"ManagerType,omitnil,omitempty" name:"ManagerType"`

	// 个人用户信息
	Person *DspmPersonUser `json:"Person,omitnil,omitempty" name:"Person"`

	// 云账号用户信息
	SubjectUser *DspmUinUser `json:"SubjectUser,omitnil,omitempty" name:"SubjectUser"`

	// 审批状态。 0-未审批 1-通过 2-拒绝
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 申请单创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type DspmApproverOrder struct {
	// 对应申请单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 身份id。
	IdentifyId *string `json:"IdentifyId,omitnil,omitempty" name:"IdentifyId"`

	// 申请人账号uin
	ApplicantUin *DspmUinUser `json:"ApplicantUin,omitnil,omitempty" name:"ApplicantUin"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 申请类型。0-关联身份 1-编辑身份 2-创建临时身份
	ApplyType *int64 `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 申请权限
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 从审批完成后开始计算的访问权限失效时间，临时账号有效。单位毫秒。
	ValidatePeriod *int64 `json:"ValidatePeriod,omitnil,omitempty" name:"ValidatePeriod"`

	// 申请原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 管理类型。0-普通成员 1-管理员
	ManagerType *int64 `json:"ManagerType,omitnil,omitempty" name:"ManagerType"`

	// 个人用户信息
	Person *DspmPersonUser `json:"Person,omitnil,omitempty" name:"Person"`

	// 云账号用户信息
	SubjectUser *DspmUinUser `json:"SubjectUser,omitnil,omitempty" name:"SubjectUser"`

	// 对应申请单创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type DspmApproverStep struct {
	// 审批人列表
	ApproverUinSet []*DspmUinUser `json:"ApproverUinSet,omitnil,omitempty" name:"ApproverUinSet"`

	// 审批人
	ApproverUin *string `json:"ApproverUin,omitnil,omitempty" name:"ApproverUin"`

	// 审批状态  0-未审批 1-通过 2-拒绝
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 审批意见
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 审批时间。
	ApproveTime *string `json:"ApproveTime,omitnil,omitempty" name:"ApproveTime"`
}

type DspmArea struct {
	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 省
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 市
	City *string `json:"City,omitnil,omitempty" name:"City"`
}

type DspmAssetAccessTopologyItem struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产账号
	AssetAccount *string `json:"AssetAccount,omitnil,omitempty" name:"AssetAccount"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账号类型
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 资产地址
	AssetIp *string `json:"AssetIp,omitnil,omitempty" name:"AssetIp"`

	// 访问来源ip地址
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 访问来源ip类型
	SourceIpType *string `json:"SourceIpType,omitnil,omitempty" name:"SourceIpType"`

	// 访问频率。次/天
	AccessFrequency *DspmFrequency `json:"AccessFrequency,omitnil,omitempty" name:"AccessFrequency"`

	// 执行SQL频率。条/小时。
	ExecSQLFrequency *DspmFrequency `json:"ExecSQLFrequency,omitnil,omitempty" name:"ExecSQLFrequency"`

	// 访问起始时间
	AccessBeginTime *string `json:"AccessBeginTime,omitnil,omitempty" name:"AccessBeginTime"`

	// 访问结束时间
	AccessEndTime *string `json:"AccessEndTime,omitnil,omitempty" name:"AccessEndTime"`

	// 账号风险数
	AccountRisk *int64 `json:"AccountRisk,omitnil,omitempty" name:"AccountRisk"`

	// 资产风险数
	AssetRisk *int64 `json:"AssetRisk,omitnil,omitempty" name:"AssetRisk"`

	// 所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 身份类型。非身份账号为null。0-未定义 2-长期身份 3-临时身份
	IdentifyType *int64 `json:"IdentifyType,omitnil,omitempty" name:"IdentifyType"`

	// 所属云账号uin用户。
	OwnerUin *DspmUinUser `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 所属个人用户信息。
	Person *DspmPersonUser `json:"Person,omitnil,omitempty" name:"Person"`

	// 账号告警数
	AccountAlarm *int64 `json:"AccountAlarm,omitnil,omitempty" name:"AccountAlarm"`

	// 资产告警数
	AssetAlarm *int64 `json:"AssetAlarm,omitnil,omitempty" name:"AssetAlarm"`
}

type DspmAssetAccount struct {
	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账号类型。 0-未定义 1-服务账号 2-个人账号 3-临时账号
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 所属对象。uin或个人id
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 权限信息
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 状态。 0-不活跃 1-活跃 2-已删除
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 访问权限生效时间。
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间。
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 是否新账号
	IsNewAccount *int64 `json:"IsNewAccount,omitnil,omitempty" name:"IsNewAccount"`

	// 身份类型。非身份账号为null。0-未定义 2-长期身份 3-临时身份
	IdentifyType *int64 `json:"IdentifyType,omitnil,omitempty" name:"IdentifyType"`

	// 所属云账号uin用户。
	OwnerUin *DspmUinUser `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 所属个人用户信息。
	Person *DspmPersonUser `json:"Person,omitnil,omitempty" name:"Person"`

	// 风险统计信息
	RiskCount *DspmRiskCount `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 预设权限。
	PresetPrivilege *DspmDbAccountPrivilege `json:"PresetPrivilege,omitnil,omitempty" name:"PresetPrivilege"`
}

type DspmAssetAccountIdentify struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 所属云账号uin用户。
	OwnerUin *DspmUinUser `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 是否管理员
	IsManager *int64 `json:"IsManager,omitnil,omitempty" name:"IsManager"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账号类型。 0-未定义 1-服务账号 2-个人账号 3-临时账号
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 权限信息
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 活跃状态。 0-不活跃 1-活跃
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 访问权限生效时间。
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间。
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 访客权限申请次数
	PersonApplyCount *int64 `json:"PersonApplyCount,omitnil,omitempty" name:"PersonApplyCount"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 风险统计信息
	RiskCount *DspmRiskCount `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 身份类型。非身份账号为null。0-未定义 2-长期身份 3-临时身份
	IdentifyType *int64 `json:"IdentifyType,omitnil,omitempty" name:"IdentifyType"`

	// 所属个人用户信息。
	Person *DspmPersonUser `json:"Person,omitnil,omitempty" name:"Person"`

	// 创建者账号uin用户。
	CreatorUin *DspmUinUser `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 预设权限。
	PresetPrivilege *DspmDbAccountPrivilege `json:"PresetPrivilege,omitnil,omitempty" name:"PresetPrivilege"`

	// 内网访问地址，如果有多个，使用';'分割
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 身份id
	IdentifyId *string `json:"IdentifyId,omitnil,omitempty" name:"IdentifyId"`

	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 资产所属账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type DspmAssetCount struct {
	// 资产个数
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// 资产有危险风险的个数
	DangerRiskCount *int64 `json:"DangerRiskCount,omitnil,omitempty" name:"DangerRiskCount"`

	// 资产有低风险的个数
	LowRiskCount *int64 `json:"LowRiskCount,omitnil,omitempty" name:"LowRiskCount"`

	// 有待处理风险的实例数
	RiskAssetCount *int64 `json:"RiskAssetCount,omitnil,omitempty" name:"RiskAssetCount"`

	// 有待处理告警的实例数
	AlarmAssetCount *int64 `json:"AlarmAssetCount,omitnil,omitempty" name:"AlarmAssetCount"`
}

type DspmAssetDataScanDetail struct {
	// <p>识别任务状态 0:未识别 1:识别中 2:识别终止 3:识别成功 4:识别失败</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>识别任务状态 0:未识别 1:识别中 2:识别终止 3:识别成功 4:识别失败</p>
	StatusInfo *string `json:"StatusInfo,omitnil,omitempty" name:"StatusInfo"`

	// <p>识别进度</p>
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// <p>最近扫描时间</p>
	LatestScanTime *string `json:"LatestScanTime,omitnil,omitempty" name:"LatestScanTime"`

	// <p>识别失败信息</p>
	ErrorInfo *string `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>数据库数量</p>
	DbCount *uint64 `json:"DbCount,omitnil,omitempty" name:"DbCount"`

	// <p>分类id集合</p>
	CategoryIds []*uint64 `json:"CategoryIds,omitnil,omitempty" name:"CategoryIds"`

	// <p>分类名称集合</p>
	CategoryNames []*string `json:"CategoryNames,omitnil,omitempty" name:"CategoryNames"`

	// <p>扫描任务配置</p>
	TaskConfig *DspmSensitiveScanTaskConfig `json:"TaskConfig,omitnil,omitempty" name:"TaskConfig"`

	// <p>识别结果分类详情</p>
	CategoryDetails []*DspmIdentifyCategoryDetail `json:"CategoryDetails,omitnil,omitempty" name:"CategoryDetails"`

	// <p>任务ID</p>
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DspmAssetDatabaseInfo struct {
	// <p>资产实例id</p>
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>总表数</p>
	TableCount *uint64 `json:"TableCount,omitnil,omitempty" name:"TableCount"`

	// <p>敏感表数</p>
	SensitiveTableCount *uint64 `json:"SensitiveTableCount,omitnil,omitempty" name:"SensitiveTableCount"`

	// <p>数据项id集合</p>
	RuleIds []*uint64 `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// <p>数据项名称集合</p>
	RuleNames []*string `json:"RuleNames,omitnil,omitempty" name:"RuleNames"`

	// <p>分类id集合</p>
	CategoryIds []*uint64 `json:"CategoryIds,omitnil,omitempty" name:"CategoryIds"`

	// <p>分类名称集合</p>
	CategoryNames []*string `json:"CategoryNames,omitnil,omitempty" name:"CategoryNames"`

	// <p>分类详情</p>
	CategoryDetails []*DspmIdentifyCategoryDetail `json:"CategoryDetails,omitnil,omitempty" name:"CategoryDetails"`
}

type DspmAssetFieldInfo struct {
	// <p>资产实例id</p>
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>schema名</p>
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// <p>表名</p>
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// <p>字段名</p>
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// <p>数据项id集合</p>
	RuleIds []*uint64 `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// <p>数据项名称集合</p>
	RuleNames []*string `json:"RuleNames,omitnil,omitempty" name:"RuleNames"`

	// <p>分类id集合</p>
	CategoryIds []*uint64 `json:"CategoryIds,omitnil,omitempty" name:"CategoryIds"`

	// <p>分类名称集合</p>
	CategoryNames []*string `json:"CategoryNames,omitnil,omitempty" name:"CategoryNames"`

	// <p>分类详情</p>
	CategoryDetails []*DspmIdentifyCategoryDetail `json:"CategoryDetails,omitnil,omitempty" name:"CategoryDetails"`

	// <p>字段注释</p>
	FieldComment *string `json:"FieldComment,omitnil,omitempty" name:"FieldComment"`
}

type DspmAssetInstance struct {
	// 资产实例Id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DspmAssetSecurityAnalyseStatus struct {
	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 资产实例Id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 是否支持敏感数据识别。0 不支持；1 支持
	IdentifyScanSupported *int64 `json:"IdentifyScanSupported,omitnil,omitempty" name:"IdentifyScanSupported"`

	// 是否支持日志投递。0-不支持 1-支持
	LogDeliverySupported *int64 `json:"LogDeliverySupported,omitnil,omitempty" name:"LogDeliverySupported"`

	// 安全分析状态（0-关闭 1-打开 2-开通中 3-关闭中）
	SecurityAnalyseStatus *int64 `json:"SecurityAnalyseStatus,omitnil,omitempty" name:"SecurityAnalyseStatus"`

	// 日志投递状态。
	// 0-投递关闭 1-投递打开 2-投递开通中 3-投递关闭中
	LogDeliveryStatus *int64 `json:"LogDeliveryStatus,omitnil,omitempty" name:"LogDeliveryStatus"`

	// 日志审计禁止开通的原因，可选值：VersionNotSupportLogSubscription, InstanceIsUpgrading,CdbRuleAuditEnabled
	LogDeliveryDisableReason *string `json:"LogDeliveryDisableReason,omitnil,omitempty" name:"LogDeliveryDisableReason"`

	// 当前实例的总日志数
	TotalAuditLogs *uint64 `json:"TotalAuditLogs,omitnil,omitempty" name:"TotalAuditLogs"`

	// 已识别敏感数据项个数
	DataScanDetailRuleCount *uint64 `json:"DataScanDetailRuleCount,omitnil,omitempty" name:"DataScanDetailRuleCount"`

	// 操作错误信息
	OperationErrorMsg *string `json:"OperationErrorMsg,omitnil,omitempty" name:"OperationErrorMsg"`
}

type DspmAssetTableInfo struct {
	// <p>资产实例id</p>
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// <p>数据库名称</p>
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// <p>schema名称</p>
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// <p>表名</p>
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// <p>字段数</p>
	FieldCount *uint64 `json:"FieldCount,omitnil,omitempty" name:"FieldCount"`

	// <p>敏感字段数</p>
	SensitiveFieldCount *uint64 `json:"SensitiveFieldCount,omitnil,omitempty" name:"SensitiveFieldCount"`

	// <p>数据项id集合</p>
	RuleIds []*uint64 `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// <p>数据项名称集合</p>
	RuleNames []*string `json:"RuleNames,omitnil,omitempty" name:"RuleNames"`

	// <p>分类id集合</p>
	CategoryIds []*uint64 `json:"CategoryIds,omitnil,omitempty" name:"CategoryIds"`

	// <p>分类名称集合</p>
	CategoryNames []*string `json:"CategoryNames,omitnil,omitempty" name:"CategoryNames"`

	// <p>分类详情</p>
	CategoryDetails []*DspmIdentifyCategoryDetail `json:"CategoryDetails,omitnil,omitempty" name:"CategoryDetails"`

	// <p>数据表id</p>
	TableId *uint64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// <p>表注释</p>
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`
}

type DspmAssetTypeCount struct {
	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 资产安全分析开启数
	OpenCount *int64 `json:"OpenCount,omitnil,omitempty" name:"OpenCount"`

	// 资产安全分析开启中数
	OpeningCount *int64 `json:"OpeningCount,omitnil,omitempty" name:"OpeningCount"`

	// 资产安全分析关闭中数
	ClosingCount *int64 `json:"ClosingCount,omitnil,omitempty" name:"ClosingCount"`

	// 资产安全分析未开启数
	CloseCount *int64 `json:"CloseCount,omitnil,omitempty" name:"CloseCount"`
}

type DspmColumnPrivilege struct {
	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 数据库列名
	Column *string `json:"Column,omitnil,omitempty" name:"Column"`

	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type DspmDatabasePrivilege struct {
	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`
}

type DspmDbAccountPrivilege struct {
	// 使用默认权限。0-未使用；1-只读权限，即SELECT权限；2-全部权限，即global级别全部权限。
	UseDefaultPrivilege *int64 `json:"UseDefaultPrivilege,omitnil,omitempty" name:"UseDefaultPrivilege"`

	// 全局权限数组。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库权限数组。
	DatabasePrivilegesList []*DspmDatabasePrivilege `json:"DatabasePrivilegesList,omitnil,omitempty" name:"DatabasePrivilegesList"`

	// 数据库中的表权限数组。
	TablePrivileges []*DspmTablePrivilege `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// 数据库表中的列权限数组。
	ColumnPrivileges []*DspmColumnPrivilege `json:"ColumnPrivileges,omitnil,omitempty" name:"ColumnPrivileges"`
}

type DspmDbAsset struct {
	// 资产实例Id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	//  资产名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 账号数
	AccountCount *int64 `json:"AccountCount,omitnil,omitempty" name:"AccountCount"`

	// 公网访问地址，如果有多个，使用';'分割
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网访问地址，如果有多个，使用';'分割
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 广域网域名地址，如果有多个，使用';'分割
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资产所在vpc的vpcid
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 资产所在vpc的vpc名
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 资产所在vpc子网的subnetid
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 资产所在vpc子网名
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 实例状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 管理者信息。
	Manager []*DspmUinUser `json:"Manager,omitnil,omitempty" name:"Manager"`

	// 是否绑定身份。0-未绑定 1-已绑定
	BindIdentify *int64 `json:"BindIdentify,omitnil,omitempty" name:"BindIdentify"`

	// 是否管理员
	IsManager *int64 `json:"IsManager,omitnil,omitempty" name:"IsManager"`

	// 风险统计信息
	RiskCount *DspmRiskCount `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 安全建议。
	// Resolve 立即解决
	// Reinforcement 加固
	// None 暂无异常
	SafetyAdvice *string `json:"SafetyAdvice,omitnil,omitempty" name:"SafetyAdvice"`

	// 日志投递状态。
	// 0-投递关闭 1-投递打开 2-投递开通中 3-投递关闭中
	LogDeliveryStatus *int64 `json:"LogDeliveryStatus,omitnil,omitempty" name:"LogDeliveryStatus"`

	// 是否支持日志投递。0-不支持 1-支持
	LogDeliverySupported *int64 `json:"LogDeliverySupported,omitnil,omitempty" name:"LogDeliverySupported"`

	// 数据扫描信息
	DataScanInfo *DspmAssetDataScanDetail `json:"DataScanInfo,omitnil,omitempty" name:"DataScanInfo"`

	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 资产所属账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 安全分析状态（0-关闭 1-打开 2-开通中 3-关闭中）
	SecurityAnalyseStatus *int64 `json:"SecurityAnalyseStatus,omitnil,omitempty" name:"SecurityAnalyseStatus"`

	// 当前实例的总日志数
	TotalAuditLogs *uint64 `json:"TotalAuditLogs,omitnil,omitempty" name:"TotalAuditLogs"`

	// 日志审计禁止开通的原因，可选值：VersionNotSupportLogSubscription, InstanceIsUpgrading, CdbRuleAuditEnabled, AssetNotExists
	LogDeliveryDisableReason *string `json:"LogDeliveryDisableReason,omitnil,omitempty" name:"LogDeliveryDisableReason"`

	// 在线日志的起始时间戳，精确到秒
	OldestOnlineLogTimestamp *uint64 `json:"OldestOnlineLogTimestamp,omitnil,omitempty" name:"OldestOnlineLogTimestamp"`

	// 在线日志的最新时间戳，精确到秒
	NewestOnlineLogTimestamp *uint64 `json:"NewestOnlineLogTimestamp,omitnil,omitempty" name:"NewestOnlineLogTimestamp"`

	// 操作错误信息
	OperationErrorMsg *string `json:"OperationErrorMsg,omitnil,omitempty" name:"OperationErrorMsg"`

	// 是否支持账号操作。0 不支持；1 支持
	AccountOptSupported *int64 `json:"AccountOptSupported,omitnil,omitempty" name:"AccountOptSupported"`

	// 实例类型
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 集群类型（MongoDB），与云接口 DescribeDBInstances 的 ClusterType 一致：0-副本集 1-分片；非 MongoDB 资产固定 0
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 是否支持敏感数据识别。0 不支持；1 支持
	IdentifyScanSupported *int64 `json:"IdentifyScanSupported,omitnil,omitempty" name:"IdentifyScanSupported"`
}

type DspmDbAssetId struct {
	// 资产实例Id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	//  资产名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资产记录id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 资产所属账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type DspmDictionary struct {
	// 字典id
	DictId *uint64 `json:"DictId,omitnil,omitempty" name:"DictId"`

	// 字典名称
	DictName *string `json:"DictName,omitnil,omitempty" name:"DictName"`
}

type DspmFrequency struct {
	// 数量。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 单位。
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type DspmIdentifyAssetStatistic struct {
	// 关联资产管理员数。
	ManagerCount *int64 `json:"ManagerCount,omitnil,omitempty" name:"ManagerCount"`

	// 关联资产普通成员数。
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`
}

type DspmIdentifyCategoryDetail struct {
	// 分类id
	CategoryId *uint64 `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 分类名称
	CategoryName *string `json:"CategoryName,omitnil,omitempty" name:"CategoryName"`

	// 数据项集合
	RuleSet []*DspmIdentifyRuleDetail `json:"RuleSet,omitnil,omitempty" name:"RuleSet"`
}

type DspmIdentifyCount struct {
	// 身份类型。0-未定义 2-长期身份 3-临时身份
	IdentifyType *int64 `json:"IdentifyType,omitnil,omitempty" name:"IdentifyType"`

	// 个数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DspmIdentifyIdItem struct {
	// 身份id。
	IdentifyId *string `json:"IdentifyId,omitnil,omitempty" name:"IdentifyId"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 身份类型。0-未定义 2-长期身份 3-临时身份
	IdentifyType *int64 `json:"IdentifyType,omitnil,omitempty" name:"IdentifyType"`

	// 所属云账号uin用户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *DspmUinUser `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建者账号uin用户。
	CreatorUin *DspmUinUser `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 状态。0-不活跃 1-活跃
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 所属个人用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Person *DspmPersonUser `json:"Person,omitnil,omitempty" name:"Person"`
}

type DspmIdentifyInfoItem struct {
	// 身份id。
	IdentifyId *string `json:"IdentifyId,omitnil,omitempty" name:"IdentifyId"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 身份类型。0-未定义 2-长期身份 3-临时身份
	IdentifyType *int64 `json:"IdentifyType,omitnil,omitempty" name:"IdentifyType"`

	// 所属云账号uin用户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *DspmUinUser `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建者账号uin用户。
	CreatorUin *DspmUinUser `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 关联资产。
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 状态。0-不活跃 1-活跃
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 所属个人用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Person *DspmPersonUser `json:"Person,omitnil,omitempty" name:"Person"`

	// 关联数据资产统计信息。
	AssetStatistic *DspmIdentifyAssetStatistic `json:"AssetStatistic,omitnil,omitempty" name:"AssetStatistic"`

	// 风险统计信息
	RiskCount *DspmRiskCount `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 安全建议。 Resolve 立即解决 Reinforcement 加固 None 暂无异常
	SafetyAdvice *string `json:"SafetyAdvice,omitnil,omitempty" name:"SafetyAdvice"`

	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 资产所属账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type DspmIdentifyRuleDetail struct {
	// <p>数据项id</p>
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// <p>数据项名称</p>
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>敏感级别id</p>
	LevelId *uint64 `json:"LevelId,omitnil,omitempty" name:"LevelId"`

	// <p>敏感级别名称</p>
	LevelName *string `json:"LevelName,omitnil,omitempty" name:"LevelName"`

	// <p>敏感程度</p>
	LevelScore *uint64 `json:"LevelScore,omitnil,omitempty" name:"LevelScore"`
}

type DspmIp struct {
	// ip地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// ip类型。public-公网 private-内网
	IpType *string `json:"IpType,omitnil,omitempty" name:"IpType"`

	// ip标记信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否已经标记信息
	IsRemarked *int64 `json:"IsRemarked,omitnil,omitempty" name:"IsRemarked"`

	// ip归属实例id
	ResourceInstanceId *string `json:"ResourceInstanceId,omitnil,omitempty" name:"ResourceInstanceId"`

	// ip所属产品
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// ip所属地域
	Area *DspmArea `json:"Area,omitnil,omitempty" name:"Area"`

	// 是否新ip地址
	IsNewIp *int64 `json:"IsNewIp,omitnil,omitempty" name:"IsNewIp"`
}

type DspmIpCount struct {
	// 访问Ip个数
	IpCount *int64 `json:"IpCount,omitnil,omitempty" name:"IpCount"`

	// 未打标公网Ip个数
	UnmarkedPublicIpCount *int64 `json:"UnmarkedPublicIpCount,omitnil,omitempty" name:"UnmarkedPublicIpCount"`

	// 内网Ip个数
	PrivateIpCount *int64 `json:"PrivateIpCount,omitnil,omitempty" name:"PrivateIpCount"`
}

type DspmPersonApplyHistoryItem struct {
	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 姓名
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 访问权限生效时间。
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间。
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 访问权限有效时间。单位毫秒。
	ValidatePeriod *int64 `json:"ValidatePeriod,omitnil,omitempty" name:"ValidatePeriod"`

	// 权限信息。
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 是否有效。0-无效；1-有效。
	ValidStatus *int64 `json:"ValidStatus,omitnil,omitempty" name:"ValidStatus"`
}

type DspmPersonIdentifyItem struct {
	// 身份id。
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type DspmPersonUser struct {
	// 个人id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 姓名
	PersonName *string `json:"PersonName,omitnil,omitempty" name:"PersonName"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

type DspmRisk struct {
	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 风险名称
	RiskName *string `json:"RiskName,omitnil,omitempty" name:"RiskName"`

	// 风险英文名称
	RiskNameEn *string `json:"RiskNameEn,omitnil,omitempty" name:"RiskNameEn"`

	// 策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 策略类别
	StrategyCategory *string `json:"StrategyCategory,omitnil,omitempty" name:"StrategyCategory"`

	// 风险等级
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 资产实例Id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 地域
	AssetRegion *string `json:"AssetRegion,omitnil,omitempty" name:"AssetRegion"`

	// 资产账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账号类型
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 风险检出时间
	DetectTime *string `json:"DetectTime,omitnil,omitempty" name:"DetectTime"`

	// 处理状态 0-未处理 1-已处置 2-已忽略
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 身份id
	IdentifyId *string `json:"IdentifyId,omitnil,omitempty" name:"IdentifyId"`

	// 所属云账号uin用户
	OwnerUin *DspmUinUser `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 所属个人用户信息
	Person *DspmPersonUser `json:"Person,omitnil,omitempty" name:"Person"`

	// 风险数据。
	RiskData *string `json:"RiskData,omitnil,omitempty" name:"RiskData"`

	// 是否资产管理员
	IsAssetManager *int64 `json:"IsAssetManager,omitnil,omitempty" name:"IsAssetManager"`

	// 数据起始时间
	DataBeginTime *string `json:"DataBeginTime,omitnil,omitempty" name:"DataBeginTime"`

	// 数据结束时间
	DataEndTime *string `json:"DataEndTime,omitnil,omitempty" name:"DataEndTime"`

	// 风险类型。risk-风险；alarm-告警。
	RiskType *string `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 资产所属账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type DspmRiskCount struct {
	// 待处理风险个数
	UnprocessedRisk *int64 `json:"UnprocessedRisk,omitnil,omitempty" name:"UnprocessedRisk"`

	// 配置风险个数
	//
	// Deprecated: ConfigurationRisk is deprecated.
	ConfigurationRisk *int64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// 基线风险个数
	//
	// Deprecated: BaselineDeviation is deprecated.
	BaselineDeviation *int64 `json:"BaselineDeviation,omitnil,omitempty" name:"BaselineDeviation"`

	// 泄露风险个数
	//
	// Deprecated: LeakDetection is deprecated.
	LeakDetection *int64 `json:"LeakDetection,omitnil,omitempty" name:"LeakDetection"`

	// SQL行为异常风险个数
	SQLBehaviorAnomaly *int64 `json:"SQLBehaviorAnomaly,omitnil,omitempty" name:"SQLBehaviorAnomaly"`

	// 权限异常风险个数
	PermissionAnomaly *int64 `json:"PermissionAnomaly,omitnil,omitempty" name:"PermissionAnomaly"`

	// 登录行为异常风险个数
	LoginBehaviorAnomaly *int64 `json:"LoginBehaviorAnomaly,omitnil,omitempty" name:"LoginBehaviorAnomaly"`

	// 攻击面风险个数
	AttackSurfaceRisk *int64 `json:"AttackSurfaceRisk,omitnil,omitempty" name:"AttackSurfaceRisk"`

	// 账号敏感操作个数
	AccountSensitiveOperation *int64 `json:"AccountSensitiveOperation,omitnil,omitempty" name:"AccountSensitiveOperation"`

	// 待处理告警个数
	UnprocessedAlarm *int64 `json:"UnprocessedAlarm,omitnil,omitempty" name:"UnprocessedAlarm"`

	// 新增事件告警
	NumOfNewAlarmEvent *int64 `json:"NumOfNewAlarmEvent,omitnil,omitempty" name:"NumOfNewAlarmEvent"`

	// 新增配置风险
	NumOfNewConfigRisk *int64 `json:"NumOfNewConfigRisk,omitnil,omitempty" name:"NumOfNewConfigRisk"`
}

type DspmRiskStrategy struct {
	// 策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 策略名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略类型
	StrategyCategory *string `json:"StrategyCategory,omitnil,omitempty" name:"StrategyCategory"`

	// 是否启用。0-禁用 1-启用
	IsEnabled *int64 `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 风险等级。
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 策略规则
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 策略内容
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 命中次数
	HitCount *int64 `json:"HitCount,omitnil,omitempty" name:"HitCount"`

	// 风险类型。risk-风险；alarm-告警。
	RiskType *string `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 资产所属账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 策略id
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type DspmRiskStrategyGroup struct {
	// 策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 策略名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略类型
	StrategyCategory *string `json:"StrategyCategory,omitnil,omitempty" name:"StrategyCategory"`

	// 是否启用。0-禁用 1-启用
	IsEnabled *int64 `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 命中次数
	HitCount *int64 `json:"HitCount,omitnil,omitempty" name:"HitCount"`

	// 风险类型。risk-风险；alarm-告警。
	RiskType *string `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// 策略列表
	StrategyList []*DspmRiskStrategy `json:"StrategyList,omitnil,omitempty" name:"StrategyList"`
}

type DspmRiskTendency struct {
	// 日期
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 未管控账号个数
	UncontrolledAccount *int64 `json:"UncontrolledAccount,omitnil,omitempty" name:"UncontrolledAccount"`

	// 配置风险个数
	//
	// Deprecated: ConfigurationRisk is deprecated.
	ConfigurationRisk *int64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// 基线风险个数
	//
	// Deprecated: BaselineRisk is deprecated.
	BaselineRisk *int64 `json:"BaselineRisk,omitnil,omitempty" name:"BaselineRisk"`

	// 泄露风险个数
	//
	// Deprecated: LeakDetectionRisk is deprecated.
	LeakDetectionRisk *int64 `json:"LeakDetectionRisk,omitnil,omitempty" name:"LeakDetectionRisk"`

	// SQL行为异常风险个数
	SQLBehaviorAnomaly *int64 `json:"SQLBehaviorAnomaly,omitnil,omitempty" name:"SQLBehaviorAnomaly"`

	// 权限异常风险个数
	PermissionAnomaly *int64 `json:"PermissionAnomaly,omitnil,omitempty" name:"PermissionAnomaly"`

	// 登录行为异常风险个数
	LoginBehaviorAnomaly *int64 `json:"LoginBehaviorAnomaly,omitnil,omitempty" name:"LoginBehaviorAnomaly"`

	// 攻击面风险风险个数
	AttackSurfaceRisk *int64 `json:"AttackSurfaceRisk,omitnil,omitempty" name:"AttackSurfaceRisk"`

	// 账号敏感操作个数
	AccountSensitiveOperation *int64 `json:"AccountSensitiveOperation,omitnil,omitempty" name:"AccountSensitiveOperation"`
}

type DspmScheduleConfig struct {
	// 调度类型: daily(按天), weekly(按周), monthly(按月)
	ScheduleType *string `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// 按天不传，按周调度配置（星期几 (1=周一, ..., 7=周日)），按月调度配置（每月第几天 (1-31)）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Day *uint64 `json:"Day,omitnil,omitempty" name:"Day"`

	// 调度时间配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 时区,默认东八区（Asia/Shanghai）
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type DspmSecurityAnalyseStatusCount struct {
	// 资产安全分析开启数
	OpenCount *int64 `json:"OpenCount,omitnil,omitempty" name:"OpenCount"`

	// 资产安全分析开启中数
	OpeningCount *int64 `json:"OpeningCount,omitnil,omitempty" name:"OpeningCount"`

	// 资产安全分析关闭中数
	ClosingCount *int64 `json:"ClosingCount,omitnil,omitempty" name:"ClosingCount"`

	// 资产安全分析未开启数
	CloseCount *int64 `json:"CloseCount,omitnil,omitempty" name:"CloseCount"`

	// 按照资产类型分组的资产安全分析状态统计数
	AssetTypeCountSet []*DspmAssetTypeCount `json:"AssetTypeCountSet,omitnil,omitempty" name:"AssetTypeCountSet"`
}

type DspmSensitiveScanTaskConfig struct {
	// 是否定时任务
	IsScheduled *bool `json:"IsScheduled,omitnil,omitempty" name:"IsScheduled"`

	// 调度周期配置
	ScheduleConfig *DspmScheduleConfig `json:"ScheduleConfig,omitnil,omitempty" name:"ScheduleConfig"`

	// 是否立即扫描
	IsRunAtOnce *bool `json:"IsRunAtOnce,omitnil,omitempty" name:"IsRunAtOnce"`
}

type DspmSupportedAssetType struct {
	// <p>产品名（用于查询）</p>
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// <p>地域列表</p>
	Regions []*RegionConfig `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>产品名-用于展示</p>
	ProductDisplayName *string `json:"ProductDisplayName,omitnil,omitempty" name:"ProductDisplayName"`

	// <p>产品组名</p>
	ProductGroup *string `json:"ProductGroup,omitnil,omitempty" name:"ProductGroup"`

	// <p>给定资产类型的资产实例总数</p>
	AssetTotal *int64 `json:"AssetTotal,omitnil,omitempty" name:"AssetTotal"`
}

type DspmTablePrivilege struct {
	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type DspmUinUser struct {
	// 账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户类型。1-主账号 2-子用户
	UserType *int64 `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type DspmWhitelistStrategy struct {
	// 白名单策略id
	WhitelistStrategyId *string `json:"WhitelistStrategyId,omitnil,omitempty" name:"WhitelistStrategyId"`

	// 策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 白名单策略名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略类型
	StrategyCategory *string `json:"StrategyCategory,omitnil,omitempty" name:"StrategyCategory"`

	// 策略规则
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 策略规则内容描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 白名单类型。risk-风险白名单；alarm-告警白名单。
	RiskType *string `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// 资产所属账号app id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 资产所属账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type EDRFilter struct {
	// <p>过滤键的名称。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>一个或者多个过滤值。</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// <p>模糊搜索</p>
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type EDRRule struct {
	// <p>策略ID</p>
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// <p>策略类型，0-系统策略/System Rule, 1-自定义策略/Custom Rule</p>
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// <p>策略名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>策略描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>内容类型 / Content Type: md5-文件MD5/File MD5, cmdline-命令行/Command Line, dns-DNS, ip_inbound-入站IP/Inbound IP, ip_outbound-出站IP/Outbound IP, custom_file-自定义文件/Custom File, process_network-进程网络/Process Network</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>执行动作 / Action: 0-告警/Alert, 1-放行/Allow, 2-告警并拦截/Alert and Block</p>
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`

	// <p>告警等级 / Alert Level: 0-无/None, 1-高危/High, 2-中危/Medium, 3-低危/Low, 4-提示/Reminder</p>
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>检测模式 / Detect Mode: 0-精准/Precise, 1-均衡/Balanced, 2-深度/Deep</p>
	DetectMode *int64 `json:"DetectMode,omitnil,omitempty" name:"DetectMode"`

	// <p>检测方式 / Detect Type: 0-主机检测/Host Detection, 1-网络检测/Network Detection</p>
	DetectType *int64 `json:"DetectType,omitnil,omitempty" name:"DetectType"`

	// <p>攻击阶段</p>
	AttackStage *string `json:"AttackStage,omitnil,omitempty" name:"AttackStage"`

	// <p>主机生效资产范围 / Effective Scope: 0-指定主机/Specified Hosts, 1-全部主机/All Hosts, 2-专业版/Professional, 3-旗舰版/Flagship, 4-专业版+旗舰版/Professional+Flagship</p>
	CWPScope *int64 `json:"CWPScope,omitnil,omitempty" name:"CWPScope"`

	// <p>主机运行时的自选主机</p>
	QUUIDS []*string `json:"QUUIDS,omitnil,omitempty" name:"QUUIDS"`

	// <p>状态 / Status: 0-开启/Enabled, 1-关闭/Disabled</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>修改时间</p>
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// <p>是否支持拦截 / Support Block: 0-不支持/Not Supported, 1-支持/Supported</p>
	SupportBlock *int64 `json:"SupportBlock,omitnil,omitempty" name:"SupportBlock"`

	// <p>MD5列表,ContentType=md5 时填充</p>
	Md5List []*string `json:"Md5List,omitnil,omitempty" name:"Md5List"`

	// <p>文件名列表,ContentType=custom_file 时填充</p>
	FileName []*string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>文件目录列表,ContentType=custom_file 时填充</p>
	FileDirectory []*string `json:"FileDirectory,omitnil,omitempty" name:"FileDirectory"`

	// <p>域名列表,ContentType=dns 时填充</p>
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// <p>出站IP列表,ContentType=ip_outbound 时填充</p>
	OutboundIP []*string `json:"OutboundIP,omitnil,omitempty" name:"OutboundIP"`

	// <p>入站IP列表,ContentType=ip_inbound 时填充</p>
	InboundIP []*string `json:"InboundIP,omitnil,omitempty" name:"InboundIP"`

	// <p>命令行规则,ContentType=cmdline 时填充</p>
	CmdLineRules *RuleContentCmdLine `json:"CmdLineRules,omitnil,omitempty" name:"CmdLineRules"`

	// <p>容器生效镜像范围 / Container Image Scope: 0-指定镜像/Specified Images, 1-全部镜像/All Images</p>
	TCSSScope *int64 `json:"TCSSScope,omitnil,omitempty" name:"TCSSScope"`

	// <p>生效镜像ID列表 / Image IDs (when TCSSScope=0)</p>
	ImageIDs []*string `json:"ImageIDs,omitnil,omitempty" name:"ImageIDs"`

	// <p>镜像名正则表达式 / Image Names Regex</p>
	ImageNamesRegex *string `json:"ImageNamesRegex,omitnil,omitempty" name:"ImageNamesRegex"`

	// <p>置信度 / Confidence: 0-低/Low, 1-中/Medium, 2-高/High</p>
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// <p>排除的主机列表 / Excluded Host QUUIDS</p>
	ExcludeQUUIDS []*string `json:"ExcludeQUUIDS,omitnil,omitempty" name:"ExcludeQUUIDS"`

	// <p>排除的镜像ID列表 / Excluded Image IDs</p>
	ExcludeImageIDs []*string `json:"ExcludeImageIDs,omitnil,omitempty" name:"ExcludeImageIDs"`

	// <p>进程网络规则 / Process network rules</p>
	ProcessNetworkRules *RuleContentProcessNetwork `json:"ProcessNetworkRules,omitnil,omitempty" name:"ProcessNetworkRules"`

	// <p>策略对应APPID</p>
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>自选实例ID范围</p>
	InstanceIDs []*string `json:"InstanceIDs,omitnil,omitempty" name:"InstanceIDs"`

	// <p>排除实例ID</p>
	ExcludeInstanceIDs []*string `json:"ExcludeInstanceIDs,omitnil,omitempty" name:"ExcludeInstanceIDs"`
}

type EdrAlertCategoryCount struct {
	// <p>告警大类</p>
	AlertCategory *string `json:"AlertCategory,omitnil,omitempty" name:"AlertCategory"`

	// <p>告警数量</p>
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type EdrAlertDetail struct {
	// <p>主键ID</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>租户ID</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>告警唯一标识</p>
	AlertId *string `json:"AlertId,omitnil,omitempty" name:"AlertId"`

	// <p>告警大类（英文枚举：VIRUS_TROJAN/ABNORMAL_LOGIN/HOST_BEHAVIOR/NETWORK_BEHAVIOR/LINK_ENGINE）</p>
	AlertCategory *string `json:"AlertCategory,omitnil,omitempty" name:"AlertCategory"`

	// <p>告警子类型（英文枚举：MALWARE_FILE/MALWARE_PROCESS/RISK_LOGIN/BRUTE_FORCE/DNS/BASH/PRIV_ESCALATION/REVERSE_SHELL/NET_ATTACK/VUL_DEFENCE/MEMORY_SHELL_INJECT/MEMORY_SHELL_SCAN/MULTI_BEHAVIOR_ATTACK）</p>
	AlertSubType *string `json:"AlertSubType,omitnil,omitempty" name:"AlertSubType"`

	// <p>关联规则ID</p>
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// <p>规则类型: 0-系统规则 1-用户自定义</p>
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// <p>告警等级（英文枚举：CRITICAL/HIGH/MEDIUM/LOW/INFO）</p>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>处理状态（英文枚举：PENDING/PROCESSED/WHITELISTED/ISOLATED/CLEANED/IGNORED/ISOLATING/RESTORING/BLOCKED/DELETED）</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>ATT&amp;CK攻击阶段</p>
	AttackStage *string `json:"AttackStage,omitnil,omitempty" name:"AttackStage"`

	// <p>检测模式（英文枚举：PRECISE/BALANCED/DEEP）</p>
	DetectMode *string `json:"DetectMode,omitnil,omitempty" name:"DetectMode"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>主机UUID</p>
	Quuid *string `json:"Quuid,omitnil,omitempty" name:"Quuid"`

	// <p>聚合事件数</p>
	EventCount *int64 `json:"EventCount,omitnil,omitempty" name:"EventCount"`

	// <p>是否付费版</p>
	IsProVersion *int64 `json:"IsProVersion,omitnil,omitempty" name:"IsProVersion"`

	// <p>告警来源（英文枚举：HOST/CONTAINER/K8S/CSIP）</p>
	AlertSource *string `json:"AlertSource,omitnil,omitempty" name:"AlertSource"`

	// <p>容器镜像ID（保留字段，恒为空串）</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>容器ID（保留字段，恒为空串）</p>
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// <p>集群ID（保留字段，恒为空串）</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>首次发现时间</p>
	FirstDetectTime *string `json:"FirstDetectTime,omitnil,omitempty" name:"FirstDetectTime"`

	// <p>最近发现时间</p>
	LatestDetectTime *string `json:"LatestDetectTime,omitnil,omitempty" name:"LatestDetectTime"`

	// <p>规则名称（规则富化）</p>
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>内容类型: md5/cmdline/dns/ip_inbound/ip_outbound/custom_file/process_network</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>实例名（资产富化）</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>公网IP（资产富化）</p>
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// <p>内网IP（资产富化）</p>
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// <p>告警详情JSON字符串（前端通过JSON.parse解析，空值为&quot;{}&quot;）</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>告警名称（子类型中英文名）</p>
	AlertName *string `json:"AlertName,omitnil,omitempty" name:"AlertName"`

	// <p>安全中心标签</p>
	CSIPTags []*CSIPTag `json:"CSIPTags,omitnil,omitempty" name:"CSIPTags"`

	// <p>危害描述（统一字段，合并原各子类型独立字段）</p>
	HarmDesc *string `json:"HarmDesc,omitnil,omitempty" name:"HarmDesc"`

	// <p>修复建议（统一字段）</p>
	SuggestScheme *string `json:"SuggestScheme,omitnil,omitempty" name:"SuggestScheme"`

	// <p>数据来源: vuldb/vdc/intel/default</p>
	HarmDescSource *string `json:"HarmDescSource,omitnil,omitempty" name:"HarmDescSource"`

	// <p>统一威胁情报标签（按子类型路由不同情报源）</p>
	ThreatTags []*string `json:"ThreatTags,omitnil,omitempty" name:"ThreatTags"`

	// <p>Base64解码后的命令（高危命令子类型独有）</p>
	BashCmdDecoded *string `json:"BashCmdDecoded,omitnil,omitempty" name:"BashCmdDecoded"`

	// <p>漏洞名称（网络攻击子类型独有）</p>
	NetVulName *string `json:"NetVulName,omitnil,omitempty" name:"NetVulName"`

	// <p>CVE编号（网络攻击子类型独有）</p>
	NetCVEId *string `json:"NetCVEId,omitnil,omitempty" name:"NetCVEId"`

	// <p>异常行为（网络攻击子类型独有）</p>
	NetAbnormalAction *string `json:"NetAbnormalAction,omitnil,omitempty" name:"NetAbnormalAction"`

	// <p>IP情报信息（为空时不返回）</p>
	IPIntel *IPIntelInfo `json:"IPIntel,omitnil,omitempty" name:"IPIntel"`

	// <p>多行为攻击规则类型分类: sequence/threshold/command</p>
	MultiBehaviorDetectionMode *string `json:"MultiBehaviorDetectionMode,omitnil,omitempty" name:"MultiBehaviorDetectionMode"`

	// <p>告警来源描述（按子类型派生，描述哪个引擎/规则检出）</p>
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// <p>处理时间参数格式：2026-05-26 19:45:48</p>
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// <p>情报富化结果来源（标识本次详情是否成功命中外部情报）；取值 &quot;VDC&quot; / &quot;IPAnalysis&quot; / &quot;BreakingTI&quot; / 空串</p>
	IntelSource *string `json:"IntelSource,omitnil,omitempty" name:"IntelSource"`

	// <p>综合研判，中英文已翻译，中：恶意/安全/未知；英：Malicious/Safe/Unknown</p>
	Verdict *string `json:"Verdict,omitnil,omitempty" name:"Verdict"`

	// <p>研判依据</p>
	VerdictBasis *string `json:"VerdictBasis,omitnil,omitempty" name:"VerdictBasis"`

	// <p>病毒名称</p>
	VirusName *string `json:"VirusName,omitnil,omitempty" name:"VirusName"`

	// <p>病毒家族</p>
	VirusFamily *string `json:"VirusFamily,omitnil,omitempty" name:"VirusFamily"`

	// <p>NetResponsePayload 响应数据包（base64 编码后的字符串）</p>
	NetResponsePayload *string `json:"NetResponsePayload,omitnil,omitempty" name:"NetResponsePayload"`

	// <p>服务进程信息（base64 编码后的 JSON 字符串）</p>
	NetSvcPs *string `json:"NetSvcPs,omitnil,omitempty" name:"NetSvcPs"`
}

type EdrAlertItem struct {
	// <p>告警表id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>APPID</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>告警ID</p>
	AlertId *string `json:"AlertId,omitnil,omitempty" name:"AlertId"`

	// <p>告警大类</p>
	AlertCategory *string `json:"AlertCategory,omitnil,omitempty" name:"AlertCategory"`

	// <p>告警子类</p>
	AlertSubType *string `json:"AlertSubType,omitnil,omitempty" name:"AlertSubType"`

	// <p>策略ID</p>
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// <p>策略类型</p>
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// <p>告警等级</p>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>告警状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>攻击阶段</p>
	AttackStage *string `json:"AttackStage,omitnil,omitempty" name:"AttackStage"`

	// <p>检测模式</p>
	DetectMode *string `json:"DetectMode,omitnil,omitempty" name:"DetectMode"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>QUUID</p>
	Quuid *string `json:"Quuid,omitnil,omitempty" name:"Quuid"`

	// <p>是否付费</p>
	IsProVersion *int64 `json:"IsProVersion,omitnil,omitempty" name:"IsProVersion"`

	// <p>告警来源</p>
	AlertSource *string `json:"AlertSource,omitnil,omitempty" name:"AlertSource"`

	// <p>镜像ID</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>容器id</p>
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// <p>集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>告警数量</p>
	EventCount *int64 `json:"EventCount,omitnil,omitempty" name:"EventCount"`

	// <p>最初发现时间</p>
	FirstDetectTime *string `json:"FirstDetectTime,omitnil,omitempty" name:"FirstDetectTime"`

	// <p>最近发现时间</p>
	LatestDetectTime *string `json:"LatestDetectTime,omitnil,omitempty" name:"LatestDetectTime"`

	// <p>规则名</p>
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>策略类型</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>实例名</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>公共IP</p>
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// <p>内网IP</p>
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// <p>该机器是否开启应用防护</p>
	RaspOpen *bool `json:"RaspOpen,omitnil,omitempty" name:"RaspOpen"`
}

type EdrAlertTarget struct {
	// <p>告警主键ID</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>告警所属账号ID（跨账号，前端必传）</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>告警唯一标识</p>
	AlertId *string `json:"AlertId,omitnil,omitempty" name:"AlertId"`

	// <p>主机UUID（可选，由列表带回透传）</p>
	Quuid *string `json:"Quuid,omitnil,omitempty" name:"Quuid"`

	// <p>实例ID（可选，由列表带回透传，用于安全中心标签富化）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>告警子类型</p>
	AlertSubType *string `json:"AlertSubType,omitnil,omitempty" name:"AlertSubType"`
}

type EdrAlertTargetForIgnore struct {
	// <p>告警主键ID</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>告警所属账号ID（跨账号，前端必传）</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>告警唯一标识</p>
	AlertId *string `json:"AlertId,omitnil,omitempty" name:"AlertId"`

	// <p>主机UUID（可选）</p>
	Quuid *string `json:"Quuid,omitnil,omitempty" name:"Quuid"`

	// <p>实例ID（可选，用于白名单写入）</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type Element struct {
	// 统计类型
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 统计对象
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ExportTask struct {
	// 任务Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 进度百分比
	Percentage *int64 `json:"Percentage,omitnil,omitempty" name:"Percentage"`

	// 任务状态：0.未开始 1.执行中 2.执行成功 3.执行超时 4.执行失败
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务更新时间
	ModifyTime *int64 `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小  字节
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 剩余时间(单位：秒)
	RemainingTime *int64 `json:"RemainingTime,omitnil,omitempty" name:"RemainingTime"`
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
	// <p>云厂商</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>云账号名称</p>
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// <p>云账号</p>
	CloudAccountId *string `json:"CloudAccountId,omitnil,omitempty" name:"CloudAccountId"`

	// <p>域名</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>IP</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>端口或者端口范围</p>
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>开放</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>风险类型</p>
	RiskType *string `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// <p>acl类型</p>
	AclType *string `json:"AclType,omitnil,omitempty" name:"AclType"`

	// <p>acl列表</p>
	AclList *string `json:"AclList,omitnil,omitempty" name:"AclList"`

	// <p>资产ID</p>
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// <p>实例名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>资产类型</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>端口服务数量</p>
	PortServiceCount *uint64 `json:"PortServiceCount,omitnil,omitempty" name:"PortServiceCount"`

	// <p>高危端口数量</p>
	HighRiskPortServiceCount *uint64 `json:"HighRiskPortServiceCount,omitnil,omitempty" name:"HighRiskPortServiceCount"`

	// <p>web应用数量</p>
	WebAppCount *uint64 `json:"WebAppCount,omitnil,omitempty" name:"WebAppCount"`

	// <p>有风险web应用数量</p>
	RiskWebAppCount *uint64 `json:"RiskWebAppCount,omitnil,omitempty" name:"RiskWebAppCount"`

	// <p>弱口令数量</p>
	WeakPasswordCount *uint64 `json:"WeakPasswordCount,omitnil,omitempty" name:"WeakPasswordCount"`

	// <p>漏洞数量</p>
	VulCount *uint64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`

	// <p>首次发现时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>最近更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>实例类型名称</p>
	AssetTypeName *string `json:"AssetTypeName,omitnil,omitempty" name:"AssetTypeName"`

	// <p>开放状态</p>
	DisplayStatus *string `json:"DisplayStatus,omitnil,omitempty" name:"DisplayStatus"`

	// <p>端口状态</p>
	DisplayRiskType *string `json:"DisplayRiskType,omitnil,omitempty" name:"DisplayRiskType"`

	// <p>扫描任务状态</p>
	ScanTaskStatus *string `json:"ScanTaskStatus,omitnil,omitempty" name:"ScanTaskStatus"`

	// <p>uuid</p>
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// <p>是否进行过安全体检</p>
	HasScan *string `json:"HasScan,omitnil,omitempty" name:"HasScan"`

	// <p>租户ID</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>租户ID字符串</p>
	AppIdStr *string `json:"AppIdStr,omitnil,omitempty" name:"AppIdStr"`

	// <p>记录ID</p>
	ExposureID *uint64 `json:"ExposureID,omitnil,omitempty" name:"ExposureID"`

	// <p>端口开放数量</p>
	PortDetectCount *uint64 `json:"PortDetectCount,omitnil,omitempty" name:"PortDetectCount"`

	// <p>端口开放结果</p>
	PortDetectResult *string `json:"PortDetectResult,omitnil,omitempty" name:"PortDetectResult"`

	// <p>标签</p>
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// <p>备注</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// <p>待治理风险数量</p>
	ToGovernedRiskCount *uint64 `json:"ToGovernedRiskCount,omitnil,omitempty" name:"ToGovernedRiskCount"`

	// <p>待治理风险内容</p>
	ToGovernedRiskContent *string `json:"ToGovernedRiskContent,omitnil,omitempty" name:"ToGovernedRiskContent"`

	// <p>资产类型图标</p>
	AssetTypeIconURL *string `json:"AssetTypeIconURL,omitnil,omitempty" name:"AssetTypeIconURL"`

	// <p>资产类型3D图标</p>
	AssetTypeIconSolidURL *string `json:"AssetTypeIconSolidURL,omitnil,omitempty" name:"AssetTypeIconSolidURL"`
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
	// 过滤条件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤条件值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否精确匹配：1 精确匹配；默认模糊匹配
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

	// 弱口令风险
	WeakPasswordRisk *uint64 `json:"WeakPasswordRisk,omitnil,omitempty" name:"WeakPasswordRisk"`
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

type HitRules struct {
	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type IPIntelInfo struct {
	// <p>情报标签（如常规木马、漏洞软件、窃密木马）</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>研判依据</p>
	Basis *string `json:"Basis,omitnil,omitempty" name:"Basis"`

	// <p>所属运营商</p>
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// <p>地理位置</p>
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// <p>家族团伙</p>
	Characteristic *string `json:"Characteristic,omitnil,omitempty" name:"Characteristic"`

	// <p>IP画像</p>
	Purpose *string `json:"Purpose,omitnil,omitempty" name:"Purpose"`

	// <p>反查域名列表</p>
	Referer []*DomainInfo `json:"Referer,omitnil,omitempty" name:"Referer"`
}

type IaCFile struct {
	// <p>ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>appid</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>文件ID</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>文件名称</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>CI/CD名称</p>
	CICDName *string `json:"CICDName,omitnil,omitempty" name:"CICDName"`

	// <p>文件路径</p>
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// <p>文件类型(1:Dockerfile,2:Terraform,3:KubernetesYaml)</p>
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>风险总计数量</p>
	RiskTotalCnt *uint64 `json:"RiskTotalCnt,omitnil,omitempty" name:"RiskTotalCnt"`

	// <p>风险等级数量(0:低危,1:中危,2:高危,3:严重)</p>
	RiskLevelCnt []*KeyValueInt `json:"RiskLevelCnt,omitnil,omitempty" name:"RiskLevelCnt"`

	// <p>扫描时间</p>
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// <p>检测状态(0:待扫描,1:检测中,2:已完成,3:检测异常)</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>扫描失败类型(0:无失败, 1:检测超时, 2:文件格式解析失败, 3:检测失败)</p>
	FailType *int64 `json:"FailType,omitnil,omitempty" name:"FailType"`
}

type IaCFileRisk struct {
	// <p>风险等级(0:低危,1:中危,2:高危,3:严重)</p>
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>风险所在行数</p>
	Line *uint64 `json:"Line,omitnil,omitempty" name:"Line"`

	// <p>规则名称</p>
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>问题描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>修复建议</p>
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`
}

type InquireInfo struct {
	// 计费项名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 购买量
	Value *uint64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type InstanceIDWithAppIdItem struct {
	// <p>APPID</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>实例ID</p>
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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

type KeySandboxCredential struct {
	// 凭证ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// 凭证名称
	CredentialName *string `json:"CredentialName,omitnil,omitempty" name:"CredentialName"`

	// 凭证类型
	// 枚举值：
	// access：常规密钥（Key/Value键值对）
	// sts：STS临时密钥凭据
	CredentialType *string `json:"CredentialType,omitnil,omitempty" name:"CredentialType"`

	// 生效机器范围
	CredentialEffectScope *CredentialEffectScope `json:"CredentialEffectScope,omitnil,omitempty" name:"CredentialEffectScope"`

	// 创建时间
	// 参数格式：YYYY-MM-DDTHH:mm:ssZ（ISO8601格式）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 参数格式：YYYY-MM-DDTHH:mm:ssZ（ISO8601格式）
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type KeyValue struct {
	// 字段
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KeyValueInt struct {
	// <p>键</p>
	Key *int64 `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>值</p>
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type Location struct {
	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 地区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 城市
	City *string `json:"City,omitnil,omitempty" name:"City"`
}

type LogCLSFilter struct {
	// <p>键</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>值</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type LogColumn struct {
	// <p>名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type LogContextInfo struct {
	// <p>主题id</p>
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// <p>上下文</p>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type LogDynamicIndex struct {
	// <p>状态</p>
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type LogFullTextInfo struct {
	// <p>大小写敏感</p>
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// <p>token</p>
	Tokenizer *string `json:"Tokenizer,omitnil,omitempty" name:"Tokenizer"`

	// <p>包含中文</p>
	ContainZH *bool `json:"ContainZH,omitnil,omitempty" name:"ContainZH"`
}

type LogHighLightItem struct {
	// <p>键</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>值</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type LogIndexRuleInfo struct {
	// <p>全文索引</p>
	FullText *LogFullTextInfo `json:"FullText,omitnil,omitempty" name:"FullText"`

	// <p>键值索引</p>
	KeyValue *LogRuleKeyValueInfo `json:"KeyValue,omitnil,omitempty" name:"KeyValue"`

	// <p>标签</p>
	Tag *LogRuleKeyValueInfo `json:"Tag,omitnil,omitempty" name:"Tag"`

	// <p>动态索引</p>
	DynamicIndex *LogDynamicIndex `json:"DynamicIndex,omitnil,omitempty" name:"DynamicIndex"`
}

type LogItem struct {
	// <p>键</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogItems struct {
	// <p>数值</p>
	Data []*LogItem `json:"Data,omitnil,omitempty" name:"Data"`
}

type LogKeyValueInfo struct {
	// <p>键</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>值</p>
	Value *LogValueInfo `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogRuleKeyValueInfo struct {
	// <p>大小写敏感</p>
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// <p>键值索引信息</p>
	KeyValues []*LogKeyValueInfo `json:"KeyValues,omitnil,omitempty" name:"KeyValues"`
}

type LogSearchErrors struct {
	// <p>主题</p>
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// <p>错误信息</p>
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// <p>错误信息</p>
	ErrorCodeStr *string `json:"ErrorCodeStr,omitnil,omitempty" name:"ErrorCodeStr"`
}

type LogSearchInfos struct {
	// <p>主题</p>
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// <p>时间间隔</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>上下文</p>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type LogSearchResult struct {
	// <p>时间</p>
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// <p>主题</p>
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// <p>主题名</p>
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// <p>源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>文件名</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>pkgid</p>
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// <p>pkglogid</p>
	PkgLogId *string `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// <p>json数据</p>
	LogJson *string `json:"LogJson,omitnil,omitempty" name:"LogJson"`

	// <p>主机名</p>
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// <p>log信息</p>
	RawLog *string `json:"RawLog,omitnil,omitempty" name:"RawLog"`

	// <p>索引状态</p>
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`

	// <p>高亮信息</p>
	HighLights []*LogHighLightItem `json:"HighLights,omitnil,omitempty" name:"HighLights"`
}

type LogSearchTopics struct {
	// <p>错误信息</p>
	Errors []*LogSearchErrors `json:"Errors,omitnil,omitempty" name:"Errors"`

	// <p>正常信息</p>
	Infos []*LogSearchInfos `json:"Infos,omitnil,omitempty" name:"Infos"`
}

type LogTopicIndexInfo struct {
	// <p>主题</p>
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// <p>状态</p>
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>规则</p>
	Rule *LogIndexRuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// <p>修改时间</p>
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// <p>是否包含</p>
	IncludeInternalFields *bool `json:"IncludeInternalFields,omitnil,omitempty" name:"IncludeInternalFields"`

	// <p>元数据标签</p>
	MetadataFlag *int64 `json:"MetadataFlag,omitnil,omitempty" name:"MetadataFlag"`
}

type LogValueInfo struct {
	// <p>类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>标签</p>
	Tokenizer *string `json:"Tokenizer,omitnil,omitempty" name:"Tokenizer"`

	// <p>sql标签</p>
	SqlFlag *bool `json:"SqlFlag,omitnil,omitempty" name:"SqlFlag"`

	// <p>包含中文</p>
	ContainZH *bool `json:"ContainZH,omitnil,omitempty" name:"ContainZH"`

	// <p>别名</p>
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type Machine struct {
	// <p>Agent状态，取值：ONLINE-在线，OFFLINE-离线，UNINSTALL-未安装</p>
	AgentStatus *string `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`

	// <p>Agent版本</p>
	AgentVersion *string `json:"AgentVersion,omitnil,omitempty" name:"AgentVersion"`

	// <p>账号AppId</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>云服务商</p>
	CloudFromEnum *string `json:"CloudFromEnum,omitnil,omitempty" name:"CloudFromEnum"`

	// <p>云标签列表</p>
	CloudTags []*Tag `json:"CloudTags,omitnil,omitempty" name:"CloudTags"`

	// <p>CSIP防护类型，取值：BASIC-基础版，PRO-专业版，ULTIMATE-旗舰版</p>
	CsipProtectType *string `json:"CsipProtectType,omitnil,omitempty" name:"CsipProtectType"`

	// <p>暴露状态</p>
	ExposedStatus *string `json:"ExposedStatus,omitnil,omitempty" name:"ExposedStatus"`

	// <p>实例ID</p>
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// <p>实例状态，取值：RUNNING-运行中，STOPPED-已关机，EXPIRED-待回收</p>
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// <p>网卡IP列表</p>
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// <p>是否为新增主机（15天内新增）</p>
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// <p>内核版本</p>
	KernelVersion *string `json:"KernelVersion,omitnil,omitempty" name:"KernelVersion"`

	// <p>最近一次离线时间（Unix时间戳）</p>
	LatestOfflineTime *int64 `json:"LatestOfflineTime,omitnil,omitempty" name:"LatestOfflineTime"`

	// <p>内网IP</p>
	MachineIp *string `json:"MachineIp,omitnil,omitempty" name:"MachineIp"`

	// <p>主机名称</p>
	MachineName *string `json:"MachineName,omitnil,omitempty" name:"MachineName"`

	// <p>操作系统</p>
	MachineOs *string `json:"MachineOs,omitnil,omitempty" name:"MachineOs"`

	// <p>外网IP</p>
	MachineWanIp *string `json:"MachineWanIp,omitnil,omitempty" name:"MachineWanIp"`

	// <p>付费模式，取值：PREPAID-预付费，POSTPAID-后付费</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>项目ID</p>
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>防护类型，取值：NONE-无防护，BASIC-基础版，PRO-专业版，ULTIMATE-旗舰版，PRO_LH-轻量版</p>
	ProtectType *string `json:"ProtectType,omitnil,omitempty" name:"ProtectType"`

	// <p>主机唯一标识</p>
	Quuid *string `json:"Quuid,omitnil,omitempty" name:"Quuid"`

	// <p>地域信息</p>
	RegionInfo *RegionInfo `json:"RegionInfo,omitnil,omitempty" name:"RegionInfo"`

	// <p>备注</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>资产标签列表</p>
	TagItems []*MiniTagItem `json:"TagItems,omitnil,omitempty" name:"TagItems"`

	// <p>标签修改信息</p>
	TagModifyInfo *AssetTagModifyAssetItem `json:"TagModifyInfo,omitnil,omitempty" name:"TagModifyInfo"`

	// <p>TAT状态，取值：ONLINE-在线，OFFLINE-离线</p>
	TatStatus *string `json:"TatStatus,omitnil,omitempty" name:"TatStatus"`

	// <p>Agent唯一标识</p>
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// <p>VPC ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>主机节点类型</p><p>枚举值：</p><ul><li>NONE： 主机节点</li><li>CLUSTER： 集群节点</li><li>CONTAINER： 容器节点</li></ul>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>容器防护状态</p><p>枚举值：</p><ul><li>Enabled： 开启防护</li><li>Disabled： 关闭防护</li><li>Unknown： 未知</li></ul>
	ContainerDefendStatus *string `json:"ContainerDefendStatus,omitnil,omitempty" name:"ContainerDefendStatus"`

	// <p>容器数量</p>
	ContainerCount *uint64 `json:"ContainerCount,omitnil,omitempty" name:"ContainerCount"`

	// <p>核数</p>
	CpuCoreCount *uint64 `json:"CpuCoreCount,omitnil,omitempty" name:"CpuCoreCount"`
}

type MachineDetail struct {
	// <p>Agent状态</p>
	AgentStatus *string `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`

	// <p>Agent版本</p>
	AgentVersion *string `json:"AgentVersion,omitnil,omitempty" name:"AgentVersion"`

	// <p>账号AppId</p>
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>资产类型名称</p>
	AssetTypeName *string `json:"AssetTypeName,omitnil,omitempty" name:"AssetTypeName"`

	// <p>系统启动时间（Unix时间戳）</p>
	BootTime *int64 `json:"BootTime,omitnil,omitempty" name:"BootTime"`

	// <p>购买时间（Unix时间戳）</p>
	BuyTime *int64 `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`

	// <p>云服务商</p>
	CloudFromEnum *string `json:"CloudFromEnum,omitnil,omitempty" name:"CloudFromEnum"`

	// <p>云标签列表</p>
	CloudTags []*Tags `json:"CloudTags,omitnil,omitempty" name:"CloudTags"`

	// <p>内核版本</p>
	CoreVersion *string `json:"CoreVersion,omitnil,omitempty" name:"CoreVersion"`

	// <p>CPU信息</p>
	Cpu *string `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>CPU负载</p>
	CpuLoad *string `json:"CpuLoad,omitnil,omitempty" name:"CpuLoad"`

	// <p>CPU核数</p>
	CpuSize *uint64 `json:"CpuSize,omitnil,omitempty" name:"CpuSize"`

	// <p>设备型号</p>
	DeviceVersion *string `json:"DeviceVersion,omitnil,omitempty" name:"DeviceVersion"`

	// <p>磁盘分区信息</p>
	Disks []*DiskPartitionInfo `json:"Disks,omitnil,omitempty" name:"Disks"`

	// <p>到期时间（Unix时间戳）</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>暴露状态</p>
	ExposedStatus *string `json:"ExposedStatus,omitnil,omitempty" name:"ExposedStatus"`

	// <p>安装时间（Unix时间戳）</p>
	InstallTime *int64 `json:"InstallTime,omitnil,omitempty" name:"InstallTime"`

	// <p>实例ID</p>
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// <p>实例状态</p>
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// <p>内核版本</p>
	KernelVersion *string `json:"KernelVersion,omitnil,omitempty" name:"KernelVersion"`

	// <p>最近一次在线时间（Unix时间戳）</p>
	LatestLiveTime *int64 `json:"LatestLiveTime,omitnil,omitempty" name:"LatestLiveTime"`

	// <p>最近一次离线时间（Unix时间戳）</p>
	LatestOfflineTime *int64 `json:"LatestOfflineTime,omitnil,omitempty" name:"LatestOfflineTime"`

	// <p>内网IP</p>
	MachineIp *string `json:"MachineIp,omitnil,omitempty" name:"MachineIp"`

	// <p>主机名称</p>
	MachineName *string `json:"MachineName,omitnil,omitempty" name:"MachineName"`

	// <p>操作系统（云采集）</p>
	MachineOs *string `json:"MachineOs,omitnil,omitempty" name:"MachineOs"`

	// <p>主机状态</p>
	MachineStatus *string `json:"MachineStatus,omitnil,omitempty" name:"MachineStatus"`

	// <p>外网IP</p>
	MachineWanIp *string `json:"MachineWanIp,omitnil,omitempty" name:"MachineWanIp"`

	// <p>内存大小(MB)</p>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// <p>内存使用率</p>
	MemoryLoad *string `json:"MemoryLoad,omitnil,omitempty" name:"MemoryLoad"`

	// <p>网卡信息</p>
	NetCards []*NetworkCardInfo `json:"NetCards,omitnil,omitempty" name:"NetCards"`

	// <p>操作系统（端采集）</p>
	OsByAgent *string `json:"OsByAgent,omitnil,omitempty" name:"OsByAgent"`

	// <p>付费模式</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>项目ID</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>已防护天数</p>
	ProtectDays *uint64 `json:"ProtectDays,omitnil,omitempty" name:"ProtectDays"`

	// <p>防护类型</p>
	ProtectType *string `json:"ProtectType,omitnil,omitempty" name:"ProtectType"`

	// <p>主机唯一标识</p>
	Quuid *string `json:"Quuid,omitnil,omitempty" name:"Quuid"`

	// <p>地域信息</p>
	RegionInfo *RegionInfo `json:"RegionInfo,omitnil,omitempty" name:"RegionInfo"`

	// <p>备注</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>序列号</p>
	SerialNumber *string `json:"SerialNumber,omitnil,omitempty" name:"SerialNumber"`

	// <p>资产标签列表</p>
	TagItems []*MiniTagItem `json:"TagItems,omitnil,omitempty" name:"TagItems"`

	// <p>标签修改信息</p>
	TagModifyInfo *AssetTagModifyAssetItem `json:"TagModifyInfo,omitnil,omitempty" name:"TagModifyInfo"`

	// <p>Agent唯一标识</p>
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// <p>VPC CIDR</p>
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// <p>VPC ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>VPC名称</p>
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// <p>主机节点类型</p><p>枚举值：</p><ul><li>NONE： 主机节点</li><li>CLUSTER： 集群节点</li><li>CONTAINER： 容器节点</li></ul>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>容器防护状态</p><p>枚举值：</p><ul><li>Enabled： 开启防护</li><li>Disabled： 关闭防护</li><li>Unknown： 未知</li></ul>
	ContainerDefendStatus *string `json:"ContainerDefendStatus,omitnil,omitempty" name:"ContainerDefendStatus"`

	// <p>集群签证md5</p>
	ClusterCaMd5 *string `json:"ClusterCaMd5,omitnil,omitempty" name:"ClusterCaMd5"`

	// <p>容器环境信息</p>
	ContainerEnvInfo *ContainerEnvInfo `json:"ContainerEnvInfo,omitnil,omitempty" name:"ContainerEnvInfo"`

	// <p>集群id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>集群名称</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type MiniTagItem struct {
	// 标签展示颜色。
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`

	// 标签描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 标签 ID。
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 标签键（中文）。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值（中文）。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 标签键（英文）。
	TagKeyEn *string `json:"TagKeyEn,omitnil,omitempty" name:"TagKeyEn"`

	// 标签值（英文）。
	TagValueEn *string `json:"TagValueEn,omitnil,omitempty" name:"TagValueEn"`
}

// Predefined struct for user
type ModifyAILinkSettingRequestParams struct {
	// <p>0 关闭AI-Link智链引擎，1 开启AI-Link智链引擎</p>
	AILinkEnable *uint64 `json:"AILinkEnable,omitnil,omitempty" name:"AILinkEnable"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>深度模式 0-关闭 1-开启</p>
	RuleScopeDeep *uint64 `json:"RuleScopeDeep,omitnil,omitempty" name:"RuleScopeDeep"`

	// <p>均衡模式 0-关闭 1-开启</p>
	RuleScopeBalanced *uint64 `json:"RuleScopeBalanced,omitnil,omitempty" name:"RuleScopeBalanced"`

	// <p>精准模式 0-关闭 1-开启</p>
	RuleScopePrecise *uint64 `json:"RuleScopePrecise,omitnil,omitempty" name:"RuleScopePrecise"`

	// <p>1 全部专业/旗舰版主机，0 自选主机列表</p>
	Scope *uint64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>自选主机Quuid列表（Scope=0时必填）</p>
	Quuids []*string `json:"Quuids,omitnil,omitempty" name:"Quuids"`

	// <p>排除主机Quuid列表（Scope=1时生效）</p>
	ExcludeQuuids []*string `json:"ExcludeQuuids,omitnil,omitempty" name:"ExcludeQuuids"`

	// <p>新增资产自动包含 0 不包含 1包含</p>
	AutoInclude *uint64 `json:"AutoInclude,omitnil,omitempty" name:"AutoInclude"`
}

type ModifyAILinkSettingRequest struct {
	*tchttp.BaseRequest
	
	// <p>0 关闭AI-Link智链引擎，1 开启AI-Link智链引擎</p>
	AILinkEnable *uint64 `json:"AILinkEnable,omitnil,omitempty" name:"AILinkEnable"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>深度模式 0-关闭 1-开启</p>
	RuleScopeDeep *uint64 `json:"RuleScopeDeep,omitnil,omitempty" name:"RuleScopeDeep"`

	// <p>均衡模式 0-关闭 1-开启</p>
	RuleScopeBalanced *uint64 `json:"RuleScopeBalanced,omitnil,omitempty" name:"RuleScopeBalanced"`

	// <p>精准模式 0-关闭 1-开启</p>
	RuleScopePrecise *uint64 `json:"RuleScopePrecise,omitnil,omitempty" name:"RuleScopePrecise"`

	// <p>1 全部专业/旗舰版主机，0 自选主机列表</p>
	Scope *uint64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>自选主机Quuid列表（Scope=0时必填）</p>
	Quuids []*string `json:"Quuids,omitnil,omitempty" name:"Quuids"`

	// <p>排除主机Quuid列表（Scope=1时生效）</p>
	ExcludeQuuids []*string `json:"ExcludeQuuids,omitnil,omitempty" name:"ExcludeQuuids"`

	// <p>新增资产自动包含 0 不包含 1包含</p>
	AutoInclude *uint64 `json:"AutoInclude,omitnil,omitempty" name:"AutoInclude"`
}

func (r *ModifyAILinkSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAILinkSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AILinkEnable")
	delete(f, "MemberId")
	delete(f, "RuleScopeDeep")
	delete(f, "RuleScopeBalanced")
	delete(f, "RuleScopePrecise")
	delete(f, "Scope")
	delete(f, "Quuids")
	delete(f, "ExcludeQuuids")
	delete(f, "AutoInclude")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAILinkSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAILinkSettingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAILinkSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAILinkSettingResponseParams `json:"Response"`
}

func (r *ModifyAILinkSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAILinkSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmRiskStatusRequestParams struct {
	// 告警或者风险id
	AlarmRiskIdSet []*CosAlarmRiskIdInfo `json:"AlarmRiskIdSet,omitnil,omitempty" name:"AlarmRiskIdSet"`

	// 风险或告警状态  1 告警 2风险
	AlarmRiskType *int64 `json:"AlarmRiskType,omitnil,omitempty" name:"AlarmRiskType"`

	// 处置状态
	HandleStatus *int64 `json:"HandleStatus,omitnil,omitempty" name:"HandleStatus"`
}

type ModifyAlarmRiskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 告警或者风险id
	AlarmRiskIdSet []*CosAlarmRiskIdInfo `json:"AlarmRiskIdSet,omitnil,omitempty" name:"AlarmRiskIdSet"`

	// 风险或告警状态  1 告警 2风险
	AlarmRiskType *int64 `json:"AlarmRiskType,omitnil,omitempty" name:"AlarmRiskType"`

	// 处置状态
	HandleStatus *int64 `json:"HandleStatus,omitnil,omitempty" name:"HandleStatus"`
}

func (r *ModifyAlarmRiskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmRiskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmRiskIdSet")
	delete(f, "AlarmRiskType")
	delete(f, "HandleStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmRiskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmRiskStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmRiskStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmRiskStatusResponseParams `json:"Response"`
}

func (r *ModifyAlarmRiskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmRiskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCosAuditMonitorAccountRequestParams struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 需要监测的appid信息
	MonitorAppIdSet []*uint64 `json:"MonitorAppIdSet,omitnil,omitempty" name:"MonitorAppIdSet"`

	// 选择存储桶映射关系
	BindBucket []*CosBucketId `json:"BindBucket,omitnil,omitempty" name:"BindBucket"`
}

type ModifyCosAuditMonitorAccountRequest struct {
	*tchttp.BaseRequest
	
	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 需要监测的appid信息
	MonitorAppIdSet []*uint64 `json:"MonitorAppIdSet,omitnil,omitempty" name:"MonitorAppIdSet"`

	// 选择存储桶映射关系
	BindBucket []*CosBucketId `json:"BindBucket,omitnil,omitempty" name:"BindBucket"`
}

func (r *ModifyCosAuditMonitorAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCosAuditMonitorAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "MonitorAppIdSet")
	delete(f, "BindBucket")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCosAuditMonitorAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCosAuditMonitorAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCosAuditMonitorAccountResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCosAuditMonitorAccountResponseParams `json:"Response"`
}

func (r *ModifyCosAuditMonitorAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCosAuditMonitorAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCosMarkInfoRequestParams struct {
	// 需要修改的存储桶列表
	BucketNameSet []*CosBucketInfo `json:"BucketNameSet,omitnil,omitempty" name:"BucketNameSet"`

	// 备注信息
	MarkInfo *string `json:"MarkInfo,omitnil,omitempty" name:"MarkInfo"`
}

type ModifyCosMarkInfoRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的存储桶列表
	BucketNameSet []*CosBucketInfo `json:"BucketNameSet,omitnil,omitempty" name:"BucketNameSet"`

	// 备注信息
	MarkInfo *string `json:"MarkInfo,omitnil,omitempty" name:"MarkInfo"`
}

func (r *ModifyCosMarkInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCosMarkInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BucketNameSet")
	delete(f, "MarkInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCosMarkInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCosMarkInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCosMarkInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCosMarkInfoResponseParams `json:"Response"`
}

func (r *ModifyCosMarkInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCosMarkInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAccessRecordRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 来源ip
	Id []*DspmAccessRecordId `json:"Id,omitnil,omitempty" name:"Id"`

	// 视图。ip或instance
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 阅读标记。 1-已阅
	Noted *int64 `json:"Noted,omitnil,omitempty" name:"Noted"`
}

type ModifyDspmAccessRecordRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 来源ip
	Id []*DspmAccessRecordId `json:"Id,omitnil,omitempty" name:"Id"`

	// 视图。ip或instance
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 阅读标记。 1-已阅
	Noted *int64 `json:"Noted,omitnil,omitempty" name:"Noted"`
}

func (r *ModifyDspmAccessRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAccessRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "Id")
	delete(f, "View")
	delete(f, "Noted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmAccessRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAccessRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmAccessRecordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmAccessRecordResponseParams `json:"Response"`
}

func (r *ModifyDspmAccessRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAccessRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmApproveStatusRequestParams struct {
	// 申请单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 状态。1-通过 2-拒绝
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 审批信息
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ModifyDspmApproveStatusRequest struct {
	*tchttp.BaseRequest
	
	// 申请单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 状态。1-通过 2-拒绝
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 审批信息
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *ModifyDspmApproveStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmApproveStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	delete(f, "Status")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmApproveStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmApproveStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmApproveStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmApproveStatusResponseParams `json:"Response"`
}

func (r *ModifyDspmApproveStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmApproveStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetAccountPrivilegesRequestParams struct {
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 权限信息
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

type ModifyDspmAssetAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 权限信息
	Privilege *DspmDbAccountPrivilege `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

func (r *ModifyDspmAssetAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "Account")
	delete(f, "Privilege")
	delete(f, "Host")
	delete(f, "RiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmAssetAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetAccountPrivilegesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmAssetAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmAssetAccountPrivilegesResponseParams `json:"Response"`
}

func (r *ModifyDspmAssetAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetAccountRequestParams struct {
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账号类型
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

type ModifyDspmAssetAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账号类型
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

func (r *ModifyDspmAssetAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "Account")
	delete(f, "MemberId")
	delete(f, "Host")
	delete(f, "AccountType")
	delete(f, "Remark")
	delete(f, "RiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmAssetAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmAssetAccountResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmAssetAccountResponseParams `json:"Response"`
}

func (r *ModifyDspmAssetAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetDataScanTaskRequestParams struct {
	// 实例id
	AssetIds []*string `json:"AssetIds,omitnil,omitempty" name:"AssetIds"`

	// 是否定时任务
	IsScheduled *bool `json:"IsScheduled,omitnil,omitempty" name:"IsScheduled"`

	// 是否同意一键授权
	IsAgreeAuth *bool `json:"IsAgreeAuth,omitnil,omitempty" name:"IsAgreeAuth"`

	// 是否立即执行
	IsRunAtOnce *bool `json:"IsRunAtOnce,omitnil,omitempty" name:"IsRunAtOnce"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 调度周期配置
	ScheduleConfig *DspmScheduleConfig `json:"ScheduleConfig,omitnil,omitempty" name:"ScheduleConfig"`
}

type ModifyDspmAssetDataScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	AssetIds []*string `json:"AssetIds,omitnil,omitempty" name:"AssetIds"`

	// 是否定时任务
	IsScheduled *bool `json:"IsScheduled,omitnil,omitempty" name:"IsScheduled"`

	// 是否同意一键授权
	IsAgreeAuth *bool `json:"IsAgreeAuth,omitnil,omitempty" name:"IsAgreeAuth"`

	// 是否立即执行
	IsRunAtOnce *bool `json:"IsRunAtOnce,omitnil,omitempty" name:"IsRunAtOnce"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 调度周期配置
	ScheduleConfig *DspmScheduleConfig `json:"ScheduleConfig,omitnil,omitempty" name:"ScheduleConfig"`
}

func (r *ModifyDspmAssetDataScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetDataScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetIds")
	delete(f, "IsScheduled")
	delete(f, "IsAgreeAuth")
	delete(f, "IsRunAtOnce")
	delete(f, "MemberId")
	delete(f, "ScheduleConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmAssetDataScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetDataScanTaskResponseParams struct {
	// 识别任务id集合
	TaskIdSet []*uint64 `json:"TaskIdSet,omitnil,omitempty" name:"TaskIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmAssetDataScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmAssetDataScanTaskResponseParams `json:"Response"`
}

func (r *ModifyDspmAssetDataScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetDataScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetLogDeliverySwitchRequestParams struct {
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 开关。1-打开 0-关闭
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ModifyDspmAssetLogDeliverySwitchRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 开关。1-打开 0-关闭
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *ModifyDspmAssetLogDeliverySwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetLogDeliverySwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "MemberId")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmAssetLogDeliverySwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetLogDeliverySwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmAssetLogDeliverySwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmAssetLogDeliverySwitchResponseParams `json:"Response"`
}

func (r *ModifyDspmAssetLogDeliverySwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetLogDeliverySwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetSecurityAnalysisSwitchRequestParams struct {
	// 实例信息
	Instances []*DspmAssetInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 开关。1-打开 0-关闭
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ModifyDspmAssetSecurityAnalysisSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 实例信息
	Instances []*DspmAssetInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 开关。1-打开 0-关闭
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *ModifyDspmAssetSecurityAnalysisSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetSecurityAnalysisSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "MemberId")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmAssetSecurityAnalysisSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmAssetSecurityAnalysisSwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmAssetSecurityAnalysisSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmAssetSecurityAnalysisSwitchResponseParams `json:"Response"`
}

func (r *ModifyDspmAssetSecurityAnalysisSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmAssetSecurityAnalysisSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmBackupSettingRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 备份日志保留时长
	BackupLogSaveTime *int64 `json:"BackupLogSaveTime,omitnil,omitempty" name:"BackupLogSaveTime"`

	// 恢复日志保留时长
	RestoreLogSaveTime *int64 `json:"RestoreLogSaveTime,omitnil,omitempty" name:"RestoreLogSaveTime"`

	// 日志最大生命周期限制
	LogMaxSaveTime *int64 `json:"LogMaxSaveTime,omitnil,omitempty" name:"LogMaxSaveTime"`

	// 在线日志最大天数限制
	OnlineLogMaxSaveTime *int64 `json:"OnlineLogMaxSaveTime,omitnil,omitempty" name:"OnlineLogMaxSaveTime"`
}

type ModifyDspmBackupSettingRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 备份日志保留时长
	BackupLogSaveTime *int64 `json:"BackupLogSaveTime,omitnil,omitempty" name:"BackupLogSaveTime"`

	// 恢复日志保留时长
	RestoreLogSaveTime *int64 `json:"RestoreLogSaveTime,omitnil,omitempty" name:"RestoreLogSaveTime"`

	// 日志最大生命周期限制
	LogMaxSaveTime *int64 `json:"LogMaxSaveTime,omitnil,omitempty" name:"LogMaxSaveTime"`

	// 在线日志最大天数限制
	OnlineLogMaxSaveTime *int64 `json:"OnlineLogMaxSaveTime,omitnil,omitempty" name:"OnlineLogMaxSaveTime"`
}

func (r *ModifyDspmBackupSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmBackupSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "BackupLogSaveTime")
	delete(f, "RestoreLogSaveTime")
	delete(f, "LogMaxSaveTime")
	delete(f, "OnlineLogMaxSaveTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmBackupSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmBackupSettingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmBackupSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmBackupSettingResponseParams `json:"Response"`
}

func (r *ModifyDspmBackupSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmBackupSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmIdentifyInfoRequestParams struct {
	// 对象。uin或person id
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyDspmIdentifyInfoRequest struct {
	*tchttp.BaseRequest
	
	// 对象。uin或person id
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyDspmIdentifyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmIdentifyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Subject")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmIdentifyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmIdentifyInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmIdentifyInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmIdentifyInfoResponseParams `json:"Response"`
}

func (r *ModifyDspmIdentifyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmIdentifyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmIpInfoRequestParams struct {
	// ip地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyDspmIpInfoRequest struct {
	*tchttp.BaseRequest
	
	// ip地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyDspmIpInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmIpInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip")
	delete(f, "MemberId")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmIpInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmIpInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmIpInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmIpInfoResponseParams `json:"Response"`
}

func (r *ModifyDspmIpInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmIpInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmPersonalIdentifyRequestParams struct {
	// 身份id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyDspmPersonalIdentifyRequest struct {
	*tchttp.BaseRequest
	
	// 身份id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyDspmPersonalIdentifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmPersonalIdentifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Phone")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmPersonalIdentifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmPersonalIdentifyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmPersonalIdentifyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmPersonalIdentifyResponseParams `json:"Response"`
}

func (r *ModifyDspmPersonalIdentifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmPersonalIdentifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmRestoreLogTaskRequestParams struct {
	// 备份日志Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyDspmRestoreLogTaskRequest struct {
	*tchttp.BaseRequest
	
	// 备份日志Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *ModifyDspmRestoreLogTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmRestoreLogTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmRestoreLogTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmRestoreLogTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmRestoreLogTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmRestoreLogTaskResponseParams `json:"Response"`
}

func (r *ModifyDspmRestoreLogTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmRestoreLogTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmRiskInfoRequestParams struct {
	// 风险id
	RiskId []*string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 风险状态 2-已忽略
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyDspmRiskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 风险id
	RiskId []*string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 风险状态 2-已忽略
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyDspmRiskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmRiskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskId")
	delete(f, "MemberId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmRiskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmRiskInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmRiskInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmRiskInfoResponseParams `json:"Response"`
}

func (r *ModifyDspmRiskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmRiskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmRiskStrategyRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 风险策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 是否启用
	IsEnabled *int64 `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 策略内容，如：{     ThresholdValue: "100" }
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 可选值：Info/Low/Medium/High
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 策略id
	StrategyId []*uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type ModifyDspmRiskStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 风险策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 是否启用
	IsEnabled *int64 `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 策略内容，如：{     ThresholdValue: "100" }
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 可选值：Info/Low/Medium/High
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 策略id
	StrategyId []*uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

func (r *ModifyDspmRiskStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmRiskStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "StrategyType")
	delete(f, "IsEnabled")
	delete(f, "Rule")
	delete(f, "RiskLevel")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmRiskStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmRiskStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmRiskStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmRiskStrategyResponseParams `json:"Response"`
}

func (r *ModifyDspmRiskStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmRiskStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmWhitelistStrategyRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 白名单id
	WhitelistStrategyId *string `json:"WhitelistStrategyId,omitnil,omitempty" name:"WhitelistStrategyId"`

	// 白名单名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyDspmWhitelistStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 白名单id
	WhitelistStrategyId *string `json:"WhitelistStrategyId,omitnil,omitempty" name:"WhitelistStrategyId"`

	// 白名单名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyDspmWhitelistStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmWhitelistStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "WhitelistStrategyId")
	delete(f, "Name")
	delete(f, "Rule")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDspmWhitelistStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDspmWhitelistStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDspmWhitelistStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDspmWhitelistStrategyResponseParams `json:"Response"`
}

func (r *ModifyDspmWhitelistStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDspmWhitelistStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEDRRuleRequestParams struct {
	// <p>策略类型 / Rule Type: 0-系统策略/System Rule, 1-自定义策略/Custom Rule</p>
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// <p>执行动作 / Action: 0-告警/Alert, 1-放行/Allow, 2-告警并拦截/Alert and Block</p>
	AlertAction *int64 `json:"AlertAction,omitnil,omitempty" name:"AlertAction"`

	// <p>生效资产 / Effective Scope: 0-指定主机/Specified Hosts, 1-全部主机/All Hosts, 2-专业版/Professional, 3-旗舰版/Flagship, 4-专业版+旗舰版/Professional+Flagship     QUUIDS        []string json:&quot;QUUIDS&quot;                                      // 主机列表 / Host QUUIDS (when Scope=0)</p>
	CWPScope *int64 `json:"CWPScope,omitnil,omitempty" name:"CWPScope"`

	// <p>容器生效镜像范围 / Container Image Scope: 0-指定镜像/Specified Images, 1-全部镜像/All Images</p>
	TCSSScope *int64 `json:"TCSSScope,omitnil,omitempty" name:"TCSSScope"`

	// <p>开关 / Status: 0-开启/Enabled, 1-关闭/Disabled</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>策略名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>内容类型 / Content Type: md5-文件MD5/File MD5, cmdline-命令行/Command Line, dns-DNS, ip_inbound-入站IP/Inbound IP, ip_outbound-出站IP/Outbound IP, custom_file-自定义文件/Custom File, process_network-进程网络/Process Network</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>告警等级 / Alert Level: 1-高危/High, 2-中危/Medium, 3-低危/Low, 4-提示/Reminder</p>
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>检测模式 / Detect Mode: 0-精准/Precise, 1-均衡/Balanced, 2-深度/Deep</p>
	DetectMode *int64 `json:"DetectMode,omitnil,omitempty" name:"DetectMode"`

	// <p>攻击阶段</p>
	AttackStage *string `json:"AttackStage,omitnil,omitempty" name:"AttackStage"`

	// <p>策略</p>
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// <p>策略描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>处理历史告警 / Handle Old Events: 0-否/No, 1-是/Yes</p>
	DealOldEvents *int64 `json:"DealOldEvents,omitnil,omitempty" name:"DealOldEvents"`

	// <p>ContentType=md5 时传入的 MD5 列表</p>
	Md5List []*string `json:"Md5List,omitnil,omitempty" name:"Md5List"`

	// <p>ContentType=custom_file 时传入的文件名列表(Base64编码)</p>
	FileName []*string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>ContentType=custom_file 时传入的文件目录列表(Base64编码)</p>
	FileDirectory []*string `json:"FileDirectory,omitnil,omitempty" name:"FileDirectory"`

	// <p>ContentType=cmdline 时传入的命令行规则，Process/PProcess/AProcess 的 Exe/Cmdline 字段需要 Base64 编码</p>
	CmdLineRules *RuleContentCmdLine `json:"CmdLineRules,omitnil,omitempty" name:"CmdLineRules"`

	// <p>ContentType=dns 时传入的域名列表(Base64编码)</p>
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// <p>ContentType=ip_outbound 时传入的出站IP列表(Base64编码)</p>
	OutboundIP []*string `json:"OutboundIP,omitnil,omitempty" name:"OutboundIP"`

	// <p>ContentType=ip_inbound 时传入的入站IP列表(Base64编码)</p>
	InboundIP []*string `json:"InboundIP,omitnil,omitempty" name:"InboundIP"`

	// <p>镜像ID列表 / Image IDs (when TCSSScope=0)</p>
	ImageIDs []*string `json:"ImageIDs,omitnil,omitempty" name:"ImageIDs"`

	// <p>ContentType=process_network 时传入的进程网络规则</p>
	ProcessNetworkRules *RuleContentProcessNetwork `json:"ProcessNetworkRules,omitnil,omitempty" name:"ProcessNetworkRules"`

	// <p>选择的多账号的APPID</p>
	TargetAppIDs []*uint64 `json:"TargetAppIDs,omitnil,omitempty" name:"TargetAppIDs"`

	// <p>告警的加白目标机器信息</p>
	Target *EdrAlertTarget `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>自选资产对应的实例ID和APPID</p>
	InstanceIDsWithAppId []*InstanceIDWithAppIdItem `json:"InstanceIDsWithAppId,omitnil,omitempty" name:"InstanceIDsWithAppId"`

	// <p>全选资产排除的实例ID和APPID</p>
	ExcludeInstanceIDsWithAppId []*InstanceIDWithAppIdItem `json:"ExcludeInstanceIDsWithAppId,omitnil,omitempty" name:"ExcludeInstanceIDsWithAppId"`
}

type ModifyEDRRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>策略类型 / Rule Type: 0-系统策略/System Rule, 1-自定义策略/Custom Rule</p>
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// <p>执行动作 / Action: 0-告警/Alert, 1-放行/Allow, 2-告警并拦截/Alert and Block</p>
	AlertAction *int64 `json:"AlertAction,omitnil,omitempty" name:"AlertAction"`

	// <p>生效资产 / Effective Scope: 0-指定主机/Specified Hosts, 1-全部主机/All Hosts, 2-专业版/Professional, 3-旗舰版/Flagship, 4-专业版+旗舰版/Professional+Flagship     QUUIDS        []string json:&quot;QUUIDS&quot;                                      // 主机列表 / Host QUUIDS (when Scope=0)</p>
	CWPScope *int64 `json:"CWPScope,omitnil,omitempty" name:"CWPScope"`

	// <p>容器生效镜像范围 / Container Image Scope: 0-指定镜像/Specified Images, 1-全部镜像/All Images</p>
	TCSSScope *int64 `json:"TCSSScope,omitnil,omitempty" name:"TCSSScope"`

	// <p>开关 / Status: 0-开启/Enabled, 1-关闭/Disabled</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// <p>策略名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>内容类型 / Content Type: md5-文件MD5/File MD5, cmdline-命令行/Command Line, dns-DNS, ip_inbound-入站IP/Inbound IP, ip_outbound-出站IP/Outbound IP, custom_file-自定义文件/Custom File, process_network-进程网络/Process Network</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>告警等级 / Alert Level: 1-高危/High, 2-中危/Medium, 3-低危/Low, 4-提示/Reminder</p>
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>检测模式 / Detect Mode: 0-精准/Precise, 1-均衡/Balanced, 2-深度/Deep</p>
	DetectMode *int64 `json:"DetectMode,omitnil,omitempty" name:"DetectMode"`

	// <p>攻击阶段</p>
	AttackStage *string `json:"AttackStage,omitnil,omitempty" name:"AttackStage"`

	// <p>策略</p>
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// <p>策略描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>处理历史告警 / Handle Old Events: 0-否/No, 1-是/Yes</p>
	DealOldEvents *int64 `json:"DealOldEvents,omitnil,omitempty" name:"DealOldEvents"`

	// <p>ContentType=md5 时传入的 MD5 列表</p>
	Md5List []*string `json:"Md5List,omitnil,omitempty" name:"Md5List"`

	// <p>ContentType=custom_file 时传入的文件名列表(Base64编码)</p>
	FileName []*string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>ContentType=custom_file 时传入的文件目录列表(Base64编码)</p>
	FileDirectory []*string `json:"FileDirectory,omitnil,omitempty" name:"FileDirectory"`

	// <p>ContentType=cmdline 时传入的命令行规则，Process/PProcess/AProcess 的 Exe/Cmdline 字段需要 Base64 编码</p>
	CmdLineRules *RuleContentCmdLine `json:"CmdLineRules,omitnil,omitempty" name:"CmdLineRules"`

	// <p>ContentType=dns 时传入的域名列表(Base64编码)</p>
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// <p>ContentType=ip_outbound 时传入的出站IP列表(Base64编码)</p>
	OutboundIP []*string `json:"OutboundIP,omitnil,omitempty" name:"OutboundIP"`

	// <p>ContentType=ip_inbound 时传入的入站IP列表(Base64编码)</p>
	InboundIP []*string `json:"InboundIP,omitnil,omitempty" name:"InboundIP"`

	// <p>镜像ID列表 / Image IDs (when TCSSScope=0)</p>
	ImageIDs []*string `json:"ImageIDs,omitnil,omitempty" name:"ImageIDs"`

	// <p>ContentType=process_network 时传入的进程网络规则</p>
	ProcessNetworkRules *RuleContentProcessNetwork `json:"ProcessNetworkRules,omitnil,omitempty" name:"ProcessNetworkRules"`

	// <p>选择的多账号的APPID</p>
	TargetAppIDs []*uint64 `json:"TargetAppIDs,omitnil,omitempty" name:"TargetAppIDs"`

	// <p>告警的加白目标机器信息</p>
	Target *EdrAlertTarget `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>自选资产对应的实例ID和APPID</p>
	InstanceIDsWithAppId []*InstanceIDWithAppIdItem `json:"InstanceIDsWithAppId,omitnil,omitempty" name:"InstanceIDsWithAppId"`

	// <p>全选资产排除的实例ID和APPID</p>
	ExcludeInstanceIDsWithAppId []*InstanceIDWithAppIdItem `json:"ExcludeInstanceIDsWithAppId,omitnil,omitempty" name:"ExcludeInstanceIDsWithAppId"`
}

func (r *ModifyEDRRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEDRRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleType")
	delete(f, "AlertAction")
	delete(f, "CWPScope")
	delete(f, "TCSSScope")
	delete(f, "Status")
	delete(f, "MemberId")
	delete(f, "Name")
	delete(f, "ContentType")
	delete(f, "Level")
	delete(f, "DetectMode")
	delete(f, "AttackStage")
	delete(f, "RuleID")
	delete(f, "Description")
	delete(f, "DealOldEvents")
	delete(f, "Md5List")
	delete(f, "FileName")
	delete(f, "FileDirectory")
	delete(f, "CmdLineRules")
	delete(f, "Domains")
	delete(f, "OutboundIP")
	delete(f, "InboundIP")
	delete(f, "ImageIDs")
	delete(f, "ProcessNetworkRules")
	delete(f, "TargetAppIDs")
	delete(f, "Target")
	delete(f, "InstanceIDsWithAppId")
	delete(f, "ExcludeInstanceIDsWithAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEDRRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEDRRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEDRRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEDRRuleResponseParams `json:"Response"`
}

func (r *ModifyEDRRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEDRRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdrAlertPermanentIgnoreRequestParams struct {
	// <p>告警定位列表（支持跨账号），最多500条</p>
	Targets []*EdrAlertTargetForIgnore `json:"Targets,omitnil,omitempty" name:"Targets"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyEdrAlertPermanentIgnoreRequest struct {
	*tchttp.BaseRequest
	
	// <p>告警定位列表（支持跨账号），最多500条</p>
	Targets []*EdrAlertTargetForIgnore `json:"Targets,omitnil,omitempty" name:"Targets"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *ModifyEdrAlertPermanentIgnoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdrAlertPermanentIgnoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Targets")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdrAlertPermanentIgnoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdrAlertPermanentIgnoreResponseParams struct {
	// <p>成功忽略的告警数</p>
	IgnoredCount *int64 `json:"IgnoredCount,omitnil,omitempty" name:"IgnoredCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdrAlertPermanentIgnoreResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdrAlertPermanentIgnoreResponseParams `json:"Response"`
}

func (r *ModifyEdrAlertPermanentIgnoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdrAlertPermanentIgnoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIaCTokenPeriodRequestParams struct {
	// <p>ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>扫描结果存储周期</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type ModifyIaCTokenPeriodRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID</p>
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>扫描结果存储周期</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *ModifyIaCTokenPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIaCTokenPeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIaCTokenPeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIaCTokenPeriodResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIaCTokenPeriodResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIaCTokenPeriodResponseParams `json:"Response"`
}

func (r *ModifyIaCTokenPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIaCTokenPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMachineRemarkRequestParams struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>备注信息</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyMachineRemarkRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>备注信息</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>集团账号的成员id</p>
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *ModifyMachineRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Remark")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMachineRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMachineRemarkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMachineRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMachineRemarkResponseParams `json:"Response"`
}

func (r *ModifyMachineRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotifyAssetConfigRequestParams struct {
	// <p>资产范围配置</p>
	Items []*NotifyAssetConfigItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type ModifyNotifyAssetConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>资产范围配置</p>
	Items []*NotifyAssetConfigItem `json:"Items,omitnil,omitempty" name:"Items"`
}

func (r *ModifyNotifyAssetConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotifyAssetConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNotifyAssetConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotifyAssetConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNotifyAssetConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNotifyAssetConfigResponseParams `json:"Response"`
}

func (r *ModifyNotifyAssetConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotifyAssetConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotifySettingAlertRequestParams struct {
	// <p>通知配置</p>
	Settings []*NotifySetting `json:"Settings,omitnil,omitempty" name:"Settings"`
}

type ModifyNotifySettingAlertRequest struct {
	*tchttp.BaseRequest
	
	// <p>通知配置</p>
	Settings []*NotifySetting `json:"Settings,omitnil,omitempty" name:"Settings"`
}

func (r *ModifyNotifySettingAlertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotifySettingAlertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Settings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNotifySettingAlertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotifySettingAlertResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNotifySettingAlertResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNotifySettingAlertResponseParams `json:"Response"`
}

func (r *ModifyNotifySettingAlertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotifySettingAlertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotifySettingRequestParams struct {
	// <p>通知模块</p><p>枚举值：</p><ul><li>AkSk： 云API风险治理</li><li>Alert： 告警中心</li><li>Agent： 客户端</li></ul>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// <p>通知设置模式</p><p>枚举值：</p><ul><li>0： 标准模式</li><li>1： 高级模式</li></ul>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>通知状态</p><p>枚举值：</p><ul><li>0： 通知关闭</li><li>1： 通知开启</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>通知开始时间</p><p>参数格式：hh:mm:ss</p>
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// <p>通知结束时间</p><p>参数格式：hh:mm:ss</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>通知选项</p><p>枚举值：</p><ul><li>CRITICAL： 告警等级：严重</li><li>HIGH： 告警等级：高危</li><li>MEDIUM： 告警等级：中危</li><li>LOW： 告警等级：低危</li><li>INFO： 告警等级：提醒</li><li>AGENT_UNINSTALL： 客户端卸载</li><li>AGENT_OFFLINE： 客户端离线</li></ul>
	Option []*string `json:"Option,omitnil,omitempty" name:"Option"`
}

type ModifyNotifySettingRequest struct {
	*tchttp.BaseRequest
	
	// <p>通知模块</p><p>枚举值：</p><ul><li>AkSk： 云API风险治理</li><li>Alert： 告警中心</li><li>Agent： 客户端</li></ul>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// <p>通知设置模式</p><p>枚举值：</p><ul><li>0： 标准模式</li><li>1： 高级模式</li></ul>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>通知状态</p><p>枚举值：</p><ul><li>0： 通知关闭</li><li>1： 通知开启</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>通知开始时间</p><p>参数格式：hh:mm:ss</p>
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// <p>通知结束时间</p><p>参数格式：hh:mm:ss</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>通知选项</p><p>枚举值：</p><ul><li>CRITICAL： 告警等级：严重</li><li>HIGH： 告警等级：高危</li><li>MEDIUM： 告警等级：中危</li><li>LOW： 告警等级：低危</li><li>INFO： 告警等级：提醒</li><li>AGENT_UNINSTALL： 客户端卸载</li><li>AGENT_OFFLINE： 客户端离线</li></ul>
	Option []*string `json:"Option,omitnil,omitempty" name:"Option"`
}

func (r *ModifyNotifySettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotifySettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Mode")
	delete(f, "Status")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Option")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNotifySettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotifySettingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNotifySettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNotifySettingResponseParams `json:"Response"`
}

func (r *ModifyNotifySettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotifySettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrganizationAccountStatusRequestParams struct {
	// 修改集团账号状态，1 开启， 0关闭
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type ModifyOrganizationAccountStatusRequest struct {
	*tchttp.BaseRequest
	
	// 修改集团账号状态，1 开启， 0关闭
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
type ModifyPolicyStatusRequestParams struct {
	// 策略id集合
	PolicyIdSet []*int64 `json:"PolicyIdSet,omitnil,omitempty" name:"PolicyIdSet"`

	// 状态值
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyPolicyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 策略id集合
	PolicyIdSet []*int64 `json:"PolicyIdSet,omitnil,omitempty" name:"PolicyIdSet"`

	// 状态值
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyPolicyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPolicyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyIdSet")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPolicyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPolicyStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPolicyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPolicyStatusResponseParams `json:"Response"`
}

func (r *ModifyPolicyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPolicyStatusResponse) FromJsonString(s string) error {
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

type NetworkCardInfo struct {
	// <p>DNS服务器</p>
	DnsServer *string `json:"DnsServer,omitnil,omitempty" name:"DnsServer"`

	// <p>网关</p>
	Gateway *string `json:"Gateway,omitnil,omitempty" name:"Gateway"`

	// <p>IP地址</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>IPv6地址</p>
	Ipv6 *string `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// <p>MAC地址</p>
	Mac *string `json:"Mac,omitnil,omitempty" name:"Mac"`

	// <p>网卡名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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

type NotifyAssetConfigItem struct {
	// <p>模块名</p>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// <p>子模块</p>
	SubModule *string `json:"SubModule,omitnil,omitempty" name:"SubModule"`

	// <p>资产范围</p><p>枚举值：</p><ul><li>0： 无含义</li><li>1： 全部</li><li>2： 自选</li><li>3： 按标签</li></ul>
	AssetRange *int64 `json:"AssetRange,omitnil,omitempty" name:"AssetRange"`

	// <p>选中的实例ID</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>剔除的实例ID</p>
	ExcludedInstanceIds []*string `json:"ExcludedInstanceIds,omitnil,omitempty" name:"ExcludedInstanceIds"`

	// <p>标签ID</p>
	TagIds []*int64 `json:"TagIds,omitnil,omitempty" name:"TagIds"`

	// <p>云标签</p>
	CloudTags []*string `json:"CloudTags,omitnil,omitempty" name:"CloudTags"`

	// <p>总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type NotifySetting struct {
	// <p>通知模块</p><p>枚举值：</p><ul><li>AkSk： 云API风险治理</li><li>Alert： 告警中心</li><li>Agent： 客户端</li></ul>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// <p>通知设置模式</p><p>枚举值：</p><ul><li>0： 标准模式</li><li>1： 高级模式</li></ul>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>通知状态</p><p>枚举值：</p><ul><li>0： 通知关闭</li><li>1： 通知开启</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>通知开始时间</p><p>参数格式：hh:mm:ss</p>
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// <p>通知结束时间</p><p>参数格式：hh:mm:ss</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>资产范围</p><p>枚举值：</p><ul><li>1： 全部主机</li><li>2： 自选主机</li><li>3： 按标签选择</li></ul>
	AssetRange *int64 `json:"AssetRange,omitnil,omitempty" name:"AssetRange"`

	// <p>通知选项</p><p>枚举值：</p><ul><li>CRITICAL： 告警等级：严重</li><li>HIGH： 告警等级：高危</li><li>MEDIUM： 告警等级：中危</li><li>LOW： 告警等级：低危</li><li>INFO： 告警等级：提醒</li><li>AGENT_UNINSTALL： 客户端卸载</li><li>AGENT_OFFLINE： 客户端离线</li></ul>
	Option []*string `json:"Option,omitnil,omitempty" name:"Option"`

	// <p>通知模块（二级模块）</p>
	SubModule *string `json:"SubModule,omitnil,omitempty" name:"SubModule"`

	// <p>处置状态等</p>
	Item []*string `json:"Item,omitnil,omitempty" name:"Item"`
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

type RegionConfig struct {
	// <p>地域</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>地域中文</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// <p>是否境外</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Foreign *int64 `json:"Foreign,omitnil,omitempty" name:"Foreign"`

	// <p>地域码</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>是否自驾云</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAutoDriveCloud *int64 `json:"IsAutoDriveCloud,omitnil,omitempty" name:"IsAutoDriveCloud"`

	// <p>是否支持nat</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportNat *int64 `json:"IsSupportNat,omitnil,omitempty" name:"IsSupportNat"`

	// <p>地区信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionArea *string `json:"RegionArea,omitnil,omitempty" name:"RegionArea"`

	// <p>地域英文</p>
	RegionNameEN *string `json:"RegionNameEN,omitnil,omitempty" name:"RegionNameEN"`
}

type RegionInfo struct {
	// <p>地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>地域编码</p>
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// <p>地域ID</p>
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>地域名称</p>
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// <p>地域英文名称</p>
	RegionNameEn *string `json:"RegionNameEn,omitnil,omitempty" name:"RegionNameEn"`
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

// Predefined struct for user
type ResetDspmAssetAccountPasswordRequestParams struct {
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

type ResetDspmAssetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

func (r *ResetDspmAssetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDspmAssetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "Account")
	delete(f, "Host")
	delete(f, "RiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetDspmAssetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDspmAssetAccountPasswordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetDspmAssetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetDspmAssetAccountPasswordResponseParams `json:"Response"`
}

func (r *ResetDspmAssetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDspmAssetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDspmExportLogRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type RetryDspmExportLogRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *RetryDspmExportLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDspmExportLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryDspmExportLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDspmExportLogResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryDspmExportLogResponse struct {
	*tchttp.BaseResponse
	Response *RetryDspmExportLogResponseParams `json:"Response"`
}

func (r *RetryDspmExportLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDspmExportLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevertDspmAssetAccountRequestParams struct {
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 是否回退权限
	PrivilegeFlag *int64 `json:"PrivilegeFlag,omitnil,omitempty" name:"PrivilegeFlag"`

	// 是否回退密码
	PasswordFlag *int64 `json:"PasswordFlag,omitnil,omitempty" name:"PasswordFlag"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

type RevertDspmAssetAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 是否回退权限
	PrivilegeFlag *int64 `json:"PrivilegeFlag,omitnil,omitempty" name:"PrivilegeFlag"`

	// 是否回退密码
	PasswordFlag *int64 `json:"PasswordFlag,omitnil,omitempty" name:"PasswordFlag"`

	// 风险id
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`
}

func (r *RevertDspmAssetAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevertDspmAssetAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	delete(f, "Account")
	delete(f, "Host")
	delete(f, "PrivilegeFlag")
	delete(f, "PasswordFlag")
	delete(f, "RiskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevertDspmAssetAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevertDspmAssetAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RevertDspmAssetAccountResponse struct {
	*tchttp.BaseResponse
	Response *RevertDspmAssetAccountResponseParams `json:"Response"`
}

func (r *RevertDspmAssetAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevertDspmAssetAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskCallRecord struct {
	// 接口名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 接口中文描述
	EventDescCN *string `json:"EventDescCN,omitnil,omitempty" name:"EventDescCN"`

	// 接口英文描述
	EventDescEN *string `json:"EventDescEN,omitnil,omitempty" name:"EventDescEN"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品中文名称
	ProductNameCN *string `json:"ProductNameCN,omitnil,omitempty" name:"ProductNameCN"`

	// 调用次数
	CallCount *int64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`
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
	// <p>首次发现时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>风险状态</p>
	RiskStatus *int64 `json:"RiskStatus,omitnil,omitempty" name:"RiskStatus"`

	// <p>风险内容</p>
	RiskContent *string `json:"RiskContent,omitnil,omitempty" name:"RiskContent"`

	// <p>云厂商</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>云厂商名称</p>
	ProviderName *string `json:"ProviderName,omitnil,omitempty" name:"ProviderName"`

	// <p>云账号</p>
	CloudAccountId *string `json:"CloudAccountId,omitnil,omitempty" name:"CloudAccountId"`

	// <p>云账号名称</p>
	CloudAccountName *string `json:"CloudAccountName,omitnil,omitempty" name:"CloudAccountName"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>风险ID</p>
	RiskId *uint64 `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// <p>风险规则ID</p>
	RiskRuleId *string `json:"RiskRuleId,omitnil,omitempty" name:"RiskRuleId"`

	// <p>风险验证状态</p>
	CheckStatus *string `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`

	// <p>用户AppID</p>
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>资产类型</p>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>风险忽略原因</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type RiskRuleInfo struct {
	// 风险检查项ID
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 云厂商名称
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 风险名称
	RiskTitle *string `json:"RiskTitle,omitnil,omitempty" name:"RiskTitle"`

	// 检查类型
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// 风险等级
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 风险危害
	RiskInfluence *string `json:"RiskInfluence,omitnil,omitempty" name:"RiskInfluence"`

	// 风险修复指引报告链接
	RiskFixAdvance *string `json:"RiskFixAdvance,omitnil,omitempty" name:"RiskFixAdvance"`

	// 边界管控
	DispositionType *string `json:"DispositionType,omitnil,omitempty" name:"DispositionType"`
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

type RuleContentCmdLine struct {
	// <p>进程命令行信息</p>
	Process *RuleContentProcessInfo `json:"Process,omitnil,omitempty" name:"Process"`

	// <p>父进程命令行信息</p>
	ParentProcess *RuleContentProcessInfo `json:"ParentProcess,omitnil,omitempty" name:"ParentProcess"`

	// <p>祖先进程命令行信息</p>
	AncestorProcess *RuleContentProcessInfo `json:"AncestorProcess,omitnil,omitempty" name:"AncestorProcess"`
}

type RuleContentProcessInfo struct {
	// <p>进程文件路径</p>
	Exe *string `json:"Exe,omitnil,omitempty" name:"Exe"`

	// <p>进程命令行</p>
	CmdLine *string `json:"CmdLine,omitnil,omitempty" name:"CmdLine"`
}

type RuleContentProcessNetwork struct {
	// <p>当前进程</p>
	Process *RuleContentProcessInfo `json:"Process,omitnil,omitempty" name:"Process"`

	// <p>目标IP（必填）: 支持单个IP/IP范围/CIDR, 支持IPv4和IPv6</p>
	DstIP *string `json:"DstIP,omitnil,omitempty" name:"DstIP"`

	// <p>父进程</p>
	ParentProcess *RuleContentProcessInfo `json:"ParentProcess,omitnil,omitempty" name:"ParentProcess"`

	// <p>目标端口列表（可选）: 支持1-65535, 为空表示不限端口</p>
	DstPorts []*uint64 `json:"DstPorts,omitnil,omitempty" name:"DstPorts"`
}

type STSCredentialOutput struct {
	// 凭据提供商标识（原文），如tencentCam、aws、aliyun等
	System *string `json:"System,omitnil,omitempty" name:"System"`

	// SecretID（打码后）
	// 补充说明：保留前3后4位，中间用***替代；长度不足7位时全部替换为***
	SecretID *string `json:"SecretID,omitnil,omitempty" name:"SecretID"`

	// SecretKey（打码后）
	// 补充说明：保留前3后4位，中间用***替代；长度不足7位时全部替换为***
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`
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

// Predefined struct for user
type SendDspmAssetLoginSmsCodeRequestParams struct {
	// 个人id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 数据库资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 主机。默认'%'
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type SendDspmAssetLoginSmsCodeRequest struct {
	*tchttp.BaseRequest
	
	// 个人id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 数据库资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 主机。默认'%'
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *SendDspmAssetLoginSmsCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendDspmAssetLoginSmsCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "AssetId")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendDspmAssetLoginSmsCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendDspmAssetLoginSmsCodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendDspmAssetLoginSmsCodeResponse struct {
	*tchttp.BaseResponse
	Response *SendDspmAssetLoginSmsCodeResponseParams `json:"Response"`
}

func (r *SendDspmAssetLoginSmsCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendDspmAssetLoginSmsCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SensitiveDetail struct {
	// 字段分类
	CategoryRule *string `json:"CategoryRule,omitnil,omitempty" name:"CategoryRule"`

	// 字段分级
	LevelRisk *string `json:"LevelRisk,omitnil,omitempty" name:"LevelRisk"`

	// 1:敏感信息字段
	// 0:非敏感字段
	IsSensitive *uint64 `json:"IsSensitive,omitnil,omitempty" name:"IsSensitive"`
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

type SkillCapabilityTag struct {
	// 能力标签标识，适合程序判定、过滤或聚合使用
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 能力标签展示名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type SkillRuleCatalogItem struct {
	// 融合规则 ID（9xxxx）
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 风险类别名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type SkillScanEngineResult struct {
	// 子引擎类型
	// 枚举值：
	// AI：AI 引擎
	// STATIC：静态分析引擎
	ScanType *string `json:"ScanType,omitnil,omitempty" name:"ScanType"`

	// 该引擎命中的规则列表
	RuleList []*SkillScanRuleHit `json:"RuleList,omitnil,omitempty" name:"RuleList"`
}

type SkillScanItem struct {
	// <p>Skill 名称</p>
	SkillName *string `json:"SkillName,omitnil,omitempty" name:"SkillName"`

	// <p>Skill 描述，帮助理解 Skill 的主要用途</p>
	SkillDescription *string `json:"SkillDescription,omitnil,omitempty" name:"SkillDescription"`

	// <p>ZIP 文件的 SHA256 Hash<br>参数格式：sha256:&lt;64位hex&gt;</p>
	ContentHash *string `json:"ContentHash,omitnil,omitempty" name:"ContentHash"`

	// <p>原始上传 ZIP 文件解压后的实际文件数，也是计费的范围，扫描成功后1个文件计为1次额度</p>
	UploadFileCount *int64 `json:"UploadFileCount,omitnil,omitempty" name:"UploadFileCount"`

	// <p>综合风险等级<br>枚举值：<br>malicious：恶意<br>suspicious：可疑<br>benign：可信</p>
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>风险主标签融合规则 ID（9xxxx），由服务端从命中的融合风险标签中生成；benign 且无规则命中时为空。展示名称可通过 RuleCatalog 获取</p>
	PrimaryRuleID *string `json:"PrimaryRuleID,omitnil,omitempty" name:"PrimaryRuleID"`

	// <p>综合处置建议，用于指导调用方优先执行下线、隔离、修复、复检等动作。历史结果中可能为空。传 Language=en-US 时返回英文文案</p>
	Mitigation *string `json:"Mitigation,omitnil,omitempty" name:"Mitigation"`

	// <p>风险综合描述，对本次检测发现的风险进行概括性说明。传 Language=en-US 时返回英文文案</p>
	RiskDescription *string `json:"RiskDescription,omitnil,omitempty" name:"RiskDescription"`

	// <p>安全评分取值范围：[0, 100]补充说明：分数越高越安全</p>
	SecurityScore *int64 `json:"SecurityScore,omitnil,omitempty" name:"SecurityScore"`

	// <p>本次扫描使用的引擎版本号</p>
	EngineVersion *int64 `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// <p>Skill 能力标签列表，描述 Skill 具备的能力特征或适用场景。不等同于风险标签，也不参与风险等级判定。传 Language=en-US 时 Name 切换为英文，ID 保持不变</p>
	CapabilityTags []*SkillCapabilityTag `json:"CapabilityTags,omitnil,omitempty" name:"CapabilityTags"`

	// <p>融合规则目录全集，包含所有融合规则类别（9xxxx），调用方可据此展示分类标签，无需本地维护映射表。传 Language=en-US 时返回英文名称</p>
	RuleCatalog []*SkillRuleCatalogItem `json:"RuleCatalog,omitnil,omitempty" name:"RuleCatalog"`

	// <p>扫描结果详情，按子引擎分组。每个元素包含 ScanType（引擎类型）和 RuleList（命中规则列表）。规则中的 RuleID 使用融合编码（9xxxx），可与 RuleCatalog 交叉引用。传 Language=en-US 时 Description 返回英文文本</p>
	ScanItems []*SkillScanEngineResult `json:"ScanItems,omitnil,omitempty" name:"ScanItems"`

	// <p>综合安全审计报告地址（签名 URL）。有效期由请求参数 ReportURLExpireHours 控制</p>
	ReportURL *string `json:"ReportURL,omitnil,omitempty" name:"ReportURL"`

	// <p>扫描完成时间。仅 Status=SUCCESS 时有值<br>参数格式：YYYY-MM-DDTHH:mm:ssZ（ISO8601格式）</p>
	ScannedAt *string `json:"ScannedAt,omitnil,omitempty" name:"ScannedAt"`

	// <p>任务创建时间。仅 Status=SCANNING 时有值<br>参数格式：YYYY-MM-DDTHH:mm:ssZ（ISO8601格式）</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>失败时间。仅 Status=FAILED 时有值<br>参数格式：YYYY-MM-DDTHH:mm:ssZ（ISO8601格式）</p>
	FailedAt *string `json:"FailedAt,omitnil,omitempty" name:"FailedAt"`

	// <p>失败原因描述。仅 Status=FAILED 时有值</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type SkillScanRuleHit struct {
	// 融合规则编号（9xxxx），可与 RuleCatalog 交叉引用
	RuleID *string `json:"RuleID,omitnil,omitempty" name:"RuleID"`

	// 当前命中规则的具体发现描述，包含文件位置、行为特征、风险点等信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type SkillState struct {
	// SKILL安装状态
	// 枚举值：
	// 0：未安装
	// 1：安装中
	// 2：已安装
	// 3：安装失败
	// 4：卸载中
	// 5：卸载失败
	SkillInstallStatus *int64 `json:"SkillInstallStatus,omitnil,omitempty" name:"SkillInstallStatus"`

	// SKILL安装/卸载操作时间
	// 参数格式：YYYY-MM-DDTHH:mm:ssZ（ISO8601格式）
	SkillInstallTime *string `json:"SkillInstallTime,omitnil,omitempty" name:"SkillInstallTime"`

	// SKILL安装/卸载结果描述信息
	SkillInstallResult *string `json:"SkillInstallResult,omitnil,omitempty" name:"SkillInstallResult"`
}

type SourceIPAsset struct {
	// 源IP id
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 源IP
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 账号所属APPID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// IP地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 调用方式
	// -1:未统计
	// 0:控制台调用
	// 1:API
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// IP类型
	// 0:账号内（未备注）
	// 1:账号外（未备注）
	// 2:账号内   (已备注)
	// 3:账号外   (已备注)
	IPType *int64 `json:"IPType,omitnil,omitempty" name:"IPType"`

	// 告警信息列表
	AccessKeyAlarmList []*AccessKeyAlarmInfo `json:"AccessKeyAlarmList,omitnil,omitempty" name:"AccessKeyAlarmList"`

	// ak信息列表
	AKInfo []*AKInfo `json:"AKInfo,omitnil,omitempty" name:"AKInfo"`

	// 调用接口数量
	ActionCount *int64 `json:"ActionCount,omitnil,omitempty" name:"ActionCount"`

	// 最近访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// IP关联实例ID，如果为空字符串，代表非账号内资产
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// IP关联实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 账号所属Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 展示状态
	ShowStatus *bool `json:"ShowStatus,omitnil,omitempty" name:"ShowStatus"`

	// 运营商字段
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// 账号外vpc信息
	VpcInfo []*SourceIPVpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// 云类型
	// 0为腾讯云
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`
}

type SourceIPVpcInfo struct {
	// 账号名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// vpc所属appid
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// vpc id
	VpcID *string `json:"VpcID,omitnil,omitempty" name:"VpcID"`

	// vpc 名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`
}

type StandardItem struct {
	// 规范ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 规范名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type StandardTerm struct {
	// 标签
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 条款
	Terms []*string `json:"Terms,omitnil,omitempty" name:"Terms"`
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
	// <p>主键ID，无业务意义仅作为唯一键</p>
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>子账号Appid</p>
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// <p>子账号UIn</p>
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// <p>子账号名称</p>
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// <p>主账号Appid</p>
	OwnerAppID *string `json:"OwnerAppID,omitnil,omitempty" name:"OwnerAppID"`

	// <p>主账号Uin</p>
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// <p>主账号名称</p>
	OwnerNickName *string `json:"OwnerNickName,omitnil,omitempty" name:"OwnerNickName"`

	// <p>所属主账号memberId信息</p>
	OwnerMemberID *string `json:"OwnerMemberID,omitnil,omitempty" name:"OwnerMemberID"`

	// <p>账户类型，0为腾讯云账户，1为AWS账户</p>
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// <p>可访问服务数量</p>
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// <p>可访问接口数量</p>
	InterfaceCount *int64 `json:"InterfaceCount,omitnil,omitempty" name:"InterfaceCount"`

	// <p>可访问资源数量</p>
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// <p>访问/行为日志数量</p>
	LogCount *int64 `json:"LogCount,omitnil,omitempty" name:"LogCount"`

	// <p>权限配置风险</p>
	ConfigRiskCount *int64 `json:"ConfigRiskCount,omitnil,omitempty" name:"ConfigRiskCount"`

	// <p>危险行为告警</p>
	ActionRiskCount *int64 `json:"ActionRiskCount,omitnil,omitempty" name:"ActionRiskCount"`

	// <p>是否接入操作审计日志</p>
	IsAccessCloudAudit *bool `json:"IsAccessCloudAudit,omitnil,omitempty" name:"IsAccessCloudAudit"`

	// <p>是否配置风险的安全体检</p>
	IsAccessCheck *bool `json:"IsAccessCheck,omitnil,omitempty" name:"IsAccessCheck"`

	// <p>是否配置用户行为管理策略</p>
	IsAccessUeba *bool `json:"IsAccessUeba,omitnil,omitempty" name:"IsAccessUeba"`

	// <p>创建时间（Unix时间戳）</p>
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
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

// Predefined struct for user
type SyncDspmAssetsRequestParams struct {
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type SyncDspmAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *SyncDspmAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncDspmAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncDspmAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncDspmAssetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncDspmAssetsResponse struct {
	*tchttp.BaseResponse
	Response *SyncDspmAssetsResponseParams `json:"Response"`
}

func (r *SyncDspmAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncDspmAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncDspmUsersRequestParams struct {

}

type SyncDspmUsersRequest struct {
	*tchttp.BaseRequest
	
}

func (r *SyncDspmUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncDspmUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncDspmUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncDspmUsersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncDspmUsersResponse struct {
	*tchttp.BaseResponse
	Response *SyncDspmUsersResponseParams `json:"Response"`
}

func (r *SyncDspmUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncDspmUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableField struct {
	// 数据库名
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 数据库视图名
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 数据库表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 数据库字段名
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 字段敏感信息
	Sensitive *SensitiveDetail `json:"Sensitive,omitnil,omitempty" name:"Sensitive"`
}

type Tag struct {
	// 标签键。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签值。
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

	// APP ID
	TargetAppId *string `json:"TargetAppId,omitnil,omitempty" name:"TargetAppId"`
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

type TrafficPluginState struct {
	// 插件安装状态（上层聚合）
	// 枚举值：
	// NONE：未安装
	// INSTALLING：安装中
	// INSTALLED：已安装
	// INSTALL_FAIL：安装失败
	InstallStatus *string `json:"InstallStatus,omitnil,omitempty" name:"InstallStatus"`

	// 插件安装细分状态。取值与 InstallStatus 对应：未安装（InstallStatus=UNINSTALL）时为空字符串；安装成功（InstallStatus=INSTALLED）时为 SUCCESS；安装失败（InstallStatus=INSTALL_FAIL）时为具体失败原因
	// 枚举值：
	// NOT_SUPPORT：环境不支持
	// CONTAINER_NOT_FOUND：容器不存在
	// REQUIRE_RESTART：需要重启
	// CA_FAILED：CA 失败
	// EBPF_FAILED：eBPF 失败
	// IPTABLE_FAILED：iptables 失败
	// REDIRECT_FAILED：流量重定向失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态文案（由 Status 根据请求语言派生的国际化描述）
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 插件最近活跃时间
	// 参数格式：YYYY-MM-DDTHH:mm:ssZ（ISO8601格式）
	ActivityTime *string `json:"ActivityTime,omitnil,omitempty" name:"ActivityTime"`
}

type TrafficRuleState struct {
	// <p>沙箱插件模块名</p>
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// <p>沙箱规则状态</p><p>枚举值：</p><ul><li>ON： 开启</li><li>OFF： 关闭</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
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
type UpdateAccessKeyAlarmStatusRequestParams struct {
	// 状态  0:未处理 1:已处理 2:已忽略
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 告警ID列表
	AlarmIDList []*int64 `json:"AlarmIDList,omitnil,omitempty" name:"AlarmIDList"`

	// 风险ID列表
	RiskIDList []*int64 `json:"RiskIDList,omitnil,omitempty" name:"RiskIDList"`
}

type UpdateAccessKeyAlarmStatusRequest struct {
	*tchttp.BaseRequest
	
	// 状态  0:未处理 1:已处理 2:已忽略
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 告警ID列表
	AlarmIDList []*int64 `json:"AlarmIDList,omitnil,omitempty" name:"AlarmIDList"`

	// 风险ID列表
	RiskIDList []*int64 `json:"RiskIDList,omitnil,omitempty" name:"RiskIDList"`
}

func (r *UpdateAccessKeyAlarmStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyAlarmStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "MemberId")
	delete(f, "AlarmIDList")
	delete(f, "RiskIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAccessKeyAlarmStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyAlarmStatusResponseParams struct {
	// 0成功，1失败
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAccessKeyAlarmStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAccessKeyAlarmStatusResponseParams `json:"Response"`
}

func (r *UpdateAccessKeyAlarmStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyAlarmStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyRemarkRequestParams struct {
	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 源IP 名称
	SourceIPList []*string `json:"SourceIPList,omitnil,omitempty" name:"SourceIPList"`

	// ak名称
	AccessKeyList []*string `json:"AccessKeyList,omitnil,omitempty" name:"AccessKeyList"`

	// 源IP的ID
	SourceIPIDList []*uint64 `json:"SourceIPIDList,omitnil,omitempty" name:"SourceIPIDList"`

	// AK的ID
	AccessKeyIDList []*uint64 `json:"AccessKeyIDList,omitnil,omitempty" name:"AccessKeyIDList"`
}

type UpdateAccessKeyRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 集团账号的成员id
	MemberId []*string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 源IP 名称
	SourceIPList []*string `json:"SourceIPList,omitnil,omitempty" name:"SourceIPList"`

	// ak名称
	AccessKeyList []*string `json:"AccessKeyList,omitnil,omitempty" name:"AccessKeyList"`

	// 源IP的ID
	SourceIPIDList []*uint64 `json:"SourceIPIDList,omitnil,omitempty" name:"SourceIPIDList"`

	// AK的ID
	AccessKeyIDList []*uint64 `json:"AccessKeyIDList,omitnil,omitempty" name:"AccessKeyIDList"`
}

func (r *UpdateAccessKeyRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Remark")
	delete(f, "MemberId")
	delete(f, "SourceIPList")
	delete(f, "AccessKeyList")
	delete(f, "SourceIPIDList")
	delete(f, "AccessKeyIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAccessKeyRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyRemarkResponseParams struct {
	// 0:成功 1:失败
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAccessKeyRemarkResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAccessKeyRemarkResponseParams `json:"Response"`
}

func (r *UpdateAccessKeyRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type UserCallRecord struct {
	// 调用源IP
	SourceIP *string `json:"SourceIP,omitnil,omitempty" name:"SourceIP"`

	// 调用类型
	// 0:控制台调用
	// 1:API
	EventType *int64 `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 调用次数
	CallCount *int64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// 调用错误码
	// 0表示成功
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 首次调用时间
	FirstCallTime *string `json:"FirstCallTime,omitnil,omitempty" name:"FirstCallTime"`

	// 最后调用时间
	LastCallTime *string `json:"LastCallTime,omitnil,omitempty" name:"LastCallTime"`

	// 调用源IP备注
	SourceIPRemark *string `json:"SourceIPRemark,omitnil,omitempty" name:"SourceIPRemark"`

	// 调用源IP地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 用户/角色名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 聚合日期
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// appid
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 运营商
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`
}

type UserDspmInfo struct {
	// APPID
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 账号昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 账号下数据库资产数量
	AssetNum *int64 `json:"AssetNum,omitnil,omitempty" name:"AssetNum"`

	// 账号下开启安全分析数据库资产数量
	UsedAssetNum *int64 `json:"UsedAssetNum,omitnil,omitempty" name:"UsedAssetNum"`

	// 是否被共享，1-被共享，2-未被共享
	IsShared *int64 `json:"IsShared,omitnil,omitempty" name:"IsShared"`

	// 是否单独购买，1-单独购买，2-未单独购买
	IsSelfBuy *int64 `json:"IsSelfBuy,omitnil,omitempty" name:"IsSelfBuy"`

	// 配额来源账号
	ShareFromAppID *int64 `json:"ShareFromAppID,omitnil,omitempty" name:"ShareFromAppID"`

	// 云类型（0：腾讯云 1:亚马逊云 2:微软云 3:谷歌云 4:阿里云 5:华为云）
	CloudType *int64 `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 账号是否隔离中
	IsIsolating *bool `json:"IsIsolating,omitnil,omitempty" name:"IsIsolating"`

	// 是否正在数据清理
	IsDataCleaning *bool `json:"IsDataCleaning,omitnil,omitempty" name:"IsDataCleaning"`
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

// Predefined struct for user
type VerifyDspmAssetLoginCodeRequestParams struct {
	// 个人id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 数据库资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 验证码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 主机。默认'%'
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type VerifyDspmAssetLoginCodeRequest struct {
	*tchttp.BaseRequest
	
	// 个人id
	PersonId *string `json:"PersonId,omitnil,omitempty" name:"PersonId"`

	// 数据库资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 验证码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 主机。默认'%'
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *VerifyDspmAssetLoginCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDspmAssetLoginCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "AssetId")
	delete(f, "Code")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyDspmAssetLoginCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyDspmAssetLoginCodeResponseParams struct {
	// 账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 密码信息
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 有效期开始时间
	ValidateStart *string `json:"ValidateStart,omitnil,omitempty" name:"ValidateStart"`

	// 有效期结束时间
	ValidateEnd *string `json:"ValidateEnd,omitnil,omitempty" name:"ValidateEnd"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyDspmAssetLoginCodeResponse struct {
	*tchttp.BaseResponse
	Response *VerifyDspmAssetLoginCodeResponseParams `json:"Response"`
}

func (r *VerifyDspmAssetLoginCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDspmAssetLoginCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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