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

package v20191022

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AlgorithmSpecification struct {
	// 镜像名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingImageName *string `json:"TrainingImageName,omitempty" name:"TrainingImageName"`

	// 输入模式File|Pipe
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingInputMode *string `json:"TrainingInputMode,omitempty" name:"TrainingInputMode"`

	// 算法名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgorithmName *string `json:"AlgorithmName,omitempty" name:"AlgorithmName"`
}

type BillingLabel struct {
	// 计费项标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 存储大小
	VolumeSize *int64 `json:"VolumeSize,omitempty" name:"VolumeSize"`

	// 计费状态
	// None: 不计费
	// StorageOnly: 仅存储计费
	// Computing: 计算和存储都计费
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ClsConfig struct {
	// 接入类型，可选项为free、customer
	Type *string `json:"Type,omitempty" name:"Type"`

	// 自定义CLS的日志集ID，只有当Type为customer时生效
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 自定义CLS的日志主题ID，只有当Type为customer时生效
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type CodeRepoSummary struct {
	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`

	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

	// Git配置
	GitConfig *GitConfig `json:"GitConfig,omitempty" name:"GitConfig"`

	// 是否有Git凭证
	NoSecret *bool `json:"NoSecret,omitempty" name:"NoSecret"`
}

type CosDataSource struct {
	// cos桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// cos文件key
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyPrefix *string `json:"KeyPrefix,omitempty" name:"KeyPrefix"`

	// 分布式数据下载方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDistributionType *string `json:"DataDistributionType,omitempty" name:"DataDistributionType"`

	// 数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *string `json:"DataType,omitempty" name:"DataType"`
}

// Predefined struct for user
type CreateCodeRepositoryRequestParams struct {
	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

	// Git相关配置
	GitConfig *GitConfig `json:"GitConfig,omitempty" name:"GitConfig"`

	// Git凭证
	GitSecret *GitSecret `json:"GitSecret,omitempty" name:"GitSecret"`
}

type CreateCodeRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

	// Git相关配置
	GitConfig *GitConfig `json:"GitConfig,omitempty" name:"GitConfig"`

	// Git凭证
	GitSecret *GitSecret `json:"GitSecret,omitempty" name:"GitSecret"`
}

func (r *CreateCodeRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CodeRepositoryName")
	delete(f, "GitConfig")
	delete(f, "GitSecret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCodeRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodeRepositoryResponseParams struct {
	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCodeRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateCodeRepositoryResponseParams `json:"Response"`
}

func (r *CreateCodeRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookInstanceRequestParams struct {
	// Notebook实例名称，不能超过63个字符
	// 规则：“^\[a-zA-Z0-9\](-\*\[a-zA-Z0-9\])\*$”
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// Notebook算力类型
	// 参考https://cloud.tencent.com/document/product/851/41239
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 数据卷大小(GB)
	// 用户持久化Notebook实例的数据
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitempty" name:"VolumeSizeInGB"`

	// 外网访问权限，可取值Enabled/Disabled
	// 开启后，Notebook实例可以具有访问外网80，443端口的权限
	DirectInternetAccess *string `json:"DirectInternetAccess,omitempty" name:"DirectInternetAccess"`

	// Root用户权限，可取值Enabled/Disabled
	// 开启后，Notebook实例可以切换至root用户执行命令
	RootAccess *string `json:"RootAccess,omitempty" name:"RootAccess"`

	// 子网ID
	// 如果需要Notebook实例访问VPC内的资源，则需要选择对应的子网
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 生命周期脚本名称
	// 必须是已存在的生命周期脚本，具体参考https://cloud.tencent.com/document/product/851/43140
	LifecycleScriptsName *string `json:"LifecycleScriptsName,omitempty" name:"LifecycleScriptsName"`

	// 默认存储库名称
	// 可以是已创建的存储库名称或者已https://开头的公共git库
	// 参考https://cloud.tencent.com/document/product/851/43139
	DefaultCodeRepository *string `json:"DefaultCodeRepository,omitempty" name:"DefaultCodeRepository"`

	// 其他存储库列表
	// 每个元素可以是已创建的存储库名称或者已https://开头的公共git库
	// 参考https://cloud.tencent.com/document/product/851/43139
	AdditionalCodeRepositories []*string `json:"AdditionalCodeRepositories,omitempty" name:"AdditionalCodeRepositories"`

	// 已弃用，请使用ClsConfig配置。
	// 是否开启CLS日志服务，可取值Enabled/Disabled，默认为Disabled
	// 开启后，Notebook运行的日志会收集到CLS中，CLS会产生费用，请根据需要选择
	ClsAccess *string `json:"ClsAccess,omitempty" name:"ClsAccess"`

	// 自动停止配置
	// 选择定时停止Notebook实例
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// 自动停止，可取值Enabled/Disabled
	// 取值为Disabled的时候StoppingCondition将被忽略
	// 取值为Enabled的时候读取StoppingCondition作为自动停止的配置
	AutoStopping *string `json:"AutoStopping,omitempty" name:"AutoStopping"`

	// 接入日志的配置，默认接入免费日志
	ClsConfig *ClsConfig `json:"ClsConfig,omitempty" name:"ClsConfig"`
}

type CreateNotebookInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Notebook实例名称，不能超过63个字符
	// 规则：“^\[a-zA-Z0-9\](-\*\[a-zA-Z0-9\])\*$”
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// Notebook算力类型
	// 参考https://cloud.tencent.com/document/product/851/41239
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 数据卷大小(GB)
	// 用户持久化Notebook实例的数据
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitempty" name:"VolumeSizeInGB"`

	// 外网访问权限，可取值Enabled/Disabled
	// 开启后，Notebook实例可以具有访问外网80，443端口的权限
	DirectInternetAccess *string `json:"DirectInternetAccess,omitempty" name:"DirectInternetAccess"`

	// Root用户权限，可取值Enabled/Disabled
	// 开启后，Notebook实例可以切换至root用户执行命令
	RootAccess *string `json:"RootAccess,omitempty" name:"RootAccess"`

	// 子网ID
	// 如果需要Notebook实例访问VPC内的资源，则需要选择对应的子网
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 生命周期脚本名称
	// 必须是已存在的生命周期脚本，具体参考https://cloud.tencent.com/document/product/851/43140
	LifecycleScriptsName *string `json:"LifecycleScriptsName,omitempty" name:"LifecycleScriptsName"`

	// 默认存储库名称
	// 可以是已创建的存储库名称或者已https://开头的公共git库
	// 参考https://cloud.tencent.com/document/product/851/43139
	DefaultCodeRepository *string `json:"DefaultCodeRepository,omitempty" name:"DefaultCodeRepository"`

	// 其他存储库列表
	// 每个元素可以是已创建的存储库名称或者已https://开头的公共git库
	// 参考https://cloud.tencent.com/document/product/851/43139
	AdditionalCodeRepositories []*string `json:"AdditionalCodeRepositories,omitempty" name:"AdditionalCodeRepositories"`

	// 已弃用，请使用ClsConfig配置。
	// 是否开启CLS日志服务，可取值Enabled/Disabled，默认为Disabled
	// 开启后，Notebook运行的日志会收集到CLS中，CLS会产生费用，请根据需要选择
	ClsAccess *string `json:"ClsAccess,omitempty" name:"ClsAccess"`

	// 自动停止配置
	// 选择定时停止Notebook实例
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// 自动停止，可取值Enabled/Disabled
	// 取值为Disabled的时候StoppingCondition将被忽略
	// 取值为Enabled的时候读取StoppingCondition作为自动停止的配置
	AutoStopping *string `json:"AutoStopping,omitempty" name:"AutoStopping"`

	// 接入日志的配置，默认接入免费日志
	ClsConfig *ClsConfig `json:"ClsConfig,omitempty" name:"ClsConfig"`
}

func (r *CreateNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookInstanceName")
	delete(f, "InstanceType")
	delete(f, "VolumeSizeInGB")
	delete(f, "DirectInternetAccess")
	delete(f, "RootAccess")
	delete(f, "SubnetId")
	delete(f, "LifecycleScriptsName")
	delete(f, "DefaultCodeRepository")
	delete(f, "AdditionalCodeRepositories")
	delete(f, "ClsAccess")
	delete(f, "StoppingCondition")
	delete(f, "AutoStopping")
	delete(f, "ClsConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookInstanceResponseParams struct {
	// Notebook实例名字
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookInstanceResponseParams `json:"Response"`
}

func (r *CreateNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookLifecycleScriptRequestParams struct {
	// Notebook生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 创建脚本，base64编码
	// base64后的脚本长度不能超过16384个字符
	CreateScript *string `json:"CreateScript,omitempty" name:"CreateScript"`

	// 启动脚本，base64编码
	// base64后的脚本长度不能超过16384个字符
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`
}

type CreateNotebookLifecycleScriptRequest struct {
	*tchttp.BaseRequest
	
	// Notebook生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 创建脚本，base64编码
	// base64后的脚本长度不能超过16384个字符
	CreateScript *string `json:"CreateScript,omitempty" name:"CreateScript"`

	// 启动脚本，base64编码
	// base64后的脚本长度不能超过16384个字符
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`
}

func (r *CreateNotebookLifecycleScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookLifecycleScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookLifecycleScriptsName")
	delete(f, "CreateScript")
	delete(f, "StartScript")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookLifecycleScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookLifecycleScriptResponseParams struct {
	// 生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNotebookLifecycleScriptResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookLifecycleScriptResponseParams `json:"Response"`
}

func (r *CreateNotebookLifecycleScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookLifecycleScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresignedNotebookInstanceUrlRequestParams struct {
	// Notebook实例名称
	// 规则：“^\[a-zA-Z0-9\](-\*\[a-zA-Z0-9\])\*$”
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// session有效时间，秒，取值范围[1800, 43200]
	SessionExpirationDurationInSeconds *int64 `json:"SessionExpirationDurationInSeconds,omitempty" name:"SessionExpirationDurationInSeconds"`
}

type CreatePresignedNotebookInstanceUrlRequest struct {
	*tchttp.BaseRequest
	
	// Notebook实例名称
	// 规则：“^\[a-zA-Z0-9\](-\*\[a-zA-Z0-9\])\*$”
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// session有效时间，秒，取值范围[1800, 43200]
	SessionExpirationDurationInSeconds *int64 `json:"SessionExpirationDurationInSeconds,omitempty" name:"SessionExpirationDurationInSeconds"`
}

func (r *CreatePresignedNotebookInstanceUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresignedNotebookInstanceUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookInstanceName")
	delete(f, "SessionExpirationDurationInSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePresignedNotebookInstanceUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresignedNotebookInstanceUrlResponseParams struct {
	// 授权url
	AuthorizedUrl *string `json:"AuthorizedUrl,omitempty" name:"AuthorizedUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePresignedNotebookInstanceUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreatePresignedNotebookInstanceUrlResponseParams `json:"Response"`
}

func (r *CreatePresignedNotebookInstanceUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresignedNotebookInstanceUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingJobRequestParams struct {
	// 算法镜像配置
	AlgorithmSpecification *AlgorithmSpecification `json:"AlgorithmSpecification,omitempty" name:"AlgorithmSpecification"`

	// 输出数据配置
	OutputDataConfig *OutputDataConfig `json:"OutputDataConfig,omitempty" name:"OutputDataConfig"`

	// 资源实例配置
	ResourceConfig *ResourceConfig `json:"ResourceConfig,omitempty" name:"ResourceConfig"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

	// 输入数据配置
	InputDataConfig []*InputDataConfig `json:"InputDataConfig,omitempty" name:"InputDataConfig"`

	// 中止条件
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// 私有网络配置
	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

	// 算法超级参数
	HyperParameters *string `json:"HyperParameters,omitempty" name:"HyperParameters"`

	// 环境变量配置
	EnvConfig []*EnvConfig `json:"EnvConfig,omitempty" name:"EnvConfig"`

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 在资源不足（ResourceInsufficient）时后台不定时尝试重新创建训练任务。可取值Enabled/Disabled
	// 默认值为Disabled即不重新尝试。设为Enabled时重新尝试有一定的时间期限，定义在 StoppingCondition 中 MaxWaitTimeInSecond中 ，默认值为1天，超过该期限创建失败。
	RetryWhenResourceInsufficient *string `json:"RetryWhenResourceInsufficient,omitempty" name:"RetryWhenResourceInsufficient"`
}

type CreateTrainingJobRequest struct {
	*tchttp.BaseRequest
	
	// 算法镜像配置
	AlgorithmSpecification *AlgorithmSpecification `json:"AlgorithmSpecification,omitempty" name:"AlgorithmSpecification"`

	// 输出数据配置
	OutputDataConfig *OutputDataConfig `json:"OutputDataConfig,omitempty" name:"OutputDataConfig"`

	// 资源实例配置
	ResourceConfig *ResourceConfig `json:"ResourceConfig,omitempty" name:"ResourceConfig"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

	// 输入数据配置
	InputDataConfig []*InputDataConfig `json:"InputDataConfig,omitempty" name:"InputDataConfig"`

	// 中止条件
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// 私有网络配置
	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

	// 算法超级参数
	HyperParameters *string `json:"HyperParameters,omitempty" name:"HyperParameters"`

	// 环境变量配置
	EnvConfig []*EnvConfig `json:"EnvConfig,omitempty" name:"EnvConfig"`

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 在资源不足（ResourceInsufficient）时后台不定时尝试重新创建训练任务。可取值Enabled/Disabled
	// 默认值为Disabled即不重新尝试。设为Enabled时重新尝试有一定的时间期限，定义在 StoppingCondition 中 MaxWaitTimeInSecond中 ，默认值为1天，超过该期限创建失败。
	RetryWhenResourceInsufficient *string `json:"RetryWhenResourceInsufficient,omitempty" name:"RetryWhenResourceInsufficient"`
}

func (r *CreateTrainingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlgorithmSpecification")
	delete(f, "OutputDataConfig")
	delete(f, "ResourceConfig")
	delete(f, "TrainingJobName")
	delete(f, "InputDataConfig")
	delete(f, "StoppingCondition")
	delete(f, "VpcConfig")
	delete(f, "HyperParameters")
	delete(f, "EnvConfig")
	delete(f, "RoleName")
	delete(f, "RetryWhenResourceInsufficient")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrainingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingJobResponseParams struct {
	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTrainingJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrainingJobResponseParams `json:"Response"`
}

func (r *CreateTrainingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataSource struct {
	// cos数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosDataSource *CosDataSource `json:"CosDataSource,omitempty" name:"CosDataSource"`

	// 文件系统输入源
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSystemDataSource *FileSystemDataSource `json:"FileSystemDataSource,omitempty" name:"FileSystemDataSource"`
}

// Predefined struct for user
type DeleteCodeRepositoryRequestParams struct {
	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`
}

type DeleteCodeRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`
}

func (r *DeleteCodeRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CodeRepositoryName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCodeRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCodeRepositoryResponseParams struct {
	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCodeRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCodeRepositoryResponseParams `json:"Response"`
}

func (r *DeleteCodeRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookInstanceRequestParams struct {
	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`
}

type DeleteNotebookInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`
}

func (r *DeleteNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNotebookInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNotebookInstanceResponseParams `json:"Response"`
}

func (r *DeleteNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookLifecycleScriptRequestParams struct {
	// 生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 是否忽略已关联的 notebook 实例强行删除生命周期脚本，默认 false
	Forcible *bool `json:"Forcible,omitempty" name:"Forcible"`
}

type DeleteNotebookLifecycleScriptRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 是否忽略已关联的 notebook 实例强行删除生命周期脚本，默认 false
	Forcible *bool `json:"Forcible,omitempty" name:"Forcible"`
}

func (r *DeleteNotebookLifecycleScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookLifecycleScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookLifecycleScriptsName")
	delete(f, "Forcible")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNotebookLifecycleScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookLifecycleScriptResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNotebookLifecycleScriptResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNotebookLifecycleScriptResponseParams `json:"Response"`
}

func (r *DeleteNotebookLifecycleScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookLifecycleScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeRepositoriesRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// instance-name - String - 是否必填：否 -（过滤条件）按照名称过滤。
	// search-by-name - String - 是否必填：否 -（过滤条件）按照名称检索，模糊匹配。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序规则。默认取Descending
	// Descending 按更新时间降序
	// Ascending 按更新时间升序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

type DescribeCodeRepositoriesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// instance-name - String - 是否必填：否 -（过滤条件）按照名称过滤。
	// search-by-name - String - 是否必填：否 -（过滤条件）按照名称检索，模糊匹配。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序规则。默认取Descending
	// Descending 按更新时间降序
	// Ascending 按更新时间升序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeCodeRepositoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeRepositoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodeRepositoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeRepositoriesResponseParams struct {
	// 存储库总数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 存储库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeRepoSet []*CodeRepoSummary `json:"CodeRepoSet,omitempty" name:"CodeRepoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCodeRepositoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCodeRepositoriesResponseParams `json:"Response"`
}

func (r *DescribeCodeRepositoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeRepositoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeRepositoryRequestParams struct {
	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`
}

type DescribeCodeRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`
}

func (r *DescribeCodeRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CodeRepositoryName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCodeRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCodeRepositoryResponseParams struct {
	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`

	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

	// Git存储配置
	GitConfig *GitConfig `json:"GitConfig,omitempty" name:"GitConfig"`

	// 是否有Git凭证
	NoSecret *bool `json:"NoSecret,omitempty" name:"NoSecret"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCodeRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCodeRepositoryResponseParams `json:"Response"`
}

func (r *DescribeCodeRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCodeRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookInstanceRequestParams struct {
	// Notebook实例名称
	// 规则：“^\[a-zA-Z0-9\](-\*\[a-zA-Z0-9\])\*$”
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`
}

type DescribeNotebookInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Notebook实例名称
	// 规则：“^\[a-zA-Z0-9\](-\*\[a-zA-Z0-9\])\*$”
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`
}

func (r *DescribeNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookInstanceResponseParams struct {
	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// Notebook算力资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 角色的资源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// 外网访问权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	DirectInternetAccess *string `json:"DirectInternetAccess,omitempty" name:"DirectInternetAccess"`

	// Root用户权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootAccess *string `json:"RootAccess,omitempty" name:"RootAccess"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 数据卷大小(GB)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitempty" name:"VolumeSizeInGB"`

	// 创建失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`

	// Notebook实例创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// Notebook实例最近修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`

	// Notebook实例日志链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogUrl *string `json:"LogUrl,omitempty" name:"LogUrl"`

	// Notebook实例状态
	// 
	// Pending: 创建中
	// Inservice: 运行中
	// Stopping: 停止中
	// Stopped: 已停止
	// Failed: 失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookInstanceStatus *string `json:"NotebookInstanceStatus,omitempty" name:"NotebookInstanceStatus"`

	// Notebook实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// notebook生命周期脚本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifecycleScriptsName *string `json:"LifecycleScriptsName,omitempty" name:"LifecycleScriptsName"`

	// 默认存储库名称
	// 可以是已创建的存储库名称或者已https://开头的公共git库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCodeRepository *string `json:"DefaultCodeRepository,omitempty" name:"DefaultCodeRepository"`

	// 其他存储库列表
	// 每个元素可以是已创建的存储库名称或者已https://开头的公共git库
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdditionalCodeRepositories []*string `json:"AdditionalCodeRepositories,omitempty" name:"AdditionalCodeRepositories"`

	// 是否开启CLS日志服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsAccess *string `json:"ClsAccess,omitempty" name:"ClsAccess"`

	// 是否预付费实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prepay *bool `json:"Prepay,omitempty" name:"Prepay"`

	// 实例运行截止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// 自动停止配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// Cls配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsConfig *ClsConfig `json:"ClsConfig,omitempty" name:"ClsConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookInstanceResponseParams `json:"Response"`
}

func (r *DescribeNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookInstancesRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序规则。默认取Descending
	// Descending 按更新时间降序
	// Ascending 按更新时间升序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// 过滤条件。
	// instance-name - String - 是否必填：否 -（过滤条件）按照名称过滤。
	// search-by-name - String - 是否必填：否 -（过滤条件）按照名称检索，模糊匹配。
	// lifecycle-name - String - 是否必填：否 -（过滤条件）按照生命周期脚本名称过滤。
	// default-code-repo-name - String - 是否必填：否 -（过滤条件）按照默认存储库名称过滤。
	// additional-code-repo-name - String - 是否必填：否 -（过滤条件）按照其他存储库名称过滤。
	// billing-status - String - 是否必填：否 - （过滤条件）按照计费状态过滤，可取以下值
	//    StorageOnly：仅存储计费的实例
	//    Computing：计算和存储都计费的实例
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 【废弃字段】排序字段
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
}

type DescribeNotebookInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序规则。默认取Descending
	// Descending 按更新时间降序
	// Ascending 按更新时间升序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// 过滤条件。
	// instance-name - String - 是否必填：否 -（过滤条件）按照名称过滤。
	// search-by-name - String - 是否必填：否 -（过滤条件）按照名称检索，模糊匹配。
	// lifecycle-name - String - 是否必填：否 -（过滤条件）按照生命周期脚本名称过滤。
	// default-code-repo-name - String - 是否必填：否 -（过滤条件）按照默认存储库名称过滤。
	// additional-code-repo-name - String - 是否必填：否 -（过滤条件）按照其他存储库名称过滤。
	// billing-status - String - 是否必填：否 - （过滤条件）按照计费状态过滤，可取以下值
	//    StorageOnly：仅存储计费的实例
	//    Computing：计算和存储都计费的实例
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 【废弃字段】排序字段
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *DescribeNotebookInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortOrder")
	delete(f, "Filters")
	delete(f, "SortBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookInstancesResponseParams struct {
	// Notebook实例列表
	NotebookInstanceSet []*NotebookInstanceSummary `json:"NotebookInstanceSet,omitempty" name:"NotebookInstanceSet"`

	// Notebook实例总数目
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNotebookInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookInstancesResponseParams `json:"Response"`
}

func (r *DescribeNotebookInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookLifecycleScriptRequestParams struct {
	// 生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`
}

type DescribeNotebookLifecycleScriptRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`
}

func (r *DescribeNotebookLifecycleScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookLifecycleScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookLifecycleScriptsName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookLifecycleScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookLifecycleScriptResponseParams struct {
	// 生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 创建脚本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateScript *string `json:"CreateScript,omitempty" name:"CreateScript"`

	// 启动脚本
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最后修改时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNotebookLifecycleScriptResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookLifecycleScriptResponseParams `json:"Response"`
}

func (r *DescribeNotebookLifecycleScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookLifecycleScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookLifecycleScriptsRequestParams struct {
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// instance-name - String - 是否必填：否 -（过滤条件）按照名称过滤。
	// search-by-name - String - 是否必填：否 -（过滤条件）按照名称检索，模糊匹配。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序规则。默认取Descending
	// Descending 按更新时间降序
	// Ascending 按更新时间升序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

type DescribeNotebookLifecycleScriptsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// instance-name - String - 是否必填：否 -（过滤条件）按照名称过滤。
	// search-by-name - String - 是否必填：否 -（过滤条件）按照名称检索，模糊匹配。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序规则。默认取Descending
	// Descending 按更新时间降序
	// Ascending 按更新时间升序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeNotebookLifecycleScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookLifecycleScriptsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookLifecycleScriptsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookLifecycleScriptsResponseParams struct {
	// Notebook生命周期脚本列表
	NotebookLifecycleScriptsSet []*NotebookLifecycleScriptsSummary `json:"NotebookLifecycleScriptsSet,omitempty" name:"NotebookLifecycleScriptsSet"`

	// Notebook生命周期脚本总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNotebookLifecycleScriptsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookLifecycleScriptsResponseParams `json:"Response"`
}

func (r *DescribeNotebookLifecycleScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookLifecycleScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSummaryRequestParams struct {

}

type DescribeNotebookSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNotebookSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSummaryResponseParams struct {
	// 实例总数
	AllInstanceCnt *int64 `json:"AllInstanceCnt,omitempty" name:"AllInstanceCnt"`

	// 计费实例总数
	BillingInstanceCnt *int64 `json:"BillingInstanceCnt,omitempty" name:"BillingInstanceCnt"`

	// 仅存储计费的实例总数
	StorageOnlyBillingInstanceCnt *int64 `json:"StorageOnlyBillingInstanceCnt,omitempty" name:"StorageOnlyBillingInstanceCnt"`

	// 计算和存储都计费的实例总数
	ComputingBillingInstanceCnt *int64 `json:"ComputingBillingInstanceCnt,omitempty" name:"ComputingBillingInstanceCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNotebookSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSummaryResponseParams `json:"Response"`
}

func (r *DescribeNotebookSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingJobRequestParams struct {
	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`
}

type DescribeTrainingJobRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`
}

func (r *DescribeTrainingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingJobName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingJobResponseParams struct {
	// 算法镜像配置
	AlgorithmSpecification *AlgorithmSpecification `json:"AlgorithmSpecification,omitempty" name:"AlgorithmSpecification"`

	// 任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

	// 算法超级参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HyperParameters *string `json:"HyperParameters,omitempty" name:"HyperParameters"`

	// 输入数据配置
	InputDataConfig []*InputDataConfig `json:"InputDataConfig,omitempty" name:"InputDataConfig"`

	// 输出数据配置
	OutputDataConfig *OutputDataConfig `json:"OutputDataConfig,omitempty" name:"OutputDataConfig"`

	// 中止条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// 计算实例配置
	ResourceConfig *ResourceConfig `json:"ResourceConfig,omitempty" name:"ResourceConfig"`

	// 私有网络配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`

	// 最近修改时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`

	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingStartTime *string `json:"TrainingStartTime,omitempty" name:"TrainingStartTime"`

	// 任务完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingEndTime *string `json:"TrainingEndTime,omitempty" name:"TrainingEndTime"`

	// 模型输出配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelArtifacts *ModelArtifacts `json:"ModelArtifacts,omitempty" name:"ModelArtifacts"`

	// 详细状态，取值范围
	// Starting：启动中
	// Downloading: 准备训练数据
	// Training: 正在训练
	// Uploading: 上传训练结果
	// Completed：已完成
	// Failed: 失败
	// MaxRuntimeExceeded: 任务超过最大运行时间
	// Stopping: 停止中
	// Stopped：已停止
	SecondaryStatus *string `json:"SecondaryStatus,omitempty" name:"SecondaryStatus"`

	// 详细状态事件记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondaryStatusTransitions []*SecondaryStatusTransition `json:"SecondaryStatusTransitions,omitempty" name:"SecondaryStatusTransitions"`

	// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 训练任务状态，取值范围
	// InProgress：运行中
	// Completed: 已完成
	// Failed: 失败
	// Stopping: 停止中
	// Stopped：已停止
	TrainingJobStatus *string `json:"TrainingJobStatus,omitempty" name:"TrainingJobStatus"`

	// 训练任务日志链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogUrl *string `json:"LogUrl,omitempty" name:"LogUrl"`

	// 训练任务实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingJobResponseParams `json:"Response"`
}

func (r *DescribeTrainingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingJobsRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 创建时间晚于
	CreationTimeAfter *string `json:"CreationTimeAfter,omitempty" name:"CreationTimeAfter"`

	// 创建时间早于
	CreationTimeBefore *string `json:"CreationTimeBefore,omitempty" name:"CreationTimeBefore"`

	// 根据名称过滤
	NameContains *string `json:"NameContains,omitempty" name:"NameContains"`

	// 根据状态过滤
	StatusEquals *string `json:"StatusEquals,omitempty" name:"StatusEquals"`

	// 过滤条件。
	// instance-name - String - 是否必填：否 -（过滤条件）按照名称过滤。
	// search-by-name - String - 是否必填：否 -（过滤条件）按照名称检索，模糊匹配。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTrainingJobsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 创建时间晚于
	CreationTimeAfter *string `json:"CreationTimeAfter,omitempty" name:"CreationTimeAfter"`

	// 创建时间早于
	CreationTimeBefore *string `json:"CreationTimeBefore,omitempty" name:"CreationTimeBefore"`

	// 根据名称过滤
	NameContains *string `json:"NameContains,omitempty" name:"NameContains"`

	// 根据状态过滤
	StatusEquals *string `json:"StatusEquals,omitempty" name:"StatusEquals"`

	// 过滤条件。
	// instance-name - String - 是否必填：否 -（过滤条件）按照名称过滤。
	// search-by-name - String - 是否必填：否 -（过滤条件）按照名称检索，模糊匹配。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTrainingJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreationTimeAfter")
	delete(f, "CreationTimeBefore")
	delete(f, "NameContains")
	delete(f, "StatusEquals")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingJobsResponseParams struct {
	// 训练任务列表
	TrainingJobSet []*TrainingJobSummary `json:"TrainingJobSet,omitempty" name:"TrainingJobSet"`

	// 训练任务总数目
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingJobsResponseParams `json:"Response"`
}

func (r *DescribeTrainingJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvConfig struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type FileSystemDataSource struct {
	// 文件系统目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	DirectoryPath *string `json:"DirectoryPath,omitempty" name:"DirectoryPath"`

	// 文件系统类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSystemType *string `json:"FileSystemType,omitempty" name:"FileSystemType"`

	// 文件系统访问模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSystemAccessMode *string `json:"FileSystemAccessMode,omitempty" name:"FileSystemAccessMode"`

	// 文件系统ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type Filter struct {
	// 过滤字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段取值
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type GitConfig struct {
	// git地址
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" name:"RepositoryUrl"`

	// 代码分支
	// 注意：此字段可能返回 null，表示取不到有效值。
	Branch *string `json:"Branch,omitempty" name:"Branch"`
}

type GitSecret struct {
	// 无秘钥，默认选项
	NoSecret *bool `json:"NoSecret,omitempty" name:"NoSecret"`

	// Git用户名密码base64编码后的字符串
	// 编码前的内容应为Json字符串，如
	// {"UserName": "用户名", "Password":"密码"}
	Secret *string `json:"Secret,omitempty" name:"Secret"`
}

type InputDataConfig struct {
	// 通道名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 数据源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *DataSource `json:"DataSource,omitempty" name:"DataSource"`

	// 输入类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputMode *string `json:"InputMode,omitempty" name:"InputMode"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
}

type ModelArtifacts struct {
	// cos输出路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosModelArtifacts *string `json:"CosModelArtifacts,omitempty" name:"CosModelArtifacts"`
}

type NotebookInstanceSummary struct {
	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最近修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`

	// notebook实例名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// notebook实例状态，取值范围：
	// Pending: 创建中
	// Inservice: 运行中
	// Stopping: 停止中
	// Stopped: 已停止
	// Failed: 失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookInstanceStatus *string `json:"NotebookInstanceStatus,omitempty" name:"NotebookInstanceStatus"`

	// 算力类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupTime *string `json:"StartupTime,omitempty" name:"StartupTime"`

	// 运行截止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deadline *string `json:"Deadline,omitempty" name:"Deadline"`

	// 自动停止配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// 是否是预付费实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prepay *bool `json:"Prepay,omitempty" name:"Prepay"`

	// 计费标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingLabel *BillingLabel `json:"BillingLabel,omitempty" name:"BillingLabel"`

	// 运行时长，秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *int64 `json:"RuntimeInSeconds,omitempty" name:"RuntimeInSeconds"`

	// 剩余时长，秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainTimeInSeconds *int64 `json:"RemainTimeInSeconds,omitempty" name:"RemainTimeInSeconds"`
}

type NotebookLifecycleScriptsSummary struct {
	// notebook生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 修改时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`
}

type OutputDataConfig struct {
	// cos输出桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosOutputBucket *string `json:"CosOutputBucket,omitempty" name:"CosOutputBucket"`

	// cos输出key前缀
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosOutputKeyPrefix *string `json:"CosOutputKeyPrefix,omitempty" name:"CosOutputKeyPrefix"`

	// 文件系统输出，如果指定了文件系统，那么Cos输出会被忽略
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSystemDataSource *FileSystemDataSource `json:"FileSystemDataSource,omitempty" name:"FileSystemDataSource"`
}

type ResourceConfig struct {
	// 计算实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 计算实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 挂载CBS大小（GB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitempty" name:"VolumeSizeInGB"`
}

type SecondaryStatusTransition struct {
	// 状态开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 状态结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 状态名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusMessage *string `json:"StatusMessage,omitempty" name:"StatusMessage"`
}

// Predefined struct for user
type StartNotebookInstanceRequestParams struct {
	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// 自动停止，可取值Enabled/Disabled
	// 取值为Disabled的时候StoppingCondition将被忽略
	// 取值为Enabled的时候读取StoppingCondition作为自动停止的配置
	AutoStopping *string `json:"AutoStopping,omitempty" name:"AutoStopping"`

	// 自动停止配置，只在AutoStopping为Enabled的时候生效
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`
}

type StartNotebookInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// 自动停止，可取值Enabled/Disabled
	// 取值为Disabled的时候StoppingCondition将被忽略
	// 取值为Enabled的时候读取StoppingCondition作为自动停止的配置
	AutoStopping *string `json:"AutoStopping,omitempty" name:"AutoStopping"`

	// 自动停止配置，只在AutoStopping为Enabled的时候生效
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`
}

func (r *StartNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartNotebookInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookInstanceName")
	delete(f, "AutoStopping")
	delete(f, "StoppingCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartNotebookInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartNotebookInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StartNotebookInstanceResponseParams `json:"Response"`
}

func (r *StartNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartNotebookInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopNotebookInstanceRequestParams struct {
	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`
}

type StopNotebookInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`
}

func (r *StopNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopNotebookInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopNotebookInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopNotebookInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StopNotebookInstanceResponseParams `json:"Response"`
}

func (r *StopNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopNotebookInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTrainingJobRequestParams struct {
	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`
}

type StopTrainingJobRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`
}

func (r *StopTrainingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTrainingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingJobName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopTrainingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTrainingJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopTrainingJobResponse struct {
	*tchttp.BaseResponse
	Response *StopTrainingJobResponseParams `json:"Response"`
}

func (r *StopTrainingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTrainingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StoppingCondition struct {
	// 最长运行运行时间（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRuntimeInSeconds *uint64 `json:"MaxRuntimeInSeconds,omitempty" name:"MaxRuntimeInSeconds"`

	// 最长等待运行时间（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxWaitTimeInSeconds *uint64 `json:"MaxWaitTimeInSeconds,omitempty" name:"MaxWaitTimeInSeconds"`
}

type TrainingJobSummary struct {
	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 最近修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" name:"LastModifiedTime"`

	// 训练任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

	// 训练任务状态，取值范围
	// InProgress：运行中
	// Completed: 已完成
	// Failed: 失败
	// Stopping: 停止中
	// Stopped：已停止
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingJobStatus *string `json:"TrainingJobStatus,omitempty" name:"TrainingJobStatus"`

	// 完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingEndTime *string `json:"TrainingEndTime,omitempty" name:"TrainingEndTime"`

	// 算了实例Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceConfig *ResourceConfig `json:"ResourceConfig,omitempty" name:"ResourceConfig"`
}

// Predefined struct for user
type UpdateCodeRepositoryRequestParams struct {
	// 查询存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

	// Git凭证
	GitSecret *GitSecret `json:"GitSecret,omitempty" name:"GitSecret"`
}

type UpdateCodeRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 查询存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

	// Git凭证
	GitSecret *GitSecret `json:"GitSecret,omitempty" name:"GitSecret"`
}

func (r *UpdateCodeRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CodeRepositoryName")
	delete(f, "GitSecret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCodeRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCodeRepositoryResponseParams struct {
	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateCodeRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCodeRepositoryResponseParams `json:"Response"`
}

func (r *UpdateCodeRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNotebookInstanceRequestParams struct {
	// Notebook实例名称
	// 规则：“^\[a-zA-Z0-9\](-\*\[a-zA-Z0-9\])\*$”
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// 角色的资源描述
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// Root访问权限
	RootAccess *string `json:"RootAccess,omitempty" name:"RootAccess"`

	// 数据卷大小(GB)
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitempty" name:"VolumeSizeInGB"`

	// 算力资源类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// notebook生命周期脚本名称
	LifecycleScriptsName *string `json:"LifecycleScriptsName,omitempty" name:"LifecycleScriptsName"`

	// 是否解绑生命周期脚本，默认 false。
	// 该值为true时，LifecycleScriptsName将被忽略
	DisassociateLifecycleScript *bool `json:"DisassociateLifecycleScript,omitempty" name:"DisassociateLifecycleScript"`

	// 默认存储库名称
	// 可以是已创建的存储库名称或者已https://开头的公共git库
	DefaultCodeRepository *string `json:"DefaultCodeRepository,omitempty" name:"DefaultCodeRepository"`

	// 其他存储库列表
	// 每个元素可以是已创建的存储库名称或者已https://开头的公共git库
	AdditionalCodeRepositories []*string `json:"AdditionalCodeRepositories,omitempty" name:"AdditionalCodeRepositories"`

	// 是否取消关联默认存储库，默认false
	// 该值为true时，DefaultCodeRepository将被忽略
	DisassociateDefaultCodeRepository *bool `json:"DisassociateDefaultCodeRepository,omitempty" name:"DisassociateDefaultCodeRepository"`

	// 是否取消关联其他存储库，默认false
	// 该值为true时，AdditionalCodeRepositories将被忽略
	DisassociateAdditionalCodeRepositories *bool `json:"DisassociateAdditionalCodeRepositories,omitempty" name:"DisassociateAdditionalCodeRepositories"`

	// 已弃用，请使用ClsConfig配置。是否开启CLS日志服务，可取值Enabled/Disabled
	ClsAccess *string `json:"ClsAccess,omitempty" name:"ClsAccess"`

	// 自动停止，可取值Enabled/Disabled
	// 取值为Disabled的时候StoppingCondition将被忽略
	// 取值为Enabled的时候读取StoppingCondition作为自动停止的配置
	AutoStopping *string `json:"AutoStopping,omitempty" name:"AutoStopping"`

	// 自动停止配置，只在AutoStopping为Enabled的时候生效
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// 接入日志的配置，默认使用免费日志服务。
	ClsConfig *ClsConfig `json:"ClsConfig,omitempty" name:"ClsConfig"`
}

type UpdateNotebookInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Notebook实例名称
	// 规则：“^\[a-zA-Z0-9\](-\*\[a-zA-Z0-9\])\*$”
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// 角色的资源描述
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// Root访问权限
	RootAccess *string `json:"RootAccess,omitempty" name:"RootAccess"`

	// 数据卷大小(GB)
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitempty" name:"VolumeSizeInGB"`

	// 算力资源类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// notebook生命周期脚本名称
	LifecycleScriptsName *string `json:"LifecycleScriptsName,omitempty" name:"LifecycleScriptsName"`

	// 是否解绑生命周期脚本，默认 false。
	// 该值为true时，LifecycleScriptsName将被忽略
	DisassociateLifecycleScript *bool `json:"DisassociateLifecycleScript,omitempty" name:"DisassociateLifecycleScript"`

	// 默认存储库名称
	// 可以是已创建的存储库名称或者已https://开头的公共git库
	DefaultCodeRepository *string `json:"DefaultCodeRepository,omitempty" name:"DefaultCodeRepository"`

	// 其他存储库列表
	// 每个元素可以是已创建的存储库名称或者已https://开头的公共git库
	AdditionalCodeRepositories []*string `json:"AdditionalCodeRepositories,omitempty" name:"AdditionalCodeRepositories"`

	// 是否取消关联默认存储库，默认false
	// 该值为true时，DefaultCodeRepository将被忽略
	DisassociateDefaultCodeRepository *bool `json:"DisassociateDefaultCodeRepository,omitempty" name:"DisassociateDefaultCodeRepository"`

	// 是否取消关联其他存储库，默认false
	// 该值为true时，AdditionalCodeRepositories将被忽略
	DisassociateAdditionalCodeRepositories *bool `json:"DisassociateAdditionalCodeRepositories,omitempty" name:"DisassociateAdditionalCodeRepositories"`

	// 已弃用，请使用ClsConfig配置。是否开启CLS日志服务，可取值Enabled/Disabled
	ClsAccess *string `json:"ClsAccess,omitempty" name:"ClsAccess"`

	// 自动停止，可取值Enabled/Disabled
	// 取值为Disabled的时候StoppingCondition将被忽略
	// 取值为Enabled的时候读取StoppingCondition作为自动停止的配置
	AutoStopping *string `json:"AutoStopping,omitempty" name:"AutoStopping"`

	// 自动停止配置，只在AutoStopping为Enabled的时候生效
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// 接入日志的配置，默认使用免费日志服务。
	ClsConfig *ClsConfig `json:"ClsConfig,omitempty" name:"ClsConfig"`
}

func (r *UpdateNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNotebookInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookInstanceName")
	delete(f, "RoleArn")
	delete(f, "RootAccess")
	delete(f, "VolumeSizeInGB")
	delete(f, "InstanceType")
	delete(f, "LifecycleScriptsName")
	delete(f, "DisassociateLifecycleScript")
	delete(f, "DefaultCodeRepository")
	delete(f, "AdditionalCodeRepositories")
	delete(f, "DisassociateDefaultCodeRepository")
	delete(f, "DisassociateAdditionalCodeRepositories")
	delete(f, "ClsAccess")
	delete(f, "AutoStopping")
	delete(f, "StoppingCondition")
	delete(f, "ClsConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateNotebookInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNotebookInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateNotebookInstanceResponseParams `json:"Response"`
}

func (r *UpdateNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNotebookInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNotebookLifecycleScriptRequestParams struct {
	// notebook生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 创建脚本，base64编码
	// base64后的脚本长度不能超过16384个字符
	CreateScript *string `json:"CreateScript,omitempty" name:"CreateScript"`

	// 启动脚本，base64编码
	// base64后的脚本长度不能超过16384个字符
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`
}

type UpdateNotebookLifecycleScriptRequest struct {
	*tchttp.BaseRequest
	
	// notebook生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 创建脚本，base64编码
	// base64后的脚本长度不能超过16384个字符
	CreateScript *string `json:"CreateScript,omitempty" name:"CreateScript"`

	// 启动脚本，base64编码
	// base64后的脚本长度不能超过16384个字符
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`
}

func (r *UpdateNotebookLifecycleScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNotebookLifecycleScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookLifecycleScriptsName")
	delete(f, "CreateScript")
	delete(f, "StartScript")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateNotebookLifecycleScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNotebookLifecycleScriptResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateNotebookLifecycleScriptResponse struct {
	*tchttp.BaseResponse
	Response *UpdateNotebookLifecycleScriptResponseParams `json:"Response"`
}

func (r *UpdateNotebookLifecycleScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNotebookLifecycleScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcConfig struct {
	// 安全组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}