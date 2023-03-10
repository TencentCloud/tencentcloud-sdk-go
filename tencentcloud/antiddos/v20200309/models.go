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

package v20200309

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AclConfig struct {
	// 协议类型, 可取值tcp, udp, all
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// 目的端口起始，可取值范围0~65535
	DPortStart *uint64 `json:"DPortStart,omitempty" name:"DPortStart"`

	// 目的端口结束，可取值范围0~65535
	DPortEnd *uint64 `json:"DPortEnd,omitempty" name:"DPortEnd"`

	// 来源端口起始，可取值范围0~65535
	SPortStart *uint64 `json:"SPortStart,omitempty" name:"SPortStart"`

	// 来源端口结束，可取值范围0~65535
	SPortEnd *uint64 `json:"SPortEnd,omitempty" name:"SPortEnd"`

	// 动作，可取值：drop， transmit， forward
	Action *string `json:"Action,omitempty" name:"Action"`

	// 策略优先级，数字越小，级别越高，该规则越靠前匹配，取值1-1000
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

type AclConfigRelation struct {
	// acl策略
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`

	// 实例列表
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type AnycastOutPackRelation struct {
	// 业务带宽(单位M)
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalBandwidth *uint64 `json:"NormalBandwidth,omitempty" name:"NormalBandwidth"`

	// 转发规则数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardRulesLimit *uint64 `json:"ForwardRulesLimit,omitempty" name:"ForwardRulesLimit"`

	// 自动续费标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
}

// Predefined struct for user
type AssociateDDoSEipAddressRequestParams struct {
	// 资源实例ID，实例ID形如：bgpip-0000011x。只能填写高防IP实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源实例ID对应的高防弹性公网IP。
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 要绑定的实例 ID。实例 ID 形如：ins-11112222。可通过登录控制台查询，也可通过 DescribeInstances 接口返回值中的InstanceId获取。
	CvmInstanceID *string `json:"CvmInstanceID,omitempty" name:"CvmInstanceID"`

	// cvm实例所在地域，例如：ap-hongkong。
	CvmRegion *string `json:"CvmRegion,omitempty" name:"CvmRegion"`
}

type AssociateDDoSEipAddressRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID，实例ID形如：bgpip-0000011x。只能填写高防IP实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源实例ID对应的高防弹性公网IP。
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 要绑定的实例 ID。实例 ID 形如：ins-11112222。可通过登录控制台查询，也可通过 DescribeInstances 接口返回值中的InstanceId获取。
	CvmInstanceID *string `json:"CvmInstanceID,omitempty" name:"CvmInstanceID"`

	// cvm实例所在地域，例如：ap-hongkong。
	CvmRegion *string `json:"CvmRegion,omitempty" name:"CvmRegion"`
}

func (r *AssociateDDoSEipAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Eip")
	delete(f, "CvmInstanceID")
	delete(f, "CvmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDDoSEipAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDDoSEipAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateDDoSEipAddressResponse struct {
	*tchttp.BaseResponse
	Response *AssociateDDoSEipAddressResponseParams `json:"Response"`
}

func (r *AssociateDDoSEipAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDDoSEipLoadBalancerRequestParams struct {
	// 资源实例ID，实例ID形如：bgpip-0000011x。只能填写高防IP实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源实例ID对应的高防弹性公网IP。
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 要绑定的负载均衡ID。负载均衡 ID 形如：lb-0000002i。可通过登录控制台查询，也可通过 DescribeLoadBalancers 接口返回值中的LoadBalancerId获取。
	LoadBalancerID *string `json:"LoadBalancerID,omitempty" name:"LoadBalancerID"`

	// CLB所在地域，例如：ap-hongkong。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`

	// CLB内网IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type AssociateDDoSEipLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID，实例ID形如：bgpip-0000011x。只能填写高防IP实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源实例ID对应的高防弹性公网IP。
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 要绑定的负载均衡ID。负载均衡 ID 形如：lb-0000002i。可通过登录控制台查询，也可通过 DescribeLoadBalancers 接口返回值中的LoadBalancerId获取。
	LoadBalancerID *string `json:"LoadBalancerID,omitempty" name:"LoadBalancerID"`

	// CLB所在地域，例如：ap-hongkong。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`

	// CLB内网IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *AssociateDDoSEipLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Eip")
	delete(f, "LoadBalancerID")
	delete(f, "LoadBalancerRegion")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDDoSEipLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDDoSEipLoadBalancerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateDDoSEipLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *AssociateDDoSEipLoadBalancerResponseParams `json:"Response"`
}

func (r *AssociateDDoSEipLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BGPIPInstance struct {
	// 资产实例的详细信息
	InstanceDetail *InstanceRelation `json:"InstanceDetail,omitempty" name:"InstanceDetail"`

	// 资产实例的规格信息
	SpecificationLimit *BGPIPInstanceSpecification `json:"SpecificationLimit,omitempty" name:"SpecificationLimit"`

	// 资产实例的使用统计信息
	Usage *BGPIPInstanceUsages `json:"Usage,omitempty" name:"Usage"`

	// 资产实例所在的地域
	Region *RegionInfo `json:"Region,omitempty" name:"Region"`

	// 资产实例的防护状态，状态码如下：
	// "idle"：正常状态(无攻击)
	// "attacking"：攻击中
	// "blocking"：封堵中
	// "creating"：创建中
	// "deblocking"：解封中
	// "isolate"：回收隔离中
	Status *string `json:"Status,omitempty" name:"Status"`

	// 购买时间
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 到期时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 资产实例的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资产实例所属的套餐包信息，
	// 注意：当资产实例不是套餐包的实例时，此字段为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackInfo *PackInfo `json:"PackInfo,omitempty" name:"PackInfo"`

	// 资产实例所属的三网套餐包详情，
	// 注意：当资产实例不是三网套餐包的实例时，此字段为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaticPackRelation *StaticPackRelation `json:"StaticPackRelation,omitempty" name:"StaticPackRelation"`

	// 区分高防IP境外线路
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 区分集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tgw *uint64 `json:"Tgw,omitempty" name:"Tgw"`

	// 高防弹性公网IP状态，包含'CREATING'(创建中),'BINDING'(绑定中),'BIND'(已绑定),'UNBINDING'(解绑中),'UNBIND'(已解绑),'OFFLINING'(释放中),'BIND_ENI'(绑定悬空弹性网卡)。只对高防弹性公网IP实例有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipAddressStatus *string `json:"EipAddressStatus,omitempty" name:"EipAddressStatus"`

	// 是否高防弹性公网IP实例，是为1，否为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipFlag *int64 `json:"EipFlag,omitempty" name:"EipFlag"`

	// 资产实例所属的高防弹性公网IP套餐包详情，
	// 注意：当资产实例不是高防弹性公网IP套餐包的实例时，此字段为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipAddressPackRelation *EipAddressPackRelation `json:"EipAddressPackRelation,omitempty" name:"EipAddressPackRelation"`

	// 高防弹性公网IP关联的实例信息。
	// 注意：当资产实例不是高防弹性公网IP实例时，此字段为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipAddressInfo *EipAddressRelation `json:"EipAddressInfo,omitempty" name:"EipAddressInfo"`

	// 建议客户接入的域名，客户可使用域名接入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 是否开启安全加速，是为1，否为0。
	DamDDoSStatus *uint64 `json:"DamDDoSStatus,omitempty" name:"DamDDoSStatus"`

	// 是否Ipv6版本的IP, 是为1，否为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	V6Flag *uint64 `json:"V6Flag,omitempty" name:"V6Flag"`

	// 是否渠道版高防IP，是为1，否为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	BGPIPChannelFlag *uint64 `json:"BGPIPChannelFlag,omitempty" name:"BGPIPChannelFlag"`

	// 资源关联标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagInfoList []*TagInfo `json:"TagInfoList,omitempty" name:"TagInfoList"`

	// 资产实例所属的全力防护套餐包详情，
	// 注意：当资产实例不是全力防护套餐包的实例时，此字段为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnycastOutPackRelation *AnycastOutPackRelation `json:"AnycastOutPackRelation,omitempty" name:"AnycastOutPackRelation"`

	// 资源实例版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceVersion *uint64 `json:"InstanceVersion,omitempty" name:"InstanceVersion"`

	// 重保实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConvoyId *string `json:"ConvoyId,omitempty" name:"ConvoyId"`

	// 带宽后付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticBandwidth *uint64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`

	// 是否为EO代播的ip: 1是，0不是
	// 注意：此字段可能返回 null，表示取不到有效值。
	EOFlag *uint64 `json:"EOFlag,omitempty" name:"EOFlag"`
}

type BGPIPInstanceSpecification struct {
	// 保底防护峰值，单位Mbps
	ProtectBandwidth *uint64 `json:"ProtectBandwidth,omitempty" name:"ProtectBandwidth"`

	// CC防护峰值，单位qps
	ProtectCCQPS *uint64 `json:"ProtectCCQPS,omitempty" name:"ProtectCCQPS"`

	// 正常业务带宽，单位Mbps
	NormalBandwidth *uint64 `json:"NormalBandwidth,omitempty" name:"NormalBandwidth"`

	// 转发规则数，单位条
	ForwardRulesLimit *uint64 `json:"ForwardRulesLimit,omitempty" name:"ForwardRulesLimit"`

	// 自动续费状态，取值[
	// 0：没有开启自动续费
	// 1：开启了自动续费
	// ]
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 高防IP线路，取值为[
	// 1：BGP线路
	// 2：电信
	// 3：联通
	// 4：移动
	// 99：第三方合作线路
	// ]
	Line *uint64 `json:"Line,omitempty" name:"Line"`

	// 弹性防护峰值，单位Mbps
	ElasticBandwidth *uint64 `json:"ElasticBandwidth,omitempty" name:"ElasticBandwidth"`
}

type BGPIPInstanceUsages struct {
	// 已使用的端口规则数，单位条
	PortRulesUsage *uint64 `json:"PortRulesUsage,omitempty" name:"PortRulesUsage"`

	// 已使用的域名规则数，单位条
	DomainRulesUsage *uint64 `json:"DomainRulesUsage,omitempty" name:"DomainRulesUsage"`

	// 最近7天的攻击次数，单位次
	Last7DayAttackCount *uint64 `json:"Last7DayAttackCount,omitempty" name:"Last7DayAttackCount"`
}

type BGPInstance struct {
	// 资产实例的详细信息
	InstanceDetail *InstanceRelation `json:"InstanceDetail,omitempty" name:"InstanceDetail"`

	// 资产实例的规格信息
	SpecificationLimit *BGPInstanceSpecification `json:"SpecificationLimit,omitempty" name:"SpecificationLimit"`

	// 资产实例的使用统计信息
	Usage *BGPInstanceUsages `json:"Usage,omitempty" name:"Usage"`

	// 资产实例所在的地域
	Region *RegionInfo `json:"Region,omitempty" name:"Region"`

	// 资产实例的防护状态，状态码如下：
	// "idle"：正常状态(无攻击)
	// "attacking"：攻击中
	// "blocking"：封堵中
	// "creating"：创建中
	// "deblocking"：解封中
	// "isolate"：回收隔离中
	Status *string `json:"Status,omitempty" name:"Status"`

	// 购买时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 到期时间
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 资产实例的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资产实例所属的套餐包信息，
	// 注意：当资产实例不是套餐包的实例时，此字段为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackInfo *PackInfo `json:"PackInfo,omitempty" name:"PackInfo"`

	// 高防包绑定的EIP属于的云产品信息
	EipProductInfos []*EipProductInfo `json:"EipProductInfos,omitempty" name:"EipProductInfos"`

	// 高防包绑定状态，取值[
	// "idle"：绑定已完成
	//  "bounding"：正在绑定中
	// "failed"：绑定失败
	// ]
	BoundStatus *string `json:"BoundStatus,omitempty" name:"BoundStatus"`

	// 四层防护严格级别
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`

	// CC防护开关
	CCEnable *int64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// 资源关联标签
	TagInfoList []*TagInfo `json:"TagInfoList,omitempty" name:"TagInfoList"`

	// 新版本1ip高防包
	IpCountNewFlag *uint64 `json:"IpCountNewFlag,omitempty" name:"IpCountNewFlag"`

	// 攻击封堵套餐标记
	VitalityVersion *uint64 `json:"VitalityVersion,omitempty" name:"VitalityVersion"`

	// 网络线路
	// 注意：此字段可能返回 null，表示取不到有效值。
	Line *uint64 `json:"Line,omitempty" name:"Line"`

	// 弹性业务带宽开关
	ElasticServiceBandwidth *uint64 `json:"ElasticServiceBandwidth,omitempty" name:"ElasticServiceBandwidth"`

	// 赠送的业务带宽
	GiftServiceBandWidth *int64 `json:"GiftServiceBandWidth,omitempty" name:"GiftServiceBandWidth"`
}

type BGPInstanceSpecification struct {
	// 保底防护峰值，单位Gbps
	ProtectBandwidth *uint64 `json:"ProtectBandwidth,omitempty" name:"ProtectBandwidth"`

	// 防护次数，单位次
	ProtectCountLimit *uint64 `json:"ProtectCountLimit,omitempty" name:"ProtectCountLimit"`

	// 防护IP数，单位个
	ProtectIPNumberLimit *uint64 `json:"ProtectIPNumberLimit,omitempty" name:"ProtectIPNumberLimit"`

	// 自动续费状态，取值[
	// 0：没有开启自动续费
	// 1：开启了自动续费
	// ]
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 联合产品标记，0代表普通高防包，1代表联合高防包
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnionPackFlag *uint64 `json:"UnionPackFlag,omitempty" name:"UnionPackFlag"`

	// 业务带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceBandWidth *uint64 `json:"ServiceBandWidth,omitempty" name:"ServiceBandWidth"`

	// 战斗服版本标记，0表示普通高防包，1表示战斗服高防包
	// 注意：此字段可能返回 null，表示取不到有效值。
	BattleEditionFlag *uint64 `json:"BattleEditionFlag,omitempty" name:"BattleEditionFlag"`

	// 渠道版标记，0表示普通高防包，1表示渠道版高防包
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelEditionFlag *uint64 `json:"ChannelEditionFlag,omitempty" name:"ChannelEditionFlag"`

	// 高防包企业版标记，0表示普通高防包；1表示企业版高防包
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnterpriseFlag *uint64 `json:"EnterpriseFlag,omitempty" name:"EnterpriseFlag"`

	// 高防包企业版弹性阈值，0表示未开启；大于0为弹性防护阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticLimit *uint64 `json:"ElasticLimit,omitempty" name:"ElasticLimit"`

	// 降配后的防护能力，单位Gbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownGradeProtect *uint64 `json:"DownGradeProtect,omitempty" name:"DownGradeProtect"`
}

type BGPInstanceUsages struct {
	// 已使用的防护次数，单位次
	ProtectCountUsage *uint64 `json:"ProtectCountUsage,omitempty" name:"ProtectCountUsage"`

	// 已防护的IP数，单位个
	ProtectIPNumberUsage *uint64 `json:"ProtectIPNumberUsage,omitempty" name:"ProtectIPNumberUsage"`

	// 最近7天的攻击次数，单位次
	Last7DayAttackCount *uint64 `json:"Last7DayAttackCount,omitempty" name:"Last7DayAttackCount"`
}

type BlackWhiteIpRelation struct {
	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// IP类型，取值[black(黑IP)，white(白IP)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// 黑白IP所属的实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`

	// ip掩码，0表示32位完整ip
	Mask *uint64 `json:"Mask,omitempty" name:"Mask"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type BoundIpInfo struct {
	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 绑定的产品分类，绑定操作为必填项，解绑操作可不填。取值[public（CVM、CLB产品），bm（黑石产品），eni（弹性网卡），vpngw（VPN网关）， natgw（NAT网关），waf（Web应用安全产品），fpc（金融产品），gaap（GAAP产品）, other(托管IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// IP所属的资源实例ID，绑定操作为必填项，解绑操作可不填。例如是弹性网卡的IP，则InstanceId填写弹性网卡的ID(eni-*); 如果绑定的是托管IP没有对应的资源实例ID，请填写"none";
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 产品分类下的子类型，绑定操作为必填项，解绑操作可不填。取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（弹性公网常规IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 运营商，绑定操作为必填项，解绑操作可不填。0：电信；1：联通；2：移动；5：BGP
	IspCode *uint64 `json:"IspCode,omitempty" name:"IspCode"`
}

type CCLevelPolicy struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 防护等级，可取值default表示默认策略，loose表示宽松，strict表示严格
	Level *string `json:"Level,omitempty" name:"Level"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type CCPrecisionPlyRecord struct {
	// 配置项类型，当前仅支持value
	FieldType *string `json:"FieldType,omitempty" name:"FieldType"`

	// 配置字段，可取值cgi， ua， cookie， referer， accept,  srcip
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 配置取值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 配置项值比对方式，可取值equal ，not_equal， include
	ValueOperator *string `json:"ValueOperator,omitempty" name:"ValueOperator"`
}

type CCPrecisionPolicy struct {
	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Ip地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 策略方式（丢弃或验证码）
	PolicyAction *string `json:"PolicyAction,omitempty" name:"PolicyAction"`

	// 策略列表
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitempty" name:"PolicyList"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type CCReqLimitPolicy struct {
	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Ip地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 策略项
	PolicyRecord *CCReqLimitPolicyRecord `json:"PolicyRecord,omitempty" name:"PolicyRecord"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type CCReqLimitPolicyRecord struct {
	// 统计周期，可取值1，10，30，60，单位秒
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 请求数，取值1~20000
	RequestNum *uint64 `json:"RequestNum,omitempty" name:"RequestNum"`

	// 频率限制策略方式，可取值alg表示验证码，drop表示丢弃
	Action *string `json:"Action,omitempty" name:"Action"`

	// 频率限制策略时长，可取值1~86400，单位秒
	ExecuteDuration *uint64 `json:"ExecuteDuration,omitempty" name:"ExecuteDuration"`

	// 策略项比对方式，可取值include表示包含，equal表示等于
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Uri，三个策略项仅可填其中之一
	Uri *string `json:"Uri,omitempty" name:"Uri"`

	// User-Agent，三个策略项仅可填其中之一
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// Cookie，三个策略项仅可填其中之一
	Cookie *string `json:"Cookie,omitempty" name:"Cookie"`
}

type CCThresholdPolicy struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Ip地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 清洗阈值
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type CcBlackWhiteIpPolicy struct {
	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// 黑白名单IP地址
	BlackWhiteIp *string `json:"BlackWhiteIp,omitempty" name:"BlackWhiteIp"`

	// 掩码
	Mask *uint64 `json:"Mask,omitempty" name:"Mask"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type CcGeoIPBlockConfig struct {
	// 区域类型，取值[
	// oversea(海外)
	// china(国内)
	// customized(自定义地区)
	// ]
	RegionType *string `json:"RegionType,omitempty" name:"RegionType"`

	// 封禁动作，取值[
	// drop(拦截)
	// alg(人机校验)
	// ]
	Action *string `json:"Action,omitempty" name:"Action"`

	// 配置ID，配置添加成功后生成；添加新配置时不用填写此字段，修改或删除配置时需要填写配置ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 当RegionType为customized时，必须填写AreaList；当RegionType为china或oversea时，AreaList为空
	AreaList []*int64 `json:"AreaList,omitempty" name:"AreaList"`
}

type CcGeoIpPolicyNew struct {
	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 用户动作，drop或alg
	Action *string `json:"Action,omitempty" name:"Action"`

	// 地域类型，分为china, oversea与customized
	RegionType *string `json:"RegionType,omitempty" name:"RegionType"`

	// 用户选择封禁的地域ID列表
	AreaList []*uint64 `json:"AreaList,omitempty" name:"AreaList"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type CertIdInsL7Rules struct {
	// 使用证书的规则列表
	L7Rules []*InsL7Rules `json:"L7Rules,omitempty" name:"L7Rules"`

	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`
}

type ConnectLimitConfig struct {
	// 基于源IP+目的IP的每秒新建数限制
	SdNewLimit *uint64 `json:"SdNewLimit,omitempty" name:"SdNewLimit"`

	// 基于目的IP的每秒新建数限制
	DstNewLimit *uint64 `json:"DstNewLimit,omitempty" name:"DstNewLimit"`

	// 基于源IP+目的IP的并发连接控制
	SdConnLimit *uint64 `json:"SdConnLimit,omitempty" name:"SdConnLimit"`

	// 基于目的IP+目的端口的并发连接控制
	DstConnLimit *uint64 `json:"DstConnLimit,omitempty" name:"DstConnLimit"`

	// 基于连接抑制触发阈值，取值范围[0,4294967295]
	BadConnThreshold *uint64 `json:"BadConnThreshold,omitempty" name:"BadConnThreshold"`

	// 异常连接检测条件，空连接防护开关，，取值范围[0,1]
	NullConnEnable *uint64 `json:"NullConnEnable,omitempty" name:"NullConnEnable"`

	// 异常连接检测条件，连接超时，，取值范围[0,65535]
	ConnTimeout *uint64 `json:"ConnTimeout,omitempty" name:"ConnTimeout"`

	// 异常连接检测条件，syn占比ack百分比，，取值范围[0,100]
	SynRate *uint64 `json:"SynRate,omitempty" name:"SynRate"`

	// 异常连接检测条件，syn阈值，取值范围[0,100]
	SynLimit *uint64 `json:"SynLimit,omitempty" name:"SynLimit"`
}

type ConnectLimitRelation struct {
	// 连接抑制配置
	ConnectLimitConfig *ConnectLimitConfig `json:"ConnectLimitConfig,omitempty" name:"ConnectLimitConfig"`

	// 连接抑制关联的实例信息
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

// Predefined struct for user
type CreateBlackWhiteIpListRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlackWhiteIpListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *CreateBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *CreateBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBoundIPRequestParams struct {
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 绑定到资源实例的IP数组，当资源实例为高防包(独享包)时，数组只允许填1个IP；当没有要绑定的IP时可以为空数组；但是BoundDevList和UnBoundDevList至少有一个不为空；
	BoundDevList []*BoundIpInfo `json:"BoundDevList,omitempty" name:"BoundDevList"`

	// 与资源实例解绑的IP数组，当资源实例为高防包(独享包)时，数组只允许填1个IP；当没有要解绑的IP时可以为空数组；但是BoundDevList和UnBoundDevList至少有一个不为空；
	UnBoundDevList []*BoundIpInfo `json:"UnBoundDevList,omitempty" name:"UnBoundDevList"`

	// 已弃用，不填
	CopyPolicy *string `json:"CopyPolicy,omitempty" name:"CopyPolicy"`
}

type CreateBoundIPRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp表示独享包；bgp-multip表示共享包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 绑定到资源实例的IP数组，当资源实例为高防包(独享包)时，数组只允许填1个IP；当没有要绑定的IP时可以为空数组；但是BoundDevList和UnBoundDevList至少有一个不为空；
	BoundDevList []*BoundIpInfo `json:"BoundDevList,omitempty" name:"BoundDevList"`

	// 与资源实例解绑的IP数组，当资源实例为高防包(独享包)时，数组只允许填1个IP；当没有要解绑的IP时可以为空数组；但是BoundDevList和UnBoundDevList至少有一个不为空；
	UnBoundDevList []*BoundIpInfo `json:"UnBoundDevList,omitempty" name:"UnBoundDevList"`

	// 已弃用，不填
	CopyPolicy *string `json:"CopyPolicy,omitempty" name:"CopyPolicy"`
}

func (r *CreateBoundIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBoundIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "BoundDevList")
	delete(f, "UnBoundDevList")
	delete(f, "CopyPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBoundIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBoundIPResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBoundIPResponse struct {
	*tchttp.BaseResponse
	Response *CreateBoundIPResponseParams `json:"Response"`
}

func (r *CreateBoundIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBoundIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCPrecisionPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP值
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议， 可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 策略方式，可取值alg表示验证码，drop表示丢弃
	PolicyAction *string `json:"PolicyAction,omitempty" name:"PolicyAction"`

	// 策略记录
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitempty" name:"PolicyList"`
}

type CreateCCPrecisionPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP值
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议， 可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 策略方式，可取值alg表示验证码，drop表示丢弃
	PolicyAction *string `json:"PolicyAction,omitempty" name:"PolicyAction"`

	// 策略记录
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitempty" name:"PolicyList"`
}

func (r *CreateCCPrecisionPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCPrecisionPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Protocol")
	delete(f, "Domain")
	delete(f, "PolicyAction")
	delete(f, "PolicyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCCPrecisionPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCPrecisionPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCCPrecisionPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCCPrecisionPolicyResponseParams `json:"Response"`
}

func (r *CreateCCPrecisionPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCPrecisionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCReqLimitPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP值
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 策略项
	Policy *CCReqLimitPolicyRecord `json:"Policy,omitempty" name:"Policy"`

	// 是否为兜底频控
	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

type CreateCCReqLimitPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP值
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 策略项
	Policy *CCReqLimitPolicyRecord `json:"Policy,omitempty" name:"Policy"`

	// 是否为兜底频控
	IsGlobal *int64 `json:"IsGlobal,omitempty" name:"IsGlobal"`
}

func (r *CreateCCReqLimitPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCReqLimitPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Protocol")
	delete(f, "Domain")
	delete(f, "Policy")
	delete(f, "IsGlobal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCCReqLimitPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCReqLimitPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCCReqLimitPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCCReqLimitPolicyResponseParams `json:"Response"`
}

func (r *CreateCCReqLimitPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCReqLimitPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcBlackWhiteIpListRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*IpSegment `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// Ip地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type CreateCcBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*IpSegment `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// Ip地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *CreateCcBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCcBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcBlackWhiteIpListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCcBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *CreateCcBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *CreateCcBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcGeoIPBlockConfigRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ip地址
	IP *string `json:"IP,omitempty" name:"IP"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议类型
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// CC区域封禁配置
	CcGeoIPBlockConfig *CcGeoIPBlockConfig `json:"CcGeoIPBlockConfig,omitempty" name:"CcGeoIPBlockConfig"`
}

type CreateCcGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ip地址
	IP *string `json:"IP,omitempty" name:"IP"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议类型
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// CC区域封禁配置
	CcGeoIPBlockConfig *CcGeoIPBlockConfig `json:"CcGeoIPBlockConfig,omitempty" name:"CcGeoIPBlockConfig"`
}

func (r *CreateCcGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IP")
	delete(f, "Domain")
	delete(f, "Protocol")
	delete(f, "CcGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCcGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcGeoIPBlockConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCcGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateCcGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *CreateCcGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSAIRequestParams struct {
	// 资源实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// AI防护开关，取值[
	// on(开启)
	// off(关闭)
	// ]
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`
}

type CreateDDoSAIRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// AI防护开关，取值[
	// on(开启)
	// off(关闭)
	// ]
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`
}

func (r *CreateDDoSAIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSAIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdList")
	delete(f, "DDoSAI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSAIResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDDoSAIResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSAIResponseParams `json:"Response"`
}

func (r *CreateDDoSAIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSAIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSBlackWhiteIpListRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*IpSegment `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateDDoSBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*IpSegment `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateDDoSBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSBlackWhiteIpListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDDoSBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *CreateDDoSBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSConnectLimitRequestParams struct {
	// 资源实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 连接抑制配置
	ConnectLimitConfig *ConnectLimitConfig `json:"ConnectLimitConfig,omitempty" name:"ConnectLimitConfig"`
}

type CreateDDoSConnectLimitRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 连接抑制配置
	ConnectLimitConfig *ConnectLimitConfig `json:"ConnectLimitConfig,omitempty" name:"ConnectLimitConfig"`
}

func (r *CreateDDoSConnectLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSConnectLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConnectLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSConnectLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSConnectLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDDoSConnectLimitResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSConnectLimitResponseParams `json:"Response"`
}

func (r *CreateDDoSConnectLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSConnectLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSGeoIPBlockConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DDoS区域封禁配置，填写参数时配置ID请为空
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitempty" name:"DDoSGeoIPBlockConfig"`
}

type CreateDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DDoS区域封禁配置，填写参数时配置ID请为空
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitempty" name:"DDoSGeoIPBlockConfig"`
}

func (r *CreateDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSGeoIPBlockConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *CreateDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSSpeedLimitConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 访问限速配置，填写参数时配置ID请为空
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitempty" name:"DDoSSpeedLimitConfig"`
}

type CreateDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 访问限速配置，填写参数时配置ID请为空
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitempty" name:"DDoSSpeedLimitConfig"`
}

func (r *CreateDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSSpeedLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSSpeedLimitConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSSpeedLimitConfigResponseParams `json:"Response"`
}

func (r *CreateDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultAlarmThresholdRequestParams struct {
	// 默认告警阈值配置
	DefaultAlarmConfig *DefaultAlarmThreshold `json:"DefaultAlarmConfig,omitempty" name:"DefaultAlarmConfig"`

	// 产品类型，取值[
	// bgp(表示高防包产品)
	// bgpip(表示高防IP产品)
	// ]
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type CreateDefaultAlarmThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 默认告警阈值配置
	DefaultAlarmConfig *DefaultAlarmThreshold `json:"DefaultAlarmConfig,omitempty" name:"DefaultAlarmConfig"`

	// 产品类型，取值[
	// bgp(表示高防包产品)
	// bgpip(表示高防IP产品)
	// ]
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *CreateDefaultAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DefaultAlarmConfig")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDefaultAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultAlarmThresholdResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDefaultAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *CreateDefaultAlarmThresholdResponseParams `json:"Response"`
}

func (r *CreateDefaultAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIPAlarmThresholdConfigRequestParams struct {
	// IP告警阈值配置列表
	IpAlarmThresholdConfigList []*IPAlarmThresholdRelation `json:"IpAlarmThresholdConfigList,omitempty" name:"IpAlarmThresholdConfigList"`
}

type CreateIPAlarmThresholdConfigRequest struct {
	*tchttp.BaseRequest
	
	// IP告警阈值配置列表
	IpAlarmThresholdConfigList []*IPAlarmThresholdRelation `json:"IpAlarmThresholdConfigList,omitempty" name:"IpAlarmThresholdConfigList"`
}

func (r *CreateIPAlarmThresholdConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPAlarmThresholdConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IpAlarmThresholdConfigList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIPAlarmThresholdConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIPAlarmThresholdConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIPAlarmThresholdConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateIPAlarmThresholdConfigResponseParams `json:"Response"`
}

func (r *CreateIPAlarmThresholdConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPAlarmThresholdConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RuleCertsRequestParams struct {
	// SSL证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// L7域名转发规则列表
	L7Rules []*InsL7Rules `json:"L7Rules,omitempty" name:"L7Rules"`
}

type CreateL7RuleCertsRequest struct {
	*tchttp.BaseRequest
	
	// SSL证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// L7域名转发规则列表
	L7Rules []*InsL7Rules `json:"L7Rules,omitempty" name:"L7Rules"`
}

func (r *CreateL7RuleCertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertId")
	delete(f, "L7Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7RuleCertsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RuleCertsResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL7RuleCertsResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7RuleCertsResponseParams `json:"Response"`
}

func (r *CreateL7RuleCertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL7RulesRequestParams struct {
	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`

	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 资源IP列表
	VipList []*string `json:"VipList,omitempty" name:"VipList"`
}

type CreateNewL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则列表
	Rules []*L7RuleEntry `json:"Rules,omitempty" name:"Rules"`

	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID列表
	IdList []*string `json:"IdList,omitempty" name:"IdList"`

	// 资源IP列表
	VipList []*string `json:"VipList,omitempty" name:"VipList"`
}

func (r *CreateNewL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "Business")
	delete(f, "IdList")
	delete(f, "VipList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNewL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL7RulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNewL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateNewL7RulesResponseParams `json:"Response"`
}

func (r *CreateNewL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePacketFilterConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 特征过滤规则
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`
}

type CreatePacketFilterConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 特征过滤规则
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`
}

func (r *CreatePacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PacketFilterConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePacketFilterConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreatePacketFilterConfigResponseParams `json:"Response"`
}

func (r *CreatePacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePortAclConfigListRequestParams struct {
	// 资源实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 端口acl策略
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`
}

type CreatePortAclConfigListRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 端口acl策略
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`
}

func (r *CreatePortAclConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePortAclConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdList")
	delete(f, "AclConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePortAclConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePortAclConfigListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePortAclConfigListResponse struct {
	*tchttp.BaseResponse
	Response *CreatePortAclConfigListResponseParams `json:"Response"`
}

func (r *CreatePortAclConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePortAclConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePortAclConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 端口acl策略
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`
}

type CreatePortAclConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 端口acl策略
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`
}

func (r *CreatePortAclConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePortAclConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AclConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePortAclConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePortAclConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePortAclConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreatePortAclConfigResponseParams `json:"Response"`
}

func (r *CreatePortAclConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePortAclConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProtocolBlockConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 协议封禁配置
	ProtocolBlockConfig *ProtocolBlockConfig `json:"ProtocolBlockConfig,omitempty" name:"ProtocolBlockConfig"`
}

type CreateProtocolBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 协议封禁配置
	ProtocolBlockConfig *ProtocolBlockConfig `json:"ProtocolBlockConfig,omitempty" name:"ProtocolBlockConfig"`
}

func (r *CreateProtocolBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtocolBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProtocolBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProtocolBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProtocolBlockConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProtocolBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateProtocolBlockConfigResponseParams `json:"Response"`
}

func (r *CreateProtocolBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtocolBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulingDomainRequestParams struct {
	// 代表是否混合云本地化的产品。
	// hybrid: 宙斯盾本地化
	// 不填写：其他
	Product *string `json:"Product,omitempty" name:"Product"`
}

type CreateSchedulingDomainRequest struct {
	*tchttp.BaseRequest
	
	// 代表是否混合云本地化的产品。
	// hybrid: 宙斯盾本地化
	// 不填写：其他
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateSchedulingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulingDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchedulingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulingDomainResponseParams struct {
	// 新创建的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSchedulingDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchedulingDomainResponseParams `json:"Response"`
}

func (r *CreateSchedulingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulingDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWaterPrintConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 水印防护配置
	WaterPrintConfig *WaterPrintConfig `json:"WaterPrintConfig,omitempty" name:"WaterPrintConfig"`
}

type CreateWaterPrintConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 水印防护配置
	WaterPrintConfig *WaterPrintConfig `json:"WaterPrintConfig,omitempty" name:"WaterPrintConfig"`
}

func (r *CreateWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WaterPrintConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWaterPrintConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateWaterPrintConfigResponseParams `json:"Response"`
}

func (r *CreateWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWaterPrintKeyRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CreateWaterPrintKeyRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreateWaterPrintKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWaterPrintKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWaterPrintKeyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWaterPrintKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateWaterPrintKeyResponseParams `json:"Response"`
}

func (r *CreateWaterPrintKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DDoSAIRelation struct {
	// AI防护开关，取值[
	// on(开启)
	// off(关闭)
	// ]
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`

	// AI防护开关所属的资源实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type DDoSGeoIPBlockConfig struct {
	// 区域类型，取值[
	// oversea(境外)
	// china(国内)
	// customized(自定义地区)
	// ]
	RegionType *string `json:"RegionType,omitempty" name:"RegionType"`

	// 封禁动作，取值[
	// drop(拦截)
	// trans(放行)
	// ]
	Action *string `json:"Action,omitempty" name:"Action"`

	// 配置ID，配置添加成功后生成；添加新配置时不用填写此字段，修改或删除配置时需要填写配置ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 当RegionType为customized时，必须填写AreaList，且最多填写128个；
	AreaList []*int64 `json:"AreaList,omitempty" name:"AreaList"`
}

type DDoSGeoIPBlockConfigRelation struct {
	// DDoS区域封禁配置
	GeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"GeoIPBlockConfig,omitempty" name:"GeoIPBlockConfig"`

	// 配置所属的资源实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type DDoSSpeedLimitConfig struct {
	// 限速模式，取值[
	// 1(基于源IP限速)
	// 2(基于目的端口限速)
	// ]
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 限速值，每种类型的限速值最多支持1个；该字段数组至少有一种限速值
	SpeedValues []*SpeedValue `json:"SpeedValues,omitempty" name:"SpeedValues"`

	// 此字段已弃用，请填写新字段DstPortList。
	DstPortScopes []*PortSegment `json:"DstPortScopes,omitempty" name:"DstPortScopes"`

	// 配置ID，配置添加成功后生成；添加新限制配置时不用填写此字段，修改或删除限速配置时需要填写配置ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// IP protocol numbers, 取值[
	// ALL(所有协议)
	// TCP(tcp协议)
	// UDP(udp协议)
	// SMP(smp协议)
	// 1;2-100(自定义协议号范围,最多8个)
	// ]
	// 注意：当自定义协议号范围时，只能填写协议号，多个范围;分隔；当填写ALL时不能再填写其他协议或协议号。
	ProtocolList *string `json:"ProtocolList,omitempty" name:"ProtocolList"`

	// 端口范围列表，最多8个，多个;分隔，范围表示用-；此端口范围必须填写；填写样式1:0-65535，样式2:80;443;1000-2000
	DstPortList *string `json:"DstPortList,omitempty" name:"DstPortList"`
}

type DDoSSpeedLimitConfigRelation struct {
	// DDoS访问限速配置
	SpeedLimitConfig *DDoSSpeedLimitConfig `json:"SpeedLimitConfig,omitempty" name:"SpeedLimitConfig"`

	// 配置所属的资源实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type DefaultAlarmThreshold struct {
	// 告警阈值类型，取值[
	// 1(入流量告警阈值)
	// 2(攻击清洗流量告警阈值)
	// ]
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警阈值，单位Mbps，取值>=0；当作为输入参数时，设置0会删除告警阈值配置；
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`
}

// Predefined struct for user
type DeleteCCLevelPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置策略的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值http
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DeleteCCLevelPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置策略的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值http
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DeleteCCLevelPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCLevelPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCLevelPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCLevelPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCCLevelPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCLevelPolicyResponseParams `json:"Response"`
}

func (r *DeleteCCLevelPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCLevelPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCPrecisionPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DeleteCCPrecisionPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteCCPrecisionPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCPrecisionPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCPrecisionPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCPrecisionPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCCPrecisionPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCPrecisionPolicyResponseParams `json:"Response"`
}

func (r *DeleteCCPrecisionPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCPrecisionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCRequestLimitPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DeleteCCRequestLimitPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteCCRequestLimitPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCRequestLimitPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCRequestLimitPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCRequestLimitPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCCRequestLimitPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCRequestLimitPolicyResponseParams `json:"Response"`
}

func (r *DeleteCCRequestLimitPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCRequestLimitPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCThresholdPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置策略的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值http
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DeleteCCThresholdPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置策略的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值http
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DeleteCCThresholdPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCThresholdPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCThresholdPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCThresholdPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCCThresholdPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCThresholdPolicyResponseParams `json:"Response"`
}

func (r *DeleteCCThresholdPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCThresholdPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcBlackWhiteIpListRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DeleteCcBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DeleteCcBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCcBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcBlackWhiteIpListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCcBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCcBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DeleteCcBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcGeoIPBlockConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// CC区域封禁配置，填写参数时配置ID不能为空
	CcGeoIPBlockConfig *CcGeoIPBlockConfig `json:"CcGeoIPBlockConfig,omitempty" name:"CcGeoIPBlockConfig"`
}

type DeleteCcGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// CC区域封禁配置，填写参数时配置ID不能为空
	CcGeoIPBlockConfig *CcGeoIPBlockConfig `json:"CcGeoIPBlockConfig,omitempty" name:"CcGeoIPBlockConfig"`
}

func (r *DeleteCcGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CcGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCcGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcGeoIPBlockConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCcGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCcGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *DeleteCcGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSBlackWhiteIpListRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*IpSegment `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DeleteDDoSBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*IpSegment `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteDDoSBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDDoSBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSBlackWhiteIpListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDDoSBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDDoSBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DeleteDDoSBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSGeoIPBlockConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DDoS区域封禁配置，填写参数时配置ID不能为空
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitempty" name:"DDoSGeoIPBlockConfig"`
}

type DeleteDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DDoS区域封禁配置，填写参数时配置ID不能为空
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitempty" name:"DDoSGeoIPBlockConfig"`
}

func (r *DeleteDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSGeoIPBlockConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDDoSGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *DeleteDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSSpeedLimitConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 访问限速配置，填写参数时配置ID不能为空
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitempty" name:"DDoSSpeedLimitConfig"`
}

type DeleteDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 访问限速配置，填写参数时配置ID不能为空
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitempty" name:"DDoSSpeedLimitConfig"`
}

func (r *DeleteDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSSpeedLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSSpeedLimitConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDDoSSpeedLimitConfigResponseParams `json:"Response"`
}

func (r *DeleteDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePacketFilterConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 特征过滤配置
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`
}

type DeletePacketFilterConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 特征过滤配置
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`
}

func (r *DeletePacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PacketFilterConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePacketFilterConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeletePacketFilterConfigResponseParams `json:"Response"`
}

func (r *DeletePacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePortAclConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 端口acl策略
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`
}

type DeletePortAclConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 端口acl策略
	AclConfig *AclConfig `json:"AclConfig,omitempty" name:"AclConfig"`
}

func (r *DeletePortAclConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePortAclConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AclConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePortAclConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePortAclConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePortAclConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeletePortAclConfigResponseParams `json:"Response"`
}

func (r *DeletePortAclConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePortAclConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWaterPrintConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DeleteWaterPrintConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWaterPrintConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWaterPrintConfigResponseParams `json:"Response"`
}

func (r *DeleteWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWaterPrintKeyRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 水印密钥ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type DeleteWaterPrintKeyRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 水印密钥ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DeleteWaterPrintKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWaterPrintKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWaterPrintKeyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWaterPrintKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWaterPrintKeyResponseParams `json:"Response"`
}

func (r *DeleteWaterPrintKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicDeviceStatusRequestParams struct {
	// IP 资源列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

type DescribeBasicDeviceStatusRequest struct {
	*tchttp.BaseRequest
	
	// IP 资源列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *DescribeBasicDeviceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDeviceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBasicDeviceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicDeviceStatusResponseParams struct {
	// 返回资源及状态，状态码：
	// 1 - 封堵状态
	// 2 - 正常状态
	// 3 - 攻击状态
	Data []*KeyValue `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBasicDeviceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBasicDeviceStatusResponseParams `json:"Response"`
}

func (r *DescribeBasicDeviceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDeviceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBgpBizTrendRequestParams struct {
	// 大禹子产品代号（bgp-multip表示高防包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计开始时间。 例：“2020-09-22 00:00:00”
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间。 例：“2020-09-22 00:00:00”
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计纬度，可取值intraffic, outtraffic, inpkg, outpkg
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 0表示固定时间，1表示自定义时间
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

type DescribeBgpBizTrendRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp-multip表示高防包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计开始时间。 例：“2020-09-22 00:00:00”
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间。 例：“2020-09-22 00:00:00”
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计纬度，可取值intraffic, outtraffic, inpkg, outpkg
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 0表示固定时间，1表示自定义时间
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

func (r *DescribeBgpBizTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBgpBizTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "InstanceId")
	delete(f, "Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBgpBizTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBgpBizTrendResponseParams struct {
	// 曲线图各个时间点的值
	DataList []*uint64 `json:"DataList,omitempty" name:"DataList"`

	// 曲线图取值个数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 统计纬度
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 返回数组最大值
	MaxData *uint64 `json:"MaxData,omitempty" name:"MaxData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBgpBizTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBgpBizTrendResponseParams `json:"Response"`
}

func (r *DescribeBgpBizTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBgpBizTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizHttpStatusRequestParams struct {
	// 统计方式，仅支持sum
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计周期，可取值60，300，1800，3600， 21600，86400，单位秒
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间。 如2020-02-01 12:04:12
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间。如2020-02-03 18:03:23
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 特定域名查询
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议及端口列表，协议可取值TCP, UDP, HTTP, HTTPS，仅统计纬度为连接数时有效
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitempty" name:"ProtoInfo"`
}

type DescribeBizHttpStatusRequest struct {
	*tchttp.BaseRequest
	
	// 统计方式，仅支持sum
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计周期，可取值60，300，1800，3600， 21600，86400，单位秒
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间。 如2020-02-01 12:04:12
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间。如2020-02-03 18:03:23
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 特定域名查询
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议及端口列表，协议可取值TCP, UDP, HTTP, HTTPS，仅统计纬度为连接数时有效
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitempty" name:"ProtoInfo"`
}

func (r *DescribeBizHttpStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizHttpStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Statistics")
	delete(f, "Business")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "Domain")
	delete(f, "ProtoInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBizHttpStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizHttpStatusResponseParams struct {
	// 业务流量http状态码统计数据
	HttpStatusMap *HttpStatusMap `json:"HttpStatusMap,omitempty" name:"HttpStatusMap"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBizHttpStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBizHttpStatusResponseParams `json:"Response"`
}

func (r *DescribeBizHttpStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizHttpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizTrendRequestParams struct {
	// 统计方式，可取值max, min, avg, sum, 如统计纬度是流量速率或包量速率，仅可取值max
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计周期，可取值60，300，1800，3600，21600，86400，单位秒
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间。 例：“2020-09-22 00:00:00”
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间。 例：“2020-09-22 00:00:00”
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计纬度，可取值connum, new_conn, inactive_conn, intraffic, outtraffic, inpkg, outpkg, qps
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计纬度为qps时，可选特定域名查询
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议及端口列表，协议可取值TCP, UDP, HTTP, HTTPS，仅统计纬度为连接数时有效
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitempty" name:"ProtoInfo"`
}

type DescribeBizTrendRequest struct {
	*tchttp.BaseRequest
	
	// 统计方式，可取值max, min, avg, sum, 如统计纬度是流量速率或包量速率，仅可取值max
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计周期，可取值60，300，1800，3600，21600，86400，单位秒
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间。 例：“2020-09-22 00:00:00”
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间。 例：“2020-09-22 00:00:00”
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 统计纬度，可取值connum, new_conn, inactive_conn, intraffic, outtraffic, inpkg, outpkg, qps
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计纬度为qps时，可选特定域名查询
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议及端口列表，协议可取值TCP, UDP, HTTP, HTTPS，仅统计纬度为连接数时有效
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitempty" name:"ProtoInfo"`
}

func (r *DescribeBizTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Statistics")
	delete(f, "Business")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "MetricName")
	delete(f, "Domain")
	delete(f, "ProtoInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBizTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizTrendResponseParams struct {
	// 曲线图各个时间点的值
	DataList []*float64 `json:"DataList,omitempty" name:"DataList"`

	// 统计纬度
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBizTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBizTrendResponseParams `json:"Response"`
}

func (r *DescribeBizTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlackWhiteIpListRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlackWhiteIpListResponseParams struct {
	// 黑名单IP列表
	BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList"`

	// 白名单IP列表
	WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DescribeBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCLevelListRequestParams struct {
	// 大禹子产品代号（bgp-multip表示高防包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeCCLevelListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp-multip表示高防包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeCCLevelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCLevelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCLevelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCLevelListResponseParams struct {
	// 分级策略列表总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 分级策略列表总数
	LevelList []*CCLevelPolicy `json:"LevelList,omitempty" name:"LevelList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCLevelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCLevelListResponseParams `json:"Response"`
}

func (r *DescribeCCLevelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCLevelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCLevelPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP值
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeCCLevelPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP值
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DescribeCCLevelPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCLevelPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCLevelPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCLevelPolicyResponseParams struct {
	// CC防护等级，可取值loose表示宽松，strict表示严格，normal表示适中， emergency表示攻击紧急， sup_loose表示超级宽松，default表示默认策略（无频控配置下发），customized表示自定义策略
	Level *string `json:"Level,omitempty" name:"Level"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCLevelPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCLevelPolicyResponseParams `json:"Response"`
}

func (r *DescribeCCLevelPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCLevelPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCPrecisionPlyListRequestParams struct {
	// 大禹子产品代号（bgpip-multip：表示高防包；bgpip：表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定特定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址，普通高防IP要传该字段
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名，普通高防IP要传该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，普通高防IP要传该字段
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeCCPrecisionPlyListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip-multip：表示高防包；bgpip：表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定特定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址，普通高防IP要传该字段
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名，普通高防IP要传该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，普通高防IP要传该字段
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DescribeCCPrecisionPlyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCPrecisionPlyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCPrecisionPlyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCPrecisionPlyListResponseParams struct {
	// 策略列表总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 策略列表详情
	PrecisionPolicyList []*CCPrecisionPolicy `json:"PrecisionPolicyList,omitempty" name:"PrecisionPolicyList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCPrecisionPlyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCPrecisionPlyListResponseParams `json:"Response"`
}

func (r *DescribeCCPrecisionPlyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCPrecisionPlyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCReqLimitPolicyListRequestParams struct {
	// 大禹子产品代号（bgp-multip表示高防包，bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址，普通高防IP要传该字段
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名，普通高防IP要传该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，普通高防IP要传该字段
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeCCReqLimitPolicyListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp-multip表示高防包，bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址，普通高防IP要传该字段
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名，普通高防IP要传该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，普通高防IP要传该字段
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DescribeCCReqLimitPolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCReqLimitPolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCReqLimitPolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCReqLimitPolicyListResponseParams struct {
	// 频率限制列表总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 频率限制列表详情
	RequestLimitPolicyList []*CCReqLimitPolicy `json:"RequestLimitPolicyList,omitempty" name:"RequestLimitPolicyList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCReqLimitPolicyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCReqLimitPolicyListResponseParams `json:"Response"`
}

func (r *DescribeCCReqLimitPolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCReqLimitPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCThresholdListRequestParams struct {
	// 大禹子产品代号（bgp-multip表示高防包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeCCThresholdListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp-multip表示高防包）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeCCThresholdListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCThresholdListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCThresholdListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCThresholdListResponseParams struct {
	// 清洗阈值策略列表总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 清洗阈值策略列表详情
	ThresholdList []*CCThresholdPolicy `json:"ThresholdList,omitempty" name:"ThresholdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCThresholdListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCThresholdListResponseParams `json:"Response"`
}

func (r *DescribeCCThresholdListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCThresholdListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCTrendRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[inqps(总请求峰值，dropqps(攻击请求峰值))，incount(请求次数), dropcount(攻击次数)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 域名，可选
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeCCTrendRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[inqps(总请求峰值，dropqps(攻击请求峰值))，incount(请求次数), dropcount(攻击次数)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 域名，可选
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeCCTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Ip")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Domain")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCTrendResponseParams struct {
	// 值个数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 值数组
	Data []*uint64 `json:"Data,omitempty" name:"Data"`

	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 指标，取值[inqps(总请求峰值，dropqps(攻击请求峰值))，incount(请求次数), dropcount(攻击次数)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCCTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCTrendResponseParams `json:"Response"`
}

func (r *DescribeCCTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcBlackWhiteIpListRequestParams struct {
	// 大禹子产品代号（bgp-multip：表示高防包；bgpip：表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 指定特定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// IP地址，普通高防IP要传该字段
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名，普通高防IP要传该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，普通高防IP要传该字段
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 筛选IP，需要筛选黑白名单IP时传该字段
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 黑白名单筛选字段，需要筛选黑白名单列表时传该字段
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

type DescribeCcBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgp-multip：表示高防包；bgpip：表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 指定特定实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// IP地址，普通高防IP要传该字段
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名，普通高防IP要传该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，普通高防IP要传该字段
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 筛选IP，需要筛选黑白名单IP时传该字段
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 黑白名单筛选字段，需要筛选黑白名单列表时传该字段
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`
}

func (r *DescribeCcBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	delete(f, "FilterIp")
	delete(f, "FilterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcBlackWhiteIpListResponseParams struct {
	// CC四层黑白名单策略列表总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// CC四层黑白名单策略列表详情
	CcBlackWhiteIpList []*CcBlackWhiteIpPolicy `json:"CcBlackWhiteIpList,omitempty" name:"CcBlackWhiteIpList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCcBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DescribeCcBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcGeoIPBlockConfigListRequestParams struct {
	// 大禹子产品代号（bgpip-multip：表示高防包；bgpip：表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定特定实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址，普通高防IP要传该字段
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名，普通高防IP要传该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，普通高防IP要传该字段
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeCcGeoIPBlockConfigListRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip-multip：表示高防包；bgpip：表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 指定特定实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址，普通高防IP要传该字段
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名，普通高防IP要传该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，普通高防IP要传该字段
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

func (r *DescribeCcGeoIPBlockConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcGeoIPBlockConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcGeoIPBlockConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcGeoIPBlockConfigListResponseParams struct {
	// CC地域封禁策略列表总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// CC地域封禁策略列表详情
	CcGeoIpPolicyList []*CcGeoIpPolicyNew `json:"CcGeoIpPolicyList,omitempty" name:"CcGeoIpPolicyList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCcGeoIPBlockConfigListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcGeoIPBlockConfigListResponseParams `json:"Response"`
}

func (r *DescribeCcGeoIPBlockConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcGeoIPBlockConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSBlackWhiteIpListRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDDoSBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDDoSBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSBlackWhiteIpListResponseParams struct {
	// 黑名单IP列表
	BlackIpList []*IpSegment `json:"BlackIpList,omitempty" name:"BlackIpList"`

	// 白名单IP列表
	WhiteIpList []*IpSegment `json:"WhiteIpList,omitempty" name:"WhiteIpList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DescribeDDoSBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSConnectLimitListRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可选参数，按照IP进行过滤
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 可选参数，按照实例id进行过滤
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`
}

type DescribeDDoSConnectLimitListRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可选参数，按照IP进行过滤
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 可选参数，按照实例id进行过滤
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`
}

func (r *DescribeDDoSConnectLimitListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSConnectLimitListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterIp")
	delete(f, "FilterInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSConnectLimitListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSConnectLimitListResponseParams struct {
	// 连接抑制配置总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 连接抑制配置详情信息
	ConfigList []*ConnectLimitRelation `json:"ConfigList,omitempty" name:"ConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSConnectLimitListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSConnectLimitListResponseParams `json:"Response"`
}

func (r *DescribeDDoSConnectLimitListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSConnectLimitListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSTrendRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeDDoSTrendRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源实例的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源实例ID，当Business为basic时，此字段不用填写（因为基础防护没有资源实例）
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDDoSTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Ip")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSTrendResponseParams struct {
	// 值个数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 值数组，攻击流量带宽单位为Mbps，包速率单位为pps
	Data []*uint64 `json:"Data,omitempty" name:"Data"`

	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDDoSTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSTrendResponseParams `json:"Response"`
}

func (r *DescribeDDoSTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultAlarmThresholdRequestParams struct {
	// 产品类型，取值[
	// bgp(表示高防包产品)
	// bgpip(表示高防IP产品)
	// ]
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 告警阈值类型搜索，取值[
	// 1(入流量告警阈值)
	// 2(攻击清洗流量告警阈值)
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitempty" name:"FilterAlarmType"`
}

type DescribeDefaultAlarmThresholdRequest struct {
	*tchttp.BaseRequest
	
	// 产品类型，取值[
	// bgp(表示高防包产品)
	// bgpip(表示高防IP产品)
	// ]
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 告警阈值类型搜索，取值[
	// 1(入流量告警阈值)
	// 2(攻击清洗流量告警阈值)
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitempty" name:"FilterAlarmType"`
}

func (r *DescribeDefaultAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "FilterAlarmType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultAlarmThresholdResponseParams struct {
	// 默认告警阈值配置
	DefaultAlarmConfigList []*DefaultAlarmThreshold `json:"DefaultAlarmConfigList,omitempty" name:"DefaultAlarmConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDefaultAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultAlarmThresholdResponseParams `json:"Response"`
}

func (r *DescribeDefaultAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7RulesBySSLCertIdRequestParams struct {
	// 域名状态，可取bindable, binded, opened, closed, all，all表示全部状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 证书ID列表
	CertIds []*string `json:"CertIds,omitempty" name:"CertIds"`
}

type DescribeL7RulesBySSLCertIdRequest struct {
	*tchttp.BaseRequest
	
	// 域名状态，可取bindable, binded, opened, closed, all，all表示全部状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 证书ID列表
	CertIds []*string `json:"CertIds,omitempty" name:"CertIds"`
}

func (r *DescribeL7RulesBySSLCertIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7RulesBySSLCertIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "CertIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7RulesBySSLCertIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7RulesBySSLCertIdResponseParams struct {
	// 证书规则集合
	CertSet []*CertIdInsL7Rules `json:"CertSet,omitempty" name:"CertSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL7RulesBySSLCertIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7RulesBySSLCertIdResponseParams `json:"Response"`
}

func (r *DescribeL7RulesBySSLCertIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7RulesBySSLCertIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBGPIPInstancesRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为20;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 资产实例ID搜索，例如，bgpip-00000001
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// 高防IP线路搜索，取值为[
	// 1：BGP线路
	// 2：电信
	// 3：联通
	// 4：移动
	// 99：第三方合作线路
	// ]
	FilterLine *uint64 `json:"FilterLine,omitempty" name:"FilterLine"`

	// 地域搜索，例如，ap-guangzhou
	FilterRegion *string `json:"FilterRegion,omitempty" name:"FilterRegion"`

	// 名称搜索
	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`

	// 是否只获取高防弹性公网IP实例。填写时，只能填写1或者0。当填写1时，表示返回高防弹性公网IP实例。当填写0时，表示返回非高防弹性公网IP实例。
	FilterEipType *int64 `json:"FilterEipType,omitempty" name:"FilterEipType"`

	// 高防弹性公网IP实例的绑定状态搜索条件，取值范围 [BINDING、 BIND、UNBINDING、UNBIND]。该搜索条件只在FilterEipType=1时才有效。
	FilterEipEipAddressStatus []*string `json:"FilterEipEipAddressStatus,omitempty" name:"FilterEipEipAddressStatus"`

	// 是否只获取安全加速实例。填写时，只能填写1或者0。当填写1时，表示返回安全加速实例。当填写0时，表示返回非安全加速实例。
	FilterDamDDoSStatus *int64 `json:"FilterDamDDoSStatus,omitempty" name:"FilterDamDDoSStatus"`

	// 获取特定状态的资源，运行中填idle，攻击中填attacking，封堵中填blocking，试用资源填trial
	FilterStatus *string `json:"FilterStatus,omitempty" name:"FilterStatus"`

	// 获取特定的实例Cname
	FilterCname *string `json:"FilterCname,omitempty" name:"FilterCname"`

	// 批量查询实例ID对应的高防IP实例资源
	FilterInstanceIdList []*string `json:"FilterInstanceIdList,omitempty" name:"FilterInstanceIdList"`

	// 标签搜索
	FilterTag *TagFilter `json:"FilterTag,omitempty" name:"FilterTag"`

	// 按照套餐类型进行过滤
	FilterPackType []*string `json:"FilterPackType,omitempty" name:"FilterPackType"`

	// 重保护航搜索
	FilterConvoy *uint64 `json:"FilterConvoy,omitempty" name:"FilterConvoy"`
}

type DescribeListBGPIPInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为20;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 资产实例ID搜索，例如，bgpip-00000001
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// 高防IP线路搜索，取值为[
	// 1：BGP线路
	// 2：电信
	// 3：联通
	// 4：移动
	// 99：第三方合作线路
	// ]
	FilterLine *uint64 `json:"FilterLine,omitempty" name:"FilterLine"`

	// 地域搜索，例如，ap-guangzhou
	FilterRegion *string `json:"FilterRegion,omitempty" name:"FilterRegion"`

	// 名称搜索
	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`

	// 是否只获取高防弹性公网IP实例。填写时，只能填写1或者0。当填写1时，表示返回高防弹性公网IP实例。当填写0时，表示返回非高防弹性公网IP实例。
	FilterEipType *int64 `json:"FilterEipType,omitempty" name:"FilterEipType"`

	// 高防弹性公网IP实例的绑定状态搜索条件，取值范围 [BINDING、 BIND、UNBINDING、UNBIND]。该搜索条件只在FilterEipType=1时才有效。
	FilterEipEipAddressStatus []*string `json:"FilterEipEipAddressStatus,omitempty" name:"FilterEipEipAddressStatus"`

	// 是否只获取安全加速实例。填写时，只能填写1或者0。当填写1时，表示返回安全加速实例。当填写0时，表示返回非安全加速实例。
	FilterDamDDoSStatus *int64 `json:"FilterDamDDoSStatus,omitempty" name:"FilterDamDDoSStatus"`

	// 获取特定状态的资源，运行中填idle，攻击中填attacking，封堵中填blocking，试用资源填trial
	FilterStatus *string `json:"FilterStatus,omitempty" name:"FilterStatus"`

	// 获取特定的实例Cname
	FilterCname *string `json:"FilterCname,omitempty" name:"FilterCname"`

	// 批量查询实例ID对应的高防IP实例资源
	FilterInstanceIdList []*string `json:"FilterInstanceIdList,omitempty" name:"FilterInstanceIdList"`

	// 标签搜索
	FilterTag *TagFilter `json:"FilterTag,omitempty" name:"FilterTag"`

	// 按照套餐类型进行过滤
	FilterPackType []*string `json:"FilterPackType,omitempty" name:"FilterPackType"`

	// 重保护航搜索
	FilterConvoy *uint64 `json:"FilterConvoy,omitempty" name:"FilterConvoy"`
}

func (r *DescribeListBGPIPInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPIPInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterIp")
	delete(f, "FilterInstanceId")
	delete(f, "FilterLine")
	delete(f, "FilterRegion")
	delete(f, "FilterName")
	delete(f, "FilterEipType")
	delete(f, "FilterEipEipAddressStatus")
	delete(f, "FilterDamDDoSStatus")
	delete(f, "FilterStatus")
	delete(f, "FilterCname")
	delete(f, "FilterInstanceIdList")
	delete(f, "FilterTag")
	delete(f, "FilterPackType")
	delete(f, "FilterConvoy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBGPIPInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBGPIPInstancesResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 高防IP资产实例列表
	InstanceList []*BGPIPInstance `json:"InstanceList,omitempty" name:"InstanceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListBGPIPInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListBGPIPInstancesResponseParams `json:"Response"`
}

func (r *DescribeListBGPIPInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPIPInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBGPInstancesRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为20;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 资产实例ID搜索，例如，bgp-00000001
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// 地域搜索，例如，ap-guangzhou
	FilterRegion *string `json:"FilterRegion,omitempty" name:"FilterRegion"`

	// 名称搜索
	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`

	// 按照线路搜索, 1: BGP; 2: 三网
	FilterLine *uint64 `json:"FilterLine,omitempty" name:"FilterLine"`

	// 状态搜索，idle：运行中；attacking：攻击中；blocking：封堵中
	FilterStatus *string `json:"FilterStatus,omitempty" name:"FilterStatus"`

	// 高防包绑定状态搜索，bounding：绑定中； failed：绑定失败
	FilterBoundStatus *string `json:"FilterBoundStatus,omitempty" name:"FilterBoundStatus"`

	// 实例id数组
	FilterInstanceIdList []*string `json:"FilterInstanceIdList,omitempty" name:"FilterInstanceIdList"`

	// 企业版搜索,  1：包含重保护航套餐下的企业版列表, 2: 不包含重保护航套餐的企业版列表
	FilterEnterpriseFlag *uint64 `json:"FilterEnterpriseFlag,omitempty" name:"FilterEnterpriseFlag"`

	// 轻量版搜索
	FilterLightFlag *uint64 `json:"FilterLightFlag,omitempty" name:"FilterLightFlag"`

	// 定制版搜索
	FilterChannelFlag *uint64 `json:"FilterChannelFlag,omitempty" name:"FilterChannelFlag"`

	// 标签搜索
	FilterTag *TagFilter `json:"FilterTag,omitempty" name:"FilterTag"`

	// 试用资源搜索，1: 应急防护资源；2：PLG试用资源
	FilterTrialFlag *uint64 `json:"FilterTrialFlag,omitempty" name:"FilterTrialFlag"`

	// 重保护航搜索
	FilterConvoy *uint64 `json:"FilterConvoy,omitempty" name:"FilterConvoy"`

	// 默认false；接口传true，返回数据中不包含高级信息，高级信息包含：InstanceList[0].Usage。
	ExcludeAdvancedInfo *bool `json:"ExcludeAdvancedInfo,omitempty" name:"ExcludeAdvancedInfo"`
}

type DescribeListBGPInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为20;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 资产实例ID搜索，例如，bgp-00000001
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// 地域搜索，例如，ap-guangzhou
	FilterRegion *string `json:"FilterRegion,omitempty" name:"FilterRegion"`

	// 名称搜索
	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`

	// 按照线路搜索, 1: BGP; 2: 三网
	FilterLine *uint64 `json:"FilterLine,omitempty" name:"FilterLine"`

	// 状态搜索，idle：运行中；attacking：攻击中；blocking：封堵中
	FilterStatus *string `json:"FilterStatus,omitempty" name:"FilterStatus"`

	// 高防包绑定状态搜索，bounding：绑定中； failed：绑定失败
	FilterBoundStatus *string `json:"FilterBoundStatus,omitempty" name:"FilterBoundStatus"`

	// 实例id数组
	FilterInstanceIdList []*string `json:"FilterInstanceIdList,omitempty" name:"FilterInstanceIdList"`

	// 企业版搜索,  1：包含重保护航套餐下的企业版列表, 2: 不包含重保护航套餐的企业版列表
	FilterEnterpriseFlag *uint64 `json:"FilterEnterpriseFlag,omitempty" name:"FilterEnterpriseFlag"`

	// 轻量版搜索
	FilterLightFlag *uint64 `json:"FilterLightFlag,omitempty" name:"FilterLightFlag"`

	// 定制版搜索
	FilterChannelFlag *uint64 `json:"FilterChannelFlag,omitempty" name:"FilterChannelFlag"`

	// 标签搜索
	FilterTag *TagFilter `json:"FilterTag,omitempty" name:"FilterTag"`

	// 试用资源搜索，1: 应急防护资源；2：PLG试用资源
	FilterTrialFlag *uint64 `json:"FilterTrialFlag,omitempty" name:"FilterTrialFlag"`

	// 重保护航搜索
	FilterConvoy *uint64 `json:"FilterConvoy,omitempty" name:"FilterConvoy"`

	// 默认false；接口传true，返回数据中不包含高级信息，高级信息包含：InstanceList[0].Usage。
	ExcludeAdvancedInfo *bool `json:"ExcludeAdvancedInfo,omitempty" name:"ExcludeAdvancedInfo"`
}

func (r *DescribeListBGPInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterIp")
	delete(f, "FilterInstanceId")
	delete(f, "FilterRegion")
	delete(f, "FilterName")
	delete(f, "FilterLine")
	delete(f, "FilterStatus")
	delete(f, "FilterBoundStatus")
	delete(f, "FilterInstanceIdList")
	delete(f, "FilterEnterpriseFlag")
	delete(f, "FilterLightFlag")
	delete(f, "FilterChannelFlag")
	delete(f, "FilterTag")
	delete(f, "FilterTrialFlag")
	delete(f, "FilterConvoy")
	delete(f, "ExcludeAdvancedInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBGPInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBGPInstancesResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 高防包资产实例列表
	InstanceList []*BGPInstance `json:"InstanceList,omitempty" name:"InstanceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListBGPInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListBGPInstancesResponseParams `json:"Response"`
}

func (r *DescribeListBGPInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBlackWhiteIpListRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

type DescribeListBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBlackWhiteIpListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 黑白IP列表
	IpList []*BlackWhiteIpRelation `json:"IpList,omitempty" name:"IpList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DescribeListBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSAIRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

type DescribeListDDoSAIRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListDDoSAIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSAIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListDDoSAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSAIResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// AI防护开关列表
	ConfigList []*DDoSAIRelation `json:"ConfigList,omitempty" name:"ConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListDDoSAIResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListDDoSAIResponseParams `json:"Response"`
}

func (r *DescribeListDDoSAIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSAIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSGeoIPBlockConfigRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

type DescribeListDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSGeoIPBlockConfigResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// DDoS区域封禁配置列表
	ConfigList []*DDoSGeoIPBlockConfigRelation `json:"ConfigList,omitempty" name:"ConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListDDoSGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *DescribeListDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSSpeedLimitConfigRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

type DescribeListDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSSpeedLimitConfigResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 访问限速配置列表
	ConfigList []*DDoSSpeedLimitConfigRelation `json:"ConfigList,omitempty" name:"ConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListDDoSSpeedLimitConfigResponseParams `json:"Response"`
}

func (r *DescribeListDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListIPAlarmConfigRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// 告警阈值类型搜索，取值[
	// 1(入流量告警阈值)
	// 2(攻击清洗流量告警阈值)
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitempty" name:"FilterAlarmType"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 高防IP实例资源的cname
	FilterCname *string `json:"FilterCname,omitempty" name:"FilterCname"`
}

type DescribeListIPAlarmConfigRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// 告警阈值类型搜索，取值[
	// 1(入流量告警阈值)
	// 2(攻击清洗流量告警阈值)
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitempty" name:"FilterAlarmType"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 高防IP实例资源的cname
	FilterCname *string `json:"FilterCname,omitempty" name:"FilterCname"`
}

func (r *DescribeListIPAlarmConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListIPAlarmConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterAlarmType")
	delete(f, "FilterIp")
	delete(f, "FilterCname")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListIPAlarmConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListIPAlarmConfigResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// IP告警阈值配置列表
	ConfigList []*IPAlarmThresholdRelation `json:"ConfigList,omitempty" name:"ConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListIPAlarmConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListIPAlarmConfigResponseParams `json:"Response"`
}

func (r *DescribeListIPAlarmConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListIPAlarmConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListListenerRequestParams struct {

}

type DescribeListListenerRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeListListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListListenerResponseParams struct {
	// 4层转发监听器列表
	Layer4Listeners []*Layer4Rule `json:"Layer4Listeners,omitempty" name:"Layer4Listeners"`

	// 7层转发监听器列表
	Layer7Listeners []*Layer7Rule `json:"Layer7Listeners,omitempty" name:"Layer7Listeners"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListListenerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListListenerResponseParams `json:"Response"`
}

func (r *DescribeListListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListPacketFilterConfigRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

type DescribeListPacketFilterConfigRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListPacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListPacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListPacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListPacketFilterConfigResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 特征过滤配置
	ConfigList []*PacketFilterRelation `json:"ConfigList,omitempty" name:"ConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListPacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListPacketFilterConfigResponseParams `json:"Response"`
}

func (r *DescribeListPacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListPacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListPortAclListRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// ip搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

type DescribeListPortAclListRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// ip搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListPortAclListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListPortAclListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListPortAclListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListPortAclListResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 端口acl策略
	AclList []*AclConfigRelation `json:"AclList,omitempty" name:"AclList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListPortAclListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListPortAclListResponseParams `json:"Response"`
}

func (r *DescribeListPortAclListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListPortAclListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListProtectThresholdConfigRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 域名搜索(查询域名与协议的CC防护阈值时使用）
	FilterDomain *string `json:"FilterDomain,omitempty" name:"FilterDomain"`

	// 协议搜索(查询域名与协议的CC防护阈值时使用）
	FilterProtocol *string `json:"FilterProtocol,omitempty" name:"FilterProtocol"`
}

type DescribeListProtectThresholdConfigRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`

	// 域名搜索(查询域名与协议的CC防护阈值时使用）
	FilterDomain *string `json:"FilterDomain,omitempty" name:"FilterDomain"`

	// 协议搜索(查询域名与协议的CC防护阈值时使用）
	FilterProtocol *string `json:"FilterProtocol,omitempty" name:"FilterProtocol"`
}

func (r *DescribeListProtectThresholdConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtectThresholdConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	delete(f, "FilterDomain")
	delete(f, "FilterProtocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListProtectThresholdConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListProtectThresholdConfigResponseParams struct {
	// 总记录数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 防护阈值配置列表
	ConfigList []*ProtectThresholdRelation `json:"ConfigList,omitempty" name:"ConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListProtectThresholdConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListProtectThresholdConfigResponseParams `json:"Response"`
}

func (r *DescribeListProtectThresholdConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtectThresholdConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListProtocolBlockConfigRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

type DescribeListProtocolBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListProtocolBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtocolBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListProtocolBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListProtocolBlockConfigResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 协议封禁配置
	ConfigList []*ProtocolBlockRelation `json:"ConfigList,omitempty" name:"ConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListProtocolBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListProtocolBlockConfigResponseParams `json:"Response"`
}

func (r *DescribeListProtocolBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtocolBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListSchedulingDomainRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为20;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 调度域名搜索
	FilterDomain *string `json:"FilterDomain,omitempty" name:"FilterDomain"`
}

type DescribeListSchedulingDomainRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为20;最大取值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 调度域名搜索
	FilterDomain *string `json:"FilterDomain,omitempty" name:"FilterDomain"`
}

func (r *DescribeListSchedulingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListSchedulingDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListSchedulingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListSchedulingDomainResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 调度域名信息列表
	DomainList []*SchedulingDomainInfo `json:"DomainList,omitempty" name:"DomainList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListSchedulingDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListSchedulingDomainResponseParams `json:"Response"`
}

func (r *DescribeListSchedulingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListSchedulingDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListWaterPrintConfigRequestParams struct {
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

type DescribeListWaterPrintConfigRequest struct {
	*tchttp.BaseRequest
	
	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 一页条数，当Limit=0时，默认一页条数为100;最大取值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 资源实例ID搜索, 支持资源实例前缀通配搜索，例如bgp-*表示获取高防包类型的资源实例
	FilterInstanceId *string `json:"FilterInstanceId,omitempty" name:"FilterInstanceId"`

	// IP搜索
	FilterIp *string `json:"FilterIp,omitempty" name:"FilterIp"`
}

func (r *DescribeListWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListWaterPrintConfigResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 水印配置列表
	ConfigList []*WaterPrintRelation `json:"ConfigList,omitempty" name:"ConfigList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeListWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListWaterPrintConfigResponseParams `json:"Response"`
}

func (r *DescribeListWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesErrHealthRequestParams struct {
	// 大禹子产品代号(bgpip表示高防IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// 规则Id列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type DescribeNewL7RulesErrHealthRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号(bgpip表示高防IP)
	Business *string `json:"Business,omitempty" name:"Business"`

	// 规则Id列表
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *DescribeNewL7RulesErrHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesErrHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewL7RulesErrHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesErrHealthResponseParams struct {
	// 异常规则列表，返回值说明: Key值为规则ID，Value值为异常IP及错误信息，多个IP用","分割
	ErrHealths []*KeyValue `json:"ErrHealths,omitempty" name:"ErrHealths"`

	// 异常规则的总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNewL7RulesErrHealthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewL7RulesErrHealthResponseParams `json:"Response"`
}

func (r *DescribeNewL7RulesErrHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesErrHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 状态搜索，选填，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList"`

	// 域名搜索，选填，当需要搜索域名请填写
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// IP搜索，选填，当需要搜索IP请填写
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 转发协议搜索，选填，取值[http, https, http/https]
	ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList"`

	// 高防IP实例的Cname
	Cname *string `json:"Cname,omitempty" name:"Cname"`
}

type DescribeNewL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 状态搜索，选填，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	StatusList []*uint64 `json:"StatusList,omitempty" name:"StatusList"`

	// 域名搜索，选填，当需要搜索域名请填写
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// IP搜索，选填，当需要搜索IP请填写
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 一页条数，填0表示不分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页起始偏移，取值为(页码-1)*一页条数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 转发协议搜索，选填，取值[http, https, http/https]
	ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList"`

	// 高防IP实例的Cname
	Cname *string `json:"Cname,omitempty" name:"Cname"`
}

func (r *DescribeNewL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "StatusList")
	delete(f, "Domain")
	delete(f, "Ip")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ProtocolList")
	delete(f, "Cname")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesResponseParams struct {
	// 转发规则列表
	Rules []*NewL7RuleEntry `json:"Rules,omitempty" name:"Rules"`

	// 健康检查配置列表
	Healths []*L7RuleHealth `json:"Healths,omitempty" name:"Healths"`

	// 总规则数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNewL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewL7RulesResponseParams `json:"Response"`
}

func (r *DescribeNewL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewAttackTrendRequestParams struct {
	// 攻击类型，取值ddos， cc
	Type *string `json:"Type,omitempty" name:"Type"`

	// 纬度，当前仅支持attackcount
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 周期，当前仅支持86400
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 防护概览攻击趋势开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 防护概览攻击趋势结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeOverviewAttackTrendRequest struct {
	*tchttp.BaseRequest
	
	// 攻击类型，取值ddos， cc
	Type *string `json:"Type,omitempty" name:"Type"`

	// 纬度，当前仅支持attackcount
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 周期，当前仅支持86400
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 防护概览攻击趋势开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 防护概览攻击趋势结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeOverviewAttackTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewAttackTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Dimension")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewAttackTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewAttackTrendResponseParams struct {
	// 攻击类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 防护概览攻击趋势起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 防护概览攻击趋势结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 周期
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 每个周期点的攻击次数
	Data []*uint64 `json:"Data,omitempty" name:"Data"`

	// 包含的周期点数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOverviewAttackTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewAttackTrendResponseParams `json:"Response"`
}

func (r *DescribeOverviewAttackTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewAttackTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewCCTrendRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp-multip表示共享包；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[inqps(总请求峰值，dropqps(攻击请求峰值))，incount(请求次数), dropcount(攻击次数)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源的IP
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeOverviewCCTrendRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp-multip表示共享包；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[inqps(总请求峰值，dropqps(攻击请求峰值))，incount(请求次数), dropcount(攻击次数)]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源的IP
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeOverviewCCTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewCCTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "IpList")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewCCTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewCCTrendResponseParams struct {
	// 值个数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 值数组
	Data []*uint64 `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOverviewCCTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewCCTrendResponseParams `json:"Response"`
}

func (r *DescribeOverviewCCTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewCCTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewDDoSEventListRequestParams struct {
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 可选按攻击状态过滤，start：攻击中；end：攻击结束
	AttackStatus *string `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 记录条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeOverviewDDoSEventListRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 可选按攻击状态过滤，start：攻击中；end：攻击结束
	AttackStatus *string `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 记录条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeOverviewDDoSEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewDDoSEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AttackStatus")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewDDoSEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewDDoSEventListResponseParams struct {
	// 记录总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 事件列表
	EventList []*OverviewDDoSEvent `json:"EventList,omitempty" name:"EventList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOverviewDDoSEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewDDoSEventListResponseParams `json:"Response"`
}

func (r *DescribeOverviewDDoSEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewDDoSEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewDDoSTrendRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP；bgp-multip表示高防包；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源实例的IP列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeOverviewDDoSTrendRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP；bgp-multip表示高防包；basic表示DDoS基础防护）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计粒度，取值[300(5分钟)，3600(小时)，86400(天)]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 统计开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 统计结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标，取值[bps(攻击流量带宽，pps(攻击包速率))]
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源实例的IP列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// 资源实例ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeOverviewDDoSTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewDDoSTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "IpList")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewDDoSTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewDDoSTrendResponseParams struct {
	// 值个数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 值数组，攻击流量带宽单位为Mbps，包速率单位为pps
	Data []*uint64 `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOverviewDDoSTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewDDoSTrendResponseParams `json:"Response"`
}

func (r *DescribeOverviewDDoSTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewDDoSTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewIndexRequestParams struct {
	// 拉取指标起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 拉取指标结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeOverviewIndexRequest struct {
	*tchttp.BaseRequest
	
	// 拉取指标起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 拉取指标结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeOverviewIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewIndexResponseParams struct {
	// IP总数
	AllIpCount *uint64 `json:"AllIpCount,omitempty" name:"AllIpCount"`

	// 高防IP总数（包含高防包+高防IP）
	AntiddosIpCount *uint64 `json:"AntiddosIpCount,omitempty" name:"AntiddosIpCount"`

	// 攻击IP总数
	AttackIpCount *uint64 `json:"AttackIpCount,omitempty" name:"AttackIpCount"`

	// 封堵IP总数
	BlockIpCount *uint64 `json:"BlockIpCount,omitempty" name:"BlockIpCount"`

	// 高防域名总数
	AntiddosDomainCount *uint64 `json:"AntiddosDomainCount,omitempty" name:"AntiddosDomainCount"`

	// 攻击域名总数
	AttackDomainCount *uint64 `json:"AttackDomainCount,omitempty" name:"AttackDomainCount"`

	// 攻击流量峰值
	MaxAttackFlow *uint64 `json:"MaxAttackFlow,omitempty" name:"MaxAttackFlow"`

	// 当前最近一条攻击中的起始时间
	NewAttackTime *string `json:"NewAttackTime,omitempty" name:"NewAttackTime"`

	// 当前最近一条攻击中的IP
	NewAttackIp *string `json:"NewAttackIp,omitempty" name:"NewAttackIp"`

	// 当前最近一条攻击中的攻击类型
	NewAttackType *string `json:"NewAttackType,omitempty" name:"NewAttackType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOverviewIndexResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewIndexResponseParams `json:"Response"`
}

func (r *DescribeOverviewIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePendingRiskInfoRequestParams struct {

}

type DescribePendingRiskInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePendingRiskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePendingRiskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePendingRiskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePendingRiskInfoResponseParams struct {
	// 是否为付费用户，true：付费用户， false：普通用户
	IsPaidUsr *bool `json:"IsPaidUsr,omitempty" name:"IsPaidUsr"`

	// 攻击中的资源数量
	AttackingCount *int64 `json:"AttackingCount,omitempty" name:"AttackingCount"`

	// 封堵中的资源数量
	BlockingCount *int64 `json:"BlockingCount,omitempty" name:"BlockingCount"`

	// 已过期的资源数量
	ExpiredCount *int64 `json:"ExpiredCount,omitempty" name:"ExpiredCount"`

	// 所有待处理风险事件总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePendingRiskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePendingRiskInfoResponseParams `json:"Response"`
}

func (r *DescribePendingRiskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePendingRiskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateDDoSEipAddressRequestParams struct {
	// 资源实例ID，实例ID形如：bgpip-0000011x。只能填写高防IP实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源实例ID对应的高防弹性公网IP。
	Eip *string `json:"Eip,omitempty" name:"Eip"`
}

type DisassociateDDoSEipAddressRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID，实例ID形如：bgpip-0000011x。只能填写高防IP实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源实例ID对应的高防弹性公网IP。
	Eip *string `json:"Eip,omitempty" name:"Eip"`
}

func (r *DisassociateDDoSEipAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDDoSEipAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Eip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateDDoSEipAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateDDoSEipAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateDDoSEipAddressResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateDDoSEipAddressResponseParams `json:"Response"`
}

func (r *DisassociateDDoSEipAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDDoSEipAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EipAddressPackRelation struct {
	// 套餐IP数量
	IpCount *uint64 `json:"IpCount,omitempty" name:"IpCount"`

	// 自动续费标记
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 当前到期时间
	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
}

type EipAddressRelation struct {
	// 高防弹性公网IP绑定的实例地区，例如hk代表香港
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipAddressRegion *string `json:"EipAddressRegion,omitempty" name:"EipAddressRegion"`

	// 绑定的资源实例ID。可能是一个CVM。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipBoundRscIns *string `json:"EipBoundRscIns,omitempty" name:"EipBoundRscIns"`

	// 绑定的弹性网卡ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipBoundRscEni *string `json:"EipBoundRscEni,omitempty" name:"EipBoundRscEni"`

	// 绑定的资源内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipBoundRscVip *string `json:"EipBoundRscVip,omitempty" name:"EipBoundRscVip"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type EipProductInfo struct {
	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 云产品类型，取值[
	// public（CVM产品），
	// bm（黑石产品），
	// eni（弹性网卡），
	// vpngw（VPN网关），
	//  natgw（NAT网关），
	// waf（Web应用安全产品），
	// fpc（金融产品），
	// gaap（GAAP产品）, 
	// other(托管IP)
	// ]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// 云产品子类型，取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（黑石弹性IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// IP所属的云产品实例ID，例如是弹性网卡的IP，InstanceId为弹性网卡的ID(eni-*); 如果是托管IP没有对应的资源实例ID,InstanceId为""
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ForwardListener struct {
	// 转发监听端口下限，取值1~65535
	FrontendPort *int64 `json:"FrontendPort,omitempty" name:"FrontendPort"`

	// 转发协议，取值[
	// TCP
	// UDP
	// ]
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`

	// 转发监听端口上限，取值1~65535
	FrontendPortEnd *int64 `json:"FrontendPortEnd,omitempty" name:"FrontendPortEnd"`
}

type HttpStatusMap struct {
	// http2xx回源状态码
	SourceHttp2xx []*float64 `json:"SourceHttp2xx,omitempty" name:"SourceHttp2xx"`

	// http5xx状态码
	Http5xx []*float64 `json:"Http5xx,omitempty" name:"Http5xx"`

	// http5xx回源状态码
	SourceHttp5xx []*float64 `json:"SourceHttp5xx,omitempty" name:"SourceHttp5xx"`

	// http404回源状态码
	SourceHttp404 []*float64 `json:"SourceHttp404,omitempty" name:"SourceHttp404"`

	// http4xx状态码
	Http4xx []*float64 `json:"Http4xx,omitempty" name:"Http4xx"`

	// http4xx回源状态码
	SourceHttp4xx []*float64 `json:"SourceHttp4xx,omitempty" name:"SourceHttp4xx"`

	// http2xx状态码
	Http2xx []*float64 `json:"Http2xx,omitempty" name:"Http2xx"`

	// http404状态码
	Http404 []*float64 `json:"Http404,omitempty" name:"Http404"`

	// http3xx回源状态码
	SourceHttp3xx []*float64 `json:"SourceHttp3xx,omitempty" name:"SourceHttp3xx"`

	// http3xx状态码
	Http3xx []*float64 `json:"Http3xx,omitempty" name:"Http3xx"`
}

type IPAlarmThresholdRelation struct {
	// 告警阈值类型，取值[
	// 1(入流量告警阈值)
	// 2(攻击清洗流量告警阈值)
	// ]
	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警阈值，单位Mbps，取值>=0；当作为输入参数时，设置0会删除告警阈值配置；
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitempty" name:"AlarmThreshold"`

	// 告警阈值所属的资源实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type IPLineInfo struct {
	// IP线路类型，取值[
	// "bgp"：BGP线路IP
	// "ctcc"：电信线路IP
	// "cucc"：联通线路IP
	// "cmcc"：移动线路IP
	// "abroad"：境外线路IP
	// ]
	Type *string `json:"Type,omitempty" name:"Type"`

	// 线路IP
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 实例对应的cname
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 资源flag，0：高防包资源，1：高防IP资源，2：非高防资源IP
	ResourceFlag *int64 `json:"ResourceFlag,omitempty" name:"ResourceFlag"`
}

type InsL7Rules struct {
	// 规则在中间状态态不可修改，只可在（0， 2， 8）状态可编辑。
	// 规则状态，0: 正常运行中, 1: 配置规则中(配置生效中), 2: 配置规则失败（配置生效失败）, 3: 删除规则中(删除生效中), 5: 删除规则失败(删除失败), 6: 等待添加规则, 7: 等待删除规则, 8: 等待上传证书, 9: 规则对应的资源不存在，被隔离, 10:等待修改规则, 11:配置修改中
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 实例ID
	InsId *string `json:"InsId,omitempty" name:"InsId"`

	// 用户AppID
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 高防端口
	VirtualPort *string `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// 证书ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`
}

type InstanceRelation struct {
	// 资源实例的IP
	EipList []*string `json:"EipList,omitempty" name:"EipList"`

	// 资源实例的ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type IpSegment struct {
	// ip地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// ip掩码，如果为32位ip，填0
	Mask *uint64 `json:"Mask,omitempty" name:"Mask"`
}

type KeyValue struct {
	// 字段名称
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段取值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type L4RuleSource struct {
	// 回源IP或域名
	Source *string `json:"Source,omitempty" name:"Source"`

	// 权重值，取值[0,100]
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 8000
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 备份源站，1: 备份源站，0: 普通源站
	// 注意：此字段可能返回 null，表示取不到有效值。
	Backup *uint64 `json:"Backup,omitempty" name:"Backup"`
}

type L7RuleEntry struct {
	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// 转发域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发协议，取值[http, https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 回源方式，取值[1(域名回源)，2(IP回源)]
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 负载均衡方式，取值[1(加权轮询)]
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// 回源列表
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList"`

	// 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// 规则状态，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 规则ID，当添加新规则时可以不用填写此字段；当修改或者删除规则时需要填写此字段；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// HTTPS协议的CC防护阈值
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// 当证书来源为自有证书时，此字段必须填写证书密钥；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// HTTPS协议的CC防护状态，取值[0(关闭), 1(开启)]
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// 是否开启Https协议使用Http回源，取值[0(关闭), 1(开启)]，不填写默认是关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpsToHttpEnable *uint64 `json:"HttpsToHttpEnable,omitempty" name:"HttpsToHttpEnable"`

	// 证书来源，当转发协议为https时必须填，取值[2(腾讯云托管证书)]，当转发协议为http时也可以填0
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 当证书来源为自有证书时，此字段必须填写证书内容；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// HTTPS协议的CC防护等级
	CCLevel *string `json:"CCLevel,omitempty" name:"CCLevel"`

	// 规则描述
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// cc防护状态，取值[0(关闭), 1(开启)]
	CCStatus *uint64 `json:"CCStatus,omitempty" name:"CCStatus"`

	// 接入端口值
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// 当证书来源为腾讯云托管证书时，此字段必须填写托管证书ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 同ruleId
	Id *string `json:"Id,omitempty" name:"Id"`

	// 智能cc开关，取值[0(关闭), 1(开启)]
	CCAIEnable *uint64 `json:"CCAIEnable,omitempty" name:"CCAIEnable"`
}

type L7RuleHealth struct {
	// 配置状态，0： 正常，1：配置中，2：配置失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// =1表示开启；=0表示关闭
	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 检查目录的URL，默认为/
	Url *string `json:"Url,omitempty" name:"Url"`

	// 检测间隔时间，单位秒
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 健康阈值，单位次
	AliveNum *uint64 `json:"AliveNum,omitempty" name:"AliveNum"`

	// 不健康阈值，单位次
	KickNum *uint64 `json:"KickNum,omitempty" name:"KickNum"`

	// HTTP请求方式，取值[HEAD,GET]
	Method *string `json:"Method,omitempty" name:"Method"`

	// 健康检查判定正常状态码，1xx =1, 2xx=2, 3xx=4, 4xx=8,5xx=16，多个状态码值加和
	StatusCode *uint64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 是否同时下发http和https规则健康检查配置
	ProtocolFlag *uint64 `json:"ProtocolFlag,omitempty" name:"ProtocolFlag"`

	// 被动探测开关，=1表示开启；=0表示关闭
	PassiveEnable *uint64 `json:"PassiveEnable,omitempty" name:"PassiveEnable"`

	// 被动探测不健康屏蔽时间
	BlockInter *uint64 `json:"BlockInter,omitempty" name:"BlockInter"`

	// 被动探测不健康统计间隔
	FailedCountInter *uint64 `json:"FailedCountInter,omitempty" name:"FailedCountInter"`

	// 被动探测不健康阈值
	FailedThreshold *uint64 `json:"FailedThreshold,omitempty" name:"FailedThreshold"`

	// 被动探测判定正常状态码，1xx =1, 2xx=2, 3xx=4, 4xx=8,5xx=16，多个状态码值加和
	PassiveStatusCode *uint64 `json:"PassiveStatusCode,omitempty" name:"PassiveStatusCode"`

	// 被动探测配置状态，0： 正常，1：配置中，2：配置失败
	PassiveStatus *uint64 `json:"PassiveStatus,omitempty" name:"PassiveStatus"`
}

type Layer4Rule struct {
	// 源站端口，取值1~65535
	BackendPort *uint64 `json:"BackendPort,omitempty" name:"BackendPort"`

	// 转发端口，取值1~65535
	FrontendPort *uint64 `json:"FrontendPort,omitempty" name:"FrontendPort"`

	// 转发协议，取值[
	// TCP(TCP协议)
	// UDP(UDP协议)
	// ]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 源站列表
	RealServers []*SourceServer `json:"RealServers,omitempty" name:"RealServers"`

	// 资源实例
	InstanceDetails []*InstanceRelation `json:"InstanceDetails,omitempty" name:"InstanceDetails"`

	// 规则所属的资源实例
	InstanceDetailRule []*RuleInstanceRelation `json:"InstanceDetailRule,omitempty" name:"InstanceDetailRule"`
}

type Layer7Rule struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发类型列表
	ProxyTypeList []*ProxyTypeInfo `json:"ProxyTypeList,omitempty" name:"ProxyTypeList"`

	// 源站列表
	RealServers []*SourceServer `json:"RealServers,omitempty" name:"RealServers"`

	// 资源实例
	InstanceDetails []*InstanceRelation `json:"InstanceDetails,omitempty" name:"InstanceDetails"`

	// 规则所属的资源实例
	InstanceDetailRule []*RuleInstanceRelation `json:"InstanceDetailRule,omitempty" name:"InstanceDetailRule"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口号
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type ListenerCcThreholdConfig struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议（可取值https）
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 开关状态（0：关闭，1：开启）
	CCEnable *int64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// cc防护阈值
	CCThreshold *int64 `json:"CCThreshold,omitempty" name:"CCThreshold"`
}

// Predefined struct for user
type ModifyCCLevelPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// CC防护等级，可取值loose表示宽松，strict表示严格，normal表示适中， emergency表示攻击紧急， sup_loose表示超级宽松，default表示默认策略（无频控配置下发），customized表示自定义策略
	Level *string `json:"Level,omitempty" name:"Level"`
}

type ModifyCCLevelPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// CC防护等级，可取值loose表示宽松，strict表示严格，normal表示适中， emergency表示攻击紧急， sup_loose表示超级宽松，default表示默认策略（无频控配置下发），customized表示自定义策略
	Level *string `json:"Level,omitempty" name:"Level"`
}

func (r *ModifyCCLevelPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCLevelPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	delete(f, "Level")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCLevelPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCLevelPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCLevelPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCLevelPolicyResponseParams `json:"Response"`
}

func (r *ModifyCCLevelPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCLevelPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCPrecisionPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略方式。可取值：alg、drop。alg指返回验证码方式验证，drop表示该访问丢弃。
	PolicyAction *string `json:"PolicyAction,omitempty" name:"PolicyAction"`

	// 策略记录
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitempty" name:"PolicyList"`
}

type ModifyCCPrecisionPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略方式。可取值：alg、drop。alg指返回验证码方式验证，drop表示该访问丢弃。
	PolicyAction *string `json:"PolicyAction,omitempty" name:"PolicyAction"`

	// 策略记录
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitempty" name:"PolicyList"`
}

func (r *ModifyCCPrecisionPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCPrecisionPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	delete(f, "PolicyAction")
	delete(f, "PolicyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCPrecisionPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCPrecisionPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCPrecisionPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCPrecisionPolicyResponseParams `json:"Response"`
}

func (r *ModifyCCPrecisionPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCPrecisionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCReqLimitPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略项
	Policy *CCReqLimitPolicyRecord `json:"Policy,omitempty" name:"Policy"`
}

type ModifyCCReqLimitPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略项
	Policy *CCReqLimitPolicyRecord `json:"Policy,omitempty" name:"Policy"`
}

func (r *ModifyCCReqLimitPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCReqLimitPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	delete(f, "Policy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCReqLimitPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCReqLimitPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCReqLimitPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCReqLimitPolicyResponseParams `json:"Response"`
}

func (r *ModifyCCReqLimitPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCReqLimitPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCThresholdPolicyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 清洗阈值，-1表示开启“默认”模式
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
}

type ModifyCCThresholdPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议，可取值HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 清洗阈值，-1表示开启“默认”模式
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
}

func (r *ModifyCCThresholdPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCThresholdPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	delete(f, "Threshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCThresholdPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCThresholdPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCCThresholdPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCThresholdPolicyResponseParams `json:"Response"`
}

func (r *ModifyCCThresholdPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCThresholdPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcBlackWhiteIpListRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*IpSegment `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

type ModifyCcBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*IpSegment `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *ModifyCcBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCcBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcBlackWhiteIpListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCcBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCcBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *ModifyCcBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSBlackWhiteIpListRequestParams struct {
	// 资源Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 当前配置的黑白名单类型，取值black时表示黑名单；取值white时表示白名单
	OldIpType *string `json:"OldIpType,omitempty" name:"OldIpType"`

	// 当前配置的Ip段，包含ip与掩码
	OldIp *IpSegment `json:"OldIp,omitempty" name:"OldIp"`

	// 修改后黑白名单类型，取值black时黑名单，取值white时白名单
	NewIpType *string `json:"NewIpType,omitempty" name:"NewIpType"`

	// 当前配置的Ip段，包含ip与掩码
	NewIp *IpSegment `json:"NewIp,omitempty" name:"NewIp"`
}

type ModifyDDoSBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// 资源Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 当前配置的黑白名单类型，取值black时表示黑名单；取值white时表示白名单
	OldIpType *string `json:"OldIpType,omitempty" name:"OldIpType"`

	// 当前配置的Ip段，包含ip与掩码
	OldIp *IpSegment `json:"OldIp,omitempty" name:"OldIp"`

	// 修改后黑白名单类型，取值black时黑名单，取值white时白名单
	NewIpType *string `json:"NewIpType,omitempty" name:"NewIpType"`

	// 当前配置的Ip段，包含ip与掩码
	NewIp *IpSegment `json:"NewIp,omitempty" name:"NewIp"`
}

func (r *ModifyDDoSBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldIpType")
	delete(f, "OldIp")
	delete(f, "NewIpType")
	delete(f, "NewIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSBlackWhiteIpListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *ModifyDDoSBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSGeoIPBlockConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DDoS区域封禁配置，填写参数时配置ID不能为空
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitempty" name:"DDoSGeoIPBlockConfig"`
}

type ModifyDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DDoS区域封禁配置，填写参数时配置ID不能为空
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitempty" name:"DDoSGeoIPBlockConfig"`
}

func (r *ModifyDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSGeoIPBlockConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *ModifyDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSLevelRequestParams struct {
	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// =get表示读取防护等级；=set表示修改防护等级
	Method *string `json:"Method,omitempty" name:"Method"`

	// 防护等级，取值[low,middle,high]；当Method=set时必填
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`
}

type ModifyDDoSLevelRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// =get表示读取防护等级；=set表示修改防护等级
	Method *string `json:"Method,omitempty" name:"Method"`

	// 防护等级，取值[low,middle,high]；当Method=set时必填
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`
}

func (r *ModifyDDoSLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Business")
	delete(f, "Method")
	delete(f, "DDoSLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSLevelResponseParams struct {
	// 防护等级，取值[low,middle,high]
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSLevelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSLevelResponseParams `json:"Response"`
}

func (r *ModifyDDoSLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSSpeedLimitConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 访问限速配置，填写参数时配置ID不能为空
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitempty" name:"DDoSSpeedLimitConfig"`
}

type ModifyDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 访问限速配置，填写参数时配置ID不能为空
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitempty" name:"DDoSSpeedLimitConfig"`
}

func (r *ModifyDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSSpeedLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSSpeedLimitConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSSpeedLimitConfigResponseParams `json:"Response"`
}

func (r *ModifyDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSThresholdRequestParams struct {
	// DDoS清洗阈值，取值[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000];
	// 当设置值为0时，表示采用默认值；
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 配置其他阈值标志位，1表示配置其他阈值
	OtherThresholdFlag *int64 `json:"OtherThresholdFlag,omitempty" name:"OtherThresholdFlag"`

	// SYN FLOOD流量阈值
	SynFloodThreshold *uint64 `json:"SynFloodThreshold,omitempty" name:"SynFloodThreshold"`

	// SYN FLOOD包量阈值
	SynFloodPktThreshold *uint64 `json:"SynFloodPktThreshold,omitempty" name:"SynFloodPktThreshold"`

	// UDP FLOOD流量阈值
	UdpFloodThreshold *uint64 `json:"UdpFloodThreshold,omitempty" name:"UdpFloodThreshold"`

	// UDP FLOOD包量阈值
	UdpFloodPktThreshold *uint64 `json:"UdpFloodPktThreshold,omitempty" name:"UdpFloodPktThreshold"`

	// ACK FLOOD流量阈值
	AckFloodThreshold *uint64 `json:"AckFloodThreshold,omitempty" name:"AckFloodThreshold"`

	// ACK FLOOD包量阈值
	AckFloodPktThreshold *uint64 `json:"AckFloodPktThreshold,omitempty" name:"AckFloodPktThreshold"`

	// SYNACK FLOOD流量阈值
	SynAckFloodThreshold *uint64 `json:"SynAckFloodThreshold,omitempty" name:"SynAckFloodThreshold"`

	// SYNACK FLOOD包量阈值
	SynAckFloodPktThreshold *uint64 `json:"SynAckFloodPktThreshold,omitempty" name:"SynAckFloodPktThreshold"`

	// RST FLOOD流量阈值
	RstFloodThreshold *uint64 `json:"RstFloodThreshold,omitempty" name:"RstFloodThreshold"`

	// RST FLOOD包量阈值
	RstFloodPktThreshold *uint64 `json:"RstFloodPktThreshold,omitempty" name:"RstFloodPktThreshold"`
}

type ModifyDDoSThresholdRequest struct {
	*tchttp.BaseRequest
	
	// DDoS清洗阈值，取值[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000];
	// 当设置值为0时，表示采用默认值；
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 大禹子产品代号（bgpip表示高防IP；bgp表示独享包；bgp-multip表示共享包；net表示高防IP专业版）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 配置其他阈值标志位，1表示配置其他阈值
	OtherThresholdFlag *int64 `json:"OtherThresholdFlag,omitempty" name:"OtherThresholdFlag"`

	// SYN FLOOD流量阈值
	SynFloodThreshold *uint64 `json:"SynFloodThreshold,omitempty" name:"SynFloodThreshold"`

	// SYN FLOOD包量阈值
	SynFloodPktThreshold *uint64 `json:"SynFloodPktThreshold,omitempty" name:"SynFloodPktThreshold"`

	// UDP FLOOD流量阈值
	UdpFloodThreshold *uint64 `json:"UdpFloodThreshold,omitempty" name:"UdpFloodThreshold"`

	// UDP FLOOD包量阈值
	UdpFloodPktThreshold *uint64 `json:"UdpFloodPktThreshold,omitempty" name:"UdpFloodPktThreshold"`

	// ACK FLOOD流量阈值
	AckFloodThreshold *uint64 `json:"AckFloodThreshold,omitempty" name:"AckFloodThreshold"`

	// ACK FLOOD包量阈值
	AckFloodPktThreshold *uint64 `json:"AckFloodPktThreshold,omitempty" name:"AckFloodPktThreshold"`

	// SYNACK FLOOD流量阈值
	SynAckFloodThreshold *uint64 `json:"SynAckFloodThreshold,omitempty" name:"SynAckFloodThreshold"`

	// SYNACK FLOOD包量阈值
	SynAckFloodPktThreshold *uint64 `json:"SynAckFloodPktThreshold,omitempty" name:"SynAckFloodPktThreshold"`

	// RST FLOOD流量阈值
	RstFloodThreshold *uint64 `json:"RstFloodThreshold,omitempty" name:"RstFloodThreshold"`

	// RST FLOOD包量阈值
	RstFloodPktThreshold *uint64 `json:"RstFloodPktThreshold,omitempty" name:"RstFloodPktThreshold"`
}

func (r *ModifyDDoSThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Threshold")
	delete(f, "Id")
	delete(f, "Business")
	delete(f, "OtherThresholdFlag")
	delete(f, "SynFloodThreshold")
	delete(f, "SynFloodPktThreshold")
	delete(f, "UdpFloodThreshold")
	delete(f, "UdpFloodPktThreshold")
	delete(f, "AckFloodThreshold")
	delete(f, "AckFloodPktThreshold")
	delete(f, "SynAckFloodThreshold")
	delete(f, "SynAckFloodPktThreshold")
	delete(f, "RstFloodThreshold")
	delete(f, "RstFloodPktThreshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSThresholdResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDDoSThresholdResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSThresholdResponseParams `json:"Response"`
}

func (r *ModifyDDoSThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainUsrNameRequestParams struct {
	// 用户CNAME
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 域名名称
	DomainUserName *string `json:"DomainUserName,omitempty" name:"DomainUserName"`
}

type ModifyDomainUsrNameRequest struct {
	*tchttp.BaseRequest
	
	// 用户CNAME
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 域名名称
	DomainUserName *string `json:"DomainUserName,omitempty" name:"DomainUserName"`
}

func (r *ModifyDomainUsrNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainUsrNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "DomainUserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainUsrNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainUsrNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDomainUsrNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainUsrNameResponseParams `json:"Response"`
}

func (r *ModifyDomainUsrNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainUsrNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewDomainRulesRequestParams struct {
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 域名转发规则
	Rule *NewL7RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

type ModifyNewDomainRulesRequest struct {
	*tchttp.BaseRequest
	
	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 域名转发规则
	Rule *NewL7RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyNewDomainRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewDomainRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNewDomainRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewDomainRulesResponseParams struct {
	// 成功码
	Success *SuccessCode `json:"Success,omitempty" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNewDomainRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNewDomainRulesResponseParams `json:"Response"`
}

func (r *ModifyNewDomainRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewDomainRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPacketFilterConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 特征过滤配置
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`
}

type ModifyPacketFilterConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 特征过滤配置
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`
}

func (r *ModifyPacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PacketFilterConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPacketFilterConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPacketFilterConfigResponseParams `json:"Response"`
}

func (r *ModifyPacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPortAclConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 旧端口acl策略
	OldAclConfig *AclConfig `json:"OldAclConfig,omitempty" name:"OldAclConfig"`

	// 新端口acl策略
	NewAclConfig *AclConfig `json:"NewAclConfig,omitempty" name:"NewAclConfig"`
}

type ModifyPortAclConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 旧端口acl策略
	OldAclConfig *AclConfig `json:"OldAclConfig,omitempty" name:"OldAclConfig"`

	// 新端口acl策略
	NewAclConfig *AclConfig `json:"NewAclConfig,omitempty" name:"NewAclConfig"`
}

func (r *ModifyPortAclConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPortAclConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldAclConfig")
	delete(f, "NewAclConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPortAclConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPortAclConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPortAclConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPortAclConfigResponseParams `json:"Response"`
}

func (r *ModifyPortAclConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPortAclConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewL7RuleEntry struct {
	// 转发协议，取值[http, https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 转发域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 负载均衡方式，取值[1(加权轮询)]
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// 回源方式，取值[1(域名回源)，2(IP回源)]
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 回源列表
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList"`

	// 区域码
	Region *uint64 `json:"Region,omitempty" name:"Region"`

	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源Ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 规则ID，当添加新规则时可以不用填写此字段；当修改或者删除规则时需要填写此字段；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则描述
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 证书来源，当转发协议为https时必须填，取值[2(腾讯云托管证书)]，当转发协议为http时也可以填0
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 当证书来源为腾讯云托管证书时，此字段必须填写托管证书ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 当证书来源为自有证书时，此字段必须填写证书内容；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 当证书来源为自有证书时，此字段必须填写证书密钥；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 规则状态，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// cc防护状态，取值[0(关闭), 1(开启)]
	CCStatus *uint64 `json:"CCStatus,omitempty" name:"CCStatus"`

	// HTTPS协议的CC防护状态，取值[0(关闭), 1(开启)]
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// HTTPS协议的CC防护阈值
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// HTTPS协议的CC防护等级
	CCLevel *string `json:"CCLevel,omitempty" name:"CCLevel"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 是否开启Https协议使用Http回源，取值[0(关闭), 1(开启)]，不填写默认是关闭
	HttpsToHttpEnable *uint64 `json:"HttpsToHttpEnable,omitempty" name:"HttpsToHttpEnable"`

	// 接入端口值
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// http强制跳转https，1表示打开，0表示关闭
	RewriteHttps *uint64 `json:"RewriteHttps,omitempty" name:"RewriteHttps"`

	// 规则配置失败时的详细错误原因(仅当Status=2时有效)，1001证书不存在，1002证书获取失败，1003证书上传失败，1004证书已过期
	ErrCode *uint64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *uint64 `json:"Version,omitempty" name:"Version"`
}

type OverviewDDoSEvent struct {
	// 事件Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// ip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 攻击类型
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`

	// 攻击状态，0：攻击中；1：攻击结束
	AttackStatus *uint64 `json:"AttackStatus,omitempty" name:"AttackStatus"`

	// 攻击流量，单位Mbps
	Mbps *uint64 `json:"Mbps,omitempty" name:"Mbps"`

	// 攻击包量，单位pps
	Pps *uint64 `json:"Pps,omitempty" name:"Pps"`

	// 业务类型，bgp-multip：高防包；bgpip：高防ip；basic：基础防护
	Business *string `json:"Business,omitempty" name:"Business"`

	// 高防实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 高防实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type PackInfo struct {
	// 套餐包的类型，取值[
	// staticpack：高防IP三网套餐包
	// insurance：保险套餐包
	// ]
	PackType *string `json:"PackType,omitempty" name:"PackType"`

	// 套餐包的ID
	PackId *string `json:"PackId,omitempty" name:"PackId"`
}

type PacketFilterConfig struct {
	// 协议，取值[tcp udp icmp all]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 起始源端口，取值0~65535
	SportStart *int64 `json:"SportStart,omitempty" name:"SportStart"`

	// 结束源端口，取值1~65535，必须大于等于起始源端口
	SportEnd *int64 `json:"SportEnd,omitempty" name:"SportEnd"`

	// 起始目的端口，取值0~65535
	DportStart *int64 `json:"DportStart,omitempty" name:"DportStart"`

	// 结束目的端口，取值1~65535，必须大于等于起始目的端口
	DportEnd *int64 `json:"DportEnd,omitempty" name:"DportEnd"`

	// 最小报文长度，取值1-1500
	PktlenMin *int64 `json:"PktlenMin,omitempty" name:"PktlenMin"`

	// 最大报文长度，取值1-1500，必须大于等于最小报文长度
	PktlenMax *int64 `json:"PktlenMax,omitempty" name:"PktlenMax"`

	// 动作，取值[
	// drop(丢弃)
	// transmit(放行)
	// drop_black(丢弃并拉黑)
	// drop_rst(拦截)
	// drop_black_rst(拦截并拉黑)
	// forward(继续防护)
	// ]
	Action *string `json:"Action,omitempty" name:"Action"`

	// 检测位置，取值[
	// begin_l3(IP头)
	// begin_l4(TCP/UDP头)
	// begin_l5(T载荷)
	// no_match(不匹配)
	// ]
	MatchBegin *string `json:"MatchBegin,omitempty" name:"MatchBegin"`

	// 检测类型，取值[
	// sunday(关键字)
	// pcre(正则表达式)
	// ]
	MatchType *string `json:"MatchType,omitempty" name:"MatchType"`

	// 检测值，关键字符串或正则表达式,取值[
	// 当检测类型为sunday时，请填写字符串或者16进制字节码，例如\x313233对应的是字符串"123"的16进制字节码;
	// 当检测类型为pcre时, 请填写正则表达式字符串;
	// ]
	Str *string `json:"Str,omitempty" name:"Str"`

	// 从检测位置开始的检测深度，取值[0,1500]
	Depth *int64 `json:"Depth,omitempty" name:"Depth"`

	// 从检测位置开始的偏移量，取值范围[0,Depth]
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 是否包含检测值，取值[
	// 0(包含)
	// 1(不包含)
	// ]
	IsNot *int64 `json:"IsNot,omitempty" name:"IsNot"`

	// 当有第二个检测条件时，与第一检测条件的且或关系，取值[
	// and(且的关系)
	// none(当没有第二个检测条件时填写此值)
	// ]
	MatchLogic *string `json:"MatchLogic,omitempty" name:"MatchLogic"`

	// 第二个检测位置，取值[
	// begin_l5(载荷)
	// no_match(不匹配)
	// ]
	MatchBegin2 *string `json:"MatchBegin2,omitempty" name:"MatchBegin2"`

	// 第二个检测类型，取值[
	// sunday(关键字)
	// pcre(正则表达式)
	// ]
	MatchType2 *string `json:"MatchType2,omitempty" name:"MatchType2"`

	// 第二个检测值，关键字符串或正则表达式,取值[
	// 当检测类型为sunday时，请填写字符串或者16进制字节码，例如\x313233对应的是字符串"123"的16进制字节码;
	// 当检测类型为pcre时, 请填写正则表达式字符串;
	// ]
	Str2 *string `json:"Str2,omitempty" name:"Str2"`

	// 从第二个检测位置开始的第二个检测深度，取值[0,1500]
	Depth2 *int64 `json:"Depth2,omitempty" name:"Depth2"`

	// 从第二个检测位置开始的偏移量，取值范围[0,Depth2]
	Offset2 *int64 `json:"Offset2,omitempty" name:"Offset2"`

	// 第二个检测是否包含检测值，取值[
	// 0(包含)
	// 1(不包含)
	// ]
	IsNot2 *int64 `json:"IsNot2,omitempty" name:"IsNot2"`

	// 特征过滤配置添加成功后自动生成的规则ID，当添加新特征过滤配置时，此字段不用填写；
	Id *string `json:"Id,omitempty" name:"Id"`

	// 大于报文长度，取值1+
	PktLenGT *int64 `json:"PktLenGT,omitempty" name:"PktLenGT"`
}

type PacketFilterRelation struct {
	// 特征过滤配置
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`

	// 特征过滤配置所属的实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type PortSegment struct {
	// 起始端口，取值1~65535
	BeginPort *uint64 `json:"BeginPort,omitempty" name:"BeginPort"`

	// 结束端口，取值1~65535，必须不小于起始端口
	EndPort *uint64 `json:"EndPort,omitempty" name:"EndPort"`
}

type ProtectThresholdRelation struct {
	// DDoS防护等级，取值[
	// low(宽松)
	// middle(适中)
	// high(严格)
	// ]
	DDoSLevel *string `json:"DDoSLevel,omitempty" name:"DDoSLevel"`

	// DDoS清洗阈值，单位Mbps
	DDoSThreshold *uint64 `json:"DDoSThreshold,omitempty" name:"DDoSThreshold"`

	// DDoS的AI防护开关，取值[
	// on(开启)
	// off(关闭)
	// ]
	DDoSAI *string `json:"DDoSAI,omitempty" name:"DDoSAI"`

	// CC清洗开关，取值[
	// 0(关闭)
	// 1(开启)
	// ]
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// CC清洗阈值，单位QPS
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// 所属的资源实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`

	// 域名与协议纬度的防护阈值
	ListenerCcThresholdList []*ListenerCcThreholdConfig `json:"ListenerCcThresholdList,omitempty" name:"ListenerCcThresholdList"`

	// SYN FLOOD流量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SynFloodThreshold *uint64 `json:"SynFloodThreshold,omitempty" name:"SynFloodThreshold"`

	// SYN FLOOD包量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SynFloodPktThreshold *uint64 `json:"SynFloodPktThreshold,omitempty" name:"SynFloodPktThreshold"`

	// UDP FLOOD流量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	UdpFloodThreshold *uint64 `json:"UdpFloodThreshold,omitempty" name:"UdpFloodThreshold"`

	// UDP FLOOD包量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	UdpFloodPktThreshold *uint64 `json:"UdpFloodPktThreshold,omitempty" name:"UdpFloodPktThreshold"`

	// ACK FLOOD流量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	AckFloodThreshold *uint64 `json:"AckFloodThreshold,omitempty" name:"AckFloodThreshold"`

	// ACK FLOOD包量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	AckFloodPktThreshold *uint64 `json:"AckFloodPktThreshold,omitempty" name:"AckFloodPktThreshold"`

	// SYNACK FLOOD流量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SynAckFloodThreshold *uint64 `json:"SynAckFloodThreshold,omitempty" name:"SynAckFloodThreshold"`

	// SYNACK FLOOD包量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SynAckFloodPktThreshold *uint64 `json:"SynAckFloodPktThreshold,omitempty" name:"SynAckFloodPktThreshold"`

	// RST FLOOD流量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RstFloodThreshold *uint64 `json:"RstFloodThreshold,omitempty" name:"RstFloodThreshold"`

	// RST FLOOD包量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RstFloodPktThreshold *uint64 `json:"RstFloodPktThreshold,omitempty" name:"RstFloodPktThreshold"`
}

type ProtocolBlockConfig struct {
	// TCP封禁，取值[0(封禁关)，1(封禁开)]
	DropTcp *int64 `json:"DropTcp,omitempty" name:"DropTcp"`

	// UDP封禁，取值[0(封禁关)，1(封禁开)]
	DropUdp *int64 `json:"DropUdp,omitempty" name:"DropUdp"`

	// ICMP封禁，取值[0(封禁关)，1(封禁开)]
	DropIcmp *int64 `json:"DropIcmp,omitempty" name:"DropIcmp"`

	// 其他协议封禁，取值[0(封禁关)，1(封禁开)]
	DropOther *int64 `json:"DropOther,omitempty" name:"DropOther"`

	// 异常空连接防护，取值[0(防护关)，1(防护开)]
	CheckExceptNullConnect *int64 `json:"CheckExceptNullConnect,omitempty" name:"CheckExceptNullConnect"`

	// ping of death防护，取值[0(防护关)，1(防护开)]
	PingOfDeath *int64 `json:"PingOfDeath,omitempty" name:"PingOfDeath"`

	// tear drop防护，取值[0(防护关)，1(防护开)]
	TearDrop *int64 `json:"TearDrop,omitempty" name:"TearDrop"`
}

type ProtocolBlockRelation struct {
	// 协议封禁配置
	ProtocolBlockConfig *ProtocolBlockConfig `json:"ProtocolBlockConfig,omitempty" name:"ProtocolBlockConfig"`

	// 协议封禁配置所属的实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}

type ProtocolPort struct {
	// 协议（tcp；udp）
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

type ProxyTypeInfo struct {
	// 转发监听端口列表，端口取值1~65535
	ProxyPorts []*int64 `json:"ProxyPorts,omitempty" name:"ProxyPorts"`

	// 转发协议，取值[
	// http(HTTP协议)
	// https(HTTPS协议)
	// ]
	ProxyType *string `json:"ProxyType,omitempty" name:"ProxyType"`
}

type RegionInfo struct {
	// 地域名称，例如，ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`
}

type RuleInstanceRelation struct {
	// 资源实例的IP
	EipList []*string `json:"EipList,omitempty" name:"EipList"`

	// 资源实例的ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源实例的Cname
	Cname *string `json:"Cname,omitempty" name:"Cname"`
}

type SchedulingDomainInfo struct {
	// 调度域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 线路IP列表
	LineIPList []*IPLineInfo `json:"LineIPList,omitempty" name:"LineIPList"`

	// 调度方式，当前仅支持优先级的方式，取值[priority]
	Method *string `json:"Method,omitempty" name:"Method"`

	// 调度域名解析记录的TTL值
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 运行状态，取值[
	// 0：未运行
	// 1：运行中
	// 2：运行异常
	// ]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 最后修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 域名名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsrDomainName *string `json:"UsrDomainName,omitempty" name:"UsrDomainName"`
}

type SourceServer struct {
	// 源站的地址（IP或者域名）
	RealServer *string `json:"RealServer,omitempty" name:"RealServer"`

	// 源站的地址类型，取值[
	// 1(域名地址)
	// 2(IP地址)
	// ]
	RsType *int64 `json:"RsType,omitempty" name:"RsType"`

	// 源站的回源权重，取值1~100
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 端口号：0~65535
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type SpeedValue struct {
	// 限速值类型，取值[
	// 1(包速率pps)
	// 2(带宽bps)
	// ]
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 值大小
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type StaticPackRelation struct {
	// 保底带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectBandwidth *uint64 `json:"ProtectBandwidth,omitempty" name:"ProtectBandwidth"`

	// 业务带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalBandwidth *uint64 `json:"NormalBandwidth,omitempty" name:"NormalBandwidth"`

	// 转发规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardRulesLimit *uint64 `json:"ForwardRulesLimit,omitempty" name:"ForwardRulesLimit"`

	// 自动续费标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`
}

type SuccessCode struct {
	// 描述
	Message *string `json:"Message,omitempty" name:"Message"`

	// 成功/错误码
	Code *string `json:"Code,omitempty" name:"Code"`
}

// Predefined struct for user
type SwitchWaterPrintConfigRequestParams struct {
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 水印开启/关闭状态，1表示开启；0表示关闭
	OpenStatus *int64 `json:"OpenStatus,omitempty" name:"OpenStatus"`

	// 是否开启代理，1开启则忽略IP+端口校验；0关闭则需要IP+端口校验
	CloudSdkProxy *int64 `json:"CloudSdkProxy,omitempty" name:"CloudSdkProxy"`
}

type SwitchWaterPrintConfigRequest struct {
	*tchttp.BaseRequest
	
	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 水印开启/关闭状态，1表示开启；0表示关闭
	OpenStatus *int64 `json:"OpenStatus,omitempty" name:"OpenStatus"`

	// 是否开启代理，1开启则忽略IP+端口校验；0关闭则需要IP+端口校验
	CloudSdkProxy *int64 `json:"CloudSdkProxy,omitempty" name:"CloudSdkProxy"`
}

func (r *SwitchWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OpenStatus")
	delete(f, "CloudSdkProxy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchWaterPrintConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwitchWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *SwitchWaterPrintConfigResponseParams `json:"Response"`
}

func (r *SwitchWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签键值列表
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type WaterPrintConfig struct {
	// 水印偏移量，取值范围[0, 100)
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 是否开启，取值[
	// 0（手动开启）
	// 1（立即运行）
	// ]
	OpenStatus *int64 `json:"OpenStatus,omitempty" name:"OpenStatus"`

	// 水印所属的转发监听器列表
	Listeners []*ForwardListener `json:"Listeners,omitempty" name:"Listeners"`

	// 水印添加成功后生成的水印密钥列表，一条水印最少1个密钥，最多2个密钥
	Keys []*WaterPrintKey `json:"Keys,omitempty" name:"Keys"`

	// 水印检查模式, 取值[
	// checkall（普通模式）
	// shortfpcheckall（精简模式）
	// ]
	Verify *string `json:"Verify,omitempty" name:"Verify"`

	// 是否开启代理，1开启则忽略IP+端口校验；0关闭则需要IP+端口校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudSdkProxy *int64 `json:"CloudSdkProxy,omitempty" name:"CloudSdkProxy"`
}

type WaterPrintKey struct {
	// 密钥版本号
	KeyVersion *string `json:"KeyVersion,omitempty" name:"KeyVersion"`

	// 密钥内容
	KeyContent *string `json:"KeyContent,omitempty" name:"KeyContent"`

	// 密钥ID
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 密钥启用状态，只有一个取值1(启用)
	KeyOpenStatus *int64 `json:"KeyOpenStatus,omitempty" name:"KeyOpenStatus"`

	// 密钥生成时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type WaterPrintRelation struct {
	// 水印配置
	WaterPrintConfig *WaterPrintConfig `json:"WaterPrintConfig,omitempty" name:"WaterPrintConfig"`

	// 水印配置所属的资源实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
}