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

package v20220217

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type BuildPacksInfo struct {
	// 基础镜像
	BaseImage *string `json:"BaseImage,omitnil,omitempty" name:"BaseImage"`

	// 启动命令
	EntryPoint *string `json:"EntryPoint,omitnil,omitempty" name:"EntryPoint"`

	// 语言
	RepoLanguage *string `json:"RepoLanguage,omitnil,omitempty" name:"RepoLanguage"`

	// 上传文件名
	UploadFilename *string `json:"UploadFilename,omitnil,omitempty" name:"UploadFilename"`

	// 语言版本
	LanguageVersion *string `json:"LanguageVersion,omitnil,omitempty" name:"LanguageVersion"`
}

type ClsInfo struct {
	// cls所属地域
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`

	// cls日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitnil,omitempty" name:"ClsLogsetId"`

	// cls日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type CreateCloudRunEnvRequestParams struct {
	// Trial,Standard,Professional,Enterprise
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 环境别名，要以a-z开头，不能包含 a-z,0-9,- 以外的字符
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 用户享有的免费额度级别，目前只能为“basic”，不传该字段或该字段为空，标识不享受免费额度。
	FreeQuota *string `json:"FreeQuota,omitnil,omitempty" name:"FreeQuota"`

	// 订单标记。建议使用方统一转大小写之后再判断。
	// QuickStart：快速启动来源
	// Activity：活动来源
	Flag *string `json:"Flag,omitnil,omitempty" name:"Flag"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网列表
	SubNetIds []*string `json:"SubNetIds,omitnil,omitempty" name:"SubNetIds"`

	// 请求key 用于防重
	ReqKey *string `json:"ReqKey,omitnil,omitempty" name:"ReqKey"`

	// 来源：wechat | cloud | weda
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 渠道：wechat | cloud | weda
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 环境ID 云开发平台必填
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type CreateCloudRunEnvRequest struct {
	*tchttp.BaseRequest
	
	// Trial,Standard,Professional,Enterprise
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 环境别名，要以a-z开头，不能包含 a-z,0-9,- 以外的字符
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 用户享有的免费额度级别，目前只能为“basic”，不传该字段或该字段为空，标识不享受免费额度。
	FreeQuota *string `json:"FreeQuota,omitnil,omitempty" name:"FreeQuota"`

	// 订单标记。建议使用方统一转大小写之后再判断。
	// QuickStart：快速启动来源
	// Activity：活动来源
	Flag *string `json:"Flag,omitnil,omitempty" name:"Flag"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网列表
	SubNetIds []*string `json:"SubNetIds,omitnil,omitempty" name:"SubNetIds"`

	// 请求key 用于防重
	ReqKey *string `json:"ReqKey,omitnil,omitempty" name:"ReqKey"`

	// 来源：wechat | cloud | weda
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 渠道：wechat | cloud | weda
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 环境ID 云开发平台必填
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
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
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 后付费订单号
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 部署信息
	DeployInfo *DeployParam `json:"DeployInfo,omitnil,omitempty" name:"DeployInfo"`

	// 服务配置信息(已废弃)
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitnil,omitempty" name:"ServerConfig"`

	// 服务配置信息
	Items []*DiffConfigItem `json:"Items,omitnil,omitempty" name:"Items"`

	// vpc 信息
	VpcInfo *CreateVpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`
}

type CreateCloudRunServerRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 部署信息
	DeployInfo *DeployParam `json:"DeployInfo,omitnil,omitempty" name:"DeployInfo"`

	// 服务配置信息(已废弃)
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitnil,omitempty" name:"ServerConfig"`

	// 服务配置信息
	Items []*DiffConfigItem `json:"Items,omitnil,omitempty" name:"Items"`

	// vpc 信息
	VpcInfo *CreateVpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`
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
	delete(f, "Items")
	delete(f, "VpcInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudRunServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudRunServerResponseParams struct {
	// 一键部署任务Id，微信云托管，暂时用不到
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CreateVpcInfo struct {
	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 1 新建 2 指定
	CreateType *int64 `json:"CreateType,omitnil,omitempty" name:"CreateType"`

	// 子网ID列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

type DatabasesInfo struct {
	// 数据库唯一标识
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 状态。包含以下取值：
	// <li>INITIALIZING：资源初始化中</li>
	// <li>RUNNING：运行中，可正常使用的状态</li>
	// <li>UNUSABLE：禁用，不可用</li>
	// <li>OVERDUE：资源过期</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type DeleteCloudRunServerRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 操作人信息
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
}

type DeleteCloudRunServerRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 操作人信息
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
}

func (r *DeleteCloudRunServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRunServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "OperatorRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudRunServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudRunServerResponseParams struct {
	// 删除结果：success / failed
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudRunServerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudRunServerResponseParams `json:"Response"`
}

func (r *DeleteCloudRunServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRunServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudRunVersionsRequestParams struct {
	// 环境 Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 是否删除服务，只有最后一个版本的时候才生效
	IsDeleteServer *bool `json:"IsDeleteServer,omitnil,omitempty" name:"IsDeleteServer"`

	// 只有删除服务的时候，才生效
	IsDeleteImage *bool `json:"IsDeleteImage,omitnil,omitempty" name:"IsDeleteImage"`

	// 删除版本的信息
	SimpleVersions []*SimpleVersion `json:"SimpleVersions,omitnil,omitempty" name:"SimpleVersions"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
}

type DeleteCloudRunVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 环境 Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 是否删除服务，只有最后一个版本的时候才生效
	IsDeleteServer *bool `json:"IsDeleteServer,omitnil,omitempty" name:"IsDeleteServer"`

	// 只有删除服务的时候，才生效
	IsDeleteImage *bool `json:"IsDeleteImage,omitnil,omitempty" name:"IsDeleteImage"`

	// 删除版本的信息
	SimpleVersions []*SimpleVersion `json:"SimpleVersions,omitnil,omitempty" name:"SimpleVersions"`

	// 操作备注
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
}

func (r *DeleteCloudRunVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRunVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "IsDeleteServer")
	delete(f, "IsDeleteImage")
	delete(f, "SimpleVersions")
	delete(f, "OperatorRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudRunVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudRunVersionsResponseParams struct {
	// succ | fail | partial
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 删除失败的版本列表
	FailVersions []*FailDeleteVersions `json:"FailVersions,omitnil,omitempty" name:"FailVersions"`

	// 删除成功的版本列表
	SuccessVersions []*SuccessDeleteVersions `json:"SuccessVersions,omitnil,omitempty" name:"SuccessVersions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudRunVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudRunVersionsResponseParams `json:"Response"`
}

func (r *DeleteCloudRunVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRunVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployParam struct {
	// 部署类型：package/image/repository/pipeline/jar/war
	DeployType *string `json:"DeployType,omitnil,omitempty" name:"DeployType"`

	// 部署类型为image时传入
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 部署类型为package时传入
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 部署类型为package时传入
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// 部署备注
	DeployRemark *string `json:"DeployRemark,omitnil,omitempty" name:"DeployRemark"`

	// 代码仓库信息
	RepoInfo *RepositoryInfo `json:"RepoInfo,omitnil,omitempty" name:"RepoInfo"`

	// 无Dockerfile时填写
	BuildPacks *BuildPacksInfo `json:"BuildPacks,omitnil,omitempty" name:"BuildPacks"`

	// 发布类型 GRAY | FULL
	ReleaseType *string `json:"ReleaseType,omitnil,omitempty" name:"ReleaseType"`
}

type DeployRecord struct {
	// 部署Id
	DeployId *string `json:"DeployId,omitnil,omitempty" name:"DeployId"`

	// 部署开始时间
	DeployTime *string `json:"DeployTime,omitnil,omitempty" name:"DeployTime"`

	// 状态：running/deploying/deploy_failed
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 部署运行Id 用户查询部署日志
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// 构建Id
	BuildId *int64 `json:"BuildId,omitnil,omitempty" name:"BuildId"`

	// 流量比例
	FlowRatio *int64 `json:"FlowRatio,omitnil,omitempty" name:"FlowRatio"`

	// 镜像url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 缩容状态 缩容为 zero 否则为空
	ScaleStatus *string `json:"ScaleStatus,omitnil,omitempty" name:"ScaleStatus"`

	// 是否分配流量
	HasTraffic *bool `json:"HasTraffic,omitnil,omitempty" name:"HasTraffic"`

	// 流量分配方式, FLOW: 百分比分配; URL_PARAMS: 匹配 query 参数; HEADERS: 匹配请求 Header
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// 当前版本是否在发布中
	IsReleasing *bool `json:"IsReleasing,omitnil,omitempty" name:"IsReleasing"`
}

// Predefined struct for user
type DescribeCloudRunDeployRecordRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`
}

type DescribeCloudRunDeployRecordRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`
}

func (r *DescribeCloudRunDeployRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunDeployRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRunDeployRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunDeployRecordResponseParams struct {
	// 部署列表
	DeployRecords []*DeployRecord `json:"DeployRecords,omitnil,omitempty" name:"DeployRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudRunDeployRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudRunDeployRecordResponseParams `json:"Response"`
}

func (r *DescribeCloudRunDeployRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunDeployRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunEnvsRequestParams struct {
	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitnil,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitnil,omitempty" name:"Channels"`
}

type DescribeCloudRunEnvsRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 指定Channels字段为可见渠道列表或不可见渠道列表
	// 如只想获取渠道A的环境 就填写IsVisible= true,Channels = ["A"], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = ["A"]
	IsVisible *bool `json:"IsVisible,omitnil,omitempty" name:"IsVisible"`

	// 渠道列表，代表可见或不可见渠道由IsVisible参数指定
	Channels []*string `json:"Channels,omitnil,omitempty" name:"Channels"`
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
	EnvList []*EnvInfo `json:"EnvList,omitnil,omitempty" name:"EnvList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCloudRunPodListRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 默认为10， 最大为50
	// 不传或传0时 取默认10
	// 大于50时取50
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 不传或传0时 会默认为1
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`
}

type DescribeCloudRunPodListRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 默认为10， 最大为50
	// 不传或传0时 取默认10
	// 大于50时取50
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 不传或传0时 会默认为1
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`
}

func (r *DescribeCloudRunPodListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunPodListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "VersionName")
	delete(f, "PageSize")
	delete(f, "PageNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRunPodListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunPodListResponseParams struct {
	// pod实例列表
	PodList []*VersionPodInstance `json:"PodList,omitnil,omitempty" name:"PodList"`

	// pod总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudRunPodListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudRunPodListResponseParams `json:"Response"`
}

func (r *DescribeCloudRunPodListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunPodListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunProcessLogRequestParams struct {
	// 环境 Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 操作 Id
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`
}

type DescribeCloudRunProcessLogRequest struct {
	*tchttp.BaseRequest
	
	// 环境 Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 操作 Id
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`
}

func (r *DescribeCloudRunProcessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunProcessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "RunId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRunProcessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunProcessLogResponseParams struct {
	// 运行日志列表
	Logs []*string `json:"Logs,omitnil,omitempty" name:"Logs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudRunProcessLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudRunProcessLogResponseParams `json:"Response"`
}

func (r *DescribeCloudRunProcessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRunProcessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunServerDetailRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`
}

type DescribeCloudRunServerDetailRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`
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
	BaseInfo *ServerBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// 服务配置信息
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitnil,omitempty" name:"ServerConfig"`

	// 在线版本信息
	OnlineVersionInfos []*OnlineVersionInfo `json:"OnlineVersionInfos,omitnil,omitempty" name:"OnlineVersionInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 默认为9， 最大为30
	// 不传或传0时 取默认9
	// 大于30时取30
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 不传或传0时 会默认为1
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 服务类型：function | container
	ServerType *string `json:"ServerType,omitnil,omitempty" name:"ServerType"`

	// vpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type DescribeCloudRunServersRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 默认为9， 最大为30
	// 不传或传0时 取默认9
	// 大于30时取30
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 不传或传0时 会默认为1
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 服务类型：function | container
	ServerType *string `json:"ServerType,omitnil,omitempty" name:"ServerType"`

	// vpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
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
	delete(f, "ServerName")
	delete(f, "ServerType")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRunServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRunServersResponseParams struct {
	// 服务列表
	ServerList []*ServerBaseInfo `json:"ServerList,omitnil,omitempty" name:"ServerList"`

	// 服务总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeEnvBaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 环境 Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
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
	EnvBaseInfo *EnvBaseInfo `json:"EnvBaseInfo,omitnil,omitempty" name:"EnvBaseInfo"`

	// 是否存在
	IsExist *bool `json:"IsExist,omitnil,omitempty" name:"IsExist"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeReleaseOrderRequestParams struct {
	// 环境 Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 发布单状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeReleaseOrderRequest struct {
	*tchttp.BaseRequest
	
	// 环境 Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 发布单状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeReleaseOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseOrderResponseParams struct {
	// 是否存在
	IsExist *bool `json:"IsExist,omitnil,omitempty" name:"IsExist"`

	// 发布单信息
	ReleaseOrderInfo *ReleaseOrderInfo `json:"ReleaseOrderInfo,omitnil,omitempty" name:"ReleaseOrderInfo"`

	// 上一次成功发布时间
	LastReleasedSuccessTime *string `json:"LastReleasedSuccessTime,omitnil,omitempty" name:"LastReleasedSuccessTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseOrderResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseOrderResponseParams `json:"Response"`
}

func (r *DescribeReleaseOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerManageTaskRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 任务Id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
}

type DescribeServerManageTaskRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 任务Id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
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
	IsExist *bool `json:"IsExist,omitnil,omitempty" name:"IsExist"`

	// 任务信息
	Task *ServerManageTaskInfo `json:"Task,omitnil,omitempty" name:"Task"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeVersionDetailRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// channel
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`
}

type DescribeVersionDetailRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// channel
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`
}

func (r *DescribeVersionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVersionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "VersionName")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVersionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVersionDetailResponseParams struct {
	// 版本名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 端口号
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// cpu 规格
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// mem 规格
	Mem *float64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 最小副本数
	MinNum *int64 `json:"MinNum,omitnil,omitempty" name:"MinNum"`

	// 最大副本数
	MaxNum *int64 `json:"MaxNum,omitnil,omitempty" name:"MaxNum"`

	// 扩缩容策略
	PolicyDetails []*HpaPolicy `json:"PolicyDetails,omitnil,omitempty" name:"PolicyDetails"`

	// Dockerfile path
	Dockerfile *string `json:"Dockerfile,omitnil,omitempty" name:"Dockerfile"`

	// 目标目录
	BuildDir *string `json:"BuildDir,omitnil,omitempty" name:"BuildDir"`

	// 环境变量
	EnvParams *string `json:"EnvParams,omitnil,omitempty" name:"EnvParams"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 日志采集路径
	LogPath *string `json:"LogPath,omitnil,omitempty" name:"LogPath"`

	// entryPoint
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntryPoint *string `json:"EntryPoint,omitnil,omitempty" name:"EntryPoint"`

	// Cmd
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// vpc conf
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcConf *VpcConf `json:"VpcConf,omitnil,omitempty" name:"VpcConf"`

	// volume conf
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumesConf []*VolumeConf `json:"VolumesConf,omitnil,omitempty" name:"VolumesConf"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVersionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVersionDetailResponseParams `json:"Response"`
}

func (r *DescribeVersionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVersionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiffConfigItem struct {
	// 配置项 Key
	// MinNum 最小副本数
	// MaxNum 最大副本数
	// PolicyDetails 扩缩容策略
	// AccessTypes 访问类型
	// TimerScale 定时扩缩容
	// InternalAccess 内网访问
	// OperationMode 运行模式 noScale | condScale | alwaysScale | custom ｜ manualScale
	// SessionAffinity 会话亲和性 open | close
	// CpuSpecs cpu 规格
	// MemSpecs mem规格
	// EnvParam 环境变量
	// LogPath 日志采集路径
	// Port 端口
	// Dockerfile dockerfile 文件名
	// BuildDir 目标目录
	// Tag 服务标签
	// LogType 日志类型 none | default | custom 
	// LogSetId 日志集Id
	// LogTopicId 日志主题ID
	// LogParseType 日志解析类型 json ｜ line
	// EntryPoint entrypoint 命令
	// Cmd cmd命令
	// VpcConf 网络信息
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 字符串类型配置项值
	// InternalAccess、OperationMode、SessionAffinity、EnvParam、LogPath、Dockerfile、BuildDir、Tag、LogType、LogSetId、LogTopicId、LogParseType
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// int 类型配置项值
	// MinNum、MaxNum、Port
	IntValue *int64 `json:"IntValue,omitnil,omitempty" name:"IntValue"`

	// bool 类型配置项值
	BoolValue *bool `json:"BoolValue,omitnil,omitempty" name:"BoolValue"`

	// 浮点型配置项值
	// CpuSpecs、MemSpecs
	FloatValue *float64 `json:"FloatValue,omitnil,omitempty" name:"FloatValue"`

	// 字符串数组配置项值
	// AccessTypes，EntryPoint，Cmd
	ArrayValue []*string `json:"ArrayValue,omitnil,omitempty" name:"ArrayValue"`

	// 扩缩容策略配置项值
	PolicyDetails []*HpaPolicy `json:"PolicyDetails,omitnil,omitempty" name:"PolicyDetails"`

	// 定时扩缩容配置项值
	TimerScale []*TimerScale `json:"TimerScale,omitnil,omitempty" name:"TimerScale"`

	// 配置内网访问时网络信息
	VpcConf *VpcConf `json:"VpcConf,omitnil,omitempty" name:"VpcConf"`

	// 存储配置信息
	VolumesConf []*VolumeConf `json:"VolumesConf,omitnil,omitempty" name:"VolumesConf"`
}

type EnvBaseInfo struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 套餐类型：Trial ｜ Standard ｜ Professional ｜ Enterprise
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// VPC Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 环境创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 环境别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 环境状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 环境地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 环境类型 tcbr ｜ run
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 子网id
	SubnetIds *string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 回收标志，为空则表示正常，recycle表示已回收
	Recycle *string `json:"Recycle,omitnil,omitempty" name:"Recycle"`
}

type EnvInfo struct {
	// 账户下该环境唯一标识
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 环境来源。包含以下取值：
	// <li>miniapp：微信小程序</li>
	// <li>qcloud ：腾讯云</li>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 环境别名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 环境状态。包含以下取值：
	// <li>NORMAL：正常可用</li>
	// <li>UNAVAILABLE：服务不可用，可能是尚未初始化或者初始化过程中</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否到期自动降为免费版
	IsAutoDegrade *bool `json:"IsAutoDegrade,omitnil,omitempty" name:"IsAutoDegrade"`

	// 环境渠道
	EnvChannel *string `json:"EnvChannel,omitnil,omitempty" name:"EnvChannel"`

	// 支付方式。包含以下取值：
	// <li> prepayment：预付费</li>
	// <li> postpaid：后付费</li>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否为默认环境
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 环境所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 环境类型：baas, run, hosting, weda,tcbr
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 数据库列表
	Databases []*DatabasesInfo `json:"Databases,omitnil,omitempty" name:"Databases"`

	// 存储列表
	Storages []*StorageInfo `json:"Storages,omitnil,omitempty" name:"Storages"`

	// 函数列表
	Functions []*FunctionInfo `json:"Functions,omitnil,omitempty" name:"Functions"`

	// 云日志服务列表
	LogServices []*LogServiceInfo `json:"LogServices,omitnil,omitempty" name:"LogServices"`

	// 静态资源信息
	StaticStorages []*StaticStorageInfo `json:"StaticStorages,omitnil,omitempty" name:"StaticStorages"`

	// 环境标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自定义日志服务
	CustomLogServices []*ClsInfo `json:"CustomLogServices,omitnil,omitempty" name:"CustomLogServices"`

	// tcb产品套餐ID，参考DescribePackages接口的返回值。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 套餐中文名称，参考DescribePackages接口的返回值。
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type FailDeleteVersions struct {
	// 删除失败版本信息
	Version *SimpleVersion `json:"Version,omitnil,omitempty" name:"Version"`

	// 删除失败错误码
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 删除失败错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 删除操作 RequestId
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FunctionInfo struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type HpaPolicy struct {
	// 扩缩容类型
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 扩缩容阈值
	PolicyThreshold *uint64 `json:"PolicyThreshold,omitnil,omitempty" name:"PolicyThreshold"`
}

type LogObject struct {
	// 日志属于的 topic ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 采集路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 日志来源设备
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 日志唯一标识
	PkgLogId *string `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`
}

type LogResObject struct {
	// 获取更多检索结果的游标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 搜索结果是否已经全部返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 日志内容信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*LogObject `json:"Results,omitnil,omitempty" name:"Results"`
}

type LogServiceInfo struct {
	// log名
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// log-id
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// topic名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// topic-id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// cls日志所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type ObjectKV struct {
	// 键值对Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 键值对Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ObjectKVPriority struct {
	// 键值对Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 键值对Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 键值对权重
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type OnlineVersionInfo struct {
	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 镜像url
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 流量
	FlowRatio *string `json:"FlowRatio,omitnil,omitempty" name:"FlowRatio"`
}

// Predefined struct for user
type OperateServerManageRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 任报Id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作类型:cancel | go_back | done
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
}

type OperateServerManageRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 任报Id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作类型:cancel | go_back | done
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 灰度类型
	GrayType *string `json:"GrayType,omitnil,omitempty" name:"GrayType"`

	// 流量类型
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// 流量策略
	VersionFlowItems []*VersionFlowInfo `json:"VersionFlowItems,omitnil,omitempty" name:"VersionFlowItems"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`

	// 流量比例
	GrayFlowRatio *int64 `json:"GrayFlowRatio,omitnil,omitempty" name:"GrayFlowRatio"`
}

type ReleaseGrayRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 灰度类型
	GrayType *string `json:"GrayType,omitnil,omitempty" name:"GrayType"`

	// 流量类型
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// 流量策略
	VersionFlowItems []*VersionFlowInfo `json:"VersionFlowItems,omitnil,omitempty" name:"VersionFlowItems"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`

	// 流量比例
	GrayFlowRatio *int64 `json:"GrayFlowRatio,omitnil,omitempty" name:"GrayFlowRatio"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ReleaseOrderInfo struct {
	// 发布单Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 当前版本
	CurrentVersion *VersionInfo `json:"CurrentVersion,omitnil,omitempty" name:"CurrentVersion"`

	// 发布版本
	ReleaseVersion *VersionInfo `json:"ReleaseVersion,omitnil,omitempty" name:"ReleaseVersion"`

	// 灰度状态
	GrayStatus *string `json:"GrayStatus,omitnil,omitempty" name:"GrayStatus"`

	// 发布状态
	ReleaseStatus *string `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// 流量值
	TrafficTypeValues []*ObjectKV `json:"TrafficTypeValues,omitnil,omitempty" name:"TrafficTypeValues"`

	// 流量类型
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// 百分比
	FlowRatio *int64 `json:"FlowRatio,omitnil,omitempty" name:"FlowRatio"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否发布中
	IsReleasing *bool `json:"IsReleasing,omitnil,omitempty" name:"IsReleasing"`
}

type RepositoryInfo struct {
	// git source
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 仓库名
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 分支名
	Branch *string `json:"Branch,omitnil,omitempty" name:"Branch"`
}

// Predefined struct for user
type SearchClsLogRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询语句，详情参考 https://cloud.tencent.com/document/product/614/47044
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 是否使用Lucene语法，默认为false
	UseLucene *bool `json:"UseLucene,omitnil,omitempty" name:"UseLucene"`

	// 日志类型
	LogType *int64 `json:"LogType,omitnil,omitempty" name:"LogType"`
}

type SearchClsLogRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询语句，详情参考 https://cloud.tencent.com/document/product/614/47044
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 是否使用Lucene语法，默认为false
	UseLucene *bool `json:"UseLucene,omitnil,omitempty" name:"UseLucene"`

	// 日志类型
	LogType *int64 `json:"LogType,omitnil,omitempty" name:"LogType"`
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
	delete(f, "LogType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClsLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClsLogResponseParams struct {
	// 日志内容结果
	LogResults *LogResObject `json:"LogResults,omitnil,omitempty" name:"LogResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ServerBaseConfig struct {
	// 环境 Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 是否开启公网访问
	OpenAccessTypes []*string `json:"OpenAccessTypes,omitnil,omitempty" name:"OpenAccessTypes"`

	// Cpu 规格
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Mem 规格
	Mem *float64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 最小副本数
	MinNum *uint64 `json:"MinNum,omitnil,omitempty" name:"MinNum"`

	// 最大副本数
	MaxNum *uint64 `json:"MaxNum,omitnil,omitempty" name:"MaxNum"`

	// 扩缩容配置
	PolicyDetails []*HpaPolicy `json:"PolicyDetails,omitnil,omitempty" name:"PolicyDetails"`

	// 日志采集路径
	CustomLogs *string `json:"CustomLogs,omitnil,omitempty" name:"CustomLogs"`

	// 环境变量
	EnvParams *string `json:"EnvParams,omitnil,omitempty" name:"EnvParams"`

	// 延迟检测时间
	InitialDelaySeconds *uint64 `json:"InitialDelaySeconds,omitnil,omitempty" name:"InitialDelaySeconds"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 服务端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 是否有Dockerfile
	HasDockerfile *bool `json:"HasDockerfile,omitnil,omitempty" name:"HasDockerfile"`

	// Dockerfile 文件名
	Dockerfile *string `json:"Dockerfile,omitnil,omitempty" name:"Dockerfile"`

	// 构建目录
	BuildDir *string `json:"BuildDir,omitnil,omitempty" name:"BuildDir"`

	// 日志类型: none | default | custom
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// cls setId
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// cls 主题id
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// 解析类型：json ｜ line
	LogParseType *string `json:"LogParseType,omitnil,omitempty" name:"LogParseType"`

	// 服务标签, function: 函数托管
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 内网访问开关 close | open
	InternalAccess *string `json:"InternalAccess,omitnil,omitempty" name:"InternalAccess"`

	// 内网域名
	InternalDomain *string `json:"InternalDomain,omitnil,omitempty" name:"InternalDomain"`

	// 运行模式
	OperationMode *string `json:"OperationMode,omitnil,omitempty" name:"OperationMode"`

	// 定时扩缩容配置
	TimerScale []*TimerScale `json:"TimerScale,omitnil,omitempty" name:"TimerScale"`

	// Dockerfile EntryPoint 参数
	EntryPoint []*string `json:"EntryPoint,omitnil,omitempty" name:"EntryPoint"`

	// Dockerfile Cmd 参数
	Cmd []*string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 会话亲和性开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionAffinity *string `json:"SessionAffinity,omitnil,omitempty" name:"SessionAffinity"`

	// Vpc 配置参数
	VpcConf *VpcConf `json:"VpcConf,omitnil,omitempty" name:"VpcConf"`

	// 存储配置信息
	VolumesConf []*VolumeConf `json:"VolumesConf,omitnil,omitempty" name:"VolumesConf"`

	// 关联镜像密钥
	LinkImageRegistry *string `json:"LinkImageRegistry,omitnil,omitempty" name:"LinkImageRegistry"`
}

type ServerBaseInfo struct {
	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 默认服务域名
	DefaultDomainName *string `json:"DefaultDomainName,omitnil,omitempty" name:"DefaultDomainName"`

	// 自定义域名
	CustomDomainName *string `json:"CustomDomainName,omitnil,omitempty" name:"CustomDomainName"`

	// 服务状态：running/deploying/deploy_failed
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 公网访问类型
	AccessTypes []*string `json:"AccessTypes,omitnil,omitempty" name:"AccessTypes"`

	// 展示自定义域名
	CustomDomainNames []*string `json:"CustomDomainNames,omitnil,omitempty" name:"CustomDomainNames"`

	// 服务类型: function 云函数2.0；container 容器服务
	ServerType *string `json:"ServerType,omitnil,omitempty" name:"ServerType"`

	// 流量类型，目前只有 FLOW
	TrafficType *string `json:"TrafficType,omitnil,omitempty" name:"TrafficType"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ServerManageTaskInfo struct {
	// 任务Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 变更类型
	ChangeType *string `json:"ChangeType,omitnil,omitempty" name:"ChangeType"`

	// 发布类型
	ReleaseType *string `json:"ReleaseType,omitnil,omitempty" name:"ReleaseType"`

	// 部署类型
	DeployType *string `json:"DeployType,omitnil,omitempty" name:"DeployType"`

	// 上一个版本名
	PreVersionName *string `json:"PreVersionName,omitnil,omitempty" name:"PreVersionName"`

	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 流水线Id
	PipelineId *int64 `json:"PipelineId,omitnil,omitempty" name:"PipelineId"`

	// 流水线任务Id
	PipelineTaskId *int64 `json:"PipelineTaskId,omitnil,omitempty" name:"PipelineTaskId"`

	// 发布单Id
	ReleaseId *int64 `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 步骤信息
	Steps []*TaskStepInfo `json:"Steps,omitnil,omitempty" name:"Steps"`

	// 失败原因
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
}

type SimpleVersion struct {
	// 要删除版本的环境 Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 要删除版本的服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 要删除版本的版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`
}

type StaticStorageInfo struct {
	// 静态CDN域名
	StaticDomain *string `json:"StaticDomain,omitnil,omitempty" name:"StaticDomain"`

	// 静态CDN默认文件夹，当前为根目录
	DefaultDirName *string `json:"DefaultDirName,omitnil,omitempty" name:"DefaultDirName"`

	// 资源状态(process/online/offline/init)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// cos所属区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// bucket信息
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

type StorageInfo struct {
	// 资源所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 桶名，存储资源的唯一标识
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// cdn 域名
	CdnDomain *string `json:"CdnDomain,omitnil,omitempty" name:"CdnDomain"`

	// 资源所属用户的腾讯云appId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

// Predefined struct for user
type SubmitServerRollbackRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 当前版本
	CurrentVersionName *string `json:"CurrentVersionName,omitnil,omitempty" name:"CurrentVersionName"`

	// 回滚版本
	RollbackVersionName *string `json:"RollbackVersionName,omitnil,omitempty" name:"RollbackVersionName"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
}

type SubmitServerRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 当前版本
	CurrentVersionName *string `json:"CurrentVersionName,omitnil,omitempty" name:"CurrentVersionName"`

	// 回滚版本
	RollbackVersionName *string `json:"RollbackVersionName,omitnil,omitempty" name:"RollbackVersionName"`

	// 操作标识
	OperatorRemark *string `json:"OperatorRemark,omitnil,omitempty" name:"OperatorRemark"`
}

func (r *SubmitServerRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitServerRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "CurrentVersionName")
	delete(f, "RollbackVersionName")
	delete(f, "OperatorRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitServerRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitServerRollbackResponseParams struct {
	// 任务Id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitServerRollbackResponse struct {
	*tchttp.BaseResponse
	Response *SubmitServerRollbackResponseParams `json:"Response"`
}

func (r *SubmitServerRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitServerRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuccessDeleteVersions struct {
	// 版本简化信息
	Version *SimpleVersion `json:"Version,omitnil,omitempty" name:"Version"`

	// 删除版本的 RequestId
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 删除版本结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TaskStepInfo struct {
	// 步骤名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 未启动："todo"
	// 运行中："running"
	// 失败："failed"
	// 成功结束："finished"
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 消耗时间：秒
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 失败原因
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`
}

type TimerScale struct {
	// 循环类型
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 循环起始
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 循环结束
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 副本个数
	ReplicaNum *uint64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`
}

// Predefined struct for user
type UpdateCloudRunServerRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 部署信息
	DeployInfo *DeployParam `json:"DeployInfo,omitnil,omitempty" name:"DeployInfo"`

	// 服务配置信息(已废弃)
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitnil,omitempty" name:"ServerConfig"`

	// 业务类型，默认tcr
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// 服务配置信息
	Items []*DiffConfigItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type UpdateCloudRunServerRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 部署信息
	DeployInfo *DeployParam `json:"DeployInfo,omitnil,omitempty" name:"DeployInfo"`

	// 服务配置信息(已废弃)
	ServerConfig *ServerBaseConfig `json:"ServerConfig,omitnil,omitempty" name:"ServerConfig"`

	// 业务类型，默认tcr
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// 服务配置信息
	Items []*DiffConfigItem `json:"Items,omitnil,omitempty" name:"Items"`
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
	delete(f, "Business")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCloudRunServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCloudRunServerResponseParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 一键部署任务Id，暂时用不到
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 是否默认版本
	IsDefaultPriority *bool `json:"IsDefaultPriority,omitnil,omitempty" name:"IsDefaultPriority"`

	// 流量比例
	FlowRatio *int64 `json:"FlowRatio,omitnil,omitempty" name:"FlowRatio"`

	// 测试KV值
	UrlParam *ObjectKV `json:"UrlParam,omitnil,omitempty" name:"UrlParam"`

	// 权重
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type VersionInfo struct {
	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 流量比例
	FlowRatio *int64 `json:"FlowRatio,omitnil,omitempty" name:"FlowRatio"`

	// 版本状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 构建Id
	BuildId *int64 `json:"BuildId,omitnil,omitempty" name:"BuildId"`

	// 上传方式
	UploadType *string `json:"UploadType,omitnil,omitempty" name:"UploadType"`

	// 操作标识
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 测试参数
	UrlParam *ObjectKV `json:"UrlParam,omitnil,omitempty" name:"UrlParam"`

	// 权重
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 是否默认
	IsDefaultPriority *bool `json:"IsDefaultPriority,omitnil,omitempty" name:"IsDefaultPriority"`

	// 流量参数
	FlowParams []*ObjectKVPriority `json:"FlowParams,omitnil,omitempty" name:"FlowParams"`

	// 最小副本数
	MinReplicas *int64 `json:"MinReplicas,omitnil,omitempty" name:"MinReplicas"`

	// 最大副本数
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 操作Id
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// 百分比
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 当前副本数
	CurrentReplicas *int64 `json:"CurrentReplicas,omitnil,omitempty" name:"CurrentReplicas"`

	// 架构类型
	Architecture *string `json:"Architecture,omitnil,omitempty" name:"Architecture"`
}

type VersionPodInstance struct {
	// 实例webshell链接
	Webshell *string `json:"Webshell,omitnil,omitempty" name:"Webshell"`

	// 实例Id
	PodId *string `json:"PodId,omitnil,omitempty" name:"PodId"`

	// 实例状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type VolumeConf struct {
	// 存储类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 对象存储桶名称
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 存储连接地址
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 存储连接用户密码
	KeyID *string `json:"KeyID,omitnil,omitempty" name:"KeyID"`

	// 存储挂载目的目录
	DstPath *string `json:"DstPath,omitnil,omitempty" name:"DstPath"`

	// 存储挂载源目录
	SrcPath *string `json:"SrcPath,omitnil,omitempty" name:"SrcPath"`
}

type VpcConf struct {
	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc 网段
	VpcCIDR *string `json:"VpcCIDR,omitnil,omitempty" name:"VpcCIDR"`

	// subnet id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// subnet 网段
	SubnetCIDR *string `json:"SubnetCIDR,omitnil,omitempty" name:"SubnetCIDR"`
}