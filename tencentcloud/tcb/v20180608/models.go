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

package v20180608

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AuthDomain struct {

	// 域名ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名类型。包含以下取值：
	// <li>SYSTEM</li>
	// <li>USER</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 状态。包含以下取值：
	// <li>ENABLE</li>
	// <li>DISABLE</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CheckTcbServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *CheckTcbServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckTcbServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckTcbServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true表示已开通
		Initialized *bool `json:"Initialized,omitempty" name:"Initialized"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckTcbServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckTcbServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CloudBaseCapabilities struct {

	// 启用安全能力项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Add []*string `json:"Add,omitempty" name:"Add" list`

	// 禁用安全能力向列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Drop []*string `json:"Drop,omitempty" name:"Drop" list`
}

type CloudBaseCodeRepoDetail struct {

	// repo的名字
	Name *CloudBaseCodeRepoName `json:"Name,omitempty" name:"Name"`

	// repo的url
	Url *string `json:"Url,omitempty" name:"Url"`
}

type CloudBaseCodeRepoName struct {

	// repo的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// repo的完整全名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullName *string `json:"FullName,omitempty" name:"FullName"`
}

type CloudBaseEsInfo struct {

	// es的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// secret名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *string `json:"Index,omitempty" name:"Index"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitempty" name:"Account"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`
}

type CloudBaseProjectVersion struct {

	// 项目名
	Name *string `json:"Name,omitempty" name:"Name"`

	// SAM json
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sam *string `json:"Sam,omitempty" name:"Sam"`

	// 来源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *CodeSource `json:"Source,omitempty" name:"Source"`

	// 创建时间, unix时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间 ,unix时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 项目状态, 枚举值: 
	//         "creatingEnv"-创建环境中
	// 	"createEnvFail"-创建环境失败
	// 	"building"-构建中
	// 	"buildFail"-构建失败
	// 	"deploying"-部署中
	// 	 "deployFail"-部署失败
	// 	 "success"-部署成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 环境变量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parameters []*KVPair `json:"Parameters,omitempty" name:"Parameters" list`

	// 项目类型, 枚举值:
	// "framework-oneclick" 控制台一键部署
	// "framework-local-oneclick" cli本地一键部署
	// "qci-extension-cicd" 内网coding ci cd
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// ci的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CIId *string `json:"CIId,omitempty" name:"CIId"`

	// cd的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CDId *string `json:"CDId,omitempty" name:"CDId"`

	// 环境id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionNum *int64 `json:"VersionNum,omitempty" name:"VersionNum"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// rc.json内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RcJson *string `json:"RcJson,omitempty" name:"RcJson"`

	// 插件配置内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddonConfig *string `json:"AddonConfig,omitempty" name:"AddonConfig"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 网络配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConfig *string `json:"NetworkConfig,omitempty" name:"NetworkConfig"`

	// 扩展id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailType *string `json:"FailType,omitempty" name:"FailType"`
}

type CloudBaseRunImageInfo struct {

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 是否公有
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 镜像tag名称
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 镜像server
	ServerAddr *string `json:"ServerAddr,omitempty" name:"ServerAddr"`

	// 镜像拉取地址
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

type CloudBaseRunImageSecretInfo struct {

	// 镜像地址
	RegistryServer *string `json:"RegistryServer,omitempty" name:"RegistryServer"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 仓库密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

type CloudBaseRunNfsVolumeSource struct {

	// NFS挂载Server
	Server *string `json:"Server,omitempty" name:"Server"`

	// Server路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 是否只读
	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// secret名称
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 临时目录
	EnableEmptyDirVolume *bool `json:"EnableEmptyDirVolume,omitempty" name:"EnableEmptyDirVolume"`
}

type CloudBaseRunSideSpec struct {

	// 容器镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerImage *string `json:"ContainerImage,omitempty" name:"ContainerImage"`

	// 容器端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerPort *int64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 容器的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// kv的json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvVar *string `json:"EnvVar,omitempty" name:"EnvVar"`

	// InitialDelaySeconds 延迟多长时间启动健康检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// CPU大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小（单位：M）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// 安全特性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Security *CloudBaseSecurityContext `json:"Security,omitempty" name:"Security"`

	// 挂载信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeMountInfos []*CloudBaseRunVolumeMount `json:"VolumeMountInfos,omitempty" name:"VolumeMountInfos" list`
}

type CloudBaseRunVolumeMount struct {

	// 资源名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 挂载路径
	MountPath *string `json:"MountPath,omitempty" name:"MountPath"`

	// 是否只读
	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// Nfs挂载信息
	NfsVolumes []*CloudBaseRunNfsVolumeSource `json:"NfsVolumes,omitempty" name:"NfsVolumes" list`
}

type CloudBaseRunVpcInfo struct {

	// vpc的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 创建类型(0=继承; 1=新建; 2=指定)
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateType *int64 `json:"CreateType,omitempty" name:"CreateType"`
}

type CloudBaseRunVpcSubnet struct {

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 子网的ipv4
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// subnet类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Target *string `json:"Target,omitempty" name:"Target"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CloudBaseSecurityContext struct {

	// 安全特性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Capabilities *CloudBaseCapabilities `json:"Capabilities,omitempty" name:"Capabilities"`
}

type CloudRunServiceSimpleVersionSnapshot struct {

	// 版本名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 版本备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// cpu规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *float64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mem *float64 `json:"Mem,omitempty" name:"Mem"`

	// 最小副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinNum *int64 `json:"MinNum,omitempty" name:"MinNum"`

	// 最大副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxNum *int64 `json:"MaxNum,omitempty" name:"MaxNum"`

	// 镜像url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 扩容策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 策略阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyThreshold *int64 `json:"PolicyThreshold,omitempty" name:"PolicyThreshold"`

	// 环境参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvParams *string `json:"EnvParams,omitempty" name:"EnvParams"`

	// 容器端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerPort *int64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 更新类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// dockerfile路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`

	// 构建路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildDir *string `json:"BuildDir,omitempty" name:"BuildDir"`

	// repo类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// 仓库
	// 注意：此字段可能返回 null，表示取不到有效值。
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// 分支
	// 注意：此字段可能返回 null，表示取不到有效值。
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 环境id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// package名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// package版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 自定义log路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomLogs *string `json:"CustomLogs,omitempty" name:"CustomLogs"`

	// 延时健康检查时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// snapshot名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 镜像信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *CloudBaseRunImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 代码仓库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeDetail *CloudBaseCodeRepoDetail `json:"CodeDetail,omitempty" name:"CodeDetail"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type CloudRunServiceVolume struct {

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// NFS的挂载方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	NFS *CloudBaseRunNfsVolumeSource `json:"NFS,omitempty" name:"NFS"`

	// secret名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 是否开启临时目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableEmptyDirVolume *bool `json:"EnableEmptyDirVolume,omitempty" name:"EnableEmptyDirVolume"`
}

type CodeSource struct {

	// 类型, 可能的枚举: "coding","package","package_url","github","gitlab","gitee","rawcode"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 工作目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkDir *string `json:"WorkDir,omitempty" name:"WorkDir"`

	// code包名, type为coding的时候需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodingPackageName *string `json:"CodingPackageName,omitempty" name:"CodingPackageName"`

	// coding版本名, type为coding的时候需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodingPackageVersion *string `json:"CodingPackageVersion,omitempty" name:"CodingPackageVersion"`

	// 源码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawCode *string `json:"RawCode,omitempty" name:"RawCode"`
}

type CommonServiceAPIRequest struct {
	*tchttp.BaseRequest

	// Service名，需要转发访问的接口名
	Service *string `json:"Service,omitempty" name:"Service"`

	// 需要转发的云API参数，要转成JSON格式
	JSONData *string `json:"JSONData,omitempty" name:"JSONData"`
}

func (r *CommonServiceAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommonServiceAPIRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommonServiceAPIResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// json格式response
		JSONResp *string `json:"JSONResp,omitempty" name:"JSONResp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommonServiceAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommonServiceAPIResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAndDeployCloudBaseProjectRequest struct {
	*tchttp.BaseRequest

	// 项目名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源
	Source *CodeSource `json:"Source,omitempty" name:"Source"`

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 项目类型, 枚举值为: framework-oneclick,qci-extension-cicd
	Type *string `json:"Type,omitempty" name:"Type"`

	// 环境变量
	Parameters []*KVPair `json:"Parameters,omitempty" name:"Parameters" list`

	// 环境别名
	EnvAlias *string `json:"EnvAlias,omitempty" name:"EnvAlias"`

	// rc.json的内容
	RcJson *string `json:"RcJson,omitempty" name:"RcJson"`

	// 插件配置内容
	AddonConfig *string `json:"AddonConfig,omitempty" name:"AddonConfig"`

	// 标签
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 网络配置
	NetworkConfig *string `json:"NetworkConfig,omitempty" name:"NetworkConfig"`

	// 免费额度的"basic", 不使用的用""
	FreeQuota *string `json:"FreeQuota,omitempty" name:"FreeQuota"`
}

func (r *CreateAndDeployCloudBaseProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAndDeployCloudBaseProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAndDeployCloudBaseProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAndDeployCloudBaseProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAndDeployCloudBaseProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuthDomainRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 安全域名
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`
}

func (r *CreateAuthDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuthDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAuthDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuthDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAuthDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCloudBaseRunResourceRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// vpc的ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表，当VpcId不为空，SubnetIds也不能为空
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`
}

func (r *CreateCloudBaseRunResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCloudBaseRunResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCloudBaseRunResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回集群创建是否成功 succ为成功。并且中间无err
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudBaseRunResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCloudBaseRunResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCloudBaseRunServerVersionRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 枚举（package/repository/image)
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// 流量占比
	FlowRatio *int64 `json:"FlowRatio,omitempty" name:"FlowRatio"`

	// Cpu的大小，单位：核
	Cpu *float64 `json:"Cpu,omitempty" name:"Cpu"`

	// Mem的大小，单位：G
	Mem *float64 `json:"Mem,omitempty" name:"Mem"`

	// 最小副本数，最小值：0
	MinNum *int64 `json:"MinNum,omitempty" name:"MinNum"`

	// 副本最大数，最大值：50
	MaxNum *int64 `json:"MaxNum,omitempty" name:"MaxNum"`

	// 策略类型(枚举值：比如cpu)
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 策略阈值
	PolicyThreshold *int64 `json:"PolicyThreshold,omitempty" name:"PolicyThreshold"`

	// 服务端口
	ContainerPort *int64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// repository的类型(coding/gitlab/github/coding)
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// Dockerfile地址
	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`

	// 构建目录
	BuildDir *string `json:"BuildDir,omitempty" name:"BuildDir"`

	// 环境变量
	EnvParams *string `json:"EnvParams,omitempty" name:"EnvParams"`

	// repository地址
	Repository *string `json:"Repository,omitempty" name:"Repository"`

	// 分支
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 代码包名字
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 代码包的版本
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// Image的详情
	ImageInfo *CloudBaseRunImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// Github等拉取代码的详情
	CodeDetail *CloudBaseCodeRepoDetail `json:"CodeDetail,omitempty" name:"CodeDetail"`

	// 私有镜像秘钥信息
	ImageSecretInfo *CloudBaseRunImageSecretInfo `json:"ImageSecretInfo,omitempty" name:"ImageSecretInfo"`

	// 私有镜像 认证名称
	ImagePullSecret *string `json:"ImagePullSecret,omitempty" name:"ImagePullSecret"`

	// 用户自定义采集日志路径
	CustomLogs *string `json:"CustomLogs,omitempty" name:"CustomLogs"`

	// 延迟多长时间开始健康检查（单位s）
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// cfs挂载信息
	MountVolumeInfo []*CloudBaseRunVolumeMount `json:"MountVolumeInfo,omitempty" name:"MountVolumeInfo" list`

	// 4 代表只能微信链路访问
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// es信息
	EsInfo *CloudBaseEsInfo `json:"EsInfo,omitempty" name:"EsInfo"`

	// 是否使用统一域名
	EnableUnion *bool `json:"EnableUnion,omitempty" name:"EnableUnion"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`

	// 服务路径
	ServerPath *string `json:"ServerPath,omitempty" name:"ServerPath"`

	// 镜像复用的key
	ImageReuseKey *string `json:"ImageReuseKey,omitempty" name:"ImageReuseKey"`

	// 容器的描述文件
	SidecarSpecs []*CloudBaseRunSideSpec `json:"SidecarSpecs,omitempty" name:"SidecarSpecs" list`

	// 安全特性
	Security *CloudBaseSecurityContext `json:"Security,omitempty" name:"Security"`

	// 服务磁盘挂载
	ServiceVolumes []*CloudRunServiceVolume `json:"ServiceVolumes,omitempty" name:"ServiceVolumes" list`
}

func (r *CreateCloudBaseRunServerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCloudBaseRunServerVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCloudBaseRunServerVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态(creating/succ)
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 版本名称（只有Result为succ的时候，才会返回VersionName)
	// 注意：此字段可能返回 null，表示取不到有效值。
		VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCloudBaseRunServerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCloudBaseRunServerVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHostingDomainRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`
}

func (r *CreateHostingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHostingDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateHostingDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHostingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateHostingDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePostpayPackageRequest struct {
	*tchttp.BaseRequest

	// 环境ID，需要系统自动创建环境时，此字段不传
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 微信 AppId，微信必传
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// 付费来源
	// <li>miniapp</li>
	// <li>qcloud</li>
	Source *string `json:"Source,omitempty" name:"Source"`

	// 用户享有的免费额度级别，目前只能为“basic”，不传该字段或该字段为空，标识不享受免费额度。
	FreeQuota *string `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// 环境创建来源，取值：
	// <li>miniapp</li>
	// <li>qcloud</li>
	// 用法同CreateEnv接口的Source参数
	// 和 Channel 参数同时传，或者同时不传；EnvId 为空时必传。
	EnvSource *string `json:"EnvSource,omitempty" name:"EnvSource"`

	// 环境别名，要以a-z开头，不能包含  a-z,0-9,-  以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 如果envsource为miniapp, channel可以为ide或api;
	// 如果envsource为qcloud, channel可以为qc_console,cocos, qq, cloudgame,dcloud,serverless_framework
	// 和 EnvSource 参数同时传，或者同时不传；EnvId 为空时必传。
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 扩展ID
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 订单标记。建议使用方统一转大小写之后再判断。
	// <li>QuickStart：快速启动来源</li>
	// <li>Activity：活动来源</li>
	Flag *string `json:"Flag,omitempty" name:"Flag"`
}

func (r *CreatePostpayPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePostpayPackageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePostpayPackageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 后付费订单号
		TranId *string `json:"TranId,omitempty" name:"TranId"`

		// 环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePostpayPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePostpayPackageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStaticStoreRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitempty" name:"EnableUnion"`
}

func (r *CreateStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStaticStoreRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建静态资源结果(succ/fail)
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStaticStoreResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DatabasesInfo struct {

	// 数据库唯一标识
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 状态。包含以下取值：
	// <li>INITIALIZING：资源初始化中</li>
	// <li>RUNNING：运行中，可正常使用的状态</li>
	// <li>UNUSABLE：禁用，不可用</li>
	// <li>OVERDUE：资源过期</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DeleteCloudBaseProjectLatestVersionRequest struct {
	*tchttp.BaseRequest

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 项目名
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *DeleteCloudBaseProjectLatestVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCloudBaseProjectLatestVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudBaseProjectLatestVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCloudBaseProjectLatestVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCloudBaseProjectLatestVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEndUserRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 用户列表，每一项都是uuid
	UserList []*string `json:"UserList,omitempty" name:"UserList" list`
}

func (r *DeleteEndUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEndUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEndUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEndUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEndUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthDomainsRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeAuthDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全域名列表列表
		Domains []*AuthDomain `json:"Domains,omitempty" name:"Domains" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseBuildServiceRequest struct {
	*tchttp.BaseRequest

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// build类型,枚举值有: cloudbaserun, framework-ci
	CIBusiness *string `json:"CIBusiness,omitempty" name:"CIBusiness"`

	// 服务版本
	ServiceVersion *string `json:"ServiceVersion,omitempty" name:"ServiceVersion"`
}

func (r *DescribeCloudBaseBuildServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseBuildServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseBuildServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 上传url
		UploadUrl *string `json:"UploadUrl,omitempty" name:"UploadUrl"`

		// 上传heder
		UploadHeaders []*KVPair `json:"UploadHeaders,omitempty" name:"UploadHeaders" list`

		// 包名
		PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

		// 包版本
		PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

		// 下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 下载Httpheader
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadHeaders []*KVPair `json:"DownloadHeaders,omitempty" name:"DownloadHeaders" list`

		// 下载链接是否过期
	// 注意：此字段可能返回 null，表示取不到有效值。
		OutDate *bool `json:"OutDate,omitempty" name:"OutDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudBaseBuildServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseBuildServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseProjectLatestVersionListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 个数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 环境id, 非必填
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 项目名称, 非必填
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目类型: framework-oneclick,qci-extension-cicd
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`

	// 标签
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`
}

func (r *DescribeCloudBaseProjectLatestVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseProjectLatestVersionListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseProjectLatestVersionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		ProjectList []*CloudBaseProjectVersion `json:"ProjectList,omitempty" name:"ProjectList" list`

		// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudBaseProjectLatestVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseProjectLatestVersionListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseRunResourceForExtendRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeCloudBaseRunResourceForExtendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseRunResourceForExtendRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseRunResourceForExtendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群状态(creating/succ)
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

		// 虚拟集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		VirtualClusterId *string `json:"VirtualClusterId,omitempty" name:"VirtualClusterId"`

		// vpc id信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// 地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Region *string `json:"Region,omitempty" name:"Region"`

		// 子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetIds []*CloudBaseRunVpcSubnet `json:"SubnetIds,omitempty" name:"SubnetIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudBaseRunResourceForExtendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseRunResourceForExtendResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseRunResourceRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeCloudBaseRunResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseRunResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseRunResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群状态(creating/succ)
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

		// 虚拟集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		VirtualClusterId *string `json:"VirtualClusterId,omitempty" name:"VirtualClusterId"`

		// vpc id信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// 地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Region *string `json:"Region,omitempty" name:"Region"`

		// 子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetIds []*CloudBaseRunVpcSubnet `json:"SubnetIds,omitempty" name:"SubnetIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudBaseRunResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseRunResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseRunServerVersionRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

func (r *DescribeCloudBaseRunServerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseRunServerVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseRunServerVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 版本名称
		VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

		// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// Dockefile的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
		DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`

		// DockerBuild的目录
	// 注意：此字段可能返回 null，表示取不到有效值。
		BuildDir *string `json:"BuildDir,omitempty" name:"BuildDir"`

		// 请使用CPUSize
		Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

		// 请使用MemSize
		Mem *int64 `json:"Mem,omitempty" name:"Mem"`

		// 副本最小值
		MinNum *int64 `json:"MinNum,omitempty" name:"MinNum"`

		// 副本最大值
		MaxNum *int64 `json:"MaxNum,omitempty" name:"MaxNum"`

		// 策略类型
		PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

		// 策略阈值
		PolicyThreshold *float64 `json:"PolicyThreshold,omitempty" name:"PolicyThreshold"`

		// 环境变量
	// 注意：此字段可能返回 null，表示取不到有效值。
		EnvParams *string `json:"EnvParams,omitempty" name:"EnvParams"`

		// 创建时间
		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

		// 更新时间
		UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

		// 版本的IP
	// 注意：此字段可能返回 null，表示取不到有效值。
		VersionIP *string `json:"VersionIP,omitempty" name:"VersionIP"`

		// 版本的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
		VersionPort *int64 `json:"VersionPort,omitempty" name:"VersionPort"`

		// 版本状态
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 代码包的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
		PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

		// 代码版本的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
		PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

		// 枚举（package/repository/image)
	// 注意：此字段可能返回 null，表示取不到有效值。
		UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

		// Repo的类型(coding/gitlab/github/coding)
	// 注意：此字段可能返回 null，表示取不到有效值。
		RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

		// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		Repo *string `json:"Repo,omitempty" name:"Repo"`

		// 分支
	// 注意：此字段可能返回 null，表示取不到有效值。
		Branch *string `json:"Branch,omitempty" name:"Branch"`

		// 服务名字
	// 注意：此字段可能返回 null，表示取不到有效值。
		ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

		// 是否对于外网开放
	// 注意：此字段可能返回 null，表示取不到有效值。
		IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

		// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// 子网实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

		// 日志采集路径
	// 注意：此字段可能返回 null，表示取不到有效值。
		CustomLogs *string `json:"CustomLogs,omitempty" name:"CustomLogs"`

		// 监听端口
	// 注意：此字段可能返回 null，表示取不到有效值。
		ContainerPort *int64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

		// 延迟多长时间开始健康检查（单位s）
	// 注意：此字段可能返回 null，表示取不到有效值。
		InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

		// 镜像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

		// CPU 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
		CpuSize *float64 `json:"CpuSize,omitempty" name:"CpuSize"`

		// MEM 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
		MemSize *float64 `json:"MemSize,omitempty" name:"MemSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudBaseRunServerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseRunServerVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseRunVersionSnapshotRequest struct {
	*tchttp.BaseRequest

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 版本历史名
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCloudBaseRunVersionSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseRunVersionSnapshotRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudBaseRunVersionSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 版本历史
	// 注意：此字段可能返回 null，表示取不到有效值。
		Snapshots []*CloudRunServiceSimpleVersionSnapshot `json:"Snapshots,omitempty" name:"Snapshots" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCloudBaseRunVersionSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCloudBaseRunVersionSnapshotResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseACLRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitempty" name:"CollectionName"`
}

func (r *DescribeDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseACLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
		AclTag *string `json:"AclTag,omitempty" name:"AclTag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseACLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadFileRequest struct {
	*tchttp.BaseRequest

	// 代码uri
	CodeUri *string `json:"CodeUri,omitempty" name:"CodeUri"`
}

func (r *DescribeDownloadFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadFileRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadFileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
		FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

		// 加密key
	// 注意：此字段可能返回 null，表示取不到有效值。
		CustomKey *string `json:"CustomKey,omitempty" name:"CustomKey"`

		// 下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadFileResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEndUserLoginStatisticRequest struct {
	*tchttp.BaseRequest

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 终端用户来源
	// <li> qcloud </li>
	// <li>miniapp</li>
	Source *string `json:"Source,omitempty" name:"Source"`
}

func (r *DescribeEndUserLoginStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEndUserLoginStatisticRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEndUserLoginStatisticResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境终端用户新增与登录统计
	// 注意：此字段可能返回 null，表示取不到有效值。
		LoginStatistics []*LoginStatistic `json:"LoginStatistics,omitempty" name:"LoginStatistics" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEndUserLoginStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEndUserLoginStatisticResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEndUserStatisticRequest struct {
	*tchttp.BaseRequest

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeEndUserStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEndUserStatisticRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEndUserStatisticResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 终端用户各平台统计
	// 注意：此字段可能返回 null，表示取不到有效值。
		PlatformStatistics []*PlatformStatistic `json:"PlatformStatistics,omitempty" name:"PlatformStatistics" list`

		// 终端用户总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEndUserStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEndUserStatisticResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEndUsersRequest struct {
	*tchttp.BaseRequest

	// 开发者的环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 可选参数，偏移量，默认 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 可选参数，拉取数量，默认 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 按照 uuid 列表过滤，最大个数为100
	UUIds []*string `json:"UUIds,omitempty" name:"UUIds" list`
}

func (r *DescribeEndUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEndUsersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEndUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户总数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 用户列表
		Users []*EndUserInfo `json:"Users,omitempty" name:"Users" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEndUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEndUsersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvFreeQuotaRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 资源类型：可选值：CDN, COS, FLEXDB, HOSTING, SCF
	// 不传则返回全部资源指标
	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes" list`
}

func (r *DescribeEnvFreeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvFreeQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvFreeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 免费抵扣配额详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		QuotaItems []*PostpayEnvQuota `json:"QuotaItems,omitempty" name:"QuotaItems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvFreeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvFreeQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvLimitRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeEnvLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvLimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境总数上限
		MaxEnvNum *int64 `json:"MaxEnvNum,omitempty" name:"MaxEnvNum"`

		// 目前环境总数
		CurrentEnvNum *int64 `json:"CurrentEnvNum,omitempty" name:"CurrentEnvNum"`

		// 免费环境数量上限
		MaxFreeEnvNum *int64 `json:"MaxFreeEnvNum,omitempty" name:"MaxFreeEnvNum"`

		// 目前免费环境数量
		CurrentFreeEnvNum *int64 `json:"CurrentFreeEnvNum,omitempty" name:"CurrentFreeEnvNum"`

		// 总计允许销毁环境次数上限
		MaxDeleteTotal *int64 `json:"MaxDeleteTotal,omitempty" name:"MaxDeleteTotal"`

		// 目前已销毁环境次数
		CurrentDeleteTotal *int64 `json:"CurrentDeleteTotal,omitempty" name:"CurrentDeleteTotal"`

		// 每月允许销毁环境次数上限
		MaxDeleteMonthly *int64 `json:"MaxDeleteMonthly,omitempty" name:"MaxDeleteMonthly"`

		// 本月已销毁环境次数
		CurrentDeleteMonthly *int64 `json:"CurrentDeleteMonthly,omitempty" name:"CurrentDeleteMonthly"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvLimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvsRequest struct {
	*tchttp.BaseRequest

	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitempty" name:"Channels" list`
}

func (r *DescribeEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境信息列表
		EnvList []*EnvInfo `json:"EnvList,omitempty" name:"EnvList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExtraPkgBillingInfoRequest struct {
	*tchttp.BaseRequest

	// 已购买增值包的环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeExtraPkgBillingInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExtraPkgBillingInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeExtraPkgBillingInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 增值包计费信息列表
		EnvInfoList []*EnvBillingInfoItem `json:"EnvInfoList,omitempty" name:"EnvInfoList" list`

		// 增值包数目
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExtraPkgBillingInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeExtraPkgBillingInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePostpayFreeQuotasRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribePostpayFreeQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePostpayFreeQuotasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePostpayFreeQuotasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 免费量资源信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		FreequotaInfoList []*FreequotaInfo `json:"FreequotaInfoList,omitempty" name:"FreequotaInfoList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePostpayFreeQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePostpayFreeQuotasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePostpayPackageFreeQuotasRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 免费额度类型标识
	FreeQuotaType *string `json:"FreeQuotaType,omitempty" name:"FreeQuotaType"`
}

func (r *DescribePostpayPackageFreeQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePostpayPackageFreeQuotasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePostpayPackageFreeQuotasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 免费量资源信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PackageFreeQuotaInfos []*PackageFreeQuotaInfo `json:"PackageFreeQuotaInfos,omitempty" name:"PackageFreeQuotaInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePostpayPackageFreeQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePostpayPackageFreeQuotasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaDataRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// <li> 指标名: </li>
	// <li> StorageSizepkg: 当月存储空间容量, 单位MB </li>
	// <li> StorageReadpkg: 当月存储读请求次数 </li>
	// <li> StorageWritepkg: 当月存储写请求次数 </li>
	// <li> StorageCdnOriginFluxpkg: 当月CDN回源流量, 单位字节 </li>
	// <li> StorageCdnOriginFluxpkgDay: 当日CDN回源流量, 单位字节 </li>
	// <li> StorageReadpkgDay: 当日存储读请求次数 </li>
	// <li> StorageWritepkgDay: 当日写请求次数 </li>
	// <li> CDNFluxpkg: 当月CDN流量, 单位为字节 </li>
	// <li> CDNFluxpkgDay: 当日CDN流量, 单位为字节 </li>
	// <li> FunctionInvocationpkg: 当月云函数调用次数 </li>
	// <li> FunctionGBspkg: 当月云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkg: 当月云函数流量, 单位千字节(KB) </li>
	// <li> FunctionInvocationpkgDay: 当日云函数调用次数 </li>
	// <li> FunctionGBspkgDay: 当日云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkgDay: 当日云函数流量, 单位千字节(KB) </li>
	// <li> DbSizepkg: 当月数据库容量大小, 单位MB </li>
	// <li> DbReadpkg: 当日数据库读请求数 </li>
	// <li> DbWritepkg: 当日数据库写请求数 </li>
	// <li> StaticFsFluxPkgDay: 当日静态托管流量 </li>
	// <li> StaticFsFluxPkg: 当月静态托管流量</li>
	// <li> StaticFsSizePkg: 当月静态托管容量 </li>
	// <li> TkeCpuUsedPkg: 当月容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeCpuUsedPkgDay: 当天容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeMemUsedPkg: 当月容器托管内存使用量，单位MB*秒 </li>
	// <li> TkeMemUsedPkgDay: 当天容器托管内存使用量，单位MB*秒 </li>
	// <li> CodingBuildTimePkgDay: 当天容器托管构建时间使用量，单位毫秒 </li>
	// <li> TkeHttpServiceNatPkgDay: 当天容器托管流量使用量，单位B </li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 资源ID, 目前仅对云函数、容器托管相关的指标有意义。云函数(FunctionInvocationpkg, FunctionGBspkg, FunctionFluxpkg)、容器托管（服务名称）。如果想查询某个云函数的指标则在ResourceId中传入函数名; 如果只想查询整个namespace的指标, 则留空或不传。
	ResourceID *string `json:"ResourceID,omitempty" name:"ResourceID"`
}

func (r *DescribeQuotaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 指标名
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// 指标的值
		Value *int64 `json:"Value,omitempty" name:"Value"`

		// 指标的附加值信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubValue *string `json:"SubValue,omitempty" name:"SubValue"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQuotaDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsQuotasRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeSmsQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsQuotasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSmsQuotasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 短信免费量信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		SmsFreeQuotaList []*SmsFreeQuota `json:"SmsFreeQuotaList,omitempty" name:"SmsFreeQuotaList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSmsQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSmsQuotasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyEnvRequest struct {
	*tchttp.BaseRequest

	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 针对预付费 删除隔离中的环境时要传true 正常环境直接跳过隔离期删除
	IsForce *bool `json:"IsForce,omitempty" name:"IsForce"`

	// 是否绕过资源检查，资源包等额外资源，默认为false，如果为true，则不检查资源是否有数据，直接删除。
	BypassCheck *bool `json:"BypassCheck,omitempty" name:"BypassCheck"`
}

func (r *DestroyEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyStaticStoreRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// cdn域名
	CdnDomain *string `json:"CdnDomain,omitempty" name:"CdnDomain"`
}

func (r *DestroyStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyStaticStoreRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 条件任务结果(succ/fail)
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyStaticStoreResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EndUserInfo struct {

	// 用户唯一ID
	UUId *string `json:"UUId,omitempty" name:"UUId"`

	// 微信ID
	WXOpenId *string `json:"WXOpenId,omitempty" name:"WXOpenId"`

	// qq ID
	QQOpenId *string `json:"QQOpenId,omitempty" name:"QQOpenId"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 昵称
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 性别
	Gender *string `json:"Gender,omitempty" name:"Gender"`

	// 头像地址
	AvatarUrl *string `json:"AvatarUrl,omitempty" name:"AvatarUrl"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否为匿名用户
	IsAnonymous *bool `json:"IsAnonymous,omitempty" name:"IsAnonymous"`

	// 是否禁用账户
	IsDisabled *bool `json:"IsDisabled,omitempty" name:"IsDisabled"`

	// 是否设置过密码
	HasPassword *bool `json:"HasPassword,omitempty" name:"HasPassword"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type EnvBillingInfoItem struct {

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// tcb产品套餐ID，参考DescribePackages接口的返回值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 自动续费标记
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" name:"IsAutoRenew"`

	// 状态。包含以下取值：
	// <li> 空字符串：初始化中</li>
	// <li> NORMAL：正常</li>
	// <li> ISOLATE：隔离</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 支付方式。包含以下取值：
	// <li> PREPAYMENT：预付费</li>
	// <li> POSTPAID：后付费</li>
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 隔离时间，最近一次隔离的时间
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// 过期时间，套餐即将到期的时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 创建时间，第一次接入计费方案的时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间，计费信息最近一次更新的时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// true表示从未升级过付费版。
	IsAlwaysFree *bool `json:"IsAlwaysFree,omitempty" name:"IsAlwaysFree"`

	// 付费渠道。
	// <li> miniapp：小程序</li>
	// <li> qcloud：腾讯云</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentChannel *string `json:"PaymentChannel,omitempty" name:"PaymentChannel"`

	// 最新的订单信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderInfo *OrderInfo `json:"OrderInfo,omitempty" name:"OrderInfo"`

	// 免费配额信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeQuota *string `json:"FreeQuota,omitempty" name:"FreeQuota"`
}

type EnvInfo struct {

	// 账户下该环境唯一标识
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 环境来源。包含以下取值：
	// <li>miniapp：微信小程序</li>
	// <li>qcloud ：腾讯云</li>
	Source *string `json:"Source,omitempty" name:"Source"`

	// 环境别名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 环境状态。包含以下取值：
	// <li>NORMAL：正常可用</li>
	// <li>UNAVAILABLE：服务不可用，可能是尚未初始化或者初始化过程中</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 数据库列表
	Databases []*DatabasesInfo `json:"Databases,omitempty" name:"Databases" list`

	// 存储列表
	Storages []*StorageInfo `json:"Storages,omitempty" name:"Storages" list`

	// 函数列表
	Functions []*FunctionInfo `json:"Functions,omitempty" name:"Functions" list`

	// tcb产品套餐ID，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 套餐中文名称，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 云日志服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogServices []*LogServiceInfo `json:"LogServices,omitempty" name:"LogServices" list`

	// 静态资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaticStorages []*StaticStorageInfo `json:"StaticStorages,omitempty" name:"StaticStorages" list`

	// 是否到期自动降为免费版
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAutoDegrade *bool `json:"IsAutoDegrade,omitempty" name:"IsAutoDegrade"`

	// 环境渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvChannel *string `json:"EnvChannel,omitempty" name:"EnvChannel"`

	// 支付方式。包含以下取值：
	// <li> prepayment：预付费</li>
	// <li> postpaid：后付费</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 是否为默认环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 环境所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type EstablishCloudBaseRunServerRequest struct {
	*tchttp.BaseRequest

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 是否开通外网访问
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 镜像仓库
	ImageRepo *string `json:"ImageRepo,omitempty" name:"ImageRepo"`

	// 服务描述
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// es信息
	EsInfo *CloudBaseEsInfo `json:"EsInfo,omitempty" name:"EsInfo"`

	// 日志类型; es/cls
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`

	// 来源方（默认值：qcloud，微信侧来源miniapp)
	Source *string `json:"Source,omitempty" name:"Source"`

	// vpc信息
	VpcInfo *CloudBaseRunVpcInfo `json:"VpcInfo,omitempty" name:"VpcInfo"`

	// 0/1=允许公网访问;2=关闭公网访问
	PublicAccess *int64 `json:"PublicAccess,omitempty" name:"PublicAccess"`
}

func (r *EstablishCloudBaseRunServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EstablishCloudBaseRunServerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EstablishCloudBaseRunServerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EstablishCloudBaseRunServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EstablishCloudBaseRunServerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FreequotaInfo struct {

	// 资源类型
	// <li>COS</li>
	// <li>CDN</li>
	// <li>FLEXDB</li>
	// <li>SCF</li>
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源指标名称
	ResourceMetric *string `json:"ResourceMetric,omitempty" name:"ResourceMetric"`

	// 资源指标免费量
	FreeQuota *int64 `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// 指标单位
	MetricUnit *string `json:"MetricUnit,omitempty" name:"MetricUnit"`

	// 免费量抵扣周期
	// <li>sum-month:以月为单位抵扣</li>
	// <li>sum-day:以天为单位抵扣</li>
	// <li>totalize:总容量抵扣</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductType *string `json:"DeductType,omitempty" name:"DeductType"`

	// 免费量类型
	// <li>basic:通用量抵扣</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeQuotaType *string `json:"FreeQuotaType,omitempty" name:"FreeQuotaType"`
}

type FunctionInfo struct {

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}

type KVPair struct {

	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogServiceInfo struct {

	// log名
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// log-id
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// topic名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// topic-id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// cls日志所属地域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type LoginStatistic struct {

	// 统计类型 新增NEWUSER 和登录 LOGIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticalType *string `json:"StatisticalType,omitempty" name:"StatisticalType"`

	// 统计周期：日DAY，周WEEK，月MONTH
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticalCycle *string `json:"StatisticalCycle,omitempty" name:"StatisticalCycle"`

	// 统计总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ModifyDatabaseACLRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitempty" name:"AclTag"`
}

func (r *ModifyDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDatabaseACLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDatabaseACLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEndUserRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// C端用户端的唯一ID
	UUId *string `json:"UUId,omitempty" name:"UUId"`

	// 帐号的状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyEndUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEndUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEndUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEndUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEndUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 环境备注名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *ModifyEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OrderInfo struct {

	// 订单号
	TranId *string `json:"TranId,omitempty" name:"TranId"`

	// 订单要切换的套餐ID
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 订单类型
	// <li>1 购买</li>
	// <li>2 续费</li>
	// <li>3 变配</li>
	TranType *string `json:"TranType,omitempty" name:"TranType"`

	// 订单状态。
	// <li>1未支付</li>
	// <li>2 支付中</li>
	// <li>3 发货中</li>
	// <li>4 发货成功</li>
	// <li>5 发货失败</li>
	// <li>6 已退款</li>
	// <li>7 已取消</li>
	// <li>100 已删除</li>
	TranStatus *string `json:"TranStatus,omitempty" name:"TranStatus"`

	// 订单更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 订单创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 付费模式.
	// <li>prepayment 预付费</li>
	// <li>postpaid 后付费</li>
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 订单绑定的扩展ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 资源初始化结果(仅当ExtensionId不为空时有效): successful(初始化成功), failed(初始化失败), doing(初始化进行中), init(准备初始化)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceReady *string `json:"ResourceReady,omitempty" name:"ResourceReady"`
}

type PackageFreeQuotaInfo struct {

	// 资源类型
	// <li>COS</li>
	// <li>CDN</li>
	// <li>FLEXDB</li>
	// <li>SCF</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源指标名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceMetric *string `json:"ResourceMetric,omitempty" name:"ResourceMetric"`

	// 资源指标免费量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeQuota *int64 `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// 指标单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricUnit *string `json:"MetricUnit,omitempty" name:"MetricUnit"`

	// 免费量抵扣周期
	// <li>sum-month:以月为单位抵扣</li>
	// <li>sum-day:以天为单位抵扣</li>
	// <li>totalize:总容量抵扣</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductType *string `json:"DeductType,omitempty" name:"DeductType"`

	// 免费量类型
	// <li>basic:通用量抵扣</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeQuotaType *string `json:"FreeQuotaType,omitempty" name:"FreeQuotaType"`
}

type PlatformStatistic struct {

	// 终端用户从属平台
	// 注意：此字段可能返回 null，表示取不到有效值。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 平台终端用户数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type PostpayEnvQuota struct {

	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 配额值
	Value *uint64 `json:"Value,omitempty" name:"Value"`

	// 配额生效时间
	// 为空表示没有时间限制
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 配额失效时间
	// 为空表示没有时间限制
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ReinstateEnvRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *ReinstateEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReinstateEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReinstateEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReinstateEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReinstateEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SmsFreeQuota struct {

	// 免费量总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeQuota *uint64 `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// 共计已使用总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedQuota *uint64 `json:"TotalUsedQuota,omitempty" name:"TotalUsedQuota"`

	// 免费周期起点，0000-00-00 00:00:00 形式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleStart *string `json:"CycleStart,omitempty" name:"CycleStart"`

	// 免费周期终点，0000-00-00 00:00:00 形式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleEnd *string `json:"CycleEnd,omitempty" name:"CycleEnd"`

	// 今天已使用总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TodayUsedQuota *uint64 `json:"TodayUsedQuota,omitempty" name:"TodayUsedQuota"`
}

type StaticStorageInfo struct {

	// 静态CDN域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaticDomain *string `json:"StaticDomain,omitempty" name:"StaticDomain"`

	// 静态CDN默认文件夹，当前为根目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDirName *string `json:"DefaultDirName,omitempty" name:"DefaultDirName"`

	// 资源状态(process/online/offline/init)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// cos所属区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// bucket信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
}

type StorageInfo struct {

	// 资源所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// 桶名，存储资源的唯一标识
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// cdn 域名
	CdnDomain *string `json:"CdnDomain,omitempty" name:"CdnDomain"`

	// 资源所属用户的腾讯云appId
	AppId *string `json:"AppId,omitempty" name:"AppId"`
}
