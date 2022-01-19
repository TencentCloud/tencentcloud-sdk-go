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

type AssociateDDoSEipAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type AssociateDDoSEipLoadBalancerRequest struct {
	*tchttp.BaseRequest

	// 资源实例ID，实例ID形如：bgpip-0000011x。只能填写高防IP实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源实例ID对应的高防弹性公网IP。
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 要绑定的负载均衡ID。负载均衡 ID 形如：lb-0000002i。可通过登录控制台查询，也可通过 DescribeLoadBalancers 接口返回值中的LoadBalancerId获取。
	LoadBalancerID *string `json:"LoadBalancerID,omitempty" name:"LoadBalancerID"`

	// Clb所在地域，例如：ap-hongkong。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDDoSEipLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateDDoSEipLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
}

type BoundIpInfo struct {

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 绑定的产品分类，取值[public（CVM、CLB产品），bm（黑石产品），eni（弹性网卡），vpngw（VPN网关）， natgw（NAT网关），waf（Web应用安全产品），fpc（金融产品），gaap（GAAP产品）, other(托管IP)]
	BizType *string `json:"BizType,omitempty" name:"BizType"`

	// IP所属的资源实例ID，当绑定新IP时必须填写此字段；例如是弹性网卡的IP，则InstanceId填写弹性网卡的ID(eni-*); 如果绑定的是托管IP没有对应的资源实例ID，请填写"none";
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 产品分类下的子类型，取值[cvm（CVM），lb（负载均衡器），eni（弹性网卡），vpngw（VPN），natgw（NAT），waf（WAF），fpc（金融），gaap（GAAP），other（托管IP），eip（黑石弹性IP）]
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 运营商，0：电信；1：联通；2：移动；5：BGP
	IspCode *uint64 `json:"IspCode,omitempty" name:"IspCode"`
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

type CreateBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateBoundIPResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功码
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateDDoSAIResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateDDoSBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateDDoSConnectLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateDefaultAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateIPAlarmThresholdConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateL7RuleCertsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功码
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreatePacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreatePortAclConfigListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreatePortAclConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateProtocolBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateSchedulingDomainRequest struct {
	*tchttp.BaseRequest
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchedulingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSchedulingDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新创建的域名
		Domain *string `json:"Domain,omitempty" name:"Domain"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateWaterPrintKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	// oversea(海外)
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

type DeleteBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest

	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`

	// IP类型，取值[black(黑名单IP), white(白名单IP)]
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DeleteBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DeleteDDoSBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeletePacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteWaterPrintKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBasicDeviceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回资源及状态，状态码：
	// 1 - 封堵状态
	// 2 - 正常状态
	// 3 - 攻击状态
		Data []*KeyValue `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBizTrendRequest struct {
	*tchttp.BaseRequest

	// 统计方式，可取值max, min, avg, sum, 如统计纬度是流量速率或包量速率，仅可取值max
	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`

	// 大禹子产品代号（bgpip表示高防IP）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 统计周期，可取值300，1800，3600，21600，86400，单位秒
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

type DescribeBizTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 曲线图各个时间点的值
		DataList []*float64 `json:"DataList,omitempty" name:"DataList"`

		// 统计纬度
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 黑名单IP列表
		BlackIpList []*string `json:"BlackIpList,omitempty" name:"BlackIpList"`

		// 白名单IP列表
		WhiteIpList []*string `json:"WhiteIpList,omitempty" name:"WhiteIpList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCCTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeDDoSBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 黑名单IP列表
		BlackIpList []*IpSegment `json:"BlackIpList,omitempty" name:"BlackIpList"`

		// 白名单IP列表
		WhiteIpList []*IpSegment `json:"WhiteIpList,omitempty" name:"WhiteIpList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDDoSConnectLimitListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 连接抑制配置总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 连接抑制配置详情信息
		ConfigList []*ConnectLimitRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDDoSTrendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeDefaultAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 默认告警阈值配置
		DefaultAlarmConfigList []*DefaultAlarmThreshold `json:"DefaultAlarmConfigList,omitempty" name:"DefaultAlarmConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeL7RulesBySSLCertIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书规则集合
		CertSet []*CertIdInsL7Rules `json:"CertSet,omitempty" name:"CertSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 获取特定状态的资源，运行中填idle，攻击中填attacking，封堵中填blocking
	FilterStatus *string `json:"FilterStatus,omitempty" name:"FilterStatus"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBGPIPInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListBGPIPInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 高防IP资产实例列表
		InstanceList []*BGPIPInstance `json:"InstanceList,omitempty" name:"InstanceList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 状态搜索，idle：允许中；attacking：攻击中；blocking：封堵中
	FilterStatus *string `json:"FilterStatus,omitempty" name:"FilterStatus"`

	// 高防包绑定状态搜索，bounding：绑定中； failed：绑定失败
	FilterBoundStatus *string `json:"FilterBoundStatus,omitempty" name:"FilterBoundStatus"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBGPInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListBGPInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 高防包资产实例列表
		InstanceList []*BGPInstance `json:"InstanceList,omitempty" name:"InstanceList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 黑白IP列表
		IpList []*BlackWhiteIpRelation `json:"IpList,omitempty" name:"IpList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListDDoSAIResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// AI防护开关列表
		ConfigList []*DDoSAIRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// DDoS区域封禁配置列表
		ConfigList []*DDoSGeoIPBlockConfigRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 访问限速配置列表
		ConfigList []*DDoSSpeedLimitConfigRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListIPAlarmConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeListIPAlarmConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// IP告警阈值配置列表
		ConfigList []*IPAlarmThresholdRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListListenerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 4层转发监听器列表
		Layer4Listeners []*Layer4Rule `json:"Layer4Listeners,omitempty" name:"Layer4Listeners"`

		// 7层转发监听器列表
		Layer7Listeners []*Layer7Rule `json:"Layer7Listeners,omitempty" name:"Layer7Listeners"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListPacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 特征过滤配置
		ConfigList []*PacketFilterRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListPortAclListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 端口acl策略
		AclList []*AclConfigRelation `json:"AclList,omitempty" name:"AclList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListProtectThresholdConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 防护阈值配置列表
		ConfigList []*ProtectThresholdRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListProtocolBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 协议封禁配置
		ConfigList []*ProtocolBlockRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListSchedulingDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 调度域名信息列表
		DomainList []*SchedulingDomainInfo `json:"DomainList,omitempty" name:"DomainList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeListWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 水印配置列表
		ConfigList []*WaterPrintRelation `json:"ConfigList,omitempty" name:"ConfigList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DisassociateDDoSEipAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
}

type InsL7Rules struct {

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

	// 规则所属的资源实例
	InstanceDetails []*InstanceRelation `json:"InstanceDetails,omitempty" name:"InstanceDetails"`
}

type Layer7Rule struct {

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发类型列表
	ProxyTypeList []*ProxyTypeInfo `json:"ProxyTypeList,omitempty" name:"ProxyTypeList"`

	// 源站列表
	RealServers []*SourceServer `json:"RealServers,omitempty" name:"RealServers"`

	// 规则所属的资源实例
	InstanceDetails []*InstanceRelation `json:"InstanceDetails,omitempty" name:"InstanceDetails"`
}

type ListenerCcThreholdConfig struct {

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 协议（可取值htttps）
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 开关状态（0：关闭，1：开启）
	CCEnable *int64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// cc防护阈值
	CCThreshold *int64 `json:"CCThreshold,omitempty" name:"CCThreshold"`
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

type ModifyDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyDomainUsrNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyL7RulesEdgeRequest struct {
	*tchttp.BaseRequest

	// 大禹子产品代号（edge表示边界防护产品）
	Business *string `json:"Business,omitempty" name:"Business"`

	// 资源ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则
	Rule *L7RuleEntry `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyL7RulesEdgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7RulesEdgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL7RulesEdgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL7RulesEdgeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功码
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyL7RulesEdgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7RulesEdgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ModifyNewDomainRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功码
		Success *SuccessCode `json:"Success,omitempty" name:"Success"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyPacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type NewL7RuleEntry struct {

	// 会话保持时间，单位秒
	KeepTime *uint64 `json:"KeepTime,omitempty" name:"KeepTime"`

	// 负载均衡方式，取值[1(加权轮询)]
	LbType *uint64 `json:"LbType,omitempty" name:"LbType"`

	// 回源列表
	SourceList []*L4RuleSource `json:"SourceList,omitempty" name:"SourceList"`

	// 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
	KeepEnable *uint64 `json:"KeepEnable,omitempty" name:"KeepEnable"`

	// 转发域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发协议，取值[http, https]
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 回源方式，取值[1(域名回源)，2(IP回源)]
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 是否开启Https协议使用Http回源，取值[0(关闭), 1(开启)]，不填写默认是关闭
	HttpsToHttpEnable *uint64 `json:"HttpsToHttpEnable,omitempty" name:"HttpsToHttpEnable"`

	// 规则状态，取值[0(规则配置成功)，1(规则配置生效中)，2(规则配置失败)，3(规则删除生效中)，5(规则删除失败)，6(规则等待配置)，7(规则等待删除)，8(规则待配置证书)]
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// HTTPS协议的CC防护等级
	CCLevel *string `json:"CCLevel,omitempty" name:"CCLevel"`

	// HTTPS协议的CC防护状态，取值[0(关闭), 1(开启)]
	CCEnable *uint64 `json:"CCEnable,omitempty" name:"CCEnable"`

	// HTTPS协议的CC防护阈值
	CCThreshold *uint64 `json:"CCThreshold,omitempty" name:"CCThreshold"`

	// 区域码
	Region *uint64 `json:"Region,omitempty" name:"Region"`

	// 规则描述
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 当证书来源为自有证书时，此字段必须填写证书内容；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

	// 规则ID，当添加新规则时可以不用填写此字段；当修改或者删除规则时需要填写此字段；
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 资源Ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 当证书来源为自有证书时，此字段必须填写证书密钥；(因已不再支持自有证书，此字段已弃用，请不用填写此字段)
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 证书来源，当转发协议为https时必须填，取值[2(腾讯云托管证书)]，当转发协议为http时也可以填0
	CertType *uint64 `json:"CertType,omitempty" name:"CertType"`

	// 接入端口值
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPort *uint64 `json:"VirtualPort,omitempty" name:"VirtualPort"`

	// cc防护状态，取值[0(关闭), 1(开启)]
	CCStatus *uint64 `json:"CCStatus,omitempty" name:"CCStatus"`

	// 当证书来源为腾讯云托管证书时，此字段必须填写托管证书ID
	SSLId *string `json:"SSLId,omitempty" name:"SSLId"`

	// 资源Id
	Id *string `json:"Id,omitempty" name:"Id"`
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
}

type PacketFilterRelation struct {

	// 特征过滤配置
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitempty" name:"PacketFilterConfig"`

	// 特征过滤配置所属的实例
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitempty" name:"InstanceDetailList"`
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

type SwitchWaterPrintConfigRequest struct {
	*tchttp.BaseRequest

	// 资源实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 水印开启/关闭状态，1表示开启；0表示关闭
	OpenStatus *int64 `json:"OpenStatus,omitempty" name:"OpenStatus"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SwitchWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
