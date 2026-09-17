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

package v20260401

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ApplyPublicIpsRequestParams struct {
	// 公网实例 ID（路由发布模式必须为 STATIC ）
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 申请Ip数量，最小为 1
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 申请的Ip类型，枚举值：ipv4、ipv6
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ApplyPublicIpsRequest struct {
	*tchttp.BaseRequest
	
	// 公网实例 ID（路由发布模式必须为 STATIC ）
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 申请Ip数量，最小为 1
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 申请的Ip类型，枚举值：ipv4、ipv6
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ApplyPublicIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyPublicIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceId")
	delete(f, "Count")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyPublicIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyPublicIpsResponseParams struct {
	// 分配的公网 IP 地址列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyPublicIpsResponse struct {
	*tchttp.BaseResponse
	Response *ApplyPublicIpsResponseParams `json:"Response"`
}

func (r *ApplyPublicIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyPublicIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeServiceRequestParams struct {
	// 可用区代码，如 ap-guangzhou-1。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type CreateEdgeNodeServiceRequest struct {
	*tchttp.BaseRequest
	
	// 可用区代码，如 ap-guangzhou-1。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *CreateEdgeNodeServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeNodeServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeNodeServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeNodeServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeNodeServiceResponseParams `json:"Response"`
}

func (r *CreateEdgeNodeServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeNodeServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// <p>可用区代码，如 ap-guangzhou-1。</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>机型规格，如 BMS5.MEDIUM8。</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>内网网络实例ID，格式如 net-xxx。</p>
	PrivateNetworkId *string `json:"PrivateNetworkId,omitnil,omitempty" name:"PrivateNetworkId"`

	// <p>公网网络实例ID，格式如 net-xxx。</p>
	PublicNetworkId *string `json:"PublicNetworkId,omitnil,omitempty" name:"PublicNetworkId"`

	// <p>实例名称。</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>镜像ID，如 img-centos-7.9。</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>创建数量，默认1，最大50。</p>
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// <p>登录密码，与SSHKey二选一</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>SSH密钥公钥字符串，与Password二选一</p>
	SSHKey *string `json:"SSHKey,omitnil,omitempty" name:"SSHKey"`

	// <p>镜像版本号，仅公共镜像有版本概念。</p>
	//
	// Deprecated: VersionNumber is deprecated.
	VersionNumber *string `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// <p>是否启用公网IPv6，默认false。启用后系统会在分配IPv4后额外分配一个IPv6地址。</p>
	//
	// Deprecated: EnableIpv6 is deprecated.
	EnableIpv6 *bool `json:"EnableIpv6,omitnil,omitempty" name:"EnableIpv6"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>可用区代码，如 ap-guangzhou-1。</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>机型规格，如 BMS5.MEDIUM8。</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>内网网络实例ID，格式如 net-xxx。</p>
	PrivateNetworkId *string `json:"PrivateNetworkId,omitnil,omitempty" name:"PrivateNetworkId"`

	// <p>公网网络实例ID，格式如 net-xxx。</p>
	PublicNetworkId *string `json:"PublicNetworkId,omitnil,omitempty" name:"PublicNetworkId"`

	// <p>实例名称。</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>镜像ID，如 img-centos-7.9。</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>创建数量，默认1，最大50。</p>
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// <p>登录密码，与SSHKey二选一</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>SSH密钥公钥字符串，与Password二选一</p>
	SSHKey *string `json:"SSHKey,omitnil,omitempty" name:"SSHKey"`

	// <p>镜像版本号，仅公共镜像有版本概念。</p>
	VersionNumber *string `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// <p>是否启用公网IPv6，默认false。启用后系统会在分配IPv4后额外分配一个IPv6地址。</p>
	EnableIpv6 *bool `json:"EnableIpv6,omitnil,omitempty" name:"EnableIpv6"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "InstanceType")
	delete(f, "PrivateNetworkId")
	delete(f, "PublicNetworkId")
	delete(f, "InstanceName")
	delete(f, "ImageId")
	delete(f, "InstanceCount")
	delete(f, "Password")
	delete(f, "SSHKey")
	delete(f, "VersionNumber")
	delete(f, "EnableIpv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// <p>创建成功的实例ID列表。</p>
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// <p>创建失败的实例个数。仅部分失败时返回，全部成功时不返回该字段。</p>
	FailedCount *uint64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancesResponseParams `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateNetworkInstanceRequestParams struct {
	// 新实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网络地址（host 位必须全为 0），必须落在以下 RFC 1918 私有范围之一：`10.0.0.0/8`、`172.16.0.0/12`、`192.168.0.0/16`
	Network *string `json:"Network,omitnil,omitempty" name:"Network"`

	// 掩码位数，上限统一为 `28`，下限随所属私有段而定：`10.0.0.0/8` 允许 `8~28`，`172.16.0.0/12` 允许 `12~28`，`192.168.0.0/16` 允许 `16~28`；需与 Network 共同构成合法网络地址（host 位全为 0）
	Mask *int64 `json:"Mask,omitnil,omitempty" name:"Mask"`
}

type CreatePrivateNetworkInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 新实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网络地址（host 位必须全为 0），必须落在以下 RFC 1918 私有范围之一：`10.0.0.0/8`、`172.16.0.0/12`、`192.168.0.0/16`
	Network *string `json:"Network,omitnil,omitempty" name:"Network"`

	// 掩码位数，上限统一为 `28`，下限随所属私有段而定：`10.0.0.0/8` 允许 `8~28`，`172.16.0.0/12` 允许 `12~28`，`192.168.0.0/16` 允许 `16~28`；需与 Network 共同构成合法网络地址（host 位全为 0）
	Mask *int64 `json:"Mask,omitnil,omitempty" name:"Mask"`
}

func (r *CreatePrivateNetworkInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateNetworkInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceName")
	delete(f, "ZoneId")
	delete(f, "Network")
	delete(f, "Mask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateNetworkInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateNetworkInstanceResponseParams struct {
	// 私网实例ID
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrivateNetworkInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrivateNetworkInstanceResponseParams `json:"Response"`
}

func (r *CreatePrivateNetworkInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateNetworkInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePublicNetworkInstanceRequestParams struct {
	// <p>可用区</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>公网实例名称</p>
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// <p>网络线路</p>
	Line *string `json:"Line,omitnil,omitempty" name:"Line"`

	// <p>路由模式</p>
	RouteMode *string `json:"RouteMode,omitnil,omitempty" name:"RouteMode"`

	// <p>公网带宽（Mbps）</p>
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// <p>BGP AS号</p>
	BgpAsNumber *int64 `json:"BgpAsNumber,omitnil,omitempty" name:"BgpAsNumber"`

	// <p>BGP认证密码</p>
	BgpPassword *string `json:"BgpPassword,omitnil,omitempty" name:"BgpPassword"`

	// <p>公网实例类型</p><p>枚举值：</p><ul><li>standard： 标准型(默认)</li><li>custom： 自定义型(暂不支持创建)</li></ul>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type CreatePublicNetworkInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>可用区</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>公网实例名称</p>
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// <p>网络线路</p>
	Line *string `json:"Line,omitnil,omitempty" name:"Line"`

	// <p>路由模式</p>
	RouteMode *string `json:"RouteMode,omitnil,omitempty" name:"RouteMode"`

	// <p>公网带宽（Mbps）</p>
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// <p>BGP AS号</p>
	BgpAsNumber *int64 `json:"BgpAsNumber,omitnil,omitempty" name:"BgpAsNumber"`

	// <p>BGP认证密码</p>
	BgpPassword *string `json:"BgpPassword,omitnil,omitempty" name:"BgpPassword"`

	// <p>公网实例类型</p><p>枚举值：</p><ul><li>standard： 标准型(默认)</li><li>custom： 自定义型(暂不支持创建)</li></ul>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

func (r *CreatePublicNetworkInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePublicNetworkInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "NetworkInstanceName")
	delete(f, "Line")
	delete(f, "RouteMode")
	delete(f, "Bandwidth")
	delete(f, "BgpAsNumber")
	delete(f, "BgpPassword")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePublicNetworkInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePublicNetworkInstanceResponseParams struct {
	// <p>公网实例 ID</p>
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePublicNetworkInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreatePublicNetworkInstanceResponseParams `json:"Response"`
}

func (r *CreatePublicNetworkInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePublicNetworkInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateNetworkInstanceRequestParams struct {
	// 私网实例Id
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`
}

type DeletePrivateNetworkInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 私网实例Id
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`
}

func (r *DeletePrivateNetworkInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateNetworkInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateNetworkInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrivateNetworkInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrivateNetworkInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrivateNetworkInstanceResponseParams `json:"Response"`
}

func (r *DeletePrivateNetworkInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateNetworkInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePublicNetworkInstanceRequestParams struct {
	// 公网实例 ID
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`
}

type DeletePublicNetworkInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 公网实例 ID
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`
}

func (r *DeletePublicNetworkInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePublicNetworkInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePublicNetworkInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePublicNetworkInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePublicNetworkInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeletePublicNetworkInstanceResponseParams `json:"Response"`
}

func (r *DeletePublicNetworkInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePublicNetworkInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTypesRequestParams struct {
	// 可用区代码，如 ap-guangzhou-1；不传则返回账号下所有可用区的机型。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 分页偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认20，最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstanceTypesRequest struct {
	*tchttp.BaseRequest
	
	// 可用区代码，如 ap-guangzhou-1；不传则返回账号下所有可用区的机型。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 分页偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认20，最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTypesResponseParams struct {
	// 机型配额列表。
	InstanceTypeQuotaSet []*InstanceTypeQuota `json:"InstanceTypeQuotaSet,omitnil,omitempty" name:"InstanceTypeQuotaSet"`

	// 返回记录数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceTypesResponseParams `json:"Response"`
}

func (r *DescribeInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// <p>实例ID列表，用于按实例ID筛选</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>实例名称，支持模糊匹配</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>可用区代码，用于筛选指定可用区的实例</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>实例状态列表，用于按状态筛选实例。可选值：allocating、running、isolating、isolated、terminating、error</p>
	InstanceStatus []*string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// <p>公网网络ID</p>
	PublicNetworkId *string `json:"PublicNetworkId,omitnil,omitempty" name:"PublicNetworkId"`

	// <p>私有网络ID</p>
	PrivateNetworkId *string `json:"PrivateNetworkId,omitnil,omitempty" name:"PrivateNetworkId"`

	// <p>公网IPv4地址列表，用于按公网IP筛选实例</p>
	PublicIps []*string `json:"PublicIps,omitnil,omitempty" name:"PublicIps"`

	// <p>偏移量，默认0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认20，最大100</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID列表，用于按实例ID筛选</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// <p>实例名称，支持模糊匹配</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>可用区代码，用于筛选指定可用区的实例</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>实例状态列表，用于按状态筛选实例。可选值：allocating、running、isolating、isolated、terminating、error</p>
	InstanceStatus []*string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// <p>公网网络ID</p>
	PublicNetworkId *string `json:"PublicNetworkId,omitnil,omitempty" name:"PublicNetworkId"`

	// <p>私有网络ID</p>
	PrivateNetworkId *string `json:"PrivateNetworkId,omitnil,omitempty" name:"PrivateNetworkId"`

	// <p>公网IPv4地址列表，用于按公网IP筛选实例</p>
	PublicIps []*string `json:"PublicIps,omitnil,omitempty" name:"PublicIps"`

	// <p>偏移量，默认0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量，默认20，最大100</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceName")
	delete(f, "Zone")
	delete(f, "InstanceStatus")
	delete(f, "PublicNetworkId")
	delete(f, "PrivateNetworkId")
	delete(f, "PublicIps")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// <p>实例详细信息列表</p>
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// <p>符合条件的实例数量</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateNetworkInstancesRequestParams struct {
	// 私网实例Id
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 新实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrivateNetworkInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 私网实例Id
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 新实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrivateNetworkInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateNetworkInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceId")
	delete(f, "NetworkInstanceName")
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateNetworkInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrivateNetworkInstancesResponseParams struct {
	// 私网实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 私网实例集合
	PrivateNetworkInstanceSet []*PrivateNetworkInstanceInfo `json:"PrivateNetworkInstanceSet,omitnil,omitempty" name:"PrivateNetworkInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrivateNetworkInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateNetworkInstancesResponseParams `json:"Response"`
}

func (r *DescribePrivateNetworkInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateNetworkInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicIpsRequestParams struct {
	// 按公网实例 ID 过滤（子串匹配，多个值取并集）
	NetworkInstanceId []*string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 按可用区/机房过滤
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 按 IP 过滤（子串匹配，多个值取并集）
	Ip []*string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 按状态过滤，可选值：`InUse`、`Unbound`（多个值取并集）
	State []*string `json:"State,omitnil,omitempty" name:"State"`

	// 按 IP 版本过滤，可选值：`Ipv4`、`Ipv6`（多个值取并集）
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// 按创建时间排序，可选值：`asc`、`desc`（默认 `desc`）
	OrderByCreateTime *string `json:"OrderByCreateTime,omitnil,omitempty" name:"OrderByCreateTime"`

	// 按更新时间排序，可选值：`asc`、`desc`（优先级高于创建时间排序）
	OrderByUpdateTime *string `json:"OrderByUpdateTime,omitnil,omitempty" name:"OrderByUpdateTime"`

	// 分页偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePublicIpsRequest struct {
	*tchttp.BaseRequest
	
	// 按公网实例 ID 过滤（子串匹配，多个值取并集）
	NetworkInstanceId []*string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 按可用区/机房过滤
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 按 IP 过滤（子串匹配，多个值取并集）
	Ip []*string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 按状态过滤，可选值：`InUse`、`Unbound`（多个值取并集）
	State []*string `json:"State,omitnil,omitempty" name:"State"`

	// 按 IP 版本过滤，可选值：`Ipv4`、`Ipv6`（多个值取并集）
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// 按创建时间排序，可选值：`asc`、`desc`（默认 `desc`）
	OrderByCreateTime *string `json:"OrderByCreateTime,omitnil,omitempty" name:"OrderByCreateTime"`

	// 按更新时间排序，可选值：`asc`、`desc`（优先级高于创建时间排序）
	OrderByUpdateTime *string `json:"OrderByUpdateTime,omitnil,omitempty" name:"OrderByUpdateTime"`

	// 分页偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePublicIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceId")
	delete(f, "ZoneId")
	delete(f, "Ip")
	delete(f, "State")
	delete(f, "Type")
	delete(f, "OrderByCreateTime")
	delete(f, "OrderByUpdateTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicIpsResponseParams struct {
	// 公网Ip总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 分配的公网 IP 地址列表
	IpInfoSet []*IpInfo `json:"IpInfoSet,omitnil,omitempty" name:"IpInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublicIpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicIpsResponseParams `json:"Response"`
}

func (r *DescribePublicIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicNetworkInstancesRequestParams struct {
	// 公网实例ID
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 公网实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// 可用区Id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePublicNetworkInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 公网实例ID
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 公网实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// 可用区Id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 分页偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认 20，最大 100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePublicNetworkInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicNetworkInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceId")
	delete(f, "NetworkInstanceName")
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicNetworkInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicNetworkInstancesResponseParams struct {
	// 公网实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 公网实例集合
	PublicNetworkInstanceSet []*PublicNetworkInstanceInfo `json:"PublicNetworkInstanceSet,omitnil,omitempty" name:"PublicNetworkInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublicNetworkInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicNetworkInstancesResponseParams `json:"Response"`
}

func (r *DescribePublicNetworkInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicNetworkInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneDataRequestParams struct {
	// 区id
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 指标名(inbw:入带宽，outbw:出带宽)
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 开始时间（UTC时间:0时区）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（UTC时间:0时区）,最多查询2天时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeZoneDataRequest struct {
	*tchttp.BaseRequest
	
	// 区id
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 指标名(inbw:入带宽，outbw:出带宽)
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 开始时间（UTC时间:0时区）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（UTC时间:0时区）,最多查询2天时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeZoneDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "MetricName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneDataResponseParams struct {
	// 统计数据,指标inbw|outbw单位为Mbps
	Data []*SwitchData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZoneDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneDataResponseParams `json:"Response"`
}

func (r *DescribeZoneDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {

}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// <p>所有地域的可用区列表。</p>
	ZoneSet []*ZoneInfo `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// <p>可用区总数量。</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailedInstance struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误码。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type Instance struct {
	// <p>实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>实例名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>绑定的物理机ID</p>
	MachineId *string `json:"MachineId,omitnil,omitempty" name:"MachineId"`

	// <p>机型规格</p>
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// <p>可用区代码</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>镜像ID</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>镜像版本号</p>
	//
	// Deprecated: VersionNumber is deprecated.
	VersionNumber *string `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// <p>实例状态，可选值：allocating、running、isolating、isolated、terminating、error</p>
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// <p>操作状态，可选值：normal、starting、stopping、stopped、rebooting</p>
	OperateStatus *string `json:"OperateStatus,omitnil,omitempty" name:"OperateStatus"`

	// <p>私有网络ID</p>
	PrivateNetworkId *string `json:"PrivateNetworkId,omitnil,omitempty" name:"PrivateNetworkId"`

	// <p>私有IPv4地址</p>
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// <p>私有IPv6地址</p>
	PrivateIpV6 *string `json:"PrivateIpV6,omitnil,omitempty" name:"PrivateIpV6"`

	// <p>公网网络ID</p>
	PublicNetworkId *string `json:"PublicNetworkId,omitnil,omitempty" name:"PublicNetworkId"`

	// <p>公网IPv4地址</p>
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// <p>公网IPv6地址</p>
	PublicIpV6 *string `json:"PublicIpV6,omitnil,omitempty" name:"PublicIpV6"`

	// <p>文件系统类型</p>
	FileSystemType *string `json:"FileSystemType,omitnil,omitempty" name:"FileSystemType"`

	// <p>创建时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>机型族标识</p>
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// <p>机型族名称</p>
	InstanceFamilyName *string `json:"InstanceFamilyName,omitnil,omitempty" name:"InstanceFamilyName"`

	// <p>CPU 型号</p>
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// <p>CPU 核数</p>
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// <p>内存大小</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`
}

type InstanceTypeQuota struct {
	// 可用区代码。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 机型规格。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 机型家族。
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// 机型族名称
	InstanceFamilyName *string `json:"InstanceFamilyName,omitnil,omitempty" name:"InstanceFamilyName"`

	// CPU核数。
	CpuCores *int64 `json:"CpuCores,omitnil,omitempty" name:"CpuCores"`

	// CPU类型。
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// 内存大小（GB）。
	MemoryGb *int64 `json:"MemoryGb,omitnil,omitempty" name:"MemoryGb"`

	// 系统盘类型。
	SystemDiskType *string `json:"SystemDiskType,omitnil,omitempty" name:"SystemDiskType"`

	// 系统盘大小（GB）。
	SystemDiskSize *int64 `json:"SystemDiskSize,omitnil,omitempty" name:"SystemDiskSize"`

	// 系统盘数量。
	SystemDiskCount *uint64 `json:"SystemDiskCount,omitnil,omitempty" name:"SystemDiskCount"`

	// 数据盘类型。
	DataDiskType *string `json:"DataDiskType,omitnil,omitempty" name:"DataDiskType"`

	// 数据盘大小（GB）。
	DataDiskSize *int64 `json:"DataDiskSize,omitnil,omitempty" name:"DataDiskSize"`

	// 数据盘数量。
	DataDiskCount *uint64 `json:"DataDiskCount,omitnil,omitempty" name:"DataDiskCount"`

	// 第二组数据盘类型
	SecondaryDataDiskType *string `json:"SecondaryDataDiskType,omitnil,omitempty" name:"SecondaryDataDiskType"`

	// 第二组数据盘大小(GB)
	SecondaryDataDiskSize *int64 `json:"SecondaryDataDiskSize,omitnil,omitempty" name:"SecondaryDataDiskSize"`

	// 第二组数据盘数量
	SecondaryDataDiskCount *int64 `json:"SecondaryDataDiskCount,omitnil,omitempty" name:"SecondaryDataDiskCount"`

	// 磁盘描述字符串（向后兼容）。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 网络接口类型。
	NetworkInterfaceType *string `json:"NetworkInterfaceType,omitnil,omitempty" name:"NetworkInterfaceType"`

	// GPU类型，无GPU时为空字符串。
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// 配额数量
	Quota *uint64 `json:"Quota,omitnil,omitempty" name:"Quota"`
}

type IpInfo struct {
	// 10.100.0.20
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// epn-asdfghjkl
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// epm-asdfghjkl
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Unbound
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Ipv4
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 2026-04-07T00:00:00
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 2026-04-07T00:00:00
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

// Predefined struct for user
type ModifyInstanceAttributeRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新的实例名称，1-60字符。与 NewPublicIp 至少传入一个。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 新的公网IP（需从该实例所绑定公网实例的可用IP中选择）。与 InstanceName 至少传入一个。
	//
	// Deprecated: NewPublicIp is deprecated.
	NewPublicIp *string `json:"NewPublicIp,omitnil,omitempty" name:"NewPublicIp"`

	// IP类型，ipv4 或 ipv6，默认 ipv4。仅在指定 NewPublicIp 时有效。
	//
	// Deprecated: IpType is deprecated.
	IpType *string `json:"IpType,omitnil,omitempty" name:"IpType"`
}

type ModifyInstanceAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新的实例名称，1-60字符。与 NewPublicIp 至少传入一个。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 新的公网IP（需从该实例所绑定公网实例的可用IP中选择）。与 InstanceName 至少传入一个。
	NewPublicIp *string `json:"NewPublicIp,omitnil,omitempty" name:"NewPublicIp"`

	// IP类型，ipv4 或 ipv6，默认 ipv4。仅在指定 NewPublicIp 时有效。
	IpType *string `json:"IpType,omitnil,omitempty" name:"IpType"`
}

func (r *ModifyInstanceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "NewPublicIp")
	delete(f, "IpType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceAttributeResponseParams `json:"Response"`
}

func (r *ModifyInstanceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateNetworkInstanceRequestParams struct {
	// 私网实例Id
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 新实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`
}

type ModifyPrivateNetworkInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 私网实例Id
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 新实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`
}

func (r *ModifyPrivateNetworkInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateNetworkInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceId")
	delete(f, "NetworkInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateNetworkInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateNetworkInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrivateNetworkInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateNetworkInstanceResponseParams `json:"Response"`
}

func (r *ModifyPrivateNetworkInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateNetworkInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPublicNetworkInstanceRequestParams struct {
	// 公网实例 ID
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 新实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`
}

type ModifyPublicNetworkInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 公网实例 ID
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 新实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`
}

func (r *ModifyPublicNetworkInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicNetworkInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceId")
	delete(f, "NetworkInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPublicNetworkInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPublicNetworkInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPublicNetworkInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPublicNetworkInstanceResponseParams `json:"Response"`
}

func (r *ModifyPublicNetworkInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicNetworkInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrivateNetworkInstanceInfo struct {
	// 私网实例ID
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 私网实例名称
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 网络地址
	Network *string `json:"Network,omitnil,omitempty" name:"Network"`

	// 网络掩码
	Mask *int64 `json:"Mask,omitnil,omitempty" name:"Mask"`

	// 关联物理机数量
	ServerCount *int64 `json:"ServerCount,omitnil,omitempty" name:"ServerCount"`

	// 可用Ip数量
	AvailableIpCount *int64 `json:"AvailableIpCount,omitnil,omitempty" name:"AvailableIpCount"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type PublicNetworkInstanceInfo struct {
	// 公网实例ID
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 公网实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkInstanceName *string `json:"NetworkInstanceName,omitnil,omitempty" name:"NetworkInstanceName"`

	// 带宽，单位Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 线路信息
	Line *string `json:"Line,omitnil,omitempty" name:"Line"`

	// 路由模式，枚举值：STATIC、BGP、OSPF
	RouteMode *string `json:"RouteMode,omitnil,omitempty" name:"RouteMode"`

	// 关联的物理服务器数量
	ServerCount *int64 `json:"ServerCount,omitnil,omitempty" name:"ServerCount"`

	// 已申请的Ipv4数量
	Ipv4Count *int64 `json:"Ipv4Count,omitnil,omitempty" name:"Ipv4Count"`

	// 已申请的Ipv6数量
	Ipv6Count *int64 `json:"Ipv6Count,omitnil,omitempty" name:"Ipv6Count"`

	// 关联的Ipv4网段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv4CidrSet []*PublicNetworkSegment `json:"Ipv4CidrSet,omitnil,omitempty" name:"Ipv4CidrSet"`

	// 关联的Ipv6网段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6CidrSet []*PublicNetworkSegment `json:"Ipv6CidrSet,omitnil,omitempty" name:"Ipv6CidrSet"`

	// 公网实例创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 公网实例修改时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type PublicNetworkSegment struct {
	// 网段Cidr
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`

	// 网关Ip
	Gateway *string `json:"Gateway,omitnil,omitempty" name:"Gateway"`
}

// Predefined struct for user
type ReleasePublicIpRequestParams struct {
	// 公网实例 ID（路由发布模式为 STATIC ）
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 待释放的Ip类型，枚举值：ipv4、ipv6
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 待释放的 Ip 地址列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`
}

type ReleasePublicIpRequest struct {
	*tchttp.BaseRequest
	
	// 公网实例 ID（路由发布模式为 STATIC ）
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil,omitempty" name:"NetworkInstanceId"`

	// 待释放的Ip类型，枚举值：ipv4、ipv6
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 待释放的 Ip 地址列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`
}

func (r *ReleasePublicIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleasePublicIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInstanceId")
	delete(f, "Type")
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleasePublicIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleasePublicIpResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleasePublicIpResponse struct {
	*tchttp.BaseResponse
	Response *ReleasePublicIpResponseParams `json:"Response"`
}

func (r *ReleasePublicIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleasePublicIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchData struct {
	// UTC时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 统计值
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type TerminateInstancesRequestParams struct {
	// <p>实例ID列表，最多100个。</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例ID列表，最多100个。</p>
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateInstancesResponseParams struct {
	// <p>销毁成功的实例ID列表。</p>
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// <p>销毁失败的实例信息列表（部分成功时返回）。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedInstanceSet []*FailedInstance `json:"FailedInstanceSet,omitnil,omitempty" name:"FailedInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateInstancesResponseParams `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneInfo struct {
	// 可用区ID。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区代码。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区中文名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 可用区英文名称。
	ZoneNameEn *string `json:"ZoneNameEn,omitnil,omitempty" name:"ZoneNameEn"`

	// 地域代码。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 区域代码。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 区域名称。
	LocationName *string `json:"LocationName,omitnil,omitempty" name:"LocationName"`
}