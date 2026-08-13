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

package v20180420

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AssetsInfo struct {
	// <p>创建时间</p>
	AddTime *uint64 `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// <p>资产 ID</p>
	Aid *uint64 `json:"Aid,omitnil,omitempty" name:"Aid"`

	// <p>数据资产 IP</p>
	AssetsIp *string `json:"AssetsIp,omitnil,omitempty" name:"AssetsIp"`

	// <p>数据资产名称</p>
	AssetsName *string `json:"AssetsName,omitnil,omitempty" name:"AssetsName"`

	// <p>数据资产端口</p>
	AssetsPort *uint64 `json:"AssetsPort,omitnil,omitempty" name:"AssetsPort"`

	// <p>数据资产类型</p>
	AssetsType *string `json:"AssetsType,omitnil,omitempty" name:"AssetsType"`

	// <p>资产版本</p>
	AssetsVersion *string `json:"AssetsVersion,omitnil,omitempty" name:"AssetsVersion"`

	// <p>是否动态</p>
	AssetsAddType *uint64 `json:"AssetsAddType,omitnil,omitempty" name:"AssetsAddType"`

	// <p>是否删除</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>最后一次修改时间</p>
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>资产的vpc</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>地域</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>审计权限</p>
	Permission *int64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>用来区分自建资产是已通过cvm还是添加ip的方式</p>
	AddType *uint64 `json:"AddType,omitnil,omitempty" name:"AddType"`

	// <p>子网Id</p>
	AssetSubnetId *string `json:"AssetSubnetId,omitnil,omitempty" name:"AssetSubnetId"`

	// <p>是否已上传数据库私钥（0 否，1 是）</p>
	UploadPem *int64 `json:"UploadPem,omitnil,omitempty" name:"UploadPem"`

	// <p>资产状态栏 0:正常 1:已删除（目前仅对tencentDB有效）</p>
	AliveStatus *int64 `json:"AliveStatus,omitnil,omitempty" name:"AliveStatus"`

	// <p>开启agent(0:关闭;1:开启)</p>
	AgentOn *uint64 `json:"AgentOn,omitnil,omitempty" name:"AgentOn"`

	// <p>开启agent(0:关闭;1:开启)</p>
	CasbOn *uint64 `json:"CasbOn,omitnil,omitempty" name:"CasbOn"`

	// <p>只读组/集群ID</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>PROXY_OFF: 未开启Casb代理;PROXY_ERROR:Casb代理接口返回异常;PROXY_BOUND:已绑定;PROXY_UNBOUND:未绑定;UNPAID:未购买;UNSUPPORTED:类型不支持;METADATA_NOT_FOUND:元数据不存在;QUOTA_EXCEEDED:Casb额度不足</p>
	Available *string `json:"Available,omitnil,omitempty" name:"Available"`

	// <p>cdbOn</p>
	CdbOn *uint64 `json:"CdbOn,omitnil,omitempty" name:"CdbOn"`

	// <p>平台位数 32位 64位</p>
	DbPlatform *string `json:"DbPlatform,omitnil,omitempty" name:"DbPlatform"`

	// <p>编码</p>
	DbCharset *string `json:"DbCharset,omitnil,omitempty" name:"DbCharset"`

	// <p>操作系统</p>
	OsPolicy *string `json:"OsPolicy,omitnil,omitempty" name:"OsPolicy"`

	// <p>是否开启双向审计</p>
	BidirectionOn *int64 `json:"BidirectionOn,omitnil,omitempty" name:"BidirectionOn"`

	// <p>最大返回行数</p>
	BidirectionMaxLine *int64 `json:"BidirectionMaxLine,omitnil,omitempty" name:"BidirectionMaxLine"`

	// <p>最大返回大小</p>
	BidirectionMaxStorage *int64 `json:"BidirectionMaxStorage,omitnil,omitempty" name:"BidirectionMaxStorage"`

	// <p>是否允许开通双向审计(1.允许；0不允许)</p>
	BidirectionAllow *int64 `json:"BidirectionAllow,omitnil,omitempty" name:"BidirectionAllow"`

	// <p>启双向审计的日志投递(1.开启;0.关闭)</p>
	BidirectionDelivery *uint64 `json:"BidirectionDelivery,omitnil,omitempty" name:"BidirectionDelivery"`

	// <p>只读状态</p>
	RoStatus *string `json:"RoStatus,omitnil,omitempty" name:"RoStatus"`

	// <p>当前资产是否开启了对当前Agent的采集策略</p>
	AgentBound *bool `json:"AgentBound,omitnil,omitempty" name:"AgentBound"`

	// <p>错误信息</p>
	CdbErrorMsg *string `json:"CdbErrorMsg,omitnil,omitempty" name:"CdbErrorMsg"`

	// <p>资产 DSGC 绑定信息</p>
	DsgcBindingInfo *DsgcBindingInfo `json:"DsgcBindingInfo,omitnil,omitempty" name:"DsgcBindingInfo"`

	// <p>绑定的规则Ids</p>
	BindingRules []*IdWithName `json:"BindingRules,omitnil,omitempty" name:"BindingRules"`

	// <p>绑定的模型Ids</p>
	BindingModels []*IdWithName `json:"BindingModels,omitnil,omitempty" name:"BindingModels"`

	// <p>所属组名</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>资产组Id</p>
	AssetGroupId *uint64 `json:"AssetGroupId,omitnil,omitempty" name:"AssetGroupId"`

	// <p>是否是新云原生审计流程</p>
	IsNewCloudAudit *bool `json:"IsNewCloudAudit,omitnil,omitempty" name:"IsNewCloudAudit"`

	// <p>1</p><p>取值范围：[0, 1]</p>
	TrafficMirrorOn *int64 `json:"TrafficMirrorOn,omitnil,omitempty" name:"TrafficMirrorOn"`

	// <p>流量镜像审计范围</p><p>枚举值：</p><ul><li>ALL： 全地域</li><li>REGION： 资产所在地域</li><li>VPC： 资产所在VPC</li></ul><p>默认值：REGION</p>
	AuditScope *string `json:"AuditScope,omitnil,omitempty" name:"AuditScope"`

	// <p>实例集群ID</p>
	InstanceGroupId *string `json:"InstanceGroupId,omitnil,omitempty" name:"InstanceGroupId"`

	// <p>该资产所在的资产组</p>
	AssetGroups []*IdWithName `json:"AssetGroups,omitnil,omitempty" name:"AssetGroups"`
}

type CdsAuditInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 续费标识
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 付费模式（数据安全审计只支持预付费：1）
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例状态： 0，未生效；1：正常运行； 2：被隔离； 3，已过期
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例被隔离时间，格式：yyyy-mm-dd HH:ii:ss
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitnil,omitempty" name:"IsolatedTimestamp"`

	// 实例创建时间，格式： yyyy-mm-dd HH:ii:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例过期时间，格式：yyyy-mm-dd HH:ii:ss
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 实例私网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 实例类型（版本）
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例域名
	Pdomain *string `json:"Pdomain,omitnil,omitempty" name:"Pdomain"`
}

// Predefined struct for user
type CreateReportPdfRequestParams struct {
	// <p>报表 Id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type CreateReportPdfRequest struct {
	*tchttp.BaseRequest
	
	// <p>报表 Id</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *CreateReportPdfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReportPdfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReportPdfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReportPdfResponseParams struct {
	// <p>下载地址</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReportPdfResponse struct {
	*tchttp.BaseResponse
	Response *CreateReportPdfResponseParams `json:"Response"`
}

func (r *CreateReportPdfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReportPdfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTimerReportRequestParams struct {
	// 任务名称 不变更为""
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 执行日期 重复周期为天：无意义周：星期几1-7月每月几号 1-31
	CntTime *int64 `json:"CntTime,omitnil,omitempty" name:"CntTime"`

	// 重复周期
	CntCycle *int64 `json:"CntCycle,omitnil,omitempty" name:"CntCycle"`

	// 发送目标
	Receivers *string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// 时间范围 1:24小时 7:近一周 30:近30天 90:近90天 180:近180天 不变更为0
	CntDay *int64 `json:"CntDay,omitnil,omitempty" name:"CntDay"`

	// 执行时间 格式15:04 到分钟
	CntDate *string `json:"CntDate,omitnil,omitempty" name:"CntDate"`

	// 报告说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 模版Id
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 报表类型
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 关联的资产数组
	AssetsId []*int64 `json:"AssetsId,omitnil,omitempty" name:"AssetsId"`

	// 报表通知 1关闭 2开启 不变更为0
	Notification *int64 `json:"Notification,omitnil,omitempty" name:"Notification"`

	// 任务起停 1:关闭 2:开启 单次报表默认为2
	MissionStart *int64 `json:"MissionStart,omitnil,omitempty" name:"MissionStart"`
}

type CreateTimerReportRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称 不变更为""
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 执行日期 重复周期为天：无意义周：星期几1-7月每月几号 1-31
	CntTime *int64 `json:"CntTime,omitnil,omitempty" name:"CntTime"`

	// 重复周期
	CntCycle *int64 `json:"CntCycle,omitnil,omitempty" name:"CntCycle"`

	// 发送目标
	Receivers *string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// 时间范围 1:24小时 7:近一周 30:近30天 90:近90天 180:近180天 不变更为0
	CntDay *int64 `json:"CntDay,omitnil,omitempty" name:"CntDay"`

	// 执行时间 格式15:04 到分钟
	CntDate *string `json:"CntDate,omitnil,omitempty" name:"CntDate"`

	// 报告说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 模版Id
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 报表类型
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 关联的资产数组
	AssetsId []*int64 `json:"AssetsId,omitnil,omitempty" name:"AssetsId"`

	// 报表通知 1关闭 2开启 不变更为0
	Notification *int64 `json:"Notification,omitnil,omitempty" name:"Notification"`

	// 任务起停 1:关闭 2:开启 单次报表默认为2
	MissionStart *int64 `json:"MissionStart,omitnil,omitempty" name:"MissionStart"`
}

func (r *CreateTimerReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTimerReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TplName")
	delete(f, "CntTime")
	delete(f, "CntCycle")
	delete(f, "Receivers")
	delete(f, "CntDay")
	delete(f, "CntDate")
	delete(f, "Remark")
	delete(f, "TemplateId")
	delete(f, "ReportType")
	delete(f, "AssetsId")
	delete(f, "Notification")
	delete(f, "MissionStart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTimerReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTimerReportResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTimerReportResponse struct {
	*tchttp.BaseResponse
	Response *CreateTimerReportResponseParams `json:"Response"`
}

func (r *CreateTimerReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTimerReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DbauditTypesInfo struct {
	// 规格描述
	InstanceVersionName *string `json:"InstanceVersionName,omitnil,omitempty" name:"InstanceVersionName"`

	// 规格名称
	InstanceVersionKey *string `json:"InstanceVersionKey,omitnil,omitempty" name:"InstanceVersionKey"`

	// 最大吞吐量
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 最大实例数
	MaxInstances *uint64 `json:"MaxInstances,omitnil,omitempty" name:"MaxInstances"`

	// 入库速率（每小时）
	InsertSpeed *uint64 `json:"InsertSpeed,omitnil,omitempty" name:"InsertSpeed"`

	// 最大在线存储量，单位：条
	OnlineStorageCapacity *uint64 `json:"OnlineStorageCapacity,omitnil,omitempty" name:"OnlineStorageCapacity"`

	// 最大归档存储量，单位：条
	ArchivingStorageCapacity *uint64 `json:"ArchivingStorageCapacity,omitnil,omitempty" name:"ArchivingStorageCapacity"`
}

// Predefined struct for user
type DescribeAssetsListRequestParams struct {
	// <p>限制数目</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>实例Id/实例名称/资产名称</p>
	SearchValues []*NameValueString `json:"SearchValues,omitnil,omitempty" name:"SearchValues"`

	// <p>数据资产类型</p>
	AssetsType *string `json:"AssetsType,omitnil,omitempty" name:"AssetsType"`

	// <p>查询的资产类型（1:cdb、2:cvm、3:others）</p>
	AssetsAddType *int64 `json:"AssetsAddType,omitnil,omitempty" name:"AssetsAddType"`

	// <p>地域</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>审计权限</p>
	Permission *int64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>状态</p>
	AliveStatus *int64 `json:"AliveStatus,omitnil,omitempty" name:"AliveStatus"`

	// <p>1.代理开启 0.代理关闭 -1.全查</p>
	CasbOn *int64 `json:"CasbOn,omitnil,omitempty" name:"CasbOn"`

	// <p>1.Agent开启 0.Agent关闭 -1.全查</p>
	AgentOn *int64 `json:"AgentOn,omitnil,omitempty" name:"AgentOn"`

	// <p>0.关闭，1.开启，2.关闭中，3.开启中 -1.全查</p>
	CdbOn *int64 `json:"CdbOn,omitnil,omitempty" name:"CdbOn"`

	// <p>扩展分类，如sensitive，指定查询支持敏感数据识别的资产</p>
	ExtendCategory *string `json:"ExtendCategory,omitnil,omitempty" name:"ExtendCategory"`

	// <p>资产组Id（Id=0 暂未分组；id&gt;0 组Id）</p>
	GroupIds []*uint64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>资产Id</p>
	Aids []*uint64 `json:"Aids,omitnil,omitempty" name:"Aids"`

	// <p>查询绑定状态（1:查询规则绑定数量；2:查询模型绑定数量）</p>
	BindingState *uint64 `json:"BindingState,omitnil,omitempty" name:"BindingState"`

	// <p>网卡是否开启流量审计</p><p>取值范围：[-1, 1]</p>
	TrafficMirrorOn *int64 `json:"TrafficMirrorOn,omitnil,omitempty" name:"TrafficMirrorOn"`
}

type DescribeAssetsListRequest struct {
	*tchttp.BaseRequest
	
	// <p>限制数目</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>实例Id/实例名称/资产名称</p>
	SearchValues []*NameValueString `json:"SearchValues,omitnil,omitempty" name:"SearchValues"`

	// <p>数据资产类型</p>
	AssetsType *string `json:"AssetsType,omitnil,omitempty" name:"AssetsType"`

	// <p>查询的资产类型（1:cdb、2:cvm、3:others）</p>
	AssetsAddType *int64 `json:"AssetsAddType,omitnil,omitempty" name:"AssetsAddType"`

	// <p>地域</p>
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// <p>审计权限</p>
	Permission *int64 `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>状态</p>
	AliveStatus *int64 `json:"AliveStatus,omitnil,omitempty" name:"AliveStatus"`

	// <p>1.代理开启 0.代理关闭 -1.全查</p>
	CasbOn *int64 `json:"CasbOn,omitnil,omitempty" name:"CasbOn"`

	// <p>1.Agent开启 0.Agent关闭 -1.全查</p>
	AgentOn *int64 `json:"AgentOn,omitnil,omitempty" name:"AgentOn"`

	// <p>0.关闭，1.开启，2.关闭中，3.开启中 -1.全查</p>
	CdbOn *int64 `json:"CdbOn,omitnil,omitempty" name:"CdbOn"`

	// <p>扩展分类，如sensitive，指定查询支持敏感数据识别的资产</p>
	ExtendCategory *string `json:"ExtendCategory,omitnil,omitempty" name:"ExtendCategory"`

	// <p>资产组Id（Id=0 暂未分组；id&gt;0 组Id）</p>
	GroupIds []*uint64 `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// <p>资产Id</p>
	Aids []*uint64 `json:"Aids,omitnil,omitempty" name:"Aids"`

	// <p>查询绑定状态（1:查询规则绑定数量；2:查询模型绑定数量）</p>
	BindingState *uint64 `json:"BindingState,omitnil,omitempty" name:"BindingState"`

	// <p>网卡是否开启流量审计</p><p>取值范围：[-1, 1]</p>
	TrafficMirrorOn *int64 `json:"TrafficMirrorOn,omitnil,omitempty" name:"TrafficMirrorOn"`
}

func (r *DescribeAssetsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SearchValues")
	delete(f, "AssetsType")
	delete(f, "AssetsAddType")
	delete(f, "RegionId")
	delete(f, "Permission")
	delete(f, "AliveStatus")
	delete(f, "CasbOn")
	delete(f, "AgentOn")
	delete(f, "CdbOn")
	delete(f, "ExtendCategory")
	delete(f, "GroupIds")
	delete(f, "Aids")
	delete(f, "BindingState")
	delete(f, "TrafficMirrorOn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetsListResponseParams struct {
	// <p>总数目</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>数据列表</p>
	List []*AssetsInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetsListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetsListResponseParams `json:"Response"`
}

func (r *DescribeAssetsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditInstanceTypeRequestParams struct {

}

type DescribeDbauditInstanceTypeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDbauditInstanceTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditInstanceTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbauditInstanceTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditInstanceTypeResponseParams struct {
	// 数据安全审计产品规格信息列表
	DbauditTypesSet []*DbauditTypesInfo `json:"DbauditTypesSet,omitnil,omitempty" name:"DbauditTypesSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDbauditInstanceTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbauditInstanceTypeResponseParams `json:"Response"`
}

func (r *DescribeDbauditInstanceTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditInstanceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditInstancesRequestParams struct {
	// 查询条件地域
	SearchRegion *string `json:"SearchRegion,omitnil,omitempty" name:"SearchRegion"`

	// 限制数目，默认10， 最大50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认1
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDbauditInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件地域
	SearchRegion *string `json:"SearchRegion,omitnil,omitempty" name:"SearchRegion"`

	// 限制数目，默认10， 最大50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认1
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDbauditInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbauditInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditInstancesResponseParams struct {
	// 总实例数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据安全审计实例信息列表
	CdsAuditInstanceSet []*CdsAuditInstance `json:"CdsAuditInstanceSet,omitnil,omitempty" name:"CdsAuditInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDbauditInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbauditInstancesResponseParams `json:"Response"`
}

func (r *DescribeDbauditInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditUsedRegionsRequestParams struct {

}

type DescribeDbauditUsedRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDbauditUsedRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditUsedRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbauditUsedRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbauditUsedRegionsResponseParams struct {
	// 可售卖地域信息列表
	RegionSet []*RegionInfo `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDbauditUsedRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbauditUsedRegionsResponseParams `json:"Response"`
}

func (r *DescribeDbauditUsedRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbauditUsedRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportListRequestParams struct {
	// 限制数目
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 报告名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 报告类型
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 报告状态
	ReportStatus *int64 `json:"ReportStatus,omitnil,omitempty" name:"ReportStatus"`

	// 报表模版id
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 需要排序的字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 排序顺序 asc desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 时间范围 1:24小时 7:近一周 30:近30天 90:近90天 180:近180天 不变更为0
	CntDay *int64 `json:"CntDay,omitnil,omitempty" name:"CntDay"`
}

type DescribeReportListRequest struct {
	*tchttp.BaseRequest
	
	// 限制数目
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 报告名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 报告类型
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 报告状态
	ReportStatus *int64 `json:"ReportStatus,omitnil,omitempty" name:"ReportStatus"`

	// 报表模版id
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 需要排序的字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 排序顺序 asc desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 时间范围 1:24小时 7:近一周 30:近30天 90:近90天 180:近180天 不变更为0
	CntDay *int64 `json:"CntDay,omitnil,omitempty" name:"CntDay"`
}

func (r *DescribeReportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Name")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ReportType")
	delete(f, "ReportStatus")
	delete(f, "TemplateId")
	delete(f, "Field")
	delete(f, "Sort")
	delete(f, "CntDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportListResponseParams struct {
	// 总数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据列表
	List []*Reports `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReportListResponseParams `json:"Response"`
}

func (r *DescribeReportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportMissionListRequestParams struct {
	// 报表名 可模糊查询
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 报表类型 1:单次报表 2:周期报表 0全查
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 报表模板 1:综合分析报告 2:等保合规报告 0全查
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 任务状态0全查 1:生成中 2:待生成 3:已生成 4:生成失败 5:已暂停
	MissionStatus *int64 `json:"MissionStatus,omitnil,omitempty" name:"MissionStatus"`

	// 排序字段 支持“NextStartTime” 与 “MissionStatus”
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// ‘desc' | 'asc'
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeReportMissionListRequest struct {
	*tchttp.BaseRequest
	
	// 报表名 可模糊查询
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 报表类型 1:单次报表 2:周期报表 0全查
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 报表模板 1:综合分析报告 2:等保合规报告 0全查
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 任务状态0全查 1:生成中 2:待生成 3:已生成 4:生成失败 5:已暂停
	MissionStatus *int64 `json:"MissionStatus,omitnil,omitempty" name:"MissionStatus"`

	// 排序字段 支持“NextStartTime” 与 “MissionStatus”
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// ‘desc' | 'asc'
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeReportMissionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportMissionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TplName")
	delete(f, "ReportType")
	delete(f, "TemplateId")
	delete(f, "MissionStatus")
	delete(f, "Field")
	delete(f, "Sort")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReportMissionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportMissionListResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 报表列表
	List []*ReportMission `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReportMissionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReportMissionListResponseParams `json:"Response"`
}

func (r *DescribeReportMissionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportMissionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DsgcBindingInfo struct {
	// dspa实例id
	DspaId *string `json:"DspaId,omitnil,omitempty" name:"DspaId"`

	// dspa绑定模板/合规组 id ComplianceGroupId
	DspaCgId *uint64 `json:"DspaCgId,omitnil,omitempty" name:"DspaCgId"`

	// dspa绑定模板/合规组名称
	DspaCgName *string `json:"DspaCgName,omitnil,omitempty" name:"DspaCgName"`

	// dspa实例状态 0 正常 1 隔离 2 销毁
	DspaStatus *uint64 `json:"DspaStatus,omitnil,omitempty" name:"DspaStatus"`

	// 模板状态 0: 正常   1: 已删除
	DspaCgStatus *uint64 `json:"DspaCgStatus,omitnil,omitempty" name:"DspaCgStatus"`
}

type IdWithName struct {
	// id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

// Predefined struct for user
type InquiryPriceDbauditInstanceRequestParams struct {
	// 实例规格，取值范围： cdsaudit，cdsaudit_adv， cdsaudit_ent 分别为合规版，高级版，企业版
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 询价类型： renew，续费；newbuy，新购
	InquiryType *string `json:"InquiryType,omitnil,omitempty" name:"InquiryType"`

	// 购买实例的时长。取值范围：1（y/m），2（y/m）,，3（y/m），4（m）， 5（m），6（m）， 7（m），8（m），9（m）， 10（m）
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位，y：年；m：月
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 实例所在地域
	ServiceRegion *string `json:"ServiceRegion,omitnil,omitempty" name:"ServiceRegion"`
}

type InquiryPriceDbauditInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例规格，取值范围： cdsaudit，cdsaudit_adv， cdsaudit_ent 分别为合规版，高级版，企业版
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 询价类型： renew，续费；newbuy，新购
	InquiryType *string `json:"InquiryType,omitnil,omitempty" name:"InquiryType"`

	// 购买实例的时长。取值范围：1（y/m），2（y/m）,，3（y/m），4（m）， 5（m），6（m）， 7（m），8（m），9（m）， 10（m）
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位，y：年；m：月
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 实例所在地域
	ServiceRegion *string `json:"ServiceRegion,omitnil,omitempty" name:"ServiceRegion"`
}

func (r *InquiryPriceDbauditInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceDbauditInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceVersion")
	delete(f, "InquiryType")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "ServiceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceDbauditInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceDbauditInstanceResponseParams struct {
	// 总价，单位：元
	TotalPrice *float64 `json:"TotalPrice,omitnil,omitempty" name:"TotalPrice"`

	// 真实价钱，预支费用的折扣价，单位：元
	RealTotalCost *float64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceDbauditInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceDbauditInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceDbauditInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceDbauditInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDbauditInstancesRenewFlagRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0，表示默认状态(用户未设置，即初始状态)；1，表示自动续费；2，表示明确不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type ModifyDbauditInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0，表示默认状态(用户未设置，即初始状态)；1，表示自动续费；2，表示明确不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

func (r *ModifyDbauditInstancesRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDbauditInstancesRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDbauditInstancesRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDbauditInstancesRenewFlagResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDbauditInstancesRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDbauditInstancesRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyDbauditInstancesRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDbauditInstancesRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NameValueString struct {
	// <p>名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RegionInfo struct {
	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域描述
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地域可用状态
	RegionState *int64 `json:"RegionState,omitnil,omitempty" name:"RegionState"`
}

type ReportMission struct {
	// 报表任务id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务名称
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 报表类型 1:单次报表 2:周期报表
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 报告说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 报表模板 1:综合分析报告 2:等保合规报告
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 包含资产
	AssetsList []*AssetsInfo `json:"AssetsList,omitnil,omitempty" name:"AssetsList"`

	// 下次启动时间
	NextStartTime *int64 `json:"NextStartTime,omitnil,omitempty" name:"NextStartTime"`

	// 任务状态 1:生成中 2:待生成3:已生成4:生成失败5:已暂停
	MissionStatus *int64 `json:"MissionStatus,omitnil,omitempty" name:"MissionStatus"`

	// 任务状态说明 仅生成中和生成失败有效
	MissionStatusMessage *string `json:"MissionStatusMessage,omitnil,omitempty" name:"MissionStatusMessage"`

	// 已生成报表数
	ReportCount *int64 `json:"ReportCount,omitnil,omitempty" name:"ReportCount"`

	// 任务起停 1:关闭 2:开启 仅周期报表有效
	MissionStart *int64 `json:"MissionStart,omitnil,omitempty" name:"MissionStart"`

	// 统计周期 1:24小时 7:近一周 30:近30天 90:近90天 180:
	CntDay *int64 `json:"CntDay,omitnil,omitempty" name:"CntDay"`

	// 重复周期 1:每天 2:每周 3:每月
	CntCycle *uint64 `json:"CntCycle,omitnil,omitempty" name:"CntCycle"`

	// 执行日期 重复周期为天：无意义 周：星期几 1-7  月每月
	CntTime *uint64 `json:"CntTime,omitnil,omitempty" name:"CntTime"`

	// 执行时间 格式15:04 到分钟
	CntDate *string `json:"CntDate,omitnil,omitempty" name:"CntDate"`

	// 创建者 0:内置 其余存放用户(uin)
	Receivers *string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// Notification  int  1关闭 2开启 不变更为0
	Notification *int64 `json:"Notification,omitnil,omitempty" name:"Notification"`
}

type Reports struct {
	// 生成时间
	AddTime *int64 `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 报告 ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 审计 ID
	InstanceId *int64 `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否已删除
	IsDelete *int64 `json:"IsDelete,omitnil,omitempty" name:"IsDelete"`

	// 发送目标
	Receivers *string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// 报告说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 报告文件
	ReportFile *string `json:"ReportFile,omitnil,omitempty" name:"ReportFile"`

	// 状态
	ReportStatus *int64 `json:"ReportStatus,omitnil,omitempty" name:"ReportStatus"`

	// 状态
	ReportTmpStatus *int64 `json:"ReportTmpStatus,omitnil,omitempty" name:"ReportTmpStatus"`

	// 报告类型
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// 发送结果
	SendResult *string `json:"SendResult,omitnil,omitempty" name:"SendResult"`

	// 发送类型
	SendType *string `json:"SendType,omitnil,omitempty" name:"SendType"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 报告名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 报表模板
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 包含资产
	AssetsList []*AssetsInfo `json:"AssetsList,omitnil,omitempty" name:"AssetsList"`

	// 时间范围 1:24小时 7:近一周 30:近30天 90:近90天 180:近180天 不变更为0
	CntDay *int64 `json:"CntDay,omitnil,omitempty" name:"CntDay"`
}