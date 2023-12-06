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

package v20201028

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccountVpcInfo struct {
	// VpcId： vpc-xadsafsdasd
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// Vpc所属地区: ap-guangzhou, ap-shanghai
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// Vpc所属账号: 123456789
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// vpc资源名称：testname
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`
}

type AccountVpcInfoOut struct {
	// VpcId： vpc-xadsafsdasd
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Vpc所属地区: ap-guangzhou, ap-shanghai
	Region *string `json:"Region,omitnil" name:"Region"`

	// Vpc所属账号: 123456789
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// vpc资源名称：testname
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`
}

type AccountVpcInfoOutput struct {
	// 关联账户的uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// vpcid
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 地域
	Region *string `json:"Region,omitnil" name:"Region"`
}

// Predefined struct for user
type AddSpecifyPrivateZoneVpcRequestParams struct {
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 本次新增的vpc信息
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 本次新增关联账户vpc信息
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 是否为同步操作
	Sync *bool `json:"Sync,omitnil" name:"Sync"`
}

type AddSpecifyPrivateZoneVpcRequest struct {
	*tchttp.BaseRequest
	
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 本次新增的vpc信息
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 本次新增关联账户vpc信息
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 是否为同步操作
	Sync *bool `json:"Sync,omitnil" name:"Sync"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 本次新增的vpc
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 本次新增的关联账号vpc
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqId *string `json:"UniqId,omitnil" name:"UniqId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Resource *string `json:"Resource,omitnil" name:"Resource"`

	// 日志表名
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 日志总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 日志列表
	DataSet []*AuditLogInfo `json:"DataSet,omitnil" name:"DataSet"`
}

type AuditLogInfo struct {
	// 时间
	Date *string `json:"Date,omitnil" name:"Date"`

	// 操作人uin
	OperatorUin *string `json:"OperatorUin,omitnil" name:"OperatorUin"`

	// 日志内容
	Content *string `json:"Content,omitnil" name:"Content"`
}

// Predefined struct for user
type CreatePrivateDNSAccountRequestParams struct {
	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitnil" name:"Account"`
}

type CreatePrivateDNSAccountRequest struct {
	*tchttp.BaseRequest
	
	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitnil" name:"Account"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitnil" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitnil" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil" name:"TTL"`
}

type CreatePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitnil" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitnil" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil" name:"TTL"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateZoneRecordResponseParams struct {
	// 记录Id
	RecordId *string `json:"RecordId,omitnil" name:"RecordId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 创建私有域的同时，为其打上标签
	TagSet []*TagInfo `json:"TagSet,omitnil" name:"TagSet"`

	// 创建私有域的同时，将其关联至VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED。默认值为ENABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil" name:"DnsForwardStatus"`

	// 创建私有域的同时，将其关联至VPC
	Vpcs []*VpcInfo `json:"Vpcs,omitnil" name:"Vpcs"`

	// 创建私有域同时绑定关联账号的VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 是否CNAME加速：ENABLED，DISABLED，默认值为ENABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil" name:"CnameSpeedupStatus"`
}

type CreatePrivateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 域名，格式必须是标准的TLD
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 创建私有域的同时，为其打上标签
	TagSet []*TagInfo `json:"TagSet,omitnil" name:"TagSet"`

	// 创建私有域的同时，将其关联至VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED。默认值为ENABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil" name:"DnsForwardStatus"`

	// 创建私有域的同时，将其关联至VPC
	Vpcs []*VpcInfo `json:"Vpcs,omitnil" name:"Vpcs"`

	// 创建私有域同时绑定关联账号的VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 是否CNAME加速：ENABLED，DISABLED，默认值为ENABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil" name:"CnameSpeedupStatus"`
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
	// 私有域ID, zone-xxxxxx
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 私有域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Date *string `json:"Date,omitnil" name:"Date"`

	// 值
	Value *int64 `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type DeleteEndPointRequestParams struct {
	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil" name:"EndPointId"`
}

type DeleteEndPointRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点ID
	EndPointId *string `json:"EndPointId,omitnil" name:"EndPointId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DeletePrivateDNSAccountRequestParams struct {
	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitnil" name:"Account"`
}

type DeletePrivateDNSAccountRequest struct {
	*tchttp.BaseRequest
	
	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitnil" name:"Account"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 记录ID（调用DescribePrivateZoneRecordList可获取到RecordId）
	RecordId *string `json:"RecordId,omitnil" name:"RecordId"`

	// 记录ID数组，RecordId 优先
	RecordIdSet []*string `json:"RecordIdSet,omitnil" name:"RecordIdSet"`
}

type DeletePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 记录ID（调用DescribePrivateZoneRecordList可获取到RecordId）
	RecordId *string `json:"RecordId,omitnil" name:"RecordId"`

	// 记录ID数组，RecordId 优先
	RecordIdSet []*string `json:"RecordIdSet,omitnil" name:"RecordIdSet"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 私有域ID数组，ZoneId 优先
	ZoneIdSet []*string `json:"ZoneIdSet,omitnil" name:"ZoneIdSet"`
}

type DeletePrivateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 私有域ID数组，ZoneId 优先
	ZoneIdSet []*string `json:"ZoneIdSet,omitnil" name:"ZoneIdSet"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 本次删除的VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 本次删除的关联账户VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 是否为同步操作
	Sync *bool `json:"Sync,omitnil" name:"Sync"`
}

type DeleteSpecifyPrivateZoneVpcRequest struct {
	*tchttp.BaseRequest
	
	// 私有域id
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 本次删除的VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 本次删除的关联账户VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 是否为同步操作
	Sync *bool `json:"Sync,omitnil" name:"Sync"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 本次删除的VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 本次删除的关联账户的VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqId *string `json:"UniqId,omitnil" name:"UniqId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	AccountUin *string `json:"AccountUin,omitnil" name:"AccountUin"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeAccountVpcListRequest struct {
	*tchttp.BaseRequest
	
	// 关联账号的uin
	AccountUin *string `json:"AccountUin,omitnil" name:"AccountUin"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// VPC 列表
	VpcSet []*AccountVpcInfoOut `json:"VpcSet,omitnil" name:"VpcSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	TimeRangeBegin *string `json:"TimeRangeBegin,omitnil" name:"TimeRangeBegin"`

	// 筛选参数：ZoneId：私有域ID；Domain：私有域；OperatorUin：操作者账号ID
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 请求量统计结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil" name:"TimeRangeEnd"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAuditLogRequest struct {
	*tchttp.BaseRequest
	
	// 请求量统计起始时间
	TimeRangeBegin *string `json:"TimeRangeBegin,omitnil" name:"TimeRangeBegin"`

	// 筛选参数：ZoneId：私有域ID；Domain：私有域；OperatorUin：操作者账号ID
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 请求量统计结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil" name:"TimeRangeEnd"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	Data []*AuditLog `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ZoneTotal *int64 `json:"ZoneTotal,omitnil" name:"ZoneTotal"`

	// 私有域关联VPC数量
	ZoneVpcCount *int64 `json:"ZoneVpcCount,omitnil" name:"ZoneVpcCount"`

	// 历史请求量总数
	RequestTotalCount *int64 `json:"RequestTotalCount,omitnil" name:"RequestTotalCount"`

	// 流量包用量
	FlowUsage []*FlowUsage `json:"FlowUsage,omitnil" name:"FlowUsage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribePrivateDNSAccountListRequestParams struct {
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribePrivateDNSAccountListRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 私有域解析账号列表
	AccountSet []*PrivateDNSAccount `json:"AccountSet,omitnil" name:"AccountSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribePrivateZoneListRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 私有域列表
	PrivateZoneSet []*PrivateZone `json:"PrivateZoneSet,omitnil" name:"PrivateZoneSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 私有域ID: zone-xxxxxx
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 过滤参数（支持使用Value、RecordType过滤）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大200，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribePrivateZoneRecordListRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID: zone-xxxxxx
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 过滤参数（支持使用Value、RecordType过滤）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页限制数目， 最大200，默认20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 解析记录列表
	RecordSet []*PrivateZoneRecord `json:"RecordSet,omitnil" name:"RecordSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 域名，格式必须是标准的TLD
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`
}

type DescribePrivateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 域名，格式必须是标准的TLD
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`
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
	PrivateZone *PrivateZone `json:"PrivateZone,omitnil" name:"PrivateZone"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ServiceStatus *string `json:"ServiceStatus,omitnil" name:"ServiceStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	TldQuota *TldQuota `json:"TldQuota,omitnil" name:"TldQuota"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeRequestDataRequestParams struct {
	// 请求量统计起始时间，格式：2020-11-22 00:00:00
	TimeRangeBegin *string `json:"TimeRangeBegin,omitnil" name:"TimeRangeBegin"`

	// 筛选参数：
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 请求量统计结束时间，格式：2020-11-22 23:59:59
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil" name:"TimeRangeEnd"`
}

type DescribeRequestDataRequest struct {
	*tchttp.BaseRequest
	
	// 请求量统计起始时间，格式：2020-11-22 00:00:00
	TimeRangeBegin *string `json:"TimeRangeBegin,omitnil" name:"TimeRangeBegin"`

	// 筛选参数：
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 请求量统计结束时间，格式：2020-11-22 23:59:59
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil" name:"TimeRangeEnd"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRequestDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRequestDataResponseParams struct {
	// 请求量统计表
	Data []*MetricData `json:"Data,omitnil" name:"Data"`

	// 请求量单位时间: Day：天，Hour：小时
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type Filter struct {
	// 参数名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 参数值数组
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FlowUsage struct {
	// 流量包类型：ZONE 私有域；TRAFFIC 解析流量包
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 流量包总额度
	TotalQuantity *int64 `json:"TotalQuantity,omitnil" name:"TotalQuantity"`

	// 流量包可用额度
	AvailableQuantity *int64 `json:"AvailableQuantity,omitnil" name:"AvailableQuantity"`
}

type MetricData struct {
	// 资源描述
	Resource *string `json:"Resource,omitnil" name:"Resource"`

	// 表名
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 表数据
	DataSet []*DatePoint `json:"DataSet,omitnil" name:"DataSet"`

	// 查询范围内的请求总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricCount *int64 `json:"MetricCount,omitnil" name:"MetricCount"`
}

// Predefined struct for user
type ModifyPrivateZoneRecordRequestParams struct {
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 记录ID
	RecordId *string `json:"RecordId,omitnil" name:"RecordId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitnil" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitnil" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil" name:"TTL"`
}

type ModifyPrivateZoneRecordRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 记录ID
	RecordId *string `json:"RecordId,omitnil" name:"RecordId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitnil" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitnil" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitnil" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil" name:"TTL"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateZoneRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil" name:"DnsForwardStatus"`

	// 是否开启CNAME加速：ENABLED， DISABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil" name:"CnameSpeedupStatus"`
}

type ModifyPrivateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil" name:"DnsForwardStatus"`

	// 是否开启CNAME加速：ENABLED， DISABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil" name:"CnameSpeedupStatus"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 私有域关联的全部VPC列表
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 私有域账号关联的全部VPC列表
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`
}

type ModifyPrivateZoneVpcRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 私有域关联的全部VPC列表
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 私有域账号关联的全部VPC列表
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`
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
	// 私有域ID, zone-xxxxxx
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 解析域关联的VPC列表
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 私有域账号关联的全部VPC列表
	AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 解析记录ID列表
	RecordIds []*int64 `json:"RecordIds,omitnil" name:"RecordIds"`

	// enabled：生效，disabled：失效
	Status *string `json:"Status,omitnil" name:"Status"`
}

type ModifyRecordsStatusRequest struct {
	*tchttp.BaseRequest
	
	// 私有域ID
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 解析记录ID列表
	RecordIds []*int64 `json:"RecordIds,omitnil" name:"RecordIds"`

	// enabled：生效，disabled：失效
	Status *string `json:"Status,omitnil" name:"Status"`
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
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 解析记录ID列表
	RecordIds []*int64 `json:"RecordIds,omitnil" name:"RecordIds"`

	// enabled：生效，disabled：失效
	Status *string `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 主账号名称
	Account *string `json:"Account,omitnil" name:"Account"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitnil" name:"Nickname"`
}

type PrivateZone struct {
	// 私有域id: zone-xxxxxxxx
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 域名所有者uin
	OwnerUin *int64 `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// 私有域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 修改时间
	UpdatedOn *string `json:"UpdatedOn,omitnil" name:"UpdatedOn"`

	// 记录数
	RecordCount *int64 `json:"RecordCount,omitnil" name:"RecordCount"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 绑定的Vpc列表
	VpcSet []*VpcInfo `json:"VpcSet,omitnil" name:"VpcSet"`

	// 私有域状态：正常解析：ENABLED, 暂停解析：SUSPEND, 锁定：FROZEN
	Status *string `json:"Status,omitnil" name:"Status"`

	// 域名递归解析状态：开通：ENABLED, 关闭，DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitnil" name:"DnsForwardStatus"`

	// 标签键值对集合
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`

	// 绑定的关联账号的vpc列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitnil" name:"AccountVpcSet"`

	// 是否自定义TLD
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCustomTld *bool `json:"IsCustomTld,omitnil" name:"IsCustomTld"`

	// CNAME加速状态：开通：ENABLED, 关闭，DISABLED
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitnil" name:"CnameSpeedupStatus"`

	// 转发规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardRuleName *string `json:"ForwardRuleName,omitnil" name:"ForwardRuleName"`

	// 转发规则类型：云上到云下，DOWN；云下到云上，UP，目前只支持DOWN
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardRuleType *string `json:"ForwardRuleType,omitnil" name:"ForwardRuleType"`

	// 转发的地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardAddress *string `json:"ForwardAddress,omitnil" name:"ForwardAddress"`

	// 终端节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPointName *string `json:"EndPointName,omitnil" name:"EndPointName"`
}

type PrivateZoneRecord struct {
	// 记录id
	RecordId *string `json:"RecordId,omitnil" name:"RecordId"`

	// 私有域id: zone-xxxxxxxx
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 子域名
	SubDomain *string `json:"SubDomain,omitnil" name:"SubDomain"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitnil" name:"RecordType"`

	// 记录值
	RecordValue *string `json:"RecordValue,omitnil" name:"RecordValue"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitnil" name:"TTL"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *int64 `json:"MX,omitnil" name:"MX"`

	// 记录状态：ENABLED
	Status *string `json:"Status,omitnil" name:"Status"`

	// 记录权重，值为1-100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// 记录创建时间
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 记录更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil" name:"UpdatedOn"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 0暂停，1启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *uint64 `json:"Enabled,omitnil" name:"Enabled"`
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
	ServiceStatus *string `json:"ServiceStatus,omitnil" name:"ServiceStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type TldQuota struct {
	// 总共额度
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 已使用额度
	Used *int64 `json:"Used,omitnil" name:"Used"`

	// 库存
	Stock *int64 `json:"Stock,omitnil" name:"Stock"`

	// 用户限额
	Quota *int64 `json:"Quota,omitnil" name:"Quota"`
}

type VpcInfo struct {
	// VpcId： vpc-xadsafsdasd
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// Vpc所属地区: ap-guangzhou, ap-shanghai
	Region *string `json:"Region,omitnil" name:"Region"`
}