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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AuthDomain struct {

	// 域名ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名类型。包含以下取值：
	// <li>SYSTEM</li>
	// <li>USER</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 状态。包含以下取值：
	// <li>ENABLE</li>
	// <li>DISABLE</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CheckTcbServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *CheckTcbServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckTcbServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckTcbServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true表示已开通
		Initialized *bool `json:"Initialized,omitempty" name:"Initialized"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckTcbServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckTcbServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommonServiceAPIRequest struct {
	*tchttp.BaseRequest

	// Service名，需要转发访问的接口名
	Service *string `json:"Service,omitempty" name:"Service"`

	// 需要转发的云API参数，要转成JSON格式
	JSONData *string `json:"JSONData,omitempty" name:"JSONData"`
}

func (r *CommonServiceAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommonServiceAPIRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommonServiceAPIResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// json格式response
		JSONResp *string `json:"JSONResp,omitempty" name:"JSONResp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommonServiceAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommonServiceAPIResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuthDomainRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 安全域名
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`
}

func (r *CreateAuthDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuthDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuthDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuthDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuthDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHostingDomainRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`
}

func (r *CreateHostingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHostingDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHostingDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHostingDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStaticStoreRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *CreateStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStaticStoreRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建静态资源结果(succ/fail)
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStaticStoreResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DatabasesInfo struct {

	// 数据库唯一标识
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 状态。包含以下取值：
	// <li>INITIALIZING：资源初始化中</li>
	// <li>RUNNING：运行中，可正常使用的状态</li>
	// <li>UNUSABLE：禁用，不可用</li>
	// <li>OVERDUE：资源过期</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DeleteEndUserRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 用户列表，每一项都是uuid
	UserList []*string `json:"UserList,omitempty" name:"UserList" list`
}

func (r *DeleteEndUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEndUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEndUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEndUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEndUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthDomainsRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeAuthDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全域名列表列表
		Domains []*AuthDomain `json:"Domains,omitempty" name:"Domains" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseACLRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitempty" name:"CollectionName"`
}

func (r *DescribeDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseACLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
		AclTag *string `json:"AclTag,omitempty" name:"AclTag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseACLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEndUsersRequest struct {
	*tchttp.BaseRequest

	// 开发者的环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 按照 uuid 列表过滤，最大个数为100
	UUIds []*string `json:"UUIds,omitempty" name:"UUIds" list`
}

func (r *DescribeEndUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEndUsersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEndUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 用户列表
		Users []*EndUserInfo `json:"Users,omitempty" name:"Users" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEndUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEndUsersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvFreeQuotaRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 资源类型：可选值：CDN, COS, FLEXDB, HOSTING, SCF
	// 不传则返回全部资源指标
	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes" list`
}

func (r *DescribeEnvFreeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvFreeQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvFreeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 免费抵扣配额详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		QuotaItems []*PostpayEnvQuota `json:"QuotaItems,omitempty" name:"QuotaItems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvFreeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvFreeQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvLimitRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeEnvLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvLimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境总数上限
		MaxEnvNum *int64 `json:"MaxEnvNum,omitempty" name:"MaxEnvNum"`

		// 目前环境总数
		CurrentEnvNum *int64 `json:"CurrentEnvNum,omitempty" name:"CurrentEnvNum"`

		// 免费环境数量上限
		MaxFreeEnvNum *int64 `json:"MaxFreeEnvNum,omitempty" name:"MaxFreeEnvNum"`

		// 目前免费环境数量
		CurrentFreeEnvNum *int64 `json:"CurrentFreeEnvNum,omitempty" name:"CurrentFreeEnvNum"`

		// 总计允许销毁环境次数上限
		MaxDeleteTotal *int64 `json:"MaxDeleteTotal,omitempty" name:"MaxDeleteTotal"`

		// 目前已销毁环境次数
		CurrentDeleteTotal *int64 `json:"CurrentDeleteTotal,omitempty" name:"CurrentDeleteTotal"`

		// 每月允许销毁环境次数上限
		MaxDeleteMonthly *int64 `json:"MaxDeleteMonthly,omitempty" name:"MaxDeleteMonthly"`

		// 本月已销毁环境次数
		CurrentDeleteMonthly *int64 `json:"CurrentDeleteMonthly,omitempty" name:"CurrentDeleteMonthly"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvLimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvsRequest struct {
	*tchttp.BaseRequest

	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitempty" name:"Channels" list`
}

func (r *DescribeEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境信息列表
		EnvList []*EnvInfo `json:"EnvList,omitempty" name:"EnvList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaDataRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

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
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源ID, 目前仅对云函数相关的指标(FunctionInvocationpkg, FunctionGBspkg, FunctionFluxpkg)有意义, 如果想查询某个云函数的指标则在ResourceId中传入函数名; 如果只想查询整个namespace的指标, 则留空或不传.
	ResourceID *string `json:"ResourceID,omitempty" name:"ResourceID"`
}

func (r *DescribeQuotaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标名
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// 指标的值
		Value *int64 `json:"Value,omitempty" name:"Value"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyEnvRequest struct {
	*tchttp.BaseRequest

	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 针对预付费 删除隔离中的环境时要传true 正常环境直接跳过隔离期删除
	IsForce *bool `json:"IsForce,omitempty" name:"IsForce"`
}

func (r *DestroyEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyStaticStoreRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// cdn域名
	CdnDomain *string `json:"CdnDomain,omitempty" name:"CdnDomain"`
}

func (r *DestroyStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyStaticStoreRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 条件任务结果(succ/fail)
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyStaticStoreResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EndUserInfo struct {

	// 用户唯一ID
	UUId *string `json:"UUId,omitempty" name:"UUId"`

	// 微信ID
	WXOpenId *string `json:"WXOpenId,omitempty" name:"WXOpenId"`

	// qq ID
	QQOpenId *string `json:"QQOpenId,omitempty" name:"QQOpenId"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 昵称
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 性别
	Gender *string `json:"Gender,omitempty" name:"Gender"`

	// 头像地址
	AvatarUrl *string `json:"AvatarUrl,omitempty" name:"AvatarUrl"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否为匿名用户
	IsAnonymous *bool `json:"IsAnonymous,omitempty" name:"IsAnonymous"`

	// 是否禁用账户
	IsDisabled *bool `json:"IsDisabled,omitempty" name:"IsDisabled"`
}

type EnvInfo struct {

	// 账户下该环境唯一标识
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 环境来源。包含以下取值：
	// <li>miniapp：微信小程序</li>
	// <li>qcloud ：腾讯云</li>
	Source *string `json:"Source,omitempty" name:"Source"`

	// 环境别名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 环境状态。包含以下取值：
	// <li>NORMAL：正常可用</li>
	// <li>UNAVAILABLE：服务不可用，可能是尚未初始化或者初始化过程中</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 数据库列表
	Databases []*DatabasesInfo `json:"Databases,omitempty" name:"Databases" list`

	// 存储列表
	Storages []*StorageInfo `json:"Storages,omitempty" name:"Storages" list`

	// 函数列表
	Functions []*FunctionInfo `json:"Functions,omitempty" name:"Functions" list`

	// tcb产品套餐ID，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 套餐中文名称，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 云日志服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogServices []*LogServiceInfo `json:"LogServices,omitempty" name:"LogServices" list`

	// 静态资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaticStorages []*StaticStorageInfo `json:"StaticStorages,omitempty" name:"StaticStorages" list`

	// 是否到期自动降为免费版
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAutoDegrade *bool `json:"IsAutoDegrade,omitempty" name:"IsAutoDegrade"`
}

type FunctionInfo struct {

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}

type LogServiceInfo struct {

	// log名
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// log-id
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// topic名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// topic-id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// cls日志所属地域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type ModifyDatabaseACLRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitempty" name:"AclTag"`
}

func (r *ModifyDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDatabaseACLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDatabaseACLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 环境备注名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *ModifyEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PostpayEnvQuota struct {

	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 配额值
	Value *uint64 `json:"Value,omitempty" name:"Value"`

	// 配额生效时间
	// 为空表示没有时间限制
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 配额失效时间
	// 为空表示没有时间限制
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ReinstateEnvRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *ReinstateEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReinstateEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReinstateEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReinstateEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReinstateEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StaticStorageInfo struct {

	// 静态CDN域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaticDomain *string `json:"StaticDomain,omitempty" name:"StaticDomain"`

	// 静态CDN默认文件夹，当前为根目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDirName *string `json:"DefaultDirName,omitempty" name:"DefaultDirName"`

	// 资源状态(process/online/offline/init)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// cos所属区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// bucket信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
}

type StorageInfo struct {

	// 资源所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// 桶名，存储资源的唯一标识
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// cdn 域名
	CdnDomain *string `json:"CdnDomain,omitempty" name:"CdnDomain"`

	// 资源所属用户的腾讯云appId
	AppId *string `json:"AppId,omitempty" name:"AppId"`
}
