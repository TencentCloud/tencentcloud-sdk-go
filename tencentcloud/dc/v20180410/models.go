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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AcceptDirectConnectTunnelRequest struct {
	*tchttp.BaseRequest

	// 物理专线拥有者接受共享专用通道申请
	DirectConnectTunnelId *string `json:"DirectConnectTunnelId,omitempty" name:"DirectConnectTunnelId"`
}

func (r *AcceptDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AcceptDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AcceptDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcceptDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	LineOperator []*string `json:"LineOperator,omitempty" name:"LineOperator" list`

	// 接入点管理的大区ID。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 接入点可用的端口类型列表。1000BASE-T代表千兆电口，1000BASE-LX代表千兆单模光口10km，1000BASE-ZX代表千兆单模光口80km,10GBASE-LR代表万兆单模光口10km,10GBASE-ZR代表万兆单模光口80km,10GBASE-LH代表万兆单模光口40km,100GBASE-LR4代表100G单模光口10km
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailablePortType []*string `json:"AvailablePortType,omitempty" name:"AvailablePortType" list`
}

type BgpPeer struct {

	// 用户侧，BGP Asn
	Asn *int64 `json:"Asn,omitempty" name:"Asn"`

	// 用户侧BGP密钥
	AuthKey *string `json:"AuthKey,omitempty" name:"AuthKey"`
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

func (r *CreateDirectConnectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 物理专线的ID。
		DirectConnectIdSet []*string `json:"DirectConnectIdSet,omitempty" name:"DirectConnectIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes" list`

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
}

func (r *CreateDirectConnectTunnelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专用通道ID
		DirectConnectTunnelIdSet []*string `json:"DirectConnectTunnelIdSet,omitempty" name:"DirectConnectTunnelIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteDirectConnectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDirectConnectTunnelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeAccessPointsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessPointsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接入点信息。
		AccessPointSet []*AccessPoint `json:"AccessPointSet,omitempty" name:"AccessPointSet" list`

		// 符合接入点数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccessPointsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectTunnelsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件:
	// 参数不支持同时指定DirectConnectTunnelIds和Filters。
	// <li> direct-connect-tunnel-name, 专用通道名称。</li>
	// <li> direct-connect-tunnel-id, 专用通道实例ID，如dcx-abcdefgh。</li>
	// <li>direct-connect-id, 物理专线实例ID，如，dc-abcdefgh。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 专用通道 ID数组
	DirectConnectTunnelIds []*string `json:"DirectConnectTunnelIds,omitempty" name:"DirectConnectTunnelIds" list`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectTunnelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectTunnelsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectTunnelsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专用通道列表
		DirectConnectTunnelSet []*DirectConnectTunnel `json:"DirectConnectTunnelSet,omitempty" name:"DirectConnectTunnelSet" list`

		// 符合专用通道数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectTunnelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectTunnelsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件:
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 物理专线 ID数组
	DirectConnectIds []*string `json:"DirectConnectIds,omitempty" name:"DirectConnectIds" list`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDirectConnectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 物理专线列表。
		DirectConnectSet []*DirectConnect `json:"DirectConnectSet,omitempty" name:"DirectConnectSet" list`

		// 符合物理专线列表数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 用户名下物理专线是否都签署了用户协议
	// 注意：此字段可能返回 null，表示取不到有效值。
		AllSignLaw *bool `json:"AllSignLaw,omitempty" name:"AllSignLaw"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDirectConnectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDirectConnectsResponse) FromJsonString(s string) error {
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
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

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
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes" list`

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
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet" list`

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
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
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
}

func (r *ModifyDirectConnectAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	RouteFilterPrefixes []*RouteFilterPrefix `json:"RouteFilterPrefixes,omitempty" name:"RouteFilterPrefixes" list`

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

func (r *ModifyDirectConnectTunnelAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectTunnelAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDirectConnectTunnelAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDirectConnectTunnelAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *RejectDirectConnectTunnelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RejectDirectConnectTunnelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RejectDirectConnectTunnelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RejectDirectConnectTunnelResponse) FromJsonString(s string) error {
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
