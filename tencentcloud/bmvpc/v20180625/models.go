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

package v20180625

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AcceptVpcPeerConnectionRequest struct {
	*tchttp.BaseRequest

	// 黑石对等连接实例ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`
}

func (r *AcceptVpcPeerConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptVpcPeerConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AcceptVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptVpcPeerConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsyncRegisterIpsRequest struct {
	*tchttp.BaseRequest

	// 私有网络的唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网唯一ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 需要注册的IP列表。
	Ips []*string `json:"Ips,omitempty" name:"Ips" list`
}

func (r *AsyncRegisterIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsyncRegisterIpsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsyncRegisterIpsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID。
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AsyncRegisterIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AsyncRegisterIpsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindEipsToNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已分配的EIP列表；AssignedEips和AutoAllocEipNum至少输入一个
	AssignedEips []*string `json:"AssignedEips,omitempty" name:"AssignedEips" list`

	// 新建EIP数目，系统将会按您的要求生产该数目个数EIP；AssignedEips和AutoAllocEipNum至少输入一个
	AutoAllocEipNum *uint64 `json:"AutoAllocEipNum,omitempty" name:"AutoAllocEipNum"`
}

func (r *BindEipsToNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindEipsToNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindEipsToNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindEipsToNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindEipsToNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindIpsToNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 部分IP信息，子网下只有该部分IP将加入NAT，仅当网关转发模式为IP方式有效
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitempty" name:"IpInfoSet" list`
}

func (r *BindIpsToNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindIpsToNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindIpsToNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindIpsToNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindIpsToNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSubnetsToNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表，子网下全部IP将加入NAT，不区分网关转发方式
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`
}

func (r *BindSubnetsToNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSubnetsToNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSubnetsToNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSubnetsToNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSubnetsToNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDockerSubnetWithVlanRequest struct {
	*tchttp.BaseRequest

	// 系统分配的私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息
	SubnetSet []*SubnetCreateInputInfo `json:"SubnetSet,omitempty" name:"SubnetSet" list`
}

func (r *CreateDockerSubnetWithVlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDockerSubnetWithVlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDockerSubnetWithVlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDockerSubnetWithVlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDockerSubnetWithVlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHostedInterfaceRequest struct {
	*tchttp.BaseRequest

	// 托管机器唯一ID 数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 私有网络ID或者私有网络统一ID，建议使用统一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID或者子网统一ID，建议使用统一ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateHostedInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHostedInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHostedInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 黑石托管机器ID
		ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostedInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHostedInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInterfacesRequest struct {
	*tchttp.BaseRequest

	// 物理机实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayRequest struct {
	*tchttp.BaseRequest

	// 转发模式，其中0表示IP方式，1表示网段方式；通过cidr方式可支持更多的IP接入到NAT网关
	ForwardMode *string `json:"ForwardMode,omitempty" name:"ForwardMode"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// NAT名称
	NatName *string `json:"NatName,omitempty" name:"NatName"`

	// 并发连接数规格；取值为1000000、3000000、10000000，分别对应小型、中型、大型NAT网关
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitempty" name:"MaxConcurrent"`

	// 子网ID列表，子网下全部IP将加入NAT，不区分网关转发方式
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 部分IP信息，子网下只有该部分IP将加入NAT，仅当网关转发模式为IP方式有效；IpInfoSet和SubnetIds中的子网ID不能同时存在
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitempty" name:"IpInfoSet" list`

	// 已分配的EIP列表, AssignedEips和AutoAllocEipNum至少输入一个
	AssignedEips []*string `json:"AssignedEips,omitempty" name:"AssignedEips" list`

	// 新建EIP数目，系统将会按您的要求生产该数目个数EIP, AssignedEips和AutoAllocEipNum至少输入一个
	AutoAllocEipNum *uint64 `json:"AutoAllocEipNum,omitempty" name:"AutoAllocEipNum"`

	// 独占标识，取值为0和1，默认值为0；0和1分别表示创建共享型NAT网关和独占NAT型网关；由于同一个VPC网络内，指向NAT集群的默认路由只有一条，因此VPC内只能创建一种类型NAT网关；创建独占型NAT网关时，需联系对应架构师进行独占NAT集群搭建，否则无法创建独占型NAT网关。
	Exclusive *uint64 `json:"Exclusive,omitempty" name:"Exclusive"`
}

func (r *CreateNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoutePoliciesRequest struct {
	*tchttp.BaseRequest

	// 路由表ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 新增的路由
	RoutePolicySet []*RoutePolicy `json:"RoutePolicySet,omitempty" name:"RoutePolicySet" list`
}

func (r *CreateRoutePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoutePoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoutePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoutePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoutePoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetRequest struct {
	*tchttp.BaseRequest

	// 系统分配的私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息
	SubnetSet []*SubnetCreateInputInfo `json:"SubnetSet,omitempty" name:"SubnetSet" list`
}

func (r *CreateSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubnetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubnetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVirtualSubnetWithVlanRequest struct {
	*tchttp.BaseRequest

	// 系统分配的私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息
	SubnetSet []*SubnetCreateInputInfo `json:"SubnetSet,omitempty" name:"SubnetSet" list`
}

func (r *CreateVirtualSubnetWithVlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVirtualSubnetWithVlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVirtualSubnetWithVlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVirtualSubnetWithVlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVirtualSubnetWithVlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcPeerConnectionRequest struct {
	*tchttp.BaseRequest

	// 本端VPC唯一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 对端VPC唯一ID
	PeerVpcId *string `json:"PeerVpcId,omitempty" name:"PeerVpcId"`

	// 对端地域，取值范围为gz,sh,bj,hk,cd,de,sh_bm,gz_bm,bj_bm,cq_bm等
	PeerRegion *string `json:"PeerRegion,omitempty" name:"PeerRegion"`

	// 对等连接名称
	VpcPeerConnectionName *string `json:"VpcPeerConnectionName,omitempty" name:"VpcPeerConnectionName"`

	// 对端账户OwnerUin（默认值为本端账户）
	PeerUin *string `json:"PeerUin,omitempty" name:"PeerUin"`

	// 跨地域必传，带宽上限值
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

func (r *CreateVpcPeerConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcPeerConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcPeerConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcRequest struct {
	*tchttp.BaseRequest

	// 私有网络的名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络的CIDR
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 私有网络的可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 子网信息
	SubnetSet []*VpcSubnetCreateInfo `json:"SubnetSet,omitempty" name:"SubnetSet" list`

	// 是否启用内网监控
	EnableMonitoring *bool `json:"EnableMonitoring,omitempty" name:"EnableMonitoring"`
}

func (r *CreateVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CustomerGateway struct {

	// 用户网关唯一ID
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`

	// 网关名称
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`

	// 公网地址
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// VPN通道引用个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpnConnNum *uint64 `json:"VpnConnNum,omitempty" name:"VpnConnNum"`
}

type DeleteCustomerGatewayRequest struct {
	*tchttp.BaseRequest

	// 对端网关ID，例如：bmcgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
}

func (r *DeleteCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCustomerGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCustomerGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteHostedInterfaceRequest struct {
	*tchttp.BaseRequest

	// 托管机器唯一ID 数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 私有网络ID或者私有网络统一ID，建议使用统一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID或者子网统一ID，建议使用统一ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DeleteHostedInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteHostedInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteHostedInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 黑石托管机器ID
		ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHostedInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteHostedInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteHostedInterfacesRequest struct {
	*tchttp.BaseRequest

	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 物理机ID
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`
}

func (r *DeleteHostedInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteHostedInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteHostedInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHostedInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteHostedInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInterfacesRequest struct {
	*tchttp.BaseRequest

	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网的唯一ID列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`
}

func (r *DeleteInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DeleteNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutePolicyRequest struct {
	*tchttp.BaseRequest

	// 路由表ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表策略ID
	RoutePolicyId *string `json:"RoutePolicyId,omitempty" name:"RoutePolicyId"`
}

func (r *DeleteRoutePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoutePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoutePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoutePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoutePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网实例ID。可通过DescribeSubnets接口返回值中的SubnetId获取。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DeleteSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubnetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubnetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVirtualIpRequest struct {
	*tchttp.BaseRequest

	// 私有网络唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 退还的IP列表。
	Ips []*string `json:"Ips,omitempty" name:"Ips" list`
}

func (r *DeleteVirtualIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVirtualIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVirtualIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVirtualIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVirtualIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcPeerConnectionRequest struct {
	*tchttp.BaseRequest

	// 黑石对等连接实例ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`
}

func (r *DeleteVpcPeerConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcPeerConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcPeerConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DeleteVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID。
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpcResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnConnectionRequest struct {
	*tchttp.BaseRequest

	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
}

func (r *DeleteVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
}

func (r *DeleteVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteVpnGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeregisterIpsRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 注销指定IP的列表
	IpSet []*string `json:"IpSet,omitempty" name:"IpSet" list`

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DeregisterIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeregisterIpsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeregisterIpsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeregisterIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeregisterIpsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewaysRequest struct {
	*tchttp.BaseRequest

	// 对端网关ID，例如：bmcgw-2wqq41m9。每次请求的实例的上限为100。参数不支持同时指定CustomerGatewayIds和Filters。
	CustomerGatewayIds []*string `json:"CustomerGatewayIds,omitempty" name:"CustomerGatewayIds" list`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定CustomerGatewayIds和Filters。
	// <li>customergateway-name - String - （过滤条件）对端网关名称。</li>
	// <li>ip-address - String - （过滤条件)对端网关地址。</li>
	// <li>customergateway-id - String - （过滤条件）对端网关唯一ID。</li>
	// <li>zone - String - （过滤条件）对端所在可用区，形如：ap-guangzhou-2。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCustomerGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 对端网关对象列表
		CustomerGatewaySet []*CustomerGateway `json:"CustomerGatewaySet,omitempty" name:"CustomerGatewaySet" list`

		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomerGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCustomerGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewaysRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// NAT名称
	NatName *string `json:"NatName,omitempty" name:"NatName"`

	// 搜索字段
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 起始值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 偏移值，默认值为 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNatGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewaysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// NAT网关信息列表
		NatGatewayInfoSet []*NatGatewayInfo `json:"NatGatewayInfoSet,omitempty" name:"NatGatewayInfoSet" list`

		// 总数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatGatewaysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSubnetsRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeNatSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatSubnetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// NAT子网信息
		NatSubnetInfoSet []*NatSubnetInfo `json:"NatSubnetInfoSet,omitempty" name:"NatSubnetInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatSubnetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoutePoliciesRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID，例如：rtb-afg8md3c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略实例ID，例如：rti-azd4dt1c。
	RoutePolicyIds []*string `json:"RoutePolicyIds,omitempty" name:"RoutePolicyIds" list`

	// 过滤条件，参数不支持同时指定RoutePolicyIds和Filters。
	// route-table-id - String - （过滤条件）路由表实例ID。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// route-policy-id - String - （过滤条件）路由策略ID。
	// route-policy-description-like - String -（过滤条件）路由项备注。
	// route-policy-type - String - （过滤条件）路由项策略类型。
	// destination-cidr-like - String - （过滤条件）路由项目的地址。
	// gateway-id-like - String - （过滤条件）路由项下一跳网关。
	// gateway-type - String - （过滤条件）路由项下一条网关类型。
	// enable - Bool - （过滤条件）路由策略是否启用。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 初始行的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页行数，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRoutePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoutePoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoutePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由策略数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 路由策略列表
		RoutePolicySet []*RoutePolicy `json:"RoutePolicySet,omitempty" name:"RoutePolicySet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoutePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoutePoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTablesRequest struct {
	*tchttp.BaseRequest

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableIds []*string `json:"RouteTableIds,omitempty" name:"RouteTableIds" list`

	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// route-table-id - String - （过滤条件）路由表实例ID。
	// route-table-name - String - （过滤条件）路由表名称。
	// route-table-id-like - String - （模糊过滤条件）路由表实例ID。
	// route-table-name-like - String - （模糊过滤条件）路由表名称。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// zone - String - （过滤条件）可用区。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 初始行的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页行数，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段, 支持按“RouteTableId”，“VpcId”, "RouteTableName", "CreateTime"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeRouteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 路由表个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 路由表列表
		RouteTableSet []*RouteTable `json:"RouteTableSet,omitempty" name:"RouteTableSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRouteTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetAvailableIpsRequest struct {
	*tchttp.BaseRequest

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CIDR前缀，例如10.0.1
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
}

func (r *DescribeSubnetAvailableIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetAvailableIpsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetAvailableIpsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用IP的范围列表
		IpSet []*string `json:"IpSet,omitempty" name:"IpSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetAvailableIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetAvailableIpsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetByDeviceRequest struct {
	*tchttp.BaseRequest

	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网类型。0: 物理机子网; 7: DOCKER子网 8: 虚拟子网
	Types []*uint64 `json:"Types,omitempty" name:"Types" list`

	// 查询的起始位置。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的个数。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubnetByDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetByDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetByDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子网个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 子网列表
		Data []*SubnetInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetByDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetByDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetByHostedDeviceRequest struct {
	*tchttp.BaseRequest

	// 托管机器ID, 如chm-xasdfx2j
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网类型。0: 物理机子网; 7: DOCKER子网 8: 虚拟子网
	Types []*uint64 `json:"Types,omitempty" name:"Types" list`

	// 查询的起始位置。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的个数。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubnetByHostedDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetByHostedDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetByHostedDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子网个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 子网列表
		Data []*SubnetInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetByHostedDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetByHostedDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest

	// 子网实例ID查询。形如：subnet-pxir56ns。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// subnet-id - String - （过滤条件）Subnet实例名称。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// cidr-block - String - （过滤条件）vpc的cidr。
	// subnet-name - String - （过滤条件）子网名称。
	// zone - String - （过滤条件）可用区。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子网列表信息
		SubnetSet []*SubnetInfo `json:"SubnetSet,omitempty" name:"SubnetSet" list`

		// 返回的子网总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubnetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务状态，其中0表示任务执行成功，1表示任务执行失败，2表示任务正在执行中
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcQuotaRequest struct {
	*tchttp.BaseRequest

	// 类型
	TypeIds []*uint64 `json:"TypeIds,omitempty" name:"TypeIds" list`
}

func (r *DescribeVpcQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额信息
		VpcQuotaSet []*VpcQuota `json:"VpcQuotaSet,omitempty" name:"VpcQuotaSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcResourceRequest struct {
	*tchttp.BaseRequest

	// 私有网络实例ID
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds" list`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// vpc-id - String - （过滤条件）私有网络实例ID，形如：vpc-f49l6u0z。
	// vpc-name - String - （过滤条件）私有网络名称。
	// zone - String - （过滤条件）可用区。
	// state - String - （过滤条件）VPC状态。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpcResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// VPC数据
		VpcResourceSet []*VpcResource `json:"VpcResourceSet,omitempty" name:"VpcResourceSet" list`

		// VPC个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcViewRequest struct {
	*tchttp.BaseRequest

	// 私有网络唯一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeVpcViewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcViewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcViewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// VPC视图信息
		VpcView *VpcViewInfo `json:"VpcView,omitempty" name:"VpcView"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcViewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds" list`

	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// vpc-name - String - （过滤条件）VPC实例名称。
	// vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。
	// cidr-block - String - （过滤条件）vpc的cidr。
	// state - String - （过滤条件）VPC状态。(pending | available).
	// zone -  String - （过滤条件）VPC的可用区。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 初始行的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页行数，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// VPC列表
		VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomerGatewayConfigurationRequest struct {
	*tchttp.BaseRequest

	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`

	// 厂商,取值 h3c，cisco
	VendorName *string `json:"VendorName,omitempty" name:"VendorName"`
}

func (r *DownloadCustomerGatewayConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCustomerGatewayConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCustomerGatewayConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置信息。
		CustomerGatewayConfiguration *string `json:"CustomerGatewayConfiguration,omitempty" name:"CustomerGatewayConfiguration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadCustomerGatewayConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadCustomerGatewayConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type IKEOptionsSpecification struct {

	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBS-192', 'AES-CBC-256', 'DES-CBC'，默认为3DES-CBC
	PropoEncryAlgorithm *string `json:"PropoEncryAlgorithm,omitempty" name:"PropoEncryAlgorithm"`

	// 认证算法：可选值：'MD5', 'SHA1'，默认为MD5
	PropoAuthenAlgorithm *string `json:"PropoAuthenAlgorithm,omitempty" name:"PropoAuthenAlgorithm"`

	// 协商模式：可选值：'AGGRESSIVE', 'MAIN'，默认为MAIN
	ExchangeMode *string `json:"ExchangeMode,omitempty" name:"ExchangeMode"`

	// 本端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS
	LocalIdentity *string `json:"LocalIdentity,omitempty" name:"LocalIdentity"`

	// 对端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS
	RemoteIdentity *string `json:"RemoteIdentity,omitempty" name:"RemoteIdentity"`

	// 本端标识，当LocalIdentity选为ADDRESS时，LocalAddress必填。localAddress默认为vpn网关公网IP
	LocalAddress *string `json:"LocalAddress,omitempty" name:"LocalAddress"`

	// 对端标识，当RemoteIdentity选为ADDRESS时，RemoteAddress必填
	RemoteAddress *string `json:"RemoteAddress,omitempty" name:"RemoteAddress"`

	// 本端标识，当LocalIdentity选为FQDN时，LocalFqdnName必填
	LocalFqdnName *string `json:"LocalFqdnName,omitempty" name:"LocalFqdnName"`

	// 对端标识，当remoteIdentity选为FQDN时，RemoteFqdnName必填
	RemoteFqdnName *string `json:"RemoteFqdnName,omitempty" name:"RemoteFqdnName"`

	// DH group，指定IKE交换密钥时使用的DH组，可选值：'GROUP1', 'GROUP2', 'GROUP5', 'GROUP14', 'GROUP24'，
	DhGroupName *string `json:"DhGroupName,omitempty" name:"DhGroupName"`

	// IKE SA Lifetime，单位：秒，设置IKE SA的生存周期，取值范围：60-604800
	IKESaLifetimeSeconds *uint64 `json:"IKESaLifetimeSeconds,omitempty" name:"IKESaLifetimeSeconds"`

	// IKE版本
	IKEVersion *string `json:"IKEVersion,omitempty" name:"IKEVersion"`
}

type IPSECOptionsSpecification struct {

	// PFS：可选值：'NULL', 'DH-GROUP1', 'DH-GROUP2', 'DH-GROUP5', 'DH-GROUP14', 'DH-GROUP24'，默认为NULL
	PfsDhGroup *string `json:"PfsDhGroup,omitempty" name:"PfsDhGroup"`

	// IPsec SA lifetime(KB)：单位KB，取值范围：2560-604800
	IPSECSaLifetimeTraffic *uint64 `json:"IPSECSaLifetimeTraffic,omitempty" name:"IPSECSaLifetimeTraffic"`

	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBC-192', 'AES-CBC-256', 'DES-CBC', 'NULL'， 默认为AES-CBC-128
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`

	// 认证算法：可选值：'MD5', 'SHA1'，默认为
	IntegrityAlgorith *string `json:"IntegrityAlgorith,omitempty" name:"IntegrityAlgorith"`

	// IPsec SA lifetime(s)：单位秒，取值范围：180-604800
	IPSECSaLifetimeSeconds *uint64 `json:"IPSECSaLifetimeSeconds,omitempty" name:"IPSECSaLifetimeSeconds"`

	// 安全协议，默认为ESP
	SecurityProto *string `json:"SecurityProto,omitempty" name:"SecurityProto"`
}

type IpInfo struct {

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// IP列表
	Ips []*string `json:"Ips,omitempty" name:"Ips" list`
}

type ModifyCustomerGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// 对端网关ID，例如：bmcgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`

	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`
}

func (r *ModifyCustomerGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomerGatewayAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomerGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCustomerGatewayAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoutePolicyRequest struct {
	*tchttp.BaseRequest

	// 路由表ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 修改的路由
	RoutePolicy *RoutePolicy `json:"RoutePolicy,omitempty" name:"RoutePolicy"`
}

func (r *ModifyRoutePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoutePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRoutePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoutePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRoutePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableRequest struct {
	*tchttp.BaseRequest

	// 路由表ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *ModifyRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyRouteTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetAttributeRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
}

func (r *ModifySubnetAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubnetAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubnetAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubnetAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetDHCPRelayRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否开启DHCP Relay
	EnableDHCP *bool `json:"EnableDHCP,omitempty" name:"EnableDHCP"`

	// DHCP服务器IP
	ServerIps []*string `json:"ServerIps,omitempty" name:"ServerIps" list`

	// 预留IP个数
	ReservedIpCount *uint64 `json:"ReservedIpCount,omitempty" name:"ReservedIpCount"`
}

func (r *ModifySubnetDHCPRelayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubnetDHCPRelayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetDHCPRelayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubnetDHCPRelayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubnetDHCPRelayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAttributeRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 是否开启内网监控，0为关闭，1为开启
	EnableMonitor *bool `json:"EnableMonitor,omitempty" name:"EnableMonitor"`
}

func (r *ModifyVpcAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcPeerConnectionRequest struct {
	*tchttp.BaseRequest

	// 黑石对等连接唯一ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`

	// 对等连接带宽
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 对等连接名称
	VpcPeerConnectionName *string `json:"VpcPeerConnectionName,omitempty" name:"VpcPeerConnectionName"`
}

func (r *ModifyVpcPeerConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcPeerConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpcPeerConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnConnectionAttributeRequest struct {
	*tchttp.BaseRequest

	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`

	// VPC实例ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPN通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`

	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey,omitempty" name:"PreShareKey"`

	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitempty" name:"SecurityPolicyDatabases" list`

	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自我保护机制，用户配置网络安全协议。
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`

	// IPSec配置，腾讯云提供IPSec安全会话设置。
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`
}

func (r *ModifyVpnConnectionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnConnectionAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnConnectionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpnConnectionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnConnectionAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayAttributeRequest struct {
	*tchttp.BaseRequest

	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`
}

func (r *ModifyVpnGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnGatewayAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpnGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyVpnGatewayAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NatGatewayInfo struct {

	// NAT网关ID
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 网关名称
	NatName *string `json:"NatName,omitempty" name:"NatName"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 网关状态，其中0表示创建中，1表示运行中，2表示创建失败
	ProductionStatus *uint64 `json:"ProductionStatus,omitempty" name:"ProductionStatus"`

	// EIP列表
	Eips []*string `json:"Eips,omitempty" name:"Eips" list`

	// 并发连接数规格，取值为1000000, 3000000, 10000000
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitempty" name:"MaxConcurrent"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 独占标识，其中0表示共享，1表示独占，默认值为0
	Exclusive *uint64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// 转发模式，其中0表示IP方式，1表示网段方式
	ForwardMode *uint64 `json:"ForwardMode,omitempty" name:"ForwardMode"`

	// 私有网络网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 网关类型，取值为 small，middle，big，分别对应小型、中型、大型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type NatSubnetInfo struct {

	// 子网名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// NAT子网类型，其中0表示绑定部分IP的NAT子网，1表示绑定全部IP的NAT子网，2表示绑定网关方式的NAT子网
	SubnetNatType *uint64 `json:"SubnetNatType,omitempty" name:"SubnetNatType"`

	// 子网网段
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
}

type RejectVpcPeerConnectionRequest struct {
	*tchttp.BaseRequest

	// 黑石对等连接实例ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`
}

func (r *RejectVpcPeerConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectVpcPeerConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RejectVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectVpcPeerConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnConnectionRequest struct {
	*tchttp.BaseRequest

	// VPC唯一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
}

func (r *ResetVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnConnectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetVpnConnectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RoutePolicy struct {

	// 目的网段
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// 下一跳类型，目前我们支持的类型有：
	// LOCAL：物理机默认路由；
	// VPN：VPN网关；
	// PEERCONNECTION：对等连接；
	// CPM：物理机自定义路由；
	// CCN：云联网；
	// TGW：公网默认路由；
	// SSLVPN : SSH SSL VPN网关。
	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`

	// 下一跳地址，这里只需要指定不同下一跳类型的网关ID，系统会自动匹配到下一跳地址。
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// 路由策略描述。
	RouteDescription *string `json:"RouteDescription,omitempty" name:"RouteDescription"`

	// 路由策略ID
	RoutePolicyId *string `json:"RoutePolicyId,omitempty" name:"RoutePolicyId"`

	// 路由类型，目前我们支持的类型有：
	// USER：用户自定义路由；
	// NETD：网络探测路由，创建网络探测实例时，系统默认下发，不可编辑与删除；
	// CCN：云联网路由，系统默认下发，不可编辑与删除。
	// 用户只能添加和编辑USER 类型的路由。
	RoutePolicyType *string `json:"RoutePolicyType,omitempty" name:"RoutePolicyType"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RouteTable struct {

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC的名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPC的CIDR
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type SecurityPolicyDatabase struct {

	// 本端网段
	LocalCidrBlock *string `json:"LocalCidrBlock,omitempty" name:"LocalCidrBlock"`

	// 对端网段
	RemoteCidrBlock []*string `json:"RemoteCidrBlock,omitempty" name:"RemoteCidrBlock" list`
}

type SubnetCreateInputInfo struct {

	// 子网名称，可任意命名，但不得超过60个字符
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网网段，子网网段必须在VPC网段内，相同VPC内子网网段不能重叠
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 是否开启子网分布式网关，默认传1，传0为关闭子网分布式网关。关闭分布式网关子网用于云服务器化子网，此子网中只能有一台物理机，同时此物理机及其上子机只能在此子网中
	DistributedFlag *uint64 `json:"DistributedFlag,omitempty" name:"DistributedFlag"`

	// 是否开启dhcp relay ，关闭为0，开启为1。默认为0
	DhcpEnable *uint64 `json:"DhcpEnable,omitempty" name:"DhcpEnable"`

	// DHCP SERVER 的IP地址数组。IP地址为相同VPC的子网内分配的IP
	DhcpServerIp []*string `json:"DhcpServerIp,omitempty" name:"DhcpServerIp" list`

	// 预留的IP个数。从该子网的最大可分配IP倒序分配N个IP 用于DHCP 动态分配使用的地址段
	IpReserve *uint64 `json:"IpReserve,omitempty" name:"IpReserve"`

	// 子网绑定的vlanId。VlanId取值范围为2000-2999。创建物理机子网，VlanId默认为5; 创建docker子网或者虚拟子网，VlanId默认会分配2000--2999未使用的数值。
	VlanId *uint64 `json:"VlanId,omitempty" name:"VlanId"`

	// 黑石子网的可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 是否25G子网，1为是，0为否。
	IsSmartNic *uint64 `json:"IsSmartNic,omitempty" name:"IsSmartNic"`
}

type SubnetInfo struct {

	// 私有网络的唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC的名称。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPC的CIDR。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 私有网络的唯一ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名称。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网CIDR。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 子网类型。0: 黑石物理机子网; 6: ccs子网; 7 Docker子网; 8: 虚拟机子网
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 可用区ID。
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子网物理机的个数
	CpmNum *uint64 `json:"CpmNum,omitempty" name:"CpmNum"`

	// 子网的VlanId。
	VlanId *uint64 `json:"VlanId,omitempty" name:"VlanId"`

	// 是否开启分布式网关 ，关闭为0，开启为1。
	DistributedFlag *uint64 `json:"DistributedFlag,omitempty" name:"DistributedFlag"`

	// 是否开启dhcp relay ，关闭为0，开启为1。默认为0。
	DhcpEnable *uint64 `json:"DhcpEnable,omitempty" name:"DhcpEnable"`

	// DHCP SERVER 的IP地址数组。IP地址为相同VPC的子网内分配的IP。
	DhcpServerIp []*string `json:"DhcpServerIp,omitempty" name:"DhcpServerIp" list`

	// 预留的IP个数。从该子网的最大可分配IP倒序分配N个IP 用于DHCP 动态分配使用的地址段。
	IpReserve *uint64 `json:"IpReserve,omitempty" name:"IpReserve"`

	// 子网中可用的IP个数
	AvailableIpNum *uint64 `json:"AvailableIpNum,omitempty" name:"AvailableIpNum"`

	// 子网中总共的IP个数
	TotalIpNum *uint64 `json:"TotalIpNum,omitempty" name:"TotalIpNum"`

	// 子网创建时间
	SubnetCreateTime *string `json:"SubnetCreateTime,omitempty" name:"SubnetCreateTime"`

	// 25G子网标识
	IsSmartNic *uint64 `json:"IsSmartNic,omitempty" name:"IsSmartNic"`
}

type UnbindEipsFromNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已分配的EIP列表
	AssignedEips []*string `json:"AssignedEips,omitempty" name:"AssignedEips" list`
}

func (r *UnbindEipsFromNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindEipsFromNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindEipsFromNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindEipsFromNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindEipsFromNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindIpsFromNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 部分IP信息；子网须以部分IP将加入NAT网关
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitempty" name:"IpInfoSet" list`
}

func (r *UnbindIpsFromNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindIpsFromNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindIpsFromNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindIpsFromNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindIpsFromNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindSubnetsFromNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表，子网不区分加入NAT网关的转发方式
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`
}

func (r *UnbindSubnetsFromNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindSubnetsFromNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindSubnetsFromNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindSubnetsFromNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindSubnetsFromNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeNatGatewayRequest struct {
	*tchttp.BaseRequest

	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 并发连接数规格；取值为1000000、3000000、10000000，分别对应小型、中型、大型NAT网关
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitempty" name:"MaxConcurrent"`
}

func (r *UpgradeNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeNatGatewayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeNatGatewayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VpcInfo struct {

	// 私有网络的唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC的名称。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPC的CIDR。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC状态
	State *string `json:"State,omitempty" name:"State"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type VpcQuota struct {

	// 配额类型ID
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// 配额
	Quota *uint64 `json:"Quota,omitempty" name:"Quota"`
}

type VpcResource struct {

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络的CIDR
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 子网个数
	SubnetNum *uint64 `json:"SubnetNum,omitempty" name:"SubnetNum"`

	// NAT个数
	NatNum *uint64 `json:"NatNum,omitempty" name:"NatNum"`

	// VPC状态
	State *string `json:"State,omitempty" name:"State"`

	// 是否开启监控
	MonitorFlag *bool `json:"MonitorFlag,omitempty" name:"MonitorFlag"`

	// 物理机个数
	CpmNum *uint64 `json:"CpmNum,omitempty" name:"CpmNum"`

	// 可用IP个数
	LeaveIpNum *uint64 `json:"LeaveIpNum,omitempty" name:"LeaveIpNum"`

	// 负载均衡个数
	LbNum *uint64 `json:"LbNum,omitempty" name:"LbNum"`

	// 流量镜像网关个数
	TrafficMirrorNum *uint64 `json:"TrafficMirrorNum,omitempty" name:"TrafficMirrorNum"`

	// 弹性IP个数
	EipNum *uint64 `json:"EipNum,omitempty" name:"EipNum"`

	// 专线网关个数
	PlgwNum *uint64 `json:"PlgwNum,omitempty" name:"PlgwNum"`

	// 专线通道个数
	PlvpNum *uint64 `json:"PlvpNum,omitempty" name:"PlvpNum"`

	// ssl vpn网关个数
	SslVpnGwNum *uint64 `json:"SslVpnGwNum,omitempty" name:"SslVpnGwNum"`

	// 对等链接个数
	VpcPeerNum *uint64 `json:"VpcPeerNum,omitempty" name:"VpcPeerNum"`

	// ipsec vpn网关个数
	IpsecVpnGwNum *uint64 `json:"IpsecVpnGwNum,omitempty" name:"IpsecVpnGwNum"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否老专区VPC
	IsOld *bool `json:"IsOld,omitempty" name:"IsOld"`
}

type VpcSubnetCreateInfo struct {

	// 子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网的CIDR
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 子网的可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type VpcSubnetViewInfo struct {

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网CIDR
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 子网下设备个数
	CpmNum *uint64 `json:"CpmNum,omitempty" name:"CpmNum"`

	// 内网负载均衡个数
	LbNum *uint64 `json:"LbNum,omitempty" name:"LbNum"`

	// 子网所在可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type VpcViewInfo struct {

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络CIDR
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 私有网络所在可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 外网负载均衡个数
	LbNum *uint64 `json:"LbNum,omitempty" name:"LbNum"`

	// 弹性公网IP个数
	EipNum *uint64 `json:"EipNum,omitempty" name:"EipNum"`

	// NAT网关个数
	NatNum *uint64 `json:"NatNum,omitempty" name:"NatNum"`

	// 子网列表
	SubnetSet []*VpcSubnetViewInfo `json:"SubnetSet,omitempty" name:"SubnetSet" list`
}
