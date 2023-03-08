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

package v20220217

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BuildPacksInfo struct {
	// 基础镜像
	BaseImage *string `json:"BaseImage,omitempty" name:"BaseImage"`

	// 启动命令
	EntryPoint *string `json:"EntryPoint,omitempty" name:"EntryPoint"`

	// 语言
	RepoLanguage *string `json:"RepoLanguage,omitempty" name:"RepoLanguage"`

	// 上传文件名
	UploadFilename *string `json:"UploadFilename,omitempty" name:"UploadFilename"`
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

// Predefined struct for user
type CreateCloudRunEnvRequestParams struct {
	// Trial,Standard,Professional,Enterprise
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

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

	// 请求key 用于防重
	ReqKey *string `json:"ReqKey,omitempty" name:"ReqKey"`

	// 来源：wechat | cloud
	Source *string `json:"Source,omitempty" name:"Source"`

	// 渠道：wechat | cloud
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

type CreateCloudRunEnvRequest struct {
	*tchttp.BaseRequest
	
	// Trial,Standard,Professional,Enterprise
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

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

	// 请求key 用于防重
	ReqKey *string `json:"ReqKey,omitempty" name:"ReqKey"`

	// 来源：wechat | cloud
	Source *string `json:"Source,omitempty" name:"Source"`

	// 渠道：wechat | cloud
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *CreateCloudRunEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRunEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageType")
	delete(f, "Alias")
	delete(f, "FreeQuota")
	delete(f, "Flag")
	delete(f, "VpcId")
	delete(f, "SubNetIds")
	delete(f, "ReqKey")
	delete(f, "Source")
	delete(f, "Channel")
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudRunEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudRunEnvResponseParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 后付费订单号
	TranId *string `json:"TranId,omitempty" name:"TranId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCloudRunEnvResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudRunEnvResponseParams `json:"Response"`
}

func (r *CreateCloudRunEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRunEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudRunServerRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 部署信息
	DeployInfo *DeployParam `json:"DeployInfo,omitempty" name:"DeployInfo"`

	// 服务配置信息
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitempty" name:"ServerConfig"`
}

type CreateCloudRunServerRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 部署信息
	DeployInfo *DeployParam `json:"DeployInfo,omitempty" name:"DeployInfo"`

	// 服务配置信息
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitempty" name:"ServerConfig"`
}

func (r *CreateCloudRunServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRunServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "DeployInfo")
	delete(f, "ServerConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudRunServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudRunServerResponseParams struct {
	// 一键部署任务Id，微信云托管，暂时用不到
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCloudRunServerResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudRunServerResponseParams `json:"Response"`
}

func (r *CreateCloudRunServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRunServerResponse) FromJsonString(s string) error {
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

type DeployParam struct {
	// 部署类型：package/image/repository/pipeline/jar/war
	DeployType *string `json:"DeployType,omitempty" name:"DeployType"`

	// 部署类型为image时传入
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 部署类型为package时传入
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 部署类型为package时传入
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 部署备注
	DeployRemark *string `json:"DeployRemark,omitempty" name:"DeployRemark"`

	// 代码仓库信息
	RepoInfo *RepositoryInfo `json:"RepoInfo,omitempty" name:"RepoInfo"`

	// 无Dockerfile时填写
	BuildPacks *BuildPacksInfo `json:"BuildPacks,omitempty" name:"BuildPacks"`

	// 发布类型 GRAY | FULL
	ReleaseType *string `json:"ReleaseType,omitempty" name:"ReleaseType"`
}

// Predefined struct for user
type DescribeCloudRunEnvsRequestParams struct {
	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitempty" name:"Channels"`
}

type DescribeCloudRunEnvsRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitempty" name:"Channels"`
}

func (r *DescribeCloudRunEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "IsVisible")
	delete(f, "Channels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRunEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunEnvsResponseParams struct {
	// 环境信息列表
	EnvList []*EnvInfo `json:"EnvList,omitempty" name:"EnvList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudRunEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudRunEnvsResponseParams `json:"Response"`
}

func (r *DescribeCloudRunEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunServerDetailRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
}

type DescribeCloudRunServerDetailRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
}

func (r *DescribeCloudRunServerDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunServerDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRunServerDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunServerDetailResponseParams struct {
	// 服务基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseInfo *ServerBaseInfo `json:"BaseInfo,omitempty" name:"BaseInfo"`

	// 服务配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitempty" name:"ServerConfig"`

	// 在线版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineVersionInfos []*OnlineVersionInfo `json:"OnlineVersionInfos,omitempty" name:"OnlineVersionInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudRunServerDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudRunServerDetailResponseParams `json:"Response"`
}

func (r *DescribeCloudRunServerDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunServerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunServersRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 默认为9， 最大为30
	// 不传或传0时 取默认9
	// 大于30时取30
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 不传或传0时 会默认为1
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
}

type DescribeCloudRunServersRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 默认为9， 最大为30
	// 不传或传0时 取默认9
	// 大于30时取30
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 不传或传0时 会默认为1
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
}

func (r *DescribeCloudRunServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "PageSize")
	delete(f, "PageNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRunServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunServersResponseParams struct {
	// 服务列表
	ServerList []*ServerBaseInfo `json:"ServerList,omitempty" name:"ServerList"`

	// 服务总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCloudRunServersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudRunServersResponseParams `json:"Response"`
}

func (r *DescribeCloudRunServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvBaseInfoRequestParams struct {
	// 环境 Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

type DescribeEnvBaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 环境 Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeEnvBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvBaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvBaseInfoResponseParams struct {
	// 环境基础信息
	EnvBaseInfo *EnvBaseInfo `json:"EnvBaseInfo,omitempty" name:"EnvBaseInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEnvBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvBaseInfoResponseParams `json:"Response"`
}

func (r *DescribeEnvBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerManageTaskRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 任务Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

type DescribeServerManageTaskRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 任务Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

func (r *DescribeServerManageTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerManageTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "TaskId")
	delete(f, "OperatorRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerManageTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerManageTaskResponseParams struct {
	// 是否存在
	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`

	// 任务信息
	Task *ServerManageTaskInfo `json:"Task,omitempty" name:"Task"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServerManageTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerManageTaskResponseParams `json:"Response"`
}

func (r *DescribeServerManageTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerManageTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvBaseInfo struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 套餐类型：Trial ｜ Standard ｜ Professional ｜ Enterprise
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// VPC Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 环境创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 环境别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 环境状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 环境地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 环境类型 tcbr ｜ run
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 子网id
	SubnetIds *string `json:"SubnetIds,omitempty" name:"SubnetIds"`
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

	// 环境类型：baas, run, hosting, weda,tcbr
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvType *string `json:"EnvType,omitempty" name:"EnvType"`

	// 数据库列表
	Databases []*DatabasesInfo `json:"Databases,omitempty" name:"Databases"`

	// 存储列表
	Storages []*StorageInfo `json:"Storages,omitempty" name:"Storages"`

	// 函数列表
	Functions []*FunctionInfo `json:"Functions,omitempty" name:"Functions"`

	// 云日志服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogServices []*LogServiceInfo `json:"LogServices,omitempty" name:"LogServices"`

	// 静态资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaticStorages []*StaticStorageInfo `json:"StaticStorages,omitempty" name:"StaticStorages"`

	// 环境标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 自定义日志服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomLogServices []*ClsInfo `json:"CustomLogServices,omitempty" name:"CustomLogServices"`

	// tcb产品套餐ID，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 套餐中文名称，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
}

type FunctionInfo struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}

type HpaPolicy struct {
	// 扩缩容类型
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 扩缩容阈值
	PolicyThreshold *uint64 `json:"PolicyThreshold,omitempty" name:"PolicyThreshold"`
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

type ObjectKV struct {
	// 键值对Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 键值对Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type OnlineVersionInfo struct {
	// 版本名
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 镜像url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowRatio *string `json:"FlowRatio,omitempty" name:"FlowRatio"`
}

// Predefined struct for user
type OperateServerManageRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 任报Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 操作类型:cancel | go_back | done
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

type OperateServerManageRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 任报Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 操作类型:cancel | go_back | done
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

func (r *OperateServerManageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OperateServerManageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "TaskId")
	delete(f, "OperateType")
	delete(f, "OperatorRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OperateServerManageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OperateServerManageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OperateServerManageResponse struct {
	*tchttp.BaseResponse
	Response *OperateServerManageResponseParams `json:"Response"`
}

func (r *OperateServerManageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OperateServerManageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseGrayRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 灰度类型
	GrayType *string `json:"GrayType,omitempty" name:"GrayType"`

	// 流量类型
	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`

	// 流量策略
	VersionFlowItems []*VersionFlowInfo `json:"VersionFlowItems,omitempty" name:"VersionFlowItems"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`

	// 流量比例
	GrayFlowRatio *int64 `json:"GrayFlowRatio,omitempty" name:"GrayFlowRatio"`
}

type ReleaseGrayRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 灰度类型
	GrayType *string `json:"GrayType,omitempty" name:"GrayType"`

	// 流量类型
	TrafficType *string `json:"TrafficType,omitempty" name:"TrafficType"`

	// 流量策略
	VersionFlowItems []*VersionFlowInfo `json:"VersionFlowItems,omitempty" name:"VersionFlowItems"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`

	// 流量比例
	GrayFlowRatio *int64 `json:"GrayFlowRatio,omitempty" name:"GrayFlowRatio"`
}

func (r *ReleaseGrayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseGrayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "GrayType")
	delete(f, "TrafficType")
	delete(f, "VersionFlowItems")
	delete(f, "OperatorRemark")
	delete(f, "GrayFlowRatio")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseGrayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseGrayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseGrayResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseGrayResponseParams `json:"Response"`
}

func (r *ReleaseGrayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseGrayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RepositoryInfo struct {
	// git source
	Source *string `json:"Source,omitempty" name:"Source"`

	// 仓库名
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// 分支名
	Branch *string `json:"Branch,omitempty" name:"Branch"`
}

type ServerBaseConfig struct {
	// 环境 Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 是否开启公网访问
	OpenAccessTypes []*string `json:"OpenAccessTypes,omitempty" name:"OpenAccessTypes"`

	// Cpu 规格
	Cpu *float64 `json:"Cpu,omitempty" name:"Cpu"`

	// Mem 规格
	Mem *float64 `json:"Mem,omitempty" name:"Mem"`

	// 最小副本数
	MinNum *uint64 `json:"MinNum,omitempty" name:"MinNum"`

	// 最大副本数
	MaxNum *uint64 `json:"MaxNum,omitempty" name:"MaxNum"`

	// 扩缩容配置
	PolicyDetails []*HpaPolicy `json:"PolicyDetails,omitempty" name:"PolicyDetails"`

	// 日志采集路径
	CustomLogs *string `json:"CustomLogs,omitempty" name:"CustomLogs"`

	// 环境变量
	EnvParams *string `json:"EnvParams,omitempty" name:"EnvParams"`

	// 延迟检测时间
	InitialDelaySeconds *uint64 `json:"InitialDelaySeconds,omitempty" name:"InitialDelaySeconds"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 服务端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 是否有Dockerfile
	HasDockerfile *bool `json:"HasDockerfile,omitempty" name:"HasDockerfile"`

	// Dockerfile 文件名
	Dockerfile *string `json:"Dockerfile,omitempty" name:"Dockerfile"`

	// 构建目录
	BuildDir *string `json:"BuildDir,omitempty" name:"BuildDir"`

	// 日志类型: none | default | custom
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// cls setId
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// cls 主题id
	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`

	// 解析类型：json ｜ line
	LogParseType *string `json:"LogParseType,omitempty" name:"LogParseType"`
}

type ServerBaseInfo struct {
	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 默认服务域名
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" name:"DefaultDomainName"`

	// 自定义域名
	CustomDomainName *string `json:"CustomDomainName,omitempty" name:"CustomDomainName"`

	// 服务状态：running/deploying/deploy_failed
	Status *string `json:"Status,omitempty" name:"Status"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 公网访问类型
	AccessTypes []*string `json:"AccessTypes,omitempty" name:"AccessTypes"`

	// 展示自定义域名
	CustomDomainNames []*string `json:"CustomDomainNames,omitempty" name:"CustomDomainNames"`
}

type ServerManageTaskInfo struct {
	// 任务Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 变更类型
	ChangeType *string `json:"ChangeType,omitempty" name:"ChangeType"`

	// 发布类型
	ReleaseType *string `json:"ReleaseType,omitempty" name:"ReleaseType"`

	// 部署类型
	DeployType *string `json:"DeployType,omitempty" name:"DeployType"`

	// 上一个版本名
	PreVersionName *string `json:"PreVersionName,omitempty" name:"PreVersionName"`

	// 版本名
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 流水线Id
	PipelineId *int64 `json:"PipelineId,omitempty" name:"PipelineId"`

	// 流水线任务Id
	PipelineTaskId *int64 `json:"PipelineTaskId,omitempty" name:"PipelineTaskId"`

	// 发布单Id
	ReleaseId *int64 `json:"ReleaseId,omitempty" name:"ReleaseId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 步骤信息
	Steps []*TaskStepInfo `json:"Steps,omitempty" name:"Steps"`

	// 失败原因
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitempty" name:"OperatorRemark"`
}

type StaticStorageInfo struct {
	// 静态CDN域名
	StaticDomain *string `json:"StaticDomain,omitempty" name:"StaticDomain"`

	// 静态CDN默认文件夹，当前为根目录
	DefaultDirName *string `json:"DefaultDirName,omitempty" name:"DefaultDirName"`

	// 资源状态(process/online/offline/init)
	Status *string `json:"Status,omitempty" name:"Status"`

	// cos所属区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// bucket信息
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

type TaskStepInfo struct {
	// 步骤名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 未启动："todo"
	// 运行中："running"
	// 失败："failed"
	// 成功结束："finished"
	Status *string `json:"Status,omitempty" name:"Status"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 消耗时间：秒
	CostTime *int64 `json:"CostTime,omitempty" name:"CostTime"`

	// 失败原因
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
}

// Predefined struct for user
type UpdateCloudRunServerRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 部署信息
	DeployInfo *DeployParam `json:"DeployInfo,omitempty" name:"DeployInfo"`

	// 服务配置信息
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitempty" name:"ServerConfig"`
}

type UpdateCloudRunServerRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 部署信息
	DeployInfo *DeployParam `json:"DeployInfo,omitempty" name:"DeployInfo"`

	// 服务配置信息
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitempty" name:"ServerConfig"`
}

func (r *UpdateCloudRunServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCloudRunServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "DeployInfo")
	delete(f, "ServerConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCloudRunServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCloudRunServerResponseParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 一键部署任务Id，暂时用不到
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateCloudRunServerResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCloudRunServerResponseParams `json:"Response"`
}

func (r *UpdateCloudRunServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCloudRunServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionFlowInfo struct {
	// 版本名
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 是否默认版本
	IsDefaultPriority *bool `json:"IsDefaultPriority,omitempty" name:"IsDefaultPriority"`

	// 流量比例
	FlowRatio *int64 `json:"FlowRatio,omitempty" name:"FlowRatio"`

	// 测试KV值
	UrlParam *ObjectKV `json:"UrlParam,omitempty" name:"UrlParam"`

	// 权重
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}