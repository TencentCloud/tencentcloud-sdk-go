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

package v20191112

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AlarmPolicy struct {
	// 用户账号
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 告警事件
	Event *string `json:"Event,omitempty" name:"Event"`

	// 告警阈值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 告警策略是否生效，0：停用，1：启用
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 在这个时间后才允许发送告警
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 在这个时间前才允许发送告警
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

// Predefined struct for user
type DescribeHSMBySubnetIdRequestParams struct {
	// Subnet标识符
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DescribeHSMBySubnetIdRequest struct {
	*tchttp.BaseRequest
	
	// Subnet标识符
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DescribeHSMBySubnetIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHSMBySubnetIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHSMBySubnetIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHSMBySubnetIdResponseParams struct {
	// HSM数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 作为查询条件的SubnetId
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHSMBySubnetIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHSMBySubnetIdResponseParams `json:"Response"`
}

func (r *DescribeHSMBySubnetIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHSMBySubnetIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHSMByVpcIdRequestParams struct {
	// VPC标识符
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeHSMByVpcIdRequest struct {
	*tchttp.BaseRequest
	
	// VPC标识符
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeHSMByVpcIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHSMByVpcIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHSMByVpcIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHSMByVpcIdResponseParams struct {
	// HSM数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 作为查询条件的VpcId
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHSMByVpcIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHSMByVpcIdResponseParams `json:"Response"`
}

func (r *DescribeHSMByVpcIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHSMByVpcIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetRequestParams struct {
	// 返回数量。Limit需要在[1, 100]之间。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量。偏移量最小为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询指定VpcId下的子网信息。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 查找关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

type DescribeSubnetRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量。Limit需要在[1, 100]之间。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量。偏移量最小为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询指定VpcId下的子网信息。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 查找关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "VpcId")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetResponseParams struct {
	// 返回的子网数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的子网实例列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetList []*Subnet `json:"SubnetList,omitempty" name:"SubnetList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubnetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetResponseParams `json:"Response"`
}

func (r *DescribeSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedHsmRequestParams struct {
	// Hsm类型，可选值all、virtulization、GHSM、EHSM、SHSM
	HsmType *string `json:"HsmType,omitempty" name:"HsmType"`
}

type DescribeSupportedHsmRequest struct {
	*tchttp.BaseRequest
	
	// Hsm类型，可选值all、virtulization、GHSM、EHSM、SHSM
	HsmType *string `json:"HsmType,omitempty" name:"HsmType"`
}

func (r *DescribeSupportedHsmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedHsmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HsmType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportedHsmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedHsmResponseParams struct {
	// 当前地域所支持的设备列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceTypes []*DeviceInfo `json:"DeviceTypes,omitempty" name:"DeviceTypes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSupportedHsmResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportedHsmResponseParams `json:"Response"`
}

func (r *DescribeSupportedHsmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedHsmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsgRequestParams struct {
	// 偏移量，当Offset和Limit均为0时将一次性返回用户所有的安全组列表。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回量，当Offset和Limit均为0时将一次性返回用户所有的安全组列表。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

type DescribeUsgRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，当Offset和Limit均为0时将一次性返回用户所有的安全组列表。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回量，当Offset和Limit均为0时将一次性返回用户所有的安全组列表。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeUsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsgResponseParams struct {
	// 用户的安全组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgList []*SgUnit `json:"SgList,omitempty" name:"SgList"`

	// 返回的安全组数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsgResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsgResponseParams `json:"Response"`
}

func (r *DescribeUsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsgRuleRequestParams struct {
	// 根据安全组Id获取安全组详情
	SgIds []*string `json:"SgIds,omitempty" name:"SgIds"`
}

type DescribeUsgRuleRequest struct {
	*tchttp.BaseRequest
	
	// 根据安全组Id获取安全组详情
	SgIds []*string `json:"SgIds,omitempty" name:"SgIds"`
}

func (r *DescribeUsgRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsgRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SgIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsgRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsgRuleResponseParams struct {
	// 安全组详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgRules []*UsgRuleDetail `json:"SgRules,omitempty" name:"SgRules"`

	// 安全组详情数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsgRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsgRuleResponseParams `json:"Response"`
}

func (r *DescribeUsgRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsgRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcRequestParams struct {
	// 返回偏移量。Offset最小为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量。Limit需要在[1, 100]之间。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

type DescribeVpcRequest struct {
	*tchttp.BaseRequest
	
	// 返回偏移量。Offset最小为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量。Limit需要在[1, 100]之间。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcResponseParams struct {
	// 可查询到的所有Vpc实例总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Vpc对象列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcList []*Vpc `json:"VpcList,omitempty" name:"VpcList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVpcResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcResponseParams `json:"Response"`
}

func (r *DescribeVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVsmAttributesRequestParams struct {
	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeVsmAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DescribeVsmAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVsmAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVsmAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVsmAttributesResponseParams struct {
	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源状态，1表示资源为正常，2表示资源处于隔离状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 资源IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 资源所属Vpc
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 资源所属子网
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 资源所属HSM的规格
	Model *string `json:"Model,omitempty" name:"Model"`

	// 资源类型，17表示EVSM，33表示GVSM，49表示SVSM
	VsmType *int64 `json:"VsmType,omitempty" name:"VsmType"`

	// 地域Id，返回腾讯云地域代码，如广州为1，北京为8
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 区域Id，返回腾讯云每个地域的可用区代码
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 资源过期时间，以时间戳形式展示。
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 安全组详情信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgList []*UsgRuleDetail `json:"SgList,omitempty" name:"SgList"`

	// 子网名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 地域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 区域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 实例是否已经过期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expired *bool `json:"Expired,omitempty" name:"Expired"`

	// 为正数表示实例距离过期时间剩余秒数，为负数表示实例已经过期多少秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainSeconds *int64 `json:"RemainSeconds,omitempty" name:"RemainSeconds"`

	// 私有虚拟网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VPC的IPv4 CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 子网的CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitempty" name:"SubnetCidrBlock"`

	// 资源所关联的Tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 资源续费标识，0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费(用户设置)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 厂商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVsmAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVsmAttributesResponseParams `json:"Response"`
}

func (r *DescribeVsmAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVsmAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVsmsRequestParams struct {
	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 设备所属的厂商名称，根据厂商来进行筛选
	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`

	// Hsm服务类型，可选virtualization、physical、GHSM、EHSM、SHSM、all
	HsmType *string `json:"HsmType,omitempty" name:"HsmType"`
}

type DescribeVsmsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 设备所属的厂商名称，根据厂商来进行筛选
	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`

	// Hsm服务类型，可选virtualization、physical、GHSM、EHSM、SHSM、all
	HsmType *string `json:"HsmType,omitempty" name:"HsmType"`
}

func (r *DescribeVsmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVsmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "TagFilters")
	delete(f, "Manufacturer")
	delete(f, "HsmType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVsmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVsmsResponseParams struct {
	// 获取实例的总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VsmList []*ResourceInfo `json:"VsmList,omitempty" name:"VsmList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVsmsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVsmsResponseParams `json:"Response"`
}

func (r *DescribeVsmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVsmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceInfo struct {
	// 厂商名称
	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`

	// 此厂商旗下的设备信息列表
	HsmTypes []*HsmInfo `json:"HsmTypes,omitempty" name:"HsmTypes"`
}

// Predefined struct for user
type GetAlarmEventRequestParams struct {

}

type GetAlarmEventRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetAlarmEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAlarmEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAlarmEventResponseParams struct {
	// 用户所有的告警策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmConfig []*AlarmPolicy `json:"AlarmConfig,omitempty" name:"AlarmConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAlarmEventResponse struct {
	*tchttp.BaseResponse
	Response *GetAlarmEventResponseParams `json:"Response"`
}

func (r *GetAlarmEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVsmMonitorInfoRequestParams struct {
	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type GetVsmMonitorInfoRequest struct {
	*tchttp.BaseRequest
	
	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

func (r *GetVsmMonitorInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVsmMonitorInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ResourceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetVsmMonitorInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVsmMonitorInfoResponseParams struct {
	// VSM监控信息
	MonitorInfo []*string `json:"MonitorInfo,omitempty" name:"MonitorInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetVsmMonitorInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetVsmMonitorInfoResponseParams `json:"Response"`
}

func (r *GetVsmMonitorInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVsmMonitorInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HsmInfo struct {
	// 加密机型号
	Model *string `json:"Model,omitempty" name:"Model"`

	// 此类型的加密机所支持的VSM类型列表
	VsmTypes []*VsmInfo `json:"VsmTypes,omitempty" name:"VsmTypes"`

	// 加密机母机类型：virtualization、GHSM、EHSM、SHSM
	HsmType *string `json:"HsmType,omitempty" name:"HsmType"`
}

// Predefined struct for user
type InquiryPriceBuyVsmRequestParams struct {
	// 需购买实例的数量
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 付费模式：0表示按需计费/后付费，1表示预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 商品的时间大小，整型参数，举例：当TimeSpan为1，TImeUnit为m时，表示询价购买时长为1个月时的价格
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 商品的时间单位，m表示月，y表示年
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 货币类型，默认为CNY
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 默认为CREATE，可选RENEW
	Type *string `json:"Type,omitempty" name:"Type"`

	// Hsm服务类型，可选值virtualization、physical、GHSM、EHSM、SHSM
	HsmType *string `json:"HsmType,omitempty" name:"HsmType"`
}

type InquiryPriceBuyVsmRequest struct {
	*tchttp.BaseRequest
	
	// 需购买实例的数量
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 付费模式：0表示按需计费/后付费，1表示预付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 商品的时间大小，整型参数，举例：当TimeSpan为1，TImeUnit为m时，表示询价购买时长为1个月时的价格
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 商品的时间单位，m表示月，y表示年
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 货币类型，默认为CNY
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 默认为CREATE，可选RENEW
	Type *string `json:"Type,omitempty" name:"Type"`

	// Hsm服务类型，可选值virtualization、physical、GHSM、EHSM、SHSM
	HsmType *string `json:"HsmType,omitempty" name:"HsmType"`
}

func (r *InquiryPriceBuyVsmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceBuyVsmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GoodsNum")
	delete(f, "PayMode")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "Currency")
	delete(f, "Type")
	delete(f, "HsmType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceBuyVsmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceBuyVsmResponseParams struct {
	// 原始总金额，浮点型参数，精确到小数点后两位，如：2000.99
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`

	// 购买的实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 商品的时间大小，整型参数，举例：当TimeSpan为1，TImeUnit为m时，表示询价购买时长为1个月时的价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 商品的时间单位，m表示月，y表示年
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 应付总金额，浮点型参数，精确到小数点后两位，如：2000.99
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceBuyVsmResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceBuyVsmResponseParams `json:"Response"`
}

func (r *InquiryPriceBuyVsmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceBuyVsmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmEventRequestParams struct {
	// 告警事件，支持CPU、MEM、TCP
	Event *string `json:"Event,omitempty" name:"Event"`

	// 告警阈值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 告警状态，0表示停用，1表示启动
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 告警开始时间，只有在这个时间后才会发送告警，当跟EndTime同时为空时表示全天告警
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 告警结束时间，只有在这个时间前才会发送告警，当跟BeginTime同时为空时表示全天告警
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ModifyAlarmEventRequest struct {
	*tchttp.BaseRequest
	
	// 告警事件，支持CPU、MEM、TCP
	Event *string `json:"Event,omitempty" name:"Event"`

	// 告警阈值
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 告警状态，0表示停用，1表示启动
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 告警开始时间，只有在这个时间后才会发送告警，当跟EndTime同时为空时表示全天告警
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 告警结束时间，只有在这个时间前才会发送告警，当跟BeginTime同时为空时表示全天告警
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyAlarmEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Event")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmEventResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmEventResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmEventResponseParams `json:"Response"`
}

func (r *ModifyAlarmEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVsmAttributesRequestParams struct {
	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// UpdateResourceName-修改资源名称,
	// UpdateSgIds-修改安全组名称,
	// UpdateNetWork-修改网络,
	// Default-默认不修改
	Type []*string `json:"Type,omitempty" name:"Type"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 安全组Id
	SgIds []*string `json:"SgIds,omitempty" name:"SgIds"`

	// 虚拟专网Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 告警开关，0表示关闭告警，1表示启用告警
	AlarmStatus *int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
}

type ModifyVsmAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 资源Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// UpdateResourceName-修改资源名称,
	// UpdateSgIds-修改安全组名称,
	// UpdateNetWork-修改网络,
	// Default-默认不修改
	Type []*string `json:"Type,omitempty" name:"Type"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 安全组Id
	SgIds []*string `json:"SgIds,omitempty" name:"SgIds"`

	// 虚拟专网Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 告警开关，0表示关闭告警，1表示启用告警
	AlarmStatus *int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
}

func (r *ModifyVsmAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVsmAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "Type")
	delete(f, "ResourceName")
	delete(f, "SgIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "AlarmStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVsmAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVsmAttributesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyVsmAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVsmAttributesResponseParams `json:"Response"`
}

func (r *ModifyVsmAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVsmAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceInfo struct {
	// 资源Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源状态，1-正常，2-隔离，3-销毁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 资源IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 资源所属Vpc
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 资源所属子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 资源所属HSM规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *string `json:"Model,omitempty" name:"Model"`

	// 云加密机类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VsmType *int64 `json:"VsmType,omitempty" name:"VsmType"`

	// 地域Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 区域Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 过期时间（Epoch Unix Timestamp）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 地域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 区域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 实例的安全组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgList []*SgUnit `json:"SgList,omitempty" name:"SgList"`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 当前实例是否已经过期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expired *bool `json:"Expired,omitempty" name:"Expired"`

	// 为正数表示实例距离过期时间还剩余多少秒，为负数表示已经过期多少秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainSeconds *int64 `json:"RemainSeconds,omitempty" name:"RemainSeconds"`

	// Vpc名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 创建者Uin账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 自动续费状态标识， 0-手动续费，1-自动续费，2-到期不续
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 厂商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manufacturer *string `json:"Manufacturer,omitempty" name:"Manufacturer"`

	// 告警状态，0：停用，1：启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmStatus *int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
}

type SgUnit struct {
	// 安全组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 安全组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgName *string `json:"SgName,omitempty" name:"SgName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgRemark *string `json:"SgRemark,omitempty" name:"SgRemark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Subnet struct {
	// VPC实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网实例ID，例如：subnet-bthucmmy。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 子网的 IPv4 CIDR。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 可用IP数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" name:"AvailableIpAddressCount"`

	// 子网的 IPv6 CIDR。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" name:"Ipv6CidrBlock"`

	// 总IP数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalIpAddressCount *int64 `json:"TotalIpAddressCount,omitempty" name:"TotalIpAddressCount"`

	// 是否为默认Subnet
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type UsgPolicy struct {
	// cidr格式地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 安全组id代表的地址集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 地址组id代表的地址集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proto *string `json:"Proto,omitempty" name:"Proto"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 服务组id代表的协议和端口集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 匹配后行为:ACCEPT/DROP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`
}

type UsgRuleDetail struct {
	// 入站规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	InBound []*UsgPolicy `json:"InBound,omitempty" name:"InBound"`

	// 出站规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutBound []*UsgPolicy `json:"OutBound,omitempty" name:"OutBound"`

	// 安全组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 安全组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgName *string `json:"SgName,omitempty" name:"SgName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgRemark *string `json:"SgRemark,omitempty" name:"SgRemark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *int64 `json:"Version,omitempty" name:"Version"`
}

type Vpc struct {
	// Vpc名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 是否为默认VPC
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type VsmInfo struct {
	// VSM类型名称
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// VSM类型值
	TypeID *int64 `json:"TypeID,omitempty" name:"TypeID"`
}