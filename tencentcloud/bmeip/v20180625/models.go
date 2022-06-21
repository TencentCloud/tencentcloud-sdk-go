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

package v20180625

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type BindEipAclsRequestParams struct {
	// 待关联的 EIP 与 ACL关系列表
	EipIdAclIdList []*EipAclMap `json:"EipIdAclIdList,omitempty" name:"EipIdAclIdList"`
}

type BindEipAclsRequest struct {
	*tchttp.BaseRequest
	
	// 待关联的 EIP 与 ACL关系列表
	EipIdAclIdList []*EipAclMap `json:"EipIdAclIdList,omitempty" name:"EipIdAclIdList"`
}

func (r *BindEipAclsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEipAclsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipIdAclIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindEipAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindEipAclsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindEipAclsResponse struct {
	*tchttp.BaseResponse
	Response *BindEipAclsResponseParams `json:"Response"`
}

func (r *BindEipAclsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEipAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindHostedRequestParams struct {
	// Eip实例ID，可通过DescribeBmEip 接口返回字段中的 eipId获取。Eip和EipId参数必须要填写一个。
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// 托管机器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type BindHostedRequest struct {
	*tchttp.BaseRequest
	
	// Eip实例ID，可通过DescribeBmEip 接口返回字段中的 eipId获取。Eip和EipId参数必须要填写一个。
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// 托管机器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *BindHostedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindHostedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindHostedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindHostedResponseParams struct {
	// 异步任务ID，可以通过EipBmQueryTask查询任务状态
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindHostedResponse struct {
	*tchttp.BaseResponse
	Response *BindHostedResponseParams `json:"Response"`
}

func (r *BindHostedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindHostedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindRsRequestParams struct {
	// Eip实例ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// 物理服务器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type BindRsRequest struct {
	*tchttp.BaseRequest
	
	// Eip实例ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// 物理服务器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *BindRsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindRsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindRsResponseParams struct {
	// 绑定黑石物理机异步任务ID，可以通过DescribeEipTask查询任务状态
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindRsResponse struct {
	*tchttp.BaseResponse
	Response *BindRsResponseParams `json:"Response"`
}

func (r *BindRsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindVpcIpRequestParams struct {
	// Eip实例ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// EIP归属VpcId，例如vpc-k7j1t2x1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 绑定的VPC内IP地址
	VpcIp *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

type BindVpcIpRequest struct {
	*tchttp.BaseRequest
	
	// Eip实例ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// EIP归属VpcId，例如vpc-k7j1t2x1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 绑定的VPC内IP地址
	VpcIp *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

func (r *BindVpcIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindVpcIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipId")
	delete(f, "VpcId")
	delete(f, "VpcIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindVpcIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindVpcIpResponseParams struct {
	// EIP绑定VPC网络IP异步任务ID，可以通过查询EIP任务状态查询任务状态
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindVpcIpResponse struct {
	*tchttp.BaseResponse
	Response *BindVpcIpResponseParams `json:"Response"`
}

func (r *BindVpcIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindVpcIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEipAclRequestParams struct {
	// ACL 名称
	AclName *string `json:"AclName,omitempty" name:"AclName"`

	// ACL 状态 0：无状态，1：有状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type CreateEipAclRequest struct {
	*tchttp.BaseRequest
	
	// ACL 名称
	AclName *string `json:"AclName,omitempty" name:"AclName"`

	// ACL 状态 0：无状态，1：有状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *CreateEipAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEipAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AclName")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEipAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEipAclResponseParams struct {
	// ACL 实例 ID
	AclId *string `json:"AclId,omitempty" name:"AclId"`

	// ACL 实例状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// ACL 实例名称
	AclName *string `json:"AclName,omitempty" name:"AclName"`

	// ACL 实例创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEipAclResponse struct {
	*tchttp.BaseResponse
	Response *CreateEipAclResponseParams `json:"Response"`
}

func (r *CreateEipAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEipAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEipRequestParams struct {
	// 申请数量，默认为1, 最大 20
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// EIP计费方式，flow-流量计费；bandwidth-带宽计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 带宽设定值（只在带宽计费时生效）
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// EIP模式，目前支持tunnel和fullnat
	SetType *string `json:"SetType,omitempty" name:"SetType"`

	// 是否使用独占集群，0：不使用，1：使用。默认为0
	Exclusive *uint64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// EIP归属私有网络ID，例如vpc-k7j1t2x1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 指定申请的IP列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

type CreateEipRequest struct {
	*tchttp.BaseRequest
	
	// 申请数量，默认为1, 最大 20
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// EIP计费方式，flow-流量计费；bandwidth-带宽计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 带宽设定值（只在带宽计费时生效）
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// EIP模式，目前支持tunnel和fullnat
	SetType *string `json:"SetType,omitempty" name:"SetType"`

	// 是否使用独占集群，0：不使用，1：使用。默认为0
	Exclusive *uint64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// EIP归属私有网络ID，例如vpc-k7j1t2x1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 指定申请的IP列表
	IpList []*string `json:"IpList,omitempty" name:"IpList"`
}

func (r *CreateEipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GoodsNum")
	delete(f, "PayMode")
	delete(f, "Bandwidth")
	delete(f, "SetType")
	delete(f, "Exclusive")
	delete(f, "VpcId")
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEipResponseParams struct {
	// EIP列表
	EipIds []*string `json:"EipIds,omitempty" name:"EipIds"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEipResponse struct {
	*tchttp.BaseResponse
	Response *CreateEipResponseParams `json:"Response"`
}

func (r *CreateEipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEipAclRequestParams struct {
	// 待删除的 ACL 实例 ID
	AclId *string `json:"AclId,omitempty" name:"AclId"`
}

type DeleteEipAclRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的 ACL 实例 ID
	AclId *string `json:"AclId,omitempty" name:"AclId"`
}

func (r *DeleteEipAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEipAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AclId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEipAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEipAclResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEipAclResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEipAclResponseParams `json:"Response"`
}

func (r *DeleteEipAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEipAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEipRequestParams struct {
	// Eip实例ID列表
	EipIds []*string `json:"EipIds,omitempty" name:"EipIds"`
}

type DeleteEipRequest struct {
	*tchttp.BaseRequest
	
	// Eip实例ID列表
	EipIds []*string `json:"EipIds,omitempty" name:"EipIds"`
}

func (r *DeleteEipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEipResponseParams struct {
	// 任务Id
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEipResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEipResponseParams `json:"Response"`
}

func (r *DeleteEipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEipAclsRequestParams struct {
	// ACL 名称，支持模糊查找
	AclName *string `json:"AclName,omitempty" name:"AclName"`

	// ACL 实例 ID 列表，数组下标从 0 开始
	AclIds []*string `json:"AclIds,omitempty" name:"AclIds"`

	// 分页参数。偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数。每一页的 EIPACL 列表数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// EIP实例ID列表
	EipIds []*string `json:"EipIds,omitempty" name:"EipIds"`

	// EIP IP地址列表
	EipIps []*string `json:"EipIps,omitempty" name:"EipIps"`

	// EIP名称列表
	EipNames []*string `json:"EipNames,omitempty" name:"EipNames"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值：0:增序(默认)，1:降序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// ACL名称列表，支持模糊查找
	AclNames []*string `json:"AclNames,omitempty" name:"AclNames"`
}

type DescribeEipAclsRequest struct {
	*tchttp.BaseRequest
	
	// ACL 名称，支持模糊查找
	AclName *string `json:"AclName,omitempty" name:"AclName"`

	// ACL 实例 ID 列表，数组下标从 0 开始
	AclIds []*string `json:"AclIds,omitempty" name:"AclIds"`

	// 分页参数。偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数。每一页的 EIPACL 列表数目
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// EIP实例ID列表
	EipIds []*string `json:"EipIds,omitempty" name:"EipIds"`

	// EIP IP地址列表
	EipIps []*string `json:"EipIps,omitempty" name:"EipIps"`

	// EIP名称列表
	EipNames []*string `json:"EipNames,omitempty" name:"EipNames"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值：0:增序(默认)，1:降序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// ACL名称列表，支持模糊查找
	AclNames []*string `json:"AclNames,omitempty" name:"AclNames"`
}

func (r *DescribeEipAclsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEipAclsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AclName")
	delete(f, "AclIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EipIds")
	delete(f, "EipIps")
	delete(f, "EipNames")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "AclNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEipAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEipAclsResponseParams struct {
	// 返回 EIPACL 列表总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// EIPACL列表
	EipAclList []*EipAcl `json:"EipAclList,omitempty" name:"EipAclList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEipAclsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEipAclsResponseParams `json:"Response"`
}

func (r *DescribeEipAclsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEipAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEipQuotaRequestParams struct {

}

type DescribeEipQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEipQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEipQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEipQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEipQuotaResponseParams struct {
	// 能拥有的EIP个数的总配额，默认是100个
	EipNumQuota *int64 `json:"EipNumQuota,omitempty" name:"EipNumQuota"`

	// 当前已使用的EIP个数，包括创建中、绑定中、已绑定、解绑中、未绑定几种状态的EIP个数总和
	CurrentEipNum *int64 `json:"CurrentEipNum,omitempty" name:"CurrentEipNum"`

	// 当天申请EIP次数
	DailyApplyCount *int64 `json:"DailyApplyCount,omitempty" name:"DailyApplyCount"`

	// 每日申请EIP的次数限制
	DailyApplyQuota *int64 `json:"DailyApplyQuota,omitempty" name:"DailyApplyQuota"`

	// BatchApplyMax
	BatchApplyMax *int64 `json:"BatchApplyMax,omitempty" name:"BatchApplyMax"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEipQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEipQuotaResponseParams `json:"Response"`
}

func (r *DescribeEipQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEipQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEipTaskRequestParams struct {
	// EIP查询任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeEipTaskRequest struct {
	*tchttp.BaseRequest
	
	// EIP查询任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeEipTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEipTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEipTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEipTaskResponseParams struct {
	// 当前任务状态码：0-成功，1-失败，2-进行中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEipTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEipTaskResponseParams `json:"Response"`
}

func (r *DescribeEipTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEipTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEipsRequestParams struct {
	// EIP实例ID列表
	EipIds []*string `json:"EipIds,omitempty" name:"EipIds"`

	// EIP IP 列表
	Eips []*string `json:"Eips,omitempty" name:"Eips"`

	// 主机实例ID 列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// EIP名称,模糊匹配
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 状态列表, 默认所有
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回EIP数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持： EipId,Eip,Status, InstanceId,CreatedAt
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增 1:递减(默认)
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 计费模式,流量：flow，带宽：bandwidth
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// EIP归属VpcId，例如vpc-k7j1t2x1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 绑定类型，-1：未绑定，0：物理机，1：nat网关，2：虚拟IP, 3:托管机器
	BindTypes []*int64 `json:"BindTypes,omitempty" name:"BindTypes"`

	// 独占标志，0：共享，1：独占
	ExclusiveTag *int64 `json:"ExclusiveTag,omitempty" name:"ExclusiveTag"`

	// EIP ACL实例ID
	AclId *string `json:"AclId,omitempty" name:"AclId"`

	// 搜索条件，是否绑定了EIP ACL， 0：未绑定，1：绑定
	BindAcl *int64 `json:"BindAcl,omitempty" name:"BindAcl"`
}

type DescribeEipsRequest struct {
	*tchttp.BaseRequest
	
	// EIP实例ID列表
	EipIds []*string `json:"EipIds,omitempty" name:"EipIds"`

	// EIP IP 列表
	Eips []*string `json:"Eips,omitempty" name:"Eips"`

	// 主机实例ID 列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// EIP名称,模糊匹配
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 状态列表, 默认所有
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回EIP数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持： EipId,Eip,Status, InstanceId,CreatedAt
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式 0:递增 1:递减(默认)
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 计费模式,流量：flow，带宽：bandwidth
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// EIP归属VpcId，例如vpc-k7j1t2x1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 绑定类型，-1：未绑定，0：物理机，1：nat网关，2：虚拟IP, 3:托管机器
	BindTypes []*int64 `json:"BindTypes,omitempty" name:"BindTypes"`

	// 独占标志，0：共享，1：独占
	ExclusiveTag *int64 `json:"ExclusiveTag,omitempty" name:"ExclusiveTag"`

	// EIP ACL实例ID
	AclId *string `json:"AclId,omitempty" name:"AclId"`

	// 搜索条件，是否绑定了EIP ACL， 0：未绑定，1：绑定
	BindAcl *int64 `json:"BindAcl,omitempty" name:"BindAcl"`
}

func (r *DescribeEipsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEipsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipIds")
	delete(f, "Eips")
	delete(f, "InstanceIds")
	delete(f, "SearchKey")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "PayMode")
	delete(f, "VpcId")
	delete(f, "BindTypes")
	delete(f, "ExclusiveTag")
	delete(f, "AclId")
	delete(f, "BindAcl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEipsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEipsResponseParams struct {
	// 返回EIP信息数组
	EipSet []*EipInfo `json:"EipSet,omitempty" name:"EipSet"`

	// 返回EIP数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEipsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEipsResponseParams `json:"Response"`
}

func (r *DescribeEipsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EipAcl struct {
	// ACL 实例 ID。
	AclId *string `json:"AclId,omitempty" name:"AclId"`

	// ACL 实例名称
	AclName *string `json:"AclName,omitempty" name:"AclName"`

	// ACL 状态。0：无状态，1：有状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// EIPACL 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// EIPACL 已关联的 eip 数目
	EipNum *int64 `json:"EipNum,omitempty" name:"EipNum"`

	// 出站规则
	OutRules []*EipAclRule `json:"OutRules,omitempty" name:"OutRules"`

	// 入站规则
	InRules []*EipAclRule `json:"InRules,omitempty" name:"InRules"`
}

type EipAclMap struct {
	// EIP 实例 ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// ACL 实例 ID
	AclId *string `json:"AclId,omitempty" name:"AclId"`
}

type EipAclRule struct {
	// 源 IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 目标端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// 协议(TCP/UDP/ICMP/ANY)
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 策略（accept/drop）
	Action *string `json:"Action,omitempty" name:"Action"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

type EipInfo struct {
	// EIP实例ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// EIP名称
	EipName *string `json:"EipName,omitempty" name:"EipName"`

	// EIP地址
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 运营商ID 0：电信； 1：联通； 2：移动； 3：教育网； 4：盈科； 5：BGP； 6：中国香港
	IspId *int64 `json:"IspId,omitempty" name:"IspId"`

	// 状态 0：创建中； 1：绑定中； 2：已绑定； 3：解绑中； 4：未绑定； 6：下线中； 9：创建失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 是否欠费隔离 1： 欠费隔离； 0： 正常。处在欠费隔离情况下的EIP不能进行任何管理操作。
	Arrears *int64 `json:"Arrears,omitempty" name:"Arrears"`

	// EIP所绑定的服务器实例ID，未绑定则为空
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 服务器别名
	InstanceAlias *string `json:"InstanceAlias,omitempty" name:"InstanceAlias"`

	// EIP解绑时间
	FreeAt *string `json:"FreeAt,omitempty" name:"FreeAt"`

	// EIP创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// EIP更新时间
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// EIP未绑定服务器时长（单位：秒）
	FreeSecond *int64 `json:"FreeSecond,omitempty" name:"FreeSecond"`

	// EIP所绑定的资源类型，-1：未绑定资源；0：黑石物理机，字段对应unInstanceId；1：Nat网关，字段对应natUid；2：云服务器字段对应vpcIp; 3: 托管机器，字段对应HInstanceId, HInstanceAlias
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// EIP计费模式，"flow"：流量计费； "bandwidth"：带宽计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// EIP带宽计费模式下的带宽上限（单位：MB）
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 最近一次操作变更的EIP计费模式，"flow"：流量计费； "bandwidth"：带宽计费
	LatestPayMode *string `json:"LatestPayMode,omitempty" name:"LatestPayMode"`

	// 最近一次操作变更的EIP计费模式对应的带宽上限值，仅在带宽计费模式下有效（单位：MB）
	LatestBandwidth *int64 `json:"LatestBandwidth,omitempty" name:"LatestBandwidth"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// EIP所绑定的NAT网关的数字ID，形如：1001,，未绑定则为空
	NatId *int64 `json:"NatId,omitempty" name:"NatId"`

	// EIP所绑定的NAT网关实例ID，形如："nat-n47xxxxx"，未绑定则为空
	NatUid *string `json:"NatUid,omitempty" name:"NatUid"`

	// EIP所绑定的云服务器IP(托管或者云服务器的IP），形如："10.1.1.3"。 注意：IP资源需要通过bmvpc模块注册或者申请后才可以绑定eip，接口使用申请子网IP和注册子网IP：,未绑定则为空
	VpcIp *string `json:"VpcIp,omitempty" name:"VpcIp"`

	// 私有网络实例ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 是否为独占类型EIP
	Exclusive *int64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// 私有网络的cidr
	VpcCidr *string `json:"VpcCidr,omitempty" name:"VpcCidr"`

	// EIP ACL实例ID
	AclId *string `json:"AclId,omitempty" name:"AclId"`

	// EIP ACL名称
	AclName *string `json:"AclName,omitempty" name:"AclName"`

	// 托管机器实例ID
	HInstanceId *string `json:"HInstanceId,omitempty" name:"HInstanceId"`

	// 托管机器别名
	HInstanceAlias *string `json:"HInstanceAlias,omitempty" name:"HInstanceAlias"`
}

type EipRsMap struct {
	// EIP实例 ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// 黑石物理机实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type ModifyEipAclRequestParams struct {
	// ACL 实例 ID
	AclId *string `json:"AclId,omitempty" name:"AclId"`

	// ACL 名称
	AclName *string `json:"AclName,omitempty" name:"AclName"`

	// ACL 状态。0：无状态 1：有状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 规则类型（in/out）。in：入站规则 out：出站规则
	Type *string `json:"Type,omitempty" name:"Type"`

	// ACL规则列表
	Rules []*EipAclRule `json:"Rules,omitempty" name:"Rules"`
}

type ModifyEipAclRequest struct {
	*tchttp.BaseRequest
	
	// ACL 实例 ID
	AclId *string `json:"AclId,omitempty" name:"AclId"`

	// ACL 名称
	AclName *string `json:"AclName,omitempty" name:"AclName"`

	// ACL 状态。0：无状态 1：有状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 规则类型（in/out）。in：入站规则 out：出站规则
	Type *string `json:"Type,omitempty" name:"Type"`

	// ACL规则列表
	Rules []*EipAclRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *ModifyEipAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEipAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AclId")
	delete(f, "AclName")
	delete(f, "Status")
	delete(f, "Type")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEipAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEipAclResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEipAclResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEipAclResponseParams `json:"Response"`
}

func (r *ModifyEipAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEipAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEipChargeRequestParams struct {
	// EIP计费方式，flow-流量计费；bandwidth-带宽计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Eip实例ID列表
	EipIds []*string `json:"EipIds,omitempty" name:"EipIds"`

	// 带宽设定值（只在带宽计费时生效）
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

type ModifyEipChargeRequest struct {
	*tchttp.BaseRequest
	
	// EIP计费方式，flow-流量计费；bandwidth-带宽计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Eip实例ID列表
	EipIds []*string `json:"EipIds,omitempty" name:"EipIds"`

	// 带宽设定值（只在带宽计费时生效）
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

func (r *ModifyEipChargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEipChargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayMode")
	delete(f, "EipIds")
	delete(f, "Bandwidth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEipChargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEipChargeResponseParams struct {
	// 修改计费模式的异步任务ID，可以通过查询EIP任务状态查询任务状态
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEipChargeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEipChargeResponseParams `json:"Response"`
}

func (r *ModifyEipChargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEipChargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEipNameRequestParams struct {
	// Eip实例ID，可通过/v2/DescribeEip 接口返回字段中的 eipId获取
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// EIP 实例别名
	EipName *string `json:"EipName,omitempty" name:"EipName"`
}

type ModifyEipNameRequest struct {
	*tchttp.BaseRequest
	
	// Eip实例ID，可通过/v2/DescribeEip 接口返回字段中的 eipId获取
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// EIP 实例别名
	EipName *string `json:"EipName,omitempty" name:"EipName"`
}

func (r *ModifyEipNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEipNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipId")
	delete(f, "EipName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEipNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEipNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEipNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEipNameResponseParams `json:"Response"`
}

func (r *ModifyEipNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEipNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindEipAclsRequestParams struct {
	// 待解关联的 EIP 与 ACL列表
	EipIdAclIdList []*EipAclMap `json:"EipIdAclIdList,omitempty" name:"EipIdAclIdList"`
}

type UnbindEipAclsRequest struct {
	*tchttp.BaseRequest
	
	// 待解关联的 EIP 与 ACL列表
	EipIdAclIdList []*EipAclMap `json:"EipIdAclIdList,omitempty" name:"EipIdAclIdList"`
}

func (r *UnbindEipAclsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindEipAclsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipIdAclIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindEipAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindEipAclsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindEipAclsResponse struct {
	*tchttp.BaseResponse
	Response *UnbindEipAclsResponseParams `json:"Response"`
}

func (r *UnbindEipAclsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindEipAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindHostedRequestParams struct {
	// 托管机器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Eip实例ID，可通过DescribeBmEip 接口返回字段中的 eipId获取。Eip和EipId参数必须要填写一个。
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// 弹性IP。Eip和EipId参数必须要填写一个。
	Eip *string `json:"Eip,omitempty" name:"Eip"`
}

type UnbindHostedRequest struct {
	*tchttp.BaseRequest
	
	// 托管机器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Eip实例ID，可通过DescribeBmEip 接口返回字段中的 eipId获取。Eip和EipId参数必须要填写一个。
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// 弹性IP。Eip和EipId参数必须要填写一个。
	Eip *string `json:"Eip,omitempty" name:"Eip"`
}

func (r *UnbindHostedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindHostedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EipId")
	delete(f, "Eip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindHostedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindHostedResponseParams struct {
	// 异步任务ID，可以通过EipBmQueryTask查询任务状态
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindHostedResponse struct {
	*tchttp.BaseResponse
	Response *UnbindHostedResponseParams `json:"Response"`
}

func (r *UnbindHostedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindHostedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindRsListRequestParams struct {
	// 物理机绑定的EIP列表
	EipRsList []*EipRsMap `json:"EipRsList,omitempty" name:"EipRsList"`
}

type UnbindRsListRequest struct {
	*tchttp.BaseRequest
	
	// 物理机绑定的EIP列表
	EipRsList []*EipRsMap `json:"EipRsList,omitempty" name:"EipRsList"`
}

func (r *UnbindRsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindRsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipRsList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindRsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindRsListResponseParams struct {
	// 解绑操作的异步任务ID，可以通过查询EIP任务状态查询任务状态
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindRsListResponse struct {
	*tchttp.BaseResponse
	Response *UnbindRsListResponseParams `json:"Response"`
}

func (r *UnbindRsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindRsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindRsRequestParams struct {
	// Eip实例ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// 物理服务器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type UnbindRsRequest struct {
	*tchttp.BaseRequest
	
	// Eip实例ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// 物理服务器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *UnbindRsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindRsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindRsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindRsResponseParams struct {
	// 解绑操作的异步任务ID，可以通过查询EIP任务状态查询任务状态
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindRsResponse struct {
	*tchttp.BaseResponse
	Response *UnbindRsResponseParams `json:"Response"`
}

func (r *UnbindRsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindRsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindVpcIpRequestParams struct {
	// Eip实例ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// EIP归属VpcId，例如vpc-k7j1t2x1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 绑定的VPC内IP地址
	VpcIp *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

type UnbindVpcIpRequest struct {
	*tchttp.BaseRequest
	
	// Eip实例ID
	EipId *string `json:"EipId,omitempty" name:"EipId"`

	// EIP归属VpcId，例如vpc-k7j1t2x1
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 绑定的VPC内IP地址
	VpcIp *string `json:"VpcIp,omitempty" name:"VpcIp"`
}

func (r *UnbindVpcIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindVpcIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EipId")
	delete(f, "VpcId")
	delete(f, "VpcIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindVpcIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindVpcIpResponseParams struct {
	// 绑定黑石物理机异步任务ID，可以通过查询EIP任务状态查询任务状态
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindVpcIpResponse struct {
	*tchttp.BaseResponse
	Response *UnbindVpcIpResponseParams `json:"Response"`
}

func (r *UnbindVpcIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindVpcIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}