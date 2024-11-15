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

package v20230508

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateWorkspaceRequestParams struct {
	// 工作空间名称, 长度限制 2~64
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 工作空间描述, 长度限制 0~255
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间规格。Standard: 2C4G, Calculation: 4C8G, Profession: 8C16G. 默认是 Standard。
	Specs *string `json:"Specs,omitnil,omitempty" name:"Specs"`

	// 工作空间基础镜像名称, 默认会使用 All In One 镜像, 长度限制 1~255
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Git 仓库. 工作空间启动时会自动克隆该仓库
	Repository *GitRepository `json:"Repository,omitnil,omitempty" name:"Repository"`

	// 环境变量. 会被注入到工作空间中
	Envs []*Env `json:"Envs,omitnil,omitempty" name:"Envs"`

	// 预装插件. 工作空间启动时, 会自动安装这些插件。长度限制: 0~10
	Extensions []*string `json:"Extensions,omitnil,omitempty" name:"Extensions"`

	// 工作空间生命周期钩子.  分为三个阶段 init, start, destroy. 分别表示工作空间数据初始化阶段, 工作空间启动阶段, 工作空间关闭阶段.  用户可以自定义 shell 命令. 
	Lifecycle *LifeCycle `json:"Lifecycle,omitnil,omitempty" name:"Lifecycle"`

	// 应用名称
	TenantAppId *int64 `json:"TenantAppId,omitnil,omitempty" name:"TenantAppId"`

	// 用户UIN
	TenantUin *string `json:"TenantUin,omitnil,omitempty" name:"TenantUin"`

	// VPCID
	TenantUniqVpcId *string `json:"TenantUniqVpcId,omitnil,omitempty" name:"TenantUniqVpcId"`

	// 子网ID
	TenantSubnetId *string `json:"TenantSubnetId,omitnil,omitempty" name:"TenantSubnetId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type CreateWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间名称, 长度限制 2~64
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 工作空间描述, 长度限制 0~255
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间规格。Standard: 2C4G, Calculation: 4C8G, Profession: 8C16G. 默认是 Standard。
	Specs *string `json:"Specs,omitnil,omitempty" name:"Specs"`

	// 工作空间基础镜像名称, 默认会使用 All In One 镜像, 长度限制 1~255
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// Git 仓库. 工作空间启动时会自动克隆该仓库
	Repository *GitRepository `json:"Repository,omitnil,omitempty" name:"Repository"`

	// 环境变量. 会被注入到工作空间中
	Envs []*Env `json:"Envs,omitnil,omitempty" name:"Envs"`

	// 预装插件. 工作空间启动时, 会自动安装这些插件。长度限制: 0~10
	Extensions []*string `json:"Extensions,omitnil,omitempty" name:"Extensions"`

	// 工作空间生命周期钩子.  分为三个阶段 init, start, destroy. 分别表示工作空间数据初始化阶段, 工作空间启动阶段, 工作空间关闭阶段.  用户可以自定义 shell 命令. 
	Lifecycle *LifeCycle `json:"Lifecycle,omitnil,omitempty" name:"Lifecycle"`

	// 应用名称
	TenantAppId *int64 `json:"TenantAppId,omitnil,omitempty" name:"TenantAppId"`

	// 用户UIN
	TenantUin *string `json:"TenantUin,omitnil,omitempty" name:"TenantUin"`

	// VPCID
	TenantUniqVpcId *string `json:"TenantUniqVpcId,omitnil,omitempty" name:"TenantUniqVpcId"`

	// 子网ID
	TenantSubnetId *string `json:"TenantSubnetId,omitnil,omitempty" name:"TenantSubnetId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

func (r *CreateWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Specs")
	delete(f, "Image")
	delete(f, "Repository")
	delete(f, "Envs")
	delete(f, "Extensions")
	delete(f, "Lifecycle")
	delete(f, "TenantAppId")
	delete(f, "TenantUin")
	delete(f, "TenantUniqVpcId")
	delete(f, "TenantSubnetId")
	delete(f, "Region")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceResponseParams struct {
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`

	// 工作空间名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceResponseParams `json:"Response"`
}

func (r *CreateWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceTokenRequestParams struct {
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`

	// token过期时间，单位是秒，默认 3600
	TokenExpiredLimitSec *uint64 `json:"TokenExpiredLimitSec,omitnil,omitempty" name:"TokenExpiredLimitSec"`

	// token 授权策略，可选值为 workspace-run-only, all。默认为 workspace-run-only
	Policies []*string `json:"Policies,omitnil,omitempty" name:"Policies"`
}

type CreateWorkspaceTokenRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`

	// token过期时间，单位是秒，默认 3600
	TokenExpiredLimitSec *uint64 `json:"TokenExpiredLimitSec,omitnil,omitempty" name:"TokenExpiredLimitSec"`

	// token 授权策略，可选值为 workspace-run-only, all。默认为 workspace-run-only
	Policies []*string `json:"Policies,omitnil,omitempty" name:"Policies"`
}

func (r *CreateWorkspaceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	delete(f, "TokenExpiredLimitSec")
	delete(f, "Policies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceTokenResponseParams struct {
	// 访问工作空间临时凭证
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// token 过期时间
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkspaceTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceTokenResponseParams `json:"Response"`
}

func (r *CreateWorkspaceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRequestParams struct {
	// 配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigResponseParams struct {
	// 配置值
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeImagesRequestParams struct {

}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// 镜像列表
	Images []*Image `json:"Images,omitnil,omitempty" name:"Images"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspacesRequestParams struct {
	// 工作空间名称过滤条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeWorkspacesRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间名称过滤条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeWorkspacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspacesResponseParams struct {
	// 工作空间列表
	Data []*WorkspaceStatusInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkspacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspacesResponseParams `json:"Response"`
}

func (r *DescribeWorkspacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Env struct {
	// 环境变量 key
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境变量 value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type GitRepository struct {
	// Git 仓库地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Git 仓库分支名或 Tag 名
	Branch *string `json:"Branch,omitnil,omitempty" name:"Branch"`
}

type Image struct {
	// 镜像名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 镜像仓库
	Repository *string `json:"Repository,omitnil,omitempty" name:"Repository"`

	// tag 列表
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type LifeCycle struct {
	// 工作空间首次初始化时执行
	Init []*LifeCycleCommand `json:"Init,omitnil,omitempty" name:"Init"`

	// 每次工作空间启动时执行
	Start []*LifeCycleCommand `json:"Start,omitnil,omitempty" name:"Start"`

	// 每次工作空间关闭时执行
	Destroy []*LifeCycleCommand `json:"Destroy,omitnil,omitempty" name:"Destroy"`
}

type LifeCycleCommand struct {
	// 指令描述
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 具体命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`
}

// Predefined struct for user
type ModifyWorkspaceRequestParams struct {
	// 工作空间 SpaceKey. 更新该工作空间的属性
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`

	// 工作空间名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 工作空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间规格。STANDARD: 2C4G, CALCULATION: 4C8G, PROFESSION: 8C16G. 默认是 STANDARD。
	Specs *string `json:"Specs,omitnil,omitempty" name:"Specs"`

	// 环境变量. 会被注入到工作空间中
	Envs []*Env `json:"Envs,omitnil,omitempty" name:"Envs"`

	// 预装插件. 工作空间启动时, 会自动安装这些插件 
	Extensions []*string `json:"Extensions,omitnil,omitempty" name:"Extensions"`

	// 工作空间生命周期钩子.  分为三个阶段 init, start, destroy. 分别表示工作空间数据初始化阶段, 工作空间启动阶段, 工作空间关闭阶段.  用户可以自定义 shell 命令. 
	Lifecycle *LifeCycle `json:"Lifecycle,omitnil,omitempty" name:"Lifecycle"`
}

type ModifyWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间 SpaceKey. 更新该工作空间的属性
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`

	// 工作空间名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 工作空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间规格。STANDARD: 2C4G, CALCULATION: 4C8G, PROFESSION: 8C16G. 默认是 STANDARD。
	Specs *string `json:"Specs,omitnil,omitempty" name:"Specs"`

	// 环境变量. 会被注入到工作空间中
	Envs []*Env `json:"Envs,omitnil,omitempty" name:"Envs"`

	// 预装插件. 工作空间启动时, 会自动安装这些插件 
	Extensions []*string `json:"Extensions,omitnil,omitempty" name:"Extensions"`

	// 工作空间生命周期钩子.  分为三个阶段 init, start, destroy. 分别表示工作空间数据初始化阶段, 工作空间启动阶段, 工作空间关闭阶段.  用户可以自定义 shell 命令. 
	Lifecycle *LifeCycle `json:"Lifecycle,omitnil,omitempty" name:"Lifecycle"`
}

func (r *ModifyWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Specs")
	delete(f, "Envs")
	delete(f, "Extensions")
	delete(f, "Lifecycle")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkspaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkspaceResponseParams `json:"Response"`
}

func (r *ModifyWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveWorkspaceRequestParams struct {
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`
}

type RemoveWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`
}

func (r *RemoveWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveWorkspaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *RemoveWorkspaceResponseParams `json:"Response"`
}

func (r *RemoveWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkspaceRequestParams struct {
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`
}

type RunWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`
}

func (r *RunWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunWorkspaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *RunWorkspaceResponseParams `json:"Response"`
}

func (r *RunWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWorkspaceRequestParams struct {
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`
}

type StopWorkspaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间 SpaceKey
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`
}

func (r *StopWorkspaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkspaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopWorkspaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWorkspaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopWorkspaceResponse struct {
	*tchttp.BaseResponse
	Response *StopWorkspaceResponseParams `json:"Response"`
}

func (r *StopWorkspaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkspaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WorkspaceStatusInfo struct {
	// 工作空间 ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 工作空间名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 工作空间标识
	SpaceKey *string `json:"SpaceKey,omitnil,omitempty" name:"SpaceKey"`

	// 工作空间状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CPU数量
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 工作空间图标
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 工作空间状态, 异常原因
	StatusReason *string `json:"StatusReason,omitnil,omitempty" name:"StatusReason"`

	// 工作空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 工作空间类型
	WorkspaceType *string `json:"WorkspaceType,omitnil,omitempty" name:"WorkspaceType"`

	// Git 仓库 HTTPS 地址
	VersionControlUrl *string `json:"VersionControlUrl,omitnil,omitempty" name:"VersionControlUrl"`

	// Git 仓库引用。指定分支使用 /refs/heads/{分支名}, 指定 Tag 用 /refs/tags/{Tag名}
	VersionControlRef *string `json:"VersionControlRef,omitnil,omitempty" name:"VersionControlRef"`

	// 最后操作时间
	LastOpsDate *string `json:"LastOpsDate,omitnil,omitempty" name:"LastOpsDate"`

	// 创建时间
	CreateDate *string `json:"CreateDate,omitnil,omitempty" name:"CreateDate"`
}