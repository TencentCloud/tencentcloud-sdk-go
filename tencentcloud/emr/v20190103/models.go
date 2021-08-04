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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BootstrapAction struct {

	// 脚本位置，支持cos上的文件，且只支持https协议。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 执行时间。
	// resourceAfter 表示在机器资源申请成功后执行。
	// clusterBefore 表示在集群初始化前执行。
	// clusterAfter 表示在集群初始化后执行。
	WhenRun *string `json:"WhenRun,omitempty" name:"WhenRun"`

	// 脚本参数
	Args []*string `json:"Args,omitempty" name:"Args"`
}

type COSSettings struct {

	// COS SecretId
	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`

	// COS SecrectKey
	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`

	// 日志存储在COS上的路径
	LogOnCosPath *string `json:"LogOnCosPath,omitempty" name:"LogOnCosPath"`
}

type CdbInfo struct {

	// 数据库实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 数据库IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 数据库端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 数据库内存规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// 数据库磁盘规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 服务标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *string `json:"Service,omitempty" name:"Service"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 付费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayType *int64 `json:"PayType,omitempty" name:"PayType"`

	// 过期标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireFlag *bool `json:"ExpireFlag,omitempty" name:"ExpireFlag"`

	// 数据库状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 续费标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAutoRenew *int64 `json:"IsAutoRenew,omitempty" name:"IsAutoRenew"`

	// 数据库字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// ZoneId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// RegionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type ClusterInstancesInfo struct {

	// ID号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ftitle *string `json:"Ftitle,omitempty" name:"Ftitle"`

	// 集群名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 用户APPID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 用户UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 集群VPCID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

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
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 添加时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 已经运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunTime *string `json:"RunTime,omitempty" name:"RunTime"`

	// 集群产品配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *EmrProductConfigOutter `json:"Config,omitempty" name:"Config"`

	// 主节点外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterIp *string `json:"MasterIp,omitempty" name:"MasterIp"`

	// EMR版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmrVersion *string `json:"EmrVersion,omitempty" name:"EmrVersion"`

	// 收费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// 交易版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeVersion *int64 `json:"TradeVersion,omitempty" name:"TradeVersion"`

	// 资源订单ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceOrderId *int64 `json:"ResourceOrderId,omitempty" name:"ResourceOrderId"`

	// 是否计费集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsTradeCluster *int64 `json:"IsTradeCluster,omitempty" name:"IsTradeCluster"`

	// 集群错误状态告警信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmInfo *string `json:"AlarmInfo,omitempty" name:"AlarmInfo"`

	// 是否采用新架构
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWoodpeckerCluster *int64 `json:"IsWoodpeckerCluster,omitempty" name:"IsWoodpeckerCluster"`

	// 元数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaDb *string `json:"MetaDb,omitempty" name:"MetaDb"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Hive元数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HiveMetaDb *string `json:"HiveMetaDb,omitempty" name:"HiveMetaDb"`

	// 集群类型:EMR,CLICKHOUSE,DRUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceClass *string `json:"ServiceClass,omitempty" name:"ServiceClass"`

	// 集群所有节点的别名序列化
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasInfo *string `json:"AliasInfo,omitempty" name:"AliasInfo"`

	// 集群版本Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`

	// 地区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ClusterSetting struct {

	// 付费方式。
	// PREPAID 包年包月。
	// POSTPAID_BY_HOUR 按量计费，默认方式。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 是否为HA集群。
	SupportHA *bool `json:"SupportHA,omitempty" name:"SupportHA"`

	// 集群所使用的安全组，目前仅支持一个。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 实例位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 实例所在VPC。
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`

	// 实例登录配置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例标签。
	TagSpecification []*string `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 元数据库配置。
	MetaDB *MetaDbInfo `json:"MetaDB,omitempty" name:"MetaDB"`

	// 实例硬件配置。
	ResourceSpec *JobFlowResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

	// 是否申请公网IP，默认为false。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`

	// 包年包月配置，只对包年包月集群生效。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 集群置放群组。
	DisasterRecoverGroupIds *string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// 是否使用cbs加密。
	CbsEncryptFlag *bool `json:"CbsEncryptFlag,omitempty" name:"CbsEncryptFlag"`

	// 是否使用远程登录，默认为false。
	RemoteTcpDefaultPort *bool `json:"RemoteTcpDefaultPort,omitempty" name:"RemoteTcpDefaultPort"`
}

type Configuration struct {

	// 配置文件名，支持SPARK、HIVE、HDFS、YARN的部分配置文件自定义。
	Classification *string `json:"Classification,omitempty" name:"Classification"`

	// 配置参数通过KV的形式传入，部分文件支持自定义，可以通过特殊的键"content"传入所有内容。
	Properties *string `json:"Properties,omitempty" name:"Properties"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：
	// <li>1：表示EMR-V1.3.1。</li>
	// <li>2：表示EMR-V2.0.1。</li>
	// <li>4：表示EMR-V2.1.0。</li>
	// <li>7：表示EMR-V3.0.0。</li>
	// <li>9：表示EMR-V2.2.0。</li>
	// <li>11：表示CLICKHOUSE-V1.0.0。</li>
	// <li>13：表示DRUID-V1.0.0。</li>
	// <li>15：表示EMR-V2.2.1。</li>
	// <li>16：表示EMR-V2.3.0。</li>
	// <li>17：表示CLICKHOUSE-V1.1.0。</li>
	// <li>19：表示EMR-V2.4.0。</li>
	// <li>20：表示EMR-V2.5.0。</li>
	// <li>22：表示CLICKHOUSE-V1.2.0。</li>
	// <li>24：表示EMR-TianQiong-V1.0.0。</li>
	// <li>25：表示EMR-V3.1.0。</li>
	// <li>26：表示DORIS-V1.0.0。</li>
	// <li>27：表示KAFKA-V1.0.0。</li>
	// <li>28：表示EMR-V3.2.0。</li>
	// <li>29：表示EMR-V2.5.1。</li>
	// <li>30：表示EMR-V2.6.0。</li>
	ProductId *uint64 `json:"ProductId,omitempty" name:"ProductId"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）对应不同可选组件列表，不同产品版本可选组件列表查询：[组件版本](https://cloud.tencent.com/document/product/589/20279) ；
	// 填写实例值：hive、flink。
	Software []*string `json:"Software,omitempty" name:"Software"`

	// 节点资源的规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// 实例名称。
	// <li>长度限制为6-36个字符。</li>
	// <li>只允许包含中文、字母、数字、-、_。</li>
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 实例登录设置。通过该参数可以设置所购买节点的登录方式密码或者密钥。
	// <li>设置密钥时，密码仅用于组件原生WebUI快捷入口登录。</li>
	// <li>未设置密钥时，密码用于登录所购节点以及组件原生WebUI快捷入口登录。</li>
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 开启COS访问需要设置的参数。
	COSSettings *COSSettings `json:"COSSettings,omitempty" name:"COSSettings"`

	// 实例所属安全组的ID，形如sg-xxxxxxxx。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的SecurityGroupId字段来获取。
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// [引导操作](https://cloud.tencent.com/document/product/589/35656)脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings"`

	// 包年包月实例是否自动续费。取值范围：
	// <li>0：表示不自动续费。</li>
	// <li>1：表示自动续费。</li>
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 客户端Token。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 是否开启集群Master节点公网。取值范围：
	// <li>NEED_MASTER_WAN：表示开启集群Master节点公网。</li>
	// <li>NOT_NEED_MASTER_WAN：表示不开启。</li>默认开启集群Master节点公网。
	NeedMasterWan *string `json:"NeedMasterWan,omitempty" name:"NeedMasterWan"`

	// 是否需要开启外网远程登录，即22号端口。在SgId不为空时，该参数无效。
	RemoteLoginAtCreate *int64 `json:"RemoteLoginAtCreate,omitempty" name:"RemoteLoginAtCreate"`

	// 是否开启安全集群。0表示不开启，非0表示开启。
	CheckSecurity *int64 `json:"CheckSecurity,omitempty" name:"CheckSecurity"`

	// 访问外部文件系统。
	ExtendFsField *string `json:"ExtendFsField,omitempty" name:"ExtendFsField"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的实例。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 分散置放群组ID列表，当前只支持指定一个。
	// 该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/213/15486 ) 的返回值中的SecurityGroupId字段来获取。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// 集群维度CBS加密盘，默认0表示不加密，1表示加密
	CbsEncrypt *uint64 `json:"CbsEncrypt,omitempty" name:"CbsEncrypt"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_NEW_META：表示集群默认创建</li>
	// <li>EMR_EXIT_META：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitempty" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitempty" name:"MetaDBInfo"`

	// 自定义应用角色。
	ApplicationRole *string `json:"ApplicationRole,omitempty" name:"ApplicationRole"`
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
	delete(f, "VPCSettings")
	delete(f, "Software")
	delete(f, "ResourceSpec")
	delete(f, "SupportHA")
	delete(f, "InstanceName")
	delete(f, "PayMode")
	delete(f, "Placement")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "LoginSettings")
	delete(f, "COSSettings")
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CustomMetaInfo struct {

	// 自定义MetaDB的JDBC连接，请以 jdbc:mysql:// 开头
	MetaDataJdbcUrl *string `json:"MetaDataJdbcUrl,omitempty" name:"MetaDataJdbcUrl"`

	// 自定义MetaDB用户名
	MetaDataUser *string `json:"MetaDataUser,omitempty" name:"MetaDataUser"`

	// 自定义MetaDB密码
	MetaDataPass *string `json:"MetaDataPass,omitempty" name:"MetaDataPass"`
}

type DescribeClusterNodesRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

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
	NodeFlag *string `json:"NodeFlag,omitempty" name:"NodeFlag"`

	// 页编号，默认值为0，表示第一页。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页返回数量，默认值为100，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源类型:支持all/host/pod，默认为all
	HardwareResourceType *string `json:"HardwareResourceType,omitempty" name:"HardwareResourceType"`

	// 支持搜索的字段
	SearchFields []*SearchItem `json:"SearchFields,omitempty" name:"SearchFields"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询到的节点总数
		TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

		// 节点详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		NodeList []*NodeHardwareInfo `json:"NodeList,omitempty" name:"NodeList"`

		// 用户所有的标签键列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

		// 资源类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		HardwareResourceTypeList []*string `json:"HardwareResourceTypeList,omitempty" name:"HardwareResourceTypeList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCvmQuotaRequest struct {
	*tchttp.BaseRequest

	// EMR集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 区ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
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

type DescribeCvmQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 后付费配额列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PostPaidQuotaSet []*QuotaEntity `json:"PostPaidQuotaSet,omitempty" name:"PostPaidQuotaSet"`

		// 竞价实例配额列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		SpotPaidQuotaSet []*QuotaEntity `json:"SpotPaidQuotaSet,omitempty" name:"SpotPaidQuotaSet"`

		// eks配额
	// 注意：此字段可能返回 null，表示取不到有效值。
		EksQuotaSet []*PodSaleSpec `json:"EksQuotaSet,omitempty" name:"EksQuotaSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeInstanceRenewNodesRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

type DescribeInstanceRenewNodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询到的节点总数
		TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

		// 节点详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		NodeList []*RenewInstancesInfo `json:"NodeList,omitempty" name:"NodeList"`

		// 用户所有的标签键列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		MetaInfo []*string `json:"MetaInfo,omitempty" name:"MetaInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群筛选策略。取值范围：
	// <li>clusterList：表示查询除了已销毁集群之外的集群列表。</li>
	// <li>monitorManage：表示查询除了已销毁、创建中以及创建失败的集群之外的集群列表。</li>
	// <li>cloudHardwareManage/componentManage：目前这两个取值为预留取值，暂时和monitorManage表示同样的含义。</li>
	DisplayStrategy *string `json:"DisplayStrategy,omitempty" name:"DisplayStrategy"`

	// 按照一个或者多个实例ID查询。实例ID形如: emr-xxxxxxxx 。(此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的 Ids.N 一节)。如果不填写实例ID，返回该APPID下所有实例列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 页编号，默认值为0，表示第一页。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页返回数量，默认值为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 建议必填-1，表示拉取所有项目下的集群。
	// 不填默认值为0，表示拉取默认项目下的集群。
	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/378/4400) 的返回值中的 projectId 字段来获取。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 排序字段。取值范围：
	// <li>clusterId：表示按照实例ID排序。</li>
	// <li>addTime：表示按照实例创建时间排序。</li>
	// <li>status：表示按照实例的状态码排序。</li>
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 按照OrderField升序或者降序进行排序。取值范围：
	// <li>0：表示降序。</li>
	// <li>1：表示升序。</li>默认值为0。
	Asc *int64 `json:"Asc,omitempty" name:"Asc"`
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

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例总数。
		TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

		// EMR实例详细信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterList []*ClusterInstancesInfo `json:"ClusterList,omitempty" name:"ClusterList"`

		// 实例关联的标签键列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeJobFlowRequest struct {
	*tchttp.BaseRequest

	// 流程任务Id，RunJobFlow接口返回的值。
	JobFlowId *int64 `json:"JobFlowId,omitempty" name:"JobFlowId"`
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

type DescribeJobFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流程任务状态，可以为以下值：
	// JobFlowInit，流程任务初始化。
	// JobFlowResourceApplied，资源申请中，通常为JobFlow需要新建集群时的状态。
	// JobFlowResourceReady，执行流程任务的资源就绪。
	// JobFlowStepsRunning，流程任务步骤已提交。
	// JobFlowStepsComplete，流程任务步骤已完成。
	// JobFlowTerminating，流程任务所需资源销毁中。
	// JobFlowFinish，流程任务已完成。
		State *string `json:"State,omitempty" name:"State"`

		// 流程任务步骤结果。
		Details []*JobResult `json:"Details,omitempty" name:"Details"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DiskGroup struct {

	// 磁盘规格。
	Spec *DiskSpec `json:"Spec,omitempty" name:"Spec"`

	// 同类型磁盘数量。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DiskSpec struct {

	// 磁盘类型。
	// LOCAL_BASIC  本地盘。
	// CLOUD_BASIC 云硬盘。
	// LOCAL_SSD 本地SSD。
	// CLOUD_SSD 云SSD。
	// CLOUD_PREMIUM 高效云盘。
	// CLOUD_HSSD 增强型云SSD。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 磁盘大小，单位GB。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type DynamicPodSpec struct {

	// 需求最小cpu核数
	RequestCpu *float64 `json:"RequestCpu,omitempty" name:"RequestCpu"`

	// 需求最大cpu核数
	LimitCpu *float64 `json:"LimitCpu,omitempty" name:"LimitCpu"`

	// 需求最小memory，单位MB
	RequestMemory *float64 `json:"RequestMemory,omitempty" name:"RequestMemory"`

	// 需求最大memory，单位MB
	LimitMemory *float64 `json:"LimitMemory,omitempty" name:"LimitMemory"`
}

type EmrProductConfigOutter struct {

	// 软件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftInfo []*string `json:"SoftInfo,omitempty" name:"SoftInfo"`

	// Master节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterNodeSize *int64 `json:"MasterNodeSize,omitempty" name:"MasterNodeSize"`

	// Core节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreNodeSize *int64 `json:"CoreNodeSize,omitempty" name:"CoreNodeSize"`

	// Task节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNodeSize *int64 `json:"TaskNodeSize,omitempty" name:"TaskNodeSize"`

	// Common节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComNodeSize *int64 `json:"ComNodeSize,omitempty" name:"ComNodeSize"`

	// Master节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterResource *OutterResource `json:"MasterResource,omitempty" name:"MasterResource"`

	// Core节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreResource *OutterResource `json:"CoreResource,omitempty" name:"CoreResource"`

	// Task节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResource *OutterResource `json:"TaskResource,omitempty" name:"TaskResource"`

	// Common节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComResource *OutterResource `json:"ComResource,omitempty" name:"ComResource"`

	// 是否使用COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnCos *bool `json:"OnCos,omitempty" name:"OnCos"`

	// 收费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// Router节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouterNodeSize *int64 `json:"RouterNodeSize,omitempty" name:"RouterNodeSize"`

	// 是否支持HA
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportHA *bool `json:"SupportHA,omitempty" name:"SupportHA"`

	// 是否支持安全模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityOn *bool `json:"SecurityOn,omitempty" name:"SecurityOn"`

	// 安全组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 是否开启Cbs加密
	// 注意：此字段可能返回 null，表示取不到有效值。
	CbsEncrypt *int64 `json:"CbsEncrypt,omitempty" name:"CbsEncrypt"`

	// 自定义应用角色。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationRole *string `json:"ApplicationRole,omitempty" name:"ApplicationRole"`

	// 安全组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

type Execution struct {

	// 任务类型，目前支持以下类型。
	// 1. “MR”，将通过hadoop jar的方式提交。
	// 2. "HIVE"，将通过hive -f的方式提交。
	// 3. "SPARK"，将通过spark-submit的方式提交。
	JobType *string `json:"JobType,omitempty" name:"JobType"`

	// 任务参数，提供除提交指令以外的参数。
	Args []*string `json:"Args,omitempty" name:"Args"`
}

type HostVolumeContext struct {

	// Pod挂载宿主机的目录。资源对宿主机的挂载点，指定的挂载点对应了宿主机的路径，该挂载点在Pod中作为数据存储目录使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumePath *string `json:"VolumePath,omitempty" name:"VolumePath"`
}

type InquirePriceRenewEmrRequest struct {
	*tchttp.BaseRequest

	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费1一个月
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 待续费集群ID列表。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitempty" name:"Currency"`
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

type InquirePriceRenewEmrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// 实例续费的时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 购买实例的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 询价的节点规格。
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 是否开启节点高可用。取值范围：
	// <li>0：表示不开启节点高可用。</li>
	// <li>1：表示开启节点高可用。</li>
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// 部署的组件列表。不同的EMR产品ID（ProductId：具体含义参考入参ProductId字段）需要选择不同的必选组件：
	// <li>ProductId为1的时候，必选组件包括：hadoop-2.7.3、knox-1.2.0、zookeeper-3.4.9</li>
	// <li>ProductId为2的时候，必选组件包括：hadoop-2.7.3、knox-1.2.0、zookeeper-3.4.9</li>
	// <li>ProductId为4的时候，必选组件包括：hadoop-2.8.4、knox-1.2.0、zookeeper-3.4.9</li>
	// <li>ProductId为7的时候，必选组件包括：hadoop-3.1.2、knox-1.2.0、zookeeper-3.4.9</li>
	Software []*string `json:"Software,omitempty" name:"Software"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`

	// hive共享元数据库类型。取值范围：
	// <li>EMR_NEW_META：表示集群默认创建</li>
	// <li>EMR_EXIT_METE：表示集群使用指定EMR-MetaDB。</li>
	// <li>USER_CUSTOM_META：表示集群使用自定义MetaDB。</li>
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// EMR-MetaDB实例
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitempty" name:"UnifyMetaInstanceId"`

	// 自定义MetaDB信息
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitempty" name:"MetaDBInfo"`

	// 产品ID，不同产品ID表示不同的EMR产品版本。取值范围：
	// <li>1：表示EMR-V1.3.1。</li>
	// <li>2：表示EMR-V2.0.1。</li>
	// <li>4：表示EMR-V2.1.0。</li>
	// <li>7：表示EMR-V3.0.0。</li>
	ProductId *uint64 `json:"ProductId,omitempty" name:"ProductId"`
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
	delete(f, "ResourceSpec")
	delete(f, "Currency")
	delete(f, "PayMode")
	delete(f, "SupportHA")
	delete(f, "Software")
	delete(f, "Placement")
	delete(f, "VPCSettings")
	delete(f, "MetaType")
	delete(f, "UnifyMetaInstanceId")
	delete(f, "MetaDBInfo")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// 购买实例的时间单位。取值范围：
	// <li>s：表示秒。</li>
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// 购买实例的时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type InquiryPriceRenewInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例续费的时长。需要结合TimeUnit一起使用。1表示续费1一个月
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 待续费节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr/static/hardware)查询。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 实例计费模式。此处只支持取值为1，表示包年包月。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 是否按量转包年包月。0：否，1：是。
	ModifyPayMode *int64 `json:"ModifyPayMode,omitempty" name:"ModifyPayMode"`
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

type InquiryPriceRenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// 实例续费的时间单位。取值范围：
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// 实例续费的时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type InquiryPriceScaleOutInstanceRequest struct {
	*tchttp.BaseRequest

	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 实例所属的可用区ID，例如100003。该参数可以通过调用 [DescribeZones](https://cloud.tencent.com/document/api/213/15707) 的返回值中的ZoneId字段来获取。
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitempty" name:"RouterCount"`

	// 扩容的Master节点数量。
	MasterCount *uint64 `json:"MasterCount,omitempty" name:"MasterCount"`
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

type InquiryPriceScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *string `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountCost *string `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// 扩容的时间单位。取值范围：
	// <li>s：表示秒。</li>
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
		Unit *string `json:"Unit,omitempty" name:"Unit"`

		// 询价的节点规格。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PriceSpec *PriceResource `json:"PriceSpec,omitempty" name:"PriceSpec"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type InquiryPriceUpdateInstanceRequest struct {
	*tchttp.BaseRequest

	// 变配的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 变配的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 节点变配的目标配置。
	UpdateSpec *UpdateInstanceSettings `json:"UpdateSpec,omitempty" name:"UpdateSpec"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 货币种类。取值范围：
	// <li>CNY：表示人民币。</li>
	Currency *string `json:"Currency,omitempty" name:"Currency"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpdateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 原价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 折扣价，单位为元。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// 变配的时间单位。取值范围：
	// <li>s：表示秒。</li>
	// <li>m：表示月份。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// 变配的时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否自动续费，默认为否。
	RenewFlag *bool `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type JobFlowResource struct {

	// 机器类型描述。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 机器类型描述，可参考CVM的该含义。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 标签KV对。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 磁盘描述列表。
	DiskGroups []*DiskGroup `json:"DiskGroups,omitempty" name:"DiskGroups"`
}

type JobFlowResourceSpec struct {

	// 主节点数量。
	MasterCount *int64 `json:"MasterCount,omitempty" name:"MasterCount"`

	// 主节点配置。
	MasterResourceSpec *JobFlowResource `json:"MasterResourceSpec,omitempty" name:"MasterResourceSpec"`

	// Core节点数量
	CoreCount *int64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Core节点配置。
	CoreResourceSpec *JobFlowResource `json:"CoreResourceSpec,omitempty" name:"CoreResourceSpec"`

	// Task节点数量。
	TaskCount *int64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// Common节点数量。
	CommonCount *int64 `json:"CommonCount,omitempty" name:"CommonCount"`

	// Task节点配置。
	TaskResourceSpec *JobFlowResource `json:"TaskResourceSpec,omitempty" name:"TaskResourceSpec"`

	// Common节点配置。
	CommonResourceSpec *JobFlowResource `json:"CommonResourceSpec,omitempty" name:"CommonResourceSpec"`
}

type JobResult struct {

	// 任务步骤名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务步骤失败时的处理策略，可以为以下值：
	// "CONTINUE"，跳过当前失败步骤，继续后续步骤。
	// “TERMINATE_CLUSTER”，终止当前及后续步骤，并销毁集群。
	// “CANCEL_AND_WAIT”，取消当前步骤并阻塞等待处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionOnFailure *string `json:"ActionOnFailure,omitempty" name:"ActionOnFailure"`

	// 当前步骤的状态，可以为以下值：
	// “JobFlowStepStatusInit”，初始化状态，等待执行。
	// “JobFlowStepStatusRunning”，任务步骤正在执行。
	// “JobFlowStepStatusFailed”，任务步骤执行失败。
	// “JobFlowStepStatusSucceed”，任务步骤执行成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobState *string `json:"JobState,omitempty" name:"JobState"`
}

type LoginSettings struct {

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Public Key
	PublicKeyId *string `json:"PublicKeyId,omitempty" name:"PublicKeyId"`
}

type MetaDbInfo struct {

	// 元数据类型。
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 统一元数据库实例ID。
	UnifyMetaInstanceId *string `json:"UnifyMetaInstanceId,omitempty" name:"UnifyMetaInstanceId"`

	// 自建元数据库信息。
	MetaDBInfo *CustomMetaInfo `json:"MetaDBInfo,omitempty" name:"MetaDBInfo"`
}

type MultiDisk struct {

	// 云盘类型
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_HSSD：表示增强型SSD云硬盘。</li>
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 云盘大小
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 该类型云盘个数
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type MultiDiskMC struct {

	// 该类型云盘个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 云盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`
}

type NewResourceSpec struct {

	// 描述Master节点资源
	MasterResourceSpec *Resource `json:"MasterResourceSpec,omitempty" name:"MasterResourceSpec"`

	// 描述Core节点资源
	CoreResourceSpec *Resource `json:"CoreResourceSpec,omitempty" name:"CoreResourceSpec"`

	// 描述Task节点资源
	TaskResourceSpec *Resource `json:"TaskResourceSpec,omitempty" name:"TaskResourceSpec"`

	// Master节点数量
	MasterCount *int64 `json:"MasterCount,omitempty" name:"MasterCount"`

	// Core节点数量
	CoreCount *int64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Task节点数量
	TaskCount *int64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// 描述Common节点资源
	CommonResourceSpec *Resource `json:"CommonResourceSpec,omitempty" name:"CommonResourceSpec"`

	// Common节点数量
	CommonCount *int64 `json:"CommonCount,omitempty" name:"CommonCount"`
}

type NodeHardwareInfo struct {

	// 用户APPID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 机器实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// master节点绑定外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 节点类型。0:common节点；1:master节点
	// ；2:core节点；3:task节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Flag *int64 `json:"Flag,omitempty" name:"Flag"`

	// 节点规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 节点核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuNum *int64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// 节点内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// 节点内存描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemDesc *string `json:"MemDesc,omitempty" name:"MemDesc"`

	// 节点所在region
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 节点所在Zone
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 释放时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeTime *string `json:"FreeTime,omitempty" name:"FreeTime"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`

	// 节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameTag *string `json:"NameTag,omitempty" name:"NameTag"`

	// 节点部署服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Services *string `json:"Services,omitempty" name:"Services"`

	// 磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// 系统盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// 付费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// 数据库IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdbIp *string `json:"CdbIp,omitempty" name:"CdbIp"`

	// 数据库端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdbPort *int64 `json:"CdbPort,omitempty" name:"CdbPort"`

	// 硬盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwDiskSize *int64 `json:"HwDiskSize,omitempty" name:"HwDiskSize"`

	// 硬盘容量描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwDiskSizeDesc *string `json:"HwDiskSizeDesc,omitempty" name:"HwDiskSizeDesc"`

	// 内存容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwMemSize *int64 `json:"HwMemSize,omitempty" name:"HwMemSize"`

	// 内存容量描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwMemSizeDesc *string `json:"HwMemSizeDesc,omitempty" name:"HwMemSizeDesc"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 节点资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmrResourceId *string `json:"EmrResourceId,omitempty" name:"EmrResourceId"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAutoRenew *int64 `json:"IsAutoRenew,omitempty" name:"IsAutoRenew"`

	// 设备标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceClass *string `json:"DeviceClass,omitempty" name:"DeviceClass"`

	// 支持变配
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mutable *int64 `json:"Mutable,omitempty" name:"Mutable"`

	// 多云盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	MCMultiDisk []*MultiDiskMC `json:"MCMultiDisk,omitempty" name:"MCMultiDisk"`

	// 数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdbNodeInfo *CdbInfo `json:"CdbNodeInfo,omitempty" name:"CdbNodeInfo"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 此节点是否可销毁，1可销毁，0不可销毁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destroyable *int64 `json:"Destroyable,omitempty" name:"Destroyable"`

	// 节点绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否是自动扩缩容节点，0为普通节点，1为自动扩缩容节点。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoFlag *int64 `json:"AutoFlag,omitempty" name:"AutoFlag"`

	// 资源类型, host/pod
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardwareResourceType *string `json:"HardwareResourceType,omitempty" name:"HardwareResourceType"`

	// 是否浮动规格，1是，0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDynamicSpec *int64 `json:"IsDynamicSpec,omitempty" name:"IsDynamicSpec"`

	// 浮动规格值json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	DynamicPodSpec *string `json:"DynamicPodSpec,omitempty" name:"DynamicPodSpec"`

	// 是否支持变更计费类型 1是，0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportModifyPayMode *int64 `json:"SupportModifyPayMode,omitempty" name:"SupportModifyPayMode"`
}

type OutterResource struct {

	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 规格名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// CPU个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type PersistentVolumeContext struct {

	// 磁盘大小，单位为GB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 磁盘类型。CLOUD_PREMIUM;CLOUD_SSD
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 磁盘数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskNum *int64 `json:"DiskNum,omitempty" name:"DiskNum"`
}

type Placement struct {

	// 实例所属项目ID。该参数可以通过调用 DescribeProject 的返回值中的 projectId 字段来获取。填0为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例所属的可用区，例如ap-guangzhou-1。该参数也可以通过调用 DescribeZones 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type PodParameter struct {

	// TKE或EKS集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

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
	Config *string `json:"Config,omitempty" name:"Config"`

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
	Parameter *string `json:"Parameter,omitempty" name:"Parameter"`
}

type PodSaleSpec struct {

	// 可售卖的资源规格，仅为以下值:"TASK","CORE","MASTER","ROUTER"。
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// Cpu核数。
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存数量，单位为GB。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 该规格资源可申请的最大数量。
	Number *uint64 `json:"Number,omitempty" name:"Number"`
}

type PodSpec struct {

	// 外部资源提供者的标识符，例如"cls-a1cd23fa"。
	ResourceProviderIdentifier *string `json:"ResourceProviderIdentifier,omitempty" name:"ResourceProviderIdentifier"`

	// 外部资源提供者类型，例如"tke",当前仅支持"tke"。
	ResourceProviderType *string `json:"ResourceProviderType,omitempty" name:"ResourceProviderType"`

	// 资源的用途，即节点类型，当前仅支持"TASK"。
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// CPU核数。
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位为GB。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 资源对宿主机的挂载点，指定的挂载点对应了宿主机的路径，该挂载点在Pod中作为数据存储目录使用。弃用
	DataVolumes []*string `json:"DataVolumes,omitempty" name:"DataVolumes"`

	// Eks集群-CPU类型，当前支持"intel"和"amd"
	CpuType *string `json:"CpuType,omitempty" name:"CpuType"`

	// Pod节点数据目录挂载信息。
	PodVolumes []*PodVolume `json:"PodVolumes,omitempty" name:"PodVolumes"`

	// 是否浮动规格，1是，0否
	IsDynamicSpec *uint64 `json:"IsDynamicSpec,omitempty" name:"IsDynamicSpec"`

	// 浮动规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DynamicPodSpec *DynamicPodSpec `json:"DynamicPodSpec,omitempty" name:"DynamicPodSpec"`

	// 代表vpc网络唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 代表vpc子网唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type PodState struct {

	// pod的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// pod uuid
	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`

	// pod的状态
	State *string `json:"State,omitempty" name:"State"`

	// pod处于该状态原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// pod所属集群
	OwnerCluster *string `json:"OwnerCluster,omitempty" name:"OwnerCluster"`

	// pod内存大小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
}

type PodVolume struct {

	// 存储类型，可为"pvc"，"hostpath"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeType *string `json:"VolumeType,omitempty" name:"VolumeType"`

	// 当VolumeType为"pvc"时，该字段生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PVCVolume *PersistentVolumeContext `json:"PVCVolume,omitempty" name:"PVCVolume"`

	// 当VolumeType为"hostpath"时，该字段生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostVolume *HostVolumeContext `json:"HostVolume,omitempty" name:"HostVolume"`
}

type PreExecuteFileSettings struct {

	// 脚本在COS上路径，已废弃
	Path *string `json:"Path,omitempty" name:"Path"`

	// 执行脚本参数
	Args []*string `json:"Args,omitempty" name:"Args"`

	// COS的Bucket名称，已废弃
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// COS的Region名称，已废弃
	Region *string `json:"Region,omitempty" name:"Region"`

	// COS的Domain数据，已废弃
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 执行顺序
	RunOrder *int64 `json:"RunOrder,omitempty" name:"RunOrder"`

	// resourceAfter 或 clusterAfter
	WhenRun *string `json:"WhenRun,omitempty" name:"WhenRun"`

	// 脚本文件名，已废弃
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`

	// 脚本的cos地址
	CosFileURI *string `json:"CosFileURI,omitempty" name:"CosFileURI"`

	// cos的SecretId
	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`

	// Cos的SecretKey
	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`

	// cos的appid，已废弃
	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type PriceResource struct {

	// 需要的规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *uint64 `json:"StorageType,omitempty" name:"StorageType"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// 核心数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 云盘列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiDisks []*MultiDisk `json:"MultiDisks,omitempty" name:"MultiDisks"`

	// 磁盘数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCnt *int64 `json:"DiskCnt,omitempty" name:"DiskCnt"`

	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 磁盘数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskNum *int64 `json:"DiskNum,omitempty" name:"DiskNum"`

	// 本地盘的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDiskNum *int64 `json:"LocalDiskNum,omitempty" name:"LocalDiskNum"`
}

type QuotaEntity struct {

	// 已使用配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedQuota *int64 `json:"UsedQuota,omitempty" name:"UsedQuota"`

	// 剩余配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainingQuota *int64 `json:"RemainingQuota,omitempty" name:"RemainingQuota"`

	// 总配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalQuota *int64 `json:"TotalQuota,omitempty" name:"TotalQuota"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type RenewInstancesInfo struct {

	// 节点资源ID
	EmrResourceId *string `json:"EmrResourceId,omitempty" name:"EmrResourceId"`

	// 节点类型。0:common节点；1:master节点
	// ；2:core节点；3:task节点
	Flag *int64 `json:"Flag,omitempty" name:"Flag"`

	// 内网IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 节点内存描述
	MemDesc *string `json:"MemDesc,omitempty" name:"MemDesc"`

	// 节点核数
	CpuNum *int64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// 硬盘大小
	DiskSize *string `json:"DiskSize,omitempty" name:"DiskSize"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 节点规格
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 磁盘类型
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`
}

type Resource struct {

	// 节点规格描述，如CVM.SA2。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 存储类型
	// 取值范围：
	// <li>4：表示云SSD。</li>
	// <li>5：表示高效云盘。</li>
	// <li>6：表示增强型SSD云硬盘。</li>
	// <li>11：表示吞吐型云硬盘。</li>
	// <li>12：表示极速型SSD云硬盘。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// 磁盘类型
	// 取值范围：
	// <li>CLOUD_SSD：表示云SSD。</li>
	// <li>CLOUD_PREMIUM：表示高效云盘。</li>
	// <li>CLOUD_BASIC：表示云硬盘。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 内存容量,单位为M
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 系统盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// 云盘列表，当数据盘为一块云盘时，直接使用DiskType和DiskSize参数，超出部分使用MultiDisks
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiDisks []*MultiDisk `json:"MultiDisks,omitempty" name:"MultiDisks"`

	// 需要绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 规格类型，如S2.MEDIUM8
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 本地盘数量，该字段已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDiskNum *uint64 `json:"LocalDiskNum,omitempty" name:"LocalDiskNum"`

	// 本地盘数量，如2
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskNum *uint64 `json:"DiskNum,omitempty" name:"DiskNum"`
}

type RunJobFlowRequest struct {
	*tchttp.BaseRequest

	// 作业名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否新创建集群。
	// true，新创建集群，则使用Instance中的参数进行集群创建。
	// false，使用已有集群，则通过InstanceId传入。
	CreateCluster *bool `json:"CreateCluster,omitempty" name:"CreateCluster"`

	// 作业流程执行步骤。
	Steps []*Step `json:"Steps,omitempty" name:"Steps"`

	// 作业流程正常完成时，集群的处理方式，可选择:
	// Terminate 销毁集群。
	// Reserve 保留集群。
	InstancePolicy *string `json:"InstancePolicy,omitempty" name:"InstancePolicy"`

	// 只有CreateCluster为true时生效，目前只支持EMR版本，例如EMR-2.2.0，不支持ClickHouse和Druid版本。
	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`

	// 只在CreateCluster为true时生效。
	// true 表示安装kerberos，false表示不安装kerberos。
	SecurityClusterFlag *bool `json:"SecurityClusterFlag,omitempty" name:"SecurityClusterFlag"`

	// 只在CreateCluster为true时生效。
	// 新建集群时，要安装的软件列表。
	Software []*string `json:"Software,omitempty" name:"Software"`

	// 引导脚本。
	BootstrapActions []*BootstrapAction `json:"BootstrapActions,omitempty" name:"BootstrapActions"`

	// 指定配置创建集群。
	Configurations []*Configuration `json:"Configurations,omitempty" name:"Configurations"`

	// 作业日志保存地址。
	LogUri *string `json:"LogUri,omitempty" name:"LogUri"`

	// 只在CreateCluster为false时生效。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 自定义应用角色，大数据应用访问外部服务时使用的角色，默认为"EME_QCSRole"。
	ApplicationRole *string `json:"ApplicationRole,omitempty" name:"ApplicationRole"`

	// 重入标签，用来可重入检查，防止在一段时间内，创建相同的流程作业。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 只在CreateCluster为true时生效，使用该配置创建集群。
	Instance *ClusterSetting `json:"Instance,omitempty" name:"Instance"`
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

type RunJobFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 作业流程ID。
		JobFlowId *int64 `json:"JobFlowId,omitempty" name:"JobFlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest

	// 扩容的时间单位。取值范围：
	// <li>s：表示秒。PayMode取值为0时，TimeUnit只能取值为s。</li>
	// <li>m：表示月份。PayMode取值为1时，TimeUnit只能取值为m。</li>
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 扩容的时长。结合TimeUnit一起使用。
	// <li>TimeUnit为s时，该参数只能填写3600，表示按量计费实例。</li>
	// <li>TimeUnit为m时，该参数填写的数字表示包年包月实例的购买时长，如1表示购买一个月</li>
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例计费模式。取值范围：
	// <li>0：表示按量计费。</li>
	// <li>1：表示包年包月。</li>
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 客户端Token。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 引导操作脚本设置。
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings"`

	// 扩容的Task节点数量。
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// 扩容的Core节点数量。
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// 扩容时不需要安装的进程。
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitempty" name:"UnNecessaryNodeList"`

	// 扩容的Router节点数量。
	RouterCount *uint64 `json:"RouterCount,omitempty" name:"RouterCount"`

	// 部署的服务。
	// <li>SoftDeployInfo和ServiceNodeInfo是同组参数，和UnNecessaryNodeList参数互斥。</li>
	// <li>建议使用SoftDeployInfo和ServiceNodeInfo组合。</li>
	SoftDeployInfo []*uint64 `json:"SoftDeployInfo,omitempty" name:"SoftDeployInfo"`

	// 启动的进程。
	ServiceNodeInfo []*uint64 `json:"ServiceNodeInfo,omitempty" name:"ServiceNodeInfo"`

	// 分散置放群组ID列表，当前仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitempty" name:"DisasterRecoverGroupIds"`

	// 扩容节点绑定标签列表。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 扩容所选资源类型，可选范围为"host","pod"，host为普通的CVM资源，Pod为TKE集群或EKS集群提供的资源
	HardwareResourceType *string `json:"HardwareResourceType,omitempty" name:"HardwareResourceType"`

	// 使用Pod资源扩容时，指定的Pod规格以及来源等信息
	PodSpec *PodSpec `json:"PodSpec,omitempty" name:"PodSpec"`

	// 使用clickhouse集群扩容时，选择的机器分组名称
	ClickHouseClusterName *string `json:"ClickHouseClusterName,omitempty" name:"ClickHouseClusterName"`

	// 使用clickhouse集群扩容时，选择的机器分组类型。new为新增，old为选择旧分组
	ClickHouseClusterType *string `json:"ClickHouseClusterType,omitempty" name:"ClickHouseClusterType"`

	// 规则扩容指定 yarn node label
	YarnNodeLabel *string `json:"YarnNodeLabel,omitempty" name:"YarnNodeLabel"`

	// POD自定义权限和自定义参数
	PodParameter *PodParameter `json:"PodParameter,omitempty" name:"PodParameter"`

	// 扩容的Master节点的数量。
	// 使用clickhouse集群扩容时，该参数不生效。
	// 使用kafka集群扩容时，该参数不生效。
	// 当HardwareResourceType=POD时，该参数不生效。
	MasterCount *uint64 `json:"MasterCount,omitempty" name:"MasterCount"`

	// 扩容后是否启动服务，true：启动，false：不启动
	StartServiceAfterScaleOut *string `json:"StartServiceAfterScaleOut,omitempty" name:"StartServiceAfterScaleOut"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID。
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

		// 客户端Token。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

		// 扩容流程ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

		// 大订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BillId *string `json:"BillId,omitempty" name:"BillId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type SearchItem struct {

	// 支持搜索的类型
	SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

	// 支持搜索的值
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

type Step struct {

	// 执行步骤名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 执行动作。
	ExecutionStep *Execution `json:"ExecutionStep,omitempty" name:"ExecutionStep"`

	// 执行失败策略。
	// 1. TERMINATE_CLUSTER 执行失败时退出并销毁集群。
	// 2. CANCEL_AND_WAIT 执行失败时阻塞等待。
	// 3. CONTINUE 执行失败时跳过并执行后续步骤。
	ActionOnFailure *string `json:"ActionOnFailure,omitempty" name:"ActionOnFailure"`

	// 指定执行Step时的用户名，非必须，默认为hadoop。
	User *string `json:"User,omitempty" name:"User"`
}

type SyncPodStateRequest struct {
	*tchttp.BaseRequest

	// EmrService中pod状态信息
	Message *PodState `json:"Message,omitempty" name:"Message"`
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

type SyncPodStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TerminateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 销毁节点ID。该参数为预留参数，用户无需配置。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
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

type TerminateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type TerminateTasksRequest struct {
	*tchttp.BaseRequest

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 待销毁节点的资源ID列表。资源ID形如：emr-vm-xxxxxxxx。有效的资源ID可通过登录[控制台](https://console.cloud.tencent.com/emr/static/hardware)查询。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
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

type TerminateTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type UpdateInstanceSettings struct {

	// 内存容量，单位为G
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// CPU核数
	CPUCores *uint64 `json:"CPUCores,omitempty" name:"CPUCores"`

	// 机器资源ID（EMR测资源标识）
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 变配机器规格
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type VPCSettings struct {

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
