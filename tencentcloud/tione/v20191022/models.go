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

func (r *CreateCodeRepositoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCodeRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存储库名称
		CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCodeRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCodeRepositoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNotebookInstanceRequest struct {
	*tchttp.BaseRequest

	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// Notebook算力类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 数据卷大小(GB)
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitempty" name:"VolumeSizeInGB"`

	// 外网访问权限，可取值Enabled/Disabled
	DirectInternetAccess *string `json:"DirectInternetAccess,omitempty" name:"DirectInternetAccess"`

	// Root用户权限，可取值Enabled/Disabled
	RootAccess *string `json:"RootAccess,omitempty" name:"RootAccess"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 生命周期脚本名称
	LifecycleScriptsName *string `json:"LifecycleScriptsName,omitempty" name:"LifecycleScriptsName"`

	// 默认存储库名称
	// 可以是已创建的存储库名称或者已https://开头的公共git库
	DefaultCodeRepository *string `json:"DefaultCodeRepository,omitempty" name:"DefaultCodeRepository"`

	// 其他存储库列表
	// 每个元素可以是已创建的存储库名称或者已https://开头的公共git库
	AdditionalCodeRepositories []*string `json:"AdditionalCodeRepositories,omitempty" name:"AdditionalCodeRepositories" list`
}

func (r *CreateNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNotebookInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Notebook实例名字
		NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNotebookInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNotebookLifecycleScriptRequest struct {
	*tchttp.BaseRequest

	// Notebook生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 创建脚本，base64编码格式
	CreateScript *string `json:"CreateScript,omitempty" name:"CreateScript"`

	// 启动脚本，base64编码格式
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`
}

func (r *CreateNotebookLifecycleScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNotebookLifecycleScriptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNotebookLifecycleScriptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生命周期脚本名称
		NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNotebookLifecycleScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNotebookLifecycleScriptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePresignedNotebookInstanceUrlRequest struct {
	*tchttp.BaseRequest

	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`

	// session有效时间，秒
	SessionExpirationDurationInSeconds *int64 `json:"SessionExpirationDurationInSeconds,omitempty" name:"SessionExpirationDurationInSeconds"`
}

func (r *CreatePresignedNotebookInstanceUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePresignedNotebookInstanceUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePresignedNotebookInstanceUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 授权url
		AuthorizedUrl *string `json:"AuthorizedUrl,omitempty" name:"AuthorizedUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePresignedNotebookInstanceUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePresignedNotebookInstanceUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTrainingJobRequest struct {
	*tchttp.BaseRequest

	// 算法镜像配置
	AlgorithmSpecification *AlgorithmSpecification `json:"AlgorithmSpecification,omitempty" name:"AlgorithmSpecification"`

	// 输入数据配置
	InputDataConfig []*InputDataConfig `json:"InputDataConfig,omitempty" name:"InputDataConfig" list`

	// 输出数据配置
	OutputDataConfig *OutputDataConfig `json:"OutputDataConfig,omitempty" name:"OutputDataConfig"`

	// 资源实例配置
	ResourceConfig *ResourceConfig `json:"ResourceConfig,omitempty" name:"ResourceConfig"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

	// 中止条件
	StoppingCondition *StoppingCondition `json:"StoppingCondition,omitempty" name:"StoppingCondition"`

	// 私有网络配置
	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

	// 算法超级参数
	HyperParameters *string `json:"HyperParameters,omitempty" name:"HyperParameters"`

	// 环境变量配置
	EnvConfig []*EnvConfig `json:"EnvConfig,omitempty" name:"EnvConfig" list`

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *CreateTrainingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTrainingJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTrainingJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 训练任务名称
		TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTrainingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type DeleteCodeRepositoryRequest struct {
	*tchttp.BaseRequest

	// 存储库名称
	CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`
}

func (r *DeleteCodeRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCodeRepositoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCodeRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存储库名称
		CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCodeRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCodeRepositoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteNotebookInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNotebookInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteNotebookLifecycleScriptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNotebookLifecycleScriptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNotebookLifecycleScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNotebookLifecycleScriptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 排序规则。默认取Descending
	// Descending 按更新时间降序
	// Ascending 按更新时间升序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeCodeRepositoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCodeRepositoriesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCodeRepositoriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存储库总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 存储库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		CodeRepoSet []*CodeRepoSummary `json:"CodeRepoSet,omitempty" name:"CodeRepoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCodeRepositoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCodeRepositoriesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeCodeRepositoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCodeRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *DescribeCodeRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCodeRepositoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNotebookInstanceRequest struct {
	*tchttp.BaseRequest

	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`
}

func (r *DescribeNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotebookInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		AdditionalCodeRepositories []*string `json:"AdditionalCodeRepositories,omitempty" name:"AdditionalCodeRepositories" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotebookInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 【废弃字段】排序字段
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *DescribeNotebookInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotebookInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNotebookInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Notebook实例列表
		NotebookInstanceSet []*NotebookInstanceSummary `json:"NotebookInstanceSet,omitempty" name:"NotebookInstanceSet" list`

		// Notebook实例总数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNotebookInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotebookInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeNotebookLifecycleScriptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNotebookLifecycleScriptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *DescribeNotebookLifecycleScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotebookLifecycleScriptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 排序规则。默认取Descending
	// Descending 按更新时间降序
	// Ascending 按更新时间升序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeNotebookLifecycleScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotebookLifecycleScriptsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNotebookLifecycleScriptsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Notebook生命周期脚本列表
		NotebookLifecycleScriptsSet []*NotebookLifecycleScriptsSummary `json:"NotebookLifecycleScriptsSet,omitempty" name:"NotebookLifecycleScriptsSet" list`

		// Notebook生命周期脚本总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNotebookLifecycleScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotebookLifecycleScriptsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeTrainingJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTrainingJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 算法镜像配置
		AlgorithmSpecification *AlgorithmSpecification `json:"AlgorithmSpecification,omitempty" name:"AlgorithmSpecification"`

		// 任务名称
		TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

		// 算法超级参数
	// 注意：此字段可能返回 null，表示取不到有效值。
		HyperParameters *string `json:"HyperParameters,omitempty" name:"HyperParameters"`

		// 输入数据配置
		InputDataConfig []*InputDataConfig `json:"InputDataConfig,omitempty" name:"InputDataConfig" list`

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

		// 详细状态
		SecondaryStatus *string `json:"SecondaryStatus,omitempty" name:"SecondaryStatus"`

		// 详细状态事件记录
	// 注意：此字段可能返回 null，表示取不到有效值。
		SecondaryStatusTransitions []*SecondaryStatusTransition `json:"SecondaryStatusTransitions,omitempty" name:"SecondaryStatusTransitions" list`

		// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

		// 任务状态
		TrainingJobStatus *string `json:"TrainingJobStatus,omitempty" name:"TrainingJobStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTrainingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTrainingJobResponse) FromJsonString(s string) error {
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
	Values []*string `json:"Values,omitempty" name:"Values" list`
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

	// notebook实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookInstanceStatus *string `json:"NotebookInstanceStatus,omitempty" name:"NotebookInstanceStatus"`

	// 算力类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 算力Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

	// cos桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosOutputBucket *string `json:"CosOutputBucket,omitempty" name:"CosOutputBucket"`

	// cos文件key
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosOutputKeyPrefix *string `json:"CosOutputKeyPrefix,omitempty" name:"CosOutputKeyPrefix"`
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

type StartNotebookInstanceRequest struct {
	*tchttp.BaseRequest

	// Notebook实例名称
	NotebookInstanceName *string `json:"NotebookInstanceName,omitempty" name:"NotebookInstanceName"`
}

func (r *StartNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartNotebookInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartNotebookInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *StopNotebookInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopNotebookInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *StopTrainingJobRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopTrainingJobResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopTrainingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopTrainingJobResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StoppingCondition struct {

	// 最长运行运行时间（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRuntimeInSeconds *uint64 `json:"MaxRuntimeInSeconds,omitempty" name:"MaxRuntimeInSeconds"`
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

func (r *UpdateCodeRepositoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateCodeRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存储库名称
		CodeRepositoryName *string `json:"CodeRepositoryName,omitempty" name:"CodeRepositoryName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateCodeRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateCodeRepositoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNotebookInstanceRequest struct {
	*tchttp.BaseRequest

	// Notebook实例名称
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
	// 如果本来就没有绑定脚本，则忽略此参数；
	// 如果本来有绑定脚本，此参数为 true 则解绑；
	// 如果本来有绑定脚本，此参数为 false，则需要额外填入 LifecycleScriptsName
	DisassociateLifecycleScript *bool `json:"DisassociateLifecycleScript,omitempty" name:"DisassociateLifecycleScript"`

	// 默认存储库名称
	// 可以是已创建的存储库名称或者已https://开头的公共git库
	DefaultCodeRepository *string `json:"DefaultCodeRepository,omitempty" name:"DefaultCodeRepository"`

	// 其他存储库列表
	// 每个元素可以是已创建的存储库名称或者已https://开头的公共git库
	AdditionalCodeRepositories []*string `json:"AdditionalCodeRepositories,omitempty" name:"AdditionalCodeRepositories" list`

	// 是否取消关联默认存储库，默认false
	// 该值为true时，DefaultCodeRepository将被忽略
	DisassociateDefaultCodeRepository *bool `json:"DisassociateDefaultCodeRepository,omitempty" name:"DisassociateDefaultCodeRepository"`

	// 是否取消关联其他存储库，默认false
	// 该值为true时，AdditionalCodeRepositories将被忽略
	DisassociateAdditionalCodeRepositories *bool `json:"DisassociateAdditionalCodeRepositories,omitempty" name:"DisassociateAdditionalCodeRepositories"`
}

func (r *UpdateNotebookInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNotebookInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNotebookInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNotebookInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNotebookInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNotebookLifecycleScriptRequest struct {
	*tchttp.BaseRequest

	// notebook生命周期脚本名称
	NotebookLifecycleScriptsName *string `json:"NotebookLifecycleScriptsName,omitempty" name:"NotebookLifecycleScriptsName"`

	// 创建脚本
	CreateScript *string `json:"CreateScript,omitempty" name:"CreateScript"`

	// 启动脚本
	StartScript *string `json:"StartScript,omitempty" name:"StartScript"`
}

func (r *UpdateNotebookLifecycleScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNotebookLifecycleScriptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateNotebookLifecycleScriptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateNotebookLifecycleScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateNotebookLifecycleScriptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VpcConfig struct {

	// 安全组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
