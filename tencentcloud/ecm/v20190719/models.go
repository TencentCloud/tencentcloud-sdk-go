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

package v20190719

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Address struct {
	// EIP的ID，是EIP的唯一标识。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// EIP名称。
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// EIP状态，包含'CREATING'(创建中),'BINDING'(绑定中),'BIND'(已绑定),'UNBINDING'(解绑中),'UNBIND'(已解绑),'OFFLINING'(释放中),'BIND_ENI'(绑定悬空弹性网卡)
	AddressStatus *string `json:"AddressStatus,omitempty" name:"AddressStatus"`

	// 外网IP地址
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`

	// 绑定的资源实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 创建时间。ISO 8601 格式：YYYY-MM-DDTHH:mm:ss.sssZ
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 绑定的弹性网卡ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 绑定的资源内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateAddressIp *string `json:"PrivateAddressIp,omitempty" name:"PrivateAddressIp"`

	// 资源隔离状态。true表示eip处于隔离状态，false表示资源处于未隔离状态
	IsArrears *bool `json:"IsArrears,omitempty" name:"IsArrears"`

	// 资源封堵状态。true表示eip处于封堵状态，false表示eip处于未封堵状态
	IsBlocked *bool `json:"IsBlocked,omitempty" name:"IsBlocked"`

	// eip是否支持直通模式。true表示eip支持直通模式，false表示资源不支持直通模式
	IsEipDirectConnection *bool `json:"IsEipDirectConnection,omitempty" name:"IsEipDirectConnection"`

	// eip资源类型，包括"CalcIP","WanIP","EIP","AnycastEIP"。其中"CalcIP"表示设备ip，“WanIP”表示普通公网ip，“EIP”表示弹性公网ip，“AnycastEip”表示加速EIP
	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`

	// eip是否在解绑后自动释放。true表示eip将会在解绑后自动释放，false表示eip在解绑后不会自动释放
	CascadeRelease *bool `json:"CascadeRelease,omitempty" name:"CascadeRelease"`

	// 运营商，CTCC电信，CUCC联通，CMCC移动
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`

	// 带宽上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

type AddressInfo struct {
	// 实例的外网ip相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIPAddressInfo *PublicIPAddressInfo `json:"PublicIPAddressInfo,omitempty" name:"PublicIPAddressInfo"`

	// 实例的内网ip相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIPAddressInfo *PrivateIPAddressInfo `json:"PrivateIPAddressInfo,omitempty" name:"PrivateIPAddressInfo"`

	// 实例的外网ipv6相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIPv6AddressInfo *PublicIPAddressInfo `json:"PublicIPv6AddressInfo,omitempty" name:"PublicIPv6AddressInfo"`
}

type AddressTemplateSpecification struct {
	// IP地址ID，例如：eipm-2uw6ujo6。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// IP地址组ID，例如：eipmg-2uw6ujo6。
	AddressGroupId *string `json:"AddressGroupId,omitempty" name:"AddressGroupId"`
}

// Predefined struct for user
type AllocateAddressesRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// EIP数量。默认值：1。
	AddressCount *uint64 `json:"AddressCount,omitempty" name:"AddressCount"`

	// CMCC：中国移动
	// CTCC：中国电信
	// CUCC：中国联通
	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`

	// 1 Mbps 至 5000 Mbps，默认值：1 Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 需要关联的标签列表。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 要绑定的实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要绑定的弹性网卡 ID。 弹性网卡 ID 形如：eni-11112222。NetworkInterfaceId 与 InstanceId 不可同时指定。弹性网卡 ID 可通过DescribeNetworkInterfaces接口返回值中的networkInterfaceId获取。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 要绑定的内网 IP。如果指定了 NetworkInterfaceId 则也必须指定 PrivateIpAddress ，表示将 EIP 绑定到指定弹性网卡的指定内网 IP 上。同时要确保指定的 PrivateIpAddress 是指定的 NetworkInterfaceId 上的一个内网 IP。指定弹性网卡的内网 IP 可通过DescribeNetworkInterfaces接口返回值中的privateIpAddress获取。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

type AllocateAddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// EIP数量。默认值：1。
	AddressCount *uint64 `json:"AddressCount,omitempty" name:"AddressCount"`

	// CMCC：中国移动
	// CTCC：中国电信
	// CUCC：中国联通
	InternetServiceProvider *string `json:"InternetServiceProvider,omitempty" name:"InternetServiceProvider"`

	// 1 Mbps 至 5000 Mbps，默认值：1 Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 需要关联的标签列表。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 要绑定的实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要绑定的弹性网卡 ID。 弹性网卡 ID 形如：eni-11112222。NetworkInterfaceId 与 InstanceId 不可同时指定。弹性网卡 ID 可通过DescribeNetworkInterfaces接口返回值中的networkInterfaceId获取。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 要绑定的内网 IP。如果指定了 NetworkInterfaceId 则也必须指定 PrivateIpAddress ，表示将 EIP 绑定到指定弹性网卡的指定内网 IP 上。同时要确保指定的 PrivateIpAddress 是指定的 NetworkInterfaceId 上的一个内网 IP。指定弹性网卡的内网 IP 可通过DescribeNetworkInterfaces接口返回值中的privateIpAddress获取。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *AllocateAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressCount")
	delete(f, "InternetServiceProvider")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "Tags")
	delete(f, "InstanceId")
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateAddressesResponseParams struct {
	// 申请到的 EIP 的唯一 ID 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet"`

	// 异步任务TaskId。可以使用DescribeTaskResult接口查询任务状态。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AllocateAddressesResponse struct {
	*tchttp.BaseResponse
	Response *AllocateAddressesResponseParams `json:"Response"`
}

func (r *AllocateAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Area struct {
	// 区域ID
	AreaId *string `json:"AreaId,omitempty" name:"AreaId"`

	// 区域名称
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
}

// Predefined struct for user
type AssignIpv6AddressesRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，形如：eni-1snva0vd。目前只支持主网卡上分配。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的IPv6地址列表，单次最多指定10个。与入参Ipv6AddressCount合并计算配额。与Ipv6AddressCount必填一个。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`

	// 自动分配IPv6地址个数，内网IP地址个数总和不能超过配数。与入参Ipv6Addresses合并计算配额。与Ipv6Addresses必填一个。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`

	// ipv6运营商如下：
	// CTCC：中国电信
	// CUCC：中国联通
	// CMCC：中国移动
	Ipv6ISP *string `json:"Ipv6ISP,omitempty" name:"Ipv6ISP"`
}

type AssignIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，形如：eni-1snva0vd。目前只支持主网卡上分配。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的IPv6地址列表，单次最多指定10个。与入参Ipv6AddressCount合并计算配额。与Ipv6AddressCount必填一个。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`

	// 自动分配IPv6地址个数，内网IP地址个数总和不能超过配数。与入参Ipv6Addresses合并计算配额。与Ipv6Addresses必填一个。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`

	// ipv6运营商如下：
	// CTCC：中国电信
	// CUCC：中国联通
	// CMCC：中国移动
	Ipv6ISP *string `json:"Ipv6ISP,omitempty" name:"Ipv6ISP"`
}

func (r *AssignIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	delete(f, "Ipv6AddressCount")
	delete(f, "Ipv6ISP")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6AddressesResponseParams struct {
	// 分配给弹性网卡的IPv6地址列表。
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssignIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *AssignIpv6AddressesResponseParams `json:"Response"`
}

func (r *AssignIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignPrivateIpAddressesRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 指定的内网IP信息，单次最多指定10个。与SecondaryPrivateIpAddressCount至少提供一个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 新申请的内网IP地址个数，与PrivateIpAddresses至少提供一个。内网IP地址个数总和不能超过配额数
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`
}

type AssignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 指定的内网IP信息，单次最多指定10个。与SecondaryPrivateIpAddressCount至少提供一个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 新申请的内网IP地址个数，与PrivateIpAddresses至少提供一个。内网IP地址个数总和不能超过配额数
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`
}

func (r *AssignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignPrivateIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "EcmRegion")
	delete(f, "PrivateIpAddresses")
	delete(f, "SecondaryPrivateIpAddressCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignPrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignPrivateIpAddressesResponseParams struct {
	// 内网IP详细信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *AssignPrivateIpAddressesResponseParams `json:"Response"`
}

func (r *AssignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssistantCidr struct {
	// VPC实例ID。形如：vpc-6v2ht8q5
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 辅助CIDR。形如：172.16.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 辅助CIDR类型（0：普通辅助CIDR，1：容器辅助CIDR），默认都是0。
	AssistantType *uint64 `json:"AssistantType,omitempty" name:"AssistantType"`

	// 辅助CIDR拆分的子网。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet"`
}

// Predefined struct for user
type AssociateAddressRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：eip-11112222。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 要绑定的实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要绑定的弹性网卡 ID。 弹性网卡 ID 形如：eni-11112222。NetworkInterfaceId 与 InstanceId 不可同时指定。弹性网卡 ID 可通过DescribeNetworkInterfaces接口返回值中的networkInterfaceId获取。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 要绑定的内网 IP。如果指定了 NetworkInterfaceId 则也必须指定 PrivateIpAddress ，表示将 EIP 绑定到指定弹性网卡的指定内网 IP 上。同时要确保指定的 PrivateIpAddress 是指定的 NetworkInterfaceId 上的一个内网 IP。指定弹性网卡的内网 IP 可通过DescribeNetworkInterfaces接口返回值中的privateIpAddress获取。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

type AssociateAddressRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：eip-11112222。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 要绑定的实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要绑定的弹性网卡 ID。 弹性网卡 ID 形如：eni-11112222。NetworkInterfaceId 与 InstanceId 不可同时指定。弹性网卡 ID 可通过DescribeNetworkInterfaces接口返回值中的networkInterfaceId获取。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 要绑定的内网 IP。如果指定了 NetworkInterfaceId 则也必须指定 PrivateIpAddress ，表示将 EIP 绑定到指定弹性网卡的指定内网 IP 上。同时要确保指定的 PrivateIpAddress 是指定的 NetworkInterfaceId 上的一个内网 IP。指定弹性网卡的内网 IP 可通过DescribeNetworkInterfaces接口返回值中的privateIpAddress获取。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *AssociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressId")
	delete(f, "InstanceId")
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateAddressResponseParams struct {
	// 异步任务TaskId。可以使用DescribeTaskResult接口查询任务状态。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *AssociateAddressResponseParams `json:"Response"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 要绑定的安全组ID，类似esg-efil73jd，只支持绑定单个安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 被绑定的实例ID，类似ein-lesecurk，支持指定多个实例，每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要绑定的安全组ID，类似esg-efil73jd，只支持绑定单个安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 被绑定的实例ID，类似ein-lesecurk，支持指定多个实例，每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachDisksRequestParams struct {
	// 云服务器实例ID。云盘将被挂载到此云服务器上，通过[DescribeInstances](/document/product/213/15728)接口查询。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 将要被挂载的弹性云盘ID。通过[DescribeDisks](/document/product/362/16315)接口查询。单次最多可挂载10块弹性云盘。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 可选参数，不传该参数则仅执行挂载操作。传入`True`时，会在挂载成功后将云硬盘设置为随云主机销毁模式，仅对按量计费云硬盘有效。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// 可选参数，用于控制云盘挂载时使用的挂载模式，目前仅对黑石裸金属机型有效。取值范围：<br><li>PF<br><li>VF
	AttachMode *string `json:"AttachMode,omitempty" name:"AttachMode"`
}

type AttachDisksRequest struct {
	*tchttp.BaseRequest
	
	// 云服务器实例ID。云盘将被挂载到此云服务器上，通过[DescribeInstances](/document/product/213/15728)接口查询。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 将要被挂载的弹性云盘ID。通过[DescribeDisks](/document/product/362/16315)接口查询。单次最多可挂载10块弹性云盘。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 可选参数，不传该参数则仅执行挂载操作。传入`True`时，会在挂载成功后将云硬盘设置为随云主机销毁模式，仅对按量计费云硬盘有效。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// 可选参数，用于控制云盘挂载时使用的挂载模式，目前仅对黑石裸金属机型有效。取值范围：<br><li>PF<br><li>VF
	AttachMode *string `json:"AttachMode,omitempty" name:"AttachMode"`
}

func (r *AttachDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DiskIds")
	delete(f, "DeleteWithInstance")
	delete(f, "AttachMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachDisksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachDisksResponse struct {
	*tchttp.BaseResponse
	Response *AttachDisksResponseParams `json:"Response"`
}

func (r *AttachDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachNetworkInterfaceRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 实例ID。形如：ein-r8hr2upy。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type AttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 实例ID。形如：ein-r8hr2upy。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *AttachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "InstanceId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachNetworkInterfaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *AttachNetworkInterfaceResponseParams `json:"Response"`
}

func (r *AttachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Backend struct {
	// 后端服务的唯一 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 后端服务的监听端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 后端服务的内网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 后端服务被绑定的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisteredTime *string `json:"RegisteredTime,omitempty" name:"RegisteredTime"`

	// 弹性网卡唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniId *string `json:"EniId,omitempty" name:"EniId"`

	// 后端服务的外网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// 后端服务的实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

// Predefined struct for user
type BatchDeregisterTargetsRequestParams struct {
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 解绑目标
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

type BatchDeregisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 解绑目标
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *BatchDeregisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeregisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeregisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeregisterTargetsResponseParams struct {
	// 解绑失败的监听器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeregisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeregisterTargetsResponseParams `json:"Response"`
}

func (r *BatchDeregisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeregisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTargetWeightRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要批量修改权重的列表
	ModifyList []*TargetsWeightRule `json:"ModifyList,omitempty" name:"ModifyList"`
}

type BatchModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要批量修改权重的列表
	ModifyList []*TargetsWeightRule `json:"ModifyList,omitempty" name:"ModifyList"`
}

func (r *BatchModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ModifyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyTargetWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTargetWeightResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyTargetWeightResponseParams `json:"Response"`
}

func (r *BatchModifyTargetWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterTargetsRequestParams struct {
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 绑定目标
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

type BatchRegisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 绑定目标
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *BatchRegisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchRegisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterTargetsResponseParams struct {
	// 绑定失败的监听器ID，如为空表示全部绑定成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchRegisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchRegisterTargetsResponseParams `json:"Response"`
}

func (r *BatchRegisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchTarget struct {
	// 监听器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 绑定端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 子机ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 弹性网卡ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`

	// 子机权重，范围[0, 100]。绑定时如果不存在，则默认为10。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type City struct {
	// 城市ID
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 城市名称
	CityName *string `json:"CityName,omitempty" name:"CityName"`
}

type Country struct {
	// 国家ID
	CountryId *string `json:"CountryId,omitempty" name:"CountryId"`

	// 国家名称
	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
}

// Predefined struct for user
type CreateDisksRequestParams struct {
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 云硬盘计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDCPAID：独享集群付费<br>各类型价格请参考云硬盘[价格总览](/document/product/362/2413)。
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘<br><li>CLOUD_HSSD：表示增强型SSD云硬盘<br><li>CLOUD_TSSD：表示极速型SSD云硬盘。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 云盘显示名称。不传则默认为“未命名”。最大长度不能超60个字节。
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// 云盘绑定的标签。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 预付费模式，即包年包月相关参数设置。通过该参数指定包年包月云盘的购买时长、是否设置自动续费等属性。<br>创建预付费云盘该参数必传，创建按小时后付费云盘无需传该参数。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// 创建云硬盘数量，不传则默认为1。单次请求最多可创建的云盘数有限制，具体参见[云硬盘使用限制](https://cloud.tencent.com/doc/product/362/5145)。
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// 可选参数。使用此参数可给云硬盘购买额外的性能。<br>当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// 云硬盘大小，单位为GB。<br><li>如果传入`SnapshotId`则可不传`DiskSize`，此时新建云盘的大小为快照大小<br><li>如果传入`SnapshotId`同时传入`DiskSize`，则云盘大小必须大于或等于快照大小<br><li>云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 可选参数，默认为False。传入True时，云盘将创建为共享型云盘。
	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 传入该参数用于创建加密云盘，取值固定为ENCRYPT。
	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`

	// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照，可通过[DescribeSnapshots](/document/product/362/15647)接口查询快照，见输出参数DiskUsage解释。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type CreateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 云硬盘计费类型。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDCPAID：独享集群付费<br>各类型价格请参考云硬盘[价格总览](/document/product/362/2413)。
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘<br><li>CLOUD_HSSD：表示增强型SSD云硬盘<br><li>CLOUD_TSSD：表示极速型SSD云硬盘。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 云盘显示名称。不传则默认为“未命名”。最大长度不能超60个字节。
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// 云盘绑定的标签。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 预付费模式，即包年包月相关参数设置。通过该参数指定包年包月云盘的购买时长、是否设置自动续费等属性。<br>创建预付费云盘该参数必传，创建按小时后付费云盘无需传该参数。
	DiskChargePrepaid *DiskChargePrepaid `json:"DiskChargePrepaid,omitempty" name:"DiskChargePrepaid"`

	// 创建云硬盘数量，不传则默认为1。单次请求最多可创建的云盘数有限制，具体参见[云硬盘使用限制](https://cloud.tencent.com/doc/product/362/5145)。
	DiskCount *uint64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// 可选参数。使用此参数可给云硬盘购买额外的性能。<br>当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// 云硬盘大小，单位为GB。<br><li>如果传入`SnapshotId`则可不传`DiskSize`，此时新建云盘的大小为快照大小<br><li>如果传入`SnapshotId`同时传入`DiskSize`，则云盘大小必须大于或等于快照大小<br><li>云盘大小取值范围参见云硬盘[产品分类](/document/product/362/2353)的说明。
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 可选参数，默认为False。传入True时，云盘将创建为共享型云盘。
	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 传入该参数用于创建加密云盘，取值固定为ENCRYPT。
	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`

	// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照，可通过[DescribeSnapshots](/document/product/362/15647)接口查询快照，见输出参数DiskUsage解释。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

func (r *CreateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "DiskChargeType")
	delete(f, "DiskType")
	delete(f, "DiskName")
	delete(f, "Tags")
	delete(f, "DiskChargePrepaid")
	delete(f, "DiskCount")
	delete(f, "ThroughputPerformance")
	delete(f, "DiskSize")
	delete(f, "Shareable")
	delete(f, "ClientToken")
	delete(f, "Encrypt")
	delete(f, "SnapshotId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisksResponseParams struct {
	// 创建的云硬盘ID列表。
	DiskIdSet []*string `json:"DiskIdSet,omitempty" name:"DiskIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDisksResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisksResponseParams `json:"Response"`
}

func (r *CreateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHaVipRequestParams struct {
	// HAVIP所在私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// HAVIP所在子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// HAVIP名称。
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`

	// 指定虚拟IP地址，必须在VPC网段内且未被占用。不指定则自动分配。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type CreateHaVipRequest struct {
	*tchttp.BaseRequest
	
	// HAVIP所在私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// HAVIP所在子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// HAVIP名称。
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`

	// 指定虚拟IP地址，必须在VPC网段内且未被占用。不指定则自动分配。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *CreateHaVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "HaVipName")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHaVipResponseParams struct {
	// HAVIP对象。
	HaVip *HaVip `json:"HaVip,omitempty" name:"HaVip"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHaVipResponse struct {
	*tchttp.BaseResponse
	Response *CreateHaVipResponseParams `json:"Response"`
}

func (r *CreateHaVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageRequestParams struct {
	// 镜像名称。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 需要制作镜像的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 镜像描述。
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 是否执行强制关机以制作镜像。取值范围：
	// TRUE：表示自动关机后制作镜像
	// FALSE：表示开机状态制作，目前不支持，需要先手动关机
	// 默认取值：FALSE。
	ForcePoweroff *string `json:"ForcePoweroff,omitempty" name:"ForcePoweroff"`
}

type CreateImageRequest struct {
	*tchttp.BaseRequest
	
	// 镜像名称。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 需要制作镜像的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 镜像描述。
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 是否执行强制关机以制作镜像。取值范围：
	// TRUE：表示自动关机后制作镜像
	// FALSE：表示开机状态制作，目前不支持，需要先手动关机
	// 默认取值：FALSE。
	ForcePoweroff *string `json:"ForcePoweroff,omitempty" name:"ForcePoweroff"`
}

func (r *CreateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageName")
	delete(f, "InstanceId")
	delete(f, "ImageDescription")
	delete(f, "ForcePoweroff")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageResponseParams `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairRequestParams struct {
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest
	
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
}

func (r *CreateKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyPairRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeyPairRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyPairResponseParams struct {
	// 密钥对信息。
	KeyPair *KeyPair `json:"KeyPair,omitempty" name:"KeyPair"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *CreateKeyPairResponseParams `json:"Response"`
}

func (r *CreateKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要将监听器创建到哪些端口，每个端口对应一个新的监听器
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// 监听器协议： TCP | UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 要创建的监听器名称列表，名称与Ports数组按序一一对应，如不需立即命名，则无需提供此参数
	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`

	// 健康检查相关参数
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 会话保持类型。不传或传NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// 批量端口段的结束端口，必须和Ports长度一样。
	EndPorts []*int64 `json:"EndPorts,omitempty" name:"EndPorts"`
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要将监听器创建到哪些端口，每个端口对应一个新的监听器
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// 监听器协议： TCP | UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 要创建的监听器名称列表，名称与Ports数组按序一一对应，如不需立即命名，则无需提供此参数
	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`

	// 健康检查相关参数
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 会话保持类型。不传或传NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// 批量端口段的结束端口，必须和Ports长度一样。
	EndPorts []*int64 `json:"EndPorts,omitempty" name:"EndPorts"`
}

func (r *CreateListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Ports")
	delete(f, "Protocol")
	delete(f, "ListenerNames")
	delete(f, "HealthCheck")
	delete(f, "SessionExpireTime")
	delete(f, "Scheduler")
	delete(f, "SessionType")
	delete(f, "EndPorts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerResponseParams struct {
	// 创建的监听器的唯一标识数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse
	Response *CreateListenerResponseParams `json:"Response"`
}

func (r *CreateListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerRequestParams struct {
	// ECM区域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 负载均衡实例的网络类型。目前只支持传入OPEN，表示公网属性。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// CMCC | CTCC | CUCC，分别对应 移动 | 电信 | 联通。
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// 负载均衡实例的名称，只在创建一个实例的时候才会生效。规则：1-50 个英文、汉字、数字、连接线“-”或下划线“_”。
	// 注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡后端目标设备所属的网络 ID，如vpc-12345678。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 创建负载均衡的个数，默认值 1。
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 负载均衡的带宽限制等信息。
	InternetAccessible *LoadBalancerInternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 安全组。
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// 仅适用于公网负载均衡。IP版本，可取值：IPV4、IPv6FullChain，默认值 IPV4。说明：取值为IPv6FullChain，表示为IPv6版本。
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// 在购买IPV6负载均衡实例的情况下，必须指定子网 ID, 此参数必填。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// ECM区域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 负载均衡实例的网络类型。目前只支持传入OPEN，表示公网属性。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// CMCC | CTCC | CUCC，分别对应 移动 | 电信 | 联通。
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// 负载均衡实例的名称，只在创建一个实例的时候才会生效。规则：1-50 个英文、汉字、数字、连接线“-”或下划线“_”。
	// 注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡后端目标设备所属的网络 ID，如vpc-12345678。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 创建负载均衡的个数，默认值 1。
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 负载均衡的带宽限制等信息。
	InternetAccessible *LoadBalancerInternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 安全组。
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// 仅适用于公网负载均衡。IP版本，可取值：IPV4、IPv6FullChain，默认值 IPV4。说明：取值为IPv6FullChain，表示为IPv6版本。
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// 在购买IPV6负载均衡实例的情况下，必须指定子网 ID, 此参数必填。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "LoadBalancerType")
	delete(f, "VipIsp")
	delete(f, "LoadBalancerName")
	delete(f, "VpcId")
	delete(f, "Number")
	delete(f, "InternetAccessible")
	delete(f, "Tags")
	delete(f, "SecurityGroups")
	delete(f, "AddressIPVersion")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerResponseParams struct {
	// 由负载均衡实例ID组成的数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancerResponseParams `json:"Response"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModuleRequestParams struct {
	// 模块名称，如视频直播模块。限制：模块名称不得以空格开头，长度不得超过60个字符。
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// 默认带宽，单位：M。范围不得超过带宽上下限，详看DescribeConfig。
	DefaultBandWidth *int64 `json:"DefaultBandWidth,omitempty" name:"DefaultBandWidth"`

	// 默认镜像，如img-qsdf3ff2。
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// 机型ID。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 默认系统盘大小，单位：G，默认大小为50G。范围不得超过系统盘上下限制，详看DescribeConfig。
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// 默认数据盘大小，单位：G。范围不得超过数据盘范围大小，详看DescribeConfig。
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// 是否关闭IP直通。取值范围：
	// true：表示关闭IP直通
	// false：表示开通IP直通
	CloseIpDirect *bool `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`

	// 标签列表。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 模块默认安全组列表
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// 默认入带宽，单位：M。范围不得超过带宽上下限，详看DescribeConfig。
	DefaultBandWidthIn *int64 `json:"DefaultBandWidthIn,omitempty" name:"DefaultBandWidthIn"`

	// 是否禁止分配外网IP
	DisableWanIp *bool `json:"DisableWanIp,omitempty" name:"DisableWanIp"`

	// 系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘信息。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type CreateModuleRequest struct {
	*tchttp.BaseRequest
	
	// 模块名称，如视频直播模块。限制：模块名称不得以空格开头，长度不得超过60个字符。
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// 默认带宽，单位：M。范围不得超过带宽上下限，详看DescribeConfig。
	DefaultBandWidth *int64 `json:"DefaultBandWidth,omitempty" name:"DefaultBandWidth"`

	// 默认镜像，如img-qsdf3ff2。
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// 机型ID。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 默认系统盘大小，单位：G，默认大小为50G。范围不得超过系统盘上下限制，详看DescribeConfig。
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// 默认数据盘大小，单位：G。范围不得超过数据盘范围大小，详看DescribeConfig。
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// 是否关闭IP直通。取值范围：
	// true：表示关闭IP直通
	// false：表示开通IP直通
	CloseIpDirect *bool `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`

	// 标签列表。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 模块默认安全组列表
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`

	// 默认入带宽，单位：M。范围不得超过带宽上下限，详看DescribeConfig。
	DefaultBandWidthIn *int64 `json:"DefaultBandWidthIn,omitempty" name:"DefaultBandWidthIn"`

	// 是否禁止分配外网IP
	DisableWanIp *bool `json:"DisableWanIp,omitempty" name:"DisableWanIp"`

	// 系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘信息。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

func (r *CreateModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleName")
	delete(f, "DefaultBandWidth")
	delete(f, "DefaultImageId")
	delete(f, "InstanceType")
	delete(f, "DefaultSystemDiskSize")
	delete(f, "DefaultDataDiskSize")
	delete(f, "CloseIpDirect")
	delete(f, "TagSpecification")
	delete(f, "SecurityGroups")
	delete(f, "DefaultBandWidthIn")
	delete(f, "DisableWanIp")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModuleResponseParams struct {
	// 模块ID，创建模块成功后分配给该模块的ID。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateModuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateModuleResponseParams `json:"Response"`
}

func (r *CreateModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkInterfaceRequestParams struct {
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// 新申请的内网IP地址个数，内网IP地址个数总和不能超过配额数。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// 指定绑定的安全组，例如：['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// 新申请的内网IP地址个数，内网IP地址个数总和不能超过配额数。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// 指定绑定的安全组，例如：['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NetworkInterfaceName")
	delete(f, "SubnetId")
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceDescription")
	delete(f, "SecondaryPrivateIpAddressCount")
	delete(f, "SecurityGroupIds")
	delete(f, "PrivateIpAddresses")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkInterfaceResponseParams struct {
	// 弹性网卡实例。
	NetworkInterface *NetworkInterface `json:"NetworkInterface,omitempty" name:"NetworkInterface"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkInterfaceResponseParams `json:"Response"`
}

func (r *CreateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteTableRequestParams struct {
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// ecm地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type CreateRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// ecm地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *CreateRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "RouteTableName")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteTableResponseParams struct {
	// 路由表对象
	RouteTable *RouteTable `json:"RouteTable,omitempty" name:"RouteTable"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateRouteTableResponseParams `json:"Response"`
}

func (r *CreateRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoutesRequestParams struct {
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 要创建的路由策略对象。
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

type CreateRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 要创建的路由策略对象。
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *CreateRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoutesResponseParams struct {
	// 新增的实例个数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 路由表对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteTableSet []*RouteTable `json:"RouteTableSet,omitempty" name:"RouteTableSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRoutesResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoutesResponseParams `json:"Response"`
}

func (r *CreateRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupPoliciesRequestParams struct {
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组规则集合。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

type CreateSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组规则集合。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupRequestParams struct {
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupResponseParams struct {
	// 安全组对象。
	SecurityGroup *SecurityGroup `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetRequestParams struct {
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网网段，子网网段必须在VPC网段内，相同VPC内子网网段不能重叠。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 子网所在的可用区ID，不同子网选择不同可用区可以做跨可用区灾备。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateSubnetRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网网段，子网网段必须在VPC网段内，相同VPC内子网网段不能重叠。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 子网所在的可用区ID，不同子网选择不同可用区可以做跨可用区灾备。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "SubnetName")
	delete(f, "CidrBlock")
	delete(f, "Zone")
	delete(f, "EcmRegion")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetResponseParams struct {
	// 子网对象。
	Subnet *Subnet `json:"Subnet,omitempty" name:"Subnet"`

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
type CreateVpcRequestParams struct {
	// vpc名称，最大长度不能超过60个字节。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// vpc的cidr，只能为10.*.0.0/16，172.[16-31].0.0/16，192.168.0.0/16这三个内网网段内。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 是否开启组播。true: 开启, false: 不开启。暂不支持
	EnableMulticast *string `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// DNS地址，最多支持4个，暂不支持
	DnsServers []*string `json:"DnsServers,omitempty" name:"DnsServers"`

	// 域名，暂不支持
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateVpcRequest struct {
	*tchttp.BaseRequest
	
	// vpc名称，最大长度不能超过60个字节。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// vpc的cidr，只能为10.*.0.0/16，172.[16-31].0.0/16，192.168.0.0/16这三个内网网段内。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 是否开启组播。true: 开启, false: 不开启。暂不支持
	EnableMulticast *string `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// DNS地址，最多支持4个，暂不支持
	DnsServers []*string `json:"DnsServers,omitempty" name:"DnsServers"`

	// 域名，暂不支持
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
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
	delete(f, "EcmRegion")
	delete(f, "EnableMulticast")
	delete(f, "DnsServers")
	delete(f, "DomainName")
	delete(f, "Tags")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcResponseParams struct {
	// Vpc对象。
	Vpc *VpcInfo `json:"Vpc,omitempty" name:"Vpc"`

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

type DataDisk struct {
	// 数据盘大小。单位GB。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 数据盘类型，取值范围：
	// - LOCAL_BASIC：本地硬盘
	// - CLOUD_PREMIUM：高性能云硬盘
	// 
	// 默认取值： LOCAL_BASIC。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

// Predefined struct for user
type DeleteHaVipRequestParams struct {
	// HAVIP唯一ID，形如：havip-9o233uri。
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

type DeleteHaVipRequest struct {
	*tchttp.BaseRequest
	
	// HAVIP唯一ID，形如：havip-9o233uri。
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

func (r *DeleteHaVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHaVipResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteHaVipResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHaVipResponseParams `json:"Response"`
}

func (r *DeleteHaVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageRequestParams struct {
	// 镜像ID列表。
	ImageIDSet []*string `json:"ImageIDSet,omitempty" name:"ImageIDSet"`
}

type DeleteImageRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID列表。
	ImageIDSet []*string `json:"ImageIDSet,omitempty" name:"ImageIDSet"`
}

func (r *DeleteImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIDSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageResponseParams `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要删除的监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要删除的监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

func (r *DeleteListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteListenerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteListenerResponseParams `json:"Response"`
}

func (r *DeleteListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerListenersRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 指定删除的监听器ID数组，若不填则删除负载均衡的所有监听器
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

type DeleteLoadBalancerListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 指定删除的监听器ID数组，若不填则删除负载均衡的所有监听器
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *DeleteLoadBalancerListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerListenersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerListenersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerListenersResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancerListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerRequestParams struct {
	// 要删除的负载均衡实例 ID数组，数组大小最大支持20
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的负载均衡实例 ID数组，数组大小最大支持20
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModuleRequestParams struct {
	// 模块ID。如：em-qn46snq8
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

type DeleteModuleRequest struct {
	*tchttp.BaseRequest
	
	// 模块ID。如：em-qn46snq8
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *DeleteModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteModuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModuleResponseParams `json:"Response"`
}

func (r *DeleteModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkInterfaceRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DeleteNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DeleteNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkInterfaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetworkInterfaceResponseParams `json:"Response"`
}

func (r *DeleteNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTableRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type DeleteRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *DeleteRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTableResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRouteTableResponseParams `json:"Response"`
}

func (r *DeleteRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoutesRequestParams struct {
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略对象。
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

type DeleteRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略对象。
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *DeleteRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoutesResponseParams `json:"Response"`
}

func (r *DeleteRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupPoliciesRequestParams struct {
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组规则集合。一个请求中只能删除单个方向的一条或多条规则。支持指定索引（PolicyIndex） 匹配删除和安全组规则匹配删除两种方式，一个请求中只能使用一种匹配方式。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

type DeleteSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组规则集合。一个请求中只能删除单个方向的一条或多条规则。支持指定索引（PolicyIndex） 匹配删除和安全组规则匹配删除两种方式，一个请求中只能使用一种匹配方式。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *DeleteSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *DeleteSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupRequestParams struct {
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

type DeleteSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityGroupResponseParams `json:"Response"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsRequestParams struct {
	// 要删除的快照ID列表，可通过[DescribeSnapshots](/document/product/362/15647)查询。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// 是否强制删除快照关联的镜像
	DeleteBindImages *bool `json:"DeleteBindImages,omitempty" name:"DeleteBindImages"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的快照ID列表，可通过[DescribeSnapshots](/document/product/362/15647)查询。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// 是否强制删除快照关联的镜像
	DeleteBindImages *bool `json:"DeleteBindImages,omitempty" name:"DeleteBindImages"`
}

func (r *DeleteSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	delete(f, "DeleteBindImages")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotsResponseParams `json:"Response"`
}

func (r *DeleteSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubnetRequestParams struct {
	// 子网实例ID。可通过DescribeSubnets接口返回值中的SubnetId获取。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DeleteSubnetRequest struct {
	*tchttp.BaseRequest
	
	// 子网实例ID。可通过DescribeSubnets接口返回值中的SubnetId获取。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
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
	delete(f, "SubnetId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubnetResponseParams struct {
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
type DeleteVpcRequestParams struct {
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DeleteVpcRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
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
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcResponseParams struct {
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
type DescribeAddressQuotaRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DescribeAddressQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressQuotaResponseParams struct {
	// 账户 EIP 配额信息。
	QuotaSet []*EipQuota `json:"QuotaSet,omitempty" name:"QuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAddressQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressQuotaResponseParams `json:"Response"`
}

func (r *DescribeAddressQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressesRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：eip-11112222。参数不支持同时指定AddressIds和Filters。
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`

	// 每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定AddressIds和Filters。详细的过滤条件如下：
	// address-id - String - 是否必填：否 - （过滤条件）按照 EIP 的唯一 ID 过滤。EIP 唯一 ID 形如：eip-11112222。
	// address-name - String - 是否必填：否 - （过滤条件）按照 EIP 名称过滤。不支持模糊过滤。
	// address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。
	// address-status - String - 是否必填：否 - （过滤条件）按照 EIP 的状态过滤。取值范围：详见EIP状态列表。
	// instance-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的实例 ID 过滤。实例 ID 形如：ins-11112222。
	// private-ip-address - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的内网 IP 过滤。
	// network-interface-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的弹性网卡 ID 过滤。弹性网卡 ID 形如：eni-11112222。
	// is-arrears - String - 是否必填：否 - （过滤条件）按照 EIP 是否欠费进行过滤。（TRUE：EIP 处于欠费状态|FALSE：EIP 费用状态正常）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：eip-11112222。参数不支持同时指定AddressIds和Filters。
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`

	// 每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定AddressIds和Filters。详细的过滤条件如下：
	// address-id - String - 是否必填：否 - （过滤条件）按照 EIP 的唯一 ID 过滤。EIP 唯一 ID 形如：eip-11112222。
	// address-name - String - 是否必填：否 - （过滤条件）按照 EIP 名称过滤。不支持模糊过滤。
	// address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。
	// address-status - String - 是否必填：否 - （过滤条件）按照 EIP 的状态过滤。取值范围：详见EIP状态列表。
	// instance-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的实例 ID 过滤。实例 ID 形如：ins-11112222。
	// private-ip-address - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的内网 IP 过滤。
	// network-interface-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的弹性网卡 ID 过滤。弹性网卡 ID 形如：eni-11112222。
	// is-arrears - String - 是否必填：否 - （过滤条件）按照 EIP 是否欠费进行过滤。（TRUE：EIP 处于欠费状态|FALSE：EIP 费用状态正常）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressesResponseParams struct {
	// 符合条件的 EIP 数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// EIP 详细信息列表。
	AddressSet []*Address `json:"AddressSet,omitempty" name:"AddressSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAddressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressesResponseParams `json:"Response"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseOverviewRequestParams struct {

}

type DescribeBaseOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBaseOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaseOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseOverviewResponseParams struct {
	// 模块数量，单位：个
	ModuleNum *int64 `json:"ModuleNum,omitempty" name:"ModuleNum"`

	// 节点数量，单位：个
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// cpu核数，单位：个
	VcpuNum *int64 `json:"VcpuNum,omitempty" name:"VcpuNum"`

	// 内存大小，单位：G
	MemoryNum *int64 `json:"MemoryNum,omitempty" name:"MemoryNum"`

	// 硬盘大小，单位：G
	StorageNum *int64 `json:"StorageNum,omitempty" name:"StorageNum"`

	// 昨日网络峰值,单位：M
	NetworkNum *int64 `json:"NetworkNum,omitempty" name:"NetworkNum"`

	// 实例数量，单位：台
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 运行中数量，单位：台
	RunningNum *int64 `json:"RunningNum,omitempty" name:"RunningNum"`

	// 安全隔离数量，单位：台
	IsolationNum *int64 `json:"IsolationNum,omitempty" name:"IsolationNum"`

	// 过期实例数量，单位：台
	ExpiredNum *int64 `json:"ExpiredNum,omitempty" name:"ExpiredNum"`

	// 即将过期实例数量，单位：台
	WillExpireNum *int64 `json:"WillExpireNum,omitempty" name:"WillExpireNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBaseOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaseOverviewResponseParams `json:"Response"`
}

func (r *DescribeBaseOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRequestParams struct {

}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigResponseParams struct {
	// 网络带宽硬盘大小的范围信息。
	NetworkStorageRange *NetworkStorageRange `json:"NetworkStorageRange,omitempty" name:"NetworkStorageRange"`

	// 镜像操作系统白名单。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageWhiteSet []*string `json:"ImageWhiteSet,omitempty" name:"ImageWhiteSet"`

	// 网络限额信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNetworkLimitConfigs []*InstanceNetworkLimitConfig `json:"InstanceNetworkLimitConfigs,omitempty" name:"InstanceNetworkLimitConfigs"`

	// 镜像限额信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageLimits *ImageLimitConfig `json:"ImageLimits,omitempty" name:"ImageLimits"`

	// 默认是否IP直通，用于模块创建，虚机购买等具有直通参数场景时的默认参数。
	DefaultIPDirect *bool `json:"DefaultIPDirect,omitempty" name:"DefaultIPDirect"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigResponseParams `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomImageTaskRequestParams struct {
	// 支持key,value查询
	// task-id: 异步任务ID
	// image-id: 镜像ID
	// image-name: 镜像名称
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCustomImageTaskRequest struct {
	*tchttp.BaseRequest
	
	// 支持key,value查询
	// task-id: 异步任务ID
	// image-id: 镜像ID
	// image-name: 镜像名称
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCustomImageTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomImageTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomImageTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomImageTaskResponseParams struct {
	// 导入任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTaskSet []*ImageTask `json:"ImageTaskSet,omitempty" name:"ImageTaskSet"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomImageTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomImageTaskResponseParams `json:"Response"`
}

func (r *DescribeCustomImageTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomImageTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultSubnetRequestParams struct {
	// ECM地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ECM可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeDefaultSubnetRequest struct {
	*tchttp.BaseRequest
	
	// ECM地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ECM可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeDefaultSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultSubnetResponseParams struct {
	// 默认子网信息，若无子网，则为空数据。
	Subnet *Subnet `json:"Subnet,omitempty" name:"Subnet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDefaultSubnetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultSubnetResponseParams `json:"Response"`
}

func (r *DescribeDefaultSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksRequestParams struct {
	// 过滤条件。参数不支持同时指定`DiskIds`和`Filters`。<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按云盘类型过滤。 (SYSTEM_DISK：表示系统盘 | DATA_DISK：表示数据盘)<br><li>disk-charge-type - Array of String - 是否必填：否 -（过滤条件）按照云硬盘计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费。)<br><li>portable - Array of String - 是否必填：否 -（过滤条件）按是否为弹性云盘过滤。 (TRUE：表示弹性云盘 | FALSE：表示非弹性云盘。)<br><li>project-id - Array of Integer - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘ID过滤。云盘ID形如：`disk-11112222`。<br><li>disk-name - Array of String - 是否必填：否 -（过滤条件）按照云盘名称过滤。<br><li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：表示SSD云硬盘 | CLOUD_HSSD：表示增强型SSD云硬盘。| CLOUD_TSSD：表示极速型云硬盘。)<br><li>disk-state - Array of String - 是否必填：否 -（过滤条件）按照云盘状态过滤。(UNATTACHED：未挂载 | ATTACHING：挂载中 | ATTACHED：已挂载 | DETACHING：解挂中 | EXPANDING：扩容中 | ROLLBACKING：回滚中 | TORECYCLE：待回收。)<br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按照云盘挂载的云主机实例ID过滤。可根据此参数查询挂载在指定云主机下的云硬盘。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/product/213/15753#ZoneInfo)过滤。<br><li>instance-ip-address - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载云主机的内网或外网IP过滤。<br><li>instance-name - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载的实例名称过滤。<br><li>tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键进行过滤。<br><li>tag-value - Array of String - 是否必填：否 -（过滤条件）照标签值进行过滤。<br><li>tag:tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 云盘列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云盘的创建时间排序<br><li>DEADLINE：依据云盘的到期时间排序<br>默认按云盘创建时间排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 云盘详情中是否需要返回云盘绑定的定期快照策略ID，TRUE表示需要返回，FALSE表示不返回。
	ReturnBindAutoSnapshotPolicy *bool `json:"ReturnBindAutoSnapshotPolicy,omitempty" name:"ReturnBindAutoSnapshotPolicy"`

	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：`disk-11112222`，此参数的具体格式可参考API[简介](/document/product/362/15633)的ids.N一节）。参数不支持同时指定`DiskIds`和`Filters`。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeDisksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。参数不支持同时指定`DiskIds`和`Filters`。<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按云盘类型过滤。 (SYSTEM_DISK：表示系统盘 | DATA_DISK：表示数据盘)<br><li>disk-charge-type - Array of String - 是否必填：否 -（过滤条件）按照云硬盘计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费。)<br><li>portable - Array of String - 是否必填：否 -（过滤条件）按是否为弹性云盘过滤。 (TRUE：表示弹性云盘 | FALSE：表示非弹性云盘。)<br><li>project-id - Array of Integer - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘ID过滤。云盘ID形如：`disk-11112222`。<br><li>disk-name - Array of String - 是否必填：否 -（过滤条件）按照云盘名称过滤。<br><li>disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：表示SSD云硬盘 | CLOUD_HSSD：表示增强型SSD云硬盘。| CLOUD_TSSD：表示极速型云硬盘。)<br><li>disk-state - Array of String - 是否必填：否 -（过滤条件）按照云盘状态过滤。(UNATTACHED：未挂载 | ATTACHING：挂载中 | ATTACHED：已挂载 | DETACHING：解挂中 | EXPANDING：扩容中 | ROLLBACKING：回滚中 | TORECYCLE：待回收。)<br><li>instance-id - Array of String - 是否必填：否 -（过滤条件）按照云盘挂载的云主机实例ID过滤。可根据此参数查询挂载在指定云主机下的云硬盘。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/product/213/15753#ZoneInfo)过滤。<br><li>instance-ip-address - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载云主机的内网或外网IP过滤。<br><li>instance-name - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载的实例名称过滤。<br><li>tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键进行过滤。<br><li>tag-value - Array of String - 是否必填：否 -（过滤条件）照标签值进行过滤。<br><li>tag:tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 云盘列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据云盘的创建时间排序<br><li>DEADLINE：依据云盘的到期时间排序<br>默认按云盘创建时间排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 云盘详情中是否需要返回云盘绑定的定期快照策略ID，TRUE表示需要返回，FALSE表示不返回。
	ReturnBindAutoSnapshotPolicy *bool `json:"ReturnBindAutoSnapshotPolicy,omitempty" name:"ReturnBindAutoSnapshotPolicy"`

	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：`disk-11112222`，此参数的具体格式可参考API[简介](/document/product/362/15633)的ids.N一节）。参数不支持同时指定`DiskIds`和`Filters`。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Offset")
	delete(f, "ReturnBindAutoSnapshotPolicy")
	delete(f, "DiskIds")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDisksResponseParams struct {
	// 符合条件的云硬盘数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 云硬盘的详细信息列表。
	DiskSet []*Disk `json:"DiskSet,omitempty" name:"DiskSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDisksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDisksResponseParams `json:"Response"`
}

func (r *DescribeDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHaVipsRequestParams struct {
	// HAVIP数组，HAVIP唯一ID，形如：havip-9o233uri。
	HaVipIds []*string `json:"HaVipIds,omitempty" name:"HaVipIds"`

	// 过滤条件，参数不支持同时指定HaVipIds和Filters。
	// havip-id - String - HAVIP唯一ID，形如：havip-9o233uri。
	// havip-name - String - HAVIP名称。
	// vpc-id - String - HAVIP所在私有网络ID。
	// subnet-id - String - HAVIP所在子网ID。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认值是0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值是20，最大是100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Ecm 区域，不填代表全部区域。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DescribeHaVipsRequest struct {
	*tchttp.BaseRequest
	
	// HAVIP数组，HAVIP唯一ID，形如：havip-9o233uri。
	HaVipIds []*string `json:"HaVipIds,omitempty" name:"HaVipIds"`

	// 过滤条件，参数不支持同时指定HaVipIds和Filters。
	// havip-id - String - HAVIP唯一ID，形如：havip-9o233uri。
	// havip-name - String - HAVIP名称。
	// vpc-id - String - HAVIP所在私有网络ID。
	// subnet-id - String - HAVIP所在子网ID。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认值是0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值是20，最大是100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Ecm 区域，不填代表全部区域。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DescribeHaVipsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHaVipsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHaVipsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHaVipsResponseParams struct {
	// 符合条件的对象数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// HAVIP对象数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HaVipSet []*HaVip `json:"HaVipSet,omitempty" name:"HaVipSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHaVipsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHaVipsResponseParams `json:"Response"`
}

func (r *DescribeHaVipsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHaVipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRequestParams struct {
	// 过滤条件，每次请求的Filters的上限为10，详细的过滤条件如下：
	// image-id - String - 是否必填： 否 - （过滤条件）按照镜像ID进行过滤
	// image-type - String - 是否必填： 否 - （过滤条件）按照镜像类型进行过滤。取值范围：
	// PRIVATE_IMAGE: 私有镜像 (本帐户创建的镜像) 
	// PUBLIC_IMAGE: 公共镜像 (腾讯云官方镜像)
	// instance-type -String - 是否必填: 否 - (过滤条件) 按机型过滤支持的镜像
	// image-name - String - 是否必填：否 - (过滤条件) 按镜像的名称模糊匹配，只能提供一个值
	// image-os - String - 是否必填：否 - (过滤条件) 按镜像系统的名称模糊匹配，只能提供一个值
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeImageRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，每次请求的Filters的上限为10，详细的过滤条件如下：
	// image-id - String - 是否必填： 否 - （过滤条件）按照镜像ID进行过滤
	// image-type - String - 是否必填： 否 - （过滤条件）按照镜像类型进行过滤。取值范围：
	// PRIVATE_IMAGE: 私有镜像 (本帐户创建的镜像) 
	// PUBLIC_IMAGE: 公共镜像 (腾讯云官方镜像)
	// instance-type -String - 是否必填: 否 - (过滤条件) 按机型过滤支持的镜像
	// image-name - String - 是否必填：否 - (过滤条件) 按镜像的名称模糊匹配，只能提供一个值
	// image-os - String - 是否必填：否 - (过滤条件) 按镜像系统的名称模糊匹配，只能提供一个值
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageResponseParams struct {
	// 镜像总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageResponseParams `json:"Response"`
}

func (r *DescribeImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImportImageOsRequestParams struct {

}

type DescribeImportImageOsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImportImageOsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportImageOsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImportImageOsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImportImageOsResponseParams struct {
	// 支持的导入镜像的操作系统类型
	ImportImageOsListSupported *ImageOsList `json:"ImportImageOsListSupported,omitempty" name:"ImportImageOsListSupported"`

	// 支持的导入镜像的操作系统版本
	ImportImageOsVersionSet []*OsVersion `json:"ImportImageOsVersionSet,omitempty" name:"ImportImageOsVersionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImportImageOsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImportImageOsResponseParams `json:"Response"`
}

func (r *DescribeImportImageOsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImportImageOsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTypeConfigRequestParams struct {

}

type DescribeInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceTypeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceTypeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTypeConfigResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 机型配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceTypeConfigResponseParams `json:"Response"`
}

func (r *DescribeInstanceTypeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceVncUrlRequestParams struct {
	// 一个操作的实例ID。可通过DescribeInstances API返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceVncUrlRequest struct {
	*tchttp.BaseRequest
	
	// 一个操作的实例ID。可通过DescribeInstances API返回值中的InstanceId获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceVncUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceVncUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceVncUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceVncUrlResponseParams struct {
	// 实例的管理终端地址。
	InstanceVncUrl *string `json:"InstanceVncUrl,omitempty" name:"InstanceVncUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceVncUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceVncUrlResponseParams `json:"Response"`
}

func (r *DescribeInstanceVncUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceVncUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDeniedActionsRequestParams struct {
	// 无
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type DescribeInstancesDeniedActionsRequest struct {
	*tchttp.BaseRequest
	
	// 无
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *DescribeInstancesDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDeniedActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesDeniedActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesDeniedActionsResponseParams struct {
	// 实例对应的禁止操作
	InstanceOperatorSet []*InstanceOperator `json:"InstanceOperatorSet,omitempty" name:"InstanceOperatorSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesDeniedActionsResponseParams `json:"Response"`
}

func (r *DescribeInstancesDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesDeniedActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 过滤条件。
	// zone      String      是否必填：否     （过滤条件）按照可用区英文标识符过滤。
	// zone-name      String      是否必填：否     （过滤条件）按照可用区中文名过滤,支持模糊匹配。
	// module-id      String      是否必填：否     （过滤条件）按照模块ID过滤。
	// instance-id      String      是否必填：否      （过滤条件）按照实例ID过滤。
	// instance-name      String      是否必填：否      （过滤条件）按照实例名称过滤,支持模糊匹配。
	// ip-address      String      是否必填：否      （过滤条件）按照实例的内网/公网IP过滤。
	// instance-uuid   string 是否必填：否 （过滤条件）按照uuid过滤实例列表。
	// instance-state  string  是否必填：否 （过滤条件）按照实例状态更新实例列表。
	// internet-service-provider      String      是否必填：否      （过滤条件）按照实例公网IP所属的运营商进行过滤。
	// tag-key      String      是否必填：否      （过滤条件）按照标签键进行过滤。
	// tag:tag-key      String      是否必填：否      （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	// instance-family      String      是否必填：否      （过滤条件）按照机型family过滤。
	// module-name      String      是否必填：否      （过滤条件）按照模块名称过滤,支持模糊匹配。
	// image-id      String      是否必填：否      （过滤条件）按照实例的镜像ID过滤。
	// vpc-id String      是否必填：否      （过滤条件）按照实例的vpc id过滤。
	// subnet-id String      是否必填：否      （过滤条件）按照实例的subnet id过滤。
	// 
	// 若不传Filters参数则表示查询所有相关的实例信息。
	// 单次请求的Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20(如果查询结果数目大于等于20)，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定排序字段。目前支持的可选值如下
	// timestamp 按实例创建时间排序。
	// 注意：目前仅支持按创建时间排序，后续可能会有扩展。
	// 如果不传，默认按实例创建时间排序
	OrderByField *string `json:"OrderByField,omitempty" name:"OrderByField"`

	// 指定排序是降序还是升序。0表示降序； 1表示升序。如果不传默认为降序
	OrderDirection *int64 `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// zone      String      是否必填：否     （过滤条件）按照可用区英文标识符过滤。
	// zone-name      String      是否必填：否     （过滤条件）按照可用区中文名过滤,支持模糊匹配。
	// module-id      String      是否必填：否     （过滤条件）按照模块ID过滤。
	// instance-id      String      是否必填：否      （过滤条件）按照实例ID过滤。
	// instance-name      String      是否必填：否      （过滤条件）按照实例名称过滤,支持模糊匹配。
	// ip-address      String      是否必填：否      （过滤条件）按照实例的内网/公网IP过滤。
	// instance-uuid   string 是否必填：否 （过滤条件）按照uuid过滤实例列表。
	// instance-state  string  是否必填：否 （过滤条件）按照实例状态更新实例列表。
	// internet-service-provider      String      是否必填：否      （过滤条件）按照实例公网IP所属的运营商进行过滤。
	// tag-key      String      是否必填：否      （过滤条件）按照标签键进行过滤。
	// tag:tag-key      String      是否必填：否      （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	// instance-family      String      是否必填：否      （过滤条件）按照机型family过滤。
	// module-name      String      是否必填：否      （过滤条件）按照模块名称过滤,支持模糊匹配。
	// image-id      String      是否必填：否      （过滤条件）按照实例的镜像ID过滤。
	// vpc-id String      是否必填：否      （过滤条件）按照实例的vpc id过滤。
	// subnet-id String      是否必填：否      （过滤条件）按照实例的subnet id过滤。
	// 
	// 若不传Filters参数则表示查询所有相关的实例信息。
	// 单次请求的Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20(如果查询结果数目大于等于20)，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定排序字段。目前支持的可选值如下
	// timestamp 按实例创建时间排序。
	// 注意：目前仅支持按创建时间排序，后续可能会有扩展。
	// 如果不传，默认按实例创建时间排序
	OrderByField *string `json:"OrderByField,omitempty" name:"OrderByField"`

	// 指定排序是降序还是升序。0表示降序； 1表示升序。如果不传默认为降序
	OrderDirection *int64 `json:"OrderDirection,omitempty" name:"OrderDirection"`
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
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 返回的实例相关信息列表的长度。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的实例相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeListenersRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要查询的负载均衡监听器 ID数组
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 要查询的监听器协议类型，取值 TCP | UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 要查询的监听器的端口
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要查询的负载均衡监听器 ID数组
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 要查询的监听器协议类型，取值 TCP | UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 要查询的监听器的端口
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersResponseParams struct {
	// 监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listeners []*Listener `json:"Listeners,omitempty" name:"Listeners"`

	// 总的监听器个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenersResponseParams `json:"Response"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalanceTaskStatusRequestParams struct {
	// 请求ID，即接口返回的 RequestId 参数
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeLoadBalanceTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 请求ID，即接口返回的 RequestId 参数
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeLoadBalanceTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalanceTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalanceTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalanceTaskStatusResponseParams struct {
	// 任务的当前状态。 0：成功，1：失败，2：进行中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalanceTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalanceTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeLoadBalanceTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalanceTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersRequestParams struct {
	// 区域。如果不传则默认查询所有区域。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 负载均衡实例 ID。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例的 VIP 地址，支持多个。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 负载均衡绑定的后端服务的内网 IP。
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitempty" name:"BackendPrivateIps"`

	// 数据偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 负载均衡是否绑定后端服务，0：没有绑定后端服务，1：绑定后端服务，-1：查询全部。 
	// 如果不传则默认查询全部。
	WithBackend *int64 `json:"WithBackend,omitempty" name:"WithBackend"`

	// 负载均衡实例所属私有网络唯一ID，如 vpc-bhqkbhdx。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。详细的过滤条件如下：
	// tag-key - String - 是否必填：否 - （过滤条件）按照标签的键过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 安全组。
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 区域。如果不传则默认查询所有区域。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 负载均衡实例 ID。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例的 VIP 地址，支持多个。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 负载均衡绑定的后端服务的内网 IP。
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitempty" name:"BackendPrivateIps"`

	// 数据偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 负载均衡是否绑定后端服务，0：没有绑定后端服务，1：绑定后端服务，-1：查询全部。 
	// 如果不传则默认查询全部。
	WithBackend *int64 `json:"WithBackend,omitempty" name:"WithBackend"`

	// 负载均衡实例所属私有网络唯一ID，如 vpc-bhqkbhdx。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。详细的过滤条件如下：
	// tag-key - String - 是否必填：否 - （过滤条件）按照标签的键过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 安全组。
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "LoadBalancerIds")
	delete(f, "LoadBalancerName")
	delete(f, "LoadBalancerVips")
	delete(f, "BackendPrivateIps")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "WithBackend")
	delete(f, "VpcId")
	delete(f, "Filters")
	delete(f, "SecurityGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersResponseParams struct {
	// 满足过滤条件的负载均衡实例总数。此数值与入参中的Limit无关。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的负载均衡实例数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancersResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleDetailRequestParams struct {
	// 模块ID，如em-qn46snq8。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

type DescribeModuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 模块ID，如em-qn46snq8。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *DescribeModuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleDetailResponseParams struct {
	// 模块的详细信息，详细见数据结构中的ModuleInfo。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Module *Module `json:"Module,omitempty" name:"Module"`

	// 模块的统计信息，详细见数据结构中的ModuleCounterInfo。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleCounter *ModuleCounter `json:"ModuleCounter,omitempty" name:"ModuleCounter"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeModuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModuleDetailResponseParams `json:"Response"`
}

func (r *DescribeModuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleRequestParams struct {
	// 过滤条件。
	// module-name - string - 是否必填：否 - （过滤条件）按照模块名称过滤。
	// module-id - string - 是否必填：否 - （过滤条件）按照模块ID过滤。
	// image-id      String      是否必填：否      （过滤条件）按照镜像ID过滤。
	// instance-family      String      是否必填：否      （过滤条件）按照机型family过滤。
	// security-group-id - string 是否必填：否 - （过滤条件）按照模块绑定的安全组id过滤。
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定排序字段。目前支持的可选值如下
	// instance-num 按实例数量排序。
	// node-num 按节点数量排序。
	// timestamp 按实例创建时间排序。
	// 如果不传，默认按实例创建时间排序
	OrderByField *string `json:"OrderByField,omitempty" name:"OrderByField"`

	// 指定排序是降序还是升序。0表示降序； 1表示升序。如果不传默认为降序
	OrderDirection *int64 `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeModuleRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// module-name - string - 是否必填：否 - （过滤条件）按照模块名称过滤。
	// module-id - string - 是否必填：否 - （过滤条件）按照模块ID过滤。
	// image-id      String      是否必填：否      （过滤条件）按照镜像ID过滤。
	// instance-family      String      是否必填：否      （过滤条件）按照机型family过滤。
	// security-group-id - string 是否必填：否 - （过滤条件）按照模块绑定的安全组id过滤。
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 指定排序字段。目前支持的可选值如下
	// instance-num 按实例数量排序。
	// node-num 按节点数量排序。
	// timestamp 按实例创建时间排序。
	// 如果不传，默认按实例创建时间排序
	OrderByField *string `json:"OrderByField,omitempty" name:"OrderByField"`

	// 指定排序是降序还是升序。0表示降序； 1表示升序。如果不传默认为降序
	OrderDirection *int64 `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleResponseParams struct {
	// 符合条件的模块数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 模块详情信息的列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleItemSet []*ModuleItem `json:"ModuleItemSet,omitempty" name:"ModuleItemSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeModuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModuleResponseParams `json:"Response"`
}

func (r *DescribeModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonthPeakNetworkRequestParams struct {
	// 月份时间(xxxx-xx) 如2021-03,默认取当前时间的上一个月份
	Month *string `json:"Month,omitempty" name:"Month"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeMonthPeakNetworkRequest struct {
	*tchttp.BaseRequest
	
	// 月份时间(xxxx-xx) 如2021-03,默认取当前时间的上一个月份
	Month *string `json:"Month,omitempty" name:"Month"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMonthPeakNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonthPeakNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonthPeakNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonthPeakNetworkResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonthNetWorkData []*MonthNetwork `json:"MonthNetWorkData,omitempty" name:"MonthNetWorkData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMonthPeakNetworkResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonthPeakNetworkResponseParams `json:"Response"`
}

func (r *DescribeMonthPeakNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonthPeakNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfacesRequestParams struct {
	// 弹性网卡实例ID查询。形如：eni-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定NetworkInterfaceIds和Filters。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds"`

	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。
	// network-interface-id - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。
	// attachment.instance-id - String - （过滤条件）绑定的云服务器实例ID，形如：ein-3nqpdn3i。
	// groups.security-group-id - String - （过滤条件）绑定的安全组实例ID，例如：sg-f9ekbxeq。
	// network-interface-name - String - （过滤条件）网卡实例名称。
	// network-interface-description - String - （过滤条件）网卡实例描述。
	// address-ip - String - （过滤条件）内网IPv4地址。
	// tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。使用请参考示例2
	// tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例3。
	// is-primary - Boolean - 是否必填：否 - （过滤条件）按照是否主网卡进行过滤。值为true时，仅过滤主网卡；值为false时，仅过滤辅助网卡；次过滤参数为提供时，同时过滤主网卡和辅助网卡。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DescribeNetworkInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID查询。形如：eni-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定NetworkInterfaceIds和Filters。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds"`

	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。
	// network-interface-id - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。
	// attachment.instance-id - String - （过滤条件）绑定的云服务器实例ID，形如：ein-3nqpdn3i。
	// groups.security-group-id - String - （过滤条件）绑定的安全组实例ID，例如：sg-f9ekbxeq。
	// network-interface-name - String - （过滤条件）网卡实例名称。
	// network-interface-description - String - （过滤条件）网卡实例描述。
	// address-ip - String - （过滤条件）内网IPv4地址。
	// tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。使用请参考示例2
	// tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例3。
	// is-primary - Boolean - 是否必填：否 - （过滤条件）按照是否主网卡进行过滤。值为true时，仅过滤主网卡；值为false时，仅过滤辅助网卡；次过滤参数为提供时，同时过滤主网卡和辅助网卡。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DescribeNetworkInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfacesResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例详细信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet,omitempty" name:"NetworkInterfaceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkInterfacesResponseParams `json:"Response"`
}

func (r *DescribeNetworkInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeRequestParams struct {
	// 过滤条件，name取值为： InstanceFamily-实例系列
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeNodeRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，name取值为： InstanceFamily-实例系列
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeResponseParams struct {
	// 节点详细信息的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSet []*Node `json:"NodeSet,omitempty" name:"NodeSet"`

	// 所有的节点数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeResponseParams `json:"Response"`
}

func (r *DescribeNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackingQuotaGroupRequestParams struct {
	// 过滤条件，name取值为：Zone-可用区， InstanceType-实例类型，DataDiskSize - 数据盘大小
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePackingQuotaGroupRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，name取值为：Zone-可用区， InstanceType-实例类型，DataDiskSize - 数据盘大小
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePackingQuotaGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackingQuotaGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePackingQuotaGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackingQuotaGroupResponseParams struct {
	// 装箱配额组
	PackingQuotaSet []*PackingQuotaGroup `json:"PackingQuotaSet,omitempty" name:"PackingQuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePackingQuotaGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribePackingQuotaGroupResponseParams `json:"Response"`
}

func (r *DescribePackingQuotaGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackingQuotaGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakBaseOverviewRequestParams struct {
	// 开始时间（xxxx-xx-xx）如2019-08-14，默认为一周之前的日期，不应与当前日期间隔超过90天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（xxxx-xx-xx）如2019-08-14，默认为昨天，不应与当前日期间隔超过90天。当开始与结束间隔不超过30天时返回1小时粒度的数据，否则返回3小时粒度的数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribePeakBaseOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间（xxxx-xx-xx）如2019-08-14，默认为一周之前的日期，不应与当前日期间隔超过90天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（xxxx-xx-xx）如2019-08-14，默认为昨天，不应与当前日期间隔超过90天。当开始与结束间隔不超过30天时返回1小时粒度的数据，否则返回3小时粒度的数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribePeakBaseOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakBaseOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePeakBaseOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakBaseOverviewResponseParams struct {
	// 基础峰值列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeakFamilyInfoSet []*PeakFamilyInfo `json:"PeakFamilyInfoSet,omitempty" name:"PeakFamilyInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePeakBaseOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribePeakBaseOverviewResponseParams `json:"Response"`
}

func (r *DescribePeakBaseOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakBaseOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakNetworkOverviewRequestParams struct {
	// 开始时间（xxxx-xx-xx）如2019-08-14，默认为一周之前的日期，不应与当前日期间隔超过30天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（xxxx-xx-xx）如2019-08-14，默认为昨天，不应与当前日期间隔超过30天。当开始与结束间隔不超过1天时会返回1分钟粒度的数据，不超过7天时返回5分钟粒度的数据，否则返回1小时粒度的数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤条件。
	// 
	// region    String      是否必填：否     （过滤条件）按照region过滤，不支持模糊匹配。注意 region 填上需要查询ecm region才能返回数据。
	// area       String      是否必填：否     （过滤条件）按照大区过滤，不支持模糊匹配。大区包括：china-central、china-east等等，可以通过DescribeNode获得所有大区；也可使用ALL_REGION表示所有地区。
	// isp         String      是否必填：否     （过滤条件）按照运营商过滤大区流量，运营商包括CTCC、CUCC和CMCC。只和area同时使用，且一次只能指定一种运营商。
	// 
	// region和area只应填写一个。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 统计周期，单位秒。取值60/300。
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type DescribePeakNetworkOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间（xxxx-xx-xx）如2019-08-14，默认为一周之前的日期，不应与当前日期间隔超过30天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（xxxx-xx-xx）如2019-08-14，默认为昨天，不应与当前日期间隔超过30天。当开始与结束间隔不超过1天时会返回1分钟粒度的数据，不超过7天时返回5分钟粒度的数据，否则返回1小时粒度的数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤条件。
	// 
	// region    String      是否必填：否     （过滤条件）按照region过滤，不支持模糊匹配。注意 region 填上需要查询ecm region才能返回数据。
	// area       String      是否必填：否     （过滤条件）按照大区过滤，不支持模糊匹配。大区包括：china-central、china-east等等，可以通过DescribeNode获得所有大区；也可使用ALL_REGION表示所有地区。
	// isp         String      是否必填：否     （过滤条件）按照运营商过滤大区流量，运营商包括CTCC、CUCC和CMCC。只和area同时使用，且一次只能指定一种运营商。
	// 
	// region和area只应填写一个。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 统计周期，单位秒。取值60/300。
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribePeakNetworkOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakNetworkOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePeakNetworkOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakNetworkOverviewResponseParams struct {
	// 网络峰值数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeakNetworkRegionSet []*PeakNetworkRegionInfo `json:"PeakNetworkRegionSet,omitempty" name:"PeakNetworkRegionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePeakNetworkOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribePeakNetworkOverviewResponseParams `json:"Response"`
}

func (r *DescribePeakNetworkOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakNetworkOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePriceRunInstanceRequestParams struct {
	// 实例的机型信息
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 系统盘信息
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例个数
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 数据盘信息
	DataDisk []*DataDisk `json:"DataDisk,omitempty" name:"DataDisk"`

	// 实例计费类型。其中：
	// 0，按资源维度后付费，计算当日用量峰值，例如CPU，内存，硬盘等，仅适用于非GNR系列机型；
	// 1，按小时后付费，单价：xx元/实例/小时，仅适用于GNR机型，如需开通该计费方式请提工单申请；
	// 2，按月后付费，单价：xx元/实例/月，仅适用于GNR机型；
	// 该字段不填时，非GNR机型会默认选择0；GNR机型默认选择2。
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

type DescribePriceRunInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例的机型信息
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 系统盘信息
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例个数
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 数据盘信息
	DataDisk []*DataDisk `json:"DataDisk,omitempty" name:"DataDisk"`

	// 实例计费类型。其中：
	// 0，按资源维度后付费，计算当日用量峰值，例如CPU，内存，硬盘等，仅适用于非GNR系列机型；
	// 1，按小时后付费，单价：xx元/实例/小时，仅适用于GNR机型，如需开通该计费方式请提工单申请；
	// 2，按月后付费，单价：xx元/实例/月，仅适用于GNR机型；
	// 该字段不填时，非GNR机型会默认选择0；GNR机型默认选择2。
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *DescribePriceRunInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePriceRunInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "SystemDisk")
	delete(f, "InstanceCount")
	delete(f, "DataDisk")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePriceRunInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePriceRunInstanceResponseParams struct {
	// 实例价格信息
	InstancePrice *InstancesPrice `json:"InstancePrice,omitempty" name:"InstancePrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePriceRunInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribePriceRunInstanceResponseParams `json:"Response"`
}

func (r *DescribePriceRunInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePriceRunInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteConflictsRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 要检查的与之冲突的目的端列表
	DestinationCidrBlocks []*string `json:"DestinationCidrBlocks,omitempty" name:"DestinationCidrBlocks"`
}

type DescribeRouteConflictsRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 要检查的与之冲突的目的端列表
	DestinationCidrBlocks []*string `json:"DestinationCidrBlocks,omitempty" name:"DestinationCidrBlocks"`
}

func (r *DescribeRouteConflictsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteConflictsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "DestinationCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteConflictsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteConflictsResponseParams struct {
	// 路由策略冲突列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteConflictSet []*RouteConflict `json:"RouteConflictSet,omitempty" name:"RouteConflictSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRouteConflictsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteConflictsResponseParams `json:"Response"`
}

func (r *DescribeRouteConflictsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteConflictsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTablesRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableIds []*string `json:"RouteTableIds,omitempty" name:"RouteTableIds"`

	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// route-table-id - String - （过滤条件）路由表实例ID。
	// route-table-name - String - （过滤条件）路由表名称。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// association.main - String - （过滤条件）是否主路由表。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// ECM 地域，传空或不传表示所有区域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DescribeRouteTablesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableIds []*string `json:"RouteTableIds,omitempty" name:"RouteTableIds"`

	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// route-table-id - String - （过滤条件）路由表实例ID。
	// route-table-name - String - （过滤条件）路由表名称。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// association.main - String - （过滤条件）是否主路由表。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// ECM 地域，传空或不传表示所有区域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
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
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTablesResponseParams struct {
	// 符合条件的实例数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 路由表列表
	// 注意：此字段可能返回 null，表示取不到有效值。
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
type DescribeSecurityGroupAssociationStatisticsRequestParams struct {
	// 安全实例ID，例如esg-33ocnj9n，可通过[DescribeSecurityGroups](https://cloud.tencent.com/document/product/1108/47697)获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type DescribeSecurityGroupAssociationStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 安全实例ID，例如esg-33ocnj9n，可通过[DescribeSecurityGroups](https://cloud.tencent.com/document/product/1108/47697)获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *DescribeSecurityGroupAssociationStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupAssociationStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupAssociationStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupAssociationStatisticsResponseParams struct {
	// 安全组关联实例统计。
	SecurityGroupAssociationStatisticsSet []*SecurityGroupAssociationStatistics `json:"SecurityGroupAssociationStatisticsSet,omitempty" name:"SecurityGroupAssociationStatisticsSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupAssociationStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupAssociationStatisticsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupAssociationStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupAssociationStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupLimitsRequestParams struct {

}

type DescribeSecurityGroupLimitsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecurityGroupLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupLimitsResponseParams struct {
	// 用户安全组配额限制。
	SecurityGroupLimitSet *SecurityGroupLimitSet `json:"SecurityGroupLimitSet,omitempty" name:"SecurityGroupLimitSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupLimitsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupPoliciesRequestParams struct {
	// 安全组实例ID，例如：esg-33ocnj9n，可通过[DescribeSecurityGroups](https://cloud.tencent.com/document/product/1108/47697)获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

type DescribeSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如：esg-33ocnj9n，可通过[DescribeSecurityGroups](https://cloud.tencent.com/document/product/1108/47697)获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupPoliciesResponseParams struct {
	// 安全组规则集合。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsRequestParams struct {
	// 安全组实例ID，例如：esg-33ocnj9n，可通过[DescribeSecurityGroups](https://cloud.tencent.com/document/product/1108/47697)获取。每次请求的实例的上限为100。参数不支持同时指定SecurityGroupIds和Filters。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 过滤条件，参数不支持同时指定SecurityGroupIds和Filters。
	// security-group-id - String - （过滤条件）安全组ID。
	// security-group-name - String - （过滤条件）安全组名称。
	// tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。
	// tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如：esg-33ocnj9n，可通过[DescribeSecurityGroups](https://cloud.tencent.com/document/product/1108/47697)获取。每次请求的实例的上限为100。参数不支持同时指定SecurityGroupIds和Filters。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 过滤条件，参数不支持同时指定SecurityGroupIds和Filters。
	// security-group-id - String - （过滤条件）安全组ID。
	// security-group-name - String - （过滤条件）安全组名称。
	// tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。
	// tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 安全组对象。
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitempty" name:"SecurityGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsRequestParams struct {
	// 要查询快照的ID列表。参数不支持同时指定`SnapshotIds`和`Filters`。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// 过滤条件。参数不支持同时指定`SnapshotIds`和`Filters`。<br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按照快照的ID过滤。快照ID形如：`snap-11112222`。<br><li>snapshot-name - Array of String - 是否必填：否 -（过滤条件）按照快照名称过滤。<br><li>snapshot-state - Array of String - 是否必填：否 -（过滤条件）按照快照状态过滤。 (NORMAL：正常 | CREATING：创建中 | ROLLBACKING：回滚中。)<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建快照的云盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)<br><li>project-id  - Array of String - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id  - Array of String - 是否必填：否 -（过滤条件）按照创建快照的云硬盘ID过滤。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/product/213/15753#ZoneInfo)过滤。<br><li>encrypt - Array of String - 是否必填：否 -（过滤条件）按是否加密盘快照过滤。 (TRUE：表示加密盘快照 | FALSE：表示非加密盘快照。)
	// <li>snapshot-type- Array of String - 是否必填：否 -（过滤条件）根据snapshot-type指定的快照类型查询对应的快照。
	// (SHARED_SNAPSHOT：表示共享过来的快照 | PRIVATE_SNAPSHOT：表示自己私有快照。)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 快照列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据快照的创建时间排序<br>默认按创建时间排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询快照的ID列表。参数不支持同时指定`SnapshotIds`和`Filters`。
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`

	// 过滤条件。参数不支持同时指定`SnapshotIds`和`Filters`。<br><li>snapshot-id - Array of String - 是否必填：否 -（过滤条件）按照快照的ID过滤。快照ID形如：`snap-11112222`。<br><li>snapshot-name - Array of String - 是否必填：否 -（过滤条件）按照快照名称过滤。<br><li>snapshot-state - Array of String - 是否必填：否 -（过滤条件）按照快照状态过滤。 (NORMAL：正常 | CREATING：创建中 | ROLLBACKING：回滚中。)<br><li>disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建快照的云盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)<br><li>project-id  - Array of String - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。<br><li>disk-id  - Array of String - 是否必填：否 -（过滤条件）按照创建快照的云硬盘ID过滤。<br><li>zone - Array of String - 是否必填：否 -（过滤条件）按照[可用区](/document/product/213/15753#ZoneInfo)过滤。<br><li>encrypt - Array of String - 是否必填：否 -（过滤条件）按是否加密盘快照过滤。 (TRUE：表示加密盘快照 | FALSE：表示非加密盘快照。)
	// <li>snapshot-type- Array of String - 是否必填：否 -（过滤条件）根据snapshot-type指定的快照类型查询对应的快照。
	// (SHARED_SNAPSHOT：表示共享过来的快照 | PRIVATE_SNAPSHOT：表示自己私有快照。)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](/document/product/362/15633)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 快照列表排序的依据字段。取值范围：<br><li>CREATE_TIME：依据快照的创建时间排序<br>默认按创建时间排序。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考API[简介](/document/product/362/15633)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 输出云盘列表的排列顺序。取值范围：<br><li>ASC：升序排列<br><li>DESC：降序排列。
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsResponseParams struct {
	// 快照的数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 快照的详情列表。
	SnapshotSet []*Snapshot `json:"SnapshotSet,omitempty" name:"SnapshotSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetsRequestParams struct {
	// 子网实例ID查询。形如：subnet-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// subnet-id - String - Subnet实例名称。
	// subnet-name - String - 子网名称。只支持单值的模糊查询。
	// cidr-block - String - 子网网段，形如: 192.168.1.0 。只支持单值的模糊查询。
	// vpc-id - String - VPC实例ID，形如：vpc-f49l6u0z。
	// vpc-cidr-block  - String - vpc网段，形如: 192.168.1.0 。只支持单值的模糊查询。
	// region - String - ECM地域
	// zone - String - 可用区。
	// tag-key - String -是否必填：否- 按照标签键进行过滤。
	// tag:tag-key - String - 是否必填：否 - 按照标签键值对进行过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 排序方式：time时间倒序, default按照网络规划排序
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// 子网实例ID查询。形如：subnet-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// subnet-id - String - Subnet实例名称。
	// subnet-name - String - 子网名称。只支持单值的模糊查询。
	// cidr-block - String - 子网网段，形如: 192.168.1.0 。只支持单值的模糊查询。
	// vpc-id - String - VPC实例ID，形如：vpc-f49l6u0z。
	// vpc-cidr-block  - String - vpc网段，形如: 192.168.1.0 。只支持单值的模糊查询。
	// region - String - ECM地域
	// zone - String - 可用区。
	// tag-key - String -是否必填：否- 按照标签键进行过滤。
	// tag:tag-key - String - 是否必填：否 - 按照标签键值对进行过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 排序方式：time时间倒序, default按照网络规划排序
	Sort *string `json:"Sort,omitempty" name:"Sort"`
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
	delete(f, "EcmRegion")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 子网对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet"`

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
type DescribeTargetHealthRequestParams struct {
	// 要查询的负载均衡实例 ID列表
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type DescribeTargetHealthRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的负载均衡实例 ID列表
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

func (r *DescribeTargetHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetHealthResponseParams struct {
	// 负载均衡实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancers []*LoadBalancerHealth `json:"LoadBalancers,omitempty" name:"LoadBalancers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTargetHealthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetHealthResponseParams `json:"Response"`
}

func (r *DescribeTargetHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetsRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器 ID列表
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 监听器协议类型
	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器端口
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DescribeTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器 ID列表
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 监听器协议类型
	Protocol *int64 `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器端口
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

func (r *DescribeTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetsResponseParams struct {
	// 监听器后端绑定的机器信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listeners []*ListenerBackend `json:"Listeners,omitempty" name:"Listeners"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetsResponseParams `json:"Response"`
}

func (r *DescribeTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 异步任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 异步任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultResponseParams struct {
	// 异步任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 执行结果，包括"SUCCESS", "FAILED", "RUNNING"
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResultResponseParams `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// 任务描述
	TaskSet []*TaskInput `json:"TaskSet,omitempty" name:"TaskSet"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务描述
	TaskSet []*TaskInput `json:"TaskSet,omitempty" name:"TaskSet"`
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
	delete(f, "TaskSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// 任务描述
	TaskSet []*TaskOutput `json:"TaskSet,omitempty" name:"TaskSet"`

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
type DescribeVpcsRequestParams struct {
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// vpc-name - String - VPC实例名称，只支持单值的模糊查询。
	// vpc-id - String - VPC实例ID形如：vpc-f49l6u0z。
	// cidr-block - String - vpc的cidr，只支持单值的模糊查询。
	// region - String - vpc的region。
	// tag-key - String -是否必填：否- 按照标签键进行过滤。
	// tag:tag-key - String - 是否必填：否 - 按照标签键值对进行过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 排序方式：time时间倒序, default按照网络规划排序
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// vpc-name - String - VPC实例名称，只支持单值的模糊查询。
	// vpc-id - String - VPC实例ID形如：vpc-f49l6u0z。
	// cidr-block - String - vpc的cidr，只支持单值的模糊查询。
	// region - String - vpc的region。
	// tag-key - String -是否必填：否- 按照标签键进行过滤。
	// tag:tag-key - String - 是否必填：否 - 按照标签键值对进行过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 排序方式：time时间倒序, default按照网络规划排序
	Sort *string `json:"Sort,omitempty" name:"Sort"`
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
	delete(f, "EcmRegion")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcsResponseParams struct {
	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 私有网络对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
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
type DetachDisksRequestParams struct {
	// 将要卸载的云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询，单次请求最多可卸载10块弹性云盘。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 对于非共享型云盘，会忽略该参数；对于共享型云盘，该参数表示要从哪个CVM实例上卸载云盘。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DetachDisksRequest struct {
	*tchttp.BaseRequest
	
	// 将要卸载的云硬盘ID， 通过[DescribeDisks](/document/product/362/16315)接口查询，单次请求最多可卸载10块弹性云盘。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`

	// 对于非共享型云盘，会忽略该参数；对于共享型云盘，该参数表示要从哪个CVM实例上卸载云盘。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DetachDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachDisksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachDisksResponse struct {
	*tchttp.BaseResponse
	Response *DetachDisksResponseParams `json:"Response"`
}

func (r *DetachDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachNetworkInterfaceRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 实例ID。形如：ein-hcs7jkg4
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type DetachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 实例ID。形如：ein-hcs7jkg4
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DetachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "InstanceId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachNetworkInterfaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *DetachNetworkInterfaceResponseParams `json:"Response"`
}

func (r *DetachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRoutesRequestParams struct {
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略ID。
	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

type DisableRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略ID。
	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *DisableRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DisableRoutesResponseParams `json:"Response"`
}

func (r *DisableRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAddressRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：eip-11112222。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 表示解绑 EIP 之后是否分配普通公网 IP。取值范围：
	// TRUE：表示解绑 EIP 之后分配普通公网 IP。
	// FALSE：表示解绑 EIP 之后不分配普通公网 IP。
	// 默认取值：FALSE。
	// 
	// 只有满足以下条件时才能指定该参数：
	// 只有在解绑主网卡的主内网 IP 上的 EIP 时才能指定该参数。
	// 解绑 EIP 后重新分配普通公网 IP 操作一个账号每天最多操作 10 次；详情可通过 DescribeAddressQuota 接口获取。
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitempty" name:"ReallocateNormalPublicIp"`
}

type DisassociateAddressRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：eip-11112222。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 表示解绑 EIP 之后是否分配普通公网 IP。取值范围：
	// TRUE：表示解绑 EIP 之后分配普通公网 IP。
	// FALSE：表示解绑 EIP 之后不分配普通公网 IP。
	// 默认取值：FALSE。
	// 
	// 只有满足以下条件时才能指定该参数：
	// 只有在解绑主网卡的主内网 IP 上的 EIP 时才能指定该参数。
	// 解绑 EIP 后重新分配普通公网 IP 操作一个账号每天最多操作 10 次；详情可通过 DescribeAddressQuota 接口获取。
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitempty" name:"ReallocateNormalPublicIp"`
}

func (r *DisassociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressId")
	delete(f, "ReallocateNormalPublicIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAddressResponseParams struct {
	// 异步任务TaskId。可以使用DescribeTaskResult接口查询任务状态。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateAddressResponseParams `json:"Response"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateInstancesKeyPairsRequestParams struct {
	// 可以通过以下方式获取可用的实例ID：
	// 通过登录控制台查询实例ID。
	// 通过调用接口 DescribeInstances ，取返回信息中的 InstanceId 获取实例ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 密钥对ID列表，每次请求批量密钥对的上限为100。密钥对ID形如：skey-11112222。
	// 
	// 可以通过以下方式获取可用的密钥ID：
	// 通过登录控制台查询密钥ID。
	// 通过调用接口 DescribeKeyPairs ，取返回信息中的 KeyId 获取密钥对ID。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再解绑密钥。取值范围：
	// TRUE：表示在正常关机失败后进行强制关机。
	// FALSE：表示在正常关机失败后不进行强制关机。
	// 
	// 默认取值：FALSE。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	
	// 可以通过以下方式获取可用的实例ID：
	// 通过登录控制台查询实例ID。
	// 通过调用接口 DescribeInstances ，取返回信息中的 InstanceId 获取实例ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 密钥对ID列表，每次请求批量密钥对的上限为100。密钥对ID形如：skey-11112222。
	// 
	// 可以通过以下方式获取可用的密钥ID：
	// 通过登录控制台查询密钥ID。
	// 通过调用接口 DescribeKeyPairs ，取返回信息中的 KeyId 获取密钥对ID。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再解绑密钥。取值范围：
	// TRUE：表示在正常关机失败后进行强制关机。
	// FALSE：表示在正常关机失败后不进行强制关机。
	// 
	// 默认取值：FALSE。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`
}

func (r *DisassociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateInstancesKeyPairsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "KeyIds")
	delete(f, "ForceStop")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateInstancesKeyPairsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateInstancesKeyPairsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateInstancesKeyPairsResponseParams `json:"Response"`
}

func (r *DisassociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateInstancesKeyPairsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// 要解绑的安全组ID，类似esg-efil73jd，只支持解绑单个安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 被解绑的实例ID，类似ein-lesecurk，支持指定多个实例 。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要解绑的安全组ID，类似esg-efil73jd，只支持解绑单个安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 被解绑的实例ID，类似ein-lesecurk，支持指定多个实例 。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Disk struct {
	// 云盘是否与挂载的实例一起销毁。<br><li>true:销毁实例时会同时销毁云盘，只支持按小时后付费云盘。<br><li>false：销毁实例时不销毁云盘。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" name:"DeleteWithInstance"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 硬盘介质类型。取值范围：<br><li>CLOUD_BASIC：表示普通云硬盘<br><li>CLOUD_PREMIUM：表示高性能云硬盘<br><li>CLOUD_SSD：表示SSD云硬盘<br><li>CLOUD_HSSD：表示增强型SSD云硬盘<br><li>CLOUD_TSSD：表示极速型SSD云硬盘。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 云盘状态。取值范围：<br><li>UNATTACHED：未挂载<br><li>ATTACHING：挂载中<br><li>ATTACHED：已挂载<br><li>DETACHING：解挂中<br><li>EXPANDING：扩容中<br><li>ROLLBACKING：回滚中<br><li>TORECYCLE：待回收<br><li>DUMPING：拷贝硬盘中。
	DiskState *string `json:"DiskState,omitempty" name:"DiskState"`

	// 云盘拥有的快照总数。
	SnapshotCount *int64 `json:"SnapshotCount,omitempty" name:"SnapshotCount"`

	// 云盘已挂载到子机，且子机与云盘都是包年包月。<br><li>true：子机设置了自动续费标识，但云盘未设置<br><li>false：云盘自动续费标识正常。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlagError *bool `json:"AutoRenewFlagError,omitempty" name:"AutoRenewFlagError"`

	// 云盘是否处于快照回滚状态。取值范围：<br><li>false:表示不处于快照回滚状态<br><li>true:表示处于快照回滚状态。
	Rollbacking *bool `json:"Rollbacking,omitempty" name:"Rollbacking"`

	// 对于非共享型云盘，该参数为空数组。对于共享型云盘，则表示该云盘当前被挂载到的CVM实例InstanceId
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 云盘是否为加密盘。取值范围：<br><li>false:表示非加密盘<br><li>true:表示加密盘。
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`

	// 云硬盘名称。
	DiskName *string `json:"DiskName,omitempty" name:"DiskName"`

	// 云硬盘因欠费销毁或者到期销毁时， 是否使用快照备份数据的标识。true表示销毁时创建快照进行数据备份。false表示直接销毁，不进行数据备份。
	BackupDisk *bool `json:"BackupDisk,omitempty" name:"BackupDisk"`

	// 与云盘绑定的标签，云盘未绑定标签则取值为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 云硬盘挂载的云主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云盘关联的定期快照ID。只有在调用DescribeDisks接口时，入参ReturnBindAutoSnapshotPolicy取值为TRUE才会返回该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds,omitempty" name:"AutoSnapshotPolicyIds"`

	// 云硬盘额外性能值，单位MB/s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`

	// 云盘是否处于类型变更中。取值范围：<br><li>false:表示云盘不处于类型变更中<br><li>true:表示云盘已发起类型变更，正处于迁移中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Migrating *bool `json:"Migrating,omitempty" name:"Migrating"`

	// 云硬盘ID。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 云盘拥有的快照总容量，单位为MB。
	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`

	// 云硬盘所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 判断预付费的云盘是否支持主动退还。<br><li>true:支持主动退还<br><li>false:不支持主动退还。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsReturnable *bool `json:"IsReturnable,omitempty" name:"IsReturnable"`

	// 云盘是否挂载到云主机上。取值范围：<br><li>false:表示未挂载<br><li>true:表示已挂载。
	Attached *bool `json:"Attached,omitempty" name:"Attached"`

	// 云硬盘大小，单位GB。
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 云盘类型变更的迁移进度，取值0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigratePercent *uint64 `json:"MigratePercent,omitempty" name:"MigratePercent"`

	// 云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// 付费模式。取值范围：<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：后付费，即按量计费。
	DiskChargeType *string `json:"DiskChargeType,omitempty" name:"DiskChargeType"`

	// 是否为弹性云盘，false表示非弹性云盘，true表示弹性云盘。
	Portable *bool `json:"Portable,omitempty" name:"Portable"`

	// 云盘是否具备创建快照的能力。取值范围：<br><li>false表示不具备<br><li>true表示具备。
	SnapshotAbility *bool `json:"SnapshotAbility,omitempty" name:"SnapshotAbility"`

	// 在云盘已挂载到实例，且实例与云盘都是包年包月的条件下，此字段才有意义。<br><li>true:云盘到期时间早于实例。<br><li>false：云盘到期时间晚于实例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadlineError *bool `json:"DeadlineError,omitempty" name:"DeadlineError"`

	// 云盘快照回滚的进度。
	RollbackPercent *uint64 `json:"RollbackPercent,omitempty" name:"RollbackPercent"`

	// 当前时间距离盘到期的天数（仅对预付费盘有意义）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DifferDaysOfDeadline *int64 `json:"DifferDaysOfDeadline,omitempty" name:"DifferDaysOfDeadline"`

	// 预付费云盘在不支持主动退还的情况下，该参数表明不支持主动退还的具体原因。取值范围：<br><li>1：云硬盘已经退还<br><li>2：云硬盘已过期<br><li>3：云盘不支持退还<br><li>8：超过可退还数量的限制。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnFailCode *int64 `json:"ReturnFailCode,omitempty" name:"ReturnFailCode"`

	// 云盘是否为共享型云盘。
	Shareable *bool `json:"Shareable,omitempty" name:"Shareable"`

	// 云硬盘的创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 云硬盘的到期时间。
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// 云盘的挂载类型。取值范围：<br><li>PF: PF挂载<br><li>VF: VF挂载
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachMode *string `json:"AttachMode,omitempty" name:"AttachMode"`
}

type DiskChargePrepaid struct {
	// 购买云盘的时长，默认单位为月，取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 需要将云盘的到期时间与挂载的子机对齐时，可传入该参数。该参数表示子机当前的到期时间，此时Period如果传入，则表示子机需要续费的时长，云盘会自动按对齐到子机续费后的到期时间续费，示例取值：2018-03-30 20:15:03。
	CurInstanceDeadline *string `json:"CurInstanceDeadline,omitempty" name:"CurInstanceDeadline"`
}

type DiskInfo struct {
	// 磁盘类型：LOCAL_BASIC
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 磁盘ID
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 磁盘大小（GB）
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type EipQuota struct {
	// 配额名称，取值范围：
	// TOTAL_EIP_QUOTA：用户当前地域下EIP的配额数；
	// DAILY_EIP_APPLY：用户当前地域下今日申购次数；
	// DAILY_PUBLIC_IP_ASSIGN：用户当前地域下，重新分配公网 IP次数。
	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`

	// 当前数量
	QuotaCurrent *uint64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`

	// 配额数量
	QuotaLimit *uint64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

// Predefined struct for user
type EnableRoutesRequestParams struct {
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略ID。
	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

type EnableRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略ID。
	RouteIds []*uint64 `json:"RouteIds,omitempty" name:"RouteIds"`
}

func (r *EnableRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableRoutesResponse struct {
	*tchttp.BaseResponse
	Response *EnableRoutesResponseParams `json:"Response"`
}

func (r *EnableRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {
	// 是否开启云镜服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// 是否开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`

	// 是否开通IP直通。若不指定该参数，则Linux镜像默认开通，windows镜像暂不支持IP直通。
	EIPDirectService *RunEIPDirectServiceEnabled `json:"EIPDirectService,omitempty" name:"EIPDirectService"`
}

type Filter struct {
	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type HaVip struct {
	// HAVIP的ID，是HAVIP的唯一标识。
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`

	// HAVIP名称。
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`

	// 虚拟IP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// HAVIP所在私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// HAVIP所在子网ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// HAVIP关联弹性网卡ID。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 被绑定的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 绑定EIP。
	AddressIp *string `json:"AddressIp,omitempty" name:"AddressIp"`

	// 状态：
	// AVAILABLE：运行中。
	// UNBIND：未绑定。
	State *string `json:"State,omitempty" name:"State"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 使用havip的业务标识。
	Business *string `json:"Business,omitempty" name:"Business"`
}

type HealthCheck struct {
	// 是否开启健康检查：1（开启）、0（关闭）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查的响应超时时间，可选值：2~60，默认值：2，单位：秒。响应超时时间要小于检查间隔时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 健康检查探测间隔时间，默认值：5，可选值：5~300，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2~10，单位：次。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发异常，可选值：2~10，单位：次。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnHealthyNum *int64 `json:"UnHealthyNum,omitempty" name:"UnHealthyNum"`

	// 自定义探测相关参数。健康检查端口，默认为后端服务的端口，除非您希望指定特定端口，否则建议留空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPort *int64 `json:"CheckPort,omitempty" name:"CheckPort"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查的输入格式，可取值：HEX或TEXT；取值为HEX时，SendContext和RecvContext的字符只能在0123456789ABCDEF中选取且长度必须是偶数位。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContextType *string `json:"ContextType,omitempty" name:"ContextType"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查发送的请求内容，只允许ASCII可见字符，最大长度限制500。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendContext *string `json:"SendContext,omitempty" name:"SendContext"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查返回的结果，只允许ASCII可见字符，最大长度限制500。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecvContext *string `json:"RecvContext,omitempty" name:"RecvContext"`

	// 自定义探测相关参数。健康检查使用的协议：TCP | CUSTOM（UDP监听器只支持CUSTOM；如果使用自定义健康检查功能，则必传）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckType *string `json:"CheckType,omitempty" name:"CheckType"`
}

type ISP struct {
	// 运营商ID
	ISPId *string `json:"ISPId,omitempty" name:"ISPId"`

	// 运营商名称
	ISPName *string `json:"ISPName,omitempty" name:"ISPName"`
}

type ISPCounter struct {
	// 运营商名称
	ProviderName *string `json:"ProviderName,omitempty" name:"ProviderName"`

	// 节点数量
	ProviderNodeNum *int64 `json:"ProviderNodeNum,omitempty" name:"ProviderNodeNum"`

	// 实例数量
	ProvederInstanceNum *int64 `json:"ProvederInstanceNum,omitempty" name:"ProvederInstanceNum"`

	// Zone实例信息结构体数组
	ZoneInstanceInfoSet []*ZoneInstanceInfo `json:"ZoneInstanceInfoSet,omitempty" name:"ZoneInstanceInfoSet"`
}

type Image struct {
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 镜像状态
	ImageState *string `json:"ImageState,omitempty" name:"ImageState"`

	// 镜像类型
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// 操作系统名称
	ImageOsName *string `json:"ImageOsName,omitempty" name:"ImageOsName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 镜像导入时间
	ImageCreateTime *string `json:"ImageCreateTime,omitempty" name:"ImageCreateTime"`

	// 操作系统位数
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 操作系统类型
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 操作系统版本
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// 操作系统平台
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 镜像所有者
	ImageOwner *int64 `json:"ImageOwner,omitempty" name:"ImageOwner"`

	// 镜像大小。单位：GB
	ImageSize *int64 `json:"ImageSize,omitempty" name:"ImageSize"`

	// 镜像来源信息
	SrcImage *SrcImage `json:"SrcImage,omitempty" name:"SrcImage"`

	// 镜像来源类型
	ImageSource *string `json:"ImageSource,omitempty" name:"ImageSource"`

	// 中间态和失败时候的任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 是否支持CloudInit
	IsSupportCloudInit *bool `json:"IsSupportCloudInit,omitempty" name:"IsSupportCloudInit"`
}

type ImageLimitConfig struct {
	// 支持的最大镜像大小，包括可导入的自定义镜像大小，中心云镜像大小，单位为GB。
	MaxImageSize *int64 `json:"MaxImageSize,omitempty" name:"MaxImageSize"`
}

type ImageOsList struct {
	// 支持的windows操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	Windows []*string `json:"Windows,omitempty" name:"Windows"`

	// 支持的linux操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	Linux []*string `json:"Linux,omitempty" name:"Linux"`
}

type ImageTask struct {
	// 镜像导入状态， PENDING, PROCESSING, SUCCESS, FAILED
	State *string `json:"State,omitempty" name:"State"`

	// 导入失败(FAILED)时， 说明失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ImageUrl struct {
	// 镜像文件COS链接，如设置私有读写，需授权腾讯云ECM运营账号访问权限。
	ImageFile *string `json:"ImageFile,omitempty" name:"ImageFile"`
}

// Predefined struct for user
type ImportCustomImageRequestParams struct {
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 导入镜像的操作系统架构，x86_64 或 i386
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 导入镜像的操作系统类型，通过DescribeImportImageOs获取
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 导入镜像的操作系统版本，通过DescribeImportImageOs获取
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 镜像启动方式，cloudinit或nbd， 默认cloudinit
	InitFlag *string `json:"InitFlag,omitempty" name:"InitFlag"`

	// 镜像文件描述，多层镜像按顺序传入
	ImageUrls []*ImageUrl `json:"ImageUrls,omitempty" name:"ImageUrls"`
}

type ImportCustomImageRequest struct {
	*tchttp.BaseRequest
	
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 导入镜像的操作系统架构，x86_64 或 i386
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 导入镜像的操作系统类型，通过DescribeImportImageOs获取
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 导入镜像的操作系统版本，通过DescribeImportImageOs获取
	OsVersion *string `json:"OsVersion,omitempty" name:"OsVersion"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 镜像启动方式，cloudinit或nbd， 默认cloudinit
	InitFlag *string `json:"InitFlag,omitempty" name:"InitFlag"`

	// 镜像文件描述，多层镜像按顺序传入
	ImageUrls []*ImageUrl `json:"ImageUrls,omitempty" name:"ImageUrls"`
}

func (r *ImportCustomImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportCustomImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageName")
	delete(f, "Architecture")
	delete(f, "OsType")
	delete(f, "OsVersion")
	delete(f, "ImageDescription")
	delete(f, "InitFlag")
	delete(f, "ImageUrls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportCustomImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportCustomImageResponseParams struct {
	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 异步任务ID，可根据DescribeCustomImageTask查询任务信息
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportCustomImageResponse struct {
	*tchttp.BaseResponse
	Response *ImportCustomImageResponseParams `json:"Response"`
}

func (r *ImportCustomImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportCustomImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportImageRequestParams struct {
	// 镜像的Id。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像的描述。
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 源地域
	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
}

type ImportImageRequest struct {
	*tchttp.BaseRequest
	
	// 镜像的Id。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像的描述。
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 源地域
	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
}

func (r *ImportImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "ImageDescription")
	delete(f, "SourceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportImageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportImageResponse struct {
	*tchttp.BaseResponse
	Response *ImportImageResponseParams `json:"Response"`
}

func (r *ImportImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称，如ens-34241f3s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例状态。取值范围：
	// PENDING：表示创建中
	// LAUNCH_FAILED：表示创建失败
	// RUNNING：表示运行中
	// STOPPED：表示关机
	// STARTING：表示开机中
	// STOPPING：表示关机中
	// REBOOTING：表示重启中
	// SHUTDOWN：表示停止待销毁
	// TERMINATING：表示销毁中。
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// 实例当前使用的镜像的信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *Image `json:"Image,omitempty" name:"Image"`

	// 实例当前所属的模块简要信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimpleModule *SimpleModule `json:"SimpleModule,omitempty" name:"SimpleModule"`

	// 实例所在的位置相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *Position `json:"Position,omitempty" name:"Position"`

	// 实例的网络相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Internet *Internet `json:"Internet,omitempty" name:"Internet"`

	// 实例的配置相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeConfig *InstanceTypeConfig `json:"InstanceTypeConfig,omitempty" name:"InstanceTypeConfig"`

	// 实例的创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// 实例最后一次操作。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperation *string `json:"LatestOperation,omitempty" name:"LatestOperation"`

	// 实例最后一次操作结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestOperationState *string `json:"LatestOperationState,omitempty" name:"LatestOperationState"`

	// 实例业务状态。取值范围：
	// NORMAL：表示正常状态的实例
	// EXPIRED：表示过期的实例
	// PROTECTIVELY_ISOLATED：表示被安全隔离的实例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`

	// 系统盘大小，单位GB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 数据盘大小，单位GB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// 实例UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UUID *string `json:"UUID,omitempty" name:"UUID"`

	// 付费方式。
	//     0为后付费。
	//     1为预付费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 过期时间。格式为yyyy-mm-dd HH:mm:ss。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 隔离时间。格式为yyyy-mm-dd HH:mm:ss。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// 是否自动续费。
	//       0为不自动续费。
	//       1为自动续费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 过期状态。
	//     NORMAL 表示机器运行正常。
	//     WILL_EXPIRE 表示即将过期。
	//     EXPIRED 表示已过期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireState *string `json:"ExpireState,omitempty" name:"ExpireState"`

	// 系统盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDisk *DiskInfo `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisks []*DiskInfo `json:"DataDisks,omitempty" name:"DataDisks"`

	// 新实例标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewFlag *int64 `json:"NewFlag,omitempty" name:"NewFlag"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// VPC属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 实例运营商字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ISP *string `json:"ISP,omitempty" name:"ISP"`

	// 物理位置信息。注意该字段目前为保留字段，均为空值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhysicalPosition *PhysicalPosition `json:"PhysicalPosition,omitempty" name:"PhysicalPosition"`
}

type InstanceFamilyConfig struct {
	// 机型名称
	InstanceFamilyName *string `json:"InstanceFamilyName,omitempty" name:"InstanceFamilyName"`

	// 机型ID
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type InstanceFamilyTypeConfig struct {
	// 实例机型系列类型Id
	InstanceFamilyType *string `json:"InstanceFamilyType,omitempty" name:"InstanceFamilyType"`

	// 实例机型系列类型名称
	InstanceFamilyTypeName *string `json:"InstanceFamilyTypeName,omitempty" name:"InstanceFamilyTypeName"`
}

type InstanceNetworkInfo struct {
	// 实例内外网ip相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressInfoSet []*AddressInfo `json:"AddressInfoSet,omitempty" name:"AddressInfoSet"`

	// 网卡ID。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 网卡名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// 主网卡属性。true为主网卡，false为辅助网卡。
	Primary *bool `json:"Primary,omitempty" name:"Primary"`
}

type InstanceNetworkLimitConfig struct {
	// cpu核数
	CpuNum *int64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// 网卡数量限制
	NetworkInterfaceLimit *int64 `json:"NetworkInterfaceLimit,omitempty" name:"NetworkInterfaceLimit"`

	// 每张网卡内网ip数量限制
	InnerIpPerNetworkInterface *int64 `json:"InnerIpPerNetworkInterface,omitempty" name:"InnerIpPerNetworkInterface"`

	// 每个实例的外网ip限制
	PublicIpPerInstance *int64 `json:"PublicIpPerInstance,omitempty" name:"PublicIpPerInstance"`
}

type InstanceOperator struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例禁止的操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeniedActions []*OperatorAction `json:"DeniedActions,omitempty" name:"DeniedActions"`
}

type InstancePricesPartDetail struct {
	// cpu的价格信息
	CpuPrice *PriceDetail `json:"CpuPrice,omitempty" name:"CpuPrice"`

	// 内存价格信息
	MemPrice *PriceDetail `json:"MemPrice,omitempty" name:"MemPrice"`

	// 磁盘价格信息
	DisksPrice *PriceDetail `json:"DisksPrice,omitempty" name:"DisksPrice"`
}

type InstanceStatistic struct {
	// 实例的类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例的个数
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

type InstanceTypeConfig struct {
	// 机型族配置信息
	InstanceFamilyConfig *InstanceFamilyConfig `json:"InstanceFamilyConfig,omitempty" name:"InstanceFamilyConfig"`

	// 机型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU核数
	Vcpu *int64 `json:"Vcpu,omitempty" name:"Vcpu"`

	// 内存大小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 主频
	Frequency *string `json:"Frequency,omitempty" name:"Frequency"`

	// 处理器型号
	CpuModelName *string `json:"CpuModelName,omitempty" name:"CpuModelName"`

	// 机型族类别配置信息
	InstanceFamilyTypeConfig *InstanceFamilyTypeConfig `json:"InstanceFamilyTypeConfig,omitempty" name:"InstanceFamilyTypeConfig"`

	// 机型额外信息 是一个json字符串，如果存在则表示特殊机型，格式如下：{"dataDiskSize":3200,"systemDiskSize":60, "systemDiskSizeShow":"系统盘默认60G","dataDiskSizeShow":"本地NVMe SSD 硬盘3200 GB"}
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`

	// GPU卡数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vgpu *float64 `json:"Vgpu,omitempty" name:"Vgpu"`

	// GPU型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuModelName *string `json:"GpuModelName,omitempty" name:"GpuModelName"`
}

type InstancesPrice struct {
	// 分部描述实例子维度的价格
	InstancePricesPartDetail *InstancePricesPartDetail `json:"InstancePricesPartDetail,omitempty" name:"InstancePricesPartDetail"`

	// 实例总价折扣
	Discount *uint64 `json:"Discount,omitempty" name:"Discount"`

	// 折扣后价格
	DiscountPrice *uint64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// 折扣前价格，原始总价
	OriginalPrice *uint64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type Internet struct {
	// 实例的内网相关信息列表。顺序为主网卡在前，辅助网卡按绑定先后顺序排列。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIPAddressSet []*PrivateIPAddressInfo `json:"PrivateIPAddressSet,omitempty" name:"PrivateIPAddressSet"`

	// 实例的公网相关信息列表。顺序为主网卡在前，辅助网卡按绑定先后顺序排列。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIPAddressSet []*PublicIPAddressInfo `json:"PublicIPAddressSet,omitempty" name:"PublicIPAddressSet"`

	// 实例网络相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNetworkInfoSet []*InstanceNetworkInfo `json:"InstanceNetworkInfoSet,omitempty" name:"InstanceNetworkInfoSet"`
}

type Ipv6Address struct {
	// IPv6地址，形如：3402:4e00:20:100:0:8cd9:2a67:71f3
	Address *string `json:"Address,omitempty" name:"Address"`

	// 是否是主IP。
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// EIP实例ID，形如：eip-hxlqja90。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 公网IP是否被封堵。
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty" name:"IsWanIpBlocked"`

	// IPv6地址状态：
	// PENDING：生产中
	// MIGRATING：迁移中
	// DELETING：删除中
	// AVAILABLE：可用的
	State *string `json:"State,omitempty" name:"State"`
}

type KeyPair struct {
	// 密钥对的ID，是密钥对的唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 密钥对名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`

	// 密钥对所属的项目ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 密钥对描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 密钥对的纯文本公钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 钥对的纯文本私钥。腾讯云不会保管私钥，请用户自行妥善保存。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 钥关联的实例ID列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds,omitempty" name:"AssociatedInstanceIds"`

	// 创建时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type Listener struct {
	// 负载均衡监听器 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 监听器的健康检查信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 请求的调度方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 会话保持时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 监听器的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 监听器的会话类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// 端口段结束端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`
}

type ListenerBackend struct {
	// 监听器 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器的协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器的端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 监听器上绑定的后端服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*Backend `json:"Targets,omitempty" name:"Targets"`
}

type ListenerHealth struct {
	// 监听器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器的协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器的端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 监听器的转发规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*RuleHealth `json:"Rules,omitempty" name:"Rules"`
}

type LoadBalancer struct {
	// 区域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 位置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *Position `json:"Position,omitempty" name:"Position"`

	// 负载均衡实例 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例的名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例的网络类型：OPEN：公网属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的 VIP 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 负载均衡实例的状态，包括
	//  0：创建中，1：正常运行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 负载均衡实例的创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 负载均衡实例的上次状态转换时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusTime *string `json:"StatusTime,omitempty" name:"StatusTime"`

	// 私有网络的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 负载均衡实例的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 负载均衡IP地址所属的ISP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// 负载均衡实例的网络属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAttributes *LoadBalancerInternetAccessible `json:"NetworkAttributes,omitempty" name:"NetworkAttributes"`

	// 安全组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecureGroups []*string `json:"SecureGroups,omitempty" name:"SecureGroups"`

	// 后端机器是否放通来自ELB的流量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`

	// 负载均衡实例的IPv6地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIPv6 *string `json:"AddressIPv6,omitempty" name:"AddressIPv6"`
}

type LoadBalancerHealth struct {
	// 负载均衡实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listeners []*ListenerHealth `json:"Listeners,omitempty" name:"Listeners"`
}

type LoadBalancerInternetAccessible struct {
	// 最大出带宽，单位Mbps。默认值10
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

// Predefined struct for user
type MigrateNetworkInterfaceRequestParams struct {
	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 弹性网卡当前绑定的ECM实例ID。形如：ein-r8hr2upy。
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`

	// 待迁移的目的ECM实例ID。
	DestinationInstanceId *string `json:"DestinationInstanceId,omitempty" name:"DestinationInstanceId"`
}

type MigrateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 弹性网卡当前绑定的ECM实例ID。形如：ein-r8hr2upy。
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" name:"SourceInstanceId"`

	// 待迁移的目的ECM实例ID。
	DestinationInstanceId *string `json:"DestinationInstanceId,omitempty" name:"DestinationInstanceId"`
}

func (r *MigrateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "SourceInstanceId")
	delete(f, "DestinationInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigrateNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigrateNetworkInterfaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MigrateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *MigrateNetworkInterfaceResponseParams `json:"Response"`
}

func (r *MigrateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigratePrivateIpAddressRequestParams struct {
	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 当前内网IP绑定的弹性网卡实例ID，例如：eni-11112222。
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitempty" name:"SourceNetworkInterfaceId"`

	// 待迁移的目的弹性网卡实例ID。
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitempty" name:"DestinationNetworkInterfaceId"`

	// 迁移的内网IP地址，例如：10.0.0.6。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

type MigratePrivateIpAddressRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 当前内网IP绑定的弹性网卡实例ID，例如：eni-11112222。
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitempty" name:"SourceNetworkInterfaceId"`

	// 待迁移的目的弹性网卡实例ID。
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitempty" name:"DestinationNetworkInterfaceId"`

	// 迁移的内网IP地址，例如：10.0.0.6。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *MigratePrivateIpAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigratePrivateIpAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "SourceNetworkInterfaceId")
	delete(f, "DestinationNetworkInterfaceId")
	delete(f, "PrivateIpAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigratePrivateIpAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigratePrivateIpAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MigratePrivateIpAddressResponse struct {
	*tchttp.BaseResponse
	Response *MigratePrivateIpAddressResponseParams `json:"Response"`
}

func (r *MigratePrivateIpAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigratePrivateIpAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressAttributeRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：eip-11112222。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 修改后的 EIP 名称。长度上限为20个字符。
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// 设定EIP是否直通，"TRUE"表示直通，"FALSE"表示非直通。注意该参数仅对EIP直通功能可见的用户可以设定。
	EipDirectConnection *string `json:"EipDirectConnection,omitempty" name:"EipDirectConnection"`
}

type ModifyAddressAttributeRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：eip-11112222。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 修改后的 EIP 名称。长度上限为20个字符。
	AddressName *string `json:"AddressName,omitempty" name:"AddressName"`

	// 设定EIP是否直通，"TRUE"表示直通，"FALSE"表示非直通。注意该参数仅对EIP直通功能可见的用户可以设定。
	EipDirectConnection *string `json:"EipDirectConnection,omitempty" name:"EipDirectConnection"`
}

func (r *ModifyAddressAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressId")
	delete(f, "AddressName")
	delete(f, "EipDirectConnection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAddressAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressAttributeResponseParams `json:"Response"`
}

func (r *ModifyAddressAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressesBandwidthRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// EIP唯一标识ID，形如'eip-xxxxxxx'
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`

	// 调整带宽目标值
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type ModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// EIP唯一标识ID，形如'eip-xxxxxxx'
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`

	// 调整带宽目标值
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *ModifyAddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressesBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressIds")
	delete(f, "InternetMaxBandwidthOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressesBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressesBandwidthResponseParams struct {
	// 异步任务TaskId。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressesBandwidthResponseParams `json:"Response"`
}

func (r *ModifyAddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultSubnetRequestParams struct {
	// ECM地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ECM可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type ModifyDefaultSubnetRequest struct {
	*tchttp.BaseRequest
	
	// ECM地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// ECM可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *ModifyDefaultSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDefaultSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDefaultSubnetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDefaultSubnetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDefaultSubnetResponseParams `json:"Response"`
}

func (r *ModifyDefaultSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDefaultSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHaVipAttributeRequestParams struct {
	// HAVIP唯一ID，形如：havip-9o233uri。
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`

	// HAVIP名称，可任意命名，但不得超过60个字符。
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`
}

type ModifyHaVipAttributeRequest struct {
	*tchttp.BaseRequest
	
	// HAVIP唯一ID，形如：havip-9o233uri。
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`

	// HAVIP名称，可任意命名，但不得超过60个字符。
	HaVipName *string `json:"HaVipName,omitempty" name:"HaVipName"`
}

func (r *ModifyHaVipAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHaVipAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	delete(f, "HaVipName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHaVipAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHaVipAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyHaVipAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHaVipAttributeResponseParams `json:"Response"`
}

func (r *ModifyHaVipAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHaVipAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAttributeRequestParams struct {
	// 镜像ID，形如img-gvbnzy6f
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 设置新的镜像名称；必须满足下列限制：
	// 不得超过20个字符。
	// - 镜像名称不能与已有镜像重复。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 设置新的镜像描述；必须满足下列限制：
	// - 不得超过60个字符。
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

type ModifyImageAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 镜像ID，形如img-gvbnzy6f
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 设置新的镜像名称；必须满足下列限制：
	// 不得超过20个字符。
	// - 镜像名称不能与已有镜像重复。
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 设置新的镜像描述；必须满足下列限制：
	// - 不得超过60个字符。
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`
}

func (r *ModifyImageAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "ImageName")
	delete(f, "ImageDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImageAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImageAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyImageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImageAttributeResponseParams `json:"Response"`
}

func (r *ModifyImageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesAttributeRequestParams struct {
	// 待修改的实例ID列表。在单次请求的过程中，请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 修改成功后显示的实例名称，不得超过60个字符，不传则名称显示为空。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 指定实例的安全组Id列表，子机将重新关联指定列表的安全组，原本关联的安全组会被解绑。限制不超过5个。
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的实例ID列表。在单次请求的过程中，请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 修改成功后显示的实例名称，不得超过60个字符，不传则名称显示为空。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 指定实例的安全组Id列表，子机将重新关联指定列表的安全组，原本关联的安全组会被解绑。限制不超过5个。
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "InstanceName")
	delete(f, "SecurityGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstancesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstancesAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstancesAttributeResponseParams `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpv6AddressesAttributeRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，形如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的IPv6地址信息。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

type ModifyIpv6AddressesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，形如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的IPv6地址信息。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

func (r *ModifyIpv6AddressesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpv6AddressesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIpv6AddressesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpv6AddressesAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyIpv6AddressesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIpv6AddressesAttributeResponseParams `json:"Response"`
}

func (r *ModifyIpv6AddressesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpv6AddressesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 新的监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 健康检查相关参数
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
}

type ModifyListenerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 新的监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 健康检查相关参数
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`
}

func (r *ModifyListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "SessionExpireTime")
	delete(f, "HealthCheck")
	delete(f, "Scheduler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyListenerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyListenerResponseParams `json:"Response"`
}

func (r *ModifyListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesRequestParams struct {
	// 负载均衡的唯一ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 网络计费及带宽相关参数
	InternetChargeInfo *LoadBalancerInternetAccessible `json:"InternetChargeInfo,omitempty" name:"InternetChargeInfo"`

	// Target是否放通来自ELB的流量。开启放通（true）：只验证ELB上的安全组；不开启放通（false）：需同时验证ELB和后端实例上的安全组。
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`
}

type ModifyLoadBalancerAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡的唯一ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 网络计费及带宽相关参数
	InternetChargeInfo *LoadBalancerInternetAccessible `json:"InternetChargeInfo,omitempty" name:"InternetChargeInfo"`

	// Target是否放通来自ELB的流量。开启放通（true）：只验证ELB上的安全组；不开启放通（false）：需同时验证ELB和后端实例上的安全组。
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`
}

func (r *ModifyLoadBalancerAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerName")
	delete(f, "InternetChargeInfo")
	delete(f, "LoadBalancerPassToTarget")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerAttributesResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleConfigRequestParams struct {
	// 模块ID。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 机型ID。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 默认数据盘大小，单位：G。范围不得超过数据盘范围大小，详看DescribeConfig。
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// 默认系统盘大小，单位：G。范围不得超过数据盘范围大小，详看DescribeConfig。
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// 系统盘
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type ModifyModuleConfigRequest struct {
	*tchttp.BaseRequest
	
	// 模块ID。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 机型ID。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 默认数据盘大小，单位：G。范围不得超过数据盘范围大小，详看DescribeConfig。
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// 默认系统盘大小，单位：G。范围不得超过数据盘范围大小，详看DescribeConfig。
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// 系统盘
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

func (r *ModifyModuleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "InstanceType")
	delete(f, "DefaultDataDiskSize")
	delete(f, "DefaultSystemDiskSize")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleConfigResponseParams `json:"Response"`
}

func (r *ModifyModuleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleDisableWanIpRequestParams struct {
	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 是否禁止分配外网ip,true：统一分配外网ip，false：禁止分配外网ip.
	DisableWanIp *bool `json:"DisableWanIp,omitempty" name:"DisableWanIp"`
}

type ModifyModuleDisableWanIpRequest struct {
	*tchttp.BaseRequest
	
	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 是否禁止分配外网ip,true：统一分配外网ip，false：禁止分配外网ip.
	DisableWanIp *bool `json:"DisableWanIp,omitempty" name:"DisableWanIp"`
}

func (r *ModifyModuleDisableWanIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleDisableWanIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "DisableWanIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleDisableWanIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleDisableWanIpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleDisableWanIpResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleDisableWanIpResponseParams `json:"Response"`
}

func (r *ModifyModuleDisableWanIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleDisableWanIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleImageRequestParams struct {
	// 默认镜像ID
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

type ModifyModuleImageRequest struct {
	*tchttp.BaseRequest
	
	// 默认镜像ID
	DefaultImageId *string `json:"DefaultImageId,omitempty" name:"DefaultImageId"`

	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *ModifyModuleImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DefaultImageId")
	delete(f, "ModuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleImageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleImageResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleImageResponseParams `json:"Response"`
}

func (r *ModifyModuleImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleIpDirectRequestParams struct {
	// 模块ID。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 是否关闭IP直通。取值范围：
	// true：表示关闭IP直通
	// false：表示开通IP直通
	CloseIpDirect *bool `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`
}

type ModifyModuleIpDirectRequest struct {
	*tchttp.BaseRequest
	
	// 模块ID。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 是否关闭IP直通。取值范围：
	// true：表示关闭IP直通
	// false：表示开通IP直通
	CloseIpDirect *bool `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`
}

func (r *ModifyModuleIpDirectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleIpDirectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "CloseIpDirect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleIpDirectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleIpDirectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleIpDirectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleIpDirectResponseParams `json:"Response"`
}

func (r *ModifyModuleIpDirectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleIpDirectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleNameRequestParams struct {
	// 模块ID。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称。
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type ModifyModuleNameRequest struct {
	*tchttp.BaseRequest
	
	// 模块ID。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称。
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *ModifyModuleNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "ModuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleNameResponseParams `json:"Response"`
}

func (r *ModifyModuleNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleNetworkRequestParams struct {
	// 模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 默认出带宽上限
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`

	// 默认入带宽上限
	DefaultBandwidthIn *int64 `json:"DefaultBandwidthIn,omitempty" name:"DefaultBandwidthIn"`
}

type ModifyModuleNetworkRequest struct {
	*tchttp.BaseRequest
	
	// 模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 默认出带宽上限
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`

	// 默认入带宽上限
	DefaultBandwidthIn *int64 `json:"DefaultBandwidthIn,omitempty" name:"DefaultBandwidthIn"`
}

func (r *ModifyModuleNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleId")
	delete(f, "DefaultBandwidth")
	delete(f, "DefaultBandwidthIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleNetworkResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleNetworkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleNetworkResponseParams `json:"Response"`
}

func (r *ModifyModuleNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleSecurityGroupsRequestParams struct {
	// 安全组列表。不超过5个。
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitempty" name:"SecurityGroupIdSet"`

	// 模块id。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

type ModifyModuleSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组列表。不超过5个。
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitempty" name:"SecurityGroupIdSet"`

	// 模块id。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *ModifyModuleSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIdSet")
	delete(f, "ModuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyModuleSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyModuleSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateIpAddressesAttributeRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// ECM 节点Region信息，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type ModifyPrivateIpAddressesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// ECM 节点Region信息，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *ModifyPrivateIpAddressesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateIpAddressesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddresses")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateIpAddressesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateIpAddressesAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrivateIpAddressesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateIpAddressesAttributeResponseParams `json:"Response"`
}

func (r *ModifyPrivateIpAddressesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateIpAddressesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRouteTableAttributeRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

type ModifyRouteTableAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *ModifyRouteTableAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRouteTableAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRouteTableAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRouteTableAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRouteTableAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRouteTableAttributeResponseParams `json:"Response"`
}

func (r *ModifyRouteTableAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRouteTableAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupAttributeRequestParams struct {
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
}

type ModifySecurityGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitempty" name:"GroupDescription"`
}

func (r *ModifySecurityGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupPoliciesRequestParams struct {
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组规则集合。 SecurityGroupPolicySet对象必须同时指定新的出（Egress）入（Ingress）站规则。 SecurityGroupPolicy对象不支持自定义索引（PolicyIndex）。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`

	// 排序安全组标识。值为True时，支持安全组排序；SortPolicys不存在或SortPolicys为False时，为修改安全组规则。
	SortPolicys *bool `json:"SortPolicys,omitempty" name:"SortPolicys"`
}

type ModifySecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组规则集合。 SecurityGroupPolicySet对象必须同时指定新的出（Egress）入（Ingress）站规则。 SecurityGroupPolicy对象不支持自定义索引（PolicyIndex）。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`

	// 排序安全组标识。值为True时，支持安全组排序；SortPolicys不存在或SortPolicys为False时，为修改安全组规则。
	SortPolicys *bool `json:"SortPolicys,omitempty" name:"SortPolicys"`
}

func (r *ModifySecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	delete(f, "SortPolicys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetAttributeRequestParams struct {
	// 子网实例ID。形如：subnet-pxir56ns。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网是否开启广播。
	EnableBroadcast *string `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`

	// 子网的标签键值
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type ModifySubnetAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 子网实例ID。形如：subnet-pxir56ns。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网是否开启广播。
	EnableBroadcast *string `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`

	// 子网的标签键值
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "SubnetId")
	delete(f, "EcmRegion")
	delete(f, "SubnetName")
	delete(f, "EnableBroadcast")
	delete(f, "Tags")
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
type ModifyTargetPortRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改端口的后端服务列表
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 后端服务绑定到监听器或转发规则的新端口
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`
}

type ModifyTargetPortRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改端口的后端服务列表
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 后端服务绑定到监听器或转发规则的新端口
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`
}

func (r *ModifyTargetPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "NewPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetPortResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTargetPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetPortResponseParams `json:"Response"`
}

func (r *ModifyTargetPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetWeightRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改权重的后端服务列表
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 后端服务新的转发权重，取值范围：0~100，默认值10。如果设置了 Targets.Weight 参数，则此参数不生效。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type ModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改权重的后端服务列表
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 后端服务新的转发权重，取值范围：0~100，默认值10。如果设置了 Targets.Weight 参数，则此参数不生效。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *ModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "Weight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetWeightResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetWeightResponseParams `json:"Response"`
}

func (r *ModifyTargetWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAttributeRequestParams struct {
	// VPC实例ID。形如：vpc-f49l6u0z。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 私有网络名称，可任意命名，但不得超过60个字符。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 私有网络描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyVpcAttributeRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。形如：vpc-f49l6u0z。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 私有网络名称，可任意命名，但不得超过60个字符。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 私有网络描述
	Description *string `json:"Description,omitempty" name:"Description"`
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
	delete(f, "EcmRegion")
	delete(f, "VpcName")
	delete(f, "Tags")
	delete(f, "Description")
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

type Module struct {
	// 模块Id。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称。
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// 模块状态：
	// NORMAL：正常。
	// DELETING：删除中 
	// DELETEFAILED：删除失败。
	ModuleState *string `json:"ModuleState,omitempty" name:"ModuleState"`

	// 默认系统盘大小。
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// 默认数据盘大小。
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// 默认机型。
	InstanceTypeConfig *InstanceTypeConfig `json:"InstanceTypeConfig,omitempty" name:"InstanceTypeConfig"`

	// 默认镜像。
	DefaultImage *Image `json:"DefaultImage,omitempty" name:"DefaultImage"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 默认出带宽。
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`

	// 标签集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// 是否关闭IP直通。
	CloseIpDirect *int64 `json:"CloseIpDirect,omitempty" name:"CloseIpDirect"`

	// 默认安全组id列表。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 默认入带宽。
	DefaultBandwidthIn *int64 `json:"DefaultBandwidthIn,omitempty" name:"DefaultBandwidthIn"`

	// 自定义脚本数据
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘信息。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type ModuleCounter struct {
	// 运营商统计信息列表
	ISPCounterSet []*ISPCounter `json:"ISPCounterSet,omitempty" name:"ISPCounterSet"`

	// 省份数量
	ProvinceNum *int64 `json:"ProvinceNum,omitempty" name:"ProvinceNum"`

	// 城市数量
	CityNum *int64 `json:"CityNum,omitempty" name:"CityNum"`

	// 节点数量
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

type ModuleItem struct {
	// 节点实例统计信息
	NodeInstanceNum *NodeInstanceNum `json:"NodeInstanceNum,omitempty" name:"NodeInstanceNum"`

	// 模块信息
	Module *Module `json:"Module,omitempty" name:"Module"`
}

type MonthNetwork struct {
	// 节点zone信息
	ZoneInfo *ZoneInfo `json:"ZoneInfo,omitempty" name:"ZoneInfo"`

	// 带宽月份 示例"202103"
	Month *string `json:"Month,omitempty" name:"Month"`

	// 带宽包id 格式如"bwp-xxxxxxxx"
	BandwidthPkgId *string `json:"BandwidthPkgId,omitempty" name:"BandwidthPkgId"`

	// 运营商简称 取值范围"CUCC;CTCC;CMCC"
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 入网带宽包峰值,取值范围0.0-xxx.xxx
	TrafficMaxIn *float64 `json:"TrafficMaxIn,omitempty" name:"TrafficMaxIn"`

	// 出网带宽包峰值,取值范围0.0-xxx.xxx
	TrafficMaxOut *float64 `json:"TrafficMaxOut,omitempty" name:"TrafficMaxOut"`

	// 计费带宽,取值范围0.0-xxx.xxx
	FeeTraffic *float64 `json:"FeeTraffic,omitempty" name:"FeeTraffic"`

	// 月计费带宽起始时间 格式"yyyy-mm-dd HH:mm:ss"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 月计费带宽结束时间 格式"yyyy-mm-dd HH:mm:ss"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 月计费带宽实际有效天数 整形必须大于等于0
	EffectiveDays *int64 `json:"EffectiveDays,omitempty" name:"EffectiveDays"`

	// 指定月份的实际天数 实例 30
	MonthDays *int64 `json:"MonthDays,omitempty" name:"MonthDays"`

	// 有效天占比 保留小数点后四位0.2134
	EffectiveDaysRate *float64 `json:"EffectiveDaysRate,omitempty" name:"EffectiveDaysRate"`

	// 计费带宽包类型 实例"Address","LoadBalance","AddressIpv6"
	BandwidthPkgType *string `json:"BandwidthPkgType,omitempty" name:"BandwidthPkgType"`
}

type NetworkInterface struct {
	// 弹性网卡实例ID，例如：eni-f1xjkw1b。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 弹性网卡名称。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// 弹性网卡描述。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// 子网实例ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 绑定的安全组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupSet []*string `json:"GroupSet,omitempty" name:"GroupSet"`

	// 是否是主网卡。
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// MAC地址。
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 弹性网卡状态：
	// PENDING：创建中
	// AVAILABLE：可用的
	// ATTACHING：绑定中
	// DETACHING：解绑中
	// DELETING：删除中
	State *string `json:"State,omitempty" name:"State"`

	// 内网IP信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet"`

	// 绑定的云服务器对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attachment *NetworkInterfaceAttachment `json:"Attachment,omitempty" name:"Attachment"`

	// 可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// IPv6地址列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet"`

	// 标签键值对。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// 网卡类型。0 - 弹性网卡；1 - evm弹性网卡。
	EniType *uint64 `json:"EniType,omitempty" name:"EniType"`

	// EcmRegion ecm区域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 网卡绑定的子机类型：cvm，eks。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Business *string `json:"Business,omitempty" name:"Business"`
}

type NetworkInterfaceAttachment struct {
	// 云主机实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 网卡在云主机实例内的序号。
	DeviceIndex *uint64 `json:"DeviceIndex,omitempty" name:"DeviceIndex"`

	// 云主机所有者账户信息。
	InstanceAccountId *string `json:"InstanceAccountId,omitempty" name:"InstanceAccountId"`

	// 绑定时间。
	AttachTime *string `json:"AttachTime,omitempty" name:"AttachTime"`
}

type NetworkStorageRange struct {
	// 网络带宽上限
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 数据盘上限
	MaxSystemDiskSize *int64 `json:"MaxSystemDiskSize,omitempty" name:"MaxSystemDiskSize"`

	// 网络带宽下限
	MinBandwidth *int64 `json:"MinBandwidth,omitempty" name:"MinBandwidth"`

	// 数据盘下限
	MinSystemDiskSize *int64 `json:"MinSystemDiskSize,omitempty" name:"MinSystemDiskSize"`

	// 最大数据盘大小
	MaxDataDiskSize *int64 `json:"MaxDataDiskSize,omitempty" name:"MaxDataDiskSize"`

	// 最小数据盘大小
	MinDataDiskSize *int64 `json:"MinDataDiskSize,omitempty" name:"MinDataDiskSize"`

	// 建议带宽
	SuggestBandwidth *int64 `json:"SuggestBandwidth,omitempty" name:"SuggestBandwidth"`

	// 建议硬盘大小
	SuggestDataDiskSize *int64 `json:"SuggestDataDiskSize,omitempty" name:"SuggestDataDiskSize"`

	// 建议系统盘大小
	SuggestSystemDiskSize *int64 `json:"SuggestSystemDiskSize,omitempty" name:"SuggestSystemDiskSize"`

	// Cpu核数峰值
	MaxVcpu *int64 `json:"MaxVcpu,omitempty" name:"MaxVcpu"`

	// Cpu核最小值
	MinVcpu *int64 `json:"MinVcpu,omitempty" name:"MinVcpu"`

	// 单次请求最大cpu核数
	MaxVcpuPerReq *int64 `json:"MaxVcpuPerReq,omitempty" name:"MaxVcpuPerReq"`

	// 带宽步长
	PerBandwidth *int64 `json:"PerBandwidth,omitempty" name:"PerBandwidth"`

	// 数据盘步长
	PerDataDisk *int64 `json:"PerDataDisk,omitempty" name:"PerDataDisk"`

	// 总模块数量
	MaxModuleNum *int64 `json:"MaxModuleNum,omitempty" name:"MaxModuleNum"`
}

type Node struct {
	// zone信息。
	ZoneInfo *ZoneInfo `json:"ZoneInfo,omitempty" name:"ZoneInfo"`

	// 国家信息。
	Country *Country `json:"Country,omitempty" name:"Country"`

	// 区域信息。
	Area *Area `json:"Area,omitempty" name:"Area"`

	// 省份信息。
	Province *Province `json:"Province,omitempty" name:"Province"`

	// 城市信息。
	City *City `json:"City,omitempty" name:"City"`

	// Region信息。
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`

	// 运营商列表。
	ISPSet []*ISP `json:"ISPSet,omitempty" name:"ISPSet"`

	// 运营商数量。
	ISPNum *int64 `json:"ISPNum,omitempty" name:"ISPNum"`
}

type NodeInstanceNum struct {
	// 节点数量
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

type OperatorAction struct {
	// 可执行操作
	Action *string `json:"Action,omitempty" name:"Action"`

	// 编码Code
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 具体信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type OsVersion struct {
	// 操作系统类型
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 支持的操作系统版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsVersions []*string `json:"OsVersions,omitempty" name:"OsVersions"`

	// 支持的操作系统架构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Architecture []*string `json:"Architecture,omitempty" name:"Architecture"`
}

type PackingQuotaGroup struct {
	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区id
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// ISP id
	ISPId *string `json:"ISPId,omitempty" name:"ISPId"`

	// 一组相互关联的装箱配额
	PackingQuotaInfos []*PackingQuotaInfo `json:"PackingQuotaInfos,omitempty" name:"PackingQuotaInfos"`
}

type PackingQuotaInfo struct {
	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 装箱配额
	PackingQuota *int64 `json:"PackingQuota,omitempty" name:"PackingQuota"`
}

type PeakBase struct {
	// CPU峰值
	PeakCpuNum *int64 `json:"PeakCpuNum,omitempty" name:"PeakCpuNum"`

	// 内存峰值
	PeakMemoryNum *int64 `json:"PeakMemoryNum,omitempty" name:"PeakMemoryNum"`

	// 硬盘峰值
	PeakStorageNum *int64 `json:"PeakStorageNum,omitempty" name:"PeakStorageNum"`

	// 记录时间
	RecordTime *string `json:"RecordTime,omitempty" name:"RecordTime"`
}

type PeakFamilyInfo struct {
	// 机型类别信息。
	InstanceFamily *InstanceFamilyTypeConfig `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// 基础数据峰值信息。
	PeakBaseSet []*PeakBase `json:"PeakBaseSet,omitempty" name:"PeakBaseSet"`
}

type PeakNetwork struct {
	// 记录时间。
	RecordTime *string `json:"RecordTime,omitempty" name:"RecordTime"`

	// 入带宽数据。
	PeakInNetwork *string `json:"PeakInNetwork,omitempty" name:"PeakInNetwork"`

	// 出带宽数据。
	PeakOutNetwork *string `json:"PeakOutNetwork,omitempty" name:"PeakOutNetwork"`

	// 计费带宽。单位bps
	ChargeNetwork *string `json:"ChargeNetwork,omitempty" name:"ChargeNetwork"`
}

type PeakNetworkRegionInfo struct {
	// region信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 网络峰值集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeakNetworkSet []*PeakNetwork `json:"PeakNetworkSet,omitempty" name:"PeakNetworkSet"`
}

type PhysicalPosition struct {
	// 机位
	// 注意：此字段可能返回 null，表示取不到有效值。
	PosId *string `json:"PosId,omitempty" name:"PosId"`

	// 机架
	// 注意：此字段可能返回 null，表示取不到有效值。
	RackId *string `json:"RackId,omitempty" name:"RackId"`

	// 交换机
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchId *string `json:"SwitchId,omitempty" name:"SwitchId"`
}

type Placement struct {
	// 云硬盘所属的[可用区](/document/product/213/15753#ZoneInfo)。该参数也可以通过调用  [DescribeZones](/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 围笼Id。作为入参时，表示对指定的CageId的资源进行操作，可为空。 作为出参时，表示资源所属围笼ID，可为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CageId *string `json:"CageId,omitempty" name:"CageId"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](/document/api/378/4400) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 独享集群名字。作为入参时，忽略。作为出参时，表示云硬盘所属的独享集群名，可为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcName *string `json:"CdcName,omitempty" name:"CdcName"`

	// 实例所属的独享集群ID。作为入参时，表示对指定的CdcId独享集群的资源进行操作，可为空。 作为出参时，表示资源所属的独享集群的ID，可为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
}

type Position struct {
	// 实例所在的Zone的信息。
	ZoneInfo *ZoneInfo `json:"ZoneInfo,omitempty" name:"ZoneInfo"`

	// 实例所在的国家的信息。
	Country *Country `json:"Country,omitempty" name:"Country"`

	// 实例所在的Area的信息。
	Area *Area `json:"Area,omitempty" name:"Area"`

	// 实例所在的省份的信息。
	Province *Province `json:"Province,omitempty" name:"Province"`

	// 实例所在的城市的信息。
	City *City `json:"City,omitempty" name:"City"`

	// 实例所在的Region的信息。
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`
}

type PriceDetail struct {
	// 表示折扣，20 表示20%，打2折
	Discount *uint64 `json:"Discount,omitempty" name:"Discount"`

	// 打折后价格，单位分
	DiscountPrice *uint64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// 折扣前价格，单位分
	OriginalPrice *uint64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`
}

type PrivateIPAddressInfo struct {
	// 实例的内网ip。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIPAddress *string `json:"PrivateIPAddress,omitempty" name:"PrivateIPAddress"`
}

type PrivateIpAddressSpecification struct {
	// 内网IP地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// 是否是主IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// 公网IP地址。
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// EIP实例ID，例如：eip-11112222。
	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`

	// 内网IP描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 公网IP是否被封堵。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty" name:"IsWanIpBlocked"`

	// IP状态：
	// PENDING：生产中
	// MIGRATING：迁移中
	// DELETING：删除中
	// AVAILABLE：可用的
	State *string `json:"State,omitempty" name:"State"`
}

type Province struct {
	// 省份Id
	ProvinceId *string `json:"ProvinceId,omitempty" name:"ProvinceId"`

	// 省份名称
	ProvinceName *string `json:"ProvinceName,omitempty" name:"ProvinceName"`
}

type PublicIPAddressInfo struct {
	// 计费模式。
	ChargeMode *string `json:"ChargeMode,omitempty" name:"ChargeMode"`

	// 实例的公网ip。
	PublicIPAddress *string `json:"PublicIPAddress,omitempty" name:"PublicIPAddress"`

	// 实例的公网ip所属的运营商。
	ISP *ISP `json:"ISP,omitempty" name:"ISP"`

	// 实例的最大出带宽上限，单位为Mbps。
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`

	// 实例的最大入带宽上限，单位为Mbps。
	MaxBandwidthIn *int64 `json:"MaxBandwidthIn,omitempty" name:"MaxBandwidthIn"`
}

// Predefined struct for user
type RebootInstancesRequestParams struct {
	// 待重启的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 是否在正常重启失败后选择强制重启实例。取值范围：
	// TRUE：表示在正常重启失败后进行强制重启；
	// FALSE：表示在正常重启失败后不进行强制重启；
	// 默认取值：FALSE。
	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`

	// 关机类型。取值范围：
	// SOFT：表示软关机
	// HARD：表示硬关机
	// SOFT_FIRST：表示优先软关机，失败再执行硬关机
	// 
	// 默认取值：SOFT。
	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 待重启的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 是否在正常重启失败后选择强制重启实例。取值范围：
	// TRUE：表示在正常重启失败后进行强制重启；
	// FALSE：表示在正常重启失败后不进行强制重启；
	// 默认取值：FALSE。
	ForceReboot *bool `json:"ForceReboot,omitempty" name:"ForceReboot"`

	// 关机类型。取值范围：
	// SOFT：表示软关机
	// HARD：表示硬关机
	// SOFT_FIRST：表示优先软关机，失败再执行硬关机
	// 
	// 默认取值：SOFT。
	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "ForceReboot")
	delete(f, "StopType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebootInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebootInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RebootInstancesResponseParams `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebootInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// Region
	Region *string `json:"Region,omitempty" name:"Region"`

	// Region名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// RegionID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

// Predefined struct for user
type ReleaseAddressesRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID 列表。
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
}

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID 列表。
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds"`
}

func (r *ReleaseAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "AddressIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseAddressesResponseParams struct {
	// 异步任务TaskId。可以使用DescribeTaskResult接口查询任务状态。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseAddressesResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseAddressesResponseParams `json:"Response"`
}

func (r *ReleaseAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseIpv6AddressesRequestParams struct {
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，形如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的IPv6地址列表，单次最多指定10个。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

type ReleaseIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，形如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的IPv6地址列表，单次最多指定10个。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitempty" name:"Ipv6Addresses"`
}

func (r *ReleaseIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIpv6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseIpv6AddressesResponseParams struct {
	// 任务ID，可以通过DescribeTaskResult查询任务状态
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseIpv6AddressesResponseParams `json:"Response"`
}

func (r *ReleaseIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePrivateIpAddressesRequestParams struct {
	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，例如：eni-11112222。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}

type RemovePrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// ECM 地域，形如ap-xian-ecm。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，例如：eni-11112222。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
}

func (r *RemovePrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePrivateIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EcmRegion")
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemovePrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePrivateIpAddressesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemovePrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *RemovePrivateIpAddressesResponseParams `json:"Response"`
}

func (r *RemovePrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRouteTableAssociationRequestParams struct {
	// 子网实例ID，例如：subnet-3x5lf5q0。可通过DescribeSubnets接口查询。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

type ReplaceRouteTableAssociationRequest struct {
	*tchttp.BaseRequest
	
	// 子网实例ID，例如：subnet-3x5lf5q0。可通过DescribeSubnets接口查询。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *ReplaceRouteTableAssociationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRouteTableAssociationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	delete(f, "RouteTableId")
	delete(f, "EcmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceRouteTableAssociationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRouteTableAssociationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceRouteTableAssociationResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceRouteTableAssociationResponseParams `json:"Response"`
}

func (r *ReplaceRouteTableAssociationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRouteTableAssociationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRoutesRequestParams struct {
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略对象。
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

type ReplaceRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由策略对象。
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *ReplaceRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceRoutesResponseParams `json:"Response"`
}

func (r *ReplaceRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPolicyRequestParams struct {
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组规则集合对象。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

type ReplaceSecurityGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如esg-33ocnj9n，可通过DescribeSecurityGroups获取
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组规则集合对象。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *ReplaceSecurityGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceSecurityGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceSecurityGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceSecurityGroupPolicyResponseParams `json:"Response"`
}

func (r *ReplaceSecurityGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesMaxBandwidthRequestParams struct {
	// 待重置带宽上限的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 修改后的最大出带宽上限。
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`

	// 修改后的最大入带宽上限。
	MaxBandwidthIn *int64 `json:"MaxBandwidthIn,omitempty" name:"MaxBandwidthIn"`
}

type ResetInstancesMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// 待重置带宽上限的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 修改后的最大出带宽上限。
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`

	// 修改后的最大入带宽上限。
	MaxBandwidthIn *int64 `json:"MaxBandwidthIn,omitempty" name:"MaxBandwidthIn"`
}

func (r *ResetInstancesMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "MaxBandwidthOut")
	delete(f, "MaxBandwidthIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesMaxBandwidthResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetInstancesMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesMaxBandwidthResponseParams `json:"Response"`
}

func (r *ResetInstancesMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordRequestParams struct {
	// 待重置密码的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 新密码，Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9]和[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]中的符号。密码不允许以/符号开头。
	// Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9]和[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]中的符号。密码不允许以/符号开头。
	// 如果实例即包含Linux实例又包含Windows实例，则密码复杂度限制按照Windows实例的限制。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否强制关机，默认为false。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// 待重置密码的实例的用户名，不得超过64个字符。若未指定用户名，则对于Linux而言，默认重置root用户的密码，对于Windows而言，默认重置administrator的密码。
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 待重置密码的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 新密码，Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9]和[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]中的符号。密码不允许以/符号开头。
	// Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9]和[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]中的符号。密码不允许以/符号开头。
	// 如果实例即包含Linux实例又包含Windows实例，则密码复杂度限制按照Windows实例的限制。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否强制关机，默认为false。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// 待重置密码的实例的用户名，不得超过64个字符。若未指定用户名，则对于Linux而言，默认重置root用户的密码，对于Windows而言，默认重置administrator的密码。
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *ResetInstancesPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "Password")
	delete(f, "ForceStop")
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesPasswordResponseParams `json:"Response"`
}

func (r *ResetInstancesPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesRequestParams struct {
	// 待重装的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 重装使用的镜像ID，若未指定，则使用各个实例当前的镜像进行重装。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 密码设置，若未指定，则后续将以站内信的形式通知密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否开启云监控和云镜服务，未指定时默认开启。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 是否保留数据盘数据，取值"true"/"false"。默认为"true"
	KeepData *string `json:"KeepData,omitempty" name:"KeepData"`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：
	// TRUE：表示保持镜像的登录设置
	// FALSE：表示不保持镜像的登录设置
	// 
	// 默认取值：FALSE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ResetInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 待重装的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 重装使用的镜像ID，若未指定，则使用各个实例当前的镜像进行重装。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 密码设置，若未指定，则后续将以站内信的形式通知密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否开启云监控和云镜服务，未指定时默认开启。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 是否保留数据盘数据，取值"true"/"false"。默认为"true"
	KeepData *string `json:"KeepData,omitempty" name:"KeepData"`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：
	// TRUE：表示保持镜像的登录设置
	// FALSE：表示不保持镜像的登录设置
	// 
	// 默认取值：FALSE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

func (r *ResetInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "ImageId")
	delete(f, "Password")
	delete(f, "EnhancedService")
	delete(f, "KeepData")
	delete(f, "KeepImageLogin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ResetInstancesResponseParams `json:"Response"`
}

func (r *ResetInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRoutesRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 路由策略。
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

type ResetRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 路由策略。
	Routes []*Route `json:"Routes,omitempty" name:"Routes"`
}

func (r *ResetRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteTableName")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ResetRoutesResponseParams `json:"Response"`
}

func (r *ResetRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Route struct {
	// 目的IPv4网段
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// 下一跳类型
	// NORMAL_CVM：普通云服务器；
	GatewayType *string `json:"GatewayType,omitempty" name:"GatewayType"`

	// 下一跳地址
	// 这里只需要指定不同下一跳类型的网关ID，系统会自动匹配到下一跳地址
	// 当 GatewayType 为 EIP 时，GatewayId 固定值 '0'
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// 路由策略唯一ID
	RouteItemId *string `json:"RouteItemId,omitempty" name:"RouteItemId"`

	// 路由策略描述
	RouteDescription *string `json:"RouteDescription,omitempty" name:"RouteDescription"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 路由类型，目前我们支持的类型有：
	// USER：用户路由；
	// NETD：网络探测路由，创建网络探测实例时，系统默认下发，不可编辑与删除；
	// CCN：云联网路由，系统默认下发，不可编辑与删除。
	// 用户只能添加和操作 USER 类型的路由。
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// 路由策略ID。IPv4路由策略ID是有意义的值，IPv6路由策略是无意义的值0。后续建议完全使用字符串唯一ID `RouteItemId`操作路由策略
	RouteId *uint64 `json:"RouteId,omitempty" name:"RouteId"`
}

type RouteConflict struct {
	// 路由表实例ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 要检查的与之冲突的目的端
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// 冲突的路由策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConflictSet []*Route `json:"ConflictSet,omitempty" name:"ConflictSet"`
}

type RouteTable struct {
	// VPC实例ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 路由表实例ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`

	// 路由表关联关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociationSet []*RouteTableAssociation `json:"AssociationSet,omitempty" name:"AssociationSet"`

	// IPv4路由策略集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteSet []*Route `json:"RouteSet,omitempty" name:"RouteSet"`

	// 是否默认路由表
	Main *bool `json:"Main,omitempty" name:"Main"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type RouteTableAssociation struct {
	// 子网实例ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 路由表实例ID
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type RuleHealth struct {
	// 本规则上绑定的后端的健康检查状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*TargetHealth `json:"Targets,omitempty" name:"Targets"`
}

type RunEIPDirectServiceEnabled struct {
	// 是否开通IP直通。取值范围：
	// TRUE：表示开通IP直通
	// FALSE：表示不开通IP直通
	// 默认取值：TRUE。
	// windows镜像目前不支持IP直通。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

// Predefined struct for user
type RunInstancesRequestParams struct {
	// 需要创建实例的可用区及创建数目及运营商的列表。在单次请求的过程中，单个region下的请求创建实例数上限为100
	ZoneInstanceCountISPSet []*ZoneInstanceCountISP `json:"ZoneInstanceCountISPSet,omitempty" name:"ZoneInstanceCountISPSet"`

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：
	// Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? / ]中的特殊符。Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? /]中的特殊符号。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 公网出带宽上限，单位：Mbps。
	// 1.如果未传该参数或者传的值为0，则使用模块下的默认值。
	// 2.如果未传该参数或者传的值为0且未指定模块，则使用InternetMaxBandwidthIn的值
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 模块ID。如果未传该参数，则必须传ImageId，InstanceType，DataDiskSize，InternetMaxBandwidthOut参数
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 镜像ID。如果未传该参数或者传的值为空，则使用模块下的默认值
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例显示名称。
	// 不指定实例显示名称则默认显示‘未命名’。
	// 购买多台实例，如果指定模式串{R:x}，表示生成数字[x, x+n-1]，其中n表示购买实例的数量，例如server\_{R:3}，购买1台时，实例显示名称为server\_3；购买2台时，实例显示名称分别为server\_3，server\_4。
	// 支持指定多个模式串{R:x}。
	// 购买多台实例，如果不指定模式串，则在实例显示名称添加后缀1、2...n，其中n表示购买实例的数量，例如server_，购买2台时，实例显示名称分别为server\_1，server\_2。
	// 如果购买的实例属于不同的地域或运营商，则上述规则在每个地域和运营商内独立计数。
	// 最多支持60个字符（包含模式串）。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 主机名称
	// 点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。
	// Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。
	// 其他类型（Linux 等）实例：字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 用于保证请求幂等性的字符串。目前为保留参数，请勿使用。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 标签列表
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 机型。如果未传该参数或者传的值为空，则使用模块下的默认值
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 数据盘大小，单位是G。如果未传该参数或者传的值为0，则使用模块下的默认值
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 系统盘大小，单位是G。如果未传该参数或者传的值为0，则使用模块下的默认值
	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 公网入带宽上限，单位：Mbps。
	// 1.如果未传该参数或者传的值为0，则使用对应模块的默认值。
	// 2.如果未传该参数或者传的值为0且未指定模块，则使用InternetMaxBandwidthOut
	InternetMaxBandwidthIn *int64 `json:"InternetMaxBandwidthIn,omitempty" name:"InternetMaxBandwidthIn"`

	// 实例计费类型。其中：
	// 0，按资源维度后付费，计算当日用量峰值，例如CPU，内存，硬盘等，仅适用于非GNR系列机型；
	// 1，按小时后付费，单价：xx元/实例/小时，仅适用于GNR机型，如需开通该计费方式请提工单申请；
	// 2，按月后付费，单价：xx元/实例/月，仅适用于GNR机型；
	// 该字段不填时，非GNR机型会默认选择0；GNR机型默认选择2。
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 密钥对。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：
	// TRUE：表示保持镜像的登录设置
	// FALSE：表示不保持镜像的登录设置
	// 
	// 默认取值：FALSE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`

	// 系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘信息。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 需要创建实例的可用区及创建数目及运营商的列表。在单次请求的过程中，单个region下的请求创建实例数上限为100
	ZoneInstanceCountISPSet []*ZoneInstanceCountISP `json:"ZoneInstanceCountISPSet,omitempty" name:"ZoneInstanceCountISPSet"`

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：
	// Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? / ]中的特殊符。Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? /]中的特殊符号。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 公网出带宽上限，单位：Mbps。
	// 1.如果未传该参数或者传的值为0，则使用模块下的默认值。
	// 2.如果未传该参数或者传的值为0且未指定模块，则使用InternetMaxBandwidthIn的值
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 模块ID。如果未传该参数，则必须传ImageId，InstanceType，DataDiskSize，InternetMaxBandwidthOut参数
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 镜像ID。如果未传该参数或者传的值为空，则使用模块下的默认值
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例显示名称。
	// 不指定实例显示名称则默认显示‘未命名’。
	// 购买多台实例，如果指定模式串{R:x}，表示生成数字[x, x+n-1]，其中n表示购买实例的数量，例如server\_{R:3}，购买1台时，实例显示名称为server\_3；购买2台时，实例显示名称分别为server\_3，server\_4。
	// 支持指定多个模式串{R:x}。
	// 购买多台实例，如果不指定模式串，则在实例显示名称添加后缀1、2...n，其中n表示购买实例的数量，例如server_，购买2台时，实例显示名称分别为server\_1，server\_2。
	// 如果购买的实例属于不同的地域或运营商，则上述规则在每个地域和运营商内独立计数。
	// 最多支持60个字符（包含模式串）。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 主机名称
	// 点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。
	// Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。
	// 其他类型（Linux 等）实例：字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 用于保证请求幂等性的字符串。目前为保留参数，请勿使用。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认公共镜像开启云监控、云安全服务
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 标签列表
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 机型。如果未传该参数或者传的值为空，则使用模块下的默认值
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 数据盘大小，单位是G。如果未传该参数或者传的值为0，则使用模块下的默认值
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" name:"DataDiskSize"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 系统盘大小，单位是G。如果未传该参数或者传的值为0，则使用模块下的默认值
	SystemDiskSize *uint64 `json:"SystemDiskSize,omitempty" name:"SystemDiskSize"`

	// 公网入带宽上限，单位：Mbps。
	// 1.如果未传该参数或者传的值为0，则使用对应模块的默认值。
	// 2.如果未传该参数或者传的值为0且未指定模块，则使用InternetMaxBandwidthOut
	InternetMaxBandwidthIn *int64 `json:"InternetMaxBandwidthIn,omitempty" name:"InternetMaxBandwidthIn"`

	// 实例计费类型。其中：
	// 0，按资源维度后付费，计算当日用量峰值，例如CPU，内存，硬盘等，仅适用于非GNR系列机型；
	// 1，按小时后付费，单价：xx元/实例/小时，仅适用于GNR机型，如需开通该计费方式请提工单申请；
	// 2，按月后付费，单价：xx元/实例/月，仅适用于GNR机型；
	// 该字段不填时，非GNR机型会默认选择0；GNR机型默认选择2。
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 密钥对。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：
	// TRUE：表示保持镜像的登录设置
	// FALSE：表示不保持镜像的登录设置
	// 
	// 默认取值：FALSE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`

	// 系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘信息。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

func (r *RunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneInstanceCountISPSet")
	delete(f, "Password")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "ModuleId")
	delete(f, "ImageId")
	delete(f, "InstanceName")
	delete(f, "HostName")
	delete(f, "ClientToken")
	delete(f, "EnhancedService")
	delete(f, "TagSpecification")
	delete(f, "UserData")
	delete(f, "InstanceType")
	delete(f, "DataDiskSize")
	delete(f, "SecurityGroupIds")
	delete(f, "SystemDiskSize")
	delete(f, "InternetMaxBandwidthIn")
	delete(f, "InstanceChargeType")
	delete(f, "KeyIds")
	delete(f, "KeepImageLogin")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunInstancesResponseParams struct {
	// 创建中的实例ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RunInstancesResponseParams `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {
	// 是否开启。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// 是否开启。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 云镜版本：0 基础版，1 专业版。目前仅支持基础版
	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type SecurityGroup struct {
	// 安全组实例ID，例如：esg-ohuuioma。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称，可任意命名，但不得超过60个字符。
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组备注，最多100个字符。
	SecurityGroupDesc *string `json:"SecurityGroupDesc,omitempty" name:"SecurityGroupDesc"`

	// 是否是默认安全组，默认安全组不支持删除。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 安全组创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 标签键值对。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`
}

type SecurityGroupAssociationStatistics struct {
	// 安全组实例ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// ECM实例数。
	ECM *int64 `json:"ECM,omitempty" name:"ECM"`

	// ECM模块数。
	Module *int64 `json:"Module,omitempty" name:"Module"`

	// 弹性网卡实例数。
	ENI *int64 `json:"ENI,omitempty" name:"ENI"`

	// 被安全组引用数。
	SG *int64 `json:"SG,omitempty" name:"SG"`

	// 负载均衡实例数。
	CLB *int64 `json:"CLB,omitempty" name:"CLB"`

	// 全量实例的绑定统计。
	InstanceStatistics []*InstanceStatistic `json:"InstanceStatistics,omitempty" name:"InstanceStatistics"`

	// 所有资源的总计数（不包含被安全组引用数）。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type SecurityGroupLimitSet struct {
	// 可创建安全组总数
	SecurityGroupLimit *int64 `json:"SecurityGroupLimit,omitempty" name:"SecurityGroupLimit"`

	// 安全组下的最大规则数
	SecurityGroupPolicyLimit *int64 `json:"SecurityGroupPolicyLimit,omitempty" name:"SecurityGroupPolicyLimit"`

	// 安全组下嵌套安全组规则数
	ReferedSecurityGroupLimit *int64 `json:"ReferedSecurityGroupLimit,omitempty" name:"ReferedSecurityGroupLimit"`

	// 单安全组关联实例数
	SecurityGroupInstanceLimit *int64 `json:"SecurityGroupInstanceLimit,omitempty" name:"SecurityGroupInstanceLimit"`

	// 实例关联安全组数
	InstanceSecurityGroupLimit *int64 `json:"InstanceSecurityGroupLimit,omitempty" name:"InstanceSecurityGroupLimit"`

	// 单安全组关联的模块数
	SecurityGroupModuleLimit *int64 `json:"SecurityGroupModuleLimit,omitempty" name:"SecurityGroupModuleLimit"`

	// 模块关联的安全组数
	ModuleSecurityGroupLimit *int64 `json:"ModuleSecurityGroupLimit,omitempty" name:"ModuleSecurityGroupLimit"`
}

type SecurityGroupPolicy struct {
	// 安全组规则索引号
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// 协议, 取值: TCP,UDP, ICMP。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口(all, 离散port, range)。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 协议端口ID或者协议端口组ID。ServiceTemplate和Protocol+Port互斥。
	ServiceTemplate *ServiceTemplateSpecification `json:"ServiceTemplate,omitempty" name:"ServiceTemplate"`

	// 网段或IP(互斥)。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 安全组实例ID，例如：esg-ohuuioma。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// IP地址ID或者ID地址组ID。
	AddressTemplate *AddressTemplateSpecification `json:"AddressTemplate,omitempty" name:"AddressTemplate"`

	// ACCEPT 或 DROP。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 安全组规则描述。
	PolicyDescription *string `json:"PolicyDescription,omitempty" name:"PolicyDescription"`

	// 修改时间，例如 2020-07-22 19：27：23
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 网段或IPv6(互斥)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`
}

type SecurityGroupPolicySet struct {
	// 安全组规则当前版本。用户每次更新安全规则版本会自动加1，防止更新的路由规则已过期，不填不考虑冲突。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 出站规则。其中出站规则和入站规则必须选一个。
	Egress []*SecurityGroupPolicy `json:"Egress,omitempty" name:"Egress"`

	// 入站规则。其中出站规则和入站规则必须选一个。
	Ingress []*SecurityGroupPolicy `json:"Ingress,omitempty" name:"Ingress"`
}

type ServiceTemplateSpecification struct {
	// 协议端口ID，例如：eppm-f5n1f8da。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 协议端口组ID，例如：eppmg-f5n1f8da。
	ServiceGroupId *string `json:"ServiceGroupId,omitempty" name:"ServiceGroupId"`
}

// Predefined struct for user
type SetLoadBalancerSecurityGroupsRequestParams struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 安全组ID构成的数组，一个负载均衡实例最多可绑定5个安全组，如果要解绑所有安全组，可不传此参数，或传入空数组
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

type SetLoadBalancerSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 安全组ID构成的数组，一个负载均衡实例最多可绑定5个安全组，如果要解绑所有安全组，可不传此参数，或传入空数组
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

func (r *SetLoadBalancerSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SecurityGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetLoadBalancerSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLoadBalancerSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetLoadBalancerSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *SetLoadBalancerSecurityGroupsResponseParams `json:"Response"`
}

func (r *SetLoadBalancerSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetSecurityGroupForLoadbalancersRequestParams struct {
	// 负载均衡实例ID数组
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 安全组ID，如 esg-12345678
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// ADD 绑定安全组；
	// DEL 解绑安全组
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
}

type SetSecurityGroupForLoadbalancersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID数组
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 安全组ID，如 esg-12345678
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// ADD 绑定安全组；
	// DEL 解绑安全组
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`
}

func (r *SetSecurityGroupForLoadbalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSecurityGroupForLoadbalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "SecurityGroup")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetSecurityGroupForLoadbalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetSecurityGroupForLoadbalancersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetSecurityGroupForLoadbalancersResponse struct {
	*tchttp.BaseResponse
	Response *SetSecurityGroupForLoadbalancersResponseParams `json:"Response"`
}

func (r *SetSecurityGroupForLoadbalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSecurityGroupForLoadbalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SimpleModule struct {
	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type Snapshot struct {
	// 快照所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 是否为跨地域复制的快照。取值范围：<br><li>true：表示为跨地域复制的快照。<br><li>false:本地域的快照。
	CopyFromRemote *bool `json:"CopyFromRemote,omitempty" name:"CopyFromRemote"`

	// 是否为永久快照。取值范围：<br><li>true：永久快照<br><li>false：非永久快照。
	IsPermanent *bool `json:"IsPermanent,omitempty" name:"IsPermanent"`

	// 快照名称，用户自定义的快照别名。调用[ModifySnapshotAttribute](/document/product/362/15650)可修改此字段。
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 快照创建进度百分比，快照创建成功后此字段恒为100。
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

	// 快照关联的镜像列表。
	Images []*Image `json:"Images,omitempty" name:"Images"`

	// 快照当前被共享数。
	ShareReference *uint64 `json:"ShareReference,omitempty" name:"ShareReference"`

	// 快照类型，目前该项取值可以为PRIVATE_SNAPSHOT或者SHARED_SNAPSHOT
	SnapshotType *string `json:"SnapshotType,omitempty" name:"SnapshotType"`

	// 创建此快照的云硬盘大小，单位GB。
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 创建此快照的云硬盘ID。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 快照正在跨地域复制的目的地域，默认取值为[]。
	CopyingToRegions []*string `json:"CopyingToRegions,omitempty" name:"CopyingToRegions"`

	// 快照ID。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 创建此快照的云硬盘类型。取值范围：<br><li>SYSTEM_DISK：系统盘<br><li>DATA_DISK：数据盘。
	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`

	// 是否为加密盘创建的快照。取值范围：<br><li>true：该快照为加密盘创建的<br><li>false:非加密盘创建的快照。
	Encrypt *bool `json:"Encrypt,omitempty" name:"Encrypt"`

	// 快照的创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 快照关联的镜像个数。
	ImageCount *uint64 `json:"ImageCount,omitempty" name:"ImageCount"`

	// 快照的状态。取值范围：<br><li>NORMAL：正常<br><li>CREATING：创建中<br><li>ROLLBACKING：回滚中<br><li>COPYING_FROM_REMOTE：跨地域复制中<br><li>CHECKING_COPIED：复制校验中<br><li>TORECYCLE：待回收。
	SnapshotState *string `json:"SnapshotState,omitempty" name:"SnapshotState"`

	// 快照的到期时间。
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// 快照开始共享的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeStartShare *string `json:"TimeStartShare,omitempty" name:"TimeStartShare"`
}

type SrcImage struct {
	// 镜像id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" name:"ImageName"`

	// 系统名称
	ImageOsName *string `json:"ImageOsName,omitempty" name:"ImageOsName"`

	// 镜像描述
	ImageDescription *string `json:"ImageDescription,omitempty" name:"ImageDescription"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 区域ID
	RegionID *int64 `json:"RegionID,omitempty" name:"RegionID"`

	// 区域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 来源实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 来源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 来源镜像类型
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`
}

// Predefined struct for user
type StartInstancesRequestParams struct {
	// 待开启的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 待开启的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StartInstancesResponseParams `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesRequestParams struct {
	// 需要关机的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 是否在正常关闭失败后选择强制关闭实例，默认为false，即否。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// 实例的关闭模式。取值范围：
	// SOFT_FIRST：表示在正常关闭失败后进行强制关闭;
	// HARD：直接强制关闭;
	// SOFT：仅软关机；
	// 默认为SOFT。
	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 需要关机的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 是否在正常关闭失败后选择强制关闭实例，默认为false，即否。
	ForceStop *bool `json:"ForceStop,omitempty" name:"ForceStop"`

	// 实例的关闭模式。取值范围：
	// SOFT_FIRST：表示在正常关闭失败后进行强制关闭;
	// HARD：直接强制关闭;
	// SOFT：仅软关机；
	// 默认为SOFT。
	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

func (r *StopInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "ForceStop")
	delete(f, "StopType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StopInstancesResponseParams `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Subnet struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网实例ID，例如：subnet-bthucmmy。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名称。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网的 IPv4 CIDR。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 是否默认子网。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 是否开启广播。
	EnableBroadcast *bool `json:"EnableBroadcast,omitempty" name:"EnableBroadcast"`

	// 路由表实例ID，例如：rtb-l2h8d7c2。
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 可用IP数。
	AvailableIpAddressCount *uint64 `json:"AvailableIpAddressCount,omitempty" name:"AvailableIpAddressCount"`

	// 子网的 IPv6 CIDR。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// 关联ACLID
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`

	// 是否为 SNAT 地址池子网。
	IsRemoteVpcSnat *bool `json:"IsRemoteVpcSnat,omitempty" name:"IsRemoteVpcSnat"`

	// 标签键值对。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// 所在区域
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// VPC的 IPv4 CIDR。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// VPC的 IPv6 CIDR。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcIpv6CidrBlock *string `json:"VpcIpv6CidrBlock,omitempty" name:"VpcIpv6CidrBlock"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type SystemDisk struct {
	// 硬盘类型。取值范围：
	// - LOCAL_BASIC：本地硬盘；
	// - CLOUD_PREMIUM：高性能云硬盘；
	// 默认取值：CLOUD_BASIC。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 硬盘ID。此参数暂不可用。
	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`

	// 硬盘容量大小。单位GB。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type Tag struct {
	// 标签健。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagInfo struct {
	// 标签的键。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签的值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagSpecification struct {
	// 资源类型，目前仅支持"instance"、"module"
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type Target struct {
	// 后端服务的监听端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 子机ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 绑定弹性网卡时需要传入此参数，代表弹性网卡的IP，弹性网卡必须先绑定至子机，然后才能绑定到负载均衡实例。注意：参数 InstanceId 和 EniIp 只能传入一个且必须传入一个。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`
}

type TargetHealth struct {
	// Target的内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitempty" name:"IP"`

	// Target绑定的端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 当前健康状态，true：健康，false：不健康（包括尚未开始探测、探测中、状态异常等几种状态）。只有处于健康状态（且权重大于0），负载均衡才会向其转发流量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthStatus *bool `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// Target的实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// 当前健康状态的详细信息。如：Alive、Dead、Unknown、Close。Alive状态为健康，Dead状态为异常，Unknown状态包括尚未开始探测、探测中、状态未知，Close为未配置健康检查。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthStatusDetail *string `json:"HealthStatusDetail,omitempty" name:"HealthStatusDetail"`
}

type TargetsWeightRule struct {
	// 负载均衡监听器 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改权重的后端机器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 后端服务新的转发权重，取值范围：0~100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type TaskInput struct {
	// 操作名，即API名称，比如：CreateImage
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type TaskOutput struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 状态描述
	Message *string `json:"Message,omitempty" name:"Message"`

	// 状态值，SUCCESS/FAILED/OPERATING
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务提交时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

// Predefined struct for user
type TerminateDisksRequestParams struct {
	// 需退还的云盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

type TerminateDisksRequest struct {
	*tchttp.BaseRequest
	
	// 需退还的云盘ID列表。
	DiskIds []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (r *TerminateDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DiskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDisksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateDisksResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDisksResponseParams `json:"Response"`
}

func (r *TerminateDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstancesRequestParams struct {
	// 待销毁的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 是否定时销毁，默认为否。
	TerminateDelay *bool `json:"TerminateDelay,omitempty" name:"TerminateDelay"`

	// 定时销毁的时间，格式形如："2019-08-05 12:01:30"，若非定时销毁，则此参数被忽略。
	TerminateTime *string `json:"TerminateTime,omitempty" name:"TerminateTime"`

	// 是否关联删除已绑定的弹性网卡和弹性IP，默认为true。
	// 当为true时，一并删除弹性网卡和弹性IP；
	// 当为false时，只销毁主机，保留弹性网卡和弹性IP。
	AssociatedResourceDestroy *bool `json:"AssociatedResourceDestroy,omitempty" name:"AssociatedResourceDestroy"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 待销毁的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 是否定时销毁，默认为否。
	TerminateDelay *bool `json:"TerminateDelay,omitempty" name:"TerminateDelay"`

	// 定时销毁的时间，格式形如："2019-08-05 12:01:30"，若非定时销毁，则此参数被忽略。
	TerminateTime *string `json:"TerminateTime,omitempty" name:"TerminateTime"`

	// 是否关联删除已绑定的弹性网卡和弹性IP，默认为true。
	// 当为true时，一并删除弹性网卡和弹性IP；
	// 当为false时，只销毁主机，保留弹性网卡和弹性IP。
	AssociatedResourceDestroy *bool `json:"AssociatedResourceDestroy,omitempty" name:"AssociatedResourceDestroy"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "TerminateDelay")
	delete(f, "TerminateTime")
	delete(f, "AssociatedResourceDestroy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateInstancesResponseParams `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {
	// 私有网络ID，形如vpc-xxx。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如subnet-xxx。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：
	// TRUE：表示用作公网网关
	// FALSE：表示不用作公网网关
	// 
	// 默认取值：FALSE。
	AsVpcGateway *bool `json:"AsVpcGateway,omitempty" name:"AsVpcGateway"`

	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 为弹性网卡指定随机生成的 IPv6 地址数量。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

type VpcInfo struct {
	// VPC名称。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPC实例ID，例如：vpc-azd4dt1c。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// VPC的IPv4 CIDR。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 是否默认VPC。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 是否开启组播。
	EnableMulticast *bool `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// DNS列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsServerSet []*string `json:"DnsServerSet,omitempty" name:"DnsServerSet"`

	// DHCP域名选项值。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// DHCP选项集ID。
	DhcpOptionsId *string `json:"DhcpOptionsId,omitempty" name:"DhcpOptionsId"`

	// 是否开启DHCP。
	EnableDhcp *bool `json:"EnableDhcp,omitempty" name:"EnableDhcp"`

	// VPC的IPv6 CIDR。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// 标签键值对
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// 辅助CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 地域中文名
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 包含子网数量
	SubnetCount *uint64 `json:"SubnetCount,omitempty" name:"SubnetCount"`

	// 包含实例数量
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`
}

type ZoneInfo struct {
	// ZoneId
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// ZoneName
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Zone
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ZoneInstanceCountISP struct {
	// 创建实例的可用区。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 在当前可用区欲创建的实例数目。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 运营商如下：
	// CTCC：中国电信
	// CUCC：中国联通
	// CMCC：中国移动
	// 多个运营商用英文分号连接";"，例如："CMCC;CUCC;CTCC"。多运营商需要开通白名单，请直接联系腾讯云客服。
	ISP *string `json:"ISP,omitempty" name:"ISP"`

	// 指定私有网络编号，SubnetId与VpcId必须同时指定或不指定
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 指定子网编号，SubnetId与VpcId必须同时指定或不指定
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 指定主网卡内网IP。条件：SubnetId与VpcId必须同时指定，并且IP数量与InstanceCount相同，多IP主机副网卡内网IP在相同子网内通过DHCP获取。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 为弹性网卡指定随机生成的IPv6地址数量，目前数量不能大于1。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" name:"Ipv6AddressCount"`
}

type ZoneInstanceInfo struct {
	// Zone名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}