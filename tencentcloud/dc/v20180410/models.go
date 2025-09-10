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

package v20180410

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AcceptDirectConnectTunnelRequestParams struct {
	// 专用通道ID。可以通过[DescribeDirectConnectTunnels](https://cloud.tencent.com/document/product/216/19819)接口获取。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`
}

type AcceptDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID。可以通过[DescribeDirectConnectTunnels](https://cloud.tencent.com/document/product/216/19819)接口获取。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	AccessPointName *string `json:"AccessPointName,omitnil,omitempty" name:"AccessPointName"`

	// 接入点唯一ID。
	AccessPointId *string `json:"AccessPointId,omitnil,omitempty" name:"AccessPointId"`

	// 接入点的状态。可用，不可用。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 接入点的位置。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 接入点支持的运营商列表。
	LineOperator []*string `json:"LineOperator,omitnil,omitempty" name:"LineOperator"`

	// 接入点管理的大区ID。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 接入点可用的端口类型列表。1000BASE-T代表千兆电口，1000BASE-LX代表千兆单模光口10km，1000BASE-ZX代表千兆单模光口80km,10GBASE-LR代表万兆单模光口10km,10GBASE-ZR代表万兆单模光口80km,10GBASE-LH代表万兆单模光口40km,100GBASE-LR4代表100G单模光口10km。
	AvailablePortType []*string `json:"AvailablePortType,omitnil,omitempty" name:"AvailablePortType"`

	// 接入点经纬度。
	Coordinate *Coordinate `json:"Coordinate,omitnil,omitempty" name:"Coordinate"`

	// 接入点所在城市。
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 接入点地域名称。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 接入点类型。VXLAN/QCPL/QCAR
	AccessPointType *string `json:"AccessPointType,omitnil,omitempty" name:"AccessPointType"`

	// 端口规格信息。
	AvailablePortInfo []*PortSpecification `json:"AvailablePortInfo,omitnil,omitempty" name:"AvailablePortInfo"`

	// 接入点地址。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 是否MACsec
	IsMacSec *bool `json:"IsMacSec,omitnil,omitempty" name:"IsMacSec"`
}

// Predefined struct for user
type ApplyInternetAddressRequestParams struct {
	// CIDR地址掩码长度
	MaskLen *int64 `json:"MaskLen,omitnil,omitempty" name:"MaskLen"`

	// 0:BGP类型地址
	// 1：中国电信
	// 2：中国移动
	// 3：中国联通
	AddrType *int64 `json:"AddrType,omitnil,omitempty" name:"AddrType"`

	// 0：IPv4
	// 1:IPv6
	AddrProto *int64 `json:"AddrProto,omitnil,omitempty" name:"AddrProto"`
}

type ApplyInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// CIDR地址掩码长度
	MaskLen *int64 `json:"MaskLen,omitnil,omitempty" name:"MaskLen"`

	// 0:BGP类型地址
	// 1：中国电信
	// 2：中国移动
	// 3：中国联通
	AddrType *int64 `json:"AddrType,omitnil,omitempty" name:"AddrType"`

	// 0：IPv4
	// 1:IPv6
	AddrProto *int64 `json:"AddrProto,omitnil,omitempty" name:"AddrProto"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProbeFailedTimes *int64 `json:"ProbeFailedTimes,omitnil,omitempty" name:"ProbeFailedTimes"`

	// 健康检查间隔
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`
}

type BGPStatus struct {
	// 腾讯侧主互联IP BGP状态
	TencentAddressBgpState *string `json:"TencentAddressBgpState,omitnil,omitempty" name:"TencentAddressBgpState"`

	// 腾讯侧备互联IP BGP状态
	TencentBackupAddressBgpState *string `json:"TencentBackupAddressBgpState,omitnil,omitempty" name:"TencentBackupAddressBgpState"`
}

type BgpPeer struct {
	// 腾讯侧BGP ASN
	CloudAsn *int64 `json:"CloudAsn,omitnil,omitempty" name:"CloudAsn"`

	// 用户侧BGP ASN
	Asn *int64 `json:"Asn,omitnil,omitempty" name:"Asn"`

	// 用户侧BGP密钥
	AuthKey *string `json:"AuthKey,omitnil,omitempty" name:"AuthKey"`
}

type CloudAttachInfo struct {
	// 敏捷上云实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 敏捷上云名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 合作伙伴的AppId
	IapId *string `json:"IapId,omitnil,omitempty" name:"IapId"`

	// 需要接入敏捷上云的IDC的地址
	IdcAddress *string `json:"IdcAddress,omitnil,omitempty" name:"IdcAddress"`

	// 需要接入敏捷上云的IDC的互联网服务提供商类型
	IdcType *string `json:"IdcType,omitnil,omitempty" name:"IdcType"`

	// 敏捷上云的带宽，单位为MB
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 联系电话
	Telephone *string `json:"Telephone,omitnil,omitempty" name:"Telephone"`

	// 敏捷上云的状态
	// available：就绪状态
	// applying：申请，待审核状态
	// pendingpay：代付款状态
	// building：建设中状态
	// confirming：待确认状态
	// isolate: 隔离状态
	// stoped：终止状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 敏捷上云申请的时间
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 敏捷上云建设完成的时间
	ReadyTime *string `json:"ReadyTime,omitnil,omitempty" name:"ReadyTime"`

	// 敏捷上云过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 备注信息
	Remarks *string `json:"Remarks,omitnil,omitempty" name:"Remarks"`

	// 敏捷上云的地域状态。
	// same-region：同地域
	// cross-region：跨地域
	RegionStatus *string `json:"RegionStatus,omitnil,omitempty" name:"RegionStatus"`

	// 用户的AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户的Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 用户注册名称
	CustomerAuthName *string `json:"CustomerAuthName,omitnil,omitempty" name:"CustomerAuthName"`

	// 物理专线实例ID
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`

	// 敏捷上云是否支持创建高速上云专线网关
	CloudAttachServiceGatewaysSupport *bool `json:"CloudAttachServiceGatewaysSupport,omitnil,omitempty" name:"CloudAttachServiceGatewaysSupport"`

	// 敏捷上云服务是否处于升降配中
	BUpdateBandwidth *bool `json:"BUpdateBandwidth,omitnil,omitempty" name:"BUpdateBandwidth"`

	// 接入地域
	ArRegion *string `json:"ArRegion,omitnil,omitempty" name:"ArRegion"`
}

type Coordinate struct {
	// 纬度
	Lat *float64 `json:"Lat,omitnil,omitempty" name:"Lat"`

	// 经度
	Lng *float64 `json:"Lng,omitnil,omitempty" name:"Lng"`
}

type CreateCasInput struct {
	// 敏捷上云名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 需要接入敏捷上云的IDC的地址
	IdcAddress *string `json:"IdcAddress,omitnil,omitempty" name:"IdcAddress"`

	// 需要接入敏捷上云的IDC的互联网服务提供商类型
	IdcType *string `json:"IdcType,omitnil,omitempty" name:"IdcType"`

	// 敏捷上云的带宽，单位为MB
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 联系电话
	Telephone *string `json:"Telephone,omitnil,omitempty" name:"Telephone"`

	// 备注信息
	Remarks *string `json:"Remarks,omitnil,omitempty" name:"Remarks"`

	// 接入地域
	ArRegion *string `json:"ArRegion,omitnil,omitempty" name:"ArRegion"`
}

// Predefined struct for user
type CreateCloudAttachServiceRequestParams struct {
	// 创建敏捷上云入参
	Data *CreateCasInput `json:"Data,omitnil,omitempty" name:"Data"`
}

type CreateCloudAttachServiceRequest struct {
	*tchttp.BaseRequest
	
	// 创建敏捷上云入参
	Data *CreateCasInput `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *CreateCloudAttachServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudAttachServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudAttachServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudAttachServiceResponseParams struct {
	// 敏捷上云服务详情
	CloudAttach *CloudAttachInfo `json:"CloudAttach,omitnil,omitempty" name:"CloudAttach"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudAttachServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudAttachServiceResponseParams `json:"Response"`
}

func (r *CreateCloudAttachServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudAttachServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectRequestParams struct {
	// 物理专线的名称。
	DirectConnectName *string `json:"DirectConnectName,omitnil,omitempty" name:"DirectConnectName"`

	// 物理专线所在的接入点。
	// 您可以通过调用[DescribeAccessPoints](https://cloud.tencent.com/document/product/216/34827)接口获取接入点ID。
	AccessPointId *string `json:"AccessPointId,omitnil,omitempty" name:"AccessPointId"`

	// 提供接入物理专线的运营商。
	// ChinaTelecom：中国电信； 
	// ChinaMobile：中国移动；
	// ChinaUnicom：中国联通；
	//  In-houseWiring：楼内线；
	// ChinaOther：中国其他；
	//  InternationalOperator：境外其他。
	LineOperator *string `json:"LineOperator,omitnil,omitempty" name:"LineOperator"`

	// 物理专线接入端口类型，取值：
	// 100Base-T：百兆电口；
	// 1000Base-T（默认值）：千兆电口；
	// 1000Base-LX：千兆单模光口（10千米）；
	// 10GBase-T：万兆电口；
	// 10GBase-LR（默认值）：万兆单模光口（10千米）。
	PortType *string `json:"PortType,omitnil,omitempty" name:"PortType"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitnil,omitempty" name:"CircuitCode"`

	// 本地数据中心的地理位置。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 物理专线接入接口带宽，单位为Mbps，默认值为1000，取值范围为 [2, 10240]。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 冗余物理专线的ID。
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitnil,omitempty" name:"RedundantDirectConnectId"`

	// 物理专线调试VLAN。默认开启VLAN，自动分配VLAN。
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联 IP。默认自动分配。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联 IP。默认自动分配。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitnil,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil,omitempty" name:"FaultReportContactNumber"`

	// 报障联系邮箱。
	FaultReportContactEmail *string `json:"FaultReportContactEmail,omitnil,omitempty" name:"FaultReportContactEmail"`

	// 物理专线申请者是否签署了用户使用协议。默认已签署。
	SignLaw *bool `json:"SignLaw,omitnil,omitempty" name:"SignLaw"`

	// 标签键值对
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否MACsec需求
	IsMacSec *bool `json:"IsMacSec,omitnil,omitempty" name:"IsMacSec"`
}

type CreateDirectConnectRequest struct {
	*tchttp.BaseRequest
	
	// 物理专线的名称。
	DirectConnectName *string `json:"DirectConnectName,omitnil,omitempty" name:"DirectConnectName"`

	// 物理专线所在的接入点。
	// 您可以通过调用[DescribeAccessPoints](https://cloud.tencent.com/document/product/216/34827)接口获取接入点ID。
	AccessPointId *string `json:"AccessPointId,omitnil,omitempty" name:"AccessPointId"`

	// 提供接入物理专线的运营商。
	// ChinaTelecom：中国电信； 
	// ChinaMobile：中国移动；
	// ChinaUnicom：中国联通；
	//  In-houseWiring：楼内线；
	// ChinaOther：中国其他；
	//  InternationalOperator：境外其他。
	LineOperator *string `json:"LineOperator,omitnil,omitempty" name:"LineOperator"`

	// 物理专线接入端口类型，取值：
	// 100Base-T：百兆电口；
	// 1000Base-T（默认值）：千兆电口；
	// 1000Base-LX：千兆单模光口（10千米）；
	// 10GBase-T：万兆电口；
	// 10GBase-LR（默认值）：万兆单模光口（10千米）。
	PortType *string `json:"PortType,omitnil,omitempty" name:"PortType"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitnil,omitempty" name:"CircuitCode"`

	// 本地数据中心的地理位置。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 物理专线接入接口带宽，单位为Mbps，默认值为1000，取值范围为 [2, 10240]。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 冗余物理专线的ID。
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitnil,omitempty" name:"RedundantDirectConnectId"`

	// 物理专线调试VLAN。默认开启VLAN，自动分配VLAN。
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联 IP。默认自动分配。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联 IP。默认自动分配。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitnil,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil,omitempty" name:"FaultReportContactNumber"`

	// 报障联系邮箱。
	FaultReportContactEmail *string `json:"FaultReportContactEmail,omitnil,omitempty" name:"FaultReportContactEmail"`

	// 物理专线申请者是否签署了用户使用协议。默认已签署。
	SignLaw *bool `json:"SignLaw,omitnil,omitempty" name:"SignLaw"`

	// 标签键值对
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否MACsec需求
	IsMacSec *bool `json:"IsMacSec,omitnil,omitempty" name:"IsMacSec"`
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
	delete(f, "FaultReportContactEmail")
	delete(f, "SignLaw")
	delete(f, "Tags")
	delete(f, "IsMacSec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectResponseParams struct {
	// 物理专线的ID。
	DirectConnectIdSet []*string `json:"DirectConnectIdSet,omitnil,omitempty" name:"DirectConnectIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 物理专线ID，例如：dc-kd7d06of。
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`

	// 专用通道名称。
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil,omitempty" name:"DirectConnectTunnelName"`

	// 物理专线owner，缺省为当前客户（物理专线 owner）
	// 共享专线时这里需要填写共享专线的开发商账号 ID。
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitnil,omitempty" name:"DirectConnectOwnerAccount"`

	// 网络类型，枚举：VPC、CCN、NAT；默认为VPC。VPC：私有网络；CCN：云联网；NAT：NAT网络）。
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 网络地域。
	NetworkRegion *string `json:"NetworkRegion,omitnil,omitempty" name:"NetworkRegion"`

	// 私有网络统一ID，在NetworkType为VPC时必填，且与专线网关所属的VPCID一致；NetworkType为其它组网类型时可不填，内部会统一处理。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 专线网关ID，例如 dcg-d545ddf。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// 专线带宽，单位：Mbps；默认是物理专线带宽值。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 路由类型，枚举：BGP、STATIC；默认为BGP 。（BGP ：BGP路由；STATIC：静态）。
	RouteType *string `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey。
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil,omitempty" name:"BgpPeer"`

	// 静态路由，用户IDC的网段地址。
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil,omitempty" name:"RouteFilterPrefixes"`

	// vlan，范围：0 ~ 3000。
	// 0：不开启子接口，默认值是非0。
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// TencentAddress，腾讯侧互联 IP。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// CustomerAddress，用户侧互联 IP。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// TencentBackupAddress，腾讯侧备用互联 IP。
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil,omitempty" name:"TencentBackupAddress"`

	// 高速上云服务ID。
	CloudAttachId *string `json:"CloudAttachId,omitnil,omitempty" name:"CloudAttachId"`

	// 是否开启BFD。
	BfdEnable *int64 `json:"BfdEnable,omitnil,omitempty" name:"BfdEnable"`

	// 是否开启NQA。
	NqaEnable *int64 `json:"NqaEnable,omitnil,omitempty" name:"NqaEnable"`

	// BFD配置信息。
	BfdInfo *BFDInfo `json:"BfdInfo,omitnil,omitempty" name:"BfdInfo"`

	// NQA配置信息。
	NqaInfo *NQAInfo `json:"NqaInfo,omitnil,omitempty" name:"NqaInfo"`

	// 标签键值对
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// 物理专线ID，例如：dc-kd7d06of。
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`

	// 专用通道名称。
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil,omitempty" name:"DirectConnectTunnelName"`

	// 物理专线owner，缺省为当前客户（物理专线 owner）
	// 共享专线时这里需要填写共享专线的开发商账号 ID。
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitnil,omitempty" name:"DirectConnectOwnerAccount"`

	// 网络类型，枚举：VPC、CCN、NAT；默认为VPC。VPC：私有网络；CCN：云联网；NAT：NAT网络）。
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 网络地域。
	NetworkRegion *string `json:"NetworkRegion,omitnil,omitempty" name:"NetworkRegion"`

	// 私有网络统一ID，在NetworkType为VPC时必填，且与专线网关所属的VPCID一致；NetworkType为其它组网类型时可不填，内部会统一处理。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 专线网关ID，例如 dcg-d545ddf。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// 专线带宽，单位：Mbps；默认是物理专线带宽值。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 路由类型，枚举：BGP、STATIC；默认为BGP 。（BGP ：BGP路由；STATIC：静态）。
	RouteType *string `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// BgpPeer，用户侧bgp信息，包括Asn和AuthKey。
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil,omitempty" name:"BgpPeer"`

	// 静态路由，用户IDC的网段地址。
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil,omitempty" name:"RouteFilterPrefixes"`

	// vlan，范围：0 ~ 3000。
	// 0：不开启子接口，默认值是非0。
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// TencentAddress，腾讯侧互联 IP。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// CustomerAddress，用户侧互联 IP。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// TencentBackupAddress，腾讯侧备用互联 IP。
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil,omitempty" name:"TencentBackupAddress"`

	// 高速上云服务ID。
	CloudAttachId *string `json:"CloudAttachId,omitnil,omitempty" name:"CloudAttachId"`

	// 是否开启BFD。
	BfdEnable *int64 `json:"BfdEnable,omitnil,omitempty" name:"BfdEnable"`

	// 是否开启NQA。
	NqaEnable *int64 `json:"NqaEnable,omitnil,omitempty" name:"NqaEnable"`

	// BFD配置信息。
	BfdInfo *BFDInfo `json:"BfdInfo,omitnil,omitempty" name:"BfdInfo"`

	// NQA配置信息。
	NqaInfo *NQAInfo `json:"NqaInfo,omitnil,omitempty" name:"NqaInfo"`

	// 标签键值对
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectTunnelResponseParams struct {
	// 专用通道ID。
	DirectConnectTunnelIdSet []*string `json:"DirectConnectTunnelIdSet,omitnil,omitempty" name:"DirectConnectTunnelIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`
}

type DeleteDirectConnectRequest struct {
	*tchttp.BaseRequest
	
	// 物理专线的ID。
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`
}

type DeleteDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 接入点所在的地域。你可以通过调用[DescribeRegions](https://cloud.tencent.com/document/product/1596/77930)接口获取地域ID。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数，支持：access-point-id、isp
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAccessPointsRequest struct {
	*tchttp.BaseRequest
	
	// 接入点所在的地域。你可以通过调用[DescribeRegions](https://cloud.tencent.com/document/product/1596/77930)接口获取地域ID。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数，支持：access-point-id、isp
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessPointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessPointsResponseParams struct {
	// 接入点信息。
	AccessPointSet []*AccessPoint `json:"AccessPointSet,omitnil,omitempty" name:"AccessPointSet"`

	// 接入点总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`
}

type DescribeDirectConnectTunnelExtraRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`
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
	// 专用通道扩展信息。
	DirectConnectTunnelExtra *DirectConnectTunnelExtra `json:"DirectConnectTunnelExtra,omitnil,omitempty" name:"DirectConnectTunnelExtra"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// direct-connect-tunnel-name, 专用通道名称。
	// direct-connect-tunnel-id, 专用通道实例ID，如：dcx-abcdefgh。
	// direct-connect-id, 物理专线实例ID，如：dc-abcdefgh。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 专用通道ID数组。
	DirectConnectTunnelIds []*string `json:"DirectConnectTunnelIds,omitnil,omitempty" name:"DirectConnectTunnelIds"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDirectConnectTunnelsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件:
	// 参数不支持同时指定DirectConnectTunnelIds和Filters。
	// direct-connect-tunnel-name, 专用通道名称。
	// direct-connect-tunnel-id, 专用通道实例ID，如：dcx-abcdefgh。
	// direct-connect-id, 物理专线实例ID，如：dc-abcdefgh。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 专用通道ID数组。
	DirectConnectTunnelIds []*string `json:"DirectConnectTunnelIds,omitnil,omitempty" name:"DirectConnectTunnelIds"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	// 专用通道列表。
	DirectConnectTunnelSet []*DirectConnectTunnel `json:"DirectConnectTunnelSet,omitnil,omitempty" name:"DirectConnectTunnelSet"`

	// 专用通道总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 过滤条件。direct-connect-id：物理专线ID，states：物理专线状态（AVAILABLE-就绪，PENDING-申请中，REJECTED-申请被拒绝，PENDINGPAY-待付款，PAID-付款完成，BUILDING-建设中，STOPED-建设终止，DELETED-删除完成）。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 物理专线 ID数组。
	DirectConnectIds []*string `json:"DirectConnectIds,omitnil,omitempty" name:"DirectConnectIds"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDirectConnectsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。direct-connect-id：物理专线ID，states：物理专线状态（AVAILABLE-就绪，PENDING-申请中，REJECTED-申请被拒绝，PENDINGPAY-待付款，PAID-付款完成，BUILDING-建设中，STOPED-建设终止，DELETED-删除完成）。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 物理专线 ID数组。
	DirectConnectIds []*string `json:"DirectConnectIds,omitnil,omitempty" name:"DirectConnectIds"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	DirectConnectSet []*DirectConnect `json:"DirectConnectSet,omitnil,omitempty" name:"DirectConnectSet"`

	// 符合物理专线列表数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 用户名下物理专线是否都签署了用户协议。
	AllSignLaw *bool `json:"AllSignLaw,omitnil,omitempty" name:"AllSignLaw"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Ipv6PrefixLen *int64 `json:"Ipv6PrefixLen,omitnil,omitempty" name:"Ipv6PrefixLen"`

	// BGP类型IPv4互联网地址配额
	Ipv4BgpQuota *int64 `json:"Ipv4BgpQuota,omitnil,omitempty" name:"Ipv4BgpQuota"`

	// 非BGP类型IPv4互联网地址配额
	Ipv4OtherQuota *int64 `json:"Ipv4OtherQuota,omitnil,omitempty" name:"Ipv4OtherQuota"`

	// BGP类型IPv4互联网地址已使用数量
	Ipv4BgpNum *int64 `json:"Ipv4BgpNum,omitnil,omitempty" name:"Ipv4BgpNum"`

	// 非BGP类型互联网地址已使用数量
	Ipv4OtherNum *int64 `json:"Ipv4OtherNum,omitnil,omitempty" name:"Ipv4OtherNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件： <li>AddrType，地址类型。0：BGP 1；1: 电信；2：移动；3：联通</li> <li>AddrProto，地址类型。0：IPv4；1:IPv6</li> <li>Status，地址状态。 0：使用中；1：已停用； 2：已退还</li> <li>Subnet，互联网公网地址。数组</li> <li>InstanceIds，互联网公网地址ID。数组</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件： <li>AddrType，地址类型。0：BGP 1；1: 电信；2：移动；3：联通</li> <li>AddrProto，地址类型。0：IPv4；1:IPv6</li> <li>Status，地址状态。 0：使用中；1：已停用； 2：已退还</li> <li>Subnet，互联网公网地址。数组</li> <li>InstanceIds，互联网公网地址ID。数组</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 互联网公网地址列表
	Subnets []*InternetAddressDetail `json:"Subnets,omitnil,omitempty" name:"Subnets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 互联网公网地址统计信息列表
	InternetAddressStatistics []*InternetAddressStatistics `json:"InternetAddressStatistics,omitnil,omitempty" name:"InternetAddressStatistics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`

	// 过滤条件：
	// route-type：路由类型，取值：BGP/STATIC；
	// route-subnet：路由cidr，取值如：192.68.1.0/24。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePublicDirectConnectTunnelRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`

	// 过滤条件：
	// route-type：路由类型，取值：BGP/STATIC；
	// route-subnet：路由cidr，取值如：192.68.1.0/24。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	// 互联网通道路由列表。
	Routes []*DirectConnectTunnelRoute `json:"Routes,omitnil,omitempty" name:"Routes"`

	// 路由总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`

	// 物理专线的名称。
	DirectConnectName *string `json:"DirectConnectName,omitnil,omitempty" name:"DirectConnectName"`

	// 物理专线的接入点ID。
	AccessPointId *string `json:"AccessPointId,omitnil,omitempty" name:"AccessPointId"`

	// 物理专线的状态。
	// 申请中：PENDING 
	// 申请驳回：REJECTED   
	// 待付款：TOPAY 
	// 已付款：PAID 
	// 建设中：ALLOCATED   
	// 已开通：AVAILABLE  
	// 删除中 ：DELETING
	// 已删除：DELETED 。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 物理专线创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 物理专线的开通时间。
	EnabledTime *string `json:"EnabledTime,omitnil,omitempty" name:"EnabledTime"`

	// 提供接入物理专线的运营商。ChinaTelecom：中国电信， ChinaMobile：中国移动，ChinaUnicom：中国联通， In-houseWiring：楼内线，ChinaOther：中国其他， InternationalOperator：境外其他。
	LineOperator *string `json:"LineOperator,omitnil,omitempty" name:"LineOperator"`

	// 本地数据中心的地理位置。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 物理专线接入接口带宽，单位为Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 用户侧物理专线接入端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）
	PortType *string `json:"PortType,omitnil,omitempty" name:"PortType"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitnil,omitempty" name:"CircuitCode"`

	// 冗余物理专线的ID。
	RedundantDirectConnectId *string `json:"RedundantDirectConnectId,omitnil,omitempty" name:"RedundantDirectConnectId"`

	// 物理专线调试VLAN。默认开启VLAN，自动分配VLAN。
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联IP。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联IP。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitnil,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil,omitempty" name:"CustomerContactNumber"`

	// 物理专线的过期时间。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 物理专线计费类型。 NON_RECURRING_CHARGE：一次性接入费用；PREPAID_BY_YEAR：按年预付费。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil,omitempty" name:"FaultReportContactNumber"`

	// 标签键值对
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 物理专线的接入点类型。
	AccessPointType *string `json:"AccessPointType,omitnil,omitempty" name:"AccessPointType"`

	// IDC所在城市
	IdcCity *string `json:"IdcCity,omitnil,omitempty" name:"IdcCity"`

	// 计费状态
	ChargeState *string `json:"ChargeState,omitnil,omitempty" name:"ChargeState"`

	// 物理专线开通时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 物理专线是否已签署用户协议
	SignLaw *bool `json:"SignLaw,omitnil,omitempty" name:"SignLaw"`

	// 物理专线是否为LocalZone
	LocalZone *bool `json:"LocalZone,omitnil,omitempty" name:"LocalZone"`

	// 该物理专线下vlan 0的专用通道数量
	VlanZeroDirectConnectTunnelCount *uint64 `json:"VlanZeroDirectConnectTunnelCount,omitnil,omitempty" name:"VlanZeroDirectConnectTunnelCount"`

	// 该物理专线下非vlan 0的专用通道数量
	OtherVlanDirectConnectTunnelCount *uint64 `json:"OtherVlanDirectConnectTunnelCount,omitnil,omitempty" name:"OtherVlanDirectConnectTunnelCount"`

	// 物理专线最小带宽
	MinBandwidth *uint64 `json:"MinBandwidth,omitnil,omitempty" name:"MinBandwidth"`

	// 建设模式
	Construct *uint64 `json:"Construct,omitnil,omitempty" name:"Construct"`

	// 物理专线的接入点名称
	AccessPointName *string `json:"AccessPointName,omitnil,omitempty" name:"AccessPointName"`

	// 是否三层架构
	IsThreeArch *bool `json:"IsThreeArch,omitnil,omitempty" name:"IsThreeArch"`

	// 是否MACsec
	IsMacSec *bool `json:"IsMacSec,omitnil,omitempty" name:"IsMacSec"`

	// 端口规格(Mbps)
	PortSpecification *uint64 `json:"PortSpecification,omitnil,omitempty" name:"PortSpecification"`
}

type DirectConnectTunnel struct {
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`

	// 物理专线ID
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`

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
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 物理专线的拥有者，开发商账号 ID
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitnil,omitempty" name:"DirectConnectOwnerAccount"`

	// 专用通道的拥有者，开发商账号 ID
	OwnerAccount *string `json:"OwnerAccount,omitnil,omitempty" name:"OwnerAccount"`

	// 网络类型，分别为VPC、BMVPC、CCN
	//  VPC：私有网络 ，BMVPC：黑石网络，CCN：云联网
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// VPC地域对应的网络名，如ap-guangzhou
	NetworkRegion *string `json:"NetworkRegion,omitnil,omitempty" name:"NetworkRegion"`

	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 专线网关 ID
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// BGP ：BGP路由 STATIC：静态 默认为 BGP 路由
	RouteType *string `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// 用户侧BGP，包括： CloudAsn，Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil,omitempty" name:"BgpPeer"`

	// 用户侧网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil,omitempty" name:"RouteFilterPrefixes"`

	// 专用通道的Vlan
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// TencentAddress，腾讯侧互联 IP
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// CustomerAddress，用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil,omitempty" name:"DirectConnectTunnelName"`

	// 专用通道创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 专用通道带宽值
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 专用通道标签值
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 关联的网络自定义探测ID
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`

	// BGP community开关
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitnil,omitempty" name:"EnableBGPCommunity"`

	// 是否为Nat通道
	NatType *int64 `json:"NatType,omitnil,omitempty" name:"NatType"`

	// VPC地域简码，如gz、cd
	VpcRegion *string `json:"VpcRegion,omitnil,omitempty" name:"VpcRegion"`

	// 是否开启BFD
	BfdEnable *int64 `json:"BfdEnable,omitnil,omitempty" name:"BfdEnable"`

	// 专用通道接入点类型
	AccessPointType *string `json:"AccessPointType,omitnil,omitempty" name:"AccessPointType"`

	// 专线网关名称
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil,omitempty" name:"DirectConnectGatewayName"`

	// VPC名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// TencentBackupAddress，腾讯侧备用互联 IP
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil,omitempty" name:"TencentBackupAddress"`

	// 专用通道关联的物理专线是否签署了用户协议
	SignLaw *bool `json:"SignLaw,omitnil,omitempty" name:"SignLaw"`

	// 高速上云服务ID
	CloudAttachId *string `json:"CloudAttachId,omitnil,omitempty" name:"CloudAttachId"`

	// 是否共享通道
	ShareOrNot *uint64 `json:"ShareOrNot,omitnil,omitempty" name:"ShareOrNot"`

	// 接入点名称
	AccessPointName *string `json:"AccessPointName,omitnil,omitempty" name:"AccessPointName"`

	// 接入点ID
	AccessPointId *string `json:"AccessPointId,omitnil,omitempty" name:"AccessPointId"`
}

type DirectConnectTunnelExtra struct {
	// 专用通道ID
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`

	// 物理专线ID
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`

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
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 物理专线的拥有者，开发商账号 ID
	DirectConnectOwnerAccount *string `json:"DirectConnectOwnerAccount,omitnil,omitempty" name:"DirectConnectOwnerAccount"`

	// 专用通道的拥有者，开发商账号 ID
	OwnerAccount *string `json:"OwnerAccount,omitnil,omitempty" name:"OwnerAccount"`

	// 网络类型，分别为VPC、BMVPC、CCN
	//  VPC：私有网络 ，BMVPC：黑石网络，CCN：云联网
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// VPC地域对应的网络名，如ap-guangzhou
	NetworkRegion *string `json:"NetworkRegion,omitnil,omitempty" name:"NetworkRegion"`

	// 私有网络统一 ID 或者黑石网络统一 ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 专线网关 ID
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil,omitempty" name:"DirectConnectGatewayId"`

	// BGP ：BGP路由 STATIC：静态 默认为 BGP 路由
	RouteType *string `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// 用户侧BGP，Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil,omitempty" name:"BgpPeer"`

	// 用户侧网段地址
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil,omitempty" name:"RouteFilterPrefixes"`

	// 互联网通道公网网段地址
	PublicAddresses []*RouteFilterPrefix `json:"PublicAddresses,omitnil,omitempty" name:"PublicAddresses"`

	// 专用通道的Vlan
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// 腾讯侧互联 IP
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 腾讯侧备用互联IP
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil,omitempty" name:"TencentBackupAddress"`

	// 用户侧互联 IP
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 专用通道名称
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil,omitempty" name:"DirectConnectTunnelName"`

	// 专用通道创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 专用通道带宽值
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 关联的网络自定义探测ID
	NetDetectId *string `json:"NetDetectId,omitnil,omitempty" name:"NetDetectId"`

	// BGP community开关
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitnil,omitempty" name:"EnableBGPCommunity"`

	// 是否为Nat通道
	NatType *int64 `json:"NatType,omitnil,omitempty" name:"NatType"`

	// VPC地域简码，如gz、cd
	VpcRegion *string `json:"VpcRegion,omitnil,omitempty" name:"VpcRegion"`

	// 是否开启BFD
	BfdEnable *int64 `json:"BfdEnable,omitnil,omitempty" name:"BfdEnable"`

	// 是否开启NQA
	NqaEnable *int64 `json:"NqaEnable,omitnil,omitempty" name:"NqaEnable"`

	// 专用通道接入点类型
	AccessPointType *string `json:"AccessPointType,omitnil,omitempty" name:"AccessPointType"`

	// 专线网关名称
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil,omitempty" name:"DirectConnectGatewayName"`

	// VPC名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 专用通道关联的物理专线是否签署了用户协议
	SignLaw *bool `json:"SignLaw,omitnil,omitempty" name:"SignLaw"`

	// BFD配置信息
	BfdInfo *BFDInfo `json:"BfdInfo,omitnil,omitempty" name:"BfdInfo"`

	// NQA配置信息
	NqaInfo *NQAInfo `json:"NqaInfo,omitnil,omitempty" name:"NqaInfo"`

	// BGP状态
	BgpStatus *BGPStatus `json:"BgpStatus,omitnil,omitempty" name:"BgpStatus"`

	// 是否开启IPv6
	IPv6Enable *int64 `json:"IPv6Enable,omitnil,omitempty" name:"IPv6Enable"`

	// 腾讯侧互联IPv6地址
	TencentIPv6Address *string `json:"TencentIPv6Address,omitnil,omitempty" name:"TencentIPv6Address"`

	// 腾讯侧备用互联IPv6地址
	TencentBackupIPv6Address *string `json:"TencentBackupIPv6Address,omitnil,omitempty" name:"TencentBackupIPv6Address"`

	// BGPv6状态
	BgpIPv6Status *BGPStatus `json:"BgpIPv6Status,omitnil,omitempty" name:"BgpIPv6Status"`

	// 用户侧互联IPv6地址
	CustomerIPv6Address *string `json:"CustomerIPv6Address,omitnil,omitempty" name:"CustomerIPv6Address"`

	// 专用通道是否支持巨帧。1 支持，0 不支持
	JumboEnable *int64 `json:"JumboEnable,omitnil,omitempty" name:"JumboEnable"`

	// 专用通道是否支持高精度BFD。1支持，0不支持
	HighPrecisionBFDEnable *int64 `json:"HighPrecisionBFDEnable,omitnil,omitempty" name:"HighPrecisionBFDEnable"`
}

type DirectConnectTunnelRoute struct {
	// 专用通道路由ID
	RouteId *string `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// 网段CIDR
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`

	// 路由类型：BGP/STATIC路由
	RouteType *string `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// ENABLE：路由启用，DISABLE：路由禁用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// ASPath信息
	ASPath []*string `json:"ASPath,omitnil,omitempty" name:"ASPath"`

	// 路由下一跳IP
	NextHop *string `json:"NextHop,omitnil,omitempty" name:"NextHop"`

	// 路由更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否配置在通道上
	ApplyOnTunnelEnable *bool `json:"ApplyOnTunnelEnable,omitnil,omitempty" name:"ApplyOnTunnelEnable"`
}

// Predefined struct for user
type DisableInternetAddressRequestParams struct {
	// 公网互联网地址ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DisableInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// 公网互联网地址ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type EnableInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// 互联网公网地址ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type InternetAddressDetail struct {
	// 互联网地址ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 互联网网络地址
	Subnet *string `json:"Subnet,omitnil,omitempty" name:"Subnet"`

	// 网络地址掩码长度
	MaskLen *int64 `json:"MaskLen,omitnil,omitempty" name:"MaskLen"`

	// 0:BGP
	// 1:电信
	// 2:移动
	// 3:联通
	AddrType *int64 `json:"AddrType,omitnil,omitempty" name:"AddrType"`

	// 0:使用中
	// 1:已停用
	// 2:已退还
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 申请时间
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 停用时间
	StopTime *string `json:"StopTime,omitnil,omitempty" name:"StopTime"`

	// 退还时间
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 用户ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 0:IPv4 1:IPv6
	AddrProto *int64 `json:"AddrProto,omitnil,omitempty" name:"AddrProto"`

	// 释放状态的IP地址保留的天数
	ReserveTime *int64 `json:"ReserveTime,omitnil,omitempty" name:"ReserveTime"`
}

type InternetAddressStatistics struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 互联网公网地址数量
	SubnetNum *int64 `json:"SubnetNum,omitnil,omitempty" name:"SubnetNum"`
}

// Predefined struct for user
type ModifyDirectConnectAttributeRequestParams struct {
	// 物理专线ID。
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`

	// 物理专线名称。
	DirectConnectName *string `json:"DirectConnectName,omitnil,omitempty" name:"DirectConnectName"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitnil,omitempty" name:"CircuitCode"`

	// 物理专线调试VLAN。
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联 IP。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联 IP。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitnil,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil,omitempty" name:"FaultReportContactNumber"`

	// 报障联系邮箱。
	FaultReportContactEmail *string `json:"FaultReportContactEmail,omitnil,omitempty" name:"FaultReportContactEmail"`

	// 物理专线申请者补签用户使用协议。
	SignLaw *bool `json:"SignLaw,omitnil,omitempty" name:"SignLaw"`

	// 物理专线带宽。
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`
}

type ModifyDirectConnectAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 物理专线ID。
	DirectConnectId *string `json:"DirectConnectId,omitnil,omitempty" name:"DirectConnectId"`

	// 物理专线名称。
	DirectConnectName *string `json:"DirectConnectName,omitnil,omitempty" name:"DirectConnectName"`

	// 运营商或者服务商为物理专线提供的电路编码。
	CircuitCode *string `json:"CircuitCode,omitnil,omitempty" name:"CircuitCode"`

	// 物理专线调试VLAN。
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// 物理专线调试腾讯侧互联 IP。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 物理专线调试用户侧互联 IP。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 物理专线申请者姓名。默认从账户体系获取。
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 物理专线申请者联系邮箱。默认从账户体系获取。
	CustomerContactMail *string `json:"CustomerContactMail,omitnil,omitempty" name:"CustomerContactMail"`

	// 物理专线申请者联系号码。默认从账户体系获取。
	CustomerContactNumber *string `json:"CustomerContactNumber,omitnil,omitempty" name:"CustomerContactNumber"`

	// 报障联系人。
	FaultReportContactPerson *string `json:"FaultReportContactPerson,omitnil,omitempty" name:"FaultReportContactPerson"`

	// 报障联系电话。
	FaultReportContactNumber *string `json:"FaultReportContactNumber,omitnil,omitempty" name:"FaultReportContactNumber"`

	// 报障联系邮箱。
	FaultReportContactEmail *string `json:"FaultReportContactEmail,omitnil,omitempty" name:"FaultReportContactEmail"`

	// 物理专线申请者补签用户使用协议。
	SignLaw *bool `json:"SignLaw,omitnil,omitempty" name:"SignLaw"`

	// 物理专线带宽。
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`
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
	delete(f, "FaultReportContactEmail")
	delete(f, "SignLaw")
	delete(f, "Bandwidth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDirectConnectAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道名称。
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil,omitempty" name:"DirectConnectTunnelName"`

	// 用户侧BGP，包括Asn，AuthKey。
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil,omitempty" name:"BgpPeer"`

	// 用户侧网段地址。
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil,omitempty" name:"RouteFilterPrefixes"`

	// 腾讯侧互联IP。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 用户侧互联IP。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 专用通道带宽值，单位为M。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 腾讯侧备用互联IP。
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil,omitempty" name:"TencentBackupAddress"`
}

type ModifyDirectConnectTunnelAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道名称。
	DirectConnectTunnelName *string `json:"DirectConnectTunnelName,omitnil,omitempty" name:"DirectConnectTunnelName"`

	// 用户侧BGP，包括Asn，AuthKey。
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil,omitempty" name:"BgpPeer"`

	// 用户侧网段地址。
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil,omitempty" name:"RouteFilterPrefixes"`

	// 腾讯侧互联IP。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 用户侧互联IP。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 专用通道带宽值，单位为M。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 腾讯侧备用互联IP。
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil,omitempty" name:"TencentBackupAddress"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道的Vlan。
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// Bgp参数，包括Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil,omitempty" name:"BgpPeer"`

	// 用户侧过滤网段地址。
	RouteFilterPrefixes *RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil,omitempty" name:"RouteFilterPrefixes"`

	// 腾讯侧互联IP。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 腾讯侧备用互联IP。
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil,omitempty" name:"TencentBackupAddress"`

	// 用户侧互联IP。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 专用通道带宽值。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// BGP community开关。
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitnil,omitempty" name:"EnableBGPCommunity"`

	// 是否开启BFD。
	BfdEnable *int64 `json:"BfdEnable,omitnil,omitempty" name:"BfdEnable"`

	// 是否开启NQA。
	NqaEnable *int64 `json:"NqaEnable,omitnil,omitempty" name:"NqaEnable"`

	// BFD配置信息。
	BfdInfo *BFDInfo `json:"BfdInfo,omitnil,omitempty" name:"BfdInfo"`

	// NQA配置信息。
	NqaInfo *NQAInfo `json:"NqaInfo,omitnil,omitempty" name:"NqaInfo"`

	// IPV6使能。0：停用IPv6；1: 启用IPv6。
	IPv6Enable *int64 `json:"IPv6Enable,omitnil,omitempty" name:"IPv6Enable"`

	// 去往用户侧的路由信息。
	CustomerIDCRoutes []*RouteFilterPrefix `json:"CustomerIDCRoutes,omitnil,omitempty" name:"CustomerIDCRoutes"`

	// 是否开启巨帧。1：开启；0：不开启。
	JumboEnable *int64 `json:"JumboEnable,omitnil,omitempty" name:"JumboEnable"`

	// 腾讯侧互联IPv6。
	TencentIPv6Address *string `json:"TencentIPv6Address,omitnil,omitempty" name:"TencentIPv6Address"`

	// 腾讯侧备用互联IPv6。
	TencentBackupIPv6Address *string `json:"TencentBackupIPv6Address,omitnil,omitempty" name:"TencentBackupIPv6Address"`

	// 用户侧互联IPv6。
	CustomerIPv6Address *string `json:"CustomerIPv6Address,omitnil,omitempty" name:"CustomerIPv6Address"`
}

type ModifyDirectConnectTunnelExtraRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`

	// 专用通道的Vlan。
	Vlan *int64 `json:"Vlan,omitnil,omitempty" name:"Vlan"`

	// Bgp参数，包括Asn，AuthKey
	BgpPeer *BgpPeer `json:"BgpPeer,omitnil,omitempty" name:"BgpPeer"`

	// 用户侧过滤网段地址。
	RouteFilterPrefixes *RouteFilterPrefix `json:"RouteFilterPrefixes,omitnil,omitempty" name:"RouteFilterPrefixes"`

	// 腾讯侧互联IP。
	TencentAddress *string `json:"TencentAddress,omitnil,omitempty" name:"TencentAddress"`

	// 腾讯侧备用互联IP。
	TencentBackupAddress *string `json:"TencentBackupAddress,omitnil,omitempty" name:"TencentBackupAddress"`

	// 用户侧互联IP。
	CustomerAddress *string `json:"CustomerAddress,omitnil,omitempty" name:"CustomerAddress"`

	// 专用通道带宽值。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// BGP community开关。
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitnil,omitempty" name:"EnableBGPCommunity"`

	// 是否开启BFD。
	BfdEnable *int64 `json:"BfdEnable,omitnil,omitempty" name:"BfdEnable"`

	// 是否开启NQA。
	NqaEnable *int64 `json:"NqaEnable,omitnil,omitempty" name:"NqaEnable"`

	// BFD配置信息。
	BfdInfo *BFDInfo `json:"BfdInfo,omitnil,omitempty" name:"BfdInfo"`

	// NQA配置信息。
	NqaInfo *NQAInfo `json:"NqaInfo,omitnil,omitempty" name:"NqaInfo"`

	// IPV6使能。0：停用IPv6；1: 启用IPv6。
	IPv6Enable *int64 `json:"IPv6Enable,omitnil,omitempty" name:"IPv6Enable"`

	// 去往用户侧的路由信息。
	CustomerIDCRoutes []*RouteFilterPrefix `json:"CustomerIDCRoutes,omitnil,omitempty" name:"CustomerIDCRoutes"`

	// 是否开启巨帧。1：开启；0：不开启。
	JumboEnable *int64 `json:"JumboEnable,omitnil,omitempty" name:"JumboEnable"`

	// 腾讯侧互联IPv6。
	TencentIPv6Address *string `json:"TencentIPv6Address,omitnil,omitempty" name:"TencentIPv6Address"`

	// 腾讯侧备用互联IPv6。
	TencentBackupIPv6Address *string `json:"TencentBackupIPv6Address,omitnil,omitempty" name:"TencentBackupIPv6Address"`

	// 用户侧互联IPv6。
	CustomerIPv6Address *string `json:"CustomerIPv6Address,omitnil,omitempty" name:"CustomerIPv6Address"`
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
	delete(f, "TencentIPv6Address")
	delete(f, "TencentBackupIPv6Address")
	delete(f, "CustomerIPv6Address")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDirectConnectTunnelExtraRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectTunnelExtraResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProbeFailedTimes *int64 `json:"ProbeFailedTimes,omitnil,omitempty" name:"ProbeFailedTimes"`

	// 健康检查间隔
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 健康检查地址
	DestinationIp *string `json:"DestinationIp,omitnil,omitempty" name:"DestinationIp"`
}

type PortSpecification struct {
	// 端口名称
	InternationalName *string `json:"InternationalName,omitnil,omitempty" name:"InternationalName"`

	// 端口规格（M）
	Specification *uint64 `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 端口类型：T-电口，X-光口
	PortType *string `json:"PortType,omitnil,omitempty" name:"PortType"`
}

// Predefined struct for user
type RejectDirectConnectTunnelRequestParams struct {
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`
}

type RejectDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest
	
	// 专用通道ID。
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitnil,omitempty" name:"DirectConnectTunnelId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ReleaseInternetAddressRequest struct {
	*tchttp.BaseRequest
	
	// 公网互联网地址ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}