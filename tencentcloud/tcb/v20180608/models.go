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

package v20180608

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AuthDomain struct {
	// 域名ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名类型。包含以下取值：
	// <li>SYSTEM</li>
	// <li>USER</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 状态。包含以下取值：
	// <li>ENABLE</li>
	// <li>DISABLE</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type BaasPackageInfo struct {
	// DAU产品套餐ID
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// DAU套餐中文名称
	PackageTitle *string `json:"PackageTitle,omitnil,omitempty" name:"PackageTitle"`

	// 套餐分组
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 套餐分组中文名
	GroupTitle *string `json:"GroupTitle,omitnil,omitempty" name:"GroupTitle"`

	// json格式化计费标签，例如：
	// {"pid":2, "cids":{"create": 2, "renew": 2, "modify": 2}, "productCode":"p_tcb_mp", "subProductCode":"sp_tcb_mp_cloudbase_dau"}
	BillTags *string `json:"BillTags,omitnil,omitempty" name:"BillTags"`

	// json格式化用户资源限制，例如：
	// {"Qps":1000,"InvokeNum":{"TimeUnit":"m", "Unit":"万次", "MaxSize": 100},"Capacity":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100}, "Cdn":{"Flux":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100}, "BackFlux":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100}},"Scf":{"Concurrency":1000,"OutFlux":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100},"MemoryUse":{"TimeUnit":"m", "Unit":"WGBS", "MaxSize": 100000}}}
	ResourceLimit *string `json:"ResourceLimit,omitnil,omitempty" name:"ResourceLimit"`

	// json格式化高级限制，例如：
	// {"CMSEnable":false,"ProvisionedConcurrencyMem":512000, "PictureProcessing":false, "SecurityAudit":false, "RealTimePush":false, "TemplateMessageBatchPush":false, "Payment":false}
	AdvanceLimit *string `json:"AdvanceLimit,omitnil,omitempty" name:"AdvanceLimit"`

	// 套餐描述
	PackageDescription *string `json:"PackageDescription,omitnil,omitempty" name:"PackageDescription"`

	// 是否对外展示
	IsExternal *bool `json:"IsExternal,omitnil,omitempty" name:"IsExternal"`
}

// Predefined struct for user
type BindCloudBaseAccessDomainRequestParams struct {
	// 服务Id，目前是指环境Id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 自定义域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 腾讯云证书Id
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 默认1，1 绑定默认Cdn，2 绑定TcbIngress（不经过cdn），4 绑定自定义cdn
	BindFlag *uint64 `json:"BindFlag,omitnil,omitempty" name:"BindFlag"`

	// 自定义cdn cname域名
	CustomCname *string `json:"CustomCname,omitnil,omitempty" name:"CustomCname"`
}

type BindCloudBaseAccessDomainRequest struct {
	*tchttp.BaseRequest
	
	// 服务Id，目前是指环境Id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 自定义域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 腾讯云证书Id
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 默认1，1 绑定默认Cdn，2 绑定TcbIngress（不经过cdn），4 绑定自定义cdn
	BindFlag *uint64 `json:"BindFlag,omitnil,omitempty" name:"BindFlag"`

	// 自定义cdn cname域名
	CustomCname *string `json:"CustomCname,omitnil,omitempty" name:"CustomCname"`
}

func (r *BindCloudBaseAccessDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudBaseAccessDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Domain")
	delete(f, "CertId")
	delete(f, "BindFlag")
	delete(f, "CustomCname")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindCloudBaseAccessDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudBaseAccessDomainResponseParams struct {
	// 服务Id，目前是指环境Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindCloudBaseAccessDomainResponse struct {
	*tchttp.BaseResponse
	Response *BindCloudBaseAccessDomainResponseParams `json:"Response"`
}

func (r *BindCloudBaseAccessDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudBaseAccessDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudBaseGWDomainRequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 是否启用多地域
	EnableRegion *bool `json:"EnableRegion,omitnil,omitempty" name:"EnableRegion"`
}

type BindCloudBaseGWDomainRequest struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 是否启用多地域
	EnableRegion *bool `json:"EnableRegion,omitnil,omitempty" name:"EnableRegion"`
}

func (r *BindCloudBaseGWDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudBaseGWDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Domain")
	delete(f, "CertId")
	delete(f, "EnableRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindCloudBaseGWDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudBaseGWDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindCloudBaseGWDomainResponse struct {
	*tchttp.BaseResponse
	Response *BindCloudBaseGWDomainResponseParams `json:"Response"`
}

func (r *BindCloudBaseGWDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudBaseGWDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTcbServiceRequestParams struct {

}

type CheckTcbServiceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CheckTcbServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTcbServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckTcbServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTcbServiceResponseParams struct {
	// true表示已开通
	Initialized *bool `json:"Initialized,omitnil,omitempty" name:"Initialized"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckTcbServiceResponse struct {
	*tchttp.BaseResponse
	Response *CheckTcbServiceResponseParams `json:"Response"`
}

func (r *CheckTcbServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTcbServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudBaseClientQPSPolicy struct {
	// UserID 或 ClientIP 或 None，如果为 None 代表不限制
	LimitBy *string `json:"LimitBy,omitnil,omitempty" name:"LimitBy"`

	// 限制值
	LimitValue *int64 `json:"LimitValue,omitnil,omitempty" name:"LimitValue"`
}

type CloudBaseGWAPI struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// API Path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// API 类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// API 名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// API创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 自定义值通用字段：
	// Type为1时，该值为空。
	// Type为2时，该值为容器的代理IP:PORT数组。
	Custom *string `json:"Custom,omitnil,omitempty" name:"Custom"`

	// 表示是否开启认证
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 云开发环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 访问类型（该参数暂不对外暴露）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 统一发布状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnionStatus *int64 `json:"UnionStatus,omitnil,omitempty" name:"UnionStatus"`

	// 域名（*表示所有域名）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 是否有路径冲突
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConflictFlag *bool `json:"ConflictFlag,omitnil,omitempty" name:"ConflictFlag"`

	// 域名状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainStatus *int64 `json:"DomainStatus,omitnil,omitempty" name:"DomainStatus"`

	// 是否开启路径透传，默认true表示短路径，即不开启(已弃用)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsShortPath *bool `json:"IsShortPath,omitnil,omitempty" name:"IsShortPath"`

	// 路径透传，默认0关闭，1开启，2关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathTransmission *uint64 `json:"PathTransmission,omitnil,omitempty" name:"PathTransmission"`

	// 跨域校验，默认0开启，1开启，2关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableCheckAcrossDomain *uint64 `json:"EnableCheckAcrossDomain,omitnil,omitempty" name:"EnableCheckAcrossDomain"`

	// 静态托管文件目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaticFileDirectory *string `json:"StaticFileDirectory,omitnil,omitempty" name:"StaticFileDirectory"`

	// QPS策略
	QPSPolicy *CloudBaseGWAPIQPSPolicy `json:"QPSPolicy,omitnil,omitempty" name:"QPSPolicy"`
}

type CloudBaseGWAPIQPSPolicy struct {
	// qps限额总量
	QPSTotal *int64 `json:"QPSTotal,omitnil,omitempty" name:"QPSTotal"`

	// 客户端限频，如果不限制，LimitBy=None
	QPSPerClient *CloudBaseClientQPSPolicy `json:"QPSPerClient,omitnil,omitempty" name:"QPSPerClient"`
}

type CloudBaseGWService struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 开启时间
	OpenTime *uint64 `json:"OpenTime,omitnil,omitempty" name:"OpenTime"`

	// 绑定状态，1 绑定中；2绑定失败；3绑定成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否被抢占, 被抢占表示域名被其他环境绑定了，需要解绑或者重新绑定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPreempted *bool `json:"IsPreempted,omitnil,omitempty" name:"IsPreempted"`

	// 是否开启多地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableRegion *bool `json:"EnableRegion,omitnil,omitempty" name:"EnableRegion"`

	// cdn CName地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// 统一域名状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnionStatus *int64 `json:"UnionStatus,omitnil,omitempty" name:"UnionStatus"`

	// CName状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CnameStatus *int64 `json:"CnameStatus,omitnil,omitempty" name:"CnameStatus"`

	// 证书Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 是否强制https
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceHttps *bool `json:"ForceHttps,omitnil,omitempty" name:"ForceHttps"`

	// icp黑名单封禁状态，0-未封禁，1-封禁
	// 注意：此字段可能返回 null，表示取不到有效值。
	IcpForbidStatus *int64 `json:"IcpForbidStatus,omitnil,omitempty" name:"IcpForbidStatus"`

	// 自定义路由规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomRoutingRules *string `json:"CustomRoutingRules,omitnil,omitempty" name:"CustomRoutingRules"`

	// 绑定类型，1绑定cdn，2源站，4自定义
	BindFlag *uint64 `json:"BindFlag,omitnil,omitempty" name:"BindFlag"`

	// TcbIngress源站cname
	OriginCname *string `json:"OriginCname,omitnil,omitempty" name:"OriginCname"`

	// 自定义cname
	CustomCname *string `json:"CustomCname,omitnil,omitempty" name:"CustomCname"`
}

type CloudBaseOption struct {
	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ClsInfo struct {
	// cls所属地域
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`

	// cls日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitnil,omitempty" name:"ClsLogsetId"`

	// cls日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ClusterDetail struct {
	// 是否开启公网访问
	IsOpenPubNetAccess *bool `json:"IsOpenPubNetAccess,omitnil,omitempty" name:"IsOpenPubNetAccess"`

	// 最大算力
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 最小算力
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// TDSQL-C集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 存储用量（单位：MB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedStorage *int64 `json:"UsedStorage,omitnil,omitempty" name:"UsedStorage"`

	// 最大存储量（单位：GB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 数据库类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 数据库类型
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 公网访问状态；open开启，opening开启中，closed关闭，closing关闭中
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 数据库集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *string `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// serverless状态
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`
}

// Predefined struct for user
type CreateAuthDomainRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 安全域名
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

type CreateAuthDomainRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 安全域名
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

func (r *CreateAuthDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuthDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuthDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuthDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuthDomainResponseParams `json:"Response"`
}

func (r *CreateAuthDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBillDealRequestParams struct {
	// 当前下单的操作类型，可取[purchase,renew,modify]三种值，分别代表新购，续费，变配。
	DealType *string `json:"DealType,omitnil,omitempty" name:"DealType"`

	// 购买的产品类型，可取[tcb-baas,tcb-promotion,tcb-package], 分别代表baas套餐、大促包、资源包
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 目标下单产品/套餐Id。
	// 对于云开发环境套餐，可通过 DescribeBaasPackageList 接口获取，对应其出参的PackageName
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 默认只下单不支付，为ture则下单并支付。
	// 如果需要下单并支付，请确保账户下有足够的余额，否则会导致下单失败。
	CreateAndPay *bool `json:"CreateAndPay,omitnil,omitempty" name:"CreateAndPay"`

	// 购买时长，与TimeUnit字段搭配使用。
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位,按各产品规则可选d(天),m(月),y(年),p(一次性)。
	// 对于 云开发环境的 新购和续费，目前仅支持 按月购买（即 TimeUnit=m）。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 资源唯一标识。
	// 在云开发环境 续费和变配 场景下必传，取值为环境ID。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 来源可选[qcloud,miniapp]，默认qcloud。
	// miniapp表示微信云开发，主要适用于[小程序云开发](https://developers.weixin.qq.com/miniprogram/dev/wxcloudservice/wxcloud/billing/price.html)。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 环境别名，用于新购云开发环境时，给云开发环境起别名。
	// 仅当 新购云开发环境（DealType=purchase 并且 ProductType=tcb-baas ）时有效。
	// 
	// ### 格式要求
	// - 可选字符： 小写字母(a~z)、数字、减号(-)
	// - 不能以 减号(-) 开头或结尾
	// - 不能有连个连续的 减号(-)
	// - 长度不超过20位
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 环境id，当购买资源包和大促包时（ProductType取值为tcb-promotion 或 tcb-package）必传，表示资源包在哪个环境下生效。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 开启超限按量。
	// 开启后，当 套餐内的资源点 和 资源包 都用尽后，会自动按量计费。
	// 详见 [计费说明](https://cloud.tencent.com/document/product/876/127357)。
	EnableExcess *bool `json:"EnableExcess,omitnil,omitempty" name:"EnableExcess"`

	// 变配目标套餐id，对于云开发环境变配场景下必传。
	// 对于云开发环境套餐，可通过 DescribeBaasPackageList 接口获取，对应其出参的PackageName
	ModifyPackageId *string `json:"ModifyPackageId,omitnil,omitempty" name:"ModifyPackageId"`

	// jsonstr附加信息
	Extension *string `json:"Extension,omitnil,omitempty" name:"Extension"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 资源类型。
	// 代表新购环境（DealType=purchase 并且 ProductType=tcb-baas ）时需要发货哪些资源。
	// 可取值：flexdb, cos, cdn, scf
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// 环境标签。
	//  代表新购环境（DealType=purchase 并且 ProductType=tcb-baas ）时需要打的标签。
	EnvTags []*Tag `json:"EnvTags,omitnil,omitempty" name:"EnvTags"`
}

type CreateBillDealRequest struct {
	*tchttp.BaseRequest
	
	// 当前下单的操作类型，可取[purchase,renew,modify]三种值，分别代表新购，续费，变配。
	DealType *string `json:"DealType,omitnil,omitempty" name:"DealType"`

	// 购买的产品类型，可取[tcb-baas,tcb-promotion,tcb-package], 分别代表baas套餐、大促包、资源包
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 目标下单产品/套餐Id。
	// 对于云开发环境套餐，可通过 DescribeBaasPackageList 接口获取，对应其出参的PackageName
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 默认只下单不支付，为ture则下单并支付。
	// 如果需要下单并支付，请确保账户下有足够的余额，否则会导致下单失败。
	CreateAndPay *bool `json:"CreateAndPay,omitnil,omitempty" name:"CreateAndPay"`

	// 购买时长，与TimeUnit字段搭配使用。
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位,按各产品规则可选d(天),m(月),y(年),p(一次性)。
	// 对于 云开发环境的 新购和续费，目前仅支持 按月购买（即 TimeUnit=m）。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 资源唯一标识。
	// 在云开发环境 续费和变配 场景下必传，取值为环境ID。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 来源可选[qcloud,miniapp]，默认qcloud。
	// miniapp表示微信云开发，主要适用于[小程序云开发](https://developers.weixin.qq.com/miniprogram/dev/wxcloudservice/wxcloud/billing/price.html)。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 环境别名，用于新购云开发环境时，给云开发环境起别名。
	// 仅当 新购云开发环境（DealType=purchase 并且 ProductType=tcb-baas ）时有效。
	// 
	// ### 格式要求
	// - 可选字符： 小写字母(a~z)、数字、减号(-)
	// - 不能以 减号(-) 开头或结尾
	// - 不能有连个连续的 减号(-)
	// - 长度不超过20位
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 环境id，当购买资源包和大促包时（ProductType取值为tcb-promotion 或 tcb-package）必传，表示资源包在哪个环境下生效。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 开启超限按量。
	// 开启后，当 套餐内的资源点 和 资源包 都用尽后，会自动按量计费。
	// 详见 [计费说明](https://cloud.tencent.com/document/product/876/127357)。
	EnableExcess *bool `json:"EnableExcess,omitnil,omitempty" name:"EnableExcess"`

	// 变配目标套餐id，对于云开发环境变配场景下必传。
	// 对于云开发环境套餐，可通过 DescribeBaasPackageList 接口获取，对应其出参的PackageName
	ModifyPackageId *string `json:"ModifyPackageId,omitnil,omitempty" name:"ModifyPackageId"`

	// jsonstr附加信息
	Extension *string `json:"Extension,omitnil,omitempty" name:"Extension"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 资源类型。
	// 代表新购环境（DealType=purchase 并且 ProductType=tcb-baas ）时需要发货哪些资源。
	// 可取值：flexdb, cos, cdn, scf
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// 环境标签。
	//  代表新购环境（DealType=purchase 并且 ProductType=tcb-baas ）时需要打的标签。
	EnvTags []*Tag `json:"EnvTags,omitnil,omitempty" name:"EnvTags"`
}

func (r *CreateBillDealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBillDealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealType")
	delete(f, "ProductType")
	delete(f, "PackageId")
	delete(f, "CreateAndPay")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "ResourceId")
	delete(f, "Source")
	delete(f, "Alias")
	delete(f, "EnvId")
	delete(f, "EnableExcess")
	delete(f, "ModifyPackageId")
	delete(f, "Extension")
	delete(f, "AutoVoucher")
	delete(f, "ResourceTypes")
	delete(f, "EnvTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBillDealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBillDealResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBillDealResponse struct {
	*tchttp.BaseResponse
	Response *CreateBillDealResponseParams `json:"Response"`
}

func (r *CreateBillDealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBillDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudBaseGWAPIRequestParams struct {
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// API Path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// API类型（1表示云函数，2表示容器）
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// API Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// APIId，如果非空，表示修改绑定Path
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// 自定义值通用字段（当Type为容器时必填）
	Custom *string `json:"Custom,omitnil,omitempty" name:"Custom"`

	// 认证开关 1为开启 2为关闭
	AuthSwitch *uint64 `json:"AuthSwitch,omitnil,omitempty" name:"AuthSwitch"`

	// 是否开启多地域
	EnableRegion *bool `json:"EnableRegion,omitnil,omitempty" name:"EnableRegion"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 访问类型："OA", "PUBLIC", "MINIAPP", "VPC" （不传默认PUBLIC+MINIAPP+VPC）
	AccessTypes []*string `json:"AccessTypes,omitnil,omitempty" name:"AccessTypes"`

	// 是否开启路径透传，默认true表示短路径，即不开启路径透传(已弃用)
	IsShortPath *bool `json:"IsShortPath,omitnil,omitempty" name:"IsShortPath"`

	// 路径透传，默认0关闭，1开启，2关闭
	PathTransmission *uint64 `json:"PathTransmission,omitnil,omitempty" name:"PathTransmission"`

	// 跨域校验，默认0开启，1开启，2关闭
	EnableCheckAcrossDomain *uint64 `json:"EnableCheckAcrossDomain,omitnil,omitempty" name:"EnableCheckAcrossDomain"`

	// 静态托管资源目录
	StaticFileDirectory *string `json:"StaticFileDirectory,omitnil,omitempty" name:"StaticFileDirectory"`
}

type CreateCloudBaseGWAPIRequest struct {
	*tchttp.BaseRequest
	
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// API Path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// API类型（1表示云函数，2表示容器）
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// API Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// APIId，如果非空，表示修改绑定Path
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// 自定义值通用字段（当Type为容器时必填）
	Custom *string `json:"Custom,omitnil,omitempty" name:"Custom"`

	// 认证开关 1为开启 2为关闭
	AuthSwitch *uint64 `json:"AuthSwitch,omitnil,omitempty" name:"AuthSwitch"`

	// 是否开启多地域
	EnableRegion *bool `json:"EnableRegion,omitnil,omitempty" name:"EnableRegion"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 访问类型："OA", "PUBLIC", "MINIAPP", "VPC" （不传默认PUBLIC+MINIAPP+VPC）
	AccessTypes []*string `json:"AccessTypes,omitnil,omitempty" name:"AccessTypes"`

	// 是否开启路径透传，默认true表示短路径，即不开启路径透传(已弃用)
	IsShortPath *bool `json:"IsShortPath,omitnil,omitempty" name:"IsShortPath"`

	// 路径透传，默认0关闭，1开启，2关闭
	PathTransmission *uint64 `json:"PathTransmission,omitnil,omitempty" name:"PathTransmission"`

	// 跨域校验，默认0开启，1开启，2关闭
	EnableCheckAcrossDomain *uint64 `json:"EnableCheckAcrossDomain,omitnil,omitempty" name:"EnableCheckAcrossDomain"`

	// 静态托管资源目录
	StaticFileDirectory *string `json:"StaticFileDirectory,omitnil,omitempty" name:"StaticFileDirectory"`
}

func (r *CreateCloudBaseGWAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudBaseGWAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Path")
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "APIId")
	delete(f, "Custom")
	delete(f, "AuthSwitch")
	delete(f, "EnableRegion")
	delete(f, "EnableUnion")
	delete(f, "Domain")
	delete(f, "AccessTypes")
	delete(f, "IsShortPath")
	delete(f, "PathTransmission")
	delete(f, "EnableCheckAcrossDomain")
	delete(f, "StaticFileDirectory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudBaseGWAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudBaseGWAPIResponseParams struct {
	// API ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudBaseGWAPIResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudBaseGWAPIResponseParams `json:"Response"`
}

func (r *CreateCloudBaseGWAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudBaseGWAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvRequestParams struct {
	// 环境别名。
	// 
	// ### 格式要求
	// - 可选字符： 小写字母(a~z)、数字、减号(-)
	// - 不能以 减号(-) 开头或结尾
	// - 不能有连个连续的 减号(-)
	// - 长度不超过20位
	// 示例值：cloud
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 云开发环境套餐Id。
	// 对于云开发环境套餐，可通过 [DescribeBaasPackageList](https://cloud.tencent.com/document/product/876/78167) 接口获取，对应其出参的PackageName。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 资源类型。代表新购环境时需要发货哪些资源。
	// 可取值以及含义：
	// - flexdb : 表示文档型数据库
	// - storage : 表示云存储
	// - function : 表示云函数
	// 
	// **该字段不可为空**
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24。
	// 默认值为1，即1个月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 环境标签。
	// 可取值通过接口 [tag:DescribeTags](https://cloud.tencent.com/document/product/651/35316) 可获取到。
	// 不传或为空则默认不打任何标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自动续费标识。取值范围：
	// - NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// - NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费（需要手动续费，可通过接口 [RenewEnv](https://cloud.tencent.com/document/product/876/128590) 来续费）
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。
	// 若该参数指定为NOTIFY_AND_AUTO_RENEW（即：自动续费），在账户余额充足的情况下，实例到期后将按月自动续费；但如果账户余额不足，将无法自动续费。请留意腾讯云短信和邮件通知。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type CreateEnvRequest struct {
	*tchttp.BaseRequest
	
	// 环境别名。
	// 
	// ### 格式要求
	// - 可选字符： 小写字母(a~z)、数字、减号(-)
	// - 不能以 减号(-) 开头或结尾
	// - 不能有连个连续的 减号(-)
	// - 长度不超过20位
	// 示例值：cloud
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 云开发环境套餐Id。
	// 对于云开发环境套餐，可通过 [DescribeBaasPackageList](https://cloud.tencent.com/document/product/876/78167) 接口获取，对应其出参的PackageName。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 资源类型。代表新购环境时需要发货哪些资源。
	// 可取值以及含义：
	// - flexdb : 表示文档型数据库
	// - storage : 表示云存储
	// - function : 表示云函数
	// 
	// **该字段不可为空**
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24。
	// 默认值为1，即1个月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 环境标签。
	// 可取值通过接口 [tag:DescribeTags](https://cloud.tencent.com/document/product/651/35316) 可获取到。
	// 不传或为空则默认不打任何标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自动续费标识。取值范围：
	// - NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// - NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费（需要手动续费，可通过接口 [RenewEnv](https://cloud.tencent.com/document/product/876/128590) 来续费）
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。
	// 若该参数指定为NOTIFY_AND_AUTO_RENEW（即：自动续费），在账户余额充足的情况下，实例到期后将按月自动续费；但如果账户余额不足，将无法自动续费。请留意腾讯云短信和邮件通知。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

func (r *CreateEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "PackageId")
	delete(f, "Resources")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "Tags")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvResponseParams struct {
	// 自动生成的环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEnvResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvResponseParams `json:"Response"`
}

func (r *CreateEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostingDomainRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type CreateHostingDomainRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

func (r *CreateHostingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostingDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Domain")
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHostingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostingDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateHostingDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateHostingDomainResponseParams `json:"Response"`
}

func (r *CreateHostingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostingDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIndex struct {
	// 索引名称
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 索引结构
	MgoKeySchema *MgoKeySchema `json:"MgoKeySchema,omitnil,omitempty" name:"MgoKeySchema"`
}

// Predefined struct for user
type CreateMySQLRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Db类型: MYSQL
	DbInstanceType *string `json:"DbInstanceType,omitnil,omitempty" name:"DbInstanceType"`

	// mysql版本
	MysqlVersion *string `json:"MysqlVersion,omitnil,omitempty" name:"MysqlVersion"`

	// vpc Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 0 区分表名大小写；1 不区分表名大小写(默认)
	LowerCaseTableNames *string `json:"LowerCaseTableNames,omitnil,omitempty" name:"LowerCaseTableNames"`
}

type CreateMySQLRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Db类型: MYSQL
	DbInstanceType *string `json:"DbInstanceType,omitnil,omitempty" name:"DbInstanceType"`

	// mysql版本
	MysqlVersion *string `json:"MysqlVersion,omitnil,omitempty" name:"MysqlVersion"`

	// vpc Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 0 区分表名大小写；1 不区分表名大小写(默认)
	LowerCaseTableNames *string `json:"LowerCaseTableNames,omitnil,omitempty" name:"LowerCaseTableNames"`
}

func (r *CreateMySQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMySQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "DbInstanceType")
	delete(f, "MysqlVersion")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "LowerCaseTableNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMySQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMySQLResponseParams struct {
	// 开通结果
	Data *CreateMySQLResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMySQLResponse struct {
	*tchttp.BaseResponse
	Response *CreateMySQLResponseParams `json:"Response"`
}

func (r *CreateMySQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMySQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMySQLResult struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CreateStaticStoreRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`
}

type CreateStaticStoreRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`
}

func (r *CreateStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaticStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "EnableUnion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStaticStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStaticStoreResponseParams struct {
	// 创建静态资源结果(succ/fail)
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *CreateStaticStoreResponseParams `json:"Response"`
}

func (r *CreateStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaticStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableRequestParams struct {
	// 数据表名；长度不超过96个字符，可以为英文字母、数字、下划线(_)和短横线(-)的组合，且不能以下划线开头
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID，如：tnt-nl7hjzasw
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// FlexDB数据库权限信息
	PermissionInfo *PermissionInfo `json:"PermissionInfo,omitnil,omitempty" name:"PermissionInfo"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type CreateTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据表名；长度不超过96个字符，可以为英文字母、数字、下划线(_)和短横线(-)的组合，且不能以下划线开头
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID，如：tnt-nl7hjzasw
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// FlexDB数据库权限信息
	PermissionInfo *PermissionInfo `json:"PermissionInfo,omitnil,omitempty" name:"PermissionInfo"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *CreateTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "Tag")
	delete(f, "PermissionInfo")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateTableResponseParams `json:"Response"`
}

func (r *CreateTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 用户名，用户名规则：1. 长度1-64字符 2. 只能包含大小写英文字母、数字和符号 . _ - 3. 只能以字母或数字开头 4. 不能重复
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户ID，最多64字符，如不传则系统自动生成
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户类型：internalUser-内部用户、externalUser-外部用户，默认internalUser（内部用户）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 密码，传入Uid时密码可不传。密码规则：1. 长度8-32字符（推荐12位以上） 2. 不能以特殊字符开头 3. 至少包含以下四项中的三项：小写字母a-z、大写字母A-Z、数字0-9、特殊字符()!@#$%^&*\|?><_-
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结），默认激活
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称，长度2-64字符
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，不能重复
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱地址，不能重复
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接，可访问的公网URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述，最多200字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 用户名，用户名规则：1. 长度1-64字符 2. 只能包含大小写英文字母、数字和符号 . _ - 3. 只能以字母或数字开头 4. 不能重复
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户ID，最多64字符，如不传则系统自动生成
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户类型：internalUser-内部用户、externalUser-外部用户，默认internalUser（内部用户）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 密码，传入Uid时密码可不传。密码规则：1. 长度8-32字符（推荐12位以上） 2. 不能以特殊字符开头 3. 至少包含以下四项中的三项：小写字母a-z、大写字母A-Z、数字0-9、特殊字符()!@#$%^&*\|?><_-
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结），默认激活
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称，长度2-64字符
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，不能重复
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱地址，不能重复
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接，可访问的公网URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述，最多200字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Name")
	delete(f, "Uid")
	delete(f, "Type")
	delete(f, "Password")
	delete(f, "UserStatus")
	delete(f, "NickName")
	delete(f, "Phone")
	delete(f, "Email")
	delete(f, "AvatarUrl")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResp struct {
	// 用户ID
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 结果返回
	Data *CreateUserResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatabasesInfo struct {
	// 数据库唯一标识
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 状态。包含以下取值：
	// <li>INITIALIZING：资源初始化中</li>
	// <li>RUNNING：运行中，可正常使用的状态</li>
	// <li>UNUSABLE：禁用，不可用</li>
	// <li>OVERDUE：资源过期</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DbInstance struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MySQL 连接器实例 ID；`"default"` 或为空表示使用 TCB 环境的默认连接器
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名；为空时使用连接器配置的默认数据库名
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`
}

// Predefined struct for user
type DeleteAuthDomainRequestParams struct {
	// 开发者的环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名ID列表，支持批量
	DomainIds []*string `json:"DomainIds,omitnil,omitempty" name:"DomainIds"`
}

type DeleteAuthDomainRequest struct {
	*tchttp.BaseRequest
	
	// 开发者的环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名ID列表，支持批量
	DomainIds []*string `json:"DomainIds,omitnil,omitempty" name:"DomainIds"`
}

func (r *DeleteAuthDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "DomainIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuthDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuthDomainResponseParams struct {
	// 删除的域名个数
	Deleted *int64 `json:"Deleted,omitnil,omitempty" name:"Deleted"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuthDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuthDomainResponseParams `json:"Response"`
}

func (r *DeleteAuthDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudBaseGWAPIRequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// API Path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// API类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// API Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义值字段（Type为2时，传递容器服务名表示需要删除JNSGW）
	Custom *string `json:"Custom,omitnil,omitempty" name:"Custom"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DeleteCloudBaseGWAPIRequest struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// API Path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// API类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// API Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义值字段（Type为2时，传递容器服务名表示需要删除JNSGW）
	Custom *string `json:"Custom,omitnil,omitempty" name:"Custom"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *DeleteCloudBaseGWAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudBaseGWAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Path")
	delete(f, "APIId")
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Custom")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudBaseGWAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudBaseGWAPIResponseParams struct {
	// 最终删除API个数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudBaseGWAPIResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudBaseGWAPIResponseParams `json:"Response"`
}

func (r *DeleteCloudBaseGWAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudBaseGWAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudBaseGWDomainRequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DeleteCloudBaseGWDomainRequest struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *DeleteCloudBaseGWDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudBaseGWDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudBaseGWDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudBaseGWDomainResponseParams struct {
	// 删除个数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudBaseGWDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudBaseGWDomainResponseParams `json:"Response"`
}

func (r *DeleteCloudBaseGWDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudBaseGWDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableRequestParams struct {
	// 待删除的表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type DeleteTableRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *DeleteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "Tag")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableResponseParams `json:"Response"`
}

func (r *DeleteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// tcb用户id列表, 一次最多支持删除100个
	Uids []*string `json:"Uids,omitnil,omitempty" name:"Uids"`
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// tcb用户id列表, 一次最多支持删除100个
	Uids []*string `json:"Uids,omitnil,omitempty" name:"Uids"`
}

func (r *DeleteUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Uids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersResp struct {
	// 成功个数
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 失败个数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`
}

// Predefined struct for user
type DeleteUsersResponseParams struct {
	// 删除用户结果
	Data *DeleteUsersResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUsersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsersResponseParams `json:"Response"`
}

func (r *DeleteUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthDomainsRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeAuthDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeAuthDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuthDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthDomainsResponseParams struct {
	// 安全域名列表
	Domains []*AuthDomain `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuthDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuthDomainsResponseParams `json:"Response"`
}

func (r *DescribeAuthDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaasPackageListRequestParams struct {
	// tcb产品套餐ID，不填拉取全量package信息。
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 套餐归属方，填写后只返回对应的套餐 包含miniapp与qcloud两种 默认为miniapp
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 套餐归属环境渠道
	EnvChannel *string `json:"EnvChannel,omitnil,omitempty" name:"EnvChannel"`

	// 拉取套餐用途：
	// 1）new 新购
	// 2）modify变配
	// 3）renew续费
	TargetAction *string `json:"TargetAction,omitnil,omitempty" name:"TargetAction"`

	// 预留字段，同一商品会对应多个类型套餐，对指标有不同侧重。
	// 计算型calculation
	// 流量型flux
	// 容量型capactiy
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 类型分组过滤。默认为["default"]
	PackageTypeList []*string `json:"PackageTypeList,omitnil,omitempty" name:"PackageTypeList"`

	// 付费渠道，与回包billTags中的计费参数相关，不填返回默认值。
	PaymentChannel *string `json:"PaymentChannel,omitnil,omitempty" name:"PaymentChannel"`
}

type DescribeBaasPackageListRequest struct {
	*tchttp.BaseRequest
	
	// tcb产品套餐ID，不填拉取全量package信息。
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 套餐归属方，填写后只返回对应的套餐 包含miniapp与qcloud两种 默认为miniapp
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 套餐归属环境渠道
	EnvChannel *string `json:"EnvChannel,omitnil,omitempty" name:"EnvChannel"`

	// 拉取套餐用途：
	// 1）new 新购
	// 2）modify变配
	// 3）renew续费
	TargetAction *string `json:"TargetAction,omitnil,omitempty" name:"TargetAction"`

	// 预留字段，同一商品会对应多个类型套餐，对指标有不同侧重。
	// 计算型calculation
	// 流量型flux
	// 容量型capactiy
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 类型分组过滤。默认为["default"]
	PackageTypeList []*string `json:"PackageTypeList,omitnil,omitempty" name:"PackageTypeList"`

	// 付费渠道，与回包billTags中的计费参数相关，不填返回默认值。
	PaymentChannel *string `json:"PaymentChannel,omitnil,omitempty" name:"PaymentChannel"`
}

func (r *DescribeBaasPackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaasPackageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageName")
	delete(f, "EnvId")
	delete(f, "Source")
	delete(f, "EnvChannel")
	delete(f, "TargetAction")
	delete(f, "GroupName")
	delete(f, "PackageTypeList")
	delete(f, "PaymentChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaasPackageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaasPackageListResponseParams struct {
	// 套餐列表
	PackageList []*BaasPackageInfo `json:"PackageList,omitnil,omitempty" name:"PackageList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBaasPackageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaasPackageListResponseParams `json:"Response"`
}

func (r *DescribeBaasPackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaasPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseBuildServiceRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// build类型,枚举值有: cloudbaserun, framework-ci
	CIBusiness *string `json:"CIBusiness,omitnil,omitempty" name:"CIBusiness"`

	// 服务版本
	ServiceVersion *string `json:"ServiceVersion,omitnil,omitempty" name:"ServiceVersion"`

	// 文件后缀
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`
}

type DescribeCloudBaseBuildServiceRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// build类型,枚举值有: cloudbaserun, framework-ci
	CIBusiness *string `json:"CIBusiness,omitnil,omitempty" name:"CIBusiness"`

	// 服务版本
	ServiceVersion *string `json:"ServiceVersion,omitnil,omitempty" name:"ServiceVersion"`

	// 文件后缀
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`
}

func (r *DescribeCloudBaseBuildServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseBuildServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServiceName")
	delete(f, "CIBusiness")
	delete(f, "ServiceVersion")
	delete(f, "Suffix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseBuildServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseBuildServiceResponseParams struct {
	// 上传url
	UploadUrl *string `json:"UploadUrl,omitnil,omitempty" name:"UploadUrl"`

	// 上传header
	UploadHeaders []*KVPair `json:"UploadHeaders,omitnil,omitempty" name:"UploadHeaders"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 包版本
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// 下载链接
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 下载Httpheader
	DownloadHeaders []*KVPair `json:"DownloadHeaders,omitnil,omitempty" name:"DownloadHeaders"`

	// 下载链接是否过期
	OutDate *bool `json:"OutDate,omitnil,omitempty" name:"OutDate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudBaseBuildServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseBuildServiceResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseBuildServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseBuildServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseGWAPIRequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// API域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// API Path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// API类型，1为云函数，2为容器
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// API名，Type为1时为云函数名，Type为2时为容器服务名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 查询的分页参数，用于设置查询的偏移位置，0表示从头开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询的分页参数，用于表示每次查询的最大返回数据量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否启用多地域
	EnableRegion *bool `json:"EnableRegion,omitnil,omitempty" name:"EnableRegion"`

	// 是否使用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`
}

type DescribeCloudBaseGWAPIRequest struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// API域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// API Path
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// API类型，1为云函数，2为容器
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// API名，Type为1时为云函数名，Type为2时为容器服务名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 查询的分页参数，用于设置查询的偏移位置，0表示从头开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询的分页参数，用于表示每次查询的最大返回数据量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否启用多地域
	EnableRegion *bool `json:"EnableRegion,omitnil,omitempty" name:"EnableRegion"`

	// 是否使用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`
}

func (r *DescribeCloudBaseGWAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseGWAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Domain")
	delete(f, "Path")
	delete(f, "APIId")
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EnableRegion")
	delete(f, "EnableUnion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseGWAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseGWAPIResponseParams struct {
	// API列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	APISet []*CloudBaseGWAPI `json:"APISet,omitnil,omitempty" name:"APISet"`

	// 是否开启http调用
	EnableService *bool `json:"EnableService,omitnil,omitempty" name:"EnableService"`

	// 查询结果的数据总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 查询的分页参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询的分页参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudBaseGWAPIResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseGWAPIResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseGWAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseGWAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseGWServiceRequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 是否启用多地域
	EnableRegion *bool `json:"EnableRegion,omitnil,omitempty" name:"EnableRegion"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`
}

type DescribeCloudBaseGWServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 是否启用多地域
	EnableRegion *bool `json:"EnableRegion,omitnil,omitempty" name:"EnableRegion"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`
}

func (r *DescribeCloudBaseGWServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseGWServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Domain")
	delete(f, "EnableRegion")
	delete(f, "EnableUnion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseGWServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseGWServiceResponseParams struct {
	// 服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceSet []*CloudBaseGWService `json:"ServiceSet,omitnil,omitempty" name:"ServiceSet"`

	// 是否开启服务
	EnableService *bool `json:"EnableService,omitnil,omitempty" name:"EnableService"`

	// 默认域名信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDomain *string `json:"DefaultDomain,omitnil,omitempty" name:"DefaultDomain"`

	// 是否开启CDN迁移
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`

	// 是否开启跨域校验，默认开启 true
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableCheckAcrossDomain *bool `json:"EnableCheckAcrossDomain,omitnil,omitempty" name:"EnableCheckAcrossDomain"`

	// 自定义路由规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomRoutingRules *string `json:"CustomRoutingRules,omitnil,omitempty" name:"CustomRoutingRules"`

	// 默认域名绑定类型，1绑定TCB-CDN，2绑定tcbingres（不经过cdn）
	AccessFlag *uint64 `json:"AccessFlag,omitnil,omitempty" name:"AccessFlag"`

	// 云接入源站域名
	OriginDomain *string `json:"OriginDomain,omitnil,omitempty" name:"OriginDomain"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudBaseGWServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseGWServiceResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseGWServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseGWServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCreateMySQLResult struct {
	// 状态 notexist | init | doing | success | fail
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 是否已被冻结（只在 Status=success时有效）
	FreezeStatus *bool `json:"FreezeStatus,omitnil,omitempty" name:"FreezeStatus"`
}

// Predefined struct for user
type DescribeCreateMySQLResultRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// OpenMysql 返回任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeCreateMySQLResultRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// OpenMysql 返回任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeCreateMySQLResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCreateMySQLResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCreateMySQLResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCreateMySQLResultResponseParams struct {
	// 查询开通结果
	Data *DescribeCreateMySQLResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCreateMySQLResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCreateMySQLResultResponseParams `json:"Response"`
}

func (r *DescribeCreateMySQLResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCreateMySQLResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseACLRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`
}

type DescribeDatabaseACLRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`
}

func (r *DescribeDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseACLResponseParams struct {
	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseACLResponseParams `json:"Response"`
}

func (r *DescribeDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvAccountCircleRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeEnvAccountCircleRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeEnvAccountCircleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvAccountCircleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvAccountCircleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvAccountCircleResponseParams struct {
	// 环境计费周期开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 环境计费周期结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvAccountCircleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvAccountCircleResponseParams `json:"Response"`
}

func (r *DescribeEnvAccountCircleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvAccountCircleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvLimitRequestParams struct {

}

type DescribeEnvLimitRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEnvLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvLimitResponseParams struct {
	// 环境总数上限
	MaxEnvNum *int64 `json:"MaxEnvNum,omitnil,omitempty" name:"MaxEnvNum"`

	// 目前环境总数
	CurrentEnvNum *int64 `json:"CurrentEnvNum,omitnil,omitempty" name:"CurrentEnvNum"`

	// 免费环境数量上限
	MaxFreeEnvNum *int64 `json:"MaxFreeEnvNum,omitnil,omitempty" name:"MaxFreeEnvNum"`

	// 目前免费环境数量
	CurrentFreeEnvNum *int64 `json:"CurrentFreeEnvNum,omitnil,omitempty" name:"CurrentFreeEnvNum"`

	// 总计允许销毁环境次数上限
	MaxDeleteTotal *int64 `json:"MaxDeleteTotal,omitnil,omitempty" name:"MaxDeleteTotal"`

	// 目前已销毁环境次数
	CurrentDeleteTotal *int64 `json:"CurrentDeleteTotal,omitnil,omitempty" name:"CurrentDeleteTotal"`

	// 每月允许销毁环境次数上限
	MaxDeleteMonthly *int64 `json:"MaxDeleteMonthly,omitnil,omitempty" name:"MaxDeleteMonthly"`

	// 本月已销毁环境次数
	CurrentDeleteMonthly *int64 `json:"CurrentDeleteMonthly,omitnil,omitempty" name:"CurrentDeleteMonthly"`

	// 微信网关体验版可购买月份数
	MaxFreeTrialNum *int64 `json:"MaxFreeTrialNum,omitnil,omitempty" name:"MaxFreeTrialNum"`

	// 微信网关体验版已购买月份数
	CurrentFreeTrialNum *int64 `json:"CurrentFreeTrialNum,omitnil,omitempty" name:"CurrentFreeTrialNum"`

	// 转支付限额总数
	ChangePayTotal *int64 `json:"ChangePayTotal,omitnil,omitempty" name:"ChangePayTotal"`

	// 当前已用转支付次数
	CurrentChangePayTotal *int64 `json:"CurrentChangePayTotal,omitnil,omitempty" name:"CurrentChangePayTotal"`

	// 转支付每月限额
	ChangePayMonthly *int64 `json:"ChangePayMonthly,omitnil,omitempty" name:"ChangePayMonthly"`

	// 本月已用转支付额度
	CurrentChangePayMonthly *int64 `json:"CurrentChangePayMonthly,omitnil,omitempty" name:"CurrentChangePayMonthly"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvLimitResponseParams `json:"Response"`
}

func (r *DescribeEnvLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvsRequestParams struct {
	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitnil,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitnil,omitempty" name:"Channels"`

	// 分页参数，单页限制个数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeEnvsRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitnil,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitnil,omitempty" name:"Channels"`

	// 分页参数，单页限制个数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "IsVisible")
	delete(f, "Channels")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvsResponseParams struct {
	// 环境信息列表
	EnvList []*EnvInfo `json:"EnvList,omitnil,omitempty" name:"EnvList"`

	// 环境个数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvsResponseParams `json:"Response"`
}

func (r *DescribeEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostingDomainTaskRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeHostingDomainTaskRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeHostingDomainTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostingDomainTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostingDomainTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostingDomainTaskResponseParams struct {
	// todo/doing/done/error
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostingDomainTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostingDomainTaskResponseParams `json:"Response"`
}

func (r *DescribeHostingDomainTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostingDomainTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySQLClusterDetailRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeMySQLClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeMySQLClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySQLClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMySQLClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySQLClusterDetailResponseParams struct {
	// 集群详情
	Data *MySQLClusterDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMySQLClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMySQLClusterDetailResponseParams `json:"Response"`
}

func (r *DescribeMySQLClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySQLClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySQLTaskStatusRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

type DescribeMySQLTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

func (r *DescribeMySQLTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySQLTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "TaskId")
	delete(f, "TaskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMySQLTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySQLTaskStatusResponseParams struct {
	// 任务状态
	Data *MySQLTaskStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMySQLTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMySQLTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeMySQLTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySQLTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaDataRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <li> 指标名: </li>
	// <li> StorageSizepkg: 当月存储空间容量, 单位MB </li>
	// <li> StorageReadpkg: 当月存储读请求次数 </li>
	// <li> StorageWritepkg: 当月存储写请求次数 </li>
	// <li> StorageCdnOriginFluxpkg: 当月CDN回源流量, 单位字节 </li>
	// <li> StorageCdnOriginFluxpkgDay: 当日CDN回源流量, 单位字节 </li>
	// <li> StorageReadpkgDay: 当日存储读请求次数 </li>
	// <li> StorageWritepkgDay: 当日写请求次数 </li>
	// <li> CDNFluxpkg: 当月CDN流量, 单位为字节 </li>
	// <li> CDNFluxpkgDay: 当日CDN流量, 单位为字节 </li>
	// <li> FunctionInvocationpkg: 当月云函数调用次数 </li>
	// <li> FunctionGBspkg: 当月云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkg: 当月云函数流量, 单位千字节(KB) </li>
	// <li> FunctionInvocationpkgDay: 当日云函数调用次数 </li>
	// <li> FunctionGBspkgDay: 当日云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkgDay: 当日云函数流量, 单位千字节(KB) </li>
	// <li> DbSizepkg: 当月数据库容量大小, 单位MB </li>
	// <li> DbReadpkg: 当日数据库读请求数 </li>
	// <li> DbWritepkg: 当日数据库写请求数 </li>
	// <li> StaticFsFluxPkgDay: 当日静态托管流量 </li>
	// <li> StaticFsFluxPkg: 当月静态托管流量</li>
	// <li> StaticFsSizePkg: 当月静态托管容量 </li>
	// <li> TkeCpuUsedPkg: 当月容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeCpuUsedPkgDay: 当天容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeMemUsedPkg: 当月容器托管内存使用量，单位MB*秒 </li>
	// <li> TkeMemUsedPkgDay: 当天容器托管内存使用量，单位MB*秒 </li>
	// <li> CodingBuildTimePkgDay: 当天容器托管构建时间使用量，单位毫秒 </li>
	// <li> TkeHttpServiceNatPkgDay: 当天容器托管流量使用量，单位B </li>
	// <li> CynosdbCcupkg: 当月微信云托管MySQL CCU使用量，单位个  （需要除以1000）</li>
	// <li> CynosdbStoragepkg: 当月微信云托管MySQL 存储使用量，单位MB  （需要除以1000）</li>
	// <li> CynosdbCcupkgDay: 当天微信云托管MySQL 存储使用量，单位个 （需要除以1000） </li>
	// <li> CynosdbStoragepkgDay: 当天微信云托管MySQL 存储使用量，单位MB （需要除以1000） </li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 资源ID, 目前仅对云函数、容器托管相关的指标有意义。云函数(FunctionInvocationpkg, FunctionGBspkg, FunctionFluxpkg)、容器托管（服务名称）。如果想查询某个云函数的指标则在ResourceId中传入函数名; 如果只想查询整个namespace的指标, 则留空或不传。
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`
}

type DescribeQuotaDataRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <li> 指标名: </li>
	// <li> StorageSizepkg: 当月存储空间容量, 单位MB </li>
	// <li> StorageReadpkg: 当月存储读请求次数 </li>
	// <li> StorageWritepkg: 当月存储写请求次数 </li>
	// <li> StorageCdnOriginFluxpkg: 当月CDN回源流量, 单位字节 </li>
	// <li> StorageCdnOriginFluxpkgDay: 当日CDN回源流量, 单位字节 </li>
	// <li> StorageReadpkgDay: 当日存储读请求次数 </li>
	// <li> StorageWritepkgDay: 当日写请求次数 </li>
	// <li> CDNFluxpkg: 当月CDN流量, 单位为字节 </li>
	// <li> CDNFluxpkgDay: 当日CDN流量, 单位为字节 </li>
	// <li> FunctionInvocationpkg: 当月云函数调用次数 </li>
	// <li> FunctionGBspkg: 当月云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkg: 当月云函数流量, 单位千字节(KB) </li>
	// <li> FunctionInvocationpkgDay: 当日云函数调用次数 </li>
	// <li> FunctionGBspkgDay: 当日云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkgDay: 当日云函数流量, 单位千字节(KB) </li>
	// <li> DbSizepkg: 当月数据库容量大小, 单位MB </li>
	// <li> DbReadpkg: 当日数据库读请求数 </li>
	// <li> DbWritepkg: 当日数据库写请求数 </li>
	// <li> StaticFsFluxPkgDay: 当日静态托管流量 </li>
	// <li> StaticFsFluxPkg: 当月静态托管流量</li>
	// <li> StaticFsSizePkg: 当月静态托管容量 </li>
	// <li> TkeCpuUsedPkg: 当月容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeCpuUsedPkgDay: 当天容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeMemUsedPkg: 当月容器托管内存使用量，单位MB*秒 </li>
	// <li> TkeMemUsedPkgDay: 当天容器托管内存使用量，单位MB*秒 </li>
	// <li> CodingBuildTimePkgDay: 当天容器托管构建时间使用量，单位毫秒 </li>
	// <li> TkeHttpServiceNatPkgDay: 当天容器托管流量使用量，单位B </li>
	// <li> CynosdbCcupkg: 当月微信云托管MySQL CCU使用量，单位个  （需要除以1000）</li>
	// <li> CynosdbStoragepkg: 当月微信云托管MySQL 存储使用量，单位MB  （需要除以1000）</li>
	// <li> CynosdbCcupkgDay: 当天微信云托管MySQL 存储使用量，单位个 （需要除以1000） </li>
	// <li> CynosdbStoragepkgDay: 当天微信云托管MySQL 存储使用量，单位MB （需要除以1000） </li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 资源ID, 目前仅对云函数、容器托管相关的指标有意义。云函数(FunctionInvocationpkg, FunctionGBspkg, FunctionFluxpkg)、容器托管（服务名称）。如果想查询某个云函数的指标则在ResourceId中传入函数名; 如果只想查询整个namespace的指标, 则留空或不传。
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`
}

func (r *DescribeQuotaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "MetricName")
	delete(f, "ResourceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaDataResponseParams struct {
	// 指标名
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 指标的值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 指标的附加值信息
	SubValue *string `json:"SubValue,omitnil,omitempty" name:"SubValue"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQuotaDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotaDataResponseParams `json:"Response"`
}

func (r *DescribeQuotaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeRuleRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 微信AppId，微信必传
	WxAppId *string `json:"WxAppId,omitnil,omitempty" name:"WxAppId"`
}

type DescribeSafeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 微信AppId，微信必传
	WxAppId *string `json:"WxAppId,omitnil,omitempty" name:"WxAppId"`
}

func (r *DescribeSafeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	delete(f, "WxAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSafeRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeRuleResponseParams struct {
	// 规则内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	// <li> CUSTOM：自定义安全规则</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSafeRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSafeRuleResponseParams `json:"Response"`
}

func (r *DescribeSafeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaticStoreRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeStaticStoreRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaticStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaticStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaticStoreResponseParams struct {
	// 静态托管资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*StaticStoreInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStaticStoreResponseParams `json:"Response"`
}

func (r *DescribeStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaticStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableRequestParams struct {
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlecDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type DescribeTableRequest struct {
	*tchttp.BaseRequest
	
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlecDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *DescribeTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "Tag")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableResponseParams struct {
	// 索引相关信息
	Indexes []*IndexInfo `json:"Indexes,omitnil,omitempty" name:"Indexes"`

	// 索引个数
	IndexNum *int64 `json:"IndexNum,omitnil,omitempty" name:"IndexNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableResponseParams `json:"Response"`
}

func (r *DescribeTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// 分页条件
	MgoLimit *int64 `json:"MgoLimit,omitnil,omitempty" name:"MgoLimit"`

	// 实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 分页条件
	MgoOffset *int64 `json:"MgoOffset,omitnil,omitempty" name:"MgoOffset"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoConnector
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`

	// 指定表名过滤，为空时返回所有表
	TableNames []*string `json:"TableNames,omitnil,omitempty" name:"TableNames"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// 分页条件
	MgoLimit *int64 `json:"MgoLimit,omitnil,omitempty" name:"MgoLimit"`

	// 实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 分页条件
	MgoOffset *int64 `json:"MgoOffset,omitnil,omitempty" name:"MgoOffset"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoConnector
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`

	// 指定表名过滤，为空时返回所有表
	TableNames []*string `json:"TableNames,omitnil,omitempty" name:"TableNames"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MgoLimit")
	delete(f, "Tag")
	delete(f, "MgoOffset")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	delete(f, "TableNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// 表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*TableInfo `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pager *Pager `json:"Pager,omitnil,omitempty" name:"Pager"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesResponseParams `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserListRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 页码，从1开始，默认1
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页数量，默认20，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 用户名，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户昵称，模糊查询
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，模糊查询
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱，模糊查询
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

type DescribeUserListRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 页码，从1开始，默认1
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页数量，默认20，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 用户名，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户昵称，模糊查询
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，模糊查询
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱，模糊查询
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

func (r *DescribeUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "Name")
	delete(f, "NickName")
	delete(f, "Phone")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserListResp struct {
	// 用户总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 用户列表
	UserList []*User `json:"UserList,omitnil,omitempty" name:"UserList"`
}

// Predefined struct for user
type DescribeUserListResponseParams struct {
	// 结果返回
	Data *DescribeUserListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserListResponseParams `json:"Response"`
}

func (r *DescribeUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 针对预付费 删除隔离中的环境时要传true 正常环境直接跳过隔离期删除
	IsForce *bool `json:"IsForce,omitnil,omitempty" name:"IsForce"`

	// 是否绕过资源检查，资源包等额外资源，默认为false，如果为true，则不检查资源是否有数据，直接删除。
	BypassCheck *bool `json:"BypassCheck,omitnil,omitempty" name:"BypassCheck"`
}

type DestroyEnvRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 针对预付费 删除隔离中的环境时要传true 正常环境直接跳过隔离期删除
	IsForce *bool `json:"IsForce,omitnil,omitempty" name:"IsForce"`

	// 是否绕过资源检查，资源包等额外资源，默认为false，如果为true，则不检查资源是否有数据，直接删除。
	BypassCheck *bool `json:"BypassCheck,omitnil,omitempty" name:"BypassCheck"`
}

func (r *DestroyEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "IsForce")
	delete(f, "BypassCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyEnvResponse struct {
	*tchttp.BaseResponse
	Response *DestroyEnvResponseParams `json:"Response"`
}

func (r *DestroyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyMySQLRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DestroyMySQLRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DestroyMySQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyMySQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyMySQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyMySQLResponseParams struct {
	// 销毁结果
	Data *DestroyMySQLResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyMySQLResponse struct {
	*tchttp.BaseResponse
	Response *DestroyMySQLResponseParams `json:"Response"`
}

func (r *DestroyMySQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyMySQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyMySQLResult struct {
	// 是否成功
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

// Predefined struct for user
type DestroyStaticStoreRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// cdn域名
	CdnDomain *string `json:"CdnDomain,omitnil,omitempty" name:"CdnDomain"`
}

type DestroyStaticStoreRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// cdn域名
	CdnDomain *string `json:"CdnDomain,omitnil,omitempty" name:"CdnDomain"`
}

func (r *DestroyStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyStaticStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CdnDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyStaticStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyStaticStoreResponseParams struct {
	// 条件任务结果(succ/fail)
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *DestroyStaticStoreResponseParams `json:"Response"`
}

func (r *DestroyStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyStaticStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DropIndex struct {
	// 索引名称
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`
}

// Predefined struct for user
type EditAuthConfigRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 手机号登录配置 "TRUE",  "FALSE", "LOGIN_ONLY"
	PhoneNumberLogin *string `json:"PhoneNumberLogin,omitnil,omitempty" name:"PhoneNumberLogin"`

	// 匿名登录配置 "TRUE",  "FALSE"
	AnonymousLogin *string `json:"AnonymousLogin,omitnil,omitempty" name:"AnonymousLogin"`

	// 用户名密码登录配置 "TRUE",  "FALSE"
	UsernameLogin *string `json:"UsernameLogin,omitnil,omitempty" name:"UsernameLogin"`
}

type EditAuthConfigRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 手机号登录配置 "TRUE",  "FALSE", "LOGIN_ONLY"
	PhoneNumberLogin *string `json:"PhoneNumberLogin,omitnil,omitempty" name:"PhoneNumberLogin"`

	// 匿名登录配置 "TRUE",  "FALSE"
	AnonymousLogin *string `json:"AnonymousLogin,omitnil,omitempty" name:"AnonymousLogin"`

	// 用户名密码登录配置 "TRUE",  "FALSE"
	UsernameLogin *string `json:"UsernameLogin,omitnil,omitempty" name:"UsernameLogin"`
}

func (r *EditAuthConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditAuthConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "PhoneNumberLogin")
	delete(f, "AnonymousLogin")
	delete(f, "UsernameLogin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EditAuthConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EditAuthConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EditAuthConfigResponse struct {
	*tchttp.BaseResponse
	Response *EditAuthConfigResponseParams `json:"Response"`
}

func (r *EditAuthConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EditAuthConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvInfo struct {
	// 账户下该环境唯一标识
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 环境来源。包含以下取值：
	// <li>miniapp：微信小程序</li>
	// <li>qcloud ：腾讯云</li>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 环境别名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 环境状态。包含以下取值：
	// <li>NORMAL：正常可用</li>
	// <li>UNAVAILABLE：服务不可用，可能是尚未初始化或者初始化过程中</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据库列表
	Databases []*DatabasesInfo `json:"Databases,omitnil,omitempty" name:"Databases"`

	// 存储列表
	Storages []*StorageInfo `json:"Storages,omitnil,omitempty" name:"Storages"`

	// 函数列表
	Functions []*FunctionInfo `json:"Functions,omitnil,omitempty" name:"Functions"`

	// tcb产品套餐ID，参考DescribePackages接口的返回值。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 套餐中文名称，参考DescribePackages接口的返回值。
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 云日志服务列表
	LogServices []*LogServiceInfo `json:"LogServices,omitnil,omitempty" name:"LogServices"`

	// 静态资源信息
	StaticStorages []*StaticStorageInfo `json:"StaticStorages,omitnil,omitempty" name:"StaticStorages"`

	// 是否到期自动降为免费版
	IsAutoDegrade *bool `json:"IsAutoDegrade,omitnil,omitempty" name:"IsAutoDegrade"`

	// 环境渠道
	EnvChannel *string `json:"EnvChannel,omitnil,omitempty" name:"EnvChannel"`

	// 支付方式。包含以下取值：
	// <li> prepayment：预付费</li>
	// <li> postpaid：后付费</li>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否为默认环境
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 环境所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 环境标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自定义日志服务
	CustomLogServices []*ClsInfo `json:"CustomLogServices,omitnil,omitempty" name:"CustomLogServices"`

	// 环境类型：baas, run, hoting, weda
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 是否是dau新套餐
	IsDauPackage *bool `json:"IsDauPackage,omitnil,omitempty" name:"IsDauPackage"`

	// 套餐类型:空\baas\tcbr
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 架构类型
	ArchitectureType *string `json:"ArchitectureType,omitnil,omitempty" name:"ArchitectureType"`

	// 回收标志，默认为空
	Recycle *string `json:"Recycle,omitnil,omitempty" name:"Recycle"`
}

type FunctionInfo struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type IndexAccesses struct {
	// 索引命中次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ops *int64 `json:"Ops,omitnil,omitempty" name:"Ops"`

	// 命中次数从何时开始计数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Since *string `json:"Since,omitnil,omitempty" name:"Since"`
}

type IndexInfo struct {
	// 索引名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 索引大小，单位: 字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 索引键值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*Indexkey `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 索引使用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Accesses *IndexAccesses `json:"Accesses,omitnil,omitempty" name:"Accesses"`

	// 是否为唯一索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unique *bool `json:"Unique,omitnil,omitempty" name:"Unique"`
}

type Indexkey struct {
	// 键名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 方向：specify 1 for ascending or -1 for descending
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type KVPair struct {
	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ListTablesRequestParams struct {
	// 每页返回数量（0-1000)
	MgoLimit *int64 `json:"MgoLimit,omitnil,omitempty" name:"MgoLimit"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 分页偏移量
	MgoOffset *int64 `json:"MgoOffset,omitnil,omitempty" name:"MgoOffset"`

	// 过滤标签数组，用于过滤表名，可选值如：HIDDEN、WEDA、WEDA_SYSTEM
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 模糊搜索查询值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 是否展示隐藏表
	ShowHidden *bool `json:"ShowHidden,omitnil,omitempty" name:"ShowHidden"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// mongo连接器信息
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type ListTablesRequest struct {
	*tchttp.BaseRequest
	
	// 每页返回数量（0-1000)
	MgoLimit *int64 `json:"MgoLimit,omitnil,omitempty" name:"MgoLimit"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 分页偏移量
	MgoOffset *int64 `json:"MgoOffset,omitnil,omitempty" name:"MgoOffset"`

	// 过滤标签数组，用于过滤表名，可选值如：HIDDEN、WEDA、WEDA_SYSTEM
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 模糊搜索查询值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 是否展示隐藏表
	ShowHidden *bool `json:"ShowHidden,omitnil,omitempty" name:"ShowHidden"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// mongo连接器信息
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *ListTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MgoLimit")
	delete(f, "Tag")
	delete(f, "MgoOffset")
	delete(f, "Filters")
	delete(f, "SearchValue")
	delete(f, "ShowHidden")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTablesResponseParams struct {
	// 表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*TableInfo `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pager *Pager `json:"Pager,omitnil,omitempty" name:"Pager"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTablesResponse struct {
	*tchttp.BaseResponse
	Response *ListTablesResponseParams `json:"Response"`
}

func (r *ListTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogObject struct {
	// 日志属于的 topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题的名字
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志时间
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 日志内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 采集路径
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 日志来源设备
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type LogResObject struct {
	// 获取更多检索结果的游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 搜索结果是否已经全部返回
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 日志内容信息
	Results []*LogObject `json:"Results,omitnil,omitempty" name:"Results"`

	// 日志聚合结果
	AnalysisRecords []*string `json:"AnalysisRecords,omitnil,omitempty" name:"AnalysisRecords"`
}

type LogServiceInfo struct {
	// log名
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// log-id
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// topic名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// topic-id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// cls日志所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// topic保存时长 默认7天
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type MgoCommandParam struct {
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 操作类型，可选类型为：UPDATE/QUERY/INSERT/DELETE/COMMAND，本操作必须按实际填写，监控会依赖该字段统计本次操作的类型，并实时减少用户配额，如果填写错误会误扣用户请求配额
	CommandType *string `json:"CommandType,omitnil,omitempty" name:"CommandType"`

	// 待执行命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`
}

type MgoIndexKeys struct {
	// 无
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 无
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type MgoKeySchema struct {
	// 索引字段
	MgoIndexKeys []*MgoIndexKeys `json:"MgoIndexKeys,omitnil,omitempty" name:"MgoIndexKeys"`

	// 是否唯一索引
	MgoIsUnique *bool `json:"MgoIsUnique,omitnil,omitempty" name:"MgoIsUnique"`

	// 是否稀疏索引
	MgoIsSparse *bool `json:"MgoIsSparse,omitnil,omitempty" name:"MgoIsSparse"`
}

// Predefined struct for user
type ModifyCloudBaseGWAPIRequestParams struct {
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// 选项列表，key取值：domain, path。
	Options []*CloudBaseOption `json:"Options,omitnil,omitempty" name:"Options"`
}

type ModifyCloudBaseGWAPIRequest struct {
	*tchttp.BaseRequest
	
	// Service ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// API ID
	APIId *string `json:"APIId,omitnil,omitempty" name:"APIId"`

	// 选项列表，key取值：domain, path。
	Options []*CloudBaseOption `json:"Options,omitnil,omitempty" name:"Options"`
}

func (r *ModifyCloudBaseGWAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudBaseGWAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "APIId")
	delete(f, "Options")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudBaseGWAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudBaseGWAPIResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudBaseGWAPIResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudBaseGWAPIResponseParams `json:"Response"`
}

func (r *ModifyCloudBaseGWAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudBaseGWAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClsTopicRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 日志生命周期，单位天，可取值范围1~3600，取值为3640时代表永久保存
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type ModifyClsTopicRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 日志生命周期，单位天，可取值范围1~3600，取值为3640时代表永久保存
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *ModifyClsTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClsTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClsTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClsTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClsTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClsTopicResponseParams `json:"Response"`
}

func (r *ModifyClsTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClsTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseACLRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`
}

type ModifyDatabaseACLRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`
}

func (r *ModifyDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	delete(f, "AclTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseACLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseACLResponseParams `json:"Response"`
}

func (r *ModifyDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvPlanRequestParams struct {
	// 所需变更套餐的环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 目标套餐Id。
	// 对于云开发环境套餐，可通过 [DescribeBaasPackageList](https://cloud.tencent.com/document/product/876/78167) 接口获取，对应其出参的PackageName
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type ModifyEnvPlanRequest struct {
	*tchttp.BaseRequest
	
	// 所需变更套餐的环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 目标套餐Id。
	// 对于云开发环境套餐，可通过 [DescribeBaasPackageList](https://cloud.tencent.com/document/product/876/78167) 接口获取，对应其出参的PackageName
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

func (r *ModifyEnvPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "PackageId")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEnvPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvPlanResponseParams `json:"Response"`
}

func (r *ModifyEnvPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 环境备注名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type ModifyEnvRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 环境备注名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

func (r *ModifyEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEnvResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvResponseParams `json:"Response"`
}

func (r *ModifyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySafeRuleRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	// <li> CUSTOM：自定义安全规则</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 安全规则内容。
	// 当 AclTag=CUSTOM 时，此参数必填。
	// 详情参考：[文档型数据库安全规则](https://docs.cloudbase.net/database/security-rules)
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type ModifySafeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	// <li> CUSTOM：自定义安全规则</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 安全规则内容。
	// 当 AclTag=CUSTOM 时，此参数必填。
	// 详情参考：[文档型数据库安全规则](https://docs.cloudbase.net/database/security-rules)
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *ModifySafeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySafeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	delete(f, "AclTag")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySafeRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySafeRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySafeRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifySafeRuleResponseParams `json:"Response"`
}

func (r *ModifySafeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySafeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 用户Id, 不做修改
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户名，用户名规则：1. 长度1-64字符 2. 只能包含大小写英文字母、数字和符号 . _ - 3. 只能以字母或数字开头 4. 不能重复，不传该字段或传空字符不修改
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户类型：0-内部用户、1-外部用户，默认0（内部用户），不传该字段或传空字符串不修改
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 密码，传入Uid时密码可不传。密码规则：1. 长度8-32字符（推荐12位以上） 2. 不能以特殊字符开头 3. 至少包含以下四项中的三项：小写字母a-z、大写字母A-Z、数字0-9、特殊字符()!@#$%^&*\|?><_-，不传该字段或传空字符串不修改
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结），默认冻结，不传该字段或传空字符串不修改
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称，长度2-64字符，不传该字段不修改，传空字符修改为空
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，11位数字，不传该字段不修改，传空字符串修改为空
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱地址，不传该字段不修改，传空字符修改为空
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接，可访问的公网URL，不传该字段不修改，传空字符串修改为空
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述，最多200字符，不传该字段不修改，传空字符修改为空
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 用户Id, 不做修改
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户名，用户名规则：1. 长度1-64字符 2. 只能包含大小写英文字母、数字和符号 . _ - 3. 只能以字母或数字开头 4. 不能重复，不传该字段或传空字符不修改
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户类型：0-内部用户、1-外部用户，默认0（内部用户），不传该字段或传空字符串不修改
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 密码，传入Uid时密码可不传。密码规则：1. 长度8-32字符（推荐12位以上） 2. 不能以特殊字符开头 3. 至少包含以下四项中的三项：小写字母a-z、大写字母A-Z、数字0-9、特殊字符()!@#$%^&*\|?><_-，不传该字段或传空字符串不修改
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结），默认冻结，不传该字段或传空字符串不修改
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称，长度2-64字符，不传该字段不修改，传空字符修改为空
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，11位数字，不传该字段不修改，传空字符串修改为空
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱地址，不传该字段不修改，传空字符修改为空
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接，可访问的公网URL，不传该字段不修改，传空字符串修改为空
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述，最多200字符，不传该字段不修改，传空字符修改为空
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Uid")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Password")
	delete(f, "UserStatus")
	delete(f, "NickName")
	delete(f, "Phone")
	delete(f, "Email")
	delete(f, "AvatarUrl")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserResp struct {
	// 是否成功
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`
}

// Predefined struct for user
type ModifyUserResponseParams struct {
	// 修改用户返回值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ModifyUserResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserResponseParams `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MongoConnector struct {
	// 连接器实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MongoDB数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`
}

type MySQLClusterDetail struct {
	// 集群ID
	DbClusterId *string `json:"DbClusterId,omitnil,omitempty" name:"DbClusterId"`

	// 网络详情
	NetInfo *MySQLNetDetail `json:"NetInfo,omitnil,omitempty" name:"NetInfo"`

	// 数据库详情
	DbInfo *ClusterDetail `json:"DbInfo,omitnil,omitempty" name:"DbInfo"`
}

type MySQLNetDetail struct {
	// 内网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateNetAddress *string `json:"PrivateNetAddress,omitnil,omitempty" name:"PrivateNetAddress"`

	// 外网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PubNetAddress *string `json:"PubNetAddress,omitnil,omitempty" name:"PubNetAddress"`

	// 网络信息（VPCID/SubnetID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Net *string `json:"Net,omitnil,omitempty" name:"Net"`

	// 是否开通公网
	PubNetAccessEnabled *bool `json:"PubNetAccessEnabled,omitnil,omitempty" name:"PubNetAccessEnabled"`

	// vpc id 
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网名
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`
}

type MySQLTaskStatus struct {
	// SUCCESS | FAILED | PENDING
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`
}

type Pager struct {
	// 分页偏移量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 文档集合总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type PermissionInfo struct {
	// "READONLY",   //公有读，私有写。所有用户可读，仅创建者及管理员可写  
	// "PRIVATE",    //私有读写，仅创建者及管理员可读写  
	// "ADMINWRITE", //所有用户可读，仅管理员可写  
	// "ADMINONLY",  //仅管理员可操作  
	// "CUSTOM",     // 安全规则
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 自定义规则
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

// Predefined struct for user
type ReinstateEnvRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type ReinstateEnvRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *ReinstateEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReinstateEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReinstateEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReinstateEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReinstateEnvResponse struct {
	*tchttp.BaseResponse
	Response *ReinstateEnvResponseParams `json:"Response"`
}

func (r *ReinstateEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReinstateEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewEnvRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 续费周期，单位：月。
	// 默认值为 1，即续费1个月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type RenewEnvRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 续费周期，单位：月。
	// 默认值为 1，即续费1个月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

func (r *RenewEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewEnvResponse struct {
	*tchttp.BaseResponse
	Response *RenewEnvResponseParams `json:"Response"`
}

func (r *RenewEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunCommandsRequestParams struct {
	// 待执行命令
	MgoCommands []*MgoCommandParam `json:"MgoCommands,omitnil,omitempty" name:"MgoCommands"`

	// 实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Mongo连接器实例信息
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type RunCommandsRequest struct {
	*tchttp.BaseRequest
	
	// 待执行命令
	MgoCommands []*MgoCommandParam `json:"MgoCommands,omitnil,omitempty" name:"MgoCommands"`

	// 实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Mongo连接器实例信息
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *RunCommandsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MgoCommands")
	delete(f, "Tag")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunCommandsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunCommandsResponseParams struct {
	// 返回结果，返回结果为一个json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunCommandsResponse struct {
	*tchttp.BaseResponse
	Response *RunCommandsResponseParams `json:"Response"`
}

func (r *RunCommandsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunSqlRequestParams struct {
	// 要执行的SQL语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 数据库连接器实例信息
	DbInstance *DbInstance `json:"DbInstance,omitnil,omitempty" name:"DbInstance"`

	// 是否只读；当 `true` 时仅允许以 `SELECT/WITH/SHOW/DESCRIBE/DESC/EXPLAIN` 开头的 SQL
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

type RunSqlRequest struct {
	*tchttp.BaseRequest
	
	// 要执行的SQL语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 数据库连接器实例信息
	DbInstance *DbInstance `json:"DbInstance,omitnil,omitempty" name:"DbInstance"`

	// 是否只读；当 `true` 时仅允许以 `SELECT/WITH/SHOW/DESCRIBE/DESC/EXPLAIN` 开头的 SQL
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

func (r *RunSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sql")
	delete(f, "EnvId")
	delete(f, "DbInstance")
	delete(f, "ReadOnly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunSqlResponseParams struct {
	// 查询结果行，每个元素为 JSON 字符串
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 列元数据信息，每个元素为 JSON 字符串，字段包含 `name/databaseType/nullable/length/precision/scale`
	Infos []*string `json:"Infos,omitnil,omitempty" name:"Infos"`

	// 受影响的行数（INSERT/UPDATE/DELETE 等语句）
	RowsAffected *int64 `json:"RowsAffected,omitnil,omitempty" name:"RowsAffected"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunSqlResponse struct {
	*tchttp.BaseResponse
	Response *RunSqlResponseParams `json:"Response"`
}

func (r *RunSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClsLogRequestParams struct {
	// 环境唯一ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 查询起始时间条件
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间条件
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询语句，例如查询云函数：(src:app OR src:system) AND log:\"START RequestId*\"， 查询云数据库：module:database，查询审批流：module:workflow，查询模型：module:model，查询用户权限：module:auth，以上仅为示例语句，实际使用时请根据具体日志内容进行调整，查询语句需严格遵循CLS（Cloud Log Service）语法规范 详细的语法规则请参考官方档：https://cloud.tencent.com/document/product/614/47044
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 是否使用Lucene语法，默认为false
	UseLucene *bool `json:"UseLucene,omitnil,omitempty" name:"UseLucene"`
}

type SearchClsLogRequest struct {
	*tchttp.BaseRequest
	
	// 环境唯一ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 查询起始时间条件
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间条件
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询语句，例如查询云函数：(src:app OR src:system) AND log:\"START RequestId*\"， 查询云数据库：module:database，查询审批流：module:workflow，查询模型：module:model，查询用户权限：module:auth，以上仅为示例语句，实际使用时请根据具体日志内容进行调整，查询语句需严格遵循CLS（Cloud Log Service）语法规范 详细的语法规则请参考官方档：https://cloud.tencent.com/document/product/614/47044
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 是否使用Lucene语法，默认为false
	UseLucene *bool `json:"UseLucene,omitnil,omitempty" name:"UseLucene"`
}

func (r *SearchClsLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryString")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseLucene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClsLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClsLogResponseParams struct {
	// 日志内容结果
	LogResults *LogResObject `json:"LogResults,omitnil,omitempty" name:"LogResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchClsLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchClsLogResponseParams `json:"Response"`
}

func (r *SearchClsLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaticStorageInfo struct {
	// 静态CDN域名
	StaticDomain *string `json:"StaticDomain,omitnil,omitempty" name:"StaticDomain"`

	// 静态CDN默认文件夹，当前为根目录
	DefaultDirName *string `json:"DefaultDirName,omitnil,omitempty" name:"DefaultDirName"`

	// 资源状态(process/online/offline/init)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// cos所属区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// bucket信息
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

type StaticStoreInfo struct {
	// 环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 静态域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdnDomain *string `json:"CdnDomain,omitnil,omitempty" name:"CdnDomain"`

	// COS桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// cos区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Regoin is deprecated.
	Regoin *string `json:"Regoin,omitnil,omitempty" name:"Regoin"`

	// 资源状态:init(初始化)/process(处理中)/online(上线)/destroying(销毁中)/offline(下线))
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type StorageInfo struct {
	// 资源所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 桶名，存储资源的唯一标识
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// cdn 域名
	CdnDomain *string `json:"CdnDomain,omitnil,omitempty" name:"CdnDomain"`

	// 资源所属用户的腾讯云appId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type TableInfo struct {
	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表中文档数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 表的大小（即表中文档总大小），单位：字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 索引数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexCount *int64 `json:"IndexCount,omitnil,omitempty" name:"IndexCount"`

	// 索引占用空间，单位：字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexSize *int64 `json:"IndexSize,omitnil,omitempty" name:"IndexSize"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type UpdateTableRequestParams struct {
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 待删除索引信息
	DropIndexes []*DropIndex `json:"DropIndexes,omitnil,omitempty" name:"DropIndexes"`

	// 待创建索引信息
	CreateIndexes []*CreateIndex `json:"CreateIndexes,omitnil,omitempty" name:"CreateIndexes"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type UpdateTableRequest struct {
	*tchttp.BaseRequest
	
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 待删除索引信息
	DropIndexes []*DropIndex `json:"DropIndexes,omitnil,omitempty" name:"DropIndexes"`

	// 待创建索引信息
	CreateIndexes []*CreateIndex `json:"CreateIndexes,omitnil,omitempty" name:"CreateIndexes"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *UpdateTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "Tag")
	delete(f, "DropIndexes")
	delete(f, "CreateIndexes")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTableResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTableResponseParams `json:"Response"`
}

func (r *UpdateTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {
	// 用户ID
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户类型：internalUser-内部用户、externalUser-外部用户
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结）
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}