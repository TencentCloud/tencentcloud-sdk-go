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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActivityInfoItem struct {
	// 活动id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 记录插入时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 记录最后一次变更时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 活动开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 活动结束时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 自定义备注信息
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type ActivityRecordItem struct {
	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 活动id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 自定义状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 自定义子状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubStatus *string `json:"SubStatus,omitempty" name:"SubStatus"`

	// 整型子状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubStatusInt *int64 `json:"SubStatusInt,omitempty" name:"SubStatusInt"`

	// 是否软删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDeleted *bool `json:"IsDeleted,omitempty" name:"IsDeleted"`
}

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

type BaasPackageInfo struct {
	// DAU产品套餐ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// DAU套餐中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageTitle *string `json:"PackageTitle,omitempty" name:"PackageTitle"`

	// 套餐分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 套餐分组中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupTitle *string `json:"GroupTitle,omitempty" name:"GroupTitle"`

	// json格式化计费标签，例如：
	// {"pid":2, "cids":{"create": 2, "renew": 2, "modify": 2}, "productCode":"p_tcb_mp", "subProductCode":"sp_tcb_mp_cloudbase_dau"}
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillTags *string `json:"BillTags,omitempty" name:"BillTags"`

	// json格式化用户资源限制，例如：
	// {"Qps":1000,"InvokeNum":{"TimeUnit":"m", "Unit":"万次", "MaxSize": 100},"Capacity":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100}, "Cdn":{"Flux":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100}, "BackFlux":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100}},"Scf":{"Concurrency":1000,"OutFlux":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100},"MemoryUse":{"TimeUnit":"m", "Unit":"WGBS", "MaxSize": 100000}}}
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceLimit *string `json:"ResourceLimit,omitempty" name:"ResourceLimit"`

	// json格式化高级限制，例如：
	// {"CMSEnable":false,"ProvisionedConcurrencyMem":512000, "PictureProcessing":false, "SecurityAudit":false, "RealTimePush":false, "TemplateMessageBatchPush":false, "Payment":false}
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceLimit *string `json:"AdvanceLimit,omitempty" name:"AdvanceLimit"`

	// 套餐描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageDescription *string `json:"PackageDescription,omitempty" name:"PackageDescription"`

	// 是否对外展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsExternal *bool `json:"IsExternal,omitempty" name:"IsExternal"`
}

type BackendServiceInfo struct {
	// 服务名称
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 服务状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type BindEnvGatewayRequestParams struct {
	// 子环境id
	SubEnvId *string `json:"SubEnvId,omitempty" name:"SubEnvId"`
}

type BindEnvGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 子环境id
	SubEnvId *string `json:"SubEnvId,omitempty" name:"SubEnvId"`
}

func (r *BindEnvGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEnvGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubEnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindEnvGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindEnvGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindEnvGatewayResponse struct {
	*tchttp.BaseResponse
	Response *BindEnvGatewayResponseParams `json:"Response"`
}

func (r *BindEnvGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEnvGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTcbServiceRequestParams struct {

}

type CheckTcbServiceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CheckTcbServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTcbServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckTcbServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTcbServiceResponseParams struct {
	// true表示已开通
	Initialized *bool `json:"Initialized,omitempty" name:"Initialized"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckTcbServiceResponse struct {
	*tchttp.BaseResponse
	Response *CheckTcbServiceResponseParams `json:"Response"`
}

func (r *CheckTcbServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTcbServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudBaseCapabilities struct {
	// 启用安全能力项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Add []*string `json:"Add,omitempty" name:"Add"`

	// 禁用安全能力向列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Drop []*string `json:"Drop,omitempty" name:"Drop"`
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
	Parameters []*KVPair `json:"Parameters,omitempty" name:"Parameters"`

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
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 网络配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConfig *string `json:"NetworkConfig,omitempty" name:"NetworkConfig"`

	// 扩展id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtensionId *string `json:"ExtensionId,omitempty" name:"ExtensionId"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailType *string `json:"FailType,omitempty" name:"FailType"`

	// 私有仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoUrl *string `json:"RepoUrl,omitempty" name:"RepoUrl"`

	// 是否私有仓库代码变更触发自动部署
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoDeployOnCodeChange *bool `json:"AutoDeployOnCodeChange,omitempty" name:"AutoDeployOnCodeChange"`

	// ci部署进度（%）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildPercent *int64 `json:"BuildPercent,omitempty" name:"BuildPercent"`
}

type CloudBaseRunEmptyDirVolumeSource struct {
	// 启用emptydir数据卷
	EnableEmptyDirVolume *bool `json:"EnableEmptyDirVolume,omitempty" name:"EnableEmptyDirVolume"`

	// "","Memory","HugePages"
	Medium *string `json:"Medium,omitempty" name:"Medium"`

	// emptydir数据卷大小
	SizeLimit *string `json:"SizeLimit,omitempty" name:"SizeLimit"`
}

type CloudBaseRunForGatewayConf struct {
	// 是否缩容到0
	IsZero *bool `json:"IsZero,omitempty" name:"IsZero"`

	// 按百分比灰度的权重
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 按请求/header参数的灰度Key
	GrayKey *string `json:"GrayKey,omitempty" name:"GrayKey"`

	// 按请求/header参数的灰度Value
	GrayValue *string `json:"GrayValue,omitempty" name:"GrayValue"`

	// 是否为默认版本(按请求/header参数)
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 访问权限，对应二进制分多段，vpc内网｜公网｜oa
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 访问的URL（域名＋路径）列表
	URLs []*string `json:"URLs,omitempty" name:"URLs"`

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 灰度类型：FLOW(权重), URL_PARAMS/HEAD_PARAMS
	GrayType *string `json:"GrayType,omitempty" name:"GrayType"`

	// CLB的IP:Port
	LbAddr *string `json:"LbAddr,omitempty" name:"LbAddr"`

	// 0:http访问服务配置信息, 1: 服务域名
	ConfigType *int64 `json:"ConfigType,omitempty" name:"ConfigType"`
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

type CloudBaseRunKVPriority struct {
	// 参数的Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 参数的Value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
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

type CloudBaseRunServerVersionItem struct {
	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 流量占比
	FlowRatio *int64 `json:"FlowRatio,omitempty" name:"FlowRatio"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 构建ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildId *int64 `json:"BuildId,omitempty" name:"BuildId"`

	// 构建方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// url中的参数路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlParam *ObjectKV `json:"UrlParam,omitempty" name:"UrlParam"`

	// 优先级（数值越小，优先级越高）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 是否是默认兜底版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefaultPriority *bool `json:"IsDefaultPriority,omitempty" name:"IsDefaultPriority"`

	// KV Params
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowParams []*CloudBaseRunKVPriority `json:"FlowParams,omitempty" name:"FlowParams"`

	// 最小副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// 最大副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// 操作记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunId *string `json:"RunId,omitempty" name:"RunId"`

	// 进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`

	// 当前副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentReplicas *int64 `json:"CurrentReplicas,omitempty" name:"CurrentReplicas"`

	// Monolithic，Microservice
	// 注意：此字段可能返回 null，表示取不到有效值。
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`
}

type CloudBaseRunServiceVolumeHostPath struct {

}

type CloudBaseRunServiceVolumeMount struct {
	// Volume 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 挂载路径
	MountPath *string `json:"MountPath,omitempty" name:"MountPath"`

	// 是否只读
	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// 子路径
	SubPath *string `json:"SubPath,omitempty" name:"SubPath"`

	// 传播挂载方式
	MountPropagation *string `json:"MountPropagation,omitempty" name:"MountPropagation"`
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
	VolumeMountInfos []*CloudBaseRunVolumeMount `json:"VolumeMountInfos,omitempty" name:"VolumeMountInfos"`
}

type CloudBaseRunVersionFlowItem struct {
	// 版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 流量占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowRatio *int64 `json:"FlowRatio,omitempty" name:"FlowRatio"`

	// 流量参数键值对（URL参数/HEADERS参数）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlParam *ObjectKV `json:"UrlParam,omitempty" name:"UrlParam"`

	// 优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 是否是默认兜底版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefaultPriority *bool `json:"IsDefaultPriority,omitempty" name:"IsDefaultPriority"`
}

type CloudBaseRunVolumeMount struct {
	// 资源名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 挂载路径
	MountPath *string `json:"MountPath,omitempty" name:"MountPath"`

	// 是否只读
	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// Nfs挂载信息
	NfsVolumes []*CloudBaseRunNfsVolumeSource `json:"NfsVolumes,omitempty" name:"NfsVolumes"`
}

type CloudBaseRunVpcInfo struct {
	// vpc的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

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

	// 是否开启临时目录逐步废弃，请使用 EmptyDir
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableEmptyDirVolume *bool `json:"EnableEmptyDirVolume,omitempty" name:"EnableEmptyDirVolume"`

	// emptydir数据卷详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmptyDir *CloudBaseRunEmptyDirVolumeSource `json:"EmptyDir,omitempty" name:"EmptyDir"`

	// 主机路径挂载信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostPath *CloudBaseRunServiceVolumeHostPath `json:"HostPath,omitempty" name:"HostPath"`
}

type ClsInfo struct {
	// cls所属地域
	ClsRegion *string `json:"ClsRegion,omitempty" name:"ClsRegion"`

	// cls日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// cls日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
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

	// 代码分支
	// 注意：此字段可能返回 null，表示取不到有效值。
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// coding项目ID，type为coding时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// coding项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

// Predefined struct for user
type CommonServiceAPIRequestParams struct {
	// Service名，需要转发访问的接口名
	Service *string `json:"Service,omitempty" name:"Service"`

	// 需要转发的云API参数，要转成JSON格式
	JSONData *string `json:"JSONData,omitempty" name:"JSONData"`

	// 指定角色
	ApiRole *string `json:"ApiRole,omitempty" name:"ApiRole"`
}

type CommonServiceAPIRequest struct {
	*tchttp.BaseRequest
	
	// Service名，需要转发访问的接口名
	Service *string `json:"Service,omitempty" name:"Service"`

	// 需要转发的云API参数，要转成JSON格式
	JSONData *string `json:"JSONData,omitempty" name:"JSONData"`

	// 指定角色
	ApiRole *string `json:"ApiRole,omitempty" name:"ApiRole"`
}

func (r *CommonServiceAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommonServiceAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "JSONData")
	delete(f, "ApiRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CommonServiceAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommonServiceAPIResponseParams struct {
	// json格式response
	JSONResp *string `json:"JSONResp,omitempty" name:"JSONResp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CommonServiceAPIResponse struct {
	*tchttp.BaseResponse
	Response *CommonServiceAPIResponseParams `json:"Response"`
}

func (r *CommonServiceAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommonServiceAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndDeployCloudBaseProjectRequestParams struct {
	// 项目名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 来源
	Source *CodeSource `json:"Source,omitempty" name:"Source"`

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 项目类型, 枚举值为: framework-oneclick,qci-extension-cicd
	Type *string `json:"Type,omitempty" name:"Type"`

	// 环境变量
	Parameters []*KVPair `json:"Parameters,omitempty" name:"Parameters"`

	// 环境别名。要以a-z开头，不能包含a-zA-z0-9-以外的字符
	EnvAlias *string `json:"EnvAlias,omitempty" name:"EnvAlias"`

	// rc.json的内容
	RcJson *string `json:"RcJson,omitempty" name:"RcJson"`

	// 插件配置内容
	AddonConfig *string `json:"AddonConfig,omitempty" name:"AddonConfig"`

	// 标签
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 网络配置
	NetworkConfig *string `json:"NetworkConfig,omitempty" name:"NetworkConfig"`

	// 用户享有的免费额度级别，目前只能为“basic”，不传该字段或该字段为空，标识不享受免费额度。
	FreeQuota *string `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// 是否代码变更触发自动部署
	AutoDeployOnCodeChange *bool `json:"AutoDeployOnCodeChange,omitempty" name:"AutoDeployOnCodeChange"`

	// 私有仓库地址
	RepoUrl *string `json:"RepoUrl,omitempty" name:"RepoUrl"`
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
	Parameters []*KVPair `json:"Parameters,omitempty" name:"Parameters"`

	// 环境别名。要以a-z开头，不能包含a-zA-z0-9-以外的字符
	EnvAlias *string `json:"EnvAlias,omitempty" name:"EnvAlias"`

	// rc.json的内容
	RcJson *string `json:"RcJson,omitempty" name:"RcJson"`

	// 插件配置内容
	AddonConfig *string `json:"AddonConfig,omitempty" name:"AddonConfig"`

	// 标签
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 网络配置
	NetworkConfig *string `json:"NetworkConfig,omitempty" name:"NetworkConfig"`

	// 用户享有的免费额度级别，目前只能为“basic”，不传该字段或该字段为空，标识不享受免费额度。
	FreeQuota *string `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// 是否代码变更触发自动部署
	AutoDeployOnCodeChange *bool `json:"AutoDeployOnCodeChange,omitempty" name:"AutoDeployOnCodeChange"`

	// 私有仓库地址
	RepoUrl *string `json:"RepoUrl,omitempty" name:"RepoUrl"`
}

func (r *CreateAndDeployCloudBaseProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndDeployCloudBaseProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Source")
	delete(f, "EnvId")
	delete(f, "Type")
	delete(f, "Parameters")
	delete(f, "EnvAlias")
	delete(f, "RcJson")
	delete(f, "AddonConfig")
	delete(f, "Tags")
	delete(f, "NetworkConfig")
	delete(f, "FreeQuota")
	delete(f, "AutoDeployOnCodeChange")
	delete(f, "RepoUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndDeployCloudBaseProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndDeployCloudBaseProjectResponseParams struct {
	// 环境Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAndDeployCloudBaseProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndDeployCloudBaseProjectResponseParams `json:"Response"`
}

func (r *CreateAndDeployCloudBaseProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndDeployCloudBaseProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuthDomainRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 安全域名
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

type CreateAuthDomainRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 安全域名
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *CreateAuthDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuthDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuthDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAuthDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuthDomainResponseParams `json:"Response"`
}

func (r *CreateAuthDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudBaseRunResourceRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// vpc的ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表，当VpcId不为空，SubnetIds也不能为空
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

type CreateCloudBaseRunResourceRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// vpc的ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表，当VpcId不为空，SubnetIds也不能为空
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

func (r *CreateCloudBaseRunResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudBaseRunResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "VpcId")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudBaseRunResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudBaseRunResourceResponseParams struct {
	// 返回集群创建是否成功 succ为成功。并且中间无err
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCloudBaseRunResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudBaseRunResourceResponseParams `json:"Response"`
}

func (r *CreateCloudBaseRunResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudBaseRunResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudBaseRunServerVersionRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 枚举（package/repository/image/jar/war)
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
	MountVolumeInfo []*CloudBaseRunVolumeMount `json:"MountVolumeInfo,omitempty" name:"MountVolumeInfo"`

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
	SidecarSpecs []*CloudBaseRunSideSpec `json:"SidecarSpecs,omitempty" name:"SidecarSpecs"`

	// 安全特性
	Security *CloudBaseSecurityContext `json:"Security,omitempty" name:"Security"`

	// 服务磁盘挂载
	ServiceVolumes []*CloudRunServiceVolume `json:"ServiceVolumes,omitempty" name:"ServiceVolumes"`

	// 是否创建JnsGw 0未传默认创建 1创建 2不创建
	IsCreateJnsGw *int64 `json:"IsCreateJnsGw,omitempty" name:"IsCreateJnsGw"`

	// 数据卷挂载参数
	ServiceVolumeMounts []*CloudBaseRunServiceVolumeMount `json:"ServiceVolumeMounts,omitempty" name:"ServiceVolumeMounts"`

	// 是否有Dockerfile：0-default has, 1-has, 2-has not
	HasDockerfile *int64 `json:"HasDockerfile,omitempty" name:"HasDockerfile"`

	// 基础镜像
	BaseImage *string `json:"BaseImage,omitempty" name:"BaseImage"`

	// 容器启动入口命令
	EntryPoint *string `json:"EntryPoint,omitempty" name:"EntryPoint"`

	// 仓库语言
	RepoLanguage *string `json:"RepoLanguage,omitempty" name:"RepoLanguage"`

	// 用户实际上传文件名（仅UploadType为jar/war时必填）
	UploadFilename *string `json:"UploadFilename,omitempty" name:"UploadFilename"`

	// 自动扩缩容策略组
	PolicyDetail []*HpaPolicy `json:"PolicyDetail,omitempty" name:"PolicyDetail"`
}

type CreateCloudBaseRunServerVersionRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 枚举（package/repository/image/jar/war)
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
	MountVolumeInfo []*CloudBaseRunVolumeMount `json:"MountVolumeInfo,omitempty" name:"MountVolumeInfo"`

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
	SidecarSpecs []*CloudBaseRunSideSpec `json:"SidecarSpecs,omitempty" name:"SidecarSpecs"`

	// 安全特性
	Security *CloudBaseSecurityContext `json:"Security,omitempty" name:"Security"`

	// 服务磁盘挂载
	ServiceVolumes []*CloudRunServiceVolume `json:"ServiceVolumes,omitempty" name:"ServiceVolumes"`

	// 是否创建JnsGw 0未传默认创建 1创建 2不创建
	IsCreateJnsGw *int64 `json:"IsCreateJnsGw,omitempty" name:"IsCreateJnsGw"`

	// 数据卷挂载参数
	ServiceVolumeMounts []*CloudBaseRunServiceVolumeMount `json:"ServiceVolumeMounts,omitempty" name:"ServiceVolumeMounts"`

	// 是否有Dockerfile：0-default has, 1-has, 2-has not
	HasDockerfile *int64 `json:"HasDockerfile,omitempty" name:"HasDockerfile"`

	// 基础镜像
	BaseImage *string `json:"BaseImage,omitempty" name:"BaseImage"`

	// 容器启动入口命令
	EntryPoint *string `json:"EntryPoint,omitempty" name:"EntryPoint"`

	// 仓库语言
	RepoLanguage *string `json:"RepoLanguage,omitempty" name:"RepoLanguage"`

	// 用户实际上传文件名（仅UploadType为jar/war时必填）
	UploadFilename *string `json:"UploadFilename,omitempty" name:"UploadFilename"`

	// 自动扩缩容策略组
	PolicyDetail []*HpaPolicy `json:"PolicyDetail,omitempty" name:"PolicyDetail"`
}

func (r *CreateCloudBaseRunServerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudBaseRunServerVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "UploadType")
	delete(f, "FlowRatio")
	delete(f, "Cpu")
	delete(f, "Mem")
	delete(f, "MinNum")
	delete(f, "MaxNum")
	delete(f, "PolicyType")
	delete(f, "PolicyThreshold")
	delete(f, "ContainerPort")
	delete(f, "ServerName")
	delete(f, "RepositoryType")
	delete(f, "DockerfilePath")
	delete(f, "BuildDir")
	delete(f, "EnvParams")
	delete(f, "Repository")
	delete(f, "Branch")
	delete(f, "VersionRemark")
	delete(f, "PackageName")
	delete(f, "PackageVersion")
	delete(f, "ImageInfo")
	delete(f, "CodeDetail")
	delete(f, "ImageSecretInfo")
	delete(f, "ImagePullSecret")
	delete(f, "CustomLogs")
	delete(f, "InitialDelaySeconds")
	delete(f, "MountVolumeInfo")
	delete(f, "AccessType")
	delete(f, "EsInfo")
	delete(f, "EnableUnion")
	delete(f, "OperatorRemark")
	delete(f, "ServerPath")
	delete(f, "ImageReuseKey")
	delete(f, "SidecarSpecs")
	delete(f, "Security")
	delete(f, "ServiceVolumes")
	delete(f, "IsCreateJnsGw")
	delete(f, "ServiceVolumeMounts")
	delete(f, "HasDockerfile")
	delete(f, "BaseImage")
	delete(f, "EntryPoint")
	delete(f, "RepoLanguage")
	delete(f, "UploadFilename")
	delete(f, "PolicyDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudBaseRunServerVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudBaseRunServerVersionResponseParams struct {
	// 状态(creating/succ)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 版本名称（只有Result为succ的时候，才会返回VersionName)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 操作记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunId *string `json:"RunId,omitempty" name:"RunId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCloudBaseRunServerVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudBaseRunServerVersionResponseParams `json:"Response"`
}

func (r *CreateCloudBaseRunServerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudBaseRunServerVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostingDomainRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostingDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Domain")
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHostingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostingDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHostingDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateHostingDomainResponseParams `json:"Response"`
}

func (r *CreateHostingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostingDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePostpayPackageRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePostpayPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "WxAppId")
	delete(f, "Source")
	delete(f, "FreeQuota")
	delete(f, "EnvSource")
	delete(f, "Alias")
	delete(f, "Channel")
	delete(f, "ExtensionId")
	delete(f, "Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePostpayPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePostpayPackageResponseParams struct {
	// 后付费订单号
	TranId *string `json:"TranId,omitempty" name:"TranId"`

	// 环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePostpayPackageResponse struct {
	*tchttp.BaseResponse
	Response *CreatePostpayPackageResponseParams `json:"Response"`
}

func (r *CreatePostpayPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePostpayPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStandaloneGatewayRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网关名
	GatewayAlias *string `json:"GatewayAlias,omitempty" name:"GatewayAlias"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 网关描述
	GatewayDesc *string `json:"GatewayDesc,omitempty" name:"GatewayDesc"`

	// 网关套餐版本
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`
}

type CreateStandaloneGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网关名
	GatewayAlias *string `json:"GatewayAlias,omitempty" name:"GatewayAlias"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 网关描述
	GatewayDesc *string `json:"GatewayDesc,omitempty" name:"GatewayDesc"`

	// 网关套餐版本
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`
}

func (r *CreateStandaloneGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStandaloneGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GatewayAlias")
	delete(f, "VpcId")
	delete(f, "SubnetIds")
	delete(f, "GatewayDesc")
	delete(f, "PackageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStandaloneGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStandaloneGatewayResponseParams struct {
	// 网关名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStandaloneGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateStandaloneGatewayResponseParams `json:"Response"`
}

func (r *CreateStandaloneGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStandaloneGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStaticStoreRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitempty" name:"EnableUnion"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaticStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "EnableUnion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStaticStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStaticStoreResponseParams struct {
	// 创建静态资源结果(succ/fail)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *CreateStaticStoreResponseParams `json:"Response"`
}

func (r *CreateStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaticStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWxCloudBaseRunEnvRequestParams struct {
	// wx应用Id
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// 环境别名，要以a-z开头，不能包含 a-z,0-9,- 以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 用户享有的免费额度级别，目前只能为“basic”，不传该字段或该字段为空，标识不享受免费额度。
	FreeQuota *string `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// 订单标记。建议使用方统一转大小写之后再判断。
	// QuickStart：快速启动来源
	// Activity：活动来源
	Flag *string `json:"Flag,omitempty" name:"Flag"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网列表
	SubNetIds []*string `json:"SubNetIds,omitempty" name:"SubNetIds"`

	// 是否打开云调用
	IsOpenCloudInvoke *bool `json:"IsOpenCloudInvoke,omitempty" name:"IsOpenCloudInvoke"`

	// 创建来源：wechat | cloud
	Source *string `json:"Source,omitempty" name:"Source"`

	// 渠道：wechat | cloud
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type CreateWxCloudBaseRunEnvRequest struct {
	*tchttp.BaseRequest
	
	// wx应用Id
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// 环境别名，要以a-z开头，不能包含 a-z,0-9,- 以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 用户享有的免费额度级别，目前只能为“basic”，不传该字段或该字段为空，标识不享受免费额度。
	FreeQuota *string `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// 订单标记。建议使用方统一转大小写之后再判断。
	// QuickStart：快速启动来源
	// Activity：活动来源
	Flag *string `json:"Flag,omitempty" name:"Flag"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网列表
	SubNetIds []*string `json:"SubNetIds,omitempty" name:"SubNetIds"`

	// 是否打开云调用
	IsOpenCloudInvoke *bool `json:"IsOpenCloudInvoke,omitempty" name:"IsOpenCloudInvoke"`

	// 创建来源：wechat | cloud
	Source *string `json:"Source,omitempty" name:"Source"`

	// 渠道：wechat | cloud
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *CreateWxCloudBaseRunEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWxCloudBaseRunEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WxAppId")
	delete(f, "Alias")
	delete(f, "FreeQuota")
	delete(f, "Flag")
	delete(f, "VpcId")
	delete(f, "SubNetIds")
	delete(f, "IsOpenCloudInvoke")
	delete(f, "Source")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWxCloudBaseRunEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWxCloudBaseRunEnvResponseParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 后付费订单号
	TranId *string `json:"TranId,omitempty" name:"TranId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWxCloudBaseRunEnvResponse struct {
	*tchttp.BaseResponse
	Response *CreateWxCloudBaseRunEnvResponseParams `json:"Response"`
}

func (r *CreateWxCloudBaseRunEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWxCloudBaseRunEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWxCloudBaseRunServerDBClusterRequestParams struct {
	// 账户密码
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 微信appid
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// mysql内核版本，支持5.7,8.0
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 0: 非大小写敏感
	// 1: 大小写敏感
	// 默认 0
	LowerCaseTableName *string `json:"LowerCaseTableName,omitempty" name:"LowerCaseTableName"`
}

type CreateWxCloudBaseRunServerDBClusterRequest struct {
	*tchttp.BaseRequest
	
	// 账户密码
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 微信appid
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// mysql内核版本，支持5.7,8.0
	DbVersion *string `json:"DbVersion,omitempty" name:"DbVersion"`

	// 0: 非大小写敏感
	// 1: 大小写敏感
	// 默认 0
	LowerCaseTableName *string `json:"LowerCaseTableName,omitempty" name:"LowerCaseTableName"`
}

func (r *CreateWxCloudBaseRunServerDBClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWxCloudBaseRunServerDBClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountPassword")
	delete(f, "EnvId")
	delete(f, "WxAppId")
	delete(f, "DbVersion")
	delete(f, "LowerCaseTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWxCloudBaseRunServerDBClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWxCloudBaseRunServerDBClusterResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWxCloudBaseRunServerDBClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateWxCloudBaseRunServerDBClusterResponseParams `json:"Response"`
}

func (r *CreateWxCloudBaseRunServerDBClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWxCloudBaseRunServerDBClusterResponse) FromJsonString(s string) error {
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

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type DeleteCloudBaseProjectLatestVersionRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 项目名
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 是否保留资源
	KeepResource *bool `json:"KeepResource,omitempty" name:"KeepResource"`
}

type DeleteCloudBaseProjectLatestVersionRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 项目名
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 是否保留资源
	KeepResource *bool `json:"KeepResource,omitempty" name:"KeepResource"`
}

func (r *DeleteCloudBaseProjectLatestVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudBaseProjectLatestVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ProjectName")
	delete(f, "KeepResource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudBaseProjectLatestVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudBaseProjectLatestVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCloudBaseProjectLatestVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudBaseProjectLatestVersionResponseParams `json:"Response"`
}

func (r *DeleteCloudBaseProjectLatestVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudBaseProjectLatestVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudBaseRunServerVersionRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 是否删除服务，只有最后一个版本的时候，才生效。
	IsDeleteServer *bool `json:"IsDeleteServer,omitempty" name:"IsDeleteServer"`

	// 只有删除服务的时候，才会起作用
	IsDeleteImage *bool `json:"IsDeleteImage,omitempty" name:"IsDeleteImage"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

type DeleteCloudBaseRunServerVersionRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 是否删除服务，只有最后一个版本的时候，才生效。
	IsDeleteServer *bool `json:"IsDeleteServer,omitempty" name:"IsDeleteServer"`

	// 只有删除服务的时候，才会起作用
	IsDeleteImage *bool `json:"IsDeleteImage,omitempty" name:"IsDeleteImage"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

func (r *DeleteCloudBaseRunServerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudBaseRunServerVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "VersionName")
	delete(f, "IsDeleteServer")
	delete(f, "IsDeleteImage")
	delete(f, "OperatorRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudBaseRunServerVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudBaseRunServerVersionResponseParams struct {
	// 返回结果，succ为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCloudBaseRunServerVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudBaseRunServerVersionResponseParams `json:"Response"`
}

func (r *DeleteCloudBaseRunServerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudBaseRunServerVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEndUserRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 用户列表，每一项都是uuid
	UserList []*string `json:"UserList,omitempty" name:"UserList"`
}

type DeleteEndUserRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 用户列表，每一项都是uuid
	UserList []*string `json:"UserList,omitempty" name:"UserList"`
}

func (r *DeleteEndUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEndUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "UserList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEndUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEndUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEndUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEndUserResponseParams `json:"Response"`
}

func (r *DeleteEndUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEndUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWxGatewayRouteRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	GatewayRouteName *string `json:"GatewayRouteName,omitempty" name:"GatewayRouteName"`
}

type DeleteWxGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	GatewayRouteName *string `json:"GatewayRouteName,omitempty" name:"GatewayRouteName"`
}

func (r *DeleteWxGatewayRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWxGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GatewayRouteName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWxGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWxGatewayRouteResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWxGatewayRouteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWxGatewayRouteResponseParams `json:"Response"`
}

func (r *DeleteWxGatewayRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWxGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActivityInfoRequestParams struct {
	// 活动id列表
	ActivityIdList []*int64 `json:"ActivityIdList,omitempty" name:"ActivityIdList"`
}

type DescribeActivityInfoRequest struct {
	*tchttp.BaseRequest
	
	// 活动id列表
	ActivityIdList []*int64 `json:"ActivityIdList,omitempty" name:"ActivityIdList"`
}

func (r *DescribeActivityInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActivityInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActivityInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActivityInfoResponseParams struct {
	// 活动详情
	ActivityInfoList []*ActivityInfoItem `json:"ActivityInfoList,omitempty" name:"ActivityInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeActivityInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActivityInfoResponseParams `json:"Response"`
}

func (r *DescribeActivityInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActivityInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActivityRecordRequestParams struct {
	// 渠道加密token
	ChannelToken *string `json:"ChannelToken,omitempty" name:"ChannelToken"`

	// 渠道来源，每个来源对应不同secretKey
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 活动id列表
	ActivityIdList []*int64 `json:"ActivityIdList,omitempty" name:"ActivityIdList"`

	// 过滤状态码，已废弃
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 状态码过滤数组，空数组时不过滤
	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`

	// 根据是否软删除进行过滤，[0]未删除, [1] 删除，不传不过滤
	IsDeletedList []*int64 `json:"IsDeletedList,omitempty" name:"IsDeletedList"`
}

type DescribeActivityRecordRequest struct {
	*tchttp.BaseRequest
	
	// 渠道加密token
	ChannelToken *string `json:"ChannelToken,omitempty" name:"ChannelToken"`

	// 渠道来源，每个来源对应不同secretKey
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 活动id列表
	ActivityIdList []*int64 `json:"ActivityIdList,omitempty" name:"ActivityIdList"`

	// 过滤状态码，已废弃
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 状态码过滤数组，空数组时不过滤
	Statuses []*int64 `json:"Statuses,omitempty" name:"Statuses"`

	// 根据是否软删除进行过滤，[0]未删除, [1] 删除，不传不过滤
	IsDeletedList []*int64 `json:"IsDeletedList,omitempty" name:"IsDeletedList"`
}

func (r *DescribeActivityRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActivityRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelToken")
	delete(f, "Channel")
	delete(f, "ActivityIdList")
	delete(f, "Status")
	delete(f, "Statuses")
	delete(f, "IsDeletedList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActivityRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActivityRecordResponseParams struct {
	// 活动记录详情
	ActivityRecords []*ActivityRecordItem `json:"ActivityRecords,omitempty" name:"ActivityRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeActivityRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActivityRecordResponseParams `json:"Response"`
}

func (r *DescribeActivityRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActivityRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthDomainsRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuthDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthDomainsResponseParams struct {
	// 安全域名列表列表
	Domains []*AuthDomain `json:"Domains,omitempty" name:"Domains"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuthDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuthDomainsResponseParams `json:"Response"`
}

func (r *DescribeAuthDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaasPackageListRequestParams struct {
	// tcb产品套餐ID，不填拉取全量package信息。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 套餐归属方，填写后只返回对应的套餐 包含miniapp与qcloud两种 默认为miniapp
	Source *string `json:"Source,omitempty" name:"Source"`

	// 套餐归属环境渠道
	EnvChannel *string `json:"EnvChannel,omitempty" name:"EnvChannel"`

	// 拉取套餐用途：
	// 1）new 新购
	// 2）modify变配
	// 3）renew续费
	TargetAction *string `json:"TargetAction,omitempty" name:"TargetAction"`

	// 预留字段，同一商品会对应多个类型套餐，对指标有不同侧重。
	// 计算型calculation
	// 流量型flux
	// 容量型capactiy
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 类型分组过滤。默认为["default"]
	PackageTypeList []*string `json:"PackageTypeList,omitempty" name:"PackageTypeList"`

	// 付费渠道，与回包billTags中的计费参数相关，不填返回默认值。
	PaymentChannel *string `json:"PaymentChannel,omitempty" name:"PaymentChannel"`
}

type DescribeBaasPackageListRequest struct {
	*tchttp.BaseRequest
	
	// tcb产品套餐ID，不填拉取全量package信息。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 套餐归属方，填写后只返回对应的套餐 包含miniapp与qcloud两种 默认为miniapp
	Source *string `json:"Source,omitempty" name:"Source"`

	// 套餐归属环境渠道
	EnvChannel *string `json:"EnvChannel,omitempty" name:"EnvChannel"`

	// 拉取套餐用途：
	// 1）new 新购
	// 2）modify变配
	// 3）renew续费
	TargetAction *string `json:"TargetAction,omitempty" name:"TargetAction"`

	// 预留字段，同一商品会对应多个类型套餐，对指标有不同侧重。
	// 计算型calculation
	// 流量型flux
	// 容量型capactiy
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 类型分组过滤。默认为["default"]
	PackageTypeList []*string `json:"PackageTypeList,omitempty" name:"PackageTypeList"`

	// 付费渠道，与回包billTags中的计费参数相关，不填返回默认值。
	PaymentChannel *string `json:"PaymentChannel,omitempty" name:"PaymentChannel"`
}

func (r *DescribeBaasPackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaasPackageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageName")
	delete(f, "EnvId")
	delete(f, "Source")
	delete(f, "EnvChannel")
	delete(f, "TargetAction")
	delete(f, "GroupName")
	delete(f, "PackageTypeList")
	delete(f, "PaymentChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaasPackageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaasPackageListResponseParams struct {
	// 套餐列表
	PackageList []*BaasPackageInfo `json:"PackageList,omitempty" name:"PackageList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBaasPackageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaasPackageListResponseParams `json:"Response"`
}

func (r *DescribeBaasPackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaasPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseBuildServiceRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// build类型,枚举值有: cloudbaserun, framework-ci
	CIBusiness *string `json:"CIBusiness,omitempty" name:"CIBusiness"`

	// 服务版本
	ServiceVersion *string `json:"ServiceVersion,omitempty" name:"ServiceVersion"`

	// 文件后缀
	Suffix *string `json:"Suffix,omitempty" name:"Suffix"`
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

	// 文件后缀
	Suffix *string `json:"Suffix,omitempty" name:"Suffix"`
}

func (r *DescribeCloudBaseBuildServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseBuildServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServiceName")
	delete(f, "CIBusiness")
	delete(f, "ServiceVersion")
	delete(f, "Suffix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseBuildServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseBuildServiceResponseParams struct {
	// 上传url
	UploadUrl *string `json:"UploadUrl,omitempty" name:"UploadUrl"`

	// 上传heder
	UploadHeaders []*KVPair `json:"UploadHeaders,omitempty" name:"UploadHeaders"`

	// 包名
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 包版本
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 下载Httpheader
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadHeaders []*KVPair `json:"DownloadHeaders,omitempty" name:"DownloadHeaders"`

	// 下载链接是否过期
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutDate *bool `json:"OutDate,omitempty" name:"OutDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseBuildServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseBuildServiceResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseBuildServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseBuildServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseProjectLatestVersionListRequestParams struct {
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
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// ci的id
	CiId *string `json:"CiId,omitempty" name:"CiId"`
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
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// ci的id
	CiId *string `json:"CiId,omitempty" name:"CiId"`
}

func (r *DescribeCloudBaseProjectLatestVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseProjectLatestVersionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "PageSize")
	delete(f, "EnvId")
	delete(f, "ProjectName")
	delete(f, "ProjectType")
	delete(f, "Tags")
	delete(f, "CiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseProjectLatestVersionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseProjectLatestVersionListResponseParams struct {
	// 项目列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectList []*CloudBaseProjectVersion `json:"ProjectList,omitempty" name:"ProjectList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseProjectLatestVersionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseProjectLatestVersionListResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseProjectLatestVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseProjectLatestVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseProjectVersionListRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 第几页,从0开始
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 起始时间 2021-03-27 12:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 终止时间 2021-03-27 12:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeCloudBaseProjectVersionListRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 第几页,从0开始
	PageNum *uint64 `json:"PageNum,omitempty" name:"PageNum"`

	// 起始时间 2021-03-27 12:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 终止时间 2021-03-27 12:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeCloudBaseProjectVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseProjectVersionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ProjectName")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseProjectVersionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseProjectVersionListResponseParams struct {
	// 版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectVersions []*CloudBaseProjectVersion `json:"ProjectVersions,omitempty" name:"ProjectVersions"`

	// 总个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseProjectVersionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseProjectVersionListResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseProjectVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseProjectVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunAllVpcsRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

type DescribeCloudBaseRunAllVpcsRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeCloudBaseRunAllVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunAllVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunAllVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunAllVpcsResponseParams struct {
	// 所有vpcid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vpcs []*string `json:"Vpcs,omitempty" name:"Vpcs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunAllVpcsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunAllVpcsResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunAllVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunAllVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunConfForGateWayRequestParams struct {
	// 环境ID
	EnvID *string `json:"EnvID,omitempty" name:"EnvID"`

	// vpc信息
	VpcID *string `json:"VpcID,omitempty" name:"VpcID"`
}

type DescribeCloudBaseRunConfForGateWayRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvID *string `json:"EnvID,omitempty" name:"EnvID"`

	// vpc信息
	VpcID *string `json:"VpcID,omitempty" name:"VpcID"`
}

func (r *DescribeCloudBaseRunConfForGateWayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunConfForGateWayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvID")
	delete(f, "VpcID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunConfForGateWayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunConfForGateWayResponseParams struct {
	// 最近更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpTime *string `json:"LastUpTime,omitempty" name:"LastUpTime"`

	// 配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CloudBaseRunForGatewayConf `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunConfForGateWayResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunConfForGateWayResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunConfForGateWayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunConfForGateWayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunOneClickTaskExternalRequestParams struct {
	// 外部任务Id 最长64字节
	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
}

type DescribeCloudBaseRunOneClickTaskExternalRequest struct {
	*tchttp.BaseRequest
	
	// 外部任务Id 最长64字节
	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
}

func (r *DescribeCloudBaseRunOneClickTaskExternalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunOneClickTaskExternalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExternalId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunOneClickTaskExternalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunOneClickTaskExternalResponseParams struct {
	// 外部任务Id
	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`

	// 弃用
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 用户uin
	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 当前阶段
	// 微信云托管环境创建阶段：envStage
	// 存储资源创建阶段：storageStage
	// 服务创建阶段：serverStage
	Stage *string `json:"Stage,omitempty" name:"Stage"`

	// 状态
	// running
	// stopped
	// failed
	// finished
	Status *string `json:"Status,omitempty" name:"Status"`

	// 失败原因
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// 用户envId
	UserEnvId *string `json:"UserEnvId,omitempty" name:"UserEnvId"`

	// 创建时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 步骤信息
	Steps []*OneClickTaskStepInfo `json:"Steps,omitempty" name:"Steps"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunOneClickTaskExternalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunOneClickTaskExternalResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunOneClickTaskExternalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunOneClickTaskExternalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunOperationTypesRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称，精确匹配
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
}

type DescribeCloudBaseRunOperationTypesRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称，精确匹配
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
}

func (r *DescribeCloudBaseRunOperationTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunOperationTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunOperationTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunOperationTypesResponseParams struct {
	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action []*string `json:"Action,omitempty" name:"Action"`

	// 服务名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName []*string `json:"ServerName,omitempty" name:"ServerName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunOperationTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunOperationTypesResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunOperationTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunOperationTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunPodListRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 分页限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 容器状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 容器名
	PodName *string `json:"PodName,omitempty" name:"PodName"`
}

type DescribeCloudBaseRunPodListRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 分页限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 容器状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 容器名
	PodName *string `json:"PodName,omitempty" name:"PodName"`
}

func (r *DescribeCloudBaseRunPodListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunPodListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "VersionName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "PodName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunPodListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunPodListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunPodListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunPodListResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunPodListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunPodListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunResourceForExtendRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunResourceForExtendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunResourceForExtendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunResourceForExtendResponseParams struct {
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
	SubnetIds []*CloudBaseRunVpcSubnet `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunResourceForExtendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunResourceForExtendResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunResourceForExtendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunResourceForExtendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunResourceRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunResourceResponseParams struct {
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
	SubnetIds []*CloudBaseRunVpcSubnet `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunResourceResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunServerDomainNameRequestParams struct {
	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 环境Id
	UserEnvId *string `json:"UserEnvId,omitempty" name:"UserEnvId"`

	// 用户Uin
	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`

	// 外部Id
	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
}

type DescribeCloudBaseRunServerDomainNameRequest struct {
	*tchttp.BaseRequest
	
	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 环境Id
	UserEnvId *string `json:"UserEnvId,omitempty" name:"UserEnvId"`

	// 用户Uin
	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`

	// 外部Id
	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
}

func (r *DescribeCloudBaseRunServerDomainNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunServerDomainNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServerName")
	delete(f, "UserEnvId")
	delete(f, "UserUin")
	delete(f, "ExternalId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunServerDomainNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunServerDomainNameResponseParams struct {
	// 公网服务域名
	PublicDomain *string `json:"PublicDomain,omitempty" name:"PublicDomain"`

	// 内部服务域名
	InternalDomain *string `json:"InternalDomain,omitempty" name:"InternalDomain"`

	// 弃用
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunServerDomainNameResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunServerDomainNameResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunServerDomainNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunServerDomainNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunServerRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 版本名字（精确匹配）
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

type DescribeCloudBaseRunServerRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 版本名字（精确匹配）
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

func (r *DescribeCloudBaseRunServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "VersionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunServerResponseParams struct {
	// 个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionItems []*CloudBaseRunServerVersionItem `json:"VersionItems,omitempty" name:"VersionItems"`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 是否对于外网开放
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 镜像仓库
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageRepo *string `json:"ImageRepo,omitempty" name:"ImageRepo"`

	// 流量配置的类型（FLOW,URL_PARAMS)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`

	// 服务创建类型，默认为空，一键部署为oneclick
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunServerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunServerResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunServerVersionRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunServerVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "VersionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunServerVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunServerVersionResponseParams struct {
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

	// Repo的类型(gitlab/github/coding)
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
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

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

	// 是否有Dockerfile：0-default has, 1-has, 2-has not
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasDockerfile *int64 `json:"HasDockerfile,omitempty" name:"HasDockerfile"`

	// 基础镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseImage *string `json:"BaseImage,omitempty" name:"BaseImage"`

	// 容器启动入口命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntryPoint *string `json:"EntryPoint,omitempty" name:"EntryPoint"`

	// 仓库语言
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoLanguage *string `json:"RepoLanguage,omitempty" name:"RepoLanguage"`

	// 自动扩缩容策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDetail []*HpaPolicy `json:"PolicyDetail,omitempty" name:"PolicyDetail"`

	// Tke集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TkeClusterInfo *TkeClusterInfo `json:"TkeClusterInfo,omitempty" name:"TkeClusterInfo"`

	// 版本工作负载类型；deployment/deamonset
	// 注意：此字段可能返回 null，表示取不到有效值。
	TkeWorkloadType *string `json:"TkeWorkloadType,omitempty" name:"TkeWorkloadType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunServerVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunServerVersionResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunServerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunServerVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunVersionRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

type DescribeCloudBaseRunVersionRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

func (r *DescribeCloudBaseRunVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "VersionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunVersionResponseParams struct {
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
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

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

	// 扩缩容策略详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDetail []*HpaPolicy `json:"PolicyDetail,omitempty" name:"PolicyDetail"`

	// Cpu的Request值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *float64 `json:"Cpu,omitempty" name:"Cpu"`

	// Mem的Request值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mem *float64 `json:"Mem,omitempty" name:"Mem"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunVersionResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunVersionRsByConditionRequestParams struct {

}

type DescribeCloudBaseRunVersionRsByConditionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCloudBaseRunVersionRsByConditionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunVersionRsByConditionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunVersionRsByConditionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunVersionRsByConditionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunVersionRsByConditionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunVersionRsByConditionResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunVersionRsByConditionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunVersionRsByConditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunVersionSnapshotRequestParams struct {
	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 版本历史名
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 偏移量。默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制大小。默认10，最大20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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

	// 偏移量。默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 限制大小。默认10，最大20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCloudBaseRunVersionSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunVersionSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServerName")
	delete(f, "VersionName")
	delete(f, "EnvId")
	delete(f, "SnapshotName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunVersionSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunVersionSnapshotResponseParams struct {
	// 版本历史
	// 注意：此字段可能返回 null，表示取不到有效值。
	Snapshots []*CloudRunServiceSimpleVersionSnapshot `json:"Snapshots,omitempty" name:"Snapshots"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunVersionSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunVersionSnapshotResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunVersionSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunVersionSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurveDataRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// <li> 指标名: </li>
	// <li> StorageRead: 存储读请求次数 </li>
	// <li> StorageWrite: 存储写请求次数 </li>
	// <li> StorageCdnOriginFlux: CDN回源流量, 单位字节 </li>
	// <li> CDNFlux: CDN回源流量, 单位字节 </li>
	// <li> FunctionInvocation: 云函数调用次数 </li>
	// <li> FunctionGBs: 云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFlux: 云函数流量, 单位千字节(KB) </li>
	// <li> FunctionError: 云函数调用错误次数 </li>
	// <li> FunctionDuration: 云函数运行时间, 单位毫秒 </li>
	// <li> DbRead: 数据库读请求数 </li>
	// <li> DbWrite: 数据库写请求数 </li>
	// <li> DbCostTime10ms: 数据库耗时在10ms~50ms请求数 </li>
	// <li> DbCostTime50ms: 数据库耗时在50ms~100ms请求数 </li>
	// <li> DbCostTime100ms: 数据库耗时在100ms以上请求数 </li>
	// <li> TkeCpuRatio: 容器CPU占用率 </li>
	// <li> TkeMemRatio: 容器内存占用率 </li>
	// <li> TkeCpuUsed: 容器CPU使用量 </li>
	// <li> TkeMemUsed: 容器内存使用量 </li>
	// <li> TkeInvokeNum: 调用量 </li>
	// <li> FunctionConcurrentExecutions: 云函数并发执行个数</li>
	// <li> FunctionIdleProvisioned: 云函数预置并发闲置量 </li>
	// <li> FunctionConcurrencyMemoryMB: 云函数并发执行内存量 </li>
	// <li> FunctionThrottle: 云函数受限次数 </li>
	// <li> FunctionProvisionedConcurrency: 云函数预置并发 </li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 开始时间，如2018-08-24 10:50:00, 开始时间需要早于结束时间至少五分钟(原因是因为目前统计粒度最小是5分钟).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如2018-08-24 10:50:00, 结束时间需要晚于开始时间至少五分钟(原因是因为目前统计粒度最小是5分钟)..
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源ID, 目前仅对云函数、容器托管相关的指标有意义。云函数(FunctionInvocation, FunctionGBs, FunctionFlux, FunctionError, FunctionDuration)、容器托管（服务名称）, 如果想查询某个云函数的指标则在ResourceId中传入函数名; 如果只想查询整个namespace的指标, 则留空或不传.如果想查询数据库某个集合相关信息，传入集合名称
	ResourceID *string `json:"ResourceID,omitempty" name:"ResourceID"`
}

type DescribeCurveDataRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// <li> 指标名: </li>
	// <li> StorageRead: 存储读请求次数 </li>
	// <li> StorageWrite: 存储写请求次数 </li>
	// <li> StorageCdnOriginFlux: CDN回源流量, 单位字节 </li>
	// <li> CDNFlux: CDN回源流量, 单位字节 </li>
	// <li> FunctionInvocation: 云函数调用次数 </li>
	// <li> FunctionGBs: 云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFlux: 云函数流量, 单位千字节(KB) </li>
	// <li> FunctionError: 云函数调用错误次数 </li>
	// <li> FunctionDuration: 云函数运行时间, 单位毫秒 </li>
	// <li> DbRead: 数据库读请求数 </li>
	// <li> DbWrite: 数据库写请求数 </li>
	// <li> DbCostTime10ms: 数据库耗时在10ms~50ms请求数 </li>
	// <li> DbCostTime50ms: 数据库耗时在50ms~100ms请求数 </li>
	// <li> DbCostTime100ms: 数据库耗时在100ms以上请求数 </li>
	// <li> TkeCpuRatio: 容器CPU占用率 </li>
	// <li> TkeMemRatio: 容器内存占用率 </li>
	// <li> TkeCpuUsed: 容器CPU使用量 </li>
	// <li> TkeMemUsed: 容器内存使用量 </li>
	// <li> TkeInvokeNum: 调用量 </li>
	// <li> FunctionConcurrentExecutions: 云函数并发执行个数</li>
	// <li> FunctionIdleProvisioned: 云函数预置并发闲置量 </li>
	// <li> FunctionConcurrencyMemoryMB: 云函数并发执行内存量 </li>
	// <li> FunctionThrottle: 云函数受限次数 </li>
	// <li> FunctionProvisionedConcurrency: 云函数预置并发 </li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 开始时间，如2018-08-24 10:50:00, 开始时间需要早于结束时间至少五分钟(原因是因为目前统计粒度最小是5分钟).
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如2018-08-24 10:50:00, 结束时间需要晚于开始时间至少五分钟(原因是因为目前统计粒度最小是5分钟)..
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 资源ID, 目前仅对云函数、容器托管相关的指标有意义。云函数(FunctionInvocation, FunctionGBs, FunctionFlux, FunctionError, FunctionDuration)、容器托管（服务名称）, 如果想查询某个云函数的指标则在ResourceId中传入函数名; 如果只想查询整个namespace的指标, 则留空或不传.如果想查询数据库某个集合相关信息，传入集合名称
	ResourceID *string `json:"ResourceID,omitempty" name:"ResourceID"`
}

func (r *DescribeCurveDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurveDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "MetricName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ResourceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCurveDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurveDataResponseParams struct {
	// 开始时间, 会根据数据的统计周期进行取整.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间, 会根据数据的统计周期进行取整.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指标名.
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计周期(单位秒), 当时间区间为1天内, 统计周期为5分钟; 当时间区间选择为1天以上, 15天以下, 统计周期为1小时; 当时间区间选择为15天以上, 180天以下, 统计周期为1天.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 有效的监控数据, 每个有效监控数据的上报时间可以从时间数组中的对应位置上获取到.
	Values []*int64 `json:"Values,omitempty" name:"Values"`

	// 时间数据, 标识监控数据Values中的点是哪个时间段上报的.
	Time []*int64 `json:"Time,omitempty" name:"Time"`

	// 有效的监控数据, 每个有效监控数据的上报时间可以从时间数组中的对应位置上获取到.
	NewValues *float64 `json:"NewValues,omitempty" name:"NewValues"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCurveDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCurveDataResponseParams `json:"Response"`
}

func (r *DescribeCurveDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurveDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseACLRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitempty" name:"CollectionName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseACLResponseParams struct {
	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitempty" name:"AclTag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseACLResponseParams `json:"Response"`
}

func (r *DescribeDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDownloadFileRequestParams struct {
	// 代码uri，格式如：extension://abcdefhhxxx.zip，对应 DescribeExtensionUploadInfo 接口的返回值
	CodeUri *string `json:"CodeUri,omitempty" name:"CodeUri"`
}

type DescribeDownloadFileRequest struct {
	*tchttp.BaseRequest
	
	// 代码uri，格式如：extension://abcdefhhxxx.zip，对应 DescribeExtensionUploadInfo 接口的返回值
	CodeUri *string `json:"CodeUri,omitempty" name:"CodeUri"`
}

func (r *DescribeDownloadFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDownloadFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CodeUri")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDownloadFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDownloadFileResponseParams struct {
	// 文件路径，该字段已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 加密key，用于计算下载加密文件的header。参考SSE-C https://cloud.tencent.com/document/product/436/7728#sse-c
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomKey *string `json:"CustomKey,omitempty" name:"CustomKey"`

	// 下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDownloadFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDownloadFileResponseParams `json:"Response"`
}

func (r *DescribeDownloadFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDownloadFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndUserLoginStatisticRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 终端用户来源
	// <li> qcloud </li>
	// <li>miniapp</li>
	Source *string `json:"Source,omitempty" name:"Source"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndUserLoginStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEndUserLoginStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndUserLoginStatisticResponseParams struct {
	// 环境终端用户新增与登录统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginStatistics []*LoginStatistic `json:"LoginStatistics,omitempty" name:"LoginStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEndUserLoginStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEndUserLoginStatisticResponseParams `json:"Response"`
}

func (r *DescribeEndUserLoginStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndUserLoginStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndUserStatisticRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndUserStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEndUserStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndUserStatisticResponseParams struct {
	// 终端用户各平台统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlatformStatistics []*PlatformStatistic `json:"PlatformStatistics,omitempty" name:"PlatformStatistics"`

	// 终端用户总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEndUserStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEndUserStatisticResponseParams `json:"Response"`
}

func (r *DescribeEndUserStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndUserStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndUsersRequestParams struct {
	// 开发者的环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 可选参数，偏移量，默认 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 可选参数，拉取数量，默认 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 按照 uuid 列表过滤，最大个数为100
	UUIds []*string `json:"UUIds,omitempty" name:"UUIds"`
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
	UUIds []*string `json:"UUIds,omitempty" name:"UUIds"`
}

func (r *DescribeEndUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "UUIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEndUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEndUsersResponseParams struct {
	// 用户总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 用户列表
	Users []*EndUserInfo `json:"Users,omitempty" name:"Users"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEndUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEndUsersResponseParams `json:"Response"`
}

func (r *DescribeEndUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEndUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvDealRegionRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 订单类型：
	// ENV_PREPAY_MINIAPP= 预付费环境(微信小程序)
	// ENV_PREPAY_CLOUD= 预付费环境(腾讯云)
	// ENV_POSTPAY = 后付费环境
	// HOSTING_PREPAY = 预付费静态托管
	// PACKAGE=套餐包
	DealType *string `json:"DealType,omitempty" name:"DealType"`

	// 下单类型：
	// CREATE = 新购
	// RENEW = 续费
	// MODIFY = 套餐调整(升级/降级)
	// REFUND = 退费
	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`

	// 下单地域：
	// ap-guangzhou = 广州地域
	// ap-shanghai = 上海地域
	// ap-beijing = 北京地域
	DealRegion *string `json:"DealRegion,omitempty" name:"DealRegion"`
}

type DescribeEnvDealRegionRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 订单类型：
	// ENV_PREPAY_MINIAPP= 预付费环境(微信小程序)
	// ENV_PREPAY_CLOUD= 预付费环境(腾讯云)
	// ENV_POSTPAY = 后付费环境
	// HOSTING_PREPAY = 预付费静态托管
	// PACKAGE=套餐包
	DealType *string `json:"DealType,omitempty" name:"DealType"`

	// 下单类型：
	// CREATE = 新购
	// RENEW = 续费
	// MODIFY = 套餐调整(升级/降级)
	// REFUND = 退费
	DealAction *string `json:"DealAction,omitempty" name:"DealAction"`

	// 下单地域：
	// ap-guangzhou = 广州地域
	// ap-shanghai = 上海地域
	// ap-beijing = 北京地域
	DealRegion *string `json:"DealRegion,omitempty" name:"DealRegion"`
}

func (r *DescribeEnvDealRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvDealRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "DealType")
	delete(f, "DealAction")
	delete(f, "DealRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvDealRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvDealRegionResponseParams struct {
	// 下单region
	Region *string `json:"Region,omitempty" name:"Region"`

	// 下单zone
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 下单regionId
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 下单zoneId
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvDealRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvDealRegionResponseParams `json:"Response"`
}

func (r *DescribeEnvDealRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvDealRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvFreeQuotaRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 资源类型：可选值：CDN, COS, FLEXDB, HOSTING, SCF
	// 不传则返回全部资源指标
	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes"`
}

type DescribeEnvFreeQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 资源类型：可选值：CDN, COS, FLEXDB, HOSTING, SCF
	// 不传则返回全部资源指标
	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes"`
}

func (r *DescribeEnvFreeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvFreeQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ResourceTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvFreeQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvFreeQuotaResponseParams struct {
	// 免费抵扣配额详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaItems []*PostpayEnvQuota `json:"QuotaItems,omitempty" name:"QuotaItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvFreeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvFreeQuotaResponseParams `json:"Response"`
}

func (r *DescribeEnvFreeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvFreeQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvLimitRequestParams struct {

}

type DescribeEnvLimitRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEnvLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvLimitResponseParams struct {
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

	// 微信网关体验版可购买月份数
	MaxFreeTrialNum *int64 `json:"MaxFreeTrialNum,omitempty" name:"MaxFreeTrialNum"`

	// 微信网关体验版已购买月份数
	CurrentFreeTrialNum *int64 `json:"CurrentFreeTrialNum,omitempty" name:"CurrentFreeTrialNum"`

	// 转支付限额总数
	ChangePayTotal *int64 `json:"ChangePayTotal,omitempty" name:"ChangePayTotal"`

	// 当前已用转支付次数
	CurrentChangePayTotal *int64 `json:"CurrentChangePayTotal,omitempty" name:"CurrentChangePayTotal"`

	// 转支付每月限额
	ChangePayMonthly *int64 `json:"ChangePayMonthly,omitempty" name:"ChangePayMonthly"`

	// 本月已用转支付额度
	CurrentChangePayMonthly *int64 `json:"CurrentChangePayMonthly,omitempty" name:"CurrentChangePayMonthly"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvLimitResponseParams `json:"Response"`
}

func (r *DescribeEnvLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvPostpaidDeductRequestParams struct {
	// 资源方列表
	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes"`

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeEnvPostpaidDeductRequest struct {
	*tchttp.BaseRequest
	
	// 资源方列表
	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes"`

	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeEnvPostpaidDeductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvPostpaidDeductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceTypes")
	delete(f, "EnvId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvPostpaidDeductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvPostpaidDeductResponseParams struct {
	// 指标抵扣详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostPaidEnvDeductInfoList []*PostPaidEnvDeductInfo `json:"PostPaidEnvDeductInfoList,omitempty" name:"PostPaidEnvDeductInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvPostpaidDeductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvPostpaidDeductResponseParams `json:"Response"`
}

func (r *DescribeEnvPostpaidDeductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvPostpaidDeductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvsRequestParams struct {
	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitempty" name:"Channels"`
}

type DescribeEnvsRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitempty" name:"Channels"`
}

func (r *DescribeEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "IsVisible")
	delete(f, "Channels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvsResponseParams struct {
	// 环境信息列表
	EnvList []*EnvInfo `json:"EnvList,omitempty" name:"EnvList"`

	// 环境个数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvsResponseParams `json:"Response"`
}

func (r *DescribeEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionUploadInfoRequestParams struct {
	// 待上传的文件
	ExtensionFiles []*ExtensionFile `json:"ExtensionFiles,omitempty" name:"ExtensionFiles"`
}

type DescribeExtensionUploadInfoRequest struct {
	*tchttp.BaseRequest
	
	// 待上传的文件
	ExtensionFiles []*ExtensionFile `json:"ExtensionFiles,omitempty" name:"ExtensionFiles"`
}

func (r *DescribeExtensionUploadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionUploadInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExtensionFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtensionUploadInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtensionUploadInfoResponseParams struct {
	// 待上传文件的信息数组
	FilesData []*ExtensionFileInfo `json:"FilesData,omitempty" name:"FilesData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExtensionUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtensionUploadInfoResponseParams `json:"Response"`
}

func (r *DescribeExtensionUploadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtensionUploadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtraPkgBillingInfoRequestParams struct {
	// 已购买增值包的环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtraPkgBillingInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtraPkgBillingInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtraPkgBillingInfoResponseParams struct {
	// 增值包计费信息列表
	EnvInfoList []*EnvBillingInfoItem `json:"EnvInfoList,omitempty" name:"EnvInfoList"`

	// 增值包数目
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExtraPkgBillingInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtraPkgBillingInfoResponseParams `json:"Response"`
}

func (r *DescribeExtraPkgBillingInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtraPkgBillingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostingDomainTaskRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

type DescribeHostingDomainTaskRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeHostingDomainTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostingDomainTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostingDomainTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostingDomainTaskResponseParams struct {
	// todo/doing/done/error
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostingDomainTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostingDomainTaskResponseParams `json:"Response"`
}

func (r *DescribeHostingDomainTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostingDomainTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostpayFreeQuotasRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostpayFreeQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePostpayFreeQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostpayFreeQuotasResponseParams struct {
	// 免费量资源信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreequotaInfoList []*FreequotaInfo `json:"FreequotaInfoList,omitempty" name:"FreequotaInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePostpayFreeQuotasResponse struct {
	*tchttp.BaseResponse
	Response *DescribePostpayFreeQuotasResponseParams `json:"Response"`
}

func (r *DescribePostpayFreeQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostpayFreeQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostpayPackageFreeQuotasRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 免费额度类型标识
	FreeQuotaType *string `json:"FreeQuotaType,omitempty" name:"FreeQuotaType"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostpayPackageFreeQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "FreeQuotaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePostpayPackageFreeQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostpayPackageFreeQuotasResponseParams struct {
	// 免费量资源信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageFreeQuotaInfos []*PackageFreeQuotaInfo `json:"PackageFreeQuotaInfos,omitempty" name:"PackageFreeQuotaInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePostpayPackageFreeQuotasResponse struct {
	*tchttp.BaseResponse
	Response *DescribePostpayPackageFreeQuotasResponseParams `json:"Response"`
}

func (r *DescribePostpayPackageFreeQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostpayPackageFreeQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaDataRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "MetricName")
	delete(f, "ResourceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaDataResponseParams struct {
	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标的值
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// 指标的附加值信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubValue *string `json:"SubValue,omitempty" name:"SubValue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeQuotaDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotaDataResponseParams `json:"Response"`
}

func (r *DescribeQuotaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsQuotasRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmsQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsQuotasResponseParams struct {
	// 短信免费量信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmsFreeQuotaList []*SmsFreeQuota `json:"SmsFreeQuotaList,omitempty" name:"SmsFreeQuotaList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSmsQuotasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmsQuotasResponseParams `json:"Response"`
}

func (r *DescribeSmsQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecialCostItemsRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeSpecialCostItemsRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSpecialCostItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecialCostItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecialCostItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecialCostItemsResponseParams struct {
	// 1分钱抵扣详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialCostItems []*SpecialCostItem `json:"SpecialCostItems,omitempty" name:"SpecialCostItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpecialCostItemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecialCostItemsResponseParams `json:"Response"`
}

func (r *DescribeSpecialCostItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecialCostItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandaloneGatewayPackageRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 套餐版本，包含starter、basic、advanced、enterprise
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`
}

type DescribeStandaloneGatewayPackageRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 套餐版本，包含starter、basic、advanced、enterprise
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`
}

func (r *DescribeStandaloneGatewayPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandaloneGatewayPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "PackageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStandaloneGatewayPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandaloneGatewayPackageResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 套餐详情
	StandaloneGatewayPackageList []*StandaloneGatewayPackageInfo `json:"StandaloneGatewayPackageList,omitempty" name:"StandaloneGatewayPackageList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStandaloneGatewayPackageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStandaloneGatewayPackageResponseParams `json:"Response"`
}

func (r *DescribeStandaloneGatewayPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandaloneGatewayPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandaloneGatewayRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网关名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 网关别名
	GatewayAlias *string `json:"GatewayAlias,omitempty" name:"GatewayAlias"`
}

type DescribeStandaloneGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网关名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 网关别名
	GatewayAlias *string `json:"GatewayAlias,omitempty" name:"GatewayAlias"`
}

func (r *DescribeStandaloneGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandaloneGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GatewayName")
	delete(f, "GatewayAlias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStandaloneGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandaloneGatewayResponseParams struct {
	// 独立网关信息列表
	StandaloneGatewayList []*StandaloneGatewayInfo `json:"StandaloneGatewayList,omitempty" name:"StandaloneGatewayList"`

	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStandaloneGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStandaloneGatewayResponseParams `json:"Response"`
}

func (r *DescribeStandaloneGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandaloneGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserActivityInfoRequestParams struct {
	// 活动id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 渠道加密token
	ChannelToken *string `json:"ChannelToken,omitempty" name:"ChannelToken"`

	// 渠道来源，每个来源对应不同secretKey
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 团id, 1元钱裂变中活动团id不为空时根据团id来查询记录，为空时查询uin最新记录
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type DescribeUserActivityInfoRequest struct {
	*tchttp.BaseRequest
	
	// 活动id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 渠道加密token
	ChannelToken *string `json:"ChannelToken,omitempty" name:"ChannelToken"`

	// 渠道来源，每个来源对应不同secretKey
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 团id, 1元钱裂变中活动团id不为空时根据团id来查询记录，为空时查询uin最新记录
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeUserActivityInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserActivityInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityId")
	delete(f, "ChannelToken")
	delete(f, "Channel")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserActivityInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserActivityInfoResponseParams struct {
	// 自定义标记，1元钱裂变需求中即代指`团id`
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 自定义备注，1元钱裂变需求中返回`团列表`，uin列表通过","拼接
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 活动剩余时间，单位为s.1元钱裂变需求中即为 time(活动过期时间)-Now()), 过期后为0，即返回必为自然数
	ActivityTimeLeft *int64 `json:"ActivityTimeLeft,omitempty" name:"ActivityTimeLeft"`

	// 拼团剩余时间，单位为s.1元钱裂变需求中即为time(成团时间)+24H-Now()，过期后为0，即返回必为自然数
	GroupTimeLeft *int64 `json:"GroupTimeLeft,omitempty" name:"GroupTimeLeft"`

	// 昵称列表,通过","拼接， 1元钱裂变活动中与Notes中uin一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickNameList *string `json:"NickNameList,omitempty" name:"NickNameList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserActivityInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserActivityInfoResponseParams `json:"Response"`
}

func (r *DescribeUserActivityInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserActivityInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWxCloudBaseRunEnvsRequestParams struct {
	// wx应用Id
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// 是否查询全地域
	AllRegions *bool `json:"AllRegions,omitempty" name:"AllRegions"`
}

type DescribeWxCloudBaseRunEnvsRequest struct {
	*tchttp.BaseRequest
	
	// wx应用Id
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// 是否查询全地域
	AllRegions *bool `json:"AllRegions,omitempty" name:"AllRegions"`
}

func (r *DescribeWxCloudBaseRunEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWxCloudBaseRunEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WxAppId")
	delete(f, "AllRegions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWxCloudBaseRunEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWxCloudBaseRunEnvsResponseParams struct {
	// env列表
	EnvList []*EnvInfo `json:"EnvList,omitempty" name:"EnvList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWxCloudBaseRunEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWxCloudBaseRunEnvsResponseParams `json:"Response"`
}

func (r *DescribeWxCloudBaseRunEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWxCloudBaseRunEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWxCloudBaseRunSubNetsRequestParams struct {
	// VPC id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 查询个数限制，不填或小于等于0，等于不限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeWxCloudBaseRunSubNetsRequest struct {
	*tchttp.BaseRequest
	
	// VPC id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 查询个数限制，不填或小于等于0，等于不限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeWxCloudBaseRunSubNetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWxCloudBaseRunSubNetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWxCloudBaseRunSubNetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWxCloudBaseRunSubNetsResponseParams struct {
	// 子网Id列表
	SubNetIds []*string `json:"SubNetIds,omitempty" name:"SubNetIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWxCloudBaseRunSubNetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWxCloudBaseRunSubNetsResponseParams `json:"Response"`
}

func (r *DescribeWxCloudBaseRunSubNetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWxCloudBaseRunSubNetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 针对预付费 删除隔离中的环境时要传true 正常环境直接跳过隔离期删除
	IsForce *bool `json:"IsForce,omitempty" name:"IsForce"`

	// 是否绕过资源检查，资源包等额外资源，默认为false，如果为true，则不检查资源是否有数据，直接删除。
	BypassCheck *bool `json:"BypassCheck,omitempty" name:"BypassCheck"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "IsForce")
	delete(f, "BypassCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyEnvResponse struct {
	*tchttp.BaseResponse
	Response *DestroyEnvResponseParams `json:"Response"`
}

func (r *DestroyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyStandaloneGatewayRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网名名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 是否强制释放
	IsForce *bool `json:"IsForce,omitempty" name:"IsForce"`
}

type DestroyStandaloneGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网名名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 是否强制释放
	IsForce *bool `json:"IsForce,omitempty" name:"IsForce"`
}

func (r *DestroyStandaloneGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyStandaloneGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GatewayName")
	delete(f, "IsForce")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyStandaloneGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyStandaloneGatewayResponseParams struct {
	// 删除独立网关状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyStandaloneGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DestroyStandaloneGatewayResponseParams `json:"Response"`
}

func (r *DestroyStandaloneGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyStandaloneGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyStaticStoreRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// cdn域名
	CdnDomain *string `json:"CdnDomain,omitempty" name:"CdnDomain"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyStaticStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CdnDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyStaticStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyStaticStoreResponseParams struct {
	// 条件任务结果(succ/fail)
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *DestroyStaticStoreResponseParams `json:"Response"`
}

func (r *DestroyStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 是否开启 `超过套餐额度部分转按量付费`
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableOverrun *bool `json:"EnableOverrun,omitempty" name:"EnableOverrun"`

	// 环境套餐类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtPackageType *string `json:"ExtPackageType,omitempty" name:"ExtPackageType"`
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
	Databases []*DatabasesInfo `json:"Databases,omitempty" name:"Databases"`

	// 存储列表
	Storages []*StorageInfo `json:"Storages,omitempty" name:"Storages"`

	// 函数列表
	Functions []*FunctionInfo `json:"Functions,omitempty" name:"Functions"`

	// tcb产品套餐ID，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 套餐中文名称，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 云日志服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogServices []*LogServiceInfo `json:"LogServices,omitempty" name:"LogServices"`

	// 静态资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaticStorages []*StaticStorageInfo `json:"StaticStorages,omitempty" name:"StaticStorages"`

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

	// 环境标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 自定义日志服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomLogServices []*ClsInfo `json:"CustomLogServices,omitempty" name:"CustomLogServices"`

	// 环境类型：baas, run, hoting, weda
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 是否是dau新套餐
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDauPackage *bool `json:"IsDauPackage,omitempty" name:"IsDauPackage"`

	// 套餐类型:空\baas\tcbr
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`
}

// Predefined struct for user
type EstablishCloudBaseRunServerRequestParams struct {
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

	// OA PUBLIC MINIAPP VPC
	OpenAccessTypes []*string `json:"OpenAccessTypes,omitempty" name:"OpenAccessTypes"`

	// 是否创建Path 0未传默认创建 1创建 2不创建
	IsCreatePath *int64 `json:"IsCreatePath,omitempty" name:"IsCreatePath"`

	// 指定创建路径（如不存在，则创建。存在，则忽略）
	ServerPath *string `json:"ServerPath,omitempty" name:"ServerPath"`
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

	// OA PUBLIC MINIAPP VPC
	OpenAccessTypes []*string `json:"OpenAccessTypes,omitempty" name:"OpenAccessTypes"`

	// 是否创建Path 0未传默认创建 1创建 2不创建
	IsCreatePath *int64 `json:"IsCreatePath,omitempty" name:"IsCreatePath"`

	// 指定创建路径（如不存在，则创建。存在，则忽略）
	ServerPath *string `json:"ServerPath,omitempty" name:"ServerPath"`
}

func (r *EstablishCloudBaseRunServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EstablishCloudBaseRunServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServiceName")
	delete(f, "IsPublic")
	delete(f, "ImageRepo")
	delete(f, "Remark")
	delete(f, "EsInfo")
	delete(f, "LogType")
	delete(f, "OperatorRemark")
	delete(f, "Source")
	delete(f, "VpcInfo")
	delete(f, "PublicAccess")
	delete(f, "OpenAccessTypes")
	delete(f, "IsCreatePath")
	delete(f, "ServerPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EstablishCloudBaseRunServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EstablishCloudBaseRunServerResponseParams struct {
	// 创建服务是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EstablishCloudBaseRunServerResponse struct {
	*tchttp.BaseResponse
	Response *EstablishCloudBaseRunServerResponseParams `json:"Response"`
}

func (r *EstablishCloudBaseRunServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EstablishCloudBaseRunServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EstablishWxGatewayRouteRequestParams struct {
	// 网关id
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// 服务名称
	GatewayRouteName *string `json:"GatewayRouteName,omitempty" name:"GatewayRouteName"`

	// 服务地址
	GatewayRouteAddr *string `json:"GatewayRouteAddr,omitempty" name:"GatewayRouteAddr"`

	// 协议类型 http/https
	GatewayRouteProtocol *string `json:"GatewayRouteProtocol,omitempty" name:"GatewayRouteProtocol"`

	// 服务描述
	GatewayRouteDesc *string `json:"GatewayRouteDesc,omitempty" name:"GatewayRouteDesc"`
}

type EstablishWxGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// 网关id
	GatewayId *string `json:"GatewayId,omitempty" name:"GatewayId"`

	// 服务名称
	GatewayRouteName *string `json:"GatewayRouteName,omitempty" name:"GatewayRouteName"`

	// 服务地址
	GatewayRouteAddr *string `json:"GatewayRouteAddr,omitempty" name:"GatewayRouteAddr"`

	// 协议类型 http/https
	GatewayRouteProtocol *string `json:"GatewayRouteProtocol,omitempty" name:"GatewayRouteProtocol"`

	// 服务描述
	GatewayRouteDesc *string `json:"GatewayRouteDesc,omitempty" name:"GatewayRouteDesc"`
}

func (r *EstablishWxGatewayRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EstablishWxGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GatewayRouteName")
	delete(f, "GatewayRouteAddr")
	delete(f, "GatewayRouteProtocol")
	delete(f, "GatewayRouteDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EstablishWxGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EstablishWxGatewayRouteResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EstablishWxGatewayRouteResponse struct {
	*tchttp.BaseResponse
	Response *EstablishWxGatewayRouteResponseParams `json:"Response"`
}

func (r *EstablishWxGatewayRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EstablishWxGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtensionFile struct {
	// 文件类型。枚举值
	// <li>FUNCTION：函数代码</li>
	// <li>STATIC：静态托管代码</li>
	// <li>SMS：短信文件</li>
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件名，长度不超过24
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type ExtensionFileInfo struct {
	// 模板里使用的地址
	CodeUri *string `json:"CodeUri,omitempty" name:"CodeUri"`

	// 上传文件的临时地址，含签名
	UploadUrl *string `json:"UploadUrl,omitempty" name:"UploadUrl"`

	// 自定义密钥。如果为空，则表示不需要加密
	CustomKey *string `json:"CustomKey,omitempty" name:"CustomKey"`

	// 文件大小限制，单位M，客户端上传前需要主动检查文件大小，超过限制的文件会被删除。
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`
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

// Predefined struct for user
type FreezeCloudBaseRunServersRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名列表
	ServerNameList []*string `json:"ServerNameList,omitempty" name:"ServerNameList"`
}

type FreezeCloudBaseRunServersRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名列表
	ServerNameList []*string `json:"ServerNameList,omitempty" name:"ServerNameList"`
}

func (r *FreezeCloudBaseRunServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeCloudBaseRunServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FreezeCloudBaseRunServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeCloudBaseRunServersResponseParams struct {
	// 批量状态状态
	// 成功：succ
	// 失败：fail
	// 部分：partial（部分成功、部分失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 冻结失败服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailServerList []*string `json:"FailServerList,omitempty" name:"FailServerList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FreezeCloudBaseRunServersResponse struct {
	*tchttp.BaseResponse
	Response *FreezeCloudBaseRunServersResponseParams `json:"Response"`
}

func (r *FreezeCloudBaseRunServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeCloudBaseRunServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionInfo struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}

type HpaPolicy struct {
	// 策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 策略阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyThreshold *int64 `json:"PolicyThreshold,omitempty" name:"PolicyThreshold"`
}

type KVPair struct {
	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogObject struct {
	// 日志属于的 topic ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题的名字
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 日志时间
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 采集路径
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 日志来源设备
	Source *string `json:"Source,omitempty" name:"Source"`
}

type LogResObject struct {
	// 获取更多检索结果的游标
	Context *string `json:"Context,omitempty" name:"Context"`

	// 搜索结果是否已经全部返回
	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

	// 日志内容信息
	Results []*LogObject `json:"Results,omitempty" name:"Results"`
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

	// topic保存时长 默认7天
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *int64 `json:"Period,omitempty" name:"Period"`
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

// Predefined struct for user
type ModifyCloudBaseRunServerFlowConfRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 流量占比
	VersionFlowItems []*CloudBaseRunVersionFlowItem `json:"VersionFlowItems,omitempty" name:"VersionFlowItems"`

	// 流量类型（URL_PARAMS / FLOW / HEADERS)
	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

type ModifyCloudBaseRunServerFlowConfRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 流量占比
	VersionFlowItems []*CloudBaseRunVersionFlowItem `json:"VersionFlowItems,omitempty" name:"VersionFlowItems"`

	// 流量类型（URL_PARAMS / FLOW / HEADERS)
	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

func (r *ModifyCloudBaseRunServerFlowConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudBaseRunServerFlowConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "VersionFlowItems")
	delete(f, "TrafficType")
	delete(f, "OperatorRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudBaseRunServerFlowConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudBaseRunServerFlowConfResponseParams struct {
	// 返回结果，succ代表成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCloudBaseRunServerFlowConfResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudBaseRunServerFlowConfResponseParams `json:"Response"`
}

func (r *ModifyCloudBaseRunServerFlowConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudBaseRunServerFlowConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudBaseRunServerVersionRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 环境变量
	EnvParams *string `json:"EnvParams,omitempty" name:"EnvParams"`

	// 最小副本数
	MinNum *string `json:"MinNum,omitempty" name:"MinNum"`

	// 最大副本数
	MaxNum *string `json:"MaxNum,omitempty" name:"MaxNum"`

	// 端口
	ContainerPort *string `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 日志采集路径
	CustomLogs *string `json:"CustomLogs,omitempty" name:"CustomLogs"`

	// 是否重设备注
	IsResetRemark *bool `json:"IsResetRemark,omitempty" name:"IsResetRemark"`

	// 修改基础信息
	BasicModify *bool `json:"BasicModify,omitempty" name:"BasicModify"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

type ModifyCloudBaseRunServerVersionRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 环境变量
	EnvParams *string `json:"EnvParams,omitempty" name:"EnvParams"`

	// 最小副本数
	MinNum *string `json:"MinNum,omitempty" name:"MinNum"`

	// 最大副本数
	MaxNum *string `json:"MaxNum,omitempty" name:"MaxNum"`

	// 端口
	ContainerPort *string `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 日志采集路径
	CustomLogs *string `json:"CustomLogs,omitempty" name:"CustomLogs"`

	// 是否重设备注
	IsResetRemark *bool `json:"IsResetRemark,omitempty" name:"IsResetRemark"`

	// 修改基础信息
	BasicModify *bool `json:"BasicModify,omitempty" name:"BasicModify"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

func (r *ModifyCloudBaseRunServerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudBaseRunServerVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "VersionName")
	delete(f, "EnvParams")
	delete(f, "MinNum")
	delete(f, "MaxNum")
	delete(f, "ContainerPort")
	delete(f, "Remark")
	delete(f, "CustomLogs")
	delete(f, "IsResetRemark")
	delete(f, "BasicModify")
	delete(f, "OperatorRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudBaseRunServerVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudBaseRunServerVersionResponseParams struct {
	// 返回结果（succ为成功）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCloudBaseRunServerVersionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudBaseRunServerVersionResponseParams `json:"Response"`
}

func (r *ModifyCloudBaseRunServerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudBaseRunServerVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClsTopicRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 日志生命周期，单位天，可取值范围1~3600，取值为3640时代表永久保存
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type ModifyClsTopicRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 日志生命周期，单位天，可取值范围1~3600，取值为3640时代表永久保存
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *ModifyClsTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClsTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClsTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClsTopicResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClsTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClsTopicResponseParams `json:"Response"`
}

func (r *ModifyClsTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClsTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseACLRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	delete(f, "AclTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseACLResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseACLResponseParams `json:"Response"`
}

func (r *ModifyDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEndUserRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// C端用户端的唯一ID
	UUId *string `json:"UUId,omitempty" name:"UUId"`

	// 帐号的状态
	// <li>ENABLE</li>
	// <li>DISABLE</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyEndUserRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// C端用户端的唯一ID
	UUId *string `json:"UUId,omitempty" name:"UUId"`

	// 帐号的状态
	// <li>ENABLE</li>
	// <li>DISABLE</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyEndUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEndUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "UUId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEndUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEndUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEndUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEndUserResponseParams `json:"Response"`
}

func (r *ModifyEndUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEndUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 环境备注名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyEnvResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvResponseParams `json:"Response"`
}

func (r *ModifyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ObjectKV struct {
	// object 的 key
	Key *string `json:"Key,omitempty" name:"Key"`

	// object key 对应的 value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type OneClickTaskStepInfo struct {
	// 未启动："todo"
	// 运行中："running"
	// 失败："failed"
	// 成功结束："finished"
	Status *string `json:"Status,omitempty" name:"Status"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 耗时：秒
	CostTime *int64 `json:"CostTime,omitempty" name:"CostTime"`

	// 失败原因
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// 步骤名
	Name *string `json:"Name,omitempty" name:"Name"`
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

	// 安装标记。建议使用方统一转大小写之后再判断。
	// <li>QuickStart：快速启动来源</li>
	// <li>Activity：活动来源</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Flag *string `json:"Flag,omitempty" name:"Flag"`

	// 下单时的参数
	ReqBody *string `json:"ReqBody,omitempty" name:"ReqBody"`
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

type PostPaidEnvDeductInfo struct {
	// 资源方
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 按量计费详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResQuota *float64 `json:"ResQuota,omitempty" name:"ResQuota"`

	// 资源包抵扣详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgQuota *float64 `json:"PkgQuota,omitempty" name:"PkgQuota"`

	// 免费额度抵扣详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeQuota *float64 `json:"FreeQuota,omitempty" name:"FreeQuota"`

	// 环境id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
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

// Predefined struct for user
type ReinstateEnvRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReinstateEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReinstateEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReinstateEnvResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReinstateEnvResponse struct {
	*tchttp.BaseResponse
	Response *ReinstateEnvResponseParams `json:"Response"`
}

func (r *ReinstateEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReinstateEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceActivityRecordRequestParams struct {
	// 活动id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 状态码
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 自定义子状态
	SubStatus *string `json:"SubStatus,omitempty" name:"SubStatus"`

	// 鉴权token
	ChannelToken *string `json:"ChannelToken,omitempty" name:"ChannelToken"`

	// 渠道名，不同渠道对应不同secretKey
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type ReplaceActivityRecordRequest struct {
	*tchttp.BaseRequest
	
	// 活动id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 状态码
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 自定义子状态
	SubStatus *string `json:"SubStatus,omitempty" name:"SubStatus"`

	// 鉴权token
	ChannelToken *string `json:"ChannelToken,omitempty" name:"ChannelToken"`

	// 渠道名，不同渠道对应不同secretKey
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *ReplaceActivityRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceActivityRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityId")
	delete(f, "Status")
	delete(f, "SubStatus")
	delete(f, "ChannelToken")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceActivityRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceActivityRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceActivityRecordResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceActivityRecordResponseParams `json:"Response"`
}

func (r *ReplaceActivityRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceActivityRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollUpdateCloudBaseRunServerVersionRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 要替换的版本名称，可以为latest
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 枚举（package/repository/image)
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// repository的类型(coding/gitlab/github)
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 流量占比
	FlowRatio *int64 `json:"FlowRatio,omitempty" name:"FlowRatio"`

	// dockerfile地址
	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`

	// 构建目录
	BuildDir *string `json:"BuildDir,omitempty" name:"BuildDir"`

	// Cpu的大小，单位：核
	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`

	// Mem的大小，单位：G
	Mem *string `json:"Mem,omitempty" name:"Mem"`

	// 最小副本数，最小值：0
	MinNum *string `json:"MinNum,omitempty" name:"MinNum"`

	// 最大副本数
	MaxNum *string `json:"MaxNum,omitempty" name:"MaxNum"`

	// 策略类型
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 策略阈值
	PolicyThreshold *string `json:"PolicyThreshold,omitempty" name:"PolicyThreshold"`

	// 环境变量
	EnvParams *string `json:"EnvParams,omitempty" name:"EnvParams"`

	// 容器端口
	ContainerPort *int64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// repository地址
	Repository *string `json:"Repository,omitempty" name:"Repository"`

	// 分支
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 代码包名字
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 代码包版本
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// Image的详情
	ImageInfo *CloudBaseRunImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// Github等拉取代码的详情
	CodeDetail *CloudBaseCodeRepoDetail `json:"CodeDetail,omitempty" name:"CodeDetail"`

	// 是否回放流量
	IsRebuild *bool `json:"IsRebuild,omitempty" name:"IsRebuild"`

	// 延迟多长时间开始健康检查（单位s）
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// cfs挂载信息
	MountVolumeInfo []*CloudBaseRunVolumeMount `json:"MountVolumeInfo,omitempty" name:"MountVolumeInfo"`

	// 是否回滚
	Rollback *bool `json:"Rollback,omitempty" name:"Rollback"`

	// 版本历史名
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 自定义采集路径
	CustomLogs *string `json:"CustomLogs,omitempty" name:"CustomLogs"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitempty" name:"EnableUnion"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`

	// 服务路径（只会第一次生效）
	ServerPath *string `json:"ServerPath,omitempty" name:"ServerPath"`

	// 是否更新Cls
	IsUpdateCls *bool `json:"IsUpdateCls,omitempty" name:"IsUpdateCls"`

	// 自动扩缩容策略组
	PolicyDetail []*HpaPolicy `json:"PolicyDetail,omitempty" name:"PolicyDetail"`
}

type RollUpdateCloudBaseRunServerVersionRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 要替换的版本名称，可以为latest
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 枚举（package/repository/image)
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// repository的类型(coding/gitlab/github)
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 流量占比
	FlowRatio *int64 `json:"FlowRatio,omitempty" name:"FlowRatio"`

	// dockerfile地址
	DockerfilePath *string `json:"DockerfilePath,omitempty" name:"DockerfilePath"`

	// 构建目录
	BuildDir *string `json:"BuildDir,omitempty" name:"BuildDir"`

	// Cpu的大小，单位：核
	Cpu *string `json:"Cpu,omitempty" name:"Cpu"`

	// Mem的大小，单位：G
	Mem *string `json:"Mem,omitempty" name:"Mem"`

	// 最小副本数，最小值：0
	MinNum *string `json:"MinNum,omitempty" name:"MinNum"`

	// 最大副本数
	MaxNum *string `json:"MaxNum,omitempty" name:"MaxNum"`

	// 策略类型
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 策略阈值
	PolicyThreshold *string `json:"PolicyThreshold,omitempty" name:"PolicyThreshold"`

	// 环境变量
	EnvParams *string `json:"EnvParams,omitempty" name:"EnvParams"`

	// 容器端口
	ContainerPort *int64 `json:"ContainerPort,omitempty" name:"ContainerPort"`

	// 服务名称
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// repository地址
	Repository *string `json:"Repository,omitempty" name:"Repository"`

	// 分支
	Branch *string `json:"Branch,omitempty" name:"Branch"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 代码包名字
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 代码包版本
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// Image的详情
	ImageInfo *CloudBaseRunImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// Github等拉取代码的详情
	CodeDetail *CloudBaseCodeRepoDetail `json:"CodeDetail,omitempty" name:"CodeDetail"`

	// 是否回放流量
	IsRebuild *bool `json:"IsRebuild,omitempty" name:"IsRebuild"`

	// 延迟多长时间开始健康检查（单位s）
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// cfs挂载信息
	MountVolumeInfo []*CloudBaseRunVolumeMount `json:"MountVolumeInfo,omitempty" name:"MountVolumeInfo"`

	// 是否回滚
	Rollback *bool `json:"Rollback,omitempty" name:"Rollback"`

	// 版本历史名
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 自定义采集路径
	CustomLogs *string `json:"CustomLogs,omitempty" name:"CustomLogs"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitempty" name:"EnableUnion"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`

	// 服务路径（只会第一次生效）
	ServerPath *string `json:"ServerPath,omitempty" name:"ServerPath"`

	// 是否更新Cls
	IsUpdateCls *bool `json:"IsUpdateCls,omitempty" name:"IsUpdateCls"`

	// 自动扩缩容策略组
	PolicyDetail []*HpaPolicy `json:"PolicyDetail,omitempty" name:"PolicyDetail"`
}

func (r *RollUpdateCloudBaseRunServerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollUpdateCloudBaseRunServerVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "VersionName")
	delete(f, "UploadType")
	delete(f, "RepositoryType")
	delete(f, "FlowRatio")
	delete(f, "DockerfilePath")
	delete(f, "BuildDir")
	delete(f, "Cpu")
	delete(f, "Mem")
	delete(f, "MinNum")
	delete(f, "MaxNum")
	delete(f, "PolicyType")
	delete(f, "PolicyThreshold")
	delete(f, "EnvParams")
	delete(f, "ContainerPort")
	delete(f, "ServerName")
	delete(f, "Repository")
	delete(f, "Branch")
	delete(f, "VersionRemark")
	delete(f, "PackageName")
	delete(f, "PackageVersion")
	delete(f, "ImageInfo")
	delete(f, "CodeDetail")
	delete(f, "IsRebuild")
	delete(f, "InitialDelaySeconds")
	delete(f, "MountVolumeInfo")
	delete(f, "Rollback")
	delete(f, "SnapshotName")
	delete(f, "CustomLogs")
	delete(f, "EnableUnion")
	delete(f, "OperatorRemark")
	delete(f, "ServerPath")
	delete(f, "IsUpdateCls")
	delete(f, "PolicyDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollUpdateCloudBaseRunServerVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollUpdateCloudBaseRunServerVersionResponseParams struct {
	// succ为成功
	Result *string `json:"Result,omitempty" name:"Result"`

	// 滚动更新的VersionName
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 操作记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunId *string `json:"RunId,omitempty" name:"RunId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RollUpdateCloudBaseRunServerVersionResponse struct {
	*tchttp.BaseResponse
	Response *RollUpdateCloudBaseRunServerVersionResponseParams `json:"Response"`
}

func (r *RollUpdateCloudBaseRunServerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollUpdateCloudBaseRunServerVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClsLogRequestParams struct {
	// 环境唯一ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 查询起始时间条件
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间条件
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询语句，详情参考 https://cloud.tencent.com/document/product/614/47044
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitempty" name:"Context"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 是否使用Lucene语法，默认为false
	UseLucene *bool `json:"UseLucene,omitempty" name:"UseLucene"`
}

type SearchClsLogRequest struct {
	*tchttp.BaseRequest
	
	// 环境唯一ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 查询起始时间条件
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间条件
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询语句，详情参考 https://cloud.tencent.com/document/product/614/47044
	QueryString *string `json:"QueryString,omitempty" name:"QueryString"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitempty" name:"Context"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 是否使用Lucene语法，默认为false
	UseLucene *bool `json:"UseLucene,omitempty" name:"UseLucene"`
}

func (r *SearchClsLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryString")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseLucene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClsLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClsLogResponseParams struct {
	// 日志内容结果
	LogResults *LogResObject `json:"LogResults,omitempty" name:"LogResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchClsLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchClsLogResponseParams `json:"Response"`
}

func (r *SearchClsLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogResponse) FromJsonString(s string) error {
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

type SpecialCostItem struct {
	// 上报日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportDate *string `json:"ReportDate,omitempty" name:"ReportDate"`

	// 腾讯云uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 资源id:环境id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 上报任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type StandaloneGatewayInfo struct {
	// 独立网关名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// CPU核心数
	CPU *float64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小，单位MB
	Mem *uint64 `json:"Mem,omitempty" name:"Mem"`

	// 套餐包版本名称
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 网关别名
	GatewayAlias *string `json:"GatewayAlias,omitempty" name:"GatewayAlias"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 网关描述
	GatewayDesc *string `json:"GatewayDesc,omitempty" name:"GatewayDesc"`

	// 网关状态
	GateWayStatus *string `json:"GateWayStatus,omitempty" name:"GateWayStatus"`

	// 服务信息
	ServiceInfo *BackendServiceInfo `json:"ServiceInfo,omitempty" name:"ServiceInfo"`

	// 公网CLBIP
	PublicClbIp *string `json:"PublicClbIp,omitempty" name:"PublicClbIp"`

	// 内网CLBIP
	InternalClbIp *string `json:"InternalClbIp,omitempty" name:"InternalClbIp"`
}

type StandaloneGatewayPackageInfo struct {
	// CPU核心数
	CPU *float64 `json:"CPU,omitempty" name:"CPU"`

	// 内存大小，单位MB
	Mem *uint64 `json:"Mem,omitempty" name:"Mem"`

	// 套餐包版本名称
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`
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

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TkeClusterInfo struct {
	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群的vpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 版本内网CLB所在子网Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionClbSubnetId *string `json:"VersionClbSubnetId,omitempty" name:"VersionClbSubnetId"`
}

// Predefined struct for user
type TurnOffStandaloneGatewayRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网关名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 服务名称列表
	ServiceNameList []*string `json:"ServiceNameList,omitempty" name:"ServiceNameList"`
}

type TurnOffStandaloneGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网关名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 服务名称列表
	ServiceNameList []*string `json:"ServiceNameList,omitempty" name:"ServiceNameList"`
}

func (r *TurnOffStandaloneGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnOffStandaloneGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GatewayName")
	delete(f, "ServiceNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TurnOffStandaloneGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnOffStandaloneGatewayResponseParams struct {
	// 关闭独立网关状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TurnOffStandaloneGatewayResponse struct {
	*tchttp.BaseResponse
	Response *TurnOffStandaloneGatewayResponseParams `json:"Response"`
}

func (r *TurnOffStandaloneGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnOffStandaloneGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnOnStandaloneGatewayRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网关名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 服务名称列表
	ServiceNameList []*string `json:"ServiceNameList,omitempty" name:"ServiceNameList"`
}

type TurnOnStandaloneGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 网关名称
	GatewayName *string `json:"GatewayName,omitempty" name:"GatewayName"`

	// 服务名称列表
	ServiceNameList []*string `json:"ServiceNameList,omitempty" name:"ServiceNameList"`
}

func (r *TurnOnStandaloneGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnOnStandaloneGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GatewayName")
	delete(f, "ServiceNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TurnOnStandaloneGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnOnStandaloneGatewayResponseParams struct {
	// 小租户网关开启状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TurnOnStandaloneGatewayResponse struct {
	*tchttp.BaseResponse
	Response *TurnOnStandaloneGatewayResponseParams `json:"Response"`
}

func (r *TurnOnStandaloneGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnOnStandaloneGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnfreezeCloudBaseRunServersRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称列表
	ServerNameList []*string `json:"ServerNameList,omitempty" name:"ServerNameList"`
}

type UnfreezeCloudBaseRunServersRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名称列表
	ServerNameList []*string `json:"ServerNameList,omitempty" name:"ServerNameList"`
}

func (r *UnfreezeCloudBaseRunServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfreezeCloudBaseRunServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerNameList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnfreezeCloudBaseRunServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnfreezeCloudBaseRunServersResponseParams struct {
	// 批量执行结果
	// 成功：succ
	// 失败：fail
	// 部分：partial（部分成功、部分失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 解冻失败列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailServerList []*string `json:"FailServerList,omitempty" name:"FailServerList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnfreezeCloudBaseRunServersResponse struct {
	*tchttp.BaseResponse
	Response *UnfreezeCloudBaseRunServersResponseParams `json:"Response"`
}

func (r *UnfreezeCloudBaseRunServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfreezeCloudBaseRunServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}