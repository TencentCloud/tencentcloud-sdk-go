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

package v20180608

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AggregationObj struct {
	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数组
	Bucket []*Bucket `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

type AlarmInfoRsp struct {
	// 近7天威胁告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackEvent []*AttackEvent `json:"AttackEvent,omitnil,omitempty" name:"AttackEvent"`
}

type AlertDetail struct {
	// 告警基础信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseInfo *AlertType `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// 告警详情，json序列化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type AlertListAggregations struct {
	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AlertListData struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 返回列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertList []*AlertType `json:"AlertList,omitnil,omitempty" name:"AlertList"`

	// 聚合参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Aggregations *AlertListAggregations `json:"Aggregations,omitnil,omitempty" name:"Aggregations"`
}

type AlertType struct {
	// 标准时间格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertTime *string `json:"AlertTime,omitnil,omitempty" name:"AlertTime"`

	// 唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertId *string `json:"AlertId,omitnil,omitempty" name:"AlertId"`

	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetPrivateIp []*string `json:"AssetPrivateIp,omitnil,omitempty" name:"AssetPrivateIp"`

	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertName *string `json:"AlertName,omitnil,omitempty" name:"AlertName"`

	// 告警级别  0:未知 1:低危 2:中危 3:高危 4:严重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 攻击字段1
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackChain *string `json:"AttackChain,omitnil,omitempty" name:"AttackChain"`

	// 攻击字段2
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackId *string `json:"AttackId,omitnil,omitempty" name:"AttackId"`

	// 关注点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Concerns []*ConcernInfo `json:"Concerns,omitnil,omitempty" name:"Concerns"`

	// 1：已防御，0,2：仅检测(0:告警类 1:拦截类 2:放行类 )
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 0/空：未知，1：未成功，2：成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackResult *int64 `json:"AttackResult,omitnil,omitempty" name:"AttackResult"`

	// //调查状态  0/空：未启用，1：调查中，2：完成调查
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventStatus *int64 `json:"EventStatus,omitnil,omitempty" name:"EventStatus"`

	// //关联事件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// //处置状态  0：未关闭，1：已关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 恶意实体
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConcernMaliciousCount *int64 `json:"ConcernMaliciousCount,omitnil,omitempty" name:"ConcernMaliciousCount"`

	// 受害者实体
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConcernVictimCount *int64 `json:"ConcernVictimCount,omitnil,omitempty" name:"ConcernVictimCount"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VictimAssetType *string `json:"VictimAssetType,omitnil,omitempty" name:"VictimAssetType"`

	// 告警子类
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 攻击技术名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackName *string `json:"AttackName,omitnil,omitempty" name:"AttackName"`

	// 外网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetPublicIp []*string `json:"AssetPublicIp,omitnil,omitempty" name:"AssetPublicIp"`

	// 攻击战术名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackTactic *string `json:"AttackTactic,omitnil,omitempty" name:"AttackTactic"`

	// 资产子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	VictimAssetSub *string `json:"VictimAssetSub,omitnil,omitempty" name:"VictimAssetSub"`

	// 资产vpc
	// 注意：此字段可能返回 null，表示取不到有效值。
	VictimAssetVpc *string `json:"VictimAssetVpc,omitnil,omitempty" name:"VictimAssetVpc"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 资产组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetGroupName []*string `json:"AssetGroupName,omitnil,omitempty" name:"AssetGroupName"`

	// 资产项目名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetProjectName *string `json:"AssetProjectName,omitnil,omitempty" name:"AssetProjectName"`

	// 失陷资产内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	VictimAssetContent []*string `json:"VictimAssetContent,omitnil,omitempty" name:"VictimAssetContent"`

	// 错误报告状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	WrongReportStatus *int64 `json:"WrongReportStatus,omitnil,omitempty" name:"WrongReportStatus"`

	// 错误报告Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WrongReportConditionId *int64 `json:"WrongReportConditionId,omitnil,omitempty" name:"WrongReportConditionId"`
}

type Asset struct {
	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 区域
	AssetRegionName *string `json:"AssetRegionName,omitnil,omitempty" name:"AssetRegionName"`

	// 所属网络
	AssetVpcid *string `json:"AssetVpcid,omitnil,omitempty" name:"AssetVpcid"`

	// 主机类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 主机状态
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 引擎版本
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 数据库标识
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 配置风险统计数
	AssetCspmRiskNum *int64 `json:"AssetCspmRiskNum,omitnil,omitempty" name:"AssetCspmRiskNum"`

	// 主机IP
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 资产唯一标识
	AssetUniqid *string `json:"AssetUniqid,omitnil,omitempty" name:"AssetUniqid"`

	// 付费类型
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 安全事件统计数
	AssetEventNum *int64 `json:"AssetEventNum,omitnil,omitempty" name:"AssetEventNum"`

	// 漏洞统计数
	AssetVulNum *int64 `json:"AssetVulNum,omitnil,omitempty" name:"AssetVulNum"`

	// 主机IP内网
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 所属分组
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaAssetDiscoverTime *string `json:"SsaAssetDiscoverTime,omitnil,omitempty" name:"SsaAssetDiscoverTime"`

	// 下线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaAssetDeleteTime *string `json:"SsaAssetDeleteTime,omitnil,omitempty" name:"SsaAssetDeleteTime"`

	// 是否是新增资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 所属子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetSubnetId *string `json:"AssetSubnetId,omitnil,omitempty" name:"AssetSubnetId"`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetSubnetName *string `json:"AssetSubnetName,omitnil,omitempty" name:"AssetSubnetName"`

	// vpc名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetVpcName *string `json:"AssetVpcName,omitnil,omitempty" name:"AssetVpcName"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`

	// 负载均衡实例的网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的vip列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitnil,omitempty" name:"LoadBalancerVips"`

	// ipv6信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetIpv6 []*string `json:"AssetIpv6,omitnil,omitempty" name:"AssetIpv6"`

	// ssh端口暴露风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSHRisk *string `json:"SSHRisk,omitnil,omitempty" name:"SSHRisk"`

	// rdp端口暴露风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	RDPRisk *string `json:"RDPRisk,omitnil,omitempty" name:"RDPRisk"`

	// 资产失陷事件风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventRisk *string `json:"EventRisk,omitnil,omitempty" name:"EventRisk"`
}

type AssetDetail struct {
	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 所属网络
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 主机类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 主机状态
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 主机IP-公网
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 引擎版本
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 标识
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 标签
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 内网IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 负载均衡示例的vip列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitnil,omitempty" name:"LoadBalancerVips"`

	// 账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationDate *string `json:"CreationDate,omitnil,omitempty" name:"CreationDate"`

	// 访问域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 资产唯一id
	AssetUniqid *string `json:"AssetUniqid,omitnil,omitempty" name:"AssetUniqid"`

	// 关联实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 配置硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 云硬盘/证书状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetStatus *string `json:"AssetStatus,omitnil,omitempty" name:"AssetStatus"`

	// 证书类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 所属项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertEndTime *string `json:"CertEndTime,omitnil,omitempty" name:"CertEndTime"`

	// nosql引擎/版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 主机IP-内网
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 证书有效期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 端口服务数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port []*string `json:"Port,omitnil,omitempty" name:"Port"`

	// 配置风险数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskConfig []*string `json:"RiskConfig,omitnil,omitempty" name:"RiskConfig"`

	// 相关待处理事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 相关待处理漏洞
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vul *string `json:"Vul,omitnil,omitempty" name:"Vul"`

	// 资产发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaAssetDiscoverTime *string `json:"SsaAssetDiscoverTime,omitnil,omitempty" name:"SsaAssetDiscoverTime"`

	// 所属子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetSubnetId *string `json:"AssetSubnetId,omitnil,omitempty" name:"AssetSubnetId"`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetSubnetName *string `json:"AssetSubnetName,omitnil,omitempty" name:"AssetSubnetName"`

	// vpc名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetVpcName *string `json:"AssetVpcName,omitnil,omitempty" name:"AssetVpcName"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// 负载均衡网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// ipv6信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetIpv6 []*string `json:"AssetIpv6,omitnil,omitempty" name:"AssetIpv6"`

	// ssh风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSHRisk *string `json:"SSHRisk,omitnil,omitempty" name:"SSHRisk"`

	// rdp风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	RDPRisk *string `json:"RDPRisk,omitnil,omitempty" name:"RDPRisk"`

	// 安全事件风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventRisk *string `json:"EventRisk,omitnil,omitempty" name:"EventRisk"`

	// 漏洞数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetVulNum *int64 `json:"AssetVulNum,omitnil,omitempty" name:"AssetVulNum"`

	// 资产事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetEventNum *int64 `json:"AssetEventNum,omitnil,omitempty" name:"AssetEventNum"`

	// cspm风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCspmRiskNum *int64 `json:"AssetCspmRiskNum,omitnil,omitempty" name:"AssetCspmRiskNum"`

	// 资产删除时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaAssetDeleteTime *string `json:"SsaAssetDeleteTime,omitnil,omitempty" name:"SsaAssetDeleteTime"`

	// 费用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetRegionName *string `json:"AssetRegionName,omitnil,omitempty" name:"AssetRegionName"`

	// vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetVpcid *string `json:"AssetVpcid,omitnil,omitempty" name:"AssetVpcid"`
}

type AssetList struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 资产数组
	List []*Asset `json:"List,omitnil,omitempty" name:"List"`
}

type AssetQueryFilter struct {
	// 查询参数
	Filter []*QueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询连接符，1 and  ，2 or
	Logic *uint64 `json:"Logic,omitnil,omitempty" name:"Logic"`
}

type AssetTypeStatistic struct {
	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 统计计数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCount *uint64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`
}

type AttackEvent struct {
	// 来源ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaSrcIp *string `json:"SsaSrcIp,omitnil,omitempty" name:"SsaSrcIp"`

	// 目标ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaDstIp *string `json:"SsaDstIp,omitnil,omitempty" name:"SsaDstIp"`

	// 目标省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaDstProvince *string `json:"SsaDstProvince,omitnil,omitempty" name:"SsaDstProvince"`

	// 目标城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaDstCity *string `json:"SsaDstCity,omitnil,omitempty" name:"SsaDstCity"`

	// 目标国家
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaDstCountry *string `json:"SsaDstCountry,omitnil,omitempty" name:"SsaDstCountry"`

	// 来源省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaSrcProvince *string `json:"SsaSrcProvince,omitnil,omitempty" name:"SsaSrcProvince"`

	// 来源国家
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaSrcCountry *string `json:"SsaSrcCountry,omitnil,omitempty" name:"SsaSrcCountry"`

	// 来源城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaSrcCity *string `json:"SsaSrcCity,omitnil,omitempty" name:"SsaSrcCity"`
}

type Bucket struct {
	// key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type CheckAssetItem struct {
	// 检查项下资产组ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 资产组实例id
	Instid *string `json:"Instid,omitnil,omitempty" name:"Instid"`

	// 处置跳转URL
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 检查任务id
	Taskid *string `json:"Taskid,omitnil,omitempty" name:"Taskid"`

	// 检查结果
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 更新时间
	Updatetime *string `json:"Updatetime,omitnil,omitempty" name:"Updatetime"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 是否忽略
	IsIgnore *int64 `json:"IsIgnore,omitnil,omitempty" name:"IsIgnore"`

	// 检查状态
	IsChecked *int64 `json:"IsChecked,omitnil,omitempty" name:"IsChecked"`

	// 资产组信息
	AssetInfo *string `json:"AssetInfo,omitnil,omitempty" name:"AssetInfo"`

	// 资产组ES的_id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 详情
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 备注内容
	Remarks *string `json:"Remarks,omitnil,omitempty" name:"Remarks"`
}

type CheckConfigDetail struct {
	// 检查项Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 检查项名称
	CheckName *string `json:"CheckName,omitnil,omitempty" name:"CheckName"`

	// 检查项内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 检查项处置方案
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 检查项帮助文档
	Doc *string `json:"Doc,omitnil,omitempty" name:"Doc"`

	// 未通过总数
	ErrorCount *int64 `json:"ErrorCount,omitnil,omitempty" name:"ErrorCount"`

	// 是否通过检查
	IsPass *int64 `json:"IsPass,omitnil,omitempty" name:"IsPass"`

	// 通过检查项
	SafeCount *int64 `json:"SafeCount,omitnil,omitempty" name:"SafeCount"`

	// 忽略检查项
	IgnoreCount *int64 `json:"IgnoreCount,omitnil,omitempty" name:"IgnoreCount"`

	// 风险检查项
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 检查项英文
	NameEn *string `json:"NameEn,omitnil,omitempty" name:"NameEn"`

	// 检查项类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// res_count
	ResCount *int64 `json:"ResCount,omitnil,omitempty" name:"ResCount"`

	// 是否忽略
	IsIgnore *int64 `json:"IsIgnore,omitnil,omitempty" name:"IsIgnore"`
}

type ComplianceCheckDetail struct {
	// 检查项ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 检查项类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 检查项类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 不通过总数
	ErrorCount *int64 `json:"ErrorCount,omitnil,omitempty" name:"ErrorCount"`

	// 检查项英文名
	NameEn *string `json:"NameEn,omitnil,omitempty" name:"NameEn"`

	// 检查项名称
	CheckName *string `json:"CheckName,omitnil,omitempty" name:"CheckName"`

	// 检查项处置方式
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 帮助文档
	Doc *string `json:"Doc,omitnil,omitempty" name:"Doc"`

	// 通过总数
	SafeCount *int64 `json:"SafeCount,omitnil,omitempty" name:"SafeCount"`

	// 检查项检查内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 是否通过检测
	IsPass *int64 `json:"IsPass,omitnil,omitempty" name:"IsPass"`

	// 忽略总数
	IgnoreCount *int64 `json:"IgnoreCount,omitnil,omitempty" name:"IgnoreCount"`

	// 风险总数
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 最近一次检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastCheckTime *string `json:"LastCheckTime,omitnil,omitempty" name:"LastCheckTime"`

	// 资产组类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// res_count
	ResCount *int64 `json:"ResCount,omitnil,omitempty" name:"ResCount"`

	// 检查项UUID
	UUID *string `json:"UUID,omitnil,omitempty" name:"UUID"`

	// 标准项
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardItem *string `json:"StandardItem,omitnil,omitempty" name:"StandardItem"`

	// 章节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Chapter *string `json:"Chapter,omitnil,omitempty" name:"Chapter"`

	// 资产类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTypeDesc *string `json:"AssetTypeDesc,omitnil,omitempty" name:"AssetTypeDesc"`

	// 是否忽略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsIgnore *uint64 `json:"IsIgnore,omitnil,omitempty" name:"IsIgnore"`

	// 风险项
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskItem *string `json:"RiskItem,omitnil,omitempty" name:"RiskItem"`

	// 合规检查项完整名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type ConcernInfo struct {
	// 关注点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConcernType *int64 `json:"ConcernType,omitnil,omitempty" name:"ConcernType"`

	// 实体类型 1: 非云上IP，2: 云上IP，3: 域名，4: IP，5: 文件，6: 进程
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntityType *int64 `json:"EntityType,omitnil,omitempty" name:"EntityType"`

	// 关注点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Concern *string `json:"Concern,omitnil,omitempty" name:"Concern"`

	// 最近数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticsCount *int64 `json:"StatisticsCount,omitnil,omitempty" name:"StatisticsCount"`

	// IP国家
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpCountry *string `json:"IpCountry,omitnil,omitempty" name:"IpCountry"`

	// IP省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpProvince *string `json:"IpProvince,omitnil,omitempty" name:"IpProvince"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 置信度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 服务商
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpIsp *string `json:"IpIsp,omitnil,omitempty" name:"IpIsp"`

	// 是否基础设施
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpInfrastructure *string `json:"IpInfrastructure,omitnil,omitempty" name:"IpInfrastructure"`

	// 威胁类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThreatType []*string `json:"ThreatType,omitnil,omitempty" name:"ThreatType"`

	// 威胁团伙
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*string `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 状态威胁情报接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 恶意标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VictimAssetType *string `json:"VictimAssetType,omitnil,omitempty" name:"VictimAssetType"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VictimAssetName *string `json:"VictimAssetName,omitnil,omitempty" name:"VictimAssetName"`

	// 注册者
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainRegistrant *string `json:"DomainRegistrant,omitnil,omitempty" name:"DomainRegistrant"`

	// 注册机构
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainRegisteredInstitution *string `json:"DomainRegisteredInstitution,omitnil,omitempty" name:"DomainRegisteredInstitution"`

	// 注册时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainRegistrationTime *string `json:"DomainRegistrationTime,omitnil,omitempty" name:"DomainRegistrationTime"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileMd5 *string `json:"FileMd5,omitnil,omitempty" name:"FileMd5"`

	// 病毒名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirusName *string `json:"VirusName,omitnil,omitempty" name:"VirusName"`

	// 文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 进程名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcName *string `json:"ProcName,omitnil,omitempty" name:"ProcName"`

	// 进程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pid *string `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 进程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcPath *string `json:"ProcPath,omitnil,omitempty" name:"ProcPath"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcUser *string `json:"ProcUser,omitnil,omitempty" name:"ProcUser"`

	// 已防御
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefendedCount *uint64 `json:"DefendedCount,omitnil,omitempty" name:"DefendedCount"`

	// 仅检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectedCount *uint64 `json:"DetectedCount,omitnil,omitempty" name:"DetectedCount"`

	// 可疑关注点字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchData *string `json:"SearchData,omitnil,omitempty" name:"SearchData"`

	// 可疑关注点字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpCountryIso *string `json:"IpCountryIso,omitnil,omitempty" name:"IpCountryIso"`

	// 可疑关注点字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpProvinceIso *string `json:"IpProvinceIso,omitnil,omitempty" name:"IpProvinceIso"`

	// 可疑关注点字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpCity *string `json:"IpCity,omitnil,omitempty" name:"IpCity"`

	// 可疑关注点字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`
}

type DataCheck struct {
	// 检查项唯一标识符uuid
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 检查项名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 检查项类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 最近一次检查时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastCheckTime *string `json:"LastCheckTime,omitnil,omitempty" name:"LastCheckTime"`

	// 初始未检测状态0, 已通过为1，未通过为2
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 0-未忽略,1-已忽略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsIgnored *uint64 `json:"IsIgnored,omitnil,omitempty" name:"IsIgnored"`

	// 有风险的资源总数，未通过数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCount *uint64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 0-检测中,1-结束检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsChecked *uint64 `json:"IsChecked,omitnil,omitempty" name:"IsChecked"`

	// 总资产数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTotal *int64 `json:"AssetTotal,omitnil,omitempty" name:"AssetTotal"`

	// 备注内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remarks *string `json:"Remarks,omitnil,omitempty" name:"Remarks"`
}

type DataCompliance struct {
	// 等保唯一标识符
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 检查项唯一标识符
	CheckItemId *string `json:"CheckItemId,omitnil,omitempty" name:"CheckItemId"`

	// 检查项名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 检查项资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 检查项类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 检查项类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 检查项标准项
	StandardItem *string `json:"StandardItem,omitnil,omitempty" name:"StandardItem"`

	// 检查项章节号
	Chapter *string `json:"Chapter,omitnil,omitempty" name:"Chapter"`

	// 最近一次检查时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastCheckTime *string `json:"LastCheckTime,omitnil,omitempty" name:"LastCheckTime"`

	// 初始未检测状态0, 已通过为1，未通过为2
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 有风险的资源总数，未通过数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCount *uint64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// 0-检测中,1-结束检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsChecked *uint64 `json:"IsChecked,omitnil,omitempty" name:"IsChecked"`

	// 检查项风险项
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskItem *string `json:"RiskItem,omitnil,omitempty" name:"RiskItem"`

	// 0-未忽略,1-已忽略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsIgnored *uint64 `json:"IsIgnored,omitnil,omitempty" name:"IsIgnored"`

	// 等保检查项完整名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 资产总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTotal *int64 `json:"AssetTotal,omitnil,omitempty" name:"AssetTotal"`

	// 忽略内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remarks *string `json:"Remarks,omitnil,omitempty" name:"Remarks"`
}

// Predefined struct for user
type DescribeAlarmStatRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeAlarmStatRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeAlarmStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmStatResponseParams struct {
	// 威胁告警信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AlarmInfoRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmStatResponseParams `json:"Response"`
}

func (r *DescribeAlarmStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDetailListRequestParams struct {
	// 查询条件，可支持的查询字段：AssetUniqid,AssetName,AssetIpAll,AssetVpcid,Tag
	Filter []*AssetQueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序条件，可支持的排序字段：
	// AssetCspmRiskNum,AssetVulNum,AssetEventNum,SsaAssetDiscoverTime
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 风险标签
	RiskTags []*string `json:"RiskTags,omitnil,omitempty" name:"RiskTags"`

	// 标签
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 页
	PageIndex *uint64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeAssetDetailListRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件，可支持的查询字段：AssetUniqid,AssetName,AssetIpAll,AssetVpcid,Tag
	Filter []*AssetQueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序条件，可支持的排序字段：
	// AssetCspmRiskNum,AssetVulNum,AssetEventNum,SsaAssetDiscoverTime
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 风险标签
	RiskTags []*string `json:"RiskTags,omitnil,omitempty" name:"RiskTags"`

	// 标签
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 页
	PageIndex *uint64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeAssetDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Sorter")
	delete(f, "RiskTags")
	delete(f, "Tags")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDetailListResponseParams struct {
	// 业务数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*AssetDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetDetailListResponseParams `json:"Response"`
}

func (r *DescribeAssetDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDetailRequestParams struct {
	// 查询过滤参数
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type DescribeAssetDetailRequest struct {
	*tchttp.BaseRequest
	
	// 查询过滤参数
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *DescribeAssetDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetDetailResponseParams struct {
	// 资产详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AssetDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetDetailResponseParams `json:"Response"`
}

func (r *DescribeAssetDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetListRequestParams struct {
	// 查询过滤参数
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type DescribeAssetListRequest struct {
	*tchttp.BaseRequest
	
	// 查询过滤参数
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *DescribeAssetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetListResponseParams struct {
	// 资产列表
	AssetList *AssetList `json:"AssetList,omitnil,omitempty" name:"AssetList"`

	// 聚合数据
	AggregationData []*AggregationObj `json:"AggregationData,omitnil,omitempty" name:"AggregationData"`

	// 命名空间数据
	NamespaceData []*string `json:"NamespaceData,omitnil,omitempty" name:"NamespaceData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetListResponseParams `json:"Response"`
}

func (r *DescribeAssetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckConfigAssetListRequestParams struct {
	// 检查项UUID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页列表数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// db搜索条件
	Search []*Filter `json:"Search,omitnil,omitempty" name:"Search"`

	// ES过滤条件
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCheckConfigAssetListRequest struct {
	*tchttp.BaseRequest
	
	// 检查项UUID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页列表数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// db搜索条件
	Search []*Filter `json:"Search,omitnil,omitempty" name:"Search"`

	// ES过滤条件
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCheckConfigAssetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckConfigAssetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Search")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCheckConfigAssetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckConfigAssetListResponseParams struct {
	// 资产列表总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 资产列表项
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAssetsList []*CheckAssetItem `json:"CheckAssetsList,omitnil,omitempty" name:"CheckAssetsList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCheckConfigAssetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCheckConfigAssetListResponseParams `json:"Response"`
}

func (r *DescribeCheckConfigAssetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckConfigAssetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckConfigDetailRequestParams struct {
	// 检查项ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeCheckConfigDetailRequest struct {
	*tchttp.BaseRequest
	
	// 检查项ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeCheckConfigDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckConfigDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCheckConfigDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckConfigDetailResponseParams struct {
	// 检查项详情
	CheckConfigDetail *CheckConfigDetail `json:"CheckConfigDetail,omitnil,omitempty" name:"CheckConfigDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCheckConfigDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCheckConfigDetailResponseParams `json:"Response"`
}

func (r *DescribeCheckConfigDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckConfigDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceAssetListRequestParams struct {
	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 检查项uuid
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 过滤条件
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询条件
	Search []*Filter `json:"Search,omitnil,omitempty" name:"Search"`
}

type DescribeComplianceAssetListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 检查项uuid
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 过滤条件
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询条件
	Search []*Filter `json:"Search,omitnil,omitempty" name:"Search"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Id")
	delete(f, "Filter")
	delete(f, "Search")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceAssetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceAssetListResponseParams struct {
	// 资产组列表
	CheckAssetsList []*CheckAssetItem `json:"CheckAssetsList,omitnil,omitempty" name:"CheckAssetsList"`

	// 资产组列表总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeComplianceDetailRequestParams struct {
	// 检查项uuid
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeComplianceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 检查项uuid
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeComplianceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceDetailResponseParams struct {
	// 合规管理检查项详情
	CheckConfigDetail *ComplianceCheckDetail `json:"CheckConfigDetail,omitnil,omitempty" name:"CheckConfigDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeComplianceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComplianceDetailResponseParams `json:"Response"`
}

func (r *DescribeComplianceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceListRequestParams struct {
	// 搜索过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeComplianceListRequest struct {
	*tchttp.BaseRequest
	
	// 搜索过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeComplianceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeComplianceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeComplianceListResponseParams struct {
	// 检查项列表
	Data []*DataCompliance `json:"Data,omitnil,omitempty" name:"Data"`

	// 总检查资产数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTotalNum *int64 `json:"AssetTotalNum,omitnil,omitempty" name:"AssetTotalNum"`

	// 总检查项
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTotalNum *int64 `json:"ConfigTotalNum,omitnil,omitempty" name:"ConfigTotalNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeComplianceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeComplianceListResponseParams `json:"Response"`
}

func (r *DescribeComplianceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeComplianceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigListRequestParams struct {
	// 搜索过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeConfigListRequest struct {
	*tchttp.BaseRequest
	
	// 搜索过滤条件
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigListResponseParams struct {
	// 检查项列表
	Data []*DataCheck `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigListResponseParams `json:"Response"`
}

func (r *DescribeConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainListRequestParams struct {
	// 起始，从0开始(只支持32位)
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// limit,最大值200(只支持32位)
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资产大类，根据此字段时返回不同的子结构,AssetBasicType(只支持32位)
	AssetBasicType *uint64 `json:"AssetBasicType,omitnil,omitempty" name:"AssetBasicType"`

	// 过滤条件
	Filter []*QueryFilterV3 `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 导出字段
	Field []*string `json:"Field,omitnil,omitempty" name:"Field"`

	// 时间范围(只支持32位)
	TimeRange *uint64 `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// 逻辑字段(只支持32位)
	Logic *uint64 `json:"Logic,omitnil,omitempty" name:"Logic"`

	// 聚合字段  
	GroupByField *string `json:"GroupByField,omitnil,omitempty" name:"GroupByField"`

	// -
	Task *string `json:"Task,omitnil,omitempty" name:"Task"`

	// 0:cfw 1:vss 2.soc 3.waf 4.cwp
	RequestFrom *uint64 `json:"RequestFrom,omitnil,omitempty" name:"RequestFrom"`
}

type DescribeDomainListRequest struct {
	*tchttp.BaseRequest
	
	// 起始，从0开始(只支持32位)
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// limit,最大值200(只支持32位)
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资产大类，根据此字段时返回不同的子结构,AssetBasicType(只支持32位)
	AssetBasicType *uint64 `json:"AssetBasicType,omitnil,omitempty" name:"AssetBasicType"`

	// 过滤条件
	Filter []*QueryFilterV3 `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 导出字段
	Field []*string `json:"Field,omitnil,omitempty" name:"Field"`

	// 时间范围(只支持32位)
	TimeRange *uint64 `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// 逻辑字段(只支持32位)
	Logic *uint64 `json:"Logic,omitnil,omitempty" name:"Logic"`

	// 聚合字段  
	GroupByField *string `json:"GroupByField,omitnil,omitempty" name:"GroupByField"`

	// -
	Task *string `json:"Task,omitnil,omitempty" name:"Task"`

	// 0:cfw 1:vss 2.soc 3.waf 4.cwp
	RequestFrom *uint64 `json:"RequestFrom,omitnil,omitempty" name:"RequestFrom"`
}

func (r *DescribeDomainListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AssetBasicType")
	delete(f, "Filter")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "Field")
	delete(f, "TimeRange")
	delete(f, "Logic")
	delete(f, "GroupByField")
	delete(f, "Task")
	delete(f, "RequestFrom")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainListResponseParams struct {
	// 无
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 无
	DomainInfoCollection []*DomainInfo `json:"DomainInfoCollection,omitnil,omitempty" name:"DomainInfoCollection"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainListResponseParams `json:"Response"`
}

func (r *DescribeDomainListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventDetailRequestParams struct {
	// 事件索引名
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 事件id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 事件来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 事件子类型
	SubEventType *uint64 `json:"SubEventType,omitnil,omitempty" name:"SubEventType"`

	// 事件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeEventDetailRequest struct {
	*tchttp.BaseRequest
	
	// 事件索引名
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 事件id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 事件来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 事件子类型
	SubEventType *uint64 `json:"SubEventType,omitnil,omitempty" name:"SubEventType"`

	// 事件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeEventDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Index")
	delete(f, "Id")
	delete(f, "Source")
	delete(f, "SubEventType")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventDetailResponseParams struct {
	// 事件详情
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventDetailResponseParams `json:"Response"`
}

func (r *DescribeEventDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeakDetectionListRequestParams struct {
	// 筛选条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页码
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeLeakDetectionListRequest struct {
	*tchttp.BaseRequest
	
	// 筛选条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页码
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeLeakDetectionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeakDetectionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Page")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLeakDetectionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeakDetectionListResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据列表
	List []*string `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLeakDetectionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLeakDetectionListResponseParams `json:"Response"`
}

func (r *DescribeLeakDetectionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeakDetectionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMappingResultsRequestParams struct {
	// 过滤条件，FilterKey 取值范围：AssetId，AssetIp，PrivateIp，Protocol，Service，OS，Process，Component，AssetType，Domain，Port，LastMappingTime，MappingType，Disposal，Vpc
	Filter []*AssetQueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序条件，SortKey取值范围：CreateTime，LastMappingTime
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 页码
	PageIndex *uint64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页大小，默认大小20
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeMappingResultsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，FilterKey 取值范围：AssetId，AssetIp，PrivateIp，Protocol，Service，OS，Process，Component，AssetType，Domain，Port，LastMappingTime，MappingType，Disposal，Vpc
	Filter []*AssetQueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序条件，SortKey取值范围：CreateTime，LastMappingTime
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 页码
	PageIndex *uint64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页大小，默认大小20
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeMappingResultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMappingResultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Sorter")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMappingResultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMappingResultsResponseParams struct {
	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *Results `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMappingResultsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMappingResultsResponseParams `json:"Response"`
}

func (r *DescribeMappingResultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMappingResultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSocAlertDetailsRequestParams struct {
	// 告警id
	AlertId *string `json:"AlertId,omitnil,omitempty" name:"AlertId"`

	// 告警时间，取Timestamp字段
	AlertTimestamp *string `json:"AlertTimestamp,omitnil,omitempty" name:"AlertTimestamp"`
}

type DescribeSocAlertDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 告警id
	AlertId *string `json:"AlertId,omitnil,omitempty" name:"AlertId"`

	// 告警时间，取Timestamp字段
	AlertTimestamp *string `json:"AlertTimestamp,omitnil,omitempty" name:"AlertTimestamp"`
}

func (r *DescribeSocAlertDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocAlertDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlertId")
	delete(f, "AlertTimestamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSocAlertDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSocAlertDetailsResponseParams struct {
	// 返回详情数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AlertDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSocAlertDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSocAlertDetailsResponseParams `json:"Response"`
}

func (r *DescribeSocAlertDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocAlertDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSocAlertListRequestParams struct {
	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 1:急需关注 2.重保监控 3.全量告警
	Scenes *int64 `json:"Scenes,omitnil,omitempty" name:"Scenes"`

	// 查询参数
	Filter []*QueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序参数
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 是否导出；默认为否，如量级超过1000，则使用单独的导出接口
	ExportFlag *bool `json:"ExportFlag,omitnil,omitempty" name:"ExportFlag"`
}

type DescribeSocAlertListRequest struct {
	*tchttp.BaseRequest
	
	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 1:急需关注 2.重保监控 3.全量告警
	Scenes *int64 `json:"Scenes,omitnil,omitempty" name:"Scenes"`

	// 查询参数
	Filter []*QueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序参数
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 是否导出；默认为否，如量级超过1000，则使用单独的导出接口
	ExportFlag *bool `json:"ExportFlag,omitnil,omitempty" name:"ExportFlag"`
}

func (r *DescribeSocAlertListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocAlertListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageIndex")
	delete(f, "Scenes")
	delete(f, "Filter")
	delete(f, "Sorter")
	delete(f, "ExportFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSocAlertListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSocAlertListResponseParams struct {
	// 业务数据
	Data *AlertListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSocAlertListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSocAlertListResponseParams `json:"Response"`
}

func (r *DescribeSocAlertListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocAlertListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSocCheckItemListRequestParams struct {
	// 查询参数,可支持的排序字段:Name,Type,AssetType,Level,Standard,IsFree
	Filter []*QueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序参数:无
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 当前页码数据，默认值为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页面索引，默认值为0
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`
}

type DescribeSocCheckItemListRequest struct {
	*tchttp.BaseRequest
	
	// 查询参数,可支持的排序字段:Name,Type,AssetType,Level,Standard,IsFree
	Filter []*QueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序参数:无
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 当前页码数据，默认值为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页面索引，默认值为0
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`
}

func (r *DescribeSocCheckItemListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocCheckItemListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Sorter")
	delete(f, "PageSize")
	delete(f, "PageIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSocCheckItemListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSocCheckItemListResponseParams struct {
	// 检查项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeSocCheckItemListRspRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSocCheckItemListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSocCheckItemListResponseParams `json:"Response"`
}

func (r *DescribeSocCheckItemListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocCheckItemListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSocCheckItemListRspRsp struct {
	// 检查项详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*SocCheckItemV1 `json:"List,omitnil,omitempty" name:"List"`

	// 检查项总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

// Predefined struct for user
type DescribeSocCheckResultListRequestParams struct {
	// 查询参数,可支持的查询参数：
	// Name,Type,AssetType,Result,PloyName,PloyId
	Filter []*QueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序参数,可支持的排序参数：CheckStatus,RiskCount
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 当前页码数据，默认值为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页面索引，默认值为0
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type DescribeSocCheckResultListRequest struct {
	*tchttp.BaseRequest
	
	// 查询参数,可支持的查询参数：
	// Name,Type,AssetType,Result,PloyName,PloyId
	Filter []*QueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 排序参数,可支持的排序参数：CheckStatus,RiskCount
	Sorter []*QuerySort `json:"Sorter,omitnil,omitempty" name:"Sorter"`

	// 当前页码数据，默认值为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页面索引，默认值为0
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *DescribeSocCheckResultListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocCheckResultListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Sorter")
	delete(f, "PageSize")
	delete(f, "PageIndex")
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSocCheckResultListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSocCheckResultListResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeSocCheckResultListRspRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSocCheckResultListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSocCheckResultListResponseParams `json:"Response"`
}

func (r *DescribeSocCheckResultListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocCheckResultListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSocCheckResultListRspRsp struct {
	// 具体检查项详情
	List []*SocCheckResult `json:"List,omitnil,omitempty" name:"List"`

	// 检查结果总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 低危个数
	LowTotal *int64 `json:"LowTotal,omitnil,omitempty" name:"LowTotal"`

	// 中危个数
	MiddleTotal *int64 `json:"MiddleTotal,omitnil,omitempty" name:"MiddleTotal"`

	// 高危个数
	HighTotal *int64 `json:"HighTotal,omitnil,omitempty" name:"HighTotal"`

	// 正常个数
	NormalTotal *int64 `json:"NormalTotal,omitnil,omitempty" name:"NormalTotal"`
}

// Predefined struct for user
type DescribeSocCspmComplianceRequestParams struct {

}

type DescribeSocCspmComplianceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSocCspmComplianceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocCspmComplianceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSocCspmComplianceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSocCspmComplianceResponseParams struct {
	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *SocComplianceInfoResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSocCspmComplianceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSocCspmComplianceResponseParams `json:"Response"`
}

func (r *DescribeSocCspmComplianceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSocCspmComplianceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDetailRequestParams struct {
	// 漏洞唯一标识符
	UniqId *string `json:"UniqId,omitnil,omitempty" name:"UniqId"`

	// 查看详情来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type DescribeVulDetailRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞唯一标识符
	UniqId *string `json:"UniqId,omitnil,omitempty" name:"UniqId"`

	// 查看详情来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
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
	delete(f, "UniqId")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulDetailResponseParams struct {
	// 漏洞类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulType *int64 `json:"VulType,omitnil,omitempty" name:"VulType"`

	// 漏洞子类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubVulType *string `json:"SubVulType,omitnil,omitempty" name:"SubVulType"`

	// cvss分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvssScore *string `json:"CvssScore,omitnil,omitempty" name:"CvssScore"`

	// cvss值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cvss *string `json:"Cvss,omitnil,omitempty" name:"Cvss"`

	// cve编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cve *string `json:"Cve,omitnil,omitempty" name:"Cve"`

	// cnvd编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cnvd *string `json:"Cnvd,omitnil,omitempty" name:"Cnvd"`

	// cnnvd编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cnnvd *string `json:"Cnnvd,omitnil,omitempty" name:"Cnnvd"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 参考
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reference *string `json:"Reference,omitnil,omitempty" name:"Reference"`

	// 修复意见
	// 注意：此字段可能返回 null，表示取不到有效值。
	Repair *string `json:"Repair,omitnil,omitempty" name:"Repair"`

	// 披露时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 漏洞名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 受影响资产唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImpactAsset *string `json:"ImpactAsset,omitnil,omitempty" name:"ImpactAsset"`

	// 受影响资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImpactAssetName *string `json:"ImpactAssetName,omitnil,omitempty" name:"ImpactAssetName"`

	// 受影响资产是否已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAssetDeleted *bool `json:"IsAssetDeleted,omitnil,omitempty" name:"IsAssetDeleted"`

	// 漏洞来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 漏洞URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulUrl *string `json:"VulUrl,omitnil,omitempty" name:"VulUrl"`

	// 资产归属
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaAssetCategory *int64 `json:"SsaAssetCategory,omitnil,omitempty" name:"SsaAssetCategory"`

	// 资产文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulPath *string `json:"VulPath,omitnil,omitempty" name:"VulPath"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeVulListRequestParams struct {
	// 查询过滤参数:(json序列化的结果）
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type DescribeVulListRequest struct {
	*tchttp.BaseRequest
	
	// 查询过滤参数:(json序列化的结果）
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *DescribeVulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulListResponseParams struct {
	// 漏洞列表
	Data *VulList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulListResponseParams `json:"Response"`
}

func (r *DescribeVulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainInfo struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 解析地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResolveAddr []*string `json:"ResolveAddr,omitnil,omitempty" name:"ResolveAddr"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType []*string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 漏洞风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskVulCount *uint64 `json:"RiskVulCount,omitnil,omitempty" name:"RiskVulCount"`

	// 敏感内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	SensitiveCount *uint64 `json:"SensitiveCount,omitnil,omitempty" name:"SensitiveCount"`

	// 挂马暗链
	// 注意：此字段可能返回 null，表示取不到有效值。
	HorseLinkCount *uint64 `json:"HorseLinkCount,omitnil,omitempty" name:"HorseLinkCount"`

	// 网页篡改
	WebModifyCount *uint64 `json:"WebModifyCount,omitnil,omitempty" name:"WebModifyCount"`

	// 上次扫描时间
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// 最近发现时间
	DiscoverTime *string `json:"DiscoverTime,omitnil,omitempty" name:"DiscoverTime"`

	// 扫描次数
	ScanTaskCount *uint64 `json:"ScanTaskCount,omitnil,omitempty" name:"ScanTaskCount"`

	// 端口
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// 弱口令
	WeekPwdCount *uint64 `json:"WeekPwdCount,omitnil,omitempty" name:"WeekPwdCount"`

	// 资产归属
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetLocation *string `json:"AssetLocation,omitnil,omitempty" name:"AssetLocation"`

	// 网络风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkRisk *uint64 `json:"NetworkRisk,omitnil,omitempty" name:"NetworkRisk"`

	// 网络攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAttack *uint64 `json:"NetworkAttack,omitnil,omitempty" name:"NetworkAttack"`

	// bot访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotVisit *uint64 `json:"BotVisit,omitnil,omitempty" name:"BotVisit"`

	// 网络访问
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAccess *uint64 `json:"NetworkAccess,omitnil,omitempty" name:"NetworkAccess"`

	// 资产创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// waf状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafStatus *uint64 `json:"WafStatus,omitnil,omitempty" name:"WafStatus"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId []*string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName []*string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 是否核心资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNotCore *uint64 `json:"IsNotCore,omitnil,omitempty" name:"IsNotCore"`

	// 是否云外资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`
}

type Filter struct {
	// 过滤键的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否需要精确匹配
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type ListDataSaEventPub struct {
	// 时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 安全事件1级分类
	EventType1 *int64 `json:"EventType1,omitnil,omitempty" name:"EventType1"`

	// 安全事件2级分类
	EventType2 *int64 `json:"EventType2,omitnil,omitempty" name:"EventType2"`

	// 安全事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 风险等级
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 安全事件状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 攻击源ip
	SrcIp *string `json:"SrcIp,omitnil,omitempty" name:"SrcIp"`

	// 攻击目标ip
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// 攻击目标端口
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`

	// 受影响资产
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 私有字段和公有字段映射的原始采集数据唯一标识的MD5值
	OldIdMd5 *string `json:"OldIdMd5,omitnil,omitempty" name:"OldIdMd5"`
}

type MappingResult struct {
	// 资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetIp *string `json:"AssetIp,omitnil,omitempty" name:"AssetIp"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 进程
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *string `json:"Process,omitnil,omitempty" name:"Process"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// 测绘时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastMappingTime *string `json:"LastMappingTime,omitnil,omitempty" name:"LastMappingTime"`

	// 处置建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisposalRecommendations *string `json:"DisposalRecommendations,omitnil,omitempty" name:"DisposalRecommendations"`

	// 处置建议详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisposalRecommendationDetails *string `json:"DisposalRecommendationDetails,omitnil,omitempty" name:"DisposalRecommendationDetails"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 测绘状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	MappingStatus *uint64 `json:"MappingStatus,omitnil,omitempty" name:"MappingStatus"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 安全防护状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityStatus []*SecurityStatus `json:"SecurityStatus,omitnil,omitempty" name:"SecurityStatus"`

	// 处置建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisposalRecommendation *int64 `json:"DisposalRecommendation,omitnil,omitempty" name:"DisposalRecommendation"`

	// 测绘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MappingType *string `json:"MappingType,omitnil,omitempty" name:"MappingType"`
}

type ObjDataSaEventPub struct {
	// Count
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// List
	List []*ListDataSaEventPub `json:"List,omitnil,omitempty" name:"List"`
}

type QueryFilter struct {
	// 过滤key
	FilterKey *string `json:"FilterKey,omitnil,omitempty" name:"FilterKey"`

	// 操作符(只支持32位)
	FilterOperatorType *int64 `json:"FilterOperatorType,omitnil,omitempty" name:"FilterOperatorType"`

	// 过滤value
	FilterValue *string `json:"FilterValue,omitnil,omitempty" name:"FilterValue"`
}

type QueryFilterV3 struct {
	// 过滤条件
	Filter *QueryFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 有无子条件
	HasSub *bool `json:"HasSub,omitnil,omitempty" name:"HasSub"`

	// 查询条件
	SubFilters []*QueryFilter `json:"SubFilters,omitnil,omitempty" name:"SubFilters"`

	// 逻辑操作(只支持32位)
	Logic *uint64 `json:"Logic,omitnil,omitempty" name:"Logic"`
}

type QuerySort struct {
	// 排序字段
	SortKey *string `json:"SortKey,omitnil,omitempty" name:"SortKey"`

	// 顺序，1升序2降序
	SortType *int64 `json:"SortType,omitnil,omitempty" name:"SortType"`
}

type Results struct {
	// 测绘类型统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Statistics []*AssetTypeStatistic `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// 测绘结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*MappingResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 测绘任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCount *uint64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 最大测绘任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskMaxCount *uint64 `json:"TaskMaxCount,omitnil,omitempty" name:"TaskMaxCount"`
}

// Predefined struct for user
type SaDivulgeScanRuleMutateRequestParams struct {
	// Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// DivulgeSoure
	DivulgeSoure *string `json:"DivulgeSoure,omitnil,omitempty" name:"DivulgeSoure"`

	// DivulgeSoureUrl
	DivulgeSoureUrl *string `json:"DivulgeSoureUrl,omitnil,omitempty" name:"DivulgeSoureUrl"`

	// RuleName
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// RuleWord
	RuleWord *string `json:"RuleWord,omitnil,omitempty" name:"RuleWord"`

	// ScanStatus
	ScanStatus *string `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// DivulgeType
	DivulgeType *string `json:"DivulgeType,omitnil,omitempty" name:"DivulgeType"`

	// RepairAdvice
	RepairAdvice *string `json:"RepairAdvice,omitnil,omitempty" name:"RepairAdvice"`
}

type SaDivulgeScanRuleMutateRequest struct {
	*tchttp.BaseRequest
	
	// Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// DivulgeSoure
	DivulgeSoure *string `json:"DivulgeSoure,omitnil,omitempty" name:"DivulgeSoure"`

	// DivulgeSoureUrl
	DivulgeSoureUrl *string `json:"DivulgeSoureUrl,omitnil,omitempty" name:"DivulgeSoureUrl"`

	// RuleName
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// RuleWord
	RuleWord *string `json:"RuleWord,omitnil,omitempty" name:"RuleWord"`

	// ScanStatus
	ScanStatus *string `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// DivulgeType
	DivulgeType *string `json:"DivulgeType,omitnil,omitempty" name:"DivulgeType"`

	// RepairAdvice
	RepairAdvice *string `json:"RepairAdvice,omitnil,omitempty" name:"RepairAdvice"`
}

func (r *SaDivulgeScanRuleMutateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaDivulgeScanRuleMutateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "DivulgeSoure")
	delete(f, "DivulgeSoureUrl")
	delete(f, "RuleName")
	delete(f, "RuleWord")
	delete(f, "ScanStatus")
	delete(f, "DivulgeType")
	delete(f, "RepairAdvice")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaDivulgeScanRuleMutateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaDivulgeScanRuleMutateResponseParams struct {
	// Data
	Data *SaDivulgeScanRuleSetList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaDivulgeScanRuleMutateResponse struct {
	*tchttp.BaseResponse
	Response *SaDivulgeScanRuleMutateResponseParams `json:"Response"`
}

func (r *SaDivulgeScanRuleMutateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaDivulgeScanRuleMutateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SaDivulgeScanRuleSetList struct {
	// Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Code
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type SaEventPubRequestParams struct {
	// 受影响资产
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 安全事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 安全事件1级分类，-1:未知 0:全部 1:攻击事件 2:侦查事件 3:僵木蠕毒 4:违规策略
	EventType1 *int64 `json:"EventType1,omitnil,omitempty" name:"EventType1"`

	// 安全事件2级分类，-1:未知 0:全部 1:DDOS事件 2:Web攻击 3:木马 4:异地登录 5:密码破解
	EventType2 *int64 `json:"EventType2,omitnil,omitempty" name:"EventType2"`

	// 风险等级，-1:未知 0:全部 1:低危 2:中危 3:高危 4:严重，可多选，如：1,2
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 安全事件状态，-1:未知 0:全部 1:待处理 2:已处理 3:误报 4:已忽略 5:已知晓 6:已信任
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询起始地址
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 私有字段和公有字段映射的原始采集数据唯一标识的MD5值
	OldIdMd5 *string `json:"OldIdMd5,omitnil,omitempty" name:"OldIdMd5"`
}

type SaEventPubRequest struct {
	*tchttp.BaseRequest
	
	// 受影响资产
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 安全事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 安全事件1级分类，-1:未知 0:全部 1:攻击事件 2:侦查事件 3:僵木蠕毒 4:违规策略
	EventType1 *int64 `json:"EventType1,omitnil,omitempty" name:"EventType1"`

	// 安全事件2级分类，-1:未知 0:全部 1:DDOS事件 2:Web攻击 3:木马 4:异地登录 5:密码破解
	EventType2 *int64 `json:"EventType2,omitnil,omitempty" name:"EventType2"`

	// 风险等级，-1:未知 0:全部 1:低危 2:中危 3:高危 4:严重，可多选，如：1,2
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 安全事件状态，-1:未知 0:全部 1:待处理 2:已处理 3:误报 4:已忽略 5:已知晓 6:已信任
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询起始地址
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 私有字段和公有字段映射的原始采集数据唯一标识的MD5值
	OldIdMd5 *string `json:"OldIdMd5,omitnil,omitempty" name:"OldIdMd5"`
}

func (r *SaEventPubRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaEventPubRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Asset")
	delete(f, "EventName")
	delete(f, "EventType1")
	delete(f, "EventType2")
	delete(f, "Level")
	delete(f, "Status")
	delete(f, "StartTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EndTime")
	delete(f, "OldIdMd5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaEventPubRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaEventPubResponseParams struct {
	// DataSaEventPub
	DataSaEventPub *ObjDataSaEventPub `json:"DataSaEventPub,omitnil,omitempty" name:"DataSaEventPub"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaEventPubResponse struct {
	*tchttp.BaseResponse
	Response *SaEventPubResponseParams `json:"Response"`
}

func (r *SaEventPubResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaEventPubResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityStatus struct {

}

type SocCheckItem struct {
	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelId *string `json:"LevelId,omitnil,omitempty" name:"LevelId"`

	// 成功数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 失败数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailCount *int64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`
}

type SocCheckItemV1 struct {
	// 检查项id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckId *string `json:"CheckId,omitnil,omitempty" name:"CheckId"`

	// 配置要求
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 检查项类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 检查对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 默认风险等级 2:低危 3:中危 4:高危
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 相关规范
	// 注意：此字段可能返回 null，表示取不到有效值。
	Standard *string `json:"Standard,omitnil,omitempty" name:"Standard"`

	// 检查项是否付费 1:免费 2:付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFree *int64 `json:"IsFree,omitnil,omitempty" name:"IsFree"`
}

type SocCheckResult struct {
	// 检查项的uuid
	CheckId *string `json:"CheckId,omitnil,omitempty" name:"CheckId"`

	// 配置要求
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 检查项的类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 检查对象
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 策略名
	PloyName *string `json:"PloyName,omitnil,omitempty" name:"PloyName"`

	// 策略id
	PloyId *int64 `json:"PloyId,omitnil,omitempty" name:"PloyId"`

	// 正常,低危,中危,高危
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 不符合数
	FailAssetNum *int64 `json:"FailAssetNum,omitnil,omitempty" name:"FailAssetNum"`

	// 总数
	TotalAssetNum *int64 `json:"TotalAssetNum,omitnil,omitempty" name:"TotalAssetNum"`

	// 处置建议url链接
	DealUrl *string `json:"DealUrl,omitnil,omitempty" name:"DealUrl"`
}

type SocComplianceInfoResp struct {
	// 合格项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*SocComplianceItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type SocComplianceItem struct {
	// 唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *string `json:"Item,omitnil,omitempty" name:"Item"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardItem *string `json:"StandardItem,omitnil,omitempty" name:"StandardItem"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 产品字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProStr *string `json:"ProStr,omitnil,omitempty" name:"ProStr"`

	// 产品数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Production []*SocProductionItem `json:"Production,omitnil,omitempty" name:"Production"`

	// 配置项数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckItems []*SocCheckItem `json:"CheckItems,omitnil,omitempty" name:"CheckItems"`
}

type SocProductionItem struct {
	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type Tag struct {
	// 数据库标识
	Fid *int64 `json:"Fid,omitnil,omitempty" name:"Fid"`

	// 标签名称字段
	Fname *string `json:"Fname,omitnil,omitempty" name:"Fname"`
}

type VulItem struct {
	// 标识
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 漏洞名称
	VulName *string `json:"VulName,omitnil,omitempty" name:"VulName"`

	// 漏洞类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 风险等级
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 处理状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 发现时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 影响资产数
	ImpactAssetNum *int64 `json:"ImpactAssetNum,omitnil,omitempty" name:"ImpactAssetNum"`

	// 影响资产id
	ImpactAsset *string `json:"ImpactAsset,omitnil,omitempty" name:"ImpactAsset"`

	// 影响资产名称
	ImpactAssetName *string `json:"ImpactAssetName,omitnil,omitempty" name:"ImpactAssetName"`

	// 漏洞描述
	VulDetail *string `json:"VulDetail,omitnil,omitempty" name:"VulDetail"`

	// 参考链接
	VulRefLink *string `json:"VulRefLink,omitnil,omitempty" name:"VulRefLink"`

	// Md5值
	OldIdMd5 *string `json:"OldIdMd5,omitnil,omitempty" name:"OldIdMd5"`

	// 漏洞唯一标识
	UniqId *string `json:"UniqId,omitnil,omitempty" name:"UniqId"`

	// 忽略时间
	OperateTime *string `json:"OperateTime,omitnil,omitempty" name:"OperateTime"`

	// 受影响资产是否下线
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAssetDeleted *string `json:"IsAssetDeleted,omitnil,omitempty" name:"IsAssetDeleted"`

	// 漏洞首次发现时间
	DiscoverTime *string `json:"DiscoverTime,omitnil,omitempty" name:"DiscoverTime"`

	// 主机源信息标识符
	OriginId *uint64 `json:"OriginId,omitnil,omitempty" name:"OriginId"`

	// 资产区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资产所属网络
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vpcid *string `json:"Vpcid,omitnil,omitempty" name:"Vpcid"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 资产子类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetSubType *string `json:"AssetSubType,omitnil,omitempty" name:"AssetSubType"`

	// 资产IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetIpAll []*string `json:"AssetIpAll,omitnil,omitempty" name:"AssetIpAll"`

	// cvm类型的公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// cvm类型的内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 漏洞来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulSource *string `json:"VulSource,omitnil,omitempty" name:"VulSource"`

	// 影响URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectedUrl *string `json:"AffectedUrl,omitnil,omitempty" name:"AffectedUrl"`

	// 资产归属
	// 注意：此字段可能返回 null，表示取不到有效值。
	SsaAssetCategory *int64 `json:"SsaAssetCategory,omitnil,omitempty" name:"SsaAssetCategory"`

	// 影响url
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulUrl *string `json:"VulUrl,omitnil,omitempty" name:"VulUrl"`

	// 是否扫描
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOpen *bool `json:"IsOpen,omitnil,omitempty" name:"IsOpen"`

	// 御知主机id
	// 注意：此字段可能返回 null，表示取不到有效值。
	YzHostId *uint64 `json:"YzHostId,omitnil,omitempty" name:"YzHostId"`

	// 漏洞描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulRepairPlan *string `json:"VulRepairPlan,omitnil,omitempty" name:"VulRepairPlan"`

	// 漏洞文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulPath *string `json:"VulPath,omitnil,omitempty" name:"VulPath"`
}

type VulList struct {
	// 列表
	List []*VulItem `json:"List,omitnil,omitempty" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}