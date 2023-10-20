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

package v20220105

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ApplyEmbedIntervalRequestParams struct {
	// 分享项目id，必选
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0
	PageId *uint64 `json:"PageId,omitnil" name:"PageId"`

	// 需要申请延期的Token
	BIToken *string `json:"BIToken,omitnil" name:"BIToken"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// panel,看板；page，页面
	Scope *string `json:"Scope,omitnil" name:"Scope"`
}

type ApplyEmbedIntervalRequest struct {
	*tchttp.BaseRequest
	
	// 分享项目id，必选
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0
	PageId *uint64 `json:"PageId,omitnil" name:"PageId"`

	// 需要申请延期的Token
	BIToken *string `json:"BIToken,omitnil" name:"BIToken"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// panel,看板；page，页面
	Scope *string `json:"Scope,omitnil" name:"Scope"`
}

func (r *ApplyEmbedIntervalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyEmbedIntervalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	delete(f, "BIToken")
	delete(f, "ExtraParam")
	delete(f, "Scope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyEmbedIntervalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyEmbedIntervalResponseParams struct {
	// 额外参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 结果数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ApplyEmbedTokenInfo `json:"Data,omitnil" name:"Data"`

	// 结果描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ApplyEmbedIntervalResponse struct {
	*tchttp.BaseResponse
	Response *ApplyEmbedIntervalResponseParams `json:"Response"`
}

func (r *ApplyEmbedIntervalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyEmbedIntervalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyEmbedTokenInfo struct {
	// 申请结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil" name:"Result"`
}

type BaseStateAction struct {
	// 编辑是否可见
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowEdit *bool `json:"ShowEdit,omitnil" name:"ShowEdit"`

	// 编辑是否可点击
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsEdit *bool `json:"IsEdit,omitnil" name:"IsEdit"`

	// 编辑按钮的文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditText *string `json:"EditText,omitnil" name:"EditText"`

	// 编辑不可用时的提示文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditTips *string `json:"EditTips,omitnil" name:"EditTips"`

	// 删除是否可见
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowDel *bool `json:"ShowDel,omitnil" name:"ShowDel"`

	// 删除是否可点击
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDel *bool `json:"IsDel,omitnil" name:"IsDel"`

	// 删除按钮的文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelText *string `json:"DelText,omitnil" name:"DelText"`

	// 删除不可用时的提示文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelTips *string `json:"DelTips,omitnil" name:"DelTips"`

	// 复制是否可见
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowCopy *bool `json:"ShowCopy,omitnil" name:"ShowCopy"`

	// 是否可见
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowView *bool `json:"ShowView,omitnil" name:"ShowView"`

	// 是否可重命名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowRename *bool `json:"ShowRename,omitnil" name:"ShowRename"`
}

type CorpUserListData struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*UserIdAndUserName `json:"List,omitnil" name:"List"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 页数
	TotalPages *int64 `json:"TotalPages,omitnil" name:"TotalPages"`
}

// Predefined struct for user
type CreateDatasourceCloudRequestParams struct {
	// 后端提供字典：域类型，1、腾讯云，2、本地
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 驱动
	DbType *string `json:"DbType,omitnil" name:"DbType"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil" name:"Charset"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil" name:"DbUser"`

	// 密码
	DbPwd *string `json:"DbPwd,omitnil" name:"DbPwd"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 公有云内网ip
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 公有云内网端口
	Vport *string `json:"Vport,omitnil" name:"Vport"`

	// vpc标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 统一vpc标识
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 区域标识（gz,bj)
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 扩展参数
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 数据源产品名
	ProdDbName *string `json:"ProdDbName,omitnil" name:"ProdDbName"`

	// 第三方数据源标识
	DataOrigin *string `json:"DataOrigin,omitnil" name:"DataOrigin"`

	// 第三方项目id
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil" name:"DataOriginProjectId"`

	// 第三方数据源id
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil" name:"DataOriginDatasourceId"`
}

type CreateDatasourceCloudRequest struct {
	*tchttp.BaseRequest
	
	// 后端提供字典：域类型，1、腾讯云，2、本地
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 驱动
	DbType *string `json:"DbType,omitnil" name:"DbType"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil" name:"Charset"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil" name:"DbUser"`

	// 密码
	DbPwd *string `json:"DbPwd,omitnil" name:"DbPwd"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 公有云内网ip
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 公有云内网端口
	Vport *string `json:"Vport,omitnil" name:"Vport"`

	// vpc标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 统一vpc标识
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 区域标识（gz,bj)
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 扩展参数
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 数据源产品名
	ProdDbName *string `json:"ProdDbName,omitnil" name:"ProdDbName"`

	// 第三方数据源标识
	DataOrigin *string `json:"DataOrigin,omitnil" name:"DataOrigin"`

	// 第三方项目id
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil" name:"DataOriginProjectId"`

	// 第三方数据源id
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil" name:"DataOriginDatasourceId"`
}

func (r *CreateDatasourceCloudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasourceCloudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "DbType")
	delete(f, "Charset")
	delete(f, "DbUser")
	delete(f, "DbPwd")
	delete(f, "DbName")
	delete(f, "SourceName")
	delete(f, "ProjectId")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "VpcId")
	delete(f, "UniqVpcId")
	delete(f, "RegionId")
	delete(f, "ExtraParam")
	delete(f, "InstanceId")
	delete(f, "ProdDbName")
	delete(f, "DataOrigin")
	delete(f, "DataOriginProjectId")
	delete(f, "DataOriginDatasourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatasourceCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceCloudResponseParams struct {
	// 成功无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *IdDTO `json:"Data,omitnil" name:"Data"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDatasourceCloudResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatasourceCloudResponseParams `json:"Response"`
}

func (r *CreateDatasourceCloudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasourceCloudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceRequestParams struct {
	// HOST
	DbHost *string `json:"DbHost,omitnil" name:"DbHost"`

	// 端口
	DbPort *uint64 `json:"DbPort,omitnil" name:"DbPort"`

	// 后端提供字典：域类型，1、腾讯云，2、本地
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 驱动
	DbType *string `json:"DbType,omitnil" name:"DbType"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil" name:"Charset"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil" name:"DbUser"`

	// 密码
	DbPwd *string `json:"DbPwd,omitnil" name:"DbPwd"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 项目id
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// catalog值
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 第三方数据源标识
	DataOrigin *string `json:"DataOrigin,omitnil" name:"DataOrigin"`

	// 第三方项目id
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil" name:"DataOriginProjectId"`

	// 第三方数据源id
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil" name:"DataOriginDatasourceId"`

	// 扩展参数
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 腾讯云私有网络统一标识
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 私有网络ip
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 私有网络端口
	Vport *string `json:"Vport,omitnil" name:"Vport"`

	// 腾讯云私有网络标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type CreateDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// HOST
	DbHost *string `json:"DbHost,omitnil" name:"DbHost"`

	// 端口
	DbPort *uint64 `json:"DbPort,omitnil" name:"DbPort"`

	// 后端提供字典：域类型，1、腾讯云，2、本地
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 驱动
	DbType *string `json:"DbType,omitnil" name:"DbType"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil" name:"Charset"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil" name:"DbUser"`

	// 密码
	DbPwd *string `json:"DbPwd,omitnil" name:"DbPwd"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 项目id
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// catalog值
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 第三方数据源标识
	DataOrigin *string `json:"DataOrigin,omitnil" name:"DataOrigin"`

	// 第三方项目id
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil" name:"DataOriginProjectId"`

	// 第三方数据源id
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil" name:"DataOriginDatasourceId"`

	// 扩展参数
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 腾讯云私有网络统一标识
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 私有网络ip
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 私有网络端口
	Vport *string `json:"Vport,omitnil" name:"Vport"`

	// 腾讯云私有网络标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

func (r *CreateDatasourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbHost")
	delete(f, "DbPort")
	delete(f, "ServiceType")
	delete(f, "DbType")
	delete(f, "Charset")
	delete(f, "DbUser")
	delete(f, "DbPwd")
	delete(f, "DbName")
	delete(f, "SourceName")
	delete(f, "ProjectId")
	delete(f, "Catalog")
	delete(f, "DataOrigin")
	delete(f, "DataOriginProjectId")
	delete(f, "DataOriginDatasourceId")
	delete(f, "ExtraParam")
	delete(f, "UniqVpcId")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceResponseParams struct {
	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *IdDTO `json:"Data,omitnil" name:"Data"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDatasourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatasourceResponseParams `json:"Response"`
}

func (r *CreateDatasourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedTokenRequestParams struct {
	// 分享项目id
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0
	PageId *uint64 `json:"PageId,omitnil" name:"PageId"`

	// page表示嵌出页面，panel表嵌出整个看板
	Scope *string `json:"Scope,omitnil" name:"Scope"`

	// 过期时间。 单位：分钟 最大值：240。即，4小时 默认值：240
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 使用者企业Id(仅用于多用户)
	UserCorpId *string `json:"UserCorpId,omitnil" name:"UserCorpId"`

	// 使用者Id(仅用于多用户)
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type CreateEmbedTokenRequest struct {
	*tchttp.BaseRequest
	
	// 分享项目id
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 分享页面id，嵌出看板时此为空值0
	PageId *uint64 `json:"PageId,omitnil" name:"PageId"`

	// page表示嵌出页面，panel表嵌出整个看板
	Scope *string `json:"Scope,omitnil" name:"Scope"`

	// 过期时间。 单位：分钟 最大值：240。即，4小时 默认值：240
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 备用字段
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 使用者企业Id(仅用于多用户)
	UserCorpId *string `json:"UserCorpId,omitnil" name:"UserCorpId"`

	// 使用者Id(仅用于多用户)
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

func (r *CreateEmbedTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	delete(f, "Scope")
	delete(f, "ExpireTime")
	delete(f, "ExtraParam")
	delete(f, "UserCorpId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmbedTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedTokenResponseParams struct {
	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *EmbedTokenInfo `json:"Data,omitnil" name:"Data"`

	// 结果描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateEmbedTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmbedTokenResponseParams `json:"Response"`
}

func (r *CreateEmbedTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// 项目名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// logo底色
	ColorCode *string `json:"ColorCode,omitnil" name:"ColorCode"`

	// 项目Logo
	Logo *string `json:"Logo,omitnil" name:"Logo"`

	// 备注
	Mark *string `json:"Mark,omitnil" name:"Mark"`

	// 是否允许用户申请
	IsApply *bool `json:"IsApply,omitnil" name:"IsApply"`

	// 默认看板
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil" name:"DefaultPanelType"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// logo底色
	ColorCode *string `json:"ColorCode,omitnil" name:"ColorCode"`

	// 项目Logo
	Logo *string `json:"Logo,omitnil" name:"Logo"`

	// 备注
	Mark *string `json:"Mark,omitnil" name:"Mark"`

	// 是否允许用户申请
	IsApply *bool `json:"IsApply,omitnil" name:"IsApply"`

	// 默认看板
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil" name:"DefaultPanelType"`
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
	delete(f, "Name")
	delete(f, "ColorCode")
	delete(f, "Logo")
	delete(f, "Mark")
	delete(f, "IsApply")
	delete(f, "DefaultPanelType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// 额外数据
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 数据
	Data *Data `json:"Data,omitnil" name:"Data"`

	// 返回信息
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type CreateUserRoleProjectRequestParams struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 角色ID列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil" name:"RoleIdList"`

	// 用户列表（废弃）
	UserList []*UserIdAndUserName `json:"UserList,omitnil" name:"UserList"`

	// 用户列表（新）
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil" name:"UserInfoList"`
}

type CreateUserRoleProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 角色ID列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil" name:"RoleIdList"`

	// 用户列表（废弃）
	UserList []*UserIdAndUserName `json:"UserList,omitnil" name:"UserList"`

	// 用户列表（新）
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil" name:"UserInfoList"`
}

func (r *CreateUserRoleProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRoleProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RoleIdList")
	delete(f, "UserList")
	delete(f, "UserInfoList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRoleProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRoleProjectResponseParams struct {
	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataId `json:"Data,omitnil" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateUserRoleProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserRoleProjectResponseParams `json:"Response"`
}

func (r *CreateUserRoleProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRoleProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRoleRequestParams struct {
	// 角色ID列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil" name:"RoleIdList"`

	// 用户列表（废弃）
	UserList []*UserIdAndUserName `json:"UserList,omitnil" name:"UserList"`

	// 用户列表（新）
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil" name:"UserInfoList"`
}

type CreateUserRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色ID列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil" name:"RoleIdList"`

	// 用户列表（废弃）
	UserList []*UserIdAndUserName `json:"UserList,omitnil" name:"UserList"`

	// 用户列表（新）
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil" name:"UserInfoList"`
}

func (r *CreateUserRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleIdList")
	delete(f, "UserList")
	delete(f, "UserInfoList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRoleResponseParams struct {
	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataId `json:"Data,omitnil" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateUserRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserRoleResponseParams `json:"Response"`
}

func (r *CreateUserRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Data struct {
	// 项目Id
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

type DataId struct {
	// 数据id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

type DatasourceInfo struct {
	// 数据库ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 域类型，1、腾讯云，2、本地
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 数据库驱动
	DbType *string `json:"DbType,omitnil" name:"DbType"`

	// HOST
	DbHost *string `json:"DbHost,omitnil" name:"DbHost"`

	// 端口
	DbPort *uint64 `json:"DbPort,omitnil" name:"DbPort"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil" name:"DbUser"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil" name:"Charset"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 修改时间
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil" name:"CreatedUser"`

	// catalog值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 连接类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectType *string `json:"ConnectType,omitnil" name:"ConnectType"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 数据源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 数据源状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 来源平台
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourcePlat *string `json:"SourcePlat,omitnil" name:"SourcePlat"`

	// 扩展参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddInfo *string `json:"AddInfo,omitnil" name:"AddInfo"`

	// 项目名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 引擎类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 数据源负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manager *string `json:"Manager,omitnil" name:"Manager"`

	// 特定操作人员白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorWhitelist *string `json:"OperatorWhitelist,omitnil" name:"OperatorWhitelist"`

	// 数据源的vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 数据源的uniqVpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 数据源的地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 操作属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	StateAction *BaseStateAction `json:"StateAction,omitnil" name:"StateAction"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil" name:"UpdatedUser"`

	// 权限列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermissionList []*PermissionGroup `json:"PermissionList,omitnil" name:"PermissionList"`

	// 权限值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthList []*string `json:"AuthList,omitnil" name:"AuthList"`

	// 第三方数据源标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataOrigin *string `json:"DataOrigin,omitnil" name:"DataOrigin"`

	// 第三方项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil" name:"DataOriginProjectId"`

	// 第三方数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil" name:"DataOriginDatasourceId"`
}

type DatasourceInfoData struct {
	// 数据源详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DatasourceInfo `json:"List,omitnil" name:"List"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *int64 `json:"TotalPages,omitnil" name:"TotalPages"`
}

// Predefined struct for user
type DeleteDatasourceRequestParams struct {
	// 数据源id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 项目id
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type DeleteDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// 数据源id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 项目id
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *DeleteDatasourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatasourceResponseParams struct {
	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 扩展
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 信息
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDatasourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDatasourceResponseParams `json:"Response"`
}

func (r *DeleteDatasourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// 项目ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 随机数
	Seed *string `json:"Seed,omitnil" name:"Seed"`

	// 默认看板
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil" name:"DefaultPanelType"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 随机数
	Seed *string `json:"Seed,omitnil" name:"Seed"`

	// 默认看板
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil" name:"DefaultPanelType"`
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
	delete(f, "Id")
	delete(f, "Seed")
	delete(f, "DefaultPanelType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// ”“
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type DeleteUserRoleProjectRequestParams struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type DeleteUserRoleProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

func (r *DeleteUserRoleProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRoleProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRoleProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRoleProjectResponseParams struct {
	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUserRoleProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserRoleProjectResponseParams `json:"Response"`
}

func (r *DeleteUserRoleProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRoleProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRoleRequestParams struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type DeleteUserRoleRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

func (r *DeleteUserRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRoleResponseParams struct {
	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUserRoleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserRoleResponseParams `json:"Response"`
}

func (r *DeleteUserRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceListRequestParams struct {
	// 无
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 返回所有页面，默认false
	AllPage *bool `json:"AllPage,omitnil" name:"AllPage"`

	// 数据库名称检索
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 无
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 无
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 搜索关键词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 过滤无权限列表的参数（0 全量，1 使用权限，2 编辑权限）
	PermissionType *int64 `json:"PermissionType,omitnil" name:"PermissionType"`
}

type DescribeDatasourceListRequest struct {
	*tchttp.BaseRequest
	
	// 无
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 返回所有页面，默认false
	AllPage *bool `json:"AllPage,omitnil" name:"AllPage"`

	// 数据库名称检索
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 无
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 无
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 搜索关键词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 过滤无权限列表的参数（0 全量，1 使用权限，2 编辑权限）
	PermissionType *int64 `json:"PermissionType,omitnil" name:"PermissionType"`
}

func (r *DescribeDatasourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AllPage")
	delete(f, "DbName")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "Keyword")
	delete(f, "PermissionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasourceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceListResponseParams struct {
	// 列表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DatasourceInfoData `json:"Data,omitnil" name:"Data"`

	// 信息
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 信息
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDatasourceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasourceListResponseParams `json:"Response"`
}

func (r *DescribeDatasourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectInfoRequestParams struct {
	// 项目Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 默认看板
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil" name:"DefaultPanelType"`
}

type DescribeProjectInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 默认看板
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil" name:"DefaultPanelType"`
}

func (r *DescribeProjectInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "DefaultPanelType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectInfoResponseParams struct {
	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 项目详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *Project `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProjectInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectInfoResponseParams `json:"Response"`
}

func (r *DescribeProjectInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectListRequestParams struct {
	// 页容，初版默认20，将来可能根据屏幕宽度动态变化
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页标
	PageNo *uint64 `json:"PageNo,omitnil" name:"PageNo"`

	// 检索模糊字段
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 是否全部展示，如果是ture，则忽略分页
	AllPage *bool `json:"AllPage,omitnil" name:"AllPage"`

	// 角色信息
	ModuleCollection *string `json:"ModuleCollection,omitnil" name:"ModuleCollection"`
}

type DescribeProjectListRequest struct {
	*tchttp.BaseRequest
	
	// 页容，初版默认20，将来可能根据屏幕宽度动态变化
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页标
	PageNo *uint64 `json:"PageNo,omitnil" name:"PageNo"`

	// 检索模糊字段
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 是否全部展示，如果是ture，则忽略分页
	AllPage *bool `json:"AllPage,omitnil" name:"AllPage"`

	// 角色信息
	ModuleCollection *string `json:"ModuleCollection,omitnil" name:"ModuleCollection"`
}

func (r *DescribeProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Keyword")
	delete(f, "AllPage")
	delete(f, "ModuleCollection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectListResponseParams struct {
	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 接口信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ProjectListData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectListResponseParams `json:"Response"`
}

func (r *DescribeProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserProjectListRequestParams struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 无
	AllPage *bool `json:"AllPage,omitnil" name:"AllPage"`

	// 无
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 无
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeUserProjectListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 无
	AllPage *bool `json:"AllPage,omitnil" name:"AllPage"`

	// 无
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 无
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeUserProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AllPage")
	delete(f, "PageNo")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserProjectListResponseParams struct {
	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CorpUserListData `json:"Data,omitnil" name:"Data"`

	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserProjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserProjectListResponseParams `json:"Response"`
}

func (r *DescribeUserProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleListRequestParams struct {
	// 页码
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 全部页码
	AllPage *bool `json:"AllPage,omitnil" name:"AllPage"`

	// 0 企业用户 1 访客 不填表示所有用户
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 模糊搜索的关键字
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

type DescribeUserRoleListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 全部页码
	AllPage *bool `json:"AllPage,omitnil" name:"AllPage"`

	// 0 企业用户 1 访客 不填表示所有用户
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 模糊搜索的关键字
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *DescribeUserRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRoleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "AllPage")
	delete(f, "UserType")
	delete(f, "Keyword")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRoleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleListResponseParams struct {
	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UserRoleListData `json:"Data,omitnil" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserRoleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserRoleListResponseParams `json:"Response"`
}

func (r *DescribeUserRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRoleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleProjectListRequestParams struct {
	// 页码
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type DescribeUserRoleProjectListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *DescribeUserRoleProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRoleProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRoleProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleProjectListResponseParams struct {
	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UserRoleListData `json:"Data,omitnil" name:"Data"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserRoleProjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserRoleProjectListResponseParams `json:"Response"`
}

func (r *DescribeUserRoleProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRoleProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmbedTokenInfo struct {
	// 信息标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 令牌
	// 注意：此字段可能返回 null，表示取不到有效值。
	BIToken *string `json:"BIToken,omitnil" name:"BIToken"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil" name:"CreatedUser"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil" name:"UpdatedUser"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// 页面Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageId *string `json:"PageId,omitnil" name:"PageId"`

	// 备用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 嵌出类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope *string `json:"Scope,omitnil" name:"Scope"`

	// 过期时间，分钟为单位，最大240
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 使用者企业Id(仅用于多用户)
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserCorpId *string `json:"UserCorpId,omitnil" name:"UserCorpId"`

	// 使用者Id(仅用于多用户)
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type IdDTO struct {
	// 请求id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// key
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 事务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranId *string `json:"TranId,omitnil" name:"TranId"`

	// 事务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranStatus *int64 `json:"TranStatus,omitnil" name:"TranStatus"`
}

// Predefined struct for user
type ModifyDatasourceCloudRequestParams struct {
	// 后端提供字典：域类型，1、腾讯云，2、本地
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 驱动
	DbType *string `json:"DbType,omitnil" name:"DbType"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil" name:"Charset"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil" name:"DbUser"`

	// 密码
	DbPwd *string `json:"DbPwd,omitnil" name:"DbPwd"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 住键
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 公有云内网ip
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 公有云内网端口
	Vport *string `json:"Vport,omitnil" name:"Vport"`

	// vpc标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 统一vpc标识
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 区域标识（gz,bj)
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 扩展参数
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 数据源产品名
	ProdDbName *string `json:"ProdDbName,omitnil" name:"ProdDbName"`

	// 第三方数据源标识
	DataOrigin *string `json:"DataOrigin,omitnil" name:"DataOrigin"`

	// 第三方项目id
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil" name:"DataOriginProjectId"`

	// 第三方数据源id
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil" name:"DataOriginDatasourceId"`
}

type ModifyDatasourceCloudRequest struct {
	*tchttp.BaseRequest
	
	// 后端提供字典：域类型，1、腾讯云，2、本地
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 驱动
	DbType *string `json:"DbType,omitnil" name:"DbType"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil" name:"Charset"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil" name:"DbUser"`

	// 密码
	DbPwd *string `json:"DbPwd,omitnil" name:"DbPwd"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 住键
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 公有云内网ip
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 公有云内网端口
	Vport *string `json:"Vport,omitnil" name:"Vport"`

	// vpc标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 统一vpc标识
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 区域标识（gz,bj)
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 扩展参数
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 数据源产品名
	ProdDbName *string `json:"ProdDbName,omitnil" name:"ProdDbName"`

	// 第三方数据源标识
	DataOrigin *string `json:"DataOrigin,omitnil" name:"DataOrigin"`

	// 第三方项目id
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil" name:"DataOriginProjectId"`

	// 第三方数据源id
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil" name:"DataOriginDatasourceId"`
}

func (r *ModifyDatasourceCloudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatasourceCloudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "DbType")
	delete(f, "Charset")
	delete(f, "DbUser")
	delete(f, "DbPwd")
	delete(f, "DbName")
	delete(f, "SourceName")
	delete(f, "ProjectId")
	delete(f, "Id")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "VpcId")
	delete(f, "UniqVpcId")
	delete(f, "RegionId")
	delete(f, "ExtraParam")
	delete(f, "InstanceId")
	delete(f, "ProdDbName")
	delete(f, "DataOrigin")
	delete(f, "DataOriginProjectId")
	delete(f, "DataOriginDatasourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatasourceCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatasourceCloudResponseParams struct {
	// 成功无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDatasourceCloudResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatasourceCloudResponseParams `json:"Response"`
}

func (r *ModifyDatasourceCloudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatasourceCloudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatasourceRequestParams struct {
	// HOST
	DbHost *string `json:"DbHost,omitnil" name:"DbHost"`

	// 端口
	DbPort *uint64 `json:"DbPort,omitnil" name:"DbPort"`

	// 后端提供字典：域类型，1、腾讯云，2、本地
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 驱动
	DbType *string `json:"DbType,omitnil" name:"DbType"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil" name:"Charset"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil" name:"DbUser"`

	// 密码
	DbPwd *string `json:"DbPwd,omitnil" name:"DbPwd"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 数据源id
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// catalog值
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 第三方数据源标识
	DataOrigin *string `json:"DataOrigin,omitnil" name:"DataOrigin"`

	// 第三方项目id
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil" name:"DataOriginProjectId"`

	// 第三方数据源id
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil" name:"DataOriginDatasourceId"`

	// 扩展参数
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 腾讯云私有网络统一标识
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 私有网络ip
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 私有网络端口
	Vport *string `json:"Vport,omitnil" name:"Vport"`

	// 腾讯云私有网络标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type ModifyDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// HOST
	DbHost *string `json:"DbHost,omitnil" name:"DbHost"`

	// 端口
	DbPort *uint64 `json:"DbPort,omitnil" name:"DbPort"`

	// 后端提供字典：域类型，1、腾讯云，2、本地
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 驱动
	DbType *string `json:"DbType,omitnil" name:"DbType"`

	// 数据库编码
	Charset *string `json:"Charset,omitnil" name:"Charset"`

	// 用户名
	DbUser *string `json:"DbUser,omitnil" name:"DbUser"`

	// 密码
	DbPwd *string `json:"DbPwd,omitnil" name:"DbPwd"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库别名
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 数据源id
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// catalog值
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 第三方数据源标识
	DataOrigin *string `json:"DataOrigin,omitnil" name:"DataOrigin"`

	// 第三方项目id
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil" name:"DataOriginProjectId"`

	// 第三方数据源id
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil" name:"DataOriginDatasourceId"`

	// 扩展参数
	ExtraParam *string `json:"ExtraParam,omitnil" name:"ExtraParam"`

	// 腾讯云私有网络统一标识
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// 私有网络ip
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 私有网络端口
	Vport *string `json:"Vport,omitnil" name:"Vport"`

	// 腾讯云私有网络标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

func (r *ModifyDatasourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatasourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbHost")
	delete(f, "DbPort")
	delete(f, "ServiceType")
	delete(f, "DbType")
	delete(f, "Charset")
	delete(f, "DbUser")
	delete(f, "DbPwd")
	delete(f, "DbName")
	delete(f, "SourceName")
	delete(f, "Id")
	delete(f, "ProjectId")
	delete(f, "Catalog")
	delete(f, "DataOrigin")
	delete(f, "DataOriginProjectId")
	delete(f, "DataOriginDatasourceId")
	delete(f, "ExtraParam")
	delete(f, "UniqVpcId")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatasourceResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDatasourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatasourceResponseParams `json:"Response"`
}

func (r *ModifyDatasourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatasourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// 项目Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 名字
	Name *string `json:"Name,omitnil" name:"Name"`

	// 颜色值
	ColorCode *string `json:"ColorCode,omitnil" name:"ColorCode"`

	// 图标
	Logo *string `json:"Logo,omitnil" name:"Logo"`

	// 备注
	Mark *string `json:"Mark,omitnil" name:"Mark"`

	// 可申请
	IsApply *bool `json:"IsApply,omitnil" name:"IsApply"`

	// 种子
	Seed *string `json:"Seed,omitnil" name:"Seed"`

	// 默认看板
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil" name:"DefaultPanelType"`

	// 2
	PanelScope *string `json:"PanelScope,omitnil" name:"PanelScope"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 名字
	Name *string `json:"Name,omitnil" name:"Name"`

	// 颜色值
	ColorCode *string `json:"ColorCode,omitnil" name:"ColorCode"`

	// 图标
	Logo *string `json:"Logo,omitnil" name:"Logo"`

	// 备注
	Mark *string `json:"Mark,omitnil" name:"Mark"`

	// 可申请
	IsApply *bool `json:"IsApply,omitnil" name:"IsApply"`

	// 种子
	Seed *string `json:"Seed,omitnil" name:"Seed"`

	// 默认看板
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil" name:"DefaultPanelType"`

	// 2
	PanelScope *string `json:"PanelScope,omitnil" name:"PanelScope"`
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
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "ColorCode")
	delete(f, "Logo")
	delete(f, "Mark")
	delete(f, "IsApply")
	delete(f, "Seed")
	delete(f, "DefaultPanelType")
	delete(f, "PanelScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 返回数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ModifyUserRoleProjectRequestParams struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`
}

type ModifyUserRoleProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`
}

func (r *ModifyUserRoleProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRoleProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserId")
	delete(f, "RoleIdList")
	delete(f, "Email")
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRoleProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleProjectResponseParams struct {
	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyUserRoleProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserRoleProjectResponseParams `json:"Response"`
}

func (r *ModifyUserRoleProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRoleProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleRequestParams struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 手机号
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 手机区号
	AreaCode *string `json:"AreaCode,omitnil" name:"AreaCode"`
}

type ModifyUserRoleRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 角色ID 列表
	RoleIdList []*int64 `json:"RoleIdList,omitnil" name:"RoleIdList"`

	// 邮箱
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 手机号
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 手机区号
	AreaCode *string `json:"AreaCode,omitnil" name:"AreaCode"`
}

func (r *ModifyUserRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "RoleIdList")
	delete(f, "Email")
	delete(f, "UserName")
	delete(f, "PhoneNumber")
	delete(f, "AreaCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleResponseParams struct {
	// 扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil" name:"Extra"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyUserRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserRoleResponseParams `json:"Response"`
}

func (r *ModifyUserRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PermissionComponent struct {
	// 权限值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleId *string `json:"ModuleId,omitnil" name:"ModuleId"`

	// 可见/可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeType *string `json:"IncludeType,omitnil" name:"IncludeType"`

	// 目标升级版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeVersionType *string `json:"UpgradeVersionType,omitnil" name:"UpgradeVersionType"`

	// 补充信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tips *string `json:"Tips,omitnil" name:"Tips"`

	// 补充信息的key值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TipsKey *string `json:"TipsKey,omitnil" name:"TipsKey"`
}

type PermissionGroup struct {
	// 分组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleGroup *string `json:"ModuleGroup,omitnil" name:"ModuleGroup"`

	// 权限列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*PermissionComponent `json:"Components,omitnil" name:"Components"`
}

type Project struct {
	// 项目ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 项目Logo
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitnil" name:"Logo"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// logo底色
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColorCode *string `json:"ColorCode,omitnil" name:"ColorCode"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil" name:"CreatedUser"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 成员个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberCount *int64 `json:"MemberCount,omitnil" name:"MemberCount"`

	// 页面个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *int64 `json:"PageCount,omitnil" name:"PageCount"`

	// 最后修改报表、简报名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyName *string `json:"LastModifyName,omitnil" name:"LastModifyName"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil" name:"Source"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Apply *bool `json:"Apply,omitnil" name:"Apply"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil" name:"UpdatedUser"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *string `json:"CorpId,omitnil" name:"CorpId"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mark *string `json:"Mark,omitnil" name:"Mark"`

	// ""
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seed *string `json:"Seed,omitnil" name:"Seed"`

	// 项目内权限列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthList []*string `json:"AuthList,omitnil" name:"AuthList"`

	// 默认看板
	// 注意：此字段可能返回 null，表示取不到有效值。
	PanelScope *string `json:"PanelScope,omitnil" name:"PanelScope"`

	// 是否被托管
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsExternalManage *bool `json:"IsExternalManage,omitnil" name:"IsExternalManage"`

	// 托管平台名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagePlatform *string `json:"ManagePlatform,omitnil" name:"ManagePlatform"`

	// 定制化参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigList []*ProjectConfigList `json:"ConfigList,omitnil" name:"ConfigList"`
}

type ProjectConfigList struct {
	// 模块组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleGroup *string `json:"ModuleGroup,omitnil" name:"ModuleGroup"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*ProjectConfigResult `json:"Components,omitnil" name:"Components"`
}

type ProjectConfigResult struct {
	// 配置名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleId *string `json:"ModuleId,omitnil" name:"ModuleId"`

	// 配置方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeType *string `json:"IncludeType,omitnil" name:"IncludeType"`

	// 额外参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil" name:"Params"`
}

type ProjectListData struct {
	// 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Project `json:"List,omitnil" name:"List"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *uint64 `json:"TotalPages,omitnil" name:"TotalPages"`
}

type UserIdAndUserName struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpId *string `json:"CorpId,omitnil" name:"CorpId"`

	// 电子邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil" name:"Email"`

	// 最后一次登录时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLogin *string `json:"LastLogin,omitnil" name:"LastLogin"`

	// 停启用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 首次登陆是否修改密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstModify *int64 `json:"FirstModify,omitnil" name:"FirstModify"`

	// 手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 手机区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCode *string `json:"AreaCode,omitnil" name:"AreaCode"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil" name:"CreatedUser"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 更改人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil" name:"UpdatedUser"`

	// 更改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// 全局角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	GlobalUserName *string `json:"GlobalUserName,omitnil" name:"GlobalUserName"`

	// 手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`
}

type UserInfo struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil" name:"Email"`

	// 手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 手机号区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCode *string `json:"AreaCode,omitnil" name:"AreaCode"`
}

type UserRoleListData struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *int64 `json:"TotalPages,omitnil" name:"TotalPages"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*UserRoleListDataUserRoleInfo `json:"List,omitnil" name:"List"`
}

type UserRoleListDataRoleInfo struct {
	// 角色名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// 角色ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *int64 `json:"RoleId,omitnil" name:"RoleId"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 项目名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 是否为全局角色（0 不是 1 是）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScopeType *int64 `json:"ScopeType,omitnil" name:"ScopeType"`

	// 角色key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleCollection *string `json:"ModuleCollection,omitnil" name:"ModuleCollection"`
}

type UserRoleListDataUserRoleInfo struct {
	// 业务ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 角色列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleList []*UserRoleListDataRoleInfo `json:"RoleList,omitnil" name:"RoleList"`

	// 角色ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleIdList []*uint64 `json:"RoleIdList,omitnil" name:"RoleIdList"`

	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 企业ID
	CorpId *string `json:"CorpId,omitnil" name:"CorpId"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil" name:"Email"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedUser *string `json:"CreatedUser,omitnil" name:"CreatedUser"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedUser *string `json:"UpdatedUser,omitnil" name:"UpdatedUser"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// 最后一次登录时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLogin *string `json:"LastLogin,omitnil" name:"LastLogin"`

	// 账号状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// 手机号区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCode *string `json:"AreaCode,omitnil" name:"AreaCode"`

	// 是否为主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootAccount *bool `json:"RootAccount,omitnil" name:"RootAccount"`

	// 是否为企业管理员
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpAdmin *bool `json:"CorpAdmin,omitnil" name:"CorpAdmin"`
}