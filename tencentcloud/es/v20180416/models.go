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

package v20180416

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CosBackup struct {

	// 是否开启cos自动备份
	IsAutoBackup *bool `json:"IsAutoBackup,omitempty" name:"IsAutoBackup"`

	// 自动备份执行时间（精确到小时）, e.g. "22:00"
	BackupTime *string `json:"BackupTime,omitempty" name:"BackupTime"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例版本（支持"5.6.4"、"6.4.3"、"6.8.2"、"7.5.1"）
	EsVersion *string `json:"EsVersion,omitempty" name:"EsVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 访问密码（密码需8到16位，至少包括两项（[a-z,A-Z],[0-9]和[-!@#$%&^*+=_:;,.?]的特殊符号）
	Password *string `json:"Password,omitempty" name:"Password"`

	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 已废弃请使用NodeInfoList
	// 节点数量（2-50个）
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 计费类型<li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li>默认值POSTPAID_BY_HOUR
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 包年包月购买时长（单位由参数TimeUnit决定）
	ChargePeriod *uint64 `json:"ChargePeriod,omitempty" name:"ChargePeriod"`

	// 自动续费标识<li>RENEW_FLAG_AUTO：自动续费</li><li>RENEW_FLAG_MANUAL：不自动续费，用户手动续费</li>ChargeType为PREPAID时需要设置，如不传递该参数，普通用户默认不自动续费，SVIP用户自动续费
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 已废弃请使用NodeInfoList
	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 已废弃请使用NodeInfoList
	// 节点磁盘类型<li>CLOUD_SSD：SSD云硬盘</li><li>CLOUD_PREMIUM：高硬能云硬盘</li>默认值CLOUD_SSD
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 已废弃请使用NodeInfoList
	// 节点磁盘容量（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 计费时长单位（ChargeType为PREPAID时需要设置，默认值为“m”，表示月，当前只支持“m”）
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 已废弃请使用NodeInfoList
	// 是否创建专用主节点<li>true：开启专用主节点</li><li>false：不开启专用主节点</li>默认值false
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitempty" name:"EnableDedicatedMaster"`

	// 已废弃请使用NodeInfoList
	// 专用主节点个数（只支持3个和5个，EnableDedicatedMaster为true时该值必传）
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitempty" name:"MasterNodeNum"`

	// 已废弃请使用NodeInfoList
	// 专用主节点类型（EnableDedicatedMaster为true时必传）<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitempty" name:"MasterNodeType"`

	// 已废弃请使用NodeInfoList
	// 专用主节点磁盘大小（单位GB，非必传，若传递则必须为50，暂不支持自定义）
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitempty" name:"MasterNodeDiskSize"`

	// 集群配置文件中的ClusterName（系统默认配置为实例ID，暂不支持自定义）
	ClusterNameInConf *string `json:"ClusterNameInConf,omitempty" name:"ClusterNameInConf"`

	// 集群部署方式<li>0：单可用区部署</li><li>1：多可用区部署</li>默认为0
	DeployMode *uint64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// 多可用区部署时可用区的详细信息(DeployMode为1时必传)
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitempty" name:"MultiZoneInfo" list`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// 节点信息列表， 用于描述集群各类节点的规格信息如节点类型，节点个数，节点规格，磁盘类型，磁盘大小等
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`

	// 节点标签信息列表
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList" list`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLogsRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志类型，默认值为1
	// <li>1, 主日志</li>
	// <li>2, 搜索慢日志</li>
	// <li>3, 索引慢日志</li>
	// <li>4, GC日志</li>
	LogType *uint64 `json:"LogType,omitempty" name:"LogType"`

	// 搜索词，支持LUCENE语法，如 level:WARN、ip:1.1.1.1、message:test-index等
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 日志开始时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志结束时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页起始值, 默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值为100，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 时间排序方式，默认值为0
	// <li>0, 降序</li>
	// <li>1, 升序</li>
	OrderByType *uint64 `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeInstanceLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的日志条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 日志详细信息列表
		InstanceLogList []*InstanceLog `json:"InstanceLogList,omitempty" name:"InstanceLogList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationsRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 起始时间, e.g. "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间, e.g. "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页起始值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstanceOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 操作记录
		Operations []*Operation `json:"Operations,omitempty" name:"Operations" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群实例所属可用区，不传则默认所有可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 集群实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 集群实例名称列表
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames" list`

	// 分页起始值, 默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，默认值20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段<li>1：实例ID</li><li>2：实例名称</li><li>3：可用区</li><li>4：创建时间</li>若orderKey未传递则按创建时间降序排序
	OrderByKey *uint64 `json:"OrderByKey,omitempty" name:"OrderByKey"`

	// 排序方式<li>0：升序</li><li>1：降序</li>若传递了orderByKey未传递orderByType, 则默认升序
	OrderByType *uint64 `json:"OrderByType,omitempty" name:"OrderByType"`

	// 节点标签信息列表
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList" list`

	// 私有网络vip列表
	IpList []*string `json:"IpList,omitempty" name:"IpList" list`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的实例个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息列表
		InstanceList []*InstanceInfo `json:"InstanceList,omitempty" name:"InstanceList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DictInfo struct {

	// 词典键值
	Key *string `json:"Key,omitempty" name:"Key"`

	// 词典名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 词典大小，单位B
	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

type EsAcl struct {

	// kibana访问黑名单
	BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList" list`

	// kibana访问白名单
	WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList" list`
}

type EsDictionaryInfo struct {

	// 启用词词典列表
	MainDict []*DictInfo `json:"MainDict,omitempty" name:"MainDict" list`

	// 停用词词典列表
	Stopwords []*DictInfo `json:"Stopwords,omitempty" name:"Stopwords" list`

	// QQ分词词典列表
	QQDict []*DictInfo `json:"QQDict,omitempty" name:"QQDict" list`

	// 同义词词典列表
	Synonym []*DictInfo `json:"Synonym,omitempty" name:"Synonym" list`

	// 更新词典类型
	UpdateType *string `json:"UpdateType,omitempty" name:"UpdateType"`
}

type EsPublicAcl struct {

	// 访问黑名单
	BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList" list`

	// 访问白名单
	WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList" list`
}

type InstanceInfo struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 用户ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 用户UIN
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 实例所属VPC的UID
	VpcUid *string `json:"VpcUid,omitempty" name:"VpcUid"`

	// 实例所属子网的UID
	SubnetUid *string `json:"SubnetUid,omitempty" name:"SubnetUid"`

	// 实例状态，0:处理中,1:正常,-1停止,-2:销毁中,-3:已销毁
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例计费模式。取值范围：  PREPAID：表示预付费，即包年包月  POSTPAID_BY_HOUR：表示后付费，即按量计费  CDHPAID：CDH付费，即只对CDH计费，不对CDH上的实例计费。
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 包年包月购买时长,单位:月
	ChargePeriod *uint64 `json:"ChargePeriod,omitempty" name:"ChargePeriod"`

	// 自动续费标识。取值范围：  NOTIFY_AND_AUTO_RENEW：通知过期且自动续费  NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费  DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费  默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点个数
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 节点CPU核数
	CpuNum *uint64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// 节点内存大小，单位GB
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 节点磁盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 节点磁盘大小，单位GB
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// ES域名
	EsDomain *string `json:"EsDomain,omitempty" name:"EsDomain"`

	// ES VIP
	EsVip *string `json:"EsVip,omitempty" name:"EsVip"`

	// ES端口
	EsPort *uint64 `json:"EsPort,omitempty" name:"EsPort"`

	// Kibana访问url
	KibanaUrl *string `json:"KibanaUrl,omitempty" name:"KibanaUrl"`

	// ES版本号
	EsVersion *string `json:"EsVersion,omitempty" name:"EsVersion"`

	// ES配置项
	EsConfig *string `json:"EsConfig,omitempty" name:"EsConfig"`

	// Kibana访问控制配置
	EsAcl *EsAcl `json:"EsAcl,omitempty" name:"EsAcl"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例最后修改操作时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例到期时间
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// 实例类型（实例类型标识，当前只有1,2两种）
	InstanceType *uint64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// Ik分词器配置
	IkConfig *EsDictionaryInfo `json:"IkConfig,omitempty" name:"IkConfig"`

	// 专用主节点配置
	MasterNodeInfo *MasterNodeInfo `json:"MasterNodeInfo,omitempty" name:"MasterNodeInfo"`

	// cos自动备份配置
	CosBackup *CosBackup `json:"CosBackup,omitempty" name:"CosBackup"`

	// 是否允许cos自动备份
	AllowCosBackup *bool `json:"AllowCosBackup,omitempty" name:"AllowCosBackup"`

	// 实例拥有的标签列表
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList" list`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// 是否为冷热集群<li>true: 冷热集群</li><li>false: 非冷热集群</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableHotWarmMode *bool `json:"EnableHotWarmMode,omitempty" name:"EnableHotWarmMode"`

	// 冷节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmNodeType *string `json:"WarmNodeType,omitempty" name:"WarmNodeType"`

	// 冷节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmNodeNum *uint64 `json:"WarmNodeNum,omitempty" name:"WarmNodeNum"`

	// 冷节点CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmCpuNum *uint64 `json:"WarmCpuNum,omitempty" name:"WarmCpuNum"`

	// 冷节点内存内存大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmMemSize *uint64 `json:"WarmMemSize,omitempty" name:"WarmMemSize"`

	// 冷节点磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmDiskType *string `json:"WarmDiskType,omitempty" name:"WarmDiskType"`

	// 冷节点磁盘大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmDiskSize *uint64 `json:"WarmDiskSize,omitempty" name:"WarmDiskSize"`

	// 集群节点信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`

	// Es公网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsPublicUrl *string `json:"EsPublicUrl,omitempty" name:"EsPublicUrl"`

	// 多可用区网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitempty" name:"MultiZoneInfo" list`

	// 部署模式<li>0：单可用区</li><li>1：多可用区</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployMode *uint64 `json:"DeployMode,omitempty" name:"DeployMode"`

	// ES公网访问状态<li>OPEN：开启</li><li>CLOSE：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccess *string `json:"PublicAccess,omitempty" name:"PublicAccess"`

	// ES公网访问控制配置
	EsPublicAcl *EsAcl `json:"EsPublicAcl,omitempty" name:"EsPublicAcl"`

	// Kibana内网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPrivateUrl *string `json:"KibanaPrivateUrl,omitempty" name:"KibanaPrivateUrl"`

	// Kibana公网访问状态<li>OPEN：开启</li><li>CLOSE：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitempty" name:"KibanaPublicAccess"`

	// Kibana内网访问状态<li>OPEN：开启</li><li>CLOSE：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitempty" name:"KibanaPrivateAccess"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityType *uint64 `json:"SecurityType,omitempty" name:"SecurityType"`
}

type InstanceLog struct {

	// 日志时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 日志级别
	Level *string `json:"Level,omitempty" name:"Level"`

	// 集群节点ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 日志内容
	Message *string `json:"Message,omitempty" name:"Message"`
}

type KeyValue struct {

	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LocalDiskInfo struct {

	// 本地盘类型<li>LOCAL_SATA：大数据型</li><li>NVME_SSD：高IO型</li>
	LocalDiskType *string `json:"LocalDiskType,omitempty" name:"LocalDiskType"`

	// 本地盘单盘大小
	LocalDiskSize *uint64 `json:"LocalDiskSize,omitempty" name:"LocalDiskSize"`

	// 本地盘块数
	LocalDiskCount *uint64 `json:"LocalDiskCount,omitempty" name:"LocalDiskCount"`
}

type MasterNodeInfo struct {

	// 是否启用了专用主节点
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitempty" name:"EnableDedicatedMaster"`

	// 专用主节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitempty" name:"MasterNodeType"`

	// 专用主节点个数
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitempty" name:"MasterNodeNum"`

	// 专用主节点CPU核数
	MasterNodeCpuNum *uint64 `json:"MasterNodeCpuNum,omitempty" name:"MasterNodeCpuNum"`

	// 专用主节点内存大小，单位GB
	MasterNodeMemSize *uint64 `json:"MasterNodeMemSize,omitempty" name:"MasterNodeMemSize"`

	// 专用主节点磁盘大小，单位GB
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitempty" name:"MasterNodeDiskSize"`

	// 专用主节点磁盘类型
	MasterNodeDiskType *string `json:"MasterNodeDiskType,omitempty" name:"MasterNodeDiskType"`
}

type NodeInfo struct {

	// 节点数量
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点类型<li>hotData: 热数据节点</li>
	// <li>warmData: 冷数据节点</li>
	// <li>dedicatedMaster: 专用主节点</li>
	// 默认值为hotData
	Type *string `json:"Type,omitempty" name:"Type"`

	// 节点磁盘类型<li>CLOUD_SSD：SSD云硬盘</li><li>CLOUD_PREMIUM：高硬能云硬盘</li>默认值CLOUD_SSD
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 节点磁盘容量（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 节点本地盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDiskInfo *LocalDiskInfo `json:"LocalDiskInfo,omitempty" name:"LocalDiskInfo"`

	// 节点磁盘块数
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// 节点磁盘是否加密 0: 不加密，1: 加密；默认不加密
	DiskEncrypt *uint64 `json:"DiskEncrypt,omitempty" name:"DiskEncrypt"`
}

type Operation struct {

	// 操作唯一id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 操作开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 操作类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 操作详情
	Detail *OperationDetail `json:"Detail,omitempty" name:"Detail"`

	// 操作结果
	Result *string `json:"Result,omitempty" name:"Result"`

	// 流程任务信息
	Tasks []*TaskDetail `json:"Tasks,omitempty" name:"Tasks" list`

	// 操作进度
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`
}

type OperationDetail struct {

	// 实例原始配置信息
	OldInfo []*KeyValue `json:"OldInfo,omitempty" name:"OldInfo" list`

	// 实例更新后配置信息
	NewInfo []*KeyValue `json:"NewInfo,omitempty" name:"NewInfo" list`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否强制重启<li>true：强制重启</li><li>false：不强制重启</li>默认false
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestartInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubTaskDetail struct {

	// 子任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子任务结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 子任务错误信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 子任务类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 子任务状态，0处理中 1成功 -1失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 升级检查失败的索引名
	FailedIndices []*string `json:"FailedIndices,omitempty" name:"FailedIndices" list`

	// 子任务结束时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 子任务等级，1警告 2失败
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

type TagInfo struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TaskDetail struct {

	// 任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务进度
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`

	// 任务完成时间
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 子任务
	SubTasks []*SubTaskDetail `json:"SubTasks,omitempty" name:"SubTasks" list`
}

type UpdateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 已废弃请使用NodeInfoList
	// 节点个数（2-50个）
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 配置项（JSON格式字符串）。当前仅支持以下配置项：<li>action.destructive_requires_name</li><li>indices.fielddata.cache.size</li><li>indices.query.bool.max_clause_count</li>
	EsConfig *string `json:"EsConfig,omitempty" name:"EsConfig"`

	// 默认用户elastic的密码（8到16位，至少包括两项（[a-z,A-Z],[0-9]和[-!@#$%&^*+=_:;,.?]的特殊符号）
	Password *string `json:"Password,omitempty" name:"Password"`

	// 访问控制列表
	EsAcl *EsAcl `json:"EsAcl,omitempty" name:"EsAcl"`

	// 已废弃请使用NodeInfoList
	// 磁盘大小（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 已废弃请使用NodeInfoList
	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 已废弃请使用NodeInfoList
	// 专用主节点个数（只支持3个或5个）
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitempty" name:"MasterNodeNum"`

	// 已废弃请使用NodeInfoList
	// 专用主节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitempty" name:"MasterNodeType"`

	// 已废弃请使用NodeInfoList
	// 专用主节点磁盘大小（单位GB系统默认配置为50GB,暂不支持自定义）
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitempty" name:"MasterNodeDiskSize"`

	// 更新配置时是否强制重启<li>true强制重启</li><li>false不强制重启</li>当前仅更新EsConfig时需要设置，默认值为false
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`

	// COS自动备份信息
	CosBackup *CosBackup `json:"CosBackup,omitempty" name:"CosBackup"`

	// 节点信息列表，可以只传递要更新的节点及其对应的规格信息。支持的操作包括<li>修改一种节点的个数</li><li>修改一种节点的节点规格及磁盘大小</li><li>增加一种节点类型（需要同时指定该节点的类型，个数，规格，磁盘等信息）</li>上述操作一次只能进行一种，且磁盘类型不支持修改
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitempty" name:"NodeInfoList" list`

	// 公网访问状态
	PublicAccess *string `json:"PublicAccess,omitempty" name:"PublicAccess"`

	// 公网访问控制列表
	EsPublicAcl *EsPublicAcl `json:"EsPublicAcl,omitempty" name:"EsPublicAcl"`

	// Kibana公网访问状态
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitempty" name:"KibanaPublicAccess"`

	// Kibana内网访问状态
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitempty" name:"KibanaPrivateAccess"`

	// ES 6.8及以上版本基础版开启或关闭用户认证
	BasicSecurityType *int64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`

	// Kibana内网端口
	KibanaPrivatePort *uint64 `json:"KibanaPrivatePort,omitempty" name:"KibanaPrivatePort"`

	// 0: 蓝绿变更方式扩容，集群不重启 （默认） 1: 磁盘解挂载扩容，集群滚动重启
	ScaleType *int64 `json:"ScaleType,omitempty" name:"ScaleType"`
}

func (r *UpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePluginsRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 需要安装的插件名列表
	InstallPluginList []*string `json:"InstallPluginList,omitempty" name:"InstallPluginList" list`

	// 需要卸载的插件名列表
	RemovePluginList []*string `json:"RemovePluginList,omitempty" name:"RemovePluginList" list`

	// 是否强制重启
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`
}

func (r *UpdatePluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePluginsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePluginsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePluginsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 目标ES版本，支持：”6.4.3“, "6.8.2"，"7.5.1"
	EsVersion *string `json:"EsVersion,omitempty" name:"EsVersion"`

	// 是否只做升级检查，默认值为false
	CheckOnly *bool `json:"CheckOnly,omitempty" name:"CheckOnly"`

	// 目标商业特性版本：<li>oss 开源版</li><li>basic 基础版</li>当前仅在5.6.4升级6.x版本时使用，默认值为basic
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLicenseRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitempty" name:"LicenseType"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitempty" name:"BasicSecurityType"`

	// 是否强制重启<li>true强制重启</li><li>false不强制重启</li> 默认值false
	ForceRestart *bool `json:"ForceRestart,omitempty" name:"ForceRestart"`
}

func (r *UpgradeLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLicenseRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLicenseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLicenseResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZoneDetail struct {

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
