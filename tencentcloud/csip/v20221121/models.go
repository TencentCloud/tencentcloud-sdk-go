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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// vpc-name
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitnil" name:"AssetName"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	Os *string `json:"Os,omitnil" name:"Os"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil" name:"AssetType"`

	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`

	// 账号数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountNum *uint64 `json:"AccountNum,omitnil" name:"AccountNum"`

	// 端口数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortNum *uint64 `json:"PortNum,omitnil" name:"PortNum"`

	// 进程数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessNum *uint64 `json:"ProcessNum,omitnil" name:"ProcessNum"`

	// 软件应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftApplicationNum *uint64 `json:"SoftApplicationNum,omitnil" name:"SoftApplicationNum"`

	// 数据库数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseNum *uint64 `json:"DatabaseNum,omitnil" name:"DatabaseNum"`

	// Web应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebApplicationNum *uint64 `json:"WebApplicationNum,omitnil" name:"WebApplicationNum"`

	// 服务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceNum *uint64 `json:"ServiceNum,omitnil" name:"ServiceNum"`

	// web框架数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebFrameworkNum *uint64 `json:"WebFrameworkNum,omitnil" name:"WebFrameworkNum"`

	// Web站点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSiteNum *uint64 `json:"WebSiteNum,omitnil" name:"WebSiteNum"`

	// Jar包数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	JarPackageNum *uint64 `json:"JarPackageNum,omitnil" name:"JarPackageNum"`

	// 启动服务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartServiceNum *uint64 `json:"StartServiceNum,omitnil" name:"StartServiceNum"`

	// 计划任务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduledTaskNum *uint64 `json:"ScheduledTaskNum,omitnil" name:"ScheduledTaskNum"`

	// 环境变量数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentVariableNum *uint64 `json:"EnvironmentVariableNum,omitnil" name:"EnvironmentVariableNum"`

	// 内核模块数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	KernelModuleNum *uint64 `json:"KernelModuleNum,omitnil" name:"KernelModuleNum"`

	// 系统安装包数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemInstallationPackageNum *uint64 `json:"SystemInstallationPackageNum,omitnil" name:"SystemInstallationPackageNum"`

	// 剩余防护时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurplusProtectDay *uint64 `json:"SurplusProtectDay,omitnil" name:"SurplusProtectDay"`

	// 客户端是否安装  1 已安装 0 未安装
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPStatus *uint64 `json:"CWPStatus,omitnil" name:"CWPStatus"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// 防护等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectLevel *string `json:"ProtectLevel,omitnil" name:"ProtectLevel"`

	// 防护时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedDay *uint64 `json:"ProtectedDay,omitnil" name:"ProtectedDay"`
}

type AssetClusterPod struct {
	// 租户id
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 租户uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 租户昵称
	Nick *string `json:"Nick,omitnil" name:"Nick"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// pod id
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`

	// pod名称
	AssetName *string `json:"AssetName,omitnil" name:"AssetName"`

	// pod创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCreateTime *string `json:"InstanceCreateTime,omitnil" name:"InstanceCreateTime"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 主机id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineId *string `json:"MachineId,omitnil" name:"MachineId"`

	// 主机名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineName *string `json:"MachineName,omitnil" name:"MachineName"`

	// pod ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodIp *string `json:"PodIp,omitnil" name:"PodIp"`

	// 关联service数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCount *int64 `json:"ServiceCount,omitnil" name:"ServiceCount"`

	// 关联容器数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerCount *int64 `json:"ContainerCount,omitnil" name:"ContainerCount"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 是否核心：1:核心，2:非核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *int64 `json:"IsCore,omitnil" name:"IsCore"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil" name:"IsNewAsset"`
}

type AssetInfoDetail struct {
	// 用户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppID *string `json:"AppID,omitnil" name:"AppID"`

	// CVE编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVEId *string `json:"CVEId,omitnil" name:"CVEId"`

	// 是扫描，0默认未扫描，1正在扫描，2扫描完成，3扫描出错
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsScan *int64 `json:"IsScan,omitnil" name:"IsScan"`

	// 影响资产数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfluenceAsset *int64 `json:"InfluenceAsset,omitnil" name:"InfluenceAsset"`

	// 未修复资产数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotRepairAsset *int64 `json:"NotRepairAsset,omitnil" name:"NotRepairAsset"`

	// 未防护资产数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotProtectAsset *int64 `json:"NotProtectAsset,omitnil" name:"NotProtectAsset"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskPercent *int64 `json:"TaskPercent,omitnil" name:"TaskPercent"`

	// 任务时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTime *int64 `json:"TaskTime,omitnil" name:"TaskTime"`

	// 扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTime *string `json:"ScanTime,omitnil" name:"ScanTime"`
}

type AssetViewPortRisk struct {
	// 端口
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitnil" name:"AffectAsset"`

	// 风险等级
	Level *string `json:"Level,omitnil" name:"Level"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 组件
	Component *string `json:"Component,omitnil" name:"Component"`

	// 服务
	Service *string `json:"Service,omitnil" name:"Service"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 处置建议,0保持现状、1限制访问、2封禁端口
	Suggestion *uint64 `json:"Suggestion,omitnil" name:"Suggestion"`

	// 状态，0未处理、1已处置、2已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 资产唯一id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 前端索引
	Index *string `json:"Index,omitnil" name:"Index"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 用户appid
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitnil" name:"Nick"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 来源
	From *string `json:"From,omitnil" name:"From"`
}

type AssetViewVULRisk struct {
	// 影响资产
	AffectAsset *string `json:"AffectAsset,omitnil" name:"AffectAsset"`

	// 风险等级
	Level *string `json:"Level,omitnil" name:"Level"`

	// 资产类型
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 组件
	Component *string `json:"Component,omitnil" name:"Component"`

	// 服务
	Service *string `json:"Service,omitnil" name:"Service"`

	// 最近识别时间
	RecentTime *string `json:"RecentTime,omitnil" name:"RecentTime"`

	// 首次识别时间
	FirstTime *string `json:"FirstTime,omitnil" name:"FirstTime"`

	// 状态，0未处理、1已处置、2已忽略
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 资产唯一id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 前端索引
	Index *string `json:"Index,omitnil" name:"Index"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 用户appid
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitnil" name:"Nick"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 漏洞类型
	VULType *string `json:"VULType,omitnil" name:"VULType"`

	// 端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 描述
	Describe *string `json:"Describe,omitnil" name:"Describe"`

	// 版本名
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// 相关信息
	References *string `json:"References,omitnil" name:"References"`

	// 版本
	AppVersion *string `json:"AppVersion,omitnil" name:"AppVersion"`

	// 漏洞url
	VULURL *string `json:"VULURL,omitnil" name:"VULURL"`

	// 漏洞名称
	VULName *string `json:"VULName,omitnil" name:"VULName"`

	// cve
	CVE *string `json:"CVE,omitnil" name:"CVE"`

	// 修复建议
	Fix *string `json:"Fix,omitnil" name:"Fix"`

	// pocid
	POCId *string `json:"POCId,omitnil" name:"POCId"`

	// 来源
	From *string `json:"From,omitnil" name:"From"`

	// 主机版本
	CWPVersion *int64 `json:"CWPVersion,omitnil" name:"CWPVersion"`

	// 是否支持修复
	IsSupportRepair *bool `json:"IsSupportRepair,omitnil" name:"IsSupportRepair"`

	// 是否支持扫描
	IsSupportDetect *bool `json:"IsSupportDetect,omitnil" name:"IsSupportDetect"`

	// 实例uuid
	InstanceUUID *string `json:"InstanceUUID,omitnil" name:"InstanceUUID"`

	// 负载
	Payload *string `json:"Payload,omitnil" name:"Payload"`

	// 应急漏洞类型，1-应急漏洞，0-非应急漏洞
	// 注意：此字段可能返回 null，表示取不到有效值。
	EMGCVulType *int64 `json:"EMGCVulType,omitnil" name:"EMGCVulType"`
}

type BugInfoDetail struct {
	// 漏洞编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 漏洞对应pocId
	// 注意：此字段可能返回 null，表示取不到有效值。
	PatchId *string `json:"PatchId,omitnil" name:"PatchId"`

	// 漏洞名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VULName *string `json:"VULName,omitnil" name:"VULName"`

	// 漏洞严重性：high,middle，low，info
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitnil" name:"Level"`

	// cvss评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSSScore *string `json:"CVSSScore,omitnil" name:"CVSSScore"`

	// cve编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVEId *string `json:"CVEId,omitnil" name:"CVEId"`

	// 漏洞标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitnil" name:"Tag"`

	// 漏洞种类，1:web应用，2:系统组件漏洞，3:配置风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	VULCategory *uint64 `json:"VULCategory,omitnil" name:"VULCategory"`

	// 漏洞影响系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImpactOs *string `json:"ImpactOs,omitnil" name:"ImpactOs"`

	// 漏洞影响组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImpactCOMPENT *string `json:"ImpactCOMPENT,omitnil" name:"ImpactCOMPENT"`

	// 漏洞影响版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImpactVersion *string `json:"ImpactVersion,omitnil" name:"ImpactVersion"`

	// 链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reference *string `json:"Reference,omitnil" name:"Reference"`

	// 漏洞描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	VULDescribe *string `json:"VULDescribe,omitnil" name:"VULDescribe"`

	// 修复建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fix *string `json:"Fix,omitnil" name:"Fix"`

	// 产品支持状态，实时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProSupport *uint64 `json:"ProSupport,omitnil" name:"ProSupport"`

	// 是否公开，0为未发布，1为发布
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPublish *uint64 `json:"IsPublish,omitnil" name:"IsPublish"`

	// 释放时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitnil" name:"ReleaseTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 漏洞子类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubCategory *string `json:"SubCategory,omitnil" name:"SubCategory"`
}

type CVMAssetVO struct {
	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitnil" name:"AssetName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 防护状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPStatus *uint64 `json:"CWPStatus,omitnil" name:"CWPStatus"`

	// 资产创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCreateTime *string `json:"AssetCreateTime,omitnil" name:"AssetCreateTime"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 私网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// vpc 名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// appid信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitnil" name:"NickName"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableArea *string `json:"AvailableArea,omitnil" name:"AvailableArea"`

	// 是否核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitnil" name:"IsCore"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 子网名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUuid *string `json:"InstanceUuid,omitnil" name:"InstanceUuid"`

	// qquid
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceQUuid *string `json:"InstanceQUuid,omitnil" name:"InstanceQUuid"`

	// os名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsName *string `json:"OsName,omitnil" name:"OsName"`

	// 分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionCount *uint64 `json:"PartitionCount,omitnil" name:"PartitionCount"`

	// cpu信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPUInfo *string `json:"CPUInfo,omitnil" name:"CPUInfo"`

	// cpu大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPUSize *uint64 `json:"CPUSize,omitnil" name:"CPUSize"`

	// cpu负载
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPULoad *string `json:"CPULoad,omitnil" name:"CPULoad"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemorySize *string `json:"MemorySize,omitnil" name:"MemorySize"`

	// 内存负载
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryLoad *string `json:"MemoryLoad,omitnil" name:"MemoryLoad"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *string `json:"DiskSize,omitnil" name:"DiskSize"`

	// 硬盘负载
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskLoad *string `json:"DiskLoad,omitnil" name:"DiskLoad"`

	// 账号数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountCount *string `json:"AccountCount,omitnil" name:"AccountCount"`

	// 进程数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessCount *string `json:"ProcessCount,omitnil" name:"ProcessCount"`

	// 软件应用
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppCount *string `json:"AppCount,omitnil" name:"AppCount"`

	// 监听端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortCount *uint64 `json:"PortCount,omitnil" name:"PortCount"`

	// 网络攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attack *uint64 `json:"Attack,omitnil" name:"Attack"`

	// 网络访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	Access *uint64 `json:"Access,omitnil" name:"Access"`

	// 网络拦截
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intercept *uint64 `json:"Intercept,omitnil" name:"Intercept"`

	// 入向峰值带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InBandwidth *string `json:"InBandwidth,omitnil" name:"InBandwidth"`

	// 出向峰值带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBandwidth *string `json:"OutBandwidth,omitnil" name:"OutBandwidth"`

	// 入向累计流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InFlow *string `json:"InFlow,omitnil" name:"InFlow"`

	// 出向累计流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutFlow *string `json:"OutFlow,omitnil" name:"OutFlow"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 恶意主动外联
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetWorkOut *uint64 `json:"NetWorkOut,omitnil" name:"NetWorkOut"`

	// 端口风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortRisk *uint64 `json:"PortRisk,omitnil" name:"PortRisk"`

	// 漏洞风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil" name:"VulnerabilityRisk"`

	// 配置风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil" name:"ConfigurationRisk"`

	// 扫描任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTask *uint64 `json:"ScanTask,omitnil" name:"ScanTask"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// memberId
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberId *string `json:"MemberId,omitnil" name:"MemberId"`

	// os全称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Os *string `json:"Os,omitnil" name:"Os"`

	// 风险服务暴露
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskExposure *int64 `json:"RiskExposure,omitnil" name:"RiskExposure"`

	// 模拟攻击工具状态。0代表未安装，1代表已安装，2代表已离线
	// 注意：此字段可能返回 null，表示取不到有效值。
	BASAgentStatus *int64 `json:"BASAgentStatus,omitnil" name:"BASAgentStatus"`

	// 1新资产；0 非新资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil" name:"IsNewAsset"`
}

type ClbListenerListInfo struct {
	// 监听器id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitnil" name:"ListenerId"`

	// 监听器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitnil" name:"ListenerName"`

	// 负载均衡Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil" name:"LoadBalancerId"`

	// 负载均衡名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil" name:"LoadBalancerName"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 负载均衡ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	VPort *int64 `json:"VPort,omitnil" name:"VPort"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 私有网络id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil" name:"NumericalVpcId"`

	// 负载均衡类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil" name:"LoadBalancerType"`

	// 监听器域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 负载均衡域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil" name:"LoadBalancerDomain"`
}

// Predefined struct for user
type CreateDomainAndIpRequestParams struct {
	// -
	Content []*string `json:"Content,omitnil" name:"Content"`
}

type CreateDomainAndIpRequest struct {
	*tchttp.BaseRequest
	
	// -
	Content []*string `json:"Content,omitnil" name:"Content"`
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
	Data *int64 `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填
	ScanAssetType *int64 `json:"ScanAssetType,omitnil" name:"ScanAssetType"`

	// 扫描项目；port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitnil" name:"ScanItem"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填
	ScanPlanType *int64 `json:"ScanPlanType,omitnil" name:"ScanPlanType"`

	// 扫描资产信息列表
	Assets []*TaskAssetObject `json:"Assets,omitnil" name:"Assets"`

	// 扫描计划详情
	ScanPlanContent *string `json:"ScanPlanContent,omitnil" name:"ScanPlanContent"`

	// ip/域名/url数组
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil" name:"SelfDefiningAssets"`

	// 高级配置
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil" name:"TaskAdvanceCFG"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式
	TaskMode *int64 `json:"TaskMode,omitnil" name:"TaskMode"`
}

type CreateRiskCenterScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 0-全扫，1-指定资产扫，2-排除资产扫，3-手动填写扫；1和2则Assets字段必填，3则SelfDefiningAssets必填
	ScanAssetType *int64 `json:"ScanAssetType,omitnil" name:"ScanAssetType"`

	// 扫描项目；port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitnil" name:"ScanItem"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义；0,2,3则ScanPlanContent必填
	ScanPlanType *int64 `json:"ScanPlanType,omitnil" name:"ScanPlanType"`

	// 扫描资产信息列表
	Assets []*TaskAssetObject `json:"Assets,omitnil" name:"Assets"`

	// 扫描计划详情
	ScanPlanContent *string `json:"ScanPlanContent,omitnil" name:"ScanPlanContent"`

	// ip/域名/url数组
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil" name:"SelfDefiningAssets"`

	// 高级配置
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil" name:"TaskAdvanceCFG"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式，默认标准模式
	TaskMode *int64 `json:"TaskMode,omitnil" name:"TaskMode"`
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
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 0,任务创建成功；小于0失败；-1为存在资产未认证
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 未认证资产列表
	UnAuthAsset []*string `json:"UnAuthAsset,omitnil" name:"UnAuthAsset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitnil" name:"AssetName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil" name:"AssetType"`

	// vpcid
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// vpc标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 资产创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCreateTime *string `json:"AssetCreateTime,omitnil" name:"AssetCreateTime"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 配置风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil" name:"ConfigurationRisk"`

	// 网络攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attack *uint64 `json:"Attack,omitnil" name:"Attack"`

	// 网络访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	Access *uint64 `json:"Access,omitnil" name:"Access"`

	// 扫描任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTask *uint64 `json:"ScanTask,omitnil" name:"ScanTask"`

	// 用户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 昵称别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitnil" name:"NickName"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 是否核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitnil" name:"IsCore"`

	// 是否新资产: 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil" name:"IsNewAsset"`
}

type DataSearchBug struct {
	// 返回查询状态
	StateCode *string `json:"StateCode,omitnil" name:"StateCode"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataBug []*BugInfoDetail `json:"DataBug,omitnil" name:"DataBug"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataAsset []*AssetInfoDetail `json:"DataAsset,omitnil" name:"DataAsset"`

	// true支持扫描。false不支持扫描
	// 注意：此字段可能返回 null，表示取不到有效值。
	VSSScan *bool `json:"VSSScan,omitnil" name:"VSSScan"`

	// 0不支持，1支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPScan *string `json:"CWPScan,omitnil" name:"CWPScan"`

	// 1支持虚拟补丁，0或空不支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFWPatch *string `json:"CFWPatch,omitnil" name:"CFWPatch"`

	// 0不支持，1支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafPatch *int64 `json:"WafPatch,omitnil" name:"WafPatch"`

	// 0不支持，1支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPFix *int64 `json:"CWPFix,omitnil" name:"CWPFix"`
}

type DbAssetInfo struct {
	// 云防状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFWStatus *uint64 `json:"CFWStatus,omitnil" name:"CFWStatus"`

	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`

	// vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil" name:"AssetType"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 私网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitnil" name:"AssetName"`

	// 云防保护版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFWProtectLevel *uint64 `json:"CFWProtectLevel,omitnil" name:"CFWProtectLevel"`

	// tag信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`
}

// Predefined struct for user
type DeleteDomainAndIpRequestParams struct {
	// -
	Content []*PublicIpDomainListKey `json:"Content,omitnil" name:"Content"`

	// 是否保留路径配置，1：保留，其他：不保留，默认不传为不保留
	RetainPath *int64 `json:"RetainPath,omitnil" name:"RetainPath"`

	// 以后是否忽略该资产，，1：忽略，其他：不忽略，默认不传为忽略
	IgnoreAsset *int64 `json:"IgnoreAsset,omitnil" name:"IgnoreAsset"`
}

type DeleteDomainAndIpRequest struct {
	*tchttp.BaseRequest
	
	// -
	Content []*PublicIpDomainListKey `json:"Content,omitnil" name:"Content"`

	// 是否保留路径配置，1：保留，其他：不保留，默认不传为不保留
	RetainPath *int64 `json:"RetainPath,omitnil" name:"RetainPath"`

	// 以后是否忽略该资产，，1：忽略，其他：不忽略，默认不传为忽略
	IgnoreAsset *int64 `json:"IgnoreAsset,omitnil" name:"IgnoreAsset"`
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
	delete(f, "Content")
	delete(f, "RetainPath")
	delete(f, "IgnoreAsset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainAndIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainAndIpResponseParams struct {
	// 删除的资产数量
	Data *int64 `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil" name:"TaskIdList"`
}

type DeleteRiskScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id 列表
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil" name:"TaskIdList"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRiskScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskScanTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeCVMAssetInfoRequestParams struct {
	// -
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`
}

type DescribeCVMAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// -
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`
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
	Data *AssetBaseInfoResponse `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeCVMAssetsRequest struct {
	*tchttp.BaseRequest
	
	// -
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CVMAssetVO `json:"Data,omitnil" name:"Data"`

	// 地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionList []*FilterDataObject `json:"RegionList,omitnil" name:"RegionList"`

	// 防护状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil" name:"DefenseStatusList"`

	// vpc枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcList []*FilterDataObject `json:"VpcList,omitnil" name:"VpcList"`

	// 资产类型枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil" name:"AssetTypeList"`

	// 操作系统枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemTypeList []*FilterDataObject `json:"SystemTypeList,omitnil" name:"SystemTypeList"`

	// ip列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitnil" name:"IpTypeList"`

	// appid列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil" name:"AppIdList"`

	// 可用区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneList []*FilterDataObject `json:"ZoneList,omitnil" name:"ZoneList"`

	// os列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsList []*FilterDataObject `json:"OsList,omitnil" name:"OsList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeClusterPodAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	Data []*AssetClusterPod `json:"Data,omitnil" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 集群pod状态枚举
	PodStatusList []*FilterDataObject `json:"PodStatusList,omitnil" name:"PodStatusList"`

	// 命名空间枚举
	NamespaceList []*FilterDataObject `json:"NamespaceList,omitnil" name:"NamespaceList"`

	// 地域枚举
	RegionList []*FilterDataObject `json:"RegionList,omitnil" name:"RegionList"`

	// 租户枚举
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil" name:"AppIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`
}

type DescribeDbAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// 资产id
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`
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
	Data *DbAssetInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeDbAssetsRequest struct {
	*tchttp.BaseRequest
	
	// -
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 资产总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DBAssetVO `json:"Data,omitnil" name:"Data"`

	// 地域枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionList []*FilterDataObject `json:"RegionList,omitnil" name:"RegionList"`

	// 资产类型枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil" name:"AssetTypeList"`

	// Vpc枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcList []*FilterDataObject `json:"VpcList,omitnil" name:"VpcList"`

	// Appid枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil" name:"AppIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeDomainAssetsRequest struct {
	*tchttp.BaseRequest
	
	// -
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DomainAssetVO `json:"Data,omitnil" name:"Data"`

	// 防护状态列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil" name:"DefenseStatusList"`

	// 资产归属地列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitnil" name:"AssetLocationList"`

	// 资产类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceTypeList []*FilterDataObject `json:"SourceTypeList,omitnil" name:"SourceTypeList"`

	// 地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionList []*FilterDataObject `json:"RegionList,omitnil" name:"RegionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeListenerListRequest struct {
	*tchttp.BaseRequest
	
	// -
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ClbListenerListInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribePublicIpAssetsRequest struct {
	*tchttp.BaseRequest
	
	// filte过滤条件
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	Data []*IpAssetListVO `json:"Data,omitnil" name:"Data"`

	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 资产归属地
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitnil" name:"AssetLocationList"`

	// ip列表枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitnil" name:"IpTypeList"`

	// 地域列表枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionList []*FilterDataObject `json:"RegionList,omitnil" name:"RegionList"`

	// 防护枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil" name:"DefenseStatusList"`

	// 资产类型枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil" name:"AssetTypeList"`

	// AppId枚举
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil" name:"AppIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeRiskCenterAssetViewPortRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤内容
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 资产视角的配置风险列表
	Data []*AssetViewPortRisk `json:"Data,omitnil" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil" name:"LevelLists"`

	// 建议列表
	SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitnil" name:"SuggestionLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil" name:"InstanceTypeLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil" name:"FromLists"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeRiskCenterAssetViewVULRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤内容
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 资产视角的漏洞风险列表
	Data []*AssetViewVULRisk `json:"Data,omitnil" name:"Data"`

	// 状态列表
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil" name:"StatusLists"`

	// 危险等级列表
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil" name:"LevelLists"`

	// 来源列表
	FromLists []*FilterDataObject `json:"FromLists,omitnil" name:"FromLists"`

	// 漏洞类型列表
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil" name:"VULTypeLists"`

	// 资产类型列表
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil" name:"InstanceTypeLists"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeScanReportListRequest struct {
	*tchttp.BaseRequest
	
	// 列表过滤条件
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 任务日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ScanTaskInfo `json:"Data,omitnil" name:"Data"`

	// 主账户ID列表
	UINList []*string `json:"UINList,omitnil" name:"UINList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeScanTaskListRequestParams struct {
	// 过滤内容
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeScanTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤内容
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskListResponseParams struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 任务日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ScanTaskInfoList `json:"Data,omitnil" name:"Data"`

	// 主账户ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UINList []*string `json:"UINList,omitnil" name:"UINList"`

	// 体检模式过滤列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskModeList []*FilterDataObject `json:"TaskModeList,omitnil" name:"TaskModeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 无
	Id *string `json:"Id,omitnil" name:"Id"`

	// id=3时传入该参数
	CVEId *string `json:"CVEId,omitnil" name:"CVEId"`
}

type DescribeSearchBugInfoRequest struct {
	*tchttp.BaseRequest
	
	// 无
	Id *string `json:"Id,omitnil" name:"Id"`

	// id=3时传入该参数
	CVEId *string `json:"CVEId,omitnil" name:"CVEId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataSearchBug `json:"Data,omitnil" name:"Data"`

	// 状态值，0：查询成功，非0：查询失败
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 状态信息，success：查询成功，fail：查询失败
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeSubnetAssetsRequestParams struct {
	// 过滤参数
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeSubnetAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	Data []*SubnetAsset `json:"Data,omitnil" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitnil" name:"RegionList"`

	// vpc列表
	VpcList []*FilterDataObject `json:"VpcList,omitnil" name:"VpcList"`

	// appid列表
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil" name:"AppIdList"`

	// 可用区列表
	ZoneList []*FilterDataObject `json:"ZoneList,omitnil" name:"ZoneList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 过滤内容
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeTaskLogListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤内容
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogListResponseParams struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 报告列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TaskLogInfo `json:"Data,omitnil" name:"Data"`

	// 待查看数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotViewNumber *int64 `json:"NotViewNumber,omitnil" name:"NotViewNumber"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 任务报告Id 列表
	ReportItemKeyList []*ReportItemKey `json:"ReportItemKeyList,omitnil" name:"ReportItemKeyList"`

	// 0: 预览， 1: 下载
	Type *int64 `json:"Type,omitnil" name:"Type"`
}

type DescribeTaskLogURLRequest struct {
	*tchttp.BaseRequest
	
	// 任务报告Id 列表
	ReportItemKeyList []*ReportItemKey `json:"ReportItemKeyList,omitnil" name:"ReportItemKeyList"`

	// 0: 预览， 1: 下载
	Type *int64 `json:"Type,omitnil" name:"Type"`
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
	delete(f, "ReportItemKeyList")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogURLResponseParams struct {
	// 返回报告临时下载url
	Data []*TaskLogURL `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeVpcAssetsRequestParams struct {
	// 过滤参数
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
}

type DescribeVpcAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filter *Filter `json:"Filter,omitnil" name:"Filter"`
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
	Data []*Vpc `json:"Data,omitnil" name:"Data"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// vpc列表
	VpcList []*FilterDataObject `json:"VpcList,omitnil" name:"VpcList"`

	// 地域列表
	RegionList []*FilterDataObject `json:"RegionList,omitnil" name:"RegionList"`

	// appid列表
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil" name:"AppIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	AssetId []*string `json:"AssetId,omitnil" name:"AssetId"`

	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName []*string `json:"AssetName,omitnil" name:"AssetName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType []*string `json:"AssetType,omitnil" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region []*string `json:"Region,omitnil" name:"Region"`

	// Waf状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	WAFStatus *uint64 `json:"WAFStatus,omitnil" name:"WAFStatus"`

	// 资产创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCreateTime *string `json:"AssetCreateTime,omitnil" name:"AssetCreateTime"`

	// Appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// 账号id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitnil" name:"NickName"`

	// 是否核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitnil" name:"IsCore"`

	// 是否云上资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCloud *uint64 `json:"IsCloud,omitnil" name:"IsCloud"`

	// 网络攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attack *uint64 `json:"Attack,omitnil" name:"Attack"`

	// 网络访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	Access *uint64 `json:"Access,omitnil" name:"Access"`

	// 网络拦截
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intercept *uint64 `json:"Intercept,omitnil" name:"Intercept"`

	// 入站峰值带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InBandwidth *string `json:"InBandwidth,omitnil" name:"InBandwidth"`

	// 出站峰值带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBandwidth *string `json:"OutBandwidth,omitnil" name:"OutBandwidth"`

	// 入站累计流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InFlow *string `json:"InFlow,omitnil" name:"InFlow"`

	// 出站累计流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutFlow *string `json:"OutFlow,omitnil" name:"OutFlow"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 端口风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortRisk *uint64 `json:"PortRisk,omitnil" name:"PortRisk"`

	// 漏洞风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil" name:"VulnerabilityRisk"`

	// 配置风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil" name:"ConfigurationRisk"`

	// 扫描任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTask *uint64 `json:"ScanTask,omitnil" name:"ScanTask"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDomain *string `json:"SubDomain,omitnil" name:"SubDomain"`

	// 解析ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	SeverIp []*string `json:"SeverIp,omitnil" name:"SeverIp"`

	// boi访问数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotCount *uint64 `json:"BotCount,omitnil" name:"BotCount"`

	// 弱口令风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeakPassword *uint64 `json:"WeakPassword,omitnil" name:"WeakPassword"`

	// 内容风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebContentRisk *uint64 `json:"WebContentRisk,omitnil" name:"WebContentRisk"`

	// tag标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// 关联实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// memberiD
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberId *string `json:"MemberId,omitnil" name:"MemberId"`

	// cc攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	CCAttack *int64 `json:"CCAttack,omitnil" name:"CCAttack"`

	// web攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebAttack *int64 `json:"WebAttack,omitnil" name:"WebAttack"`

	// 风险服务暴露数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceRisk *uint64 `json:"ServiceRisk,omitnil" name:"ServiceRisk"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil" name:"IsNewAsset"`

	// 待确认资产的随机三级域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyDomain *string `json:"VerifyDomain,omitnil" name:"VerifyDomain"`

	// 待确认资产的TXT记录内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyTXTRecord *string `json:"VerifyTXTRecord,omitnil" name:"VerifyTXTRecord"`

	// 待确认资产的认证状态，0-待认证，1-认证成功，2-认证中，3-txt认证失败，4-人工认证失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyStatus *int64 `json:"VerifyStatus,omitnil" name:"VerifyStatus"`

	// bot访问数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotAccessCount *int64 `json:"BotAccessCount,omitnil" name:"BotAccessCount"`
}

type Filter struct {
	// 查询数量限制
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询偏移位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序采用升序还是降序 升:asc 降 desc
	Order *string `json:"Order,omitnil" name:"Order"`

	// 需排序的字段
	By *string `json:"By,omitnil" name:"By"`

	// 过滤的列及内容
	Filters []*WhereFilter `json:"Filters,omitnil" name:"Filters"`

	// 可填无， 日志使用查询时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 可填无， 日志使用查询时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type FilterDataObject struct {
	// 英文翻译
	Value *string `json:"Value,omitnil" name:"Value"`

	// 中文翻译
	Text *string `json:"Text,omitnil" name:"Text"`
}

type IpAssetListVO struct {
	// 资产id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`

	// 资产name
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitnil" name:"AssetName"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 云防状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFWStatus *uint64 `json:"CFWStatus,omitnil" name:"CFWStatus"`

	// 资产创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCreateTime *string `json:"AssetCreateTime,omitnil" name:"AssetCreateTime"`

	// 公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 公网ip类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpType *uint64 `json:"PublicIpType,omitnil" name:"PublicIpType"`

	// vpc
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// vpc名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitnil" name:"NickName"`

	// 核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitnil" name:"IsCore"`

	// 云上
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCloud *uint64 `json:"IsCloud,omitnil" name:"IsCloud"`

	// 网络攻击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attack *uint64 `json:"Attack,omitnil" name:"Attack"`

	// 网络访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	Access *uint64 `json:"Access,omitnil" name:"Access"`

	// 网络拦截
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intercept *uint64 `json:"Intercept,omitnil" name:"Intercept"`

	// 入向带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InBandwidth *string `json:"InBandwidth,omitnil" name:"InBandwidth"`

	// 出向带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBandwidth *string `json:"OutBandwidth,omitnil" name:"OutBandwidth"`

	// 入向流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InFlow *string `json:"InFlow,omitnil" name:"InFlow"`

	// 出向流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutFlow *string `json:"OutFlow,omitnil" name:"OutFlow"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 端口风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortRisk *uint64 `json:"PortRisk,omitnil" name:"PortRisk"`

	// 漏洞风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil" name:"VulnerabilityRisk"`

	// 配置风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil" name:"ConfigurationRisk"`

	// 扫描任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanTask *uint64 `json:"ScanTask,omitnil" name:"ScanTask"`

	// 弱口令
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeakPassword *uint64 `json:"WeakPassword,omitnil" name:"WeakPassword"`

	// 内容风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebContentRisk *uint64 `json:"WebContentRisk,omitnil" name:"WebContentRisk"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// eip主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// memberid信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberId *string `json:"MemberId,omitnil" name:"MemberId"`

	// 风险服务暴露
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskExposure *int64 `json:"RiskExposure,omitnil" name:"RiskExposure"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil" name:"IsNewAsset"`

	// 资产认证状态，0-待认证，1-认证成功，2-认证中，3+-认证失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyStatus *int64 `json:"VerifyStatus,omitnil" name:"VerifyStatus"`
}

type PublicIpDomainListKey struct {
	// 资产值
	Asset *string `json:"Asset,omitnil" name:"Asset"`
}

type ReportItemKey struct {
	// 日志Id列表
	TaskLogList []*string `json:"TaskLogList,omitnil" name:"TaskLogList"`
}

type ScanTaskInfo struct {
	// 任务日志Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务日志名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 任务状态码：1等待开始  2正在扫描  3扫描出错 4扫描完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitnil" name:"Progress"`

	// 对应的展示时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTime *string `json:"TaskTime,omitnil" name:"TaskTime"`

	// 报表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 报表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportName *string `json:"ReportName,omitnil" name:"ReportName"`

	// 扫描计划，0-周期任务,1-立即扫描,2-定时扫描,3-自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanPlan *int64 `json:"ScanPlan,omitnil" name:"ScanPlan"`

	// 关联的资产数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetCount *int64 `json:"AssetCount,omitnil" name:"AssetCount"`

	// APP ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 用户主账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UIN *string `json:"UIN,omitnil" name:"UIN"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`
}

type ScanTaskInfoList struct {
	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// corn
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanPlanContent *string `json:"ScanPlanContent,omitnil" name:"ScanPlanContent"`

	// 0-周期任务,1-立即扫描,2-定时扫描,3-自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 排除扫描资产信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil" name:"SelfDefiningAssets"`

	// 预估时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PredictTime *int64 `json:"PredictTime,omitnil" name:"PredictTime"`

	// 预估完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PredictEndTime *string `json:"PredictEndTime,omitnil" name:"PredictEndTime"`

	// 报告数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportNumber *int64 `json:"ReportNumber,omitnil" name:"ReportNumber"`

	// 资产数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetNumber *int64 `json:"AssetNumber,omitnil" name:"AssetNumber"`

	// 扫描状态 0 初始值  1正在扫描  2扫描完成  3扫描出错
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *int64 `json:"ScanStatus,omitnil" name:"ScanStatus"`

	// 任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *float64 `json:"Percent,omitnil" name:"Percent"`

	// port/poc/weakpass/webcontent/configrisk
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanItem *string `json:"ScanItem,omitnil" name:"ScanItem"`

	// 0-全扫，1-指定资产扫，2-排除资产扫
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanAssetType *int64 `json:"ScanAssetType,omitnil" name:"ScanAssetType"`

	// vss子任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VSSTaskId *string `json:"VSSTaskId,omitnil" name:"VSSTaskId"`

	// cspm子任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CSPMTaskId *string `json:"CSPMTaskId,omitnil" name:"CSPMTaskId"`

	// 主机漏扫子任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPPOCId *string `json:"CWPPOCId,omitnil" name:"CWPPOCId"`

	// 主机基线子任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPBlId *string `json:"CWPBlId,omitnil" name:"CWPBlId"`

	// vss子任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	VSSTaskProcess *int64 `json:"VSSTaskProcess,omitnil" name:"VSSTaskProcess"`

	// cspm子任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	CSPMTaskProcess *uint64 `json:"CSPMTaskProcess,omitnil" name:"CSPMTaskProcess"`

	// 主机漏扫子任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPPOCProcess *int64 `json:"CWPPOCProcess,omitnil" name:"CWPPOCProcess"`

	// 主机基线子任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	CWPBlProcess *uint64 `json:"CWPBlProcess,omitnil" name:"CWPBlProcess"`

	// 异常状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *int64 `json:"ErrorCode,omitnil" name:"ErrorCode"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *string `json:"ErrorInfo,omitnil" name:"ErrorInfo"`

	// 周期任务开始的天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartDay *int64 `json:"StartDay,omitnil" name:"StartDay"`

	// 扫描频率,单位天,1-每天,7-每周,30-月,0-扫描一次
	// 注意：此字段可能返回 null，表示取不到有效值。
	Frequency *int64 `json:"Frequency,omitnil" name:"Frequency"`

	// 完成次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompleteNumber *int64 `json:"CompleteNumber,omitnil" name:"CompleteNumber"`

	// 已完成资产个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompleteAssetNumber *int64 `json:"CompleteAssetNumber,omitnil" name:"CompleteAssetNumber"`

	// 风险数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCount *int64 `json:"RiskCount,omitnil" name:"RiskCount"`

	// 资产
	// 注意：此字段可能返回 null，表示取不到有效值。
	Assets []*TaskAssetObject `json:"Assets,omitnil" name:"Assets"`

	// 用户Appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 用户主账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UIN *string `json:"UIN,omitnil" name:"UIN"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 体检模式，0-标准模式，1-快速模式，2-高级模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskMode *int64 `json:"TaskMode,omitnil" name:"TaskMode"`

	// 扫描来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanFrom *string `json:"ScanFrom,omitnil" name:"ScanFrom"`

	// 是否限免体检0不是，1是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFree *int64 `json:"IsFree,omitnil" name:"IsFree"`

	// 是否可以删除，1-可以，0-不可以，对应多账户管理使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDelete *int64 `json:"IsDelete,omitnil" name:"IsDelete"`
}

// Predefined struct for user
type StopRiskCenterTaskRequestParams struct {
	// 任务id 列表
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil" name:"TaskIdList"`
}

type StopRiskCenterTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id 列表
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil" name:"TaskIdList"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRiskCenterTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRiskCenterTaskResponseParams struct {
	// Status为0， 停止成功
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type SubnetAsset struct {
	// appid
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 资产ID
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`

	// 资产名
	AssetName *string `json:"AssetName,omitnil" name:"AssetName"`

	// 区域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络名
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// 昵称
	Nick *string `json:"Nick,omitnil" name:"Nick"`

	// cidr
	CIDR *string `json:"CIDR,omitnil" name:"CIDR"`

	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// cvm数
	CVM *int64 `json:"CVM,omitnil" name:"CVM"`

	// 可用ip数
	AvailableIp *int64 `json:"AvailableIp,omitnil" name:"AvailableIp"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 配置风险
	ConfigureRisk *int64 `json:"ConfigureRisk,omitnil" name:"ConfigureRisk"`

	// 任务数
	ScanTask *int64 `json:"ScanTask,omitnil" name:"ScanTask"`

	// 最后扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 是否核心
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitnil" name:"IsCore"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil" name:"IsNewAsset"`
}

type Tag struct {
	// 标签名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 标签内容
	Value *string `json:"Value,omitnil" name:"Value"`
}

type TaskAdvanceCFG struct {
	// 漏洞风险高级配置
	VulRisk []*TaskCenterVulRiskInputParam `json:"VulRisk,omitnil" name:"VulRisk"`

	// 弱口令风险高级配置
	WeakPwdRisk []*TaskCenterWeakPwdRiskInputParam `json:"WeakPwdRisk,omitnil" name:"WeakPwdRisk"`

	// 配置风险高级配置
	CFGRisk []*TaskCenterCFGRiskInputParam `json:"CFGRisk,omitnil" name:"CFGRisk"`
}

type TaskAssetObject struct {
	// 资产名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetName *string `json:"AssetName,omitnil" name:"AssetName"`

	// 	资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 资产分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil" name:"AssetType"`

	// ip/域名/资产id，数据库id等
	Asset *string `json:"Asset,omitnil" name:"Asset"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`
}

type TaskCenterCFGRiskInputParam struct {
	// 检测项ID
	ItemId *string `json:"ItemId,omitnil" name:"ItemId"`

	// 是否开启，0-不开启，1-开启
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`
}

type TaskCenterVulRiskInputParam struct {
	// 风险ID
	RiskId *string `json:"RiskId,omitnil" name:"RiskId"`

	// 是否开启，0-不开启，1-开启
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`
}

type TaskCenterWeakPwdRiskInputParam struct {
	// 检测项ID
	CheckItemId *int64 `json:"CheckItemId,omitnil" name:"CheckItemId"`

	// 是否开启，0-不开启，1-开启
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`
}

type TaskIdListKey struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type TaskLogInfo struct {
	// 报告名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskLogName *string `json:"TaskLogName,omitnil" name:"TaskLogName"`

	// 报告ID
	TaskLogId *string `json:"TaskLogId,omitnil" name:"TaskLogId"`

	// 关联资产个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetsNumber *int64 `json:"AssetsNumber,omitnil" name:"AssetsNumber"`

	// 安全风险数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskNumber *int64 `json:"RiskNumber,omitnil" name:"RiskNumber"`

	// 报告生成时间,任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil" name:"Time"`

	// 任务状态码：0 初始值  1正在扫描  2扫描完成  3扫描出错，4停止，5暂停，6该任务已被重启过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 关联任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 扫描开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 任务中心扫描任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCenterTaskId *string `json:"TaskCenterTaskId,omitnil" name:"TaskCenterTaskId"`

	// 租户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 主账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UIN *string `json:"UIN,omitnil" name:"UIN"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`
}

type TaskLogURL struct {
	// 报告下载临时链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	URL *string `json:"URL,omitnil" name:"URL"`

	// 任务报告id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogId *string `json:"LogId,omitnil" name:"LogId"`

	// 任务报告名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskLogName *string `json:"TaskLogName,omitnil" name:"TaskLogName"`

	// APP ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil" name:"AppId"`
}

type Vpc struct {
	// 子网(只支持32位)
	Subnet *uint64 `json:"Subnet,omitnil" name:"Subnet"`

	// 互通vpc(只支持32位)
	ConnectedVpc *uint64 `json:"ConnectedVpc,omitnil" name:"ConnectedVpc"`

	// 资产id
	AssetId *string `json:"AssetId,omitnil" name:"AssetId"`

	// region区域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 云服务器(只支持32位)
	CVM *uint64 `json:"CVM,omitnil" name:"CVM"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// dns域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DNS []*string `json:"DNS,omitnil" name:"DNS"`

	// 资产名称
	AssetName *string `json:"AssetName,omitnil" name:"AssetName"`

	// cidr网段
	CIDR *string `json:"CIDR,omitnil" name:"CIDR"`

	// 资产创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// appid
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 昵称
	Nick *string `json:"Nick,omitnil" name:"Nick"`

	// 是否新资产 1新
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil" name:"IsNewAsset"`

	// 是否核心资产1是 2不是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCore *uint64 `json:"IsCore,omitnil" name:"IsCore"`
}

type WhereFilter struct {
	// 过滤的项
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤的值
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 中台定义：
	// 1等于 2大于 3小于 4大于等于 5小于等于 6不等于 9模糊匹配 13非模糊匹配 14按位与
	// 精确匹配填 7 模糊匹配填9 兼容 中台定的结构
	OperatorType *int64 `json:"OperatorType,omitnil" name:"OperatorType"`
}