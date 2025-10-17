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

package v20201028

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccountVpcInfo struct {
	// VpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Vpc所属地区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Vpc所属账号
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// vpc资源名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`
}

type AccountVpcInfoOut struct {
	// VpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Vpc所属地区: ap-guangzhou, ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Vpc所属账号: 123456789
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// vpc资源名称：testname
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`
}

type AccountVpcInfoOutput struct {
	// 关联账户的uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// vpcid
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type AddSpecifyPrivateZoneVpcRequestParams struct {
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 本次新增的vpc信息
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 本次新增关联账户vpc信息
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 是否为同步操作
	Sync *bool `json:"Sync,omitnil,omitempty" name:"Sync"`
}

type AddSpecifyPrivateZoneVpcRequest struct {
	*tchttp.BaseRequest
	
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 本次新增的vpc信息
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 本次新增关联账户vpc信息
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 是否为同步操作
	Sync *bool `json:"Sync,omitnil,omitempty" name:"Sync"`
}

func (r *AddSpecifyPrivateZoneVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpecifyPrivateZoneVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "VpcSet")
	delete(f, "AccountVpcSet")
	delete(f, "Sync")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSpecifyPrivateZoneVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpecifyPrivateZoneVpcResponseParams struct {
	// zone id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 本次新增的vpc
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 本次新增的关联账号vpc
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 唯一id
	UniqId *string `json:"UniqId,omitnil,omitempty" name:"UniqId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddSpecifyPrivateZoneVpcResponse struct {
	*tchttp.BaseResponse
	Response *AddSpecifyPrivateZoneVpcResponseParams `json:"Response"`
}

func (r *AddSpecifyPrivateZoneVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpecifyPrivateZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditLog struct {
	// 日志类型
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 日志表名
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 日志总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志列表
	DataSet []*AuditLogInfo `json:"DataSet,omitnil,omitempty" name:"DataSet"`
}

type AuditLogInfo struct {
	// 时间
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 操作人uin
	OperatorUin *string `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// 日志内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type CreateExtendEndpointRequestParams struct {

}

type CreateExtendEndpointRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateExtendEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExtendEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExtendEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExtendEndpointResponseParams struct {
	// 终端节点id
	EndpointId *string `json:"EndpointId,omitnil,omitempty" name:"EndpointId"`

	// 终端节点名称
	EndpointName *string `json:"EndpointName,omitnil,omitempty" name:"EndpointName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExtendEndpointResponse struct {
	*tchttp.BaseResponse
	Response *CreateExtendEndpointResponseParams `json:"Response"`
}

func (r *CreateExtendEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExtendEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardRuleRequestParams struct {
	// 转发规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 转发规则类型：云上到云下DOWN，云下到云上UP
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 私有域ID，可在私有域列表页面查看
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

type CreateForwardRuleRequest struct {
	*tchttp.BaseRequest
	
	// 转发规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 转发规则类型：云上到云下DOWN，云下到云上UP
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 私有域ID，可在私有域列表页面查看
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

func (r *CreateForwardRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "RuleType")
	delete(f, "ZoneId")
	delete(f, "EndPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateForwardRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateForwardRuleResponseParams struct {
	// 转发规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转发规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 转发规则类型
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateForwardRuleResponseParams `json:"Response"`
}

func (r *CreateForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateForwardRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInboundEndpointRequestParams struct {
	// 名称
	EndpointName *string `json:"EndpointName,omitnil,omitempty" name:"EndpointName"`

	// 地域
	EndpointRegion *string `json:"EndpointRegion,omitnil,omitempty" name:"EndpointRegion"`

	// vpcid
	EndpointVpc *string `json:"EndpointVpc,omitnil,omitempty" name:"EndpointVpc"`

	// 子网信息
	SubnetIp []*SubnetIpInfo `json:"SubnetIp,omitnil,omitempty" name:"SubnetIp"`
}

type CreateInboundEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 名称
	EndpointName *string `json:"EndpointName,omitnil,omitempty" name:"EndpointName"`

	// 地域
	EndpointRegion *string `json:"EndpointRegion,omitnil,omitempty" name:"EndpointRegion"`

	// vpcid
	EndpointVpc *string `json:"EndpointVpc,omitnil,omitempty" name:"EndpointVpc"`

	// 子网信息
	SubnetIp []*SubnetIpInfo `json:"SubnetIp,omitnil,omitempty" name:"SubnetIp"`
}

func (r *CreateInboundEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInboundEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndpointName")
	delete(f, "EndpointRegion")
	delete(f, "EndpointVpc")
	delete(f, "SubnetIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInboundEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInboundEndpointResponseParams struct {
	// 终端节点ID
	EndpointId *string `json:"EndpointId,omitnil,omitempty" name:"EndpointId"`

	// 名称
	EndpointName *string `json:"EndpointName,omitnil,omitempty" name:"EndpointName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInboundEndpointResponse struct {
	*tchttp.BaseResponse
	Response *CreateInboundEndpointResponseParams `json:"Response"`
}

func (r *CreateInboundEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInboundEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateDNSAccountRequestParams struct {
	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

type CreatePrivateDNSAccountRequest struct {
	*tchttp.BaseRequest
	
	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

func (r *CreatePrivateDNSAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateDNSAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateDNSAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateDNSAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrivateDNSAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrivateDNSAccountResponseParams `json:"Response"`
}

func (r *CreatePrivateDNSAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateDNSAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateZoneRecordRequestParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreatePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreatePrivateZoneRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordType")
	delete(f, "SubDomain")
	delete(f, "RecordValue")
	delete(f, "Weight")
	delete(f, "MX")
	delete(f, "TTL")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateZoneRecordResponseParams struct {
	// 记录Id
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrivateZoneRecordResponseParams `json:"Response"`
}

func (r *CreatePrivateZoneRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateZoneRequestParams struct {
	// 域名，格式必须是标准的TLD
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 创建私有域的同时，为其打上标签
	TagSet []*TagInfo `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 创建私有域的同时，将其关联至VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED。默认值为ENABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil,omitempty" name:"DnsForwardStatus"`

	// 创建私有域的同时，将其关联至VPC
	//
	// Deprecated: Vpcs is deprecated.
	Vpcs []*VpcInfo `json:"Vpcs,omitnil,omitempty" name:"Vpcs"`

	// 创建私有域同时绑定关联账号的VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 是否CNAME加速：ENABLED，DISABLED，默认值为ENABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil,omitempty" name:"CnameSpeedupStatus"`
}

type CreatePrivateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 域名，格式必须是标准的TLD
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 创建私有域的同时，为其打上标签
	TagSet []*TagInfo `json:"TagSet,omitnil,omitempty" name:"TagSet"`

	// 创建私有域的同时，将其关联至VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED。默认值为ENABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil,omitempty" name:"DnsForwardStatus"`

	// 创建私有域的同时，将其关联至VPC
	Vpcs []*VpcInfo `json:"Vpcs,omitnil,omitempty" name:"Vpcs"`

	// 创建私有域同时绑定关联账号的VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 是否CNAME加速：ENABLED，DISABLED，默认值为ENABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil,omitempty" name:"CnameSpeedupStatus"`
}

func (r *CreatePrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "TagSet")
	delete(f, "VpcSet")
	delete(f, "Remark")
	delete(f, "DnsForwardStatus")
	delete(f, "Vpcs")
	delete(f, "AccountVpcSet")
	delete(f, "CnameSpeedupStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateZoneResponseParams struct {
	// 私有域ID, zone-12sa5ce78
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 私有域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrivateZoneResponseParams `json:"Response"`
}

func (r *CreatePrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatePoint struct {
	// 时间
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type DeleteEndPointRequestParams struct {
	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

type DeleteEndPointRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

func (r *DeleteEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEndPointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEndPointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEndPointResponseParams `json:"Response"`
}

func (r *DeleteEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardRuleRequestParams struct {
	// 转发规则ID数组
	RuleIdSet []*string `json:"RuleIdSet,omitnil,omitempty" name:"RuleIdSet"`
}

type DeleteForwardRuleRequest struct {
	*tchttp.BaseRequest
	
	// 转发规则ID数组
	RuleIdSet []*string `json:"RuleIdSet,omitnil,omitempty" name:"RuleIdSet"`
}

func (r *DeleteForwardRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteForwardRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteForwardRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteForwardRuleResponseParams `json:"Response"`
}

func (r *DeleteForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteForwardRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInboundEndpointRequestParams struct {
	// 终端节点ID
	EndpointId *string `json:"EndpointId,omitnil,omitempty" name:"EndpointId"`
}

type DeleteInboundEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点ID
	EndpointId *string `json:"EndpointId,omitnil,omitempty" name:"EndpointId"`
}

func (r *DeleteInboundEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInboundEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndpointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInboundEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInboundEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInboundEndpointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInboundEndpointResponseParams `json:"Response"`
}

func (r *DeleteInboundEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInboundEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateDNSAccountRequestParams struct {
	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

type DeletePrivateDNSAccountRequest struct {
	*tchttp.BaseRequest
	
	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitnil,omitempty" name:"Account"`
}

func (r *DeletePrivateDNSAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateDNSAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateDNSAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateDNSAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrivateDNSAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrivateDNSAccountResponseParams `json:"Response"`
}

func (r *DeletePrivateDNSAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateDNSAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateZoneRecordRequestParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 记录ID（调用DescribePrivateZoneRecordList可获取到RecordId）
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 记录ID数组，RecordId 优先
	RecordIdSet []*string `json:"RecordIdSet,omitnil,omitempty" name:"RecordIdSet"`
}

type DeletePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 记录ID（调用DescribePrivateZoneRecordList可获取到RecordId）
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 记录ID数组，RecordId 优先
	RecordIdSet []*string `json:"RecordIdSet,omitnil,omitempty" name:"RecordIdSet"`
}

func (r *DeletePrivateZoneRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordId")
	delete(f, "RecordIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateZoneRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrivateZoneRecordResponseParams `json:"Response"`
}

func (r *DeletePrivateZoneRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateZoneRequestParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 私有域ID数组，ZoneId 优先
	ZoneIdSet []*string `json:"ZoneIdSet,omitnil,omitempty" name:"ZoneIdSet"`
}

type DeletePrivateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 私有域ID数组，ZoneId 优先
	ZoneIdSet []*string `json:"ZoneIdSet,omitnil,omitempty" name:"ZoneIdSet"`
}

func (r *DeletePrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ZoneIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateZoneResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrivateZoneResponseParams `json:"Response"`
}

func (r *DeletePrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpecifyPrivateZoneVpcRequestParams struct {
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 本次删除的VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 本次删除的关联账户VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 是否为同步操作
	Sync *bool `json:"Sync,omitnil,omitempty" name:"Sync"`
}

type DeleteSpecifyPrivateZoneVpcRequest struct {
	*tchttp.BaseRequest
	
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 本次删除的VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 本次删除的关联账户VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 是否为同步操作
	Sync *bool `json:"Sync,omitnil,omitempty" name:"Sync"`
}

func (r *DeleteSpecifyPrivateZoneVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpecifyPrivateZoneVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "VpcSet")
	delete(f, "AccountVpcSet")
	delete(f, "Sync")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSpecifyPrivateZoneVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpecifyPrivateZoneVpcResponseParams struct {
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 本次删除的VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 本次删除的关联账户的VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 唯一id
	UniqId *string `json:"UniqId,omitnil,omitempty" name:"UniqId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSpecifyPrivateZoneVpcResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSpecifyPrivateZoneVpcResponseParams `json:"Response"`
}

func (r *DeleteSpecifyPrivateZoneVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpecifyPrivateZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountVpcListRequestParams struct {
	// 关联账号的uin
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAccountVpcListRequest struct {
	*tchttp.BaseRequest
	
	// 关联账号的uin
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAccountVpcListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountVpcListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountVpcListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountVpcListResponseParams struct {
	// VPC数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// VPC 列表
	VpcSet []*AccountVpcInfoOut `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountVpcListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountVpcListResponseParams `json:"Response"`
}

func (r *DescribeAccountVpcListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountVpcListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogRequestParams struct {
	// 请求量统计起始时间
	TimeRangeBegin *string `json:"TimeRangeBegin,omitnil,omitempty" name:"TimeRangeBegin"`

	// 筛选参数：ZoneId：私有域ID；Domain：私有域；OperatorUin：操作者账号ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 请求量统计结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAuditLogRequest struct {
	*tchttp.BaseRequest
	
	// 请求量统计起始时间
	TimeRangeBegin *string `json:"TimeRangeBegin,omitnil,omitempty" name:"TimeRangeBegin"`

	// 筛选参数：ZoneId：私有域ID；Domain：私有域；OperatorUin：操作者账号ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 请求量统计结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAuditLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRangeBegin")
	delete(f, "Filters")
	delete(f, "TimeRangeEnd")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogResponseParams struct {
	// 操作日志列表
	Data []*AuditLog `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogResponseParams `json:"Response"`
}

func (r *DescribeAuditLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDashboardRequestParams struct {

}

type DescribeDashboardRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDashboardResponseParams struct {
	// 私有域解析总数
	ZoneTotal *int64 `json:"ZoneTotal,omitnil,omitempty" name:"ZoneTotal"`

	// 私有域关联VPC数量
	ZoneVpcCount *int64 `json:"ZoneVpcCount,omitnil,omitempty" name:"ZoneVpcCount"`

	// 历史请求量总数
	RequestTotalCount *int64 `json:"RequestTotalCount,omitnil,omitempty" name:"RequestTotalCount"`

	// 流量包用量
	FlowUsage []*FlowUsage `json:"FlowUsage,omitnil,omitempty" name:"FlowUsage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDashboardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDashboardResponseParams `json:"Response"`
}

func (r *DescribeDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtendEndpointListRequestParams struct {
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数，支持EndpointName,EndpointId
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeExtendEndpointListRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数，支持EndpointName,EndpointId
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeExtendEndpointListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendEndpointListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtendEndpointListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtendEndpointListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtendEndpointListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtendEndpointListResponseParams `json:"Response"`
}

func (r *DescribeExtendEndpointListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendEndpointListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForwardRuleListRequestParams struct {
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeForwardRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeForwardRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeForwardRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForwardRuleListResponseParams struct {
	// 私有域数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 私有域列表
	ForwardRuleSet []*ForwardRule `json:"ForwardRuleSet,omitnil,omitempty" name:"ForwardRuleSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeForwardRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeForwardRuleListResponseParams `json:"Response"`
}

func (r *DescribeForwardRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForwardRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInboundEndpointListRequestParams struct {
	// 分页偏移量，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数，支持EndPointName，EndpointName，EndpointId
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeInboundEndpointListRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数，支持EndPointName，EndpointName，EndpointId
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeInboundEndpointListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInboundEndpointListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInboundEndpointListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInboundEndpointListResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 终端节点信息
	InboundEndpointSet []*InboundEndpointSet `json:"InboundEndpointSet,omitnil,omitempty" name:"InboundEndpointSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInboundEndpointListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInboundEndpointListResponseParams `json:"Response"`
}

func (r *DescribeInboundEndpointListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInboundEndpointListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateDNSAccountListRequestParams struct {
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrivateDNSAccountListRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrivateDNSAccountListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateDNSAccountListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateDNSAccountListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateDNSAccountListResponseParams struct {
	// 私有域解析账号数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 私有域解析账号列表
	AccountSet []*PrivateDNSAccount `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrivateDNSAccountListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateDNSAccountListResponseParams `json:"Response"`
}

func (r *DescribePrivateDNSAccountListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateDNSAccountListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateZoneListRequestParams struct {
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrivateZoneListRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrivateZoneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateZoneListResponseParams struct {
	// 私有域数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 私有域列表
	PrivateZoneSet []*PrivateZone `json:"PrivateZoneSet,omitnil,omitempty" name:"PrivateZoneSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrivateZoneListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateZoneListResponseParams `json:"Response"`
}

func (r *DescribePrivateZoneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateZoneRecordListRequestParams struct {
	// 私有域ID: zone-12c5a6e8
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤参数（支持使用Value、RecordType过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大200，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrivateZoneRecordListRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID: zone-12c5a6e8
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过滤参数（支持使用Value、RecordType过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目， 最大200，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrivateZoneRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateZoneRecordListResponseParams struct {
	// 解析记录数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 解析记录列表
	RecordSet []*PrivateZoneRecord `json:"RecordSet,omitnil,omitempty" name:"RecordSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrivateZoneRecordListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateZoneRecordListResponseParams `json:"Response"`
}

func (r *DescribePrivateZoneRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateZoneRequestParams struct {
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type DescribePrivateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *DescribePrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateZoneResponseParams struct {
	// 私有域详情
	PrivateZone *PrivateZone `json:"PrivateZone,omitnil,omitempty" name:"PrivateZone"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateZoneResponseParams `json:"Response"`
}

func (r *DescribePrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateZoneServiceRequestParams struct {

}

type DescribePrivateZoneServiceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePrivateZoneServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateZoneServiceResponseParams struct {
	// 私有域解析服务开通状态。ENABLED已开通，DISABLED未开通
	ServiceStatus *string `json:"ServiceStatus,omitnil,omitempty" name:"ServiceStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrivateZoneServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateZoneServiceResponseParams `json:"Response"`
}

func (r *DescribePrivateZoneServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaUsageRequestParams struct {

}

type DescribeQuotaUsageRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeQuotaUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotaUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaUsageResponseParams struct {
	// Tld额度使用情况
	TldQuota *TldQuota `json:"TldQuota,omitnil,omitempty" name:"TldQuota"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQuotaUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotaUsageResponseParams `json:"Response"`
}

func (r *DescribeQuotaUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordRequestParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 记录ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`
}

type DescribeRecordRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 记录ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`
}

func (r *DescribeRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordResponseParams struct {
	// 记录信息
	RecordInfo *RecordInfo `json:"RecordInfo,omitnil,omitempty" name:"RecordInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordResponseParams `json:"Response"`
}

func (r *DescribeRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRequestDataRequestParams struct {
	// 请求量统计起始时间，格式：2020-11-22 00:00:00
	TimeRangeBegin *string `json:"TimeRangeBegin,omitnil,omitempty" name:"TimeRangeBegin"`

	// 筛选参数：
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 请求量统计结束时间，格式：2020-11-22 23:59:59
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// 是否导出：true导出，false不导出
	Export *bool `json:"Export,omitnil,omitempty" name:"Export"`
}

type DescribeRequestDataRequest struct {
	*tchttp.BaseRequest
	
	// 请求量统计起始时间，格式：2020-11-22 00:00:00
	TimeRangeBegin *string `json:"TimeRangeBegin,omitnil,omitempty" name:"TimeRangeBegin"`

	// 筛选参数：
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 请求量统计结束时间，格式：2020-11-22 23:59:59
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// 是否导出：true导出，false不导出
	Export *bool `json:"Export,omitnil,omitempty" name:"Export"`
}

func (r *DescribeRequestDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRequestDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRangeBegin")
	delete(f, "Filters")
	delete(f, "TimeRangeEnd")
	delete(f, "Export")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRequestDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRequestDataResponseParams struct {
	// 请求量统计表
	Data []*MetricData `json:"Data,omitnil,omitempty" name:"Data"`

	// 请求量单位时间: Day：天，Hour：小时
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 导出数据下载地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRequestDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRequestDataResponseParams `json:"Response"`
}

func (r *DescribeRequestDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRequestDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndPointServiceInfo struct {
	// ip
	EndPointVip *string `json:"EndPointVip,omitnil,omitempty" name:"EndPointVip"`

	// 子网id
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 网络侧状态，0 可用，1 待接受，3 拒绝
	EndPointState *uint64 `json:"EndPointState,omitnil,omitempty" name:"EndPointState"`

	// 状态，1表示可用，0表示删除
	EndPointStatus *uint64 `json:"EndPointStatus,omitnil,omitempty" name:"EndPointStatus"`

	// 备注信息
	EndPointRemark *string `json:"EndPointRemark,omitnil,omitempty" name:"EndPointRemark"`

	// 网络侧隔离状态，1：已隔离，0：未隔离
	EndPointIsolateFlag *uint64 `json:"EndPointIsolateFlag,omitnil,omitempty" name:"EndPointIsolateFlag"`
}

type Filter struct {
	// 参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值数组
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FlowUsage struct {
	// 流量包类型：ZONE 私有域；TRAFFIC 解析流量包
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 流量包总额度
	TotalQuantity *int64 `json:"TotalQuantity,omitnil,omitempty" name:"TotalQuantity"`

	// 流量包可用额度
	AvailableQuantity *int64 `json:"AvailableQuantity,omitnil,omitempty" name:"AvailableQuantity"`
}

type ForwardRule struct {
	// 私有域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 转发规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则id
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转发规则类型：云上到云下DOWN、云下到云上DOWN
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 终端节点名称
	EndPointName *string `json:"EndPointName,omitnil,omitempty" name:"EndPointName"`

	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// 转发地址
	ForwardAddress []*string `json:"ForwardAddress,omitnil,omitempty" name:"ForwardAddress"`

	// 私有域绑定的vpc列表
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 绑定的私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 标签
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type InboundEndpointSet struct {
	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`

	// 名称
	EndPointName *string `json:"EndPointName,omitnil,omitempty" name:"EndPointName"`

	// vpcid
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 终端节点信息
	EndPointService []*EndPointServiceInfo `json:"EndPointService,omitnil,omitempty" name:"EndPointService"`
}

type MetricData struct {
	// 资源描述
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 表名
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 表数据
	DataSet []*DatePoint `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 查询范围内的请求总量
	MetricCount *int64 `json:"MetricCount,omitnil,omitempty" name:"MetricCount"`
}

// Predefined struct for user
type ModifyForwardRuleRequestParams struct {
	// 转发规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转发规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

type ModifyForwardRuleRequest struct {
	*tchttp.BaseRequest
	
	// 转发规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转发规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil,omitempty" name:"EndPointId"`
}

func (r *ModifyForwardRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "EndPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyForwardRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyForwardRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyForwardRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyForwardRuleResponseParams `json:"Response"`
}

func (r *ModifyForwardRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyForwardRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInboundEndpointRequestParams struct {
	// 终端节点ID
	EndpointId *string `json:"EndpointId,omitnil,omitempty" name:"EndpointId"`

	// 终端节点名称
	EndpointName *string `json:"EndpointName,omitnil,omitempty" name:"EndpointName"`
}

type ModifyInboundEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点ID
	EndpointId *string `json:"EndpointId,omitnil,omitempty" name:"EndpointId"`

	// 终端节点名称
	EndpointName *string `json:"EndpointName,omitnil,omitempty" name:"EndpointName"`
}

func (r *ModifyInboundEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInboundEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndpointId")
	delete(f, "EndpointName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInboundEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInboundEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInboundEndpointResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInboundEndpointResponseParams `json:"Response"`
}

func (r *ModifyInboundEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInboundEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateZoneRecordRequestParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 记录ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyPrivateZoneRecordRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 记录ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyPrivateZoneRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordId")
	delete(f, "RecordType")
	delete(f, "SubDomain")
	delete(f, "RecordValue")
	delete(f, "Weight")
	delete(f, "MX")
	delete(f, "TTL")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateZoneRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateZoneRecordResponseParams `json:"Response"`
}

func (r *ModifyPrivateZoneRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateZoneRequestParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil,omitempty" name:"DnsForwardStatus"`

	// 是否开启CNAME加速：ENABLED， DISABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil,omitempty" name:"CnameSpeedupStatus"`
}

type ModifyPrivateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil,omitempty" name:"DnsForwardStatus"`

	// 是否开启CNAME加速：ENABLED， DISABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil,omitempty" name:"CnameSpeedupStatus"`
}

func (r *ModifyPrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Remark")
	delete(f, "DnsForwardStatus")
	delete(f, "CnameSpeedupStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateZoneResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateZoneResponseParams `json:"Response"`
}

func (r *ModifyPrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateZoneVpcRequestParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 私有域关联的全部VPC列表
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 私有域账号关联的全部VPC列表
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`
}

type ModifyPrivateZoneVpcRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 私有域关联的全部VPC列表
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 私有域账号关联的全部VPC列表
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`
}

func (r *ModifyPrivateZoneVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "VpcSet")
	delete(f, "AccountVpcSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateZoneVpcResponseParams struct {
	// 私有域ID, zone-12e45ds6
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 解析域关联的VPC列表
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 私有域账号关联的全部VPC列表
	AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrivateZoneVpcResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateZoneVpcResponseParams `json:"Response"`
}

func (r *ModifyPrivateZoneVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordsStatusRequestParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 解析记录ID列表
	RecordIds []*int64 `json:"RecordIds,omitnil,omitempty" name:"RecordIds"`

	// enabled：生效，disabled：失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyRecordsStatusRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 解析记录ID列表
	RecordIds []*int64 `json:"RecordIds,omitnil,omitempty" name:"RecordIds"`

	// enabled：生效，disabled：失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyRecordsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordIds")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordsStatusResponseParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 解析记录ID列表
	RecordIds []*int64 `json:"RecordIds,omitnil,omitempty" name:"RecordIds"`

	// enabled：生效，disabled：失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRecordsStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordsStatusResponseParams `json:"Response"`
}

func (r *ModifyRecordsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrivateDNSAccount struct {
	// 主账号Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 主账号名称
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`
}

type PrivateZone struct {
	// 私有域id: zone-xxxxxxxx
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 域名所有者uin
	OwnerUin *int64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 私有域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 修改时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 记录数
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 绑定的Vpc列表
	VpcSet []*VpcInfo `json:"VpcSet,omitnil,omitempty" name:"VpcSet"`

	// 私有域绑定VPC状态，未关联vpc：SUSPEND，已关联VPC：ENABLED
	// ，关联VPC失败：FAILED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名递归解析状态：开通：ENABLED, 关闭，DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil,omitempty" name:"DnsForwardStatus"`

	// 标签键值对集合
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 绑定的关联账号的vpc列表
	AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitnil,omitempty" name:"AccountVpcSet"`

	// 是否自定义TLD
	IsCustomTld *bool `json:"IsCustomTld,omitnil,omitempty" name:"IsCustomTld"`

	// CNAME加速状态：开通：ENABLED, 关闭，DISABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil,omitempty" name:"CnameSpeedupStatus"`

	// 转发规则名称
	ForwardRuleName *string `json:"ForwardRuleName,omitnil,omitempty" name:"ForwardRuleName"`

	// 转发规则类型：云上到云下，DOWN；云下到云上，UP，目前只支持DOWN
	ForwardRuleType *string `json:"ForwardRuleType,omitnil,omitempty" name:"ForwardRuleType"`

	// 转发的地址
	ForwardAddress *string `json:"ForwardAddress,omitnil,omitempty" name:"ForwardAddress"`

	// 终端节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPointName *string `json:"EndPointName,omitnil,omitempty" name:"EndPointName"`

	// 已删除的vpc
	DeletedVpcSet []*VpcInfo `json:"DeletedVpcSet,omitnil,omitempty" name:"DeletedVpcSet"`
}

type PrivateZoneRecord struct {
	// 记录id
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 私有域id: zone-xxxxxxxx
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 子域名
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录值
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录状态：ENABLED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 记录权重，值为1-100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 记录创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 记录更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 0暂停，1启用
	Enabled *uint64 `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

// Predefined struct for user
type QueryAsyncBindVpcStatusRequestParams struct {
	// 唯一ID
	UniqId *string `json:"UniqId,omitnil,omitempty" name:"UniqId"`
}

type QueryAsyncBindVpcStatusRequest struct {
	*tchttp.BaseRequest
	
	// 唯一ID
	UniqId *string `json:"UniqId,omitnil,omitempty" name:"UniqId"`
}

func (r *QueryAsyncBindVpcStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAsyncBindVpcStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UniqId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAsyncBindVpcStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryAsyncBindVpcStatusResponseParams struct {
	// processing 处理中，success 执行成功，
	// failed 执行失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryAsyncBindVpcStatusResponse struct {
	*tchttp.BaseResponse
	Response *QueryAsyncBindVpcStatusResponseParams `json:"Response"`
}

func (r *QueryAsyncBindVpcStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAsyncBindVpcStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordInfo struct {
	// 记录id
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 私有域id: zone-xxxxxxxx
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 子域名
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录值
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 记录创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 记录更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 0暂停，1启用
	Enabled *uint64 `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type SubnetIpInfo struct {
	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// ip
	SubnetVip *string `json:"SubnetVip,omitnil,omitempty" name:"SubnetVip"`
}

// Predefined struct for user
type SubscribePrivateZoneServiceRequestParams struct {

}

type SubscribePrivateZoneServiceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *SubscribePrivateZoneServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubscribePrivateZoneServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubscribePrivateZoneServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubscribePrivateZoneServiceResponseParams struct {
	// 私有域解析服务开通状态
	ServiceStatus *string `json:"ServiceStatus,omitnil,omitempty" name:"ServiceStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubscribePrivateZoneServiceResponse struct {
	*tchttp.BaseResponse
	Response *SubscribePrivateZoneServiceResponseParams `json:"Response"`
}

func (r *SubscribePrivateZoneServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubscribePrivateZoneServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TldQuota struct {
	// 总共额度
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 已使用额度
	Used *int64 `json:"Used,omitnil,omitempty" name:"Used"`

	// 库存
	Stock *int64 `json:"Stock,omitnil,omitempty" name:"Stock"`

	// 用户限额
	Quota *int64 `json:"Quota,omitnil,omitempty" name:"Quota"`
}

type VpcInfo struct {
	// VpcId
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Vpc所属地区: ap-guangzhou, ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}