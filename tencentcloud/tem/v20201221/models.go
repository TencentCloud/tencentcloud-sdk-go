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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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

// Predefined struct for user
type CreateCosTokenRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
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
type CreateCosTokenV2RequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// optType 1上传  2查询
	OptType *int64 `json:"OptType,omitempty" name:"OptType"`

	// 来源 channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 充当deployVersion入参
	TimeVersion *string `json:"TimeVersion,omitempty" name:"TimeVersion"`
}

type CreateCosTokenV2Request struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// optType 1上传  2查询
	OptType *int64 `json:"OptType,omitempty" name:"OptType"`

	// 来源 channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 充当deployVersion入参
	TimeVersion *string `json:"TimeVersion,omitempty" name:"TimeVersion"`
}

func (r *CreateCosTokenV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "PkgName")
	delete(f, "OptType")
	delete(f, "SourceChannel")
	delete(f, "TimeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosTokenV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosTokenV2ResponseParams struct {
	// 成功时为CosToken对象，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CosToken `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCosTokenV2Response struct {
	*tchttp.BaseResponse
	Response *CreateCosTokenV2ResponseParams `json:"Response"`
}

func (r *CreateCosTokenV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceRequestParams struct {
	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 命名空间描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// K8s version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 是否开启tsw服务
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 命名空间描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// K8s version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 是否开启tsw服务
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`
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
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "Description")
	delete(f, "K8sVersion")
	delete(f, "SourceChannel")
	delete(f, "EnableTswTraceService")
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
type CreateResourceRequestParams struct {
	// 命名空间 Id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 资源类型，目前支持文件系统：CFS；日志服务：CLS；注册中心：TSE_SRE
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源 Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间 Id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 资源类型，目前支持文件系统：CFS；日志服务：CLS；注册中心：TSE_SRE
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源 Id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
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
	delete(f, "NamespaceId")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "SourceChannel")
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

// Predefined struct for user
type CreateServiceV2RequestParams struct {
	// 服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

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

	// 服务所在子网
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
}

type CreateServiceV2Request struct {
	*tchttp.BaseRequest
	
	// 服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

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

	// 服务所在子网
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
}

func (r *CreateServiceV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceName")
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceV2ResponseParams struct {
	// 服务code
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateServiceV2Response struct {
	*tchttp.BaseResponse
	Response *CreateServiceV2ResponseParams `json:"Response"`
}

func (r *CreateServiceV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIngressRequestParams struct {
	// tem NamespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// eks namespace 名
	EksNamespace *string `json:"EksNamespace,omitempty" name:"EksNamespace"`

	// ingress 规则名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DeleteIngressRequest struct {
	*tchttp.BaseRequest
	
	// tem NamespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// eks namespace 名
	EksNamespace *string `json:"EksNamespace,omitempty" name:"EksNamespace"`

	// ingress 规则名
	Name *string `json:"Name,omitempty" name:"Name"`

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
	delete(f, "NamespaceId")
	delete(f, "EksNamespace")
	delete(f, "Name")
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
type DeployServiceV2RequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 容器端口
	ContainerPort *uint64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 初始化 pod 数
	InitPodNum *uint64 `json:"InitPodNum,omitempty" name:"InitPodNum"`

	// cpu规格
	CpuSpec *float64 `json:"CpuSpec,omitempty" name:"CpuSpec"`

	// 内存规格
	MemorySpec *float64 `json:"MemorySpec,omitempty" name:"MemorySpec"`

	// 环境ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 镜像仓库
	ImgRepo *string `json:"ImgRepo,omitempty" name:"ImgRepo"`

	// 版本描述信息
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// 启动参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 弹性伸缩配置，不传默认不启用弹性伸缩配置
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
	// - KONA：使用 kona jdk。
	// - OPEN：使用 open jdk。
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

	// 服务端口映射
	PortMappings []*PortMapping `json:"PortMappings,omitempty" name:"PortMappings"`

	// 是否添加默认注册中心配置
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitempty" name:"UseRegistryDefaultConfig"`

	// 挂载配置信息
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitempty" name:"SettingConfs"`

	// eks 访问设置
	EksService *EksService `json:"EksService,omitempty" name:"EksService"`

	// 要回滚到的历史版本id
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 启动后执行的脚本
	PostStart *string `json:"PostStart,omitempty" name:"PostStart"`

	// 停止前执行的脚本
	PreStop *string `json:"PreStop,omitempty" name:"PreStop"`

	// 分批发布策略配置
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// 存活探针配置
	Liveness *HealthCheckConfig `json:"Liveness,omitempty" name:"Liveness"`

	// 就绪探针配置
	Readiness *HealthCheckConfig `json:"Readiness,omitempty" name:"Readiness"`
}

type DeployServiceV2Request struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 容器端口
	ContainerPort *uint64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 初始化 pod 数
	InitPodNum *uint64 `json:"InitPodNum,omitempty" name:"InitPodNum"`

	// cpu规格
	CpuSpec *float64 `json:"CpuSpec,omitempty" name:"CpuSpec"`

	// 内存规格
	MemorySpec *float64 `json:"MemorySpec,omitempty" name:"MemorySpec"`

	// 环境ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 镜像仓库
	ImgRepo *string `json:"ImgRepo,omitempty" name:"ImgRepo"`

	// 版本描述信息
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// 启动参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 弹性伸缩配置，不传默认不启用弹性伸缩配置
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
	// - KONA：使用 kona jdk。
	// - OPEN：使用 open jdk。
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

	// 服务端口映射
	PortMappings []*PortMapping `json:"PortMappings,omitempty" name:"PortMappings"`

	// 是否添加默认注册中心配置
	UseRegistryDefaultConfig *bool `json:"UseRegistryDefaultConfig,omitempty" name:"UseRegistryDefaultConfig"`

	// 挂载配置信息
	SettingConfs []*MountedSettingConf `json:"SettingConfs,omitempty" name:"SettingConfs"`

	// eks 访问设置
	EksService *EksService `json:"EksService,omitempty" name:"EksService"`

	// 要回滚到的历史版本id
	VersionId *string `json:"VersionId,omitempty" name:"VersionId"`

	// 启动后执行的脚本
	PostStart *string `json:"PostStart,omitempty" name:"PostStart"`

	// 停止前执行的脚本
	PreStop *string `json:"PreStop,omitempty" name:"PreStop"`

	// 分批发布策略配置
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// 存活探针配置
	Liveness *HealthCheckConfig `json:"Liveness,omitempty" name:"Liveness"`

	// 就绪探针配置
	Readiness *HealthCheckConfig `json:"Readiness,omitempty" name:"Readiness"`
}

func (r *DeployServiceV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployServiceV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ContainerPort")
	delete(f, "InitPodNum")
	delete(f, "CpuSpec")
	delete(f, "MemorySpec")
	delete(f, "NamespaceId")
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
	delete(f, "PortMappings")
	delete(f, "UseRegistryDefaultConfig")
	delete(f, "SettingConfs")
	delete(f, "EksService")
	delete(f, "VersionId")
	delete(f, "PostStart")
	delete(f, "PreStop")
	delete(f, "DeployStrategyConf")
	delete(f, "Liveness")
	delete(f, "Readiness")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployServiceV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployServiceV2ResponseParams struct {
	// 版本ID（前端可忽略）
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployServiceV2Response struct {
	*tchttp.BaseResponse
	Response *DeployServiceV2ResponseParams `json:"Response"`
}

func (r *DeployServiceV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployServiceV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployStrategyConf struct {
	// 总分批数
	TotalBatchCount *int64 `json:"TotalBatchCount,omitempty" name:"TotalBatchCount"`

	// beta分批实例数
	BetaBatchNum *int64 `json:"BetaBatchNum,omitempty" name:"BetaBatchNum"`

	// 分批策略：0-全自动，1-全手动，beta分批一定是手动的，这里的策略指定的是剩余批次
	DeployStrategyType *int64 `json:"DeployStrategyType,omitempty" name:"DeployStrategyType"`

	// 每批暂停间隔
	BatchInterval *int64 `json:"BatchInterval,omitempty" name:"BatchInterval"`
}

// Predefined struct for user
type DescribeIngressRequestParams struct {
	// tem namespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// eks namespace 名
	EksNamespace *string `json:"EksNamespace,omitempty" name:"EksNamespace"`

	// ingress 规则名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type DescribeIngressRequest struct {
	*tchttp.BaseRequest
	
	// tem namespaceId
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// eks namespace 名
	EksNamespace *string `json:"EksNamespace,omitempty" name:"EksNamespace"`

	// ingress 规则名
	Name *string `json:"Name,omitempty" name:"Name"`

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
	delete(f, "NamespaceId")
	delete(f, "EksNamespace")
	delete(f, "Name")
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
	// namespace id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// namespace
	EksNamespace *string `json:"EksNamespace,omitempty" name:"EksNamespace"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// ingress 规则名列表
	Names []*string `json:"Names,omitempty" name:"Names"`
}

type DescribeIngressesRequest struct {
	*tchttp.BaseRequest
	
	// namespace id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// namespace
	EksNamespace *string `json:"EksNamespace,omitempty" name:"EksNamespace"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// ingress 规则名列表
	Names []*string `json:"Names,omitempty" name:"Names"`
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
	delete(f, "NamespaceId")
	delete(f, "EksNamespace")
	delete(f, "SourceChannel")
	delete(f, "Names")
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
type DescribeNamespacesRequestParams struct {
	// 分页limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页下标
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 来源source
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
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

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesResponseParams struct {
	// 返回结果
	Result *NamespacePage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespacesResponseParams `json:"Response"`
}

func (r *DescribeNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedIngressesRequestParams struct {
	// 环境 id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// EKS namespace
	EksNamespace *string `json:"EksNamespace,omitempty" name:"EksNamespace"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

type DescribeRelatedIngressesRequest struct {
	*tchttp.BaseRequest
	
	// 环境 id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// EKS namespace
	EksNamespace *string `json:"EksNamespace,omitempty" name:"EksNamespace"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
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
	delete(f, "NamespaceId")
	delete(f, "EksNamespace")
	delete(f, "SourceChannel")
	delete(f, "ServiceId")
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
type DescribeServiceRunPodListV2RequestParams struct {
	// 环境id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 服务名id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

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

type DescribeServiceRunPodListV2Request struct {
	*tchttp.BaseRequest
	
	// 环境id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 服务名id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

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

func (r *DescribeServiceRunPodListV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceRunPodListV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "PodName")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceRunPodListV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceRunPodListV2ResponseParams struct {
	// 返回结果
	Result *DescribeRunPodPage `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceRunPodListV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeServiceRunPodListV2ResponseParams `json:"Response"`
}

func (r *DescribeServiceRunPodListV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceRunPodListV2Response) FromJsonString(s string) error {
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
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

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
type GenerateDownloadUrlRequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 需要下载的包版本
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// 来源 channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type GenerateDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 需要下载的包版本
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// 来源 channel
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *GenerateDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "PkgName")
	delete(f, "DeployVersion")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateDownloadUrlResponseParams struct {
	// 包下载临时链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *GenerateDownloadUrlResponseParams `json:"Response"`
}

func (r *GenerateDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDownloadUrlResponse) FromJsonString(s string) error {
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
	Rules []*IngressRule `json:"Rules,omitempty" name:"Rules"`

	// clb ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClbId *string `json:"ClbId,omitempty" name:"ClbId"`

	// tls 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tls []*IngressTls `json:"Tls,omitempty" name:"Tls"`

	// eks clusterId
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
type ModifyNamespaceRequestParams struct {
	// 环境id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 命名空间描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网网络
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
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
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
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
	delete(f, "Description")
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceResponseParams struct {
	// 成功时为命名空间ID，失败为null
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
type ModifyServiceInfoRequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

type ModifyServiceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
}

func (r *ModifyServiceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Description")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceInfoResponseParams struct {
	// 成功与否
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyServiceInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceInfoResponseParams `json:"Response"`
}

func (r *ModifyServiceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountedSettingConf struct {
	// 配置名称
	ConfigDataName *string `json:"ConfigDataName,omitempty" name:"ConfigDataName"`

	// 挂载路径
	MountedPath *string `json:"MountedPath,omitempty" name:"MountedPath"`

	// 配置内容
	Data []*Pair `json:"Data,omitempty" name:"Data"`
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
}

type Pair struct {
	// 建
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PortMapping struct {
	// 端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 映射端口
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// 协议栈 TCP/UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

// Predefined struct for user
type RestartServiceRunPodRequestParams struct {
	// 环境id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 服务名id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

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

type RestartServiceRunPodRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 服务名id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

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

func (r *RestartServiceRunPodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartServiceRunPodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NamespaceId")
	delete(f, "ServiceId")
	delete(f, "PodName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "SourceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartServiceRunPodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartServiceRunPodResponseParams struct {
	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartServiceRunPodResponse struct {
	*tchttp.BaseResponse
	Response *RestartServiceRunPodResponseParams `json:"Response"`
}

func (r *RestartServiceRunPodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartServiceRunPodResponse) FromJsonString(s string) error {
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

	// 是否开启tsw
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`
}