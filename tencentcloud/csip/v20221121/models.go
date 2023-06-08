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
}

type Tag struct {
	// 标签名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签内容
	Value *string `json:"Value,omitempty" name:"Value"`
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
}

type WhereFilter struct {
	// 过滤的项
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤的值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 精确匹配填 7 模糊匹配填9 ， 兼容 中台定的结构
	OperatorType *int64 `json:"OperatorType,omitempty" name:"OperatorType"`
}