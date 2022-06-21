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

package v20201207

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApolloEnvParam struct {
	// 环境名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 环境内引擎的节点规格 ID
	EngineResourceSpec *string `json:"EngineResourceSpec,omitempty" name:"EngineResourceSpec"`

	// 环境内引擎的节点数量
	EngineNodeNum *int64 `json:"EngineNodeNum,omitempty" name:"EngineNodeNum"`

	// 配置存储空间大小，以GB为单位
	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`

	// VPC ID。在 VPC 的子网内分配一个 IP 作为 ConfigServer 的访问地址
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID。在 VPC 的子网内分配一个 IP 作为 ConfigServer 的访问地址
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type BoundK8SInfo struct {
	// 绑定的kubernetes集群ID
	BoundClusterId *string `json:"BoundClusterId,omitempty" name:"BoundClusterId"`

	// 绑定的kubernetes的集群类型，分tke和eks两种
	// 注意：此字段可能返回 null，表示取不到有效值。
	BoundClusterType *string `json:"BoundClusterType,omitempty" name:"BoundClusterType"`

	// 服务同步模式，all为全量同步，demand为按需同步
	// 注意：此字段可能返回 null，表示取不到有效值。
	SyncMode *string `json:"SyncMode,omitempty" name:"SyncMode"`
}

// Predefined struct for user
type CreateEngineRequestParams struct {
	// 引擎类型。参考值：
	// - zookeeper
	// - nacos
	// - consul
	// - apollo
	// - eureka
	// - polaris
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 引擎的开源版本。每种引擎支持的开源版本不同，请参考产品文档或者控制台购买页
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 引擎的产品版本。参考值：
	// - STANDARD： 标准版
	// 
	// 引擎各版本及可选择的规格、节点数说明：
	// apollo - STANDARD版本
	// 规格列表：spec-qcr53kf1t（1C2G）,spec-qdr53kf2w（2C4G）
	// 节点数：1，2，3，4，5
	// 
	// eureka - STANDARD版本
	// 规格列表：spec-qvj6k7t4q（1C2G）,spec-qcr53kfjt（2C4G）,spec-qvj6k7t4m（4G8G）,spec-qcr54kfjt（8C16G）,spec-qcr55kfjt（16C32G）
	// 节点数：3，4，5
	EngineProductVersion *string `json:"EngineProductVersion,omitempty" name:"EngineProductVersion"`

	// 引擎所在地域。参考值说明：
	// 中国区 参考值：
	// - ap-guangzhou：广州
	// - ap-beijing：北京
	// - ap-chengdu：成都
	// - ap-chongqing：重庆
	// - ap-nanjing：南京
	// - ap-shanghai：上海
	// - ap-hongkong：香港
	// - ap-taipei：台北
	// 亚太区 参考值：
	// - ap-jakarta：雅加达
	// - ap-singapore：新加坡
	// 北美区 参考值
	// - na-toronto：多伦多
	// 金融专区 参考值
	// - ap-beijing-fsi：北京金融
	// - ap-shanghai-fsi：上海金融
	// - ap-shenzhen-fsi：深圳金融
	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`

	// 引擎名称。参考值：
	// - eurek-test
	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`

	// 付费类型。参考值：
	// - 0：后付费
	// - 1：预付费
	TradeType *int64 `json:"TradeType,omitempty" name:"TradeType"`

	// 引擎的节点规格 ID。参见EngineProductVersion字段说明
	EngineResourceSpec *string `json:"EngineResourceSpec,omitempty" name:"EngineResourceSpec"`

	// 引擎的节点数量。参见EngineProductVersion字段说明
	EngineNodeNum *int64 `json:"EngineNodeNum,omitempty" name:"EngineNodeNum"`

	// VPC ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - vpc-conz6aix
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - subnet-ahde9me9
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Apollo 环境配置参数列表。参数说明：
	// 如果创建Apollo类型，此参数为必填的环境信息列表，最多可选4个环境。环境信息参数说明：
	// - Name：环境名。参考值：prod, dev, fat, uat
	// - EngineResourceSpec：环境内引擎的节点规格ID。参见EngineProductVersion参数说明
	// - EngineNodeNum：环境内引擎的节点数量。参见EngineProductVersion参数说明，其中prod环境支持的节点数为2，3，4，5
	// - StorageCapacity：配置存储空间大小，以GB为单位，步长为5.参考值：35
	// - VpcId：VPC ID。参考值：vpc-conz6aix
	// - SubnetId：子网 ID。参考值：subnet-ahde9me9
	ApolloEnvParams []*ApolloEnvParam `json:"ApolloEnvParams,omitempty" name:"ApolloEnvParams"`

	// 引擎的标签列表。用户自定义的key/value形式，无参考值
	EngineTags []*InstanceTagInfo `json:"EngineTags,omitempty" name:"EngineTags"`

	// 引擎的初始帐号信息。可设置参数：
	// - Name：控制台初始用户名
	// - Password：控制台初始密码
	// - Token：引擎接口的管理员 Token
	EngineAdmin *EngineAdmin `json:"EngineAdmin,omitempty" name:"EngineAdmin"`

	// 预付费时长，以月为单位
	PrepaidPeriod *int64 `json:"PrepaidPeriod,omitempty" name:"PrepaidPeriod"`

	// 自动续费标记，仅预付费使用。参考值：
	// - 0：不自动续费
	// - 1：自动续费
	PrepaidRenewFlag *int64 `json:"PrepaidRenewFlag,omitempty" name:"PrepaidRenewFlag"`
}

type CreateEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎类型。参考值：
	// - zookeeper
	// - nacos
	// - consul
	// - apollo
	// - eureka
	// - polaris
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 引擎的开源版本。每种引擎支持的开源版本不同，请参考产品文档或者控制台购买页
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 引擎的产品版本。参考值：
	// - STANDARD： 标准版
	// 
	// 引擎各版本及可选择的规格、节点数说明：
	// apollo - STANDARD版本
	// 规格列表：spec-qcr53kf1t（1C2G）,spec-qdr53kf2w（2C4G）
	// 节点数：1，2，3，4，5
	// 
	// eureka - STANDARD版本
	// 规格列表：spec-qvj6k7t4q（1C2G）,spec-qcr53kfjt（2C4G）,spec-qvj6k7t4m（4G8G）,spec-qcr54kfjt（8C16G）,spec-qcr55kfjt（16C32G）
	// 节点数：3，4，5
	EngineProductVersion *string `json:"EngineProductVersion,omitempty" name:"EngineProductVersion"`

	// 引擎所在地域。参考值说明：
	// 中国区 参考值：
	// - ap-guangzhou：广州
	// - ap-beijing：北京
	// - ap-chengdu：成都
	// - ap-chongqing：重庆
	// - ap-nanjing：南京
	// - ap-shanghai：上海
	// - ap-hongkong：香港
	// - ap-taipei：台北
	// 亚太区 参考值：
	// - ap-jakarta：雅加达
	// - ap-singapore：新加坡
	// 北美区 参考值
	// - na-toronto：多伦多
	// 金融专区 参考值
	// - ap-beijing-fsi：北京金融
	// - ap-shanghai-fsi：上海金融
	// - ap-shenzhen-fsi：深圳金融
	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`

	// 引擎名称。参考值：
	// - eurek-test
	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`

	// 付费类型。参考值：
	// - 0：后付费
	// - 1：预付费
	TradeType *int64 `json:"TradeType,omitempty" name:"TradeType"`

	// 引擎的节点规格 ID。参见EngineProductVersion字段说明
	EngineResourceSpec *string `json:"EngineResourceSpec,omitempty" name:"EngineResourceSpec"`

	// 引擎的节点数量。参见EngineProductVersion字段说明
	EngineNodeNum *int64 `json:"EngineNodeNum,omitempty" name:"EngineNodeNum"`

	// VPC ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - vpc-conz6aix
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - subnet-ahde9me9
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// Apollo 环境配置参数列表。参数说明：
	// 如果创建Apollo类型，此参数为必填的环境信息列表，最多可选4个环境。环境信息参数说明：
	// - Name：环境名。参考值：prod, dev, fat, uat
	// - EngineResourceSpec：环境内引擎的节点规格ID。参见EngineProductVersion参数说明
	// - EngineNodeNum：环境内引擎的节点数量。参见EngineProductVersion参数说明，其中prod环境支持的节点数为2，3，4，5
	// - StorageCapacity：配置存储空间大小，以GB为单位，步长为5.参考值：35
	// - VpcId：VPC ID。参考值：vpc-conz6aix
	// - SubnetId：子网 ID。参考值：subnet-ahde9me9
	ApolloEnvParams []*ApolloEnvParam `json:"ApolloEnvParams,omitempty" name:"ApolloEnvParams"`

	// 引擎的标签列表。用户自定义的key/value形式，无参考值
	EngineTags []*InstanceTagInfo `json:"EngineTags,omitempty" name:"EngineTags"`

	// 引擎的初始帐号信息。可设置参数：
	// - Name：控制台初始用户名
	// - Password：控制台初始密码
	// - Token：引擎接口的管理员 Token
	EngineAdmin *EngineAdmin `json:"EngineAdmin,omitempty" name:"EngineAdmin"`

	// 预付费时长，以月为单位
	PrepaidPeriod *int64 `json:"PrepaidPeriod,omitempty" name:"PrepaidPeriod"`

	// 自动续费标记，仅预付费使用。参考值：
	// - 0：不自动续费
	// - 1：自动续费
	PrepaidRenewFlag *int64 `json:"PrepaidRenewFlag,omitempty" name:"PrepaidRenewFlag"`
}

func (r *CreateEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineType")
	delete(f, "EngineVersion")
	delete(f, "EngineProductVersion")
	delete(f, "EngineRegion")
	delete(f, "EngineName")
	delete(f, "TradeType")
	delete(f, "EngineResourceSpec")
	delete(f, "EngineNodeNum")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ApolloEnvParams")
	delete(f, "EngineTags")
	delete(f, "EngineAdmin")
	delete(f, "PrepaidPeriod")
	delete(f, "PrepaidRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEngineResponseParams struct {
	// 引擎实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEngineResponse struct {
	*tchttp.BaseResponse
	Response *CreateEngineResponseParams `json:"Response"`
}

func (r *CreateEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEngineRequestParams struct {
	// 引擎实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DeleteEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEngineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEngineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEngineResponseParams `json:"Response"`
}

func (r *DeleteEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSREInstanceAccessAddressRequestParams struct {
	// 注册引擎实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DescribeSREInstanceAccessAddressRequest struct {
	*tchttp.BaseRequest
	
	// 注册引擎实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DescribeSREInstanceAccessAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstanceAccessAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSREInstanceAccessAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSREInstanceAccessAddressResponseParams struct {
	// 内网访问地址
	IntranetAddress *string `json:"IntranetAddress,omitempty" name:"IntranetAddress"`

	// 公网访问地址
	InternetAddress *string `json:"InternetAddress,omitempty" name:"InternetAddress"`

	// apollo多环境公网ip
	EnvAddressInfos []*EnvAddressInfo `json:"EnvAddressInfos,omitempty" name:"EnvAddressInfos"`

	// 控制台公网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsoleInternetAddress *string `json:"ConsoleInternetAddress,omitempty" name:"ConsoleInternetAddress"`

	// 控制台内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsoleIntranetAddress *string `json:"ConsoleIntranetAddress,omitempty" name:"ConsoleIntranetAddress"`

	// 客户端公网带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetBandWidth *int64 `json:"InternetBandWidth,omitempty" name:"InternetBandWidth"`

	// 控制台公网带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsoleInternetBandWidth *int64 `json:"ConsoleInternetBandWidth,omitempty" name:"ConsoleInternetBandWidth"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSREInstanceAccessAddressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSREInstanceAccessAddressResponseParams `json:"Response"`
}

func (r *DescribeSREInstanceAccessAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstanceAccessAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSREInstancesRequestParams struct {
	// 请求过滤参数
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询类型
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 调用方来源
	QuerySource *string `json:"QuerySource,omitempty" name:"QuerySource"`
}

type DescribeSREInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 请求过滤参数
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询类型
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 调用方来源
	QuerySource *string `json:"QuerySource,omitempty" name:"QuerySource"`
}

func (r *DescribeSREInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "QueryType")
	delete(f, "QuerySource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSREInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSREInstancesResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例记录
	Content []*SREInstance `json:"Content,omitempty" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSREInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSREInstancesResponseParams `json:"Response"`
}

func (r *DescribeSREInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineAdmin struct {
	// 控制台初始用户名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 控制台初始密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 引擎接口的管理员 Token
	Token *string `json:"Token,omitempty" name:"Token"`
}

type EnvAddressInfo struct {
	// 环境名
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// 是否开启config公网
	EnableConfigInternet *bool `json:"EnableConfigInternet,omitempty" name:"EnableConfigInternet"`

	// config公网ip
	ConfigInternetServiceIp *string `json:"ConfigInternetServiceIp,omitempty" name:"ConfigInternetServiceIp"`
}

type EnvInfo struct {
	// 环境名称
	EnvName *string `json:"EnvName,omitempty" name:"EnvName"`

	// 环境对应的网络信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitempty" name:"VpcInfos"`

	// 云硬盘容量
	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`

	// 运行状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// Admin service 访问地址
	AdminServiceIp *string `json:"AdminServiceIp,omitempty" name:"AdminServiceIp"`

	// Config service访问地址
	ConfigServiceIp *string `json:"ConfigServiceIp,omitempty" name:"ConfigServiceIp"`

	// 是否开启config-server公网
	EnableConfigInternet *bool `json:"EnableConfigInternet,omitempty" name:"EnableConfigInternet"`

	// config-server公网访问地址
	ConfigInternetServiceIp *string `json:"ConfigInternetServiceIp,omitempty" name:"ConfigInternetServiceIp"`

	// 规格ID
	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`

	// 环境的节点数
	EnvReplica *int64 `json:"EnvReplica,omitempty" name:"EnvReplica"`

	// 环境运行的节点数
	RunningCount *int64 `json:"RunningCount,omitempty" name:"RunningCount"`

	// 环境别名
	AliasEnvName *string `json:"AliasEnvName,omitempty" name:"AliasEnvName"`
}

type Filter struct {
	// 过滤参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤参数值
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InstanceTagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type KVPair struct {
	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SREInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 版本号
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 状态, 枚举值:creating/create_fail/running/updating/update_fail/restarting/restart_fail/destroying/destroy_fail
	Status *string `json:"Status,omitempty" name:"Status"`

	// 规格ID
	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`

	// 副本数
	Replica *int64 `json:"Replica,omitempty" name:"Replica"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// Vpc iD
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 是否开启持久化存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableStorage *bool `json:"EnableStorage,omitempty" name:"EnableStorage"`

	// 数据存储方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 云硬盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageCapacity *int64 `json:"StorageCapacity,omitempty" name:"StorageCapacity"`

	// 计费方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paymode *string `json:"Paymode,omitempty" name:"Paymode"`

	// EKS集群的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EKSClusterID *string `json:"EKSClusterID,omitempty" name:"EKSClusterID"`

	// 集群创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 环境配置信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvInfos []*EnvInfo `json:"EnvInfos,omitempty" name:"EnvInfos"`

	// 引擎所在的区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`

	// 注册引擎是否开启公网
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`

	// 私有网络列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcInfos []*VpcInfo `json:"VpcInfos,omitempty" name:"VpcInfos"`

	// 服务治理相关信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGovernanceInfos []*ServiceGovernanceInfo `json:"ServiceGovernanceInfos,omitempty" name:"ServiceGovernanceInfos"`

	// 实例的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*KVPair `json:"Tags,omitempty" name:"Tags"`

	// 引擎实例是否开启控制台公网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableConsoleInternet *bool `json:"EnableConsoleInternet,omitempty" name:"EnableConsoleInternet"`

	// 引擎实例是否开启控制台内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableConsoleIntranet *bool `json:"EnableConsoleIntranet,omitempty" name:"EnableConsoleIntranet"`

	// 引擎实例是否展示参数配置页面
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigInfoVisible *bool `json:"ConfigInfoVisible,omitempty" name:"ConfigInfoVisible"`

	// 引擎实例控制台默认密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsoleDefaultPwd *string `json:"ConsoleDefaultPwd,omitempty" name:"ConsoleDefaultPwd"`

	// 交易付费类型，0后付费/1预付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeType *int64 `json:"TradeType,omitempty" name:"TradeType"`

	// 自动续费标记：0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 预付费到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurDeadline *string `json:"CurDeadline,omitempty" name:"CurDeadline"`

	// 隔离开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`
}

type ServiceGovernanceInfo struct {
	// 引擎所在的地域
	EngineRegion *string `json:"EngineRegion,omitempty" name:"EngineRegion"`

	// 服务治理引擎绑定的kubernetes集群信息
	BoundK8SInfos []*BoundK8SInfo `json:"BoundK8SInfos,omitempty" name:"BoundK8SInfos"`

	// 服务治理引擎绑定的网络信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitempty" name:"VpcInfos"`

	// 当前实例鉴权是否开启
	AuthOpen *bool `json:"AuthOpen,omitempty" name:"AuthOpen"`

	// 该实例支持的功能，鉴权就是 Auth
	Features []*string `json:"Features,omitempty" name:"Features"`

	// 主账户名默认为 polaris，该值为主账户的默认密码
	MainPassword *string `json:"MainPassword,omitempty" name:"MainPassword"`
}

type VpcInfo struct {
	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetAddress *string `json:"IntranetAddress,omitempty" name:"IntranetAddress"`
}