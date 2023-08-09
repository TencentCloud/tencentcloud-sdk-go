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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
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
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type AssetBaseInfoResponse struct {
	// vpc-id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc-name
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	Os *string `json:"Os,omitempty" name:"Os"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 账号数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountNum *uint64 `json:"AccountNum,omitempty" name:"AccountNum"`

	// 端口数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortNum *uint64 `json:"PortNum,omitempty" name:"PortNum"`

	// 进程数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessNum *uint64 `json:"ProcessNum,omitempty" name:"ProcessNum"`

	// 软件应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftApplicationNum *uint64 `json:"SoftApplicationNum,omitempty" name:"SoftApplicationNum"`

	// 数据库数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseNum *uint64 `json:"DatabaseNum,omitempty" name:"DatabaseNum"`

	// Web应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebApplicationNum *uint64 `json:"WebApplicationNum,omitempty" name:"WebApplicationNum"`

	// 服务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceNum *uint64 `json:"ServiceNum,omitempty" name:"ServiceNum"`

	// web框架数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebFrameworkNum *uint64 `json:"WebFrameworkNum,omitempty" name:"WebFrameworkNum"`

	// Web站点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSiteNum *uint64 `json:"WebSiteNum,omitempty" name:"WebSiteNum"`

	// Jar包数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	JarPackageNum *uint64 `json:"JarPackageNum,omitempty" name:"JarPackageNum"`

	// 启动服务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartServiceNum *uint64 `json:"StartServiceNum,omitempty" name:"StartServiceNum"`

	// 计划任务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduledTaskNum *uint64 `json:"ScheduledTaskNum,omitempty" name:"ScheduledTaskNum"`

	// 环境变量数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentVariableNum *uint64 `json:"EnvironmentVariableNum,omitempty" name:"EnvironmentVariableNum"`

	// 内核模块数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	KernelModuleNum *uint64 `json:"KernelModuleNum,omitempty" name:"KernelModuleNum"`

	// 系统安装包数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemInstallationPackageNum *uint64 `json:"SystemInstallationPackageNum,omitempty" name:"SystemInstallationPackageNum"`

	// 剩余防护时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurplusProtectDay *uint64 `json:"SurplusProtectDay,omitempty" name:"SurplusProtectDay"`

	// 客户端是否安装  1 已安装 0 未安装
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPStatus *uint64 `json:"CWPStatus,omitempty" name:"CWPStatus"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// 防护等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectLevel *string `json:"ProtectLevel,omitempty" name:"ProtectLevel"`

	// 防护时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedDay *uint64 `json:"ProtectedDay,omitempty" name:"ProtectedDay"`
}

type AssetClusterPod struct {
	// 租户id
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 租户uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 租户昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// pod id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// pod名称
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// pod创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" name:"InstanceCreateTime"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 主机id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineId *string `json:"MachineId,omitempty" name:"MachineId"`

	// 主机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`

	// pod ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`

	// 关联service数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`

	// 关联容器数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerCount *int64 `json:"ContainerCount,omitempty" name:"ContainerCount"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 是否核心：1:核心，2:非核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *int64 `json:"IsCore,omitempty" name:"IsCore"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitempty" name:"IsNewAsset"`
}

type AssetViewPortRisk struct {
	// 端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitempty" name:"AffectAsset"`

	// 风险等级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 组件
	Component *string `json:"Component,omitempty" name:"Component"`

	// 服务
	Service *string `json:"Service,omitempty" name:"Service"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`

	// 处置建议,0保持现状、1限制访问、2封禁端口
	Suggestion *uint64 `json:"Suggestion,omitempty" name:"Suggestion"`

	// 状态，0未处理、1已处置、2已忽略
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 资产唯一id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 前端索引
	Index *string `json:"Index,omitempty" name:"Index"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 用户appid
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 来源
	From *string `json:"From,omitempty" name:"From"`
}

type AssetViewVULRisk struct {
	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitempty" name:"AffectAsset"`

	// 风险等级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 组件
	Component *string `json:"Component,omitempty" name:"Component"`

	// 服务
	Service *string `json:"Service,omitempty" name:"Service"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitempty" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`

	// 状态，0未处理、1已处置、2已忽略
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 资产唯一id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 前端索引
	Index *string `json:"Index,omitempty" name:"Index"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 用户appid
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 漏洞类型
	VULType *string `json:"VULType,omitempty" name:"VULType"`

	// 端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// 描述
	Describe *string `json:"Describe,omitempty" name:"Describe"`

	// 版本名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 相关信息
	References *string `json:"References,omitempty" name:"References"`

	// 版本
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 漏洞url
	VULURL *string `json:"VULURL,omitempty" name:"VULURL"`

	// 漏洞名称
	VULName *string `json:"VULName,omitempty" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitempty" name:"CVE"`

	// 修复建议
	Fix *string `json:"Fix,omitempty" name:"Fix"`

	// pocid
	POCId *string `json:"POCId,omitempty" name:"POCId"`

	// 来源
	From *string `json:"From,omitempty" name:"From"`

	// 主机版本
	CWPVersion *int64 `json:"CWPVersion,omitempty" name:"CWPVersion"`

	// 是否支持修复
	IsSupportRepair *bool `json:"IsSupportRepair,omitempty" name:"IsSupportRepair"`

	// 是否支持扫描
	IsSupportDetect *bool `json:"IsSupportDetect,omitempty" name:"IsSupportDetect"`

	// 实例uuid
	InstanceUUID *string `json:"InstanceUUID,omitempty" name:"InstanceUUID"`

	// 负载
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞
	// 注意：此字段可能返回 null，表示取不到有效值。
	EMGCVulType *int64 `json:"EMGCVulType,omitempty" name:"EMGCVulType"`
}

type CVMAssetVO struct {
	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 防护状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPStatus *uint64 `json:"CWPStatus,omitempty" name:"CWPStatus"`

	// 资产创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 私网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc 名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// appid信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableArea *string `json:"AvailableArea,omitempty" name:"AvailableArea"`

	// 是否核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUuid *string `json:"InstanceUuid,omitempty" name:"InstanceUuid"`

	// qquid
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceQUuid *string `json:"InstanceQUuid,omitempty" name:"InstanceQUuid"`

	// os名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// cpu信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPUInfo *string `json:"CPUInfo,omitempty" name:"CPUInfo"`

	// cpu大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPUSize *uint64 `json:"CPUSize,omitempty" name:"CPUSize"`

	// cpu负载
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPULoad *string `json:"CPULoad,omitempty" name:"CPULoad"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemorySize *string `json:"MemorySize,omitempty" name:"MemorySize"`

	// 内存负载
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryLoad *string `json:"MemoryLoad,omitempty" name:"MemoryLoad"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`

	// 硬盘负载
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskLoad *string `json:"DiskLoad,omitempty" name:"DiskLoad"`

	// 账号数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountCount *string `json:"AccountCount,omitempty" name:"AccountCount"`

	// 进程数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessCount *string `json:"ProcessCount,omitempty" name:"ProcessCount"`

	// 软件应用
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppCount *string `json:"AppCount,omitempty" name:"AppCount"`

	// 监听端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortCount *uint64 `json:"PortCount,omitempty" name:"PortCount"`

	// 网络攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`

	// 网络访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	Access *uint64 `json:"Access,omitempty" name:"Access"`

	// 网络拦截
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intercept *uint64 `json:"Intercept,omitempty" name:"Intercept"`

	// 入向峰值带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InBandwidth *string `json:"InBandwidth,omitempty" name:"InBandwidth"`

	// 出向峰值带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBandwidth *string `json:"OutBandwidth,omitempty" name:"OutBandwidth"`

	// 入向累计流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InFlow *string `json:"InFlow,omitempty" name:"InFlow"`

	// 出向累计流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutFlow *string `json:"OutFlow,omitempty" name:"OutFlow"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 恶意主动外联
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetWorkOut *uint64 `json:"NetWorkOut,omitempty" name:"NetWorkOut"`

	// 端口风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortRisk *uint64 `json:"PortRisk,omitempty" name:"PortRisk"`

	// 漏洞风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitempty" name:"VulnerabilityRisk"`

	// 配置风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitempty" name:"ConfigurationRisk"`

	// 扫描任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTask *uint64 `json:"ScanTask,omitempty" name:"ScanTask"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// memberId
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// os全称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Os *string `json:"Os,omitempty" name:"Os"`

	// 风险服务暴露
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskExposure *int64 `json:"RiskExposure,omitempty" name:"RiskExposure"`

	// 模拟攻击工具状态。0代表未安装，1代表已安装，2代表已离线
	// 注意：此字段可能返回 null，表示取不到有效值。
	BASAgentStatus *int64 `json:"BASAgentStatus,omitempty" name:"BASAgentStatus"`

	// 1新资产；0 非新资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitempty" name:"IsNewAsset"`
}

type ClbListenerListInfo struct {
	// 监听器id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 负载均衡Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 负载均衡ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	VPort *int64 `json:"VPort,omitempty" name:"VPort"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 私有网络id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumericalVpcId *int64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`

	// 负载均衡类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 监听器域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

// Predefined struct for user
type CreateDomainAndIpRequestParams struct {
	// -
	Content []*string `json:"Content,omitempty" name:"Content"`
}

type CreateDomainAndIpRequest struct {
	*tchttp.BaseRequest
	
	// -
	Content []*string `json:"Content,omitempty" name:"Content"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainAndIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainAndIpResponseParams struct {
	// 返回创建成功的数量
	Data *int64 `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填
	ScanAssetType *int64 `json:"ScanAssetType,omitempty" name:"ScanAssetType"`

	// 扫描项目；port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitempty" name:"ScanItem"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填
	ScanPlanType *int64 `json:"ScanPlanType,omitempty" name:"ScanPlanType"`

	// 扫描资产信息列表
	Assets []*TaskAssetObject `json:"Assets,omitempty" name:"Assets"`

	// 扫描计划详情
	ScanPlanContent *string `json:"ScanPlanContent,omitempty" name:"ScanPlanContent"`

	// ip/域名/url数组
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitempty" name:"SelfDefiningAssets"`

	// 高级配置
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitempty" name:"TaskAdvanceCFG"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式
	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
}

type CreateRiskCenterScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填
	ScanAssetType *int64 `json:"ScanAssetType,omitempty" name:"ScanAssetType"`

	// 扫描项目；port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitempty" name:"ScanItem"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填
	ScanPlanType *int64 `json:"ScanPlanType,omitempty" name:"ScanPlanType"`

	// 扫描资产信息列表
	Assets []*TaskAssetObject `json:"Assets,omitempty" name:"Assets"`

	// 扫描计划详情
	ScanPlanContent *string `json:"ScanPlanContent,omitempty" name:"ScanPlanContent"`

	// ip/域名/url数组
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitempty" name:"SelfDefiningAssets"`

	// 高级配置
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitempty" name:"TaskAdvanceCFG"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式
	TaskMode *int64 `json:"TaskMode,omitempty" name:"TaskMode"`
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
	delete(f, "Assets")
	delete(f, "ScanPlanContent")
	delete(f, "SelfDefiningAssets")
	delete(f, "TaskAdvanceCFG")
	delete(f, "TaskMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRiskCenterScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRiskCenterScanTaskResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 0,任务创建成功；小于0失败；-1为存在资产未认证
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 未认证资产列表
	UnAuthAsset []*string `json:"UnAuthAsset,omitempty" name:"UnAuthAsset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type DBAssetVO struct {
	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// vpcid
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 资产创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 配置风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitempty" name:"ConfigurationRisk"`

	// 网络攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`

	// 网络访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	Access *uint64 `json:"Access,omitempty" name:"Access"`

	// 扫描任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTask *uint64 `json:"ScanTask,omitempty" name:"ScanTask"`

	// 用户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 昵称别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 是否核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`

	// 是否新资产: 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitempty" name:"IsNewAsset"`
}

type DbAssetInfo struct {
	// 云防状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFWStatus *uint64 `json:"CFWStatus,omitempty" name:"CFWStatus"`

	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 私网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 云防保护版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFWProtectLevel *uint64 `json:"CFWProtectLevel,omitempty" name:"CFWProtectLevel"`

	// tag信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
}

// Predefined struct for user
type DescribeCVMAssetInfoRequestParams struct {
	// -
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

type DescribeCVMAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// -
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
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
	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AssetBaseInfoResponse `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// -
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeCVMAssetsRequest struct {
	*tchttp.BaseRequest
	
	// -
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCVMAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetsResponseParams struct {
	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CVMAssetVO `json:"Data,omitempty" name:"Data"`

	// 地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`

	// 防护状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitempty" name:"DefenseStatusList"`

	// vpc枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`

	// 资产类型枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`

	// 操作系统枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemTypeList []*FilterDataObject `json:"SystemTypeList,omitempty" name:"SystemTypeList"`

	// ip列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitempty" name:"IpTypeList"`

	// appid列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`

	// 可用区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneList []*FilterDataObject `json:"ZoneList,omitempty" name:"ZoneList"`

	// os列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsList []*FilterDataObject `json:"OsList,omitempty" name:"OsList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeClusterPodAssetsRequestParams struct {
	// 过滤
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeClusterPodAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPodAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPodAssetsResponseParams struct {
	// 列表
	Data []*AssetClusterPod `json:"Data,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群pod状态枚举
	PodStatusList []*FilterDataObject `json:"PodStatusList,omitempty" name:"PodStatusList"`

	// 命名空间枚举
	NamespaceList []*FilterDataObject `json:"NamespaceList,omitempty" name:"NamespaceList"`

	// 地域枚举
	RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`

	// 租户枚举
	AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

type DescribeDbAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// 资产id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DbAssetInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// -
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeDbAssetsRequest struct {
	*tchttp.BaseRequest
	
	// -
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetsResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 资产总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DBAssetVO `json:"Data,omitempty" name:"Data"`

	// 地域枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`

	// 资产类型枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`

	// Vpc枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`

	// Appid枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// -
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeDomainAssetsRequest struct {
	*tchttp.BaseRequest
	
	// -
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainAssetsResponseParams struct {
	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DomainAssetVO `json:"Data,omitempty" name:"Data"`

	// 防护状态列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitempty" name:"DefenseStatusList"`

	// 资产归属地列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitempty" name:"AssetLocationList"`

	// 资产类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceTypeList []*FilterDataObject `json:"SourceTypeList,omitempty" name:"SourceTypeList"`

	// 地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeListenerListRequestParams struct {
	// -
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeListenerListRequest struct {
	*tchttp.BaseRequest
	
	// -
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerListResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ClbListenerListInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribePublicIpAssetsRequestParams struct {
	// filte过滤条件
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribePublicIpAssetsRequest struct {
	*tchttp.BaseRequest
	
	// filte过滤条件
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicIpAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicIpAssetsResponseParams struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*IpAssetListVO `json:"Data,omitempty" name:"Data"`

	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 资产归属地
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitempty" name:"AssetLocationList"`

	// ip列表枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitempty" name:"IpTypeList"`

	// 地域列表枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`

	// 防护枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitempty" name:"DefenseStatusList"`

	// 资产类型枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitempty" name:"AssetTypeList"`

	// AppId枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeRiskCenterAssetViewPortRiskListRequestParams struct {
	// 过滤内容
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeRiskCenterAssetViewPortRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤内容
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewPortRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewPortRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资产视角的配置风险列表
	Data []*AssetViewPortRisk `json:"Data,omitempty" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitempty" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`

	// 建议列表
	SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitempty" name:"SuggestionLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 过滤内容
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeRiskCenterAssetViewVULRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤内容
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewVULRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewVULRiskListResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资产视角的漏洞风险列表
	Data []*AssetViewVULRisk `json:"Data,omitempty" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitempty" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitempty" name:"LevelLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitempty" name:"FromLists"`

	// 漏洞类型列表
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitempty" name:"VULTypeLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitempty" name:"InstanceTypeLists"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeScanReportListRequestParams struct {
	// 列表过滤条件
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeScanReportListRequest struct {
	*tchttp.BaseRequest
	
	// 列表过滤条件
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanReportListResponseParams struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ScanTaskInfo `json:"Data,omitempty" name:"Data"`

	// 主账户ID列表
	UINList []*string `json:"UINList,omitempty" name:"UINList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeSubnetAssetsRequestParams struct {
	// 过滤参数
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeSubnetAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetAssetsResponseParams struct {
	// 列表
	Data []*SubnetAsset `json:"Data,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`

	// vpc列表
	VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`

	// appid列表
	AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`

	// 可用区列表
	ZoneList []*FilterDataObject `json:"ZoneList,omitempty" name:"ZoneList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeVpcAssetsRequestParams struct {
	// 过滤参数
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
}

type DescribeVpcAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filter *Filter `json:"Filter,omitempty" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAssetsResponseParams struct {
	// 列表
	Data []*Vpc `json:"Data,omitempty" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// vpc列表
	VpcList []*FilterDataObject `json:"VpcList,omitempty" name:"VpcList"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitempty" name:"RegionList"`

	// appid列表
	AppIdList []*FilterDataObject `json:"AppIdList,omitempty" name:"AppIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type DomainAssetVO struct {
	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId []*string `json:"AssetId,omitempty" name:"AssetId"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName []*string `json:"AssetName,omitempty" name:"AssetName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType []*string `json:"AssetType,omitempty" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region []*string `json:"Region,omitempty" name:"Region"`

	// Waf状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	WAFStatus *uint64 `json:"WAFStatus,omitempty" name:"WAFStatus"`

	// 资产创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`

	// Appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 账号id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 是否核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`

	// 是否云上资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCloud *uint64 `json:"IsCloud,omitempty" name:"IsCloud"`

	// 网络攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`

	// 网络访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	Access *uint64 `json:"Access,omitempty" name:"Access"`

	// 网络拦截
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intercept *uint64 `json:"Intercept,omitempty" name:"Intercept"`

	// 入站峰值带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InBandwidth *string `json:"InBandwidth,omitempty" name:"InBandwidth"`

	// 出站峰值带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBandwidth *string `json:"OutBandwidth,omitempty" name:"OutBandwidth"`

	// 入站累计流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InFlow *string `json:"InFlow,omitempty" name:"InFlow"`

	// 出站累计流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutFlow *string `json:"OutFlow,omitempty" name:"OutFlow"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 端口风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortRisk *uint64 `json:"PortRisk,omitempty" name:"PortRisk"`

	// 漏洞风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitempty" name:"VulnerabilityRisk"`

	// 配置风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitempty" name:"ConfigurationRisk"`

	// 扫描任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTask *uint64 `json:"ScanTask,omitempty" name:"ScanTask"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 解析ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	SeverIp []*string `json:"SeverIp,omitempty" name:"SeverIp"`

	// boi访问数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotCount *uint64 `json:"BotCount,omitempty" name:"BotCount"`

	// 弱口令风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeakPassword *uint64 `json:"WeakPassword,omitempty" name:"WeakPassword"`

	// 内容风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebContentRisk *uint64 `json:"WebContentRisk,omitempty" name:"WebContentRisk"`

	// tag标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// 关联实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// memberiD
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// cc攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	CCAttack *int64 `json:"CCAttack,omitempty" name:"CCAttack"`

	// web攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebAttack *int64 `json:"WebAttack,omitempty" name:"WebAttack"`

	// 风险服务暴露数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceRisk *uint64 `json:"ServiceRisk,omitempty" name:"ServiceRisk"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitempty" name:"IsNewAsset"`

	// 待确认资产的随机三级域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyDomain *string `json:"VerifyDomain,omitempty" name:"VerifyDomain"`

	// 待确认资产的TXT记录内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyTXTRecord *string `json:"VerifyTXTRecord,omitempty" name:"VerifyTXTRecord"`

	// 待确认资产的认证状态，0-待认证，1-认证成功，2-认证中，3-txt认证失败，4-人工认证失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyStatus *int64 `json:"VerifyStatus,omitempty" name:"VerifyStatus"`
}

type Filter struct {
	// 查询数量限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序采用升序还是降序 升:asc 降 desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 需排序的字段
	By *string `json:"By,omitempty" name:"By"`

	// 过滤的列及内容
	Filters []*WhereFilter `json:"Filters,omitempty" name:"Filters"`

	// 可填无， 日志使用查询时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 可填无， 日志使用查询时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type FilterDataObject struct {
	// 英文翻译
	Value *string `json:"Value,omitempty" name:"Value"`

	// 中文翻译
	Text *string `json:"Text,omitempty" name:"Text"`
}

type IpAssetListVO struct {
	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 资产name
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 云防状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFWStatus *uint64 `json:"CFWStatus,omitempty" name:"CFWStatus"`

	// 资产创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCreateTime *string `json:"AssetCreateTime,omitempty" name:"AssetCreateTime"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 公网ip类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpType *uint64 `json:"PublicIpType,omitempty" name:"PublicIpType"`

	// vpc
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`

	// 云上
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCloud *uint64 `json:"IsCloud,omitempty" name:"IsCloud"`

	// 网络攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attack *uint64 `json:"Attack,omitempty" name:"Attack"`

	// 网络访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	Access *uint64 `json:"Access,omitempty" name:"Access"`

	// 网络拦截
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intercept *uint64 `json:"Intercept,omitempty" name:"Intercept"`

	// 入向带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InBandwidth *string `json:"InBandwidth,omitempty" name:"InBandwidth"`

	// 出向带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBandwidth *string `json:"OutBandwidth,omitempty" name:"OutBandwidth"`

	// 入向流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InFlow *string `json:"InFlow,omitempty" name:"InFlow"`

	// 出向流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutFlow *string `json:"OutFlow,omitempty" name:"OutFlow"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 端口风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortRisk *uint64 `json:"PortRisk,omitempty" name:"PortRisk"`

	// 漏洞风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitempty" name:"VulnerabilityRisk"`

	// 配置风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitempty" name:"ConfigurationRisk"`

	// 扫描任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTask *uint64 `json:"ScanTask,omitempty" name:"ScanTask"`

	// 弱口令
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeakPassword *uint64 `json:"WeakPassword,omitempty" name:"WeakPassword"`

	// 内容风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebContentRisk *uint64 `json:"WebContentRisk,omitempty" name:"WebContentRisk"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// eip主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// memberid信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 风险服务暴露
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskExposure *int64 `json:"RiskExposure,omitempty" name:"RiskExposure"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitempty" name:"IsNewAsset"`

	// 资产认证状态，0-待认证，1-认证成功，2-认证中，3+-认证失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyStatus *int64 `json:"VerifyStatus,omitempty" name:"VerifyStatus"`
}

type ScanTaskInfo struct {
	// 任务日志Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务日志名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务状态码：1等待开始  2正在扫描  3扫描出错 4扫描完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 对应的展示时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTime *string `json:"TaskTime,omitempty" name:"TaskTime"`

	// 报表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportId *string `json:"ReportId,omitempty" name:"ReportId"`

	// 报表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportName *string `json:"ReportName,omitempty" name:"ReportName"`

	// 扫描计划，0-周期任务,1-立即扫描,2-定时扫描,3-自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanPlan *int64 `json:"ScanPlan,omitempty" name:"ScanPlan"`

	// 关联的资产数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCount *int64 `json:"AssetCount,omitempty" name:"AssetCount"`

	// APP ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 用户主账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UIN *string `json:"UIN,omitempty" name:"UIN"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type SubnetAsset struct {
	// appid
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 资产ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// 昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// cidr
	CIDR *string `json:"CIDR,omitempty" name:"CIDR"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// cvm数
	CVM *int64 `json:"CVM,omitempty" name:"CVM"`

	// 可用ip数
	AvailableIp *int64 `json:"AvailableIp,omitempty" name:"AvailableIp"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 配置风险
	ConfigureRisk *int64 `json:"ConfigureRisk,omitempty" name:"ConfigureRisk"`

	// 任务数
	ScanTask *int64 `json:"ScanTask,omitempty" name:"ScanTask"`

	// 最后扫描时间
	LastScanTime *string `json:"LastScanTime,omitempty" name:"LastScanTime"`

	// 是否核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitempty" name:"IsNewAsset"`
}

type Tag struct {
	// 标签名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签内容
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TaskAdvanceCFG struct {
	// 漏洞风险高级配置
	VulRisk []*TaskCenterVulRiskInputParam `json:"VulRisk,omitempty" name:"VulRisk"`

	// 弱口令风险高级配置
	WeakPwdRisk []*TaskCenterWeakPwdRiskInputParam `json:"WeakPwdRisk,omitempty" name:"WeakPwdRisk"`

	// 配置风险高级配置
	CFGRisk []*TaskCenterCFGRiskInputParam `json:"CFGRisk,omitempty" name:"CFGRisk"`
}

type TaskAssetObject struct {
	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 	资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 资产分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// ip/域名/资产id，数据库id等
	Asset *string `json:"Asset,omitempty" name:"Asset"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type TaskCenterCFGRiskInputParam struct {
	// 检测项ID
	ItemId *string `json:"ItemId,omitempty" name:"ItemId"`

	// 是否开启，0-不开启，1-开启
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type TaskCenterVulRiskInputParam struct {
	// 风险ID
	RiskId *string `json:"RiskId,omitempty" name:"RiskId"`

	// 是否开启，0-不开启，1-开启
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type TaskCenterWeakPwdRiskInputParam struct {
	// 检测项ID
	CheckItemId *int64 `json:"CheckItemId,omitempty" name:"CheckItemId"`

	// 是否开启，0-不开启，1-开启
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type Vpc struct {
	// 子网(只支持32位)
	Subnet *uint64 `json:"Subnet,omitempty" name:"Subnet"`

	// 互通vpc(只支持32位)
	ConnectedVpc *uint64 `json:"ConnectedVpc,omitempty" name:"ConnectedVpc"`

	// 资产id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// region区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 云服务器(只支持32位)
	CVM *uint64 `json:"CVM,omitempty" name:"CVM"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// dns域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DNS []*string `json:"DNS,omitempty" name:"DNS"`

	// 资产名称
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// cidr网段
	CIDR *string `json:"CIDR,omitempty" name:"CIDR"`

	// 资产创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// appid
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitempty" name:"IsNewAsset"`

	// 是否核心资产1是 2不是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitempty" name:"IsCore"`
}

type WhereFilter struct {
	// 过滤的项
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤的值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 中台定义：
	// 1等于 2大于 3小于 4大于等于 5小于等于 6不等于 9模糊匹配 13非模糊匹配 14按位与
	// 精确匹配填 7 模糊匹配填9 兼容 中台定的结构
	OperatorType *int64 `json:"OperatorType,omitempty" name:"OperatorType"`
}