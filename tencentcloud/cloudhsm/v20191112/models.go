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

package v20191112

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeHSMBySubnetIdRequest struct {
	*tchttp.BaseRequest

	// Subnet标识符
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DescribeHSMBySubnetIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHSMBySubnetIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHSMBySubnetIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// HSM数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 作为查询条件的SubnetId
		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHSMBySubnetIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHSMBySubnetIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHSMByVpcIdRequest struct {
	*tchttp.BaseRequest

	// VPC标识符
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeHSMByVpcIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHSMByVpcIdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHSMByVpcIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// HSM数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 作为查询条件的VpcId
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHSMByVpcIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHSMByVpcIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetRequest struct {
	*tchttp.BaseRequest

	// 返回数量。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询指定VpcId下的子网信息。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 查找关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的子网数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的子网实例列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetList []*Subnet `json:"SubnetList,omitempty" name:"SubnetList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsgRequest struct {
	*tchttp.BaseRequest

	// 偏移量，当Offset和Limit均为0时将一次性返回用户所有的安全组列表。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回量，当Offset和Limit均为0时将一次性返回用户所有的安全组列表。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeUsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户的安全组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		SgList []*SgUnit `json:"SgList,omitempty" name:"SgList" list`

		// 返回的安全组数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsgRuleRequest struct {
	*tchttp.BaseRequest

	// 根据安全组Id获取安全组详情
	SgIds []*string `json:"SgIds,omitempty" name:"SgIds" list`
}

func (r *DescribeUsgRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsgRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsgRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		SgRules []*UsgRuleDetail `json:"SgRules,omitempty" name:"SgRules" list`

		// 安全组详情数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsgRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsgRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRequest struct {
	*tchttp.BaseRequest

	// 返回偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可查询到的所有Vpc实例总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Vpc对象列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		VpcList []*Vpc `json:"VpcList,omitempty" name:"VpcList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVsmAttributesRequest struct {
	*tchttp.BaseRequest

	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DescribeVsmAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVsmAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVsmAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源Id
		ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

		// 资源名称
		ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

		// 资源状态
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 资源IP
		Vip *string `json:"Vip,omitempty" name:"Vip"`

		// 资源所属Vpc
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// 资源所属子网
		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

		// 资源所属HSM的规格
		Model *string `json:"Model,omitempty" name:"Model"`

		// 资源类型
		VsmType *int64 `json:"VsmType,omitempty" name:"VsmType"`

		// 地域Id
		RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

		// 区域Id
		ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

		// 过期时间
		ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

		// 安全组详情信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		SgList []*UsgRuleDetail `json:"SgList,omitempty" name:"SgList" list`

		// 子网名
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

		// 地域名
	// 注意：此字段可能返回 null，表示取不到有效值。
		RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

		// 区域名
	// 注意：此字段可能返回 null，表示取不到有效值。
		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

		// 实例是否已经过期
	// 注意：此字段可能返回 null，表示取不到有效值。
		Expired *bool `json:"Expired,omitempty" name:"Expired"`

		// 为正数表示实例距离过期时间剩余秒数，为负数表示实例已经过期多少秒
	// 注意：此字段可能返回 null，表示取不到有效值。
		RemainSeconds *int64 `json:"RemainSeconds,omitempty" name:"RemainSeconds"`

		// 私有虚拟网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

		// VPC的IPv4 CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
		VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

		// 子网的CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetCidrBlock *string `json:"SubnetCidrBlock,omitempty" name:"SubnetCidrBlock"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVsmAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVsmAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVsmsRequest struct {
	*tchttp.BaseRequest

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeVsmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVsmsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVsmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取实例的总个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		VsmList []*ResourceInfo `json:"VsmList,omitempty" name:"VsmList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVsmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVsmsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceBuyVsmRequest struct {
	*tchttp.BaseRequest

	// 需购买实例的数量
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 付费模式：0表示按需计费/后付费，1表示预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 商品的时间大小
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 商品的时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 货币类型，默认为CNY
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 默认为CREATE，可选RENEW
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *InquiryPriceBuyVsmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceBuyVsmRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceBuyVsmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总金额
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`

		// 购买的实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

		// 商品的时间大小
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// 商品的时间单位
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// 原始总金额
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceBuyVsmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceBuyVsmResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVsmAttributesRequest struct {
	*tchttp.BaseRequest

	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// UpdateResourceName-修改资源名称,
	// UpdateSgIds-修改安全组名称,
	// UpdateNetWork-修改网络,
	// Default-默认不修改
	Type []*string `json:"Type,omitempty" name:"Type" list`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 安全组Id
	SgIds []*string `json:"SgIds,omitempty" name:"SgIds" list`

	// VpcId
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *ModifyVsmAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVsmAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVsmAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVsmAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVsmAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResourceInfo struct {

	// 资源Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 资源IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 资源所属Vpc
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 资源所属子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 资源所属HSM规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *string `json:"Model,omitempty" name:"Model"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VsmType *int64 `json:"VsmType,omitempty" name:"VsmType"`

	// 地域Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 区域Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 地域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 区域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 实例的安全组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgList []*SgUnit `json:"SgList,omitempty" name:"SgList" list`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 当前实例是否已经过期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expired *bool `json:"Expired,omitempty" name:"Expired"`

	// 为正数表示实例距离过期时间还剩余多少秒，为负数表示已经过期多少秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainSeconds *int64 `json:"RemainSeconds,omitempty" name:"RemainSeconds"`

	// Vpc名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type SgUnit struct {

	// 安全组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 安全组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgName *string `json:"SgName,omitempty" name:"SgName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgRemark *string `json:"SgRemark,omitempty" name:"SgRemark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Subnet struct {

	// VPC实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网实例ID，例如：subnet-bthucmmy。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网的 IPv4 CIDR。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 可用IP数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" name:"AvailableIpAddressCount"`

	// 子网的 IPv6 CIDR。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// 总IP数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalIpAddressCount *int64 `json:"TotalIpAddressCount,omitempty" name:"TotalIpAddressCount"`

	// 是否为默认Subnet
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type UsgPolicy struct {

	// cidr格式地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 安全组id代表的地址集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 地址组id代表的地址集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 服务组id代表的协议和端口集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 匹配后行为:ACCEPT/DROP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`
}

type UsgRuleDetail struct {

	// 入站规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	InBound []*UsgPolicy `json:"InBound,omitempty" name:"InBound" list`

	// 出站规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBound []*UsgPolicy `json:"OutBound,omitempty" name:"OutBound" list`

	// 安全组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 安全组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgName *string `json:"SgName,omitempty" name:"SgName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgRemark *string `json:"SgRemark,omitempty" name:"SgRemark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type Vpc struct {

	// Vpc名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 是否为默认VPC
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}
