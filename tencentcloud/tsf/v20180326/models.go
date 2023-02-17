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

package v20180326

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddClusterInstancesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 重装系统密码设置
	Password *string `json:"Password,omitempty" name:"Password"`

	// 重装系统，关联密钥设置
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组设置
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 云主机导入方式，虚拟机集群必填，容器集群不填写此字段，R：重装TSF系统镜像，M：手动安装agent
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`

	// 镜像定制类型
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// 镜像特征ID列表
	FeatureIdList []*string `json:"FeatureIdList,omitempty" name:"FeatureIdList"`

	// 实例额外需要设置参数信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// 安全组 ID 列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type AddClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 重装系统密码设置
	Password *string `json:"Password,omitempty" name:"Password"`

	// 重装系统，关联密钥设置
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组设置
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 云主机导入方式，虚拟机集群必填，容器集群不填写此字段，R：重装TSF系统镜像，M：手动安装agent
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`

	// 镜像定制类型
	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`

	// 镜像特征ID列表
	FeatureIdList []*string `json:"FeatureIdList,omitempty" name:"FeatureIdList"`

	// 实例额外需要设置参数信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`

	// 安全组 ID 列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *AddClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	delete(f, "OsName")
	delete(f, "ImageId")
	delete(f, "Password")
	delete(f, "KeyId")
	delete(f, "SgId")
	delete(f, "InstanceImportMode")
	delete(f, "OsCustomizeType")
	delete(f, "FeatureIdList")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterInstancesResponseParams struct {
	// 添加云主机的返回列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AddInstanceResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AddClusterInstancesResponseParams `json:"Response"`
}

func (r *AddClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddInstanceResult struct {
	// 添加集群失败的节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`

	// 添加集群成功的节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds"`

	// 添加集群超时的节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutInstanceIds []*string `json:"TimeoutInstanceIds,omitempty" name:"TimeoutInstanceIds"`

	// 失败的节点的失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReasons []*string `json:"FailedReasons,omitempty" name:"FailedReasons"`
}

// Predefined struct for user
type AddInstancesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 重装系统密码设置
	Password *string `json:"Password,omitempty" name:"Password"`

	// 重装系统，关联密钥设置
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组设置
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 云主机导入方式，虚拟机集群必填，容器集群不填写此字段，R：重装TSF系统镜像，M：手动安装agent
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`

	// 安全组id
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 重装系统密码设置
	Password *string `json:"Password,omitempty" name:"Password"`

	// 重装系统，关联密钥设置
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组设置
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 云主机导入方式，虚拟机集群必填，容器集群不填写此字段，R：重装TSF系统镜像，M：手动安装agent
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`

	// 安全组id
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

func (r *AddInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	delete(f, "OsName")
	delete(f, "ImageId")
	delete(f, "Password")
	delete(f, "KeyId")
	delete(f, "SgId")
	delete(f, "InstanceImportMode")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddInstancesResponseParams struct {
	// 添加云主机是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AddInstancesResponseParams `json:"Response"`
}

func (r *AddInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvanceSettings struct {
	// 子任务单机并发数限制，默认值为2
	SubTaskConcurrency *int64 `json:"SubTaskConcurrency,omitempty" name:"SubTaskConcurrency"`
}

type AgentProfile struct {
	// Agent类型
	AgentType *string `json:"AgentType,omitempty" name:"AgentType"`

	// Agent版本号
	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
}

type ApiDefinitionDescr struct {
	// 对象名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 对象属性列表
	Properties []*PropertyField `json:"Properties,omitempty" name:"Properties"`
}

type ApiDetailInfo struct {
	// API ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// API 请求路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// Api 映射路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathMapping *string `json:"PathMapping,omitempty" name:"PathMapping"`

	// 请求方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 所属分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 是否禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// 发布状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseStatus *string `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`

	// 开启限流
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitStatus *string `json:"RateLimitStatus,omitempty" name:"RateLimitStatus"`

	// 是否开启mock
	// 注意：此字段可能返回 null，表示取不到有效值。
	MockStatus *string `json:"MockStatus,omitempty" name:"MockStatus"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleasedTime *string `json:"ReleasedTime,omitempty" name:"ReleasedTime"`

	// 所属分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// API 超时，单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// Api所在服务host
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`

	// API类型。 ms ： 微服务API； external :外部服务Api
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// Api描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// API路径匹配类型。normal：普通API；wildcard：通配API。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiMatchType *string `json:"ApiMatchType,omitempty" name:"ApiMatchType"`

	// RPC 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RpcExt *string `json:"RpcExt,omitempty" name:"RpcExt"`

	// 部署组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// md5
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// RPC 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RpcType *string `json:"RpcType,omitempty" name:"RpcType"`
}

type ApiDetailResponse struct {
	// API 请求参数
	Request []*ApiRequestDescr `json:"Request,omitempty" name:"Request"`

	// API 响应参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Response []*ApiResponseDescr `json:"Response,omitempty" name:"Response"`

	// API 复杂结构定义
	Definitions []*ApiDefinitionDescr `json:"Definitions,omitempty" name:"Definitions"`

	// API 的 content type
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestContentType *string `json:"RequestContentType,omitempty" name:"RequestContentType"`

	// API  能否调试
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanRun *bool `json:"CanRun,omitempty" name:"CanRun"`

	// API 状态 0:离线 1:在线，默认0
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// API 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ApiGroupInfo struct {
	// Api Group Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Api Group 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组上下文
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`

	// 鉴权类型。 secret： 密钥鉴权； none:无鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 发布状态, drafted: 未发布。 released: 发布
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分组创建时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 分组更新时间 如:2019-06-20 15:51:28
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// api分组已绑定的网关部署组
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindedGatewayDeployGroups []*GatewayDeployGroup `json:"BindedGatewayDeployGroups,omitempty" name:"BindedGatewayDeployGroups"`

	// api 个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiCount *int64 `json:"ApiCount,omitempty" name:"ApiCount"`

	// 访问group的ACL类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclMode *string `json:"AclMode,omitempty" name:"AclMode"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 分组类型。 ms： 微服务分组； external:外部Api分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// 网关实例的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayInstanceType *string `json:"GatewayInstanceType,omitempty" name:"GatewayInstanceType"`

	// 网关实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 命名空间参数key值
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceNameKey *string `json:"NamespaceNameKey,omitempty" name:"NamespaceNameKey"`

	// 微服务名参数key值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceNameKey *string `json:"ServiceNameKey,omitempty" name:"ServiceNameKey"`

	// 命名空间参数位置，path，header或query，默认是path
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceNameKeyPosition *string `json:"NamespaceNameKeyPosition,omitempty" name:"NamespaceNameKeyPosition"`

	// 微服务名参数位置，path，header或query，默认是path
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceNameKeyPosition *string `json:"ServiceNameKeyPosition,omitempty" name:"ServiceNameKeyPosition"`
}

type ApiInfo struct {
	// 命名空间Id，若为外部API,为固定值："namespace-external"
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 服务Id，若为外部API,为固定值："ms-external"
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// API path
	Path *string `json:"Path,omitempty" name:"Path"`

	// Api 请求
	Method *string `json:"Method,omitempty" name:"Method"`

	// 请求映射
	PathMapping *string `json:"PathMapping,omitempty" name:"PathMapping"`

	// api所在服务host,限定外部Api填写。格式: `http://127.0.0.1:8080`
	Host *string `json:"Host,omitempty" name:"Host"`

	// api描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ApiRateLimitRule struct {
	// rule Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// API ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 限流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 最大限流qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// 生效/禁用, enabled/disabled
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// 规则内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleContent *string `json:"RuleContent,omitempty" name:"RuleContent"`

	// Tsf Rule ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRuleId *string `json:"TsfRuleId,omitempty" name:"TsfRuleId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
}

type ApiRequestDescr struct {
	// 参数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数位置
	In *string `json:"In,omitempty" name:"In"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数是否必须
	Required *bool `json:"Required,omitempty" name:"Required"`

	// 参数的默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
}

type ApiResponseDescr struct {
	// 参数描述
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ApiUseStatisticsEntity struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 次数
	Count *string `json:"Count,omitempty" name:"Count"`

	// 比率
	Ratio *string `json:"Ratio,omitempty" name:"Ratio"`
}

type ApiVersionArray struct {
	// App ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// App 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// App 包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`
}

type ApplicationAttribute struct {
	// 总实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 运行实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 应用下部署组个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupCount *int64 `json:"GroupCount,omitempty" name:"GroupCount"`
}

type ApplicationForPage struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 编程语言
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgLang *string `json:"ProgLang,omitempty" name:"ProgLang"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 应用资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// 应用runtime类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationRuntimeType *string `json:"ApplicationRuntimeType,omitempty" name:"ApplicationRuntimeType"`

	// Apigateway的serviceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApigatewayServiceId *string `json:"ApigatewayServiceId,omitempty" name:"ApigatewayServiceId"`

	// 应用备注名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationRemarkName *string `json:"ApplicationRemarkName,omitempty" name:"ApplicationRemarkName"`

	// 服务配置信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceConfigList []*ServiceConfig `json:"ServiceConfigList,omitempty" name:"ServiceConfigList"`

	// IgnoreCreateImageRepository
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCreateImageRepository *bool `json:"IgnoreCreateImageRepository,omitempty" name:"IgnoreCreateImageRepository"`
}

// Predefined struct for user
type AssociateBusinessLogConfigRequestParams struct {
	// TSF分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 日志配置项ID列表
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`
}

type AssociateBusinessLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// TSF分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 日志配置项ID列表
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`
}

func (r *AssociateBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateBusinessLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ConfigIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateBusinessLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateBusinessLogConfigResponseParams struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *AssociateBusinessLogConfigResponseParams `json:"Response"`
}

func (r *AssociateBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateBusinessLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateConfigWithGroupRequestParams struct {
	// 配置项id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组信息
	Groups []*GroupInfo `json:"Groups,omitempty" name:"Groups"`

	// 是否选择全部投递，1 表示全部，0或不填表示非全部
	SelectAll *int64 `json:"SelectAll,omitempty" name:"SelectAll"`

	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 模糊搜索关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

type AssociateConfigWithGroupRequest struct {
	*tchttp.BaseRequest
	
	// 配置项id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组信息
	Groups []*GroupInfo `json:"Groups,omitempty" name:"Groups"`

	// 是否选择全部投递，1 表示全部，0或不填表示非全部
	SelectAll *int64 `json:"SelectAll,omitempty" name:"SelectAll"`

	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 模糊搜索关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *AssociateConfigWithGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateConfigWithGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "Groups")
	delete(f, "SelectAll")
	delete(f, "NamespaceId")
	delete(f, "ClusterId")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateConfigWithGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateConfigWithGroupResponseParams struct {
	// 绑定是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateConfigWithGroupResponse struct {
	*tchttp.BaseResponse
	Response *AssociateConfigWithGroupResponseParams `json:"Response"`
}

func (r *AssociateConfigWithGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateConfigWithGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindApiGroupRequestParams struct {
	// 分组绑定网关列表
	GroupGatewayList []*GatewayGroupIds `json:"GroupGatewayList,omitempty" name:"GroupGatewayList"`
}

type BindApiGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分组绑定网关列表
	GroupGatewayList []*GatewayGroupIds `json:"GroupGatewayList,omitempty" name:"GroupGatewayList"`
}

func (r *BindApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindApiGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupGatewayList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindApiGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindApiGroupResponseParams struct {
	// 返回结果，成功失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *BindApiGroupResponseParams `json:"Response"`
}

func (r *BindApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindApiGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindPluginRequestParams struct {
	// 分组/API绑定插件列表
	PluginInstanceList []*GatewayPluginBoundParam `json:"PluginInstanceList,omitempty" name:"PluginInstanceList"`
}

type BindPluginRequest struct {
	*tchttp.BaseRequest
	
	// 分组/API绑定插件列表
	PluginInstanceList []*GatewayPluginBoundParam `json:"PluginInstanceList,omitempty" name:"PluginInstanceList"`
}

func (r *BindPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginInstanceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindPluginResponseParams struct {
	// 返回结果，成功失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindPluginResponse struct {
	*tchttp.BaseResponse
	Response *BindPluginResponseParams `json:"Response"`
}

func (r *BindPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BusinesLogConfigAssociatedGroup struct {
	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署组所属应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组所属应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 部署组所属应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 部署组所属命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 部署组所属命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 部署组所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组所属集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 部署组所属集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 部署组关联日志配置时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedTime *string `json:"AssociatedTime,omitempty" name:"AssociatedTime"`
}

type BusinessLogConfig struct {
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项日志路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`

	// 配置项描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigDesc *string `json:"ConfigDesc,omitempty" name:"ConfigDesc"`

	// 配置项标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTags *string `json:"ConfigTags,omitempty" name:"ConfigTags"`

	// 配置项对应的ES管道
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigPipeline *string `json:"ConfigPipeline,omitempty" name:"ConfigPipeline"`

	// 配置项创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigCreateTime *string `json:"ConfigCreateTime,omitempty" name:"ConfigCreateTime"`

	// 配置项更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigUpdateTime *string `json:"ConfigUpdateTime,omitempty" name:"ConfigUpdateTime"`

	// 配置项解析规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigSchema *BusinessLogConfigSchema `json:"ConfigSchema,omitempty" name:"ConfigSchema"`

	// 配置项关联部署组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigAssociatedGroups []*BusinesLogConfigAssociatedGroup `json:"ConfigAssociatedGroups,omitempty" name:"ConfigAssociatedGroups"`
}

type BusinessLogConfigSchema struct {
	// 解析规则类型
	SchemaType *int64 `json:"SchemaType,omitempty" name:"SchemaType"`

	// 解析规则内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaContent *string `json:"SchemaContent,omitempty" name:"SchemaContent"`

	// 解析规则时间格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaDateFormat *string `json:"SchemaDateFormat,omitempty" name:"SchemaDateFormat"`

	// 解析规则对应的多行匹配规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaMultilinePattern *string `json:"SchemaMultilinePattern,omitempty" name:"SchemaMultilinePattern"`

	// 解析规则创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaCreateTime *string `json:"SchemaCreateTime,omitempty" name:"SchemaCreateTime"`

	// 用户填写的解析规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaPatternLayout *string `json:"SchemaPatternLayout,omitempty" name:"SchemaPatternLayout"`
}

type BusinessLogV2 struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 日志时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIp *string `json:"InstanceIp,omitempty" name:"InstanceIp"`

	// 日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

// Predefined struct for user
type ChangeApiUsableStatusRequestParams struct {
	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 切换状态，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`
}

type ChangeApiUsableStatusRequest struct {
	*tchttp.BaseRequest
	
	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 切换状态，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`
}

func (r *ChangeApiUsableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeApiUsableStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiId")
	delete(f, "UsableStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeApiUsableStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeApiUsableStatusResponseParams struct {
	// API 信息
	Result *ApiDetailInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChangeApiUsableStatusResponse struct {
	*tchttp.BaseResponse
	Response *ChangeApiUsableStatusResponseParams `json:"Response"`
}

func (r *ChangeApiUsableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeApiUsableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {
	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群所属私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 集群CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 集群总CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalCpu *float64 `json:"ClusterTotalCpu,omitempty" name:"ClusterTotalCpu"`

	// 集群总内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalMem *float64 `json:"ClusterTotalMem,omitempty" name:"ClusterTotalMem"`

	// 集群已使用CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedCpu *float64 `json:"ClusterUsedCpu,omitempty" name:"ClusterUsedCpu"`

	// 集群已使用内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedMem *float64 `json:"ClusterUsedMem,omitempty" name:"ClusterUsedMem"`

	// 集群机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 集群可用的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 集群正常状态的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalInstanceCount *int64 `json:"NormalInstanceCount,omitempty" name:"NormalInstanceCount"`

	// 删除标记：true：可以删除；false：不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群所属TSF地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionName *string `json:"TsfRegionName,omitempty" name:"TsfRegionName"`

	// 集群所属TSF可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`

	// 集群所属TSF可用区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneName *string `json:"TsfZoneName,omitempty" name:"TsfZoneName"`

	// 集群不可删除的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlagReason *string `json:"DeleteFlagReason,omitempty" name:"DeleteFlagReason"`

	// 集群最大CPU限制，单位：核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterLimitCpu *float64 `json:"ClusterLimitCpu,omitempty" name:"ClusterLimitCpu"`

	// 集群最大内存限制，单位：G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterLimitMem *float64 `json:"ClusterLimitMem,omitempty" name:"ClusterLimitMem"`

	// 集群可用的服务实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunServiceInstanceCount *int64 `json:"RunServiceInstanceCount,omitempty" name:"RunServiceInstanceCount"`

	// 集群所属子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 返回给前端的控制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationInfo *OperationInfo `json:"OperationInfo,omitempty" name:"OperationInfo"`

	// 集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
}

type ClusterV2 struct {
	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群所属私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 集群CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 集群总CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalCpu *float64 `json:"ClusterTotalCpu,omitempty" name:"ClusterTotalCpu"`

	// 集群总内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalMem *float64 `json:"ClusterTotalMem,omitempty" name:"ClusterTotalMem"`

	// 集群已使用CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedCpu *float64 `json:"ClusterUsedCpu,omitempty" name:"ClusterUsedCpu"`

	// 集群已使用内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedMem *float64 `json:"ClusterUsedMem,omitempty" name:"ClusterUsedMem"`

	// 集群机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 集群运行中的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 集群正常状态的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalInstanceCount *int64 `json:"NormalInstanceCount,omitempty" name:"NormalInstanceCount"`

	// 删除标记：true：可以删除；false：不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群所属TSF地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionName *string `json:"TsfRegionName,omitempty" name:"TsfRegionName"`

	// 集群所属TSF可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`

	// 集群所属TSF可用区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneName *string `json:"TsfZoneName,omitempty" name:"TsfZoneName"`

	// 集群不可删除的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlagReason *string `json:"DeleteFlagReason,omitempty" name:"DeleteFlagReason"`

	// 集群所属私有网络子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 集群剩余 cpu limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterLimitCpu *string `json:"ClusterLimitCpu,omitempty" name:"ClusterLimitCpu"`

	// 集群剩余 memory limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterLimitMem *string `json:"ClusterLimitMem,omitempty" name:"ClusterLimitMem"`

	// 运行服务实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunServiceInstanceCount *int64 `json:"RunServiceInstanceCount,omitempty" name:"RunServiceInstanceCount"`

	// 给前端的按钮控制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationInfo *OperationInfo `json:"OperationInfo,omitempty" name:"OperationInfo"`

	// 容器集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 部署组总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupCount *uint64 `json:"GroupCount,omitempty" name:"GroupCount"`

	// 运行中部署组数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunGroupCount *uint64 `json:"RunGroupCount,omitempty" name:"RunGroupCount"`

	// 停止中部署组数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopGroupCount *uint64 `json:"StopGroupCount,omitempty" name:"StopGroupCount"`

	// 异常部署组数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AbnormalGroupCount *uint64 `json:"AbnormalGroupCount,omitempty" name:"AbnormalGroupCount"`

	// 集群备注名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterRemarkName *string `json:"ClusterRemarkName,omitempty" name:"ClusterRemarkName"`

	// api地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	KuberneteApiServer *string `json:"KuberneteApiServer,omitempty" name:"KuberneteApiServer"`

	// K : kubeconfig, S : service account
	// 注意：此字段可能返回 null，表示取不到有效值。
	KuberneteNativeType *string `json:"KuberneteNativeType,omitempty" name:"KuberneteNativeType"`

	// native secret
	// 注意：此字段可能返回 null，表示取不到有效值。
	KuberneteNativeSecret *string `json:"KuberneteNativeSecret,omitempty" name:"KuberneteNativeSecret"`
}

type Config struct {
	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项版本描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 配置项类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 删除标识，true：可以删除；false：不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 配置项版本数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionCount *int64 `json:"ConfigVersionCount,omitempty" name:"ConfigVersionCount"`
}

type ConfigRelease struct {
	// 配置项发布ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type ConfigReleaseLog struct {
	// 配置项发布日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 发布状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseStatus *string `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`

	// 上次发布的配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigId *string `json:"LastConfigId,omitempty" name:"LastConfigId"`

	// 上次发布的配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigName *string `json:"LastConfigName,omitempty" name:"LastConfigName"`

	// 上次发布的配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigVersion *string `json:"LastConfigVersion,omitempty" name:"LastConfigVersion"`

	// 回滚标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollbackFlag *bool `json:"RollbackFlag,omitempty" name:"RollbackFlag"`
}

type ConfigTemplate struct {
	// 配置模板Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`

	// 配置模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateName *string `json:"ConfigTemplateName,omitempty" name:"ConfigTemplateName"`

	// 配置模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateDesc *string `json:"ConfigTemplateDesc,omitempty" name:"ConfigTemplateDesc"`

	// 配置模板对应的微服务框架
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateType *string `json:"ConfigTemplateType,omitempty" name:"ConfigTemplateType"`

	// 配置模板数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigTemplateValue *string `json:"ConfigTemplateValue,omitempty" name:"ConfigTemplateValue"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ContainGroup struct {
	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 部署组备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// KubeInjectEnable值
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeInjectEnable *bool `json:"KubeInjectEnable,omitempty" name:"KubeInjectEnable"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
}

type ContainGroupResult struct {
	// 部署组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ContainGroup `json:"Content,omitempty" name:"Content"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ContainerEvent struct {
	// 第一次出现的时间，以 ms 为单位的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTimestamp *int64 `json:"FirstTimestamp,omitempty" name:"FirstTimestamp"`

	// 最后一次出现的时间，以 ms 为单位的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" name:"LastTimestamp"`

	// 级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 详细描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 出现次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type ContainerGroupDeploy struct {
	// 部署组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 已启动实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 业务容器初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 业务容器最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 业务容器初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 业务容器最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 0:公网 1:集群内访问 2：NodePort
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 端口映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts"`

	// 更新方式：0:快速更新 1:滚动更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// jvm参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// agent容器初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentCpuRequest *string `json:"AgentCpuRequest,omitempty" name:"AgentCpuRequest"`

	// agent容器最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentCpuLimit *string `json:"AgentCpuLimit,omitempty" name:"AgentCpuLimit"`

	// agent容器初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentMemRequest *string `json:"AgentMemRequest,omitempty" name:"AgentMemRequest"`

	// agent容器最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentMemLimit *string `json:"AgentMemLimit,omitempty" name:"AgentMemLimit"`

	// istioproxy容器初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	IstioCpuRequest *string `json:"IstioCpuRequest,omitempty" name:"IstioCpuRequest"`

	// istioproxy容器最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	IstioCpuLimit *string `json:"IstioCpuLimit,omitempty" name:"IstioCpuLimit"`

	// istioproxy容器初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	IstioMemRequest *string `json:"IstioMemRequest,omitempty" name:"IstioMemRequest"`

	// istioproxy容器最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	IstioMemLimit *string `json:"IstioMemLimit,omitempty" name:"IstioMemLimit"`

	// 部署组的环境变量数组，这里没有展示 tsf 使用的环境变量，只展示了用户设置的环境变量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envs []*Env `json:"Envs,omitempty" name:"Envs"`

	// 健康检查配置信息，若不指定该参数，则默认不设置健康检查。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`

	// 是否部署Agent容器
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployAgent *bool `json:"DeployAgent,omitempty" name:"DeployAgent"`

	// 部署组备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 是否创建 k8s service
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisableService *bool `json:"DisableService,omitempty" name:"DisableService"`

	// service 是否为 headless 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadlessService *bool `json:"HeadlessService,omitempty" name:"HeadlessService"`

	// TcrRepoInfo值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcrRepoInfo *TcrRepoInfo `json:"TcrRepoInfo,omitempty" name:"TcrRepoInfo"`

	// 数据卷信息，list
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeInfos []*VolumeInfo `json:"VolumeInfos,omitempty" name:"VolumeInfos"`

	// 数据卷挂载信息，list
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeMountInfos []*VolumeMountInfo `json:"VolumeMountInfos,omitempty" name:"VolumeMountInfos"`

	// KubeInjectEnable值
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeInjectEnable *bool `json:"KubeInjectEnable,omitempty" name:"KubeInjectEnable"`

	// 仓库类型 (person, tcr)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// 预热配置设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmupSetting *WarmupSetting `json:"WarmupSetting,omitempty" name:"WarmupSetting"`

	// Envoy网关服务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayConfig *GatewayConfig `json:"GatewayConfig,omitempty" name:"GatewayConfig"`
}

type ContainerGroupDetail struct {
	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 已启动实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 负载均衡ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	LbIp *string `json:"LbIp,omitempty" name:"LbIp"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// Service ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// NodePort端口，只有公网和NodePort访问方式才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePort *int64 `json:"NodePort,omitempty" name:"NodePort"`

	// 最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 0:公网 1:集群内访问 2：NodePort
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *uint64 `json:"AccessType,omitempty" name:"AccessType"`

	// 更新方式：0:快速更新 1:滚动更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 端口数组对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts"`

	// 环境变量数组对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envs []*Env `json:"Envs,omitempty" name:"Envs"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// pod错误信息描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 部署组资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 部署组实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 部署组更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// kubernetes滚动更新策略的MaxSurge参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSurge *string `json:"MaxSurge,omitempty" name:"MaxSurge"`

	// kubernetes滚动更新策略的MaxUnavailable参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxUnavailable *string `json:"MaxUnavailable,omitempty" name:"MaxUnavailable"`

	// 部署组健康检查设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`

	// 允许PlainYamlDeploy
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowPlainYamlDeploy *bool `json:"AllowPlainYamlDeploy,omitempty" name:"AllowPlainYamlDeploy"`

	// 是否不等于ServiceConfig
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNotEqualServiceConfig *bool `json:"IsNotEqualServiceConfig,omitempty" name:"IsNotEqualServiceConfig"`

	// 仓库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

// Predefined struct for user
type ContinueRunFailedTaskBatchRequestParams struct {
	// 批次ID。
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

type ContinueRunFailedTaskBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批次ID。
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *ContinueRunFailedTaskBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueRunFailedTaskBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ContinueRunFailedTaskBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ContinueRunFailedTaskBatchResponseParams struct {
	// 成功或失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ContinueRunFailedTaskBatchResponse struct {
	*tchttp.BaseResponse
	Response *ContinueRunFailedTaskBatchResponseParams `json:"Response"`
}

func (r *ContinueRunFailedTaskBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueRunFailedTaskBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosCredentials struct {
	// 会话Token
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// 临时应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpAppId *string `json:"TmpAppId,omitempty" name:"TmpAppId"`

	// 临时调用者身份ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 所在域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type CosDownloadInfo struct {
	// 桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 鉴权信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Credentials *CosCredentials `json:"Credentials,omitempty" name:"Credentials"`
}

type CosUploadInfo struct {
	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 目标地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 鉴权信息
	Credentials *CosCredentials `json:"Credentials,omitempty" name:"Credentials"`
}

// Predefined struct for user
type CreateAllGatewayApiAsyncRequestParams struct {
	// API分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

type CreateAllGatewayApiAsyncRequest struct {
	*tchttp.BaseRequest
	
	// API分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *CreateAllGatewayApiAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllGatewayApiAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "MicroserviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAllGatewayApiAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAllGatewayApiAsyncResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAllGatewayApiAsyncResponse struct {
	*tchttp.BaseResponse
	Response *CreateAllGatewayApiAsyncResponseParams `json:"Response"`
}

func (r *CreateAllGatewayApiAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllGatewayApiAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiGroupRequestParams struct {
	// 分组名称, 不能包含中文
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组上下文
	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`

	// 鉴权类型。secret： 密钥鉴权； none:无鉴权
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// 分组类型,默认ms。 ms： 微服务分组； external:外部Api分组
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 命名空间参数key值
	NamespaceNameKey *string `json:"NamespaceNameKey,omitempty" name:"NamespaceNameKey"`

	// 微服务名参数key值
	ServiceNameKey *string `json:"ServiceNameKey,omitempty" name:"ServiceNameKey"`

	// 命名空间参数位置，path，header或query，默认是path
	NamespaceNameKeyPosition *string `json:"NamespaceNameKeyPosition,omitempty" name:"NamespaceNameKeyPosition"`

	// 微服务名参数位置，path，header或query，默认是path
	ServiceNameKeyPosition *string `json:"ServiceNameKeyPosition,omitempty" name:"ServiceNameKeyPosition"`
}

type CreateApiGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分组名称, 不能包含中文
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组上下文
	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`

	// 鉴权类型。secret： 密钥鉴权； none:无鉴权
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// 分组类型,默认ms。 ms： 微服务分组； external:外部Api分组
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 命名空间参数key值
	NamespaceNameKey *string `json:"NamespaceNameKey,omitempty" name:"NamespaceNameKey"`

	// 微服务名参数key值
	ServiceNameKey *string `json:"ServiceNameKey,omitempty" name:"ServiceNameKey"`

	// 命名空间参数位置，path，header或query，默认是path
	NamespaceNameKeyPosition *string `json:"NamespaceNameKeyPosition,omitempty" name:"NamespaceNameKeyPosition"`

	// 微服务名参数位置，path，header或query，默认是path
	ServiceNameKeyPosition *string `json:"ServiceNameKeyPosition,omitempty" name:"ServiceNameKeyPosition"`
}

func (r *CreateApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupContext")
	delete(f, "AuthType")
	delete(f, "Description")
	delete(f, "GroupType")
	delete(f, "GatewayInstanceId")
	delete(f, "NamespaceNameKey")
	delete(f, "ServiceNameKey")
	delete(f, "NamespaceNameKeyPosition")
	delete(f, "ServiceNameKeyPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiGroupResponseParams struct {
	// API分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiGroupResponseParams `json:"Response"`
}

func (r *CreateApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiRateLimitRuleRequestParams struct {
	// Api Id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// qps值
	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// 开启/禁用，enabled/disabled, 不传默认开启
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`
}

type CreateApiRateLimitRuleRequest struct {
	*tchttp.BaseRequest
	
	// Api Id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// qps值
	MaxQps *uint64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// 开启/禁用，enabled/disabled, 不传默认开启
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`
}

func (r *CreateApiRateLimitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiRateLimitRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiId")
	delete(f, "MaxQps")
	delete(f, "UsableStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiRateLimitRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiRateLimitRuleResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApiRateLimitRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiRateLimitRuleResponseParams `json:"Response"`
}

func (r *CreateApiRateLimitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiRateLimitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationRequestParams struct {
	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型，V：虚拟机应用；C：容器应用；S：serverless应用
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用微服务类型，M：service mesh应用；N：普通应用；G：网关应用
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用描述
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 应用日志配置项，废弃参数
	ApplicationLogConfig *string `json:"ApplicationLogConfig,omitempty" name:"ApplicationLogConfig"`

	// 应用资源类型，废弃参数
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// 应用runtime类型
	ApplicationRuntimeType *string `json:"ApplicationRuntimeType,omitempty" name:"ApplicationRuntimeType"`

	// 需要绑定的数据集ID
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`

	// 服务配置信息列表
	ServiceConfigList []*ServiceConfig `json:"ServiceConfigList,omitempty" name:"ServiceConfigList"`

	// 忽略创建镜像仓库
	IgnoreCreateImageRepository *bool `json:"IgnoreCreateImageRepository,omitempty" name:"IgnoreCreateImageRepository"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型，V：虚拟机应用；C：容器应用；S：serverless应用
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用微服务类型，M：service mesh应用；N：普通应用；G：网关应用
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用描述
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 应用日志配置项，废弃参数
	ApplicationLogConfig *string `json:"ApplicationLogConfig,omitempty" name:"ApplicationLogConfig"`

	// 应用资源类型，废弃参数
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// 应用runtime类型
	ApplicationRuntimeType *string `json:"ApplicationRuntimeType,omitempty" name:"ApplicationRuntimeType"`

	// 需要绑定的数据集ID
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`

	// 服务配置信息列表
	ServiceConfigList []*ServiceConfig `json:"ServiceConfigList,omitempty" name:"ServiceConfigList"`

	// 忽略创建镜像仓库
	IgnoreCreateImageRepository *bool `json:"IgnoreCreateImageRepository,omitempty" name:"IgnoreCreateImageRepository"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationName")
	delete(f, "ApplicationType")
	delete(f, "MicroserviceType")
	delete(f, "ApplicationDesc")
	delete(f, "ApplicationLogConfig")
	delete(f, "ApplicationResourceType")
	delete(f, "ApplicationRuntimeType")
	delete(f, "ProgramId")
	delete(f, "ServiceConfigList")
	delete(f, "IgnoreCreateImageRepository")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationResponseParams struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationResponseParams `json:"Response"`
}

func (r *CreateApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 分配给集群容器和服务IP的CIDR
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 集群备注
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群所属TSF地域
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF可用区
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 集群中每个Node上最大的Pod数量。取值范围4～256。不为2的幂值时会向上取最接近的2的幂值。
	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`

	// 集群最大的service数量。取值范围32～32768，不为2的幂值时会向上取最接近的2的幂值。
	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`

	// 需要绑定的数据集ID
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`

	// api地址
	KuberneteApiServer *string `json:"KuberneteApiServer,omitempty" name:"KuberneteApiServer"`

	// K : kubeconfig, S : service account
	KuberneteNativeType *string `json:"KuberneteNativeType,omitempty" name:"KuberneteNativeType"`

	// native secret
	KuberneteNativeSecret *string `json:"KuberneteNativeSecret,omitempty" name:"KuberneteNativeSecret"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 分配给集群容器和服务IP的CIDR
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 集群备注
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群所属TSF地域
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF可用区
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`

	// 集群中每个Node上最大的Pod数量。取值范围4～256。不为2的幂值时会向上取最接近的2的幂值。
	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`

	// 集群最大的service数量。取值范围32～32768，不为2的幂值时会向上取最接近的2的幂值。
	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`

	// 需要绑定的数据集ID
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`

	// api地址
	KuberneteApiServer *string `json:"KuberneteApiServer,omitempty" name:"KuberneteApiServer"`

	// K : kubeconfig, S : service account
	KuberneteNativeType *string `json:"KuberneteNativeType,omitempty" name:"KuberneteNativeType"`

	// native secret
	KuberneteNativeSecret *string `json:"KuberneteNativeSecret,omitempty" name:"KuberneteNativeSecret"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterName")
	delete(f, "ClusterType")
	delete(f, "VpcId")
	delete(f, "ClusterCIDR")
	delete(f, "ClusterDesc")
	delete(f, "TsfRegionId")
	delete(f, "TsfZoneId")
	delete(f, "SubnetId")
	delete(f, "ClusterVersion")
	delete(f, "MaxNodePodNum")
	delete(f, "MaxClusterServiceNum")
	delete(f, "ProgramId")
	delete(f, "KuberneteApiServer")
	delete(f, "KuberneteNativeType")
	delete(f, "KuberneteNativeSecret")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigRequestParams struct {
	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项值
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项值类型
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// Base64编码的配置项
	EncodeWithBase64 *bool `json:"EncodeWithBase64,omitempty" name:"EncodeWithBase64"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项值
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项值类型
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// Base64编码的配置项
	EncodeWithBase64 *bool `json:"EncodeWithBase64,omitempty" name:"EncodeWithBase64"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigName")
	delete(f, "ConfigVersion")
	delete(f, "ConfigValue")
	delete(f, "ApplicationId")
	delete(f, "ConfigVersionDesc")
	delete(f, "ConfigType")
	delete(f, "EncodeWithBase64")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigResponseParams struct {
	// true：创建成功；false：创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigResponseParams `json:"Response"`
}

func (r *CreateConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigTemplateRequestParams struct {
	// 配置模板名称
	ConfigTemplateName *string `json:"ConfigTemplateName,omitempty" name:"ConfigTemplateName"`

	// 配置模板对应的微服务框架
	ConfigTemplateType *string `json:"ConfigTemplateType,omitempty" name:"ConfigTemplateType"`

	// 配置模板数据
	ConfigTemplateValue *string `json:"ConfigTemplateValue,omitempty" name:"ConfigTemplateValue"`

	// 配置模板描述
	ConfigTemplateDesc *string `json:"ConfigTemplateDesc,omitempty" name:"ConfigTemplateDesc"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateConfigTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 配置模板名称
	ConfigTemplateName *string `json:"ConfigTemplateName,omitempty" name:"ConfigTemplateName"`

	// 配置模板对应的微服务框架
	ConfigTemplateType *string `json:"ConfigTemplateType,omitempty" name:"ConfigTemplateType"`

	// 配置模板数据
	ConfigTemplateValue *string `json:"ConfigTemplateValue,omitempty" name:"ConfigTemplateValue"`

	// 配置模板描述
	ConfigTemplateDesc *string `json:"ConfigTemplateDesc,omitempty" name:"ConfigTemplateDesc"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateConfigTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigTemplateName")
	delete(f, "ConfigTemplateType")
	delete(f, "ConfigTemplateValue")
	delete(f, "ConfigTemplateDesc")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigTemplateResponseParams struct {
	// true：创建成功；false：创建失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateConfigTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigTemplateResponseParams `json:"Response"`
}

func (r *CreateConfigTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContainGroupRequestParams struct {
	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 分组所属命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 分组名称字段，长度1~60，字母或下划线开头，可包含字母数字下划线
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 数组对象，见下方定义
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 最大分配 CPU 核数，对应 K8S limit
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配内存 MiB 数，对应 K8S limit
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 分组备注字段，长度应不大于200字符
	GroupComment *string `json:"GroupComment,omitempty" name:"GroupComment"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 滚动更新必填，更新间隔
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 初始分配的 CPU 核数，对应 K8S request
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配的内存 MiB 数，对应 K8S request
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 部署组资源类型
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// agent 容器分配的 CPU 核数，对应 K8S 的 request
	AgentCpuRequest *string `json:"AgentCpuRequest,omitempty" name:"AgentCpuRequest"`

	// agent 容器最大的 CPU 核数，对应 K8S 的 limit
	AgentCpuLimit *string `json:"AgentCpuLimit,omitempty" name:"AgentCpuLimit"`

	// agent 容器分配的内存 MiB 数，对应 K8S 的 request
	AgentMemRequest *string `json:"AgentMemRequest,omitempty" name:"AgentMemRequest"`

	// agent 容器最大的内存 MiB 数，对应 K8S 的 limit
	AgentMemLimit *string `json:"AgentMemLimit,omitempty" name:"AgentMemLimit"`

	// istioproxy 容器分配的 CPU 核数，对应 K8S 的 request
	IstioCpuRequest *string `json:"IstioCpuRequest,omitempty" name:"IstioCpuRequest"`

	// istioproxy 容器最大的 CPU 核数，对应 K8S 的 limit
	IstioCpuLimit *string `json:"IstioCpuLimit,omitempty" name:"IstioCpuLimit"`

	// istioproxy 容器分配的内存 MiB 数，对应 K8S 的 request
	IstioMemRequest *string `json:"IstioMemRequest,omitempty" name:"IstioMemRequest"`

	// istioproxy 容器最大的内存 MiB 数，对应 K8S 的 limit
	IstioMemLimit *string `json:"IstioMemLimit,omitempty" name:"IstioMemLimit"`
}

type CreateContainGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 分组所属命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 分组名称字段，长度1~60，字母或下划线开头，可包含字母数字下划线
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 数组对象，见下方定义
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 最大分配 CPU 核数，对应 K8S limit
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配内存 MiB 数，对应 K8S limit
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 分组备注字段，长度应不大于200字符
	GroupComment *string `json:"GroupComment,omitempty" name:"GroupComment"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 滚动更新必填，更新间隔
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 初始分配的 CPU 核数，对应 K8S request
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配的内存 MiB 数，对应 K8S request
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 部署组资源类型
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// agent 容器分配的 CPU 核数，对应 K8S 的 request
	AgentCpuRequest *string `json:"AgentCpuRequest,omitempty" name:"AgentCpuRequest"`

	// agent 容器最大的 CPU 核数，对应 K8S 的 limit
	AgentCpuLimit *string `json:"AgentCpuLimit,omitempty" name:"AgentCpuLimit"`

	// agent 容器分配的内存 MiB 数，对应 K8S 的 request
	AgentMemRequest *string `json:"AgentMemRequest,omitempty" name:"AgentMemRequest"`

	// agent 容器最大的内存 MiB 数，对应 K8S 的 limit
	AgentMemLimit *string `json:"AgentMemLimit,omitempty" name:"AgentMemLimit"`

	// istioproxy 容器分配的 CPU 核数，对应 K8S 的 request
	IstioCpuRequest *string `json:"IstioCpuRequest,omitempty" name:"IstioCpuRequest"`

	// istioproxy 容器最大的 CPU 核数，对应 K8S 的 limit
	IstioCpuLimit *string `json:"IstioCpuLimit,omitempty" name:"IstioCpuLimit"`

	// istioproxy 容器分配的内存 MiB 数，对应 K8S 的 request
	IstioMemRequest *string `json:"IstioMemRequest,omitempty" name:"IstioMemRequest"`

	// istioproxy 容器最大的内存 MiB 数，对应 K8S 的 limit
	IstioMemLimit *string `json:"IstioMemLimit,omitempty" name:"IstioMemLimit"`
}

func (r *CreateContainGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContainGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "NamespaceId")
	delete(f, "GroupName")
	delete(f, "InstanceNum")
	delete(f, "AccessType")
	delete(f, "ProtocolPorts")
	delete(f, "ClusterId")
	delete(f, "CpuLimit")
	delete(f, "MemLimit")
	delete(f, "GroupComment")
	delete(f, "UpdateType")
	delete(f, "UpdateIvl")
	delete(f, "CpuRequest")
	delete(f, "MemRequest")
	delete(f, "GroupResourceType")
	delete(f, "SubnetId")
	delete(f, "AgentCpuRequest")
	delete(f, "AgentCpuLimit")
	delete(f, "AgentMemRequest")
	delete(f, "AgentMemLimit")
	delete(f, "IstioCpuRequest")
	delete(f, "IstioCpuLimit")
	delete(f, "IstioMemRequest")
	delete(f, "IstioMemLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateContainGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContainGroupResponseParams struct {
	// 返回创建成功的部署组ID，返回null表示失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateContainGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateContainGroupResponseParams `json:"Response"`
}

func (r *CreateContainGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContainGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileConfigRequestParams struct {
	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项文件名
	ConfigFileName *string `json:"ConfigFileName,omitempty" name:"ConfigFileName"`

	// 配置项文件内容（原始内容编码需要 utf-8 格式，如果 ConfigFileCode 为 gbk，后台会进行转换）
	ConfigFileValue *string `json:"ConfigFileValue,omitempty" name:"ConfigFileValue"`

	// 配置项关联应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 发布路径
	ConfigFilePath *string `json:"ConfigFilePath,omitempty" name:"ConfigFilePath"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项文件编码，utf-8 或 gbk。注：如果选择 gbk，需要新版本 tsf-consul-template （公有云虚拟机需要使用 1.32 tsf-agent，容器需要从文档中获取最新的 tsf-consul-template-docker.tar.gz）的支持
	ConfigFileCode *string `json:"ConfigFileCode,omitempty" name:"ConfigFileCode"`

	// 后置命令
	ConfigPostCmd *string `json:"ConfigPostCmd,omitempty" name:"ConfigPostCmd"`

	// Base64编码的配置项
	EncodeWithBase64 *bool `json:"EncodeWithBase64,omitempty" name:"EncodeWithBase64"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateFileConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项文件名
	ConfigFileName *string `json:"ConfigFileName,omitempty" name:"ConfigFileName"`

	// 配置项文件内容（原始内容编码需要 utf-8 格式，如果 ConfigFileCode 为 gbk，后台会进行转换）
	ConfigFileValue *string `json:"ConfigFileValue,omitempty" name:"ConfigFileValue"`

	// 配置项关联应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 发布路径
	ConfigFilePath *string `json:"ConfigFilePath,omitempty" name:"ConfigFilePath"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项文件编码，utf-8 或 gbk。注：如果选择 gbk，需要新版本 tsf-consul-template （公有云虚拟机需要使用 1.32 tsf-agent，容器需要从文档中获取最新的 tsf-consul-template-docker.tar.gz）的支持
	ConfigFileCode *string `json:"ConfigFileCode,omitempty" name:"ConfigFileCode"`

	// 后置命令
	ConfigPostCmd *string `json:"ConfigPostCmd,omitempty" name:"ConfigPostCmd"`

	// Base64编码的配置项
	EncodeWithBase64 *bool `json:"EncodeWithBase64,omitempty" name:"EncodeWithBase64"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigName")
	delete(f, "ConfigVersion")
	delete(f, "ConfigFileName")
	delete(f, "ConfigFileValue")
	delete(f, "ApplicationId")
	delete(f, "ConfigFilePath")
	delete(f, "ConfigVersionDesc")
	delete(f, "ConfigFileCode")
	delete(f, "ConfigPostCmd")
	delete(f, "EncodeWithBase64")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileConfigResponseParams struct {
	// true：创建成功；false：创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileConfigResponseParams `json:"Response"`
}

func (r *CreateFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGatewayApiRequestParams struct {
	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Api信息
	ApiList []*ApiInfo `json:"ApiList,omitempty" name:"ApiList"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateGatewayApiRequest struct {
	*tchttp.BaseRequest
	
	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Api信息
	ApiList []*ApiInfo `json:"ApiList,omitempty" name:"ApiList"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateGatewayApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGatewayApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ApiList")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGatewayApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGatewayApiResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGatewayApiResponse struct {
	*tchttp.BaseResponse
	Response *CreateGatewayApiResponseParams `json:"Response"`
}

func (r *CreateGatewayApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGatewayApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupRequestParams struct {
	// 部署组所属的应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组所属命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 部署组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组描述
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 部署组资源类型
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 部署组备注
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组所属的应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组所属命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 部署组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组描述
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 部署组资源类型
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 部署组备注
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "NamespaceId")
	delete(f, "GroupName")
	delete(f, "ClusterId")
	delete(f, "GroupDesc")
	delete(f, "GroupResourceType")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupResponseParams struct {
	// groupId， null表示创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupResponseParams `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaneRequestParams struct {
	// 泳道名称
	LaneName *string `json:"LaneName,omitempty" name:"LaneName"`

	// 泳道备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 泳道部署组信息
	LaneGroupList []*LaneGroup `json:"LaneGroupList,omitempty" name:"LaneGroupList"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateLaneRequest struct {
	*tchttp.BaseRequest
	
	// 泳道名称
	LaneName *string `json:"LaneName,omitempty" name:"LaneName"`

	// 泳道备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 泳道部署组信息
	LaneGroupList []*LaneGroup `json:"LaneGroupList,omitempty" name:"LaneGroupList"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateLaneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaneName")
	delete(f, "Remark")
	delete(f, "LaneGroupList")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLaneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaneResponseParams struct {
	// 泳道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLaneResponse struct {
	*tchttp.BaseResponse
	Response *CreateLaneResponseParams `json:"Response"`
}

func (r *CreateLaneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaneRuleRequestParams struct {
	// 泳道规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 泳道规则备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 泳道规则标签列表
	RuleTagList []*LaneRuleTag `json:"RuleTagList,omitempty" name:"RuleTagList"`

	// 泳道规则标签关系
	RuleTagRelationship *string `json:"RuleTagRelationship,omitempty" name:"RuleTagRelationship"`

	// 泳道Id
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateLaneRuleRequest struct {
	*tchttp.BaseRequest
	
	// 泳道规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 泳道规则备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 泳道规则标签列表
	RuleTagList []*LaneRuleTag `json:"RuleTagList,omitempty" name:"RuleTagList"`

	// 泳道规则标签关系
	RuleTagRelationship *string `json:"RuleTagRelationship,omitempty" name:"RuleTagRelationship"`

	// 泳道Id
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateLaneRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaneRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "Remark")
	delete(f, "RuleTagList")
	delete(f, "RuleTagRelationship")
	delete(f, "LaneId")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLaneRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaneRuleResponseParams struct {
	// 泳道规则Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLaneRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateLaneRuleResponseParams `json:"Response"`
}

func (r *CreateLaneRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaneRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMicroserviceRequestParams struct {
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务描述信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

type CreateMicroserviceRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务描述信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

func (r *CreateMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMicroserviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "MicroserviceName")
	delete(f, "MicroserviceDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMicroserviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMicroserviceResponseParams struct {
	// 新增微服务是否成功。
	// true：操作成功。
	// false：操作失败。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *CreateMicroserviceResponseParams `json:"Response"`
}

func (r *CreateMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMicroserviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMicroserviceWithDetailRespRequestParams struct {
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务描述信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

type CreateMicroserviceWithDetailRespRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务描述信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

func (r *CreateMicroserviceWithDetailRespRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMicroserviceWithDetailRespRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "MicroserviceName")
	delete(f, "MicroserviceDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMicroserviceWithDetailRespRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMicroserviceWithDetailRespResponseParams struct {
	// id
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMicroserviceWithDetailRespResponse struct {
	*tchttp.BaseResponse
	Response *CreateMicroserviceWithDetailRespResponseParams `json:"Response"`
}

func (r *CreateMicroserviceWithDetailRespResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMicroserviceWithDetailRespResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceRequestParams struct {
	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间描述
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 命名空间资源类型(默认值为DEF)
	NamespaceResourceType *string `json:"NamespaceResourceType,omitempty" name:"NamespaceResourceType"`

	// 是否是全局命名空间(默认是DEF，表示普通命名空间；GLOBAL表示全局命名空间)
	NamespaceType *string `json:"NamespaceType,omitempty" name:"NamespaceType"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 是否开启高可用
	IsHaEnable *string `json:"IsHaEnable,omitempty" name:"IsHaEnable"`

	// 需要绑定的数据集ID
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间描述
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 命名空间资源类型(默认值为DEF)
	NamespaceResourceType *string `json:"NamespaceResourceType,omitempty" name:"NamespaceResourceType"`

	// 是否是全局命名空间(默认是DEF，表示普通命名空间；GLOBAL表示全局命名空间)
	NamespaceType *string `json:"NamespaceType,omitempty" name:"NamespaceType"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 是否开启高可用
	IsHaEnable *string `json:"IsHaEnable,omitempty" name:"IsHaEnable"`

	// 需要绑定的数据集ID
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceName")
	delete(f, "ClusterId")
	delete(f, "NamespaceDesc")
	delete(f, "NamespaceResourceType")
	delete(f, "NamespaceType")
	delete(f, "NamespaceId")
	delete(f, "IsHaEnable")
	delete(f, "ProgramId")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceResponseParams struct {
	// 成功时为命名空间ID，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNamespaceResponseParams `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePathRewritesRequestParams struct {
	// 路径重写列表
	PathRewrites *PathRewriteCreateObject `json:"PathRewrites,omitempty" name:"PathRewrites"`
}

type CreatePathRewritesRequest struct {
	*tchttp.BaseRequest
	
	// 路径重写列表
	PathRewrites *PathRewriteCreateObject `json:"PathRewrites,omitempty" name:"PathRewrites"`
}

func (r *CreatePathRewritesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePathRewritesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PathRewrites")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePathRewritesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePathRewritesResponseParams struct {
	// true/false
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePathRewritesResponse struct {
	*tchttp.BaseResponse
	Response *CreatePathRewritesResponseParams `json:"Response"`
}

func (r *CreatePathRewritesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePathRewritesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePublicConfigRequestParams struct {
	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项值，总是接收yaml格式的内容
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项类型
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// Base64编码的配置项
	EncodeWithBase64 *bool `json:"EncodeWithBase64,omitempty" name:"EncodeWithBase64"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreatePublicConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项值，总是接收yaml格式的内容
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项类型
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// Base64编码的配置项
	EncodeWithBase64 *bool `json:"EncodeWithBase64,omitempty" name:"EncodeWithBase64"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreatePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePublicConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigName")
	delete(f, "ConfigVersion")
	delete(f, "ConfigValue")
	delete(f, "ConfigVersionDesc")
	delete(f, "ConfigType")
	delete(f, "EncodeWithBase64")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePublicConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePublicConfigResponseParams struct {
	// true：创建成功；false：创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreatePublicConfigResponseParams `json:"Response"`
}

func (r *CreatePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePublicConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryRequestParams struct {
	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 仓库类型（默认仓库：default，私有仓库：private）
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 仓库所在桶名称
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 仓库所在桶地域
	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`

	// 目录
	Directory *string `json:"Directory,omitempty" name:"Directory"`

	// 仓库描述
	RepositoryDesc *string `json:"RepositoryDesc,omitempty" name:"RepositoryDesc"`
}

type CreateRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 仓库类型（默认仓库：default，私有仓库：private）
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 仓库所在桶名称
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 仓库所在桶地域
	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`

	// 目录
	Directory *string `json:"Directory,omitempty" name:"Directory"`

	// 仓库描述
	RepositoryDesc *string `json:"RepositoryDesc,omitempty" name:"RepositoryDesc"`
}

func (r *CreateRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepositoryName")
	delete(f, "RepositoryType")
	delete(f, "BucketName")
	delete(f, "BucketRegion")
	delete(f, "Directory")
	delete(f, "RepositoryDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryResponseParams struct {
	// 创建仓库是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateRepositoryResponseParams `json:"Response"`
}

func (r *CreateRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFlowRequestParams struct {
	// 工作流名称
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 触发方式
	TriggerRule *TaskRule `json:"TriggerRule,omitempty" name:"TriggerRule"`

	// 工作流任务节点列表
	FlowEdges []*TaskFlowEdge `json:"FlowEdges,omitempty" name:"FlowEdges"`

	// 工作流执行超时时间
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateTaskFlowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流名称
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 触发方式
	TriggerRule *TaskRule `json:"TriggerRule,omitempty" name:"TriggerRule"`

	// 工作流任务节点列表
	FlowEdges []*TaskFlowEdge `json:"FlowEdges,omitempty" name:"FlowEdges"`

	// 工作流执行超时时间
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateTaskFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowName")
	delete(f, "TriggerRule")
	delete(f, "FlowEdges")
	delete(f, "TimeOut")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFlowResponseParams struct {
	// 工作流 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTaskFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskFlowResponseParams `json:"Response"`
}

func (r *CreateTaskFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// 任务名称，任务长度64字符
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务内容，长度限制65536个字节
	TaskContent *string `json:"TaskContent,omitempty" name:"TaskContent"`

	// 执行类型，unicast/broadcast
	ExecuteType *string `json:"ExecuteType,omitempty" name:"ExecuteType"`

	// 任务类型,java
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务超时时间， 时间单位 ms
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 触发规则
	TaskRule *TaskRule `json:"TaskRule,omitempty" name:"TaskRule"`

	// 重试次数，0 <= RetryCount<= 10
	RetryCount *uint64 `json:"RetryCount,omitempty" name:"RetryCount"`

	// 重试间隔， 0 <= RetryInterval <= 600000， 时间单位 ms
	RetryInterval *uint64 `json:"RetryInterval,omitempty" name:"RetryInterval"`

	// 分片数量
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 分片参数
	ShardArguments []*ShardArgument `json:"ShardArguments,omitempty" name:"ShardArguments"`

	// 判断任务成功的操作符
	SuccessOperator *string `json:"SuccessOperator,omitempty" name:"SuccessOperator"`

	// 判断任务成功率的阈值，如100
	SuccessRatio *string `json:"SuccessRatio,omitempty" name:"SuccessRatio"`

	// 高级设置
	AdvanceSettings *AdvanceSettings `json:"AdvanceSettings,omitempty" name:"AdvanceSettings"`

	// 任务参数，长度限制10000个字符
	TaskArgument *string `json:"TaskArgument,omitempty" name:"TaskArgument"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称，任务长度64字符
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务内容，长度限制65536个字节
	TaskContent *string `json:"TaskContent,omitempty" name:"TaskContent"`

	// 执行类型，unicast/broadcast
	ExecuteType *string `json:"ExecuteType,omitempty" name:"ExecuteType"`

	// 任务类型,java
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务超时时间， 时间单位 ms
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 触发规则
	TaskRule *TaskRule `json:"TaskRule,omitempty" name:"TaskRule"`

	// 重试次数，0 <= RetryCount<= 10
	RetryCount *uint64 `json:"RetryCount,omitempty" name:"RetryCount"`

	// 重试间隔， 0 <= RetryInterval <= 600000， 时间单位 ms
	RetryInterval *uint64 `json:"RetryInterval,omitempty" name:"RetryInterval"`

	// 分片数量
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 分片参数
	ShardArguments []*ShardArgument `json:"ShardArguments,omitempty" name:"ShardArguments"`

	// 判断任务成功的操作符
	SuccessOperator *string `json:"SuccessOperator,omitempty" name:"SuccessOperator"`

	// 判断任务成功率的阈值，如100
	SuccessRatio *string `json:"SuccessRatio,omitempty" name:"SuccessRatio"`

	// 高级设置
	AdvanceSettings *AdvanceSettings `json:"AdvanceSettings,omitempty" name:"AdvanceSettings"`

	// 任务参数，长度限制10000个字符
	TaskArgument *string `json:"TaskArgument,omitempty" name:"TaskArgument"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "TaskContent")
	delete(f, "ExecuteType")
	delete(f, "TaskType")
	delete(f, "TimeOut")
	delete(f, "GroupId")
	delete(f, "TaskRule")
	delete(f, "RetryCount")
	delete(f, "RetryInterval")
	delete(f, "ShardCount")
	delete(f, "ShardArguments")
	delete(f, "SuccessOperator")
	delete(f, "SuccessRatio")
	delete(f, "AdvanceSettings")
	delete(f, "TaskArgument")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskResponseParams `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUnitRuleRequestParams struct {
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则项列表
	UnitRuleItemList []*UnitRuleItem `json:"UnitRuleItemList,omitempty" name:"UnitRuleItemList"`
}

type CreateUnitRuleRequest struct {
	*tchttp.BaseRequest
	
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则项列表
	UnitRuleItemList []*UnitRuleItem `json:"UnitRuleItemList,omitempty" name:"UnitRuleItemList"`
}

func (r *CreateUnitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUnitRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayInstanceId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "UnitRuleItemList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUnitRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUnitRuleResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUnitRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateUnitRuleResponseParams `json:"Response"`
}

func (r *CreateUnitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUnitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CurvePoint struct {
	// 当前坐标 X轴的值 当前是日期格式:"yyyy-MM-dd HH:mm:ss"
	Label *string `json:"Label,omitempty" name:"Label"`

	// 当前坐标 Y轴的值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 该坐标点时间戳
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
}

// Predefined struct for user
type DeleteApiGroupRequestParams struct {
	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DeleteApiGroupRequest struct {
	*tchttp.BaseRequest
	
	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApiGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiGroupResponseParams struct {
	// 成功失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApiGroupResponseParams `json:"Response"`
}

func (r *DeleteApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DeleteApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DeleteApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationResponseParams struct {
	// 删除应用操作是否成功。
	// true：操作成功。
	// false：操作失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationResponseParams `json:"Response"`
}

func (r *DeleteApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否只解绑，不删除容器集群，默认不传则删除容器集群。
	Unbind *bool `json:"Unbind,omitempty" name:"Unbind"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否只解绑，不删除容器集群，默认不传则删除容器集群。
	Unbind *bool `json:"Unbind,omitempty" name:"Unbind"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Unbind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// 删除集群操作是否成功。
	// true：操作成功。
	// false：操作失败。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterResponseParams `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigRequestParams struct {
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

type DeleteConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigResponseParams struct {
	// true：删除成功；false：删除失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigResponseParams `json:"Response"`
}

func (r *DeleteConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigTemplateRequestParams struct {
	// 无
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
}

type DeleteConfigTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 无
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
}

func (r *DeleteConfigTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigTemplateResponseParams struct {
	// true：删除成功；false：删除失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteConfigTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigTemplateResponseParams `json:"Response"`
}

func (r *DeleteConfigTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContainerGroupRequestParams struct {
	// 部署组ID，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DeleteContainerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContainerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteContainerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteContainerGroupResponseParams struct {
	// 删除操作是否成功：
	// true：成功
	// false：失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteContainerGroupResponseParams `json:"Response"`
}

func (r *DeleteContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteContainerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileConfigRequestParams struct {
	// 文件配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

type DeleteFileConfigRequest struct {
	*tchttp.BaseRequest
	
	// 文件配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileConfigResponseParams struct {
	// 删除结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileConfigResponseParams `json:"Response"`
}

func (r *DeleteFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// 删除部署组操作是否成功。
	// true：操作成功。
	// false：操作失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImageTag struct {
	// 仓库名，如/tsf/nginx
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 版本号:如V1
	TagName *string `json:"TagName,omitempty" name:"TagName"`
}

// Predefined struct for user
type DeleteImageTagsRequestParams struct {
	// 镜像版本数组
	ImageTags []*DeleteImageTag `json:"ImageTags,omitempty" name:"ImageTags"`

	// 企业: tcr ；个人: personal或者不填
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
}

type DeleteImageTagsRequest struct {
	*tchttp.BaseRequest
	
	// 镜像版本数组
	ImageTags []*DeleteImageTag `json:"ImageTags,omitempty" name:"ImageTags"`

	// 企业: tcr ；个人: personal或者不填
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`
}

func (r *DeleteImageTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageTags")
	delete(f, "RepoType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageTagsResponseParams struct {
	// 批量删除操作是否成功。
	// true：成功。
	// false：失败。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImageTagsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageTagsResponseParams `json:"Response"`
}

func (r *DeleteImageTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaneRequestParams struct {
	// 泳道Idl
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`
}

type DeleteLaneRequest struct {
	*tchttp.BaseRequest
	
	// 泳道Idl
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`
}

func (r *DeleteLaneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLaneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaneResponseParams struct {
	// true / false
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLaneResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLaneResponseParams `json:"Response"`
}

func (r *DeleteLaneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaneRuleRequestParams struct {
	// 泳道规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteLaneRuleRequest struct {
	*tchttp.BaseRequest
	
	// 泳道规则Id
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteLaneRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaneRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLaneRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaneRuleResponseParams struct {
	// 操作状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLaneRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLaneRuleResponseParams `json:"Response"`
}

func (r *DeleteLaneRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaneRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMicroserviceRequestParams struct {
	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

type DeleteMicroserviceRequest struct {
	*tchttp.BaseRequest
	
	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *DeleteMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMicroserviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MicroserviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMicroserviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMicroserviceResponseParams struct {
	// 删除微服务是否成功。
	// true：操作成功。
	// false：操作失败。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMicroserviceResponseParams `json:"Response"`
}

func (r *DeleteMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMicroserviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceRequestParams struct {
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceResponseParams struct {
	// 删除命名空间是否成功。
	// true：删除成功。
	// false：删除失败。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNamespaceResponseParams `json:"Response"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePathRewritesRequestParams struct {
	// 路径重写规则IDs
	PathRewriteIds []*string `json:"PathRewriteIds,omitempty" name:"PathRewriteIds"`
}

type DeletePathRewritesRequest struct {
	*tchttp.BaseRequest
	
	// 路径重写规则IDs
	PathRewriteIds []*string `json:"PathRewriteIds,omitempty" name:"PathRewriteIds"`
}

func (r *DeletePathRewritesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePathRewritesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PathRewriteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePathRewritesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePathRewritesResponseParams struct {
	// true/false
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePathRewritesResponse struct {
	*tchttp.BaseResponse
	Response *DeletePathRewritesResponseParams `json:"Response"`
}

func (r *DeletePathRewritesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePathRewritesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePkgsRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 需要删除的程序包ID列表
	PkgIds []*string `json:"PkgIds,omitempty" name:"PkgIds"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 程序包仓库id
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

type DeletePkgsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 需要删除的程序包ID列表
	PkgIds []*string `json:"PkgIds,omitempty" name:"PkgIds"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 程序包仓库id
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

func (r *DeletePkgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePkgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PkgIds")
	delete(f, "RepositoryType")
	delete(f, "RepositoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePkgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePkgsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePkgsResponse struct {
	*tchttp.BaseResponse
	Response *DeletePkgsResponseParams `json:"Response"`
}

func (r *DeletePkgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePkgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePublicConfigRequestParams struct {
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

type DeletePublicConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeletePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePublicConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePublicConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePublicConfigResponseParams struct {
	// true：删除成功；false：删除失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeletePublicConfigResponseParams `json:"Response"`
}

func (r *DeletePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePublicConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryRequestParams struct {
	// 仓库ID
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

type DeleteRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 仓库ID
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

func (r *DeleteRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepositoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryResponseParams struct {
	// 删除仓库是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRepositoryResponseParams `json:"Response"`
}

func (r *DeleteRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServerlessGroupRequestParams struct {
	// groupId，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DeleteServerlessGroupRequest struct {
	*tchttp.BaseRequest
	
	// groupId，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServerlessGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServerlessGroupResponseParams struct {
	// 结果true：成功；false：失败。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServerlessGroupResponseParams `json:"Response"`
}

func (r *DeleteServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskResponseParams struct {
	// 删除成功or失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskResponseParams `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUnitNamespacesRequestParams struct {
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 单元化命名空间ID数组
	UnitNamespaceList []*string `json:"UnitNamespaceList,omitempty" name:"UnitNamespaceList"`
}

type DeleteUnitNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 单元化命名空间ID数组
	UnitNamespaceList []*string `json:"UnitNamespaceList,omitempty" name:"UnitNamespaceList"`
}

func (r *DeleteUnitNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUnitNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayInstanceId")
	delete(f, "UnitNamespaceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUnitNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUnitNamespacesResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUnitNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUnitNamespacesResponseParams `json:"Response"`
}

func (r *DeleteUnitNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUnitNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUnitRuleRequestParams struct {
	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DeleteUnitRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteUnitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUnitRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUnitRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUnitRuleResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUnitRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUnitRuleResponseParams `json:"Response"`
}

func (r *DeleteUnitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUnitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeliveryConfigBindGroup struct {
	// 配置id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置名
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 采集路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectPath []*string `json:"CollectPath,omitempty" name:"CollectPath"`

	// 关联部署组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*GroupInfo `json:"Groups,omitempty" name:"Groups"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// KafkaVIp
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaVIp *string `json:"KafkaVIp,omitempty" name:"KafkaVIp"`

	// KafkaAddress
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaAddress *string `json:"KafkaAddress,omitempty" name:"KafkaAddress"`

	// KafkaVPort
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaVPort *string `json:"KafkaVPort,omitempty" name:"KafkaVPort"`

	// Topic
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// LineRule
	// 注意：此字段可能返回 null，表示取不到有效值。
	LineRule *string `json:"LineRule,omitempty" name:"LineRule"`

	// CustomRule
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomRule *string `json:"CustomRule,omitempty" name:"CustomRule"`

	// EnableGlobalLineRule
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableGlobalLineRule *bool `json:"EnableGlobalLineRule,omitempty" name:"EnableGlobalLineRule"`

	// EnableAuth
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableAuth *bool `json:"EnableAuth,omitempty" name:"EnableAuth"`

	// Username
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitempty" name:"Username"`

	// Password
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// KafkaInfos
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaInfos []*DeliveryKafkaInfo `json:"KafkaInfos,omitempty" name:"KafkaInfos"`
}

type DeliveryConfigBindGroups struct {
	// 公共条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*DeliveryConfigBindGroup `json:"Content,omitempty" name:"Content"`
}

type DeliveryKafkaInfo struct {
	// 投递kafka的topic
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 采集日志的path
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path []*string `json:"Path,omitempty" name:"Path"`

	// default，默认换行符分行
	// time，按时间分行
	// custom, 选了custom那么CustomRule就要填入具体的自定义值
	// 注意：此字段可能返回 null，表示取不到有效值。
	LineRule *string `json:"LineRule,omitempty" name:"LineRule"`

	// 自定义的分行值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomRule *string `json:"CustomRule,omitempty" name:"CustomRule"`
}

// Predefined struct for user
type DeployContainerGroupRequestParams struct {
	// 部署组ID，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 镜像版本名称,如v1
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 镜像server
	Server *string `json:"Server,omitempty" name:"Server"`

	// 旧版镜像名，如/tsf/nginx
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 业务容器最大的 CPU 核数，对应 K8S 的 limit；不填时默认为 request 的 2 倍
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 业务容器最大的内存 MiB 数，对应 K8S 的 limit；不填时默认为 request 的 2 倍
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// jvm参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 业务容器分配的 CPU 核数，对应 K8S 的 request，默认0.25
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 业务容器分配的内存 MiB 数，对应 K8S 的 request，默认640 MiB
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 是否不立即启动
	DoNotStart *bool `json:"DoNotStart,omitempty" name:"DoNotStart"`

	// （优先使用）新版镜像名，如/tsf/nginx
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 滚动更新必填，更新间隔
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// agent 容器分配的 CPU 核数，对应 K8S 的 request
	AgentCpuRequest *string `json:"AgentCpuRequest,omitempty" name:"AgentCpuRequest"`

	// agent 容器最大的 CPU 核数，对应 K8S 的 limit
	AgentCpuLimit *string `json:"AgentCpuLimit,omitempty" name:"AgentCpuLimit"`

	// agent 容器分配的内存 MiB 数，对应 K8S 的 request
	AgentMemRequest *string `json:"AgentMemRequest,omitempty" name:"AgentMemRequest"`

	// agent 容器最大的内存 MiB 数，对应 K8S 的 limit
	AgentMemLimit *string `json:"AgentMemLimit,omitempty" name:"AgentMemLimit"`

	// istioproxy 容器分配的 CPU 核数，对应 K8S 的 request
	IstioCpuRequest *string `json:"IstioCpuRequest,omitempty" name:"IstioCpuRequest"`

	// istioproxy 容器最大的 CPU 核数，对应 K8S 的 limit
	IstioCpuLimit *string `json:"IstioCpuLimit,omitempty" name:"IstioCpuLimit"`

	// istioproxy 容器分配的内存 MiB 数，对应 K8S 的 request
	IstioMemRequest *string `json:"IstioMemRequest,omitempty" name:"IstioMemRequest"`

	// istioproxy 容器最大的内存 MiB 数，对应 K8S 的 limit
	IstioMemLimit *string `json:"IstioMemLimit,omitempty" name:"IstioMemLimit"`

	// kubernetes滚动更新策略的MaxSurge参数
	MaxSurge *string `json:"MaxSurge,omitempty" name:"MaxSurge"`

	// kubernetes滚动更新策略的MaxUnavailable参数
	MaxUnavailable *string `json:"MaxUnavailable,omitempty" name:"MaxUnavailable"`

	// 健康检查配置信息，若不指定该参数，则默认不设置健康检查。
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`

	// 部署组应用运行的环境变量。若不指定该参数，则默认不设置额外的环境变量。
	Envs []*Env `json:"Envs,omitempty" name:"Envs"`

	// 容器部署组的网络设置。
	ServiceSetting *ServiceSetting `json:"ServiceSetting,omitempty" name:"ServiceSetting"`

	// 是否部署 agent 容器。若不指定该参数，则默认不部署 agent 容器。
	DeployAgent *bool `json:"DeployAgent,omitempty" name:"DeployAgent"`

	// 节点调度策略。若不指定改参数，则默认不使用节点调度策略。
	SchedulingStrategy *SchedulingStrategy `json:"SchedulingStrategy,omitempty" name:"SchedulingStrategy"`

	// 是否进行增量部署，默认为false，全量更新
	IncrementalDeployment *bool `json:"IncrementalDeployment,omitempty" name:"IncrementalDeployment"`

	// tcr或者不填
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// 数据卷信息-废弃，请用VolumeInfoList参数
	VolumeInfos *VolumeInfo `json:"VolumeInfos,omitempty" name:"VolumeInfos"`

	// 数据卷挂载点信息-废弃，请用VolumeMountInfoList参数
	VolumeMountInfos *VolumeMountInfo `json:"VolumeMountInfos,omitempty" name:"VolumeMountInfos"`

	// 数据卷信息，list
	VolumeInfoList []*VolumeInfo `json:"VolumeInfoList,omitempty" name:"VolumeInfoList"`

	// 数据卷挂载点信息，list
	VolumeMountInfoList []*VolumeMountInfo `json:"VolumeMountInfoList,omitempty" name:"VolumeMountInfoList"`

	// 是否清除数据卷信息，默认false
	VolumeClean *bool `json:"VolumeClean,omitempty" name:"VolumeClean"`

	// javaagent信息: SERVICE_AGENT/OT_AGENT
	AgentProfileList []*AgentProfile `json:"AgentProfileList,omitempty" name:"AgentProfileList"`

	// 预热配置信息
	WarmupSetting *WarmupSetting `json:"WarmupSetting,omitempty" name:"WarmupSetting"`
}

type DeployContainerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 镜像版本名称,如v1
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 镜像server
	Server *string `json:"Server,omitempty" name:"Server"`

	// 旧版镜像名，如/tsf/nginx
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 业务容器最大的 CPU 核数，对应 K8S 的 limit；不填时默认为 request 的 2 倍
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 业务容器最大的内存 MiB 数，对应 K8S 的 limit；不填时默认为 request 的 2 倍
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// jvm参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 业务容器分配的 CPU 核数，对应 K8S 的 request，默认0.25
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 业务容器分配的内存 MiB 数，对应 K8S 的 request，默认640 MiB
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 是否不立即启动
	DoNotStart *bool `json:"DoNotStart,omitempty" name:"DoNotStart"`

	// （优先使用）新版镜像名，如/tsf/nginx
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 滚动更新必填，更新间隔
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// agent 容器分配的 CPU 核数，对应 K8S 的 request
	AgentCpuRequest *string `json:"AgentCpuRequest,omitempty" name:"AgentCpuRequest"`

	// agent 容器最大的 CPU 核数，对应 K8S 的 limit
	AgentCpuLimit *string `json:"AgentCpuLimit,omitempty" name:"AgentCpuLimit"`

	// agent 容器分配的内存 MiB 数，对应 K8S 的 request
	AgentMemRequest *string `json:"AgentMemRequest,omitempty" name:"AgentMemRequest"`

	// agent 容器最大的内存 MiB 数，对应 K8S 的 limit
	AgentMemLimit *string `json:"AgentMemLimit,omitempty" name:"AgentMemLimit"`

	// istioproxy 容器分配的 CPU 核数，对应 K8S 的 request
	IstioCpuRequest *string `json:"IstioCpuRequest,omitempty" name:"IstioCpuRequest"`

	// istioproxy 容器最大的 CPU 核数，对应 K8S 的 limit
	IstioCpuLimit *string `json:"IstioCpuLimit,omitempty" name:"IstioCpuLimit"`

	// istioproxy 容器分配的内存 MiB 数，对应 K8S 的 request
	IstioMemRequest *string `json:"IstioMemRequest,omitempty" name:"IstioMemRequest"`

	// istioproxy 容器最大的内存 MiB 数，对应 K8S 的 limit
	IstioMemLimit *string `json:"IstioMemLimit,omitempty" name:"IstioMemLimit"`

	// kubernetes滚动更新策略的MaxSurge参数
	MaxSurge *string `json:"MaxSurge,omitempty" name:"MaxSurge"`

	// kubernetes滚动更新策略的MaxUnavailable参数
	MaxUnavailable *string `json:"MaxUnavailable,omitempty" name:"MaxUnavailable"`

	// 健康检查配置信息，若不指定该参数，则默认不设置健康检查。
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`

	// 部署组应用运行的环境变量。若不指定该参数，则默认不设置额外的环境变量。
	Envs []*Env `json:"Envs,omitempty" name:"Envs"`

	// 容器部署组的网络设置。
	ServiceSetting *ServiceSetting `json:"ServiceSetting,omitempty" name:"ServiceSetting"`

	// 是否部署 agent 容器。若不指定该参数，则默认不部署 agent 容器。
	DeployAgent *bool `json:"DeployAgent,omitempty" name:"DeployAgent"`

	// 节点调度策略。若不指定改参数，则默认不使用节点调度策略。
	SchedulingStrategy *SchedulingStrategy `json:"SchedulingStrategy,omitempty" name:"SchedulingStrategy"`

	// 是否进行增量部署，默认为false，全量更新
	IncrementalDeployment *bool `json:"IncrementalDeployment,omitempty" name:"IncrementalDeployment"`

	// tcr或者不填
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// 数据卷信息-废弃，请用VolumeInfoList参数
	VolumeInfos *VolumeInfo `json:"VolumeInfos,omitempty" name:"VolumeInfos"`

	// 数据卷挂载点信息-废弃，请用VolumeMountInfoList参数
	VolumeMountInfos *VolumeMountInfo `json:"VolumeMountInfos,omitempty" name:"VolumeMountInfos"`

	// 数据卷信息，list
	VolumeInfoList []*VolumeInfo `json:"VolumeInfoList,omitempty" name:"VolumeInfoList"`

	// 数据卷挂载点信息，list
	VolumeMountInfoList []*VolumeMountInfo `json:"VolumeMountInfoList,omitempty" name:"VolumeMountInfoList"`

	// 是否清除数据卷信息，默认false
	VolumeClean *bool `json:"VolumeClean,omitempty" name:"VolumeClean"`

	// javaagent信息: SERVICE_AGENT/OT_AGENT
	AgentProfileList []*AgentProfile `json:"AgentProfileList,omitempty" name:"AgentProfileList"`

	// 预热配置信息
	WarmupSetting *WarmupSetting `json:"WarmupSetting,omitempty" name:"WarmupSetting"`
}

func (r *DeployContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployContainerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "TagName")
	delete(f, "InstanceNum")
	delete(f, "Server")
	delete(f, "Reponame")
	delete(f, "CpuLimit")
	delete(f, "MemLimit")
	delete(f, "JvmOpts")
	delete(f, "CpuRequest")
	delete(f, "MemRequest")
	delete(f, "DoNotStart")
	delete(f, "RepoName")
	delete(f, "UpdateType")
	delete(f, "UpdateIvl")
	delete(f, "AgentCpuRequest")
	delete(f, "AgentCpuLimit")
	delete(f, "AgentMemRequest")
	delete(f, "AgentMemLimit")
	delete(f, "IstioCpuRequest")
	delete(f, "IstioCpuLimit")
	delete(f, "IstioMemRequest")
	delete(f, "IstioMemLimit")
	delete(f, "MaxSurge")
	delete(f, "MaxUnavailable")
	delete(f, "HealthCheckSettings")
	delete(f, "Envs")
	delete(f, "ServiceSetting")
	delete(f, "DeployAgent")
	delete(f, "SchedulingStrategy")
	delete(f, "IncrementalDeployment")
	delete(f, "RepoType")
	delete(f, "VolumeInfos")
	delete(f, "VolumeMountInfos")
	delete(f, "VolumeInfoList")
	delete(f, "VolumeMountInfoList")
	delete(f, "VolumeClean")
	delete(f, "AgentProfileList")
	delete(f, "WarmupSetting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployContainerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployContainerGroupResponseParams struct {
	// 部署容器应用是否成功。
	// true：成功。
	// false：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeployContainerGroupResponseParams `json:"Response"`
}

func (r *DeployContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployContainerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 部署组启动参数
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 部署应用描述信息
	DeployDesc *string `json:"DeployDesc,omitempty" name:"DeployDesc"`

	// 是否允许强制启动
	ForceStart *bool `json:"ForceStart,omitempty" name:"ForceStart"`

	// 是否开启健康检查
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`

	// 开启健康检查时，配置健康检查
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`

	// 部署方式，0表示快速更新，1表示滚动更新
	UpdateType *uint64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 是否启用beta批次
	DeployBetaEnable *bool `json:"DeployBetaEnable,omitempty" name:"DeployBetaEnable"`

	// 滚动发布每个批次参与的实例比率
	DeployBatch []*float64 `json:"DeployBatch,omitempty" name:"DeployBatch"`

	// 滚动发布的执行方式
	DeployExeMode *string `json:"DeployExeMode,omitempty" name:"DeployExeMode"`

	// 滚动发布每个批次的时间间隔
	DeployWaitTime *uint64 `json:"DeployWaitTime,omitempty" name:"DeployWaitTime"`

	// 启动脚本 base64编码
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`

	// 停止脚本 base64编码
	StopScript *string `json:"StopScript,omitempty" name:"StopScript"`

	// 是否进行增量部署，默认为false，全量更新
	IncrementalDeployment *bool `json:"IncrementalDeployment,omitempty" name:"IncrementalDeployment"`

	// JDK名称: konaJDK或openJDK
	JdkName *string `json:"JdkName,omitempty" name:"JdkName"`

	// JDK版本: 8或11 (openJDK只支持8)
	JdkVersion *string `json:"JdkVersion,omitempty" name:"JdkVersion"`

	// 部署agent的类型、版本
	AgentProfileList []*AgentProfile `json:"AgentProfileList,omitempty" name:"AgentProfileList"`

	// 预热参数配置
	WarmupSetting *WarmupSetting `json:"WarmupSetting,omitempty" name:"WarmupSetting"`
}

type DeployGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 部署组启动参数
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 部署应用描述信息
	DeployDesc *string `json:"DeployDesc,omitempty" name:"DeployDesc"`

	// 是否允许强制启动
	ForceStart *bool `json:"ForceStart,omitempty" name:"ForceStart"`

	// 是否开启健康检查
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`

	// 开启健康检查时，配置健康检查
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`

	// 部署方式，0表示快速更新，1表示滚动更新
	UpdateType *uint64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 是否启用beta批次
	DeployBetaEnable *bool `json:"DeployBetaEnable,omitempty" name:"DeployBetaEnable"`

	// 滚动发布每个批次参与的实例比率
	DeployBatch []*float64 `json:"DeployBatch,omitempty" name:"DeployBatch"`

	// 滚动发布的执行方式
	DeployExeMode *string `json:"DeployExeMode,omitempty" name:"DeployExeMode"`

	// 滚动发布每个批次的时间间隔
	DeployWaitTime *uint64 `json:"DeployWaitTime,omitempty" name:"DeployWaitTime"`

	// 启动脚本 base64编码
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`

	// 停止脚本 base64编码
	StopScript *string `json:"StopScript,omitempty" name:"StopScript"`

	// 是否进行增量部署，默认为false，全量更新
	IncrementalDeployment *bool `json:"IncrementalDeployment,omitempty" name:"IncrementalDeployment"`

	// JDK名称: konaJDK或openJDK
	JdkName *string `json:"JdkName,omitempty" name:"JdkName"`

	// JDK版本: 8或11 (openJDK只支持8)
	JdkVersion *string `json:"JdkVersion,omitempty" name:"JdkVersion"`

	// 部署agent的类型、版本
	AgentProfileList []*AgentProfile `json:"AgentProfileList,omitempty" name:"AgentProfileList"`

	// 预热参数配置
	WarmupSetting *WarmupSetting `json:"WarmupSetting,omitempty" name:"WarmupSetting"`
}

func (r *DeployGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "PkgId")
	delete(f, "StartupParameters")
	delete(f, "DeployDesc")
	delete(f, "ForceStart")
	delete(f, "EnableHealthCheck")
	delete(f, "HealthCheckSettings")
	delete(f, "UpdateType")
	delete(f, "DeployBetaEnable")
	delete(f, "DeployBatch")
	delete(f, "DeployExeMode")
	delete(f, "DeployWaitTime")
	delete(f, "StartScript")
	delete(f, "StopScript")
	delete(f, "IncrementalDeployment")
	delete(f, "JdkName")
	delete(f, "JdkVersion")
	delete(f, "AgentProfileList")
	delete(f, "WarmupSetting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployGroupResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TaskId `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeployGroupResponseParams `json:"Response"`
}

func (r *DeployGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiDetailRequestParams struct {
	// 微服务id
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 包版本
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeApiDetailRequest struct {
	*tchttp.BaseRequest
	
	// 微服务id
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 包版本
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApiDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MicroserviceId")
	delete(f, "Path")
	delete(f, "Method")
	delete(f, "PkgVersion")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiDetailResponseParams struct {
	// API 详情
	Result *ApiDetailResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiDetailResponseParams `json:"Response"`
}

func (r *DescribeApiDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiGroupRequestParams struct {
	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeApiGroupRequest struct {
	*tchttp.BaseRequest
	
	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiGroupResponseParams struct {
	// API分组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiGroupInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiGroupResponseParams `json:"Response"`
}

func (r *DescribeApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiGroupsRequestParams struct {
	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分组类型。 ms： 微服务分组； external:外部Api分组
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// 鉴权类型。 secret： 密钥鉴权； none:无鉴权
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 发布状态, drafted: 未发布。 released: 发布
	Status *string `json:"Status,omitempty" name:"Status"`

	// 排序字段："created_time"或"group_context"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型：0(ASC)或1(DESC)
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`
}

type DescribeApiGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分组类型。 ms： 微服务分组； external:外部Api分组
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// 鉴权类型。 secret： 密钥鉴权； none:无鉴权
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 发布状态, drafted: 未发布。 released: 发布
	Status *string `json:"Status,omitempty" name:"Status"`

	// 排序字段："created_time"或"group_context"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型：0(ASC)或1(DESC)
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`
}

func (r *DescribeApiGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupType")
	delete(f, "AuthType")
	delete(f, "Status")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "GatewayInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiGroupsResponseParams struct {
	// 翻页结构体
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageApiGroupInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiGroupsResponseParams `json:"Response"`
}

func (r *DescribeApiGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiRateLimitRulesRequestParams struct {
	// Api ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type DescribeApiRateLimitRulesRequest struct {
	*tchttp.BaseRequest
	
	// Api ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DescribeApiRateLimitRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiRateLimitRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiRateLimitRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiRateLimitRulesResponseParams struct {
	// 限流结果
	Result []*ApiRateLimitRule `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiRateLimitRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiRateLimitRulesResponseParams `json:"Response"`
}

func (r *DescribeApiRateLimitRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiRateLimitRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiUseDetailRequestParams struct {
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组Api ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeApiUseDetailRequest struct {
	*tchttp.BaseRequest
	
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组Api ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeApiUseDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiUseDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayDeployGroupId")
	delete(f, "ApiId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiUseDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiUseDetailResponseParams struct {
	// 日使用统计对象
	Result *GroupApiUseStatistics `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiUseDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiUseDetailResponseParams `json:"Response"`
}

func (r *DescribeApiUseDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiUseDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiVersionsRequestParams struct {
	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// API 请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`
}

type DescribeApiVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// API 请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`
}

func (r *DescribeApiVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MicroserviceId")
	delete(f, "Path")
	delete(f, "Method")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiVersionsResponseParams struct {
	// API版本列表
	Result []*ApiVersionArray `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiVersionsResponseParams `json:"Response"`
}

func (r *DescribeApiVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationAttributeRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeApplicationAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationAttributeResponseParams struct {
	// 应用列表其它字段返回参数
	Result *ApplicationAttribute `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationAttributeResponseParams `json:"Response"`
}

func (r *DescribeApplicationAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationBusinessLogConfigRequestParams struct {
	// TSF应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeApplicationBusinessLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// TSF应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationBusinessLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationBusinessLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationBusinessLogConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationBusinessLogConfigResponseParams `json:"Response"`
}

func (r *DescribeApplicationBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationBusinessLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationResponseParams struct {
	// 应用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApplicationForPage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationResponseParams `json:"Response"`
}

func (r *DescribeApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsRequestParams struct {
	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 应用类型
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用的微服务类型
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用资源类型数组
	ApplicationResourceTypeList []*string `json:"ApplicationResourceTypeList,omitempty" name:"ApplicationResourceTypeList"`

	// IdList
	ApplicationIdList []*string `json:"ApplicationIdList,omitempty" name:"ApplicationIdList"`
}

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 应用类型
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用的微服务类型
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用资源类型数组
	ApplicationResourceTypeList []*string `json:"ApplicationResourceTypeList,omitempty" name:"ApplicationResourceTypeList"`

	// IdList
	ApplicationIdList []*string `json:"ApplicationIdList,omitempty" name:"ApplicationIdList"`
}

func (r *DescribeApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ApplicationType")
	delete(f, "MicroserviceType")
	delete(f, "ApplicationResourceTypeList")
	delete(f, "ApplicationIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsResponseParams struct {
	// 应用分页列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageApplication `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationsResponseParams `json:"Response"`
}

func (r *DescribeApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicResourceUsageRequestParams struct {
	// 是否无视权限查询全租户的，默认 true。注：无论 true 还是 false，PackageSpaceUsed 和 ConsulInstanceCount  都是全租户的
	All *bool `json:"All,omitempty" name:"All"`
}

type DescribeBasicResourceUsageRequest struct {
	*tchttp.BaseRequest
	
	// 是否无视权限查询全租户的，默认 true。注：无论 true 还是 false，PackageSpaceUsed 和 ConsulInstanceCount  都是全租户的
	All *bool `json:"All,omitempty" name:"All"`
}

func (r *DescribeBasicResourceUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicResourceUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "All")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBasicResourceUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicResourceUsageResponseParams struct {
	// TSF基本资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *OverviewBasicResourceUsage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBasicResourceUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBasicResourceUsageResponseParams `json:"Response"`
}

func (r *DescribeBasicResourceUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicResourceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBusinessLogConfigRequestParams struct {
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

type DescribeBusinessLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBusinessLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBusinessLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBusinessLogConfigResponseParams struct {
	// 日志配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *BusinessLogConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBusinessLogConfigResponseParams `json:"Response"`
}

func (r *DescribeBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBusinessLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBusinessLogConfigsRequestParams struct {
	// 偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量，取值范围[1, 50]，默认值为10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊匹配关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`

	// 无
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`
}

type DescribeBusinessLogConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量，取值范围[1, 50]，默认值为10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊匹配关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`

	// 无
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`
}

func (r *DescribeBusinessLogConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBusinessLogConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "DisableProgramAuthCheck")
	delete(f, "ConfigIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBusinessLogConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBusinessLogConfigsResponseParams struct {
	// 业务日志配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageBusinessLogConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBusinessLogConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBusinessLogConfigsResponseParams `json:"Response"`
}

func (r *DescribeBusinessLogConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBusinessLogConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstancesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SearchWord")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstancesResponseParams struct {
	// 集群机器实例分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageInstance `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInstancesResponseParams `json:"Response"`
}

func (r *DescribeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 搜索词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// idList
	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 搜索词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// idList
	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterType")
	delete(f, "ClusterIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// Cluster分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageClusterV2 `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigReleaseLogsRequestParams struct {
	// 部署组ID，不传入时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeConfigReleaseLogsRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID，不传入时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeConfigReleaseLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigReleaseLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NamespaceId")
	delete(f, "ClusterId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigReleaseLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigReleaseLogsResponseParams struct {
	// 分页的配置项发布历史列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageConfigReleaseLog `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigReleaseLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigReleaseLogsResponseParams `json:"Response"`
}

func (r *DescribeConfigReleaseLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigReleaseLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigReleasesRequestParams struct {
	// 配置项名称，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 部署组ID，不传入时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 配置ID，不传入时查询全量
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeConfigReleasesRequest struct {
	*tchttp.BaseRequest
	
	// 配置项名称，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 部署组ID，不传入时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 配置ID，不传入时查询全量
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeConfigReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigReleasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigName")
	delete(f, "GroupId")
	delete(f, "NamespaceId")
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ConfigId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigReleasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigReleasesResponseParams struct {
	// 分页的配置发布信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageConfigRelease `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigReleasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigReleasesResponseParams `json:"Response"`
}

func (r *DescribeConfigReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigReleasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRequestParams struct {
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigResponseParams struct {
	// 配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *Config `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigResponseParams `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigSummaryRequestParams struct {
	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 查询关键字，模糊查询：应用名称，配置项名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 按时间排序：creation_time；按名称排序：config_name
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序传 0，降序传 1
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 无
	ConfigTagList []*string `json:"ConfigTagList,omitempty" name:"ConfigTagList"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`

	// 无
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`
}

type DescribeConfigSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 查询关键字，模糊查询：应用名称，配置项名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 按时间排序：creation_time；按名称排序：config_name
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序传 0，降序传 1
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 无
	ConfigTagList []*string `json:"ConfigTagList,omitempty" name:"ConfigTagList"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`

	// 无
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`
}

func (r *DescribeConfigSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "ConfigTagList")
	delete(f, "DisableProgramAuthCheck")
	delete(f, "ConfigIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigSummaryResponseParams struct {
	// 配置项分页对象
	Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigSummaryResponseParams `json:"Response"`
}

func (r *DescribeConfigSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigTemplateRequestParams struct {
	// 无
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
}

type DescribeConfigTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 无
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`
}

func (r *DescribeConfigTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigTemplateResponseParams struct {
	// Result
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ConfigTemplate `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigTemplateResponseParams `json:"Response"`
}

func (r *DescribeConfigTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigsRequestParams struct {
	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 配置项ID，不传入时查询全量，高优先级
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项ID列表，不传入时查询全量，低优先级
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`

	// 配置项名称，精确查询，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本，精确查询，不传入时查询全量
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 配置项ID，不传入时查询全量，高优先级
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项ID列表，不传入时查询全量，低优先级
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`

	// 配置项名称，精确查询，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本，精确查询，不传入时查询全量
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ConfigId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ConfigIdList")
	delete(f, "ConfigName")
	delete(f, "ConfigVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigsResponseParams struct {
	// 分页后的配置项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigsResponseParams `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerEventsRequestParams struct {
	// event 的资源类型, group 或者 instance
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// event 的资源 id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 当类型是 instance 时需要
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeContainerEventsRequest struct {
	*tchttp.BaseRequest
	
	// event 的资源类型, group 或者 instance
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// event 的资源 id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 当类型是 instance 时需要
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeContainerEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerEventsResponseParams struct {
	// events 分页列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageContainerEvent `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeContainerEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerEventsResponseParams `json:"Response"`
}

func (r *DescribeContainerEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerGroupDeployInfoRequestParams struct {
	// 实例所属 groupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeContainerGroupDeployInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例所属 groupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeContainerGroupDeployInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerGroupDeployInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerGroupDeployInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerGroupDeployInfoResponseParams struct {
	// 获取部署组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ContainerGroupDeploy `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeContainerGroupDeployInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerGroupDeployInfoResponseParams `json:"Response"`
}

func (r *DescribeContainerGroupDeployInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerGroupDeployInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerGroupDetailRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeContainerGroupDetailRequest struct {
	*tchttp.BaseRequest
	
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeContainerGroupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerGroupDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerGroupDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerGroupDetailResponseParams struct {
	// 容器部署组详情
	Result *ContainerGroupDetail `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeContainerGroupDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerGroupDetailResponseParams `json:"Response"`
}

func (r *DescribeContainerGroupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerGroupDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerGroupsRequestParams struct {
	// 搜索字段，模糊搜索groupName字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 分组所属应用ID。必填
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段，默认为 createTime字段，支持id， name， createTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间 ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

type DescribeContainerGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索字段，模糊搜索groupName字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 分组所属应用ID。必填
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段，默认为 createTime字段，支持id， name， createTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间 ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DescribeContainerGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "ApplicationId")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContainerGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContainerGroupsResponseParams struct {
	// 查询的权限数据对象
	Result *ContainGroupResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeContainerGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContainerGroupsResponseParams `json:"Response"`
}

func (r *DescribeContainerGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContainerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCreateGatewayApiStatusRequestParams struct {
	// 请求方法
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

type DescribeCreateGatewayApiStatusRequest struct {
	*tchttp.BaseRequest
	
	// 请求方法
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *DescribeCreateGatewayApiStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCreateGatewayApiStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "MicroserviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCreateGatewayApiStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCreateGatewayApiStatusResponseParams struct {
	// 是否已完成导入任务
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCreateGatewayApiStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCreateGatewayApiStatusResponseParams `json:"Response"`
}

func (r *DescribeCreateGatewayApiStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCreateGatewayApiStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeliveryConfigByGroupIdRequestParams struct {
	// 部署组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeDeliveryConfigByGroupIdRequest struct {
	*tchttp.BaseRequest
	
	// 部署组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeDeliveryConfigByGroupIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliveryConfigByGroupIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeliveryConfigByGroupIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeliveryConfigByGroupIdResponseParams struct {
	// 投递kafak配置项
	Result *SimpleKafkaDeliveryConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeliveryConfigByGroupIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeliveryConfigByGroupIdResponseParams `json:"Response"`
}

func (r *DescribeDeliveryConfigByGroupIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliveryConfigByGroupIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeliveryConfigRequestParams struct {
	// 投递配置id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

type DescribeDeliveryConfigRequest struct {
	*tchttp.BaseRequest
	
	// 投递配置id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeDeliveryConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliveryConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeliveryConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeliveryConfigResponseParams struct {
	// 投递kafka配置
	Result *KafkaDeliveryConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeliveryConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeliveryConfigResponseParams `json:"Response"`
}

func (r *DescribeDeliveryConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliveryConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeliveryConfigsRequestParams struct {
	// 关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 搜索条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDeliveryConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 搜索条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDeliveryConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliveryConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeliveryConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeliveryConfigsResponseParams struct {
	// 投递项关联部署组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DeliveryConfigBindGroups `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeliveryConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeliveryConfigsResponseParams `json:"Response"`
}

func (r *DescribeDeliveryConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeliveryConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDownloadInfoRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 程序包仓库ID
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`
}

type DescribeDownloadInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 程序包仓库ID
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`
}

func (r *DescribeDownloadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDownloadInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PkgId")
	delete(f, "RepositoryId")
	delete(f, "RepositoryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDownloadInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDownloadInfoResponseParams struct {
	// COS鉴权信息
	Result *CosDownloadInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDownloadInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDownloadInfoResponseParams `json:"Response"`
}

func (r *DescribeDownloadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDownloadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnabledUnitRuleRequestParams struct {
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`
}

type DescribeEnabledUnitRuleRequest struct {
	*tchttp.BaseRequest
	
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`
}

func (r *DescribeEnabledUnitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnabledUnitRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnabledUnitRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnabledUnitRuleResponseParams struct {
	// 单元化规则对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UnitRule `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnabledUnitRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnabledUnitRuleResponseParams `json:"Response"`
}

func (r *DescribeEnabledUnitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnabledUnitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileConfigReleasesRequestParams struct {
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeFileConfigReleasesRequest struct {
	*tchttp.BaseRequest
	
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFileConfigReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileConfigReleasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "ConfigName")
	delete(f, "GroupId")
	delete(f, "NamespaceId")
	delete(f, "ClusterId")
	delete(f, "ApplicationId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileConfigReleasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileConfigReleasesResponseParams struct {
	// 配置项发布信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageFileConfigRelease `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileConfigReleasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileConfigReleasesResponseParams `json:"Response"`
}

func (r *DescribeFileConfigReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileConfigReleasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileConfigsRequestParams struct {
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项ID列表
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

type DescribeFileConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项ID列表
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *DescribeFileConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "ConfigIdList")
	delete(f, "ConfigName")
	delete(f, "ApplicationId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ConfigVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileConfigsResponseParams struct {
	// 文件配置项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageFileConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileConfigsResponseParams `json:"Response"`
}

func (r *DescribeFileConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLastBatchStateRequestParams struct {
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type DescribeFlowLastBatchStateRequest struct {
	*tchttp.BaseRequest
	
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowLastBatchStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLastBatchStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowLastBatchStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLastBatchStateResponseParams struct {
	// 工作流批次最新状态
	Result *TaskFlowLastBatchState `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowLastBatchStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowLastBatchStateResponseParams `json:"Response"`
}

func (r *DescribeFlowLastBatchStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLastBatchStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayAllGroupApisRequestParams struct {
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 搜索关键字，支持分组名称或API Path
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

type DescribeGatewayAllGroupApisRequest struct {
	*tchttp.BaseRequest
	
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 搜索关键字，支持分组名称或API Path
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeGatewayAllGroupApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayAllGroupApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayDeployGroupId")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayAllGroupApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayAllGroupApisResponseParams struct {
	// 网关分组和API列表信息
	Result *GatewayVo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGatewayAllGroupApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayAllGroupApisResponseParams `json:"Response"`
}

func (r *DescribeGatewayAllGroupApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayAllGroupApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayApisRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页的记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字，支持 API path
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`
}

type DescribeGatewayApisRequest struct {
	*tchttp.BaseRequest
	
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页的记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字，支持 API path
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`
}

func (r *DescribeGatewayApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "GatewayDeployGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayApisResponseParams struct {
	// 翻页结构
	Result *TsfPageApiDetailInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGatewayApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayApisResponseParams `json:"Response"`
}

func (r *DescribeGatewayApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayMonitorOverviewRequestParams struct {
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`
}

type DescribeGatewayMonitorOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`
}

func (r *DescribeGatewayMonitorOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayMonitorOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayDeployGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayMonitorOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayMonitorOverviewResponseParams struct {
	// 监控概览对象
	Result *MonitorOverview `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGatewayMonitorOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayMonitorOverviewResponseParams `json:"Response"`
}

func (r *DescribeGatewayMonitorOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayMonitorOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupAttributeRequestParams struct {
	// 部署组ID字段
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID字段
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupAttributeResponseParams struct {
	// 虚拟机部署组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *VmGroupOther `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupAttributeResponseParams `json:"Response"`
}

func (r *DescribeGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupBindedGatewaysRequestParams struct {
	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

type DescribeGroupBindedGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// API 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeGroupBindedGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupBindedGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupBindedGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupBindedGatewaysResponseParams struct {
	// 翻页结构体
	Result *TsfPageGatewayDeployGroup `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupBindedGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupBindedGatewaysResponseParams `json:"Response"`
}

func (r *DescribeGroupBindedGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupBindedGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupBusinessLogConfigsRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeGroupBusinessLogConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupBusinessLogConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupBusinessLogConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupBusinessLogConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupBusinessLogConfigsResponseParams struct {
	// 业务日志配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageBusinessLogConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupBusinessLogConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupBusinessLogConfigsResponseParams `json:"Response"`
}

func (r *DescribeGroupBusinessLogConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupBusinessLogConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupGatewaysRequestParams struct {
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

type DescribeGroupGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeGroupGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayDeployGroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupGatewaysResponseParams struct {
	// API分组信息
	Result *TsfPageApiGroupInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupGatewaysResponseParams `json:"Response"`
}

func (r *DescribeGroupGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupInstancesRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "SearchWord")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupInstancesResponseParams struct {
	// 部署组机器信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageInstance `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupInstancesResponseParams `json:"Response"`
}

func (r *DescribeGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupReleaseRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeGroupReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupReleaseResponseParams struct {
	// 部署组发布的相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *GroupRelease `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupReleaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupReleaseResponseParams `json:"Response"`
}

func (r *DescribeGroupReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupResponseParams struct {
	// 虚拟机部署组详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *VmGroup `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupResponseParams `json:"Response"`
}

func (r *DescribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupUseDetailRequestParams struct {
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定top的条数,默认为10
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type DescribeGroupUseDetailRequest struct {
	*tchttp.BaseRequest
	
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定top的条数,默认为10
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

func (r *DescribeGroupUseDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupUseDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayDeployGroupId")
	delete(f, "GroupId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Count")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupUseDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupUseDetailResponseParams struct {
	// 日使用统计对象
	Result *GroupDailyUseStatistics `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupUseDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupUseDetailResponseParams `json:"Response"`
}

func (r *DescribeGroupUseDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupUseDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupsRequestParams struct {
	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组资源类型列表
	GroupResourceTypeList []*string `json:"GroupResourceTypeList,omitempty" name:"GroupResourceTypeList"`

	// 部署组状态过滤字段
	Status *string `json:"Status,omitempty" name:"Status"`

	// 无
	GroupIdList []*string `json:"GroupIdList,omitempty" name:"GroupIdList"`
}

type DescribeGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组资源类型列表
	GroupResourceTypeList []*string `json:"GroupResourceTypeList,omitempty" name:"GroupResourceTypeList"`

	// 部署组状态过滤字段
	Status *string `json:"Status,omitempty" name:"Status"`

	// 无
	GroupIdList []*string `json:"GroupIdList,omitempty" name:"GroupIdList"`
}

func (r *DescribeGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "ApplicationId")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NamespaceId")
	delete(f, "ClusterId")
	delete(f, "GroupResourceTypeList")
	delete(f, "Status")
	delete(f, "GroupIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupsResponseParams struct {
	// 虚拟机部署组分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageVmGroup `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupsResponseParams `json:"Response"`
}

func (r *DescribeGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupsWithPluginRequestParams struct {
	// 插件ID
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 绑定/未绑定: true / false
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`
}

type DescribeGroupsWithPluginRequest struct {
	*tchttp.BaseRequest
	
	// 插件ID
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 绑定/未绑定: true / false
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`
}

func (r *DescribeGroupsWithPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsWithPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "Bound")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "GatewayInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGroupsWithPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGroupsWithPluginResponseParams struct {
	// API分组信息列表
	Result *TsfPageApiGroupInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGroupsWithPluginResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGroupsWithPluginResponseParams `json:"Response"`
}

func (r *DescribeGroupsWithPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGroupsWithPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRepositoryRequestParams struct {
	// 仓库名，搜索关键字,不带命名空间的
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 企业: tcr ；个人: personal或者不填
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// TcrRepoInfo值
	TcrRepoInfo *TcrRepoInfo `json:"TcrRepoInfo,omitempty" name:"TcrRepoInfo"`
}

type DescribeImageRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名，搜索关键字,不带命名空间的
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 企业: tcr ；个人: personal或者不填
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// TcrRepoInfo值
	TcrRepoInfo *TcrRepoInfo `json:"TcrRepoInfo,omitempty" name:"TcrRepoInfo"`
}

func (r *DescribeImageRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RepoType")
	delete(f, "ApplicationId")
	delete(f, "TcrRepoInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRepositoryResponseParams struct {
	// 查询的权限数据对象
	Result *ImageRepositoryResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageRepositoryResponseParams `json:"Response"`
}

func (r *DescribeImageRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageTagsRequestParams struct {
	// 应用Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 不填和0:查询 1:不查询
	QueryImageIdFlag *int64 `json:"QueryImageIdFlag,omitempty" name:"QueryImageIdFlag"`

	// 可用于搜索的 tag 名字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 企业: tcr ；个人: personal或者不填
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// TcrRepoInfo值
	TcrRepoInfo *TcrRepoInfo `json:"TcrRepoInfo,omitempty" name:"TcrRepoInfo"`
}

type DescribeImageTagsRequest struct {
	*tchttp.BaseRequest
	
	// 应用Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 不填和0:查询 1:不查询
	QueryImageIdFlag *int64 `json:"QueryImageIdFlag,omitempty" name:"QueryImageIdFlag"`

	// 可用于搜索的 tag 名字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 企业: tcr ；个人: personal或者不填
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// TcrRepoInfo值
	TcrRepoInfo *TcrRepoInfo `json:"TcrRepoInfo,omitempty" name:"TcrRepoInfo"`
}

func (r *DescribeImageTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "QueryImageIdFlag")
	delete(f, "SearchWord")
	delete(f, "RepoType")
	delete(f, "TcrRepoInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageTagsResponseParams struct {
	// 查询的权限数据对象
	Result *ImageTagsResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageTagsResponseParams `json:"Response"`
}

func (r *DescribeImageTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInovcationIndicatorsRequestParams struct {
	// 维度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 调用方服务名
	CallerServiceName *string `json:"CallerServiceName,omitempty" name:"CallerServiceName"`

	// 被调方服务名
	CalleeServiceName *string `json:"CalleeServiceName,omitempty" name:"CalleeServiceName"`

	// 调用方接口名
	CallerInterfaceName *string `json:"CallerInterfaceName,omitempty" name:"CallerInterfaceName"`

	// 被调方接口名
	CalleeInterfaceName *string `json:"CalleeInterfaceName,omitempty" name:"CalleeInterfaceName"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInovcationIndicatorsRequest struct {
	*tchttp.BaseRequest
	
	// 维度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 调用方服务名
	CallerServiceName *string `json:"CallerServiceName,omitempty" name:"CallerServiceName"`

	// 被调方服务名
	CalleeServiceName *string `json:"CalleeServiceName,omitempty" name:"CalleeServiceName"`

	// 调用方接口名
	CallerInterfaceName *string `json:"CallerInterfaceName,omitempty" name:"CallerInterfaceName"`

	// 被调方接口名
	CalleeInterfaceName *string `json:"CalleeInterfaceName,omitempty" name:"CalleeInterfaceName"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInovcationIndicatorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInovcationIndicatorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Dimension")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "NamespaceId")
	delete(f, "ServiceId")
	delete(f, "CallerServiceName")
	delete(f, "CalleeServiceName")
	delete(f, "CallerInterfaceName")
	delete(f, "CalleeInterfaceName")
	delete(f, "ApplicationId")
	delete(f, "GroupId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInovcationIndicatorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInovcationIndicatorsResponseParams struct {
	// 服务调用监控指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *InvocationIndicator `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInovcationIndicatorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInovcationIndicatorsResponseParams `json:"Response"`
}

func (r *DescribeInovcationIndicatorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInovcationIndicatorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 机器列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *InstanceEnrichedInfoPage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInvocationMetricDataCurveRequestParams struct {
	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询时间粒度，单位秒可选值：60、3600、86400
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 查询指标维度
	MetricDimensions []*MetricDimension `json:"MetricDimensions,omitempty" name:"MetricDimensions"`

	// 查询指标名
	Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`

	// 视图视角。可选值：SERVER, CLIENT。默认为SERVER
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 类型。组件监控使用，可选值：SQL 或者 NoSQL
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeInvocationMetricDataCurveRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询时间粒度，单位秒可选值：60、3600、86400
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 查询指标维度
	MetricDimensions []*MetricDimension `json:"MetricDimensions,omitempty" name:"MetricDimensions"`

	// 查询指标名
	Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`

	// 视图视角。可选值：SERVER, CLIENT。默认为SERVER
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 类型。组件监控使用，可选值：SQL 或者 NoSQL
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeInvocationMetricDataCurveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationMetricDataCurveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	delete(f, "MetricDimensions")
	delete(f, "Metrics")
	delete(f, "Kind")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationMetricDataCurveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationMetricDataCurveResponseParams struct {
	// 指标监控数据曲线集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*MetricDataCurve `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInvocationMetricDataCurveResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationMetricDataCurveResponseParams `json:"Response"`
}

func (r *DescribeInvocationMetricDataCurveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationMetricDataCurveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationMetricDataDimensionRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始index
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 聚合维度
	DimensionName *string `json:"DimensionName,omitempty" name:"DimensionName"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 维度
	MetricDimensionValues []*MetricDimensionValue `json:"MetricDimensionValues,omitempty" name:"MetricDimensionValues"`
}

type DescribeInvocationMetricDataDimensionRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始index
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 聚合维度
	DimensionName *string `json:"DimensionName,omitempty" name:"DimensionName"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 维度
	MetricDimensionValues []*MetricDimensionValue `json:"MetricDimensionValues,omitempty" name:"MetricDimensionValues"`
}

func (r *DescribeInvocationMetricDataDimensionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationMetricDataDimensionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DimensionName")
	delete(f, "SearchWord")
	delete(f, "MetricDimensionValues")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationMetricDataDimensionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationMetricDataDimensionResponseParams struct {
	// 维度
	Result *TsfPageDimension `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInvocationMetricDataDimensionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationMetricDataDimensionResponseParams `json:"Response"`
}

func (r *DescribeInvocationMetricDataDimensionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationMetricDataDimensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationMetricDataPointRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 维度
	MetricDimensionValues []*MetricDimensionValue `json:"MetricDimensionValues,omitempty" name:"MetricDimensionValues"`

	// 指标
	Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`

	// 调用视角。可选值：SERVER, CLIENT。默认为SERVER
	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

type DescribeInvocationMetricDataPointRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 维度
	MetricDimensionValues []*MetricDimensionValue `json:"MetricDimensionValues,omitempty" name:"MetricDimensionValues"`

	// 指标
	Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`

	// 调用视角。可选值：SERVER, CLIENT。默认为SERVER
	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

func (r *DescribeInvocationMetricDataPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationMetricDataPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricDimensionValues")
	delete(f, "Metrics")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationMetricDataPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationMetricDataPointResponseParams struct {
	// 单值指标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*MetricDataSingleValue `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInvocationMetricDataPointResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationMetricDataPointResponseParams `json:"Response"`
}

func (r *DescribeInvocationMetricDataPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationMetricDataPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationMetricScatterPlotRequestParams struct {
	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询时间粒度，单位秒。可选值：60、3600、86400。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 查询指标维度
	MetricDimensions []*MetricDimension `json:"MetricDimensions,omitempty" name:"MetricDimensions"`

	// 查询指标名
	Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`

	// 视图视角。可选值：SERVER, CLIENT。默认为SERVER
	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

type DescribeInvocationMetricScatterPlotRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询时间粒度，单位秒。可选值：60、3600、86400。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 查询指标维度
	MetricDimensions []*MetricDimension `json:"MetricDimensions,omitempty" name:"MetricDimensions"`

	// 查询指标名
	Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`

	// 视图视角。可选值：SERVER, CLIENT。默认为SERVER
	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

func (r *DescribeInvocationMetricScatterPlotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationMetricScatterPlotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	delete(f, "MetricDimensions")
	delete(f, "Metrics")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationMetricScatterPlotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationMetricScatterPlotResponseParams struct {
	// 多值时间抽统计指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *InvocationMetricScatterPlot `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInvocationMetricScatterPlotResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationMetricScatterPlotResponseParams `json:"Response"`
}

func (r *DescribeInvocationMetricScatterPlotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationMetricScatterPlotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJvmMonitorRequestParams struct {
	// 查询的实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例所属应用Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 时间粒度,单位:秒
	TimeGranularity *int64 `json:"TimeGranularity,omitempty" name:"TimeGranularity"`

	// 查询数据起始时间格式(yyyy-MM-dd HH:mm:ss)
	From *string `json:"From,omitempty" name:"From"`

	// 查询数据结束时间格式(yyyy-MM-dd HH:mm:ss)
	To *string `json:"To,omitempty" name:"To"`

	// 查询的监控图列表,以返回值属性名作为入参
	RequiredPictures []*string `json:"RequiredPictures,omitempty" name:"RequiredPictures"`

	// 扩展字段
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type DescribeJvmMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 查询的实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例所属应用Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 时间粒度,单位:秒
	TimeGranularity *int64 `json:"TimeGranularity,omitempty" name:"TimeGranularity"`

	// 查询数据起始时间格式(yyyy-MM-dd HH:mm:ss)
	From *string `json:"From,omitempty" name:"From"`

	// 查询数据结束时间格式(yyyy-MM-dd HH:mm:ss)
	To *string `json:"To,omitempty" name:"To"`

	// 查询的监控图列表,以返回值属性名作为入参
	RequiredPictures []*string `json:"RequiredPictures,omitempty" name:"RequiredPictures"`

	// 扩展字段
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeJvmMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJvmMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ApplicationId")
	delete(f, "TimeGranularity")
	delete(f, "From")
	delete(f, "To")
	delete(f, "RequiredPictures")
	delete(f, "Tag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJvmMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJvmMonitorResponseParams struct {
	// Java实例jvm监控数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *JvmMonitorData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeJvmMonitorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJvmMonitorResponseParams `json:"Response"`
}

func (r *DescribeJvmMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJvmMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaneRulesRequestParams struct {
	// 每页展示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 泳道规则ID（用于精确搜索）
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 无
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

type DescribeLaneRulesRequest struct {
	*tchttp.BaseRequest
	
	// 每页展示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 泳道规则ID（用于精确搜索）
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 无
	RuleIdList []*string `json:"RuleIdList,omitempty" name:"RuleIdList"`
}

func (r *DescribeLaneRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaneRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SearchWord")
	delete(f, "RuleId")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLaneRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaneRulesResponseParams struct {
	// 泳道规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *LaneRules `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLaneRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLaneRulesResponseParams `json:"Response"`
}

func (r *DescribeLaneRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaneRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLanesRequestParams struct {
	// 每页展示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 无
	LaneIdList []*string `json:"LaneIdList,omitempty" name:"LaneIdList"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`
}

type DescribeLanesRequest struct {
	*tchttp.BaseRequest
	
	// 每页展示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 无
	LaneIdList []*string `json:"LaneIdList,omitempty" name:"LaneIdList"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`
}

func (r *DescribeLanesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLanesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SearchWord")
	delete(f, "LaneIdList")
	delete(f, "DisableProgramAuthCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLanesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLanesResponseParams struct {
	// 泳道列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *LaneInfos `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLanesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLanesResponseParams `json:"Response"`
}

func (r *DescribeLanesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLanesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMicroserviceRequestParams struct {
	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 可选，根据部署组ID进行过滤
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 过滤条件。多个 filter 之间是与关系，单个 filter 多个 value 之间是或关系。filter name 取值有：id（实例id）、name（实例名）、lan-ip（内网ip）、node-ip（所在节点ip）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeMicroserviceRequest struct {
	*tchttp.BaseRequest
	
	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 可选，根据部署组ID进行过滤
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 过滤条件。多个 filter 之间是与关系，单个 filter 多个 value 之间是或关系。filter name 取值有：id（实例id）、name（实例名）、lan-ip（内网ip）、node-ip（所在节点ip）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMicroserviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MicroserviceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMicroserviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMicroserviceResponseParams struct {
	// 微服务详情实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageMsInstance `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMicroserviceResponseParams `json:"Response"`
}

func (r *DescribeMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMicroserviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMicroservicesRequestParams struct {
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 状态过滤，online、offline、single_online
	Status []*string `json:"Status,omitempty" name:"Status"`

	// IdList
	MicroserviceIdList []*string `json:"MicroserviceIdList,omitempty" name:"MicroserviceIdList"`

	// 搜索的服务名列表
	MicroserviceNameList []*string `json:"MicroserviceNameList,omitempty" name:"MicroserviceNameList"`
}

type DescribeMicroservicesRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 状态过滤，online、offline、single_online
	Status []*string `json:"Status,omitempty" name:"Status"`

	// IdList
	MicroserviceIdList []*string `json:"MicroserviceIdList,omitempty" name:"MicroserviceIdList"`

	// 搜索的服务名列表
	MicroserviceNameList []*string `json:"MicroserviceNameList,omitempty" name:"MicroserviceNameList"`
}

func (r *DescribeMicroservicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMicroservicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "SearchWord")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "MicroserviceIdList")
	delete(f, "MicroserviceNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMicroservicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMicroservicesResponseParams struct {
	// 微服务分页列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageMicroservice `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMicroservicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMicroservicesResponseParams `json:"Response"`
}

func (r *DescribeMicroservicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMicroservicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsApiListRequestParams struct {
	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 每页的数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeMsApiListRequest struct {
	*tchttp.BaseRequest
	
	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 每页的数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMsApiListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsApiListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MicroserviceId")
	delete(f, "SearchWord")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMsApiListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsApiListResponseParams struct {
	// 相应结果
	Result *TsfApiListResponse `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMsApiListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMsApiListResponseParams `json:"Response"`
}

func (r *DescribeMsApiListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsApiListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewInvocationRequestParams struct {
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 监控统计类型，可选值：SumReqAmount、AvgFailureRate、AvgTimeCost，分别对应请求量、请求错误率、平均响应耗时
	Type *string `json:"Type,omitempty" name:"Type"`

	// 监控统计数据粒度，可选值：60、3600、86400，分别对应1分钟、1小时、1天
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 查询开始时间，默认为当天的 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，默认为当前时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeOverviewInvocationRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 监控统计类型，可选值：SumReqAmount、AvgFailureRate、AvgTimeCost，分别对应请求量、请求错误率、平均响应耗时
	Type *string `json:"Type,omitempty" name:"Type"`

	// 监控统计数据粒度，可选值：60、3600、86400，分别对应1分钟、1小时、1天
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 查询开始时间，默认为当天的 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，默认为当前时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeOverviewInvocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewInvocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "Type")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewInvocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewInvocationResponseParams struct {
	// 监控统计数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*MetricDataPoint `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOverviewInvocationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewInvocationResponseParams `json:"Response"`
}

func (r *DescribeOverviewInvocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewInvocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePathRewriteRequestParams struct {
	// 路径重写规则ID
	PathRewriteId *string `json:"PathRewriteId,omitempty" name:"PathRewriteId"`
}

type DescribePathRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 路径重写规则ID
	PathRewriteId *string `json:"PathRewriteId,omitempty" name:"PathRewriteId"`
}

func (r *DescribePathRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePathRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PathRewriteId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePathRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePathRewriteResponseParams struct {
	// 路径重写规则对象
	Result *PathRewrite `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePathRewriteResponse struct {
	*tchttp.BaseResponse
	Response *DescribePathRewriteResponseParams `json:"Response"`
}

func (r *DescribePathRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePathRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePathRewritesRequestParams struct {
	// 网关部署组ID
	GatewayGroupId *string `json:"GatewayGroupId,omitempty" name:"GatewayGroupId"`

	// 根据正则表达式或替换的内容模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 每页数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribePathRewritesRequest struct {
	*tchttp.BaseRequest
	
	// 网关部署组ID
	GatewayGroupId *string `json:"GatewayGroupId,omitempty" name:"GatewayGroupId"`

	// 根据正则表达式或替换的内容模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 每页数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePathRewritesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePathRewritesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayGroupId")
	delete(f, "SearchWord")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePathRewritesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePathRewritesResponseParams struct {
	// 路径重写翻页对象
	Result *PathRewritePage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePathRewritesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePathRewritesResponseParams `json:"Response"`
}

func (r *DescribePathRewritesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePathRewritesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePkgsRequestParams struct {
	// 应用ID（只传入应用ID，返回该应用下所有软件包信息）
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 查询关键字（支持根据包ID，包名，包版本号搜索）
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序关键字（默认为"UploadTime"：上传时间）
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序：0/降序：1（默认降序）
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// 查询起始偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 程序包仓库id
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`

	// 程序包类型数组支持（fatjar jar war tar.gz zip）
	PackageTypeList []*string `json:"PackageTypeList,omitempty" name:"PackageTypeList"`
}

type DescribePkgsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID（只传入应用ID，返回该应用下所有软件包信息）
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 查询关键字（支持根据包ID，包名，包版本号搜索）
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序关键字（默认为"UploadTime"：上传时间）
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序：0/降序：1（默认降序）
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// 查询起始偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 程序包仓库id
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`

	// 程序包类型数组支持（fatjar jar war tar.gz zip）
	PackageTypeList []*string `json:"PackageTypeList,omitempty" name:"PackageTypeList"`
}

func (r *DescribePkgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePkgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SearchWord")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RepositoryType")
	delete(f, "RepositoryId")
	delete(f, "PackageTypeList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePkgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePkgsResponseParams struct {
	// 符合查询程序包信息列表
	Result *PkgList `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePkgsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePkgsResponseParams `json:"Response"`
}

func (r *DescribePkgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePkgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginInstancesRequestParams struct {
	// 分组或者API的ID
	ScopeValue *string `json:"ScopeValue,omitempty" name:"ScopeValue"`

	// 绑定: true; 未绑定: false
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页展示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 插件类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

type DescribePluginInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 分组或者API的ID
	ScopeValue *string `json:"ScopeValue,omitempty" name:"ScopeValue"`

	// 绑定: true; 未绑定: false
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 翻页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页展示的条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 插件类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribePluginInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScopeValue")
	delete(f, "Bound")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "SearchWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginInstancesResponseParams struct {
	// 插件信息列表
	Result *TsfPageGatewayPlugin `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePluginInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginInstancesResponseParams `json:"Response"`
}

func (r *DescribePluginInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePodInstancesRequestParams struct {
	// 实例所属groupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤字段
	PodNameList []*string `json:"PodNameList,omitempty" name:"PodNameList"`
}

type DescribePodInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例所属groupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤字段
	PodNameList []*string `json:"PodNameList,omitempty" name:"PodNameList"`
}

func (r *DescribePodInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePodInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PodNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePodInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePodInstancesResponseParams struct {
	// 查询的权限数据对象
	Result *GroupPodResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePodInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePodInstancesResponseParams `json:"Response"`
}

func (r *DescribePodInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePodInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProgramsRequestParams struct {
	// 模糊查询数据集ID，数据集名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeProgramsRequest struct {
	*tchttp.BaseRequest
	
	// 模糊查询数据集ID，数据集名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeProgramsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProgramsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProgramsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProgramsResponseParams struct {
	// 数据集列表
	Result *PagedProgram `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProgramsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProgramsResponseParams `json:"Response"`
}

func (r *DescribeProgramsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProgramsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigReleaseLogsRequestParams struct {
	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePublicConfigReleaseLogsRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePublicConfigReleaseLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigReleaseLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicConfigReleaseLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigReleaseLogsResponseParams struct {
	// 分页后的公共配置项发布历史列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageConfigReleaseLog `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePublicConfigReleaseLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicConfigReleaseLogsResponseParams `json:"Response"`
}

func (r *DescribePublicConfigReleaseLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigReleaseLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigReleasesRequestParams struct {
	// 配置项名称，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 配置项ID，不传入时查询全量
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

type DescribePublicConfigReleasesRequest struct {
	*tchttp.BaseRequest
	
	// 配置项名称，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 配置项ID，不传入时查询全量
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribePublicConfigReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigReleasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigName")
	delete(f, "NamespaceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicConfigReleasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigReleasesResponseParams struct {
	// 公共配置发布信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageConfigRelease `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePublicConfigReleasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicConfigReleasesResponseParams `json:"Response"`
}

func (r *DescribePublicConfigReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigReleasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigRequestParams struct {
	// 需要查询的配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

type DescribePublicConfigRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigResponseParams struct {
	// 全局配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *Config `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicConfigResponseParams `json:"Response"`
}

func (r *DescribePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigSummaryRequestParams struct {
	// 查询关键字，模糊查询：配置项名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 按时间排序：creation_time；按名称排序：config_name
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序传 0，降序传 1
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 无
	ConfigTagList []*string `json:"ConfigTagList,omitempty" name:"ConfigTagList"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`

	// 无
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`
}

type DescribePublicConfigSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 查询关键字，模糊查询：配置项名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 按时间排序：creation_time；按名称排序：config_name
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序传 0，降序传 1
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 无
	ConfigTagList []*string `json:"ConfigTagList,omitempty" name:"ConfigTagList"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`

	// 无
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`
}

func (r *DescribePublicConfigSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "ConfigTagList")
	delete(f, "DisableProgramAuthCheck")
	delete(f, "ConfigIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicConfigSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigSummaryResponseParams struct {
	// 分页的全局配置统计信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePublicConfigSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicConfigSummaryResponseParams `json:"Response"`
}

func (r *DescribePublicConfigSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigsRequestParams struct {
	// 配置项ID，不传入时查询全量，高优先级
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项ID列表，不传入时查询全量，低优先级
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`

	// 配置项名称，精确查询，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本，精确查询，不传入时查询全量
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

type DescribePublicConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 配置项ID，不传入时查询全量，高优先级
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项ID列表，不传入时查询全量，低优先级
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`

	// 配置项名称，精确查询，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本，精确查询，不传入时查询全量
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *DescribePublicConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ConfigIdList")
	delete(f, "ConfigName")
	delete(f, "ConfigVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicConfigsResponseParams struct {
	// 分页后的全局配置项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePublicConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicConfigsResponseParams `json:"Response"`
}

func (r *DescribePublicConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleasedConfigRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeReleasedConfigRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeReleasedConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleasedConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleasedConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleasedConfigResponseParams struct {
	// 已发布的配置内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReleasedConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleasedConfigResponseParams `json:"Response"`
}

func (r *DescribeReleasedConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleasedConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoriesRequestParams struct {
	// 查询关键字（按照仓库名称搜索）
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 查询起始偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 仓库类型（默认仓库：default，私有仓库：private）
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`
}

type DescribeRepositoriesRequest struct {
	*tchttp.BaseRequest
	
	// 查询关键字（按照仓库名称搜索）
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 查询起始偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 仓库类型（默认仓库：default，私有仓库：private）
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`
}

func (r *DescribeRepositoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RepositoryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepositoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoriesResponseParams struct {
	// 符合查询仓库信息列表
	Result *RepositoryList `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRepositoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoriesResponseParams `json:"Response"`
}

func (r *DescribeRepositoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryRequestParams struct {
	// 仓库ID
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

type DescribeRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 仓库ID
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

func (r *DescribeRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepositoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryResponseParams struct {
	// 查询的仓库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *RepositoryInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoryResponseParams `json:"Response"`
}

func (r *DescribeRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleApplicationsRequestParams struct {
	// 应用ID列表
	ApplicationIdList []*string `json:"ApplicationIdList,omitempty" name:"ApplicationIdList"`

	// 应用类型
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 微服务类型
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 资源类型数组
	ApplicationResourceTypeList []*string `json:"ApplicationResourceTypeList,omitempty" name:"ApplicationResourceTypeList"`

	// 通过id和name进行关键词过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`
}

type DescribeSimpleApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID列表
	ApplicationIdList []*string `json:"ApplicationIdList,omitempty" name:"ApplicationIdList"`

	// 应用类型
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 微服务类型
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 资源类型数组
	ApplicationResourceTypeList []*string `json:"ApplicationResourceTypeList,omitempty" name:"ApplicationResourceTypeList"`

	// 通过id和name进行关键词过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`
}

func (r *DescribeSimpleApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationIdList")
	delete(f, "ApplicationType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "MicroserviceType")
	delete(f, "ApplicationResourceTypeList")
	delete(f, "SearchWord")
	delete(f, "DisableProgramAuthCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSimpleApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleApplicationsResponseParams struct {
	// 简单应用分页对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageSimpleApplication `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSimpleApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSimpleApplicationsResponseParams `json:"Response"`
}

func (r *DescribeSimpleApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleClustersRequestParams struct {
	// 需要查询的集群ID列表，不填或不传入时查询所有内容
	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`

	// 需要查询的集群类型，不填或不传入时查询所有内容
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 查询偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 对id和name进行关键词过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`
}

type DescribeSimpleClustersRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的集群ID列表，不填或不传入时查询所有内容
	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`

	// 需要查询的集群类型，不填或不传入时查询所有内容
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 查询偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 对id和name进行关键词过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`
}

func (r *DescribeSimpleClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIdList")
	delete(f, "ClusterType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "DisableProgramAuthCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSimpleClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleClustersResponseParams struct {
	// TSF集群分页对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageCluster `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSimpleClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSimpleClustersResponseParams `json:"Response"`
}

func (r *DescribeSimpleClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleGroupsRequestParams struct {
	// 部署组ID列表，不填写时查询全量
	GroupIdList []*string `json:"GroupIdList,omitempty" name:"GroupIdList"`

	// 应用ID，不填写时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 集群ID，不填写时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间ID，不填写时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 部署组ID，不填写时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 模糊查询，部署组名称，不填写时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 部署组类型，精确过滤字段，M：service mesh, P：原生应用， G：网关应用
	AppMicroServiceType *string `json:"AppMicroServiceType,omitempty" name:"AppMicroServiceType"`
}

type DescribeSimpleGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID列表，不填写时查询全量
	GroupIdList []*string `json:"GroupIdList,omitempty" name:"GroupIdList"`

	// 应用ID，不填写时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 集群ID，不填写时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间ID，不填写时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 部署组ID，不填写时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 模糊查询，部署组名称，不填写时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 部署组类型，精确过滤字段，M：service mesh, P：原生应用， G：网关应用
	AppMicroServiceType *string `json:"AppMicroServiceType,omitempty" name:"AppMicroServiceType"`
}

func (r *DescribeSimpleGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIdList")
	delete(f, "ApplicationId")
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "GroupId")
	delete(f, "SearchWord")
	delete(f, "AppMicroServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSimpleGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleGroupsResponseParams struct {
	// 简单部署组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageSimpleGroup `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSimpleGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSimpleGroupsResponseParams `json:"Response"`
}

func (r *DescribeSimpleGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleNamespacesRequestParams struct {
	// 命名空间ID列表，不传入时查询全量
	NamespaceIdList []*string `json:"NamespaceIdList,omitempty" name:"NamespaceIdList"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 查询资源类型列表
	NamespaceResourceTypeList []*string `json:"NamespaceResourceTypeList,omitempty" name:"NamespaceResourceTypeList"`

	// 通过id和name进行过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 查询的命名空间类型列表
	NamespaceTypeList []*string `json:"NamespaceTypeList,omitempty" name:"NamespaceTypeList"`

	// 通过命名空间名精确过滤
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 通过是否是默认命名空间过滤，不传表示拉取全部命名空间。0：默认，命名空间。1：非默认命名空间
	IsDefault *string `json:"IsDefault,omitempty" name:"IsDefault"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`
}

type DescribeSimpleNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID列表，不传入时查询全量
	NamespaceIdList []*string `json:"NamespaceIdList,omitempty" name:"NamespaceIdList"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 查询资源类型列表
	NamespaceResourceTypeList []*string `json:"NamespaceResourceTypeList,omitempty" name:"NamespaceResourceTypeList"`

	// 通过id和name进行过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 查询的命名空间类型列表
	NamespaceTypeList []*string `json:"NamespaceTypeList,omitempty" name:"NamespaceTypeList"`

	// 通过命名空间名精确过滤
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 通过是否是默认命名空间过滤，不传表示拉取全部命名空间。0：默认，命名空间。1：非默认命名空间
	IsDefault *string `json:"IsDefault,omitempty" name:"IsDefault"`

	// 无
	DisableProgramAuthCheck *bool `json:"DisableProgramAuthCheck,omitempty" name:"DisableProgramAuthCheck"`
}

func (r *DescribeSimpleNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceIdList")
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "NamespaceId")
	delete(f, "NamespaceResourceTypeList")
	delete(f, "SearchWord")
	delete(f, "NamespaceTypeList")
	delete(f, "NamespaceName")
	delete(f, "IsDefault")
	delete(f, "DisableProgramAuthCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSimpleNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleNamespacesResponseParams struct {
	// 命名空间分页列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageNamespace `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSimpleNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSimpleNamespacesResponseParams `json:"Response"`
}

func (r *DescribeSimpleNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStatisticsRequestParams struct {
	// 类型：Interface、Service、Group、Instance、SQL、NoSQL
	Type *string `json:"Type,omitempty" name:"Type"`

	// 步长，单位s：60、3600、86400
	TimeStep *uint64 `json:"TimeStep,omitempty" name:"TimeStep"`

	// 偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量，取值范围[1, 50]，默认值为10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间Id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 排序字段:AvgTimeConsuming[默认]、RequestCount、ErrorRate。实例监控还支持 CpuPercent
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式：ASC:0、DESC:1
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// 开始时间：年月日 时分秒2020-05-12 14:43:12
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间：年月日 时分秒2020-05-12 14:43:12
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 维度
	MetricDimensionValues []*MetricDimensionValue `json:"MetricDimensionValues,omitempty" name:"MetricDimensionValues"`

	// 聚合关键词
	BucketKey *string `json:"BucketKey,omitempty" name:"BucketKey"`

	// 数据库
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 命名空间id数组
	NamespaceIdList []*string `json:"NamespaceIdList,omitempty" name:"NamespaceIdList"`
}

type DescribeStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 类型：Interface、Service、Group、Instance、SQL、NoSQL
	Type *string `json:"Type,omitempty" name:"Type"`

	// 步长，单位s：60、3600、86400
	TimeStep *uint64 `json:"TimeStep,omitempty" name:"TimeStep"`

	// 偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量，取值范围[1, 50]，默认值为10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间Id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 排序字段:AvgTimeConsuming[默认]、RequestCount、ErrorRate。实例监控还支持 CpuPercent
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式：ASC:0、DESC:1
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// 开始时间：年月日 时分秒2020-05-12 14:43:12
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间：年月日 时分秒2020-05-12 14:43:12
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 维度
	MetricDimensionValues []*MetricDimensionValue `json:"MetricDimensionValues,omitempty" name:"MetricDimensionValues"`

	// 聚合关键词
	BucketKey *string `json:"BucketKey,omitempty" name:"BucketKey"`

	// 数据库
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 命名空间id数组
	NamespaceIdList []*string `json:"NamespaceIdList,omitempty" name:"NamespaceIdList"`
}

func (r *DescribeStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "TimeStep")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NamespaceId")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "ServiceName")
	delete(f, "SearchWord")
	delete(f, "MetricDimensionValues")
	delete(f, "BucketKey")
	delete(f, "DbName")
	delete(f, "NamespaceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStatisticsResponseParams struct {
	// 查询服务统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ServiceStatisticsResults `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStatisticsResponseParams `json:"Response"`
}

func (r *DescribeStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务历史ID
	TaskLogId *string `json:"TaskLogId,omitempty" name:"TaskLogId"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务历史ID
	TaskLogId *string `json:"TaskLogId,omitempty" name:"TaskLogId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TaskLogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
	// 任务详情
	Result *TaskRecord `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLastStatusRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTaskLastStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskLastStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLastStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLastStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLastStatusResponseParams struct {
	// 任务上一次执行状态
	Result *TaskLastExecuteStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskLastStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLastStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskLastStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLastStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRecordsRequestParams struct {
	// 翻页偏移量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询单页数量。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊查询关键字，支持任务ID和任务名称。
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 任务启用状态。enabled/disabled
	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`

	// 分组ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 任务类型。
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务触发类型，UNICAST、BROADCAST。
	ExecuteType *string `json:"ExecuteType,omitempty" name:"ExecuteType"`

	// 无
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

type DescribeTaskRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 翻页偏移量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询单页数量。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊查询关键字，支持任务ID和任务名称。
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 任务启用状态。enabled/disabled
	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`

	// 分组ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 任务类型。
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务触发类型，UNICAST、BROADCAST。
	ExecuteType *string `json:"ExecuteType,omitempty" name:"ExecuteType"`

	// 无
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DescribeTaskRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "TaskState")
	delete(f, "GroupId")
	delete(f, "TaskType")
	delete(f, "ExecuteType")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRecordsResponseParams struct {
	// 任务记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TaskRecordPage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskRecordsResponseParams `json:"Response"`
}

func (r *DescribeTaskRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnitApiUseDetailRequestParams struct {
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组Api ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网关实例ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 监控统计数据粒度
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type DescribeUnitApiUseDetailRequest struct {
	*tchttp.BaseRequest
	
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关分组Api ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的日期,格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 网关实例ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 监控统计数据粒度
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribeUnitApiUseDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnitApiUseDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayDeployGroupId")
	delete(f, "ApiId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "GatewayInstanceId")
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnitApiUseDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnitApiUseDetailResponseParams struct {
	// 单元化使用统计对象
	Result *GroupUnitApiUseStatistics `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnitApiUseDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnitApiUseDetailResponseParams `json:"Response"`
}

func (r *DescribeUnitApiUseDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnitApiUseDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnitNamespacesRequestParams struct {
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 根据命名空间名或ID模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeUnitNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 根据命名空间名或ID模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUnitNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnitNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayInstanceId")
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnitNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnitNamespacesResponseParams struct {
	// 单元化命名空间对象列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageUnitNamespace `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnitNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnitNamespacesResponseParams `json:"Response"`
}

func (r *DescribeUnitNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnitNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnitRuleRequestParams struct {
	// 单元化规则ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeUnitRuleRequest struct {
	*tchttp.BaseRequest
	
	// 单元化规则ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeUnitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnitRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnitRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnitRuleResponseParams struct {
	// 单元化规则对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UnitRule `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnitRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnitRuleResponseParams `json:"Response"`
}

func (r *DescribeUnitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnitRulesRequestParams struct {
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 根据规则名或备注内容模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 启用状态, disabled: 未发布， enabled: 发布
	Status *string `json:"Status,omitempty" name:"Status"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeUnitRulesRequest struct {
	*tchttp.BaseRequest
	
	// 网关实体ID
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 根据规则名或备注内容模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 启用状态, disabled: 未发布， enabled: 发布
	Status *string `json:"Status,omitempty" name:"Status"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUnitRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnitRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayInstanceId")
	delete(f, "SearchWord")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnitRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnitRulesResponseParams struct {
	// 分页列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*TsfPageUnitRule `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnitRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnitRulesResponseParams `json:"Response"`
}

func (r *DescribeUnitRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnitRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadInfoRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 程序包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 程序包版本
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 程序包类型
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 程序包介绍
	PkgDesc *string `json:"PkgDesc,omitempty" name:"PkgDesc"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 程序包仓库id
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

type DescribeUploadInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 程序包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 程序包版本
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 程序包类型
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 程序包介绍
	PkgDesc *string `json:"PkgDesc,omitempty" name:"PkgDesc"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 程序包仓库id
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

func (r *DescribeUploadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PkgName")
	delete(f, "PkgVersion")
	delete(f, "PkgType")
	delete(f, "PkgDesc")
	delete(f, "RepositoryType")
	delete(f, "RepositoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUploadInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadInfoResponseParams struct {
	// COS上传信息
	Result *CosUploadInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUploadInfoResponseParams `json:"Response"`
}

func (r *DescribeUploadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsableUnitNamespacesRequestParams struct {
	// 根据命名空间名或ID模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeUsableUnitNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// 根据命名空间名或ID模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 翻页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 翻页查询每页记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUsableUnitNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsableUnitNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchWord")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsableUnitNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsableUnitNamespacesResponseParams struct {
	// 单元化命名空间对象列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageUnitNamespace `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsableUnitNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsableUnitNamespacesResponseParams `json:"Response"`
}

func (r *DescribeUsableUnitNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsableUnitNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableTaskFlowRequestParams struct {
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type DisableTaskFlowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DisableTaskFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableTaskFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableTaskFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableTaskFlowResponseParams struct {
	// true成功，false: 失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableTaskFlowResponse struct {
	*tchttp.BaseResponse
	Response *DisableTaskFlowResponseParams `json:"Response"`
}

func (r *DisableTaskFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableTaskFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DisableTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DisableTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableTaskResponseParams struct {
	// 操作成功 or 失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableTaskResponse struct {
	*tchttp.BaseResponse
	Response *DisableTaskResponseParams `json:"Response"`
}

func (r *DisableTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUnitRouteRequestParams struct {
	// 网关实体ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DisableUnitRouteRequest struct {
	*tchttp.BaseRequest
	
	// 网关实体ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DisableUnitRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUnitRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableUnitRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUnitRouteResponseParams struct {
	// 返回结果，成功失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableUnitRouteResponse struct {
	*tchttp.BaseResponse
	Response *DisableUnitRouteResponseParams `json:"Response"`
}

func (r *DisableUnitRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUnitRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUnitRuleRequestParams struct {
	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DisableUnitRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DisableUnitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUnitRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableUnitRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUnitRuleResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableUnitRuleResponse struct {
	*tchttp.BaseResponse
	Response *DisableUnitRuleResponseParams `json:"Response"`
}

func (r *DisableUnitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUnitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateBusinessLogConfigRequestParams struct {
	// 业务日志配置项ID列表
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`

	// TSF分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DisassociateBusinessLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 业务日志配置项ID列表
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList"`

	// TSF分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DisassociateBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateBusinessLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigIdList")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateBusinessLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateBusinessLogConfigResponseParams struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateBusinessLogConfigResponseParams `json:"Response"`
}

func (r *DisassociateBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateBusinessLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateKafkaConfigRequestParams struct {
	// 配置项id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组id
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

type DisassociateKafkaConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项id
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组id
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
}

func (r *DisassociateKafkaConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateKafkaConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateKafkaConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateKafkaConfigResponseParams struct {
	// 解除绑定是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateKafkaConfigResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateKafkaConfigResponseParams `json:"Response"`
}

func (r *DisassociateKafkaConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateKafkaConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DraftApiGroupRequestParams struct {
	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DraftApiGroupRequest struct {
	*tchttp.BaseRequest
	
	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DraftApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DraftApiGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DraftApiGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DraftApiGroupResponseParams struct {
	// true: 成功, false: 失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DraftApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *DraftApiGroupResponseParams `json:"Response"`
}

func (r *DraftApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DraftApiGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTaskFlowRequestParams struct {
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type EnableTaskFlowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *EnableTaskFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTaskFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableTaskFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTaskFlowResponseParams struct {
	// true成功，false: 失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableTaskFlowResponse struct {
	*tchttp.BaseResponse
	Response *EnableTaskFlowResponseParams `json:"Response"`
}

func (r *EnableTaskFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTaskFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTaskRequestParams struct {
	// 启用任务
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type EnableTaskRequest struct {
	*tchttp.BaseRequest
	
	// 启用任务
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *EnableTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTaskResponseParams struct {
	// 操作成功or失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableTaskResponse struct {
	*tchttp.BaseResponse
	Response *EnableTaskResponseParams `json:"Response"`
}

func (r *EnableTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableUnitRouteRequestParams struct {
	// 网关实体ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type EnableUnitRouteRequest struct {
	*tchttp.BaseRequest
	
	// 网关实体ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *EnableUnitRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableUnitRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableUnitRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableUnitRouteResponseParams struct {
	// 返回结果，成功失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableUnitRouteResponse struct {
	*tchttp.BaseResponse
	Response *EnableUnitRouteResponseParams `json:"Response"`
}

func (r *EnableUnitRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableUnitRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableUnitRuleRequestParams struct {
	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type EnableUnitRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *EnableUnitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableUnitRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableUnitRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableUnitRuleResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableUnitRuleResponse struct {
	*tchttp.BaseResponse
	Response *EnableUnitRuleResponseParams `json:"Response"`
}

func (r *EnableUnitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableUnitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Env struct {
	// 环境变量名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 环境变量值
	Value *string `json:"Value,omitempty" name:"Value"`

	// k8s ValueFrom
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueFrom *ValueFrom `json:"ValueFrom,omitempty" name:"ValueFrom"`
}

// Predefined struct for user
type ExecuteTaskFlowRequestParams struct {
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type ExecuteTaskFlowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *ExecuteTaskFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteTaskFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskFlowResponseParams struct {
	// 工作流批次ID
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExecuteTaskFlowResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteTaskFlowResponseParams `json:"Response"`
}

func (r *ExecuteTaskFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskRequestParams struct {
	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type ExecuteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ExecuteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteTaskResponseParams struct {
	// 成功/失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExecuteTaskResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteTaskResponseParams `json:"Response"`
}

func (r *ExecuteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 扩容的机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

type ExpandGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 扩容的机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

func (r *ExpandGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpandGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandGroupResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TaskId `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExpandGroupResponse struct {
	*tchttp.BaseResponse
	Response *ExpandGroupResponseParams `json:"Response"`
}

func (r *ExpandGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldRef struct {
	// k8s 的 FieldPath
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldPath *string `json:"FieldPath,omitempty" name:"FieldPath"`
}

type FileConfig struct {
	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项版本描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFileName *string `json:"ConfigFileName,omitempty" name:"ConfigFileName"`

	// 配置项文件内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFileValue *string `json:"ConfigFileValue,omitempty" name:"ConfigFileValue"`

	// 配置项文件编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFileCode *string `json:"ConfigFileCode,omitempty" name:"ConfigFileCode"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 配置项归属应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 删除标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 配置项版本数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionCount *int64 `json:"ConfigVersionCount,omitempty" name:"ConfigVersionCount"`

	// 配置项最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 发布路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFilePath *string `json:"ConfigFilePath,omitempty" name:"ConfigFilePath"`

	// 后置命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigPostCmd *string `json:"ConfigPostCmd,omitempty" name:"ConfigPostCmd"`

	// 配置项文件长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigFileValueLength *uint64 `json:"ConfigFileValueLength,omitempty" name:"ConfigFileValueLength"`
}

type FileConfigRelease struct {
	// 配置项发布ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type Filter struct {
	// 过滤条件名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤条件匹配值，几个条件间是或关系
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type GatewayApiGroupVo struct {
	// 分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组下API个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupApiCount *uint64 `json:"GroupApiCount,omitempty" name:"GroupApiCount"`

	// 分组API列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupApis []*GatewayGroupApiVo `json:"GroupApis,omitempty" name:"GroupApis"`

	// 网关实例的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayInstanceType *string `json:"GatewayInstanceType,omitempty" name:"GatewayInstanceType"`

	// 网关实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`
}

type GatewayConfig struct {

}

type GatewayDeployGroup struct {
	// 网关部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployGroupId *string `json:"DeployGroupId,omitempty" name:"DeployGroupId"`

	// 网关部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployGroupName *string `json:"DeployGroupName,omitempty" name:"DeployGroupName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用分类：V：虚拟机应用，C：容器应用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 部署组应用状态,取值: Running、Waiting、Paused、Updating、RollingBack、Abnormal、Unknown
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *string `json:"GroupStatus,omitempty" name:"GroupStatus"`

	// 集群类型，C ：容器，V：虚拟机
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type GatewayGroupApiVo struct {
	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API 请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// API 微服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// API 请求方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type GatewayGroupIds struct {
	// 网关部署组ID
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 分组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type GatewayPlugin struct {
	// 网关插件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 插件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 插件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 插件描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 发布状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type GatewayPluginBoundParam struct {
	// 插件id
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 插件绑定到的对象类型:group/api
	ScopeType *string `json:"ScopeType,omitempty" name:"ScopeType"`

	// 插件绑定到的对象主键值，例如分组的ID/API的ID
	ScopeValue *string `json:"ScopeValue,omitempty" name:"ScopeValue"`
}

type GatewayVo struct {
	// 网关部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayDeployGroupId *string `json:"GatewayDeployGroupId,omitempty" name:"GatewayDeployGroupId"`

	// 网关部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayDeployGroupName *string `json:"GatewayDeployGroupName,omitempty" name:"GatewayDeployGroupName"`

	// API 分组个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

	// API 分组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*GatewayApiGroupVo `json:"Groups,omitempty" name:"Groups"`
}

type GroupApiUseStatistics struct {
	// 总调用数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopStatusCode []*ApiUseStatisticsEntity `json:"TopStatusCode,omitempty" name:"TopStatusCode"`

	// 平均错误率
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopTimeCost []*ApiUseStatisticsEntity `json:"TopTimeCost,omitempty" name:"TopTimeCost"`

	// 分位值对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quantile *QuantileEntity `json:"Quantile,omitempty" name:"Quantile"`
}

type GroupDailyUseStatistics struct {
	// 总调用数
	TopReqAmount []*GroupUseStatisticsEntity `json:"TopReqAmount,omitempty" name:"TopReqAmount"`

	// 平均错误率
	TopFailureRate []*GroupUseStatisticsEntity `json:"TopFailureRate,omitempty" name:"TopFailureRate"`

	// 平均响应耗时
	TopAvgTimeCost []*GroupUseStatisticsEntity `json:"TopAvgTimeCost,omitempty" name:"TopAvgTimeCost"`
}

type GroupInfo struct {
	// 部署组id
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 绑定时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateTime *string `json:"AssociateTime,omitempty" name:"AssociateTime"`
}

type GroupPod struct {
	// 实例名称(对应到kubernetes的pod名称)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 实例ID(对应到kubernetes的pod id)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// 实例状态，请参考后面的实例以及容器的状态定义。启动中（pod 未 ready）：Starting；运行中：Running；异常：Abnormal；停止：Stopped；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例处于当前状态的原因，例如容器下载镜像失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 主机IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 实例中容器的重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// 实例中已就绪容器的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadyCount *int64 `json:"ReadyCount,omitempty" name:"ReadyCount"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 实例启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 服务实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

	// 机器实例可使用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 机器实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 节点实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" name:"NodeInstanceId"`
}

type GroupPodResult struct {
	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*GroupPod `json:"Content,omitempty" name:"Content"`
}

type GroupRelease struct {
	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 程序包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 镜像名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 镜像版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 已发布的全局配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicConfigReleaseList []*ConfigRelease `json:"PublicConfigReleaseList,omitempty" name:"PublicConfigReleaseList"`

	// 已发布的应用配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseList []*ConfigRelease `json:"ConfigReleaseList,omitempty" name:"ConfigReleaseList"`

	// 已发布的文件配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileConfigReleaseList []*FileConfigRelease `json:"FileConfigReleaseList,omitempty" name:"FileConfigReleaseList"`
}

type GroupUnitApiDailyUseStatistics struct {
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 该API在该命名空间下的总调用次数
	SumReqAmount *string `json:"SumReqAmount,omitempty" name:"SumReqAmount"`

	// 该API在该命名空间下的平均错误率
	AvgFailureRate *string `json:"AvgFailureRate,omitempty" name:"AvgFailureRate"`

	// 该API在该命名空间下的平均响应时间
	AvgTimeCost *string `json:"AvgTimeCost,omitempty" name:"AvgTimeCost"`

	// 监控数据曲线点位图Map集合
	MetricDataPointMap *MetricDataPointMap `json:"MetricDataPointMap,omitempty" name:"MetricDataPointMap"`

	// 状态码分布详情
	TopStatusCode []*ApiUseStatisticsEntity `json:"TopStatusCode,omitempty" name:"TopStatusCode"`

	// 耗时分布详情
	TopTimeCost []*ApiUseStatisticsEntity `json:"TopTimeCost,omitempty" name:"TopTimeCost"`

	// 分位值对象
	Quantile *QuantileEntity `json:"Quantile,omitempty" name:"Quantile"`
}

type GroupUnitApiUseStatistics struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询网关API监控明细对象集合
	Content []*GroupUnitApiDailyUseStatistics `json:"Content,omitempty" name:"Content"`
}

type GroupUseStatisticsEntity struct {
	// API 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiPath *string `json:"ApiPath,omitempty" name:"ApiPath"`

	// 服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// API ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type HealthCheckConfig struct {
	// 健康检查路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`
}

type HealthCheckSetting struct {
	// 健康检查方法。HTTP：通过 HTTP 接口检查；CMD：通过执行命令检查；TCP：通过建立 TCP 连接检查。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 容器延时启动健康检查的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialDelaySeconds *uint64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// 每次健康检查响应的最大超时时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutSeconds *uint64 `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`

	// 进行健康检查的时间间隔。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodSeconds *uint64 `json:"PeriodSeconds,omitempty" name:"PeriodSeconds"`

	// 表示后端容器从失败到成功的连续健康检查成功次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessThreshold *uint64 `json:"SuccessThreshold,omitempty" name:"SuccessThreshold"`

	// 表示后端容器从成功到失败的连续健康检查成功次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureThreshold *uint64 `json:"FailureThreshold,omitempty" name:"FailureThreshold"`

	// HTTP 健康检查方法使用的检查协议。支持HTTP、HTTPS。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 健康检查端口，范围 1~65535 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// HTTP 健康检查接口的请求路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 执行命令检查方式，执行的命令。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Command []*string `json:"Command,omitempty" name:"Command"`

	// TSF_DEFAULT：tsf 默认就绪探针。K8S_NATIVE：k8s 原生探针。不填默认为 k8s 原生探针。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type HealthCheckSettings struct {
	// 存活健康检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessProbe *HealthCheckSetting `json:"LivenessProbe,omitempty" name:"LivenessProbe"`

	// 就绪健康检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadinessProbe *HealthCheckSetting `json:"ReadinessProbe,omitempty" name:"ReadinessProbe"`
}

type ImageRepository struct {
	// 仓库名,含命名空间,如tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Repotype *string `json:"Repotype,omitempty" name:"Repotype"`

	// 镜像版本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`

	// 是否公共,1:公有,0:私有
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPublic *int64 `json:"IsPublic,omitempty" name:"IsPublic"`

	// 是否被用户收藏。true：是，false：否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`

	// 是否是腾讯云官方仓库。 是否是腾讯云官方仓库。true：是，false：否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`

	// 被所有用户收藏次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

	// 拉取次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`

	// 描述内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// TcrRepoInfo值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcrRepoInfo *TcrRepoInfo `json:"TcrRepoInfo,omitempty" name:"TcrRepoInfo"`

	// TcrBindingId值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcrBindingId *int64 `json:"TcrBindingId,omitempty" name:"TcrBindingId"`

	// applicationid值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// ApplicationName值（废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *ScalableRule `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// ApplicationName值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationNameReal *string `json:"ApplicationNameReal,omitempty" name:"ApplicationNameReal"`

	// 是否公共,1:公有,0:私有
	// 注意：此字段可能返回 null，表示取不到有效值。
	Public *int64 `json:"Public,omitempty" name:"Public"`
}

type ImageRepositoryResult struct {
	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 镜像服务器地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ImageRepository `json:"Content,omitempty" name:"Content"`
}

type ImageTag struct {
	// 仓库名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 版本名称
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 版本ID
	TagId *string `json:"TagId,omitempty" name:"TagId"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 大小
	Size *string `json:"Size,omitempty" name:"Size"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 镜像制作者
	Author *string `json:"Author,omitempty" name:"Author"`

	// CPU架构
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// Docker客户端版本
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	Os *string `json:"Os,omitempty" name:"Os"`

	// push时间
	PushTime *string `json:"PushTime,omitempty" name:"PushTime"`

	// 单位为字节
	SizeByte *int64 `json:"SizeByte,omitempty" name:"SizeByte"`

	// TcrRepoInfo值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcrRepoInfo *TcrRepoInfo `json:"TcrRepoInfo,omitempty" name:"TcrRepoInfo"`
}

type ImageTagsResult struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库名,含命名空间,如tsf/ngin
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 镜像服务器地址
	Server *string `json:"Server,omitempty" name:"Server"`

	// 列表信息
	Content []*ImageTag `json:"Content,omitempty" name:"Content"`
}

type IndicatorCoord struct {
	// 指标横坐标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoordX *string `json:"CoordX,omitempty" name:"CoordX"`

	// 指标纵坐标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoordY *string `json:"CoordY,omitempty" name:"CoordY"`

	// 指标标签，用于标识附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoordTag *string `json:"CoordTag,omitempty" name:"CoordTag"`
}

type Instance struct {
	// 机器实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 机器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 机器内网地址IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 机器外网地址IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 机器描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// VM的状态 虚机：虚机的状态 容器：Pod所在虚机的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// VM的可使用状态 虚机：虚机是否能够作为资源使用 容器：虚机是否能够作为资源部署POD
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 服务下的服务实例的状态 虚机：应用是否可用 + Agent状态 容器：Pod状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

	// 标识此instance是否已添加在tsf中
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountInTsf *int64 `json:"CountInTsf,omitempty" name:"CountInTsf"`

	// 机器所属部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器所属应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 机器所属应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 机器实例在CVM的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCreatedTime *string `json:"InstanceCreatedTime,omitempty" name:"InstanceCreatedTime"`

	// 机器实例在CVM的过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceExpiredTime *string `json:"InstanceExpiredTime,omitempty" name:"InstanceExpiredTime"`

	// 机器实例在CVM的计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 机器实例总CPU信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTotalCpu *float64 `json:"InstanceTotalCpu,omitempty" name:"InstanceTotalCpu"`

	// 机器实例总内存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTotalMem *float64 `json:"InstanceTotalMem,omitempty" name:"InstanceTotalMem"`

	// 机器实例使用的CPU信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUsedCpu *float64 `json:"InstanceUsedCpu,omitempty" name:"InstanceUsedCpu"`

	// 机器实例使用的内存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUsedMem *float64 `json:"InstanceUsedMem,omitempty" name:"InstanceUsedMem"`

	// 机器实例Limit CPU信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLimitCpu *float64 `json:"InstanceLimitCpu,omitempty" name:"InstanceLimitCpu"`

	// 机器实例Limit 内存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLimitMem *float64 `json:"InstanceLimitMem,omitempty" name:"InstanceLimitMem"`

	// 包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePkgVersion *string `json:"InstancePkgVersion,omitempty" name:"InstancePkgVersion"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 机器实例业务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例执行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationState *int64 `json:"OperationState,omitempty" name:"OperationState"`

	// NamespaceId Ns ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// InstanceZoneId 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceZoneId *string `json:"InstanceZoneId,omitempty" name:"InstanceZoneId"`

	// InstanceImportMode 导入模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`

	// ApplicationType应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// ApplicationResourceType 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// sidecar状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceSidecarStatus *string `json:"ServiceSidecarStatus,omitempty" name:"ServiceSidecarStatus"`

	// 部署组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// NS名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 健康检查原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// agent版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`

	// 容器母机实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" name:"NodeInstanceId"`
}

type InstanceAdvancedSettings struct {
	// 数据盘挂载点, 默认不挂载数据盘. 已格式化的 ext3，ext4，xfs 文件系统的数据盘将直接挂载，其他文件系统或未格式化的数据盘将自动格式化为ext4 并挂载，请注意备份数据! 无数据盘或有多块数据盘的云主机此设置不生效。
	// 注意，注意，多盘场景请使用下方的DataDisks数据结构，设置对应的云盘类型、云盘大小、挂载路径、是否格式化等信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`

	// dockerd --graph 指定值, 默认为 /var/lib/docker
	// 注意：此字段可能返回 null，表示取不到有效值。
	DockerGraphPath *string `json:"DockerGraphPath,omitempty" name:"DockerGraphPath"`
}

type InstanceEnrichedInfo struct {
	// 机器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 机器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 机器内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 机器外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 机器所在VPC
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 机器运行状态 Pending Running Stopped Rebooting Starting Stopping Abnormal Unknown
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 机器可用状态（表示机器上的Agent在线）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 机器所在部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type InstanceEnrichedInfoPage struct {
	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*InstanceEnrichedInfo `json:"Content,omitempty" name:"Content"`
}

type InvocationIndicator struct {
	// 总请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationQuantity *int64 `json:"InvocationQuantity,omitempty" name:"InvocationQuantity"`

	// 请求成功率，百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationSuccessRate *float64 `json:"InvocationSuccessRate,omitempty" name:"InvocationSuccessRate"`

	// 请求平均耗时，单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationAvgDuration *float64 `json:"InvocationAvgDuration,omitempty" name:"InvocationAvgDuration"`

	// 成功请求数时间分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationSuccessDistribution []*IndicatorCoord `json:"InvocationSuccessDistribution,omitempty" name:"InvocationSuccessDistribution"`

	// 失败请求数时间分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationFailedDistribution []*IndicatorCoord `json:"InvocationFailedDistribution,omitempty" name:"InvocationFailedDistribution"`

	// 状态码分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationStatusDistribution []*IndicatorCoord `json:"InvocationStatusDistribution,omitempty" name:"InvocationStatusDistribution"`

	// 时延分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationDurationDistribution []*IndicatorCoord `json:"InvocationDurationDistribution,omitempty" name:"InvocationDurationDistribution"`

	// 并发请求次数时间分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationQuantityDistribution []*IndicatorCoord `json:"InvocationQuantityDistribution,omitempty" name:"InvocationQuantityDistribution"`
}

type InvocationMetricScatterPlot struct {
	// 时间轴截止时间，GMT，精确到毫秒
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 时间粒度
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 时间轴开始时间，GMT，精确到毫秒
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 多值数据点集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataPoints []*MultiValueDataPoints `json:"DataPoints,omitempty" name:"DataPoints"`
}

type JvmMonitorData struct {
	// 堆内存监控图,三条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeapMemory *MemoryPicture `json:"HeapMemory,omitempty" name:"HeapMemory"`

	// 非堆内存监控图,三条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	NonHeapMemory *MemoryPicture `json:"NonHeapMemory,omitempty" name:"NonHeapMemory"`

	// 伊甸园区监控图,三条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdenSpace *MemoryPicture `json:"EdenSpace,omitempty" name:"EdenSpace"`

	// 幸存者区监控图,三条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurvivorSpace *MemoryPicture `json:"SurvivorSpace,omitempty" name:"SurvivorSpace"`

	// 老年代监控图,三条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldSpace *MemoryPicture `json:"OldSpace,omitempty" name:"OldSpace"`

	// 元空间监控图,三条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaSpace *MemoryPicture `json:"MetaSpace,omitempty" name:"MetaSpace"`

	// 线程监控图,三条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThreadPicture *ThreadPicture `json:"ThreadPicture,omitempty" name:"ThreadPicture"`

	// youngGC增量监控图,一条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	YoungGC []*CurvePoint `json:"YoungGC,omitempty" name:"YoungGC"`

	// fullGC增量监控图,一条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullGC []*CurvePoint `json:"FullGC,omitempty" name:"FullGC"`

	// cpu使用率,一条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuUsage []*CurvePoint `json:"CpuUsage,omitempty" name:"CpuUsage"`

	// 加载类数,一条线
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassCount []*CurvePoint `json:"ClassCount,omitempty" name:"ClassCount"`
}

type KafkaDeliveryConfig struct {
	// 配置项id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 采集路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectPath []*string `json:"CollectPath,omitempty" name:"CollectPath"`

	// kafka vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaVIp *string `json:"KafkaVIp,omitempty" name:"KafkaVIp"`

	// kafka vport
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaVPort *string `json:"KafkaVPort,omitempty" name:"KafkaVPort"`

	// kafka topic
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 换行规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	LineRule *string `json:"LineRule,omitempty" name:"LineRule"`

	// 是否需要认证
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableAuth *bool `json:"EnableAuth,omitempty" name:"EnableAuth"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 投递的topic和path
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaInfos []*DeliveryKafkaInfo `json:"KafkaInfos,omitempty" name:"KafkaInfos"`

	// 是否应用单行规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableGlobalLineRule *bool `json:"EnableGlobalLineRule,omitempty" name:"EnableGlobalLineRule"`

	// 自定义分行规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomRule *string `json:"CustomRule,omitempty" name:"CustomRule"`

	// KafkaAddress
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaAddress *string `json:"KafkaAddress,omitempty" name:"KafkaAddress"`
}

type LaneGroup struct {
	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 是否入口应用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Entrance *bool `json:"Entrance,omitempty" name:"Entrance"`

	// 泳道部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LaneGroupId *string `json:"LaneGroupId,omitempty" name:"LaneGroupId"`

	// 泳道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`

	// 部署组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type LaneInfo struct {
	// 泳道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`

	// 泳道名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LaneName *string `json:"LaneName,omitempty" name:"LaneName"`

	// 泳道备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 泳道部署组
	// 注意：此字段可能返回 null，表示取不到有效值。
	LaneGroupList []*LaneGroup `json:"LaneGroupList,omitempty" name:"LaneGroupList"`

	// 是否入口应用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Entrance *bool `json:"Entrance,omitempty" name:"Entrance"`

	// 泳道已经关联部署组的命名空间列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceIdList []*string `json:"NamespaceIdList,omitempty" name:"NamespaceIdList"`
}

type LaneInfos struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 泳道信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*LaneInfo `json:"Content,omitempty" name:"Content"`
}

type LaneRule struct {
	// 泳道规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 泳道规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 泳道规则标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTagList []*LaneRuleTag `json:"RuleTagList,omitempty" name:"RuleTagList"`

	// 泳道规则标签关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTagRelationship *string `json:"RuleTagRelationship,omitempty" name:"RuleTagRelationship"`

	// 泳道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`

	// 开启状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type LaneRuleTag struct {
	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagId *string `json:"TagId,omitempty" name:"TagId"`

	// 标签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 标签操作符
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagOperator *string `json:"TagOperator,omitempty" name:"TagOperator"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 泳道规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LaneRuleId *string `json:"LaneRuleId,omitempty" name:"LaneRuleId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type LaneRules struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 泳道规则列表
	Content []*LaneRule `json:"Content,omitempty" name:"Content"`
}

type MemoryPicture struct {
	// 内存最大值
	Max []*CurvePoint `json:"Max,omitempty" name:"Max"`

	// 已用内存大小
	Used []*CurvePoint `json:"Used,omitempty" name:"Used"`

	// 系统分配内存大小
	Committed []*CurvePoint `json:"Committed,omitempty" name:"Committed"`
}

type Metric struct {
	// 指标名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指标计算方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Function *string `json:"Function,omitempty" name:"Function"`
}

type MetricDataCurve struct {
	// 指标名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标计算方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricFunction *string `json:"MetricFunction,omitempty" name:"MetricFunction"`

	// 指标数据点集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricDataPoints []*MetricDataPoint `json:"MetricDataPoints,omitempty" name:"MetricDataPoints"`
}

type MetricDataPoint struct {
	// 数据点键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 数据点值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 数据点标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type MetricDataPointMap struct {
	// 总调用次数监控数据点集合
	SumReqAmount []*MetricDataPoint `json:"SumReqAmount,omitempty" name:"SumReqAmount"`

	// 平均错误率监控数据点集合
	AvgFailureRate []*MetricDataPoint `json:"AvgFailureRate,omitempty" name:"AvgFailureRate"`

	// 平均响应时间监控数据点集合
	AvgTimeCost []*MetricDataPoint `json:"AvgTimeCost,omitempty" name:"AvgTimeCost"`
}

type MetricDataSingleValue struct {
	// 指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricFunction *string `json:"MetricFunction,omitempty" name:"MetricFunction"`

	// 指标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricDataValue *string `json:"MetricDataValue,omitempty" name:"MetricDataValue"`

	// 日环比
	// 注意：此字段可能返回 null，表示取不到有效值。
	DailyPercent *float64 `json:"DailyPercent,omitempty" name:"DailyPercent"`
}

type MetricDimension struct {
	// 指标维度名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指标维度取值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type MetricDimensionValue struct {
	// 维度名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度值
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type Microservice struct {
	// 微服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务的运行实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 微服务的离线实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	CriticalInstanceCount *int64 `json:"CriticalInstanceCount,omitempty" name:"CriticalInstanceCount"`
}

// Predefined struct for user
type ModifyApplicationRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用备注
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 应用备注名
	ApplicationRemarkName *string `json:"ApplicationRemarkName,omitempty" name:"ApplicationRemarkName"`

	// 服务配置信息列表
	ServiceConfigList []*ServiceConfig `json:"ServiceConfigList,omitempty" name:"ServiceConfigList"`
}

type ModifyApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用备注
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 应用备注名
	ApplicationRemarkName *string `json:"ApplicationRemarkName,omitempty" name:"ApplicationRemarkName"`

	// 服务配置信息列表
	ServiceConfigList []*ServiceConfig `json:"ServiceConfigList,omitempty" name:"ServiceConfigList"`
}

func (r *ModifyApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ApplicationName")
	delete(f, "ApplicationDesc")
	delete(f, "ApplicationRemarkName")
	delete(f, "ServiceConfigList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationResponseParams struct {
	// true/false
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationResponseParams `json:"Response"`
}

func (r *ModifyApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 备注名
	ClusterRemarkName *string `json:"ClusterRemarkName,omitempty" name:"ClusterRemarkName"`
}

type ModifyClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 备注名
	ClusterRemarkName *string `json:"ClusterRemarkName,omitempty" name:"ClusterRemarkName"`
}

func (r *ModifyClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	delete(f, "ClusterDesc")
	delete(f, "ClusterRemarkName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterResponseParams struct {
	// 更新集群详情操作是否成功。
	// true： 操作成功。
	// false：操作失败。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterResponseParams `json:"Response"`
}

func (r *ModifyClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContainerGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// ProtocolPorts数组
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 部署组备注
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type ModifyContainerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// ProtocolPorts数组
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 部署组备注
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *ModifyContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContainerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "AccessType")
	delete(f, "ProtocolPorts")
	delete(f, "UpdateType")
	delete(f, "UpdateIvl")
	delete(f, "SubnetId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyContainerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContainerGroupResponseParams struct {
	// 更新部署组是否成功。
	// true：成功。
	// false：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyContainerGroupResponseParams `json:"Response"`
}

func (r *ModifyContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContainerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContainerReplicasRequestParams struct {
	// 部署组ID，部署组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

type ModifyContainerReplicasRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID，部署组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

func (r *ModifyContainerReplicasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContainerReplicasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyContainerReplicasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyContainerReplicasResponseParams struct {
	// 结果true：成功；false：失败；
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyContainerReplicasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyContainerReplicasResponseParams `json:"Response"`
}

func (r *ModifyContainerReplicasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyContainerReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署组描述
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 部署组备注
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署组描述
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 部署组备注
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "GroupDesc")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupResponseParams struct {
	// 更新部署组详情是否成功。
	// true：操作成功。
	// false：操作失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGroupResponseParams `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaneRequestParams struct {
	// 泳道ID
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`

	// 泳道名称
	LaneName *string `json:"LaneName,omitempty" name:"LaneName"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyLaneRequest struct {
	*tchttp.BaseRequest
	
	// 泳道ID
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`

	// 泳道名称
	LaneName *string `json:"LaneName,omitempty" name:"LaneName"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyLaneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaneId")
	delete(f, "LaneName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLaneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaneResponseParams struct {
	// 操作状态
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLaneResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLaneResponseParams `json:"Response"`
}

func (r *ModifyLaneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaneRuleRequestParams struct {
	// 泳道规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 泳道规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 泳道规则备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 泳道规则标签列表
	RuleTagList []*LaneRuleTag `json:"RuleTagList,omitempty" name:"RuleTagList"`

	// 泳道规则标签关系
	RuleTagRelationship *string `json:"RuleTagRelationship,omitempty" name:"RuleTagRelationship"`

	// 泳道ID
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`

	// 开启状态
	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type ModifyLaneRuleRequest struct {
	*tchttp.BaseRequest
	
	// 泳道规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 泳道规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 泳道规则备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 泳道规则标签列表
	RuleTagList []*LaneRuleTag `json:"RuleTagList,omitempty" name:"RuleTagList"`

	// 泳道规则标签关系
	RuleTagRelationship *string `json:"RuleTagRelationship,omitempty" name:"RuleTagRelationship"`

	// 泳道ID
	LaneId *string `json:"LaneId,omitempty" name:"LaneId"`

	// 开启状态
	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyLaneRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaneRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "Remark")
	delete(f, "RuleTagList")
	delete(f, "RuleTagRelationship")
	delete(f, "LaneId")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLaneRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaneRuleResponseParams struct {
	// 操作状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLaneRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLaneRuleResponseParams `json:"Response"`
}

func (r *ModifyLaneRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaneRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMicroserviceRequestParams struct {
	// 微服务 ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务备注信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

type ModifyMicroserviceRequest struct {
	*tchttp.BaseRequest
	
	// 微服务 ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务备注信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

func (r *ModifyMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMicroserviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MicroserviceId")
	delete(f, "MicroserviceDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMicroserviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMicroserviceResponseParams struct {
	// 修改微服务详情是否成功。
	// true：操作成功。
	// false：操作失败。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMicroserviceResponseParams `json:"Response"`
}

func (r *ModifyMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMicroserviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceRequestParams struct {
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 命名空间备注
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 是否开启高可用
	IsHaEnable *string `json:"IsHaEnable,omitempty" name:"IsHaEnable"`
}

type ModifyNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 命名空间备注
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 是否开启高可用
	IsHaEnable *string `json:"IsHaEnable,omitempty" name:"IsHaEnable"`
}

func (r *ModifyNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "NamespaceName")
	delete(f, "NamespaceDesc")
	delete(f, "IsHaEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceResponseParams struct {
	// Result
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNamespaceResponseParams `json:"Response"`
}

func (r *ModifyNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPathRewriteRequestParams struct {
	// 路径重写规则ID
	PathRewriteId *string `json:"PathRewriteId,omitempty" name:"PathRewriteId"`

	// 正则表达式
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// 替换的内容
	Replacement *string `json:"Replacement,omitempty" name:"Replacement"`

	// 是否屏蔽映射后路径，Y: 是 N: 否
	Blocked *string `json:"Blocked,omitempty" name:"Blocked"`

	// 规则顺序，越小优先级越高
	Order *int64 `json:"Order,omitempty" name:"Order"`
}

type ModifyPathRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 路径重写规则ID
	PathRewriteId *string `json:"PathRewriteId,omitempty" name:"PathRewriteId"`

	// 正则表达式
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// 替换的内容
	Replacement *string `json:"Replacement,omitempty" name:"Replacement"`

	// 是否屏蔽映射后路径，Y: 是 N: 否
	Blocked *string `json:"Blocked,omitempty" name:"Blocked"`

	// 规则顺序，越小优先级越高
	Order *int64 `json:"Order,omitempty" name:"Order"`
}

func (r *ModifyPathRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPathRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PathRewriteId")
	delete(f, "Regex")
	delete(f, "Replacement")
	delete(f, "Blocked")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPathRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPathRewriteResponseParams struct {
	// true/false
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPathRewriteResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPathRewriteResponseParams `json:"Response"`
}

func (r *ModifyPathRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPathRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务内容
	TaskContent *string `json:"TaskContent,omitempty" name:"TaskContent"`

	// 任务执行类型
	ExecuteType *string `json:"ExecuteType,omitempty" name:"ExecuteType"`

	// 触发规则
	TaskRule *TaskRule `json:"TaskRule,omitempty" name:"TaskRule"`

	// 超时时间，单位 ms
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分片数量
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 分片参数
	ShardArguments []*ShardArgument `json:"ShardArguments,omitempty" name:"ShardArguments"`

	// 高级设置
	AdvanceSettings *AdvanceSettings `json:"AdvanceSettings,omitempty" name:"AdvanceSettings"`

	// 判断任务成功的操作符 GT/GTE
	SuccessOperator *string `json:"SuccessOperator,omitempty" name:"SuccessOperator"`

	// 判断任务成功率的阈值
	SuccessRatio *int64 `json:"SuccessRatio,omitempty" name:"SuccessRatio"`

	// 重试次数
	RetryCount *uint64 `json:"RetryCount,omitempty" name:"RetryCount"`

	// 重试间隔
	RetryInterval *uint64 `json:"RetryInterval,omitempty" name:"RetryInterval"`

	// 任务参数，长度限制10000个字符
	TaskArgument *string `json:"TaskArgument,omitempty" name:"TaskArgument"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

type ModifyTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务内容
	TaskContent *string `json:"TaskContent,omitempty" name:"TaskContent"`

	// 任务执行类型
	ExecuteType *string `json:"ExecuteType,omitempty" name:"ExecuteType"`

	// 触发规则
	TaskRule *TaskRule `json:"TaskRule,omitempty" name:"TaskRule"`

	// 超时时间，单位 ms
	TimeOut *uint64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分片数量
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 分片参数
	ShardArguments []*ShardArgument `json:"ShardArguments,omitempty" name:"ShardArguments"`

	// 高级设置
	AdvanceSettings *AdvanceSettings `json:"AdvanceSettings,omitempty" name:"AdvanceSettings"`

	// 判断任务成功的操作符 GT/GTE
	SuccessOperator *string `json:"SuccessOperator,omitempty" name:"SuccessOperator"`

	// 判断任务成功率的阈值
	SuccessRatio *int64 `json:"SuccessRatio,omitempty" name:"SuccessRatio"`

	// 重试次数
	RetryCount *uint64 `json:"RetryCount,omitempty" name:"RetryCount"`

	// 重试间隔
	RetryInterval *uint64 `json:"RetryInterval,omitempty" name:"RetryInterval"`

	// 任务参数，长度限制10000个字符
	TaskArgument *string `json:"TaskArgument,omitempty" name:"TaskArgument"`

	// 无
	ProgramIdList []*string `json:"ProgramIdList,omitempty" name:"ProgramIdList"`
}

func (r *ModifyTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TaskName")
	delete(f, "TaskType")
	delete(f, "TaskContent")
	delete(f, "ExecuteType")
	delete(f, "TaskRule")
	delete(f, "TimeOut")
	delete(f, "GroupId")
	delete(f, "ShardCount")
	delete(f, "ShardArguments")
	delete(f, "AdvanceSettings")
	delete(f, "SuccessOperator")
	delete(f, "SuccessRatio")
	delete(f, "RetryCount")
	delete(f, "RetryInterval")
	delete(f, "TaskArgument")
	delete(f, "ProgramIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskResponseParams struct {
	// 更新是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskResponseParams `json:"Response"`
}

func (r *ModifyTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUploadInfoRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 调用DescribeUploadInfo接口时返回的软件包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// COS返回上传结果（默认为0：成功，其他值表示失败）
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 程序包MD5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 程序包大小（单位字节）
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 程序包仓库id
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

type ModifyUploadInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 调用DescribeUploadInfo接口时返回的软件包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// COS返回上传结果（默认为0：成功，其他值表示失败）
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 程序包MD5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 程序包大小（单位字节）
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 程序包仓库类型
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 程序包仓库id
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`
}

func (r *ModifyUploadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUploadInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PkgId")
	delete(f, "Result")
	delete(f, "Md5")
	delete(f, "Size")
	delete(f, "RepositoryType")
	delete(f, "RepositoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUploadInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUploadInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUploadInfoResponseParams `json:"Response"`
}

func (r *ModifyUploadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUploadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorOverview struct {
	// 近24小时调用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationCountOfDay *string `json:"InvocationCountOfDay,omitempty" name:"InvocationCountOfDay"`

	// 总调用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationCount *string `json:"InvocationCount,omitempty" name:"InvocationCount"`

	// 近24小时调用错误数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCountOfDay *string `json:"ErrorCountOfDay,omitempty" name:"ErrorCountOfDay"`

	// 总调用错误数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCount *string `json:"ErrorCount,omitempty" name:"ErrorCount"`

	// 近24小时调用成功率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessRatioOfDay *string `json:"SuccessRatioOfDay,omitempty" name:"SuccessRatioOfDay"`

	// 总调用成功率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessRatio *string `json:"SuccessRatio,omitempty" name:"SuccessRatio"`
}

type MsApiArray struct {
	// API 请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 方法描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// API状态 0:离线 1:在线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type MsInstance struct {
	// 机器实例ID信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 机器实例名称信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 服务运行的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 机器实例内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 机器实例外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 机器可用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 服务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 机器TSF可用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 健康检查URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" name:"HealthCheckUrl"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 应用程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationPackageVersion *string `json:"ApplicationPackageVersion,omitempty" name:"ApplicationPackageVersion"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 服务状态，passing 在线，critical 离线
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`

	// 注册时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrationTime *int64 `json:"RegistrationTime,omitempty" name:"RegistrationTime"`

	// 上次心跳时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastHeartbeatTime *int64 `json:"LastHeartbeatTime,omitempty" name:"LastHeartbeatTime"`

	// 实例注册id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrationId *string `json:"RegistrationId,omitempty" name:"RegistrationId"`

	// 屏蔽状态，hidden 为屏蔽，unhidden 为未屏蔽
	// 注意：此字段可能返回 null，表示取不到有效值。
	HiddenStatus *string `json:"HiddenStatus,omitempty" name:"HiddenStatus"`
}

type MultiValue struct {
	// 数据点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type MultiValueDataPoints struct {
	// 多值数据点
	Points []*MultiValue `json:"Points,omitempty" name:"Points"`

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 多值数据点key列表，每个值表示当前数据点所在区域的下限
	PointKeys []*string `json:"PointKeys,omitempty" name:"PointKeys"`
}

type Namespace struct {
	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceCode *string `json:"NamespaceCode,omitempty" name:"NamespaceCode"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 命名空间描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 默认命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *string `json:"IsDefault,omitempty" name:"IsDefault"`

	// 命名空间状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" name:"NamespaceStatus"`

	// 删除标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群数组，仅携带集群ID，集群名称，集群类型等基础信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterList []*Cluster `json:"ClusterList,omitempty" name:"ClusterList"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceResourceType *string `json:"NamespaceResourceType,omitempty" name:"NamespaceResourceType"`

	// 命名空间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceType *string `json:"NamespaceType,omitempty" name:"NamespaceType"`

	// 是否开启高可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsHaEnable *string `json:"IsHaEnable,omitempty" name:"IsHaEnable"`

	// KubeInjectEnable值
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeInjectEnable *bool `json:"KubeInjectEnable,omitempty" name:"KubeInjectEnable"`
}

// Predefined struct for user
type OperateApplicationTcrBindingRequestParams struct {
	// bind 或 unbind
	Command *string `json:"Command,omitempty" name:"Command"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// TcrRepoInfo值
	TcrRepoInfo *TcrRepoInfo `json:"TcrRepoInfo,omitempty" name:"TcrRepoInfo"`
}

type OperateApplicationTcrBindingRequest struct {
	*tchttp.BaseRequest
	
	// bind 或 unbind
	Command *string `json:"Command,omitempty" name:"Command"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// TcrRepoInfo值
	TcrRepoInfo *TcrRepoInfo `json:"TcrRepoInfo,omitempty" name:"TcrRepoInfo"`
}

func (r *OperateApplicationTcrBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OperateApplicationTcrBindingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Command")
	delete(f, "ApplicationId")
	delete(f, "TcrRepoInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OperateApplicationTcrBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OperateApplicationTcrBindingResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OperateApplicationTcrBindingResponse struct {
	*tchttp.BaseResponse
	Response *OperateApplicationTcrBindingResponseParams `json:"Response"`
}

func (r *OperateApplicationTcrBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OperateApplicationTcrBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationInfo struct {
	// 初始化按钮的控制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Init *OperationInfoDetail `json:"Init,omitempty" name:"Init"`

	// 添加实例按钮的控制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddInstance *OperationInfoDetail `json:"AddInstance,omitempty" name:"AddInstance"`

	// 销毁机器的控制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destroy *OperationInfoDetail `json:"Destroy,omitempty" name:"Destroy"`
}

type OperationInfoDetail struct {
	// 不显示的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisabledReason *string `json:"DisabledReason,omitempty" name:"DisabledReason"`

	// 该按钮是否可点击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 是否显示该按钮
	// 注意：此字段可能返回 null，表示取不到有效值。
	Supported *bool `json:"Supported,omitempty" name:"Supported"`
}

type OverviewBasicResourceUsage struct {
	// 应用总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationCount *int64 `json:"ApplicationCount,omitempty" name:"ApplicationCount"`

	// 命名空间总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceCount *int64 `json:"NamespaceCount,omitempty" name:"NamespaceCount"`

	// 部署组个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupCount *int64 `json:"GroupCount,omitempty" name:"GroupCount"`

	// 程序包存储空间用量，单位字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageSpaceUsed *int64 `json:"PackageSpaceUsed,omitempty" name:"PackageSpaceUsed"`

	// 已注册实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsulInstanceCount *int64 `json:"ConsulInstanceCount,omitempty" name:"ConsulInstanceCount"`
}

type PagedProgram struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据集列表
	Content []*Program `json:"Content,omitempty" name:"Content"`
}

type PathRewrite struct {
	// 路径重写规则ID
	PathRewriteId *string `json:"PathRewriteId,omitempty" name:"PathRewriteId"`

	// 网关部署组ID
	GatewayGroupId *string `json:"GatewayGroupId,omitempty" name:"GatewayGroupId"`

	// 正则表达式
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// 替换的内容
	Replacement *string `json:"Replacement,omitempty" name:"Replacement"`

	// 是否屏蔽映射后路径，Y: 是 N: 否
	Blocked *string `json:"Blocked,omitempty" name:"Blocked"`

	// 规则顺序，越小优先级越高
	Order *int64 `json:"Order,omitempty" name:"Order"`
}

type PathRewriteCreateObject struct {
	// 网关部署组ID
	GatewayGroupId *string `json:"GatewayGroupId,omitempty" name:"GatewayGroupId"`

	// 正则表达式
	Regex *string `json:"Regex,omitempty" name:"Regex"`

	// 替换的内容
	Replacement *string `json:"Replacement,omitempty" name:"Replacement"`

	// 是否屏蔽映射后路径，Y: 是 N: 否
	Blocked *string `json:"Blocked,omitempty" name:"Blocked"`

	// 规则顺序，越小优先级越高
	Order *int64 `json:"Order,omitempty" name:"Order"`
}

type PathRewritePage struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 路径重写规则列表
	Content []*PathRewrite `json:"Content,omitempty" name:"Content"`
}

type PkgBind struct {
	// 应用id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type PkgInfo struct {
	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 程序包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 程序包类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 程序包描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgDesc *string `json:"PkgDesc,omitempty" name:"PkgDesc"`

	// 上传时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadTime *string `json:"UploadTime,omitempty" name:"UploadTime"`

	// 程序包MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 程序包状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgPubStatus *int64 `json:"PkgPubStatus,omitempty" name:"PkgPubStatus"`

	// 程序包关联关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgBindInfo []*PkgBind `json:"PkgBindInfo,omitempty" name:"PkgBindInfo"`
}

type PkgList struct {
	// 程序包总量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 程序包信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*PkgInfo `json:"Content,omitempty" name:"Content"`

	// 程序包仓库id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`

	// 程序包仓库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 程序包仓库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`
}

type Ports struct {
	// 服务端口
	TargetPort *uint64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// 端口协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type Program struct {
	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`

	// 数据集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramName *string `json:"ProgramName,omitempty" name:"ProgramName"`

	// 数据集描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDesc *string `json:"ProgramDesc,omitempty" name:"ProgramDesc"`

	// 删除标识，true: 可以删除; false: 不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 数据项列表，无值时返回空数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramItemList []*ProgramItem `json:"ProgramItemList,omitempty" name:"ProgramItemList"`
}

type ProgramItem struct {
	// 数据项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramItemId *string `json:"ProgramItemId,omitempty" name:"ProgramItemId"`

	// 资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *Resource `json:"Resource,omitempty" name:"Resource"`

	// 数据值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueList []*string `json:"ValueList,omitempty" name:"ValueList"`

	// 全选标识，true: 全选；false: 非全选
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAll *bool `json:"IsAll,omitempty" name:"IsAll"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 删除标识，true: 可删除；false: 不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramId *string `json:"ProgramId,omitempty" name:"ProgramId"`
}

type PropertyField struct {
	// 属性名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 属性描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ProtocolPort struct {
	// TCP UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 服务端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 容器端口
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// 主机端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePort *int64 `json:"NodePort,omitempty" name:"NodePort"`
}

type QuantileEntity struct {
	// 最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxValue *string `json:"MaxValue,omitempty" name:"MaxValue"`

	// 最小值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinValue *string `json:"MinValue,omitempty" name:"MinValue"`

	// 五分位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	FifthPositionValue *string `json:"FifthPositionValue,omitempty" name:"FifthPositionValue"`

	// 九分位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	NinthPositionValue *string `json:"NinthPositionValue,omitempty" name:"NinthPositionValue"`
}

// Predefined struct for user
type ReassociateBusinessLogConfigRequestParams struct {
	// 原关联日志配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 新关联日志配置ID
	NewConfigId *string `json:"NewConfigId,omitempty" name:"NewConfigId"`

	// TSF应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// TSF部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type ReassociateBusinessLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 原关联日志配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 新关联日志配置ID
	NewConfigId *string `json:"NewConfigId,omitempty" name:"NewConfigId"`

	// TSF应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// TSF部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ReassociateBusinessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReassociateBusinessLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "NewConfigId")
	delete(f, "ApplicationId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReassociateBusinessLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReassociateBusinessLogConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReassociateBusinessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *ReassociateBusinessLogConfigResponseParams `json:"Response"`
}

func (r *ReassociateBusinessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReassociateBusinessLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RedoTaskBatchRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

type RedoTaskBatchRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *RedoTaskBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedoTaskBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RedoTaskBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RedoTaskBatchResponseParams struct {
	// 批次ID
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RedoTaskBatchResponse struct {
	*tchttp.BaseResponse
	Response *RedoTaskBatchResponseParams `json:"Response"`
}

func (r *RedoTaskBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedoTaskBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RedoTaskExecuteRequestParams struct {
	// 任务批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务执行ID
	ExecuteId *string `json:"ExecuteId,omitempty" name:"ExecuteId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type RedoTaskExecuteRequest struct {
	*tchttp.BaseRequest
	
	// 任务批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务执行ID
	ExecuteId *string `json:"ExecuteId,omitempty" name:"ExecuteId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RedoTaskExecuteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedoTaskExecuteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	delete(f, "ExecuteId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RedoTaskExecuteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RedoTaskExecuteResponseParams struct {
	// 成功失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RedoTaskExecuteResponse struct {
	*tchttp.BaseResponse
	Response *RedoTaskExecuteResponseParams `json:"Response"`
}

func (r *RedoTaskExecuteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedoTaskExecuteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RedoTaskFlowBatchRequestParams struct {
	// 工作流批次 ID
	FlowBatchId *string `json:"FlowBatchId,omitempty" name:"FlowBatchId"`
}

type RedoTaskFlowBatchRequest struct {
	*tchttp.BaseRequest
	
	// 工作流批次 ID
	FlowBatchId *string `json:"FlowBatchId,omitempty" name:"FlowBatchId"`
}

func (r *RedoTaskFlowBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedoTaskFlowBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowBatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RedoTaskFlowBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RedoTaskFlowBatchResponseParams struct {
	// 工作流批次历史 ID
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RedoTaskFlowBatchResponse struct {
	*tchttp.BaseResponse
	Response *RedoTaskFlowBatchResponseParams `json:"Response"`
}

func (r *RedoTaskFlowBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedoTaskFlowBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RedoTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type RedoTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RedoTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedoTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RedoTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RedoTaskResponseParams struct {
	// 操作成功or失败
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RedoTaskResponse struct {
	*tchttp.BaseResponse
	Response *RedoTaskResponseParams `json:"Response"`
}

func (r *RedoTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RedoTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseApiGroupRequestParams struct {
	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type ReleaseApiGroupRequest struct {
	*tchttp.BaseRequest
	
	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ReleaseApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseApiGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseApiGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseApiGroupResponseParams struct {
	// 成功/失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseApiGroupResponseParams `json:"Response"`
}

func (r *ReleaseApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseApiGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseConfigRequestParams struct {
	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

type ReleaseConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *ReleaseConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "GroupId")
	delete(f, "ReleaseDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseConfigResponseParams struct {
	// true：发布成功；false：发布失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseConfigResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseConfigResponseParams `json:"Response"`
}

func (r *ReleaseConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseFileConfigRequestParams struct {
	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

type ReleaseFileConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *ReleaseFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseFileConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "GroupId")
	delete(f, "ReleaseDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseFileConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseFileConfigResponseParams struct {
	// 发布结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseFileConfigResponseParams `json:"Response"`
}

func (r *ReleaseFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseFileConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleasePublicConfigRequestParams struct {
	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

type ReleasePublicConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *ReleasePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleasePublicConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "NamespaceId")
	delete(f, "ReleaseDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleasePublicConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleasePublicConfigResponseParams struct {
	// true：发布成功；false：发布失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleasePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *ReleasePublicConfigResponseParams `json:"Response"`
}

func (r *ReleasePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleasePublicConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveInstancesRequestParams struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机 ID 列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

type RemoveInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机 ID 列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

func (r *RemoveInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveInstancesResponseParams struct {
	// 集群移除机器是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RemoveInstancesResponseParams `json:"Response"`
}

func (r *RemoveInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RepositoryInfo struct {
	// 仓库ID
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`

	// 仓库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 仓库类型（默认仓库：default，私有仓库：private）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 仓库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepositoryDesc *string `json:"RepositoryDesc,omitempty" name:"RepositoryDesc"`

	// 仓库是否正在被使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUsed *bool `json:"IsUsed,omitempty" name:"IsUsed"`

	// 仓库创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 仓库桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 仓库桶所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`

	// 仓库目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Directory *string `json:"Directory,omitempty" name:"Directory"`
}

type RepositoryList struct {
	// 仓库总量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*RepositoryInfo `json:"Content,omitempty" name:"Content"`
}

type Resource struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceCode *string `json:"ResourceCode,omitempty" name:"ResourceCode"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源所属产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCode *string `json:"ServiceCode,omitempty" name:"ServiceCode"`

	// 选取资源使用的Action
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceAction *string `json:"ResourceAction,omitempty" name:"ResourceAction"`

	// 资源数据查询的ID字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdField *string `json:"IdField,omitempty" name:"IdField"`

	// 资源数据查询的名称字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameField *string `json:"NameField,omitempty" name:"NameField"`

	// 资源数据查询的ID过滤字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelectIdsField *string `json:"SelectIdsField,omitempty" name:"SelectIdsField"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *int64 `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 删除标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 资源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceDesc *string `json:"ResourceDesc,omitempty" name:"ResourceDesc"`

	// 是否可以选择全部
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanSelectAll *bool `json:"CanSelectAll,omitempty" name:"CanSelectAll"`

	// 资源数据查询的模糊查询字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchWordField *string `json:"SearchWordField,omitempty" name:"SearchWordField"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type ResourceFieldRef struct {
	// k8s 的 Resource
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *string `json:"Resource,omitempty" name:"Resource"`
}

// Predefined struct for user
type RevocationConfigRequestParams struct {
	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

type RevocationConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

func (r *RevocationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevocationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigReleaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevocationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevocationConfigResponseParams struct {
	// true：回滚成功；false：回滚失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RevocationConfigResponse struct {
	*tchttp.BaseResponse
	Response *RevocationConfigResponseParams `json:"Response"`
}

func (r *RevocationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevocationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevocationPublicConfigRequestParams struct {
	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

type RevocationPublicConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

func (r *RevocationPublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevocationPublicConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigReleaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevocationPublicConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevocationPublicConfigResponseParams struct {
	// true：撤销成功；false：撤销失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RevocationPublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *RevocationPublicConfigResponseParams `json:"Response"`
}

func (r *RevocationPublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevocationPublicConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevokeFileConfigRequestParams struct {
	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

type RevokeFileConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

func (r *RevokeFileConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeFileConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigReleaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevokeFileConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevokeFileConfigResponseParams struct {
	// 撤回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RevokeFileConfigResponse struct {
	*tchttp.BaseResponse
	Response *RevokeFileConfigResponseParams `json:"Response"`
}

func (r *RevokeFileConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeFileConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackConfigRequestParams struct {
	// 配置项发布历史ID
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 回滚描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

type RollbackConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置项发布历史ID
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 回滚描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *RollbackConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigReleaseLogId")
	delete(f, "ReleaseDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackConfigResponseParams struct {
	// true：回滚成功；false：回滚失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RollbackConfigResponse struct {
	*tchttp.BaseResponse
	Response *RollbackConfigResponseParams `json:"Response"`
}

func (r *RollbackConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScalableRule struct {
	// RuleId值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Name值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// ExpandVmCountLimit值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpandVmCountLimit *int64 `json:"ExpandVmCountLimit,omitempty" name:"ExpandVmCountLimit"`

	// ShrinkVmCountLimit值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShrinkVmCountLimit *int64 `json:"ShrinkVmCountLimit,omitempty" name:"ShrinkVmCountLimit"`

	// GroupCount值
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupCount *int64 `json:"GroupCount,omitempty" name:"GroupCount"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否关闭指标伸缩, 默认0, 0:打开指标伸缩 1:关闭指标伸缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisableMetricAS *uint64 `json:"DisableMetricAS,omitempty" name:"DisableMetricAS"`

	// 开启定时伸缩规则, 默认0, 0:关闭定时伸缩 1:开启定时伸缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableCronAS *uint64 `json:"EnableCronAS,omitempty" name:"EnableCronAS"`
}

type SchedulingStrategy struct {
	// NONE：不使用调度策略；CROSS_AZ：跨可用区部署
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

// Predefined struct for user
type SearchBusinessLogRequestParams struct {
	// 日志配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 机器实例ID，不传表示全部实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 请求偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量，取值范围[1, 200]，默认值为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序规则，默认值"time"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords"`

	// 部署组ID列表，不传表示全部部署组
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 检索类型，取值"LUCENE", "REGEXP", "NORMAL"
	SearchWordType *string `json:"SearchWordType,omitempty" name:"SearchWordType"`

	// 批量请求类型，取值"page"或"scroll"
	BatchType *string `json:"BatchType,omitempty" name:"BatchType"`

	// 游标ID
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`
}

type SearchBusinessLogRequest struct {
	*tchttp.BaseRequest
	
	// 日志配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 机器实例ID，不传表示全部实例
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 请求偏移量，取值范围大于等于0，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 单页请求配置数量，取值范围[1, 200]，默认值为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序规则，默认值"time"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，取值"asc"或"desc"，默认值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords"`

	// 部署组ID列表，不传表示全部部署组
	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`

	// 检索类型，取值"LUCENE", "REGEXP", "NORMAL"
	SearchWordType *string `json:"SearchWordType,omitempty" name:"SearchWordType"`

	// 批量请求类型，取值"page"或"scroll"
	BatchType *string `json:"BatchType,omitempty" name:"BatchType"`

	// 游标ID
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`
}

func (r *SearchBusinessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchBusinessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "InstanceIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "SearchWords")
	delete(f, "GroupIds")
	delete(f, "SearchWordType")
	delete(f, "BatchType")
	delete(f, "ScrollId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchBusinessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchBusinessLogResponseParams struct {
	// 业务日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageBusinessLogV2 `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchBusinessLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchBusinessLogResponseParams `json:"Response"`
}

func (r *SearchBusinessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchBusinessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchStdoutLogRequestParams struct {
	// 机器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 单页请求配置数量，取值范围[1, 500]，默认值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 请求偏移量，取值范围大于等于0，默认值为
	// 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序规则，默认值"time"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，取值"asc"或"desc"，默认
	// 值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 检索类型，取值"LUCENE", "REGEXP",
	// "NORMAL"
	SearchWordType *string `json:"SearchWordType,omitempty" name:"SearchWordType"`

	// 批量请求类型，取值"page"或"scroll"，默认
	// 值"page"
	BatchType *string `json:"BatchType,omitempty" name:"BatchType"`

	// 游标ID
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`
}

type SearchStdoutLogRequest struct {
	*tchttp.BaseRequest
	
	// 机器实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 单页请求配置数量，取值范围[1, 500]，默认值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 检索关键词
	SearchWords []*string `json:"SearchWords,omitempty" name:"SearchWords"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 请求偏移量，取值范围大于等于0，默认值为
	// 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序规则，默认值"time"
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，取值"asc"或"desc"，默认
	// 值"desc"
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 检索类型，取值"LUCENE", "REGEXP",
	// "NORMAL"
	SearchWordType *string `json:"SearchWordType,omitempty" name:"SearchWordType"`

	// 批量请求类型，取值"page"或"scroll"，默认
	// 值"page"
	BatchType *string `json:"BatchType,omitempty" name:"BatchType"`

	// 游标ID
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`
}

func (r *SearchStdoutLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchStdoutLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SearchWords")
	delete(f, "StartTime")
	delete(f, "GroupId")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "SearchWordType")
	delete(f, "BatchType")
	delete(f, "ScrollId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchStdoutLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchStdoutLogResponseParams struct {
	// 标准输出日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TsfPageStdoutLogV2 `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchStdoutLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchStdoutLogResponseParams `json:"Response"`
}

func (r *SearchStdoutLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchStdoutLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceConfig struct {
	// 服务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 端口信息列表
	Ports []*Ports `json:"Ports,omitempty" name:"Ports"`

	// 健康检查配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheck *HealthCheckConfig `json:"HealthCheck,omitempty" name:"HealthCheck"`
}

type ServiceSetting struct {
	// 0:公网, 1:集群内访问, 2：NodePort, 3: VPC 内网访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 容器端口映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 是否创建 k8s service，默认为 false
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisableService *bool `json:"DisableService,omitempty" name:"DisableService"`

	// service 是否为 headless 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadlessService *bool `json:"HeadlessService,omitempty" name:"HeadlessService"`

	// 当为 true 且 DisableService 也为 true 时，会删除之前创建的 service，请小心使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowDeleteService *bool `json:"AllowDeleteService,omitempty" name:"AllowDeleteService"`

	// 开启SessionAffinity，true为开启，false为不开启，默认为false
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenSessionAffinity *bool `json:"OpenSessionAffinity,omitempty" name:"OpenSessionAffinity"`

	// SessionAffinity会话时间，默认10800
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionAffinityTimeoutSeconds *int64 `json:"SessionAffinityTimeoutSeconds,omitempty" name:"SessionAffinityTimeoutSeconds"`
}

type ServiceStatisticsResult struct {
	// 请求模版路径:type为接口时返回，服务时不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法:type为接口时返回，服务时不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 微服务Id
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 请求数
	RequestCount *uint64 `json:"RequestCount,omitempty" name:"RequestCount"`

	// 请求错误率，不带百分号
	ErrorRate *float64 `json:"ErrorRate,omitempty" name:"ErrorRate"`

	// 平均响应耗时ms
	AvgTimeConsuming *float64 `json:"AvgTimeConsuming,omitempty" name:"AvgTimeConsuming"`

	// 响应耗时曲线
	MetricDataCurves []*MetricDataCurve `json:"MetricDataCurves,omitempty" name:"MetricDataCurves"`

	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例name
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 部署组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组name
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署组类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 部署组是否存在
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupExist *int64 `json:"GroupExist,omitempty" name:"GroupExist"`

	// 实例是否存在，仅限cvm
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceExist *int64 `json:"InstanceExist,omitempty" name:"InstanceExist"`

	// 应用id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// cpu使用率
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuPercent *int64 `json:"CpuPercent,omitempty" name:"CpuPercent"`

	// 已用堆大小,单位KB
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeapUsed *int64 `json:"HeapUsed,omitempty" name:"HeapUsed"`

	// 数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Script值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Script *string `json:"Script,omitempty" name:"Script"`

	// 数据库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbType *string `json:"DbType,omitempty" name:"DbType"`

	// Apdex值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Apdex *float64 `json:"Apdex,omitempty" name:"Apdex"`

	// Qps值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *float64 `json:"Qps,omitempty" name:"Qps"`

	// 实例在线数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceOnlineCount *int64 `json:"InstanceOnlineCount,omitempty" name:"InstanceOnlineCount"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTotalCount *int64 `json:"InstanceTotalCount,omitempty" name:"InstanceTotalCount"`

	// normal/error
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// normal/warn/error
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorRateLevel *string `json:"ErrorRateLevel,omitempty" name:"ErrorRateLevel"`

	// normal/warn/error
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgTimeConsumingLevel *string `json:"AvgTimeConsumingLevel,omitempty" name:"AvgTimeConsumingLevel"`

	// normal/warn/error
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApdexLevel *string `json:"ApdexLevel,omitempty" name:"ApdexLevel"`
}

type ServiceStatisticsResults struct {
	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ServiceStatisticsResult `json:"Content,omitempty" name:"Content"`

	// 条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ShardArgument struct {
	// 分片参数 KEY，整形, 范围 [1,1000]
	ShardKey *uint64 `json:"ShardKey,omitempty" name:"ShardKey"`

	// 分片参数 VALUE
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardValue *string `json:"ShardValue,omitempty" name:"ShardValue"`
}

// Predefined struct for user
type ShrinkGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type ShrinkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ShrinkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShrinkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ShrinkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ShrinkGroupResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TaskId `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ShrinkGroupResponse struct {
	*tchttp.BaseResponse
	Response *ShrinkGroupResponseParams `json:"Response"`
}

func (r *ShrinkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShrinkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ShrinkInstancesRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 下线机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

type ShrinkInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 下线机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

func (r *ShrinkInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShrinkInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ShrinkInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ShrinkInstancesResponseParams struct {
	// 任务ID
	Result *TaskId `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ShrinkInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ShrinkInstancesResponseParams `json:"Response"`
}

func (r *ShrinkInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShrinkInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SimpleApplication struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// ApplicationDesc
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// ProgLang
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgLang *string `json:"ProgLang,omitempty" name:"ProgLang"`

	// ApplicationResourceType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// UpdateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// ApigatewayServiceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApigatewayServiceId *string `json:"ApigatewayServiceId,omitempty" name:"ApigatewayServiceId"`

	// ApplicationRuntimeType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationRuntimeType *string `json:"ApplicationRuntimeType,omitempty" name:"ApplicationRuntimeType"`
}

type SimpleGroup struct {
	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 启动参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 部署组资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 应用微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppMicroServiceType *string `json:"AppMicroServiceType,omitempty" name:"AppMicroServiceType"`
}

type SimpleKafkaDeliveryConfig struct {
	// 配置项id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
}

// Predefined struct for user
type StartContainerGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type StartContainerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StartContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartContainerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartContainerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartContainerGroupResponseParams struct {
	// 启动操作是否成功。
	// true：启动成功
	// false：启动失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *StartContainerGroupResponseParams `json:"Response"`
}

func (r *StartContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartContainerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type StartGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StartGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartGroupResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TaskId `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartGroupResponse struct {
	*tchttp.BaseResponse
	Response *StartGroupResponseParams `json:"Response"`
}

func (r *StartGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StdoutLogV2 struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 日志时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIp *string `json:"InstanceIp,omitempty" name:"InstanceIp"`
}

// Predefined struct for user
type StopContainerGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type StopContainerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StopContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopContainerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopContainerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopContainerGroupResponseParams struct {
	// 停止操作是否成功。
	// true：停止成功
	// false：停止失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *StopContainerGroupResponseParams `json:"Response"`
}

func (r *StopContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopContainerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopGroupRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type StopGroupRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StopGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopGroupResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TaskId `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopGroupResponse struct {
	*tchttp.BaseResponse
	Response *StopGroupResponseParams `json:"Response"`
}

func (r *StopGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTaskBatchRequestParams struct {
	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 参数ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type StopTaskBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 参数ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopTaskBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTaskBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopTaskBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTaskBatchResponseParams struct {
	// 操作成功 or 失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopTaskBatchResponse struct {
	*tchttp.BaseResponse
	Response *StopTaskBatchResponseParams `json:"Response"`
}

func (r *StopTaskBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTaskBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTaskExecuteRequestParams struct {
	// 任务执行ID
	ExecuteId *string `json:"ExecuteId,omitempty" name:"ExecuteId"`

	// 任务批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type StopTaskExecuteRequest struct {
	*tchttp.BaseRequest
	
	// 任务执行ID
	ExecuteId *string `json:"ExecuteId,omitempty" name:"ExecuteId"`

	// 任务批次ID
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopTaskExecuteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTaskExecuteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExecuteId")
	delete(f, "BatchId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopTaskExecuteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTaskExecuteResponseParams struct {
	// 操作成功 or 失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopTaskExecuteResponse struct {
	*tchttp.BaseResponse
	Response *StopTaskExecuteResponseParams `json:"Response"`
}

func (r *StopTaskExecuteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTaskExecuteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskFlowEdge struct {
	// 节点 ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 子节点 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildNodeId *string `json:"ChildNodeId,omitempty" name:"ChildNodeId"`

	// 是否核心任务,Y/N
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreNode *string `json:"CoreNode,omitempty" name:"CoreNode"`

	// 边类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeType *string `json:"EdgeType,omitempty" name:"EdgeType"`

	// 任务节点类型
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// X轴坐标位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PositionX *string `json:"PositionX,omitempty" name:"PositionX"`

	// Y轴坐标位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PositionY *string `json:"PositionY,omitempty" name:"PositionY"`

	// 图 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GraphId *string `json:"GraphId,omitempty" name:"GraphId"`

	// 工作流 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务历史ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskLogId *string `json:"TaskLogId,omitempty" name:"TaskLogId"`
}

type TaskFlowLastBatchState struct {
	// 批次ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowBatchId *string `json:"FlowBatchId,omitempty" name:"FlowBatchId"`

	// 批次历史ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowBatchLogId *string `json:"FlowBatchLogId,omitempty" name:"FlowBatchLogId"`

	// 状态,WAITING/SUCCESS/FAILED/RUNNING/TERMINATING
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`
}

type TaskId struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type TaskLastExecuteStatus struct {
	// 批次ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 运行状态，RUNNING/SUCCESS/FAIL/HALF/TERMINATED
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`

	// 批次历史ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchLogId *string `json:"BatchLogId,omitempty" name:"BatchLogId"`
}

type TaskRecord struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 执行类型
	ExecuteType *string `json:"ExecuteType,omitempty" name:"ExecuteType"`

	// 任务内容，长度限制65535字节
	TaskContent *string `json:"TaskContent,omitempty" name:"TaskContent"`

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 超时时间
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryCount *int64 `json:"RetryCount,omitempty" name:"RetryCount"`

	// 重试间隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryInterval *int64 `json:"RetryInterval,omitempty" name:"RetryInterval"`

	// 触发规则
	TaskRule *TaskRule `json:"TaskRule,omitempty" name:"TaskRule"`

	// 是否启用任务,ENABLED/DISABLED
	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 判断任务成功的操作符
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessOperator *string `json:"SuccessOperator,omitempty" name:"SuccessOperator"`

	// 判断任务成功的阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessRatio *int64 `json:"SuccessRatio,omitempty" name:"SuccessRatio"`

	// 分片数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardCount *int64 `json:"ShardCount,omitempty" name:"ShardCount"`

	// 高级设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceSettings *AdvanceSettings `json:"AdvanceSettings,omitempty" name:"AdvanceSettings"`

	// 分片参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardArguments []*ShardArgument `json:"ShardArguments,omitempty" name:"ShardArguments"`

	// 所属工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongFlowIds []*string `json:"BelongFlowIds,omitempty" name:"BelongFlowIds"`

	// 任务历史ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskLogId *string `json:"TaskLogId,omitempty" name:"TaskLogId"`

	// 触发类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *string `json:"TriggerType,omitempty" name:"TriggerType"`

	// 任务参数，长度限制10000个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskArgument *string `json:"TaskArgument,omitempty" name:"TaskArgument"`
}

type TaskRecordPage struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务记录列表
	Content []*TaskRecord `json:"Content,omitempty" name:"Content"`
}

type TaskRule struct {
	// 触发规则类型, Cron/Repeat
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// Cron类型规则，cron表达式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expression *string `json:"Expression,omitempty" name:"Expression"`

	// 时间间隔， 单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepeatInterval *uint64 `json:"RepeatInterval,omitempty" name:"RepeatInterval"`
}

type TcrRepoInfo struct {
	// 地域（填数字）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 仓库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

// Predefined struct for user
type TerminateTaskFlowBatchRequestParams struct {
	// 工作流批次 ID
	FlowBatchId *string `json:"FlowBatchId,omitempty" name:"FlowBatchId"`
}

type TerminateTaskFlowBatchRequest struct {
	*tchttp.BaseRequest
	
	// 工作流批次 ID
	FlowBatchId *string `json:"FlowBatchId,omitempty" name:"FlowBatchId"`
}

func (r *TerminateTaskFlowBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTaskFlowBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowBatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateTaskFlowBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateTaskFlowBatchResponseParams struct {
	// 是否停止成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateTaskFlowBatchResponse struct {
	*tchttp.BaseResponse
	Response *TerminateTaskFlowBatchResponseParams `json:"Response"`
}

func (r *TerminateTaskFlowBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateTaskFlowBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ThreadPicture struct {
	// 总线程数
	ThreadCount []*CurvePoint `json:"ThreadCount,omitempty" name:"ThreadCount"`

	// 活跃线程数
	ThreadActive []*CurvePoint `json:"ThreadActive,omitempty" name:"ThreadActive"`

	// 守护线程数
	DeamonThreadCount []*CurvePoint `json:"DeamonThreadCount,omitempty" name:"DeamonThreadCount"`
}

type TsfApiListResponse struct {
	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*MsApiArray `json:"Content,omitempty" name:"Content"`
}

type TsfPageApiDetailInfo struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API 信息列表
	Content []*ApiDetailInfo `json:"Content,omitempty" name:"Content"`
}

type TsfPageApiGroupInfo struct {
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API分组信息
	Content []*ApiGroupInfo `json:"Content,omitempty" name:"Content"`
}

type TsfPageApplication struct {
	// 应用总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 应用信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ApplicationForPage `json:"Content,omitempty" name:"Content"`
}

type TsfPageBusinessLogConfig struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 业务日志配置项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*BusinessLogConfig `json:"Content,omitempty" name:"Content"`
}

type TsfPageBusinessLogV2 struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 业务日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*BusinessLogV2 `json:"Content,omitempty" name:"Content"`

	// 游标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 查询状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type TsfPageCluster struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Cluster `json:"Content,omitempty" name:"Content"`
}

type TsfPageClusterV2 struct {
	// 集群总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ClusterV2 `json:"Content,omitempty" name:"Content"`
}

type TsfPageConfig struct {
	// TsfPageConfig
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置项列表
	Content []*Config `json:"Content,omitempty" name:"Content"`
}

type TsfPageConfigRelease struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置项发布信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ConfigRelease `json:"Content,omitempty" name:"Content"`
}

type TsfPageConfigReleaseLog struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置项发布日志数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ConfigReleaseLog `json:"Content,omitempty" name:"Content"`
}

type TsfPageContainerEvent struct {
	// 返回个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// events 数组
	Content []*ContainerEvent `json:"Content,omitempty" name:"Content"`
}

type TsfPageDimension struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 维度
	Content []*string `json:"Content,omitempty" name:"Content"`
}

type TsfPageFileConfig struct {
	// 总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 文件配置数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*FileConfig `json:"Content,omitempty" name:"Content"`
}

type TsfPageFileConfigRelease struct {
	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*FileConfigRelease `json:"Content,omitempty" name:"Content"`
}

type TsfPageGatewayDeployGroup struct {
	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 记录实体列表
	Content []*GatewayDeployGroup `json:"Content,omitempty" name:"Content"`
}

type TsfPageGatewayPlugin struct {
	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 记录实体列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*GatewayPlugin `json:"Content,omitempty" name:"Content"`
}

type TsfPageInstance struct {
	// 机器实例总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 机器实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Instance `json:"Content,omitempty" name:"Content"`
}

type TsfPageMicroservice struct {
	// 微服务总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 微服务列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Microservice `json:"Content,omitempty" name:"Content"`
}

type TsfPageMsInstance struct {
	// 微服务实例总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 微服务实例列表内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*MsInstance `json:"Content,omitempty" name:"Content"`
}

type TsfPageNamespace struct {
	// 命名空间总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 命名空间列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Namespace `json:"Content,omitempty" name:"Content"`
}

type TsfPageSimpleApplication struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 简单应用列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*SimpleApplication `json:"Content,omitempty" name:"Content"`
}

type TsfPageSimpleGroup struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 简单部署组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*SimpleGroup `json:"Content,omitempty" name:"Content"`
}

type TsfPageStdoutLogV2 struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 标准输出日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*StdoutLogV2 `json:"Content,omitempty" name:"Content"`

	// 游标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScrollId *string `json:"ScrollId,omitempty" name:"ScrollId"`

	// 查询状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type TsfPageUnitNamespace struct {
	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 记录实体列表
	Content []*UnitNamespace `json:"Content,omitempty" name:"Content"`
}

type TsfPageUnitRule struct {
	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 记录实体列表
	Content []*UnitRule `json:"Content,omitempty" name:"Content"`
}

type TsfPageVmGroup struct {
	// 虚拟机部署组总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 虚拟机部署组列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*VmGroupSimple `json:"Content,omitempty" name:"Content"`
}

// Predefined struct for user
type UnbindApiGroupRequestParams struct {
	// 分组网关id列表
	GroupGatewayList []*GatewayGroupIds `json:"GroupGatewayList,omitempty" name:"GroupGatewayList"`
}

type UnbindApiGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分组网关id列表
	GroupGatewayList []*GatewayGroupIds `json:"GroupGatewayList,omitempty" name:"GroupGatewayList"`
}

func (r *UnbindApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindApiGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupGatewayList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindApiGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindApiGroupResponseParams struct {
	// 返回结果，成功失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *UnbindApiGroupResponseParams `json:"Response"`
}

func (r *UnbindApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindApiGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnitNamespace struct {
	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间Name
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 单元化命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 网关实体ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
}

type UnitRule struct {
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 网关实体ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" name:"GatewayInstanceId"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 使用状态：enabled/disabled
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 规则项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitRuleItemList []*UnitRuleItem `json:"UnitRuleItemList,omitempty" name:"UnitRuleItemList"`

	// CreatedTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// UpdatedTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
}

type UnitRuleItem struct {
	// 逻辑关系：AND/OR
	Relationship *string `json:"Relationship,omitempty" name:"Relationship"`

	// 目的地命名空间ID
	DestNamespaceId *string `json:"DestNamespaceId,omitempty" name:"DestNamespaceId"`

	// 目的地命名空间名称
	DestNamespaceName *string `json:"DestNamespaceName,omitempty" name:"DestNamespaceName"`

	// 规则项名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 单元化规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitRuleId *string `json:"UnitRuleId,omitempty" name:"UnitRuleId"`

	// 规则顺序，越小优先级越高：默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitRuleTagList []*UnitRuleTag `json:"UnitRuleTagList,omitempty" name:"UnitRuleTagList"`
}

type UnitRuleTag struct {
	// 标签类型 : U(用户标签)
	TagType *string `json:"TagType,omitempty" name:"TagType"`

	// 标签名
	TagField *string `json:"TagField,omitempty" name:"TagField"`

	// 操作符:IN/NOT_IN/EQUAL/NOT_EQUAL/REGEX
	TagOperator *string `json:"TagOperator,omitempty" name:"TagOperator"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 单元化规则项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitRuleItemId *string `json:"UnitRuleItemId,omitempty" name:"UnitRuleItemId"`

	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`
}

// Predefined struct for user
type UpdateApiGroupRequestParams struct {
	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Api 分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Api 分组描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 鉴权类型
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 分组上下文
	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`

	// 命名空间参数key值
	NamespaceNameKey *string `json:"NamespaceNameKey,omitempty" name:"NamespaceNameKey"`

	// 微服务名参数key值
	ServiceNameKey *string `json:"ServiceNameKey,omitempty" name:"ServiceNameKey"`

	// 命名空间参数位置，path，header或query，默认是path
	NamespaceNameKeyPosition *string `json:"NamespaceNameKeyPosition,omitempty" name:"NamespaceNameKeyPosition"`

	// 微服务名参数位置，path，header或query，默认是path
	ServiceNameKeyPosition *string `json:"ServiceNameKeyPosition,omitempty" name:"ServiceNameKeyPosition"`
}

type UpdateApiGroupRequest struct {
	*tchttp.BaseRequest
	
	// Api 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Api 分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// Api 分组描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 鉴权类型
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 分组上下文
	GroupContext *string `json:"GroupContext,omitempty" name:"GroupContext"`

	// 命名空间参数key值
	NamespaceNameKey *string `json:"NamespaceNameKey,omitempty" name:"NamespaceNameKey"`

	// 微服务名参数key值
	ServiceNameKey *string `json:"ServiceNameKey,omitempty" name:"ServiceNameKey"`

	// 命名空间参数位置，path，header或query，默认是path
	NamespaceNameKeyPosition *string `json:"NamespaceNameKeyPosition,omitempty" name:"NamespaceNameKeyPosition"`

	// 微服务名参数位置，path，header或query，默认是path
	ServiceNameKeyPosition *string `json:"ServiceNameKeyPosition,omitempty" name:"ServiceNameKeyPosition"`
}

func (r *UpdateApiGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "Description")
	delete(f, "AuthType")
	delete(f, "GroupContext")
	delete(f, "NamespaceNameKey")
	delete(f, "ServiceNameKey")
	delete(f, "NamespaceNameKeyPosition")
	delete(f, "ServiceNameKeyPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApiGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiGroupResponseParams struct {
	// 返回结果，true: 成功, false: 失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApiGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApiGroupResponseParams `json:"Response"`
}

func (r *UpdateApiGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiRateLimitRuleRequestParams struct {
	// 限流规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 开启/禁用，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// qps值，开启限流规则时，必填
	MaxQps *int64 `json:"MaxQps,omitempty" name:"MaxQps"`
}

type UpdateApiRateLimitRuleRequest struct {
	*tchttp.BaseRequest
	
	// 限流规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 开启/禁用，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// qps值，开启限流规则时，必填
	MaxQps *int64 `json:"MaxQps,omitempty" name:"MaxQps"`
}

func (r *UpdateApiRateLimitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiRateLimitRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "UsableStatus")
	delete(f, "MaxQps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApiRateLimitRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiRateLimitRuleResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApiRateLimitRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApiRateLimitRuleResponseParams `json:"Response"`
}

func (r *UpdateApiRateLimitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiRateLimitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiRateLimitRulesRequestParams struct {
	// API ID 列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// 开启/禁用，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// QPS值。开启限流规则时，必填
	MaxQps *int64 `json:"MaxQps,omitempty" name:"MaxQps"`
}

type UpdateApiRateLimitRulesRequest struct {
	*tchttp.BaseRequest
	
	// API ID 列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// 开启/禁用，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// QPS值。开启限流规则时，必填
	MaxQps *int64 `json:"MaxQps,omitempty" name:"MaxQps"`
}

func (r *UpdateApiRateLimitRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiRateLimitRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiIds")
	delete(f, "UsableStatus")
	delete(f, "MaxQps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApiRateLimitRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiRateLimitRulesResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApiRateLimitRulesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApiRateLimitRulesResponseParams `json:"Response"`
}

func (r *UpdateApiRateLimitRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiRateLimitRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiTimeoutsRequestParams struct {
	// API ID 列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// 开启/禁用，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// 超时时间，单位毫秒，开启API超时时，必填
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

type UpdateApiTimeoutsRequest struct {
	*tchttp.BaseRequest
	
	// API ID 列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// 开启/禁用，enabled/disabled
	UsableStatus *string `json:"UsableStatus,omitempty" name:"UsableStatus"`

	// 超时时间，单位毫秒，开启API超时时，必填
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}

func (r *UpdateApiTimeoutsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiTimeoutsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiIds")
	delete(f, "UsableStatus")
	delete(f, "Timeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApiTimeoutsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiTimeoutsResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApiTimeoutsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApiTimeoutsResponseParams `json:"Response"`
}

func (r *UpdateApiTimeoutsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiTimeoutsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConfigTemplateRequestParams struct {
	// 配置模板id
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`

	// 配置模板名称
	ConfigTemplateName *string `json:"ConfigTemplateName,omitempty" name:"ConfigTemplateName"`

	// 配置模板对应的微服务框架
	ConfigTemplateType *string `json:"ConfigTemplateType,omitempty" name:"ConfigTemplateType"`

	// 配置模板数据
	ConfigTemplateValue *string `json:"ConfigTemplateValue,omitempty" name:"ConfigTemplateValue"`

	// 配置模板描述
	ConfigTemplateDesc *string `json:"ConfigTemplateDesc,omitempty" name:"ConfigTemplateDesc"`
}

type UpdateConfigTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 配置模板id
	ConfigTemplateId *string `json:"ConfigTemplateId,omitempty" name:"ConfigTemplateId"`

	// 配置模板名称
	ConfigTemplateName *string `json:"ConfigTemplateName,omitempty" name:"ConfigTemplateName"`

	// 配置模板对应的微服务框架
	ConfigTemplateType *string `json:"ConfigTemplateType,omitempty" name:"ConfigTemplateType"`

	// 配置模板数据
	ConfigTemplateValue *string `json:"ConfigTemplateValue,omitempty" name:"ConfigTemplateValue"`

	// 配置模板描述
	ConfigTemplateDesc *string `json:"ConfigTemplateDesc,omitempty" name:"ConfigTemplateDesc"`
}

func (r *UpdateConfigTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConfigTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigTemplateId")
	delete(f, "ConfigTemplateName")
	delete(f, "ConfigTemplateType")
	delete(f, "ConfigTemplateValue")
	delete(f, "ConfigTemplateDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateConfigTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConfigTemplateResponseParams struct {
	// 结果true：成功；false：失败；
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateConfigTemplateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateConfigTemplateResponseParams `json:"Response"`
}

func (r *UpdateConfigTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConfigTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGatewayApiRequestParams struct {
	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API 路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// Api 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 请求映射
	PathMapping *string `json:"PathMapping,omitempty" name:"PathMapping"`

	// api所在服务host
	Host *string `json:"Host,omitempty" name:"Host"`

	// api描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
}

type UpdateGatewayApiRequest struct {
	*tchttp.BaseRequest
	
	// API ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API 路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// Api 请求方法
	Method *string `json:"Method,omitempty" name:"Method"`

	// 请求映射
	PathMapping *string `json:"PathMapping,omitempty" name:"PathMapping"`

	// api所在服务host
	Host *string `json:"Host,omitempty" name:"Host"`

	// api描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *UpdateGatewayApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGatewayApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiId")
	delete(f, "Path")
	delete(f, "Method")
	delete(f, "PathMapping")
	delete(f, "Host")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGatewayApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGatewayApiResponseParams struct {
	// 返回结果，成功失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateGatewayApiResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGatewayApiResponseParams `json:"Response"`
}

func (r *UpdateGatewayApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGatewayApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateHealthCheckSettingsRequestParams struct {
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 是否能使健康检查
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`

	// 健康检查配置
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`
}

type UpdateHealthCheckSettingsRequest struct {
	*tchttp.BaseRequest
	
	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 是否能使健康检查
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`

	// 健康检查配置
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`
}

func (r *UpdateHealthCheckSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateHealthCheckSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "EnableHealthCheck")
	delete(f, "HealthCheckSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateHealthCheckSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateHealthCheckSettingsResponseParams struct {
	// 更新健康检查配置操作是否成功。
	// true：操作成功。
	// false：操作失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateHealthCheckSettingsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateHealthCheckSettingsResponseParams `json:"Response"`
}

func (r *UpdateHealthCheckSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateHealthCheckSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRepositoryRequestParams struct {
	// 仓库ID
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`

	// 仓库描述
	RepositoryDesc *string `json:"RepositoryDesc,omitempty" name:"RepositoryDesc"`
}

type UpdateRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 仓库ID
	RepositoryId *string `json:"RepositoryId,omitempty" name:"RepositoryId"`

	// 仓库描述
	RepositoryDesc *string `json:"RepositoryDesc,omitempty" name:"RepositoryDesc"`
}

func (r *UpdateRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepositoryId")
	delete(f, "RepositoryDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRepositoryResponseParams struct {
	// 更新仓库是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRepositoryResponseParams `json:"Response"`
}

func (r *UpdateRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUnitRuleRequestParams struct {
	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则项列表
	UnitRuleItemList []*UnitRuleItem `json:"UnitRuleItemList,omitempty" name:"UnitRuleItemList"`
}

type UpdateUnitRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 规则项列表
	UnitRuleItemList []*UnitRuleItem `json:"UnitRuleItemList,omitempty" name:"UnitRuleItemList"`
}

func (r *UpdateUnitRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUnitRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "UnitRuleItemList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUnitRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUnitRuleResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateUnitRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUnitRuleResponseParams `json:"Response"`
}

func (r *UpdateUnitRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUnitRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValueFrom struct {
	// k8s env 的 FieldRef
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldRef *FieldRef `json:"FieldRef,omitempty" name:"FieldRef"`

	// k8s env 的 ResourceFieldRef
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceFieldRef *ResourceFieldRef `json:"ResourceFieldRef,omitempty" name:"ResourceFieldRef"`
}

type VmGroup struct {
	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *string `json:"GroupStatus,omitempty" name:"GroupStatus"`

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 程序包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 程序包版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 部署组机器数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 部署组运行中机器数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 部署组启动参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 部署组创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 部署组更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 部署组停止机器数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffInstanceCount *int64 `json:"OffInstanceCount,omitempty" name:"OffInstanceCount"`

	// 部署组描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 部署组资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 部署组更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 部署应用描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployDesc *string `json:"DeployDesc,omitempty" name:"DeployDesc"`

	// 滚动发布的更新方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateType *uint64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 发布是否启用beta批次
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployBetaEnable *bool `json:"DeployBetaEnable,omitempty" name:"DeployBetaEnable"`

	// 滚动发布的批次比例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployBatch []*float64 `json:"DeployBatch,omitempty" name:"DeployBatch"`

	// 滚动发布的批次执行方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployExeMode *string `json:"DeployExeMode,omitempty" name:"DeployExeMode"`

	// 滚动发布的每个批次的等待时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployWaitTime *uint64 `json:"DeployWaitTime,omitempty" name:"DeployWaitTime"`

	// 是否开启了健康检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitempty" name:"EnableHealthCheck"`

	// 健康检查配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`

	// 程序包类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 启动脚本 base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`

	// 停止脚本 base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopScript *string `json:"StopScript,omitempty" name:"StopScript"`

	// 部署组备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// javaagent信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentProfileList []*AgentProfile `json:"AgentProfileList,omitempty" name:"AgentProfileList"`

	// 预热属性配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmupSetting *WarmupSetting `json:"WarmupSetting,omitempty" name:"WarmupSetting"`

	// Envoy网关配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayConfig *GatewayConfig `json:"GatewayConfig,omitempty" name:"GatewayConfig"`
}

type VmGroupOther struct {
	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 程序包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 部署组实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 部署组运行中实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 部署组中停止实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffInstanceCount *int64 `json:"OffInstanceCount,omitempty" name:"OffInstanceCount"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *string `json:"GroupStatus,omitempty" name:"GroupStatus"`

	// 服务配置信息是否匹配
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNotEqualServiceConfig *bool `json:"IsNotEqualServiceConfig,omitempty" name:"IsNotEqualServiceConfig"`

	// HealthCheckSettings
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckSettings *HealthCheckSettings `json:"HealthCheckSettings,omitempty" name:"HealthCheckSettings"`
}

type VmGroupSimple struct {
	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 部署组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 部署组更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组启动参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 部署组创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 部署组资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 部署组更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 部署应用描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployDesc *string `json:"DeployDesc,omitempty" name:"DeployDesc"`

	// 部署组备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type VolumeInfo struct {
	// 数据卷类型
	VolumeType *string `json:"VolumeType,omitempty" name:"VolumeType"`

	// 数据卷名称
	VolumeName *string `json:"VolumeName,omitempty" name:"VolumeName"`

	// 数据卷配置
	VolumeConfig *string `json:"VolumeConfig,omitempty" name:"VolumeConfig"`
}

type VolumeMountInfo struct {
	// 挂载数据卷名称
	VolumeMountName *string `json:"VolumeMountName,omitempty" name:"VolumeMountName"`

	// 挂载路径
	VolumeMountPath *string `json:"VolumeMountPath,omitempty" name:"VolumeMountPath"`

	// 挂载子路径
	VolumeMountSubPath *string `json:"VolumeMountSubPath,omitempty" name:"VolumeMountSubPath"`

	// 读写，1：读 2：读写
	ReadOrWrite *string `json:"ReadOrWrite,omitempty" name:"ReadOrWrite"`
}

type WarmupSetting struct {
	// 是否开启预热
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 预热时间
	WarmupTime *uint64 `json:"WarmupTime,omitempty" name:"WarmupTime"`

	// 预热曲率，取值 1~5
	Curvature *uint64 `json:"Curvature,omitempty" name:"Curvature"`

	// 是否开启预热保护，在开启保护的情况下，超过 50% 的节点处于预热中，则会中止预热
	EnabledProtection *bool `json:"EnabledProtection,omitempty" name:"EnabledProtection"`
}