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

package v20190103

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddUsersForUserManagerRequestParams struct {
	// 集群字符串ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 用户信息列表
	UserManagerUserList []*UserInfoForUserManager `json:"UserManagerUserList,omitnil" name:"UserManagerUserList"`
}

type AddUsersForUserManagerRequest struct {
	*tchttp.BaseRequest
	
	// 集群字符串ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 用户信息列表
	UserManagerUserList []*UserInfoForUserManager `json:"UserManagerUserList,omitnil" name:"UserManagerUserList"`
}

func (r *AddUsersForUserManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersForUserManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserManagerUserList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUsersForUserManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUsersForUserManagerResponseParams struct {
	// 添加成功的用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessUserList []*string `json:"SuccessUserList,omitnil" name:"SuccessUserList"`

	// 添加失败的用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedUserList []*string `json:"FailedUserList,omitnil" name:"FailedUserList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddUsersForUserManagerResponse struct {
	*tchttp.BaseResponse
	Response *AddUsersForUserManagerResponseParams `json:"Response"`
}

func (r *AddUsersForUserManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersForUserManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AllNodeResourceSpec struct {
	// 描述Master节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterResourceSpec *NodeResourceSpec `json:"MasterResourceSpec,omitnil" name:"MasterResourceSpec"`

	// 描述Core节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreResourceSpec *NodeResourceSpec `json:"CoreResourceSpec,omitnil" name:"CoreResourceSpec"`

	// 描述Taskr节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResourceSpec *NodeResourceSpec `json:"TaskResourceSpec,omitnil" name:"TaskResourceSpec"`

	// 描述Common节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonResourceSpec *NodeResourceSpec `json:"CommonResourceSpec,omitnil" name:"CommonResourceSpec"`

	// Master节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterCount *int64 `json:"MasterCount,omitnil" name:"MasterCount"`

	// Corer节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreCount *int64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// Task节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCount *int64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// Common节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonCount *int64 `json:"CommonCount,omitnil" name:"CommonCount"`
}

type ApplicationStatics struct {
	// 队列名
	Queue *string `json:"Queue,omitnil" name:"Queue"`

	// 用户名
	User *string `json:"User,omitnil" name:"User"`

	// 作业类型
	ApplicationType *string `json:"ApplicationType,omitnil" name:"ApplicationType"`

	// SumMemorySeconds含义
	SumMemorySeconds *int64 `json:"SumMemorySeconds,omitnil" name:"SumMemorySeconds"`

	// SumVCoreSeconds含义
	SumVCoreSeconds *int64 `json:"SumVCoreSeconds,omitnil" name:"SumVCoreSeconds"`

	// SumHDFSBytesWritten（带单位）
	SumHDFSBytesWritten *string `json:"SumHDFSBytesWritten,omitnil" name:"SumHDFSBytesWritten"`

	// SumHDFSBytesRead（待单位）
	SumHDFSBytesRead *string `json:"SumHDFSBytesRead,omitnil" name:"SumHDFSBytesRead"`

	// 作业数
	CountApps *int64 `json:"CountApps,omitnil" name:"CountApps"`
}

type BootstrapAction struct {
	// 脚本位置，支持cos上的文件，且只支持https协议。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 执行时间。
	// resourceAfter 表示在机器资源申请成功后执行。
	// clusterBefore 表示在集群初始化前执行。
	// clusterAfter 表示在集群初始化后执行。
	WhenRun *string `json:"WhenRun,omitnil" name:"WhenRun"`

	// 脚本参数
	Args []*string `json:"Args,omitnil" name:"Args"`
}

type COSSettings struct {
	// COS SecretId
	CosSecretId *string `json:"CosSecretId,omitnil" name:"CosSecretId"`

	// COS SecrectKey
	CosSecretKey *string `json:"CosSecretKey,omitnil" name:"CosSecretKey"`

	// 日志存储在COS上的路径
	LogOnCosPath *string `json:"LogOnCosPath,omitnil" name:"LogOnCosPath"`
}

type CdbInfo struct {
	// 数据库实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 数据库IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 数据库端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 数据库内存规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// 数据库磁盘规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// 服务标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *string `json:"Service,omitnil" name:"Service"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyTime *string `json:"ApplyTime,omitnil" name:"ApplyTime"`

	// 付费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayType *int64 `json:"PayType,omitnil" name:"PayType"`

	// 过期标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireFlag *bool `json:"ExpireFlag,omitnil" name:"ExpireFlag"`

	// 数据库状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 续费标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAutoRenew *int64 `json:"IsAutoRenew,omitnil" name:"IsAutoRenew"`

	// 数据库字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialNo *string `json:"SerialNo,omitnil" name:"SerialNo"`

	// ZoneId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// RegionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`
}

type ClusterExternalServiceInfo struct {
	// 依赖关系，0:被其他集群依赖，1:依赖其他集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependType *int64 `json:"DependType,omitnil" name:"DependType"`

	// 共用组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *string `json:"Service,omitnil" name:"Service"`

	// 共用集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 共用集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *int64 `json:"ClusterStatus,omitnil" name:"ClusterStatus"`
}

type ClusterIDToFlowID struct {
	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 流程id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *uint64 `json:"FlowId,omitnil" name:"FlowId"`
}

type ClusterInstancesInfo struct {
	// ID号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ftitle *string `json:"Ftitle,omitnil" name:"Ftitle"`

	// 集群名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// 地区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 用户APPID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 用户UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 集群VPCID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *int64 `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *int64 `json:"SubnetId,omitnil" name:"SubnetId"`

	// 实例的状态码。取值范围：
	// <li>2：表示集群运行中。</li>
	// <li>3：表示集群创建中。</li>
	// <li>4：表示集群扩容中。</li>
	// <li>5：表示集群增加router节点中。</li>
	// <li>6：表示集群安装组件中。</li>
	// <li>7：表示集群执行命令中。</li>
	// <li>8：表示重启服务中。</li>
	// <li>9：表示进入维护中。</li>
	// <li>10：表示服务暂停中。</li>
	// <li>11：表示退出维护中。</li>
	// <li>12：表示退出暂停中。</li>
	// <li>13：表示配置下发中。</li>
	// <li>14：表示销毁集群中。</li>
	// <li>15：表示销毁core节点中。</li>
	// <li>16：销毁task节点中。</li>
	// <li>17：表示销毁router节点中。</li>
	// <li>18：表示更改webproxy密码中。</li>
	// <li>19：表示集群隔离中。</li>
	// <li>20：表示集群冲正中。</li>
	// <li>21：表示集群回收中。</li>
	// <li>22：表示变配等待中。</li>
	// <li>23：表示集群已隔离。</li>
	// <li>24：表示缩容节点中。</li>
	// <li>33：表示集群等待退费中。</li>
	// <li>34：表示集群已退费。</li>
	// <li>301：表示创建失败。</li>
	// <li>302：表示扩容失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 添加时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`

	// 已经运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunTime *string `json:"RunTime,omitnil" name:"RunTime"`

	// 集群产品配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *EmrProductConfigOutter `json:"Config,omitnil" name:"Config"`

	// 主节点外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterIp *string `json:"MasterIp,omitnil" name:"MasterIp"`

	// EMR版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmrVersion *string `json:"EmrVersion,omitnil" name:"EmrVersion"`

	// 收费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitnil" name:"ChargeType"`

	// 交易版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeVersion *int64 `json:"TradeVersion,omitnil" name:"TradeVersion"`

	// 资源订单ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceOrderId *int64 `json:"ResourceOrderId,omitnil" name:"ResourceOrderId"`

	// 是否计费集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsTradeCluster *int64 `json:"IsTradeCluster,omitnil" name:"IsTradeCluster"`

	// 集群错误状态告警信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmInfo *string `json:"AlarmInfo,omitnil" name:"AlarmInfo"`

	// 是否采用新架构
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWoodpeckerCluster *int64 `json:"IsWoodpeckerCluster,omitnil" name:"IsWoodpeckerCluster"`

	// 元数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaDb *string `json:"MetaDb,omitnil" name:"MetaDb"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Hive元数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HiveMetaDb *string `json:"HiveMetaDb,omitnil" name:"HiveMetaDb"`

	// 集群类型:EMR,CLICKHOUSE,DRUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceClass *string `json:"ServiceClass,omitnil" name:"ServiceClass"`

	// 集群所有节点的别名序列化
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasInfo *string `json:"AliasInfo,omitnil" name:"AliasInfo"`

	// 集群版本Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *int64 `json:"ProductId,omitnil" name:"ProductId"`

	// 地区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 场景名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// 场景化集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneServiceClass *string `json:"SceneServiceClass,omitnil" name:"SceneServiceClass"`

	// 场景化EMR版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneEmrVersion *string `json:"SceneEmrVersion,omitnil" name:"SceneEmrVersion"`

	// 场景化集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// vpc name
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// subnet name
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 集群依赖关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterExternalServiceInfo []*ClusterExternalServiceInfo `json:"ClusterExternalServiceInfo,omitnil" name:"ClusterExternalServiceInfo"`

	// 集群vpcid 字符串类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 子网id 字符串类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopologyInfoList []*TopologyInfo `json:"TopologyInfoList,omitnil" name:"TopologyInfoList"`

	// 是否是跨AZ集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMultiZoneCluster *bool `json:"IsMultiZoneCluster,omitnil" name:"IsMultiZoneCluster"`

	// 是否开通异常节点自动补偿
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCvmReplace *bool `json:"IsCvmReplace,omitnil" name:"IsCvmReplace"`
}

type ClusterSetting struct {
	// 付费方式。
	// PREPAID 包年包月。
	// POSTPAID_BY_HOUR 按量计费，默认方式。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 是否为HA集群。
	SupportHA *bool `json:"SupportHA,omitnil" name:"SupportHA"`

	// 集群所使用的安全组，目前仅支持一个。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 实例位置。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 实例所在VPC。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// 实例登录配置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 实例标签，示例：["{\"TagKey\":\"test-tag1\",\"TagValue\":\"001\"}","{\"TagKey\":\"test-tag2\",\"TagValue\":\"002\"}"]。
	TagSpecification []*string `json:"TagSpecification,omitnil" name:"TagSpecification"`

	// 元数据库配置。
	MetaDB *MetaDbInfo `json:"MetaDB,omitnil" name:"MetaDB"`

	// 实例硬件配置。
	ResourceSpec *JobFlowResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// 是否申请公网IP，默认为false。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitnil" name:"PublicIpAssigned"`

	// 包年包月配置，只对包年包月集群生效。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 集群置放群组。
	DisasterRecoverGroupIds *string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// 是否使用cbs加密。
	CbsEncryptFlag *bool `json:"CbsEncryptFlag,omitnil" name:"CbsEncryptFlag"`

	// 是否使用远程登录，默认为false。
	RemoteTcpDefaultPort *bool `json:"RemoteTcpDefaultPort,omitnil" name:"RemoteTcpDefaultPort"`
}

type ComponentBasicRestartInfo struct {
	// 进程名，必填，如NameNode
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// 操作的IP列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpList []*string `json:"IpList,omitnil" name:"IpList"`
}

type Configuration struct {
	// 配置文件名，支持SPARK、HIVE、HDFS、YARN的部分配置文件自定义。
	Classification *string `json:"Classification,omitnil" name:"Classification"`

	// 配置参数通过KV的形式传入，部分文件支持自定义，可以通过特殊的键"content"传入所有内容。
	Properties *string `json:"Properties,omitnil" name:"Properties"`
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// EMR产品版本名称如EMR-V2.3.0 表示2.3.0版本的EMR， 当前支持产品版本名称查询：[产品版本名称](https://cloud.tencent.com/document/product/589/66338)
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// 是否开启节点高可用。取值范围：
	// <li>true：表示开启节点高可用。</li>
	// <li>false：表示不开启节点高可用。</li>
	EnableSupportHAFlag *bool `json:"EnableSupportHAFlag,omitnil" name:"EnableSupportHAFlag"`

	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：按小时后付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 实例登录设置。通过该参数可以设置所购买节点的登录方式密码或者密钥。
	// <li>设置密钥时，密码仅用于组件原生WebUI快捷入口登录。</li>
	// <li>未设置密钥时，密码用于登录所购节点以及组件原生WebUI快捷入口登录。</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 集群应用场景以及支持部署组件配置
	SceneSoftwareConfig *SceneSoftwareConfig `json:"SceneSoftwareConfig,omitnil" name:"SceneSoftwareConfig"`

	// 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 实例所属安全组的ID，形如sg-xxxxxxxx。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的SecurityGroupId字段来获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil" name:"ScriptBootstrapActionConfig"`

	// 唯一随机标识，时效性为5分钟，需要调用者指定 防止客户端重复创建资源，例如 a9a90aa6-751a-41b6-aad6-fae360632808
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 是否开启集群Master节点公网。取值范围：
	// <li>NEED_MASTER_WAN：表示开启集群Master节点公网。</li>
	// <li>NOT_NEED_MASTER_WAN：表示不开启。</li>默认开启集群Master节点公网。
	NeedMasterWan *string `json:"NeedMasterWan,omitnil" name:"NeedMasterWan"`

	// 是否开启外网远程登录。（在SecurityGroupId不为空时，该参数无效）不填默认为不开启 取值范围：
	// <li>true：表示开启</li>
	// <li>false：表示不开启</li>
	EnableRemoteLoginFlag *bool `json:"EnableRemoteLoginFlag,omitnil" name:"EnableRemoteLoginFlag"`

	// 是否开启Kerberos认证。默认不开启 取值范围：
	// <li>true：表示开启</li>
	// <li>false：表示不开启</li>
	EnableKerberosFlag *bool `json:"EnableKerberosFlag,omitnil" name:"EnableKerberosFlag"`

	// [自定义软件配置](https://cloud.tencent.com/document/product/589/35655?from_cn_redirect=1)
	CustomConf *string `json:"CustomConf,omitnil" name:"CustomConf"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的实例。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/product/213/17810)的返回值中的DisasterRecoverGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// 是否开启集群维度CBS加密。默认不加密 取值范围：
	// <li>true：表示加密</li>
	// <li>false：表示不加密</li>
	EnableCbsEncryptFlag *bool `json:"EnableCbsEncryptFlag,omitnil" name:"EnableCbsEncryptFlag"`

	// MetaDB信息，当MetaType选择EMR_NEW_META时，MetaDataJdbcUrl MetaDataUser MetaDataPass UnifyMetaInstanceId不用填
	// 当MetaType选择EMR_EXIT_META时，填写UnifyMetaInstanceId
	// 当MetaType选择USER_CUSTOM_META时，填写MetaDataJdbcUrl MetaDataUser MetaDataPass
	MetaDBInfo *CustomMetaDBInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// 共享组件信息
	DependService []*DependService `json:"DependService,omitnil" name:"DependService"`

	// 节点资源的规格，有几个可用区，就填几个，按顺序第一个为主可用区，第二个为备可用区，第三个为仲裁可用区。如果没有开启跨AZ，则长度为1即可。
	ZoneResourceConfiguration []*ZoneResourceConfiguration `json:"ZoneResourceConfiguration,omitnil" name:"ZoneResourceConfiguration"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// EMR产品版本名称如EMR-V2.3.0 表示2.3.0版本的EMR， 当前支持产品版本名称查询：[产品版本名称](https://cloud.tencent.com/document/product/589/66338)
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// 是否开启节点高可用。取值范围：
	// <li>true：表示开启节点高可用。</li>
	// <li>false：表示不开启节点高可用。</li>
	EnableSupportHAFlag *bool `json:"EnableSupportHAFlag,omitnil" name:"EnableSupportHAFlag"`

	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：按小时后付费。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 实例登录设置。通过该参数可以设置所购买节点的登录方式密码或者密钥。
	// <li>设置密钥时，密码仅用于组件原生WebUI快捷入口登录。</li>
	// <li>未设置密钥时，密码用于登录所购节点以及组件原生WebUI快捷入口登录。</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 集群应用场景以及支持部署组件配置
	SceneSoftwareConfig *SceneSoftwareConfig `json:"SceneSoftwareConfig,omitnil" name:"SceneSoftwareConfig"`

	// 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 实例所属安全组的ID，形如sg-xxxxxxxx。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的SecurityGroupId字段来获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil" name:"ScriptBootstrapActionConfig"`

	// 唯一随机标识，时效性为5分钟，需要调用者指定 防止客户端重复创建资源，例如 a9a90aa6-751a-41b6-aad6-fae360632808
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 是否开启集群Master节点公网。取值范围：
	// <li>NEED_MASTER_WAN：表示开启集群Master节点公网。</li>
	// <li>NOT_NEED_MASTER_WAN：表示不开启。</li>默认开启集群Master节点公网。
	NeedMasterWan *string `json:"NeedMasterWan,omitnil" name:"NeedMasterWan"`

	// 是否开启外网远程登录。（在SecurityGroupId不为空时，该参数无效）不填默认为不开启 取值范围：
	// <li>true：表示开启</li>
	// <li>false：表示不开启</li>
	EnableRemoteLoginFlag *bool `json:"EnableRemoteLoginFlag,omitnil" name:"EnableRemoteLoginFlag"`

	// 是否开启Kerberos认证。默认不开启 取值范围：
	// <li>true：表示开启</li>
	// <li>false：表示不开启</li>
	EnableKerberosFlag *bool `json:"EnableKerberosFlag,omitnil" name:"EnableKerberosFlag"`

	// [自定义软件配置](https://cloud.tencent.com/document/product/589/35655?from_cn_redirect=1)
	CustomConf *string `json:"CustomConf,omitnil" name:"CustomConf"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的实例。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/product/213/17810)的返回值中的DisasterRecoverGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// 是否开启集群维度CBS加密。默认不加密 取值范围：
	// <li>true：表示加密</li>
	// <li>false：表示不加密</li>
	EnableCbsEncryptFlag *bool `json:"EnableCbsEncryptFlag,omitnil" name:"EnableCbsEncryptFlag"`

	// MetaDB信息，当MetaType选择EMR_NEW_META时，MetaDataJdbcUrl MetaDataUser MetaDataPass UnifyMetaInstanceId不用填
	// 当MetaType选择EMR_EXIT_META时，填写UnifyMetaInstanceId
	// 当MetaType选择USER_CUSTOM_META时，填写MetaDataJdbcUrl MetaDataUser MetaDataPass
	MetaDBInfo *CustomMetaDBInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// 共享组件信息
	DependService []*DependService `json:"DependService,omitnil" name:"DependService"`

	// 节点资源的规格，有几个可用区，就填几个，按顺序第一个为主可用区，第二个为备可用区，第三个为仲裁可用区。如果没有开启跨AZ，则长度为1即可。
	ZoneResourceConfiguration []*ZoneResourceConfiguration `json:"ZoneResourceConfiguration,omitnil" name:"ZoneResourceConfiguration"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductVersion")
	delete(f, "EnableSupportHAFlag")
	delete(f, "InstanceName")
	delete(f, "InstanceChargeType")
	delete(f, "LoginSettings")
	delete(f, "SceneSoftwareConfig")
	delete(f, "InstanceChargePrepaid")
	delete(f, "SecurityGroupIds")
	delete(f, "ScriptBootstrapActionConfig")
	delete(f, "ClientToken")
	delete(f, "NeedMasterWan")
	delete(f, "EnableRemoteLoginFlag")
	delete(f, "EnableKerberosFlag")
	delete(f, "CustomConf")
	delete(f, "Tags")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "EnableCbsEncryptFlag")
	delete(f, "MetaDBInfo")
	delete(f, "DependService")
	delete(f, "ZoneResourceConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：
	// <li>16：表示EMR-V2.3.0。</li>
	// <li>20：表示EMR-V2.5.0。</li>
	// <li>25：表示EMR-V3.1.0。</li>
	// <li>27：表示KAFKA-V1.0.0。</li>
	// <li>30：表示EMR-V2.6.0。</li>
	// <li>33 :   表示EMR-V3.2.1。</li>
	// <li>34 :   表示EMR-V3.3.0。</li>
	// <li>36 :   表示STARROCKS-V1.0.0。</li>
	// <li>37 :   表示EMR-V3.4.0。</li>
	// <li>38 :   表示EMR-V2.7.0。</li>
	// <li>39 :   表示STARROCKS-V1.1.0。</li>
	// <li>41 :   表示DRUID-V1.1.0。</li>
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）对应不同可选组件列表，不同产品版本可选组件列表查询：[组件版本](https://cloud.tencent.com/document/product/589/20279) ；
	// 填写实例值：hive、flink。
	Software []*string `json:"Software,omitnil" name:"Software"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitnil" name:"SupportHA"`

	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 实例登录设置。通过该参数可以设置所购买节点的登录方式密码或者密钥。
	// <li>设置密钥时，密码仅用于组件原生WebUI快捷入口登录。</li>
	// <li>未设置密钥时，密码用于登录所购节点以及组件原生WebUI快捷入口登录。</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// 节点资源的规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// 开启COS访问需要设置的参数。
	COSSettings *COSSettings `json:"COSSettings,omitnil" name:"COSSettings"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 实例所属安全组的ID，形如sg-xxxxxxxx。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的SecurityGroupId字段来获取。
	SgId *string `json:"SgId,omitnil" name:"SgId"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil" name:"PreExecutedFileSettings"`

	// 包年包月实例是否自动续费。取值范围：
	// <li>0：表示不自动续费。</li>
	// <li>1：表示自动续费。</li>
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// 客户端Token。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 是否开启集群Master节点公网。取值范围：
	// <li>NEED_MASTER_WAN：表示开启集群Master节点公网。</li>
	// <li>NOT_NEED_MASTER_WAN：表示不开启。</li>默认开启集群Master节点公网。
	NeedMasterWan *string `json:"NeedMasterWan,omitnil" name:"NeedMasterWan"`

	// 是否需要开启外网远程登录，即22号端口。在SgId不为空时，该参数无效。
	RemoteLoginAtCreate *int64 `json:"RemoteLoginAtCreate,omitnil" name:"RemoteLoginAtCreate"`

	// 是否开启安全集群。0表示不开启，非0表示开启。
	CheckSecurity *int64 `json:"CheckSecurity,omitnil" name:"CheckSecurity"`

	// 访问外部文件系统。
	ExtendFsField *string `json:"ExtendFsField,omitnil" name:"ExtendFsField"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的实例。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/213/15486 ) 的返回值中的SecurityGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// 集群维度CBS加密盘，默认0表示不加密，1表示加密
	CbsEncrypt *uint64 `json:"CbsEncrypt,omitnil" name:"CbsEncrypt"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_DEFAULT_META：表示集群默认创建</li>
	// <li>EMR_EXIST_META：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// 自定义应用角色。
	ApplicationRole *string `json:"ApplicationRole,omitnil" name:"ApplicationRole"`

	// 场景化取值：
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// 共享组件信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil" name:"ExternalService"`

	// 如果为0，则MultiZone、MultiDeployStrategy、MultiZoneSettings是disable的状态，如果为1，则废弃ResourceSpec，使用MultiZoneSettings。
	VersionID *int64 `json:"VersionID,omitnil" name:"VersionID"`

	// true表示开启跨AZ部署；仅为新建集群时的用户参数，后续不支持调整。
	MultiZone *bool `json:"MultiZone,omitnil" name:"MultiZone"`

	// 节点资源的规格，有几个可用区，就填几个，按顺序第一个为主可用区，第二个为备可用区，第三个为仲裁可用区。如果没有开启跨AZ，则长度为1即可。
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil" name:"MultiZoneSettings"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：
	// <li>16：表示EMR-V2.3.0。</li>
	// <li>20：表示EMR-V2.5.0。</li>
	// <li>25：表示EMR-V3.1.0。</li>
	// <li>27：表示KAFKA-V1.0.0。</li>
	// <li>30：表示EMR-V2.6.0。</li>
	// <li>33 :   表示EMR-V3.2.1。</li>
	// <li>34 :   表示EMR-V3.3.0。</li>
	// <li>36 :   表示STARROCKS-V1.0.0。</li>
	// <li>37 :   表示EMR-V3.4.0。</li>
	// <li>38 :   表示EMR-V2.7.0。</li>
	// <li>39 :   表示STARROCKS-V1.1.0。</li>
	// <li>41 :   表示DRUID-V1.1.0。</li>
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）对应不同可选组件列表，不同产品版本可选组件列表查询：[组件版本](https://cloud.tencent.com/document/product/589/20279) ；
	// 填写实例值：hive、flink。
	Software []*string `json:"Software,omitnil" name:"Software"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitnil" name:"SupportHA"`

	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 实例登录设置。通过该参数可以设置所购买节点的登录方式密码或者密钥。
	// <li>设置密钥时，密码仅用于组件原生WebUI快捷入口登录。</li>
	// <li>未设置密钥时，密码用于登录所购节点以及组件原生WebUI快捷入口登录。</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil" name:"LoginSettings"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// 节点资源的规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// 开启COS访问需要设置的参数。
	COSSettings *COSSettings `json:"COSSettings,omitnil" name:"COSSettings"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 实例所属安全组的ID，形如sg-xxxxxxxx。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的SecurityGroupId字段来获取。
	SgId *string `json:"SgId,omitnil" name:"SgId"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil" name:"PreExecutedFileSettings"`

	// 包年包月实例是否自动续费。取值范围：
	// <li>0：表示不自动续费。</li>
	// <li>1：表示自动续费。</li>
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// 客户端Token。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 是否开启集群Master节点公网。取值范围：
	// <li>NEED_MASTER_WAN：表示开启集群Master节点公网。</li>
	// <li>NOT_NEED_MASTER_WAN：表示不开启。</li>默认开启集群Master节点公网。
	NeedMasterWan *string `json:"NeedMasterWan,omitnil" name:"NeedMasterWan"`

	// 是否需要开启外网远程登录，即22号端口。在SgId不为空时，该参数无效。
	RemoteLoginAtCreate *int64 `json:"RemoteLoginAtCreate,omitnil" name:"RemoteLoginAtCreate"`

	// 是否开启安全集群。0表示不开启，非0表示开启。
	CheckSecurity *int64 `json:"CheckSecurity,omitnil" name:"CheckSecurity"`

	// 访问外部文件系统。
	ExtendFsField *string `json:"ExtendFsField,omitnil" name:"ExtendFsField"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的实例。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/213/15486 ) 的返回值中的SecurityGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// 集群维度CBS加密盘，默认0表示不加密，1表示加密
	CbsEncrypt *uint64 `json:"CbsEncrypt,omitnil" name:"CbsEncrypt"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_DEFAULT_META：表示集群默认创建</li>
	// <li>EMR_EXIST_META：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// 自定义应用角色。
	ApplicationRole *string `json:"ApplicationRole,omitnil" name:"ApplicationRole"`

	// 场景化取值：
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// 共享组件信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil" name:"ExternalService"`

	// 如果为0，则MultiZone、MultiDeployStrategy、MultiZoneSettings是disable的状态，如果为1，则废弃ResourceSpec，使用MultiZoneSettings。
	VersionID *int64 `json:"VersionID,omitnil" name:"VersionID"`

	// true表示开启跨AZ部署；仅为新建集群时的用户参数，后续不支持调整。
	MultiZone *bool `json:"MultiZone,omitnil" name:"MultiZone"`

	// 节点资源的规格，有几个可用区，就填几个，按顺序第一个为主可用区，第二个为备可用区，第三个为仲裁可用区。如果没有开启跨AZ，则长度为1即可。
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil" name:"MultiZoneSettings"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Software")
	delete(f, "SupportHA")
	delete(f, "InstanceName")
	delete(f, "PayMode")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "LoginSettings")
	delete(f, "VPCSettings")
	delete(f, "ResourceSpec")
	delete(f, "COSSettings")
	delete(f, "Placement")
	delete(f, "SgId")
	delete(f, "PreExecutedFileSettings")
	delete(f, "AutoRenew")
	delete(f, "ClientToken")
	delete(f, "NeedMasterWan")
	delete(f, "RemoteLoginAtCreate")
	delete(f, "CheckSecurity")
	delete(f, "ExtendFsField")
	delete(f, "Tags")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "CbsEncrypt")
	delete(f, "MetaType")
	delete(f, "UnifyMetaInstanceId")
	delete(f, "MetaDBInfo")
	delete(f, "ApplicationRole")
	delete(f, "SceneName")
	delete(f, "ExternalService")
	delete(f, "VersionID")
	delete(f, "MultiZone")
	delete(f, "MultiZoneSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomMetaDBInfo struct {
	// 自定义MetaDB的JDBC连接，示例: jdbc:mysql://10.10.10.10:3306/dbname
	MetaDataJdbcUrl *string `json:"MetaDataJdbcUrl,omitnil" name:"MetaDataJdbcUrl"`

	// 自定义MetaDB用户名
	MetaDataUser *string `json:"MetaDataUser,omitnil" name:"MetaDataUser"`

	// 自定义MetaDB密码
	MetaDataPass *string `json:"MetaDataPass,omitnil" name:"MetaDataPass"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_DEFAULT_META：表示集群默认创建</li>
	// <li>EMR_EXIST_META：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`
}

type CustomMetaInfo struct {
	// 自定义MetaDB的JDBC连接，请以 jdbc:mysql:// 开头
	MetaDataJdbcUrl *string `json:"MetaDataJdbcUrl,omitnil" name:"MetaDataJdbcUrl"`

	// 自定义MetaDB用户名
	MetaDataUser *string `json:"MetaDataUser,omitnil" name:"MetaDataUser"`

	// 自定义MetaDB密码
	MetaDataPass *string `json:"MetaDataPass,omitnil" name:"MetaDataPass"`
}

type CustomServiceDefine struct {
	// 自定义参数key
	Name *string `json:"Name,omitnil" name:"Name"`

	// 自定义参数value
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type DeleteUserManagerUserListRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 集群用户名列表
	UserNameList []*string `json:"UserNameList,omitnil" name:"UserNameList"`

	// tke/eks集群id，容器集群传
	TkeClusterId *string `json:"TkeClusterId,omitnil" name:"TkeClusterId"`

	// 默认空，容器版传"native"
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// 用户组
	UserGroupList []*UserAndGroup `json:"UserGroupList,omitnil" name:"UserGroupList"`
}

type DeleteUserManagerUserListRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 集群用户名列表
	UserNameList []*string `json:"UserNameList,omitnil" name:"UserNameList"`

	// tke/eks集群id，容器集群传
	TkeClusterId *string `json:"TkeClusterId,omitnil" name:"TkeClusterId"`

	// 默认空，容器版传"native"
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// 用户组
	UserGroupList []*UserAndGroup `json:"UserGroupList,omitnil" name:"UserGroupList"`
}

func (r *DeleteUserManagerUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserManagerUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserNameList")
	delete(f, "TkeClusterId")
	delete(f, "DisplayStrategy")
	delete(f, "UserGroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserManagerUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserManagerUserListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUserManagerUserListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserManagerUserListResponseParams `json:"Response"`
}

func (r *DeleteUserManagerUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserManagerUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DependService struct {
	// 共用组件名
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 共用组件集群
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

// Predefined struct for user
type DescribeClusterNodesRequestParams struct {
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点标识，取值为：
	// <li>all：表示获取全部类型节点，cdb信息除外。</li>
	// <li>master：表示获取master节点信息。</li>
	// <li>core：表示获取core节点信息。</li>
	// <li>task：表示获取task节点信息。</li>
	// <li>common：表示获取common节点信息。</li>
	// <li>router：表示获取router节点信息。</li>
	// <li>db：表示获取正常状态的cdb信息。</li>
	// <li>recyle：表示获取回收站隔离中的节点信息，包括cdb信息。</li>
	// <li>renew：表示获取所有待续费的节点信息，包括cdb信息，自动续费节点不会返回。</li>
	// 注意：现在只支持以上取值，输入其他值会导致错误。
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// 页编号，默认值为0，表示第一页。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页返回数量，默认值为100，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 资源类型:支持all/host/pod，默认为all
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// 支持搜索的字段
	SearchFields []*SearchItem `json:"SearchFields,omitnil" name:"SearchFields"`

	// 无
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 无
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`
}

type DescribeClusterNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点标识，取值为：
	// <li>all：表示获取全部类型节点，cdb信息除外。</li>
	// <li>master：表示获取master节点信息。</li>
	// <li>core：表示获取core节点信息。</li>
	// <li>task：表示获取task节点信息。</li>
	// <li>common：表示获取common节点信息。</li>
	// <li>router：表示获取router节点信息。</li>
	// <li>db：表示获取正常状态的cdb信息。</li>
	// <li>recyle：表示获取回收站隔离中的节点信息，包括cdb信息。</li>
	// <li>renew：表示获取所有待续费的节点信息，包括cdb信息，自动续费节点不会返回。</li>
	// 注意：现在只支持以上取值，输入其他值会导致错误。
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// 页编号，默认值为0，表示第一页。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页返回数量，默认值为100，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 资源类型:支持all/host/pod，默认为all
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// 支持搜索的字段
	SearchFields []*SearchItem `json:"SearchFields,omitnil" name:"SearchFields"`

	// 无
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 无
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`
}

func (r *DescribeClusterNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeFlag")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "HardwareResourceType")
	delete(f, "SearchFields")
	delete(f, "OrderField")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodesResponseParams struct {
	// 查询到的节点总数
	TotalCnt *int64 `json:"TotalCnt,omitnil" name:"TotalCnt"`

	// 节点详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*NodeHardwareInfo `json:"NodeList,omitnil" name:"NodeList"`

	// 用户所有的标签键列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// 资源类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardwareResourceTypeList []*string `json:"HardwareResourceTypeList,omitnil" name:"HardwareResourceTypeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClusterNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterNodesResponseParams `json:"Response"`
}

func (r *DescribeClusterNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmQuotaRequestParams struct {
	// EMR集群ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 区ID
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`
}

type DescribeCvmQuotaRequest struct {
	*tchttp.BaseRequest
	
	// EMR集群ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 区ID
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`
}

func (r *DescribeCvmQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCvmQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmQuotaResponseParams struct {
	// 后付费配额列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostPaidQuotaSet []*QuotaEntity `json:"PostPaidQuotaSet,omitnil" name:"PostPaidQuotaSet"`

	// 竞价实例配额列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpotPaidQuotaSet []*QuotaEntity `json:"SpotPaidQuotaSet,omitnil" name:"SpotPaidQuotaSet"`

	// eks配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	EksQuotaSet []*PodSaleSpec `json:"EksQuotaSet,omitnil" name:"EksQuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCvmQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCvmQuotaResponseParams `json:"Response"`
}

func (r *DescribeCvmQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmrApplicationStaticsRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 起始时间，时间戳（秒）
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，时间戳（秒）
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 过滤的队列名
	Queues []*string `json:"Queues,omitnil" name:"Queues"`

	// 过滤的用户名
	Users []*string `json:"Users,omitnil" name:"Users"`

	// 过滤的作业类型
	ApplicationTypes []*string `json:"ApplicationTypes,omitnil" name:"ApplicationTypes"`

	// 分组字段，可选：queue, user, applicationType
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// 排序字段，可选：sumMemorySeconds, sumVCoreSeconds, sumHDFSBytesWritten, sumHDFSBytesRead
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 是否顺序排序，0-逆序，1-正序
	IsAsc *int64 `json:"IsAsc,omitnil" name:"IsAsc"`

	// 页号
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 页容量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeEmrApplicationStaticsRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 起始时间，时间戳（秒）
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，时间戳（秒）
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 过滤的队列名
	Queues []*string `json:"Queues,omitnil" name:"Queues"`

	// 过滤的用户名
	Users []*string `json:"Users,omitnil" name:"Users"`

	// 过滤的作业类型
	ApplicationTypes []*string `json:"ApplicationTypes,omitnil" name:"ApplicationTypes"`

	// 分组字段，可选：queue, user, applicationType
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// 排序字段，可选：sumMemorySeconds, sumVCoreSeconds, sumHDFSBytesWritten, sumHDFSBytesRead
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 是否顺序排序，0-逆序，1-正序
	IsAsc *int64 `json:"IsAsc,omitnil" name:"IsAsc"`

	// 页号
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 页容量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeEmrApplicationStaticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmrApplicationStaticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Queues")
	delete(f, "Users")
	delete(f, "ApplicationTypes")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "IsAsc")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEmrApplicationStaticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEmrApplicationStaticsResponseParams struct {
	// 作业统计信息
	Statics []*ApplicationStatics `json:"Statics,omitnil" name:"Statics"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 可选择的队列名
	Queues []*string `json:"Queues,omitnil" name:"Queues"`

	// 可选择的用户名
	Users []*string `json:"Users,omitnil" name:"Users"`

	// 可选择的作业类型
	ApplicationTypes []*string `json:"ApplicationTypes,omitnil" name:"ApplicationTypes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEmrApplicationStaticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEmrApplicationStaticsResponseParams `json:"Response"`
}

func (r *DescribeEmrApplicationStaticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEmrApplicationStaticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHiveQueriesRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 起始时间秒
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeHiveQueriesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 起始时间秒
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeHiveQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHiveQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHiveQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHiveQueriesResponseParams struct {
	// 总条数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 结果列表
	Results []*HiveQuery `json:"Results,omitnil" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHiveQueriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHiveQueriesResponseParams `json:"Response"`
}

func (r *DescribeHiveQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHiveQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImpalaQueriesRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 起始时间秒
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeImpalaQueriesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 起始时间秒
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeImpalaQueriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImpalaQueriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImpalaQueriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImpalaQueriesResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 结果列表
	Results []*ImpalaQuery `json:"Results,omitnil" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeImpalaQueriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImpalaQueriesResponseParams `json:"Response"`
}

func (r *DescribeImpalaQueriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImpalaQueriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRenewNodesRequestParams struct {
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceRenewNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceRenewNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRenewNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRenewNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRenewNodesResponseParams struct {
	// 查询到的节点总数
	TotalCnt *int64 `json:"TotalCnt,omitnil" name:"TotalCnt"`

	// 节点详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*RenewInstancesInfo `json:"NodeList,omitnil" name:"NodeList"`

	// 用户所有的标签键列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaInfo []*string `json:"MetaInfo,omitnil" name:"MetaInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceRenewNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceRenewNodesResponseParams `json:"Response"`
}

func (r *DescribeInstanceRenewNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRenewNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesListRequestParams struct {
	// 集群筛选策略。取值范围：<li>clusterList：表示查询除了已销毁集群之外的集群列表。</li><li>monitorManage：表示查询除了已销毁、创建中以及创建失败的集群之外的集群列表。</li><li>cloudHardwareManage/componentManage：目前这两个取值为预留取值，暂时和monitorManage表示同样的含义。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// 页编号，默认值为0，表示第一页。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页返回数量，默认值为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。取值范围：<li>clusterId：表示按照实例ID排序。</li><li>addTime：表示按照实例创建时间排序。</li><li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：<li>0：表示降序。</li><li>1：表示升序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`

	// 自定义查询
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

type DescribeInstancesListRequest struct {
	*tchttp.BaseRequest
	
	// 集群筛选策略。取值范围：<li>clusterList：表示查询除了已销毁集群之外的集群列表。</li><li>monitorManage：表示查询除了已销毁、创建中以及创建失败的集群之外的集群列表。</li><li>cloudHardwareManage/componentManage：目前这两个取值为预留取值，暂时和monitorManage表示同样的含义。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// 页编号，默认值为0，表示第一页。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页返回数量，默认值为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。取值范围：<li>clusterId：表示按照实例ID排序。</li><li>addTime：表示按照实例创建时间排序。</li><li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：<li>0：表示降序。</li><li>1：表示升序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`

	// 自定义查询
	Filters []*Filters `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeInstancesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStrategy")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Asc")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesListResponseParams struct {
	// 符合条件的实例总数。
	TotalCnt *int64 `json:"TotalCnt,omitnil" name:"TotalCnt"`

	// 集群实例列表
	InstancesList []*EmrListInstance `json:"InstancesList,omitnil" name:"InstancesList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesListResponseParams `json:"Response"`
}

func (r *DescribeInstancesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 集群筛选策略。取值范围：
	// <li>clusterList：表示查询除了已销毁集群之外的集群列表。</li>
	// <li>monitorManage：表示查询除了已销毁、创建中以及创建失败的集群之外的集群列表。</li>
	// <li>cloudHardwareManage/componentManage：目前这两个取值为预留取值，暂时和monitorManage表示同样的含义。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// 按照一个或者多个实例ID查询。实例ID形如: emr-xxxxxxxx 。(此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的 Ids.N 一节)。如果不填写实例ID，返回该APPID下所有实例列表。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 页编号，默认值为0，表示第一页。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页返回数量，默认值为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 建议必填-1，表示拉取所有项目下的集群。
	// 不填默认值为0，表示拉取默认项目下的集群。
	// 实例所属项目ID。该参数可以通过调用 [DescribeProjects](https://cloud.tencent.com/document/product/651/78725) 的返回值中的 projectId 字段来获取。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 排序字段。取值范围：
	// <li>clusterId：表示按照实例ID排序。</li>
	// <li>addTime：表示按照实例创建时间排序。</li>
	// <li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：
	// <li>0：表示降序。</li>
	// <li>1：表示升序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群筛选策略。取值范围：
	// <li>clusterList：表示查询除了已销毁集群之外的集群列表。</li>
	// <li>monitorManage：表示查询除了已销毁、创建中以及创建失败的集群之外的集群列表。</li>
	// <li>cloudHardwareManage/componentManage：目前这两个取值为预留取值，暂时和monitorManage表示同样的含义。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitnil" name:"DisplayStrategy"`

	// 按照一个或者多个实例ID查询。实例ID形如: emr-xxxxxxxx 。(此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的 Ids.N 一节)。如果不填写实例ID，返回该APPID下所有实例列表。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 页编号，默认值为0，表示第一页。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页返回数量，默认值为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 建议必填-1，表示拉取所有项目下的集群。
	// 不填默认值为0，表示拉取默认项目下的集群。
	// 实例所属项目ID。该参数可以通过调用 [DescribeProjects](https://cloud.tencent.com/document/product/651/78725) 的返回值中的 projectId 字段来获取。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 排序字段。取值范围：
	// <li>clusterId：表示按照实例ID排序。</li>
	// <li>addTime：表示按照实例创建时间排序。</li>
	// <li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：
	// <li>0：表示降序。</li>
	// <li>1：表示升序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitnil" name:"Asc"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStrategy")
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProjectId")
	delete(f, "OrderField")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 符合条件的实例总数。
	TotalCnt *int64 `json:"TotalCnt,omitnil" name:"TotalCnt"`

	// EMR实例详细信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterList []*ClusterInstancesInfo `json:"ClusterList,omitnil" name:"ClusterList"`

	// 实例关联的标签键列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobFlowRequestParams struct {
	// 流程任务Id，RunJobFlow接口返回的值。
	JobFlowId *int64 `json:"JobFlowId,omitnil" name:"JobFlowId"`
}

type DescribeJobFlowRequest struct {
	*tchttp.BaseRequest
	
	// 流程任务Id，RunJobFlow接口返回的值。
	JobFlowId *int64 `json:"JobFlowId,omitnil" name:"JobFlowId"`
}

func (r *DescribeJobFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobFlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobFlowResponseParams struct {
	// 流程任务状态，可以为以下值：
	// JobFlowInit，流程任务初始化。
	// JobFlowResourceApplied，资源申请中，通常为JobFlow需要新建集群时的状态。
	// JobFlowResourceReady，执行流程任务的资源就绪。
	// JobFlowStepsRunning，流程任务步骤已提交。
	// JobFlowStepsComplete，流程任务步骤已完成。
	// JobFlowTerminating，流程任务所需资源销毁中。
	// JobFlowFinish，流程任务已完成。
	State *string `json:"State,omitnil" name:"State"`

	// 流程任务步骤结果。
	Details []*JobResult `json:"Details,omitnil" name:"Details"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeJobFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobFlowResponseParams `json:"Response"`
}

func (r *DescribeJobFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceScheduleRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeResourceScheduleRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeResourceScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceScheduleResponseParams struct {
	// 资源调度功能是否开启
	OpenSwitch *bool `json:"OpenSwitch,omitnil" name:"OpenSwitch"`

	// 正在使用的资源调度器
	Scheduler *string `json:"Scheduler,omitnil" name:"Scheduler"`

	// 公平调度器的信息
	FSInfo *string `json:"FSInfo,omitnil" name:"FSInfo"`

	// 容量调度器的信息
	CSInfo *string `json:"CSInfo,omitnil" name:"CSInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeResourceScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceScheduleResponseParams `json:"Response"`
}

func (r *DescribeResourceScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersForUserManagerRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 分页的大小
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 查询用户列表过滤器
	UserManagerFilter *UserManagerFilter `json:"UserManagerFilter,omitnil" name:"UserManagerFilter"`

	// 是否需要keytab文件的信息，仅对开启kerberos的集群有效，默认为false
	NeedKeytabInfo *bool `json:"NeedKeytabInfo,omitnil" name:"NeedKeytabInfo"`
}

type DescribeUsersForUserManagerRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 分页的大小
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 查询用户列表过滤器
	UserManagerFilter *UserManagerFilter `json:"UserManagerFilter,omitnil" name:"UserManagerFilter"`

	// 是否需要keytab文件的信息，仅对开启kerberos的集群有效，默认为false
	NeedKeytabInfo *bool `json:"NeedKeytabInfo,omitnil" name:"NeedKeytabInfo"`
}

func (r *DescribeUsersForUserManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersForUserManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "UserManagerFilter")
	delete(f, "NeedKeytabInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersForUserManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersForUserManagerResponseParams struct {
	// 总数
	TotalCnt *int64 `json:"TotalCnt,omitnil" name:"TotalCnt"`

	// 用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserManagerUserList []*UserManagerUserBriefInfo `json:"UserManagerUserList,omitnil" name:"UserManagerUserList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUsersForUserManagerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersForUserManagerResponseParams `json:"Response"`
}

func (r *DescribeUsersForUserManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersForUserManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYarnApplicationsRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 起始时间秒
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeYarnApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 起始时间秒
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间秒，EndTime-StartTime不得超过1天秒数86400
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页起始偏移，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页大小，合法范围[1,100]
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeYarnApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYarnApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeYarnApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeYarnApplicationsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 结果列表
	Results []*YarnApplication `json:"Results,omitnil" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeYarnApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeYarnApplicationsResponseParams `json:"Response"`
}

func (r *DescribeYarnApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeYarnApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskGroup struct {
	// 磁盘规格。
	Spec *DiskSpec `json:"Spec,omitnil" name:"Spec"`

	// 同类型磁盘数量。
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type DiskSpec struct {
	// 磁盘类型。
	// LOCAL_BASIC  本地盘。
	// CLOUD_BASIC 云硬盘。
	// LOCAL_SSD 本地SSD。
	// CLOUD_SSD 云SSD。
	// CLOUD_PREMIUM 高效云盘。
	// CLOUD_HSSD 增强型云SSD。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 磁盘大小，单位GB。
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`
}

type DiskSpecInfo struct {
	// 磁盘数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 系统盘类型 取值范围：
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_BASIC：表示云硬盘。</li>
	// <li>LOCAL_BASIC：表示本地盘。</li>
	// <li>LOCAL_SSD：表示本地SSD。</li>
	// 
	// 数据盘类型 取值范围：
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_BASIC：表示云硬盘。</li>
	// <li>LOCAL_BASIC：表示本地盘。</li>
	// <li>LOCAL_SSD：表示本地SSD。</li>
	// <li>CLOUD_HSSD：表示增强型SSD云硬盘。</li>
	// <li>CLOUD_THROUGHPUT：表示吞吐型云硬盘。</li>
	// <li>CLOUD_TSSD：表示极速型SSD云硬盘。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 数据容量，单位为GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`
}

type DynamicPodSpec struct {
	// 需求最小cpu核数
	RequestCpu *float64 `json:"RequestCpu,omitnil" name:"RequestCpu"`

	// 需求最大cpu核数
	LimitCpu *float64 `json:"LimitCpu,omitnil" name:"LimitCpu"`

	// 需求最小memory，单位MB
	RequestMemory *float64 `json:"RequestMemory,omitnil" name:"RequestMemory"`

	// 需求最大memory，单位MB
	LimitMemory *float64 `json:"LimitMemory,omitnil" name:"LimitMemory"`
}

type EmrListInstance struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// 集群名字
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 集群地域
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 用户APPID
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`

	// 运行时间
	RunTime *string `json:"RunTime,omitnil" name:"RunTime"`

	// 集群IP
	MasterIp *string `json:"MasterIp,omitnil" name:"MasterIp"`

	// 集群版本
	EmrVersion *string `json:"EmrVersion,omitnil" name:"EmrVersion"`

	// 集群计费类型
	ChargeType *uint64 `json:"ChargeType,omitnil" name:"ChargeType"`

	// emr ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *uint64 `json:"RegionId,omitnil" name:"RegionId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *uint64 `json:"SubnetId,omitnil" name:"SubnetId"`

	// 网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *uint64 `json:"VpcId,omitnil" name:"VpcId"`

	// 地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 实例标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 告警信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmInfo *string `json:"AlarmInfo,omitnil" name:"AlarmInfo"`

	// 是否是woodpecker集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWoodpeckerCluster *uint64 `json:"IsWoodpeckerCluster,omitnil" name:"IsWoodpeckerCluster"`

	// Vpc中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 子网中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 字符串VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 字符串子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterClass *string `json:"ClusterClass,omitnil" name:"ClusterClass"`

	// 是否为跨AZ集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMultiZoneCluster *bool `json:"IsMultiZoneCluster,omitnil" name:"IsMultiZoneCluster"`

	// 是否手戳集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsHandsCluster *bool `json:"IsHandsCluster,omitnil" name:"IsHandsCluster"`

	// 体外客户端组件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutSideSoftInfo []*SoftDependInfo `json:"OutSideSoftInfo,omitnil" name:"OutSideSoftInfo"`

	// 当前集群的应用场景是否支持体外客户端
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportOutsideCluster *bool `json:"IsSupportOutsideCluster,omitnil" name:"IsSupportOutsideCluster"`
}

type EmrPrice struct {
	// 刊例价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// 折扣价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountCost *string `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 询价配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceSpec *PriceResource `json:"PriceSpec,omitnil" name:"PriceSpec"`

	// 是否支持竞价实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportSpotPaid *bool `json:"SupportSpotPaid,omitnil" name:"SupportSpotPaid"`
}

type EmrProductConfigOutter struct {
	// 软件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftInfo []*string `json:"SoftInfo,omitnil" name:"SoftInfo"`

	// Master节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterNodeSize *int64 `json:"MasterNodeSize,omitnil" name:"MasterNodeSize"`

	// Core节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreNodeSize *int64 `json:"CoreNodeSize,omitnil" name:"CoreNodeSize"`

	// Task节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNodeSize *int64 `json:"TaskNodeSize,omitnil" name:"TaskNodeSize"`

	// Common节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComNodeSize *int64 `json:"ComNodeSize,omitnil" name:"ComNodeSize"`

	// Master节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterResource *OutterResource `json:"MasterResource,omitnil" name:"MasterResource"`

	// Core节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreResource *OutterResource `json:"CoreResource,omitnil" name:"CoreResource"`

	// Task节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResource *OutterResource `json:"TaskResource,omitnil" name:"TaskResource"`

	// Common节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComResource *OutterResource `json:"ComResource,omitnil" name:"ComResource"`

	// 是否使用COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnCos *bool `json:"OnCos,omitnil" name:"OnCos"`

	// 收费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitnil" name:"ChargeType"`

	// Router节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouterNodeSize *int64 `json:"RouterNodeSize,omitnil" name:"RouterNodeSize"`

	// 是否支持HA
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportHA *bool `json:"SupportHA,omitnil" name:"SupportHA"`

	// 是否支持安全模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityOn *bool `json:"SecurityOn,omitnil" name:"SecurityOn"`

	// 安全组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroup *string `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// 是否开启Cbs加密
	// 注意：此字段可能返回 null，表示取不到有效值。
	CbsEncrypt *int64 `json:"CbsEncrypt,omitnil" name:"CbsEncrypt"`

	// 自定义应用角色。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationRole *string `json:"ApplicationRole,omitnil" name:"ApplicationRole"`

	// 安全组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroups []*string `json:"SecurityGroups,omitnil" name:"SecurityGroups"`

	// SSH密钥Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicKeyId *string `json:"PublicKeyId,omitnil" name:"PublicKeyId"`
}

type Execution struct {
	// 任务类型，目前支持以下类型。
	// 1. “MR”，将通过hadoop jar的方式提交。
	// 2. "HIVE"，将通过hive -f的方式提交。
	// 3. "SPARK"，将通过spark-submit的方式提交。
	JobType *string `json:"JobType,omitnil" name:"JobType"`

	// 任务参数，提供除提交指令以外的参数。
	Args []*string `json:"Args,omitnil" name:"Args"`
}

type ExternalService struct {
	// 共用组件类型，EMR/CUSTOM
	ShareType *string `json:"ShareType,omitnil" name:"ShareType"`

	// 自定义参数集合
	CustomServiceDefineList []*CustomServiceDefine `json:"CustomServiceDefineList,omitnil" name:"CustomServiceDefineList"`

	// 共用组件名
	Service *string `json:"Service,omitnil" name:"Service"`

	// 共用组件集群
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type Filters struct {
	// 字段名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤字段值
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type HiveQuery struct {
	// 查询语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Statement *string `json:"Statement,omitnil" name:"Statement"`

	// 执行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitnil" name:"Duration"`

	// 开始时间毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil" name:"State"`

	// 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil" name:"User"`

	// appId列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobIds []*string `json:"JobIds,omitnil" name:"JobIds"`

	// 执行引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEngine *string `json:"ExecutionEngine,omitnil" name:"ExecutionEngine"`

	// 查询ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`
}

type HostVolumeContext struct {
	// Pod挂载宿主机的目录。资源对宿主机的挂载点，指定的挂载点对应了宿主机的路径，该挂载点在Pod中作为数据存储目录使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumePath *string `json:"VolumePath,omitnil" name:"VolumePath"`
}

type ImpalaQuery struct {
	// 执行语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Statement *string `json:"Statement,omitnil" name:"Statement"`

	// 查询ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitnil" name:"Duration"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 执行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil" name:"State"`

	// 获取行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowsFetched *int64 `json:"RowsFetched,omitnil" name:"RowsFetched"`

	// 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil" name:"User"`

	// 默认DB
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDB *string `json:"DefaultDB,omitnil" name:"DefaultDB"`

	// 执行的Coordinator节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coordinator *string `json:"Coordinator,omitnil" name:"Coordinator"`

	// 单节点内存峰值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxNodePeakMemoryUsage *string `json:"MaxNodePeakMemoryUsage,omitnil" name:"MaxNodePeakMemoryUsage"`

	// 查询类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 扫描的HDFS行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanHDFSRows *int64 `json:"ScanHDFSRows,omitnil" name:"ScanHDFSRows"`

	// 扫描的Kudu行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanKUDURows *int64 `json:"ScanKUDURows,omitnil" name:"ScanKUDURows"`

	// 扫描的总行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanRowsTotal *int64 `json:"ScanRowsTotal,omitnil" name:"ScanRowsTotal"`

	// 读取的总字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalBytesRead *int64 `json:"TotalBytesRead,omitnil" name:"TotalBytesRead"`

	// 发送的总字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalBytesSent *int64 `json:"TotalBytesSent,omitnil" name:"TotalBytesSent"`

	// CPU总时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCpuTime *int64 `json:"TotalCpuTime,omitnil" name:"TotalCpuTime"`

	// 内部数据发送总量(Bytes)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalInnerBytesSent *int64 `json:"TotalInnerBytesSent,omitnil" name:"TotalInnerBytesSent"`

	// 内部扫描数据发送总量(Bytes)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalScanBytesSent *int64 `json:"TotalScanBytesSent,omitnil" name:"TotalScanBytesSent"`

	// 预估单节点内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	EstimatedPerHostMemBytes *int64 `json:"EstimatedPerHostMemBytes,omitnil" name:"EstimatedPerHostMemBytes"`

	// 从缓存中获取的数据行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumRowsFetchedFromCache *int64 `json:"NumRowsFetchedFromCache,omitnil" name:"NumRowsFetchedFromCache"`
}

// Predefined struct for user
type InquirePriceRenewEmrRequestParams struct {
	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费一个月
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 待续费集群ID列表。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`
}

type InquirePriceRenewEmrRequest struct {
	*tchttp.BaseRequest
	
	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费一个月
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 待续费集群ID列表。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`
}

func (r *InquirePriceRenewEmrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewEmrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeSpan")
	delete(f, "InstanceId")
	delete(f, "Placement")
	delete(f, "PayMode")
	delete(f, "TimeUnit")
	delete(f, "Currency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewEmrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewEmrResponseParams struct {
	// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCost *float64 `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountCost *float64 `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 实例续费的时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceRenewEmrResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewEmrResponseParams `json:"Response"`
}

func (r *InquirePriceRenewEmrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewEmrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateInstanceRequestParams struct {
	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitnil" name:"SupportHA"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）需要选择不同的必选组件：
	// <li>ProductId为1的时候，必选组件包括：hadoop-2.7.3、knox-1.2.0、zookeeper-3.4.9</li>
	// <li>ProductId为2的时候，必选组件包括：hadoop-2.7.3、knox-1.2.0、zookeeper-3.4.9</li>
	// <li>ProductId为4的时候，必选组件包括：hadoop-2.8.4、knox-1.2.0、zookeeper-3.4.9</li>
	// <li>ProductId为7的时候，必选组件包括：hadoop-3.1.2、knox-1.2.0、zookeeper-3.4.9</li>
	Software []*string `json:"Software,omitnil" name:"Software"`

	// 询价的节点规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_NEW_META：表示集群默认创建</li>
	// <li>EMR_EXIT_METE：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：
	// <li>1：表示EMR-V1.3.1。</li>
	// <li>2：表示EMR-V2.0.1。</li>
	// <li>4：表示EMR-V2.1.0。</li>
	// <li>7：表示EMR-V3.0.0。</li>
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// 场景化取值：
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// 共用组件信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil" name:"ExternalService"`

	// 当前默认值为0，跨AZ特性支持后为1
	VersionID *uint64 `json:"VersionID,omitnil" name:"VersionID"`

	// 可用区的规格信息
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil" name:"MultiZoneSettings"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitnil" name:"SupportHA"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）需要选择不同的必选组件：
	// <li>ProductId为1的时候，必选组件包括：hadoop-2.7.3、knox-1.2.0、zookeeper-3.4.9</li>
	// <li>ProductId为2的时候，必选组件包括：hadoop-2.7.3、knox-1.2.0、zookeeper-3.4.9</li>
	// <li>ProductId为4的时候，必选组件包括：hadoop-2.8.4、knox-1.2.0、zookeeper-3.4.9</li>
	// <li>ProductId为7的时候，必选组件包括：hadoop-3.1.2、knox-1.2.0、zookeeper-3.4.9</li>
	Software []*string `json:"Software,omitnil" name:"Software"`

	// 询价的节点规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_NEW_META：表示集群默认创建</li>
	// <li>EMR_EXIT_METE：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`

	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：
	// <li>1：表示EMR-V1.3.1。</li>
	// <li>2：表示EMR-V2.0.1。</li>
	// <li>4：表示EMR-V2.1.0。</li>
	// <li>7：表示EMR-V3.0.0。</li>
	ProductId *uint64 `json:"ProductId,omitnil" name:"ProductId"`

	// 场景化取值：
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`

	// 共用组件信息
	ExternalService []*ExternalService `json:"ExternalService,omitnil" name:"ExternalService"`

	// 当前默认值为0，跨AZ特性支持后为1
	VersionID *uint64 `json:"VersionID,omitnil" name:"VersionID"`

	// 可用区的规格信息
	MultiZoneSettings []*MultiZoneSetting `json:"MultiZoneSettings,omitnil" name:"MultiZoneSettings"`
}

func (r *InquiryPriceCreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "Currency")
	delete(f, "PayMode")
	delete(f, "SupportHA")
	delete(f, "Software")
	delete(f, "ResourceSpec")
	delete(f, "Placement")
	delete(f, "VPCSettings")
	delete(f, "MetaType")
	delete(f, "UnifyMetaInstanceId")
	delete(f, "MetaDBInfo")
	delete(f, "ProductId")
	delete(f, "SceneName")
	delete(f, "ExternalService")
	delete(f, "VersionID")
	delete(f, "MultiZoneSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateInstanceResponseParams struct {
	// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCost *float64 `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountCost *float64 `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。</li>
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 购买实例的时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 价格清单
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceList []*ZoneDetailPriceResult `json:"PriceList,omitnil" name:"PriceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstanceRequestParams struct {
	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费一个月
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 待续费节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr)查询。
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 是否按量转包年包月。0：否，1：是。
	ModifyPayMode *int64 `json:"ModifyPayMode,omitnil" name:"ModifyPayMode"`
}

type InquiryPriceRenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费一个月
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 待续费节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr)查询。
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 是否按量转包年包月。0：否，1：是。
	ModifyPayMode *int64 `json:"ModifyPayMode,omitnil" name:"ModifyPayMode"`
}

func (r *InquiryPriceRenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeSpan")
	delete(f, "ResourceIds")
	delete(f, "Placement")
	delete(f, "PayMode")
	delete(f, "TimeUnit")
	delete(f, "Currency")
	delete(f, "ModifyPayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstanceResponseParams struct {
	// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCost *float64 `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountCost *float64 `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 实例续费的时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceRenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceScaleOutInstanceRequestParams struct {
	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 实例所属的可用区ID，例如100003。该参数可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/213/15707) 的返回值中的ZoneId字段来获取。
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitnil" name:"RouterCount"`

	// 扩容的Master节点数量。
	MasterCount *uint64 `json:"MasterCount,omitnil" name:"MasterCount"`
}

type InquiryPriceScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 实例所属的可用区ID，例如100003。该参数可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/213/15707) 的返回值中的ZoneId字段来获取。
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitnil" name:"RouterCount"`

	// 扩容的Master节点数量。
	MasterCount *uint64 `json:"MasterCount,omitnil" name:"MasterCount"`
}

func (r *InquiryPriceScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "ZoneId")
	delete(f, "PayMode")
	delete(f, "InstanceId")
	delete(f, "CoreCount")
	delete(f, "TaskCount")
	delete(f, "Currency")
	delete(f, "RouterCount")
	delete(f, "MasterCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceScaleOutInstanceResponseParams struct {
	// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountCost *string `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。</li>
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 询价的节点规格。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceSpec *PriceResource `json:"PriceSpec,omitnil" name:"PriceSpec"`

	// 对应入参MultipleResources中多个规格的询价结果，其它出参返回的是第一个规格的询价结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultipleEmrPrice []*EmrPrice `json:"MultipleEmrPrice,omitnil" name:"MultipleEmrPrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceScaleOutInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpdateInstanceRequestParams struct {
	// 变配的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 变配的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 节点变配的目标配置。
	UpdateSpec *UpdateInstanceSettings `json:"UpdateSpec,omitnil" name:"UpdateSpec"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 批量变配资源ID列表
	ResourceIdList []*string `json:"ResourceIdList,omitnil" name:"ResourceIdList"`
}

type InquiryPriceUpdateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 变配的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 变配的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 节点变配的目标配置。
	UpdateSpec *UpdateInstanceSettings `json:"UpdateSpec,omitnil" name:"UpdateSpec"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 批量变配资源ID列表
	ResourceIdList []*string `json:"ResourceIdList,omitnil" name:"ResourceIdList"`
}

func (r *InquiryPriceUpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpdateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "UpdateSpec")
	delete(f, "PayMode")
	delete(f, "Placement")
	delete(f, "Currency")
	delete(f, "ResourceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpdateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpdateInstanceResponseParams struct {
	// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCost *float64 `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountCost *float64 `json:"DiscountCost,omitnil" name:"DiscountCost"`

	// 变配的时间单位。取值范围：
	// <li>s：表示秒。</li>
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 变配的时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 价格详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceDetail []*PriceDetail `json:"PriceDetail,omitnil" name:"PriceDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceUpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpdateInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceUpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpdateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// 包年包月时间，默认为1，单位：月。
	// 取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11, 12, 24, 36, 48, 60。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 是否自动续费，默认为否。
	// <li>true：是</li>
	// <li>false：否</li>
	RenewFlag *bool `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type JobFlowResource struct {
	// 机器类型描述。
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// 机器类型描述，可参考CVM的该含义。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 标签KV对。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 磁盘描述列表。
	DiskGroups []*DiskGroup `json:"DiskGroups,omitnil" name:"DiskGroups"`
}

type JobFlowResourceSpec struct {
	// 主节点数量。
	MasterCount *int64 `json:"MasterCount,omitnil" name:"MasterCount"`

	// 主节点配置。
	MasterResourceSpec *JobFlowResource `json:"MasterResourceSpec,omitnil" name:"MasterResourceSpec"`

	// Core节点数量
	CoreCount *int64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// Core节点配置。
	CoreResourceSpec *JobFlowResource `json:"CoreResourceSpec,omitnil" name:"CoreResourceSpec"`

	// Task节点数量。
	TaskCount *int64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// Common节点数量。
	CommonCount *int64 `json:"CommonCount,omitnil" name:"CommonCount"`

	// Task节点配置。
	TaskResourceSpec *JobFlowResource `json:"TaskResourceSpec,omitnil" name:"TaskResourceSpec"`

	// Common节点配置。
	CommonResourceSpec *JobFlowResource `json:"CommonResourceSpec,omitnil" name:"CommonResourceSpec"`
}

type JobResult struct {
	// 任务步骤名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 任务步骤失败时的处理策略，可以为以下值：
	// "CONTINUE"，跳过当前失败步骤，继续后续步骤。
	// “TERMINATE_CLUSTER”，终止当前及后续步骤，并销毁集群。
	// “CANCEL_AND_WAIT”，取消当前步骤并阻塞等待处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionOnFailure *string `json:"ActionOnFailure,omitnil" name:"ActionOnFailure"`

	// 当前步骤的状态，可以为以下值：
	// “JobFlowStepStatusInit”，初始化状态，等待执行。
	// “JobFlowStepStatusRunning”，任务步骤正在执行。
	// “JobFlowStepStatusFailed”，任务步骤执行失败。
	// “JobFlowStepStatusSucceed”，任务步骤执行成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobState *string `json:"JobState,omitnil" name:"JobState"`

	// YARN任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`
}

type LoginSettings struct {
	// 实例登录密码，8-16个字符，包含大写字母、小写字母、数字和特殊字符四种，特殊符号仅支持!@%^*，密码第一位不能为特殊字符
	Password *string `json:"Password,omitnil" name:"Password"`

	// 密钥ID。关联密钥后，就可以通过对应的私钥来访问实例；PublicKeyId可通过接口[DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699)获取
	PublicKeyId *string `json:"PublicKeyId,omitnil" name:"PublicKeyId"`
}

type MetaDbInfo struct {
	// 元数据类型。
	MetaType *string `json:"MetaType,omitnil" name:"MetaType"`

	// 统一元数据库实例ID。
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitnil" name:"UnifyMetaInstanceId"`

	// 自建元数据库信息。
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitnil" name:"MetaDBInfo"`
}

// Predefined struct for user
type ModifyResourcePoolsRequestParams struct {
	// emr集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 取值范围：
	// <li>fair:代表公平调度标识</li>
	// <li>capacity:代表容量调度标识</li>
	Key *string `json:"Key,omitnil" name:"Key"`
}

type ModifyResourcePoolsRequest struct {
	*tchttp.BaseRequest
	
	// emr集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 取值范围：
	// <li>fair:代表公平调度标识</li>
	// <li>capacity:代表容量调度标识</li>
	Key *string `json:"Key,omitnil" name:"Key"`
}

func (r *ModifyResourcePoolsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePoolsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Key")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePoolsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePoolsResponseParams struct {
	// false表示不是草稿，提交刷新请求成功
	IsDraft *bool `json:"IsDraft,omitnil" name:"IsDraft"`

	// 扩展字段，暂时没用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyResourcePoolsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePoolsResponseParams `json:"Response"`
}

func (r *ModifyResourcePoolsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceScheduleConfigRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 业务标识，fair表示编辑公平的配置项，fairPlan表示编辑执行计划，capacity表示编辑容量的配置项
	Key *string `json:"Key,omitnil" name:"Key"`

	// 修改后的模块消息
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ModifyResourceScheduleConfigRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 业务标识，fair表示编辑公平的配置项，fairPlan表示编辑执行计划，capacity表示编辑容量的配置项
	Key *string `json:"Key,omitnil" name:"Key"`

	// 修改后的模块消息
	Value *string `json:"Value,omitnil" name:"Value"`
}

func (r *ModifyResourceScheduleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceScheduleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Key")
	delete(f, "Value")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceScheduleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceScheduleConfigResponseParams struct {
	// true为草稿，表示还没有刷新资源池
	IsDraft *bool `json:"IsDraft,omitnil" name:"IsDraft"`

	// 校验错误信息，如果不为空，则说明校验失败，配置没有成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 返回数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyResourceScheduleConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceScheduleConfigResponseParams `json:"Response"`
}

func (r *ModifyResourceScheduleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceScheduleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceSchedulerRequestParams struct {
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 老的调度器:fair
	OldValue *string `json:"OldValue,omitnil" name:"OldValue"`

	// 新的调度器:capacity
	NewValue *string `json:"NewValue,omitnil" name:"NewValue"`
}

type ModifyResourceSchedulerRequest struct {
	*tchttp.BaseRequest
	
	// emr集群的英文id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 老的调度器:fair
	OldValue *string `json:"OldValue,omitnil" name:"OldValue"`

	// 新的调度器:capacity
	NewValue *string `json:"NewValue,omitnil" name:"NewValue"`
}

func (r *ModifyResourceSchedulerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceSchedulerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldValue")
	delete(f, "NewValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceSchedulerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceSchedulerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyResourceSchedulerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceSchedulerResponseParams `json:"Response"`
}

func (r *ModifyResourceSchedulerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceSchedulerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceTags struct {
	// 集群id 或者 cvm id
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源6段式表达式
	Resource *string `json:"Resource,omitnil" name:"Resource"`

	// 资源前缀
	ResourcePrefix *string `json:"ResourcePrefix,omitnil" name:"ResourcePrefix"`

	// ap-beijing
	ResourceRegion *string `json:"ResourceRegion,omitnil" name:"ResourceRegion"`

	// emr
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 删除的标签列表
	DeleteTags []*Tag `json:"DeleteTags,omitnil" name:"DeleteTags"`

	// 添加的标签列表
	AddTags []*Tag `json:"AddTags,omitnil" name:"AddTags"`

	// 修改的标签列表
	ModifyTags []*Tag `json:"ModifyTags,omitnil" name:"ModifyTags"`
}

// Predefined struct for user
type ModifyResourcesTagsRequestParams struct {
	// 标签类型，取值Cluster或者Node
	ModifyType *string `json:"ModifyType,omitnil" name:"ModifyType"`

	// 标签信息
	ModifyResourceTagsInfoList []*ModifyResourceTags `json:"ModifyResourceTagsInfoList,omitnil" name:"ModifyResourceTagsInfoList"`
}

type ModifyResourcesTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签类型，取值Cluster或者Node
	ModifyType *string `json:"ModifyType,omitnil" name:"ModifyType"`

	// 标签信息
	ModifyResourceTagsInfoList []*ModifyResourceTags `json:"ModifyResourceTagsInfoList,omitnil" name:"ModifyResourceTagsInfoList"`
}

func (r *ModifyResourcesTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcesTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModifyType")
	delete(f, "ModifyResourceTagsInfoList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcesTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcesTagsResponseParams struct {
	// 成功的资源id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessList []*string `json:"SuccessList,omitnil" name:"SuccessList"`

	// 失败的资源id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailList []*string `json:"FailList,omitnil" name:"FailList"`

	// 部分成功的资源id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartSuccessList []*string `json:"PartSuccessList,omitnil" name:"PartSuccessList"`

	// 集群id与流程id的映射列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterToFlowIdList []*ClusterIDToFlowID `json:"ClusterToFlowIdList,omitnil" name:"ClusterToFlowIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyResourcesTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcesTagsResponseParams `json:"Response"`
}

func (r *ModifyResourcesTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcesTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserManagerPwdRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 密码
	PassWord *string `json:"PassWord,omitnil" name:"PassWord"`
}

type ModifyUserManagerPwdRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 密码
	PassWord *string `json:"PassWord,omitnil" name:"PassWord"`
}

func (r *ModifyUserManagerPwdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserManagerPwdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "PassWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserManagerPwdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserManagerPwdResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyUserManagerPwdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserManagerPwdResponseParams `json:"Response"`
}

func (r *ModifyUserManagerPwdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserManagerPwdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultiDisk struct {
	// 云盘类型
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_HSSD：表示增强型SSD云硬盘。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 云盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// 该类型云盘个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type MultiDiskMC struct {
	// 该类型云盘个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 云盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`
}

type MultiZoneSetting struct {
	// "master"、"standby"、"third-party"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneTag *string `json:"ZoneTag,omitnil" name:"ZoneTag"`

	// 无
	VPCSettings *VPCSettings `json:"VPCSettings,omitnil" name:"VPCSettings"`

	// 无
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 无
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`
}

type NewResourceSpec struct {
	// 描述Master节点资源
	MasterResourceSpec *Resource `json:"MasterResourceSpec,omitnil" name:"MasterResourceSpec"`

	// 描述Core节点资源
	CoreResourceSpec *Resource `json:"CoreResourceSpec,omitnil" name:"CoreResourceSpec"`

	// 描述Task节点资源
	TaskResourceSpec *Resource `json:"TaskResourceSpec,omitnil" name:"TaskResourceSpec"`

	// Master节点数量
	MasterCount *int64 `json:"MasterCount,omitnil" name:"MasterCount"`

	// Core节点数量
	CoreCount *int64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// Task节点数量
	TaskCount *int64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// 描述Common节点资源
	CommonResourceSpec *Resource `json:"CommonResourceSpec,omitnil" name:"CommonResourceSpec"`

	// Common节点数量
	CommonCount *int64 `json:"CommonCount,omitnil" name:"CommonCount"`
}

type NodeDetailPriceResult struct {
	// 节点类型 master core task common router mysql
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// 节点组成部分价格详情
	PartDetailPrice []*PartDetailPriceItem `json:"PartDetailPrice,omitnil" name:"PartDetailPrice"`
}

type NodeHardwareInfo struct {
	// 用户APPID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialNo *string `json:"SerialNo,omitnil" name:"SerialNo"`

	// 机器实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *string `json:"OrderNo,omitnil" name:"OrderNo"`

	// master节点绑定外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitnil" name:"WanIp"`

	// 节点类型。0:common节点；1:master节点
	// ；2:core节点；3:task节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Flag *int64 `json:"Flag,omitnil" name:"Flag"`

	// 节点规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// 节点核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuNum *int64 `json:"CpuNum,omitnil" name:"CpuNum"`

	// 节点内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// 节点内存描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemDesc *string `json:"MemDesc,omitnil" name:"MemDesc"`

	// 节点所在region
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// 节点所在Zone
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyTime *string `json:"ApplyTime,omitnil" name:"ApplyTime"`

	// 释放时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeTime *string `json:"FreeTime,omitnil" name:"FreeTime"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *string `json:"DiskSize,omitnil" name:"DiskSize"`

	// 节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameTag *string `json:"NameTag,omitnil" name:"NameTag"`

	// 节点部署服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Services *string `json:"Services,omitnil" name:"Services"`

	// 磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitnil" name:"StorageType"`

	// 系统盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitnil" name:"RootSize"`

	// 付费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitnil" name:"ChargeType"`

	// 数据库IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdbIp *string `json:"CdbIp,omitnil" name:"CdbIp"`

	// 数据库端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdbPort *int64 `json:"CdbPort,omitnil" name:"CdbPort"`

	// 硬盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwDiskSize *int64 `json:"HwDiskSize,omitnil" name:"HwDiskSize"`

	// 硬盘容量描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwDiskSizeDesc *string `json:"HwDiskSizeDesc,omitnil" name:"HwDiskSizeDesc"`

	// 内存容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwMemSize *int64 `json:"HwMemSize,omitnil" name:"HwMemSize"`

	// 内存容量描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwMemSizeDesc *string `json:"HwMemSizeDesc,omitnil" name:"HwMemSizeDesc"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 节点资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmrResourceId *string `json:"EmrResourceId,omitnil" name:"EmrResourceId"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAutoRenew *int64 `json:"IsAutoRenew,omitnil" name:"IsAutoRenew"`

	// 设备标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceClass *string `json:"DeviceClass,omitnil" name:"DeviceClass"`

	// 支持变配
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mutable *int64 `json:"Mutable,omitnil" name:"Mutable"`

	// 多云盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	MCMultiDisk []*MultiDiskMC `json:"MCMultiDisk,omitnil" name:"MCMultiDisk"`

	// 数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdbNodeInfo *CdbInfo `json:"CdbNodeInfo,omitnil" name:"CdbNodeInfo"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 此节点是否可销毁，1可销毁，0不可销毁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destroyable *int64 `json:"Destroyable,omitnil" name:"Destroyable"`

	// 节点绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 是否是自动扩缩容节点，0为普通节点，1为自动扩缩容节点。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoFlag *int64 `json:"AutoFlag,omitnil" name:"AutoFlag"`

	// 资源类型, host/pod
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// 是否浮动规格，1是，0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDynamicSpec *int64 `json:"IsDynamicSpec,omitnil" name:"IsDynamicSpec"`

	// 浮动规格值json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	DynamicPodSpec *string `json:"DynamicPodSpec,omitnil" name:"DynamicPodSpec"`

	// 是否支持变更计费类型 1是，0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportModifyPayMode *int64 `json:"SupportModifyPayMode,omitnil" name:"SupportModifyPayMode"`

	// 系统盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootStorageType *int64 `json:"RootStorageType,omitnil" name:"RootStorageType"`

	// 可用区信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetInfo *SubnetInfo `json:"SubnetInfo,omitnil" name:"SubnetInfo"`

	// 客户端
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clients *string `json:"Clients,omitnil" name:"Clients"`

	// 系统当前时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentTime *string `json:"CurrentTime,omitnil" name:"CurrentTime"`

	// 是否用于联邦 ,1是，0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFederation *int64 `json:"IsFederation,omitnil" name:"IsFederation"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceClient *string `json:"ServiceClient,omitnil" name:"ServiceClient"`

	// 该实例是否开启实例保护，true为开启 false为关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisableApiTermination *bool `json:"DisableApiTermination,omitnil" name:"DisableApiTermination"`

	// 0表示老计费，1表示新计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeVersion *int64 `json:"TradeVersion,omitnil" name:"TradeVersion"`

	// 各组件状态，Zookeeper:STARTED,ResourceManager:STARTED，STARTED已启动，STOPED已停止
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServicesStatus *string `json:"ServicesStatus,omitnil" name:"ServicesStatus"`
}

type NodeResourceSpec struct {
	// 规格类型，如S2.MEDIUM8
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 系统盘，系统盘个数不超过1块
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDisk []*DiskSpecInfo `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 需要绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 云数据盘，云数据盘总个数不超过15块
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisk []*DiskSpecInfo `json:"DataDisk,omitnil" name:"DataDisk"`

	// 本地数据盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDataDisk []*DiskSpecInfo `json:"LocalDataDisk,omitnil" name:"LocalDataDisk"`
}

type OpScope struct {
	// 操作范围，要操作的服务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInfoList []*ServiceBasicRestartInfo `json:"ServiceInfoList,omitnil" name:"ServiceInfoList"`
}

type OutterResource struct {
	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// 规格名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitnil" name:"StorageType"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 系统盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitnil" name:"RootSize"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// CPU个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`
}

type PartDetailPriceItem struct {
	// 类型包括：节点->node、系统盘->rootDisk、云数据盘->dataDisk、metaDB
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 单价（原价）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *float64 `json:"Price,omitnil" name:"Price"`

	// 单价（折扣价）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealCost *float64 `json:"RealCost,omitnil" name:"RealCost"`

	// 总价（折扣价）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *float64 `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 折扣
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *float64 `json:"Policy,omitnil" name:"Policy"`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsNum *int64 `json:"GoodsNum,omitnil" name:"GoodsNum"`
}

type PersistentVolumeContext struct {
	// 磁盘大小，单位为GB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *uint64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 磁盘类型。CLOUD_PREMIUM;CLOUD_SSD
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 磁盘数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskNum *int64 `json:"DiskNum,omitnil" name:"DiskNum"`
}

type Placement struct {
	// 实例所属的可用区，例如ap-guangzhou-1。该参数也可以通过调用[DescribeZones](https://cloud.tencent.com/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 实例所属项目ID。该参数可以通过调用[DescribeProject](https://cloud.tencent.com/document/api/651/78725) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type PodNewParameter struct {
	// TKE或EKS集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 自定义权限
	// 如：
	// {
	//   "apiVersion": "v1",
	//   "clusters": [
	//     {
	//       "cluster": {
	//         "certificate-authority-data": "xxxxxx==",
	//         "server": "https://xxxxx.com"
	//       },
	//       "name": "cls-xxxxx"
	//     }
	//   ],
	//   "contexts": [
	//     {
	//       "context": {
	//         "cluster": "cls-xxxxx",
	//         "user": "100014xxxxx"
	//       },
	//       "name": "cls-a44yhcxxxxxxxxxx"
	//     }
	//   ],
	//   "current-context": "cls-a4xxxx-context-default",
	//   "kind": "Config",
	//   "preferences": {},
	//   "users": [
	//     {
	//       "name": "100014xxxxx",
	//       "user": {
	//         "client-certificate-data": "xxxxxx",
	//         "client-key-data": "xxxxxx"
	//       }
	//     }
	//   ]
	// }
	Config *string `json:"Config,omitnil" name:"Config"`

	// 自定义参数
	// 如：
	// {
	//     "apiVersion": "apps/v1",
	//     "kind": "Deployment",
	//     "metadata": {
	//       "name": "test-deployment",
	//       "labels": {
	//         "app": "test"
	//       }
	//     },
	//     "spec": {
	//       "replicas": 3,
	//       "selector": {
	//         "matchLabels": {
	//           "app": "test-app"
	//         }
	//       },
	//       "template": {
	//         "metadata": {
	//           "annotations": {
	//             "your-organization.com/department-v1": "test-example-v1",
	//             "your-organization.com/department-v2": "test-example-v2"
	//           },
	//           "labels": {
	//             "app": "test-app",
	//             "environment": "production"
	//           }
	//         },
	//         "spec": {
	//           "nodeSelector": {
	//             "your-organization/node-test": "test-node"
	//           },
	//           "containers": [
	//             {
	//               "name": "nginx",
	//               "image": "nginx:1.14.2",
	//               "ports": [
	//                 {
	//                   "containerPort": 80
	//                 }
	//               ]
	//             }
	//           ],
	//           "affinity": {
	//             "nodeAffinity": {
	//               "requiredDuringSchedulingIgnoredDuringExecution": {
	//                 "nodeSelectorTerms": [
	//                   {
	//                     "matchExpressions": [
	//                       {
	//                         "key": "disk-type",
	//                         "operator": "In",
	//                         "values": [
	//                           "ssd",
	//                           "sas"
	//                         ]
	//                       },
	//                       {
	//                         "key": "cpu-num",
	//                         "operator": "Gt",
	//                         "values": [
	//                           "6"
	//                         ]
	//                       }
	//                     ]
	//                   }
	//                 ]
	//               }
	//             }
	//           }
	//         }
	//       }
	//     }
	//   }
	Parameter *string `json:"Parameter,omitnil" name:"Parameter"`
}

type PodNewSpec struct {
	// 外部资源提供者的标识符，例如"cls-a1cd23fa"。
	ResourceProviderIdentifier *string `json:"ResourceProviderIdentifier,omitnil" name:"ResourceProviderIdentifier"`

	// 外部资源提供者类型，例如"tke",当前仅支持"tke"。
	ResourceProviderType *string `json:"ResourceProviderType,omitnil" name:"ResourceProviderType"`

	// 资源的用途，即节点类型，当前仅支持"TASK"。
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// CPU核数。
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存大小，单位为GB。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// Eks集群-CPU类型，当前支持"intel"和"amd"
	CpuType *string `json:"CpuType,omitnil" name:"CpuType"`

	// Pod节点数据目录挂载信息。
	PodVolumes []*PodVolume `json:"PodVolumes,omitnil" name:"PodVolumes"`

	// 是否浮动规格，默认否
	// <li>true：代表是</li>
	// <li>false：代表否</li>
	EnableDynamicSpecFlag *bool `json:"EnableDynamicSpecFlag,omitnil" name:"EnableDynamicSpecFlag"`

	// 浮动规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DynamicPodSpec *DynamicPodSpec `json:"DynamicPodSpec,omitnil" name:"DynamicPodSpec"`

	// 代表vpc网络唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 代表vpc子网唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// pod name
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil" name:"PodName"`
}

type PodParameter struct {
	// TKE或EKS集群ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 自定义权限
	// 如：
	// {
	//   "apiVersion": "v1",
	//   "clusters": [
	//     {
	//       "cluster": {
	//         "certificate-authority-data": "xxxxxx==",
	//         "server": "https://xxxxx.com"
	//       },
	//       "name": "cls-xxxxx"
	//     }
	//   ],
	//   "contexts": [
	//     {
	//       "context": {
	//         "cluster": "cls-xxxxx",
	//         "user": "100014xxxxx"
	//       },
	//       "name": "cls-a44yhcxxxxxxxxxx"
	//     }
	//   ],
	//   "current-context": "cls-a4xxxx-context-default",
	//   "kind": "Config",
	//   "preferences": {},
	//   "users": [
	//     {
	//       "name": "100014xxxxx",
	//       "user": {
	//         "client-certificate-data": "xxxxxx",
	//         "client-key-data": "xxxxxx"
	//       }
	//     }
	//   ]
	// }
	Config *string `json:"Config,omitnil" name:"Config"`

	// 自定义参数
	// 如：
	// {
	//     "apiVersion": "apps/v1",
	//     "kind": "Deployment",
	//     "metadata": {
	//       "name": "test-deployment",
	//       "labels": {
	//         "app": "test"
	//       }
	//     },
	//     "spec": {
	//       "replicas": 3,
	//       "selector": {
	//         "matchLabels": {
	//           "app": "test-app"
	//         }
	//       },
	//       "template": {
	//         "metadata": {
	//           "annotations": {
	//             "your-organization.com/department-v1": "test-example-v1",
	//             "your-organization.com/department-v2": "test-example-v2"
	//           },
	//           "labels": {
	//             "app": "test-app",
	//             "environment": "production"
	//           }
	//         },
	//         "spec": {
	//           "nodeSelector": {
	//             "your-organization/node-test": "test-node"
	//           },
	//           "containers": [
	//             {
	//               "name": "nginx",
	//               "image": "nginx:1.14.2",
	//               "ports": [
	//                 {
	//                   "containerPort": 80
	//                 }
	//               ]
	//             }
	//           ],
	//           "affinity": {
	//             "nodeAffinity": {
	//               "requiredDuringSchedulingIgnoredDuringExecution": {
	//                 "nodeSelectorTerms": [
	//                   {
	//                     "matchExpressions": [
	//                       {
	//                         "key": "disk-type",
	//                         "operator": "In",
	//                         "values": [
	//                           "ssd",
	//                           "sas"
	//                         ]
	//                       },
	//                       {
	//                         "key": "cpu-num",
	//                         "operator": "Gt",
	//                         "values": [
	//                           "6"
	//                         ]
	//                       }
	//                     ]
	//                   }
	//                 ]
	//               }
	//             }
	//           }
	//         }
	//       }
	//     }
	//   }
	Parameter *string `json:"Parameter,omitnil" name:"Parameter"`
}

type PodSaleSpec struct {
	// 可售卖的资源规格，仅为以下值:"TASK","CORE","MASTER","ROUTER"。
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// Cpu核数。
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存数量，单位为GB。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 该规格资源可申请的最大数量。
	Number *uint64 `json:"Number,omitnil" name:"Number"`
}

type PodSpec struct {
	// 外部资源提供者的标识符，例如"cls-a1cd23fa"。
	ResourceProviderIdentifier *string `json:"ResourceProviderIdentifier,omitnil" name:"ResourceProviderIdentifier"`

	// 外部资源提供者类型，例如"tke",当前仅支持"tke"。
	ResourceProviderType *string `json:"ResourceProviderType,omitnil" name:"ResourceProviderType"`

	// 资源的用途，即节点类型，当前仅支持"TASK"。
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// CPU核数。
	Cpu *uint64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存大小，单位为GB。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 资源对宿主机的挂载点，指定的挂载点对应了宿主机的路径，该挂载点在Pod中作为数据存储目录使用。弃用
	DataVolumes []*string `json:"DataVolumes,omitnil" name:"DataVolumes"`

	// Eks集群-CPU类型，当前支持"intel"和"amd"
	CpuType *string `json:"CpuType,omitnil" name:"CpuType"`

	// Pod节点数据目录挂载信息。
	PodVolumes []*PodVolume `json:"PodVolumes,omitnil" name:"PodVolumes"`

	// 是否浮动规格，1是，0否
	IsDynamicSpec *uint64 `json:"IsDynamicSpec,omitnil" name:"IsDynamicSpec"`

	// 浮动规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DynamicPodSpec *DynamicPodSpec `json:"DynamicPodSpec,omitnil" name:"DynamicPodSpec"`

	// 代表vpc网络唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 代表vpc子网唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// pod name
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil" name:"PodName"`
}

type PodSpecInfo struct {
	// 使用Pod资源扩容时，指定的Pod规格以及来源等信息
	PodSpec *PodNewSpec `json:"PodSpec,omitnil" name:"PodSpec"`

	// POD自定义权限和自定义参数
	PodParameter *PodNewParameter `json:"PodParameter,omitnil" name:"PodParameter"`
}

type PodState struct {
	// pod的名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// pod uuid
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// pod的状态
	State *string `json:"State,omitnil" name:"State"`

	// pod处于该状态原因
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// pod所属集群
	OwnerCluster *string `json:"OwnerCluster,omitnil" name:"OwnerCluster"`

	// pod内存大小
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`
}

type PodVolume struct {
	// 存储类型，可为"pvc"，"hostpath"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeType *string `json:"VolumeType,omitnil" name:"VolumeType"`

	// 当VolumeType为"pvc"时，该字段生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PVCVolume *PersistentVolumeContext `json:"PVCVolume,omitnil" name:"PVCVolume"`

	// 当VolumeType为"hostpath"时，该字段生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostVolume *HostVolumeContext `json:"HostVolume,omitnil" name:"HostVolume"`
}

type PreExecuteFileSettings struct {
	// 脚本在COS上路径，已废弃
	Path *string `json:"Path,omitnil" name:"Path"`

	// 执行脚本参数
	Args []*string `json:"Args,omitnil" name:"Args"`

	// COS的Bucket名称，已废弃
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`

	// COS的Region名称，已废弃
	Region *string `json:"Region,omitnil" name:"Region"`

	// COS的Domain数据，已废弃
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 执行顺序
	RunOrder *int64 `json:"RunOrder,omitnil" name:"RunOrder"`

	// resourceAfter 或 clusterAfter
	WhenRun *string `json:"WhenRun,omitnil" name:"WhenRun"`

	// 脚本文件名，已废弃
	CosFileName *string `json:"CosFileName,omitnil" name:"CosFileName"`

	// 脚本的cos地址
	CosFileURI *string `json:"CosFileURI,omitnil" name:"CosFileURI"`

	// cos的SecretId
	CosSecretId *string `json:"CosSecretId,omitnil" name:"CosSecretId"`

	// Cos的SecretKey
	CosSecretKey *string `json:"CosSecretKey,omitnil" name:"CosSecretKey"`

	// cos的appid，已废弃
	AppId *string `json:"AppId,omitnil" name:"AppId"`
}

type PriceDetail struct {
	// 节点ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 价格计算公式
	Formula *string `json:"Formula,omitnil" name:"Formula"`

	// 原价
	OriginalCost *float64 `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// 折扣价
	DiscountCost *float64 `json:"DiscountCost,omitnil" name:"DiscountCost"`
}

type PriceResource struct {
	// 需要的规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *uint64 `json:"StorageType,omitnil" name:"StorageType"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 系统盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitnil" name:"RootSize"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// 核心数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 云盘列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiDisks []*MultiDisk `json:"MultiDisks,omitnil" name:"MultiDisks"`

	// 磁盘数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCnt *int64 `json:"DiskCnt,omitnil" name:"DiskCnt"`

	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 磁盘数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskNum *int64 `json:"DiskNum,omitnil" name:"DiskNum"`

	// 本地盘的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDiskNum *int64 `json:"LocalDiskNum,omitnil" name:"LocalDiskNum"`
}

type QuotaEntity struct {
	// 已使用配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedQuota *int64 `json:"UsedQuota,omitnil" name:"UsedQuota"`

	// 剩余配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainingQuota *int64 `json:"RemainingQuota,omitnil" name:"RemainingQuota"`

	// 总配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalQuota *int64 `json:"TotalQuota,omitnil" name:"TotalQuota"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

type RenewInstancesInfo struct {
	// 节点资源ID
	EmrResourceId *string `json:"EmrResourceId,omitnil" name:"EmrResourceId"`

	// 节点类型。0:common节点；1:master节点
	// ；2:core节点；3:task节点
	Flag *int64 `json:"Flag,omitnil" name:"Flag"`

	// 内网IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 节点内存描述
	MemDesc *string `json:"MemDesc,omitnil" name:"MemDesc"`

	// 节点核数
	CpuNum *int64 `json:"CpuNum,omitnil" name:"CpuNum"`

	// 硬盘大小
	DiskSize *string `json:"DiskSize,omitnil" name:"DiskSize"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 节点规格
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// 磁盘类型
	StorageType *int64 `json:"StorageType,omitnil" name:"StorageType"`
}

type Resource struct {
	// 节点规格描述，如CVM.SA2。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// 存储类型
	// 取值范围：
	// <li>4：表示云SSD。</li>
	// <li>5：表示高效云盘。</li>
	// <li>6：表示增强型SSD云硬盘。</li>
	// <li>11：表示吞吐型云硬盘。</li>
	// <li>12：表示极速型SSD云硬盘。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitnil" name:"StorageType"`

	// 磁盘类型
	// 取值范围：
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_BASIC：表示云硬盘。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 内存容量,单位为M
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitnil" name:"MemSize"`

	// CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// 数据盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 系统盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitnil" name:"RootSize"`

	// 云盘列表，当数据盘为一块云盘时，直接使用DiskType和DiskSize参数，超出部分使用MultiDisks
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiDisks []*MultiDisk `json:"MultiDisks,omitnil" name:"MultiDisks"`

	// 需要绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 规格类型，如S2.MEDIUM8
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 本地盘数量，该字段已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDiskNum *uint64 `json:"LocalDiskNum,omitnil" name:"LocalDiskNum"`

	// 本地盘数量，如2
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskNum *uint64 `json:"DiskNum,omitnil" name:"DiskNum"`
}

// Predefined struct for user
type RunJobFlowRequestParams struct {
	// 作业名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否新创建集群。
	// true，新创建集群，则使用Instance中的参数进行集群创建。
	// false，使用已有集群，则通过InstanceId传入。
	CreateCluster *bool `json:"CreateCluster,omitnil" name:"CreateCluster"`

	// 作业流程执行步骤。
	Steps []*Step `json:"Steps,omitnil" name:"Steps"`

	// 作业流程正常完成时，集群的处理方式，可选择:
	// Terminate 销毁集群。
	// Reserve 保留集群。
	InstancePolicy *string `json:"InstancePolicy,omitnil" name:"InstancePolicy"`

	// 只有CreateCluster为true时生效，目前只支持EMR版本，例如EMR-2.2.0，不支持ClickHouse和Druid版本。
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// 只在CreateCluster为true时生效。
	// true 表示安装kerberos，false表示不安装kerberos。
	SecurityClusterFlag *bool `json:"SecurityClusterFlag,omitnil" name:"SecurityClusterFlag"`

	// 只在CreateCluster为true时生效。
	// 新建集群时，要安装的软件列表。
	Software []*string `json:"Software,omitnil" name:"Software"`

	// 引导脚本。
	BootstrapActions []*BootstrapAction `json:"BootstrapActions,omitnil" name:"BootstrapActions"`

	// 指定配置创建集群。
	Configurations []*Configuration `json:"Configurations,omitnil" name:"Configurations"`

	// 作业日志保存地址。
	LogUri *string `json:"LogUri,omitnil" name:"LogUri"`

	// 只在CreateCluster为false时生效。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 自定义应用角色，大数据应用访问外部服务时使用的角色，默认为"EME_QCSRole"。
	ApplicationRole *string `json:"ApplicationRole,omitnil" name:"ApplicationRole"`

	// 重入标签，用来可重入检查，防止在一段时间内，创建相同的流程作业。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 只在CreateCluster为true时生效，使用该配置创建集群。
	Instance *ClusterSetting `json:"Instance,omitnil" name:"Instance"`
}

type RunJobFlowRequest struct {
	*tchttp.BaseRequest
	
	// 作业名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否新创建集群。
	// true，新创建集群，则使用Instance中的参数进行集群创建。
	// false，使用已有集群，则通过InstanceId传入。
	CreateCluster *bool `json:"CreateCluster,omitnil" name:"CreateCluster"`

	// 作业流程执行步骤。
	Steps []*Step `json:"Steps,omitnil" name:"Steps"`

	// 作业流程正常完成时，集群的处理方式，可选择:
	// Terminate 销毁集群。
	// Reserve 保留集群。
	InstancePolicy *string `json:"InstancePolicy,omitnil" name:"InstancePolicy"`

	// 只有CreateCluster为true时生效，目前只支持EMR版本，例如EMR-2.2.0，不支持ClickHouse和Druid版本。
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// 只在CreateCluster为true时生效。
	// true 表示安装kerberos，false表示不安装kerberos。
	SecurityClusterFlag *bool `json:"SecurityClusterFlag,omitnil" name:"SecurityClusterFlag"`

	// 只在CreateCluster为true时生效。
	// 新建集群时，要安装的软件列表。
	Software []*string `json:"Software,omitnil" name:"Software"`

	// 引导脚本。
	BootstrapActions []*BootstrapAction `json:"BootstrapActions,omitnil" name:"BootstrapActions"`

	// 指定配置创建集群。
	Configurations []*Configuration `json:"Configurations,omitnil" name:"Configurations"`

	// 作业日志保存地址。
	LogUri *string `json:"LogUri,omitnil" name:"LogUri"`

	// 只在CreateCluster为false时生效。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 自定义应用角色，大数据应用访问外部服务时使用的角色，默认为"EME_QCSRole"。
	ApplicationRole *string `json:"ApplicationRole,omitnil" name:"ApplicationRole"`

	// 重入标签，用来可重入检查，防止在一段时间内，创建相同的流程作业。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 只在CreateCluster为true时生效，使用该配置创建集群。
	Instance *ClusterSetting `json:"Instance,omitnil" name:"Instance"`
}

func (r *RunJobFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunJobFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "CreateCluster")
	delete(f, "Steps")
	delete(f, "InstancePolicy")
	delete(f, "ProductVersion")
	delete(f, "SecurityClusterFlag")
	delete(f, "Software")
	delete(f, "BootstrapActions")
	delete(f, "Configurations")
	delete(f, "LogUri")
	delete(f, "InstanceId")
	delete(f, "ApplicationRole")
	delete(f, "ClientToken")
	delete(f, "Instance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunJobFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunJobFlowResponseParams struct {
	// 作业流程ID。
	JobFlowId *int64 `json:"JobFlowId,omitnil" name:"JobFlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RunJobFlowResponse struct {
	*tchttp.BaseResponse
	Response *RunJobFlowResponseParams `json:"Response"`
}

func (r *RunJobFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunJobFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutClusterRequestParams struct {
	// 节点计费模式。取值范围：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：按小时后付费。</li>
	// <li>SPOTPAID：竞价付费（仅支持TASK节点）。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 集群实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 扩容节点类型以及数量
	ScaleOutNodeConfig *ScaleOutNodeConfig `json:"ScaleOutNodeConfig,omitnil" name:"ScaleOutNodeConfig"`

	// 唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，例如 a9a90aa6-751a-41b6-aad6-fae36063280
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil" name:"ScriptBootstrapActionConfig"`

	// 扩容部署服务，新增节点将默认继承当前节点类型中所部署服务，部署服务含默认可选服务，该参数仅支持可选服务填写，如：存量task节点已部署HDFS、YARN、impala；使用api扩容task节不部署impala时，部署服务仅填写HDFS、YARN。[组件名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)。
	SoftDeployInfo []*int64 `json:"SoftDeployInfo,omitnil" name:"SoftDeployInfo"`

	// 部署进程，默认部署扩容服务的全部进程，支持修改部署进程，如：当前task节点部署服务为：HDFS、YARN、impala，默认部署服务为：DataNode,NodeManager,ImpalaServer，若用户需修改部署进程信息，部署进程：	DataNode,NodeManager,ImpalaServerCoordinator或DataNode,NodeManager,ImpalaServerExecutor。[进程名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)。
	ServiceNodeInfo []*int64 `json:"ServiceNodeInfo,omitnil" name:"ServiceNodeInfo"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/product/213/17810)的返回值中的DisasterRecoverGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// 扩容节点绑定标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 扩容所选资源类型，可选范围为"host","pod"，host为普通的CVM资源，Pod为TKE集群或EKS集群提供的资源
	HardwareSourceType *string `json:"HardwareSourceType,omitnil" name:"HardwareSourceType"`

	// Pod相关资源信息
	PodSpecInfo *PodSpecInfo `json:"PodSpecInfo,omitnil" name:"PodSpecInfo"`

	// 使用clickhouse集群扩容时，选择的机器分组名称
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil" name:"ClickHouseClusterName"`

	// 使用clickhouse集群扩容时，选择的机器分组类型。new为新增，old为选择旧分组
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil" name:"ClickHouseClusterType"`

	// 扩容指定 Yarn Node Label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil" name:"YarnNodeLabel"`

	// 扩容后是否启动服务，默认取值否
	// <li>true：是</li>
	// <li>false：否</li>
	EnableStartServiceFlag *bool `json:"EnableStartServiceFlag,omitnil" name:"EnableStartServiceFlag"`

	// 规格设置
	ResourceSpec *NodeResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// 实例所属的可用区，例如ap-guangzhou-1。该参数也可以通过调用[DescribeZones](https://cloud.tencent.com/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 子网，默认是集群创建时的子网
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type ScaleOutClusterRequest struct {
	*tchttp.BaseRequest
	
	// 节点计费模式。取值范围：
	// <li>PREPAID：预付费，即包年包月。</li>
	// <li>POSTPAID_BY_HOUR：按小时后付费。</li>
	// <li>SPOTPAID：竞价付费（仅支持TASK节点）。</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 集群实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 扩容节点类型以及数量
	ScaleOutNodeConfig *ScaleOutNodeConfig `json:"ScaleOutNodeConfig,omitnil" name:"ScaleOutNodeConfig"`

	// 唯一随机标识，时效5分钟，需要调用者指定 防止客户端重新创建资源，例如 a9a90aa6-751a-41b6-aad6-fae36063280
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	ScriptBootstrapActionConfig []*ScriptBootstrapActionConfig `json:"ScriptBootstrapActionConfig,omitnil" name:"ScriptBootstrapActionConfig"`

	// 扩容部署服务，新增节点将默认继承当前节点类型中所部署服务，部署服务含默认可选服务，该参数仅支持可选服务填写，如：存量task节点已部署HDFS、YARN、impala；使用api扩容task节不部署impala时，部署服务仅填写HDFS、YARN。[组件名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)。
	SoftDeployInfo []*int64 `json:"SoftDeployInfo,omitnil" name:"SoftDeployInfo"`

	// 部署进程，默认部署扩容服务的全部进程，支持修改部署进程，如：当前task节点部署服务为：HDFS、YARN、impala，默认部署服务为：DataNode,NodeManager,ImpalaServer，若用户需修改部署进程信息，部署进程：	DataNode,NodeManager,ImpalaServerCoordinator或DataNode,NodeManager,ImpalaServerExecutor。[进程名对应的映射关系表](https://cloud.tencent.com/document/product/589/98760)。
	ServiceNodeInfo []*int64 `json:"ServiceNodeInfo,omitnil" name:"ServiceNodeInfo"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/product/213/17810)的返回值中的DisasterRecoverGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// 扩容节点绑定标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 扩容所选资源类型，可选范围为"host","pod"，host为普通的CVM资源，Pod为TKE集群或EKS集群提供的资源
	HardwareSourceType *string `json:"HardwareSourceType,omitnil" name:"HardwareSourceType"`

	// Pod相关资源信息
	PodSpecInfo *PodSpecInfo `json:"PodSpecInfo,omitnil" name:"PodSpecInfo"`

	// 使用clickhouse集群扩容时，选择的机器分组名称
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil" name:"ClickHouseClusterName"`

	// 使用clickhouse集群扩容时，选择的机器分组类型。new为新增，old为选择旧分组
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil" name:"ClickHouseClusterType"`

	// 扩容指定 Yarn Node Label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil" name:"YarnNodeLabel"`

	// 扩容后是否启动服务，默认取值否
	// <li>true：是</li>
	// <li>false：否</li>
	EnableStartServiceFlag *bool `json:"EnableStartServiceFlag,omitnil" name:"EnableStartServiceFlag"`

	// 规格设置
	ResourceSpec *NodeResourceSpec `json:"ResourceSpec,omitnil" name:"ResourceSpec"`

	// 实例所属的可用区，例如ap-guangzhou-1。该参数也可以通过调用[DescribeZones](https://cloud.tencent.com/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 子网，默认是集群创建时的子网
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

func (r *ScaleOutClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceChargeType")
	delete(f, "InstanceId")
	delete(f, "ScaleOutNodeConfig")
	delete(f, "ClientToken")
	delete(f, "InstanceChargePrepaid")
	delete(f, "ScriptBootstrapActionConfig")
	delete(f, "SoftDeployInfo")
	delete(f, "ServiceNodeInfo")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "Tags")
	delete(f, "HardwareSourceType")
	delete(f, "PodSpecInfo")
	delete(f, "ClickHouseClusterName")
	delete(f, "ClickHouseClusterType")
	delete(f, "YarnNodeLabel")
	delete(f, "EnableStartServiceFlag")
	delete(f, "ResourceSpec")
	delete(f, "Zone")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutClusterResponseParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 客户端Token。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 扩容流程ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ScaleOutClusterResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutClusterResponseParams `json:"Response"`
}

func (r *ScaleOutClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 客户端Token。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 引导操作脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil" name:"PreExecutedFileSettings"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// 扩容时不需要安装的进程。
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitnil" name:"UnNecessaryNodeList"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitnil" name:"RouterCount"`

	// 部署的服务。
	// <li>SoftDeployInfo和ServiceNodeInfo是同组参数，和UnNecessaryNodeList参数互斥。</li>
	// <li>建议使用SoftDeployInfo和ServiceNodeInfo组合。</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitnil" name:"SoftDeployInfo"`

	// 启动的进程。
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitnil" name:"ServiceNodeInfo"`

	// 分散置放群组ID列表，当前仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// 扩容节点绑定标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 扩容所选资源类型，可选范围为"host","pod"，host为普通的CVM资源，Pod为TKE集群或EKS集群提供的资源
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// 使用Pod资源扩容时，指定的Pod规格以及来源等信息
	PodSpec *PodSpec `json:"PodSpec,omitnil" name:"PodSpec"`

	// 使用clickhouse集群扩容时，选择的机器分组名称
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil" name:"ClickHouseClusterName"`

	// 使用clickhouse集群扩容时，选择的机器分组类型。new为新增，old为选择旧分组
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil" name:"ClickHouseClusterType"`

	// 规则扩容指定 yarn node label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil" name:"YarnNodeLabel"`

	// POD自定义权限和自定义参数
	PodParameter *PodParameter `json:"PodParameter,omitnil" name:"PodParameter"`

	// 扩容的Master节点的数量。
	// 使用clickhouse集群扩容时，该参数不生效。
	// 使用kafka集群扩容时，该参数不生效。
	// 当HardwareResourceType=POD时，该参数不生效。
	MasterCount *uint64 `json:"MasterCount,omitnil" name:"MasterCount"`

	// 扩容后是否启动服务，true：启动，false：不启动
	StartServiceAfterScaleOut *string `json:"StartServiceAfterScaleOut,omitnil" name:"StartServiceAfterScaleOut"`

	// 可用区，默认是集群的主可用区
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 子网，默认是集群创建时的子网
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 预设配置组
	ScaleOutServiceConfAssign *string `json:"ScaleOutServiceConfAssign,omitnil" name:"ScaleOutServiceConfAssign"`

	// 0表示关闭自动续费，1表示开启自动续费
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 客户端Token。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 引导操作脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitnil" name:"PreExecutedFileSettings"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitnil" name:"TaskCount"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitnil" name:"CoreCount"`

	// 扩容时不需要安装的进程。
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitnil" name:"UnNecessaryNodeList"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitnil" name:"RouterCount"`

	// 部署的服务。
	// <li>SoftDeployInfo和ServiceNodeInfo是同组参数，和UnNecessaryNodeList参数互斥。</li>
	// <li>建议使用SoftDeployInfo和ServiceNodeInfo组合。</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitnil" name:"SoftDeployInfo"`

	// 启动的进程。
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitnil" name:"ServiceNodeInfo"`

	// 分散置放群组ID列表，当前仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil" name:"DisasterRecoverGroupIds"`

	// 扩容节点绑定标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 扩容所选资源类型，可选范围为"host","pod"，host为普通的CVM资源，Pod为TKE集群或EKS集群提供的资源
	HardwareResourceType *string `json:"HardwareResourceType,omitnil" name:"HardwareResourceType"`

	// 使用Pod资源扩容时，指定的Pod规格以及来源等信息
	PodSpec *PodSpec `json:"PodSpec,omitnil" name:"PodSpec"`

	// 使用clickhouse集群扩容时，选择的机器分组名称
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitnil" name:"ClickHouseClusterName"`

	// 使用clickhouse集群扩容时，选择的机器分组类型。new为新增，old为选择旧分组
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitnil" name:"ClickHouseClusterType"`

	// 规则扩容指定 yarn node label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitnil" name:"YarnNodeLabel"`

	// POD自定义权限和自定义参数
	PodParameter *PodParameter `json:"PodParameter,omitnil" name:"PodParameter"`

	// 扩容的Master节点的数量。
	// 使用clickhouse集群扩容时，该参数不生效。
	// 使用kafka集群扩容时，该参数不生效。
	// 当HardwareResourceType=POD时，该参数不生效。
	MasterCount *uint64 `json:"MasterCount,omitnil" name:"MasterCount"`

	// 扩容后是否启动服务，true：启动，false：不启动
	StartServiceAfterScaleOut *string `json:"StartServiceAfterScaleOut,omitnil" name:"StartServiceAfterScaleOut"`

	// 可用区，默认是集群的主可用区
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 子网，默认是集群创建时的子网
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 预设配置组
	ScaleOutServiceConfAssign *string `json:"ScaleOutServiceConfAssign,omitnil" name:"ScaleOutServiceConfAssign"`

	// 0表示关闭自动续费，1表示开启自动续费
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "InstanceId")
	delete(f, "PayMode")
	delete(f, "ClientToken")
	delete(f, "PreExecutedFileSettings")
	delete(f, "TaskCount")
	delete(f, "CoreCount")
	delete(f, "UnNecessaryNodeList")
	delete(f, "RouterCount")
	delete(f, "SoftDeployInfo")
	delete(f, "ServiceNodeInfo")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "Tags")
	delete(f, "HardwareResourceType")
	delete(f, "PodSpec")
	delete(f, "ClickHouseClusterName")
	delete(f, "ClickHouseClusterType")
	delete(f, "YarnNodeLabel")
	delete(f, "PodParameter")
	delete(f, "MasterCount")
	delete(f, "StartServiceAfterScaleOut")
	delete(f, "ZoneId")
	delete(f, "SubnetId")
	delete(f, "ScaleOutServiceConfAssign")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`

	// 客户端Token。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 扩容流程ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// 大订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillId *string `json:"BillId,omitnil" name:"BillId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstanceResponseParams `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleOutNodeConfig struct {
	// 扩容节点类型取值范围：
	//   <li>MASTER</li>
	//   <li>TASK</li>
	//   <li>CORE</li>
	//   <li>ROUTER</li>
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// 扩容节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`
}

type SceneSoftwareConfig struct {
	// 部署的组件列表。不同的EMR产品版本ProductVersion 对应不同可选组件列表，不同产品版本可选组件列表查询：[组件版本](https://cloud.tencent.com/document/product/589/20279) ；
	// 填写实例值：hive、flink。
	Software []*string `json:"Software,omitnil" name:"Software"`

	// 默认Hadoop-Default,[场景查询](https://cloud.tencent.com/document/product/589/14624)场景化取值范围：
	// Hadoop-Kudu
	// Hadoop-Zookeeper
	// Hadoop-Presto
	// Hadoop-Hbase
	// Hadoop-Default
	SceneName *string `json:"SceneName,omitnil" name:"SceneName"`
}

type ScriptBootstrapActionConfig struct {
	// 脚本的cos地址，参照格式：https://beijing-111111.cos.ap-beijing.myqcloud.com/data/test.sh查询cos存储桶列表：[存储桶列表](https://console.cloud.tencent.com/cos/bucket)
	CosFileURI *string `json:"CosFileURI,omitnil" name:"CosFileURI"`

	// 引导脚步执行时机范围
	// <li>resourceAfter：节点初始化后</li>
	// <li>clusterAfter：集群启动后</li>
	// <li>clusterBefore：集群启动前</li>
	ExecutionMoment *string `json:"ExecutionMoment,omitnil" name:"ExecutionMoment"`

	// 执行脚本参数，参数格式请遵循标准Shell规范
	Args []*string `json:"Args,omitnil" name:"Args"`

	// 脚本文件名
	CosFileName *string `json:"CosFileName,omitnil" name:"CosFileName"`
}

type SearchItem struct {
	// 支持搜索的类型
	SearchType *string `json:"SearchType,omitnil" name:"SearchType"`

	// 支持搜索的值
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
}

type ServiceBasicRestartInfo struct {
	// 服务名，必填，如HDFS
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 如果没传，则表示所有进程
	ComponentInfoList []*ComponentBasicRestartInfo `json:"ComponentInfoList,omitnil" name:"ComponentInfoList"`
}

type ShortNodeInfo struct {
	// 节点类型，Master/Core/Task/Router/Common
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// 节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSize *uint64 `json:"NodeSize,omitnil" name:"NodeSize"`
}

type SoftDependInfo struct {
	// 组件名称
	SoftName *string `json:"SoftName,omitnil" name:"SoftName"`

	// 是否必选
	Required *bool `json:"Required,omitnil" name:"Required"`
}

// Predefined struct for user
type StartStopServiceOrMonitorRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 操作类型，当前支持
	// <li>StartService：启动服务</li>
	// <li>StopService：停止服务</li>
	// <li>StartMonitor：退出维护</li>
	// <li>StopMonitor：进入维护</li>
	// <li>RestartService：重启服务 如果操作类型选择重启服务 StrategyConfig操作策略则是必填项</li>
	OpType *string `json:"OpType,omitnil" name:"OpType"`

	// 操作范围
	OpScope *OpScope `json:"OpScope,omitnil" name:"OpScope"`

	// 操作策略
	StrategyConfig *StrategyConfig `json:"StrategyConfig,omitnil" name:"StrategyConfig"`
}

type StartStopServiceOrMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 操作类型，当前支持
	// <li>StartService：启动服务</li>
	// <li>StopService：停止服务</li>
	// <li>StartMonitor：退出维护</li>
	// <li>StopMonitor：进入维护</li>
	// <li>RestartService：重启服务 如果操作类型选择重启服务 StrategyConfig操作策略则是必填项</li>
	OpType *string `json:"OpType,omitnil" name:"OpType"`

	// 操作范围
	OpScope *OpScope `json:"OpScope,omitnil" name:"OpScope"`

	// 操作策略
	StrategyConfig *StrategyConfig `json:"StrategyConfig,omitnil" name:"StrategyConfig"`
}

func (r *StartStopServiceOrMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStopServiceOrMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OpType")
	delete(f, "OpScope")
	delete(f, "StrategyConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartStopServiceOrMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStopServiceOrMonitorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartStopServiceOrMonitorResponse struct {
	*tchttp.BaseResponse
	Response *StartStopServiceOrMonitorResponseParams `json:"Response"`
}

func (r *StartStopServiceOrMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStopServiceOrMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Step struct {
	// 执行步骤名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 执行动作。
	ExecutionStep *Execution `json:"ExecutionStep,omitnil" name:"ExecutionStep"`

	// 执行失败策略。
	// 1. TERMINATE_CLUSTER 执行失败时退出并销毁集群。
	// 2. CONTINUE 执行失败时跳过并执行后续步骤。
	ActionOnFailure *string `json:"ActionOnFailure,omitnil" name:"ActionOnFailure"`

	// 指定执行Step时的用户名，非必须，默认为hadoop。
	User *string `json:"User,omitnil" name:"User"`
}

type StrategyConfig struct {
	// 0:关闭滚动重启
	// 1:开启滚动启动
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollingRestartSwitch *int64 `json:"RollingRestartSwitch,omitnil" name:"RollingRestartSwitch"`

	// 滚动重启每批次的重启数量，最大重启台数为 99999 台
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchSize *int64 `json:"BatchSize,omitnil" name:"BatchSize"`

	// 滚动重启每批停止等待时间 ,最大间隔为 5 分钟 单位是秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeWait *int64 `json:"TimeWait,omitnil" name:"TimeWait"`

	// 操作失败处理策略，0:失败阻塞, 1:失败自动跳过
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealOnFail *int64 `json:"DealOnFail,omitnil" name:"DealOnFail"`
}

type SubnetInfo struct {
	// 子网信息（名字）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 子网信息（ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

// Predefined struct for user
type SyncPodStateRequestParams struct {
	// EmrService中pod状态信息
	Message *PodState `json:"Message,omitnil" name:"Message"`
}

type SyncPodStateRequest struct {
	*tchttp.BaseRequest
	
	// EmrService中pod状态信息
	Message *PodState `json:"Message,omitnil" name:"Message"`
}

func (r *SyncPodStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPodStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Message")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncPodStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncPodStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SyncPodStateResponse struct {
	*tchttp.BaseResponse
	Response *SyncPodStateResponseParams `json:"Response"`
}

func (r *SyncPodStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPodStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

// Predefined struct for user
type TerminateClusterNodesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 销毁资源列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil" name:"CvmInstanceIds"`

	// 销毁节点类型取值范围：
	//   <li>MASTER</li>
	//   <li>TASK</li>
	//   <li>CORE</li>
	//   <li>ROUTER</li>
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// 优雅缩容开关
	//   <li>true:开启</li>
	//   <li>false:不开启</li>
	GraceDownFlag *bool `json:"GraceDownFlag,omitnil" name:"GraceDownFlag"`

	// 优雅缩容等待时间,时间范围60到1800  单位秒
	GraceDownTime *int64 `json:"GraceDownTime,omitnil" name:"GraceDownTime"`
}

type TerminateClusterNodesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 销毁资源列表
	CvmInstanceIds []*string `json:"CvmInstanceIds,omitnil" name:"CvmInstanceIds"`

	// 销毁节点类型取值范围：
	//   <li>MASTER</li>
	//   <li>TASK</li>
	//   <li>CORE</li>
	//   <li>ROUTER</li>
	NodeFlag *string `json:"NodeFlag,omitnil" name:"NodeFlag"`

	// 优雅缩容开关
	//   <li>true:开启</li>
	//   <li>false:不开启</li>
	GraceDownFlag *bool `json:"GraceDownFlag,omitnil" name:"GraceDownFlag"`

	// 优雅缩容等待时间,时间范围60到1800  单位秒
	GraceDownTime *int64 `json:"GraceDownTime,omitnil" name:"GraceDownTime"`
}

func (r *TerminateClusterNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateClusterNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CvmInstanceIds")
	delete(f, "NodeFlag")
	delete(f, "GraceDownFlag")
	delete(f, "GraceDownTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateClusterNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateClusterNodesResponseParams struct {
	// 缩容流程ID。
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateClusterNodesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateClusterNodesResponseParams `json:"Response"`
}

func (r *TerminateClusterNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstanceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 销毁节点ID。该参数为预留参数，用户无需配置。
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

type TerminateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 销毁节点ID。该参数为预留参数，用户无需配置。
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

func (r *TerminateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateInstanceResponseParams `json:"Response"`
}

func (r *TerminateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTasksRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 待销毁节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr)查询。
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

type TerminateTasksRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 待销毁节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr)查询。
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

func (r *TerminateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTasksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateTasksResponse struct {
	*tchttp.BaseResponse
	Response *TerminateTasksResponseParams `json:"Response"`
}

func (r *TerminateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopologyInfo struct {
	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 可用区信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetInfoList []*SubnetInfo `json:"SubnetInfoList,omitnil" name:"SubnetInfoList"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInfoList []*ShortNodeInfo `json:"NodeInfoList,omitnil" name:"NodeInfoList"`
}

type UpdateInstanceSettings struct {
	// 内存容量，单位为G
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// CPU核数
	CPUCores *uint64 `json:"CPUCores,omitnil" name:"CPUCores"`

	// 机器资源ID（EMR测资源标识）
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 变配机器规格
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`
}

type UserAndGroup struct {
	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户组
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`
}

type UserInfoForUserManager struct {
	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户所属的组
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`

	// 密码
	PassWord *string `json:"PassWord,omitnil" name:"PassWord"`

	// 备注
	ReMark *string `json:"ReMark,omitnil" name:"ReMark"`
}

type UserManagerFilter struct {
	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`
}

type UserManagerUserBriefInfo struct {
	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户所属的组
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`

	// Manager表示管理员、NormalUser表示普通用户
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 用户创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 是否可以下载用户对应的keytab文件，对开启kerberos的集群才有意义
	SupportDownLoadKeyTab *bool `json:"SupportDownLoadKeyTab,omitnil" name:"SupportDownLoadKeyTab"`

	// keytab文件的下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownLoadKeyTabUrl *string `json:"DownLoadKeyTabUrl,omitnil" name:"DownLoadKeyTabUrl"`
}

type VPCSettings struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type VirtualPrivateCloud struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type YarnApplication struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil" name:"User"`

	// 应用名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Queue *string `json:"Queue,omitnil" name:"Queue"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitnil" name:"ApplicationType"`

	// 运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElapsedTime *string `json:"ElapsedTime,omitnil" name:"ElapsedTime"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil" name:"State"`

	// 最终状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalStatus *string `json:"FinalStatus,omitnil" name:"FinalStatus"`

	// 进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitnil" name:"Progress"`

	// 开始时间毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartedTime *int64 `json:"StartedTime,omitnil" name:"StartedTime"`

	// 结束时间毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedTime *int64 `json:"FinishedTime,omitnil" name:"FinishedTime"`

	// 申请内存MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedMB *int64 `json:"AllocatedMB,omitnil" name:"AllocatedMB"`

	// 申请VCores
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocatedVCores *int64 `json:"AllocatedVCores,omitnil" name:"AllocatedVCores"`

	// 运行的Containers数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningContainers *int64 `json:"RunningContainers,omitnil" name:"RunningContainers"`

	// 内存MB*时间秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemorySeconds *int64 `json:"MemorySeconds,omitnil" name:"MemorySeconds"`

	// VCores*时间秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	VCoreSeconds *int64 `json:"VCoreSeconds,omitnil" name:"VCoreSeconds"`

	// 队列资源占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueUsagePercentage *float64 `json:"QueueUsagePercentage,omitnil" name:"QueueUsagePercentage"`

	// 集群资源占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsagePercentage *float64 `json:"ClusterUsagePercentage,omitnil" name:"ClusterUsagePercentage"`

	// 预占用的内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreemptedResourceMB *int64 `json:"PreemptedResourceMB,omitnil" name:"PreemptedResourceMB"`

	// 预占用的VCore
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreemptedResourceVCores *int64 `json:"PreemptedResourceVCores,omitnil" name:"PreemptedResourceVCores"`

	// 预占的非应用程序主节点容器数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumNonAMContainerPreempted *int64 `json:"NumNonAMContainerPreempted,omitnil" name:"NumNonAMContainerPreempted"`

	// AM预占用的容器数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumAMContainerPreempted *int64 `json:"NumAMContainerPreempted,omitnil" name:"NumAMContainerPreempted"`

	// Map总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MapsTotal *int64 `json:"MapsTotal,omitnil" name:"MapsTotal"`

	// 完成的Map数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MapsCompleted *int64 `json:"MapsCompleted,omitnil" name:"MapsCompleted"`

	// Reduce总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReducesTotal *int64 `json:"ReducesTotal,omitnil" name:"ReducesTotal"`

	// 完成的Reduce数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReducesCompleted *int64 `json:"ReducesCompleted,omitnil" name:"ReducesCompleted"`

	// 平均Map时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgMapTime *int64 `json:"AvgMapTime,omitnil" name:"AvgMapTime"`

	// 平均Reduce时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgReduceTime *int64 `json:"AvgReduceTime,omitnil" name:"AvgReduceTime"`

	// 平均Shuffle时间毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgShuffleTime *int64 `json:"AvgShuffleTime,omitnil" name:"AvgShuffleTime"`

	// 平均Merge时间毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgMergeTime *int64 `json:"AvgMergeTime,omitnil" name:"AvgMergeTime"`

	// 失败的Reduce执行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReduceAttempts *int64 `json:"FailedReduceAttempts,omitnil" name:"FailedReduceAttempts"`

	// Kill的Reduce执行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KilledReduceAttempts *int64 `json:"KilledReduceAttempts,omitnil" name:"KilledReduceAttempts"`

	// 成功的Reduce执行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessfulReduceAttempts *int64 `json:"SuccessfulReduceAttempts,omitnil" name:"SuccessfulReduceAttempts"`

	// 失败的Map执行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedMapAttempts *int64 `json:"FailedMapAttempts,omitnil" name:"FailedMapAttempts"`

	// Kill的Map执行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KilledMapAttempts *int64 `json:"KilledMapAttempts,omitnil" name:"KilledMapAttempts"`

	// 成功的Map执行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessfulMapAttempts *int64 `json:"SuccessfulMapAttempts,omitnil" name:"SuccessfulMapAttempts"`

	// GC毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	GcTimeMillis *int64 `json:"GcTimeMillis,omitnil" name:"GcTimeMillis"`

	// Map使用的VCore毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	VCoreMillisMaps *int64 `json:"VCoreMillisMaps,omitnil" name:"VCoreMillisMaps"`

	// Map使用的内存毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MbMillisMaps *int64 `json:"MbMillisMaps,omitnil" name:"MbMillisMaps"`

	// Reduce使用的VCore毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	VCoreMillisReduces *int64 `json:"VCoreMillisReduces,omitnil" name:"VCoreMillisReduces"`

	// Reduce使用的内存毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MbMillisReduces *int64 `json:"MbMillisReduces,omitnil" name:"MbMillisReduces"`

	// 启动Map的总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalLaunchedMaps *int64 `json:"TotalLaunchedMaps,omitnil" name:"TotalLaunchedMaps"`

	// 启动Reduce的总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalLaunchedReduces *int64 `json:"TotalLaunchedReduces,omitnil" name:"TotalLaunchedReduces"`

	// Map输入记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MapInputRecords *int64 `json:"MapInputRecords,omitnil" name:"MapInputRecords"`

	// Map输出记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MapOutputRecords *int64 `json:"MapOutputRecords,omitnil" name:"MapOutputRecords"`

	// Reduce输入记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReduceInputRecords *int64 `json:"ReduceInputRecords,omitnil" name:"ReduceInputRecords"`

	// Reduce输出记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReduceOutputRecords *int64 `json:"ReduceOutputRecords,omitnil" name:"ReduceOutputRecords"`

	// HDFS写入字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HDFSBytesWritten *int64 `json:"HDFSBytesWritten,omitnil" name:"HDFSBytesWritten"`

	// HDFS读取字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HDFSBytesRead *int64 `json:"HDFSBytesRead,omitnil" name:"HDFSBytesRead"`
}

type ZoneDetailPriceResult struct {
	// 可用区Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 不同节点的价格详情
	NodeDetailPrice []*NodeDetailPriceResult `json:"NodeDetailPrice,omitnil" name:"NodeDetailPrice"`
}

type ZoneResourceConfiguration struct {
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitnil" name:"VirtualPrivateCloud"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Placement *Placement `json:"Placement,omitnil" name:"Placement"`

	// 所有节点资源的规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllNodeResourceSpec *AllNodeResourceSpec `json:"AllNodeResourceSpec,omitnil" name:"AllNodeResourceSpec"`

	// 如果是单可用区，ZoneTag可以不用填， 如果是双AZ部署，第一个可用区ZoneTag选择master，第二个可用区ZoneTag选择standby，如果是三AZ部署，第一个可用区ZoneTag选择master，第二个可用区ZoneTag选择standby，第三个可用区ZoneTag选择third-party，取值范围：
	//   <li>master</li>
	//   <li>standby</li>
	//   <li>third-party</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneTag *string `json:"ZoneTag,omitnil" name:"ZoneTag"`
}