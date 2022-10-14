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

package v20180410

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AcceptDirectConnectTunnelRequestParams struct {
	// 物理专线拥有者接受共享专用通道申请
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

type AcceptDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// 物理专线拥有者接受共享专用通道申请
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *AcceptDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptDirectConnectTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcceptDirectConnectTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptDirectConnectTunnelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AcceptDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *AcceptDirectConnectTunnelResponseParams `json:"Response"`
}

func (r *AcceptDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessPoint struct {
	// 接入点的名称。
	AccessPointName *string `json:"AccessPointName,omitempty" name:"AccessPointName"`

	// 接入点唯一ID。
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 接入点的状态。可用，不可用。
	State *string `json:"State,omitempty" name:"State"`

	// 接入点的位置。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 接入点支持的运营商列表。
	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 接入点管理的大区ID。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 接入点可用的端口类型列表。1000BASE-T代表千兆电口，1000BASE-LX代表千兆单模光口10km，1000BASE-ZX代表千兆单模光口80km,10GBASE-LR代表万兆单模光口10km,10GBASE-ZR代表万兆单模光口80km,10GBASE-LH代表万兆单模光口40km,100GBASE-LR4代表100G单模光口10km
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailablePortType []*string `json:"AvailablePortType,omitempty" name:"AvailablePortType"`

	// 接入点经纬度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coordinate *Coordinate `json:"Coordinate,omitempty" name:"Coordinate"`

	// 接入点所在城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 接入点地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 接入点类型。VXLAN/QCPL/QCAR
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessPointType *string `json:"AccessPointType,omitempty" name:"AccessPointType"`
}

// Predefined struct for user
type ApplyInternetAddressRequestParams struct {
	// CIDR地址掩码长度
	MaskLen *int64 `json:"MaskLen,omitempty" name:"MaskLen"`

	// 0:BGP类型地址
	// 1：中国电信
	// 2：中国移动
	// 3：中国联通
	AddrType *int64 `json:"AddrType,omitempty" name:"AddrType"`

	// 0：IPv4
	// 1:IPv6
	AddrProto *int64 `json:"AddrProto,omitempty" name:"AddrProto"`
}

type ApplyInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// CIDR地址掩码长度
	MaskLen *int64 `json:"MaskLen,omitempty" name:"MaskLen"`

	// 0:BGP类型地址
	// 1：中国电信
	// 2：中国移动
	// 3：中国联通
	AddrType *int64 `json:"AddrType,omitempty" name:"AddrType"`

	// 0：IPv4
	// 1:IPv6
	AddrProto *int64 `json:"AddrProto,omitempty" name:"AddrProto"`
}

func (r *ApplyInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaskLen")
	delete(f, "AddrType")
	delete(f, "AddrProto")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyInternetAddressResponseParams struct {
	// 互联网公网地址ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *ApplyInternetAddressResponseParams `json:"Response"`
}

func (r *ApplyInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BFDInfo struct {
	// 健康检查次数
	ProbeFailedTimes *int64 `json:"ProbeFailedTimes,omitempty" name:"ProbeFailedTimes"`

	// 健康检查间隔
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
}

type BGPStatus struct {
	// 腾讯侧主互联IP BGP状态
	TencentAddressBgpState *string `json:"TencentAddressBgpState,omitempty" name:"TencentAddressBgpState"`

	// 腾讯侧备互联IP BGP状态
	TencentBackupAddressBgpState *string `json:"TencentBackupAddressBgpState,omitempty" name:"TencentBackupAddressBgpState"`
}

type BgpPeer struct {
	// 用户侧，BGP Asn
	Asn *int64 `json:"Asn,omitempty" name:"Asn"`

	// 用户侧BGP密钥
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`
}

type Coordinate struct {
	// 纬度
	Lat *float64 `json:"Lat,omitempty" name:"Lat"`

	// 经度
	Lng *float64 `json:"Lng,omitempty" name:"Lng"`
}

// Predefined struct for user
type CreateDirectConnectRequestParams struct {
	// 物理专线的名称。
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// 物理专线所在的接入点。
	// 您可以通过调用 DescribeAccessPoints接口获取地域ID。所选择的接入点必须存在且处于可接入的状态。
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 提供接入物理专线的运营商。ChinaTelecom：中国电信， ChinaMobile：中国移动，ChinaUnicom：中国联通， In-houseWiring：楼内线，ChinaOther：中国其他， InternationalOperator：境外其他。
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 物理专线接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。
	PortType *string `json:"PortType,omitempty" name:"PortType"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// 本地数据中心的地理位置。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 物理专线接入接口带宽，单位为Mbps，默认值为1000，取值范围为 [2, 10240]。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 冗余物理专线的ID。
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`

	// 物理专线调试VLAN。默认开启VLAN，自动分配VLAN。
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联 IP。默认自动分配。
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联 IP。默认自动分配。
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`

	// 物理专线申请者是否签署了用户使用协议。默认已签署
	SignLaw *bool `json:"SignLaw,omitempty" name:"SignLaw"`
}

type CreateDirectConnectRequest struct {
	*tchttp.BaseRequest
	
	// 物理专线的名称。
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// 物理专线所在的接入点。
	// 您可以通过调用 DescribeAccessPoints接口获取地域ID。所选择的接入点必须存在且处于可接入的状态。
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 提供接入物理专线的运营商。ChinaTelecom：中国电信， ChinaMobile：中国移动，ChinaUnicom：中国联通， In-houseWiring：楼内线，ChinaOther：中国其他， InternationalOperator：境外其他。
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 物理专线接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。
	PortType *string `json:"PortType,omitempty" name:"PortType"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// 本地数据中心的地理位置。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 物理专线接入接口带宽，单位为Mbps，默认值为1000，取值范围为 [2, 10240]。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 冗余物理专线的ID。
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`

	// 物理专线调试VLAN。默认开启VLAN，自动分配VLAN。
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联 IP。默认自动分配。
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联 IP。默认自动分配。
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`

	// 物理专线申请者是否签署了用户使用协议。默认已签署
	SignLaw *bool `json:"SignLaw,omitempty" name:"SignLaw"`
}

func (r *CreateDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectName")
	delete(f, "AccessPointId")
	delete(f, "LineOperator")
	delete(f, "PortType")
	delete(f, "CircuitCode")
	delete(f, "Location")
	delete(f, "Bandwidth")
	delete(f, "RedundantDirectConnectId")
	delete(f, "Vlan")
	delete(f, "TencentAddress")
	delete(f, "CustomerAddress")
	delete(f, "CustomerName")
	delete(f, "CustomerContactMail")
	delete(f, "CustomerContactNumber")
	delete(f, "FaultReportContactPerson")
	delete(f, "FaultReportContactNumber")
	delete(f, "SignLaw")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectResponseParams struct {
	// 物理专线的ID。
	DirectConnectIdSet []*string `json:"DirectConnectIdSet,omitempty" name:"DirectConnectIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *CreateDirectConnectResponseParams `json:"Response"`
}

func (r *CreateDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectTunnelRequestParams struct {
	// 专线 ID，例如：dc-kd7d06of
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 物理专线 owner，缺省为当前客户（物理专线 owner）
	// 共享专线时这里需要填写共享专线的开发商账号 ID
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`

	// 网络类型，分别为VPC、BMVPC，CCN，默认是VPC
	// VPC：私有网络
	// BMVPC：黑石网络
	// CCN：云联网
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 网络地域
	NetworkRegion *string `json:"NetworkRegion,omitempty" name:"NetworkRegion"`

	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 专线网关 ID，例如 dcg-d545ddf
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// 专线带宽，单位：Mbps
	// 默认是物理专线带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// BGP ：BGP路由
	// STATIC：静态
	// 默认为 BGP 路由
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 静态路由，用户IDC的网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes"`

	// vlan，范围：0 ~ 3000
	// 0：不开启子接口
	// 默认值是非0
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// TencentAddress，腾讯侧互联 IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// CustomerAddress，用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// TencentBackupAddress，腾讯侧备用互联 IP
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`

	// 高速上云服务ID
	CloudAttachId *string `json:"CloudAttachId,omitempty" name:"CloudAttachId"`

	// 是否开启BFD
	BfdEnable *int64 `json:"BfdEnable,omitempty" name:"BfdEnable"`

	// 是否开启NQA
	NqaEnable *int64 `json:"NqaEnable,omitempty" name:"NqaEnable"`

	// BFD配置信息
	BfdInfo *BFDInfo `json:"BfdInfo,omitempty" name:"BfdInfo"`

	// NQA配置信息
	NqaInfo *NQAInfo `json:"NqaInfo,omitempty" name:"NqaInfo"`
}

type CreateDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// 专线 ID，例如：dc-kd7d06of
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 物理专线 owner，缺省为当前客户（物理专线 owner）
	// 共享专线时这里需要填写共享专线的开发商账号 ID
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`

	// 网络类型，分别为VPC、BMVPC，CCN，默认是VPC
	// VPC：私有网络
	// BMVPC：黑石网络
	// CCN：云联网
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 网络地域
	NetworkRegion *string `json:"NetworkRegion,omitempty" name:"NetworkRegion"`

	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 专线网关 ID，例如 dcg-d545ddf
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// 专线带宽，单位：Mbps
	// 默认是物理专线带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// BGP ：BGP路由
	// STATIC：静态
	// 默认为 BGP 路由
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 静态路由，用户IDC的网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes"`

	// vlan，范围：0 ~ 3000
	// 0：不开启子接口
	// 默认值是非0
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// TencentAddress，腾讯侧互联 IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// CustomerAddress，用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// TencentBackupAddress，腾讯侧备用互联 IP
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`

	// 高速上云服务ID
	CloudAttachId *string `json:"CloudAttachId,omitempty" name:"CloudAttachId"`

	// 是否开启BFD
	BfdEnable *int64 `json:"BfdEnable,omitempty" name:"BfdEnable"`

	// 是否开启NQA
	NqaEnable *int64 `json:"NqaEnable,omitempty" name:"NqaEnable"`

	// BFD配置信息
	BfdInfo *BFDInfo `json:"BfdInfo,omitempty" name:"BfdInfo"`

	// NQA配置信息
	NqaInfo *NQAInfo `json:"NqaInfo,omitempty" name:"NqaInfo"`
}

func (r *CreateDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectId")
	delete(f, "DirectConnectTunnelName")
	delete(f, "DirectConnectOwnerAccount")
	delete(f, "NetworkType")
	delete(f, "NetworkRegion")
	delete(f, "VpcId")
	delete(f, "DirectConnectGatewayId")
	delete(f, "Bandwidth")
	delete(f, "RouteType")
	delete(f, "BgpPeer")
	delete(f, "RouteFilterPrefixes")
	delete(f, "Vlan")
	delete(f, "TencentAddress")
	delete(f, "CustomerAddress")
	delete(f, "TencentBackupAddress")
	delete(f, "CloudAttachId")
	delete(f, "BfdEnable")
	delete(f, "NqaEnable")
	delete(f, "BfdInfo")
	delete(f, "NqaInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectTunnelResponseParams struct {
	// 专用通道ID
	DirectConnectTunnelIdSet []*string `json:"DirectConnectTunnelIdSet,omitempty" name:"DirectConnectTunnelIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *CreateDirectConnectTunnelResponseParams `json:"Response"`
}

func (r *CreateDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectRequestParams struct {
	// 物理专线的ID。
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
}

type DeleteDirectConnectRequest struct {
	*tchttp.BaseRequest
	
	// 物理专线的ID。
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
}

func (r *DeleteDirectConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDirectConnectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDirectConnectResponseParams `json:"Response"`
}

func (r *DeleteDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectTunnelRequestParams struct {
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

type DeleteDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *DeleteDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDirectConnectTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectTunnelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDirectConnectTunnelResponseParams `json:"Response"`
}

func (r *DeleteDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessPointsRequestParams struct {
	// 接入点所在的地域。使用DescribeRegions查询
	// 
	// 您可以通过调用 DescribeRegions接口获取地域ID。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAccessPointsRequest struct {
	*tchttp.BaseRequest
	
	// 接入点所在的地域。使用DescribeRegions查询
	// 
	// 您可以通过调用 DescribeRegions接口获取地域ID。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccessPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessPointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegionId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessPointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessPointsResponseParams struct {
	// 接入点信息。
	AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet"`

	// 符合接入点数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccessPointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessPointsResponseParams `json:"Response"`
}

func (r *DescribeAccessPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectTunnelExtraRequestParams struct {
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

type DescribeDirectConnectTunnelExtraRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *DescribeDirectConnectTunnelExtraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectTunnelExtraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDirectConnectTunnelExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectTunnelExtraResponseParams struct {
	// 专用通道扩展信息
	DirectConnectTunnelExtra *DirectConnectTunnelExtra `json:"DirectConnectTunnelExtra,omitempty" name:"DirectConnectTunnelExtra"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDirectConnectTunnelExtraResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDirectConnectTunnelExtraResponseParams `json:"Response"`
}

func (r *DescribeDirectConnectTunnelExtraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectTunnelExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectTunnelsRequestParams struct {
	// 过滤条件:
	// 参数不支持同时指定DirectConnectTunnelIds和Filters。
	// <li> direct-connect-tunnel-name, 专用通道名称。</li>
	// <li> direct-connect-tunnel-id, 专用通道实例ID，如dcx-abcdefgh。</li>
	// <li>direct-connect-id, 物理专线实例ID，如，dc-abcdefgh。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 专用通道 ID数组
	DirectConnectTunnelIds []*string `json:"DirectConnectTunnelIds,omitempty" name:"DirectConnectTunnelIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDirectConnectTunnelsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件:
	// 参数不支持同时指定DirectConnectTunnelIds和Filters。
	// <li> direct-connect-tunnel-name, 专用通道名称。</li>
	// <li> direct-connect-tunnel-id, 专用通道实例ID，如dcx-abcdefgh。</li>
	// <li>direct-connect-id, 物理专线实例ID，如，dc-abcdefgh。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 专用通道 ID数组
	DirectConnectTunnelIds []*string `json:"DirectConnectTunnelIds,omitempty" name:"DirectConnectTunnelIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectTunnelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectTunnelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "DirectConnectTunnelIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDirectConnectTunnelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectTunnelsResponseParams struct {
	// 专用通道列表
	DirectConnectTunnelSet []*DirectConnectTunnel `json:"DirectConnectTunnelSet,omitempty" name:"DirectConnectTunnelSet"`

	// 符合专用通道数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDirectConnectTunnelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDirectConnectTunnelsResponseParams `json:"Response"`
}

func (r *DescribeDirectConnectTunnelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectTunnelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectsRequestParams struct {
	// 过滤条件:
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 物理专线 ID数组
	DirectConnectIds []*string `json:"DirectConnectIds,omitempty" name:"DirectConnectIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDirectConnectsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件:
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 物理专线 ID数组
	DirectConnectIds []*string `json:"DirectConnectIds,omitempty" name:"DirectConnectIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "DirectConnectIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDirectConnectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectsResponseParams struct {
	// 物理专线列表。
	DirectConnectSet []*DirectConnect `json:"DirectConnectSet,omitempty" name:"DirectConnectSet"`

	// 符合物理专线列表数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 用户名下物理专线是否都签署了用户协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllSignLaw *bool `json:"AllSignLaw,omitempty" name:"AllSignLaw"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDirectConnectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDirectConnectsResponseParams `json:"Response"`
}

func (r *DescribeDirectConnectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressQuotaRequestParams struct {

}

type DescribeInternetAddressQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInternetAddressQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternetAddressQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressQuotaResponseParams struct {
	// IPv6互联网公网允许的最小前缀长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6PrefixLen *int64 `json:"Ipv6PrefixLen,omitempty" name:"Ipv6PrefixLen"`

	// BGP类型IPv4互联网地址配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv4BgpQuota *int64 `json:"Ipv4BgpQuota,omitempty" name:"Ipv4BgpQuota"`

	// 非BGP类型IPv4互联网地址配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv4OtherQuota *int64 `json:"Ipv4OtherQuota,omitempty" name:"Ipv4OtherQuota"`

	// BGP类型IPv4互联网地址已使用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv4BgpNum *int64 `json:"Ipv4BgpNum,omitempty" name:"Ipv4BgpNum"`

	// 非BGP类型互联网地址已使用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv4OtherNum *int64 `json:"Ipv4OtherNum,omitempty" name:"Ipv4OtherNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInternetAddressQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternetAddressQuotaResponseParams `json:"Response"`
}

func (r *DescribeInternetAddressQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件：
	// <li>AddrType, 地址类型。0：BGP 1; 1: 电信， 2：移动， 3：联通</li>
	// <li>AddrProto地址类型。0：IPv4 1:IPv6</li>
	// <li>Status 地址状态。 0：使用中， 1：已停用， 2：已退还</li>
	// <li>Subnet 互联网公网地址，数组</li>
	// <InstanceIds>互联网公网地址ID，数组</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件：
	// <li>AddrType, 地址类型。0：BGP 1; 1: 电信， 2：移动， 3：联通</li>
	// <li>AddrProto地址类型。0：IPv4 1:IPv6</li>
	// <li>Status 地址状态。 0：使用中， 1：已停用， 2：已退还</li>
	// <li>Subnet 互联网公网地址，数组</li>
	// <InstanceIds>互联网公网地址ID，数组</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressResponseParams struct {
	// 互联网公网地址数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 互联网公网地址列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subnets []*InternetAddressDetail `json:"Subnets,omitempty" name:"Subnets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternetAddressResponseParams `json:"Response"`
}

func (r *DescribeInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressStatisticsRequestParams struct {

}

type DescribeInternetAddressStatisticsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInternetAddressStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternetAddressStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternetAddressStatisticsResponseParams struct {
	// 互联网公网地址统计信息数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 互联网公网地址统计信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetAddressStatistics []*InternetAddressStatistics `json:"InternetAddressStatistics,omitempty" name:"InternetAddressStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInternetAddressStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternetAddressStatisticsResponseParams `json:"Response"`
}

func (r *DescribeInternetAddressStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternetAddressStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicDirectConnectTunnelRoutesRequestParams struct {
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 过滤条件：
	// route-type：路由类型，取值：BGP/STATIC
	// route-subnet：路由cidr，取值如：192.68.1.0/24
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePublicDirectConnectTunnelRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 过滤条件：
	// route-type：路由类型，取值：BGP/STATIC
	// route-subnet：路由cidr，取值如：192.68.1.0/24
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePublicDirectConnectTunnelRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicDirectConnectTunnelRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicDirectConnectTunnelRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicDirectConnectTunnelRoutesResponseParams struct {
	// 互联网通道路由列表
	Routes []*DirectConnectTunnelRoute `json:"Routes,omitempty" name:"Routes"`

	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePublicDirectConnectTunnelRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicDirectConnectTunnelRoutesResponseParams `json:"Response"`
}

func (r *DescribePublicDirectConnectTunnelRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicDirectConnectTunnelRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DirectConnect struct {
	// 物理专线ID。
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 物理专线的名称。
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// 物理专线的接入点ID。
	AccessPointId *string `json:"AccessPointId,omitempty" name:"AccessPointId"`

	// 物理专线的状态。
	// 申请中：PENDING 
	// 申请驳回：REJECTED   
	// 待付款：TOPAY 
	// 已付款：PAID 
	// 建设中：ALLOCATED   
	// 已开通：AVAILABLE  
	// 删除中 ：DELETING
	// 已删除：DELETED 。
	State *string `json:"State,omitempty" name:"State"`

	// 物理专线创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 物理专线的开通时间。
	EnabledTime *string `json:"EnabledTime,omitempty" name:"EnabledTime"`

	// 提供接入物理专线的运营商。ChinaTelecom：中国电信， ChinaMobile：中国移动，ChinaUnicom：中国联通， In-houseWiring：楼内线，ChinaOther：中国其他， InternationalOperator：境外其他。
	LineOperator *string `json:"LineOperator,omitempty" name:"LineOperator"`

	// 本地数据中心的地理位置。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 物理专线接入接口带宽，单位为Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 用户侧物理专线接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）
	PortType *string `json:"PortType,omitempty" name:"PortType"`

	// 运营商或者服务商为物理专线提供的电路编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// 冗余物理专线的ID。
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitempty" name:"RedundantDirectConnectId"`

	// 物理专线调试VLAN。默认开启VLAN，自动分配VLAN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// 物理专线的过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 物理专线计费类型。 NON_RECURRING_CHARGE：一次性接入费用；PREPAID_BY_YEAR：按年预付费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 报障联系人。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`

	// 标签键值对
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// 物理专线的接入点类型。
	AccessPointType *string `json:"AccessPointType,omitempty" name:"AccessPointType"`

	// IDC所在城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdcCity *string `json:"IdcCity,omitempty" name:"IdcCity"`

	// 计费状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeState *string `json:"ChargeState,omitempty" name:"ChargeState"`

	// 物理专线开通时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 物理专线是否已签署用户协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignLaw *bool `json:"SignLaw,omitempty" name:"SignLaw"`

	// 物理专线是否为LocalZone
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalZone *bool `json:"LocalZone,omitempty" name:"LocalZone"`

	// 该物理专线下vlan 0的专用通道数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VlanZeroDirectConnectTunnelCount *uint64 `json:"VlanZeroDirectConnectTunnelCount,omitempty" name:"VlanZeroDirectConnectTunnelCount"`

	// 该物理专线下非vlan 0的专用通道数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherVlanDirectConnectTunnelCount *uint64 `json:"OtherVlanDirectConnectTunnelCount,omitempty" name:"OtherVlanDirectConnectTunnelCount"`

	// 物理专线最小带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinBandwidth *uint64 `json:"MinBandwidth,omitempty" name:"MinBandwidth"`
}

type DirectConnectTunnel struct {
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 物理专线ID
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 专用通道状态
	// AVAILABLE:就绪或者已连接
	// PENDING:申请中
	// ALLOCATING:配置中
	// ALLOCATED:配置完成
	// ALTERING:修改中
	// DELETING:删除中
	// DELETED:删除完成
	// COMFIRMING:待接受
	// REJECTED:拒绝
	State *string `json:"State,omitempty" name:"State"`

	// 物理专线的拥有者，开发商账号 ID
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`

	// 专用通道的拥有者，开发商账号 ID
	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`

	// 网络类型，分别为VPC、BMVPC、CCN
	//  VPC：私有网络 ，BMVPC：黑石网络，CCN：云联网
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// VPC地域对应的网络名，如ap-guangzhou
	NetworkRegion *string `json:"NetworkRegion,omitempty" name:"NetworkRegion"`

	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 专线网关 ID
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// BGP ：BGP路由 STATIC：静态 默认为 BGP 路由
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// 用户侧BGP，Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 用户侧网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes"`

	// 专用通道的Vlan
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// TencentAddress，腾讯侧互联 IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// CustomerAddress，用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 专用通道创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 专用通道带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 专用通道标签值
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`

	// 关联的网络自定义探测ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`

	// BGP community开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitempty" name:"EnableBGPCommunity"`

	// 是否为Nat通道
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatType *int64 `json:"NatType,omitempty" name:"NatType"`

	// VPC地域简码，如gz、cd
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`

	// 是否开启BFD
	// 注意：此字段可能返回 null，表示取不到有效值。
	BfdEnable *int64 `json:"BfdEnable,omitempty" name:"BfdEnable"`

	// 专用通道接入点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessPointType *string `json:"AccessPointType,omitempty" name:"AccessPointType"`

	// 专线网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`

	// VPC名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// TencentBackupAddress，腾讯侧备用互联 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`

	// 专用通道关联的物理专线是否签署了用户协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignLaw *bool `json:"SignLaw,omitempty" name:"SignLaw"`

	// 高速上云服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudAttachId *string `json:"CloudAttachId,omitempty" name:"CloudAttachId"`
}

type DirectConnectTunnelExtra struct {
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 物理专线ID
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 专用通道状态
	// AVAILABLE:就绪或者已连接
	// PENDING:申请中
	// ALLOCATING:配置中
	// ALLOCATED:配置完成
	// ALTERING:修改中
	// DELETING:删除中
	// DELETED:删除完成
	// COMFIRMING:待接受
	// REJECTED:拒绝
	State *string `json:"State,omitempty" name:"State"`

	// 物理专线的拥有者，开发商账号 ID
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitempty" name:"DirectConnectOwnerAccount"`

	// 专用通道的拥有者，开发商账号 ID
	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`

	// 网络类型，分别为VPC、BMVPC、CCN
	//  VPC：私有网络 ，BMVPC：黑石网络，CCN：云联网
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// VPC地域对应的网络名，如ap-guangzhou
	NetworkRegion *string `json:"NetworkRegion,omitempty" name:"NetworkRegion"`

	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 专线网关 ID
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`

	// BGP ：BGP路由 STATIC：静态 默认为 BGP 路由
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// 用户侧BGP，Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 用户侧网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes"`

	// 互联网通道公网网段地址
	PublicAddresses []*RouteFilterPrefix `json:"PublicAddresses,omitempty" name:"PublicAddresses"`

	// 专用通道的Vlan
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 腾讯侧互联 IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 腾讯侧备用互联IP
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`

	// 用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 专用通道创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 专用通道带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 关联的网络自定义探测ID
	NetDetectId *string `json:"NetDetectId,omitempty" name:"NetDetectId"`

	// BGP community开关
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitempty" name:"EnableBGPCommunity"`

	// 是否为Nat通道
	NatType *int64 `json:"NatType,omitempty" name:"NatType"`

	// VPC地域简码，如gz、cd
	VpcRegion *string `json:"VpcRegion,omitempty" name:"VpcRegion"`

	// 是否开启BFD
	BfdEnable *int64 `json:"BfdEnable,omitempty" name:"BfdEnable"`

	// 是否开启NQA
	NqaEnable *int64 `json:"NqaEnable,omitempty" name:"NqaEnable"`

	// 专用通道接入点类型
	AccessPointType *string `json:"AccessPointType,omitempty" name:"AccessPointType"`

	// 专线网关名称
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`

	// VPC名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 专用通道关联的物理专线是否签署了用户协议
	SignLaw *bool `json:"SignLaw,omitempty" name:"SignLaw"`

	// BFD配置信息
	BfdInfo *BFDInfo `json:"BfdInfo,omitempty" name:"BfdInfo"`

	// NQA配置信息
	NqaInfo *NQAInfo `json:"NqaInfo,omitempty" name:"NqaInfo"`

	// BGP状态
	BgpStatus *BGPStatus `json:"BgpStatus,omitempty" name:"BgpStatus"`

	// 是否开启IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6Enable *int64 `json:"IPv6Enable,omitempty" name:"IPv6Enable"`

	// 腾讯侧互联IPv6地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	TencentIPv6Address *string `json:"TencentIPv6Address,omitempty" name:"TencentIPv6Address"`

	// 腾讯侧备用互联IPv6地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	TencentBackupIPv6Address *string `json:"TencentBackupIPv6Address,omitempty" name:"TencentBackupIPv6Address"`

	// BGPv6状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	BgpIPv6Status *BGPStatus `json:"BgpIPv6Status,omitempty" name:"BgpIPv6Status"`

	// 用户侧互联IPv6地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomerIPv6Address *string `json:"CustomerIPv6Address,omitempty" name:"CustomerIPv6Address"`

	// 专用通道是否支持巨帧。1 支持，0 不支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	JumboEnable *int64 `json:"JumboEnable,omitempty" name:"JumboEnable"`

	// 专用通道是否支持高精度BFD。1支持，0不支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighPrecisionBFDEnable *int64 `json:"HighPrecisionBFDEnable,omitempty" name:"HighPrecisionBFDEnable"`
}

type DirectConnectTunnelRoute struct {
	// 专用通道路由ID
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`

	// 网段CIDR
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`

	// 路由类型：BGP/STATIC路由
	RouteType *string `json:"RouteType,omitempty" name:"RouteType"`

	// ENABLE：路由启用，DISABLE：路由禁用
	Status *string `json:"Status,omitempty" name:"Status"`

	// ASPath信息
	ASPath []*string `json:"ASPath,omitempty" name:"ASPath"`

	// 路由下一跳IP
	NextHop *string `json:"NextHop,omitempty" name:"NextHop"`
}

// Predefined struct for user
type DisableInternetAddressRequestParams struct {
	// 公网互联网地址ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DisableInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// 公网互联网地址ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DisableInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableInternetAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *DisableInternetAddressResponseParams `json:"Response"`
}

func (r *DisableInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableInternetAddressRequestParams struct {
	// 互联网公网地址ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type EnableInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// 互联网公网地址ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *EnableInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableInternetAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *EnableInternetAddressResponseParams `json:"Response"`
}

func (r *EnableInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InternetAddressDetail struct {
	// 互联网地址ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 互联网网络地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`

	// 网络地址掩码长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaskLen *int64 `json:"MaskLen,omitempty" name:"MaskLen"`

	// 0:BGP
	// 1:电信
	// 2:移动
	// 3:联通
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddrType *int64 `json:"AddrType,omitempty" name:"AddrType"`

	// 0:使用中
	// 1:已停用
	// 2:已退还
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 停用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopTime *string `json:"StopTime,omitempty" name:"StopTime"`

	// 退还时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 0:IPv4 1:IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddrProto *int64 `json:"AddrProto,omitempty" name:"AddrProto"`

	// 释放状态的IP地址保留的天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReserveTime *int64 `json:"ReserveTime,omitempty" name:"ReserveTime"`
}

type InternetAddressStatistics struct {
	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 互联网公网地址数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetNum *int64 `json:"SubnetNum,omitempty" name:"SubnetNum"`
}

// Predefined struct for user
type ModifyDirectConnectAttributeRequestParams struct {
	// 物理专线的ID。
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 物理专线名称。
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// 物理专线调试VLAN。
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联 IP。
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联 IP。
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`

	// 物理专线申请者补签用户使用协议
	SignLaw *bool `json:"SignLaw,omitempty" name:"SignLaw"`

	// 物理专线带宽
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

type ModifyDirectConnectAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 物理专线的ID。
	DirectConnectId *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`

	// 物理专线名称。
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitempty" name:"CircuitCode"`

	// 物理专线调试VLAN。
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联 IP。
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联 IP。
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitempty" name:"FaultReportContactNumber"`

	// 物理专线申请者补签用户使用协议
	SignLaw *bool `json:"SignLaw,omitempty" name:"SignLaw"`

	// 物理专线带宽
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

func (r *ModifyDirectConnectAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectId")
	delete(f, "DirectConnectName")
	delete(f, "CircuitCode")
	delete(f, "Vlan")
	delete(f, "TencentAddress")
	delete(f, "CustomerAddress")
	delete(f, "CustomerName")
	delete(f, "CustomerContactMail")
	delete(f, "CustomerContactNumber")
	delete(f, "FaultReportContactPerson")
	delete(f, "FaultReportContactNumber")
	delete(f, "SignLaw")
	delete(f, "Bandwidth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDirectConnectAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDirectConnectAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDirectConnectAttributeResponseParams `json:"Response"`
}

func (r *ModifyDirectConnectAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectTunnelAttributeRequestParams struct {
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 用户侧BGP，包括Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 用户侧网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes"`

	// 腾讯侧互联IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 用户侧互联IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 专用通道带宽值，单位为M。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 腾讯侧备用互联IP
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`
}

type ModifyDirectConnectTunnelAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitempty" name:"DirectConnectTunnelName"`

	// 用户侧BGP，包括Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 用户侧网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes"`

	// 腾讯侧互联IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 用户侧互联IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 专用通道带宽值，单位为M。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 腾讯侧备用互联IP
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`
}

func (r *ModifyDirectConnectTunnelAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectTunnelAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	delete(f, "DirectConnectTunnelName")
	delete(f, "BgpPeer")
	delete(f, "RouteFilterPrefixes")
	delete(f, "TencentAddress")
	delete(f, "CustomerAddress")
	delete(f, "Bandwidth")
	delete(f, "TencentBackupAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDirectConnectTunnelAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectTunnelAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDirectConnectTunnelAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDirectConnectTunnelAttributeResponseParams `json:"Response"`
}

func (r *ModifyDirectConnectTunnelAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectTunnelAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectTunnelExtraRequestParams struct {
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道的Vlan
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 用户侧BGP，Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 用户侧过滤网段地址
	RouteFilterPrefixes *RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes"`

	// 腾讯侧互联IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 腾讯侧备用互联IP
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`

	// 用户侧互联IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 专用通道带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// BGP community开关
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitempty" name:"EnableBGPCommunity"`

	// 是否开启BFD
	BfdEnable *int64 `json:"BfdEnable,omitempty" name:"BfdEnable"`

	// 是否开启NQA
	NqaEnable *int64 `json:"NqaEnable,omitempty" name:"NqaEnable"`

	// BFD配置信息
	BfdInfo *BFDInfo `json:"BfdInfo,omitempty" name:"BfdInfo"`

	// NQA配置信息
	NqaInfo *NQAInfo `json:"NqaInfo,omitempty" name:"NqaInfo"`

	// 0：停用IPv6
	// 1: 启用IPv6
	IPv6Enable *int64 `json:"IPv6Enable,omitempty" name:"IPv6Enable"`

	// 去往用户侧的路由信息
	CustomerIDCRoutes []*RouteFilterPrefix `json:"CustomerIDCRoutes,omitempty" name:"CustomerIDCRoutes"`

	// 是否开启巨帧
	// 1：开启
	// 0：不开启
	JumboEnable *int64 `json:"JumboEnable,omitempty" name:"JumboEnable"`
}

type ModifyDirectConnectTunnelExtraRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道的Vlan
	Vlan *int64 `json:"Vlan,omitempty" name:"Vlan"`

	// 用户侧BGP，Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitempty" name:"BgpPeer"`

	// 用户侧过滤网段地址
	RouteFilterPrefixes *RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes"`

	// 腾讯侧互联IP
	TencentAddress *string `json:"TencentAddress,omitempty" name:"TencentAddress"`

	// 腾讯侧备用互联IP
	TencentBackupAddress *string `json:"TencentBackupAddress,omitempty" name:"TencentBackupAddress"`

	// 用户侧互联IP
	CustomerAddress *string `json:"CustomerAddress,omitempty" name:"CustomerAddress"`

	// 专用通道带宽值
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// BGP community开关
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitempty" name:"EnableBGPCommunity"`

	// 是否开启BFD
	BfdEnable *int64 `json:"BfdEnable,omitempty" name:"BfdEnable"`

	// 是否开启NQA
	NqaEnable *int64 `json:"NqaEnable,omitempty" name:"NqaEnable"`

	// BFD配置信息
	BfdInfo *BFDInfo `json:"BfdInfo,omitempty" name:"BfdInfo"`

	// NQA配置信息
	NqaInfo *NQAInfo `json:"NqaInfo,omitempty" name:"NqaInfo"`

	// 0：停用IPv6
	// 1: 启用IPv6
	IPv6Enable *int64 `json:"IPv6Enable,omitempty" name:"IPv6Enable"`

	// 去往用户侧的路由信息
	CustomerIDCRoutes []*RouteFilterPrefix `json:"CustomerIDCRoutes,omitempty" name:"CustomerIDCRoutes"`

	// 是否开启巨帧
	// 1：开启
	// 0：不开启
	JumboEnable *int64 `json:"JumboEnable,omitempty" name:"JumboEnable"`
}

func (r *ModifyDirectConnectTunnelExtraRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectTunnelExtraRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	delete(f, "Vlan")
	delete(f, "BgpPeer")
	delete(f, "RouteFilterPrefixes")
	delete(f, "TencentAddress")
	delete(f, "TencentBackupAddress")
	delete(f, "CustomerAddress")
	delete(f, "Bandwidth")
	delete(f, "EnableBGPCommunity")
	delete(f, "BfdEnable")
	delete(f, "NqaEnable")
	delete(f, "BfdInfo")
	delete(f, "NqaInfo")
	delete(f, "IPv6Enable")
	delete(f, "CustomerIDCRoutes")
	delete(f, "JumboEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDirectConnectTunnelExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectTunnelExtraResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDirectConnectTunnelExtraResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDirectConnectTunnelExtraResponseParams `json:"Response"`
}

func (r *ModifyDirectConnectTunnelExtraResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectTunnelExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NQAInfo struct {
	// 健康检查次数
	ProbeFailedTimes *int64 `json:"ProbeFailedTimes,omitempty" name:"ProbeFailedTimes"`

	// 健康检查间隔
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 健康检查地址
	DestinationIp *string `json:"DestinationIp,omitempty" name:"DestinationIp"`
}

// Predefined struct for user
type RejectDirectConnectTunnelRequestParams struct {
	// 无
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

type RejectDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// 无
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *RejectDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectDirectConnectTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectTunnelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RejectDirectConnectTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectDirectConnectTunnelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RejectDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *RejectDirectConnectTunnelResponseParams `json:"Response"`
}

func (r *RejectDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectDirectConnectTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseInternetAddressRequestParams struct {
	// 公网互联网地址ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ReleaseInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// 公网互联网地址ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ReleaseInternetAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseInternetAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseInternetAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseInternetAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseInternetAddressResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseInternetAddressResponseParams `json:"Response"`
}

func (r *ReleaseInternetAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseInternetAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteFilterPrefix struct {
	// 用户侧网段地址
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}