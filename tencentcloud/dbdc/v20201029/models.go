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

package v20201029

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddNodesToDBCustomClusterRequestParams struct {
	// <p>集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>需上架的节点 ID 列表</p>
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`

	// <p>节点上架后重设的操作系统镜像ID</p><p>取值参考：可通过&quot;DescribeDBCustomImages&quot;接口获取支持的镜像。</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。</p><p>入参限制：若选择密钥方式，KeyIds 仅支持单个 ID。三种方式必须且仅可以设置其中一种。</p>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`
}

type AddNodesToDBCustomClusterRequest struct {
	*tchttp.BaseRequest
	
	// <p>集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>需上架的节点 ID 列表</p>
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`

	// <p>节点上架后重设的操作系统镜像ID</p><p>取值参考：可通过&quot;DescribeDBCustomImages&quot;接口获取支持的镜像。</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。</p><p>入参限制：若选择密钥方式，KeyIds 仅支持单个 ID。三种方式必须且仅可以设置其中一种。</p>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`
}

func (r *AddNodesToDBCustomClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodesToDBCustomClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodeIds")
	delete(f, "ImageId")
	delete(f, "LoginSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNodesToDBCustomClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodesToDBCustomClusterResponseParams struct {
	// <p>上架节点的任务ID</p>
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddNodesToDBCustomClusterResponse struct {
	*tchttp.BaseResponse
	Response *AddNodesToDBCustomClusterResponseParams `json:"Response"`
}

func (r *AddNodesToDBCustomClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodesToDBCustomClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiServerNetwork struct {
	// <p>API Server的访问私有网络ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>API Server的访问私有网络子网ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

// Predefined struct for user
type CheckRoleAuthorizedRequestParams struct {
	// <p>待检测的角色名字</p>
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type CheckRoleAuthorizedRequest struct {
	*tchttp.BaseRequest
	
	// <p>待检测的角色名字</p>
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *CheckRoleAuthorizedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRoleAuthorizedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckRoleAuthorizedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckRoleAuthorizedResponseParams struct {
	// <p>角色权限状态</p><p>枚举值：</p><ul><li>AUTHORIZED： 已经创建授权</li><li>NEED_GRANT： 未授权</li><li>ERROR： 报错</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>出错的错误信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckRoleAuthorizedResponse struct {
	*tchttp.BaseResponse
	Response *CheckRoleAuthorizedResponseParams `json:"Response"`
}

func (r *CheckRoleAuthorizedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRoleAuthorizedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerNetwork struct {
	// <p>容器网络的虚拟网络ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>容器网络的虚拟网络子网列表</p>
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

// Predefined struct for user
type CreateDBCustomClusterRequestParams struct {
	// <p>容器网络，在此集群的所有 POD 与此网络连通</p>
	ContainerNetwork *ContainerNetwork `json:"ContainerNetwork,omitnil,omitempty" name:"ContainerNetwork"`

	// <p>集群名称</p><p>入参限制：最长128个字符，只能为中文，英文，下划线。</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>集群的API Server的网络信息</p><p>入参限制：必须为此账号下拥有的网络地址，可以与容器网络保持一致。</p>
	ApiServerNetwork *ApiServerNetwork `json:"ApiServerNetwork,omitnil,omitempty" name:"ApiServerNetwork"`

	// <p>集群描述</p>
	ClusterDescription *string `json:"ClusterDescription,omitnil,omitempty" name:"ClusterDescription"`

	// <p>集群标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>客户端Token</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

type CreateDBCustomClusterRequest struct {
	*tchttp.BaseRequest
	
	// <p>容器网络，在此集群的所有 POD 与此网络连通</p>
	ContainerNetwork *ContainerNetwork `json:"ContainerNetwork,omitnil,omitempty" name:"ContainerNetwork"`

	// <p>集群名称</p><p>入参限制：最长128个字符，只能为中文，英文，下划线。</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>集群的API Server的网络信息</p><p>入参限制：必须为此账号下拥有的网络地址，可以与容器网络保持一致。</p>
	ApiServerNetwork *ApiServerNetwork `json:"ApiServerNetwork,omitnil,omitempty" name:"ApiServerNetwork"`

	// <p>集群描述</p>
	ClusterDescription *string `json:"ClusterDescription,omitnil,omitempty" name:"ClusterDescription"`

	// <p>集群标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>客户端Token</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

func (r *CreateDBCustomClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBCustomClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContainerNetwork")
	delete(f, "ClusterName")
	delete(f, "ApiServerNetwork")
	delete(f, "ClusterDescription")
	delete(f, "Tags")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBCustomClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBCustomClusterResponseParams struct {
	// <p>本次创建的集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>本次创建集群的任务ID</p>
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBCustomClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBCustomClusterResponseParams `json:"Response"`
}

func (r *CreateDBCustomClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBCustomClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBCustomNodesRequestParams struct {
	// <p>产品支持的可用区</p><p>枚举值：</p><ul><li>ap-shanghai-5： 上海五区</li><li>ap-shanghai-8： 上海八区</li><li>ap-nanjing-3： 南京三区</li></ul>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>镜像ID</p><p>参数格式：img-xxxxxxx</p><p>入参限制：必须为当前账号下DB Custom 产品拥有的镜像</p><p>取值参考：可通过&quot;DescribeDBCustomImages&quot;接口获取支持的镜像。</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>为节点打通SSH连接的VPC 网络ID。</p><p>参数格式：vpc-b4zgtest</p><p>入参限制：必须是当前账号下拥有的VPC 网络ID，且不能跨地域。</p><p>取值参考：可通过【查询VPC列表】接口获取：https://cloud.tencent.com/document/product/215/15778</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>为节点打通SSH连接的VPC 子网 ID。 </p><p>参数格式：subnet-t13dtest</p><p>入参限制：必须是VPC之下的子网，子网必须与可用区对应。</p><p>取值参考：可通过【查询子网列表】接口获取：https://cloud.tencent.com/document/product/215/15784</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>购买时长(月): 1/2/3/4/5/6/7/8/9/10/11/12/24/36</p><p>取值范围：[1, 36]</p><p>单位：月</p><p>默认值：1</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>节点机型</p><p>枚举值：</p><ul><li>DB.AT5.32XLARGE512： 高IO型服务器：128核心512GB内存，8*7180GB本地NvME SSDB。</li><li>DB.AT5.64XLARGE1152： 高IO型服务器：256核心1152GB内存，12*7180GB本地NvME SSDB。</li><li>DB.AT5.128XLARGE2304： 高IO型服务器：512核心2304GB内存，24*7180GB本地NvME SSDB。</li><li>DB.AT5.16XLARGE256： 高IO型服务器：64核心256GB内存，4*7180GB本地NvME SSDB。</li><li>DB.AT5.8XLARGE128： 高IO型服务器：32核心128GB内存，2*7180GB本地NvME SSDB。</li></ul>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>购买的节点数量</p><p>取值范围：[1, 20]</p>
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// <p>实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。</p><p>入参限制：若选择密钥方式，KeyIds 仅支持单个 ID。三种方式必须且仅可以设置其中一种。</p>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// <p>自动续费配置</p><p>枚举值：</p><ul><li>1： 自动续费</li><li>2： 不自动续费</li></ul><p>默认值：不自动续费</p>
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>节点名称</p><p>入参限制：最多128个字符</p>
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// <p>是否使用代金券自动抵扣</p><p>枚举值：</p><ul><li>1： 使用</li><li>0： 不使用</li></ul><p>默认值：0</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>代金券ID</p><p>入参限制：必须为当前账号下拥有的未抵扣的代金券ID。</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

type CreateDBCustomNodesRequest struct {
	*tchttp.BaseRequest
	
	// <p>产品支持的可用区</p><p>枚举值：</p><ul><li>ap-shanghai-5： 上海五区</li><li>ap-shanghai-8： 上海八区</li><li>ap-nanjing-3： 南京三区</li></ul>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>镜像ID</p><p>参数格式：img-xxxxxxx</p><p>入参限制：必须为当前账号下DB Custom 产品拥有的镜像</p><p>取值参考：可通过&quot;DescribeDBCustomImages&quot;接口获取支持的镜像。</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>为节点打通SSH连接的VPC 网络ID。</p><p>参数格式：vpc-b4zgtest</p><p>入参限制：必须是当前账号下拥有的VPC 网络ID，且不能跨地域。</p><p>取值参考：可通过【查询VPC列表】接口获取：https://cloud.tencent.com/document/product/215/15778</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>为节点打通SSH连接的VPC 子网 ID。 </p><p>参数格式：subnet-t13dtest</p><p>入参限制：必须是VPC之下的子网，子网必须与可用区对应。</p><p>取值参考：可通过【查询子网列表】接口获取：https://cloud.tencent.com/document/product/215/15784</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>购买时长(月): 1/2/3/4/5/6/7/8/9/10/11/12/24/36</p><p>取值范围：[1, 36]</p><p>单位：月</p><p>默认值：1</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>节点机型</p><p>枚举值：</p><ul><li>DB.AT5.32XLARGE512： 高IO型服务器：128核心512GB内存，8*7180GB本地NvME SSDB。</li><li>DB.AT5.64XLARGE1152： 高IO型服务器：256核心1152GB内存，12*7180GB本地NvME SSDB。</li><li>DB.AT5.128XLARGE2304： 高IO型服务器：512核心2304GB内存，24*7180GB本地NvME SSDB。</li><li>DB.AT5.16XLARGE256： 高IO型服务器：64核心256GB内存，4*7180GB本地NvME SSDB。</li><li>DB.AT5.8XLARGE128： 高IO型服务器：32核心128GB内存，2*7180GB本地NvME SSDB。</li></ul>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>购买的节点数量</p><p>取值范围：[1, 20]</p>
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// <p>实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。</p><p>入参限制：若选择密钥方式，KeyIds 仅支持单个 ID。三种方式必须且仅可以设置其中一种。</p>
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// <p>自动续费配置</p><p>枚举值：</p><ul><li>1： 自动续费</li><li>2： 不自动续费</li></ul><p>默认值：不自动续费</p>
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>节点名称</p><p>入参限制：最多128个字符</p>
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// <p>是否使用代金券自动抵扣</p><p>枚举值：</p><ul><li>1： 使用</li><li>0： 不使用</li></ul><p>默认值：0</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>代金券ID</p><p>入参限制：必须为当前账号下拥有的未抵扣的代金券ID。</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// <p>标签</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

func (r *CreateDBCustomNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBCustomNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "ImageId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Period")
	delete(f, "NodeType")
	delete(f, "NodeCount")
	delete(f, "LoginSettings")
	delete(f, "AutoRenew")
	delete(f, "NodeName")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "Tags")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBCustomNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBCustomNodesResponseParams struct {
	// <p>节点ID列表</p>
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`

	// <p>创建节点的任务ID</p>
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBCustomNodesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBCustomNodesResponseParams `json:"Response"`
}

func (r *CreateDBCustomNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBCustomNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBCustomCluster struct {
	// <p>集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>集群名称</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>集群支持的地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>集群规模</p><p>默认值：L500</p>
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// <p>DB Custom 集群状态</p><p>枚举值：</p><ul><li>Creating： 创建中</li><li>Running： 运行中</li><li>Destroying： 销毁中</li></ul>
	ClusterStatus *string `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// <p>集群版本号</p>
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// <p>集群中的节点数量</p><p>单位：台</p>
	ClusterNodeNum *int64 `json:"ClusterNodeNum,omitnil,omitempty" name:"ClusterNodeNum"`

	// <p>集群描述</p>
	ClusterDescription *string `json:"ClusterDescription,omitnil,omitempty" name:"ClusterDescription"`

	// <p>创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>集群的标签信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DBCustomClusterNode struct {
	// <p>节点ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>节点名称</p>
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// <p>节点内网IP地址</p>
	LanIP *string `json:"LanIP,omitnil,omitempty" name:"LanIP"`

	// <p>节点SSH 访问的Endpoint。格式为IP:Port.</p>
	SSHEndpoint *string `json:"SSHEndpoint,omitnil,omitempty" name:"SSHEndpoint"`

	// <p>节点在集群中的实例状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>节点所属的地域</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>节点类型</p><p>枚举值：</p><ul><li>DB.AT5.32XLARGE512： 高IO型服务器：128核心512GB内存，8*7180GB本地NvME SSDB。</li><li>DB.AT5.64XLARGE1152： 高IO型服务器：256核心1152GB内存，12*7180GB本地NvME SSDB。</li><li>DB.AT5.128XLARGE2304： 高IO型服务器：512核心2304GB内存，24*7180GB本地NvME SSDB。</li><li>DB.AT5.16XLARGE256： 高IO型服务器：64核心256GB内存，4*7180GB本地NvME SSDB。</li><li>DB.AT5.8XLARGE128： 高IO型服务器：32核心128GB内存，2*7180GB本地NvME SSDB。</li></ul>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`
}

type DBCustomImage struct {
	// <p>镜像ID</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>操作系统名称</p>
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// <p>镜像类型</p><p>枚举值：</p><ul><li>PUBLIC_IMAGE： 公共镜像 (腾讯云官方镜像)</li><li>PRIVATE_IMAGE： 私有镜像 (客户专属镜像)</li></ul>
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// <p>操作系统架构</p><p>枚举值：</p><ul><li>x86_64： X86 64位架构</li><li>arm64： ARM 64位机构</li></ul>
	Architecture *string `json:"Architecture,omitnil,omitempty" name:"Architecture"`
}

type DBCustomNode struct {
	// <p>节点ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>节点名称</p>
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// <p>访问此节点的SSH Endpoint，格式为IP:Port</p>
	SSHEndpoint *string `json:"SSHEndpoint,omitnil,omitempty" name:"SSHEndpoint"`

	// <p>节点的内网通信IP地址</p>
	LanIP *string `json:"LanIP,omitnil,omitempty" name:"LanIP"`

	// <p>节点所属的集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>节点所属的可用区</p>
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// <p>节点类型</p><p>枚举值：</p><ul><li>DB.AT5.32XLARGE512： 高IO型服务器：128核心512GB内存，8*7180GB本地NvME SSDB。</li><li>DB.AT5.64XLARGE1152： 高IO型服务器：256核心1152GB内存，12*7180GB本地NvME SSDB。</li><li>DB.AT5.128XLARGE2304： 高IO型服务器：512核心2304GB内存，24*7180GB本地NvME SSDB。</li><li>DB.AT5.16XLARGE256： 高IO型服务器：64核心256GB内存，4*7180GB本地NvME SSDB。</li><li>DB.AT5.8XLARGE128： 高IO型服务器：32核心128GB内存，2*7180GB本地NvME SSDB。</li></ul>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// <p>节点CPU大小</p><p>单位：核</p>
	CPU *int64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// <p>节点内存</p><p>单位：GiB</p>
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// <p>系统盘信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// <p>数据盘信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// <p>节点的操作系统名称</p>
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// <p>节点的操作系统镜像ID</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>节点SSH Endpoint 所属的VPC ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>节点SSH Endpoint 所属的VPC 子网ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>节点状态</p><p>枚举值：</p><ul><li>Creating： 创建中</li><li>Running： 运行中</li><li>Isolating： 隔离中</li><li>Isolated： 已隔离</li><li>Activating： 解除隔离中</li><li>Destroying： 销毁中</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>付费类型</p><p>枚举值：</p><ul><li>PREPAID： 包年包月</li></ul>
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// <p>节点到期时间</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>节点创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>节点隔离时间</p>
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// <p>节点标签信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>节点是否自动续费标记</p><p>枚举值：</p><ul><li>1： 自动续费</li><li>2： 不自动续费</li></ul>
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>交换机ID（已加密）</p>
	SwitchId *string `json:"SwitchId,omitnil,omitempty" name:"SwitchId"`

	// <p>机架ID（已加密）</p>
	RackId *string `json:"RackId,omitnil,omitempty" name:"RackId"`

	// <p>底层物理机IP（已加密）</p>
	HostIp *string `json:"HostIp,omitnil,omitempty" name:"HostIp"`
}

type DBInstanceDetail struct {
	// DB实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// DB实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// DB实例状态,-1:已隔离, 0:创建中, 1:流程中, 2:运行中, 3:未初始化
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// DB实例状态描述,-1:已隔离, 0:创建中, 1:流程中, 2:运行中, 3:未初始化
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// DB实例版本
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// Vip信息
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Vip使用的端口号
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 字符串型的私有网络ID
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// 字符串型的私有网络子网ID
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// 是否为分布式版本,0:否,1:是
	Shard *int64 `json:"Shard,omitnil,omitempty" name:"Shard"`

	// DB实例节点数
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// CPU规格(单位:核数)
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存规格(单位:GB)
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘规格(单位:GB)
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// 分布式类型的实例的分片数
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Db所在主机列表, 格式: m1,s1|m2,s2
	DbHosts *string `json:"DbHosts,omitnil,omitempty" name:"DbHosts"`

	// 主机角色, 1:主, 2:从, 3:主+从
	HostRole *int64 `json:"HostRole,omitnil,omitempty" name:"HostRole"`

	// DB引擎，MySQL,Percona,MariaDB
	DbEngine *string `json:"DbEngine,omitnil,omitempty" name:"DbEngine"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 可用区列表
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`
}

type DataDisk struct {
	// <p>磁盘类型</p><p>枚举值：</p><ul><li>CLOUD_HSSD： 增强型云硬盘</li><li>LOCAL_NVME： 本地硬盘</li></ul>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>磁盘大小</p><p>单位：GiB</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// <p>磁盘名称</p>
	DiskName *string `json:"DiskName,omitnil,omitempty" name:"DiskName"`
}

// Predefined struct for user
type DescribeDBCustomClusterDetailRequestParams struct {
	// <p>DB Custom 集群ID</p><p>入参限制：必须为此账号拥有的DB Custom集群</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeDBCustomClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>DB Custom 集群ID</p><p>入参限制：必须为此账号拥有的DB Custom集群</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeDBCustomClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBCustomClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomClusterDetailResponseParams struct {
	// <p>集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>集群名称</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>集群描述</p>
	ClusterDescription *string `json:"ClusterDescription,omitnil,omitempty" name:"ClusterDescription"`

	// <p>集群所属地域</p><p>枚举值：</p><ul><li>ap-shanghai： 上海地域</li><li>ap-nanjing： 南京地域</li></ul>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>DB Custom 集群状态</p><p>枚举值：</p><ul><li>Creating： 创建中</li><li>Running： 运行中</li><li>Destroying： 销毁中</li><li>Initializing： 初始化中</li></ul>
	ClusterStatus *string `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// <p>集群版本</p><p>枚举值：</p><ul><li>1.34.1： 集群版本1.34.1</li></ul><p>默认值：1.34.1</p>
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// <p>集群下的节点数量</p>
	ClusterNodeNum *int64 `json:"ClusterNodeNum,omitnil,omitempty" name:"ClusterNodeNum"`

	// <p>集群规模</p>
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// <p>创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>集群标签信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>集群的API Server的网络信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiServerNetwork *ApiServerNetwork `json:"ApiServerNetwork,omitnil,omitempty" name:"ApiServerNetwork"`

	// <p>容器网络，在此集群中的所有Pod将与此网络连通</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerNetwork *ContainerNetwork `json:"ContainerNetwork,omitnil,omitempty" name:"ContainerNetwork"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBCustomClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBCustomClusterDetailResponseParams `json:"Response"`
}

func (r *DescribeDBCustomClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomClusterKubeconfigRequestParams struct {
	// <p>集群ID</p><p>入参限制：必须为当前节点拥有的DB Custom 集群</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeDBCustomClusterKubeconfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>集群ID</p><p>入参限制：必须为当前节点拥有的DB Custom 集群</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeDBCustomClusterKubeconfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomClusterKubeconfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBCustomClusterKubeconfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomClusterKubeconfigResponseParams struct {
	// <p>集群APIServer信息</p>
	Kubeconfig *string `json:"Kubeconfig,omitnil,omitempty" name:"Kubeconfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBCustomClusterKubeconfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBCustomClusterKubeconfigResponseParams `json:"Response"`
}

func (r *DescribeDBCustomClusterKubeconfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomClusterKubeconfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomClusterNodesRequestParams struct {
	// <p>DB Custom 集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>查询筛选条件。支持的条件有：</p><ul><li>node-name：DB Custom 节点名称。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分页偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBCustomClusterNodesRequest struct {
	*tchttp.BaseRequest
	
	// <p>DB Custom 集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>查询筛选条件。支持的条件有：</p><ul><li>node-name：DB Custom 节点名称。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分页偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBCustomClusterNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomClusterNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBCustomClusterNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomClusterNodesResponseParams struct {
	// <p>集群下总的节点数量</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>分页后节点列表信息</p>
	NodeSet []*DBCustomClusterNode `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBCustomClusterNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBCustomClusterNodesResponseParams `json:"Response"`
}

func (r *DescribeDBCustomClusterNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomClusterNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomClustersRequestParams struct {
	// <p>按照一个或者多个 ClusterId 查询。</p><p>入参限制：每次请求的数量上限为100</p>
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// <p>查询筛选条件。支持的筛选条件包括：</p><ul><li>cluster-name：DB Custom 集群名称，精确匹配。</li><li>cluster-status：DB Custom 集群状态（Creating，Running，Destroying）。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>根据标签的 Key 和 Value 筛选 DB Custom 集群</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>集群列表分页偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBCustomClustersRequest struct {
	*tchttp.BaseRequest
	
	// <p>按照一个或者多个 ClusterId 查询。</p><p>入参限制：每次请求的数量上限为100</p>
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// <p>查询筛选条件。支持的筛选条件包括：</p><ul><li>cluster-name：DB Custom 集群名称，精确匹配。</li><li>cluster-status：DB Custom 集群状态（Creating，Running，Destroying）。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>根据标签的 Key 和 Value 筛选 DB Custom 集群</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>集群列表分页偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBCustomClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Filters")
	delete(f, "Tags")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBCustomClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomClustersResponseParams struct {
	// <p>总集群数量</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>集群列表信息</p>
	ClusterSet []*DBCustomCluster `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBCustomClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBCustomClustersResponseParams `json:"Response"`
}

func (r *DescribeDBCustomClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomImagesRequestParams struct {
	// <p>偏移量</p><p>默认值：0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBCustomImagesRequest struct {
	*tchttp.BaseRequest
	
	// <p>偏移量</p><p>默认值：0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBCustomImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBCustomImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomImagesResponseParams struct {
	// <p>总镜像数量</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>支持的镜像列表信息</p>
	ImageSet []*DBCustomImage `json:"ImageSet,omitnil,omitempty" name:"ImageSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBCustomImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBCustomImagesResponseParams `json:"Response"`
}

func (r *DescribeDBCustomImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomNodesRequestParams struct {
	// <p>按照一个或者多个 NodeId 查询。</p><p>入参限制：每次请求的数量上限为100</p>
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`

	// <p>查询筛选条件。支持的筛选条件包括：</p><ul><li>cluster-id：按 DB Custom 集群进行过滤。</li><li>node-name：按 DB Custom 节点名称进行过滤，精确匹配。</li><li>status：按 DB Custom 节点状态进行过滤。（可选值：Creating，Running，Isolating，Isolated，Activating（解除隔离中），Destroying）</li><li>zone：按 DB Custom 节点所在可用区进行过滤。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>根据标签的 Key 和 Value 筛选 DB Custom 节点</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>分页偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBCustomNodesRequest struct {
	*tchttp.BaseRequest
	
	// <p>按照一个或者多个 NodeId 查询。</p><p>入参限制：每次请求的数量上限为100</p>
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`

	// <p>查询筛选条件。支持的筛选条件包括：</p><ul><li>cluster-id：按 DB Custom 集群进行过滤。</li><li>node-name：按 DB Custom 节点名称进行过滤，精确匹配。</li><li>status：按 DB Custom 节点状态进行过滤。（可选值：Creating，Running，Isolating，Isolated，Activating（解除隔离中），Destroying）</li><li>zone：按 DB Custom 节点所在可用区进行过滤。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>根据标签的 Key 和 Value 筛选 DB Custom 节点</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>分页偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBCustomNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeIds")
	delete(f, "Filters")
	delete(f, "Tags")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBCustomNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomNodesResponseParams struct {
	// <p>当前账号下拥有的DB Custom 节点总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>当前账号下拥有的DB Custom 节点列表信息</p>
	NodeSet []*DBCustomNode `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBCustomNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBCustomNodesResponseParams `json:"Response"`
}

func (r *DescribeDBCustomNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomTaskStatusRequestParams struct {
	// <p>DB Custom 任务ID</p>
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeDBCustomTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>DB Custom 任务ID</p>
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeDBCustomTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBCustomTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCustomTaskStatusResponseParams struct {
	// <p>任务 ID</p><p>枚举值：</p><ul><li>Running： 运行中</li><li>Succeeded： 成功</li><li>Failed： 失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBCustomTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBCustomTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeDBCustomTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCustomTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群主机Id
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 实例类型,0:mariadb, 1:tdsql
	ShardType []*int64 `json:"ShardType,omitnil,omitempty" name:"ShardType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群主机Id
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 实例类型,0:mariadb, 1:tdsql
	ShardType []*int64 `json:"ShardType,omitnil,omitempty" name:"ShardType"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "HostId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ShardType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// 独享集群内的DB实例列表
	Instances []*DBInstanceDetail `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 独享集群内的DB实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostListRequestParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分配状态过滤，0-可分配，1-禁止分配
	AssignStatus []*int64 `json:"AssignStatus,omitnil,omitempty" name:"AssignStatus"`
}

type DescribeHostListRequest struct {
	*tchttp.BaseRequest
	
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分配状态过滤，0-可分配，1-禁止分配
	AssignStatus []*int64 `json:"AssignStatus,omitnil,omitempty" name:"AssignStatus"`
}

func (r *DescribeHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AssignStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostListResponseParams struct {
	// 主机总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主机详情
	Hosts []*HostDetail `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostListResponseParams `json:"Response"`
}

func (r *DescribeHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetail struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 产品ID, 0:CDB, 1:TDSQL
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 集群类型, 0:公有云, 1:金融围笼, 2:CDC集群
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 主机类型, 0:物理机, 1:CVM机型, 2:CDC机型
	HostType *int64 `json:"HostType,omitnil,omitempty" name:"HostType"`

	// 自动续费标志, 0:未设置, 1:自动续费, 2:到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 集群状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 主机数
	HostNum *int64 `json:"HostNum,omitnil,omitempty" name:"HostNum"`

	// DB实例数
	DbNum *int64 `json:"DbNum,omitnil,omitempty" name:"DbNum"`

	// 分配策略, 0:紧凑, 1:均匀
	AssignStrategy *int64 `json:"AssignStrategy,omitnil,omitempty" name:"AssignStrategy"`

	// 总主机CPU(单位:核数)
	CpuSpec *int64 `json:"CpuSpec,omitnil,omitempty" name:"CpuSpec"`

	// 总已分配CPU(单位:核数)
	CpuAssigned *int64 `json:"CpuAssigned,omitnil,omitempty" name:"CpuAssigned"`

	// 总可分配CPU(单位:核数)
	CpuAssignable *int64 `json:"CpuAssignable,omitnil,omitempty" name:"CpuAssignable"`

	// 总主机内存(单位:GB)
	MemorySpec *int64 `json:"MemorySpec,omitnil,omitempty" name:"MemorySpec"`

	// 总已分配内存(单位:GB)
	MemoryAssigned *int64 `json:"MemoryAssigned,omitnil,omitempty" name:"MemoryAssigned"`

	// 总可分配内存(单位:GB)
	MemoryAssignable *int64 `json:"MemoryAssignable,omitnil,omitempty" name:"MemoryAssignable"`

	// 总机器磁盘(单位:GB)
	DiskSpec *int64 `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 总已分配磁盘(单位:GB)
	DiskAssigned *int64 `json:"DiskAssigned,omitnil,omitempty" name:"DiskAssigned"`

	// 总可分配磁盘(单位:GB)
	DiskAssignable *int64 `json:"DiskAssignable,omitnil,omitempty" name:"DiskAssignable"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 金融围笼ID
	FenceId *string `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 所属集群ID(默认集群为空)
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例标签
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// CPU类型：Intel/AMD,Hygon
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// 可用区列表
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`
}

// Predefined struct for user
type DescribeInstanceDetailRequestParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDetailResponseParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 产品ID, 0:CDB, 1:TDSQL
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 集群类型, 0:公有云, 1:金融围笼
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 主机类型, 0:物理机, 1:cvm本地盘, 2:cvm云盘
	HostType *int64 `json:"HostType,omitnil,omitempty" name:"HostType"`

	// 自动续费标志, 0:未设置, 1:自动续费, 2:到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 集群状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 主机数
	HostNum *int64 `json:"HostNum,omitnil,omitempty" name:"HostNum"`

	// Db实例数
	DbNum *int64 `json:"DbNum,omitnil,omitempty" name:"DbNum"`

	// 分配策略, 0:紧凑, 1:均匀
	AssignStrategy *int64 `json:"AssignStrategy,omitnil,omitempty" name:"AssignStrategy"`

	// 总主机CPU(单位:核)
	CpuSpec *int64 `json:"CpuSpec,omitnil,omitempty" name:"CpuSpec"`

	// 总已分配CPU(单位:核)
	CpuAssigned *int64 `json:"CpuAssigned,omitnil,omitempty" name:"CpuAssigned"`

	// 总可分配CPU(单位:核)
	CpuAssignable *int64 `json:"CpuAssignable,omitnil,omitempty" name:"CpuAssignable"`

	// 总主机内存(单位:GB)
	MemorySpec *int64 `json:"MemorySpec,omitnil,omitempty" name:"MemorySpec"`

	// 总已分配内存(单位:GB)
	MemoryAssigned *int64 `json:"MemoryAssigned,omitnil,omitempty" name:"MemoryAssigned"`

	// 总可分配内存(单位:GB)
	MemoryAssignable *int64 `json:"MemoryAssignable,omitnil,omitempty" name:"MemoryAssignable"`

	// 总机器磁盘(单位:GB)
	DiskSpec *int64 `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 总已分配磁盘(单位:GB)
	DiskAssigned *int64 `json:"DiskAssigned,omitnil,omitempty" name:"DiskAssigned"`

	// 总可分配磁盘(单位:GB)
	DiskAssignable *int64 `json:"DiskAssignable,omitnil,omitempty" name:"DiskAssignable"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 金融围笼ID
	FenceId *string `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 所属集群ID(默认集群为空)
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 独享集群的标签信息
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// CPU类型，Intel/AMD,Hygon
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// 可用区列表
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListRequestParams struct {
	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，createTime,instancename两者之一
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序规则，desc,asc两者之一
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 按产品过滤，0:CDB, 1:TDSQL
	ProductId []*int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 按实例ID过滤
	InstanceId []*string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按实例名称过滤
	InstanceName []*string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 按金融围笼ID过滤
	FenceId []*string `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 按实例状态过滤, -1:已隔离, 0:创建中, 1:运行中, 2:扩容中, 3:删除中
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 按所属集群ID过滤
	ClusterId []*string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 分页返回数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，createTime,instancename两者之一
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序规则，desc,asc两者之一
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 按产品过滤，0:CDB, 1:TDSQL
	ProductId []*int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 按实例ID过滤
	InstanceId []*string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按实例名称过滤
	InstanceName []*string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 按金融围笼ID过滤
	FenceId []*string `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 按实例状态过滤, -1:已隔离, 0:创建中, 1:运行中, 2:扩容中, 3:删除中
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 按所属集群ID过滤
	ClusterId []*string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "SortBy")
	delete(f, "ProductId")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "FenceId")
	delete(f, "Status")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListResponseParams struct {
	// 独享集群列表
	Instances []*DescribeInstanceDetail `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 独享集群实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceListResponseParams `json:"Response"`
}

func (r *DescribeInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 集群类型: 0 一主一备, 1 一主两备...N-1 一主N备
	InstanceTypes []*int64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 产品ID:  0 MYSQL，1 TDSQL
	ProductIds []*int64 `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`

	// 集群uuid: 如 dbdc-q810131s
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否按金融围笼标志搜索
	FenceFlag *bool `json:"FenceFlag,omitnil,omitempty" name:"FenceFlag"`

	// 按实例名字模糊匹配
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 每页数目, 整型
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码, 整型
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 排序字段，枚举：createtime,groupname
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式: asc升序, desc降序
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 集群状态: -2 已删除, -1 已隔离, 0 创建中, 1 运行中, 2 扩容中, 3 删除中
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群类型: 0 一主一备, 1 一主两备...N-1 一主N备
	InstanceTypes []*int64 `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 产品ID:  0 MYSQL，1 TDSQL
	ProductIds []*int64 `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`

	// 集群uuid: 如 dbdc-q810131s
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否按金融围笼标志搜索
	FenceFlag *bool `json:"FenceFlag,omitnil,omitempty" name:"FenceFlag"`

	// 按实例名字模糊匹配
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 每页数目, 整型
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码, 整型
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 排序字段，枚举：createtime,groupname
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式: asc升序, desc降序
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 集群状态: -2 已删除, -1 已隔离, 0 创建中, 1 运行中, 2 扩容中, 3 删除中
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
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
	delete(f, "InstanceTypes")
	delete(f, "ProductIds")
	delete(f, "InstanceIds")
	delete(f, "FenceFlag")
	delete(f, "InstanceName")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "InstanceStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 集群数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群扩展信息
	Instances []*InstanceExpand `json:"Instances,omitnil,omitempty" name:"Instances"`

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
type DestroyDBCustomClusterRequestParams struct {
	// <p>待销毁的集群ID</p><p>入参限制：待销毁的集群必须无任何节点在此集群中。</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DestroyDBCustomClusterRequest struct {
	*tchttp.BaseRequest
	
	// <p>待销毁的集群ID</p><p>入参限制：待销毁的集群必须无任何节点在此集群中。</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DestroyDBCustomClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBCustomClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyDBCustomClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDBCustomClusterResponseParams struct {
	// <p>销毁集群的任务ID</p>
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyDBCustomClusterResponse struct {
	*tchttp.BaseResponse
	Response *DestroyDBCustomClusterResponseParams `json:"Response"`
}

func (r *DestroyDBCustomClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBCustomClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDBCustomNodeRequestParams struct {
	// <p>DB Custom 节点ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type DestroyDBCustomNodeRequest struct {
	*tchttp.BaseRequest
	
	// <p>DB Custom 节点ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

func (r *DestroyDBCustomNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBCustomNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyDBCustomNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDBCustomNodeResponseParams struct {
	// <p>任务ID</p>
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyDBCustomNodeResponse struct {
	*tchttp.BaseResponse
	Response *DestroyDBCustomNodeResponseParams `json:"Response"`
}

func (r *DestroyDBCustomNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBCustomNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceInfo struct {
	// 设备ID
	DeviceId *int64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备No
	DeviceNo *string `json:"DeviceNo,omitnil,omitempty" name:"DeviceNo"`

	// 设备类型
	DevClass *string `json:"DevClass,omitnil,omitempty" name:"DevClass"`

	// 设备总内存，单位GB
	MaxMemory *float64 `json:"MaxMemory,omitnil,omitempty" name:"MaxMemory"`

	// 设备总磁盘，单位GB
	MaxDisk *float64 `json:"MaxDisk,omitnil,omitempty" name:"MaxDisk"`

	// 设备剩余内存，单位GB
	RestMemory *float64 `json:"RestMemory,omitnil,omitempty" name:"RestMemory"`

	// 设备剩余磁盘，单位GB
	RestDisk *float64 `json:"RestDisk,omitnil,omitempty" name:"RestDisk"`

	// 设备机器个数
	RawDeviceNum *uint64 `json:"RawDeviceNum,omitnil,omitempty" name:"RawDeviceNum"`

	// 数据库实例个数
	InstanceNum *uint64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`
}

type Filter struct {
	// <p>筛选条件</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>过滤字段对应的参数值</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type HostDetail struct {
	// 主机Id
	HostId *string `json:"HostId,omitnil,omitempty" name:"HostId"`

	// 主机名称
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 主机状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 分配DB实例状态,0:可分配,1:不可分配
	AssignStatus *int64 `json:"AssignStatus,omitnil,omitempty" name:"AssignStatus"`

	// 主机类型, 0:物理机, 1:cvm本地盘, 2:cvm云盘
	HostType *int64 `json:"HostType,omitnil,omitempty" name:"HostType"`

	// DB实例数
	DbNum *int64 `json:"DbNum,omitnil,omitempty" name:"DbNum"`

	// 主机CPU(单位:核数)
	CpuSpec *int64 `json:"CpuSpec,omitnil,omitempty" name:"CpuSpec"`

	// 已分配CPU(单位:核数)
	CpuAssigned *int64 `json:"CpuAssigned,omitnil,omitempty" name:"CpuAssigned"`

	// 可分配CPU(单位:核数)
	CpuAssignable *int64 `json:"CpuAssignable,omitnil,omitempty" name:"CpuAssignable"`

	// 主机内存(单位:GB)
	MemorySpec *int64 `json:"MemorySpec,omitnil,omitempty" name:"MemorySpec"`

	// 已分配内存(单位:GB)
	MemoryAssigned *int64 `json:"MemoryAssigned,omitnil,omitempty" name:"MemoryAssigned"`

	// 可分配内存(单位:GB)
	MemoryAssignable *int64 `json:"MemoryAssignable,omitnil,omitempty" name:"MemoryAssignable"`

	// 主机磁盘(单位:GB)
	DiskSpec *int64 `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 已分配磁盘(单位:GB)
	DiskAssigned *int64 `json:"DiskAssigned,omitnil,omitempty" name:"DiskAssigned"`

	// 可分配磁盘(GB)
	DiskAssignable *int64 `json:"DiskAssignable,omitnil,omitempty" name:"DiskAssignable"`

	// CPU分配比
	CpuRatio *float64 `json:"CpuRatio,omitnil,omitempty" name:"CpuRatio"`

	// 内存分配比
	MemoryRatio *float64 `json:"MemoryRatio,omitnil,omitempty" name:"MemoryRatio"`

	// 磁盘分配比
	DiskRatio *float64 `json:"DiskRatio,omitnil,omitempty" name:"DiskRatio"`

	// 机型名称
	MachineName *string `json:"MachineName,omitnil,omitempty" name:"MachineName"`

	// 机型类别
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// 计费标签
	PidTag *string `json:"PidTag,omitnil,omitempty" name:"PidTag"`

	// 计费ID
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 可用区列表
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`
}

type InstanceDetail struct {
	// 集群状态，0：运行中，1：不在运行
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 读写集群剩余内存容量，单位GB
	ReadWriteTotalLeaveMemory *float64 `json:"ReadWriteTotalLeaveMemory,omitnil,omitempty" name:"ReadWriteTotalLeaveMemory"`

	// 读写集群剩余磁盘容量，单位GB
	ReadWriteTotalLeaveDisk *float64 `json:"ReadWriteTotalLeaveDisk,omitnil,omitempty" name:"ReadWriteTotalLeaveDisk"`

	// 读写集群总内存容量，单位GB
	ReadWriteTotalMemory *float64 `json:"ReadWriteTotalMemory,omitnil,omitempty" name:"ReadWriteTotalMemory"`

	// 读写集群总磁盘容量，单位GB
	ReadWriteTotalDisk *float64 `json:"ReadWriteTotalDisk,omitnil,omitempty" name:"ReadWriteTotalDisk"`

	// 只读集群剩余内存容量，单位GB
	ReadOnlyTotalLeaveMemory *float64 `json:"ReadOnlyTotalLeaveMemory,omitnil,omitempty" name:"ReadOnlyTotalLeaveMemory"`

	// 只读集群剩余磁盘容量，单位GB
	ReadOnlyTotalLeaveDisk *float64 `json:"ReadOnlyTotalLeaveDisk,omitnil,omitempty" name:"ReadOnlyTotalLeaveDisk"`

	// 只读集群总内存容量，单位GB
	ReadOnlyTotalMemory *float64 `json:"ReadOnlyTotalMemory,omitnil,omitempty" name:"ReadOnlyTotalMemory"`

	// 只读集群总磁盘容量，单位GB
	ReadOnlyTotalDisk *float64 `json:"ReadOnlyTotalDisk,omitnil,omitempty" name:"ReadOnlyTotalDisk"`

	// 集群设备详情
	InstanceDeviceInfos []*InstanceDeviceInfo `json:"InstanceDeviceInfos,omitnil,omitempty" name:"InstanceDeviceInfos"`
}

type InstanceDeviceInfo struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 读写设备组
	ReadWriteDevice []*DeviceInfo `json:"ReadWriteDevice,omitnil,omitempty" name:"ReadWriteDevice"`

	// 只读设备组
	ReadOnlyDevice []*DeviceInfo `json:"ReadOnlyDevice,omitnil,omitempty" name:"ReadOnlyDevice"`

	// 空闲设备组
	FreeDevice []*DeviceInfo `json:"FreeDevice,omitnil,omitempty" name:"FreeDevice"`
}

type InstanceExpand struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用户ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群类型： 0：一主一备，1：一主两备
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 集群状态: 0 集群创建中, 1 集群有效, 2 集群扩容中, 3 集群删除中, 4 集群缩容中 -1 集群已隔离 -2 集群已删除
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例自动续费标识： 0正常续费 1自动续费 2到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 机型
	Machine *string `json:"Machine,omitnil,omitempty" name:"Machine"`

	// 过期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 集群信息
	InstanceDetail *InstanceDetail `json:"InstanceDetail,omitnil,omitempty" name:"InstanceDetail"`

	// 计费侧的产品ID
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`
}

// Predefined struct for user
type IsolateDBCustomNodeRequestParams struct {
	// <p>DB Custom 节点ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type IsolateDBCustomNodeRequest struct {
	*tchttp.BaseRequest
	
	// <p>DB Custom 节点ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

func (r *IsolateDBCustomNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBCustomNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBCustomNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBCustomNodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDBCustomNodeResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDBCustomNodeResponseParams `json:"Response"`
}

func (r *IsolateDBCustomNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBCustomNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginSettings struct {
	// <p>实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下： Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) <code>~ ! @ # $ % ^ &amp; * - + = | { } [ ] : ; &#39; , . ? / ]中的特殊符号。 Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( )</code> ~ ! @ # $ % ^ &amp; * - + = | { } [ ] : ; &#39; , . ? /]中的特殊符号。</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口 DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。</p><p>入参限制：当前仅支持设置单个 ID。</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// <p>保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为true。</p><p>枚举值：</p><ul><li>true： 表示保持镜像的登录设置</li><li>false： 表示不保持镜像的登录设置</li></ul>
	KeepImageLogin *string `json:"KeepImageLogin,omitnil,omitempty" name:"KeepImageLogin"`
}

// Predefined struct for user
type ModifyDBCustomClusterTagsRequestParams struct {
	// <p>DB Custom 集群ID</p><p>参数格式：dbcc-xxxxxxxx</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>为 DB Custom 集群绑定的标签信息</p><p>入参限制：参考标签平台的限制策略</p>
	AddTags []*Tag `json:"AddTags,omitnil,omitempty" name:"AddTags"`

	// <p>为 DB Custom 集群删除的标签Key</p>
	DeleteTagKeys []*string `json:"DeleteTagKeys,omitnil,omitempty" name:"DeleteTagKeys"`
}

type ModifyDBCustomClusterTagsRequest struct {
	*tchttp.BaseRequest
	
	// <p>DB Custom 集群ID</p><p>参数格式：dbcc-xxxxxxxx</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>为 DB Custom 集群绑定的标签信息</p><p>入参限制：参考标签平台的限制策略</p>
	AddTags []*Tag `json:"AddTags,omitnil,omitempty" name:"AddTags"`

	// <p>为 DB Custom 集群删除的标签Key</p>
	DeleteTagKeys []*string `json:"DeleteTagKeys,omitnil,omitempty" name:"DeleteTagKeys"`
}

func (r *ModifyDBCustomClusterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBCustomClusterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AddTags")
	delete(f, "DeleteTagKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBCustomClusterTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBCustomClusterTagsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBCustomClusterTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBCustomClusterTagsResponseParams `json:"Response"`
}

func (r *ModifyDBCustomClusterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBCustomClusterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBCustomNodeTagsRequestParams struct {
	// <p>DB Custom 节点ID</p><p>参数格式：dbcn-0zan5xxk</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>为节点绑定的标签信息</p><p>入参限制：参考标签侧的限制</p>
	AddTags []*Tag `json:"AddTags,omitnil,omitempty" name:"AddTags"`

	// <p>需要删除的标签Key</p>
	DeleteTagKeys []*string `json:"DeleteTagKeys,omitnil,omitempty" name:"DeleteTagKeys"`
}

type ModifyDBCustomNodeTagsRequest struct {
	*tchttp.BaseRequest
	
	// <p>DB Custom 节点ID</p><p>参数格式：dbcn-0zan5xxk</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>为节点绑定的标签信息</p><p>入参限制：参考标签侧的限制</p>
	AddTags []*Tag `json:"AddTags,omitnil,omitempty" name:"AddTags"`

	// <p>需要删除的标签Key</p>
	DeleteTagKeys []*string `json:"DeleteTagKeys,omitnil,omitempty" name:"DeleteTagKeys"`
}

func (r *ModifyDBCustomNodeTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBCustomNodeTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "AddTags")
	delete(f, "DeleteTagKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBCustomNodeTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBCustomNodeTagsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBCustomNodeTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBCustomNodeTagsResponseParams `json:"Response"`
}

func (r *ModifyDBCustomNodeTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBCustomNodeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameRequestParams struct {
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 独享集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 独享集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveNodesFromDBCustomClusterRequestParams struct {
	// <p>DB Custom 集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>要下架的 DB Custom 节点ID列表</p>
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`
}

type RemoveNodesFromDBCustomClusterRequest struct {
	*tchttp.BaseRequest
	
	// <p>DB Custom 集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>要下架的 DB Custom 节点ID列表</p>
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`
}

func (r *RemoveNodesFromDBCustomClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNodesFromDBCustomClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveNodesFromDBCustomClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveNodesFromDBCustomClusterResponseParams struct {
	// <p>任务ID</p>
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveNodesFromDBCustomClusterResponse struct {
	*tchttp.BaseResponse
	Response *RemoveNodesFromDBCustomClusterResponseParams `json:"Response"`
}

func (r *RemoveNodesFromDBCustomClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNodesFromDBCustomClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBCustomNodeRequestParams struct {
	// <p>节点ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>续费周期</p><p>取值范围：[1, 36]</p><p>单位：月</p><p>默认值：1</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>是否开启自动续费</p><p>枚举值：</p><ul><li>0： 不自动续费</li><li>1： 自动续费</li></ul><p>默认值：1</p>
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>是否自动使用代金券</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>代金券ID</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

type RenewDBCustomNodeRequest struct {
	*tchttp.BaseRequest
	
	// <p>节点ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>续费周期</p><p>取值范围：[1, 36]</p><p>单位：月</p><p>默认值：1</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>是否开启自动续费</p><p>枚举值：</p><ul><li>0： 不自动续费</li><li>1： 自动续费</li></ul><p>默认值：1</p>
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// <p>是否自动使用代金券</p>
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>代金券ID</p>
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

func (r *RenewDBCustomNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBCustomNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "Period")
	delete(f, "AutoRenew")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDBCustomNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBCustomNodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewDBCustomNodeResponse struct {
	*tchttp.BaseResponse
	Response *RenewDBCustomNodeResponseParams `json:"Response"`
}

func (r *RenewDBCustomNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBCustomNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type SystemDisk struct {
	// <p>磁盘类型</p><p>枚举值：</p><ul><li>CLOUD_HSSD： 增强型云硬盘</li></ul>
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// <p>磁盘大小</p><p>单位：GiB</p>
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type Tag struct {
	// <p>标签键</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>标签值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}