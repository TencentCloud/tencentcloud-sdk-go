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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AcceptVpcPeerConnectionRequestParams struct {
	// 黑石对等连接实例ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`
}

type AcceptVpcPeerConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 黑石对等连接实例ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`
}

func (r *AcceptVpcPeerConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptVpcPeerConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcPeerConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcceptVpcPeerConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptVpcPeerConnectionResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AcceptVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *AcceptVpcPeerConnectionResponseParams `json:"Response"`
}

func (r *AcceptVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptVpcPeerConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsyncRegisterIpsRequestParams struct {
	// 私有网络的唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网唯一ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 需要注册的IP列表。
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

type AsyncRegisterIpsRequest struct {
	*tchttp.BaseRequest
	
	// 私有网络的唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网唯一ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 需要注册的IP列表。
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

func (r *AsyncRegisterIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsyncRegisterIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Ips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AsyncRegisterIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AsyncRegisterIpsResponseParams struct {
	// 任务ID。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AsyncRegisterIpsResponse struct {
	*tchttp.BaseResponse
	Response *AsyncRegisterIpsResponseParams `json:"Response"`
}

func (r *AsyncRegisterIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AsyncRegisterIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindEipsToNatGatewayRequestParams struct {
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已分配的EIP列表；AssignedEips和AutoAllocEipNum至少输入一个
	AssignedEips []*string `json:"AssignedEips,omitempty" name:"AssignedEips"`

	// 新建EIP数目，系统将会按您的要求生产该数目个数EIP；AssignedEips和AutoAllocEipNum至少输入一个
	AutoAllocEipNum *uint64 `json:"AutoAllocEipNum,omitempty" name:"AutoAllocEipNum"`
}

type BindEipsToNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已分配的EIP列表；AssignedEips和AutoAllocEipNum至少输入一个
	AssignedEips []*string `json:"AssignedEips,omitempty" name:"AssignedEips"`

	// 新建EIP数目，系统将会按您的要求生产该数目个数EIP；AssignedEips和AutoAllocEipNum至少输入一个
	AutoAllocEipNum *uint64 `json:"AutoAllocEipNum,omitempty" name:"AutoAllocEipNum"`
}

func (r *BindEipsToNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEipsToNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "VpcId")
	delete(f, "AssignedEips")
	delete(f, "AutoAllocEipNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindEipsToNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindEipsToNatGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindEipsToNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *BindEipsToNatGatewayResponseParams `json:"Response"`
}

func (r *BindEipsToNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEipsToNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindIpsToNatGatewayRequestParams struct {
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 部分IP信息，子网下只有该部分IP将加入NAT，仅当网关转发模式为IP方式有效
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitempty" name:"IpInfoSet"`
}

type BindIpsToNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 部分IP信息，子网下只有该部分IP将加入NAT，仅当网关转发模式为IP方式有效
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitempty" name:"IpInfoSet"`
}

func (r *BindIpsToNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindIpsToNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "VpcId")
	delete(f, "IpInfoSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindIpsToNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindIpsToNatGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindIpsToNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *BindIpsToNatGatewayResponseParams `json:"Response"`
}

func (r *BindIpsToNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindIpsToNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSubnetsToNatGatewayRequestParams struct {
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表，子网下全部IP将加入NAT，不区分网关转发方式
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

type BindSubnetsToNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表，子网下全部IP将加入NAT，不区分网关转发方式
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

func (r *BindSubnetsToNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSubnetsToNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "VpcId")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindSubnetsToNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSubnetsToNatGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindSubnetsToNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *BindSubnetsToNatGatewayResponseParams `json:"Response"`
}

func (r *BindSubnetsToNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSubnetsToNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomerGatewayRequestParams struct {
	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`

	// 对端网关公网IP。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 可用区ID
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type CreateCustomerGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`

	// 对端网关公网IP。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 可用区ID
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *CreateCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomerGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayName")
	delete(f, "IpAddress")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomerGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomerGatewayResponseParams struct {
	// 对端网关对象
	CustomerGateway *CustomerGateway `json:"CustomerGateway,omitempty" name:"CustomerGateway"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomerGatewayResponseParams `json:"Response"`
}

func (r *CreateCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDockerSubnetWithVlanRequestParams struct {
	// 系统分配的私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息
	SubnetSet []*SubnetCreateInputInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

type CreateDockerSubnetWithVlanRequest struct {
	*tchttp.BaseRequest
	
	// 系统分配的私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息
	SubnetSet []*SubnetCreateInputInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

func (r *CreateDockerSubnetWithVlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDockerSubnetWithVlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDockerSubnetWithVlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDockerSubnetWithVlanResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDockerSubnetWithVlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateDockerSubnetWithVlanResponseParams `json:"Response"`
}

func (r *CreateDockerSubnetWithVlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDockerSubnetWithVlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostedInterfaceRequestParams struct {
	// 托管机器唯一ID 数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 私有网络ID或者私有网络统一ID，建议使用统一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID或者子网统一ID，建议使用统一ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type CreateHostedInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// 托管机器唯一ID 数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 私有网络ID或者私有网络统一ID，建议使用统一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID或者子网统一ID，建议使用统一ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateHostedInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostedInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHostedInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostedInterfaceResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 黑石托管机器ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHostedInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateHostedInterfaceResponseParams `json:"Response"`
}

func (r *CreateHostedInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostedInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInterfacesRequestParams struct {
	// 物理机实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type CreateInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// 物理机实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInterfacesResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *CreateInterfacesResponseParams `json:"Response"`
}

func (r *CreateInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayRequestParams struct {
	// 转发模式，其中0表示IP方式，1表示网段方式；通过cidr方式可支持更多的IP接入到NAT网关
	ForwardMode *string `json:"ForwardMode,omitempty" name:"ForwardMode"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// NAT名称
	NatName *string `json:"NatName,omitempty" name:"NatName"`

	// 并发连接数规格；取值为1000000、3000000、10000000，分别对应小型、中型、大型NAT网关
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitempty" name:"MaxConcurrent"`

	// 子网ID列表，子网下全部IP将加入NAT，不区分网关转发方式
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 部分IP信息，子网下只有该部分IP将加入NAT，仅当网关转发模式为IP方式有效；IpInfoSet和SubnetIds中的子网ID不能同时存在
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitempty" name:"IpInfoSet"`

	// 已分配的EIP列表, AssignedEips和AutoAllocEipNum至少输入一个
	AssignedEips []*string `json:"AssignedEips,omitempty" name:"AssignedEips"`

	// 新建EIP数目，系统将会按您的要求生产该数目个数EIP, AssignedEips和AutoAllocEipNum至少输入一个
	AutoAllocEipNum *uint64 `json:"AutoAllocEipNum,omitempty" name:"AutoAllocEipNum"`

	// 独占标识，取值为0和1，默认值为0；0和1分别表示创建共享型NAT网关和独占NAT型网关；由于同一个VPC网络内，指向NAT集群的默认路由只有一条，因此VPC内只能创建一种类型NAT网关；创建独占型NAT网关时，需联系对应架构师进行独占NAT集群搭建，否则无法创建独占型NAT网关。
	Exclusive *uint64 `json:"Exclusive,omitempty" name:"Exclusive"`
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
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 部分IP信息，子网下只有该部分IP将加入NAT，仅当网关转发模式为IP方式有效；IpInfoSet和SubnetIds中的子网ID不能同时存在
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitempty" name:"IpInfoSet"`

	// 已分配的EIP列表, AssignedEips和AutoAllocEipNum至少输入一个
	AssignedEips []*string `json:"AssignedEips,omitempty" name:"AssignedEips"`

	// 新建EIP数目，系统将会按您的要求生产该数目个数EIP, AssignedEips和AutoAllocEipNum至少输入一个
	AutoAllocEipNum *uint64 `json:"AutoAllocEipNum,omitempty" name:"AutoAllocEipNum"`

	// 独占标识，取值为0和1，默认值为0；0和1分别表示创建共享型NAT网关和独占NAT型网关；由于同一个VPC网络内，指向NAT集群的默认路由只有一条，因此VPC内只能创建一种类型NAT网关；创建独占型NAT网关时，需联系对应架构师进行独占NAT集群搭建，否则无法创建独占型NAT网关。
	Exclusive *uint64 `json:"Exclusive,omitempty" name:"Exclusive"`
}

func (r *CreateNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ForwardMode")
	delete(f, "VpcId")
	delete(f, "NatName")
	delete(f, "MaxConcurrent")
	delete(f, "SubnetIds")
	delete(f, "IpInfoSet")
	delete(f, "AssignedEips")
	delete(f, "AutoAllocEipNum")
	delete(f, "Exclusive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatGatewayResponseParams `json:"Response"`
}

func (r *CreateNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoutePoliciesRequestParams struct {
	// 路由表ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 新增的路由
	RoutePolicySet []*RoutePolicy `json:"RoutePolicySet,omitempty" name:"RoutePolicySet"`
}

type CreateRoutePoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 新增的路由
	RoutePolicySet []*RoutePolicy `json:"RoutePolicySet,omitempty" name:"RoutePolicySet"`
}

func (r *CreateRoutePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoutePoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RoutePolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoutePoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoutePoliciesResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRoutePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoutePoliciesResponseParams `json:"Response"`
}

func (r *CreateRoutePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoutePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetRequestParams struct {
	// 系统分配的私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息
	SubnetSet []*SubnetCreateInputInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

type CreateSubnetRequest struct {
	*tchttp.BaseRequest
	
	// 系统分配的私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息
	SubnetSet []*SubnetCreateInputInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

func (r *CreateSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSubnetResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubnetResponseParams `json:"Response"`
}

func (r *CreateSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVirtualSubnetWithVlanRequestParams struct {
	// 系统分配的私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息
	SubnetSet []*SubnetCreateInputInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

type CreateVirtualSubnetWithVlanRequest struct {
	*tchttp.BaseRequest
	
	// 系统分配的私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息
	SubnetSet []*SubnetCreateInputInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

func (r *CreateVirtualSubnetWithVlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVirtualSubnetWithVlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVirtualSubnetWithVlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVirtualSubnetWithVlanResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVirtualSubnetWithVlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateVirtualSubnetWithVlanResponseParams `json:"Response"`
}

func (r *CreateVirtualSubnetWithVlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVirtualSubnetWithVlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcPeerConnectionRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcPeerConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "PeerVpcId")
	delete(f, "PeerRegion")
	delete(f, "VpcPeerConnectionName")
	delete(f, "PeerUin")
	delete(f, "Bandwidth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcPeerConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcPeerConnectionResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcPeerConnectionResponseParams `json:"Response"`
}

func (r *CreateVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcPeerConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcRequestParams struct {
	// 私有网络的名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络的CIDR
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 私有网络的可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 子网信息
	SubnetSet []*VpcSubnetCreateInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`

	// 是否启用内网监控
	EnableMonitoring *bool `json:"EnableMonitoring,omitempty" name:"EnableMonitoring"`
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
	SubnetSet []*VpcSubnetCreateInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`

	// 是否启用内网监控
	EnableMonitoring *bool `json:"EnableMonitoring,omitempty" name:"EnableMonitoring"`
}

func (r *CreateVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcName")
	delete(f, "CidrBlock")
	delete(f, "Zone")
	delete(f, "SubnetSet")
	delete(f, "EnableMonitoring")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVpcResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcResponseParams `json:"Response"`
}

func (r *CreateVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// Predefined struct for user
type DeleteCustomerGatewayRequestParams struct {
	// 对端网关ID，例如：bmcgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomerGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomerGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomerGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomerGatewayResponseParams `json:"Response"`
}

func (r *DeleteCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHostedInterfaceRequestParams struct {
	// 托管机器唯一ID 数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 私有网络ID或者私有网络统一ID，建议使用统一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID或者子网统一ID，建议使用统一ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DeleteHostedInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// 托管机器唯一ID 数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 私有网络ID或者私有网络统一ID，建议使用统一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID或者子网统一ID，建议使用统一ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DeleteHostedInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHostedInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHostedInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHostedInterfaceResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 黑石托管机器ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteHostedInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHostedInterfaceResponseParams `json:"Response"`
}

func (r *DeleteHostedInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHostedInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHostedInterfacesRequestParams struct {
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 物理机ID
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

type DeleteHostedInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 物理机ID
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

func (r *DeleteHostedInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHostedInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHostedInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHostedInterfacesResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteHostedInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHostedInterfacesResponseParams `json:"Response"`
}

func (r *DeleteHostedInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHostedInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInterfacesRequestParams struct {
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网的唯一ID列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

type DeleteInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网的唯一ID列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

func (r *DeleteInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInterfacesResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInterfacesResponseParams `json:"Response"`
}

func (r *DeleteInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayRequestParams struct {
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNatGatewayResponseParams `json:"Response"`
}

func (r *DeleteNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoutePolicyRequestParams struct {
	// 路由表ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表策略ID
	RoutePolicyId *string `json:"RoutePolicyId,omitempty" name:"RoutePolicyId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoutePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RoutePolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoutePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoutePolicyResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRoutePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoutePolicyResponseParams `json:"Response"`
}

func (r *DeleteRoutePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoutePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubnetRequestParams struct {
	// 私有网络ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网实例ID。可通过DescribeSubnets接口返回值中的SubnetId获取。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubnetResponseParams struct {
	// 异步任务ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSubnetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSubnetResponseParams `json:"Response"`
}

func (r *DeleteSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVirtualIpRequestParams struct {
	// 私有网络唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 退还的IP列表。
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

type DeleteVirtualIpRequest struct {
	*tchttp.BaseRequest
	
	// 私有网络唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 退还的IP列表。
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

func (r *DeleteVirtualIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVirtualIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Ips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVirtualIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVirtualIpResponseParams struct {
	// 异步任务ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVirtualIpResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVirtualIpResponseParams `json:"Response"`
}

func (r *DeleteVirtualIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVirtualIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcPeerConnectionRequestParams struct {
	// 黑石对等连接实例ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcPeerConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcPeerConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcPeerConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcPeerConnectionResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcPeerConnectionResponseParams `json:"Response"`
}

func (r *DeleteVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcPeerConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcRequestParams struct {
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcResponseParams struct {
	// 异步任务ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVpcResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcResponseParams `json:"Response"`
}

func (r *DeleteVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnConnectionRequestParams struct {
	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnConnectionResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnConnectionResponseParams `json:"Response"`
}

func (r *DeleteVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnGatewayResponseParams `json:"Response"`
}

func (r *DeleteVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterIpsRequestParams struct {
	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 注销指定IP的列表
	IpSet []*string `json:"IpSet,omitempty" name:"IpSet"`

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DeregisterIpsRequest struct {
	*tchttp.BaseRequest
	
	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 注销指定IP的列表
	IpSet []*string `json:"IpSet,omitempty" name:"IpSet"`

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DeregisterIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "IpSet")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterIpsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeregisterIpsResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterIpsResponseParams `json:"Response"`
}

func (r *DeregisterIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewaysRequestParams struct {
	// 对端网关ID，例如：bmcgw-2wqq41m9。每次请求的实例的上限为100。参数不支持同时指定CustomerGatewayIds和Filters。
	CustomerGatewayIds []*string `json:"CustomerGatewayIds,omitempty" name:"CustomerGatewayIds"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定CustomerGatewayIds和Filters。
	// <li>customergateway-name - String - （过滤条件）对端网关名称。</li>
	// <li>ip-address - String - （过滤条件)对端网关地址。</li>
	// <li>customergateway-id - String - （过滤条件）对端网关唯一ID。</li>
	// <li>zone - String - （过滤条件）对端所在可用区，形如：ap-guangzhou-2。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段, 支持"CreateTime"排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeCustomerGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// 对端网关ID，例如：bmcgw-2wqq41m9。每次请求的实例的上限为100。参数不支持同时指定CustomerGatewayIds和Filters。
	CustomerGatewayIds []*string `json:"CustomerGatewayIds,omitempty" name:"CustomerGatewayIds"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定CustomerGatewayIds和Filters。
	// <li>customergateway-name - String - （过滤条件）对端网关名称。</li>
	// <li>ip-address - String - （过滤条件)对端网关地址。</li>
	// <li>customergateway-id - String - （过滤条件）对端网关唯一ID。</li>
	// <li>zone - String - （过滤条件）对端所在可用区，形如：ap-guangzhou-2。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段, 支持"CreateTime"排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeCustomerGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewaysResponseParams struct {
	// 对端网关对象列表
	CustomerGatewaySet []*CustomerGateway `json:"CustomerGatewaySet,omitempty" name:"CustomerGatewaySet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomerGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerGatewaysResponseParams `json:"Response"`
}

func (r *DescribeCustomerGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaysRequestParams struct {
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

	// NAT所在可用区，形如：ap-guangzhou-2。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 排序字段, 支持"CreateTime"排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
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

	// NAT所在可用区，形如：ap-guangzhou-2。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 排序字段, 支持"CreateTime"排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeNatGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "NatName")
	delete(f, "SearchKey")
	delete(f, "VpcId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Zone")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaysResponseParams struct {
	// NAT网关信息列表
	NatGatewayInfoSet []*NatGatewayInfo `json:"NatGatewayInfoSet,omitempty" name:"NatGatewayInfoSet"`

	// 总数目
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNatGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatGatewaysResponseParams `json:"Response"`
}

func (r *DescribeNatGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatSubnetsRequestParams struct {
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatSubnetsResponseParams struct {
	// NAT子网信息
	NatSubnetInfoSet []*NatSubnetInfo `json:"NatSubnetInfoSet,omitempty" name:"NatSubnetInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNatSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatSubnetsResponseParams `json:"Response"`
}

func (r *DescribeNatSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoutePoliciesRequestParams struct {
	// 路由表实例ID，例如：rtb-afg8md3c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略实例ID，例如：rti-azd4dt1c。
	RoutePolicyIds []*string `json:"RoutePolicyIds,omitempty" name:"RoutePolicyIds"`

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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 初始行的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页行数，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRoutePoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-afg8md3c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略实例ID，例如：rti-azd4dt1c。
	RoutePolicyIds []*string `json:"RoutePolicyIds,omitempty" name:"RoutePolicyIds"`

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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 初始行的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页行数，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRoutePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoutePoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RoutePolicyIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoutePoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoutePoliciesResponseParams struct {
	// 路由策略数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 路由策略列表
	RoutePolicySet []*RoutePolicy `json:"RoutePolicySet,omitempty" name:"RoutePolicySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoutePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoutePoliciesResponseParams `json:"Response"`
}

func (r *DescribeRoutePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoutePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTablesRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableIds []*string `json:"RouteTableIds,omitempty" name:"RouteTableIds"`

	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// route-table-id - String - （过滤条件）路由表实例ID。
	// route-table-name - String - （过滤条件）路由表名称。
	// route-table-id-like - String - （模糊过滤条件）路由表实例ID。
	// route-table-name-like - String - （模糊过滤条件）路由表名称。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// zone - String - （过滤条件）可用区。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 初始行的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页行数，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段, 支持按“RouteTableId”，“VpcId”, "RouteTableName", "CreateTime"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeRouteTablesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableIds []*string `json:"RouteTableIds,omitempty" name:"RouteTableIds"`

	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// route-table-id - String - （过滤条件）路由表实例ID。
	// route-table-name - String - （过滤条件）路由表名称。
	// route-table-id-like - String - （模糊过滤条件）路由表实例ID。
	// route-table-name-like - String - （模糊过滤条件）路由表名称。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// zone - String - （过滤条件）可用区。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTablesResponseParams struct {
	// 路由表个数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 路由表列表
	RouteTableSet []*RouteTable `json:"RouteTableSet,omitempty" name:"RouteTableSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteTablesResponseParams `json:"Response"`
}

func (r *DescribeRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetAvailableIpsRequestParams struct {
	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CIDR前缀，例如10.0.1
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetAvailableIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	delete(f, "Cidr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetAvailableIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetAvailableIpsResponseParams struct {
	// 可用IP的范围列表
	IpSet []*string `json:"IpSet,omitempty" name:"IpSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubnetAvailableIpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetAvailableIpsResponseParams `json:"Response"`
}

func (r *DescribeSubnetAvailableIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetAvailableIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetByDeviceRequestParams struct {
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网类型。0: 物理机子网; 7: DOCKER子网 8: 虚拟子网
	Types []*uint64 `json:"Types,omitempty" name:"Types"`

	// 查询的起始位置。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的个数。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSubnetByDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 物理机ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网类型。0: 物理机子网; 7: DOCKER子网 8: 虚拟子网
	Types []*uint64 `json:"Types,omitempty" name:"Types"`

	// 查询的起始位置。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的个数。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubnetByDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetByDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Types")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetByDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetByDeviceResponseParams struct {
	// 子网个数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 子网列表
	Data []*SubnetInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubnetByDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetByDeviceResponseParams `json:"Response"`
}

func (r *DescribeSubnetByDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetByDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetByHostedDeviceRequestParams struct {
	// 托管机器ID, 如chm-xasdfx2j
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网类型。0: 物理机子网; 7: DOCKER子网 8: 虚拟子网
	Types []*uint64 `json:"Types,omitempty" name:"Types"`

	// 查询的起始位置。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的个数。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSubnetByHostedDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 托管机器ID, 如chm-xasdfx2j
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网类型。0: 物理机子网; 7: DOCKER子网 8: 虚拟子网
	Types []*uint64 `json:"Types,omitempty" name:"Types"`

	// 查询的起始位置。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的个数。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSubnetByHostedDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetByHostedDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Types")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetByHostedDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetByHostedDeviceResponseParams struct {
	// 子网个数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 子网列表
	Data []*SubnetInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubnetByHostedDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetByHostedDeviceResponseParams `json:"Response"`
}

func (r *DescribeSubnetByHostedDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetByHostedDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetsRequestParams struct {
	// 子网实例ID查询。形如：subnet-pxir56ns。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// subnet-id - String - （过滤条件）Subnet实例名称。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// cidr-block - String - （过滤条件）vpc的cidr。
	// subnet-name - String - （过滤条件）子网名称。
	// zone - String - （过滤条件）可用区。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段, 支持按“CreateTime”，“VlanId”
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// 子网实例ID查询。形如：subnet-pxir56ns。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// subnet-id - String - （过滤条件）Subnet实例名称。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// cidr-block - String - （过滤条件）vpc的cidr。
	// subnet-name - String - （过滤条件）子网名称。
	// zone - String - （过滤条件）可用区。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段, 支持按“CreateTime”，“VlanId”
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetsResponseParams struct {
	// 子网列表信息
	SubnetSet []*SubnetInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`

	// 返回的子网总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetsResponseParams `json:"Response"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// 任务状态，其中0表示任务执行成功，1表示任务执行失败，2表示任务正在执行中
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcPeerConnectionsRequestParams struct {
	// 对等连接实例ID
	VpcPeerConnectionIds []*string `json:"VpcPeerConnectionIds,omitempty" name:"VpcPeerConnectionIds"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定VpcPeerConnectionIds和Filters。
	// 过滤条件，参数不支持同时指定VpcPeerConnectionIds和Filters。
	// <li>peer-name - String - （过滤条件）对等连接名称。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeVpcPeerConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 对等连接实例ID
	VpcPeerConnectionIds []*string `json:"VpcPeerConnectionIds,omitempty" name:"VpcPeerConnectionIds"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定VpcPeerConnectionIds和Filters。
	// 过滤条件，参数不支持同时指定VpcPeerConnectionIds和Filters。
	// <li>peer-name - String - （过滤条件）对等连接名称。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeVpcPeerConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcPeerConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcPeerConnectionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcPeerConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcPeerConnectionsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 对等连接实例。
	VpcPeerConnectionSet []*VpcPeerConnection `json:"VpcPeerConnectionSet,omitempty" name:"VpcPeerConnectionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpcPeerConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcPeerConnectionsResponseParams `json:"Response"`
}

func (r *DescribeVpcPeerConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcPeerConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcQuotaRequestParams struct {
	// 类型
	TypeIds []*uint64 `json:"TypeIds,omitempty" name:"TypeIds"`
}

type DescribeVpcQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 类型
	TypeIds []*uint64 `json:"TypeIds,omitempty" name:"TypeIds"`
}

func (r *DescribeVpcQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcQuotaResponseParams struct {
	// 配额信息
	VpcQuotaSet []*VpcQuota `json:"VpcQuotaSet,omitempty" name:"VpcQuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpcQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcQuotaResponseParams `json:"Response"`
}

func (r *DescribeVpcQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcResourceRequestParams struct {
	// 私有网络实例ID
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// vpc-id - String - （过滤条件）私有网络实例ID，形如：vpc-f49l6u0z。
	// vpc-name - String - （过滤条件）私有网络名称。
	// zone - String - （过滤条件）可用区。
	// state - String - （过滤条件）VPC状态。available: 运营中; pending: 创建中; failed: 创建失败; deleting: 删除中
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeVpcResourceRequest struct {
	*tchttp.BaseRequest
	
	// 私有网络实例ID
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// vpc-id - String - （过滤条件）私有网络实例ID，形如：vpc-f49l6u0z。
	// vpc-name - String - （过滤条件）私有网络名称。
	// zone - String - （过滤条件）可用区。
	// state - String - （过滤条件）VPC状态。available: 运营中; pending: 创建中; failed: 创建失败; deleting: 删除中
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeVpcResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcResourceResponseParams struct {
	// VPC数据
	VpcResourceSet []*VpcResource `json:"VpcResourceSet,omitempty" name:"VpcResourceSet"`

	// VPC个数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpcResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcResourceResponseParams `json:"Response"`
}

func (r *DescribeVpcResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcViewRequestParams struct {
	// 私有网络唯一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcViewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcViewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcViewResponseParams struct {
	// VPC视图信息
	VpcView *VpcViewInfo `json:"VpcView,omitempty" name:"VpcView"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpcViewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcViewResponseParams `json:"Response"`
}

func (r *DescribeVpcViewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcsRequestParams struct {
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// vpc-name - String - （过滤条件）VPC实例名称。
	// vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。
	// cidr-block - String - （过滤条件）vpc的cidr。
	// state - String - （过滤条件）VPC状态。(pending | available).
	// zone -  String - （过滤条件）VPC的可用区。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 初始行的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页行数，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// vpc-name - String - （过滤条件）VPC实例名称。
	// vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。
	// cidr-block - String - （过滤条件）vpc的cidr。
	// state - String - （过滤条件）VPC状态。(pending | available).
	// zone -  String - （过滤条件）VPC的可用区。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 初始行的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页行数，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcsResponseParams struct {
	// VPC列表
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpcsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcsResponseParams `json:"Response"`
}

func (r *DescribeVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnConnectionsRequestParams struct {
	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnConnectionIds和Filters。
	VpnConnectionIds []*string `json:"VpnConnectionIds,omitempty" name:"VpnConnectionIds"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定VpnConnectionIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>state - String - （过滤条件 VPN状态：creating，available，createfailed，changing，changefailed，deleting，deletefailed。</li>
	// <li>zone - String - （过滤条件）VPN所在可用区，形如：ap-guangzhou-2。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// VPN网关实例ID
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// VPN通道名称
	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`

	// 排序字段, 支持"CreateTime"排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeVpnConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnConnectionIds和Filters。
	VpnConnectionIds []*string `json:"VpnConnectionIds,omitempty" name:"VpnConnectionIds"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定VpnConnectionIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>state - String - （过滤条件 VPN状态：creating，available，createfailed，changing，changefailed，deleting，deletefailed。</li>
	// <li>zone - String - （过滤条件）VPN所在可用区，形如：ap-guangzhou-2。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// VPN网关实例ID
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// VPN通道名称
	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`

	// 排序字段, 支持"CreateTime"排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeVpnConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnConnectionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "VpnGatewayId")
	delete(f, "VpnConnectionName")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnConnectionsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// VPN通道实例。
	VpnConnectionSet []*VpnConnection `json:"VpnConnectionSet,omitempty" name:"VpnConnectionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpnConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnConnectionsResponseParams `json:"Response"`
}

func (r *DescribeVpnConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaysRequestParams struct {
	// VPN网关实例ID。形如：bmvpngw-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnGatewayIds和Filters。
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitempty" name:"VpnGatewayIds"`

	// 过滤条件，参数不支持同时指定VpnGatewayIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>state - String - （过滤条件 VPN状态：creating，available，createfailed，changing，changefailed，deleting，deletefailed。</li>
	// <li>zone - String - （过滤条件）VPN所在可用区，形如：ap-guangzhou-2。</li>
	// <li>vpngw-name - String - （过滤条件）vpn网关名称。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 请求对象个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段, 支持"CreateTime"排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeVpnGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。形如：bmvpngw-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnGatewayIds和Filters。
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitempty" name:"VpnGatewayIds"`

	// 过滤条件，参数不支持同时指定VpnGatewayIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>state - String - （过滤条件 VPN状态：creating，available，createfailed，changing，changefailed，deleting，deletefailed。</li>
	// <li>zone - String - （过滤条件）VPN所在可用区，形如：ap-guangzhou-2。</li>
	// <li>vpngw-name - String - （过滤条件）vpn网关名称。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 请求对象个数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段, 支持"CreateTime"排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方向, “asc”、“desc”
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeVpnGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaysResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// VPN网关实例详细信息列表。
	VpnGatewaySet []*VpnGateway `json:"VpnGatewaySet,omitempty" name:"VpnGatewaySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpnGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnGatewaysResponseParams `json:"Response"`
}

func (r *DescribeVpnGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCustomerGatewayConfigurationRequestParams struct {
	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`

	// 厂商,取值 h3c，cisco
	VendorName *string `json:"VendorName,omitempty" name:"VendorName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCustomerGatewayConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnConnectionId")
	delete(f, "VendorName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadCustomerGatewayConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCustomerGatewayConfigurationResponseParams struct {
	// 配置信息。
	CustomerGatewayConfiguration *string `json:"CustomerGatewayConfiguration,omitempty" name:"CustomerGatewayConfiguration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadCustomerGatewayConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DownloadCustomerGatewayConfigurationResponseParams `json:"Response"`
}

func (r *DownloadCustomerGatewayConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCustomerGatewayConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type IKEOptionsSpecification struct {
	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBC-192', 'AES-CBC-256', 'DES-CBC'，默认为3DES-CBC
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

	// 报文封装模式:默认为Tunnel
	EncapMode *string `json:"EncapMode,omitempty" name:"EncapMode"`
}

type IpInfo struct {
	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// IP列表
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

// Predefined struct for user
type ModifyCustomerGatewayAttributeRequestParams struct {
	// 对端网关ID，例如：bmcgw-2wqq41m9，可通过DescribeCustomerGateways接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`

	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomerGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayId")
	delete(f, "CustomerGatewayName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomerGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomerGatewayAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCustomerGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomerGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyCustomerGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomerGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoutePolicyRequestParams struct {
	// 路由表ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 修改的路由
	RoutePolicy *RoutePolicy `json:"RoutePolicy,omitempty" name:"RoutePolicy"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoutePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RoutePolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoutePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoutePolicyResponseParams struct {
	// 异步任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRoutePolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoutePolicyResponseParams `json:"Response"`
}

func (r *ModifyRoutePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoutePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRouteTableRequestParams struct {
	// 路由表ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRouteTableResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRouteTableResponseParams `json:"Response"`
}

func (r *ModifyRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetAttributeRequestParams struct {
	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubnetAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "SubnetName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubnetAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubnetAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubnetAttributeResponseParams `json:"Response"`
}

func (r *ModifySubnetAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubnetAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetDHCPRelayRequestParams struct {
	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否开启DHCP Relay
	EnableDHCP *bool `json:"EnableDHCP,omitempty" name:"EnableDHCP"`

	// DHCP服务器IP
	ServerIps []*string `json:"ServerIps,omitempty" name:"ServerIps"`

	// 预留IP个数
	ReservedIpCount *uint64 `json:"ReservedIpCount,omitempty" name:"ReservedIpCount"`
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
	ServerIps []*string `json:"ServerIps,omitempty" name:"ServerIps"`

	// 预留IP个数
	ReservedIpCount *uint64 `json:"ReservedIpCount,omitempty" name:"ReservedIpCount"`
}

func (r *ModifySubnetDHCPRelayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubnetDHCPRelayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "EnableDHCP")
	delete(f, "ServerIps")
	delete(f, "ReservedIpCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubnetDHCPRelayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetDHCPRelayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubnetDHCPRelayResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubnetDHCPRelayResponseParams `json:"Response"`
}

func (r *ModifySubnetDHCPRelayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubnetDHCPRelayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAttributeRequestParams struct {
	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 是否开启内网监控，0为关闭，1为开启
	EnableMonitor *bool `json:"EnableMonitor,omitempty" name:"EnableMonitor"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "VpcName")
	delete(f, "EnableMonitor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVpcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcPeerConnectionRequestParams struct {
	// 黑石对等连接唯一ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`

	// 对等连接带宽
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 对等连接名称
	VpcPeerConnectionName *string `json:"VpcPeerConnectionName,omitempty" name:"VpcPeerConnectionName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcPeerConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcPeerConnectionId")
	delete(f, "Bandwidth")
	delete(f, "VpcPeerConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcPeerConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcPeerConnectionResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcPeerConnectionResponseParams `json:"Response"`
}

func (r *ModifyVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcPeerConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnConnectionAttributeRequestParams struct {
	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`

	// VPC实例ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPN通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`

	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey,omitempty" name:"PreShareKey"`

	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitempty" name:"SecurityPolicyDatabases"`

	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自我保护机制，用户配置网络安全协议。
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`

	// IPSec配置，腾讯云提供IPSec安全会话设置。
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`
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
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitempty" name:"SecurityPolicyDatabases"`

	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自我保护机制，用户配置网络安全协议。
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`

	// IPSec配置，腾讯云提供IPSec安全会话设置。
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`
}

func (r *ModifyVpnConnectionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnConnectionAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnConnectionId")
	delete(f, "VpcId")
	delete(f, "VpnConnectionName")
	delete(f, "PreShareKey")
	delete(f, "SecurityPolicyDatabases")
	delete(f, "IKEOptionsSpecification")
	delete(f, "IPSECOptionsSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnConnectionAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnConnectionAttributeResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVpnConnectionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnConnectionAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpnConnectionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnConnectionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayAttributeRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "VpnGatewayName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVpnGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpnGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 网关创建状态，其中0表示创建中，1表示运行中，2表示创建失败
	ProductionStatus *uint64 `json:"ProductionStatus,omitempty" name:"ProductionStatus"`

	// EIP列表
	Eips []*string `json:"Eips,omitempty" name:"Eips"`

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

	// 网关启用状态，1为禁用，0为启用。
	State *uint64 `json:"State,omitempty" name:"State"`

	// 私有网络整型ID
	IntVpcId *uint64 `json:"IntVpcId,omitempty" name:"IntVpcId"`

	// NAT资源ID
	NatResourceId *uint64 `json:"NatResourceId,omitempty" name:"NatResourceId"`
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

// Predefined struct for user
type RejectVpcPeerConnectionRequestParams struct {
	// 黑石对等连接实例ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectVpcPeerConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcPeerConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RejectVpcPeerConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectVpcPeerConnectionResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RejectVpcPeerConnectionResponse struct {
	*tchttp.BaseResponse
	Response *RejectVpcPeerConnectionResponseParams `json:"Response"`
}

func (r *RejectVpcPeerConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectVpcPeerConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetVpnConnectionRequestParams struct {
	// VPC唯一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPN通道实例ID。形如：bmvpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetVpnConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "VpnConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetVpnConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetVpnConnectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *ResetVpnConnectionResponseParams `json:"Response"`
}

func (r *ResetVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	RemoteCidrBlock []*string `json:"RemoteCidrBlock,omitempty" name:"RemoteCidrBlock"`
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
	DhcpServerIp []*string `json:"DhcpServerIp,omitempty" name:"DhcpServerIp"`

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

	// 子网可用区ID。
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
	DhcpServerIp []*string `json:"DhcpServerIp,omitempty" name:"DhcpServerIp"`

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

	// 子网可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC所在可用区ID
	VpcZoneId *uint64 `json:"VpcZoneId,omitempty" name:"VpcZoneId"`

	// VPC所在可用区
	VpcZone *string `json:"VpcZone,omitempty" name:"VpcZone"`

	// 是否开启广播，关闭为0，开启为1。
	BroadcastFlag *uint64 `json:"BroadcastFlag,omitempty" name:"BroadcastFlag"`
}

// Predefined struct for user
type UnbindEipsFromNatGatewayRequestParams struct {
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已分配的EIP列表
	AssignedEips []*string `json:"AssignedEips,omitempty" name:"AssignedEips"`
}

type UnbindEipsFromNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 已分配的EIP列表
	AssignedEips []*string `json:"AssignedEips,omitempty" name:"AssignedEips"`
}

func (r *UnbindEipsFromNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindEipsFromNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "VpcId")
	delete(f, "AssignedEips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindEipsFromNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindEipsFromNatGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindEipsFromNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *UnbindEipsFromNatGatewayResponseParams `json:"Response"`
}

func (r *UnbindEipsFromNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindEipsFromNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindIpsFromNatGatewayRequestParams struct {
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 部分IP信息；子网须以部分IP将加入NAT网关
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitempty" name:"IpInfoSet"`
}

type UnbindIpsFromNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 部分IP信息；子网须以部分IP将加入NAT网关
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitempty" name:"IpInfoSet"`
}

func (r *UnbindIpsFromNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindIpsFromNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "VpcId")
	delete(f, "IpInfoSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindIpsFromNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindIpsFromNatGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindIpsFromNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *UnbindIpsFromNatGatewayResponseParams `json:"Response"`
}

func (r *UnbindIpsFromNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindIpsFromNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindSubnetsFromNatGatewayRequestParams struct {
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表，子网不区分加入NAT网关的转发方式
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

type UnbindSubnetsFromNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表，子网不区分加入NAT网关的转发方式
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

func (r *UnbindSubnetsFromNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindSubnetsFromNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "VpcId")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindSubnetsFromNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindSubnetsFromNatGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindSubnetsFromNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *UnbindSubnetsFromNatGatewayResponseParams `json:"Response"`
}

func (r *UnbindSubnetsFromNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindSubnetsFromNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeNatGatewayRequestParams struct {
	// NAT网关ID，例如：nat-kdm476mp
	NatId *string `json:"NatId,omitempty" name:"NatId"`

	// 私有网络ID，例如：vpc-kd7d06of
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 并发连接数规格；取值为1000000、3000000、10000000，分别对应小型、中型、大型NAT网关
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitempty" name:"MaxConcurrent"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatId")
	delete(f, "VpcId")
	delete(f, "MaxConcurrent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeNatGatewayResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeNatGatewayResponseParams `json:"Response"`
}

func (r *UpgradeNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 整型私有网络ID。
	IntVpcId *uint64 `json:"IntVpcId,omitempty" name:"IntVpcId"`
}

type VpcPeerConnection struct {
	// 本端VPC唯一ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 对端VPC唯一ID
	PeerVpcId *string `json:"PeerVpcId,omitempty" name:"PeerVpcId"`

	// 本端APPID
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 对端APPID
	PeerAppId *string `json:"PeerAppId,omitempty" name:"PeerAppId"`

	// 对等连接唯一ID
	VpcPeerConnectionId *string `json:"VpcPeerConnectionId,omitempty" name:"VpcPeerConnectionId"`

	// 对等连接名称
	VpcPeerConnectionName *string `json:"VpcPeerConnectionName,omitempty" name:"VpcPeerConnectionName"`

	// 对等连接状态。pending:申请中,available:运行中,expired:已过期,rejected:已拒绝,deleted:已删除
	State *string `json:"State,omitempty" name:"State"`

	// 本端VPC所属可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcZone *string `json:"VpcZone,omitempty" name:"VpcZone"`

	// 对端VPC所属可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerVpcZone *string `json:"PeerVpcZone,omitempty" name:"PeerVpcZone"`

	// 本端Uin
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 对端Uin
	PeerUin *uint64 `json:"PeerUin,omitempty" name:"PeerUin"`

	// 对等连接类型
	PeerType *uint64 `json:"PeerType,omitempty" name:"PeerType"`

	// 对等连接带宽
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 本端VPC地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 对端VPC地域
	PeerRegion *string `json:"PeerRegion,omitempty" name:"PeerRegion"`

	// 是否允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *uint64 `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

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

	// 云联网服务个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnServiceNum *uint64 `json:"CcnServiceNum,omitempty" name:"CcnServiceNum"`

	// VPC允许创建的对等连接个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcPeerLimitToAllRegion *uint64 `json:"VpcPeerLimitToAllRegion,omitempty" name:"VpcPeerLimitToAllRegion"`

	// VPC允许创建的同地域的对等连接的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcPeerLimitToSameRegion *uint64 `json:"VpcPeerLimitToSameRegion,omitempty" name:"VpcPeerLimitToSameRegion"`

	// 整型私有网络ID
	IntVpcId *uint64 `json:"IntVpcId,omitempty" name:"IntVpcId"`
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
	SubnetSet []*VpcSubnetViewInfo `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

type VpnConnection struct {
	// 通道实例ID。
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" name:"VpnConnectionId"`

	// 通道名称。
	VpnConnectionName *string `json:"VpnConnectionName,omitempty" name:"VpnConnectionName"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// 对端网关实例ID。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`

	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey,omitempty" name:"PreShareKey"`

	// 通道传输协议。
	VpnProto *string `json:"VpnProto,omitempty" name:"VpnProto"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 通道的生产状态
	State *string `json:"State,omitempty" name:"State"`

	// 通道连接状态
	NetStatus *string `json:"NetStatus,omitempty" name:"NetStatus"`

	// SPD。
	SecurityPolicyDatabaseSet []*SecurityPolicyDatabase `json:"SecurityPolicyDatabaseSet,omitempty" name:"SecurityPolicyDatabaseSet"`

	// IKE选项。
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitempty" name:"IKEOptionsSpecification"`

	// IPSEC选项。
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitempty" name:"IPSECOptionsSpecification"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// VPC名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPN网关名称
	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`

	// 对端网关名称
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`

	// IPSEC VPN通道路由策略目的端地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestinationCidr []*string `json:"DestinationCidr,omitempty" name:"DestinationCidr"`

	// IPSEC VPN通道路由策略源端地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceCidr []*string `json:"SourceCidr,omitempty" name:"SourceCidr"`
}

type VpnGateway struct {
	// 网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 网关实例名称。
	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`

	// VPC网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// VPC名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 网关出带宽。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 网关实例状态
	State *string `json:"State,omitempty" name:"State"`

	// 网关公网IP。
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 可用区，如：ap-guangzhou
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPN网关的通道数
	VpnConnNum *uint64 `json:"VpnConnNum,omitempty" name:"VpnConnNum"`
}