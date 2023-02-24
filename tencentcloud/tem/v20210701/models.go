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

package v20210701

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Autoscaler struct {
	// 弹性伸缩最小实例数
	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// 弹性伸缩最大实例数
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// 指标弹性伸缩策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// 定时弹性伸缩策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronHorizontalAutoscaler []*CronHorizontalAutoscaler `json:"CronHorizontalAutoscaler,omitempty" name:"CronHorizontalAutoscaler"`

	// 弹性伸缩ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`

	// 弹性伸缩名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoscalerName *string `json:"AutoscalerName,omitempty" name:"AutoscalerName"`

	// 弹性伸缩描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// 启用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableDate *string `json:"EnableDate,omitempty" name:"EnableDate"`

	// 是否启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type ConfigData struct {
	// 配置名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 关联的服务列表
	RelatedApplications []*TemService `json:"RelatedApplications,omitempty" name:"RelatedApplications"`

	// 配置条目
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

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

// Predefined struct for user
type CreateApplicationAutoscalerRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略
	Autoscaler *Autoscaler `json:"Autoscaler,omitempty" name:"Autoscaler"`
}

type CreateApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略
	Autoscaler *Autoscaler `json:"Autoscaler,omitempty" name:"Autoscaler"`
}

func (r *CreateApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "Autoscaler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationAutoscalerResponseParams struct {
	// 弹性伸缩策略组合ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *CreateApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationRequestParams struct {
	// 应用名
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否使用默认镜像服务 1-是，0-否
	UseDefaultImageService *int64 `json:"UseDefaultImageService,omitempty" name:"UseDefaultImageService"`

	// 如果是绑定仓库，绑定的仓库类型，0-个人版，1-企业版
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`

	// 企业版镜像服务的实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 绑定镜像服务器地址
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// 绑定镜像仓库名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 应用所在子网
	SubnetList []*string `json:"SubnetList,omitempty" name:"SubnetList"`

	// 编程语言 
	// - JAVA
	// - OTHER
	CodingLanguage *string `json:"CodingLanguage,omitempty" name:"CodingLanguage"`

	// 部署方式 
	// - IMAGE
	// - JAR
	// - WAR
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// 是否开启 Java 应用的 APM 自动上报功能，1 表示启用；0 表示关闭
	EnableTracing *int64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// 使用默认镜像服务额外参数
	UseDefaultImageServiceParameters *UseDefaultRepoParameters `json:"UseDefaultImageServiceParameters,omitempty" name:"UseDefaultImageServiceParameters"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用名
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否使用默认镜像服务 1-是，0-否
	UseDefaultImageService *int64 `json:"UseDefaultImageService,omitempty" name:"UseDefaultImageService"`

	// 如果是绑定仓库，绑定的仓库类型，0-个人版，1-企业版
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`

	// 企业版镜像服务的实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 绑定镜像服务器地址
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// 绑定镜像仓库名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 应用所在子网
	SubnetList []*string `json:"SubnetList,omitempty" name:"SubnetList"`

	// 编程语言 
	// - JAVA
	// - OTHER
	CodingLanguage *string `json:"CodingLanguage,omitempty" name:"CodingLanguage"`

	// 部署方式 
	// - IMAGE
	// - JAR
	// - WAR
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// 是否开启 Java 应用的 APM 自动上报功能，1 表示启用；0 表示关闭
	EnableTracing *int64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// 使用默认镜像服务额外参数
	UseDefaultImageServiceParameters *UseDefaultRepoParameters `json:"UseDefaultImageServiceParameters,omitempty" name:"UseDefaultImageServiceParameters"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "Description")
	delete(f, "UseDefaultImageService")
	delete(f, "RepoType")
	delete(f, "InstanceId")
	delete(f, "RepoServer")
	delete(f, "RepoName")
	delete(f, "SourceChannel")
	delete(f, "SubnetList")
	delete(f, "CodingLanguage")
	delete(f, "DeployMode")
	delete(f, "EnableTracing")
	delete(f, "UseDefaultImageServiceParameters")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationResponseParams struct {
	// 服务code
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
type CreateApplicationServiceRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 访问方式详情
	Service *ServicePortMapping `json:"Service,omitempty" name:"Service"`
}

type CreateApplicationServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 访问方式详情
	Service *ServicePortMapping `json:"Service,omitempty" name:"Service"`
}

func (r *CreateApplicationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "Service")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationServiceResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationServiceResponseParams `json:"Response"`
}

func (r *CreateApplicationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigDataRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 配置信息
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

type CreateConfigDataRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 配置信息
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

func (r *CreateConfigDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "SourceChannel")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigDataResponseParams struct {
	// 创建是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateConfigDataResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigDataResponseParams `json:"Response"`
}

func (r *CreateConfigDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosTokenRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// optType 1上传  2查询
	OptType *int64 `json:"OptType,omitempty" name:"OptType"`

	// 来源 channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 充当deployVersion入参
	TimeVersion *string `json:"TimeVersion,omitempty" name:"TimeVersion"`
}

type CreateCosTokenRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// optType 1上传  2查询
	OptType *int64 `json:"OptType,omitempty" name:"OptType"`

	// 来源 channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 充当deployVersion入参
	TimeVersion *string `json:"TimeVersion,omitempty" name:"TimeVersion"`
}

func (r *CreateCosTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PkgName")
	delete(f, "OptType")
	delete(f, "SourceChannel")
	delete(f, "TimeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosTokenResponseParams struct {
	// 成功时为CosToken对象，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CosToken `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCosTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosTokenResponseParams `json:"Response"`
}

func (r *CreateCosTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvironmentRequestParams struct {
	// 环境名称
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 环境描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// K8s version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 是否开启tsw服务
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 环境类型：test、pre、prod
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 创建环境的region
	CreateRegion *string `json:"CreateRegion,omitempty" name:"CreateRegion"`

	// 是否创建私有网络
	SetupVpc *bool `json:"SetupVpc,omitempty" name:"SetupVpc"`

	// 是否创建 Prometheus 实例
	SetupPrometheus *bool `json:"SetupPrometheus,omitempty" name:"SetupPrometheus"`

	// prometheus 实例 id
	PrometheusId *string `json:"PrometheusId,omitempty" name:"PrometheusId"`

	// apm id
	ApmId *string `json:"ApmId,omitempty" name:"ApmId"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 环境名称
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 环境描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// K8s version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 是否开启tsw服务
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 环境类型：test、pre、prod
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 创建环境的region
	CreateRegion *string `json:"CreateRegion,omitempty" name:"CreateRegion"`

	// 是否创建私有网络
	SetupVpc *bool `json:"SetupVpc,omitempty" name:"SetupVpc"`

	// 是否创建 Prometheus 实例
	SetupPrometheus *bool `json:"SetupPrometheus,omitempty" name:"SetupPrometheus"`

	// prometheus 实例 id
	PrometheusId *string `json:"PrometheusId,omitempty" name:"PrometheusId"`

	// apm id
	ApmId *string `json:"ApmId,omitempty" name:"ApmId"`
}

func (r *CreateEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentName")
	delete(f, "Description")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "K8sVersion")
	delete(f, "SourceChannel")
	delete(f, "EnableTswTraceService")
	delete(f, "Tags")
	delete(f, "EnvType")
	delete(f, "CreateRegion")
	delete(f, "SetupVpc")
	delete(f, "SetupPrometheus")
	delete(f, "PrometheusId")
	delete(f, "ApmId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvironmentResponseParams struct {
	// 成功时为环境ID，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvironmentResponseParams `json:"Response"`
}

func (r *CreateEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogConfigRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收集类型，container_stdout 为标准输出；container_file 为文件；
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 日志集 ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题 ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志提取模式，minimalist_log 为单行全文；multiline_log 为多行全文；json_log 为 json格式；fullregex_log 为单行正则；multiline_fullregex_log 为多行正则
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 首行正则表达式，当LogType=multiline_log 时生效
	BeginningRegex *string `json:"BeginningRegex,omitempty" name:"BeginningRegex"`

	// 收集文件目录，当 InputType=container_file 时生效
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// 收集文件名模式，当 InputType=container_file 时生效
	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`

	// 导出规则
	ExtractRule *LogConfigExtractRule `json:"ExtractRule,omitempty" name:"ExtractRule"`
}

type CreateLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收集类型，container_stdout 为标准输出；container_file 为文件；
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 日志集 ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题 ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志提取模式，minimalist_log 为单行全文；multiline_log 为多行全文；json_log 为 json格式；fullregex_log 为单行正则；multiline_fullregex_log 为多行正则
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 首行正则表达式，当LogType=multiline_log 时生效
	BeginningRegex *string `json:"BeginningRegex,omitempty" name:"BeginningRegex"`

	// 收集文件目录，当 InputType=container_file 时生效
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// 收集文件名模式，当 InputType=container_file 时生效
	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`

	// 导出规则
	ExtractRule *LogConfigExtractRule `json:"ExtractRule,omitempty" name:"ExtractRule"`
}

func (r *CreateLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "InputType")
	delete(f, "ApplicationId")
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "LogType")
	delete(f, "BeginningRegex")
	delete(f, "LogPath")
	delete(f, "FilePattern")
	delete(f, "ExtractRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogConfigResponseParams struct {
	// 创建是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogConfigResponseParams `json:"Response"`
}

func (r *CreateLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceRequestParams struct {
	// 环境 Id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 资源类型，目前支持文件系统：CFS；日志服务：CLS；注册中心：TSE_SRE
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源 Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 资源来源，目前支持：existing，已有资源；creating，自动创建
	ResourceFrom *string `json:"ResourceFrom,omitempty" name:"ResourceFrom"`

	// 设置 resource 的额外配置
	ResourceConfig *string `json:"ResourceConfig,omitempty" name:"ResourceConfig"`
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest
	
	// 环境 Id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 资源类型，目前支持文件系统：CFS；日志服务：CLS；注册中心：TSE_SRE
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源 Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 资源来源，目前支持：existing，已有资源；creating，自动创建
	ResourceFrom *string `json:"ResourceFrom,omitempty" name:"ResourceFrom"`

	// 设置 resource 的额外配置
	ResourceConfig *string `json:"ResourceConfig,omitempty" name:"ResourceConfig"`
}

func (r *CreateResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "SourceChannel")
	delete(f, "ResourceFrom")
	delete(f, "ResourceConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceResponseParams struct {
	// 成功与否
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceResponseParams `json:"Response"`
}

func (r *CreateResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronHorizontalAutoscaler struct {
	// 定时伸缩策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 策略周期
	// * * *，三个范围，第一个是天，第二个是月，第三个是周，中间用空格隔开
	// 例子：
	// * * * （每天）
	// * * 0-3 （每周日到周三）
	// 1,11,21 * *（每个月1号，11号，21号）
	Period *string `json:"Period,omitempty" name:"Period"`

	// 定时伸缩策略明细
	Schedules []*CronHorizontalAutoscalerSchedule `json:"Schedules,omitempty" name:"Schedules"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 策略优先级，值越大优先级越高，0为最小值
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type CronHorizontalAutoscalerSchedule struct {
	// 触发事件，小时分钟，用:分割
	// 例如
	// 00:00（零点零分触发）
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// 目标实例数（不大于50）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetReplicas *int64 `json:"TargetReplicas,omitempty" name:"TargetReplicas"`
}

// Predefined struct for user
type DeleteApplicationAutoscalerRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

type DeleteApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

func (r *DeleteApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "AutoscalerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationAutoscalerResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *DeleteApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationRequestParams struct {
	// 服务Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 当服务没有任何运行版本时，是否删除此服务
	DeleteApplicationIfNoRunningVersion *bool `json:"DeleteApplicationIfNoRunningVersion,omitempty" name:"DeleteApplicationIfNoRunningVersion"`
}

type DeleteApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 服务Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 当服务没有任何运行版本时，是否删除此服务
	DeleteApplicationIfNoRunningVersion *bool `json:"DeleteApplicationIfNoRunningVersion,omitempty" name:"DeleteApplicationIfNoRunningVersion"`
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
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "DeleteApplicationIfNoRunningVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationResponseParams struct {
	// 返回结果
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
type DeleteApplicationServiceRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 访问方式服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type DeleteApplicationServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 访问方式服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DeleteApplicationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	delete(f, "ServiceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationServiceResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationServiceResponseParams `json:"Response"`
}

func (r *DeleteApplicationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIngressRequestParams struct {
	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境 namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// ingress 规则名
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DeleteIngressRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境 namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// ingress 规则名
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DeleteIngressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIngressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterNamespace")
	delete(f, "IngressName")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIngressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIngressResponseParams struct {
	// 是否删除成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIngressResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIngressResponseParams `json:"Response"`
}

func (r *DeleteIngressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployApplicationRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 初始化 pod 数
	InitPodNum *uint64 `json:"InitPodNum,omitempty" name:"InitPodNum"`

	// cpu规格
	CpuSpec *float64 `json:"CpuSpec,omitempty" name:"CpuSpec"`

	// 内存规格
	MemorySpec *float64 `json:"MemorySpec,omitempty" name:"MemorySpec"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 镜像仓库
	ImgRepo *string `json:"ImgRepo,omitempty" name:"ImgRepo"`

	// 版本描述信息
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// 启动参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 弹性伸缩配置（已废弃，请使用HorizontalAutoscaler设置弹性策略）
	EsInfo *EsInfo `json:"EsInfo,omitempty" name:"EsInfo"`

	// 环境变量配置
	EnvConf []*Pair `json:"EnvConf,omitempty" name:"EnvConf"`

	// 日志配置
	LogConfs []*string `json:"LogConfs,omitempty" name:"LogConfs"`

	// 数据卷配置
	StorageConfs []*StorageConf `json:"StorageConfs,omitempty" name:"StorageConfs"`

	// 数据卷挂载配置
	StorageMountConfs []*StorageMountConf `json:"StorageMountConfs,omitempty" name:"StorageMountConfs"`

	// 部署类型。
	// - JAR：通过 jar 包部署
	// - WAR：通过 war 包部署
	// - IMAGE：通过镜像部署
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// 部署类型为 IMAGE 时，该参数表示镜像 tag。
	// 部署类型为 JAR/WAR 时，该参数表示包版本号。
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// 包名。使用 JAR 包或者 WAR 包部署的时候必填。
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// JDK 版本。
	// - KONA:8：使用 kona jdk 8。
	// - OPEN:8：使用 open jdk 8。
	// - KONA:11：使用 kona jdk 11。
	// - OPEN:11：使用 open jdk 11。
	JdkVersion *string `json:"JdkVersion,omitempty" name:"JdkVersion"`

	// 安全组ID s
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 日志输出配置
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitempty" name:"LogOutputConf"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 版本描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 镜像命令
	ImageCommand *string `json:"ImageCommand,omitempty" name:"ImageCommand"`

	// 镜像命令参数
	ImageArgs []*string `json:"ImageArgs,omitempty" name:"ImageArgs"`

	// 是否添加默认注册中心配置
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitempty" name:"UseRegistryDefaultConfig"`

	// 挂载配置信息
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitempty" name:"SettingConfs"`

	// 应用访问设置
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// 要回滚到的历史版本id
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 启动后执行的脚本
	PostStart *string `json:"PostStart,omitempty" name:"PostStart"`

	// 停止前执行的脚本
	PreStop *string `json:"PreStop,omitempty" name:"PreStop"`

	// 存活探针配置
	Liveness *HealthCheckConfig `json:"Liveness,omitempty" name:"Liveness"`

	// 就绪探针配置
	Readiness *HealthCheckConfig `json:"Readiness,omitempty" name:"Readiness"`

	// 分批发布策略配置
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// 弹性策略（已弃用，请使用弹性伸缩策略组合相关接口）
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// 定时弹性策略（已弃用，请使用弹性伸缩策略组合相关接口）
	CronHorizontalAutoscaler []*CronHorizontalAutoscaler `json:"CronHorizontalAutoscaler,omitempty" name:"CronHorizontalAutoscaler"`

	// 是否启用log，1为启用，0为不启用
	LogEnable *int64 `json:"LogEnable,omitempty" name:"LogEnable"`

	// （除开镜像配置）配置是否修改
	ConfEdited *bool `json:"ConfEdited,omitempty" name:"ConfEdited"`

	// 是否开启应用加速
	SpeedUp *bool `json:"SpeedUp,omitempty" name:"SpeedUp"`

	// 启动探针配置
	StartupProbe *HealthCheckConfig `json:"StartupProbe,omitempty" name:"StartupProbe"`

	// 操作系统版本；
	// 当选择openjdk时，可选参数：
	// - ALPINE
	// - CENTOS
	// 当选择konajdk时，可选参数：
	// - ALPINE
	// - TENCENTOS
	OsFlavour *string `json:"OsFlavour,omitempty" name:"OsFlavour"`

	// metrics业务指标监控配置
	EnablePrometheusConf *EnablePrometheusConf `json:"EnablePrometheusConf,omitempty" name:"EnablePrometheusConf"`

	// 1：开始自动apm采集（skywalking）；
	// 0：关闭apm采集；
	EnableTracing *int64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// 1：开始自动metrics采集（open-telemetry）；
	// 0：关闭metrics采集；
	EnableMetrics *int64 `json:"EnableMetrics,omitempty" name:"EnableMetrics"`

	// 镜像部署时，选择的tcr实例id
	TcrInstanceId *string `json:"TcrInstanceId,omitempty" name:"TcrInstanceId"`

	// 镜像部署时，选择的镜像服务器地址
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// 镜像部署时，仓库类型：0：个人仓库；1：企业版；2：公共仓库；3：tem托管仓库；4：demo仓库
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`
}

type DeployApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 初始化 pod 数
	InitPodNum *uint64 `json:"InitPodNum,omitempty" name:"InitPodNum"`

	// cpu规格
	CpuSpec *float64 `json:"CpuSpec,omitempty" name:"CpuSpec"`

	// 内存规格
	MemorySpec *float64 `json:"MemorySpec,omitempty" name:"MemorySpec"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 镜像仓库
	ImgRepo *string `json:"ImgRepo,omitempty" name:"ImgRepo"`

	// 版本描述信息
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// 启动参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 弹性伸缩配置（已废弃，请使用HorizontalAutoscaler设置弹性策略）
	EsInfo *EsInfo `json:"EsInfo,omitempty" name:"EsInfo"`

	// 环境变量配置
	EnvConf []*Pair `json:"EnvConf,omitempty" name:"EnvConf"`

	// 日志配置
	LogConfs []*string `json:"LogConfs,omitempty" name:"LogConfs"`

	// 数据卷配置
	StorageConfs []*StorageConf `json:"StorageConfs,omitempty" name:"StorageConfs"`

	// 数据卷挂载配置
	StorageMountConfs []*StorageMountConf `json:"StorageMountConfs,omitempty" name:"StorageMountConfs"`

	// 部署类型。
	// - JAR：通过 jar 包部署
	// - WAR：通过 war 包部署
	// - IMAGE：通过镜像部署
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// 部署类型为 IMAGE 时，该参数表示镜像 tag。
	// 部署类型为 JAR/WAR 时，该参数表示包版本号。
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// 包名。使用 JAR 包或者 WAR 包部署的时候必填。
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// JDK 版本。
	// - KONA:8：使用 kona jdk 8。
	// - OPEN:8：使用 open jdk 8。
	// - KONA:11：使用 kona jdk 11。
	// - OPEN:11：使用 open jdk 11。
	JdkVersion *string `json:"JdkVersion,omitempty" name:"JdkVersion"`

	// 安全组ID s
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 日志输出配置
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitempty" name:"LogOutputConf"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 版本描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 镜像命令
	ImageCommand *string `json:"ImageCommand,omitempty" name:"ImageCommand"`

	// 镜像命令参数
	ImageArgs []*string `json:"ImageArgs,omitempty" name:"ImageArgs"`

	// 是否添加默认注册中心配置
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitempty" name:"UseRegistryDefaultConfig"`

	// 挂载配置信息
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitempty" name:"SettingConfs"`

	// 应用访问设置
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// 要回滚到的历史版本id
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 启动后执行的脚本
	PostStart *string `json:"PostStart,omitempty" name:"PostStart"`

	// 停止前执行的脚本
	PreStop *string `json:"PreStop,omitempty" name:"PreStop"`

	// 存活探针配置
	Liveness *HealthCheckConfig `json:"Liveness,omitempty" name:"Liveness"`

	// 就绪探针配置
	Readiness *HealthCheckConfig `json:"Readiness,omitempty" name:"Readiness"`

	// 分批发布策略配置
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// 弹性策略（已弃用，请使用弹性伸缩策略组合相关接口）
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// 定时弹性策略（已弃用，请使用弹性伸缩策略组合相关接口）
	CronHorizontalAutoscaler []*CronHorizontalAutoscaler `json:"CronHorizontalAutoscaler,omitempty" name:"CronHorizontalAutoscaler"`

	// 是否启用log，1为启用，0为不启用
	LogEnable *int64 `json:"LogEnable,omitempty" name:"LogEnable"`

	// （除开镜像配置）配置是否修改
	ConfEdited *bool `json:"ConfEdited,omitempty" name:"ConfEdited"`

	// 是否开启应用加速
	SpeedUp *bool `json:"SpeedUp,omitempty" name:"SpeedUp"`

	// 启动探针配置
	StartupProbe *HealthCheckConfig `json:"StartupProbe,omitempty" name:"StartupProbe"`

	// 操作系统版本；
	// 当选择openjdk时，可选参数：
	// - ALPINE
	// - CENTOS
	// 当选择konajdk时，可选参数：
	// - ALPINE
	// - TENCENTOS
	OsFlavour *string `json:"OsFlavour,omitempty" name:"OsFlavour"`

	// metrics业务指标监控配置
	EnablePrometheusConf *EnablePrometheusConf `json:"EnablePrometheusConf,omitempty" name:"EnablePrometheusConf"`

	// 1：开始自动apm采集（skywalking）；
	// 0：关闭apm采集；
	EnableTracing *int64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// 1：开始自动metrics采集（open-telemetry）；
	// 0：关闭metrics采集；
	EnableMetrics *int64 `json:"EnableMetrics,omitempty" name:"EnableMetrics"`

	// 镜像部署时，选择的tcr实例id
	TcrInstanceId *string `json:"TcrInstanceId,omitempty" name:"TcrInstanceId"`

	// 镜像部署时，选择的镜像服务器地址
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// 镜像部署时，仓库类型：0：个人仓库；1：企业版；2：公共仓库；3：tem托管仓库；4：demo仓库
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`
}

func (r *DeployApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "InitPodNum")
	delete(f, "CpuSpec")
	delete(f, "MemorySpec")
	delete(f, "EnvironmentId")
	delete(f, "ImgRepo")
	delete(f, "VersionDesc")
	delete(f, "JvmOpts")
	delete(f, "EsInfo")
	delete(f, "EnvConf")
	delete(f, "LogConfs")
	delete(f, "StorageConfs")
	delete(f, "StorageMountConfs")
	delete(f, "DeployMode")
	delete(f, "DeployVersion")
	delete(f, "PkgName")
	delete(f, "JdkVersion")
	delete(f, "SecurityGroupIds")
	delete(f, "LogOutputConf")
	delete(f, "SourceChannel")
	delete(f, "Description")
	delete(f, "ImageCommand")
	delete(f, "ImageArgs")
	delete(f, "UseRegistryDefaultConfig")
	delete(f, "SettingConfs")
	delete(f, "Service")
	delete(f, "VersionId")
	delete(f, "PostStart")
	delete(f, "PreStop")
	delete(f, "Liveness")
	delete(f, "Readiness")
	delete(f, "DeployStrategyConf")
	delete(f, "HorizontalAutoscaler")
	delete(f, "CronHorizontalAutoscaler")
	delete(f, "LogEnable")
	delete(f, "ConfEdited")
	delete(f, "SpeedUp")
	delete(f, "StartupProbe")
	delete(f, "OsFlavour")
	delete(f, "EnablePrometheusConf")
	delete(f, "EnableTracing")
	delete(f, "EnableMetrics")
	delete(f, "TcrInstanceId")
	delete(f, "RepoServer")
	delete(f, "RepoType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployApplicationResponseParams struct {
	// 版本ID（前端可忽略）
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployApplicationResponse struct {
	*tchttp.BaseResponse
	Response *DeployApplicationResponseParams `json:"Response"`
}

func (r *DeployApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployServiceBatchDetail struct {
	// 旧实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldPodList *DeployServicePodDetail `json:"OldPodList,omitempty" name:"OldPodList"`

	// 新实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewPodList *DeployServicePodDetail `json:"NewPodList,omitempty" name:"NewPodList"`

	// 当前批次状态："WaitForTimeExceed", "WaitForResume", "Deploying", "Finish", "NotStart"
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchStatus *string `json:"BatchStatus,omitempty" name:"BatchStatus"`

	// 该批次预计旧实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodNum *int64 `json:"PodNum,omitempty" name:"PodNum"`

	// 批次id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchIndex *int64 `json:"BatchIndex,omitempty" name:"BatchIndex"`

	// 旧实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldPods []*DeployServicePodDetail `json:"OldPods,omitempty" name:"OldPods"`

	// 新实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewPods []*DeployServicePodDetail `json:"NewPods,omitempty" name:"NewPods"`

	// =0：手动确认批次；>0：下一批次开始时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextBatchStartTime *int64 `json:"NextBatchStartTime,omitempty" name:"NextBatchStartTime"`
}

type DeployServicePodDetail struct {
	// pod Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// pod状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodStatus []*string `json:"PodStatus,omitempty" name:"PodStatus"`

	// pod版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodVersion *string `json:"PodVersion,omitempty" name:"PodVersion"`

	// pod创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// pod所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// webshell地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Webshell *string `json:"Webshell,omitempty" name:"Webshell"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type DeployStrategyConf struct {
	// 总分批数
	TotalBatchCount *int64 `json:"TotalBatchCount,omitempty" name:"TotalBatchCount"`

	// beta分批实例数
	BetaBatchNum *int64 `json:"BetaBatchNum,omitempty" name:"BetaBatchNum"`

	// 分批策略：0-全自动，1-全手动，2-beta分批，beta批一定是手动的，3-首次发布
	DeployStrategyType *int64 `json:"DeployStrategyType,omitempty" name:"DeployStrategyType"`

	// 每批暂停间隔
	BatchInterval *int64 `json:"BatchInterval,omitempty" name:"BatchInterval"`

	// 最小可用实例数
	MinAvailable *int64 `json:"MinAvailable,omitempty" name:"MinAvailable"`

	// 是否强制发布
	Force *bool `json:"Force,omitempty" name:"Force"`
}

// Predefined struct for user
type DescribeApplicationAutoscalerListRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeApplicationAutoscalerListRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeApplicationAutoscalerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationAutoscalerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationAutoscalerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationAutoscalerListResponseParams struct {
	// 弹性伸缩策略组合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*Autoscaler `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationAutoscalerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationAutoscalerListResponseParams `json:"Response"`
}

func (r *DescribeApplicationAutoscalerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationAutoscalerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationInfoRequestParams struct {
	// 服务版本ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type DescribeApplicationInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务版本ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *DescribeApplicationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationInfoResponseParams struct {
	// 返回结果
	Result *TemServiceVersionInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationInfoResponseParams `json:"Response"`
}

func (r *DescribeApplicationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationPodsRequestParams struct {
	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 单页条数，默认值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页下标，默认值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 实例状态 
	// - Running 
	// - Pending 
	// - Error
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例名字
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeApplicationPodsRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 单页条数，默认值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页下标，默认值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 实例状态 
	// - Running 
	// - Pending 
	// - Error
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例名字
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeApplicationPodsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationPodsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ApplicationId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "PodName")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationPodsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationPodsResponseParams struct {
	// 返回结果
	Result *DescribeRunPodPage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationPodsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationPodsResponseParams `json:"Response"`
}

func (r *DescribeApplicationPodsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationServiceListRequestParams struct {
	// namespace id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 服务ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// xx
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeApplicationServiceListRequest struct {
	*tchttp.BaseRequest
	
	// namespace id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 服务ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// xx
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeApplicationServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationServiceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationServiceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationServiceListResponseParams struct {
	// 应用 EKS Service 列表
	Result *EksService `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationServiceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationServiceListResponseParams `json:"Response"`
}

func (r *DescribeApplicationServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsRequestParams struct {
	// 命名空间ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 查询过滤器
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	SortInfo *SortType `json:"SortInfo,omitempty" name:"SortInfo"`
}

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 查询过滤器
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	SortInfo *SortType `json:"SortInfo,omitempty" name:"SortInfo"`
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
	delete(f, "EnvironmentId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SourceChannel")
	delete(f, "ApplicationId")
	delete(f, "Keyword")
	delete(f, "Filters")
	delete(f, "SortInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsResponseParams struct {
	// 返回结果
	Result *ServicePage `json:"Result,omitempty" name:"Result"`

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
type DescribeApplicationsStatusRequestParams struct {
	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type DescribeApplicationsStatusRequest struct {
	*tchttp.BaseRequest
	
	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *DescribeApplicationsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsStatusResponseParams struct {
	// 返回结果
	Result []*ServiceVersionBrief `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationsStatusResponseParams `json:"Response"`
}

func (r *DescribeApplicationsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigDataListPage struct {
	// 记录
	Records []*ConfigData `json:"Records,omitempty" name:"Records"`

	// 分页游标，用以查询下一页
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`

	// 剩余数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainingCount *int64 `json:"RemainingCount,omitempty" name:"RemainingCount"`
}

// Predefined struct for user
type DescribeConfigDataListRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 查询游标
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`

	// 分页 limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeConfigDataListRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 查询游标
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`

	// 分页 limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigDataListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDataListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "ContinueToken")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigDataListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigDataListResponseParams struct {
	// 配置列表
	Result *DescribeConfigDataListPage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigDataListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigDataListResponseParams `json:"Response"`
}

func (r *DescribeConfigDataListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDataListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigDataRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeConfigDataRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeConfigDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigDataResponseParams struct {
	// 配置
	Result *ConfigData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigDataResponseParams `json:"Response"`
}

func (r *DescribeConfigDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeployApplicationDetailRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 版本部署id
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

type DescribeDeployApplicationDetailRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 版本部署id
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *DescribeDeployApplicationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployApplicationDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeployApplicationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeployApplicationDetailResponseParams struct {
	// 分批发布结果详情
	Result *TemDeployApplicationDetailInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeployApplicationDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeployApplicationDetailResponseParams `json:"Response"`
}

func (r *DescribeDeployApplicationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployApplicationDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentRequestParams struct {
	// 命名空间id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源Channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源Channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentResponseParams struct {
	// 环境信息
	Result *NamespaceInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentStatusRequestParams struct {
	// 命名空间id
	EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`

	// 来源Channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeEnvironmentStatusRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间id
	EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`

	// 来源Channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeEnvironmentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentIds")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentStatusResponseParams struct {
	// 返回状态列表
	Result []*NamespaceStatusInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvironmentStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentStatusResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// 分页limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页下标
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 来源source
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 查询过滤器
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	SortInfo *SortType `json:"SortInfo,omitempty" name:"SortInfo"`

	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// 分页limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页下标
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 来源source
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 查询过滤器
	Filters []*QueryFilter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段
	SortInfo *SortType `json:"SortInfo,omitempty" name:"SortInfo"`

	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SourceChannel")
	delete(f, "Filters")
	delete(f, "SortInfo")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// 返回结果
	Result *NamespacePage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressRequestParams struct {
	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// ingress 规则名
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeIngressRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// ingress 规则名
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DescribeIngressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIngressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterNamespace")
	delete(f, "IngressName")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIngressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressResponseParams struct {
	// Ingress 规则配置
	Result *IngressInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIngressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIngressResponseParams `json:"Response"`
}

func (r *DescribeIngressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressesRequestParams struct {
	// 环境 id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境 namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// ingress 规则名列表
	IngressNames []*string `json:"IngressNames,omitempty" name:"IngressNames"`
}

type DescribeIngressesRequest struct {
	*tchttp.BaseRequest
	
	// 环境 id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境 namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// ingress 规则名列表
	IngressNames []*string `json:"IngressNames,omitempty" name:"IngressNames"`
}

func (r *DescribeIngressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIngressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterNamespace")
	delete(f, "SourceChannel")
	delete(f, "IngressNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIngressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIngressesResponseParams struct {
	// ingress 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*IngressInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIngressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIngressesResponseParams `json:"Response"`
}

func (r *DescribeIngressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIngressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogConfigRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogConfigResponseParams struct {
	// 配置
	Result *LogConfig `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogConfigResponseParams `json:"Response"`
}

func (r *DescribeLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePagedLogConfigListRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 规则名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页大小，默认 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页游标
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`
}

type DescribePagedLogConfigListRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 规则名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页大小，默认 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页游标
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`
}

func (r *DescribePagedLogConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePagedLogConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ApplicationId")
	delete(f, "ApplicationName")
	delete(f, "Name")
	delete(f, "Limit")
	delete(f, "ContinueToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePagedLogConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePagedLogConfigListResponseParams struct {
	// 日志收集配置列表
	Result *LogConfigListPage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePagedLogConfigListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePagedLogConfigListResponseParams `json:"Response"`
}

func (r *DescribePagedLogConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePagedLogConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedIngressesRequestParams struct {
	// 环境 id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境 namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DescribeRelatedIngressesRequest struct {
	*tchttp.BaseRequest
	
	// 环境 id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境 namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeRelatedIngressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelatedIngressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterNamespace")
	delete(f, "SourceChannel")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRelatedIngressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedIngressesResponseParams struct {
	// ingress 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*IngressInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRelatedIngressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRelatedIngressesResponseParams `json:"Response"`
}

func (r *DescribeRelatedIngressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelatedIngressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRunPodPage struct {
	// 分页下标
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 请求id
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// 条目
	PodList []*RunVersionPod `json:"PodList,omitempty" name:"PodList"`
}

// Predefined struct for user
type DestroyConfigDataRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DestroyConfigDataRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DestroyConfigDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyConfigDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyConfigDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyConfigDataResponseParams struct {
	// 返回结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyConfigDataResponse struct {
	*tchttp.BaseResponse
	Response *DestroyConfigDataResponseParams `json:"Response"`
}

func (r *DestroyConfigDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyConfigDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvironmentRequestParams struct {
	// 命名空间ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Namespace
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DestroyEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Namespace
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *DestroyEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvironmentResponseParams struct {
	// 返回结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *DestroyEnvironmentResponseParams `json:"Response"`
}

func (r *DestroyEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyLogConfigRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type DestroyLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DestroyLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyLogConfigResponseParams struct {
	// 返回结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DestroyLogConfigResponseParams `json:"Response"`
}

func (r *DestroyLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApplicationAutoscalerRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

type DisableApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

func (r *DisableApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "AutoscalerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApplicationAutoscalerResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *DisableApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *DisableApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EksService struct {
	// service name
	Name *string `json:"Name,omitempty" name:"Name"`

	// 可用端口
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// yaml 内容
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 版本名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 内网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIp []*string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// 外网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalIp *string `json:"ExternalIp,omitempty" name:"ExternalIp"`

	// 访问类型，可选值：
	// - EXTERNAL（公网访问）
	// - VPC（vpc内访问）
	// - CLUSTER（集群内访问）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 子网ID，只在类型为vpc访问时才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 负载均衡ID，只在外网访问和vpc内访问才有值，默认自动创建
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalanceId *string `json:"LoadBalanceId,omitempty" name:"LoadBalanceId"`

	// 端口映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortMappings []*PortMapping `json:"PortMappings,omitempty" name:"PortMappings"`

	// 每种类型访问配置详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServicePortMappingList []*ServicePortMapping `json:"ServicePortMappingList,omitempty" name:"ServicePortMappingList"`

	// 刷新复写所有类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlushAll *bool `json:"FlushAll,omitempty" name:"FlushAll"`

	// 1: 下次部署自动注入注册中心信息；0：不注入
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableRegistryNextDeploy *int64 `json:"EnableRegistryNextDeploy,omitempty" name:"EnableRegistryNextDeploy"`

	// 返回应用id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 所有服务IP是否已经ready
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllIpDone *bool `json:"AllIpDone,omitempty" name:"AllIpDone"`

	// clb 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalDomain *string `json:"ExternalDomain,omitempty" name:"ExternalDomain"`
}

// Predefined struct for user
type EnableApplicationAutoscalerRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

type EnableApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`
}

func (r *EnableApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "AutoscalerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableApplicationAutoscalerResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *EnableApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *EnableApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnablePrometheusConf struct {
	// 应用开放的监听端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 业务指标暴露的url path
	Path *string `json:"Path,omitempty" name:"Path"`
}

type EsInfo struct {
	// 最小实例数
	MinAliveInstances *int64 `json:"MinAliveInstances,omitempty" name:"MinAliveInstances"`

	// 最大实例数
	MaxAliveInstances *int64 `json:"MaxAliveInstances,omitempty" name:"MaxAliveInstances"`

	// 弹性策略,1:cpu，2:内存
	EsStrategy *int64 `json:"EsStrategy,omitempty" name:"EsStrategy"`

	// 弹性扩缩容条件值
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// 版本Id
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`
}

// Predefined struct for user
type GenerateApplicationPackageDownloadUrlRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 需要下载的包版本
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// 来源 channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type GenerateApplicationPackageDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 需要下载的包版本
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// 来源 channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *GenerateApplicationPackageDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateApplicationPackageDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "PkgName")
	delete(f, "DeployVersion")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateApplicationPackageDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateApplicationPackageDownloadUrlResponseParams struct {
	// 包下载临时链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateApplicationPackageDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *GenerateApplicationPackageDownloadUrlResponseParams `json:"Response"`
}

func (r *GenerateApplicationPackageDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateApplicationPackageDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheckConfig struct {
	// 支持的健康检查类型，如 HttpGet，TcpSocket，Exec
	Type *string `json:"Type,omitempty" name:"Type"`

	// 仅当健康检查类型为 HttpGet 时有效，表示协议类型，如 HTTP，HTTPS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 仅当健康检查类型为 HttpGet 时有效，表示请求路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 仅当健康检查类型为 Exec 时有效，表示执行的脚本内容
	Exec *string `json:"Exec,omitempty" name:"Exec"`

	// 仅当健康检查类型为 HttpGet\TcpSocket 时有效，表示请求路径
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 检查延迟开始时间，单位为秒，默认为 0
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// 超时时间，单位为秒，默认为 1
	TimeoutSeconds *int64 `json:"TimeoutSeconds,omitempty" name:"TimeoutSeconds"`

	// 间隔时间，单位为秒，默认为 10
	PeriodSeconds *int64 `json:"PeriodSeconds,omitempty" name:"PeriodSeconds"`
}

type HorizontalAutoscaler struct {
	// 最小实例数（可以不传）
	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// 最大实例数（可以不传）
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// 指标度量
	// CPU（CPU使用率，%）
	// MEMORY（内存使用率，%）
	// CPU_CORE_USED（CPU使用量，core）
	// MEMORY_SIZE_USED(内存使用量，MiB)
	// NETWORK_BANDWIDTH_RECEIVE(网络入带宽，MBps)
	// NETWORK_BANDWIDTH_TRANSMIT(网络出带宽，MBps)
	// NETWORK_TRAFFIC_RECEIVE(网络入流量，MiB/s)
	// NETWORK_TRAFFIC_TRANSMIT(网络出流量，MiB/s)
	// NETWORK_PACKETS_RECEIVE(网络入包量，Count/s)
	// NETWORK_PACKETS_TRANSMIT(网络出包量，Count/s)
	// FS_IOPS_WRITE(磁盘写次数，Count/s)
	// FS_IOPS_READ(磁盘读次数，Count/s)
	// FS_SIZE_WRITE(磁盘写大小，MiB/s)
	// FS_SIZE_READ(磁盘读大小，MiB/s)
	Metrics *string `json:"Metrics,omitempty" name:"Metrics"`

	// 阈值（整数）
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 阈值（小数，优先使用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoubleThreshold *float64 `json:"DoubleThreshold,omitempty" name:"DoubleThreshold"`
}

type IngressInfo struct {
	// 环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境namespace
	ClusterNamespace *string `json:"ClusterNamespace,omitempty" name:"ClusterNamespace"`

	// ip version
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// ingress name
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// rules 配置
	Rules []*IngressRule `json:"Rules,omitempty" name:"Rules"`

	// clb ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClbId *string `json:"ClbId,omitempty" name:"ClbId"`

	// tls 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tls []*IngressTls `json:"Tls,omitempty" name:"Tls"`

	// 环境集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// clb ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否混合 https，默认 false，可选值 true 代表有 https 协议监听
	Mixed *bool `json:"Mixed,omitempty" name:"Mixed"`

	// 重定向模式，可选值：
	// - AUTO（自动重定向http到https）
	// - NONE（不使用重定向）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewriteType *string `json:"RewriteType,omitempty" name:"RewriteType"`

	// clb 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type IngressRule struct {
	// ingress rule value
	Http *IngressRuleValue `json:"Http,omitempty" name:"Http"`

	// host 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 协议，选项为 http， https，默认为 http
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
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
	Paths []*IngressRulePath `json:"Paths,omitempty" name:"Paths"`
}

type IngressTls struct {
	// host 数组, 空数组表示全部域名的默认证书
	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`

	// secret name，如使用证书，则填空字符串
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// SSL Certificate Id
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type LogConfig struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收集类型，container_stdout 为标准输出；container_file 为文件；
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 日志集 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志提取模式，minimalist_log 为单行全文；multiline_log 为多行全文；  fullregex_log 为单行正则； multiline_fullregex_log 为多行正则； json_log 为 json；
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 首行正则表达式，当 LogType 为多行全文、多行正则时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginningRegex *string `json:"BeginningRegex,omitempty" name:"BeginningRegex"`

	// 收集文件目录，当 InputType=container_file 时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`

	// 收集文件名模式，当 InputType=container_file 时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// 应用 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 导出规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtractRule *LogConfigExtractRule `json:"ExtractRule,omitempty" name:"ExtractRule"`
}

type LogConfigExtractRule struct {
	// 首行正则表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginningRegex *string `json:"BeginningRegex,omitempty" name:"BeginningRegex"`

	// 提取结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*string `json:"Keys,omitempty" name:"Keys"`

	// 过滤键
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterKeys []*string `json:"FilterKeys,omitempty" name:"FilterKeys"`

	// 过滤值
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterRegex []*string `json:"FilterRegex,omitempty" name:"FilterRegex"`

	// 日志正则表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogRegex *string `json:"LogRegex,omitempty" name:"LogRegex"`

	// 时间字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`

	// 时间格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`

	// 是否上传解析失败日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnMatchUpload *string `json:"UnMatchUpload,omitempty" name:"UnMatchUpload"`

	// 解析失败日志的键名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnMatchedKey *string `json:"UnMatchedKey,omitempty" name:"UnMatchedKey"`

	// tracking
	// 注意：此字段可能返回 null，表示取不到有效值。
	Backtracking *string `json:"Backtracking,omitempty" name:"Backtracking"`

	// 分隔符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`
}

type LogConfigListPage struct {
	// 记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Records []*LogConfig `json:"Records,omitempty" name:"Records"`

	// 翻页游标
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContinueToken *string `json:"ContinueToken,omitempty" name:"ContinueToken"`
}

type LogOutputConf struct {
	// 日志消费端类型
	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`

	// cls日志集
	ClsLogsetName *string `json:"ClsLogsetName,omitempty" name:"ClsLogsetName"`

	// cls日志主题
	ClsLogTopicId *string `json:"ClsLogTopicId,omitempty" name:"ClsLogTopicId"`

	// cls日志集id
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// cls日志名称
	ClsLogTopicName *string `json:"ClsLogTopicName,omitempty" name:"ClsLogTopicName"`
}

// Predefined struct for user
type ModifyApplicationAutoscalerRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`

	// 弹性伸缩策略
	Autoscaler *Autoscaler `json:"Autoscaler,omitempty" name:"Autoscaler"`
}

type ModifyApplicationAutoscalerRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 弹性伸缩策略ID
	AutoscalerId *string `json:"AutoscalerId,omitempty" name:"AutoscalerId"`

	// 弹性伸缩策略
	Autoscaler *Autoscaler `json:"Autoscaler,omitempty" name:"Autoscaler"`
}

func (r *ModifyApplicationAutoscalerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationAutoscalerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "AutoscalerId")
	delete(f, "Autoscaler")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationAutoscalerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationAutoscalerResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationAutoscalerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationAutoscalerResponseParams `json:"Response"`
}

func (r *ModifyApplicationAutoscalerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationAutoscalerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationInfoRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 是否开启调用链,（此参数已弃用）
	EnableTracing *uint64 `json:"EnableTracing,omitempty" name:"EnableTracing"`
}

type ModifyApplicationInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 是否开启调用链,（此参数已弃用）
	EnableTracing *uint64 `json:"EnableTracing,omitempty" name:"EnableTracing"`
}

func (r *ModifyApplicationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "Description")
	delete(f, "SourceChannel")
	delete(f, "EnableTracing")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationInfoResponseParams struct {
	// 成功与否
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationInfoResponseParams `json:"Response"`
}

func (r *ModifyApplicationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationReplicasRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 实例数量
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type ModifyApplicationReplicasRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 实例数量
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *ModifyApplicationReplicasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationReplicasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "Replicas")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationReplicasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationReplicasResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationReplicasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationReplicasResponseParams `json:"Response"`
}

func (r *ModifyApplicationReplicasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationServiceRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 全量访问方式设置
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// 单条访问方式设置
	Data *ServicePortMapping `json:"Data,omitempty" name:"Data"`
}

type ModifyApplicationServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 全量访问方式设置
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// 单条访问方式设置
	Data *ServicePortMapping `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyApplicationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "SourceChannel")
	delete(f, "Service")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationServiceResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationServiceResponseParams `json:"Response"`
}

func (r *ModifyApplicationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigDataRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 配置信息
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

type ModifyConfigDataRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 配置信息
	Data []*Pair `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyConfigDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "SourceChannel")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigDataResponseParams struct {
	// 编辑是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyConfigDataResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConfigDataResponseParams `json:"Response"`
}

func (r *ModifyConfigDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvironmentRequestParams struct {
	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境名称
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 环境描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网网络
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境类型：test、pre、prod
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`
}

type ModifyEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境名称
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 环境描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网网络
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境类型：test、pre、prod
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`
}

func (r *ModifyEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "EnvironmentName")
	delete(f, "Description")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "SourceChannel")
	delete(f, "EnvType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvironmentResponseParams struct {
	// 成功时为环境ID，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvironmentResponseParams `json:"Response"`
}

func (r *ModifyEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIngressRequestParams struct {
	// Ingress 规则配置
	Ingress *IngressInfo `json:"Ingress,omitempty" name:"Ingress"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIngressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ingress")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIngressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIngressResponseParams struct {
	// 创建成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyIngressResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIngressResponseParams `json:"Response"`
}

func (r *ModifyIngressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIngressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogConfigRequestParams struct {
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 日志收集配置信息
	Data *LogConfig `json:"Data,omitempty" name:"Data"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type ModifyLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 配置名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 日志收集配置信息
	Data *LogConfig `json:"Data,omitempty" name:"Data"`

	// 应用 ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *ModifyLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Name")
	delete(f, "Data")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLogConfigResponseParams struct {
	// 编辑是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLogConfigResponseParams `json:"Response"`
}

func (r *ModifyLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountedSettingConf struct {
	// 配置名称
	ConfigDataName *string `json:"ConfigDataName,omitempty" name:"ConfigDataName"`

	// 挂载路径
	MountedPath *string `json:"MountedPath,omitempty" name:"MountedPath"`

	// 配置内容
	Data []*Pair `json:"Data,omitempty" name:"Data"`

	// 加密配置名称
	SecretDataName *string `json:"SecretDataName,omitempty" name:"SecretDataName"`
}

type NamespaceInfo struct {
	// ID 信息
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 名字（已弃用）
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// vpc id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// subnet id 数组
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// 环境名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// APM 资源 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApmInstanceId *string `json:"ApmInstanceId,omitempty" name:"ApmInstanceId"`

	// 环境是否上锁，1为上锁，0则未上锁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locked *int64 `json:"Locked,omitempty" name:"Locked"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 环境类型：test、pre、prod
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`
}

type NamespacePage struct {
	// 分页内容
	Records []*TemNamespaceInfo `json:"Records,omitempty" name:"Records"`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 条目数
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 页数
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 当前条目
	// 注意：此字段可能返回 null，表示取不到有效值。
	Current *int64 `json:"Current,omitempty" name:"Current"`
}

type NamespaceStatusInfo struct {
	// 命名空间id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 命名空间名称
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// TCB envId | EKS clusterId
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 环境状态
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 环境启动状态（不在启动中为null）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentStartingStatus *TemEnvironmentStartingStatus `json:"EnvironmentStartingStatus,omitempty" name:"EnvironmentStartingStatus"`

	// 环境停止状态（不在停止中为null）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentStoppingStatus *TemEnvironmentStoppingStatus `json:"EnvironmentStoppingStatus,omitempty" name:"EnvironmentStoppingStatus"`
}

type NodeInfo struct {
	// node名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// node可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// node子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 可用IP数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableIpCount *string `json:"AvailableIpCount,omitempty" name:"AvailableIpCount"`

	// cidr块
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
}

type Pair struct {
	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 类型，default 为自定义，reserved 为系统变量，referenced 为引用配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 配置名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *string `json:"Config,omitempty" name:"Config"`

	// 加密配置名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Secret *string `json:"Secret,omitempty" name:"Secret"`
}

type PortMapping struct {
	// 端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 映射端口
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// 协议栈 TCP/UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// k8s service名称
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type QueryFilter struct {
	// 查询字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 查询字段值
	Value []*string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type RestartApplicationPodRequestParams struct {
	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 名字
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 单页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页下标
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// pod状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type RestartApplicationPodRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 名字
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 单页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页下标
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// pod状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *RestartApplicationPodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartApplicationPodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ApplicationId")
	delete(f, "PodName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartApplicationPodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartApplicationPodResponseParams struct {
	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartApplicationPodResponse struct {
	*tchttp.BaseResponse
	Response *RestartApplicationPodResponseParams `json:"Response"`
}

func (r *RestartApplicationPodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartApplicationPodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartApplicationRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type RestartApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *RestartApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartApplicationResponseParams struct {
	// 返回结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartApplicationResponse struct {
	*tchttp.BaseResponse
	Response *RestartApplicationResponseParams `json:"Response"`
}

func (r *RestartApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeDeployApplicationRequestParams struct {
	// 需要开始下一批次的服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type ResumeDeployApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 需要开始下一批次的服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *ResumeDeployApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeDeployApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeDeployApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeDeployApplicationResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResumeDeployApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ResumeDeployApplicationResponseParams `json:"Response"`
}

func (r *ResumeDeployApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeDeployApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevertDeployApplicationRequestParams struct {
	// 需要回滚的服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 需要回滚的服务所在环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type RevertDeployApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 需要回滚的服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 需要回滚的服务所在环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *RevertDeployApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevertDeployApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevertDeployApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevertDeployApplicationResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RevertDeployApplicationResponse struct {
	*tchttp.BaseResponse
	Response *RevertDeployApplicationResponseParams `json:"Response"`
}

func (r *RevertDeployApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevertDeployApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollingUpdateApplicationByVersionRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 更新版本，IMAGE 部署为 tag 值；JAR/WAR 部署 为 Version
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// JAR/WAR 包名，仅 JAR/WAR 部署时必填
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 请求来源平台，含 IntelliJ，Coding
	From *string `json:"From,omitempty" name:"From"`

	// 部署策略，AUTO 为全自动；BETA 为小批量验证后自动；MANUAL 为全手动；
	DeployStrategyType *string `json:"DeployStrategyType,omitempty" name:"DeployStrategyType"`

	// 发布批次数
	TotalBatchCount *int64 `json:"TotalBatchCount,omitempty" name:"TotalBatchCount"`

	// 批次间隔时间
	BatchInterval *int64 `json:"BatchInterval,omitempty" name:"BatchInterval"`

	// 小批量验证批次的实例数
	BetaBatchNum *int64 `json:"BetaBatchNum,omitempty" name:"BetaBatchNum"`

	// 发布过程中保障的最小可用实例数
	MinAvailable *int64 `json:"MinAvailable,omitempty" name:"MinAvailable"`
}

type RollingUpdateApplicationByVersionRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 更新版本，IMAGE 部署为 tag 值；JAR/WAR 部署 为 Version
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// JAR/WAR 包名，仅 JAR/WAR 部署时必填
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 请求来源平台，含 IntelliJ，Coding
	From *string `json:"From,omitempty" name:"From"`

	// 部署策略，AUTO 为全自动；BETA 为小批量验证后自动；MANUAL 为全手动；
	DeployStrategyType *string `json:"DeployStrategyType,omitempty" name:"DeployStrategyType"`

	// 发布批次数
	TotalBatchCount *int64 `json:"TotalBatchCount,omitempty" name:"TotalBatchCount"`

	// 批次间隔时间
	BatchInterval *int64 `json:"BatchInterval,omitempty" name:"BatchInterval"`

	// 小批量验证批次的实例数
	BetaBatchNum *int64 `json:"BetaBatchNum,omitempty" name:"BetaBatchNum"`

	// 发布过程中保障的最小可用实例数
	MinAvailable *int64 `json:"MinAvailable,omitempty" name:"MinAvailable"`
}

func (r *RollingUpdateApplicationByVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollingUpdateApplicationByVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	delete(f, "DeployVersion")
	delete(f, "PackageName")
	delete(f, "From")
	delete(f, "DeployStrategyType")
	delete(f, "TotalBatchCount")
	delete(f, "BatchInterval")
	delete(f, "BetaBatchNum")
	delete(f, "MinAvailable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollingUpdateApplicationByVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollingUpdateApplicationByVersionResponseParams struct {
	// 版本ID
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RollingUpdateApplicationByVersionResponse struct {
	*tchttp.BaseResponse
	Response *RollingUpdateApplicationByVersionResponseParams `json:"Response"`
}

func (r *RollingUpdateApplicationByVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollingUpdateApplicationByVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunVersionPod struct {
	// shell地址
	Webshell *string `json:"Webshell,omitempty" name:"Webshell"`

	// pod的id
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例的ip
	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 部署版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// pod是否就绪
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ready *bool `json:"Ready,omitempty" name:"Ready"`

	// 容器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerState *string `json:"ContainerState,omitempty" name:"ContainerState"`

	// 实例所在节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInfo *NodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 是否健康
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unhealthy *bool `json:"Unhealthy,omitempty" name:"Unhealthy"`

	// 不健康时的提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnhealthyWarningMsg *string `json:"UnhealthyWarningMsg,omitempty" name:"UnhealthyWarningMsg"`

	// 版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`
}

type ServicePage struct {
	// 条目
	Records []*TemService `json:"Records,omitempty" name:"Records"`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 条目
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 页数
	Pages *int64 `json:"Pages,omitempty" name:"Pages"`

	// 当前条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Current *int64 `json:"Current,omitempty" name:"Current"`
}

type ServicePortMapping struct {
	// 服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 集群内访问vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// 集群外方位vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalIp *string `json:"ExternalIp,omitempty" name:"ExternalIp"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// LoadBalance Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalanceId *string `json:"LoadBalanceId,omitempty" name:"LoadBalanceId"`

	// yaml 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`

	// 暴露端口列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// 端口映射数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortMappingItemList []*ServicePortMappingItem `json:"PortMappingItemList,omitempty" name:"PortMappingItemList"`

	// clb domain
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalDomain *string `json:"ExternalDomain,omitempty" name:"ExternalDomain"`
}

type ServicePortMappingItem struct {
	// 应用访问端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 应用监听端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// 协议类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type ServiceVersionBrief struct {
	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 是否启动弹性 -- 已废弃
	EnableEs *int64 `json:"EnableEs,omitempty" name:"EnableEs"`

	// 当前实例
	CurrentInstances *int64 `json:"CurrentInstances,omitempty" name:"CurrentInstances"`

	// version的id
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 日志输出配置 -- 已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitempty" name:"LogOutputConf"`

	// 期望实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectedInstances *int64 `json:"ExpectedInstances,omitempty" name:"ExpectedInstances"`

	// 部署方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// 建构任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildTaskId *string `json:"BuildTaskId,omitempty" name:"BuildTaskId"`

	// 环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 环境name
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 服务name
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 是否正在发布中
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnderDeploying *bool `json:"UnderDeploying,omitempty" name:"UnderDeploying"`

	// 分批次部署状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchDeployStatus *string `json:"BatchDeployStatus,omitempty" name:"BatchDeployStatus"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInfos []*NodeInfo `json:"NodeInfos,omitempty" name:"NodeInfos"`

	// 实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodList *DescribeRunPodPage `json:"PodList,omitempty" name:"PodList"`

	// 工作负载信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadInfo *WorkloadInfo `json:"WorkloadInfo,omitempty" name:"WorkloadInfo"`

	// 创建日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 地域id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type SortType struct {
	// 排序字段名称
	Key *string `json:"Key,omitempty" name:"Key"`

	// 0：升序，1：倒序
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

// Predefined struct for user
type StopApplicationRequestParams struct {
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

type StopApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 环境ID
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *StopApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SourceChannel")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopApplicationResponseParams struct {
	// 返回结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopApplicationResponse struct {
	*tchttp.BaseResponse
	Response *StopApplicationResponseParams `json:"Response"`
}

func (r *StopApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageConf struct {
	// 存储卷名称
	StorageVolName *string `json:"StorageVolName,omitempty" name:"StorageVolName"`

	// 存储卷路径
	StorageVolPath *string `json:"StorageVolPath,omitempty" name:"StorageVolPath"`

	// 存储卷IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageVolIp *string `json:"StorageVolIp,omitempty" name:"StorageVolIp"`
}

type StorageMountConf struct {
	// 数据卷名
	VolumeName *string `json:"VolumeName,omitempty" name:"VolumeName"`

	// 数据卷绑定路径
	MountPath *string `json:"MountPath,omitempty" name:"MountPath"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TemDeployApplicationDetailInfo struct {
	// 分批发布策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 当前状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// beta分批详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BetaBatchDetail *DeployServiceBatchDetail `json:"BetaBatchDetail,omitempty" name:"BetaBatchDetail"`

	// 其他分批详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherBatchDetail []*DeployServiceBatchDetail `json:"OtherBatchDetail,omitempty" name:"OtherBatchDetail"`

	// 老版本pod列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldVersionPodList *DescribeRunPodPage `json:"OldVersionPodList,omitempty" name:"OldVersionPodList"`

	// 当前批次id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentBatchIndex *int64 `json:"CurrentBatchIndex,omitempty" name:"CurrentBatchIndex"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 当前批次状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentBatchStatus *string `json:"CurrentBatchStatus,omitempty" name:"CurrentBatchStatus"`

	// 新版本version
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewDeployVersion *string `json:"NewDeployVersion,omitempty" name:"NewDeployVersion"`

	// 旧版本version
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldDeployVersion *string `json:"OldDeployVersion,omitempty" name:"OldDeployVersion"`

	// 包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewVersionPackageInfo *string `json:"NewVersionPackageInfo,omitempty" name:"NewVersionPackageInfo"`

	// 下一批次开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextBatchStartTime *int64 `json:"NextBatchStartTime,omitempty" name:"NextBatchStartTime"`
}

type TemEnvironmentStartingStatus struct {
	// 需要启动的应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationNumNeedToStart *int64 `json:"ApplicationNumNeedToStart,omitempty" name:"ApplicationNumNeedToStart"`

	// 已经启动的应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartedApplicationNum *int64 `json:"StartedApplicationNum,omitempty" name:"StartedApplicationNum"`

	// 启动失败的应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartFailedApplicationNum *int64 `json:"StartFailedApplicationNum,omitempty" name:"StartFailedApplicationNum"`
}

type TemEnvironmentStoppingStatus struct {
	// 需要停止的应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationNumNeedToStop *int64 `json:"ApplicationNumNeedToStop,omitempty" name:"ApplicationNumNeedToStop"`

	// 已经停止的应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppedApplicationNum *int64 `json:"StoppedApplicationNum,omitempty" name:"StoppedApplicationNum"`

	// 停止失败的应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopFailedApplicationNum *int64 `json:"StopFailedApplicationNum,omitempty" name:"StopFailedApplicationNum"`
}

type TemNamespaceInfo struct {
	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 环境名称
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 区域名称
	Region *string `json:"Region,omitempty" name:"Region"`

	// 环境描述
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

	// 应用数
	ApplicationNum *int64 `json:"ApplicationNum,omitempty" name:"ApplicationNum"`

	// 运行实例数
	RunInstancesNum *int64 `json:"RunInstancesNum,omitempty" name:"RunInstancesNum"`

	// 子网络
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 环境集群 status
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 是否开启tsw
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`

	// 环境锁，1为上锁，0则为上锁
	Locked *int64 `json:"Locked,omitempty" name:"Locked"`

	// 用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 用户SubAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 资源是否有权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAuthority *bool `json:"HasAuthority,omitempty" name:"HasAuthority"`

	// 环境类型: test、pre、prod
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 地域码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

type TemService struct {
	// 主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 命名空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// 修改人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// tcr个人版or企业版
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoType *int64 `json:"RepoType,omitempty" name:"RepoType"`

	// 企业版实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 镜像仓库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 编程语言
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodingLanguage *string `json:"CodingLanguage,omitempty" name:"CodingLanguage"`

	// 部署方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// 环境名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 服务当前运行环境的实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveVersions []*ServiceVersionBrief `json:"ActiveVersions,omitempty" name:"ActiveVersions"`

	// 是否启用链路追踪
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableTracing *uint64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否有资源权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAuthority *bool `json:"HasAuthority,omitempty" name:"HasAuthority"`
}

type TemServiceVersionInfo struct {
	// 主键
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署方式
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// jdk版本
	JdkVersion *string `json:"JdkVersion,omitempty" name:"JdkVersion"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 部署版本
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// 发布方式
	PublishMode *string `json:"PublishMode,omitempty" name:"PublishMode"`

	// 启动参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 初始实例
	InitPodNum *int64 `json:"InitPodNum,omitempty" name:"InitPodNum"`

	// cpu规格
	CpuSpec *float64 `json:"CpuSpec,omitempty" name:"CpuSpec"`

	// 内存规格
	MemorySpec *float64 `json:"MemorySpec,omitempty" name:"MemorySpec"`

	// 镜像路径
	ImgRepo *string `json:"ImgRepo,omitempty" name:"ImgRepo"`

	// 镜像名称
	ImgName *string `json:"ImgName,omitempty" name:"ImgName"`

	// 镜像版本
	ImgVersion *string `json:"ImgVersion,omitempty" name:"ImgVersion"`

	// 弹性配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsInfo *EsInfo `json:"EsInfo,omitempty" name:"EsInfo"`

	// 环境配置
	EnvConf []*Pair `json:"EnvConf,omitempty" name:"EnvConf"`

	// 存储配置
	StorageConfs []*StorageConf `json:"StorageConfs,omitempty" name:"StorageConfs"`

	// 运行状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 私有网络
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网网络
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 创建时间
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 修改时间
	ModifyDate *string `json:"ModifyDate,omitempty" name:"ModifyDate"`

	// 挂载配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageMountConfs []*StorageMountConf `json:"StorageMountConfs,omitempty" name:"StorageMountConfs"`

	// 版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 日志输出配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogOutputConf *LogOutputConf `json:"LogOutputConf,omitempty" name:"LogOutputConf"`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 服务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationDescription *string `json:"ApplicationDescription,omitempty" name:"ApplicationDescription"`

	// 环境名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 公网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicDomain *string `json:"PublicDomain,omitempty" name:"PublicDomain"`

	// 是否开通公网访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnablePublicAccess *bool `json:"EnablePublicAccess,omitempty" name:"EnablePublicAccess"`

	// 现有的实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentInstances *int64 `json:"CurrentInstances,omitempty" name:"CurrentInstances"`

	// 期望的实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectedInstances *int64 `json:"ExpectedInstances,omitempty" name:"ExpectedInstances"`

	// 编程语言
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodingLanguage *string `json:"CodingLanguage,omitempty" name:"CodingLanguage"`

	// 程序包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 是否启用弹性伸缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsEnable *int64 `json:"EsEnable,omitempty" name:"EsEnable"`

	// 弹性策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsStrategy *int64 `json:"EsStrategy,omitempty" name:"EsStrategy"`

	// 镜像tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTag *string `json:"ImageTag,omitempty" name:"ImageTag"`

	// 是否启用log
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogEnable *int64 `json:"LogEnable,omitempty" name:"LogEnable"`

	// 最小实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinAliveInstances *string `json:"MinAliveInstances,omitempty" name:"MinAliveInstances"`

	// 安全组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 镜像命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageCommand *string `json:"ImageCommand,omitempty" name:"ImageCommand"`

	// 镜像命令参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageArgs []*string `json:"ImageArgs,omitempty" name:"ImageArgs"`

	// 是否使用默认注册中心配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitempty" name:"UseRegistryDefaultConfig"`

	// eks 访问设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *EksService `json:"Service,omitempty" name:"Service"`

	// 挂载配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitempty" name:"SettingConfs"`

	// log path数组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogConfs []*string `json:"LogConfs,omitempty" name:"LogConfs"`

	// 启动后立即执行的脚本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostStart *string `json:"PostStart,omitempty" name:"PostStart"`

	// 停止前执行的脚本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreStop *string `json:"PreStop,omitempty" name:"PreStop"`

	// 存活探针配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Liveness *HealthCheckConfig `json:"Liveness,omitempty" name:"Liveness"`

	// 就绪探针配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Readiness *HealthCheckConfig `json:"Readiness,omitempty" name:"Readiness"`

	// 弹性策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// 定时弹性策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronHorizontalAutoscaler []*CronHorizontalAutoscaler `json:"CronHorizontalAutoscaler,omitempty" name:"CronHorizontalAutoscaler"`

	// 应用实际可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 最新部署时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastDeployDate *string `json:"LastDeployDate,omitempty" name:"LastDeployDate"`

	// 最新部署成功时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastDeploySuccessDate *string `json:"LastDeploySuccessDate,omitempty" name:"LastDeploySuccessDate"`

	// 应用所在node信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInfos []*NodeInfo `json:"NodeInfos,omitempty" name:"NodeInfos"`

	// image类型 -0 为demo -1为正常image
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageType *int64 `json:"ImageType,omitempty" name:"ImageType"`

	// 是否启用调用链组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableTracing *uint64 `json:"EnableTracing,omitempty" name:"EnableTracing"`

	// 是否开启调用链上报，只有 EnableTracing=1 时生效（参数已弃用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableTracingReport *uint64 `json:"EnableTracingReport,omitempty" name:"EnableTracingReport"`

	// 镜像类型：0-个人镜像、1-企业镜像、2-公有镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoType *uint64 `json:"RepoType,omitempty" name:"RepoType"`

	// 分批发布子状态：batch_updating、batch_updating_waiting_confirm
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchDeployStatus *string `json:"BatchDeployStatus,omitempty" name:"BatchDeployStatus"`

	// APM 资源 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApmInstanceId *string `json:"ApmInstanceId,omitempty" name:"ApmInstanceId"`

	// 工作负载信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkloadInfo *WorkloadInfo `json:"WorkloadInfo,omitempty" name:"WorkloadInfo"`

	// 是否启用应用加速
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeedUp *bool `json:"SpeedUp,omitempty" name:"SpeedUp"`

	// 启动检测探针配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupProbe *HealthCheckConfig `json:"StartupProbe,omitempty" name:"StartupProbe"`

	// 操作系统版本，可选参数：
	// - ALPINE
	// - CENTOS
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsFlavour *string `json:"OsFlavour,omitempty" name:"OsFlavour"`

	// 镜像仓库server
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoServer *string `json:"RepoServer,omitempty" name:"RepoServer"`

	// 是否正在发布中
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnderDeploying *bool `json:"UnderDeploying,omitempty" name:"UnderDeploying"`

	// 监控业务指标监控
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnablePrometheusConf *EnablePrometheusConf `json:"EnablePrometheusConf,omitempty" name:"EnablePrometheusConf"`

	// 是否为手动停止
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppedManually *bool `json:"StoppedManually,omitempty" name:"StoppedManually"`

	// tcr实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcrInstanceId *string `json:"TcrInstanceId,omitempty" name:"TcrInstanceId"`

	// 1：开始自动metrics采集（open-telemetry）；
	// 0：关闭metrics采集；
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableMetrics *int64 `json:"EnableMetrics,omitempty" name:"EnableMetrics"`

	// 用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 用户SubAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 应用分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 是否启用注册中心
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableRegistry *int64 `json:"EnableRegistry,omitempty" name:"EnableRegistry"`

	// 弹性伸缩数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoscalerList []*Autoscaler `json:"AutoscalerList,omitempty" name:"AutoscalerList"`

	// 修改人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Modifier *string `json:"Modifier,omitempty" name:"Modifier"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 部署策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodList *DescribeRunPodPage `json:"PodList,omitempty" name:"PodList"`

	// 发布时配置是否有修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfEdited *bool `json:"ConfEdited,omitempty" name:"ConfEdited"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreStopEncoded *string `json:"PreStopEncoded,omitempty" name:"PreStopEncoded"`

	// 是否编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostStartEncoded *string `json:"PostStartEncoded,omitempty" name:"PostStartEncoded"`
}

type UseDefaultRepoParameters struct {
	// 企业版实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnterpriseInstanceName *string `json:"EnterpriseInstanceName,omitempty" name:"EnterpriseInstanceName"`

	// 企业版收费类型  0 按量收费   1 包年包月
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnterpriseInstanceChargeType *int64 `json:"EnterpriseInstanceChargeType,omitempty" name:"EnterpriseInstanceChargeType"`

	// 企业版规格：basic-基础班 ，standard-标准版，premium-高级版
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnterpriseInstanceType *string `json:"EnterpriseInstanceType,omitempty" name:"EnterpriseInstanceType"`
}

type WorkloadInfo struct {
	// 资源 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// Ready实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadyReplicas *int64 `json:"ReadyReplicas,omitempty" name:"ReadyReplicas"`

	// 实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// Updated实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedReplicas *int64 `json:"UpdatedReplicas,omitempty" name:"UpdatedReplicas"`

	// UpdatedReady实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedReadyReplicas *int64 `json:"UpdatedReadyReplicas,omitempty" name:"UpdatedReadyReplicas"`

	// 更新版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateRevision *string `json:"UpdateRevision,omitempty" name:"UpdateRevision"`

	// 当前版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentRevision *string `json:"CurrentRevision,omitempty" name:"CurrentRevision"`
}