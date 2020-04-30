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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *AllocateAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 申请到的 EIP 的唯一 ID 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AddressSet []*string `json:"AddressSet,omitempty" name:"AddressSet" list`

		// 异步任务TaskId。可以使用DescribeTaskResult接口查询任务状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Area struct {

	// 区域ID
	AreaId *string `json:"AreaId,omitempty" name:"AreaId"`

	// 区域名称
	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
}

type AssignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 指定的内网IP信息，单次最多指定10个。与SecondaryPrivateIpAddressCount至少提供一个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// 新申请的内网IP地址个数，与PrivateIpAddresses至少提供一个。内网IP地址个数总和不能超过配额数
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`
}

func (r *AssignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignPrivateIpAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内网IP详细信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet" list`
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

func (r *AssociateAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务TaskId。可以使用DescribeTaskResult接口查询任务状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 实例ID。形如：ein-r8hr2upy。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *AttachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *CreateModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模块ID，创建模块成功后分配给该模块的ID。
		ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`

	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitempty" name:"NetworkInterfaceDescription"`

	// 新申请的内网IP地址个数，内网IP地址个数总和不能超过配数。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`

	// 指定绑定的安全组，例如：['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 弹性网卡实例。
		NetworkInterface *NetworkInterface `json:"NetworkInterface,omitempty" name:"NetworkInterface"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
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

		// 子网对象。
		Subnet *Subnet `json:"Subnet,omitempty" name:"Subnet"`

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

type CreateVpcRequest struct {
	*tchttp.BaseRequest

	// vpc名称，最大长度不能超过60个字节。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// vpc的cidr，只能为10.0.0.0/16，172.16.0.0/16，192.168.0.0/16这三个内网网段内。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 是否开启组播。true: 开启, false: 不开启。
	EnableMulticast *string `json:"EnableMulticast,omitempty" name:"EnableMulticast"`

	// DNS地址，最多支持4个
	DnsServers []*string `json:"DnsServers,omitempty" name:"DnsServers" list`

	// 域名
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
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

		// Vpc对象。
		Vpc *VpcInfo `json:"Vpc,omitempty" name:"Vpc"`

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

type DeleteImageRequest struct {
	*tchttp.BaseRequest

	// 镜像ID列表。
	ImageIDSet []*string `json:"ImageIDSet,omitempty" name:"ImageIDSet" list`
}

func (r *DeleteImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleRequest struct {
	*tchttp.BaseRequest

	// 模块ID。如：es-qn46snq8
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *DeleteModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DeleteNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteSubnetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

func (r *DeleteVpcRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DescribeAddressQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账户 EIP 配额信息。
		QuotaSet []*EipQuota `json:"QuotaSet,omitempty" name:"QuotaSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：eip-11112222。参数不支持同时指定AddressIds和Filters。
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`

	// 每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定AddressIds和Filters。详细的过滤条件如下：
	// address-id - String - 是否必填：否 - （过滤条件）按照 EIP 的唯一 ID 过滤。EIP 唯一 ID 形如：eip-11112222。
	// address-name - String - 是否必填：否 - （过滤条件）按照 EIP 名称过滤。不支持模糊过滤。
	// address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。
	// address-status - String - 是否必填：否 - （过滤条件）按照 EIP 的状态过滤。取值范围：详见EIP状态列表。
	// instance-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的实例 ID 过滤。实例 ID 形如：ins-11112222。
	// private-ip-address - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的内网 IP 过滤。
	// network-interface-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的弹性网卡 ID 过滤。弹性网卡 ID 形如：eni-11112222。
	// is-arrears - String - 是否必填：否 - （过滤条件）按照 EIP 是否欠费进行过滤。（TRUE：EIP 处于欠费状态|FALSE：EIP 费用状态正常）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的 EIP 数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// EIP 详细信息列表。
		AddressSet []*Address `json:"AddressSet,omitempty" name:"AddressSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBaseOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *DescribeBaseOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBaseOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络带宽硬盘大小的范围信息。
		NetworkStorageRange *NetworkStorageRange `json:"NetworkStorageRange,omitempty" name:"NetworkStorageRange"`

		// 镜像操作系统白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImageWhiteSet []*string `json:"ImageWhiteSet,omitempty" name:"ImageWhiteSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，每次请求的Filters的上限为10，详细的过滤条件如下：
	// image-id - String - 是否必填： 否 - （过滤条件）按照镜像ID进行过滤
	// image-type - String - 是否必填： 否 - （过滤条件）按照镜像类型进行过滤。取值范围：
	// PRIVATE_IMAGE: 私有镜像 (本帐户创建的镜像) 
	// PUBLIC_IMAGE: 公共镜像 (腾讯云官方镜像)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 镜像总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 镜像数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImageSet []*Image `json:"ImageSet,omitempty" name:"ImageSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceTypeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 机型配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeInstanceVncUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceVncUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例的管理终端地址。
		InstanceVncUrl *string `json:"InstanceVncUrl,omitempty" name:"InstanceVncUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceVncUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceVncUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDeniedActionsRequest struct {
	*tchttp.BaseRequest

	// 无
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`
}

func (r *DescribeInstancesDeniedActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesDeniedActionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDeniedActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例对应的禁止操作
		InstanceOperatorSet []*InstanceOperator `json:"InstanceOperatorSet,omitempty" name:"InstanceOperatorSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesDeniedActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesDeniedActionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// zone      String      是否必填：否     （过滤条件）按照可用区中文名过滤,支持模糊匹配。
	// module-id      String      是否必填：否     （过滤条件）按照模块ID过滤。
	// instance-id      String      是否必填：否      （过滤条件）按照实例ID过滤。
	// instance-name      String      是否必填：否      （过滤条件）按照实例名称过滤,支持模糊匹配。
	// ip-address      String      是否必填：否      （过滤条件）按照实例的内网/公网IP过滤。
	// instance-uuid   string 是否必填：否 （过滤条件）按照uuid过滤实例列表。
	// instance-state  string  是否必填：否 （过滤条件）按照实例状态更新实例列表。
	// internet-service-provider      String      是否必填：否      （过滤条件）按照实例公网IP所属的运营商进行过滤。
	// tag-key      String      是否必填：否      （过滤条件）按照标签键进行过滤。
	// tag:tag-key      String      是否必填：否      （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	// 若不传Filters参数则表示查询所有相关的实例信息。
	// 单次请求的Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20(如果查询结果数目大于等于20)，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

		// 返回的实例相关信息列表的长度。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的实例相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet" list`

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

type DescribeModuleDetailRequest struct {
	*tchttp.BaseRequest

	// 模块ID，如em-qn46snq8。
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *DescribeModuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模块的详细信息，详细见数据结构中的ModuleInfo。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Module *Module `json:"Module,omitempty" name:"Module"`

		// 模块的统计信息，详细见数据结构中的ModuleCounterInfo。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ModuleCounter *ModuleCounter `json:"ModuleCounter,omitempty" name:"ModuleCounter"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleRequest struct {
	*tchttp.BaseRequest

	// 过滤条件。
	// module-name - string - 是否必填：否 - （过滤条件）按照模块名称过滤。
	// module-id - string - 是否必填：否 - （过滤条件）按照模块ID过滤。
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeModuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的模块数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 模块详情信息的列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ModuleItemSet []*ModuleItem `json:"ModuleItemSet,omitempty" name:"ModuleItemSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesRequest struct {
	*tchttp.BaseRequest

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID查询。形如：eni-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定NetworkInterfaceIds和Filters。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" name:"NetworkInterfaceIds" list`

	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。
	// network-interface-id - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。
	// attachment.instance-id - String - （过滤条件）绑定的云服务器实例ID，形如：ins-3nqpdn3i。
	// groups.security-group-id - String - （过滤条件）绑定的安全组实例ID，例如：sg-f9ekbxeq。
	// network-interface-name - String - （过滤条件）网卡实例名称。
	// network-interface-description - String - （过滤条件）网卡实例描述。
	// address-ip - String - （过滤条件）内网IPv4地址。
	// tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。使用请参考示例2
	// tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例3。
	// is-primary - Boolean - 是否必填：否 - （过滤条件）按照是否主网卡进行过滤。值为true时，仅过滤主网卡；值为false时，仅过滤辅助网卡；次过滤参数为提供时，同时过滤主网卡和辅助网卡。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNetworkInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkInterfacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet,omitempty" name:"NetworkInterfaceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNetworkInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNetworkInterfacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 节点详细信息的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		NodeSet []*Node `json:"NodeSet,omitempty" name:"NodeSet" list`

		// 所有的节点数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePeakBaseOverviewRequest struct {
	*tchttp.BaseRequest

	// 开始时间（xxxx-xx-xx）如2019-08-14，默认为一周之前的日期。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（xxxx-xx-xx）如2019-08-14，默认为昨天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribePeakBaseOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePeakBaseOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePeakBaseOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 基础峰值列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PeakFamilyInfoSet []*PeakFamilyInfo `json:"PeakFamilyInfoSet,omitempty" name:"PeakFamilyInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePeakBaseOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePeakBaseOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePeakNetworkOverviewRequest struct {
	*tchttp.BaseRequest

	// 开始时间（xxxx-xx-xx）如2019-08-14，默认为一周之前的日期。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间（xxxx-xx-xx）如2019-08-14，默认为昨天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 过滤条件。
	// region    String      是否必填：否     （过滤条件）按照region过滤,不支持模糊匹配。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribePeakNetworkOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePeakNetworkOverviewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePeakNetworkOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网络峰值数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PeakNetworkRegionSet []*PeakNetworkRegionInfo `json:"PeakNetworkRegionSet,omitempty" name:"PeakNetworkRegionSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePeakNetworkOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePeakNetworkOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 子网实例ID查询。形如：subnet-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// subnet-id - String - （过滤条件）Subnet实例名称。
	// vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。
	// cidr-block - String - （过滤条件）子网网段，形如: 192.168.1.0 。
	// is-default - Boolean - （过滤条件）是否是默认子网。
	// is-remote-vpc-snat - Boolean - （过滤条件）是否为VPC SNAT地址池子网。
	// subnet-name - String - （过滤条件）子网名称。
	// zone - String - （过滤条件）可用区。
	// tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。
	// tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *string `json:"Limit,omitempty" name:"Limit"`
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

		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 子网对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetSet []*Subnet `json:"SubnetSet,omitempty" name:"SubnetSet" list`

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

func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务ID。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 执行结果，包括"SUCCESS", "FAILED", "RUNNING"
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds" list`

	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// vpc-name - String - （过滤条件）VPC实例名称。
	// is-default - String - （过滤条件）是否默认VPC。
	// vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。
	// cidr-block - String - （过滤条件）vpc的cidr。
	// tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。
	// tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
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

		// 符合条件的对象数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 私有网络对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
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

type DetachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 实例ID。形如：ein-hcs7jkg4
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`
}

func (r *DetachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DisassociateAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务TaskId。可以使用DescribeTaskResult接口查询任务状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type EnhancedService struct {

	// 是否开启云镜服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// 是否开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type Filter struct {

	// 过滤字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段内容数组
	Values []*string `json:"Values,omitempty" name:"Values" list`
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
	ZoneInstanceInfoSet []*ZoneInstanceInfo `json:"ZoneInstanceInfoSet,omitempty" name:"ZoneInstanceInfoSet" list`
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

func (r *ImportImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

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
	DataDisks []*DiskInfo `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 新实例标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewFlag *int64 `json:"NewFlag,omitempty" name:"NewFlag"`
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

type InstanceOperator struct {

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例禁止的操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeniedActions []*OperatorAction `json:"DeniedActions,omitempty" name:"DeniedActions" list`
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

	// 机型额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtInfo *string `json:"ExtInfo,omitempty" name:"ExtInfo"`
}

type Internet struct {

	// 实例的内网相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIPAddressSet []*PrivateIPAddressInfo `json:"PrivateIPAddressSet,omitempty" name:"PrivateIPAddressSet" list`

	// 实例的公网相关信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIPAddressSet []*PublicIPAddressInfo `json:"PublicIPAddressSet,omitempty" name:"PublicIPAddressSet" list`
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

type MigrateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest

	// ECM 地域
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

func (r *MigrateNetworkInterfaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MigrateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigrateNetworkInterfaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MigratePrivateIpAddressRequest struct {
	*tchttp.BaseRequest

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 当内网IP绑定的弹性网卡实例ID，例如：eni-11112222。
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

func (r *MigratePrivateIpAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MigratePrivateIpAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigratePrivateIpAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MigratePrivateIpAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyAddressAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// EIP唯一标识ID，形如'eip-xxxxxxx'
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`

	// 调整带宽目标值
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

func (r *ModifyAddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressesBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务TaskId。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressesBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest

	// 待修改的实例ID列表。在单次请求的过程中，请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// 修改成功后显示的实例名称，不得超过60个字符，不传则名称显示为空。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyModuleImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyModuleNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleNetworkRequest struct {
	*tchttp.BaseRequest

	// 模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 默认带宽上限
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`
}

func (r *ModifyModuleNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleNetworkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleNetworkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModuleNetworkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type ModifyVpcAttributeRequest struct {
	*tchttp.BaseRequest

	// VPC实例ID。形如：vpc-f49l6u0z。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 私有网络名称，可任意命名，但不得超过60个字符。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
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

type Module struct {

	// 模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`

	// 模块状态：
	// NORMAL：正常
	// DELETING：删除中 
	// DELETEFAILED：删除失败
	ModuleState *string `json:"ModuleState,omitempty" name:"ModuleState"`

	// 默认系统盘大小
	DefaultSystemDiskSize *int64 `json:"DefaultSystemDiskSize,omitempty" name:"DefaultSystemDiskSize"`

	// 默认数据盘大小
	DefaultDataDiskSize *int64 `json:"DefaultDataDiskSize,omitempty" name:"DefaultDataDiskSize"`

	// 默认机型
	InstanceTypeConfig *InstanceTypeConfig `json:"InstanceTypeConfig,omitempty" name:"InstanceTypeConfig"`

	// 默认镜像
	DefaultImage *Image `json:"DefaultImage,omitempty" name:"DefaultImage"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 默认带宽
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" name:"DefaultBandwidth"`
}

type ModuleCounter struct {

	// 运营商统计信息列表
	ISPCounterSet []*ISPCounter `json:"ISPCounterSet,omitempty" name:"ISPCounterSet" list`

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
	GroupSet []*string `json:"GroupSet,omitempty" name:"GroupSet" list`

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
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitempty" name:"PrivateIpAddressSet" list`

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
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitempty" name:"Ipv6AddressSet" list`

	// 标签键值对。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// 网卡类型。0 - 弹性网卡；1 - evm弹性网卡。
	EniType *uint64 `json:"EniType,omitempty" name:"EniType"`
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

	// zone信息
	ZoneInfo *ZoneInfo `json:"ZoneInfo,omitempty" name:"ZoneInfo"`

	// 国家信息
	Country *Country `json:"Country,omitempty" name:"Country"`

	// 区域信息
	Area *Area `json:"Area,omitempty" name:"Area"`

	// 省份信息
	Province *Province `json:"Province,omitempty" name:"Province"`

	// 城市信息
	City *City `json:"City,omitempty" name:"City"`

	// Region信息
	RegionInfo *RegionInfo `json:"RegionInfo,omitempty" name:"RegionInfo"`

	// 运营商列表
	ISPSet []*ISP `json:"ISPSet,omitempty" name:"ISPSet" list`

	// 运营商数量
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
	PeakBaseSet []*PeakBase `json:"PeakBaseSet,omitempty" name:"PeakBaseSet" list`
}

type PeakNetwork struct {

	// 记录时间。
	RecordTime *string `json:"RecordTime,omitempty" name:"RecordTime"`

	// 入带宽数据。
	PeakInNetwork *string `json:"PeakInNetwork,omitempty" name:"PeakInNetwork"`

	// 出带宽数据。
	PeakOutNetwork *string `json:"PeakOutNetwork,omitempty" name:"PeakOutNetwork"`
}

type PeakNetworkRegionInfo struct {

	// region信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 网络峰值集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeakNetworkSet []*PeakNetwork `json:"PeakNetworkSet,omitempty" name:"PeakNetworkSet" list`
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

	// 实例的最大出带宽上限。
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest

	// 待重启的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

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

func (r *RebootInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest

	// ECM 地域
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 标识 EIP 的唯一 ID 列表。
	AddressIds []*string `json:"AddressIds,omitempty" name:"AddressIds" list`
}

func (r *ReleaseAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务TaskId。可以使用DescribeTaskResult接口查询任务状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemovePrivateIpAddressesRequest struct {
	*tchttp.BaseRequest

	// ECM 地域。
	EcmRegion *string `json:"EcmRegion,omitempty" name:"EcmRegion"`

	// 弹性网卡实例ID，例如：eni-11112222。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses" list`
}

func (r *RemovePrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemovePrivateIpAddressesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemovePrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemovePrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemovePrivateIpAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesMaxBandwidthRequest struct {
	*tchttp.BaseRequest

	// 待重置带宽上限的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// 修改后的最大带宽上限。
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`
}

func (r *ResetInstancesMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest

	// 待重置密码的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

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

func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesRequest struct {
	*tchttp.BaseRequest

	// 待重装的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// 重装使用的镜像ID，若未指定，则使用各个实例当前的镜像进行重装。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 密码设置，若未指定，则后续将以站内信的形式通知密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否开启云监控和云镜服务，未指定时默认开启。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

func (r *ResetInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest

	// 需要创建实例的可用区及创建数目及运营商的列表。在单次请求的过程中，单个region下的请求创建实例数上限为100
	ZoneInstanceCountISPSet []*ZoneInstanceCountISP `json:"ZoneInstanceCountISPSet,omitempty" name:"ZoneInstanceCountISPSet" list`

	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：
	// Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? / ]中的特殊符。Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & - + = | { } [ ] : ; ' , . ? /]中的特殊符号。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 公网出带宽上限，单位：Mbps
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 镜像ID，不传则使用模块下的默认值
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
	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification" list`

	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB
	UserData *string `json:"UserData,omitempty" name:"UserData"`
}

func (r *RunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建中的实例ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type SimpleModule struct {

	// 模块ID
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 模块名称
	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
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
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest

	// 待开启的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest

	// 需要关机的实例ID列表。在单次请求的过程中，单个region下的请求实例数上限为100。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

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

func (r *StopInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`
}

type Tag struct {

	// 标签的键。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签的值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagSpecification struct {

	// 资源类型，目前仅支持"instance"
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 待销毁的实例ID列表。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

	// 是否定时销毁，默认为否。
	TerminateDelay *bool `json:"TerminateDelay,omitempty" name:"TerminateDelay"`

	// 定时销毁的时间，格式形如："2019-08-05 12:01:30"，若非定时销毁，则此参数被忽略。
	TerminateTime *string `json:"TerminateTime,omitempty" name:"TerminateTime"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	DnsServerSet []*string `json:"DnsServerSet,omitempty" name:"DnsServerSet" list`

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
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

	// 辅助CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitempty" name:"AssistantCidrSet" list`
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

	// 运营商。
	ISP *string `json:"ISP,omitempty" name:"ISP"`
}

type ZoneInstanceInfo struct {

	// Zone名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}
