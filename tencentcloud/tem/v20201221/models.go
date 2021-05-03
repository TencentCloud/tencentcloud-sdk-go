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

package v20201221

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CosToken struct {

	// 唯一请求 ID
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// 存储桶桶名
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 存储桶所在区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 临时密钥的SecretId
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥的SecretKey
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 临时密钥的 sessionToken
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// 临时密钥获取的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 临时密钥的 expiredTime
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 包完整路径
	FullPath *string `json:"FullPath,omitempty" name:"FullPath"`
}

type CreateCosTokenRequest struct {
	*tchttp.BaseRequest

	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务版本ID
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// optType 1上传  2查询
	OptType *int64 `json:"OptType,omitempty" name:"OptType"`

	// 来源 channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *CreateCosTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "VersionId")
	delete(f, "PkgName")
	delete(f, "OptType")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return errors.New("CreateCosTokenRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCosTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功时为CosToken对象，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *CosToken `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCosTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 命名空间描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// K8s version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceName")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "Description")
	delete(f, "K8sVersion")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return errors.New("CreateNamespaceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功时为命名空间ID，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespacesRequest struct {
	*tchttp.BaseRequest

	// 分页limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页下标
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 来源source
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return errors.New("DescribeNamespacesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *NamespacePage `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IngressInfo struct {

	// tem namespaceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// eks namespace
	EksNamespace *string `json:"EksNamespace,omitempty" name:"EksNamespace"`

	// ip version
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// ingress name
	Name *string `json:"Name,omitempty" name:"Name"`

	// rules 配置
	Rules []*IngressRule `json:"Rules,omitempty" name:"Rules" list`

	// clb ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClbId *string `json:"ClbId,omitempty" name:"ClbId"`

	// tls 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tls []*IngressTls `json:"Tls,omitempty" name:"Tls" list`

	// eks clusterId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// clb ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type IngressRule struct {

	// ingress rule value
	Http *IngressRuleValue `json:"Http,omitempty" name:"Http"`

	// host 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`
}

type IngressRuleBackend struct {

	// eks service 名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// eks service 端口
	ServicePort *int64 `json:"ServicePort,omitempty" name:"ServicePort"`
}

type IngressRulePath struct {

	// path 信息
	Path *string `json:"Path,omitempty" name:"Path"`

	// backend 配置
	Backend *IngressRuleBackend `json:"Backend,omitempty" name:"Backend"`
}

type IngressRuleValue struct {

	// rule 整体配置
	Paths []*IngressRulePath `json:"Paths,omitempty" name:"Paths" list`
}

type IngressTls struct {

	// host 数组
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts" list`

	// secret name
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`
}

type ModifyIngressRequest struct {
	*tchttp.BaseRequest

	// Ingress 规则配置
	Ingress *IngressInfo `json:"Ingress,omitempty" name:"Ingress"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *ModifyIngressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIngressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ingress")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return errors.New("ModifyIngressRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIngressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建成功
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIngressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNamespaceRequest struct {
	*tchttp.BaseRequest

	// 环境id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 命名空间描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网网络
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *ModifyNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "NamespaceName")
	delete(f, "Description")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return errors.New("ModifyNamespaceRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功时为命名空间ID，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NamespacePage struct {

	// 分页内容
	Records []*TemNamespaceInfo `json:"Records,omitempty" name:"Records" list`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 条目数
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 页数
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`
}

type TemNamespaceInfo struct {

	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 区域名称
	Region *string `json:"Region,omitempty" name:"Region"`

	// 命名空间描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 状态,1:已销毁;0:正常
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// vpc网络
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 创建时间
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 修改时间
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// 修改人
	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`

	// 创建人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 服务数
	ServiceNum *int64 `json:"ServiceNum,omitempty" name:"ServiceNum"`

	// 运行实例数
	RunInstancesNum *int64 `json:"RunInstancesNum,omitempty" name:"RunInstancesNum"`

	// 子网络
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// tcb环境状态
	TcbEnvStatus *string `json:"TcbEnvStatus,omitempty" name:"TcbEnvStatus"`

	// eks cluster status
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
}
