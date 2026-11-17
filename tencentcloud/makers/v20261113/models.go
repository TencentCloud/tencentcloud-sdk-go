// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20261113

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CosCredentials struct {
	// <p>临时 SecretId</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// <p>临时 SecretKey</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// <p>临时 Token</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

// Predefined struct for user
type CreateDeploymentRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>触发来源。Github、ReDeploy、Api、Webhook、Upload、Coding、Gitee、Gitlab</p>
	ViaMeta *string `json:"ViaMeta,omitnil,omitempty" name:"ViaMeta"`

	// <p>环境 slug。Production、Preview 或自定义环境</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>代码来源。上传制品传 Upload</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>制品类型。Zip 或 Folder</p>
	DistType *string `json:"DistType,omitnil,omitempty" name:"DistType"`

	// <p>已上传制品在临时桶中的对象路径</p>
	TempBucketPath *string `json:"TempBucketPath,omitnil,omitempty" name:"TempBucketPath"`

	// <p>构建分支</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>提交哈希</p>
	RepoCommitHash *string `json:"RepoCommitHash,omitnil,omitempty" name:"RepoCommitHash"`

	// <p>提交说明</p>
	RepoCommitMsg *string `json:"RepoCommitMsg,omitnil,omitempty" name:"RepoCommitMsg"`

	// <p>构建来源</p>
	BuildFrom *string `json:"BuildFrom,omitnil,omitempty" name:"BuildFrom"`

	// <p>模板 ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>构建规格。当前服务端仍从 edgeone.json 读取，此字段先保留</p>
	CiSpec *string `json:"CiSpec,omitnil,omitempty" name:"CiSpec"`
}

type CreateDeploymentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>触发来源。Github、ReDeploy、Api、Webhook、Upload、Coding、Gitee、Gitlab</p>
	ViaMeta *string `json:"ViaMeta,omitnil,omitempty" name:"ViaMeta"`

	// <p>环境 slug。Production、Preview 或自定义环境</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>代码来源。上传制品传 Upload</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>制品类型。Zip 或 Folder</p>
	DistType *string `json:"DistType,omitnil,omitempty" name:"DistType"`

	// <p>已上传制品在临时桶中的对象路径</p>
	TempBucketPath *string `json:"TempBucketPath,omitnil,omitempty" name:"TempBucketPath"`

	// <p>构建分支</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>提交哈希</p>
	RepoCommitHash *string `json:"RepoCommitHash,omitnil,omitempty" name:"RepoCommitHash"`

	// <p>提交说明</p>
	RepoCommitMsg *string `json:"RepoCommitMsg,omitnil,omitempty" name:"RepoCommitMsg"`

	// <p>构建来源</p>
	BuildFrom *string `json:"BuildFrom,omitnil,omitempty" name:"BuildFrom"`

	// <p>模板 ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>构建规格。当前服务端仍从 edgeone.json 读取，此字段先保留</p>
	CiSpec *string `json:"CiSpec,omitnil,omitempty" name:"CiSpec"`
}

func (r *CreateDeploymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeploymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "ViaMeta")
	delete(f, "Env")
	delete(f, "Provider")
	delete(f, "DistType")
	delete(f, "TempBucketPath")
	delete(f, "RepoBranch")
	delete(f, "RepoCommitHash")
	delete(f, "RepoCommitMsg")
	delete(f, "BuildFrom")
	delete(f, "TemplateId")
	delete(f, "CiSpec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeploymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeploymentResponseParams struct {
	// <p>部署 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeploymentResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeploymentResponseParams `json:"Response"`
}

func (r *CreateDeploymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>创建来源。Template、Custom、CLI、Guest</p>
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// <p>代码来源。Upload、Github、Gitee、Gitlab、Coding</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>接入方标识</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>加速区域。mainland、overseas、global</p>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>Node.js 版本。14.21.3、16.20.2、18.20.4、20.18.0、22.11.0、22.17.1、22.21.1、24.5.0、24.11.0、24.18.0</p>
	NodeJsVersion *string `json:"NodeJsVersion,omitnil,omitempty" name:"NodeJsVersion"`

	// <p>项目根目录</p>
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>构建产物目录</p>
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>构建命令</p>
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>安装命令</p>
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>框架</p>
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>模板 ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>仓库地址</p>
	RepoUrl *string `json:"RepoUrl,omitnil,omitempty" name:"RepoUrl"`

	// <p>仓库所有者</p>
	RepoOwner *string `json:"RepoOwner,omitnil,omitempty" name:"RepoOwner"`

	// <p>仓库名称</p>
	RepoName *string `json:"RepoName,omitnil,omitempty" name:"RepoName"`

	// <p>生产分支</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>初始环境变量</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// <p>Coding 命名空间。Provider 为 Coding 时使用</p>
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// <p>Coding 项目 ID。Provider 为 Coding 时使用</p>
	CodingProjectId *int64 `json:"CodingProjectId,omitnil,omitempty" name:"CodingProjectId"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>创建来源。Template、Custom、CLI、Guest</p>
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// <p>代码来源。Upload、Github、Gitee、Gitlab、Coding</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>接入方标识</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>加速区域。mainland、overseas、global</p>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>Node.js 版本。14.21.3、16.20.2、18.20.4、20.18.0、22.11.0、22.17.1、22.21.1、24.5.0、24.11.0、24.18.0</p>
	NodeJsVersion *string `json:"NodeJsVersion,omitnil,omitempty" name:"NodeJsVersion"`

	// <p>项目根目录</p>
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>构建产物目录</p>
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>构建命令</p>
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>安装命令</p>
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>框架</p>
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>模板 ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>仓库地址</p>
	RepoUrl *string `json:"RepoUrl,omitnil,omitempty" name:"RepoUrl"`

	// <p>仓库所有者</p>
	RepoOwner *string `json:"RepoOwner,omitnil,omitempty" name:"RepoOwner"`

	// <p>仓库名称</p>
	RepoName *string `json:"RepoName,omitnil,omitempty" name:"RepoName"`

	// <p>生产分支</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>初始环境变量</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// <p>Coding 命名空间。Provider 为 Coding 时使用</p>
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// <p>Coding 项目 ID。Provider 为 Coding 时使用</p>
	CodingProjectId *int64 `json:"CodingProjectId,omitnil,omitempty" name:"CodingProjectId"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Name")
	delete(f, "Channel")
	delete(f, "Provider")
	delete(f, "Source")
	delete(f, "Area")
	delete(f, "NodeJsVersion")
	delete(f, "RootDir")
	delete(f, "OutputDir")
	delete(f, "BuildCmd")
	delete(f, "InstallCmd")
	delete(f, "Framework")
	delete(f, "TemplateId")
	delete(f, "RepoUrl")
	delete(f, "RepoOwner")
	delete(f, "RepoName")
	delete(f, "RepoBranch")
	delete(f, "EnvVars")
	delete(f, "Namespace")
	delete(f, "CodingProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// <p>项目 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneCustomDomainRequestParams struct {
	// <p>站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>自定义域名</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>环境 slug。不传时由后端决定</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type CreateZoneCustomDomainRequest struct {
	*tchttp.BaseRequest
	
	// <p>站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>自定义域名</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>环境 slug。不传时由后端决定</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *CreateZoneCustomDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneCustomDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "Domain")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateZoneCustomDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneCustomDomainResponseParams struct {
	// <p>域名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>域名配置 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateZoneCustomDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateZoneCustomDomainResponseParams `json:"Response"`
}

func (r *CreateZoneCustomDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneCustomDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomDomainInput struct {
	// <p>域名</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>证书 ID</p>
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// <p>校验状态。Verify、Pass、Fail</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>是否不申请证书</p>
	NoCertificate *bool `json:"NoCertificate,omitnil,omitempty" name:"NoCertificate"`
}

type CustomDomainItem struct {
	// <p>域名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

// Predefined struct for user
type DeleteProjectCustomDomainsRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>要解绑的域名</p>
	CustomDomains []*string `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`
}

type DeleteProjectCustomDomainsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>要解绑的域名</p>
	CustomDomains []*string `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`
}

func (r *DeleteProjectCustomDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectCustomDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "CustomDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectCustomDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectCustomDomainsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectCustomDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectCustomDomainsResponseParams `json:"Response"`
}

func (r *DeleteProjectCustomDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectCustomDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectEnvsRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>要删除的环境变量。传 Key（或 Id），Value 必填但不参与匹配</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type DeleteProjectEnvsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>要删除的环境变量。传 Key（或 Id），Value 必填但不参与匹配</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

func (r *DeleteProjectEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "EnvVars")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectEnvsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectEnvsResponseParams `json:"Response"`
}

func (r *DeleteProjectEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectResponseParams `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Deployment struct {
	// <p>部署 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`

	// <p>项目 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Success、Failed、Process、Invalid、Pending、Timeout、Cancelled</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>部署被移除前的原始状态</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalStatus *string `json:"OriginalStatus,omitnil,omitempty" name:"OriginalStatus"`

	// <p>构建退出码，0 表示成功</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>构建耗时，单位毫秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildCost *string `json:"BuildCost,omitnil,omitempty" name:"BuildCost"`

	// <p>触发来源。Github、ReDeploy、Api、Coding、Gitee、Upload、Gitlab</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViaMeta *string `json:"ViaMeta,omitnil,omitempty" name:"ViaMeta"`

	// <p>环境 slug</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>分支</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>提交哈希</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoCommitHash *string `json:"RepoCommitHash,omitnil,omitempty" name:"RepoCommitHash"`

	// <p>提交说明</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoCommitMsg *string `json:"RepoCommitMsg,omitnil,omitempty" name:"RepoCommitMsg"`

	// <p>本次部署的访问地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// <p>项目默认域名地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectUrl *string `json:"ProjectUrl,omitnil,omitempty" name:"ProjectUrl"`

	// <p>是否为当前生产版本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedInProd *bool `json:"UsedInProd,omitnil,omitempty" name:"UsedInProd"`

	// <p>封面地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// <p>路由配置，JSON 字符串</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *string `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// <p>本次部署使用的仓库地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DplRepoUrl *string `json:"DplRepoUrl,omitnil,omitempty" name:"DplRepoUrl"`

	// <p>创建时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// <p>更新时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type DeploymentFilter struct {
	// <p>过滤字段。ProjectId、DeploymentId、Name、Status、StartTime、EndTime、RepoBranch</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>过滤值</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type DescribeCosTempTokenRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目名称。临时路径按它拼接</p>
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// <p>项目 ID。不传 ProjectName 时用来反查项目名</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeCosTempTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目名称。临时路径按它拼接</p>
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// <p>项目 ID。不传 ProjectName 时用来反查项目名</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeCosTempTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosTempTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectName")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosTempTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosTempTokenResponseParams struct {
	// <p>COS 桶</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// <p>COS 地域</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>上传路径，格式 appId/projectName/时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// <p>临时密钥过期时间，Unix 秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// <p>临时密钥过期时间，ISO 8601</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expiration *string `json:"Expiration,omitnil,omitempty" name:"Expiration"`

	// <p>临时密钥</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Credentials *CosCredentials `json:"Credentials,omitnil,omitempty" name:"Credentials"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosTempTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosTempTokenResponseParams `json:"Response"`
}

func (r *DescribeCosTempTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosTempTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentLogRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>部署 ID</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

type DescribeDeploymentLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>部署 ID</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

func (r *DescribeDeploymentLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "DeploymentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeploymentLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentLogResponseParams struct {
	// <p>构建日志地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogUrl *string `json:"LogUrl,omitnil,omitempty" name:"LogUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeploymentLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeploymentLogResponseParams `json:"Response"`
}

func (r *DescribeDeploymentLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentsRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>按部署 ID 查询</p>
	DeploymentIds []*string `json:"DeploymentIds,omitnil,omitempty" name:"DeploymentIds"`

	// <p>过滤条件</p>
	Filters []*DeploymentFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>偏移</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>CreatedOn 或 ModifiedOn</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Asc 或 Desc</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeDeploymentsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>按部署 ID 查询</p>
	DeploymentIds []*string `json:"DeploymentIds,omitnil,omitempty" name:"DeploymentIds"`

	// <p>过滤条件</p>
	Filters []*DeploymentFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>偏移</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>CreatedOn 或 ModifiedOn</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Asc 或 Desc</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeDeploymentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "DeploymentIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeploymentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentsResponseParams struct {
	// <p>总数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>部署列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployments []*Deployment `json:"Deployments,omitnil,omitempty" name:"Deployments"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeploymentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeploymentsResponseParams `json:"Response"`
}

func (r *DescribeDeploymentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// <p>总数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>环境列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environments []*Environment `json:"Environments,omitnil,omitempty" name:"Environments"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeProjectEnvsRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>只返回该环境 slug 的变量</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeProjectEnvsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>只返回该环境 slug 的变量</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeProjectEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectEnvsResponseParams struct {
	// <p>环境变量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectEnvsResponseParams `json:"Response"`
}

func (r *DescribeProjectEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>页码偏移，从 0 开始。后端跳过 Offset * Limit 条</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件。Name 取值 ProjectId、Name、Status、Provider</p>
	Filters []*ProjectFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>按项目 ID 查询</p>
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// <p>排序字段。CreatedOn、ModifiedOn。默认 CreatedOn</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>排序方向。Asc、Desc、ASC、DESC。默认 DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>来源。drop 时只返回 Upload 项目</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>按项目名称模糊搜索</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>页码偏移，从 0 开始。后端跳过 Offset * Limit 条</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量，最大 100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤条件。Name 取值 ProjectId、Name、Status、Provider</p>
	Filters []*ProjectFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>按项目 ID 查询</p>
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// <p>排序字段。CreatedOn、ModifiedOn。默认 CreatedOn</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>排序方向。Asc、Desc、ASC、DESC。默认 DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>来源。drop 时只返回 Upload 项目</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>按项目名称模糊搜索</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "ProjectIds")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Source")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsResponseParams struct {
	// <p>总数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>项目列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Projects []*Project `json:"Projects,omitnil,omitempty" name:"Projects"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectsResponseParams `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneCustomDomainsRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>只查询这些域名。不传则返回项目下的域名</p>
	CustomDomains []*CustomDomainInput `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`
}

type DescribeZoneCustomDomainsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>只查询这些域名。不传则返回项目下的域名</p>
	CustomDomains []*CustomDomainInput `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`
}

func (r *DescribeZoneCustomDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneCustomDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "CustomDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneCustomDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneCustomDomainsResponseParams struct {
	// <p>总数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>域名列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PagesDomains []*ZoneDomain `json:"PagesDomains,omitnil,omitempty" name:"PagesDomains"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZoneCustomDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneCustomDomainsResponseParams `json:"Response"`
}

func (r *DescribeZoneCustomDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneCustomDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsVerification struct {
	// <p>记录类型，TXT</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// <p>主机记录</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// <p>记录值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
}

type EnvVar struct {
	// <p>变量名，最长 255</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>变量值，最长 1000</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>变量 ID。创建时不传</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>备注，最长 255</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// <p>生效环境 slug，例如 Production、Preview。同 Key 的多行环境不能重叠</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Env []*string `json:"Env,omitnil,omitempty" name:"Env"`
}

type Environment struct {
	// <p>环境标识，创建后不可改</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Slug *string `json:"Slug,omitnil,omitempty" name:"Slug"`

	// <p>环境名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Production、Preview、Custom</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>关联分支，支持通配</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Branches []*string `json:"Branches,omitnil,omitempty" name:"Branches"`

	// <p>是否自动部署。Production 恒为 true</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoDeploy *bool `json:"AutoDeploy,omitnil,omitempty" name:"AutoDeploy"`

	// <p>绑定的自定义域名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomDomains []*string `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`

	// <p>环境变量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvVars []*EnvironmentEnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// <p>创建时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// <p>更新时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type EnvironmentEnvVar struct {
	// <p>变量 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>变量名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>变量值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>备注</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// <p>生效环境 slug</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Env []*string `json:"Env,omitnil,omitempty" name:"Env"`
}

type FileVerification struct {
	// <p>校验文件路径</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>校验文件内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ForceRedirectConfig struct {
	// <p>on 或 off</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// <p>301、302、307、308</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitnil,omitempty" name:"RedirectStatusCode"`
}

type HstsConfig struct {
	// <p>on 或 off</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// <p>有效期，秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *int64 `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// <p>on 或 off</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeSubDomains *string `json:"IncludeSubDomains,omitnil,omitempty" name:"IncludeSubDomains"`

	// <p>on 或 off</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preload *string `json:"Preload,omitnil,omitempty" name:"Preload"`
}

// Predefined struct for user
type ModifyProjectEnvsRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>要写入的环境变量</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type ModifyProjectEnvsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>要写入的环境变量</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

func (r *ModifyProjectEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "EnvVars")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectEnvsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectEnvsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectEnvsResponseParams `json:"Response"`
}

func (r *ModifyProjectEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>项目名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>项目根目录</p>
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>构建产物目录</p>
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>构建命令</p>
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>安装命令</p>
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>框架</p>
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>封面地址</p>
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// <p>Node.js 版本</p>
	NodeJsVersion *string `json:"NodeJsVersion,omitnil,omitempty" name:"NodeJsVersion"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers 站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>项目名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>项目根目录</p>
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>构建产物目录</p>
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>构建命令</p>
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>安装命令</p>
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>框架</p>
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>封面地址</p>
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// <p>Node.js 版本</p>
	NodeJsVersion *string `json:"NodeJsVersion,omitnil,omitempty" name:"NodeJsVersion"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "RootDir")
	delete(f, "OutputDir")
	delete(f, "BuildCmd")
	delete(f, "InstallCmd")
	delete(f, "Framework")
	delete(f, "CoverUrl")
	delete(f, "NodeJsVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectResponseParams `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneCustomDomainEnvRequestParams struct {
	// <p>站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>域名</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>目标环境 slug</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type ModifyZoneCustomDomainEnvRequest struct {
	*tchttp.BaseRequest
	
	// <p>站点 ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>项目 ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>域名</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>目标环境 slug</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *ModifyZoneCustomDomainEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneCustomDomainEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "Domain")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneCustomDomainEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneCustomDomainEnvResponseParams struct {
	// <p>域名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>切换后的环境 slug</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyZoneCustomDomainEnvResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneCustomDomainEnvResponseParams `json:"Response"`
}

func (r *ModifyZoneCustomDomainEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneCustomDomainEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NsVerification struct {
	// <p>NS 服务器</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`
}

type OwnershipVerification struct {
	// <p>DNS 校验</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsVerification *DnsVerification `json:"DnsVerification,omitnil,omitempty" name:"DnsVerification"`

	// <p>文件校验</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileVerification *FileVerification `json:"FileVerification,omitnil,omitempty" name:"FileVerification"`

	// <p>NS 校验</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NsVerification *NsVerification `json:"NsVerification,omitnil,omitempty" name:"NsVerification"`
}

type Project struct {
	// <p>项目 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>项目名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>项目状态。Normal、Ban、Unlink、Process、Pending、Offline</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>站点 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>加速区域。mainland、overseas、global</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>代码来源。Github、Gitee、Coding、Upload、Gitlab、Bitbucket、Cnb</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>创建来源。Template、Custom、CLI、Guest</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// <p>预设域名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetDomain *string `json:"PresetDomain,omitnil,omitempty" name:"PresetDomain"`

	// <p>预设域名状态</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetDomainStatus *string `json:"PresetDomainStatus,omitnil,omitempty" name:"PresetDomainStatus"`

	// <p>项目根目录</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>构建产物目录</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>构建命令</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>安装命令</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>框架</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>Node.js 版本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodejsVersion *string `json:"NodejsVersion,omitnil,omitempty" name:"NodejsVersion"`

	// <p>仓库地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoUrl *string `json:"RepoUrl,omitnil,omitempty" name:"RepoUrl"`

	// <p>仓库所有者</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoOwner *string `json:"RepoOwner,omitnil,omitempty" name:"RepoOwner"`

	// <p>仓库名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitnil,omitempty" name:"RepoName"`

	// <p>生产分支</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>预览分支</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoPreBranches []*string `json:"RepoPreBranches,omitnil,omitempty" name:"RepoPreBranches"`

	// <p>部署地域</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regions *Regions `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>环境变量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvVars []*ProjectEnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// <p>自定义域名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomDomains []*CustomDomainItem `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`

	// <p>最近一次部署</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployment *ProjectDeployment `json:"Deployment,omitnil,omitempty" name:"Deployment"`

	// <p>是否 Agent 项目。1 是，0 否</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAgent *int64 `json:"IsAgent,omitnil,omitempty" name:"IsAgent"`

	// <p>Agent 框架。非 Agent 项目为空</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentFramework *string `json:"AgentFramework,omitnil,omitempty" name:"AgentFramework"`

	// <p>是否命中用户级泛域名。1 是，0 否</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsTld *int64 `json:"IsTld,omitnil,omitempty" name:"IsTld"`

	// <p>创建时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// <p>更新时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type ProjectDeployment struct {
	// <p>部署 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`

	// <p>项目 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>部署状态</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>构建耗时，单位毫秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildCost *string `json:"BuildCost,omitnil,omitempty" name:"BuildCost"`

	// <p>触发来源</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViaMeta *string `json:"ViaMeta,omitnil,omitempty" name:"ViaMeta"`

	// <p>环境 slug</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>分支</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>提交哈希</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoCommitHash *string `json:"RepoCommitHash,omitnil,omitempty" name:"RepoCommitHash"`

	// <p>提交说明</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoCommitMsg *string `json:"RepoCommitMsg,omitnil,omitempty" name:"RepoCommitMsg"`

	// <p>预览地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// <p>封面地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// <p>创建时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// <p>更新时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`

	// <p>本次部署使用的仓库地址</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DplRepoUrl *string `json:"DplRepoUrl,omitnil,omitempty" name:"DplRepoUrl"`
}

type ProjectEnvVar struct {
	// <p>变量名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>变量值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ProjectFilter struct {
	// <p>过滤字段。ProjectId、Name、Status、Provider</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>过滤值</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Regions struct {
	// <p>中国大陆地域</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mainland []*string `json:"Mainland,omitnil,omitempty" name:"Mainland"`

	// <p>海外地域</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Overseas []*string `json:"Overseas,omitnil,omitempty" name:"Overseas"`
}

type SwitchConfig struct {
	// <p>on 或 off</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ZoneDomain struct {
	// <p>Custom 或 Preset</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>域名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>域名状态</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>绑定的环境 slug，默认 Production</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>站点 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>加速区域</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>CNAME</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// <p>主域名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDomain *string `json:"MainDomain,omitnil,omitempty" name:"MainDomain"`

	// <p>IPv6</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *SwitchConfig `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// <p>HTTP/2</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http2 *SwitchConfig `json:"Http2,omitnil,omitempty" name:"Http2"`

	// <p>OCSP</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcspStapling *SwitchConfig `json:"OcspStapling,omitnil,omitempty" name:"OcspStapling"`

	// <p>强制 HTTPS</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirectConfig `json:"ForceRedirect,omitnil,omitempty" name:"ForceRedirect"`

	// <p>HSTS</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hsts *HstsConfig `json:"Hsts,omitnil,omitempty" name:"Hsts"`

	// <p>归属校验</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`
}