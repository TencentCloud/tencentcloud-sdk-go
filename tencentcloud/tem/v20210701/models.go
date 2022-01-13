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

	// 是否启用调用链功能
	EnableTracing *int64 `json:"EnableTracing,omitempty" name:"EnableTracing"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务code
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 环境名称
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 私有网络名称
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 子网列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 环境描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// K8s version
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 是否开启tsw服务
	EnableTswTraceService *bool `json:"EnableTswTraceService,omitempty" name:"EnableTswTraceService"`
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
	delete(f, "Vpc")
	delete(f, "SubnetIds")
	delete(f, "Description")
	delete(f, "K8sVersion")
	delete(f, "SourceChannel")
	delete(f, "EnableTswTraceService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功时为环境ID，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功与否
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteIngressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否删除成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 弹性策略
	HorizontalAutoscaler []*HorizontalAutoscaler `json:"HorizontalAutoscaler,omitempty" name:"HorizontalAutoscaler"`

	// 定时弹性策略
	CronHorizontalAutoscaler []*CronHorizontalAutoscaler `json:"CronHorizontalAutoscaler,omitempty" name:"CronHorizontalAutoscaler"`

	// 是否启用log，1为启用，0为不启用
	LogEnable *int64 `json:"LogEnable,omitempty" name:"LogEnable"`

	// （除开镜像配置）配置是否修改
	ConfEdited *bool `json:"ConfEdited,omitempty" name:"ConfEdited"`

	// 是否开启应用加速
	SpeedUp *bool `json:"SpeedUp,omitempty" name:"SpeedUp"`

	// 启动探针配置
	StartupProbe *HealthCheckConfig `json:"StartupProbe,omitempty" name:"StartupProbe"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeployApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 版本ID（前端可忽略）
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
}

type DeployStrategyConf struct {

	// 总分批数
	TotalBatchCount *int64 `json:"TotalBatchCount,omitempty" name:"TotalBatchCount"`

	// beta分批实例数
	BetaBatchNum *int64 `json:"BetaBatchNum,omitempty" name:"BetaBatchNum"`

	// 分批策略：0-全自动，1-全手动，2-beta分批，beta批一定是手动的
	DeployStrategyType *int64 `json:"DeployStrategyType,omitempty" name:"DeployStrategyType"`

	// 每批暂停间隔
	BatchInterval *int64 `json:"BatchInterval,omitempty" name:"BatchInterval"`

	// 最小可用实例数
	MinAvailable *int64 `json:"MinAvailable,omitempty" name:"MinAvailable"`
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

type DescribeApplicationPodsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *DescribeRunPodPage `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDeployApplicationDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分批发布结果详情
		Result *TemDeployApplicationDetailInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 分页limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页下标
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 来源source
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *NamespacePage `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeIngressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Ingress 规则配置
		Result *IngressInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeIngressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ingress 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result []*IngressInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeRelatedIngressesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ingress 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result []*IngressInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GenerateApplicationPackageDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 包下载临时链接
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 最小实例数
	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// 最大实例数
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// 指标度量（CPU or MEMORY）
	Metrics *string `json:"Metrics,omitempty" name:"Metrics"`

	// 阈值（百分比）
	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
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

type ModifyApplicationInfoRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 来源渠道
	SourceChannel *int64 `json:"SourceChannel,omitempty" name:"SourceChannel"`

	// 是否开启调用链, 0 为关闭，1位开启
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

type ModifyApplicationInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功与否
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyApplicationReplicasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功时为环境ID，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIngressResponse) FromJsonString(s string) error {
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
}

type PortMapping struct {

	// 端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 映射端口
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// 协议栈 TCP/UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
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

type RestartApplicationPodResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type RestartApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ResumeDeployApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type RevertDeployApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollingUpdateApplicationByVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RollingUpdateApplicationByVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 版本ID
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type StopApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
}
